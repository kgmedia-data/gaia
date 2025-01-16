package log

import (
	"context"
	"fmt"
	"reflect"
	"slices"

	"cloud.google.com/go/logging"
	"github.com/kgmedia-data/gaia/pkg/msg"
	"github.com/kgmedia-data/gaia/pkg/pub"
	"github.com/sirupsen/logrus"
)

type GCPProcessor struct {
	logName   string
	projectId string
	labels    map[string]string
	logger    *logging.Logger
}

func NewGCPProcessor(logName string, projectId string, labels map[string]string) *GCPProcessor {
	// eg. GCP_LOG_NAME is "gia"
	ctx := context.Background()
	// eg. Project ID is "kgdata-aiml"
	client, err := logging.NewClient(ctx, projectId)
	if err != nil {
		logrus.Fatalf("Failed to create client: %v", err)
	}

	logger := client.Logger(logName)

	return &GCPProcessor{
		logName:   logName,
		projectId: projectId,
		logger:    logger,
		labels:    labels,
	}
}

func (p *GCPProcessor) error(err error, method string, params ...interface{}) error {
	message := fmt.Errorf("LoggingProcessor.(%v)(%v) %w", method, params, err)
	logrus.Error(message)
	return message
}

func (p *GCPProcessor) Execute(message msg.Message[string]) error {
	logrus.Infoln("LoggerProcessor.Execute", message.Data)
	// send to GCP

	var logMapping = map[string]logging.Severity{
		"info":  logging.Info,
		"error": logging.Error,
		"fatal": logging.Critical,
		"panic": logging.Critical,
	}

	payload := map[string]any{
		"data":   message.Data,
		"labels": p.labels,
		"fields": message.Attribute,
	}

	p.logger.Log(logging.Entry{Payload: payload, Severity: logMapping[message.Attribute["level"]]})

	return nil
}

var logLevelMappings = map[logrus.Level]logging.Severity{
	logrus.TraceLevel: logging.Default,
	logrus.DebugLevel: logging.Debug,
	logrus.InfoLevel:  logging.Info,
	logrus.WarnLevel:  logging.Warning,
	logrus.ErrorLevel: logging.Error,
	logrus.FatalLevel: logging.Critical,
	logrus.PanicLevel: logging.Critical,
}

type GcpLogHook struct {
	pub pub.IPublisher[string]
}

func NewExtraFieldHook(pub pub.IPublisher[string]) *GcpLogHook {

	return &GcpLogHook{
		pub: pub,
	}
}

func (h *GcpLogHook) error(err error, method string, params ...interface{}) error {
	message := fmt.Errorf("NewExtraFieldHook.(%v)(%v) %w", method, params, err)
	logrus.Error(message)
	return message
}

func (h *GcpLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *GcpLogHook) Fire(entry *logrus.Entry) error {

	level := logLevelMappings[entry.Level]
	things := []logging.Severity{logging.Critical}

	attr := make(map[string]string)
	attr["level"] = entry.Level.String()

	// additional data
	for key, value := range entry.Data {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		attr[strKey] = strValue
	}

	// validate key GCP is valid
	insertIntoGCP := false
	if entry.Data["gcp"] != nil {
		insertIntoGCP = reflect.ValueOf(entry.Data["gcp"]).Bool()
	}

	if slices.Contains(things, level) || insertIntoGCP {
		err := h.pub.Publish(msg.Message[string]{Data: entry.Message, Attribute: attr})
		if err != nil {
			return h.error(err, "Hook Fire")
		}
	}

	return nil
}

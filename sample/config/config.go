package config

import (
	"reflect"
	"strings"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Rest   RestConfig   `mapstructure:"rest"`
	Metric MetricConfig `mapstructure:"metric"`
	GcpLog GcpLog       `mapstructure:"gcpLog"`
}

type RestConfig struct {
	Server ServerConfig   `mapstructure:"server"`
	Db     DatabaseConfig `mapstructure:"db"`
	Secret string         `mapstructure:"secret"`
}

type GcpLog struct {
	ProjectId string            `mapstructure:"projectId"`
	LogName   string            `mapstructure:"logName"`
	Labels    map[string]string `mapstructure:"labels"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Key  string `mapstructure:"key"`
}

type MetricConfig struct {
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	DataStore  string `mapstructure:"datastore"`
	NumberConn int    `mapstructure:"nConn"`
}

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	BindEnvs(Config{})
}

func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func LoadConfig[T any](path string) (T, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	var config T

	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn(err, "Failed to read config file, using values from environment variable")
		// rest.Server.Host -> REST_SERVER_HOST
		// rest.Server.Key -> REST_SERVER_KEY
		BindEnvs(config)
	}

	var conf T
	if err := viper.Unmarshal(&conf); err != nil {
		return conf, errors.Trace(err)
	}

	return conf, nil
}

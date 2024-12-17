package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/google"
)

type VertexRest struct {
	request         *resty.Request
	config          VertexAIConfig
	endpoint        string
	tokenExpiration time.Time
}

func NewVertexRest() (*VertexRest, error) {

	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	client := resty.New()
	request := client.NewRequest()
	request.SetHeader("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return &VertexRest{
		request:         request,
		config:          *GenerateVertexDefaultConfig(),
		tokenExpiration: time.Now().Add(1 * time.Hour),
	}, nil
}

func GenerateVertexDefaultConfig() *VertexAIConfig {
	config := VertexAIConfig{
		Model: "gemini-1.5-flash-002",
		GenerationConfig: GenerationConfig{
			Temperature:      1,
			TopP:             1,
			CandidateCount:   1,
			MaxOutputTokens:  1024,
			ResponseMimeType: "application/json",
		},
		SystemInstruction: SystemInstruction{},
		Contents:          Contents{},
		Labels:            map[string]string{},
		SafetySettings: []SafetySettings{
			{
				Category:  "HARM_CATEGORY_HATE_SPEECH",
				Threshold: "OFF",
			},
			{
				Category:  "HARM_CATEGORY_DANGEROUS_CONTENT",
				Threshold: "OFF",
			},
			{
				Category:  "HARM_CATEGORY_SEXUALLY_EXPLICIT",
				Threshold: "OFF",
			},
			{
				Category:  "HARM_CATEGORY_HARASSMENT",
				Threshold: "OFF",
			},
		},
	}

	return &config
}

func getAccessToken() (string, error) {
	ctx := context.Background()
	// read file into array of bytes
	filename := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	creds, err := google.CredentialsFromJSON(ctx, []byte(data),
		"https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return "", err
	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil

}

func (s VertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("VertexRestModel.(%v)(%v) %w", method, params, err)
}

func (s *VertexRest) RenewToken() error {
	newToken, err := getAccessToken()
	if err != nil {
		return s.error(err, "RenewToken")
	}

	s.request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", newToken))
	s.tokenExpiration = time.Now().Add(1 * time.Hour)

	return nil
}

func (s *VertexRest) GetResponse() (*resty.Response, error) {
	if time.Now().After(s.tokenExpiration) {
		if err := s.RenewToken(); err != nil {
			return nil, s.error(err, "GetResponse")
		}
	}

	resp, err := s.request.
		SetBody(s.config).
		Post(s.endpoint)
	if err != nil {
		return nil, s.error(err, "GetResponse", s.config)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get summary: status %s, response: %s", resp.Status(), resp.Body())
	}
	return resp, nil
}

func (s *VertexRest) ParseResponse(resp *resty.Response) (OutputVertex, error) {
	result := OutputVertex{}
	err := json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return OutputVertex{}, s.error(err, "ParseResponse", resp)
	}

	return result, nil
}

func ParseSingleResponseVertex[T any](resp *resty.Response, result *T) error {
	outputVertex := OutputVertex{}
	if err := json.Unmarshal(resp.Body(), &outputVertex); err != nil {
		return fmt.Errorf("ParseSingleResponse - resp.Body(): %w (%v)", err, resp.Body())
	}

	data := outputVertex.Candidates[0].Content.Parts[0].Text
	if err := json.Unmarshal([]byte(data), result); err != nil {
		return fmt.Errorf("ParseSingleResponse - data: %w (%v)", err, data)
	}

	return nil
}

func (s *VertexRest) SetEndpoint(endpoint string) *VertexRest {
	s.endpoint = endpoint
	return s
}
func (s *VertexRest) SetModel(model string) *VertexRest {
	s.config.Model = model
	return s
}

func (s *VertexRest) SetTemperature(temp float64) *VertexRest {
	s.config.GenerationConfig.Temperature = temp
	return s
}

func (s *VertexRest) SetTopP(topP int) *VertexRest {
	s.config.GenerationConfig.TopP = topP
	return s
}

func (s *VertexRest) SetMaxOutputTokens(maxOutputTokens int) *VertexRest {
	s.config.GenerationConfig.MaxOutputTokens = maxOutputTokens
	return s
}

func (s *VertexRest) SetResponseSchema(schema map[string]interface{}) *VertexRest {
	s.config.GenerationConfig.ResponseSchema = schema
	return s
}

func (s *VertexRest) AddSystemInstruction(instruction string) *VertexRest {
	s.config.SystemInstruction.Parts = append(s.config.SystemInstruction.Parts, InstructionPart{Text: instruction})
	return s
}

func (s *VertexRest) AddContent(prompt string, role string) *VertexRest {
	role = strings.ToUpper(role)
	if role != "USER" && role != "MODEL" {
		return s
	}
	s.config.Contents.Role = role
	s.config.Contents.Parts = append(s.config.Contents.Parts, InstructionPart{Text: prompt})
	return s
}

func (s *VertexRest) SetContent(prompt, role string) *VertexRest {
	role = strings.ToUpper(role)
	if role != "USER" && role != "MODEL" {
		return s
	}
	s.config.Contents.Role = role
	s.config.Contents.Parts = []InstructionPart{{Text: prompt}}
	return s
}

func (s *VertexRest) AddLabel(key, value string) *VertexRest {
	if key == "" {
		return s
	}
	if value != "" {
		s.config.Labels[key] = value
	} else {
		delete(s.config.Labels, key)
	}
	return s
}

func (s *VertexRest) SetSafetySettings(safetySettings map[string]string) *VertexRest {

	for k, v := range safetySettings {
		for idx, setting := range s.config.SafetySettings {
			if setting.Category == k {
				s.config.SafetySettings[idx].Threshold = v
				break
			}
		}
	}
	return s
}

func (s *VertexRest) AddTools(tools []Tools) *VertexRest {
	s.config.Tools = append(s.config.Tools, tools...)
	return s
}

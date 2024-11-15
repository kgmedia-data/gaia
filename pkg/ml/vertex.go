package ml

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/google"
)

type VertexRestModel struct {
	request *resty.Request
	config  VertexAIConfig
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

func NewVertexRestModel() (*VertexRestModel, error) {

	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	client := resty.New()
	request := client.NewRequest()
	request.SetHeader("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return &VertexRestModel{
		request: request,
		config:  *GenerateVertexDefaultConfig(),
	}, nil
}

func GenerateVertexDefaultConfig() *VertexAIConfig {
	config := VertexAIConfig{
		Model: "gemini-1.5-flash-002",
		GenerationConfig: GenerationConfig{
			Temperature:      1,
			TopP:             1,
			TopK:             1,
			CandidateCount:   1,
			MaxOutputTokens:  1024,
			ResponseMimeType: "application/json",
		},
		SystemInstruction: SystemInstruction{},
		Contents: Contents{
			Role: "MODEL",
		},
		Labels: map[string]string{},
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

func (s *VertexRestModel) SetModel(model string) *VertexRestModel {
	s.config.Model = model
	return s
}

func (s *VertexRestModel) SetTemperature(temp float64) *VertexRestModel {
	s.config.GenerationConfig.Temperature = temp
	return s
}

func (s *VertexRestModel) SetTopP(topP int) *VertexRestModel {
	s.config.GenerationConfig.TopP = topP
	return s
}

func (s *VertexRestModel) SetMaxOutputTokens(maxOutputTokens int) *VertexRestModel {
	s.config.GenerationConfig.MaxOutputTokens = maxOutputTokens
	return s
}

func (s *VertexRestModel) SetResponseSchema(schema map[string]interface{}) *VertexRestModel {
	s.config.GenerationConfig.ResponseSchema = schema
	return s
}

func (s *VertexRestModel) AddSystemInstruction(instruction string) *VertexRestModel {
	s.config.SystemInstruction.Parts = append(s.config.SystemInstruction.Parts, InstructionPart{Text: instruction})
	return s
}

func (s *VertexRestModel) AddContent(prompt string) *VertexRestModel {
	s.config.Contents.Parts = append(s.config.Contents.Parts, InstructionPart{Text: prompt})
	return s
}

func (s *VertexRestModel) SetContent(prompt string) *VertexRestModel {
	s.config.Contents.Parts = []InstructionPart{{Text: prompt}}
	return s
}

func (s *VertexRestModel) AddLabels(labels map[string]string) *VertexRestModel {
	for k, v := range labels {
		s.config.Labels[k] = v
	}
	return s
}

func (s *VertexRestModel) SetSafetySettings(safetySettings map[string]string) *VertexRestModel {
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

func (s *VertexRestModel) AddTools(tools []Tools) *VertexRestModel {
	s.config.Tools = append(s.config.Tools, tools...)
	return s
}

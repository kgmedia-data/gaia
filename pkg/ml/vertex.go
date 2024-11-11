package ml

import (
	"context"
	"encoding/json"
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

func loadConfigVertex(filename string) (*VertexAIConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config VertexAIConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func NewVertexRestModel(configDir string) (*VertexRestModel, error) {
	client := resty.New()
	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}
	request := client.NewRequest()
	request.SetHeader("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	config, err := loadConfigVertex(configDir)
	if err != nil {
		return &VertexRestModel{}, err
	}

	return &VertexRestModel{
		request: request,
		config:  *config,
	}, nil
}

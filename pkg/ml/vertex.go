package ml

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/go-resty/resty/v2"
)

type VertexRestModel struct {
	request *resty.Request
	config  VertexAIConfig
}

func getAccessToken() (string, error) {
	cmd := exec.Command("gcloud", "auth", "print-access-token", "--quiet")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	token := strings.TrimSpace(string(output))
	return token, nil
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

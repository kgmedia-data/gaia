package ml

import (
	"fmt"
)

type SpokespersonVertexRest struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewSpokespersonVertexRest(projectID, location string, projectLabel ProjectLabel) (*SpokespersonVertexRest, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("You are a tools for extracting spokesperson, it's quotes, and sentiment from news articles.").
		SetResponseSchema(map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "string",
					},
					"name": map[string]interface{}{
						"type": "string",
					},
					"job_title": map[string]interface{}{
						"type": "string",
					},
					"sentiment": map[string]interface{}{
						"type": "string",
						"enum": []string{"positive", "negative", "neutral"},
					},
					"statements": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "string",
						},
					},
				},
				"required": []string{"id", "name", "statements", "job_title", "sentiment"},
			},
		}).
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		SetEndpoint(fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, vertex.config.Model))

	return &SpokespersonVertexRest{vertex: vertex}, nil
}

func (s SpokespersonVertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("SpokespersonVertexRest.(%v)(%v) %w", method, params, err)
}

func (s *SpokespersonVertexRest) Infer(content string) ([]Spokesperson, error) {

	s.vertex.SetContent(content, "USER")

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}
	result := []Spokesperson{}
	err = ParseSingleResponseVertex(resp, &result)
	if err != nil {
		return nil, s.error(err, "Infer")
	}
	return result, nil
}

func (s *SpokespersonVertexRest) InferBatch(content map[string]string) (map[string][]Spokesperson, error) {

	for i, c := range content {
		s.vertex.AddContent(fmt.Sprintf("ID %v: , Content: %v\n", i, c), "USER")
	}

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}
	var spokespersons []Spokesperson
	if err = ParseSingleResponseVertex(resp, &spokespersons); err != nil {
		return nil, s.error(err, "InferBatch - ParseSingleResponseVertex")
	}

	finalResult := make(map[string][]Spokesperson, len(content))
	for _, spokesperson := range spokespersons {
		finalResult[spokesperson.ID] = append(finalResult[spokesperson.ID], spokesperson)
	}

	return finalResult, nil
}

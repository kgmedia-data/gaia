package ml

import (
	"fmt"
	"sort"
)

type SentimentVertex struct {
	Vertex *VertexRest
}

func (Vertex *VertexRest) NewSentimentVertexRest(projectID, location string, projectLabel ProjectLabel) (*SentimentVertex, error) {

	Vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("You are a sentiment analysis model, which will be given input of Indonesian language news article.").
		SetResponseSchema(map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "string",
					},
					"sentiment": map[string]interface{}{
						"type": "string",
						"enum": []string{"positive", "negative", "neutral"},
					},
				},
				"required": []string{"id", "sentiment"},
			},
		}).
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		SetEndpoint(fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, Vertex.config.Model))

	return &SentimentVertex{Vertex: Vertex}, nil
}

func (s SentimentVertex) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("SentimentVertexRest.(%v)(%v) %w", method, params, err)
}

func (s *SentimentVertex) Infer(contents string) ([]Sentiment, error) {

	s.Vertex.SetContent(contents, "USER")

	resp, err := s.Vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}

	result := []Sentiment{}
	err = ParseSingleResponseVertex(resp, &result)
	if err != nil {
		return nil, s.error(err, "Infer - ParseSingleResponseVertex")
	}

	return result, nil
}

func (s *SentimentVertex) InferBatch(contents map[string]string) ([]Sentiment, error) {
	s.Vertex.ResetContentsParts()
	for i, c := range contents {
		s.Vertex.AddContent(fmt.Sprintf("ID %v: , Content: %v\n", i, c), "USER")
	}
	resp, err := s.Vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}

	sentiments := []Sentiment{}
	err = ParseSingleResponseVertex(resp, &sentiments)
	if err != nil {
		return nil, s.error(err, "Infer - ParseSingleResponseVertex")
	}
	sort.Slice(sentiments, func(i, j int) bool {
		return sentiments[i].ID < sentiments[j].ID
	})
	return sentiments, nil
}

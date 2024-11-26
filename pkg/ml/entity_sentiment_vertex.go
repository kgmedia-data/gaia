package ml

import (
	"fmt"
)

type EntitySentimentVertex struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewEntitySentimentVertexRest(projectID, location string, projectLabel ProjectLabel) (*EntitySentimentVertex, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("You are an entity-based sentiment analysis model, which will be given input of Indonesian language news article, and a list of entity keywords for matching the sentiment. input will be like this: text: x. entity: y,z. 'name' is entity name. only extract the entities from input. sentiment is either positive, negative, or neutral. is_mentioned is true if the name is written in the text (such as name='shopee', if text contains 'shopi' or 'shope', return true) or any keywords that is relevant is mentioned in the text (such as name='toyota', if text contains 'fortuner', return true), else false.").
		SetResponseSchema(map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"entity_id": map[string]interface{}{
						"type": "integer",
					},
					"sentiment": map[string]interface{}{
						"type": "string",
					},
					"is_mentioned": map[string]interface{}{
						"type": "boolean",
					},
				},
				"required": []string{"entity_id", "is_mentioned", "sentiment"},
			},
		}).
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		SetEndpoint(fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, vertex.config.Model))

	return &EntitySentimentVertex{vertex: vertex}, nil
}

func (s EntitySentimentVertex) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("SentimentVertexRest.(%v)(%v) %w", method, params, err)
}

func (s *EntitySentimentVertex) Infer(contents string) ([]EntitySentiment, error) {

	s.vertex.SetContent(contents, "USER")

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, s.error(err, "Infer - GetResponse")
	}

	result := []EntitySentiment{}
	err = ParseSingleResponseVertex(resp, &result)
	if err != nil {
		return nil, s.error(err, "Infer - ParseSingleResponseVertex")
	}

	return result, nil
}

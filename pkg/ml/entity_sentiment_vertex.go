package ml

import (
	"fmt"
)

type EntitySentimentVertex struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewEntitySentimentVertexRest(projectID, location string, projectLabel ProjectLabel) (*EntitySentimentVertex, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1.2).
		SetMaxOutputTokens(8192).
		AddSystemInstruction(`You are an entity-based sentiment analysis model, which will be given input of Indonesian language news articles and a list of entity name and keywords for matching the sentiment. Your task is to analyze what is the citizen sentiment towards the entity name within the text, or the things, or events that is related to it. If the entity name is doing achievement or going to the right direction, or suggestion, give positive and vice versa.

		Give weighting more into the positive sentiments. Don't classify as negative unless you're sure. Don't easily classify as neutral. If the entity is shown taking corrective actions, enforcing rules, or handling issues transparently, then sentiment should be neutral or positive, even if the article contains some negative words. 


		For example:
		If a police institution is investigating an officer, sentiment should be neutral or positive because the institution is taking action.
		If a company is addressing customer complaints or launching an improvement, sentiment should be neutral or positive.
		
		is_mentioned is true if the name is written in the text (such as name='shopee', if text contains 'shopi' or 'shope', return true) or any keywords that is relevant is mentioned in the text (such as name='toyota', if text contains 'fortuner', return true), else false.

		please include the reason you set each sentiment in column sentiment_reason
`).
		SetResponseSchema(map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"entity_id": map[string]interface{}{
						"type": "string",
					},
					"sentiment": map[string]interface{}{
						"type": "string",
						"enum": []string{"positive", "negative", "neutral"},
					},
					"is_mentioned": map[string]interface{}{
						"type": "boolean",
					},
					"sentiment_reason": map[string]interface{}{
						"type": "string",
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

func (s *EntitySentimentVertex) Infer(contents string) ([]EntitySentiment, OutputVertex, error) {

	s.vertex.SetContent(contents, "USER")

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, OutputVertex{}, s.error(err, "Infer - GetResponse")
	}

	result, outputVertex := []EntitySentiment{}, OutputVertex{}
	err = ParseSingleResponseVertex(resp, &result, &outputVertex)
	if err != nil {
		return nil, OutputVertex{}, s.error(err, "Infer - ParseSingleResponseVertex")
	}

	return result, outputVertex, nil
}

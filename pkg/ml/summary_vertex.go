package ml

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SummaryVertexRest struct {
	endpoint string
	vertex   *VertexRestModel
}

func NewSummaryVertexRest(projectID, location string, vertex *VertexRestModel) (*SummaryVertexRest, error) {

	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, vertex.config.Model)

	return &SummaryVertexRest{
		endpoint: endpoint,
		vertex:   vertex,
	}, nil
}

func (s SummaryVertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("SummarizeVertex.(%v)(%v) %w", method, params, err)
}

func (s *SummaryVertexRest) BatchSummarize(language string, minSentences, maxSentences int, input []Summary) ([]Summary, error) {
	SYSTEM_INSTRUCTION := "You are a tools for summarization of title of news articles."

	contents_text := s.generateContentsText(language, minSentences, maxSentences, input)

	requestBody, err := json.Marshal(map[string]interface{}{
		"generation_config": map[string]interface{}{
			"temperature":      s.vertex.config.Parameters.Temperature,
			"topP":             s.vertex.config.Parameters.TopP,
			"topK":             s.vertex.config.Parameters.TopK,
			"candidateCount":   s.vertex.config.Parameters.CandidateCount,
			"maxOutputTokens":  s.vertex.config.Parameters.MaxOutputTokens,
			"responseMimeType": "application/json",
			"responseSchema": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"group_id": map[string]interface{}{
							"type": "string",
						},
						"content": map[string]interface{}{
							"type": "string",
						},
					},
					"required": []string{"content", "group_id"},
				},
			},
		},
		"system_instruction": map[string]interface{}{
			"parts": []map[string]string{
				{
					"text": SYSTEM_INSTRUCTION,
				},
			},
		},
		"contents": map[string]interface{}{
			"role": "MODEL",
			"parts": map[string]string{
				"text": contents_text,
			},
		},
		"labels": map[string]string{
			"project": "medeab",
			"env":     "prod",
			"task":    "summarization",
		},
		"safety_settings": map[string]interface{}{
			"category":  "HARM_CATEGORY_HATE_SPEECH",
			"threshold": "BLOCK_ONLY_HIGH",
			"method":    "SEVERITY",
		},
	})
	if err != nil {
		return nil, s.error(err, "BatchSummarize")
	}

	resp, err := s.vertex.request.SetBody(requestBody).Post(s.endpoint)
	if err != nil {
		return nil, s.error(err, "BatchSummarize", requestBody)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get summary: status %s", resp.Status())
	}

	// Parse Response Body to JSON
	result := OutputVertex{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, s.error(err, "BatchSummarize", requestBody)
	}

	// Parse JSON to []summary
	summary := []Summary{}
	err = json.Unmarshal([]byte(result.Candidates[0].Content.Parts[0].Text), &summary)
	if err != nil {
		return nil, s.error(err, "BatchSummarize", requestBody)
	}

	return summary, nil
}

func (s *SummaryVertexRest) generateContentsText(language string, minSentences, maxSentences int, input []Summary) (contents string) {
	contents = fmt.Sprintf("Create a summary of given indonesian news article title. Ignore the html tag and unrelated characters. For each group_id, ONLY generate 1 summary of 1 paragraph (e.g. main-1 only generated in 1 summary), which SHOULD contain only between %d and %d sentences. if there are 5 unique group_id, then also return summary of 5 unique group_id. Write the summarization in %s. \nHere are the input: ", minSentences, maxSentences, language)

	for idx, data := range input {
		contents += fmt.Sprintf("group_id: %s, content %d: %s\n", data.GroupID, idx+1, data.Content)
	}
	return contents
}

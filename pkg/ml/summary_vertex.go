package ml

import (
	"fmt"
)

type SummaryVertexRest struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewSummaryVertexRest(projectID, location string, projectLabel ProjectLabel) (*SummaryVertexRest, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("You are a tools for summarization of title of news articles.").
		SetResponseSchema(map[string]interface{}{
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
		}).
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		SetEndpoint(fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, vertex.config.Model))

	return &SummaryVertexRest{vertex: vertex}, nil
}

func (s SummaryVertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("SummarizeVertexRest.(%v)(%v) %w", method, params, err)
}

func (s *SummaryVertexRest) BatchSummarize(content string) ([]Summary, OutputVertex, error) {
	s.vertex.SetContent(content, "USER")

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, OutputVertex{}, s.error(err, "BatchSummarize")
	}

	result, outputVertex := []Summary{}, OutputVertex{}
	err = ParseSingleResponseVertex(resp, &result, &outputVertex)
	if err != nil {
		return nil, OutputVertex{}, s.error(err, "BatchSummarize")
	}

	return result, outputVertex, nil
}

func (s *SummaryVertexRest) ProcessAndBatchSummarize(language string, minSentences, maxSentences int, input []Summary) ([]Summary, OutputVertex, error) {

	contents_text := s.generateContentsText(language, minSentences, maxSentences, input)
	s.vertex.SetContent(contents_text, "USER")

	resp, err := s.vertex.GetResponse()
	if err != nil {
		return nil, OutputVertex{}, s.error(err, "BatchSummarize")
	}

	result, outputVertex := []Summary{}, OutputVertex{}
	err = ParseSingleResponseVertex(resp, &result, &outputVertex)
	if err != nil {
		return nil, OutputVertex{}, s.error(err, "BatchSummarize")
	}

	return result, outputVertex, nil
}

func (s *SummaryVertexRest) generateContentsText(language string, minSentences, maxSentences int, input []Summary) (contents string) {
	contents = fmt.Sprintf("Create a summary of given indonesian news article title. Ignore the html tag and unrelated characters. For each group_id, ONLY generate 1 summary of 1 paragraph (e.g. main-1 only generated in 1 summary), which SHOULD contain only between %d and %d sentences. if there are 5 unique group_id, then also return summary of 5 unique group_id. Write the summarization in %s. \nHere are the input: ", minSentences, maxSentences, language)

	for idx, data := range input {
		contents += fmt.Sprintf("group_id: %s, content %d: %s\n", data.GroupID, idx+1, data.Content)
	}
	return contents
}

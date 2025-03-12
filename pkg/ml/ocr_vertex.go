package ml

import "fmt"

type OCRVertexRest struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewOCRVertexRest(projectID, location string, projectLabel ProjectLabel) (*OCRVertexRest, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("You are a tools for OCR of images. Crop the newspaper/magazine sequentially.").
		SetResponseSchema(map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type": "string",
					},
					"title": map[string]interface{}{
						"type": "string",
					},
					"content": map[string]interface{}{
						"type": "string",
					},
				},
				"required": []string{"id", "content", "title"},
			},
		}).
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		SetEndpoint(fmt.Sprintf("https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", location, projectID, location, vertex.config.Model))
	return &OCRVertexRest{vertex: vertex}, nil
}
func (o OCRVertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("OCRVertexRest.(%v)(%v) %w", method, params, err)
}

func (o *OCRVertexRest) Infer(imageURL, imageType string) ([]ScannedText, OutputVertex, error) {

	o.vertex.ResetContentsParts().
		AddFileData(imageURL, imageType).
		AddContent("Process this indonesian newspaper/magazine article. Keep the quotation marks, but omit the quotes (repeated someone's saying written in a bigger font), image captions, advertisements, or non article text. Take the article titles and contents from the whole page. The content should not contain the title. Separate each paragraph with new line. Preprocess the text so it become tidy, fix the typo if any.", "USER")

	resp, err := o.vertex.
		GetResponse()

	if err != nil {
		return nil, OutputVertex{}, o.error(err, "Infer")
	}
	result, outputVertex := []ScannedText{}, OutputVertex{}
	err = ParseSingleResponseVertex(resp, &result, &outputVertex)
	if err != nil {
		return nil, OutputVertex{}, o.error(err, "Infer")
	}
	return result, outputVertex, nil
}

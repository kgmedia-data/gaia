package ml

import "fmt"

type OCRVertexRest struct {
	vertex *VertexRest
}

func (vertex *VertexRest) NewOCRVertexRest(projectID, location string, projectLabel ProjectLabel) (*OCRVertexRest, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("You are a tools for OCR of images. Crop the newspaper sequentially, and return the cropped image (as image) and article content. Omit the quotes, image captions, advertisement, or some paginations. Make sure to process until the end of each article. Separate each paragraph with new line. Preprocess the text so it become tidy. Only return the full article and coordinate of cropped image.").
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
					"crop_coordinate": map[string]interface{}{
						"type": "OBJECT",
						"properties": map[string]interface{}{
							"x1": map[string]interface{}{
								"type": "integer",
							},
							"x2": map[string]interface{}{
								"type": "integer",
							},
							"y1": map[string]interface{}{
								"type": "integer",
							},
							"y2": map[string]interface{}{
								"type": "integer",
							},
						},
					},
				},
				"required": []string{"id", "content", "title", "crop_coordinate"},
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

func (o *OCRVertexRest) Infer(imageURL, imageType string) ([]ScannedText, error) {

	o.vertex.AddFileData(imageURL, imageType).
		AddContent("process this indonesian newspaper", "USER")

	fmt.Printf("Parts: %+v\n", o.vertex.config.Contents.Parts)

	resp, err := o.vertex.
		GetResponse()

	if err != nil {
		return nil, o.error(err, "Infer")
	}
	result := []ScannedText{}
	err = ParseSingleResponseVertex(resp, &result)
	if err != nil {
		return nil, o.error(err, "Infer")
	}
	return result, nil
}

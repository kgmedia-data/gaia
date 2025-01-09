package ml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfer_OCRRestVertex(t *testing.T) {

	projectID := "kgdata-aiml"
	location := "asia-southeast1"
	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "ocr",
	}

	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	ocr, err := vertex.NewOCRVertexRest(projectID, location, projectLabel)
	assert.NoError(t, err)
	// fmt.Printf("model vertex config: %+v\n\n", ocr.vertex.config)

	resp, err := ocr.Infer("gs://kgdata-aiml-medea/printed_paper/publisher=pos-belitung/dt=2024-06-18/pages/01.jpg", "image/jpeg")
	fmt.Printf("resp: %+v\n", resp)
	assert.NoError(t, err)
	// assert.Equal(t, 5, len(resp))
}

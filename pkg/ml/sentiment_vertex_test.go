package ml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfer_SentimentVertex(t *testing.T) {
	projectID := "kgdata-aiml"
	location := "asia-southeast1"
	vertex, err := NewVertexRest()
	assert.NoError(t, err)

	projectLabel := ProjectLabel{
		ProjectName: "medeab",
		EnvName:     "dev",
		TaskName:    "sentiment",
	}

	model, err := vertex.NewSentimentVertexRest(projectID, location, projectLabel)
	assert.NoError(t, err)
	model.Vertex.SetSystemInstruction("You are a sentiment analysis model, which will be given input of Indonesian language news article. classify positive for news titles that are good for promoting SDGs, negative for news titles that are bad for promoting SDGs, and neutral for news titles that are neither good nor bad for promoting SDGs.")

	contents := make(map[string]string)
	contents["0"] = "kebakaran hutan membuat saya kesal"
	contents["1"] = "saya cinta prabowo"
	contents["2"] = "pemerintah melakukan reboisasi"
	contents["3"] = "saya senang dengan kebakaran hutan"

	resp, err := model.InferBatch(contents)
	fmt.Println("resp", resp)
	assert.NoError(t, err)
}

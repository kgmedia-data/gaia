package ml

import (
	"fmt"
)

type RAGVertexRest struct {
	vertex *VertexRest
}

func NewRAGVertexRest(projectID, location, dataStores string, projectLabel ProjectLabel, vertex *VertexRest) (*RAGVertexRest, error) {

	vertex.SetModel("gemini-1.5-flash-002").
		SetTemperature(1).
		SetMaxOutputTokens(8192).
		AddSystemInstruction("you are an agent that help users understands data in our dashboard better. use bahasa indonesia unless user wants to speak in english. you can give user 6 recommendation actions to do with agent e.g (about latest trend, campaigns, user's competitors, etc), use grounding as the main source of information. If necessary, seek clarifying details. if users asks about data analysis or visualization, process the req uest using code interpreter. show the source of grounding.").
		AddLabel("project", projectLabel.ProjectName).
		AddLabel("env", projectLabel.EnvName).
		AddLabel("task", projectLabel.TaskName).
		AddTools([]Tools{
			{
				Retrieval: Retrieval{
					VertexAISearch: VertexAISearch{
						Datastore: fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/dataStores/%s", projectID, location, dataStores),
					},
				},
			},
		}).
		SetResponseSchema(nil).
		SetEndpoint(fmt.Sprintf("https://aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:generateContent", projectID, location, vertex.config.Model))

	return &RAGVertexRest{
		vertex: vertex,
	}, nil
}

func (r RAGVertexRest) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("RAGVertexRest.(%v)(%v) %w", method, params, err)
}

func (r *RAGVertexRest) Infer(content, role string) (*OutputVertex, error) {
	r.vertex.SetContent(content, role)
	resp, err := r.vertex.GetResponse()
	if err != nil {
		return nil, r.error(err, "Infer", r.vertex.config)
	}

	// Parse Response Body to JSON
	result, err := r.vertex.ParseResponse(resp)
	if err != nil {
		return nil, r.error(err, "BatchSummarize", r.vertex.config)
	}

	return &result, nil
}

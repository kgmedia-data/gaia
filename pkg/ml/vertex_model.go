package ml

// ======================= Vertex AI Config ===============================
type VertexAIConfig struct {
	Model             string            `json:"model"`
	GenerationConfig  GenerationConfig  `json:"generation_config"`
	SystemInstruction SystemInstruction `json:"systemInstruction"`
	Contents          Contents          `json:"contents"`
	Labels            map[string]string `json:"labels"`
	SafetySettings    []SafetySettings  `json:"safetySettings,omitempty"`
	Tools             []Tools           `json:"tools"`
}

type GenerationConfig struct {
	Temperature      float64                `json:"temperature,omitempty"`
	TopP             int                    `json:"topP,omitempty"`
	TopK             *int                   `json:"topK,omitempty"`
	CandidateCount   int                    `json:"candidateCount,omitempty"`
	MaxOutputTokens  int                    `json:"maxOutputTokens,omitempty"`
	ResponseMimeType string                 `json:"responseMimeType,omitempty"`
	ResponseSchema   map[string]interface{} `json:"responseSchema,omitempty"`
}

type SystemInstruction struct {
	Parts []InstructionPart `json:"parts"`
}

type Contents struct {
	Role  string            `json:"role"`
	Parts []InstructionPart `json:"parts"`
}

type InstructionPart struct {
	Text string `json:"text"`
}

type SafetySettings struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

type Tools struct {
	Retrieval Retrieval `json:"retrieval"`
}

type Retrieval struct {
	VertexAISearch VertexAISearch `json:"vertexAiSearch"`
}

type VertexAISearch struct {
	Datastore string `json:"datastore"`
}

type ProjectLabel struct {
	ProjectName string
	EnvName     string
	TaskName    string
}

// ======================= Vertex AI Output ===============================
type OutputVertex struct {
	Candidates []struct {
		AvgLogprobs float64 `json:"avgLogprobs"`
		Content     struct {
			Role  string `json:"role"`
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
		GroundingMetadata struct {
			GroundingChunks []struct {
				RetrievedContext struct {
					Text string `json:"text"`
				} `json:"retrievedContext"`
			} `json:"groundingChunks"`
		} `json:"groundingMetadata"`
		FinishReason  string `json:"finishReason"`
		SafetyRatings []struct {
			Category         string  `json:"category"`
			Probability      string  `json:"probability"`
			ProbabilityScore float64 `json:"probabilityScore"`
			Severity         string  `json:"severity"`
			SeverityScore    float64 `json:"severityScore"`
		} `json:"safetyRatings"`
	} `json:"candidates"`
	UsageMetadata struct {
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		PromptTokenCount     int `json:"promptTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
	} `json:"usageMetadata"`
	ModelVersion string `json:"modelVersion"`
}

type Summary struct {
	GroupID string `json:"group_id"`
	Content string `json:"content"`
}

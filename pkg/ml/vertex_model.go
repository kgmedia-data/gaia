package ml

type Summary struct {
	GroupID string `json:"group_id"`
	Content string `json:"content"`
}

// Vertex AI Config
type VertexAIConfig struct {
	Model             string            `json:"model"`
	GenerationConfig  GenerationConfig  `json:"generation_config"`
	SystemInstruction SystemInstruction `json:"systemInstruction"`
	Contents          Contents          `json:"contents"`
	Labels            map[string]string `json:"labels"`
	SafetySettings    []SafetySettings  `json:"safetySettings,omitempty"`
}

type GenerationConfig struct {
	Temperature      float64                `json:"temperature"`
	TopP             int                    `json:"topP"`
	TopK             int                    `json:"topK"`
	CandidateCount   int                    `json:"candidateCount"`
	MaxOutputTokens  int                    `json:"maxOutputTokens"`
	ResponseMimeType string                 `json:"responseMimeType"`
	ResponseSchema   map[string]interface{} `json:"responseSchema"`
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

// Vertex AI Output
type OutputVertex struct {
	Candidates []struct {
		AvgLogprobs float64 `json:"avgLogprobs"`
		Content     struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
		Role          string `json:"role"`
		FinishReason  string `json:"finishReason"`
		SafetyRatings []struct {
			Category         string  `json:"category"`
			Probability      string  `json:"probability"`
			ProbabilityScore float64 `json:"probabilityScore"`
			Severity         string  `json:"severity"`
			SeverityScore    float64 `json:"severityScore"`
		} `json:"safetyRatings"`
	} `json:"candidates"`
	ModelVersion  string `json:"modelVersion"`
	UsageMetadata struct {
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		PromptTokenCount     int `json:"promptTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
	} `json:"usageMetadata"`
}

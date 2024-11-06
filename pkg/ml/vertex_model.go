package ml

type Summary struct {
	GroupID string `json:"group_id"`
	Content string `json:"content"`
}

// Vertex AI Config
type VertexAIConfig struct {
	Model             string            `json:"model"`
	Title             string            `json:"title"`
	Description       string            `json:"description"`
	Parameters        Parameters        `json:"parameters"`
	SystemInstruction SystemInstruction `json:"systemInstruction"`
	Prompt            Prompt            `json:"prompt"`
	InputPrefixes     []string          `json:"inputPrefixes"`
	OutputPrefixes    []string          `json:"outputPrefixes"`
}

type Parameters struct {
	StopSequences    []string `json:"stopSequences"`
	Temperature      float64  `json:"temperature"`
	TokenLimits      int      `json:"tokenLimits"`
	TopP             int      `json:"topP"`
	TopK             int      `json:"topK"`
	CandidateCount   int      `json:"candidateCount"`
	MaxOutputTokens  int      `json:"maxOutputTokens"`
	ResponseMimeType string   `json:"responseMimeType"`
}

type SystemInstruction struct {
	Parts []InstructionPart `json:"parts"`
}

type Prompt struct {
	Parts []InstructionPart `json:"parts"`
}

type InstructionPart struct {
	Text string `json:"text"`
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

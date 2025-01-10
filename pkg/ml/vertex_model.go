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
	Text     *string   `json:"text,omitempty"`
	FileData *FileData `json:"fileData,omitempty"`
}

type FileData struct {
	FileUri  string `json:"fileUri,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
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

type EntitySentiment struct {
	ID          int    `json:"id"`
	EntityID    int    `json:"entity_id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Sentiment   string `json:"sentiment"`
	IsMentioned bool   `json:"is_mentioned"`
}

type Sentiment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Sentiment string `json:"sentiment"`
}

type ScannedText struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	CropCoordinate CropCoordinate `json:"crop_coordinate"`
}
type CropCoordinate struct {
	X1 int `json:"x1"`
	X2 int `json:"x2"`
	Y1 int `json:"y1"`
	Y2 int `json:"y2"`
}

func (c *CropCoordinate) IsValid() bool {
	return c.X1 >= 0 && c.X2 >= 0 && c.Y1 >= 0 && c.Y2 >= 0
}

type Spokesperson struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	JobTitle   string   `json:"job_title"`
	Sentiment  string   `json:"sentiment"`
	Statements []string `json:"statements"`
}

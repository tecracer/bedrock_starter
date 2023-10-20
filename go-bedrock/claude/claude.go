package claude

type ModelId string

var ModelIdClaudeV1 ModelId = "anthropic.claude-v1"
var ModelIdClaudeInstantV1 ModelId = "anthropic.claude-instant-v1"

type ClaudeV1 struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature"`
	TopK              int      `json:"top_k"`
	TopP              float64  `json:"top_p"`
	StopSequences     []string `json:"stop_sequences"`
	AnthropicVersion  ModelId  `json:"anthropic_version"`
}

type ResponseError struct {
	Error struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

type Response struct {
	Completion string `json:"completion"`
	StopReason string `json:"stop_reason"`
	Model      string `json:"model"`
}

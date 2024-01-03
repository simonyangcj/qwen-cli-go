package api

import (
	"qwen-cli/cmd/qwen/app/config"
)

type Input struct {
	Conversation []Message `json:"messages,omitempty"`
}

type Message struct {
	// Message content
	Content string `json:"content"`
	Role    string `json:"role"`
}

type ApiParameter config.ParameterConfig

type Response struct {
	Output Output `json:"output"`
}

type Output struct {
	// checkout api doc https://help.aliyun.com/zh/dashscope/developer-reference/api-details?spm=5176.28630291.0.0.5ef87eb5UNmxpe&disableWebsiteRedirect=true
	Text         string   `json:"text"`          // only working on result_format=text
	FinishReason string   `json:"finish_reason"` // only working on result_format=text
	Choices      []Choice `json:"choices"`       // only working on result_format=message
	OutputTokens int      `json:"output_tokens"`
	InputTokens  int      `json:"input_tokens"`
	RequestID    string   `json:"request_id"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Request struct {
	Input      *Input        `json:"input"`
	Parameters *ApiParameter `json:"parameters,omitempty"`
	Model      string        `json:"model"`
}

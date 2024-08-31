package openai

import (
	"context"

	goopenai "github.com/sashabaranov/go-openai"
)

func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request goopenai.ChatCompletionRequest,
	langfuseTraceOptions ...LangfuseTraceOption,
) (response goopenai.ChatCompletionResponse, err error) {
	// record langfuse tracing events
	return c.OpenAIClient.CreateChatCompletion(ctx, request)
}

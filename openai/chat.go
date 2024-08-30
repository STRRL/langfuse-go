package openai

import (
	"context"

	goopenai "github.com/sashabaranov/go-openai"
)

func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request goopenai.ChatCompletionRequest,
) (response goopenai.ChatCompletionResponse, err error) {
	return c.OpenAIClient.CreateChatCompletion(ctx, request)
}

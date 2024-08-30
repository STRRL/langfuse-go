package openai

import (
	"context"

	goopenai "github.com/sashabaranov/go-openai"
)

func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request goopenai.ChatCompletionRequest,
) (stream *goopenai.ChatCompletionStream, err error) {
	return c.OpenAIClient.CreateChatCompletionStream(
		ctx,
		request,
	)
}

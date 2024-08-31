package openai

import (
	"net/http"

	goopenai "github.com/sashabaranov/go-openai"
)

type Streamable interface {
	goopenai.ChatCompletionStreamResponse | goopenai.CompletionResponse
}

type StreamReader[T Streamable] interface {
	Recv() (T, error)
	Close() error
	GetRateLimitHeaders() goopenai.RateLimitHeaders
	Header() http.Header
}

type ChatCompletionStream StreamReader[goopenai.ChatCompletionStreamResponse]

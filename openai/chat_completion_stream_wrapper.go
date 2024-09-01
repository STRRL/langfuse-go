package openai

import (
	hlangfuse "github.com/henomis/langfuse-go"
	"github.com/henomis/langfuse-go/model"
	goopenai "github.com/sashabaranov/go-openai"
	"io"
	"net/http"
)

var _ ChatCompletionStream = &chatCompletionStreamWrapper{}

type chatCompletionStreamWrapper struct {
	traceID       string
	observationID string
	upstream      *goopenai.ChatCompletionStream

	aggregator     *streamResponseAggregator
	langfuseClient *hlangfuse.Langfuse
}

func NewChatCompletionStreamWrapper(traceID string, observationID string, upstream *goopenai.ChatCompletionStream, langfuseClient *hlangfuse.Langfuse) ChatCompletionStream {
	return &chatCompletionStreamWrapper{
		traceID:        traceID,
		observationID:  observationID,
		upstream:       upstream,
		aggregator:     newStreamResponseAggregator(),
		langfuseClient: langfuseClient,
	}
}

func (c *chatCompletionStreamWrapper) Recv() (goopenai.ChatCompletionStreamResponse, error) {
	result, err := c.upstream.Recv()
	if err != nil {
		if io.EOF == err {
			aggregatedResponse, ttft := c.aggregator.Done()
			_, err := c.langfuseClient.GenerationEnd(&model.Generation{
				TraceID:             c.traceID,
				ID:                  c.observationID,
				Output:              aggregatedResponse,
				CompletionStartTime: &ttft,
			})
			if err != nil {
				// TODO: logging
			}
		}

		return result, err
	} else {
		c.aggregator.Append(result)
	}
	return result, err
}

func (c *chatCompletionStreamWrapper) Close() error {
	return c.upstream.Close()
}

func (c *chatCompletionStreamWrapper) GetRateLimitHeaders() goopenai.RateLimitHeaders {
	return c.upstream.GetRateLimitHeaders()
}

func (c *chatCompletionStreamWrapper) Header() http.Header {
	return c.upstream.Header()
}

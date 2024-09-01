package openai

import (
	"context"
	"time"

	"github.com/henomis/langfuse-go/model"
	goopenai "github.com/sashabaranov/go-openai"
)

func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request goopenai.ChatCompletionRequest,
	langfuseTraceOptions ...LangfuseTraceOption,
) (ChatCompletionStream, error) {
	traceID := ""
	options := ApplyTraceOptions(langfuseTraceOptions)
	if options != nil {
		if options.TraceID == "" {
			// create a new trace
			traceModel := model.Trace{
				Name:   "chat-completion-stream",
				Public: true,
			}

			traceModel.UserID = options.UserID
			traceModel.SessionID = options.SessionID
			traceModel.Release = options.Release
			traceModel.Metadata = options.Metadata
			traceModel.Tags = options.Tags

			t, err := c.LangfuseClient.Trace(&traceModel)
			if err != nil {
				return nil, err
			}
			traceID = t.ID
		} else {
			traceID = options.TraceID
		}
	}

	requestStartTime := time.Now()

	g, err := c.LangfuseClient.Generation(&model.Generation{
		Name:    "chat-completion-stream",
		TraceID: traceID,
		Input:   request,
		Model:   request.Model,
		ModelParameters: model.M{
			"stream":      true,
			"temperature": request.Temperature,
			"top_p":       request.TopP,
			"max_tokens":  request.MaxTokens,
			"stop":        request.Stop,
			"n":           request.N,
		},
		StartTime: &requestStartTime,
	}, nil)
	if err != nil {
		return nil, err
	}
	stream, err := c.OpenAIClient.CreateChatCompletionStream(
		ctx,
		request,
	)

	if err != nil {
		return nil, err
	}

	return NewChatCompletionStreamWrapper(traceID, g.ID, stream, c.LangfuseClient), nil
}

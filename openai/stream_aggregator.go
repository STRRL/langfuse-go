package openai

import goopenai "github.com/sashabaranov/go-openai"

type AggregatedResponse struct {
	Completion string              `json:"completion,omitempty"`
	ToolCallID string              `json:"tool_call_id,omitempty"`
	ToolCalls  []goopenai.ToolCall `json:"tool_calls,omitempty"`
}

type streamResponseAggregator struct {
	buf []goopenai.ChatCompletionStreamResponse
}

func newStreamResponseAggregator() *streamResponseAggregator {
	return &streamResponseAggregator{}
}

func (s *streamResponseAggregator) Append(item goopenai.ChatCompletionStreamResponse) {
	s.buf = append(s.buf, item)
}

func (s *streamResponseAggregator) Done() AggregatedResponse {
	result := AggregatedResponse{}
	for _, item := range s.buf {
		if item.Choices == nil || len(item.Choices) == 0 {
			continue
		}

		choice := item.Choices[0]

		if choice.Delta.Content != "" {
			result.Completion += choice.Delta.Content
		}

		if choice.Delta.ToolCalls != nil {
			for _, toolCall := range choice.Delta.ToolCalls {
				if toolCall.Index == nil {
					continue
				}
				for len(result.ToolCalls) <= *toolCall.Index {
					result.ToolCalls = append(result.ToolCalls, goopenai.ToolCall{})
				}
				currentToolCall := &result.ToolCalls[*toolCall.Index]
				if toolCall.ID != "" {
					currentToolCall.ID = toolCall.ID
				}
				if toolCall.Type != "" {
					currentToolCall.Type = toolCall.Type
				}

				if toolCall.Function.Name != "" {
					currentToolCall.Function.Name = toolCall.Function.Name
				}
				if toolCall.Function.Arguments != "" {
					currentToolCall.Function.Arguments += toolCall.Function.Arguments
				}

				if result.ToolCallID == "" && currentToolCall.ID != "" {
					result.ToolCallID = currentToolCall.ID
				}
			}
		}
	}

	return result
}
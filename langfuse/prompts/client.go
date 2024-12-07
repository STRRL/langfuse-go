package prompts

import (
	"context"

	"github.com/STRRL/langfuse-go/langfuse/openapi"
)

type PromptsClientInterface interface {
	PromptsList(ctx context.Context, params *openapi.PromptsListParams, reqEditors ...openapi.RequestEditorFn) (*openapi.PromptsListResponse, error)
	PromptsGet(ctx context.Context, promptName string, params *openapi.PromptsGetParams, reqEditors ...openapi.RequestEditorFn) (*openapi.PromptsGetResponse, error)
}

var _ PromptsClientInterface = &PromptsClient{}

type PromptsClient struct {
	c openapi.ClientWithResponsesInterface
}

func NewPromptsClient(c openapi.ClientWithResponsesInterface) *PromptsClient {
	return &PromptsClient{c: c}
}

func (c *PromptsClient) PromptsList(ctx context.Context, params *openapi.PromptsListParams, reqEditors ...openapi.RequestEditorFn) (*openapi.PromptsListResponse, error) {
	return c.c.PromptsListWithResponse(ctx, params, reqEditors...)
}

func (c *PromptsClient) PromptsGet(ctx context.Context, promptName string, params *openapi.PromptsGetParams, reqEditors ...openapi.RequestEditorFn) (*openapi.PromptsGetResponse, error) {
	return c.c.PromptsGetWithResponse(ctx, promptName, params, reqEditors...)
}

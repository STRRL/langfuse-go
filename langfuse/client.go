package langfuse

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/STRRL/langfuse-go/langfuse/openapi"
	"github.com/STRRL/langfuse-go/langfuse/prompts"
)

const (
	BASE_URL    = "https://cloud.langfuse.com"
	BASE_URL_EU = "https://eu.cloud.langfuse.com"
	BASE_URL_US = "https://us.cloud.langfuse.com"
)

var _ prompts.PromptsClientInterface = &LangfuseClient{}

type LangfuseClient struct {
	publicKey string
	secretKey string
	prompts.PromptsClientInterface
}

// NewLangfuseClient creates a new client with public and secret keys for authentication
func NewLangfuseClient(publicKey, secretKey string) (*LangfuseClient, error) {
	// Set basic auth header
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", publicKey, secretKey)))
	c, err := openapi.NewClientWithResponses(BASE_URL,
		openapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Basic "+auth)
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return &LangfuseClient{
		publicKey:              publicKey,
		secretKey:              secretKey,
		PromptsClientInterface: prompts.NewPromptsClient(c),
	}, nil
}

// NewLangfuseClientWithCustomBaseURL creates a new client with public and secret keys for authentication
func NewLangfuseClientWithCustomBaseURL(publicKey, secretKey, baseURL string) (*LangfuseClient, error) {
	// Set basic auth header
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", publicKey, secretKey)))
	c, err := openapi.NewClientWithResponses(baseURL,
		openapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Basic "+auth)
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return &LangfuseClient{
		publicKey:              publicKey,
		secretKey:              secretKey,
		PromptsClientInterface: prompts.NewPromptsClient(c),
	}, nil
}

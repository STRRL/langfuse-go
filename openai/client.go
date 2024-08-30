package openai

import (
	"github.com/STRRL/langfuse-go/langfuse"
	goopenai "github.com/sashabaranov/go-openai"
)

type Client struct {
	OpenAIClient   goopenai.Client
	LangfuseClient langfuse.Client
}

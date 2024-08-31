package openai

import (
	hlangfuse "github.com/henomis/langfuse-go"
	goopenai "github.com/sashabaranov/go-openai"
)

type Client struct {
	OpenAIClient   *goopenai.Client
	LangfuseClient *hlangfuse.Langfuse
}

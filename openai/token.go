package openai

import (
	"fmt"
	"github.com/pkoukk/tiktoken-go"
	goopenai "github.com/sashabaranov/go-openai"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var tkm *tiktoken.Tiktoken

func init() {
	var err error
	tkm, err = loadCl00k()
	if err != nil {
		err = fmt.Errorf("encoding for model failed: %v", err)
		log.Println(err)
	}
}

func downloadOrCache(url string, path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	http.Get(url)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

const (
	tkm_tokenPerMessage = 3
)

// GPT3 & GPT4 use CL100K encoding
func loadCl00k() (*tiktoken.Tiktoken, error) {
	bpeLoader := tiktoken.NewDefaultBpeLoader()
	tempPath, err := os.MkdirTemp("", "langfuse-go")
	if err != nil {
		return nil, err
	}
	path := tempPath + "/cl100k_base.tiktoken"
	err = downloadOrCache("https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken", path)
	ranks, err := bpeLoader.LoadTiktokenBpe(path)
	if err != nil {
		return nil, err
	}
	special_tokens := map[string]int{
		tiktoken.ENDOFTEXT:   100257,
		tiktoken.FIM_PREFIX:  100258,
		tiktoken.FIM_MIDDLE:  100259,
		tiktoken.FIM_SUFFIX:  100260,
		tiktoken.ENDOFPROMPT: 100276,
	}
	enc := &tiktoken.Encoding{
		Name:           tiktoken.MODEL_CL100K_BASE,
		PatStr:         `(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`,
		MergeableRanks: ranks,
		SpecialTokens:  special_tokens,
	}
	pbe, err := tiktoken.NewCoreBPE(enc.MergeableRanks, enc.SpecialTokens, enc.PatStr)
	if err != nil {
		return nil, err
	}
	specialTokensSet := map[string]any{}
	for k := range enc.SpecialTokens {
		specialTokensSet[k] = true
	}
	return tiktoken.NewTiktoken(pbe, enc, specialTokensSet), nil
}

func GetOpenAITokenCount(messages []goopenai.ChatCompletionMessage) (numTokens int) {
	if tkm == nil {
		return -1
	}

	for _, message := range messages {
		numTokens += tkm_tokenPerMessage
		if message.Content != "" {
			numTokens += len(tkm.Encode(message.Content, nil, nil))
		}
		for _, item := range message.MultiContent {
			if item.Type == goopenai.ChatMessagePartTypeText {
				numTokens += len(tkm.Encode(item.Text, nil, nil))
			} else if item.Type == goopenai.ChatMessagePartTypeImageURL {
				numTokens += 75
			} else {
				slog.WarnContext(nil, "unsupported message content type", "message", message)
			}
		}

		for _, toolCall := range message.ToolCalls {
			numTokens += len(tkm.Encode(toolCall.ID, nil, nil))
			numTokens += len(tkm.Encode(string(toolCall.Type), nil, nil))
			numTokens += len(tkm.Encode(toolCall.Function.Name, nil, nil))
			numTokens += len(tkm.Encode(toolCall.Function.Arguments, nil, nil))
		}

		numTokens += len(tkm.Encode(string(message.Role), nil, nil))
	}
	return numTokens
}

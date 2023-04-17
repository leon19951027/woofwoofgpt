package openai

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/leon19951027/woofwoofgpt/config"

	oai "github.com/sashabaranov/go-openai"
)

func ApplyCfg(cfg *config.Cfg) {

}

type ChatMessages struct {
	Ms []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Message string `json:"message"`
}

func Chat(messages *ChatMessages, apiKey string, baseUrl string, responseChan chan string, doneChan chan bool) {
	config := oai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl

	var completionMessages []oai.ChatCompletionMessage
	c := oai.NewClientWithConfig(config)
	fmt.Println(messages.Ms)
	for _, msg := range messages.Ms {
		ccm := oai.ChatCompletionMessage{
			Content: msg.Message,
			Role:    oai.ChatMessageRoleUser,
		}
		completionMessages = append(completionMessages, ccm)

	}
	req := oai.ChatCompletionRequest{
		Model:     oai.GPT3Dot5Turbo,
		MaxTokens: 2000,
		Messages:  completionMessages,
		Stream:    true,
	}
	ctx := context.Background()

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {

			fmt.Println("\nStream finished")
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		//	fmt.Printf(response.Choices[0].Delta.Content)
		responseChan <- response.Choices[0].Delta.Content
	}

	doneChan <- true
}

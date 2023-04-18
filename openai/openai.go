package openai

import (
	"context"
	"errors"
	"fmt"
	"io"

	oai "github.com/sashabaranov/go-openai"
)

type ChatMessages struct {
	Ms []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Message string `json:"message"`
	Role    string `json:"role"`
}

func Chat(messages *ChatMessages, apiKey string, baseUrl string, responseChan chan string, doneChan chan bool) {
	config := oai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl

	var completionMessages []oai.ChatCompletionMessage
	c := oai.NewClientWithConfig(config)

	for _, msg := range messages.Ms {

		ccm := oai.ChatCompletionMessage{
			Content: msg.Message,
		}
		switch msg.Role {
		case "system":
			ccm.Role = oai.ChatMessageRoleSystem
		default:
			ccm.Role = oai.ChatMessageRoleUser
		}

		completionMessages = append(completionMessages, ccm)

	}
	fmt.Println(completionMessages)
	req := oai.ChatCompletionRequest{
		Model:     oai.GPT3Dot5Turbo,
		MaxTokens: 4000,
		Messages:  completionMessages,
		Stream:    true,
	}
	ctx := context.Background()

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		responseChan <- err.Error()
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\n OpenAI Stream finished")
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			responseChan <- err.Error()
			doneChan <- true
			return
		}
		responseChan <- response.Choices[0].Delta.Content
	}
	doneChan <- true
}

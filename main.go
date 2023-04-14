package main

import (
	"myopenai/config"
	"myopenai/web"
)

func main() {
	cfg := config.ReadConf()
	websvc := &web.Web{}
	websvc.ApplyCfg(cfg)

	// config := openai.DefaultConfig("")

	// config.BaseURL = "https://woofgpt.uk/v1"

	// c := openai.NewClientWithConfig(config)

	// req := openai.ChatCompletionRequest{
	// 	Model:     openai.GPT3Dot5Turbo,
	// 	MaxTokens: 500,
	// 	Messages: []openai.ChatCompletionMessage{
	// 		{
	// 			Role:    openai.ChatMessageRoleUser,
	// 			Content: "使用golang写helloworld",
	// 		},
	// 	},
	// 	Stream: true,
	// }
	// ctx := context.Background()

	// stream, err := c.CreateChatCompletionStream(ctx, req)
	// if err != nil {
	// 	fmt.Printf("ChatCompletionStream error: %v\n", err)
	// 	return
	// }
	// defer stream.Close()

	// fmt.Printf("Stream response: ")

	// for {
	// 	response, err := stream.Recv()
	// 	if errors.Is(err, io.EOF) {

	// 		fmt.Println("\nStream finished")
	// 		break
	// 	}

	// 	if err != nil {
	// 		fmt.Printf("\nStream error: %v\n", err)
	// 		return
	// 	}

	// 	fmt.Printf(response.Choices[0].Delta.Content)
	// }

	// req2 := openai.ChatCompletionRequest{

	// 	Model:     openai.GPT3Dot5Turbo,
	// 	MaxTokens: 500,
	// 	Messages: []openai.ChatCompletionMessage{
	// 		{
	// 			Role:    openai.ChatMessageRoleUser,
	// 			Content: "使用golang写helloworld",
	// 		},
	// 		{
	// 			Role:    openai.ChatMessageRoleUser,
	// 			Content: "我刚才问你什么",
	// 		},
	// 	},
	// 	Stream: true,
	// }
	// ctx2 := context.Background()

	// stream2, err := c.CreateChatCompletionStream(ctx2, req2)
	// if err != nil {
	// 	fmt.Printf("ChatCompletionStream error: %v\n", err)
	// 	return
	// }
	// defer stream2.Close()

	// fmt.Printf("Stream2 response: ")
	// for {
	// 	response, err := stream2.Recv()
	// 	if errors.Is(err, io.EOF) {
	// 		fmt.Println("\nStream finished")
	// 		return
	// 	}

	// 	if err != nil {
	// 		fmt.Printf("\nStream error: %v\n", err)
	// 		return
	// 	}

	// 	fmt.Printf(response.Choices[0].Delta.Content)
	// }
}

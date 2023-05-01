package main

import (
	"context"
	"fmt"
	"os"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	c := openai.NewClient(os.Getenv("OPENAI_APIKEY"))  // set your API key as a environemnt variable
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,  //pick models from here https://github.com/sashabaranov/go-openai/blob/master/completion.go#L19
		MaxTokens: 30,
		Prompt:    "Say this is a test",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println("Result = ",resp.Choices[0].Text)
	fmt.Println("Finish Reason = ",resp.Choices[0].FinishReason)
	fmt.Println("Total Token = ", resp.Usage.TotalTokens)
}
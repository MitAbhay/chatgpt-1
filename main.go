package main

import (
	"fmt"
	"context"
	"log"
	"github.com/joho/godotenv"
	"github.com/PullRequestInc/go-gpt3"
	"os"
)

func main() {	
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")

	if(apiKey == "") {
		log.Fatalln("API_KEY not found")
	}

	cntxt := context.Background()
	client := gpt3.NewClient(apiKey)

	resp , err := client.Completion(cntxt, gpt3.CompletionRequest{
            Prompt: []string{"My name is"},
			MaxTokens: gpt3.IntPtr(30),
			Stop: []string{"."},
			Echo: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Choices[0].Text)

} 
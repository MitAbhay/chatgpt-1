package main

import (
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
}
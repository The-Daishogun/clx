package main

import (
	"clx/utils"
	"fmt"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

const (
	endpoint = "https://api.groq.com/openai/v1"
)

var client *openai.Client
var config *utils.Config

func init() {
	config = utils.GetConfig()
	err := config.CheckConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	openai_config := openai.DefaultConfig(config.Token)
	openai_config.BaseURL = endpoint
	client = openai.NewClientWithConfig(openai_config)
}

func main() {
	stdInput := utils.ReadStdIn()
	if len(os.Args) < 2 && stdInput == "" {
		fmt.Println("Usage: clx <prompt>")
		os.Exit(1)
	}
	phrase := strings.Join(os.Args[1:], " ")
	if err := utils.AskAI(client, config, phrase, stdInput); err != nil {
		log.Fatal(err)
	}
}

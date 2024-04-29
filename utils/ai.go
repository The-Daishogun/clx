package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func getSystemPrompt() string {
	systemInfo := fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
	prompt := fmt.Sprintf(`
		You are CLX, a CLI code generator. Respond with the CLI command to generate the code with only one short sentence description in first line.
		If CLI command is multiple lines, separate each line with a newline character.
		Do not write any markdown. Do not write any code.
		System Info: %s
		First line is the description in one sentence.
		Example output:
		Building and installing a Go binary
		go build main.go
		go install main
	`, systemInfo)
	return prompt
}

func AskAI(client *openai.Client, config *Config, prompt string) error {
	req := openai.ChatCompletionRequest{
		Model: config.ModelID,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: getSystemPrompt()},
			{Role: "user", Content: prompt},
		},
	}

	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return err
	}
	defer stream.Close()

	firstLine := true
	for {
		response, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Print("\n\n")
				return nil
			}
			return err
		}

		content := response.Choices[0].Delta.Content
		lines := strings.Split(content, "\n")

		for i, line := range lines {
			if line != "" {
				if firstLine {
					fmt.Printf("\x1b[1;35m%s\x1b[0m", line)
				} else {
					fmt.Print(line)
				}
			}

			if !firstLine && len(lines) > 1 && i != 0 {
				fmt.Print("\n\x1b[1;32m$ \x1b[0m")
			}

			if firstLine && len(lines) > 1 {
				fmt.Print("\n")
				firstLine = false
			}
		}

	}

}

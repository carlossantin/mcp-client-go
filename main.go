package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/carlossantin/mcp-agents-go/config"
	"github.com/tmc/langchaingo/llms"
)

func main() {

	ctx := context.Background()

	// Setup from configuration file
	err := config.SetupFromFile(ctx, "config.yaml")
	if err != nil {
		panic(err)
	}

	// Get an agent and generate content
	agent, ok := config.SysConfig.Agents["my-agent"]
	if !ok {
		panic("Agent not found")
	}

	scanner := bufio.NewScanner(os.Stdin)
	msgs := []llms.MessageContent{}
	for {
		fmt.Print("\nEnter your question: ")
		if !scanner.Scan() {
			break
		}
		question := scanner.Text()
		if question == "quit" || question == "exit" {
			break
		}

		fmt.Print("\nGenerating content...\n")

		msgs = append(msgs, llms.MessageContent{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextContent{Text: question}},
		})

		// Comment the following line to enable streaming responses
		// var finalResp string
		// finalResp, msgs = agent.GenerateContent(ctx, msgs, false)
		// fmt.Println(finalResp)

		// Comment the following lines to disable streaming responses
		var textResp <-chan string
		var msgsResp <-chan llms.MessageContent
		textResp, msgsResp = agent.GenerateContentAsStreaming(ctx, msgs, true)

		// Process both channels concurrently
		go func() {
			for resp := range msgsResp {
				msgs = append(msgs, resp)
			}
		}()

		for resp := range textResp {
			fmt.Print(resp)
		}
		fmt.Println()
	}
}

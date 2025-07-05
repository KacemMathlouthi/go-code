package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KacemMathlouthi/go-code/config"
	"github.com/KacemMathlouthi/go-code/utils"
	"github.com/openai/openai-go"
)

func GetLlmResponse(user_prompt string) (string, error) {
	client := config.GetOpenAIClient()
	ctx := context.Background()

	// Get system prompt
	systemPrompt, err := GetSystemPrompt()
	if err != nil {
		return "", fmt.Errorf("failed to get system prompt: %v", err)
	}

	param := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(user_prompt),
		},
		Seed:  openai.Int(0),
		Model: openai.ChatModelGPT4oMini,
	}

	completion, err := client.Chat.Completions.New(ctx, param)

	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}

func GetLlmResponseWithTools(user_prompt string) (string, error) {
	client := config.GetOpenAIClient()
	ctx := context.Background()

	// Get system prompt
	systemPrompt, err := GetSystemPrompt()
	if err != nil {
		return "", fmt.Errorf("failed to get system prompt: %v", err)
	}

	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(user_prompt),
		},
		Tools: utils.ToolsDefinitions,
		Model: openai.ChatModelGPT4oMini,
	}

	// Make initial chat completion request
	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}

	toolCalls := completion.Choices[0].Message.ToolCalls

	// Return early if there are no tool calls
	if len(toolCalls) == 0 {
		return completion.Choices[0].Message.Content, nil
	}

	// If there are tool calls, continue the conversation
	params.Messages = append(params.Messages, completion.Choices[0].Message.ToParam())

	for _, toolCall := range toolCalls {
		// Parse tool arguments
		var toolArgs map[string]string
		if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &toolArgs); err != nil {
			return "", fmt.Errorf("failed to parse tool arguments: %v", err)
		}

		// Execute the tool
		toolResult, err := utils.ExecuteTool(toolCall.Function.Name, toolArgs)
		if err != nil {
			return "", fmt.Errorf("failed to execute tool %s: %v", toolCall.Function.Name, err)
		}

		// Add tool result to messages
		params.Messages = append(params.Messages, openai.ToolMessage(toolResult, toolCall.ID))
	}

	// Make final completion request with tool results
	finalCompletion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}

	return finalCompletion.Choices[0].Message.Content, nil
}

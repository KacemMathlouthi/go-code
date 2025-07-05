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

func GetLlmResponseWithTools(conversationHistory []openai.ChatCompletionMessageParamUnion) (string, error) {
	client := config.GetOpenAIClient()
	ctx := context.Background()

	// Get system prompt
	systemPrompt, err := GetSystemPrompt()
	if err != nil {
		return "", fmt.Errorf("failed to get system prompt: %v", err)
	}

	// Build messages array with system prompt and conversation history
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemPrompt),
	}

	// Add conversation history (which should already include the current user message)
	messages = append(messages, conversationHistory...)

	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Tools:    utils.ToolsDefinitions,
		Model:    openai.ChatModelGPT4oMini,
	}

	// Multi-step tool calling loop
	maxIterations := 10 // Prevent infinite loops
	for iteration := 0; iteration < maxIterations; iteration++ {
		// Make chat completion request
		completion, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			return "", err
		}

		// Add the assistant's response to the conversation
		params.Messages = append(params.Messages, completion.Choices[0].Message.ToParam())

		toolCalls := completion.Choices[0].Message.ToolCalls

		// If there are no tool calls, we're done
		if len(toolCalls) == 0 {
			return completion.Choices[0].Message.Content, nil
		}

		// Execute all tool calls in this iteration
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

		// Continue to next iteration to see if the LLM wants to make more tool calls
	}

	// If we've reached max iterations, make one final request without tools
	params.Tools = nil
	finalCompletion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}

	return finalCompletion.Choices[0].Message.Content, nil
}

package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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
		Model: openai.ChatModelGPT4_1Mini,
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
		Model:    openai.ChatModelGPT4_1Mini,
	}

	// Log the start of tool-enabled LLM request
	utils.LogInfo("Starting tool-enabled LLM request", "llm", map[string]interface{}{
		"model":               string(openai.ChatModelGPT4_1Mini),
		"conversation_length": len(conversationHistory),
		"tools_available":     len(utils.ToolsDefinitions),
	})

	// Multi-step tool calling loop
	maxIterations := 10 // Prevent infinite loops
	for iteration := 0; iteration < maxIterations; iteration++ {
		utils.LogDebug("LLM iteration", "llm", map[string]interface{}{
			"iteration":      iteration + 1,
			"max_iterations": maxIterations,
		})

		// Make chat completion request
		start := time.Now()
		completion, err := client.Chat.Completions.New(ctx, params)
		duration := time.Since(start)

		if err != nil {
			utils.LogError("LLM request failed in iteration", "llm", map[string]interface{}{
				"iteration": iteration + 1,
				"error":     err.Error(),
			})
			return "", err
		}

		// Log LLM response
		utils.LogLLMResponse(completion.Choices[0].Message.Content, string(openai.ChatModelGPT4_1Mini), duration)

		// Add the assistant's response to the conversation
		params.Messages = append(params.Messages, completion.Choices[0].Message.ToParam())

		toolCalls := completion.Choices[0].Message.ToolCalls

		// If there are no tool calls, we're done
		if len(toolCalls) == 0 {
			utils.LogInfo("LLM completed without tool calls", "llm", map[string]interface{}{
				"iterations_used": iteration + 1,
			})
			return completion.Choices[0].Message.Content, nil
		}

		utils.LogInfo("LLM requested tool calls", "llm", map[string]interface{}{
			"iteration":        iteration + 1,
			"tool_calls_count": len(toolCalls),
		})

		// Execute all tool calls in this iteration
		for i, toolCall := range toolCalls {
			utils.LogDebug("Processing tool call", "tool", map[string]interface{}{
				"iteration":  iteration + 1,
				"tool_index": i + 1,
				"tool_name":  toolCall.Function.Name,
			})

			// Parse tool arguments
			var toolArgs map[string]string
			if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &toolArgs); err != nil {
				utils.LogError("Failed to parse tool arguments", "tool", map[string]interface{}{
					"tool_name": toolCall.Function.Name,
					"arguments": toolCall.Function.Arguments,
					"error":     err.Error(),
				})
				return "", fmt.Errorf("failed to parse tool arguments: %v", err)
			}

			// Log tool call
			utils.LogToolCall(toolCall.Function.Name, toolArgs)

			// Execute the tool
			toolStart := time.Now()
			toolResult, err := utils.ExecuteTool(toolCall.Function.Name, toolArgs)
			toolDuration := time.Since(toolStart)

			// Log tool result
			utils.LogToolResult(toolCall.Function.Name, toolResult, err)
			utils.LogDebug("Tool execution completed", "tool", map[string]interface{}{
				"tool_name":   toolCall.Function.Name,
				"duration":    toolDuration.String(),
				"duration_ms": toolDuration.Milliseconds(),
			})

			if err != nil {
				utils.LogError("Tool execution failed", "tool", map[string]interface{}{
					"tool_name": toolCall.Function.Name,
					"error":     err.Error(),
				})
				return "", fmt.Errorf("failed to execute tool %s: %v", toolCall.Function.Name, err)
			}

			// Add tool result to messages
			params.Messages = append(params.Messages, openai.ToolMessage(toolResult, toolCall.ID))
		}

		// Continue to next iteration to see if the LLM wants to make more tool calls
	}

	// If we've reached max iterations, make one final request without tools
	utils.LogWarning("Reached max iterations, making final request without tools", "llm", map[string]interface{}{
		"max_iterations": maxIterations,
	})

	params.Tools = nil
	finalCompletion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		utils.LogError("Final LLM request failed", "llm", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}

	utils.LogInfo("LLM completed with max iterations", "llm", map[string]interface{}{
		"iterations_used": maxIterations,
	})

	return finalCompletion.Choices[0].Message.Content, nil
}

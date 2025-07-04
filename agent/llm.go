package agent

import (
	"context"

	"github.com/KacemMathlouthi/go-code/config"
	"github.com/openai/openai-go"
)

func GetLlmResponse(user_prompt string) (string, error) {
	client := config.GetOpenAIClient()
	ctx := context.Background()

	param := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
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

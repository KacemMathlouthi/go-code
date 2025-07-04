package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/azure"
)

var openAIClient *openai.Client

// AzureOpenAIConfig
type AzureOpenAIConfig struct {
	APIVersion     string
	Endpoint       string
	APIKey         string
	DeploymentName string
}

func loadEnvConfig() *AzureOpenAIConfig {
	_ = godotenv.Load()
	config := &AzureOpenAIConfig{
		APIVersion:     os.Getenv("AZURE_API_VERSION"),
		Endpoint:       os.Getenv("AZURE_ENDPOINT"),
		APIKey:         os.Getenv("AZURE_API_KEY"),
		DeploymentName: os.Getenv("AZURE_DEPLOYMENT_NAME"),
	}
	return config
}

func GetOpenAIClient() *openai.Client {
	config := loadEnvConfig()

	if openAIClient == nil {
		client := openai.NewClient(
			azure.WithEndpoint(config.Endpoint, config.APIVersion),
			azure.WithAPIKey(config.APIKey),
		)
		openAIClient = &client
	}

	return openAIClient
}

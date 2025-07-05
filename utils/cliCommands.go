package utils

import (
	"fmt"

	"github.com/KacemMathlouthi/go-code/config"
)

const asciiArt = `
         ▄▄▄▄▄                                                              
       ▄▀  ▄█▀▀ ▀▀▄        ▐██▀▀██ ██▀███     ██▀▀██▌███▀██ ███▀██ ▀██▀██▀
     █▀█ ▐▌ ▄█▀▀█▄▐▌       ▐██  ▀▀ ██  ██     ██  ▀▀ ██▌ ██ ▐██ ██▌ ██  ▀
    █  █ ▐█▀ ▀█▄▄ ▀█       ▐██ ▄█▄ ██▄▄██     ██     ██▌▌██ ▐██ ██▌ ████
    ▀█ ▀▀█▌   █▌ █  █      ▐██ ▐██ ██  ██     ██  ▄█▌██▌ ██ ▐██ ██▌ ██ ▄▄
     ▐▀▀▄▄▄█▀▀▐▌ █▄█▀      ▐██▄███ ██▄▄██     ██▄▄██ ███▄██ ▐██▄██▌▄██▄██▌
      █▄   ▄▄▀▀ ▄█
        ▀▀▀▀█▄▄▀▀
`

func GetStartupText() {
	fmt.Print(asciiArt)
	fmt.Println("Welcome! I'm your coding agent. Ask me to create, fix or explain anything!")
	fmt.Println("Type '--help' to see the available commands.")
	fmt.Println()
}

func GetHelpText() {
	fmt.Println("Available commands:")
	fmt.Println("  - Type any text to get a response from the AI agent")
	fmt.Println("  - Type '--clear' to clear conversation history")
	fmt.Println("  - Type '--config' to show the current llm model and tools")
	fmt.Println("  - Type '--help' to show this help message")
	fmt.Println("  - Type '--quit' to exit")
	fmt.Println()
}

func GetConfigText() {
	AzureOpenAIConfig := config.LoadEnvConfig()

	fmt.Printf("Current LLM model: %v\n", AzureOpenAIConfig.DeploymentName)
	fmt.Printf("Current API version: %v\n", AzureOpenAIConfig.APIVersion)
	fmt.Printf("Current API key: %v\n", AzureOpenAIConfig.APIKey)
	fmt.Printf("Current API endpoint: %v\n", AzureOpenAIConfig.Endpoint)

	fmt.Println("Available tools:")
	fmt.Println("  - list: List files in the current directory")
	fmt.Println("  - pwd: Print the current working directory")
	fmt.Println("  - tree: Print the directory tree")
	fmt.Println("  - grep: Search for a pattern in a file")
	fmt.Println("  - shell: Execute a shell command")
	fmt.Println("  - write_file: Write to a file")
	fmt.Println("  - read_file: Read a file")
	fmt.Println("  - delete_file: Delete a file")
}

/*
Copyright © 2025 Kacem Mathlouthi <kacem.math47@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KacemMathlouthi/go-code/agent"
	"github.com/KacemMathlouthi/go-code/utils"
	"github.com/openai/openai-go"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-code",
	Short: "A coding agent in the terminal.",
	Long: `A Coding Agent in the terminal. 
	The agent can execute shell commands, read and write files, and more. 
	It can contribute to your codebase by writing code, fixing bugs, and more.`,
	Run: runInteractive,
}

func runInteractive(cmd *cobra.Command, args []string) {
	// Initialize logger
	if err := utils.InitLogger(); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer utils.CloseLogger()

	utils.GetStartupText()
	scanner := bufio.NewScanner(os.Stdin)
	conversationHistory := []openai.ChatCompletionMessageParamUnion{}

	for {
		fmt.Print(utils.FormatPrompt())
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "--quit" {
			fmt.Println(utils.ColorGreen + utils.ColorBold + "👋 Goodbye!" + utils.ColorReset)
			break
		}

		if strings.ToLower(input) == "--help" {
			utils.GetHelpText()
			continue
		}

		if strings.ToLower(input) == "--config" {
			utils.GetConfigText()
			continue
		}

		if strings.ToLower(input) == "--clear" {
			conversationHistory = []openai.ChatCompletionMessageParamUnion{}
			utils.ClearScreen()
			continue
		}

		if input == "" {
			fmt.Println(utils.ColorYellow + "Please enter some text." + utils.ColorReset)
			continue
		}

		// Log user input
		utils.LogInfo("User input received", "interaction", map[string]interface{}{
			"input_length":        len(input),
			"conversation_length": len(conversationHistory),
		})

		// Display user input in a formatted box
		fmt.Println(utils.FormatUserInput(input))

		// Add user message to conversation history
		conversationHistory = append(conversationHistory, openai.UserMessage(input))

		output, err := agent.GetLlmResponseWithTools(conversationHistory)
		if err != nil {
			utils.LogError("LLM response failed", "interaction", map[string]interface{}{
				"error": err.Error(),
			})
			fmt.Println(utils.FormatError(err.Error()))
			continue
		}

		// Add assistant response to conversation history
		conversationHistory = append(conversationHistory, openai.AssistantMessage(output))

		// Display AI response in a formatted box with markdown rendering
		fmt.Println(utils.FormatAIResponse(output))
		fmt.Println()
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-code.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

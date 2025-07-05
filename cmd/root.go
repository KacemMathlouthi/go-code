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
	utils.GetStartupText()
	scanner := bufio.NewScanner(os.Stdin)
	conversationHistory := []openai.ChatCompletionMessageParamUnion{}

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "--quit" {
			fmt.Println("Goodbye!")
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
			// Clear the terminal screen
			fmt.Print("\033[H\033[2J")
			fmt.Println("Terminal cleared!")
			fmt.Println()
			continue
		}

		if input == "" {
			fmt.Println("Please enter some text.")
			continue
		}

		// Add user message to conversation history
		conversationHistory = append(conversationHistory, openai.UserMessage(input))

		output, err := agent.GetLlmResponseWithTools(conversationHistory)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Add assistant response to conversation history
		conversationHistory = append(conversationHistory, openai.AssistantMessage(output))

		fmt.Printf("Response: %s\n", output)
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

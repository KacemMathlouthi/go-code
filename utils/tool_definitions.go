package utils

import (
	"github.com/openai/openai-go"
)

var ToolsDefinitions = []openai.ChatCompletionToolParam{
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "shell",
			Description: openai.String("Run a shell command and return its output, useful for flexible system operations, using package managers, installing dependencies, creating  projects..."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"command": map[string]interface{}{
						"type":        "string",
						"description": "The exact shell command to be executed (e.g., 'ls -la', 'cat file.txt'). Be cautious with destructive commands.",
					},
				},
				"required": []string{"command"},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "grep",
			Description: openai.String("Search for a specific text pattern inside a file using regular expressions."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"pattern": map[string]interface{}{
						"type":        "string",
						"description": "The regular expression pattern to search for (e.g., 'func', '^import').",
					},
					"path": map[string]interface{}{
						"type":        "string",
						"description": "Path to the file where the pattern should be searched.",
					},
				},
				"required": []string{"pattern", "path"},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "tree",
			Description: openai.String("Generate a visual tree-like structure of directories and files from a given path."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "Directory path to start building the tree structure (e.g., '.', './src').",
						"default":     ".",
					},
				},
				"required": []string{"path"},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "list",
			Description: openai.String("List all files and directories in the current working directory."),
			Parameters: openai.FunctionParameters{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "pwd",
			Description: openai.String("Return the current working directory of the environment."),
			Parameters: openai.FunctionParameters{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "delete_file",
			Description: openai.String("Delete a specific file from the file system."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "Path to the file that should be deleted (e.g., './temp.log'). Make sure the file is not needed anymore.",
					},
				},
				"required": []string{"path"},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "read_file",
			Description: openai.String("Read and return the full content of a specified file."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "Path to the file whose contents will be read (e.g., './README.md').",
					},
				},
				"required": []string{"path"},
			},
		},
	},
	{
		Function: openai.FunctionDefinitionParam{
			Name:        "write_file",
			Description: openai.String("Create or overwrite a file with the given content at the specified path."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "The file path where the content will be written (e.g., './output.txt').",
					},
					"content": map[string]interface{}{
						"type":        "string",
						"description": "The full string content to write into the file.",
					},
				},
				"required": []string{"path", "content"},
			},
		},
	},
}

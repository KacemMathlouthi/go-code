package utils

import (
	"fmt"

	"github.com/KacemMathlouthi/go-code/tools"
)

func ExecuteTool(toolName string, toolArgs map[string]string) (string, error) {
	switch toolName {
	case "shell":
		result, err := tools.Shell(toolArgs["command"])
		if err != nil {
			return "", fmt.Errorf("Error executing shell command: %v", err)
		}
		return result, nil

	case "grep":
		result, err := tools.Grep(toolArgs["pattern"], toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("Error executing grep command: %v", err)
		}
		return result, nil

	case "tree":
		result, err := tools.Tree(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("Error executing tree command: %v", err)
		}
		return result, nil

	case "list":
		result, err := tools.List(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("Error executing list command: %v", err)
		}
		return result, nil

	case "pwd":
		result, err := tools.Pwd()
		if err != nil {
			return "", fmt.Errorf("Error executing pwd command: %v", err)
		}
		return result, nil

	case "delete_file":
		err := tools.DeleteFile(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("File at the path: %v is not found or can't be deleted, got error: %v", toolArgs["path"], err)
		}
		return fmt.Sprintf("File at the path %v was succesfully deleted", toolArgs["path"]), nil

	case "read_file":
		result, err := tools.ReadFile(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("Error reading file: %v", err)
		}
		return result, nil

	case "write_file":
		err := tools.WriteFile(toolArgs["path"], toolArgs["content"])
		if err != nil {
			return "", fmt.Errorf("Error writing file: %v", err)
		}
		return fmt.Sprintf("File at the path %v was succesfully written", toolArgs["path"]), nil

	default:
		return "", fmt.Errorf("Tool %v not found", toolName)
	}
}

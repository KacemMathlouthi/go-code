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
			return "", fmt.Errorf("error executing shell command: %v", err)
		}
		return result, nil

	case "grep":
		result, err := tools.Grep(toolArgs["pattern"], toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("error executing grep command: %v", err)
		}
		return result, nil

	case "tree":
		result, err := tools.Tree(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("error executing tree command: %v", err)
		}
		return result, nil

	case "list":
		result, err := tools.List()
		if err != nil {
			return "", fmt.Errorf("error executing list command: %v", err)
		}
		return result, nil

	case "pwd":
		result, err := tools.Pwd()
		if err != nil {
			return "", fmt.Errorf("error executing pwd command: %v", err)
		}
		return result, nil

	case "delete_file":
		err := tools.DeleteFile(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("file at the path: %v is not found or can't be deleted, got error: %v", toolArgs["path"], err)
		}
		return fmt.Sprintf("File at the path %v was succesfully deleted", toolArgs["path"]), nil

	case "read_file":
		result, err := tools.ReadFile(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("error reading file: %v", err)
		}
		return result, nil

	case "write_file":
		err := tools.WriteFile(toolArgs["path"], toolArgs["content"])
		if err != nil {
			return "", fmt.Errorf("error writing file: %v", err)
		}
		return fmt.Sprintf("File at the path %v was succesfully written", toolArgs["path"]), nil

	case "mkdir":
		err := tools.Mkdir(toolArgs["path"])
		if err != nil {
			return "", fmt.Errorf("error creating directory: %v", err)
		}
		return fmt.Sprintf("Directory at the path %v was successfully created", toolArgs["path"]), nil

	default:
		return "", fmt.Errorf("tool %v not found", toolName)
	}
}

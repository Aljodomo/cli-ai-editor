package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type ChatGptRequestProcessor struct {
}

// ProcessRequest generates changes based on user input
func (dp *ChatGptRequestProcessor) ProcessRequest(request string) ([]FileChange, error) {
	fmt.Println("Okay. Let me think about it.")

	// Clear previous changes
	var changes []FileChange

	// get current directory where this programm was started
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not get current dir:", err)
		return nil, err
	}

	// get all files in current directory
	files, err := GetFileNamesInDirectory(dir)
	if err != nil {
		fmt.Println("Could not get files in directory:", err)
		return nil, err
	}

	systemPrompt := "You are a editing software. There are those files in the current directory: \n"
	for _, file := range files {
		systemPrompt += file + "\n"
	}
	systemPrompt += "You can create or delete files.\n"
	systemPrompt += "Only return the operation and the file name per line and NOTHING else.\n"
	systemPrompt += "Example: \n\nCREATE new-file.txt\nDELETE old-file.txt"

	response := AskChatGPT(systemPrompt, request)

	// Parse the response
	changes, err = parseInfoFileChanges(response, files)

	if err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}

	return changes, nil
}

func parseInfoFileChanges(response string, files []string) ([]FileChange, error) {

	var changes []FileChange

	// Split the response into lines
	lines := strings.Split(response, "\n")

	for _, line := range lines {
		// Trim any leading or trailing spaces
		line = strings.TrimSpace(line)

		// Split the line into words
		words := strings.Split(line, " ")

		// If there are not enough words, skip this line
		if len(words) < 2 {
			continue
		}

		// Get the operation and file name
		operation := words[0]
		fileName := words[1]

		// Check if the operation is valid
		if operation != "CREATE" && operation != "DELETE" {
			return nil, fmt.Errorf("invalid operation: %s", operation)
		}

		// Check if the file already exists
		if operation == "CREATE" && slices.Contains(files, fileName) {
			return nil, fmt.Errorf("file already exists: %s", fileName)
		}

		// Check if the file does not exist
		if operation == "DELETE" && !slices.Contains(files, fileName) {
			return nil, fmt.Errorf("file does not exist: %s", fileName)
		}

		// Add the file change
		changes = append(changes, FileChange{Operation: Operation(operation), FilePath: fileName})
	}

	return changes, nil
}

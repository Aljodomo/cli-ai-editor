package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
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
	systemPrompt += "Only return a json list of objects with the format {operation: 'CREATE' | 'DELETE' | 'EDIT', relative_file_path: string, file_content: string} and NOTHING else.\n"
	systemPrompt += "file_path can never start with a / because they represent relative paths.\n"
	systemPrompt += "Example:\n"
	systemPrompt += "[{\"operation\": \"CREATE\", \"file_path\": \"docs/new-file.txt\", \"file_content\": \"This is a new file\"}]"

	response := AskChatGPT(systemPrompt, request)
	fmt.Println("Response from ChatGPT:", response)

	// Parse the response
	changes, err = parseFileChanges(response, files)

	if err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}

	return changes, nil
}

func parseFileChanges(response string, files []string) ([]FileChange, error) {

	var changes []FileChange

	err := json.Unmarshal([]byte(response), &changes)
	if err != nil {
		return nil, err
	}

	// check if all file paths are valid
	for _, change := range changes {
		switch change.Operation {
		case EDIT:
		case DELETE:
			if !slices.Contains(files, change.FilePath) {
				return nil, fmt.Errorf("file does not exist: %s", change.FilePath)
			}
		case CREATE:
			if slices.Contains(files, change.FilePath) {
				return nil, fmt.Errorf("file already exists: %s", change.FilePath)
			}
		}
	}

	return changes, nil
}

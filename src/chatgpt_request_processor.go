package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

//go:embed system-prompt.md
var systemPrompt string

type assistantResponse struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type assistantReadPayload = []string
type assistantWritePayload = []FileChange

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

	response := AskChatGPT(systemPrompt, request)

	// Parse the response
	changes, err = parseFileChanges(response, files)

	if err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}

	return changes, nil
}

func parseAssistantResponse(response string) (*assistantResponse, error) {
	var assistantResponse assistantResponse

	err := json.Unmarshal([]byte(response), &assistantResponse)
	if err != nil {
		return nil, err
	}

	return &assistantResponse, nil
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

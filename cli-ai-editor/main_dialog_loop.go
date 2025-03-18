package main

import (
	"fmt"
	"strings"
)

type Operation string

const (
	EDIT   Operation = "Edit"
	CREATE           = "Create"
	DELETE           = "Delete"
)

// FileChange represents a file operation with path
type FileChange struct {
	Operation   Operation
	FilePath    string
	FileContent string
}

type RequestProcessor interface {
	ProcessRequest(request string) error
	GetFileChanges() []FileChange
}

type FileChangeExecuter interface {
	ExecuteFileChanges(changes []FileChange) error
}

// DisplayProposedChanges shows all proposed changes
func displayProposedChanges(fileChanges []FileChange) {
	fmt.Println("\nI want make those changes: ")

	for _, change := range fileChanges {
		fmt.Printf("%s - %s\n", change.Operation, change.FilePath)
	}

	fmt.Println()
}

// RunMainDialogLoop executes the main interaction loop
func RunMainDialogLoop(processor RequestProcessor, executer FileChangeExecuter) {

	// Initial greeting
	fmt.Println("What can I do for you?")

	// Get initial request
	request := GetUserInput("")

	// Main interaction loop
	for {
		// Process request and propose changes
		err := processor.ProcessRequest(request)

		if err != nil {
			fmt.Println("Error processing request:", err)
			break
		}

		// Display proposed changes
		displayProposedChanges(processor.GetFileChanges())

		// Get confirmation
		confirmation := RequestUserConfirmation("Is this okay?")

		if confirmation == true {

			err = executer.ExecuteFileChanges(processor.GetFileChanges())

			if err != nil {
				fmt.Println("Error executing changes:", err)
				break
			}

			request := GetUserInput("Okay all done. Please check git and tell me if you need something changed or type 'thanks' to exit")

			if strings.ToLower(request) == "thanks" {
				fmt.Println("It was a pleasure. Goodbye!")
				break
			}
		} else {
			// User rejected changes
			fmt.Println("Changes rejected. Let's try again.")
			request = GetUserInput("What would you like me to do instead? ")
		}
	}
}

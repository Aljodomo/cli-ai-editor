package main

import (
	"fmt"
	"os"
	"slices"
)

// RequestProcessor handles the request processing and change management
type TestRequestProcessor struct {
}

// ProcessRequest generates changes based on user input
func (dp *TestRequestProcessor) ProcessRequest(request string) ([]FileChange, error) {
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

	mySpecialFile := "docs/my-special-file.txt"
	if slices.Contains(files, mySpecialFile) {
		changes = append(changes, FileChange{Operation: DELETE, FilePath: mySpecialFile})
	} else {
		changes = append(changes, FileChange{Operation: CREATE, FilePath: mySpecialFile})
	}

	return changes, nil
}

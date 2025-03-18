package main

import (
	"fmt"
	"os"
	"slices"
)

// RequestProcessor handles the request processing and change management
type TestRequestProcessor struct {
	changes []FileChange
}

// ProcessRequest generates changes based on user input
func (dp *TestRequestProcessor) ProcessRequest(request string) error {
	fmt.Println("Okay. Let me think about it.")

	// Clear previous changes
	dp.changes = []FileChange{}

	// get current directory where this programm was started
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not get current dir:", err)
		return err
	}

	// get all files in current directory
	files, err := GetFileNamesInDirectory(dir)
	if err != nil {
		fmt.Println("Could not get files in directory:", err)
		return err
	}

	mySpecialFile := "my-special-file.txt"
	if slices.Contains(files, mySpecialFile) {
		dp.changes = append(dp.changes, FileChange{Operation: DELETE, FilePath: mySpecialFile, FileContent: "This is a special file"})
	} else {
		dp.changes = append(dp.changes, FileChange{Operation: CREATE, FilePath: mySpecialFile})
	}

	return nil
}

// GetFileChanges returns the proposed changes
func (dp *TestRequestProcessor) GetFileChanges() []FileChange {
	return dp.changes
}

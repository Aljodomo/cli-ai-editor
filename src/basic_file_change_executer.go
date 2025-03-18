package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type BasicFileChangeExecuter struct {
}

func (b *BasicFileChangeExecuter) ExecuteFileChanges(changes []FileChange) error {

	for _, change := range changes {
		switch change.Operation {
		case EDIT:
			err := WriteFileContent(change.FilePath, change.FileContent)
			if err != nil {
				return err
			}
			fmt.Println("Edited file:", change.FilePath)
		case CREATE:
			dir := filepath.Dir(change.FilePath)
			if dir != "." {
				err := os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return err
				}
			}
			err := WriteFileContent(change.FilePath, change.FileContent)
			if err != nil {
				return err
			}
			fmt.Println("Created file:", change.FilePath)
		case DELETE:
			err := DeleteFile(change.FilePath)
			if err != nil {
				return err
			}
			fmt.Println("Deleted file:", change.FilePath)
		default:
			return fmt.Errorf("unknown operation: %s", change.Operation)
		}
	}
	return nil
}

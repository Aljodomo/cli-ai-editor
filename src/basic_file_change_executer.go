package main

import (
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
		case DELETE:
			err := DeleteFile(change.FilePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

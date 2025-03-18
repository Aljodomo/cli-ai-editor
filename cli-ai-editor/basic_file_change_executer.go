package main

type BasicFileChangeExecuter struct {
}

func (b *BasicFileChangeExecuter) ExecuteFileChanges(changes []FileChange) error {

	for _, change := range changes {
		switch change.Operation {
		case EDIT:
		case CREATE:
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

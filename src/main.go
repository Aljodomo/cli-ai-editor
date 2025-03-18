package main

func main() {
	RunMainDialogLoop(&ChatGptRequestProcessor{}, &BasicFileChangeExecuter{})
}

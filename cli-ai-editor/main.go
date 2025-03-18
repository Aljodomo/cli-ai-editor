package main

func main() {
	RunMainDialogLoop(&TestRequestProcessor{}, &BasicFileChangeExecuter{})
}

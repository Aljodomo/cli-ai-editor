package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RequestUserConfirmation(message string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", message)

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			return false
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}

		fmt.Println("Please respond with 'y' or 'n'")
	}
}

// getUserInput reads input from the terminal with optional prompt
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	if prompt != "" {
		fmt.Println(prompt)
	}
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

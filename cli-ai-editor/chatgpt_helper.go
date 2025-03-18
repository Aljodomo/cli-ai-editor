package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func AskChatGPT(systemPrompt string, userPrompt string) string {
	openAiApiKeyEnv := "OPENAI_API_KEY"
	openAiApiKey, ok := os.LookupEnv(openAiApiKeyEnv)
	if ok == false {
		log.Fatalf(fmt.Sprintf("Env %s is not set", openAiApiKeyEnv))
	}
	slog.Debug("OpenAi api key found")
	slog.Debug("System prompt: %s", systemPrompt)
	slog.Debug("User prompt: %s", userPrompt)
	openAiRequestBody := createOpenAiRequestBody(systemPrompt, userPrompt)

	requestBodyReader := toByteReader(openAiRequestBody)

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", requestBodyReader)

	if err != nil {
		log.Fatalf("Could not create request: %s", err)
	}

	client := &http.Client{}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAiApiKey))

	resp, err := client.Do(request)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Could not close body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Could not read resonse body: %s", err)
	}

	openAiResponseBody := OpenAiResponseBody{}
	err = json.Unmarshal(body, &openAiResponseBody)
	if err != nil {
		log.Fatalf("Could not unmarshal response body: %s", err)
	}

	responseMessage := openAiResponseBody.Choices[0].Message.Content
	return responseMessage
}

func toByteReader(it any) *bytes.Reader {
	marshal, err := json.Marshal(it)

	if err != nil {
		log.Fatalf("Could not marshall object: %s", err)
	}

	return bytes.NewReader(marshal)
}

func createOpenAiRequestBody(systemPrompt string, userPrompt string) OpenAiRequestBody {
	return OpenAiRequestBody{
		Model:    "gpt-4o-mini",
		Messages: []OpenAiRequestBodyMessage{{Role: "system", Content: systemPrompt}, {Role: "user", Content: userPrompt}},
	}
}

type OpenAiRequestBodyMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAiRequestBody struct {
	Model    string                     `json:"model"`
	Messages []OpenAiRequestBodyMessage `json:"messages"`
}

type OpenAiResponseBody struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
			Refusal any    `json:"refusal"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

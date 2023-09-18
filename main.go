package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

const (
	apiKey = "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"
	apiURL = "https://api.openai.com/v1/chat/completions"
)

func extractCodeBlocks(input string) []string {
	codeBlockRegex := regexp.MustCompile("```go[\\s\\S]*?```")
	matches := codeBlockRegex.FindAllString(input, -1)
	return matches
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Message      Message `json:"message"`
}
type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
type Response struct {
	Choices []Choice `json:"choices"`
	Created int      `json:"created"`
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Object  string   `json:"object"`
	Usage   Usage    `json:"usage"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go SOURCE_FILE_PATH OUTPUT_FILE_PATH")
		return
	}

	sourceFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	jsonData, err := os.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Printf("Error reading the JSON file: %v\n", err)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	payload := data

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	if err := os.WriteFile(outputFilePath, responseJSON, 0644); err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		return
	}

	fmt.Println("success: Response JSON file successfully saved.")

	codeBlocks := extractCodeBlocks(response.Choices[0].Message.Content)
	for idx, codeBlock := range codeBlocks {
		fmt.Printf("Response:%v\n%v\n", idx+1, codeBlock[6:len(codeBlock)-3])
		goSource := codeBlock[6 : len(codeBlock)-3]
		goSourceFilePath := fmt.Sprintf("%v.%v.go", outputFilePath, idx+1)
		if err := os.WriteFile(goSourceFilePath, []byte(goSource), 0644); err != nil {
			fmt.Printf("Error writing Go source file: %v\n", err)
			return
		}
	}

	fmt.Println("success: Go source files successfully saved.")
}

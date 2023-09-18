package gpt

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
	if len(matches) == 0 {
		fmt.Println(input)
		return []string{input}
	}

	var ans []string
	for _, codeBlock := range matches {
		ans = append(ans, codeBlock[6:len(codeBlock)-3])
		fmt.Println(codeBlock[6 : len(codeBlock)-3])
	}

	return ans
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

func GPT(sourceFilePath string, outputFilePath string) ([]string, error) {
	jsonData, err := os.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Printf("Error reading the JSON file: %v\n", err)
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return nil, err
	}

	payload := data

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return nil, err
	}

	if err := os.WriteFile(outputFilePath, responseJSON, 0644); err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		return nil, err
	}

	fmt.Println("success: Response JSON file successfully saved.")

	codeBlocks := extractCodeBlocks(response.Choices[0].Message.Content)

	return codeBlocks, nil
	// for idx, codeBlock := range codeBlocks {
	// 	fmt.Printf("Response:%v\n%v\n", idx+1, codeBlock[6:len(codeBlock)-3])
	// 	goSource := codeBlock[6 : len(codeBlock)-3]
	// 	goSourceFilePath := fmt.Sprintf("%v.%v.go", outputFilePath, idx+1)
	// 	if err := os.WriteFile(goSourceFilePath, []byte(goSource), 0644); err != nil {
	// 		fmt.Printf("Error writing Go source file: %v\n", err)
	// 		return
	// 	}
	// }

	// fmt.Println("success: Go source files successfully saved.")
}

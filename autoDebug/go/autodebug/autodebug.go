package autodebug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/gpt"
	"os"
	"os/exec"
)

type Query struct {
	Model            string        `json:"model"`
	Messages         []gpt.Message `json:"messages"`
	Temperature      float64       `json:"temperature"`
	MaxTokens        int           `json:"max_tokens"`
	TopP             float64       `json:"top_p"`
	FrequencyPenalty float64       `json:"frequency_penalty"`
	PresencePenalty  float64       `json:"presence_penalty"`
}

func AutoDebug(logDirPath string, queryPath string) (string, error) {
	code, err := initTrial(logDirPath, queryPath)
	if err == nil {
		return code, nil
	}
	for i := 0; i < 3; i++ {
		code, err = tryAgain(logDirPath, queryPath, code, err, i+1)
		if err == nil {
			return code, nil
		}
		if code == "" {
			return "", fmt.Errorf("failed to generate code: %v", err)
		}
	}
	return code, nil
}

func initTrial(logDirPath string, queryPath string) (string, error) {
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return "", err
	}
	query := Query{
		Model: "gpt-3.5-turbo",
		Messages: []gpt.Message{
			{
				Role:    "system",
				Content: "You are an AI assistant and always write output of your response in Golang. Generate code that solve given problems.",
			},
			{
				Role:    "user",
				Content: string(queryBytes),
			},
		},
		Temperature:      0.5,
		MaxTokens:        1000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	jsonBytes, err := json.Marshal(query)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return "", err
	}

	if err := os.WriteFile(logDirPath+"/0_query.json", jsonBytes, 0644); err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		return "", err
	}

	codes, err := gpt.GPT(logDirPath+"/0_query.json", logDirPath+"/0_response.json")
	if err != nil {
		return "", err
	}
	err = executeGoProgram(codes[0])

	return codes[0], err
}

func tryAgain(logDirPath string, queryPath string, code string, err_ error, idx int) (string, error) {
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return "", err
	}
	query := Query{
		Model: "gpt-3.5-turbo",
		Messages: []gpt.Message{
			{
				Role:    "system",
				Content: "You are an AI assistant and always write output of your response in Golang. Generate code that solve given problems considering error logs.",
			},
			{
				Role:    "user",
				Content: string(queryBytes),
			},
			{
				Role:    "assistant",
				Content: code,
			},
			{
				Role:    "user",
				Content: "Given code results in error! Fix it.\n" + err_.Error(),
			},
		},
		Temperature:      0.5,
		MaxTokens:        1000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	jsonBytes, err := json.Marshal(query)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return "", err
	}

	if err := os.WriteFile(fmt.Sprintf("%v/%v_query.json", logDirPath, idx), jsonBytes, 0644); err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		return "", err
	}

	codes, err := gpt.GPT(fmt.Sprintf("%v/%v_query.json", logDirPath, idx), fmt.Sprintf("%v/%v_response.json", logDirPath, idx))
	if err != nil {
		return "", err
	}
	err = executeGoProgram(codes[0])

	return codes[0], err
}

func executeGoProgram(src string) error {
	tmpFile := "temp.go"
	err := createTempFile(tmpFile, src)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	cmd := exec.Command("go", "run", tmpFile)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("execution error: %v\nStderr: %s", err, stderr.String())
	}

	return nil
}

func createTempFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// func main() {
// 	src :=
// 		`package main
// 		import "fmt"
// 		func main() {
// 			fmt.Println("hello world!")
// 		}`

// 	err := executeGoProgram(src)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

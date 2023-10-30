package v3

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var (
	messages      []openai.ChatCompletionMessage
	newMessages   []openai.ChatCompletionMessage
	readFiles     []string
	ignoreFiles   []string
	readFilesAll  []string
	selectedFiles []string
)

func makeMessagesListedFiles() error {
	err := addFiles()
	if err != nil {
		return err
	}
	return nil
}

func makeMessagesToBrainStorm() error {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: "List the additional features to be satisfied direction. ",
	})

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: "Generate task list with following style: `1. <package1>\n- ...\n- ...\n 2. <package2>\n- ...\n- ... \n...` ",
	})
	return nil
}

func makeMessagesToSelectFile(direction string, filePath string) ([]openai.ChatCompletionMessage, error) {
	newMessages := append(messages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: fmt.Sprintf("direction: %s", direction),
	})

	addFileContent(filePath, &newMessages)

	newMessages = append(newMessages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: fmt.Sprintf("Will %s need to be modified or to be referred when generating code? (y/n)", filePath),
	})

	return newMessages, nil
}

func makeMessagesToGenerateCodeWithTask(task string) error {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: "Based on the above directory structure, file contents, generate a whole code to implement the features bellow. Do not omits. ",
	})

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: task,
	})

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: "Generate the only file that is edited, with following style: \nfile path\n```\nfile contents\n```\n",
	})

	return nil
}

func makeMessagesToAddDirection(direction string) error {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: fmt.Sprintf("direction: %s", direction),
	})
	return nil
}

func GenAI(direction string) {
	fmt.Println("GenAI v2\ndirection: ", direction)

	var err error

	readFiles, err = makeReadList()
	if err != nil {
		log.Println(err)
		return
	}
	ignoreFiles, err = makeIgnoreList()
	if err != nil {
		log.Println(err)
		return
	}

	if err = makeMessagesListedFiles(); err != nil {
		log.Println(err)
		return
	}

	if err = selectFiles(direction); err != nil {
		log.Println(err)
		return
	}

	readFiles = selectedFiles
	messages = []openai.ChatCompletionMessage{}
	if err = makeMessagesListedFiles(); err != nil {
		log.Println(err)
		return
	}

	err = makeMessagesToAddDirection(direction)
	if err != nil {
		log.Println(err)
		return
	}

	if err = makeMessagesToBrainStorm(); err != nil {
		log.Println(err)
		return
	}
	err = openaiAPI("task list")
	if err != nil {
		log.Println(err)
		return
	}

	taskList := parseTaskList("out1")
	if len(taskList) == 0 {
		log.Println("task list is empty")
		return
	}
	fmt.Println(strings.Join(taskList, "\n========================\n"))

	for idx, task := range taskList {
		if err = makeMessagesToGenerateCodeWithTask(task); err != nil {
			log.Println(err)
			return
		}
		if idx == 0 {
			err = openaiAPI("code init")
		} else {
			err = openaiAPI("code repeat")
		}
		if err != nil {
			log.Println(err)
			return
		}
	}

	// if err = makeMessagesToGenerateCode(); err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = openaiAPI("code init")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	parseFiles("out2")
}

func openaiAPI(outputType string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load .env file")
	}
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	res, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       "gpt-4",
			Messages:    messages,
			Temperature: 0.03,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create chat completion: %w", err)
	}

	log.Println("prompt_tokens: ", res.Usage.PromptTokens, "completion_tokens: ", res.Usage.CompletionTokens, "total_tokens: ", res.Usage.TotalTokens)

	switch outputType {
	case "task list":
		messages = append(messages, res.Choices[0].Message)

		file, err := os.OpenFile("out1", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to open the file for output: %w", err)
		}
		defer file.Close()

		_, err = fmt.Fprint(file, res.Choices[0].Message.Content)
		if err != nil {
			return fmt.Errorf("failed to write to the file for output: %w", err)
		}

		return nil

	case "code init":
		messages = append(messages, res.Choices[0].Message)

		file, err := os.OpenFile("out2", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open the file for output: %w", err)
		}
		defer file.Close()

		_, err = fmt.Fprint(file, res.Choices[0].Message.Content)
		if err != nil {
			return fmt.Errorf("failed to write to the file for output: %w", err)
		}

		return nil

	case "code repeat":
		messages = append(messages, res.Choices[0].Message)

		file, err := os.OpenFile("out2", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open the file for output: %w", err)
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "\n\n%s", res.Choices[0].Message.Content)
		if err != nil {
			return fmt.Errorf("failed to write to the file for output: %w", err)
		}

		return nil

	default:
		return fmt.Errorf("invalid outputType: %s", outputType)
	}
}

func openaiAPIv2(messages []openai.ChatCompletionMessage, outputType string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("failed to load .env file")
	}
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	res, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       "gpt-4",
			Messages:    messages,
			Temperature: 0.03,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	log.Println("prompt_tokens: ", res.Usage.PromptTokens, "completion_tokens: ", res.Usage.CompletionTokens, "total_tokens: ", res.Usage.TotalTokens)

	switch outputType {
	case "y/n":
		return res.Choices[0].Message.Content, nil

	default:
		return "", fmt.Errorf("invalid outputType: %s", outputType)
	}
}

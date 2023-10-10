package genai

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var messages []openai.ChatCompletionMessage
var readFiles []string
var ignoreFiles []string

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
		Content: "List tasks you need to do to complete the code by comparing the definition of api and database with the existing code. ",
	})
	return nil
}

func makeMessagesToGenerateCode() error {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: "Based on the above directory structure, file contents, generate a whole code which can serve both the user and task at the same time. ",
	})

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: "Generate the only file that is edited, with following style: \nfile path\n```\nfile contents\n```\n",
	})

	return nil
}

func GenAI() {
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

	if err = makeMessagesToBrainStorm(); err != nil {
		log.Println(err)
		return
	}

	if err = makeMessagesToGenerateCode(); err != nil {
		log.Println(err)
		return
	}

	err = openaiAPI(false)
	if err != nil {
		log.Println(err)
		return
	}
	parseFiles("out")
}

func openaiAPI(continueConversation bool) error {
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

	if continueConversation {
		messages = append(messages, res.Choices[0].Message)
		return nil
	} else {
		file, err := os.OpenFile("out", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to open the file for output: %w", err)
		}
		defer file.Close()

		_, err = fmt.Fprint(file, res.Choices[0].Message.Content)
		if err != nil {
			return fmt.Errorf("failed to write to the file for output: %w", err)
		}

		return nil
	}
}

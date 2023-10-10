package genai

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/sashabaranov/go-openai"
)

func visitFile(fp string, fi os.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !fi.IsDir() && isInReadFiles(fp) && !isInIgnoreFiles(fp) {
		err = addFileContent(fp)
		if err != nil {
			return err
		}
	}
	return nil
}

func addFiles() error {
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: "This is the content of each file.",
	})
	err := filepath.WalkDir("..", visitFile)

	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}
	return nil
}

func addFileContent(filePath string) error {
	fmt.Println(filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: filePath + "\n```\n" + string(file) + "\n```\n",
	})

	return nil
}

func parseTaskList(outputFilePath string) []string {
	file, err := os.ReadFile(outputFilePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	input := string(file)

	// 正規表現パターン
	pattern := "([0-9]+)\\.\\s(.*)((\\n(\\s+)\\-\\s(.*))+)"

	// 正規表現をコンパイル
	re := regexp.MustCompile(pattern)

	// マッチした部分を格納するスライス
	matches := re.FindAllStringSubmatch(input, -1)

	// マッチした内容を格納する配列
	var results []string

	// マッチした部分を処理
	for _, match := range matches {
		result := match[0]
		results = append(results, result)
	}

	fmt.Println(len(results))

	return results
}

func parseFiles(outputFilePath string) {
	file, err := os.ReadFile(outputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(file)

	// 正規表現パターン
	pattern := "\\n*(.*?):?\\n```.*\\n([\\s\\S]*?)\\n```"

	// 正規表現をコンパイル
	re := regexp.MustCompile(pattern)

	// マッチした部分を格納するスライス
	matches := re.FindAllStringSubmatch(input, -1)

	// マッチした内容を格納する配列
	var results []struct {
		FileName string
		Contents string
	}

	// マッチした部分を処理
	for _, match := range matches {
		if len(match) == 3 {
			fileName := match[1]
			contents := match[2]
			result := struct {
				FileName string
				Contents string
			}{fileName, contents}
			results = append(results, result)
		}
	}

	// 結果を出力
	for _, result := range results {
		fmt.Println(result.FileName)
		file, err := os.OpenFile(result.FileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("failed to open the file for output: " + err.Error())
			return
		}
		defer file.Close()

		_, err = fmt.Fprint(file, result.Contents)
		if err != nil {
			fmt.Println("failed to write to the file for output: " + err.Error())
			return
		}
	}
}

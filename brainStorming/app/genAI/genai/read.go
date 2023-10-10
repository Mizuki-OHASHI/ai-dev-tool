package genai

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isInReadFiles(filePath string) bool {
	for _, ignoreFile := range readFiles {
		matched, err := regexp.MatchString(ignoreFile, filePath)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if matched {
			return true
		}
	}
	return false
}

func isInIgnoreFiles(filePath string) bool {
	for _, ignoreFile := range ignoreFiles {
		matched, err := regexp.MatchString(ignoreFile, filePath)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if matched {
			return true
		}
	}
	return false
}

func makeReadList() ([]string, error) {
	file, err := os.Open(".read")
	if err != nil {
		return nil, fmt.Errorf("failed to open .read file: %w", err)
	}
	defer file.Close()

	var paths []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			escapedPath := escapeRegexMetaChars(line)
			pathPattern := ".*" + escapedPath + ".*"
			paths = append(paths, pathPattern)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan .read file: %w", err)
	}

	return paths, nil
}

func makeIgnoreList() ([]string, error) {
	file, err := os.Open(".readignore")
	if err != nil {
		return nil, fmt.Errorf("failed to open .readignore file: %w", err)
	}
	defer file.Close()

	var paths []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			escapedPath := escapeRegexMetaChars(line)
			pathPattern := ".*" + escapedPath + ".*"
			paths = append(paths, pathPattern)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan .readignore file: %w", err)
	}

	return paths, nil
}

func escapeRegexMetaChars(pattern string) string {
	metaChars := []string{"\\", ".", "^", "$", "*", "+", "?", "{", "}", "[", "]", "(", ")", "|"}
	for _, c := range metaChars {
		pattern = regexp.MustCompile(`\`+c).ReplaceAllString(pattern, `\`+c)
	}
	return pattern
}

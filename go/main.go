package main

import (
	"fmt"
	"main/autodebug"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go LOG_DIR_PATH QUERY_PATH")
		os.Exit(1)
	}

	logDirPath := os.Args[1]
	queryPath := os.Args[2]
	autodebug.AutoDebug(logDirPath, queryPath)
}

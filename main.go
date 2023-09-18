package main

import (
	"fmt"
	"main/autodebug"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go LOG_DIR_PATH QUERY_PATH")
		os.Exit(1)
	}

	logDirPath := os.Args[1]
	queryPath := os.Args[2]
	autodebug.AutoDebug(logDirPath, queryPath)
}

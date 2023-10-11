package main

import (
	"fmt"
	genAIv1 "genai/genai/v1"
	genAIv2 "genai/genai/v2"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		genAIv1.GenAI()
	} else {
		switch os.Args[1] {
		case "direction":
			var direction string
			fmt.Scan(&direction)
			genAIv2.GenAI(direction)
			return
		default:
			log.Fatalf("invalid argument: %v", os.Args[1:])
		}

	}
}

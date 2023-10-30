package main

import (
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
		case "--direction":
			if len(os.Args) != 3 {
				log.Fatalf("invalid argument: %v", os.Args[1:])
			}
			genAIv2.GenAI(os.Args[2])
			return
		default:
			log.Fatalf("invalid argument: %v", os.Args[1:])
		}
	}
}

package v2

import (
	"log"
	"time"
)

func selectFiles(direction string) error {
	var err error
	for _, filepath := range readFilesAll {
		newMessages, err = makeMessagesToSelectFile(direction, filepath)
		if err != nil {
			return err
		}
		answer, err := openaiAPIv2(newMessages, "y/n")
		if err != nil {
			return err
		}
		selected, err := parseYesNo(answer)
		if err != nil {
			return err
		}
		if selected {
			selectedFiles = append(selectedFiles, filepath)
			log.Printf("selected: %s", filepath)
		} else {
			log.Printf("not selected: %s", filepath)
		}
		time.Sleep(time.Second * 10)
	}
	return nil
}

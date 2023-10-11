package v2

import "log"

func selectFiles(direction string) error {
	var err error
	for _, filepath := range readFiles {
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
		}
	}
	return nil
}

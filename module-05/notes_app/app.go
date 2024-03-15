package main

import (
	"bufio"
	"com/notes_app/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	note1, err := note.New(title, content)

	if err != nil {
		fmt.Println("ERROR! ", err)
		return
	}

	note1.Display()

	err = note1.Save()

	if err != nil {
		fmt.Println("ERROR! Saving Note ", err)
		return
	}

	fmt.Println("Note saved successfully!")

}

func getNoteData() (string, string) {
	title := getUserInput("Please enter the title:")
	content := getUserInput("Please enter the content:")
	return title, content
}

func getUserInput(promptText string) string {
	fmt.Printf("%v ", promptText)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("ERROR! ", err)
		return ""
	}

	text := strings.TrimSuffix(value, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

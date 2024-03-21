package main

import (
	"bufio"
	"com/notes_and_todos_app/note"
	"com/notes_and_todos_app/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("ERROR! ", err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	note1, err := note.New(title, content)

	if err != nil {
		fmt.Println("ERROR! ", err)
		return
	}

	note1.Display()
	err = saveData(note1)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("ERROR! Saving Data ", err)
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
}

func getTodoData() string {
	return getUserInput("Please enter the text:")
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

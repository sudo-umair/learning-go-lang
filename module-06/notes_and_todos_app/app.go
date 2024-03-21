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

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func main() {

	printSomething("Welcome to the Notes and Todos App!")
	printSomething(2)
	printSomething(2.5)
	printSomething(true)

	todoText := getTodoData()
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("ERROR! ", err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	title, content := getNoteData()
	note1, err := note.New(title, content)

	if err != nil {
		fmt.Println("ERROR! ", err)
		return
	}

	err = outputData(note1)

	if err != nil {
		return
	}

}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	default:
		fmt.Println("Unknown type: ", value)
	}
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
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

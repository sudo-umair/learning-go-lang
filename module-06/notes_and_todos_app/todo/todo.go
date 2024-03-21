package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {

	if text == "" {
		fmt.Println("Invalid input")
		return nil, errors.New("invalid todo details")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (n *Todo) Display() {
	fmt.Printf("Your todo: \"%v\"\n", n.Text)
}

func (n *Todo) Save() error {
	fileName := "my_todo.json"
	json, err := json.Marshal(&n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLinesFromFile(fileName string) ([]string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.New("error opening file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("error reading file")
	}

	file.Close()
	return lines, nil
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("error creating file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("error writing json")
	}

	file.Close()
	return nil
}

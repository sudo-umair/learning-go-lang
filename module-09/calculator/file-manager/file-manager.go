package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.OpenFile(fm.InputFilePath, os.O_RDONLY, 0644)
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

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
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

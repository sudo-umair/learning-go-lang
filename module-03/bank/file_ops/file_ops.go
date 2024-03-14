package file_ops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	balanceText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(balanceText), 0644)

	// https://www.redhat.com/sysadmin/linux-file-permissions-explained
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse value")
	}

	// https://golang.org/pkg/strconv/#ParseFloat
	return balance, nil
}

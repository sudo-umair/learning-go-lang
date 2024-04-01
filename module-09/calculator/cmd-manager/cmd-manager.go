package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices. Confirm every price with Enter")

	var prices []string
	for {
		fmt.Print("Price: ")
		var price string
		fmt.Scanln(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

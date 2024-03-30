package main

import (
	"fmt"
)

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}

	fmt.Println("Websites:", websites)
	fmt.Println("Facebook:", websites["Facebook"])

	websites["Twitter"] = "https://www.twitter.com"
	fmt.Println("Websites:", websites)

	delete(websites, "Facebook")
	fmt.Println("Websites:", websites)
}

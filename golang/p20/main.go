package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello, World!")
	err := os.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error on writting to file")
		os.Exit(1)
	}
}

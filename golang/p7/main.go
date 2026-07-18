package main

import "fmt"

func save() error {
	return nil
}

func main() {
	if err := save(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Saved successfully")
}

package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter number:")
	fmt.Scan(&i)
	fmt.Printf("Your number is: %d\n", i)
	for j := 0; j < i; j++ {
		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	}
}

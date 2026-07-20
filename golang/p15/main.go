package main

import "fmt"

var number = 20

func Changenumber(number *int) int {
	*number = 19
	return *number
}

func changeNumber(n int) {
	n = 19
}
func main() {
	fmt.Println(Changenumber(&number))
	number := 20
	changeNumber(number)
	fmt.Println(number)
}

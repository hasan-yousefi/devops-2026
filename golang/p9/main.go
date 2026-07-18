package main

import "fmt"

func sum(a int, b int) int {
	c := a + b
	return c
}

func IsEven(s int) bool {
	if s%2 == 0 {
		return true
	} else {
		return false
	}

}
func main() {
	s := sum(6, 6)
	fmt.Println(IsEven(s))
}

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func ChangUser(user *User) {
	user.Name = "john"
}

func main() {
	user := User{Name: "bob", Age: 20}
	ChangUser(&user)
	fmt.Println(user.Name)
}

package main

import "fmt"

type Animal interface {
	speak(word string)
}

type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (d Dog) speak(word string) {
	fmt.Printf("%s says: %s\n", d.name, word)
}

func (c Cat) speak(word string) {
	fmt.Printf("%s says: %s\n", c.name, word)
}

func Makespeak(animal Animal, word string) {
	animal.speak(word)
}

func main() {
	dog := Dog{name: "Buddy"}
	cat := Cat{name: "Whiskers"}
	Makespeak(dog, "Woof!")
	Makespeak(cat, "Meow!")
}

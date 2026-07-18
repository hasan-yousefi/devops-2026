package main

import "fmt"

type flyer interface {
    fly()
}

type Bird struct {
    Name string
}

type Airplane struct {
    Model string
}

func (b Bird) fly() {
    fmt.Printf("bird %s can fly", b.Name)
}

func (p Airplane) fly() {
    fmt.Printf("airplane %s will be fly", p.Model)
}

func makefly(f flyer) {
    f.fly()
}

func main() {
    bird := Bird{Name: "birdname"}
    plane := Airplane{Model: "boing"}
    makefly(bird)
    makefly(plane)
}
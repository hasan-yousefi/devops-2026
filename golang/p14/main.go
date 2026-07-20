package main

import "fmt"

type Todo struct {
	Title string
	Done  bool
}

func (t Todo) IsFinished() bool {
	return t.Done
}

func (t *Todo) Complete() bool {
	t.Done = true
	return t.Done
}

func main() {
	task1 := Todo{Title: "Learn Go", Done: false}
	fmt.Println("Task 1:", task1.Title, "Done:", task1.IsFinished())
	task1.Complete()
	fmt.Println("Task 1:", task1.Title, "Done:", task1.IsFinished())

	tasks := map[string]*Todo{
		"task2": {Title: "Write Code", Done: false},
		"task3": {Title: "Test Code", Done: true},
	}
	for _, status := range tasks {
		if !status.IsFinished() {
			fmt.Printf("title: %s is not finished yet\n", status.Title)
		}
	}
}

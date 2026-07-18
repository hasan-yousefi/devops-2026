package main

import "fmt"

func main() {
	counts := make(map[string]int)
	information := make(map[string]string)
	biggest := 0
	ages := []int{12, 20, 65, 50, 25}
	names := []string{"bob", "rad", "john", "robert", "bob", "garry", "john"}

	for _, name := range names {
		counts[name]++
	}

	fmt.Println("Counts:")
	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}

	seen := make(map[string]bool)
	ageIndex := 0
	for _, name := range names {
		if seen[name] {
			continue
		}
		seen[name] = true

		if ageIndex < len(ages) {
			information[name] = fmt.Sprintf("%s is %d years old", name, ages[ageIndex])
			ageIndex++
		} else {
			information[name] = fmt.Sprintf("%s has no age assigned", name)
		}
	}

	fmt.Println("Information:")
	for name, info := range information {
		fmt.Printf("%s: %s\n", name, info)
	}

	for _, age := range ages {
		if age > biggest {
			biggest = age
		}
	}
	fmt.Printf("The biggest age is: %d\n", biggest)
}

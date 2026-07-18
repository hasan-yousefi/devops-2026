package main
import "fmt"



func main() {
  // array (like tuple in python)
  age := [...]int{12, 4, 5}

  for i := 0; i < len(age); i++ {
    fmt.Println(age[i])
  }
  // slice doesn't have a fixed size (like list in python)
  Numbers := []int{2, 3}  
  Numbers = append(Numbers, 5, 7)

  fmt.Println("Prime Numbers:", Numbers)

  primeNumbers := []int{2, 3, 5, 7}
  numbers := []int{1, 2, 3}
  // copy elements of primeNumbers to numbers
  copy(numbers, primeNumbers)
  fmt.Println("Numbers:", numbers)
  for index, value := range numbers {
    fmt.Println(index, value)
  }
  // create a map (like dictionary in Python)
  Marks := map[string]float32{"ali": 85, "javad": 80, "ahmad": 81}
  Marks["hsn"] = 98
  fmt.Println(Marks)

  type Person struct {
    name string
    age  int
  }
  person1 := Person{ "John", 25}
  fmt.Println(person1)
}
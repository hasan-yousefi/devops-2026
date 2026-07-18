package main
import "fmt"

func main() {
  dayOfWeek := "sunday"

  switch dayOfWeek {
    case "Saturday", "Sunday":
      fmt.Println("Weekend")

    case "Monday","Tuesday","Wednesday","Thursday","Friday":
      fmt.Println("Weekday")

    default:
      fmt.Println("Invalid day")
  }
    numbers := [5]int{21, 24, 27, 30, 33}
 
  // use range to iterate over the elements of array
  for index, item := range numbers {
    fmt.Printf("numbers[%d] = %d \n", index, item)
  }
}
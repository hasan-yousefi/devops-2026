package main
import "fmt"


// Define Global Variable
var arrayOfString = [...]string{"Hello", "github"}


func checkAGE(name string, age int) bool {
	fmt.Printf("%s I need check your age", name)
	if age < 20{
		return false
	}else{
		return true
	}
}

func main() {
	var name string
	var age int
	var sex string
	fmt.Println("Enter The Name Age Sex")
	fmt.Scanf("%s %d %s", &name, &age, &sex)
	fmt.Printf("Hi %s you are %d and %s",name, age, sex)
	acception := checkAGE(name, age)
	fmt.Println(acception)
	gen := func(age int, sex string) int{
		if sex == "male"{
			return 180 - age
		}else{
			return 160 - age
		}
	}
	fmt.Println("The maximun heart beat:", gen(age,sex))
	fmt.Println(arrayOfString)

}
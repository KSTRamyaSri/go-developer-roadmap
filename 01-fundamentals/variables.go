package fundamentals

import "fmt"

const college = "Pragati Engineering College"

var GlobalVar string = "I am a global public variable"
var privateVar string = "I am a private global variable"

func Variables() {

	// Variables
	var name string = "Sri"
	var age int = 20
	var percentage float64 = 93.5
	var isStudent bool = true

	// Short variable declaration
	course := "Computer Science"

	// Zero values
	var number int
	var price float64
	var text string
	var status bool

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Percentage:", percentage)
	fmt.Println("Student:", isStudent)
	fmt.Println("Course:", course)

	// Constant
	fmt.Println("College:", college)

	// Zero values
	fmt.Println("Zero value of int:", number)
	fmt.Println("Zero value of float64:", price)
	fmt.Println("Zero value of string:", text)
	fmt.Println("Zero value of bool:", status)
}
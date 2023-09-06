package main

import (
	"fmt"
	"learning-go/exercises"
)

func main() {
	//fmt.Println("Hello world")

	/*variables.ShowIntegers()
	variables.ShowStrings()
	variables.ShowBooleans()
	variables.ShowFloats()
	variables.ShowTime()*/

	// Functions
	// fmt.Println(functions.IntConverter(15))

	// Destructuring
	//state, value := functions.IntConverter(16)
	//fmt.Println(state)
	//fmt.Println(value)

	//conditionals.ShowOS()
	//conditionals.ShowOSWithSwitch()

	num, msg := exercises.Evaluator("13")
	fmt.Println(num)
	fmt.Println(msg)
}

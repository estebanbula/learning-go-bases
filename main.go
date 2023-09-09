package main

import (
	"learning-go/poo/controllers"
	"learning-go/poo/models"
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

	//num, msg := exercises.Evaluator("13")
	//fmt.Println(num)
	//fmt.Println(msg)

	// Key inputs
	//inputs.NumInput()

	// Iterators
	//iterators.ForIterator()

	//exercises.MultiplyTable()

	// Creating files
	//files.CreateFile()

	// Clousures
	//functions.CallClousure()

	// Recursivity
	//functions.Pow(2)

	/* Arrays */
	//arraysslices.ShowArray()

	/* Slices */
	//arraysslices.ShowSlice()
	//arraysslices.Capacity()

	/* Maps */
	//maps.ShowMap()

	/* POO */
	//controllers.NewUser()

	Matt := new(models.Man)
	controllers.PersonBreathing(Matt)

	Ana := new(models.Female)
	controllers.PersonBreathing(Ana)
}

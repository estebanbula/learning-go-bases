package functions

import "fmt"

// Clousure is a function
func Table(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number + sequence
	}
}

func CallClousure() {
	tableOf := 2
	MyTable := Table(tableOf)
	for i := 0; i < 10; i++ {
		fmt.Println(MyTable())
	}
}

package functions

import "fmt"

func Calc() {
	var num1 = 23
	var num2 = 34

	sum := func(num3 int, num4 int) int {
		return num3 + num4
	}

	fmt.Println(sum(num1, num2))
}

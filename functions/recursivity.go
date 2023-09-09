package functions

import "fmt"

func Pow(value int) {

	if value > 100000 {
		return
	}

	fmt.Println(value)
	Pow(value * 2)
}

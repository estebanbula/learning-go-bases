package variables

import "fmt"

func ShowFloats() {

	// Declaration with type
	var floatOf32 float32 = 23.3
	// Declaration without type
	var floatOf64 = 2012.3

	fmt.Println("Float of 32: ", floatOf32)
	fmt.Println("Float of 64: ", floatOf64)
}

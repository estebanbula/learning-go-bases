package variables

import "fmt"

func ShowIntegers() {
	// Definition of variables

	// * Note: int32 and int64 is not the same

	// 1. Definition with 'var':
	var commonInt int
	fmt.Println("Common int: ", commonInt)

	// 2. Definition with the type
	intOf32 := int32(20)
	fmt.Println("32 bits int: ", intOf32)
	intOf64 := int64(30)
	fmt.Println("64 bits int: ", intOf64)
}

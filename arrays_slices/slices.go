package arraysslices

import "fmt"

var slice []int = []int{2, 3, 5}

func ShowSlice() {
	fmt.Println(slice)

	names := [10]string{"Esteban", "Juan", "Andres", "José", "Juana", "Ana", "Lina", "Paula", "Jesús", "Marta"}
	names3toEnd := names[3:]
	namesStartto5 := names[:5]
	names5to7 := names[5:7]

	fmt.Println(names3toEnd, len(names3toEnd))
	fmt.Println(namesStartto5, len(namesStartto5))
	fmt.Println(names5to7, len(names5to7))
}

func Capacity() {
	/* The make function helps with the creation of structures like arrays and slices */
	/* The first argument is the type of element, the second one is the initial size of the element and the third one is the total capacity */
	elements := make([]int, 10, 20)
	fmt.Printf("Current size %d, Total capacity %d \n", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 150; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Current size %d, Total capacity %d", len(nums), cap(nums))
}

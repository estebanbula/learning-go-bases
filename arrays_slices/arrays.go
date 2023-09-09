package arraysslices

import "fmt"

var table [10]int = [10]int{0, 10, 0, 30}

func ShowArray() {
	table[7] = 70
	table[2] = 20

	names := [10]string{"Esteban", "Juan", "Andres", "José", "Juana", "Ana", "Lina", "Paula", "Jesús", "Marta"}

	fmt.Println(table)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	/* Matriz */

	var matriz [6][6]int

	fmt.Println(matriz)
}

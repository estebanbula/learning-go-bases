package maps

import (
	"fmt"
)

func ShowMap() {
	// Definition with make

	countries := make(map[string]string)

	countries["Mexico"] = "Mexico City"
	countries["Argentina"] = "Buenos Aires"
	countries["Colombia"] = "Bogotá"

	fmt.Println(countries)
	fmt.Println(len(countries))

	// Assignation
	championship := map[string]int{
		"Barcelona":   30,
		"Real Madrid": 28,
		"Junior":      20,
	}

	fmt.Println(championship)

	// Destructuring of key and value
	for team, points := range championship {
		fmt.Printf("Team %s, has %d points \n", team, points)
	}

	// Deleting data of the map
	delete(championship, "Junior")

	// We can ask for a key in the map and get the value
	points, exist := championship["Juventus"]
	fmt.Printf("The team exists? %t, and has %d points", exist, points)
}

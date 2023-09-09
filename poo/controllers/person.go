package controllers

import (
	"fmt"
	"learning-go/poo/interfaces"
)

func PersonBreathing(person interfaces.Person) {

	person.Breathe()
	fmt.Printf("I'm a %s, and i'm breathing \n", person.Genre())
}

package variables

import "fmt"

func ShowStrings() {
	var commonString string = "Var string"
	valueString := string("new string")
	fmt.Println("Var string: ", commonString)
	fmt.Println("Value string: ", valueString)
}

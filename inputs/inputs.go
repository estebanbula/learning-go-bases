package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var label string
var err error

func NumInput() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write the first number: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Write the second number: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Write the label: ")
	if scanner.Scan() {
		label = scanner.Text()
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println(label, num1*num2)
}

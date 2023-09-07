package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplyTable() string {
	scanner := bufio.NewScanner(os.Stdin)
	var content string

	fmt.Println("Please, write a number between 1 and 10: ")
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err.Error())
		}

		for {
			if num < 1 || num > 10 {
				fmt.Println("The number must be higher than 1 and lower than 10, your number is ", num)
				break
			} else {
				for i := 0; i < 11; i++ {
					content += fmt.Sprintln(num, " x ", i, " = ", num*i)
				}
				break
			}
		}
	}

	return content
}

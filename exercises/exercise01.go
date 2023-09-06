package exercises

import "strconv"

func Evaluator(value string) (int, string) {
	num, err := strconv.Atoi(value)

	if err != nil {
		return 0, "We have and error: " + err.Error()
	}

	if num > 100 {
		return num, "Is bigger than 100"
	} else {
		return num, "Is lower than 100"
	}
}

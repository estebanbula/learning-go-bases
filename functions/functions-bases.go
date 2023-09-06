package functions

import "strconv"

// IntConverter This function can return two types of values
func IntConverter(num int) (bool, string) {
	return true, strconv.Itoa(num)
}

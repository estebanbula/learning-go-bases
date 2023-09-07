package iterators

import "fmt"

func ForIterator() {
	var num = 0
	for {
		num++
		fmt.Println(num)
		if num == 10 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

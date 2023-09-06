package variables

import (
	"fmt"
	"time"
)

func ShowTime() {
	var date time.Time
	date = time.Now()
	fmt.Println("Date: ", date)
}

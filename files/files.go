package files

import (
	"fmt"
	"learning-go/exercises"
	"os"
)

func CreateFile() {
	var content = exercises.MultiplyTable()
	file, err := os.Create("./files/archive/table.txt")

	if err != nil {
		panic(err.Error())
	}

	_, errFile := fmt.Fprintln(file, content)
	if errFile != nil {
		panic(errFile.Error())
	}

	er := file.Close()
	if er != nil {
		return
	}
}

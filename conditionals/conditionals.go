package conditionals

import (
	"fmt"
	"runtime"
)

func ShowOS() {
	if runtime.GOOS == "windows" {
		fmt.Println("Is Windows")
	} else {
		fmt.Println("Is Linux or OSX")
	}
}

func ShowOSWithSwitch() {
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Is Windows")
	case "linux":
		fmt.Println("Is Linux")
	}
}

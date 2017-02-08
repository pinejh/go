package console

import "fmt"

func Log(print ...interface{}) {
	for _, n := range print {
		fmt.Print(n)
	}
	fmt.Print("\n")
}

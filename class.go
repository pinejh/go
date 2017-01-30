package main

import "fmt"

func main() {
	var array [3][5]int
	for i := 0; i < len(array); i++ {
		for z := 0; z < len(array[i]); z++ {
			fmt.Print(z, " ")
		}
		fmt.Print("\n")
	}
}

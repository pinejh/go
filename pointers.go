package main

import "fmt"

func main() {
	a, b := 5, 10
	log(a, " ", b)
	swap(&a, &b)
	log(a, " ", b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func log(print ...interface{}) {
	for _, n := range print {
		fmt.Print(n)
	}
	fmt.Print("\n")
}

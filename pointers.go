package main

import "console"

func main() {
	a, b := 5, 10
	console.Log(a, " ", b)
	swap(&a, &b)
	console.Log(a, " ", b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

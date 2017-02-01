package main

import "fmt"

func main() {
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println(factorial(input))
}

func factorial(n float64) float64 {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}

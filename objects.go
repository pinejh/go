package main

import "fmt"

func main() {
	v := vector{2, 10}
	fmt.Println("Vector 1:", v.x, v.y)
	v2 := vector{x: 3, y: 1}
	fmt.Println("Vector 2:", v2.x, v2.y)
}

type vector struct {
	x float64
	y float64
}

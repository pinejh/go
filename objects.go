package main

import "fmt"

func main() {
	v := vector{2, 10}
	fmt.Println("Vector 1:", v.x, v.y)
	v2 := vector{x: 3, y: 1}
	fmt.Println("Vector 2:", v2.x, v2.y)
	zero(v)
	fmt.Println("Zero Vector 1:", v.x, v.y)
}

type vector struct {
	x float64
	y float64
}

func zero(v *vector) {
	v.x = 0
	v.y = 0
}

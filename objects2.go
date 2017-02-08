package main

import (
	"fmt"
	"math"
)

func main() {
	circle := &Circle{1}
	fmt.Println(circle.area(), circle.circ())
	rect := &Rect{2, 3}
	fmt.Println(rect.area(), rect.perimeter())

}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) circ() float64 {
	return 2 * math.Pi * c.radius
}

type Rect struct {
	width  float64
	length float64
}

func (r *Rect) area() float64 {
	return r.width * r.length
}

func (r *Rect) perimeter() float64 {
	return r.width*2 + r.length*2
}

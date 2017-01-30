package main

import "fmt"

func main() {
	var first int
	var operator string
	var second int
	fmt.Scanf("%d", &first)
	fmt.Scanf("%s", &operator)
	fmt.Scanf("%d", &second)
	if operator == "*" {
		fmt.Println(first * second)
	} else if operator == "+" {
		fmt.Println(first + second)
	} else if operator == "-" {
		fmt.Println(first - second)
	} else if operator == "/" {
		fmt.Println(first/second, "(rounded as an integer)")
	} else {
		fmt.Println(first, operator, second)
	}
}

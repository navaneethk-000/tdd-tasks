package main

import "fmt"

func Largest(a, b, c int) int {

	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	}
	return c
}

func main() {
	fmt.Println(Largest(45, 88, 74))
}

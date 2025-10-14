package main

import "fmt"

func Largest(a, b, c int) int {
	if a > b {
		if a > c {
			return a

		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func main() {
	fmt.Println(Largest(4, 8, 7))
}

package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		// @note Go doesn't have a ternary operator, using if/else syntax is the idiomatic way!
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

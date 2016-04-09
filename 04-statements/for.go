package main

import "fmt"

func main() {
	i := 1
	for i <= 20 { // this is equal as `while` statement in C
		fmt.Println(i)
		i++
	}

	for j := 20; j >= 1; j-- {
		fmt.Println(j)
	}
}

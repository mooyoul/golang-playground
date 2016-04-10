package main

import "fmt"

func main() {
	tokens := [3]string{
		"wow",
		"such cool",
		"awesome",
	}

	for index, value := range tokens {
		fmt.Println(index, value)
	}

	// if index is not necessary, you can use underscore(_) keyword
	for _, value := range tokens {
		fmt.Println(value)
	}
}

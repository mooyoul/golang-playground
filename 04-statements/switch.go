package main

import "fmt"

func main() {
	age := 32

	switch {
	case age <= 20:
		fmt.Println("age is between 0 and 20")
	case age <= 40:
		fmt.Println("age is between 20 and 40")
	case age <= 60:
		fmt.Println("age is between 40 and 60")
	default:
		fmt.Println("age is over 60")
	}

	n := 5
	switch n % 3 {
	case 1:
		fmt.Println("n % 3: 1")
	case 2:
		fmt.Println("n % 3: 2")
	default:
		fmt.Println("n % 3: 0")
	}
}

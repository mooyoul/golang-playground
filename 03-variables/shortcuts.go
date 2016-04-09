package main

import "fmt"

func main() {
	var x = "Hello World" // Type can be optional if initial value exists
	y := 3                // `var` keyword can be optional if `:=` keyword was used

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

}

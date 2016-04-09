package main

import "fmt"

func main() {
	var x string = "Hello world"
	fmt.Println(x)
	x = "wow"
	fmt.Println(x)
	x = "such cool"
	fmt.Println(x)

	var y = "this"
	y += " is"
	y += " string"

	fmt.Println(y)
	fmt.Println("'this is string' == y: ", "this is string" == y)
}

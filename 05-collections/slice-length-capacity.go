package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("numbers slice length: ", len(numbers))
	fmt.Println("numbers slice capacity: ", cap(numbers))

	var bytes []byte
	bytes = make([]byte, 5, 10)
	fmt.Println("bytes slice length: ", len(bytes))
	fmt.Println("bytes slice capacity: ", cap(bytes))
}

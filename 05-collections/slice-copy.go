package main

import "fmt"

func main() {
	large := []int{1, 2, 3, 4, 5}
	fmt.Println("large: ", large)

	var buffer []int
	buffer = make([]int, 3)
	fmt.Println("buffer: ", buffer)

	copy(buffer, large)
	fmt.Println("buffer (copied): ", buffer)
}

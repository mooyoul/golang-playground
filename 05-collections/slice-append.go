package main

import "fmt"

func main() {
	small := []int{1, 2, 3}
	fmt.Println("small: ", small)

	large := append(small, 4, 5)
	fmt.Println("large: ", large)

	// note that append 'copies' elements
	small[0] = 100
	fmt.Println("small (after small[0] = 100): ", small)
	fmt.Println("large (after small[0] = 100): ", large)

	large[2] = 300
	fmt.Println("small (after large[2] = 300): ", small)
	fmt.Println("large (after large[2] = 300): ", large)
}

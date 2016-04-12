package main

import "fmt"

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("numbers: ", numbers)

	oddNumbers := Filter(numbers, func(n int) bool {
		if n%2 == 0 {
			return false
		} else {
			return true
		}
	})

	fmt.Println("odd numbers: ", oddNumbers)
}

package main

import "fmt"

func main() {

	// @see http://blog.golang.org/go-slices-usage-and-internals
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	// or you can create slice using `make` function
	// below two lines equal as `bytes := []byte {0, 0, 0, 0, 0}`
	var bytes []byte
	bytes = make([]byte, 5, 5)
	fmt.Println(bytes)
}

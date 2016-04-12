package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("letters: ", letters)

	chunks1 := letters[1:3]
	fmt.Println("letters[1:3]: ", chunks1)

	chunks2 := letters[3:] // equal as `letters[3:len(letters)]`
	fmt.Println("letters[3:]", chunks2)

	chunks3 := letters[:3] // equal as `letters[0:3]`
	fmt.Println("letters[:3]", chunks3)

	chunks4 := letters[:] // equal as `letters[0:len(letters)`
	fmt.Println("letters[:]", chunks4)

	chunks5 := chunks1[:1] // slice can be sliced
	fmt.Println("letters[1:3][:1]: ", chunks5)

	// note that slice references elements
	fmt.Println("original letters: ", letters)

	// modify element value in slice
	chunks5[0] = "such cool"

	fmt.Println("After modify slice: ", letters)
}

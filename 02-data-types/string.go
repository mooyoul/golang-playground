package main

import "fmt"

func main() {
	// Get string length
	fmt.Println("Length of 'Hello world': ", len("Hello world"))

	/**
	See multi bytes guide on Libraries section
	@see https://blog.golang.org/strings#TOC_7.
	*/
	fmt.Println("Length of 한글", len("한글")) // Multi-bytes string (Korean)
	fmt.Println("Length of 🏃", len("🏃"))   // note that is a single :run: emoji

	// Accessing single character
	fmt.Println("Hello World"[1]) // will print char code at specified index

	// Concatenate string
	fmt.Println("Hello " + "World")
}

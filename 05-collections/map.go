package main

import "fmt"

func main() {
	hash := make(map[string]int)
	hash["foo"] = 1
	hash["bar"] = 2

	fmt.Println(hash)

	tags := map[string]string{
		"h1":  "Heading Level 1",
		"div": "division",
		"p":   "paragraph",
	}

	fmt.Println(tags)

	name, ok := tags["h1"]
	fmt.Println(name, ok)

	name, ok = tags["h2"]
	fmt.Println(name, ok)
}

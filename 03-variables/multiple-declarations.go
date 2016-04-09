package main

import "fmt"

func main() {
	var (
		x = 1
		y = 2
		z = 10
	)

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

	const (
		FOO = "BAR"
		BAR = true
		BAZ = 123
	)

	fmt.Println("FOO: ", FOO)
	fmt.Println("BAR: ", BAR)
	fmt.Println("BAZ: ", BAZ)

}

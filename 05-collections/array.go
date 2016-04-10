package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	total := 0.0

	for i := 0; i < 5; i++ {
		total += x[i]
	}

	fmt.Println("Avg of array x: ", total/5)

	var y = [5]int{
		10,
		20,
		30,
		40,
		50,
	}

	total2 := 0
	for j := 0; j < 5; j++ {
		total2 += y[j]
	}
	fmt.Println("Avg of array y: ", total2/5)
}

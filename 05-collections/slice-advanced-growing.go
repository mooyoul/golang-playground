package main

import "fmt"

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		newSlice := make([]byte, n)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func main() {
	bytes := []byte{1, 2, 3, 4}
	fmt.Println("bytes: ", bytes)
	fmt.Println("length of bytes: ", len(bytes))
	fmt.Println("capacity of bytes: ", cap(bytes))

	bytes = AppendByte(bytes, 5, 6, 7)
	fmt.Println("bytes: ", bytes)
	fmt.Println("length of bytes: ", len(bytes))
	fmt.Println("capacity of bytes: ", cap(bytes))
}

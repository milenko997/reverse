package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
	fmt.Println(ints)
}

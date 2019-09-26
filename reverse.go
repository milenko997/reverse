package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(ints)/2; i++ {
		end := len(ints) - i - 1
		ints[i], ints[end] = ints[end], ints[i]
	}
	fmt.Println(ints)
}

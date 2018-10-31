package main

import (
	"fmt"
)

func main() {
	arr := [8]int{-2, -3, -4, -1, -2, 1, 5, -3}

	globalMaxima := 0
	localMaxima := 0
	for i := 1; i <= len(arr); i++ {
		elem := arr[i-1]
		localMaxima = max(localMaxima+elem, elem)
		if localMaxima > globalMaxima {
			globalMaxima = localMaxima
		}
	}

	fmt.Println("max sum of sub array: ", globalMaxima)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

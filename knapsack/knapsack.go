package main

import "fmt"

var val [6]int
var wt [6]int

//C is target capacity
const C = 10

var mat [len(wt) + 1][C + 1]int //matrix mat[7][11] zeroth row and zerowth column used as dummy initialization

func main() {
	val = [6]int{2, 3, 8, 10, 12, 13}
	wt = [6]int{1, 2, 3, 4, 5, 6}

	for i := 1; i <= len(wt); i++ { //rows
		currWeight := wt[i-1]
		currValue := val[i-1]

		for j := 1; j <= C; j++ { //columns
			if currWeight == j {
				maxVal := max(currValue+recurr(i-1, j-i), mat[i-1][j])
				mat[i][j] = maxVal
				continue
			}

			if currWeight < j {
				optVal := max(currValue+recurr(i-1, j-currWeight), mat[i-1][j])
				mat[i][j] = optVal
				continue
			}

			if currWeight > j {
				optVal := mat[i-1][j]
				mat[i][j] = optVal
			}
		}
	}

	fmt.Println("Max value in the knapsack: ", mat[6][10])
}

func recurr(prevRow, prevCol int) int {
	return mat[prevRow][prevCol]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

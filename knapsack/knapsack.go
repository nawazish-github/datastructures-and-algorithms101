package main

import "fmt"

var val [6]int
var wt [6]int
var mat [7][11]int

const C = 10

func main() {
	val = [6]int{2, 3, 8, 10, 12, 13}
	wt = [6]int{1, 2, 3, 4, 5, 6}

	//idx := 0
	for i := 1; i <= 6; i++ { //rows
		currWeight := wt[i-1]
		currValue := val[i-1]

		for j := 1; j <= 10; j++ { //columns
			if currWeight == j {
				//maxVal := max(currValue+recurr(col, prevRow), mat[i-1][j])
				if i == 0 {
					mat[i][j] = 0
					continue
				}
				maxVal := max(currValue+recurr(i-1, j-i), mat[i-1][j])
				mat[i][j] = maxVal
				continue
			}

			if currWeight < j {
				optVal := max(currValue+recurr(i-1, j-currWeight), mat[i-1][j])
				//prevOptVal := mat[i-1][j]
				mat[i][j] = optVal
				continue
			}

			if currWeight > j {
				//maxVal := max(currValue+recurr(i-1, i-j), mat[i-1][j])
				optVal := mat[i-1][j]
				mat[i][j] = optVal
			}
		}
	}

	fmt.Println(mat)
}

func recurr(prevRow, prevCol int) int {
	if prevRow < 0 {
		return 0
	}
	return mat[prevRow][prevCol]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

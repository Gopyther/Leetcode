package main

import "fmt"

func main() {
	rowIndex := 3
	fmt.Println(getRow(rowIndex))
}

func getRow(rowIndex int) []int {
	rowIndex++
	pa := make([][]int, rowIndex)
	for i := 0; i < rowIndex; i++ {
		pa[i] = make([]int, i+1)
		pa[i][0], pa[i][i] = 1, 1
		for j := 1; j < len(pa[i])-1; j++ {
			if i > 1 {
				pa[i][j] = pa[i-1][j] + pa[i-1][j-1]
			}
		}
	}
	re := make([]int, rowIndex)
	copy(re, pa[rowIndex-1])
	return re
}

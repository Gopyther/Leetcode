package main

import "fmt"

func main() {
	rowIndex := 3
	fmt.Println(getRow(rowIndex))
}

func getRow(rowIndex int) []int {
	arr := make([]int, rowIndex+1)
	arr[0] = 1
	for i := 0; i <= rowIndex; i++ {
		for j := i; j < 0; j-- {
			arr[j] += arr[j-1]
		}
	}
	return arr
}

// func getRow(rowIndex int) []int {
// 	prev := make([]int, 1)
// 	prev[0] = 1
// 	if rowIndex == 0 {
// 		return prev
// 	}

// 	row := 0
// 	for {
// 		row++
// 		arr := make([]int, row+1)
// 		arr[0], arr[row] = 1, 1
// 		for ind := 1; ind < row; ind++ {
// 			arr[ind] = prev[ind] + prev[ind-1]
// 		}

// 		prev = arr

// 		if row == rowIndex {
// 			break
// 		}
// 	}
// 	return prev
// }

// func getRow(rowIndex int) []int {
// 	rowIndex++
// 	pa := make([][]int, rowIndex)
// 	for i := 0; i < rowIndex; i++ {
// 		pa[i] = make([]int, i+1)
// 		pa[i][0], pa[i][i] = 1, 1
// 		for j := 1; j < len(pa[i])-1; j++ {
// 			if i > 1 {
// 				pa[i][j] = pa[i-1][j] + pa[i-1][j-1]
// 			}
// 		}
// 	}
// 	re := make([]int, rowIndex)
// 	copy(re, pa[rowIndex-1])
// 	return re
// }

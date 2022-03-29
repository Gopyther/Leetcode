package main

import "fmt"

func main() {
	numRows := 1
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	temp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		temp[i] = make([]int, i+1)
		temp[i][0], temp[i][i] = 1, 1
		if i > 1 {
			for j := 1; j < len(temp[i])-1; j++ {
				temp[i][j] = temp[i-1][j] + temp[i-1][j-1]
			}
		}
	}
	return temp
}

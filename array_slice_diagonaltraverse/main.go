package main

import (
	"fmt"
)

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// mat := [][]int{{3}, {2}}
	fmt.Println(findDiagonalOrder(mat))

}

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])

	result := make([]int, m*n)
	row, col, d := 0, 0, 1

	for i := 0; i < m*n; i++ {
		result[i] = mat[row][col]
		row -= d
		col += d

		if row >= m {
			row = m - 1
			col += 2
			d = -d
		}
		if col >= n {
			col = n - 1
			row += 2
			d = -d
		}
		if row < 0 {
			row = 0
			d = -d
		}
		if col < 0 {
			col = 0
			d = -d
		}

	}
	return result
}

// func findDiagonalOrder(mat [][]int) []int {
// 	temp := make([]int, 0)
// 	for i := range mat {
// 		for j := range mat[0] {
// 			fmt.Println(i, j)
// 			temp = append(temp, mat[i][j])
// 		}

// 	}
// 	out := make([]int, 0)

// 	for i := 0; i<= len(mat)+len(mat[0])-2; i++{
// 		if i%2 == 0{
// 			for j:=0; j<len(mat); j++{
// 				out = append(out, )
// 			}
// 		}
// 	}

// 	return out
// }

// func findDiagonalOrder(mat [][]int) []int {
// 	column := len(mat)
// 	row := len(mat[0])
// 	total := column * row
// 	cnt := 1
// 	x, y := 0, -1
// 	dir := 1 //1 is up, 0 is down
// 	temp := []int{}
// 	for cnt <= total {
// 		if x == column-1 && cnt <= total {
// 			y++
// 			cnt++
// 			temp = append(temp, mat[y][x])
// 			dir = 0
// 		}
// 		if y == row-1 && cnt <= total {
// 			x++
// 			cnt++
// 			temp = append(temp, mat[y][x])
// 			dir = 1
// 		}
// 		if y == 0 && cnt <= total {
// 			x++
// 			cnt++
// 			temp = append(temp, mat[y][x])
// 			dir = 0
// 		}
// 		if x == 0 && cnt <= total {
// 			y++
// 			cnt++
// 			temp = append(temp, mat[y][x])
// 			dir = 1
// 		}
// 		if dir == 1 && x != column-1 && y != 0 && cnt <= total {
// 			y--
// 			x++
// 			cnt++
// 			temp = append(temp, mat[y][x])
// 		}
// 		if dir == 0 && y != row-1 && x != 0 && cnt <= total {
// 			y++
// 			x--
// 			cnt++
// 			temp = append(temp, mat[y][x])
// 		}

// 	}

// 	return temp
// }

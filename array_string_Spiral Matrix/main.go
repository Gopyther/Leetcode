package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0)
	top, left, bottom, right := 0, 0, m-1, n-1
	for len(res) < m*n {
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		if top != bottom {
			for j := right - 1; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
		}
		if left != right {
			for i := bottom - 1; i > top; i-- {
				res = append(res, matrix[i][left])
			}
		}
		top++
		left++
		bottom--
		right--
	}
	return res
}

// func spiralOrder(matrix [][]int) []int {
// 	m, n, d := len(matrix), len(matrix[0]), "right"
// 	row, col := 0, -1
// 	result := make([]int, m*n)
// 	right, down, left, up := n, m, 0, 0
// 	i := 0

// 	for i < len(result) {
// 		if d == "right" {
// 			for j := col + 1; j < right; j++ {
// 				if i == len(result) {
// 					break
// 				}
// 				col = j
// 				result[i] = matrix[row][col]
// 				i++

// 			}
// 			up++
// 			d = "down"
// 		}

// 		if d == "down" {
// 			for j := row + 1; j < down; j++ {
// 				if i == len(result) {
// 					break
// 				}
// 				row = j
// 				result[i] = matrix[row][col]
// 				i++

// 			}
// 			right--
// 			d = "left"
// 		}

// 		if d == "left" {
// 			for j := col - 1; j > left-1; j-- {
// 				if i == len(result) {
// 					break
// 				}
// 				col = j
// 				result[i] = matrix[row][col]
// 				i++

// 			}
// 			down--

// 			d = "up"
// 		}

// 		if d == "up" {
// 			for j := row - 1; j > up-1; j-- {
// 				if i == len(result) {
// 					break
// 				}
// 				row = j
// 				result[i] = matrix[row][col]
// 				i++
// 			}
// 			left++
// 			d = "right"

// 		}

// 	}

// 	return result
// }

package main

func main() {
	arr := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	// arr := []int{1, 1, 1, 1, 1, 1, 1, 0, 1}
	duplicateZeros(arr)
}

func duplicateZeros(arr []int) {
	for x := 0; x < len(arr)-1; x++ {
		if arr[x] == 0 {
			arr = shift(arr, x)
			arr[x+1] = 0
			x++
		}
	}
}

func shift(arr []int, x int) []int {
	for y := len(arr) - 2; y >= x+1; y-- {
		arr[y+1] = arr[y]
	}
	return arr
}

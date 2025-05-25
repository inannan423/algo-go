package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	res := []int{}
	count := 0

	for count < len(matrix)*len(matrix[0]) {
		// 向右
		for i := left; i <= right && top <= bottom; i++ {
			res = append(res, matrix[top][i])
			count++
		}
		top++
		// 向下
		for i := top; i <= bottom && left <= right; i++ {
			res = append(res, matrix[i][right])
			count++
		}
		right--
		// 向左
		for i := right; i >= left && top <= bottom; i-- {
			res = append(res, matrix[bottom][i])
			count++
		}
		bottom--
		// 向上
		for i := bottom; i >= top && left <= right; i-- {
			res = append(res, matrix[i][left])
			count++
		}
		left++
	}

	return res
}

func main() {
	// Example usage
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := spiralOrder(matrix)
	for _, v := range result {
		print(v, " ")
	}

	fmt.Println("--")
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	result2 := spiralOrder(matrix2)
	for _, v := range result2 {
		print(v, " ")
	}
}

package main

import (
	"fmt"
)

func trap(height []int) int {
	// 前缀数组
	prefixArr := make([]int, len(height))
	prefixMax := 0
	for i := 0; i < len(height); i++ {
		prefixArr[i] = prefixMax
		prefixMax = max(prefixMax, height[i])
	}

	// 后缀数组
	subArr := make([]int, len(height))
	subMax := 0
	for i := len(height) - 1; i >= 0; i-- {
		subArr[i] = subMax
		subMax = max(subMax, height[i])
	}

	sum := 0

	for i := 0; i < len(height); i++ {
		t := min(prefixArr[i], subArr[i]) - height[i]
		if t > 0 {
			sum += t
		}
	}

	return sum
}

func main() {
	// Usage example
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	sum := trap(height)
	fmt.Println(sum)
}

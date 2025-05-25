package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	res := [][]int{}

	i := 0
	for i < len(nums)-2 {
		// 跳过相同的数字
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}

		if i >= len(nums)-1 {
			break
		}

		// 固定一个
		target := -nums[i]

		left := i + 1
		right := len(nums) - 1

		// 在后面的区间内找两数之和
		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 跳过相同的 left 或者 right
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for right > left && nums[right-1] == nums[right] {
					right--
				}

				left++
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
		i++
	}

	return res
}

func main() {
	// Example usage
	nums := []int{0, 0, 0, 0}
	result := threeSum(nums)
	for _, triplet := range result {
		println(triplet[0], triplet[1], triplet[2])
	}
}

package main

// 快速排序
func quickSelectSort(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	p := nums[0]
	pleft := []int{}
	pright := []int{}

	for i := 1; i < len(nums); i++ {
		if nums[i] < p {
			pleft = append(pleft, nums[i])
		} else {
			pright = append(pright, nums[i])
		}
	}

	pleft = quickSelectSort(pleft)
	pright = quickSelectSort(pright)
	pleft = append(pleft, p)
	pleft = append(pleft, pright...)
	return pleft
}

// 快速选择
func quickSelect(nums []int, target int) (bool, int) {
	if len(nums) == 0 {
		return false, -1
	}

	p := nums[0]
	pleft := []int{}
	pright := []int{}

	for i := 1; i < len(nums); i++ {
		if nums[i] < p {
			pleft = append(pleft, nums[i])
		} else {
			pright = append(pright, nums[i])
		}
	}

	pos := len(pleft)

	if pos == target {
		return true, p
	} else if pos < target {
		// 在右子数组查找
		find, ans := quickSelect(pright, target-len(pleft)-1)
		if find {
			return true, ans
		}
	} else {
		find, ans := quickSelect(pleft, target)
		if find {
			return true, ans
		}
	}

	return false, -1
}

func findKthLargest(nums []int, k int) int {
	find, ans := quickSelect(nums, len(nums)-k)
	if find {
		return ans
	}
	return -1
}

func main() {
	// Example usage
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	println(result) // Output: 5 (the 2nd largest element in the array)

	nums1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	result = findKthLargest(nums1, k)
	println(result) // Output: 5 (the 2nd largest element in the array)
}

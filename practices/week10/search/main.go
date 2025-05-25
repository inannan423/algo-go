package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		// 右边是有序的
		if nums[mid] <= nums[right] {
			// 在右半部分
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else if target == nums[mid] {
				return mid
			} else {
				right = mid - 1
			}
		} else { // 左边是有序的
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else if target == nums[mid] {
				return mid
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func main() {
	// Example usage
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := search(nums, target)
	println(result) // Output: 4 (index of target in nums)
}

package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	// 这道题提示不会有重复值，这种情况说明并未旋转
	if nums[left] < nums[right] {
		return nums[left]
	}
	for left < right {
		mid := left + ((right - left) >> 1)
		// 中位数可能是最小值
		if nums[mid] < nums[right] {
			right = mid
			// 中位数肯定不是最小值
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

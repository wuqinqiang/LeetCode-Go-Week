package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 7
	fmt.Printf("target %d in %d\n", target, Search(nums, target))
}

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		midValue := nums[mid]
		if midValue == target {
			return mid
		}
		// 说明此数组从中位数到右半区是有序的
		if midValue < nums[right] {
			// 说明目标值在(mid,right]之间
			if midValue < target && target <= nums[right] {
				left = mid + 1
				// 否则就在[left,mid-1]
			} else {
				right = mid - 1
			}
			// 说明此数组从左区间到中位数之间是有序的
		} else {
			if midValue > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

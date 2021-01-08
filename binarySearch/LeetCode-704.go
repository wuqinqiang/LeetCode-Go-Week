package main

import "fmt"

func main() {
	items := []int{1, 3, 6, 8, 9, 12, 16, 18, 20, 30, 35, 50}
	index := search(items, 3)
	fmt.Printf("3 is in:%v\n", index)
	index2 := search(items, 60)
	fmt.Printf("60 is in:%v\n", index2)
	index3 := search(items, 30)
	fmt.Printf("30 is in:%v\n", index3)
}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		// mid = (left + right) / 2
		// 获取中位数
		mid = left + ((right - left) >> 1)
		// 如果中位数的值等于目标值，找到了，直接返回
		if nums[mid] == target {
			return mid
			// 如果当前中位数的值 > 目标值，说明值只可能存在中位数的左区间，
			// 并且不包括中位数
		} else if nums[mid] > target {
			right = mid - 1
			// 否则当前中位数的值 < 目标值，说明值只可能存在中位数的右区间,
			// 并且不包括中位数
		} else {
			left = mid + 1
		}
	}
	return -1
}

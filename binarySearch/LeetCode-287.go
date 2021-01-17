package main

import "fmt"

func main() {
	target := []int{3, 1, 3, 4, 2}
	fmt.Printf("值是：%v\n", findDuplicate(target))
}

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	res := -1 //结果
	for left <= right {
		mid := left + ((right - left) >> 1)
		amount := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] <= mid { //统计不大于当前中位数的个数
				amount++
			}
		}
		if amount <= mid { // 如果统计个数不大于中位数，说明重复的值在右区间
			left = mid + 1
		} else { //否则重复值在左区间，并且包括中位数
			res = mid
			right = mid - 1
		}
	}
	//返回结果
	return res
}

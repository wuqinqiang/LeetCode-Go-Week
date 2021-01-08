package main

import "fmt"

func main() {
	items := []int{1, 3, 6, 8, 9, 12, 12, 12, 16, 16, 18, 20, 30, 35, 50}
	target := 12
	fmt.Printf("值是：%v\n", searchRange(items, target))
}
func searchRange(nums []int, target int) (res []int) {
	return append([]int{}, findFirstIndex(nums, target), findLastIndex(nums, target))
}

// 查找元素第一个
func findFirstIndex(nums []int, target int) int {
	var mid int
	left, right := 0, len(nums)-1
	for left <= right {
		//如果数字大会造成溢出
		// mid = (left + right)/2
		//使用位运算
		mid = left + ((right - left) >> 1)
		// 如果当前中位数的值比目标值大，说明目标值只可能存在中位数的左区间(不包括中位数)
		if nums[mid] > target {
			right = mid - 1
			// 如果当前中位数的值比目标值小，说明目标值可能只可能存在中位数的右区间
		} else if nums[mid] < target {
			left = mid + 1
		} else { //说明 此时中位数的值等于目标值，但是不能确定它就是相同目标值的第一个
			//在相等的情况下，如果当前中位数索引处是0，或者当前中位数上一个索引位置的值不等于目标值，那肯定就它了，程序结束
			if mid == 0 || (nums[mid-1] != target) {
				return mid
			} else { //否则的话 肯定在左边，就往左区间再挤挤
				right = mid - 1
			}
		}
	}
	// 如果都没找到，那就返回-1
	return -1
}

// 查找元素最后一个
func findLastIndex(nums []int, target int) int {
	var mid int
	left, right := 0, len(nums)-1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else { //说明 此时中位数的值等于目标值，但是不能确定它就是相同目标值的最后一个
			//在相等的情况下，如果当前中位数索引处于最后一个位置，或者当前中位数下一个索引位置的值不等于目标值，那肯定就它了，程序结束
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else { //否则的话，肯定再右边 就往右区间再挤挤
				left = mid + 1
			}
		}
	}
	// 如果都没找到，那就返回-1
	return -1
}

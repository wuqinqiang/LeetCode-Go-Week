package main

import "fmt"

func main() {
	n := 7
	fmt.Printf("总共需要:%d\n", climbStairs(n))
}
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	res := make([]int, len(nums), len(nums))
	res[0] = nums[0]
	res[1] = nums[1]
	if res[0] > res[1] {
		res[1] = res[0]
	}
	for i := 2; i < len(nums); i++ {
		temp := res[i-2] + nums[i]
		res[i] = max(temp, res[i-1])

	}
	return res[len(res)-1]
}
func rob2(nums []int) int {
	a := 0 // 奇数值
	b := 0 //偶数值
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			b = max(a, b+nums[i])
		} else {
			a = max(b, a+nums[i])
		}
	}
	return max(a, b)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

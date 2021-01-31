package main

import "fmt"

func main() {
	n := 7
	fmt.Printf("总共需要:%d\n", climbStairs(n))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	res := make([]int, n)
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	a, b, ret := 1, 1, 0
	for i := 2; i <= n; i++ {
		ret = a + b
		b = a
		a = ret
	}
	return ret
}

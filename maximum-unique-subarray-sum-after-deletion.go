package main

import "fmt"

func maxSum(nums []int) int {
	set := map[int]bool{}
	var (
		ans int
		mx  int = nums[0]
	)
	for _, num := range nums {
		if !set[num] {
			set[num] = true
			if num > 0 {
				ans += num
			}
		}
		mx = max(mx, num)
	}
	fmt.Println(ans, mx)
	if mx < 0 {
		return mx
	}
	return ans
}

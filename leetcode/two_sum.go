package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		got, ok := m[target-v]
		if ok {
			ans[0] = got
			ans[1] = i
			break
		}
		m[v] = i
	}
	return ans
}

func main() {
	nums := []int{3,3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
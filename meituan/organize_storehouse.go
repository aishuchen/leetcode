package main

import (
	"fmt"
)

func main() {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println(0)
		//fmt.Println(err)
		return
	}
	if count == 0 {
		fmt.Println(0)
		return
	}
	saves := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&saves[i])
	}
	takes := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&takes[i])
	}

	//fmt.Println(saves)
	//fmt.Println(takes)

	for _, idx := range takes {
		idx = idx - 1
		fmt.Println(max(sum(saves[0:idx]), sum(saves[idx+1: count])))
		saves[idx] = 0
	}
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

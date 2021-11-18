package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	for i := 0; i < len(s); i++ {
		lastIdx := strings.LastIndexByte(s, s[i])
		if lastIdx == -1 || lastIdx == i || i == lastIdx - 1 {
			lastIdx = len(s)
		}
		sub := s[i: lastIdx]
		fmt.Println(sub)
		// 统计去重字符数
		countMap := make(map[int32]int8, len(sub))
		for _, ch := range s[i:lastIdx] {
			if _, ok := countMap[ch]; !ok {
				countMap[ch] = 0
			}
		}
		if len(countMap) == len(s[i:lastIdx]) {
			return len(countMap)
		}
	}
	return 0
}

func main()  {
	fmt.Println(lengthOfLongestSubstring("abb"))
}
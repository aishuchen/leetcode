package main

import (
	"fmt"
	"strings"
)

func maxProduct(words []string) int {
	var ans int
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}
	for i, x := range masks {
		for j, y := range masks[:i] {
			length := len(words[i]) * len(words[j])
			if x&y == 0 && length > ans {
				ans = length
			}
		}
	}
	return ans
}

func main()  {
	var input string
	fmt.Scan(&input)
	input = strings.ReplaceAll(input, "\"", "")
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	words := strings.Split(input, ",")
	fmt.Println(words)
	fmt.Println(maxProduct(words))
}

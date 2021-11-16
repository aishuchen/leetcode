package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var count int
	inputSave := make([]interface{}, count)
	inputTake := make([]interface{}, count)
	_, err := fmt.Scanln(&count)
	if err != nil {
		//fmt.Println(0)
		return
	}
	_, err = fmt.Scanln(inputSave...)
	if err != nil {
		//fmt.Println(0)
	}
	_, err = fmt.Scanln(inputTake...)
	if err != nil {
		//fmt.Println(0)
	}
	//save := handleInputs(inputSave)
	//if len(save) != count {
	//	fmt.Println(0)
	//}
	//take := handleInputs(inputTake)
	//if len(take) != count {
	//	fmt.Println(0)
	//}
	fmt.Println(count)
	fmt.Println(inputSave)
	//fmt.Println(inputTake)
	//for _, takeIdx := range take {
	//	idx := takeIdx - 1
	//	save[idx] = 0
	//	fmt.Println(max(sum(save[0:idx]), sum(save[idx: count])))
	//}
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

func handleInputs(inputs string)[]int {
	ss := strings.Split(inputs, " ")
	ret := make([]int, len(ss))
	for idx, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		ret[idx] = i
	}
	return ret
}

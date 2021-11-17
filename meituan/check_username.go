package main

import "fmt"

func printWrong(){
	fmt.Println("Wrong")
}

func printAccept() {
	fmt.Println("Accept")
}

// 思路：
// 1.第一位在 A-z 之间 <==> ascii -> [65, 90] || [97, 122]
// 2.其余的 在 0-9 A-z 之间 <==> ascii -> [48, 57] || [65, 90] || [97, 122]
// 3.满足1的基础上满足2中的属于数字即可
func checkNames(names []string) {
	for _, name := range names {
		nameAscii := []rune(name)
		if checkName(nameAscii) {
			printAccept()
		} else {
			printWrong()
		}

	}
}

func checkName(name []rune) bool {
	first := name[0]
	hasNum := false
	if !(((65<= first) && (first <= 90)) || ((97 <= first) && (first <= 122))) {
		return false
	}
	for i := 1; i < len(name); i++ {
		other := name[i]
		if !(((65<= other) && (other <= 90)) || ((97 <= other) && (other <= 122))) {
			if !(48 <= other && other <= 57) {
				return false
			}
			hasNum = true
		}
	}
	return hasNum
}

func main() {
	var inputLen int
	_, err := fmt.Scan(&inputLen)
	if err != nil {
		printWrong()
	}

	inputs := make([]string, inputLen)
	for i := 0; i < inputLen; i++ {
		var item string
		fmt.Scan(&item)
		inputs[i] = item
	}
	checkNames(inputs)

}


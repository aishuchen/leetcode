package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	chain(recursion(l1, 0, 0) + recursion(l2, 0, 0), ans)
	return ans
}

func chain(num int, l *ListNode) {
	n := num % 10
	if n == 0 && num < 10 {
		l.Val = num
		l.Next = nil
		return
	}
	l.Val = n
	num = (num - n) / 10
	l.Next = new(ListNode)
	chain(num, l.Next)
}

func recursion(l *ListNode, depth, ret int) int {
	if l.Next == nil {
		return l.Val * pow(10, depth) + ret
	}
	ret += l.Val * pow(10, depth)
	depth++
	return recursion(l.Next, depth, ret)
}

func pow(base, p int) int {
	ret := 1
	for i := 0; i < p; i++ {
		ret *= base
	}
	return ret
}

func main()  {
	l := new(ListNode)
	chain(465, l)
	fmt.Println(l)
	fmt.Println(recursion(l, 0, 0))
}

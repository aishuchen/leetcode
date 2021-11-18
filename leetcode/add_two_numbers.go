package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret, head *ListNode
	//buf := make([]int, 0, 100)
	var more int
	for {
		if l1 == nil && l2 == nil && more == 0 {
			break
		}
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		s := v1 + v2 + more
		if s >= 10 {
			s = s - 10
			more = 1
		} else {
			more = 0
		}
		if ret == nil {
			head = &ListNode{Val: s}
			ret = head
		} else {
			head.Next = &ListNode{Val: s}
			head = head.Next
		}
	}
	return ret
}

func chain(nums []int) *ListNode {
	var ret = &ListNode{Val: nums[len(nums)-1]}
	for i:=len(nums)-2; i >= 0 ; i-- {
		node := &ListNode{
			Val: nums[i],
			Next: ret,
		}
		ret = node
	}
	return ret
}

func printNode(l *ListNode) {
	if l == nil {
		fmt.Println()
		return
	}
	fmt.Print(l.Val, " ")
	printNode(l.Next)
}

func main()  {
	nums1 := []int{9,9,9,9,9,9,9}
	l1 := chain(nums1)
	printNode(l1)

	nums2 := []int{9,9,9,9}
	l2 := chain(nums2)
	printNode(l2)

	printNode(addTwoNumbers(l1,l2))
}

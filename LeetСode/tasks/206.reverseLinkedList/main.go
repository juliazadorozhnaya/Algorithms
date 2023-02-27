package main

import "fmt"

// ListNode - структура для односвзяного списка
type ListNode struct {
	Next *ListNode
	Val  int
}

// Print - печать списков
func (head *ListNode) Print() {
	cur := head
	for cur != nil {
		splitter := ""
		if cur != head {
			splitter = "->"
		}
		fmt.Printf("%s%d", splitter, cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

// Reverse - метод "разворачивания" списка
func (head *ListNode) Reverse() *ListNode {
	current := head
	var prev, next *ListNode
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n1.Print()
	n1.Reverse().Print()
}

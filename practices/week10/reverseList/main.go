package main

import (
	"bufio"
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代法
func reverseListIter(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head

	p := head.Next

	for p != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}

	head.Next = nil
	return prev
}

func reverseListRecu(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecu(head.Next)

	// head 接到原来的下一个的后面
	head.Next.Next = head
	head.Next = nil // 断开连接

	return newHead
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var length int
	_, err := fmt.Fscan(reader, &length)
	if err != nil {
		panic("error")
	}

	root := &ListNode{}
	t := root
	for i := 0; i < length; i++ {
		var val int
		fmt.Fscan(reader, &val)
		node := &ListNode{
			Val: val,
		}
		t.Next = node
		t = node
	}

	r := reverseListRecu(root.Next)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	t := head
	// 计算节点数量
	for t != nil {
		n++
		t = t.Next
	}

	// 创建 dummy 节点，辅助
	dummy := &ListNode{Next: head}
	p0 := dummy        // p0 表示当前需要翻转的链表的前一个节点（用来辅助连通）
	curr := p0.Next    // 当前的节点
	var prev *ListNode // 前一个节点

	for ; k <= n; n -= k {
		// 遍历 k 次翻转
		for i := 0; i < k; i++ {
			next := curr.Next // 先记录 curr 的下一个节点
			curr.Next = prev  // 翻转当前节点指向前一个节点
			prev = curr       // 把当前节点设为前一个节点为下次操作做准备
			curr = next       // 记录下一个需要操作的节点
		}

		next := p0.Next
		p0.Next.Next = curr // 修改当前反转的队尾连上下一个 k 的节点
		p0.Next = prev      // p0 的 next 要改为翻转后的头节点
		p0 = next           // 将 p0 移到翻转后的尾节点
	}

	return dummy.Next
}

func main() {
	// Example usage:
	head := &ListNode{Val: 1}
	t := head
	for i := 2; i <= 10; i++ {
		t.Next = &ListNode{Val: i}
		t = t.Next
	}
	k := 3
	result := reverseKGroup(head, k)
	for result != nil {
		print(result.Val, " ")
		result = result.Next
	}
}

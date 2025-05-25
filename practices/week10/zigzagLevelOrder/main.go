package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(nodes []*TreeNode) {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := [][]*TreeNode{{root}}
	res := [][]int{}

	level := 0

	for len(queue) > 0 {
		temp := []int{}
		nodes := queue[0]
		queue = queue[1:]

		t := []*TreeNode{}

		for _, node := range nodes {
			if node.Left != nil {
				t = append(t, node.Left)
			}
			if node.Right != nil {
				t = append(t, node.Right)
			}
		}

		if level%2 != 0 {
			reverse(nodes)
		}

		for _, node := range nodes {
			temp = append(temp, node.Val)
		}

		if len(t) != 0 {
			queue = append(queue, t)
		}

		res = append(res, temp)
		level++
	}

	return res
}

func main() {
	// example usage
	// [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(zigzagLevelOrder(root))

	fmt.Println("___________")

	// [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Right = &TreeNode{Val: 5}
	root2.Right.Left = &TreeNode{Val: 6}
	root2.Right.Right = &TreeNode{Val: 7}
	root2.Left.Left.Left = &TreeNode{Val: 8}
	root2.Left.Left.Right = &TreeNode{Val: 9}
	root2.Left.Right.Left = &TreeNode{Val: 10}
	root2.Left.Right.Right = &TreeNode{Val: 11}
	root2.Right.Left.Left = &TreeNode{Val: 12}
	root2.Right.Left.Right = &TreeNode{Val: 13}
	root2.Right.Right.Left = &TreeNode{Val: 14}
	root2.Right.Right.Right = &TreeNode{Val: 15}
	fmt.Println(zigzagLevelOrder(root2))
}

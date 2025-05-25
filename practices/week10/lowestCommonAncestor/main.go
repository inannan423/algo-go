package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果在两侧找到了说明当前这个就是最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 两边都没有就是没有
	// 或者出现一个，那么那个就是要的结果（因为另一个只能在更深了）
	if left != nil {
		return left
	}

	return right
}

func main() {
	// example usage
	root := &TreeNode{Val: 3}
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 1}
	root.Left = p
	root.Right = q
	p.Left = &TreeNode{Val: 6}
	p.Right = &TreeNode{Val: 2}
	q.Left = &TreeNode{Val: 0}
	q.Right = &TreeNode{Val: 8}
	lca := lowestCommonAncestor(root, p, q)
	if lca != nil {
		println(lca.Val) // should print 3
	} else {
		println("No common ancestor found")
	}
}

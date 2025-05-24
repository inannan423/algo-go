package main

import (
	"fmt"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 判断是不是有效字符串
func isValid(s string) bool {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			if count > 0 {
				count--
			} else {
				// 如果有 ) 但是没有可以匹配的 ( 那么就是无效的
				return false
			}
		}
	}

	return count == 0 // 如果有剩下的 ( 那么是无效的字符串
}

func removeInvalidParentheses(s string) []string {
	queue := []string{s}
	visited := map[string]bool{}
	visited[s] = true

	res := []string{}

	for len(queue) > 0 {
		level := len(queue) // 当前层有多少需要遍历的
		found := false
		// 遍历这一层
		for x := 0; x < level; x++ {
			// 出队
			str := queue[0]
			queue = queue[1:]

			if isValid(str) {
				found = true
				res = append(res, str)
			}

			// 如果找到了，那么不再生成新的层，新的层肯定要删除更多括号
			if found {
				continue
			}

			// 生成新的层，遍历当前的字符串去除字符串的某一位加入队列
			for i := 0; i < len(str); i++ {
				// 跳过非括号字符处理
				if str[i] != '(' && str[i] != ')' {
					continue
				}
				// 去除这一位加入 queue（前提是没有遍历过）
				newStr := str[:i] + str[i+1:]
				if !visited[newStr] {
					visited[newStr] = true
					queue = append(queue, newStr)
				}
			}
		}

		// 不再遍历下一层
		if found {
			break
		}
	}

	return res
}

func main() {
	s := "()())()"
	fmt.Println(removeInvalidParentheses(s))
}

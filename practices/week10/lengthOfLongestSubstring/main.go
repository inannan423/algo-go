package main

import (
	"bufio"
	"fmt"
	"os"
)

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int, len(s))
	left := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok && val >= left {
			left = val + 1
		}

		maxLen = max(maxLen, i-left+1)
		m[s[i]] = i
	}

	return maxLen
}

func main() {
	// 从标准输入读取数据
	reader := bufio.NewReader(os.Stdin)
	var s string
	_, err := fmt.Fscan(reader, &s)
	if err != nil {
		panic(err)
	}

	// 调用函数并输出结果
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

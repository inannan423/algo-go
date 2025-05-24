package main

import (
	"fmt"
)

// 检查是否符合回文
func checkPalindrome(s string, a, b int) (string, int) {
	for a >= 0 && b < len(s) {
		if s[a] == s[b] {
			a--
			b++
		} else {
			break
		}
	}

	return s[a+1 : b], b - a
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	maxRes := ""
	for i := 0; i < len(s)-1; i++ {
		signle, slen := checkPalindrome(s, i, i)
		double, dlen := checkPalindrome(s, i, i+1)
		if slen > dlen {
			if slen > len(maxRes) {
				maxRes = signle
			}
		} else {
			if dlen > len(maxRes) {
				maxRes = double
			}
		}
	}

	return maxRes
}

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// var s string
	// _, err := fmt.Fscan(reader, &s)

	// if err != nil {
	// 	panic("error")
	// }

	// fmt.Println(longestPalindrome(s))
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

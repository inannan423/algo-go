package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	a := -prices[0]
	b := 0

	for i := 1; i < len(prices); i++ {
		// 持有：继续持有或者买入
		a = max(a, -prices[i])
		// 不持有：卖出(前一天持有的例利润 + 卖出的价格）或者继续不持有
		b = max(a+prices[i], b)
	}

	return b // 最后一天不能持有
}

func main() {
	s := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(s))
	s1 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(s1))
}

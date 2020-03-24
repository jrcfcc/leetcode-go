package arrays

import (
	"fmt"
	"strings"
)

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

解题:
生成输入数组，存储从 1 到 N 的数字。
计算从 0 到 (N - 1)! 的所有阶乘数。
在 (0, N! - 1) 区间内，k 重复减 1。
计算 k 的阶乘，使用阶乘系数构造排列。
返回排列字符串。
*/
func getPermutation(n int, k int) string {
	//存储阶乘数
	factorial := make([]int, n+1)
	factorial[0] = 1
	tokens := make([]int, n)
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1]*i
		tokens[i-1] = i
	}

	k--
	var b strings.Builder
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b.String()
}

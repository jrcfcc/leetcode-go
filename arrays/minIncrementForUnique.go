package arrays

import "sort"

/*
给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。
返回使 A 中的每个值都是唯一的最少操作次数。

先排序,然后将后面的数变为比前面的数大一
*/
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	var opt = 0
	for i:=1;i<len(A);i++{
		if A[i] == A[i-1] {
			opt++
			A[i]++
		}else if A[i] < A[i-1] {
			opt = opt + A[i-1] - A[i] + 1
			A[i] = A[i-1] + 1
		}
	}
	return opt
}

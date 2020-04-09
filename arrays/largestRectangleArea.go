package arrays

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/
/*
解法一 暴力法,查找每个柱子能画出的最大面积
高为当前柱子的高度,宽为从当前柱子往左右两边扩展的大于等于当前柱子高度的左右边界之查
*/
func largestRectangleArea(heights []int) int {
	var l = len(heights)
	var res = 0
	for i := 0; i < l; i++ {
		var left, right = i, i
		var currHeight = heights[i]
		for left > 0 && heights[left-1] >= currHeight {
			left--
		}
		for right < l-1 && heights[right+1] >= currHeight {
			right++
		}
		currArea := currHeight * (right - left + 1)
		if currArea > res {
			res = currArea
		}
	}
	return res
}

//优化,使用单调升序栈 + 哨兵,栈存储索引
func largestRectangleArea2(heights []int) int {
	var res = 0
	//首尾加一个0作为哨兵,避免为空判断
	heights = append(heights, 0)
	heights = append([]int{-1}, heights...)
	var l = len(heights)
	var stack = make([]int, 1, l)
	for i := 1; i < l; {
		//栈顶小于当前高度,入栈
		for i < l && heights[stack[len(stack)-1]] < heights[i] {
			stack = append(stack, i)
			i++
		}
		//当前栈顶元素的上一个元素小于栈顶,
		//宽度就是当前索引i减去栈顶上一个元素记录的索引stack[len(stack)-2]-1
		//i - stack[len(stack)-2] -1
		currArea := heights[stack[len(stack)-1]] * (i - stack[len(stack)-2] - 1)
		if currArea > res {
			res = currArea
		}
		stack = stack[:len(stack)-1]
	}
	return res
}

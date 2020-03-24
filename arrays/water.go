package arrays

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，
下雨之后能接多少雨水。
*/
/*
解法1: 直接按问题描述进行。对于数组中的每个元素，
我们按列找出下雨后水能达到的最高位置，等于两边最大高度的较小值减去当前高度的值。
左右两边的元素可以忽略,肯定不能存水,因为缺少墙
从当前元素开始扫描,可以忽略差值是否大于0的问题
*/

/*
解法2: 还是按列求能存的雨水,利用动态规划,一波解决
*/

func trap2(height []int) int {
	var l = len(height)
	var ans = 0

	var minner = 0
	for i:=1;i<l-1;i++ {
		var maxLeft,maxRight = 0,0
		for j:=i;j>=0;j--{
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		for j:=i;j<l;j++{
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		minner = maxLeft
		if maxLeft > maxRight {
			minner = maxRight
		}
		ans += minner - height[i]
	}
	return ans
}

func trap(height []int) int {
	var l = len(height)
	var ans = 0
	var maxLeft,maxRight = make([]int,l),make([]int,l)
	//遍历找到所有元素的左边最大值
	for i:=1;i<l-1;i++{
		maxLeft[i] = maxLeft[i-1]
		if height[i-1] > maxLeft[i] {
			maxLeft[i] = height[i-1]
		}
	}
	//遍历找到所有元素的右边最大值
	for i:=l-2;i>=0;i--{
		maxRight[i] = maxRight[i+1]
		if height[i+1] > maxRight[i] {
			maxRight[i] = height[i+1]
		}
	}
	var min = 0
	//取该列左右最大值中的较小值,减去当前列,就是该列能存储的雨水
	for i:=1;i<l-1;i++ {
		min = maxLeft[i]
		if maxLeft[i] > maxRight[i] {
			min = maxRight[i]
		}
		if min > height[i] {
			ans = ans + (min - height[i])
		}
	}
	return ans
}

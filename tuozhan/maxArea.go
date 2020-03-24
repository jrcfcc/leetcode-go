package tuozhan

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。

双指针法:
这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。此外，两线段距离越远，得到的面积就越大。

我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。 此外，我们会使用变量 maxarea 来持续存储到目前为止所获得的最大面积。
在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxarea，并将指向较短线段的指针向较长线段那端移动一步。

*/
func maxArea(height []int) int {
	var l = len(height)
	var begin,end = 0,l-1
	var max = 0
	for begin < end {
		h := height[begin]
		if height[end] < h {
			h = height[end]
		}
		area := h * (end - begin)
		if area > max {
			max = area
		}
		if height[begin] < height[end] {
			begin++
		}else{
			end--
		}
	}
	return max
}

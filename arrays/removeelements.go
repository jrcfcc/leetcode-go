package arrays

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，
并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
func removeElement(nums []int, val int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var begin,end = 0,length-1
	for begin < end {
		//从后往前找不等于val的元素
		for begin < end && nums[end] == val {
			end--
		}
		//从前往后找等于val的元素
		for begin < end && nums[begin] != val {
			begin++
		}
		//将等于val的元素移到数组后面
		nums[begin],nums[end] = nums[end],nums[begin]
	}
	if nums[0] == val {
		return 0
	}
	return end + 1
}

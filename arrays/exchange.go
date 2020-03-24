package arrays

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，
所有偶数位于数组的后半部分。
*/
func exchange(nums []int) []int {
	var l = len(nums)
	var left,right = 0,l-1
	for left < right {
		//找到左边的第一个偶数
		for left < right && nums[left] & 1 == 1 {
			left++
		}
		//找到右边的第一个奇数
		for left < right && nums[right] & 1 == 0 {
			right--
		}
		//交换
		nums[right],nums[left] = nums[left],nums[right]
	}
	return nums
}

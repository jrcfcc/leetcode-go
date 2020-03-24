package arrays


/*
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

解题思路:
解决方法：滑动窗口
算法：

每个（连续）增加的子序列是不相交的，并且每当 nums[i-1]>=nums[i] 时，
每个此类子序列的边界都会出现。当它这样做时，它标志着在 nums[i] 处开始一个新的递增子序列，
我们将这样的 i 存储在变量 anchor 中。
例如，如果 nums=[7，8，9，1，2，3]，那么 anchor 从 0 开始（nums[anchor]=7），
并再次设置为 anchor=3（nums[anchor]=1）。无论 anchor 的值如何，
我们都会记录 i-anchor+1 的候选答案、子数组 nums[anchor]、nums[anchor+1]、…、nums[i]
的长度，并且我们的答案会得到适当的更新。

*/
func findLengthOfLCIS(nums []int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var maxLength = 1
	var currLength = 1
	for i:=1;i<length;i++{
		if nums[i-1] >= nums[i] {
			currLength = 1
		}else{
			currLength++
			if maxLength < currLength {
				maxLength = currLength
			}
		}
	}
	return maxLength
}

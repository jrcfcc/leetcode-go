package arrays

import "sort"

/*
题目描述:
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。


标签：数组遍历
首先对数组进行排序，排序后固定一个数 nums[i]，再使用左右指针指向 nums[i]后面的两端，
数字分别为 nums[L] 和 nums[R]，计算三个数的和 sum 判断是否满足为 0，满足则添加进结果集
如果 nums[i]大于 0，则三数之和必然无法等于 0，结束循环
如果 nums[i] == nums[i-1]，则说明该数字重复，会导致结果重复，所以应该跳过
当 sum == 0 时，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
当 sum == 0 时，nums[R] == nums[R-1] 则会导致结果重复，应该跳过，R--
*/
func threeSum(nums []int) [][]int {
	var length = len(nums)
	if length < 3{
		return nil
	}
	var rtn = [][]int{}
	sort.Ints(nums)
	for i := 0 ;i<length; i++{
		//nums[i]是三个数中最小的,且是从小到大循环的
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var l = i+1
		var r = length - 1
		for l< r {
			if nums[i] + nums[l] + nums[r] > 0 {
				r--
			}else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			}else{
				rtn = append(rtn,append([]int{},nums[i],nums[l],nums[r]))
				//去除重复值
				for l < r && nums[l] == nums [l+1] {
					l++
				}
				//去除重复值
				for l < r && nums[r] == nums [r-1] {
					r--
				}
				//找到后,两个指针都进一步
				l++
				r--
			}
		}
	}
	return rtn
}

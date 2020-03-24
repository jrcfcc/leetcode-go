package arrays

import "sort"

/*
四数之和
基于三数之和
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var l = len(nums)
	if l < 4 {
		return nil
	}
	if nums[l-4] + nums[l-3] + nums[l-2] + nums[l-1] < target {
		return nil
	}
	var res = make([][]int,0)
	for i:=0;i<=l-4;i++{
		//防止重复值
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		if nums[i] > target {
			break
		}
		for j:=i+1;j<=l-3;j++{
			//防止重复值
			if j>i+1 && nums[j] == nums[j-1]{
				continue
			}
			var left,right = j+1,l-1
			for left < right {
				var v = nums[i] + nums[j] + nums[left] + nums[right]
				if v == target {
					res = append(res,[]int{nums[i],nums[j],nums[left],nums[right]})
					for left < right && nums[left] == nums[left+1]{
						left++
					}
					for left < right && nums[right] == nums[right-1]{
						right--
					}
					//找到正确的值后,两个指针同时前进
					left++
					right--
				}
				if v > target {
					right--
				}
				if v < target {
					left++
				}
			}
		}
	}
	return res
}

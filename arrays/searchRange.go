package arrays

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。
找出给定目标值在数组中的开始位置和结束位置。
示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。
*/
func searchRange(nums []int, target int) []int {
	var length = len(nums)
	if length == 0 {
		return []int{-1,-1}
	}
	if target < nums[0] {
		return []int{-1,-1}
	}
	if target > nums[length-1] {
		return []int{-1,-1}
	}
	var res = make([]int,2)
	leftIdx := findBound(nums,target,length,true)
	res[0] = leftIdx
	if leftIdx != -1 {
		res[1] = findBound(nums,target,length,false)
	}else{
		res[1] = -1
	}
	return res
}

//查找边界
func findBound(nums []int, target,length int,searchLeft bool) int{
	var left,right = 0,length -1
	for left <= right {
		var mid = left + ((right - left)>>1)
		if nums[mid] == target {
			if searchLeft {
				//查找左边界则缩小右边界值
				right = mid - 1
			}else{
				//查找右边界则缩小左边界值
				left = mid + 1
			}
		}else if nums[mid] < target {
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	if searchLeft {
		if nums[left] != target {
			return -1
		}
		return left
	}
	if nums[right] != target {
		return -1
	}
	return right
}

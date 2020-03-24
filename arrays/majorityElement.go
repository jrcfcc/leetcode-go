package arrays

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于  n/2  的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func majorityElement(nums []int) int {
	var n = len(nums)
	if n == 1 {
		return nums[0]
	}
	var set = make(map[int]int)
	for i:=0;i<n;i++{
		if v,ok := set[nums[i]];ok{
			v++
			if v > n >> 1 {
				return nums[i]
			}
			set[nums[i]] = v
		}else{
			set[nums[i]] = 1
		}
	}
	return 0
}

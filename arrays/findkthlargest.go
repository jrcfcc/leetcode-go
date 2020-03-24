package arrays

import "math/rand"

/*
题目描述:在未排序的数组中找到第 k 个最大的元素。请注意，
你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

解题思路:最简单的思路就是数组排好序,然后返回nums[len - k]
这样的话复杂度,就取决于排序算法的复杂读

解题思路二:可以利用快排算法中的减而治之(分而治之的特例)的思想
我们最终要找的是nums中第k大的元素,而不是将整个数组排序
我们每次先找到待排序数组中的一个privt,这个数处于它正确的位置上
所谓正确的顺序就是数组排完序后privt的位置不会变
如果 i < privt 那么一定有nums[i] < nums[privt]
如果 i > privt 那么一定有nums[i] > nums[privt]
整个算法复杂度就相等于二分查找的复杂度
*/
func FindKthLargest(nums []int, k int) int {
	var left,right = 0,len(nums) - 1
	var target = len(nums) - k
	for {
		mid := parttiion(nums,left,right)
		if mid == target {
			return nums[mid]
		}
		if mid < target {
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
}

//查找排序区间的中位数
func parttiion(nums []int,left,right int) int {
	//随机挑选比较的基准值
	if right > left {
		random := left + 1 + rand.Intn(right - left)
		nums[left],nums[random] = nums[random],nums[left]
	}
	var pivot = nums[left]
	var j = left
	for i := left + 1;i<=right;i++ {
		//将小于基准值的都移到pivot的左边
		if nums[i] < pivot {
			j++
			nums[j],nums[i] = nums[i],nums[j]
		}
	}
	//将pivot放到它应该在的位置
	nums[j],nums[left] = nums[left],nums[j]
	return j
}

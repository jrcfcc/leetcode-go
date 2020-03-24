package arrays

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。
二分查找 查找特定值
1 2 3 4 5 6 7 可以大致分为两类，
第一类 2 3 4 5 6 7 1 这种，也就是 nums[start] <= nums[mid]。此例子中就是 2 <= 5。
这种情况下，前半部分有序。因此如果 nums[start] <=target<nums[mid]，则在前半部分找，否则去后半部分找。
第二类 6 7 1 2 3 4 5 这种，也就是 nums[start] > nums[mid]。此例子中就是 6 > 2。
这种情况下，后半部分有序。因此如果 nums[mid] <target<=nums[end]，则在后半部分找，否则去前半部分找。
*/
func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	var left,right = 0,len(nums)-1
	for left <= right {
		var mid = left + ((right-left)>>1) //防止left+right 溢出
		if target == nums[mid] {
			return mid
		}
		//如果前半部分有序,此处小于等于是为了防止left = mid的情况
		if nums[left] <= nums[mid] {
			//target处于有序的左半区
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			//target处于右半区
			}else{
				left = mid + 1
			}
		}
		//如果后半部分有序,此处小于等于是为了防止mid = right的情况
		if nums[mid] < nums[right] {
			//target处于有序的右半区
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			//target处于左半区
			}else{
				right = mid - 1
			}
		}
	}
	return -1
}

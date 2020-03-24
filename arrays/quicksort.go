package arrays

import "math/rand"

func sortArray(nums []int) []int{
	var length = len(nums)
	if length < 2 {
		return nums
	}
	quicksort(nums,0,length-1)
	return nums
}

//快速排序
func quicksort(nums []int,left,right int){
	if left >= right {
		return
	}
	//取随机索引作为比对的基准值
	var random = left + 1 + rand.Intn(right - left)
	nums[left],nums[random] = nums[random],nums[left]
	var tmp = nums[left]
	var begin,end = left,right
	for begin < end {
		//指针先从右向左扫描
		if nums[end] >= tmp && begin < end {
			end--
		}
		//找到一个小于基准值的数后就赋值给left
		nums[begin] = nums[end]
		//指针从左向右扫描
		if nums[begin] <= tmp && begin < end {
			begin++
		}
		//找到一个大于基准值的数后就复制给right
		nums[end] = nums[begin]
	}
	//将tmp值赋值给左右指针相遇的点
	nums[begin] = tmp
	//tmp的左右两边继续分割
	quicksort(nums,left,begin-1)
	quicksort(nums,begin+1,right)
}

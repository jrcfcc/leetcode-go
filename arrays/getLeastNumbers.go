package arrays

import "math/rand"

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，
则最小的4个数字是1、2、3、4。
*/
func getLeastNumbers(arr []int, k int) []int {
	var l = len(arr)
	if l <= k{
		return arr
	}
	buildMaxHeap(arr[:k],k)
	//比较堆顶元素与剩下的值,如果nums[i]比堆顶元素大就不管,
	//如果比堆顶元素小,就交换堆顶与nums[i],然后调整堆
	for i:=k;i<l;i++{
		if arr[i] < arr[0] {
			arr[i],arr[0] = arr[0],arr[i]
			buildMaxHeap(arr[:k],k)
		}
	}
	return arr[:k]
}

/*
利用快排,当选择的数位于k-1时，说明左边的数+自身就是前k小的数
*/
func quickChose(arr []int, left,right,k int) {
	var base = left + 1 + rand.Intn(right-left)
	arr[base],arr[left] = arr[left],arr[base]
	var tmp = arr[left]
	var begin,end = left,right
	for begin < end {
		for begin < end && arr[end] >= tmp {
			end--
		}
		arr[begin] = arr[end]
		for begin < end && arr[begin] <= tmp {
			begin++
		}
		arr[end] = arr[begin]
	}
	arr[begin] = tmp
	//选中的数比前k个数大,只需要对左边进行快排
	if begin > k - 1 {
		quickChose(arr,0,begin-1,k)
	}
	//选中的数比前k个数小,只需要begin到最右进行选择
	if begin < k - 1 {
		quickChose(arr,begin+1,right,k)
	}
	if begin == k - 1 {
		return
	}
}

//前k个数构建大顶堆,
func buildMaxHeap(arr []int,k int) {
	//取最后一个非叶子节点,也就是l/2 -1
	last := k/2 - 1
	//从最后一个叶子节点逐步向上调整
	for i:=last;i>=0;i--{
		adjustHead(arr,i,k)
	}
}

func adjustHead(arr []int,i,length int) {
	var left = i << 1 + 1
	var right = left + 1
	var max = i
	if left < length && arr[left] > arr[max] {
		max = left
	}
	if right < length && arr[right] > arr[max] {
		max = right
	}
	if max != i {
		arr[i],arr[max] = arr[max],arr[i]
		//调整后的节点是否满足大顶堆的性质
		adjustHead(arr,max,length)
	}
}

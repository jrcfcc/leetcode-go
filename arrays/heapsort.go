package arrays

/*
堆是一种非线性结构，（本篇随笔主要分析堆的数组实现）可以把堆看作一个数组，
也可以被看作一个完全二叉树，通俗来讲堆其实就是利用完全二叉树的结构来维护的一维数组
按照堆的特点可以把堆分为大顶堆和小顶堆
大顶堆：每个结点的值都大于或等于其左右孩子结点的值
小顶堆：每个结点的值都小于或等于其左右孩子结点的值


我们用简单的公式来描述一下堆的定义就是：（读者可以对照上图的数组来理解下面两个公式）
大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]

先了解下堆排序的基本思想：
将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。
将其与末尾元素进行交换，此时末尾就为最大值。然后将剩余n-1个元素重新构造成一个堆，
这样会得到n个元素的次小值，如此反复执行，便能得到一个有序序列了，
建立最大堆时是从最后一个非叶子节点开始从下往上调整的（这句话可能不好太理解），
下面会举一个例子来理解堆排序的基本思想

堆排序过程（假设我们想要升序的排列）
第一步：先n个元素的无序序列，构建成大顶堆
第二步：将根节点与最后一个元素交换位置，（将最大元素"沉"到数组末端）
第三步：交换过后可能不再满足大顶堆的条件，所以需要将剩下的n-1个元素重新构建成大顶堆
第四步：重复第二步、第三步直到整个数组排序完成

那么最后一个非叶子节点就是：长度/2-1
*/

func headSort(nums []int) []int{
	var length = len(nums)
	if length <= 1 {
		return nums
	}
	for i:=length;i>0;i--{
		buildHeap(nums,i)
		swap(nums,0,i-1)
	}
	return nums
}

//构造大顶堆
func buildHeap(nums []int,length int) {
	var lastChildIndex = length/2 - 1
	//从最后一个非叶子节点往上逐步调整
	for i:=lastChildIndex;i>=0;i--{
		adjudgeHeap(nums,i,length)
	}
}

//调整堆
func adjudgeHeap(nums []int,i,length int){
	var left = 2 * i + 1
	var right = left + 1
	var max = i //记录根,左,右中最大值的下标
	//根节点小于左子树
	if left < length && nums[max] < nums[left] {
		//交换根节点跟左子节点的值
		max = left
	}
	//最大值小于右子树
	if right < length && nums[max] < nums[right] {
		max = right
	}
	if max != i {
		//将根节点与最大值交换
		swap(nums,i,max)
		//调整后,看调整后的堆是否满足大顶堆的定义
		adjudgeHeap(nums,max,length)
	}
}

func swap(nums []int,i,j int) {
	nums[i],nums[j] = nums[j],nums[i]
}

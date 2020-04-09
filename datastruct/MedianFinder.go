package datastruct

import "container/heap"

/*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，
那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
*/
//大顶堆
type MaxHeap []int
//小顶堆
type MinHeap []int

func (max MaxHeap) Len() int {
	return len(max)
}

func (max MaxHeap) Less(i,j int) bool {
	return max[i] > max[j]
}

func (max MaxHeap) Swap(i,j int) {
	max[i],max[j] = max[j],max[i]
}

func (max *MaxHeap) Push(x interface{}) {
	*max = append(*max,x.(int))
}

func (max *MaxHeap) Pop() interface{}{
	old := *max
	l := len(old)
	x := old[l-1]
	*max = old[:l-1]
	return x
}

func (max MaxHeap) Top() int{
	return max[0]
}

func (min MinHeap) Len() int {
	return len(min)
}

func (min MinHeap) Less(i,j int) bool {
	return min[i] < min[j]
}

func (min MinHeap) Swap(i,j int) {
	min[i],min[j] = min[j],min[i]
}

func (min *MinHeap) Push(x interface{}) {
	*min = append(*min,x.(int))
}

func (min MinHeap) Top() int{
	return min[0]
}

func (min *MinHeap) Pop() interface{}{
	old := *min
	l := len(old)
	x := old[l-1]
	*min = old[:l-1]
	return x
}

type MedianFinder struct {
	max *MaxHeap
	min *MinHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	var max = make(MaxHeap,0)
	var min = make(MinHeap,0)
	return MedianFinder{
		max: &max,
		min: &min,
	}
}


func (this *MedianFinder) AddNum(num int)  {
	//将元素加入到大顶堆
	heap.Push(this.max,num)
	//将大顶堆的堆顶加入到小顶堆
	heap.Push(this.min,this.max.Top())
	heap.Remove(this.max,0)
	//如果小顶堆元素比大顶堆的多,就把小顶堆的堆头放到大顶堆
	if this.min.Len() > this.max.Len() {
		heap.Push(this.max,this.min.Top())
		heap.Remove(this.min,0)
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() == this.min.Len() {
		return float64(this.max.Top() + this.min.Top()) * 0.5
	}
	return float64(this.max.Top())
}

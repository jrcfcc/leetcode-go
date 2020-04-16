package stack

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



进阶：

你能在线性时间复杂度内解决此题吗？



示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

使用单调队列
*/

//双端单调队列
type MonotonicQueue []int

//插入元素,并将队头前小于x的元素都出队
func (mq *MonotonicQueue) push(x int) {
	var old = *mq
	for old.len() > 0 && old.tail() < x {
		old = old[:old.len()-1]
	}
	old = append(old, x)
	*mq = old
}

//返回队头元素
func (mq MonotonicQueue) top() int {
	if len(mq) > 0 {
		return mq[0]
	}
	return -1
}

//返回队头元素
func (mq MonotonicQueue) tail() int {
	if len(mq) > 0 {
		return mq[len(mq)-1]
	}
	return -1
}

//弹出队头元素
func (mq *MonotonicQueue) pop(x int) {
	var old = *mq
	if len(old) > 0 && old.top() == x {
		old = old[1:]
		*mq = old
	}
}

//返回队列长度
func (mq MonotonicQueue) len() int {
	return len(mq)
}

func maxSlidingWindow(nums []int, k int) []int {
	var l = len(nums)
	if l == 0 {
		return nil
	}
	var res = make([]int, 0)
	var mq = make(MonotonicQueue, 0)
	for i := 0; i < l; i++ {
		if i < k-1 {
			mq.push(nums[i])
			continue
		}
		mq.push(nums[i])
		res = append(res, mq.top())
		mq.pop(nums[i-k+1])
	}
	return res
}

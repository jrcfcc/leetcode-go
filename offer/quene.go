package offer

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。
(若队列中没有元素，deleteHead 操作返回 -1 )
*/
type CQueue struct {
   Stack []int
}


func Constructor() CQueue {
	return CQueue{
		Stack:make([]int,0),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.Stack = append(this.Stack, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.Stack) == 0 {
		return -1
	}
	head := this.Stack[0]
	this.Stack = this.Stack[1:]
	return head
}

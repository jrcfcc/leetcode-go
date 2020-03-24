package datastruct


type MinStack struct {
	s []int
	helper []int //辅助栈,用来存储栈中的最小值
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{s: []int{},helper: []int{}}
}


func (this *MinStack) Push(x int)  {
	this.s = append(this.s,x)
	//辅助栈为空,或者入栈元素小于辅助栈的栈顶时,才入栈
	if len(this.helper) == 0 || x <= this.helper[len(this.helper)-1] {
		this.helper = append(this.helper,x)
	}
}


func (this *MinStack) Pop()  {
	top := this.Top()
	this.s = this.s[0:len(this.s)-1]
	//出栈元素等于辅助栈的栈顶时,才出栈
	if top == this.helper[len(this.helper)-1] {
		this.helper = this.helper[0:len(this.helper)-1]
	}
}


func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}


func (this *MinStack) GetMin() int {
	return this.helper[len(this.helper)-1]
}

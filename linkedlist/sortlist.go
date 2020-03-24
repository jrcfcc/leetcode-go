package linkedlist

/*
对链表进行排序,要求时间复杂度和空间复杂度分别为o(nlogn)和o(1)
解题思路:自顶向上的归并排序
归并排序:先切割再合并,
先切割成n/2个包含2个元素的最小单元,然后对最小单元排序,然后将2个最小单元合并为一个4个元素的单元
切割:将链表切割成只包含两个元素的单元,然后排序
合并:将两个相邻的已排好序的链表合并,新建一个临时链表,取两个链表表头的较小值放到临时表的next节点
重复切割合并的过程,直到整个链表处理完
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var rtn = &ListNode{}
	rtn.Next = head
	var length = 0
	var curr = head
	for curr != nil {
		length++
		curr = curr.Next
	}
	//循环一次排序合并后,i左移一位
	for i:=1;i<length;i<<=1{
		//此步骤是为了进行下次合并时方便找到链表头部
		var pre = rtn
		curr = rtn.Next
		/*
		*循环分割链表,将链表分割成length/i份
		*然后两两合并
		*/
		for curr != nil {
			var left = curr
			//第一次切割,left = 左1
			var right = cut(left,i)
			//第二次切割,right = 左2,curr=head.next3
			curr = cut(right,i)
			//合并左1和左2,将结果连接到返回链表
			pre.Next = merge(left,right)
			//将pre指向返回链表得最后一个节点
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return rtn.Next
}

//按步长分割,
func cut(head *ListNode,step int) *ListNode{
	if head == nil {
		return head
	}
	var rtn = &ListNode{}
	//当循环到步长时,head的下一个节点应该为空,left = head ,right = head.next
	for i:=1;head.Next != nil && i<step;i++{
		head = head.Next
	}
	//返回的是右边的节点,源节点已经切掉了右边的
	rtn = head.Next
	head.Next = nil
	return rtn
}

//合并,按大小依次填充到返回链表中
func merge(left,right *ListNode) *ListNode{
	var rtn = &ListNode{}
	var curr = rtn
	for left != nil && right != nil {
		if left.Val <= right.Val {
			curr.Next = &ListNode{Val:left.Val}
			left = left.Next
		}else{
			curr.Next = &ListNode{Val:right.Val}
			right = right.Next
		}
		curr = curr.Next
	}
	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}
	return rtn.Next
}

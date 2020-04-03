package linkedlist


/*
解题思路：迭代
定义两个指针： pre 和 cur ；pre 在前 cur 在后。
每次让 pre 的 next 指向 cur ，实现一次局部反转
局部反转完成之后， pre 和 cur 同时往前移动一个位置
循环上述过程，直至 pre 到达链表尾部
*/
func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	var pre = head
	for pre != nil {
		//零时变量存储原链表的下一个节点
		var t = pre.Next
		//反转链表的下一个节点设置为原链表的当前
		pre.Next = curr
		//curr向前移动
		curr = pre
		//pre向前移动
		pre = t
	}
	return curr
}

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	var n,l = 0,head
	if head == nil {
		return nil
	}
	var list = make([]*ListNode,0)
	list = append(list,l)
	for l != nil {
		n++
		l = l.Next
		if n % k == 0 {
			list = append(list,l)
		}
	}
	var c = n / k
	var begin,end *ListNode
	for i := 0 ;i < c;i++ {
		head := reverse(list[i],k)
		if begin == nil {
			begin = head
		}
		if end == nil {
			end = list[i]
		}else{
			end.Next = head
			end = list[i]
		}
	}
	if c < len(list) {
		end.Next = list[len(list)-1]
	}
	return begin
}

func reverse(head *ListNode,k int) *ListNode {
	var curr *ListNode
	var pre = head
	for pre != nil && k > 0{
		var next = pre.Next
		pre.Next = curr
		curr = pre
		pre = next
		k--
	}
	return curr
}

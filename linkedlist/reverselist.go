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

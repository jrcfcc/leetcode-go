package linkedlist

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.

双指针
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var first,second = head,head.Next
	if second == nil {
		return head
	}
	var third = second.Next
	second.Next = first
	first.Next = third
	var rtn = &ListNode{Next:second}
	var curr = first
	for curr != nil && curr.Next != nil {
		//第一个节点
		var l1 = curr.Next
		//第二个节点
		var l2 = l1.Next
		if l2 == nil {
			break
		}
		var l3 = l2.Next
		l1.Next = l3
		l2.Next = l1
		curr.Next = l2
		curr = l1
	}
	return rtn.Next
}

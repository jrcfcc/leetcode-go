package linkedlist

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var p,q,res = head,head,&ListNode{}
	var listLength = 0
	for p != nil {
		listLength++
		p = p.Next
	}
	k = k % listLength
	if k == 0 {
		return head
	}
	var move = listLength - k
	for move > 1 && q != nil {
		q = q.Next
		move--
	}
	currHead := q.Next
	res.Next = currHead
	q.Next = nil
	for currHead.Next != nil {
		currHead = currHead.Next
	}
	currHead.Next = head
	return res.Next
}

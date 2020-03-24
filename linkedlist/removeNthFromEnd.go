package linkedlist

/*
题目:给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var root,k = head,0
	for head != nil {
		head = head.Next
		k++
	}
	//如果要移除头节点
	if k == n {
		return root.Next
	}
	var curr,now,target = root,1,k-n
	for curr != nil {
		if now == target {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
		now++
	}
	return root
}

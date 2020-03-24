package linkedlist

/*
给定一个链表，判断链表中是否有环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
*/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var fast,slow = head,head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if fast == slow {
			break
		}
	}
	if fast == nil {
		return false
	}
	return true
}

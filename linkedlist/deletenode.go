package linkedlist

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
*/
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	//删除的是头节点
	if head.Val == val {
		return head.Next
	}
	var curr = head
	for curr != nil && curr.Next != nil{
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			return head
		}
		var tmp = curr
		curr = curr.Next
		//删除的是尾节点
		if curr.Next == nil && curr.Val == val{
			tmp.Next = nil
			return head
		}
	}
	//没有要删除的节点
	return head
}

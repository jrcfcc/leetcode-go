package linkedlist

/*
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
双指针法:
先将p,q分别指向head
然后将p向前移动k次
然后将p,q一起移动
当p指向null的时候,q指向的就是第k个节点
*/
func kthToLast(head *ListNode, k int) int {
	var p,q =  head,head
	for k > 0 {
		p = p.Next
		k--
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q.Val
}

package linkedlist


type  ListNode struct {
    Val int
    Next *ListNode
}

/*
合并两个有序链表
解题思路:递归调用
比较两个链表的链表头,取值较小的那个作为链表头,
next则为值小的链表的剩下节点与另一个链表的合并结果
依次递归调用,直到其中一个链表为空
list1[0]<list2[0] :  list1[0]+merge(list1[1:],list2)
list2[0]<list1[0] :  list2[0]+merge(list1,list2[1:])
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next = mergeTwoLists(l2.Next,l1)
		return l2
	}
}

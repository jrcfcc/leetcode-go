package linkedlist

/*
获取环形链表得入环点
解题思路:
map 记录每个节点,当某个节点出现重复时,就是入环点
快慢指针:快指针走两步,慢指针走一步,有环则快慢指针一定会相遇
相遇后,再设置两个指针分别指向表头和相遇点,一起向下走,再次相遇得地方就是入环点
*/
func detectCycle(head *ListNode) *ListNode {
	var intersect = getIntersect(head)
	if intersect == nil {
		return nil
	}
	var ptr1 = head
	var ptr2 = intersect
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return ptr1
}

//获取快慢指针相遇点
func getIntersect(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var fast,slow = head,head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return fast
		}
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var set = make(map[*ListNode]struct{})
	for head != nil {
		if _,ok:=set[head];ok{
			return head
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return nil
}

package linkedlist

/*
编写一个程序，找到两个单链表相交的起始节点。
双指针:
创建两个指针 pA 和 pB，分别初始化为链表 A 和 B 的头结点。然后让它们向后逐结点遍历。
当 pA 到达链表的尾部时，将它重定位到链表 B 的头结点 (你没看错，就是链表 B);
类似的，当 pB 到达链表的尾部时，将它重定位到链表 A 的头结点。
若在某一时刻 pA 和 pB 相遇，则 pA/pB 为相交结点。
原理就是:
A和B两个链表长度可能不同，但是A+B和B+A的长度是相同的，所以遍历A+B和遍历B+A一定是同时结束。
如果A,B相交的话A和B有一段尾巴是相同的，所以两个遍历的指针一定会同时到达交点
如果A,B不相交的话两个指针就会同时到达A+B（B+A）的尾节点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var ptra,ptrb = headA,headB
	//要么到达交点,要么ptra，ptrb,同时到达各自链表的尾节点
	for ptra != ptrb {
		if ptra == nil {
			ptra = headB
		}else{
			ptra = ptra.Next
		}
		if ptrb == nil {
			ptrb = headA
		}else{
			ptrb = ptrb.Next
		}
	}
	return ptra
}

/*
map存储所有节点,重复得节点就是两个链表的交点
*/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var m = make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _,ok:=m[headB];ok{
			return headB
		}
		m[headA] = struct{}{}
		headB = headB.Next
	}
	return nil
}


package linkedlist

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度

将k个链表 两两配对合并,一直到只剩一个链表
*/
func mergeKLists(lists []*ListNode) *ListNode {
	var l = len(lists)
	if l == 0 {
		return nil
	}
	for len(lists) > 1 {
		//链表两两合并
		for i:=0;i<len(lists)/2;i++{
			l1 := 2 * i
			l2 := l1 + 1
			list := mergeTwoLists(lists[l1],lists[l2])
			lists[l1/2] = list
		}
		var l = len(lists)
		var last = l >> 1
		//如果原链表是奇数,将最后一个链表放到l/2 + 1的位置上
		if l & 1 == 1{
			last++
			lists[last-1] = lists[l-1]
		}
		//保留合并后的链表
		lists = lists[:last]
	}
	return lists[0]
}



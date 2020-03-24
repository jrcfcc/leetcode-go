package offer

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var res = make([]int,0)
	for head != nil {
		res = append(res,head.Val)
		head = head.Next
	}
	var l = len(res)
	var left,right = 0,l-1
	for left < right {
		res[left],res[right] = res[right],res[left]
		left++
		right--
	}
	return res
}

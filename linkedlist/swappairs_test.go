package linkedlist

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	var head = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
			},
		},
	}
	swapPairs(head)
}
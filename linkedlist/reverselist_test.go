package linkedlist

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	var head = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next:&ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reverseKGroup(head,2)
}
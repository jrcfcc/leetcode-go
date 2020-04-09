package stack

import "testing"

func Test_validateStackSequences(t *testing.T) {
	var pushed = []int{1,2,3,4,5}
	var popped = []int{4,5,3,2,1}
	validateStackSequences(pushed,popped)
}
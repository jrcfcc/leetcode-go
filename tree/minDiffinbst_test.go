package tree

import "testing"

func Test_minDiffInBST(t *testing.T) {
	var left12 = &TreeNode{Val:12}
	var right49 = &TreeNode{Val:49}
	var right48 = &TreeNode{Val:48,Left:left12,Right:right49}
	var left0 = &TreeNode{Val:0}
	var root1 = &TreeNode{Val:   1, Left:  left0, Right: right48,}
	minDiffInBST(root1)
}
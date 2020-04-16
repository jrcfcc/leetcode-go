package tree

/*
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
*/
func sortedArrayToBST(nums []int) *TreeNode {
	var l = len(nums)
	if l == 0 {
		return nil
	}
	return BST(nums,0,l-1)
}

func BST(nums []int,start,end int) *TreeNode {
	if start > end {
		return nil
	}
	var mid = start + ((end - start + 1) >> 1)
	node := &TreeNode{Val:nums[mid]}
	node.Left = BST(nums,start,mid -1)
	node.Right = BST(nums,mid + 1,end)
	return node
}

package tree


/*
给定一个二叉搜索树的根结点 root，返回树中任意两节点的差的最小值。
*/
func minDiffInBST(root *TreeNode) int {
	res = make([]int,0)
	dfs(root)
	var ans = int(^uint32(0) >> 1)
	for i:=len(res)-1;i>=1;i--{
		diff := res[i] - res[i-1]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	res = append(res,root.Val)
	dfs(root.Right)
}

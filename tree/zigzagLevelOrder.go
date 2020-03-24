package tree

/*
题目:二叉树的锯齿形层次遍历
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

解题思路:
1、广度优先搜索
2、遇到index为奇数的行，翻转填充
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	if nil == root {
		return arr
	}
	//模拟队列,存储一层节点
	queue := []*TreeNode{root}
	var isLeftOrder = true
	for len(queue) != 0 {
		//记录本层的遍历值
		curr := make([]int, len(queue))
		for i := 0; i < len(curr); i++ {
			//从左到右广度遍历
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			//如果是偶数层,反转一次,将遍历顺序改成从右到左
			if isLeftOrder {
				curr[i] = queue[0].Val
			}else{
				//reverse
				curr[len(curr)-1-i] = queue[0].Val
			}
			//出队
			queue = queue[1:]
		}
		isLeftOrder = !isLeftOrder
		arr = append(arr, curr)
	}
	//fmt.Println(arr)
	return arr
}

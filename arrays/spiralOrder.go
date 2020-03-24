package arrays

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

解法:
这里的方法不需要记录已经走过的路径，所以执行用时和内存消耗都相对较小

首先设定上下左右边界
其次向右移动到最右，此时第一行因为已经使用过了，可以将其从图中删去，体现在代码中就是重新定义上边界
判断若重新定义后，上下边界交错，表明螺旋矩阵遍历结束，跳出循环，返回答案
若上下边界不交错，则遍历还未结束，接着向下向左向上移动，操作过程与第一，二步同理
不断循环以上步骤，直到某两条边界交错，跳出循环，返回答案
*/
func spiralOrder(matrix [][]int) []int {
	var row = len(matrix)
	if row == 0 {
		return nil
	}
	var col = len(matrix[0])
	var u,d,l,r = 0,row-1,0,col-1
	var res = make([]int,0,row*col)
	for {
		//从左边界向右
		for i:=l;i<=r;i++{
			res = append(res,matrix[u][i])
		}
		//上边界下移
		u++
		if u > d {
			break
		}
		//上边界向下
		for i:=u;i<=d;i++{
			res = append(res,matrix[i][r])
		}
		//右边界左移
		r--
		if l > r {
			break
		}
		//从右边界向左
		for i:=r;i>=l;i--{
			res = append(res,matrix[d][i])
		}
		//下边界上移
		d--
		if u > d {
			break
		}
		//从下边界向上
		for i:=d;i>=u;i--{
			res = append(res,matrix[i][l])
		}
		//左边界右移
		l++
		if l > r {
			break
		}
	}
	return res
}

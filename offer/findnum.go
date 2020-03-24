package offer


/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，
输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

解题思路:从matrix[0][n] 开始扫描,如果大于target,就列左移
如果小于target,就行下移
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	var row = len(matrix)
	if row == 0 {
		return false
	}
	var col = len(matrix[0])
	var i,j = 0,col-1
	for i < row && j >= 0 {
		if matrix[i][j] > target {
			j--
		}else if matrix[i][j] < target {
			i++
		}else{
			return true
		}
	}
	return false
}

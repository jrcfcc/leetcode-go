package dynamic


/*
题目描述:在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

解题思路:动态规划
我们用 0 初始化另一个矩阵 dp，维数和原始矩阵维数相同；
dp(i,j) 表示的是由 1 组成的最大正方形的边长；
从 (0,0)开始，对原始矩阵中的每一个 1，我们将当前元素的值更新为
dp(i, j)=min(dp(i−1, j), dp(i−1, j−1), dp(i, j−1))+1

*/
func maximalSquare(matrix [][]byte) int {
	var row = len(matrix)
	if row == 0 {
		return 0
	}
	var col = len(matrix[0])
	var dp = make([][]int,row+1)
	for i:=0;i<=row;i++{
		dp[i] = make([]int,col+1)
	}
	var maxsqlen = 0
	for i:=1;i<=row;i++{
		for j:=1;j<=col;j++{
			if matrix[i-1][j-1] == '1' {
				min := dp[i-1][j]
				if dp[i-1][j-1] < min {
					min = dp[i-1][j-1]
				}
				if dp[i][j-1] < min {
					min = dp[i][j-1]
				}
				dp[i][j] = min + 1
				if maxsqlen < dp[i][j] {
					maxsqlen = dp[i][j]
				}
			}
		}
	}
	return maxsqlen * maxsqlen
}

package dynamic

/*
题目:给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

解题:动态规划
base case dp[0][0] = triangle[0][0]
dp[i][j] 表示第i行,第j列元素的最小路径和
最小路径和等于上一层的相邻的两个元素的最小路径和加上当前元素
dp[i][j] = min(dp[i-1][j],dp[i-1][j-1]) + triangle[i][j]
*/
const MAX = 1 << 31 - 1
func minimumTotal(triangle [][]int) int {
	var row = len(triangle)
	if row == 0 {
		return 0
	}
	var dp = make([][]int,row)
	for i:=0;i<row;i++{
		dp[i] = make([]int,i+1)
	}
	dp[0][0] = triangle[0][0]
	var minTotal = MAX
	for i:=1;i<row;i++{
		for j:=0;j<=i;j++{
			if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}else if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			}else{
				var min = dp[i-1][j]
				if dp[i-1][j-1] < min {
					min = dp[i-1][j-1]
				}
				dp[i][j] = min + triangle[i][j]
			}
		}
	}
	for i:=0;i<row;i++{
		if minTotal > dp[row-1][i]{
			minTotal = dp[row-1][i]
		}
	}
	return minTotal
}

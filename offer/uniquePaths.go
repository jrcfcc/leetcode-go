package offer

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？

动态规划:
dp[i][j] 表示从(0,0)到i,j的最多路径
状态转移
dp[i][j] 等于从它的上方到达的路径加从它的左方到达的路径
dp[i][j] = dp[i-1][j] + dp[i][j-1])
*/

func uniquePaths(m int, n int) int {
	var dp = make([][]int,m)

	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
		dp[i][0] =1
	}
	for i:=0;i<n;i++{
		dp[0][i] =1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

package dynamic

/*
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符

动态规划
dp[i,j] 表示s1[0..i]到s2[0...j]的最短编辑距离
如果s1[i] == s2[j],不需要编辑,也就是dp[i][j] = dp[i-1][j-1]
如果s1[i] != s2[j],那么dp[i][j]就等于插入,删除,替换操作中的最小值+1
dp[i][j-1]在s2[j]前插入一个字符使得相等,j前移
dp[i-1][j]删除s1[i]使得两者相等,i前移
dp[i-1][j-1],直接替换其中一个字符,i,j都继续前移
dp[i,j] = min(dp[i][j-1],dp[i][j],dp[i][j]) + 1

base case
dp[i][0] = i,dp[0][j] = j
*/
func minDistance(word1 string, word2 string) int {
	var l1,l2 = len(word1),len(word2)
	var dp = make([][]int,l1+1)
	for i:= 0;i<=l1;i++{
		dp[i] = make([]int,l2+1)
		dp[i][0] = i
	}
	for j:=0;j<=l2;j++{
		dp[0][j] = j
	}
	for i:=1;i<=l1;i++{
		for j:=1;j<=l2;j++{
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				var min = dp[i][j-1]
				if dp[i-1][j] < min {
					min = dp[i-1][j]
				}
				if dp[i-1][j-1] < min {
					min = dp[i-1][j-1]
				}
				dp[i][j] = min + 1
			}
		}
	}
	return dp[l1][l2]
}

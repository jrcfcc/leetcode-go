package dynamic


/*
寻找最长回文子串
动态规划
dp[i,j]表示si->sj是否是一个回文子串
dp[i,j] = dp[i+1,j-1] && si == sj
base case
单个回文
dp[i,i] = true
两个字符的回文
dp[i,i+1] && si == si+1
*/
func longestPalindrome(s string) string {
	var l = len(s)
	if l <= 1 {
		return s
	}
	var dp = make([][]bool,l)
	for i:=0;i<l;i++{
		dp[i] = make([]bool,l)
		dp[i][i] = true
	}
	var max,start = 1,0
	for j:=1;j<l;j++{
		for i:=0;i<j;i++{
			if s[i] == s[j] {
				if j-i<3{
					dp[i][j] = true
				}else{
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] {
				if max < j-i+1 {
					max = j-i+1
					start = i
				}
			}
		}
	}
	return s[start:start+max]
}

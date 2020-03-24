package change


/*

*/
func waysToChange(n int) int {
	var coins = []int{1,5,10,25}
	var dp = make([]int,n+1)
	dp[0] = 0
	for i:=1;i< len(dp);i++{
		dp[i] = n + 1
	}
	for i:=1;i<len(dp);i++{
		for _,coin := range coins {
			if i - coin < 0 {
				continue
			}
			if dp[i] > dp[i-1] + 1 {
				dp[i] =  dp[i-1] + 1
			}
		}
	}
	return dp[n]%1000000007
}

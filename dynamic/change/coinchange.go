package change


/*
凑零钱:
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。

解题思路:动态规划
状态,amount
最优独立子结构:要知道凑成amount最少需要多少个硬币,只需要知道凑成amount-1需要的最少硬币数
也就是dp[n] 表示的是,凑成金额n需要的最小硬币数
base case 也就是结束递归的条件,就是凑成0 需要0个硬币
dp[n] = 0  n = 0
dp[n] = -1 n < 0
dp[n] = min(dp[n-coin]) + 1 coin in coins

*/
func coinChange(coins []int, amount int) int {
	var dp = make([]int,amount + 1)
	for k,_ := range dp {
		dp[k] = amount + 1
	}
	dp[0] = 0
	for i:=1;i<len(dp);i++{
		for _,coin := range coins {
			//该硬币无法凑出需要的金额,直接返回
			if i - coin < 0 {
				continue
			}
			if dp[i] > 1 + dp[i-coin] {
				dp[i] = 1 + dp[i-coin]
			}
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

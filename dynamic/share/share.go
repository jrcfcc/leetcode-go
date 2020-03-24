package share

/*
题目:给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票），
设计一个算法来计算你所能获取的最大利润。
注意你不能在买入股票前卖出股票。

状态:
这个问题的「状态」有三个，第一个是天数，第二个是允许交易的最大次数，
第三个是当前的持有状态（即之前说的 rest 的状态，我们不妨用 1 表示持有，0 表示没有持有）。
然后我们用一个三维数组就可以装下这几种状态的全部组合
dp[i][k][0 or 1]
0 <= i <= n-1, 1 <= k <= K
n 为天数，大 K 为最多交易数
此问题共 n × K × 2 种状态，全部穷举就能搞定。
for 0 <= i < n:
    for 1 <= k <= K:
        for s in {0, 1}:
            dp[i][k][s] = max(buy, sell, rest)
我们想求的最终答案是 dp[n - 1][K][0]，即最后一天，最多允许 K 次交易，最多获得多少利润。

本题中,只允许一次交易,所以k=1,可以直接简化
状态转移方程:
dp[i][0] = max(dp[i-1][0],dp[i-1][1] + price[i])
今天没有持有股票,可能是昨天就没有,今天什么也没做,或者是昨天持有了,今天卖了两种可能,
取两种可能利益最大的值
dp[i][1] = max(dp[i-1][1],dp[i-1][0] - price[i])
今天持有股票,可能是昨天也持有,今天啥也没做,或者昨天没有持有,但是今天买了两种可能
base case
交易日从0开始
dp[-1][0] = 0 还没开始,利润为0
dp[-1][1] = 负无穷, 因为还没有开始买卖,不可能持有股票
dp[0][0] = 0 没有持有股票,利润为0
*/
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func maxProfit(prices []int) int {
	var n = len(prices)
	if n <= 0 {
		return 0
	}
	//dp[-1][0] = 0  dp[-1][1] = INT_MIN -1表示还未开始
	var dp_i_0,dp_i_1 = 0,INT_MIN
	var temp = dp_i_0
	for i:=0;i<n;i++ {
		temp = dp_i_0
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp_i_0 = max(dp_i_0,dp_i_1 + prices[i])
		// dp[i][1] = max(dp[i-1][1],dp[i-1][0] -prices[i])
		dp_i_1 = max(dp_i_1,temp - prices[i])
	}
	return dp_i_0
}

func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}


/*
限制只能买卖一次股票
dp[i][0] = max(dp[i-1][0],dp[i-1][1] + price[i])
今天没有持有股票,可能是昨天就没有,今天什么也没做,或者是昨天持有了,今天卖了两种可能,
取两种可能利益最大的值
dp[i][1] = max(dp[i-1][1],dp[i-1][0] - price[i])
今天持有股票,可能是昨天持有,今天啥也没做,或者昨天没有持有,但是今天买了两种可能
*/
func maxProfit2(prices []int) int {
	var n = len(prices)
	if n <= 0 {
		return 0
	}
	//dp[-1][0] = 0  dp[-1][1] = INT_MIN -1表示还未开始
	var dp_i_0,dp_i_1 = 0,INT_MIN
	for i:=0;i<n;i++ {
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp_i_0 = max(dp_i_0,dp_i_1 + prices[i])
		// dp[i][1] = max(dp[i-1][1], -prices[i])
		dp_i_1 = max(dp_i_1,- prices[i])
	}
	return dp_i_0
}
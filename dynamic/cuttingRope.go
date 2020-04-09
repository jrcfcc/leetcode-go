package dynamic

import (
	"math/big"
)

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m] 。
请问 k[0]*k[1]*...*k[m] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

动态规划:
状态 : 要切分绳子的长度n
dp[n] : 该长度n的绳子切分m段后各段长度的乘积
base case： dp[2] = 1
状态转移方程： dp[n] = max(dp[n],max( dp[n-len])*len, (n-len)*len );
*/
func cuttingRope(n int) int {
	var dp = make([]int,n)
	dp[1] = 1
	dp[2] = 1
	for i:=3;i<=n;i++{
		for j:=1;j<i;j++{

			dp[i] = dp[i-j]*j
			if (i-j) * j > dp[i] {
				dp[i] = (i-j) * j
			}
		}
	}
	return dp[n]
}

/*
当 n <= 3 时，按照贪心规则应直接保留原数字，但由于题目要求必须拆分，因此必须拆出一个 1，即直接返回 n - 1；
求 n 除以 3 的整数部分 a 和余数部分 b；
当 b == 0 时，直接返回 3^a；
当 b == 1 时，要将一个 1 + 3 转换为 2+2，此时返回 3^{a-1} * 4
当 b == 2时，返回 3^a * 2
*/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n-1
	}
	var a,b = n/3,n%3
	if b == 0 {
		x := pow(big.NewInt(3),a)
		return int(x.Mod(x,big.NewInt(1000000007)).Int64())
	}
	if b == 1 {
		x := pow(big.NewInt(3),a-1)
		x.Mul(x,big.NewInt(4))
		return int(x.Mod(x,big.NewInt(1000000007)).Int64())
	}
	x := pow(big.NewInt(3),a)
	x.Mul(x,big.NewInt(2))
	return int(x.Mod(x,big.NewInt(1000000007)).Int64())
}

func pow(x *big.Int,y int) *big.Int {
	if y == 0 {
		return big.NewInt(1)
	}
	var res = big.NewInt(0).Set(x)
	for i:=y;i>1;i--{
		res = big.NewInt(0).Mul(x,res)
	}
	return res
}

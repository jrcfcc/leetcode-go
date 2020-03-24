package dynamic

/*
题目描述:给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素）
返回其最大和。

类似的，我们设计动态规划算法，不是需要一个 dp 数组吗？
我们可以假设 $dp[0...i-1]$ 都已经被算出来了，然后问自己：怎么通过这些结果算出 dp[i]？
直接拿最长递增子序列这个问题举例你就明白了。不过，首先要定义清楚 dp 数组的含义，
即 dp[i] 的值到底代表着什么？

解法二、动态规划
按照排列组合的数学算法，9 个数字，以第 i 个数字结尾的串，有 i 种组合，一共有45个组合

(sum_{i=1} ^9 i )(∑i=19i)，如果有n个数字，时间复杂度是$O(n^2)，这样的时间复杂度是明显不能接受的。

我们把目光落到动态规划上面来，首先需要把这个问题分解成最优子问题来解。最主要的思路就是将上面的45个组合进行
，分解成数量较少的几个子问题。在这里我们一共有9个数字，顺理成章的我们把组合分解成9个小组的组合。

第一个子组合是以第一个数字结尾的连续序列，也就是 [-2]，最大值-2

第二个子组合是以第二个数字结尾的连续序列，也就是 [-2,1], [1]，最大值1

第三个子组合是以第三个数字结尾的连续序列，也就是 [-2,1,3], [1,3], [3]，最大值4

以此类推。。。

如果我们能够得到每一个子组合的最优解，也就是子序列的最大值，整体的最大值就可以通过比较这9个子组合的最大值来得到了。现在我们找到了最优子问题，重叠子问题在哪呢？那就得细心比较一下每个子问题。

从第二个子组合和第三个子组合可以看到，组合 3 只是在组合 2 的基础上每一个数组后面添加第 3 个数字，也就是数字 3，然后增加一个只有第三个数字的数组 [3] 。这样两个组合之间的关系就出现了，可是我们不关心这个序列是怎么生成的，只是关心最大值之间的关系。

下面我们看组合 3 的组成，我们将子组合 3 分成两种情况：
继承子组合二得到的序列，也就是[-2,1,3], [1,3] （最大值 1 = 第二个组合的最大值 + 第三个数字）
单独第三个数字的序列，也就是[3] （最大值 2 = 第三个数字）
如果 第二个序列的最大值 大于0，那么最大值 1 就比最大值 2 要大，反之最大值 2 较大。这样，我们就通过第二个组合的最大值和第三个数字，就得到了第三个组合的最大值。因为第二个组合的结果被重复用到了，所以符合这个重叠子问题的定义。通俗来讲这个问题就变成了，第 i 个子组合的最大值可以通过第i-1个子组合的最大值和第 i 个数字获得，如果第 i-1 个子组合的最大值没法给第 i 个数字带来正增益，我们就抛弃掉前面的子组合，自己就是最大的了。

步骤一、定义状态 -> 定义数组元素的含义
定义 dp[i] 为以 i 结尾子串的最大值
步骤二、状态转移方程 -> 找出数组元素间的关系式

​


步骤三、初始化 -> 找出初始条件
dp[0] = nums[0];
步骤四、状态压缩 -> 优化数组空间
每次状态的更新只依赖于前一个状态，就是说 dp[i] 的更新只取决于 dp[i-1] , 我们只用一个存储空间保存上一次的状态即可。
步骤五、选出结果
有的题目结果是 dp[i] 。
本题结果是 dp[0]...dp[i] 中最大值。
*/
const MIN = ^int(^uint(0) >> 1)
func maxSubArray(nums []int) int {
	var max = MIN
	var sum = 0
	for _,v := range nums {
		if sum > 0 {
			sum += v
		}else{
			sum = v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
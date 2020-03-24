package strchar

/*
判断数字是否是回文数字

第二个想法是将数字本身反转，然后将反转后的数字与原始数字进行比较，
如果它们是相同的，那么这个数字就是回文。
但是，如果反转后的数字大于 int.MAX，我们将遇到整数溢出问题。

按照第二个想法，为了避免数字反转可能导致的溢出问题，为什么不考虑只反转 int 数字的一半？
毕竟，如果该数字是回文，其后半部分反转后应该与原始数字的前半部分相同。
例如，输入 1221，我们可以将数字 “1221” 的后半部分从 “21” 反转为 “12”，
并将其与前半部分 “12” 进行比较，因为二者相同，我们得知数字 1221 是回文。
让我们看看如何将这个想法转化为一个算法。

算法

首先，我们应该处理一些临界情况。所有负数都不可能是回文，例如：-123不是回文因为 -不等于 3。所以我们可以对所有负数返回 false。

现在，让我们来考虑如何反转后半部分的数字。
对于数字 1221，如果执行 1221 % 10，我们将得到最后一位数字 1，
要得到倒数第二位数字，我们可以先通过除以 10 把最后一位数字从 1221 中移除，
1221 / 10 = 122，再求出上一步结果除以 10 的余数，122 % 10 = 2，
就可以得到倒数第二位数字。如果我们把最后一位数字乘以 10，再加上倒数第二位数字，
1 * 10 + 2 = 12，就得到了我们想要的反转后的数字。如果继续这个过程，
我们将得到更多位数的反转数字。

现在的问题是，我们如何知道反转数字的位数已经达到原始数字位数的一半？

我们将原始数字除以 10，然后给反转后的数字乘上 10，所以，
当原始数字小于反转后的数字时，就意味着我们已经处理了一半位数的数字。

*/
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	var reverse = 0
	for x  > reverse {
		var last = x % 10
		x = x / 10
		reverse = reverse * 10 + last
	}
	//左右相等,或者原数字有奇数位去除最后一位后相等
	if x == reverse || x  == reverse / 10  {
		return true
	}
	return false
}


/*
判断字符串是否是回文
*/
func isPalindromeStr(s string) bool {
	var l = len(s)
	if l <= 1 {
		return true
	}
	var head,tail = 0,l-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

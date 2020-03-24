package tuozhan

/*
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	//开始二分查找
	var left,right = 1,x>>1
	var pow,mid = 0,0
	for left <= right {
		mid = left + (right-left)>>1
		pow = mid * mid
		if pow > x {
			right = mid - 1
		}
		if pow < x {
			if (mid + 1) * (mid + 1) > x {
				return mid
			}
			left = mid + 1
		}
		if pow == x {
			return mid
		}
	}
	return 0
}

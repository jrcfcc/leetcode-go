package arrays

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
边界移动
*/
func generateMatrix(n int) [][]int {
	var u,d,l,r = 0,n-1,0,n-1
	var res = make([][]int,n)
	for i:=0;i<n;i++{
		res[i] = make([]int,n)
	}
	var num = 1
	for {
		for i:=l;i<=r;i++{
			res[u][i] = num
			num++
		}
		u++
		if u > d {
			break
		}
		for i:=u;i<=d;i++{
			res[i][r] = num
			num++
		}
		r--
		if l > r {
			break
		}
		for i:=r;i>=l;i--{
			res[d][i] = num
			num++
		}
		d--
		if u > d {
			break
		}
		for i:=d;i>=u;i--{
			res[i][l] = num
			num++
		}
		l++
		if l > r {
			break
		}
	}
	return res
}

package matrix

/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，
因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

解题思路:
算法原理：
深度优先搜索： 可以理解为暴力法遍历矩阵中所有字符串可能性。DFS 通过递归，
先朝一个方向搜到底，再回溯至上个节点，沿另一个方向搜索，以此类推。
剪枝： 在搜索中，遇到 这条路不可能和目标字符串匹配成功 的情况
（例如：此矩阵元素和目标字符不同、此元素已被访问），则应立即返回，称之为 可行性剪枝 。
算法剖析：
递归参数： 当前元素在矩阵 board 中的行列索引 i 和 j ，当前目标字符在 word 中的索引 k 。
终止条件：
返回 false ： ① 行或列索引越界 或 ② 当前矩阵元素与目标字符不同 或 ③ 当前矩阵元素已访问过 （③ 可合并至 ② ） 。
返回 true ： 字符串 word 已全部匹配，即 k = len(word) - 1 。
递推工作：
标记当前矩阵元素： 将 board[i][j] 值暂存于变量 tmp ，并修改为字符 '/' ，代表此元素已访问过，防止之后搜索时重复访问。
搜索下一单元格： 朝当前元素的 上、下、左、右 四个方向开启下层递归，使用 或 连接 （代表只需一条可行路径） ，并记录结果至 res 。
还原当前矩阵元素： 将 tmp 暂存值还原至 board[i][j] 元素。
回溯返回值： 返回 res ，代表是否搜索到目标字符串。
*/
var row,col,l int
var s string
func exist(board [][]byte, word string) bool {
	row = len(board)
	col = len(board[0])
	l = len(word)
	s = word
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if searchword(board,i,j,0) {
				return true
			}
		}
	}
	return false
}

func searchword(board [][]byte,i,j,k int) bool {
	if i < 0 || i > row - 1 || j < 0 || j > col - 1 {
		return  false
	}
	if board[i][j] != s[k] {
		return false
	}
	if k == l - 1 {
		return true
	}
	var tmp = board[i][j]
	board[i][j] = '/'

	res := searchword(board,i-1,j,k+1) ||
		searchword(board,i+1,j,k+1) ||
		searchword(board,i,j-1,k+1) ||
		searchword(board,i,j+1,k+1)
	board[i][j] = tmp
	return res
}

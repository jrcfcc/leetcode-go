package arrays



/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个岛屿是由四个方向
(水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

解题,递归查找
两层循环处理,从值为1的点开始向四方递归查找
找到后,将1置为0,相当于标识出此节点已查找过
*/
var long,width int
var gridCopy [][]int
var area int
func maxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	var maxArea = 0
	gridCopy = grid
	long = len(grid)
	width = len(grid[0])
	for i:=0;i<long;i++{
		for j:=0;j<width;j++{
			if grid[i][j] == 1 {
				area = 0
				getArea(i,j)
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func getArea(i,j int) {
	if i < 0 || j < 0 || i >= long || j >= width {
		return
	}
	if gridCopy[i][j] == 0  {
		return
	}
	gridCopy[i][j] = 0
	area++
	getArea(i+1,j)
	getArea(i-1,j)
	getArea(i,j+1)
	getArea(i,j-1)
}



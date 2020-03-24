package tuozhan

import "strconv"

/*
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。
请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
*/
type state struct {
	x int
	y int
}

func canMeasureWater(x int, y int, z int) bool {
	if x + y < z || z < 0 {
		return false
	}
	if z == 0 {
		return true
	}
	var set = make(map[string]struct{})
	var quene = make([]state,0,100)
	var initState = state{}
	quene = append(quene,initState)
	set[strconv.Itoa(0)+"_"+strconv.Itoa(0)] = struct{}{}
	//使用广度优先搜索
	for len(quene) > 0 {
		var head = quene[0]
		quene = quene[1:]
		//终止条件
		if head.x == z || head.y == z || head.x + head.y == z {
			return true
		}
		//从当前状态获取所有可能的下一步状态
		nextState := getNextState(head.x,head.y,x,y)
		for _,state := range nextState {
			//没有访问过的状态,加到队列中
			if _,ok:=set[strconv.Itoa(state.x)+"_"+strconv.Itoa(state.y)];!ok {
				quene = append(quene,state)
				set[strconv.Itoa(state.x)+"_"+strconv.Itoa(state.y)] = struct{}{}
			}
		}
	}
	return false
}

func getNextState(currX,currY,x,y int) []state{
	var nextState = make([]state,0,8)
	//获取所有可能的状态
	//将x装满
	if currX < x {
		nextState = append(nextState,state{x: x,y:currY})
	}
	//将y装满
	if currY < y {
		nextState = append(nextState,state{x:currX,y: y,})
	}
	//将x倒空
	if currX > 0 && currY !=0 {
		nextState = append(nextState,state{x: 0,y:currY})
	}
	//将y倒空
	if currY > 0 && currX !=0{
		nextState = append(nextState,state{x: currX,y:0})
	}
	//将x倒到y,y倒满,x有剩
	if currX > y - currY {
		nextState = append(nextState,state{x: currX-(y-currX),y:y})
	}
	//将x倒到y,x倒空,y未满
	if currX <= y - currY {
		nextState = append(nextState,state{x: 0,y:currY + currX})
	}
	//将y倒到x,x倒满,y有剩
	if currY > x - currX {
		nextState = append(nextState,state{x: x,y:currY-(x-currX)})
	}
	//将y倒到x,y倒空,x未满
	if currY <= x - currX {
		nextState = append(nextState,state{x: currY + currX,y:0})
	}
	return  nextState
}

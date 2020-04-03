package matrix


/*
班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。
如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。
所谓的朋友圈，是指所有朋友的集合。
给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，
表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。
你必须输出所有学生中的已知的朋友圈总数。

解题:并查集
*/

/*
并查集:
isConnected:父节点一样则说明是联通的
union:联通两个节点,即将其中一个树的层高比较矮的根节点指向层高高的根节点
find : 查找根节点,每次找到后,直接将当前节点指向根节点,降低层高,增加根节点的子节点个数
*/
var parent []int  //parent[i] 表示i这个节点的父节点索引 ,当parent[i] = i时,表明找到了根节点
var rank []int   //rank[i] 表示以i作为根节点的联通树的层数
var count = 0
func findCircleNum(M [][]int) int {
	var length = len(M)
	count = length
	parent = make([]int,length)
	rank = make([]int,length)
	for i:=0;i<length;i++{
		//parent全部初始为-1,这样当最后所有元素合并完后,还为-1的就是根节点,只需要统计根节点的个数就能知道朋友圈的个数
		parent[i],rank[i] = i,1
	}
	for i:=0;i<length;i++{
		for j:=0;j<i;j++{
			if M[i][j] == 1 {
				union(i,j)
			}
		}
	}
	return count
}

//判断两个元素是否联通
func isConnected(p,q int) bool {
	//如果两个元素的根节点相同,说明他们是联通的
	if find(p) == find(q) {
		return true
	}
	return false
}

//查找两个元素的根节点
func find(p int) int {
	//跳出循环时说明找到了根节点
	var queryIndex = p
	for parent[p] != p {
		p = parent[p]
	}
	parent[queryIndex] = p
	return p
}

//联通两个节点,层级低的指向层级高的
func union(p,q int) {
	pRoot,qRoot := find(p),find(q)
	if pRoot == qRoot {
		return
	}

	//低层级的根节点指向高层级节点的根节点
	if rank[p] <= rank[q] {
		parent[pRoot] = qRoot
		rank[q] += 1
	} else{
		//两边层级相等,随便指向,被指向的树层高+1
		parent[qRoot] = pRoot
		rank[p] += 1
	}
	count--
}
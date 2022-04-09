package other

/*
	循环依赖检测
	现有n个编译项，编号为0 ~ n-1。给定一个二维数组，表示编译项之间有依赖关系。如[0, 1]表示1依赖于0。
*/

// 拓扑排序
// 拓扑排序算法过程：
//
//1、选择图中一个入度为0的点，记录下来
//2、在图中删除该点和所有以它为起点的边
//3、重复1和2，直到图为空或没有入度为0的点。

// 使用的是邻接表 、用inDegree记录每个节点的入度、使用queue进行BSF遍历图。
// 把入度为0的加入到队列，队列的顺序其实就是拓扑排序。然后从队列取出加入队列中。
func haveCircularDependency(prerequisites [][]int, n int) []int {
	graph := make([][]int, n)
	res := make([]int, 0)
	inDegree := make([]int, n)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		if graph[a] == nil {
			graph[a] = make([]int, 0)
		}
		graph[a] = append(graph[a], b)
		inDegree[b]++
	}
	queue := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		res = append(res, t)
		for i := 0; i < len(graph[t]); i++ {
			v := graph[t][i]
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(res) == n {
		return res
	}
	return nil
}

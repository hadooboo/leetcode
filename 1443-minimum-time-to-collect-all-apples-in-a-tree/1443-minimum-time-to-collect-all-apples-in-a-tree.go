type Node struct {
	Val      int
	Children []*Node
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	nodes := make([]*Node, n)
	for i := range nodes {
		nodes[i] = &Node{
			Val:      i,
			Children: make([]*Node, 0),
		}
	}
	for _, edge := range edges {
		nodes[edge[0]].Children = append(nodes[edge[0]].Children, nodes[edge[1]])
		nodes[edge[1]].Children = append(nodes[edge[1]].Children, nodes[edge[0]])
	}

	res, _ := doBFS(nodes, hasApple, -1, 0)
	return res
}

func doBFS(nodes []*Node, hasApple []bool, prev, curr int) (int, bool) {
	res := 0
	for _, child := range nodes[curr].Children {
		if child.Val == prev {
			continue
		}
		r, ok := doBFS(nodes, hasApple, curr, child.Val)
		if ok {
			res += r + 2
		}
	}
	if res == 0 {
		return 0, hasApple[curr]
	}

	return res, true
}
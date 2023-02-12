func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for _, v := range roads {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}

	var res int64 = 0
	for _, v := range adj[0] {
		r, _ := doBFS(adj, v, 0, seats)
		res += r
	}
	return res
}

func doBFS(adj [][]int, curr, prev int, seats int) (int64, int) {
	var res int64 = 0
	var cnt int = 1
	for _, v := range adj[curr] {
		if v == prev {
			continue
		}
		r, s := doBFS(adj, v, curr, seats)
		res += r
		cnt += s
	}
	return res + int64((cnt+seats-1)/seats), cnt
}
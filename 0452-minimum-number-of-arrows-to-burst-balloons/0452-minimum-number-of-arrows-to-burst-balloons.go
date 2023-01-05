import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res, p := 1, points[0][1]

	for i := 1; i < len(points); i++ {
		if p < points[i][0] {
			p = points[i][1]
			res++
		}
	}

	return res
}
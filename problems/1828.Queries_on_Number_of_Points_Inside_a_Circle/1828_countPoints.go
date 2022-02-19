package problems

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, 0)
	// each query
	for _, q := range queries {
		count := 0
		// each point
		for _, p := range points {
			// ( (xj - xi) + (yj - yi) ) ^ 1/2
			d1 := float64(p[0] - q[0])
			d2 := float64(p[1] - q[1])
			r := float64(q[2])
			if math.Sqrt(math.Pow(d1, 2)+math.Pow(d2, 2)) <= r {
				count++
			}
		}
		res = append(res, count)
	}

	return res
}

func countPoints2(points [][]int, queries [][]int) []int {
	res := make([]int, 0)
	// each query
	for _, q := range queries {
		count := 0
		// each point
		for _, p := range points {
			d1 := p[0] - q[0]
			d2 := p[1] - q[1]
			r := q[2] * q[2]
			if d1*d1+d2*d2 <= r {
				count++
			}
		}
		res = append(res, count)
	}

	return res
}

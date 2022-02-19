package problems

func fib(n int) int {
	F := map[int]int{}
	F[0], F[1] = 0, 1
	for i := 2; i <= n; i++ {
		F[i] = F[i-1] + F[i-2]
	}
	return F[n]
}

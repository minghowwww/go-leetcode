package getOffer

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b, c := 0, 0, 1
	for i := 2; i <= n; i++ {
		a = b
		b = c
		c = (a + b) % (1e9 + 7)
	}
	return c
}

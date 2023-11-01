package problem509

func fib(n int) int {
	if n < 2 {
		return n
	}

	f1, f2 := 1, 0
	for i := 2; i <= n; i++ {
		f1, f2 = f1+f2, f1
	}

	return f1
}

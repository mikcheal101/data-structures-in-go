package fibonacciseries


type FibonacciSeries struct {
	FibonacciNumber int
}

func (Fibonacci *FibonacciSeries) FibonacciConstructor(n int) {
	if n == 1 || n == 2 {
		Fibonacci.FibonacciNumber = 1
	} else {
		Fibonacci.FibonacciConstructor(n - 1)
		leftSide := Fibonacci.FibonacciNumber

		Fibonacci.FibonacciConstructor(n - 2)
		rightSide := Fibonacci.FibonacciNumber

		Fibonacci.FibonacciNumber = leftSide + rightSide
	}
}

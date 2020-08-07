package diffsquares

// SquareOfSum - the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares - the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference - the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

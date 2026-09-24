package phyllotaxis

import "math"

const (
	// GoldenRatio is phi = (1 + sqrt(5)) / 2
	GoldenRatio = 1.6180339887498948482

	// GoldenAngleDegrees is 360 * (1 - 1/phi) ~= 137.507764 deg
	GoldenAngleDegrees = 137.5077640500378546

	// GoldenAngleRadians is the golden angle in radians
	GoldenAngleRadians = GoldenAngleDegrees * math.Pi / 180.0
)

// Fibonacci returns the n-th Fibonacci number (0-indexed: F(0)=0, F(1)=1, F(2)=1, F(3)=2...)
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// IsFibonacciPair checks whether n and m are consecutive Fibonacci numbers
func IsFibonacciPair(n, m int) bool {
	if n > m {
		n, m = m, n
	}
	a, b := 1, 1
	for b < m {
		a, b = b, a+b
		if a == n && b == m {
			return true
		}
	}
	return false
}

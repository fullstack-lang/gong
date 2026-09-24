package phyllotaxis_test

import (
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/geom/phyllotaxis"
)

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i, exp := range expected {
		got := phyllotaxis.Fibonacci(i)
		if got != exp {
			t.Errorf("Fibonacci(%d) = %d, expected %d", i, got, exp)
		}
	}
}

func TestIsFibonacciPair(t *testing.T) {
	if !phyllotaxis.IsFibonacciPair(3, 5) {
		t.Errorf("expected (3, 5) to be a Fibonacci pair")
	}
	if !phyllotaxis.IsFibonacciPair(8, 13) {
		t.Errorf("expected (8, 13) to be a Fibonacci pair")
	}
	if phyllotaxis.IsFibonacciPair(4, 7) {
		t.Errorf("expected (4, 7) not to be a Fibonacci pair")
	}
}

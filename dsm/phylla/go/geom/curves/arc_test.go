package curves_test

import (
	"math"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/geom/curves"
)

func TestComputeArcCenterFromEndpoints(t *testing.T) {
	// Semicircle with radius 10 from (-10, 0) to (10, 0)
	cx, cy, r := curves.ComputeArcCenterFromEndpoints(-10, 0, 10, 0, 10, false, false)
	if math.Abs(cx) > 1e-4 {
		t.Errorf("expected cx=0, got %f", cx)
	}
	if math.Abs(r-10) > 1e-4 {
		t.Errorf("expected r=10, got %f", r)
	}
	_ = cy
}

func TestEvalArcY(t *testing.T) {
	// Top of circle centered at (0, 0) with radius 10: at x=0, y should be 10
	y := curves.EvalArcY(-10, 0, 10, 0, 0, 0, 10, 0)
	if math.Abs(y-10) > 1e-4 {
		t.Errorf("expected y=10 at x=0, got %f", y)
	}
}

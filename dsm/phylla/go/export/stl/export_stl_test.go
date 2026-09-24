package stl_test

import (
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/export/stl"
)

func TestWriteFacet(t *testing.T) {
	var sb strings.Builder
	v1 := stl.Vector3{X: 0, Y: 0, Z: 0}
	v2 := stl.Vector3{X: 1, Y: 0, Z: 0}
	v3 := stl.Vector3{X: 0, Y: 1, Z: 0}

	stl.WriteFacet(&sb, v1, v2, v3)
	out := sb.String()

	if !strings.Contains(out, "facet normal") {
		t.Errorf("expected 'facet normal' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "outer loop") || !strings.Contains(out, "endloop") {
		t.Errorf("expected outer loop in output, got:\n%s", out)
	}
}

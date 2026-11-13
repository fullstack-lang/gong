package models

import (
	"math"
	"testing"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func TestBuildSVGRectAnchoredPaths_StrokeIcon24(t *testing.T) {
	raw := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
		<title>Ignored title</title>
		<defs><path d="M0 0 L10 10"/></defs>
		<rect x="2" y="4" width="20" height="16" rx="2"/>
		<line x1="2" y1="10" x2="22" y2="10"/>
	</svg>`

	paths := buildSVGRectAnchoredPaths("Payment Gateway", raw, svg.RECT_TOP_LEFT, 10.0, 11.0, 18.0, "#757575")
	if len(paths) != 2 {
		t.Fatalf("expected 2 RectAnchoredPaths (title and defs skipped), got %d", len(paths))
	}

	wantScale := 18.0 / 24.0
	for i, p := range paths {
		if math.Abs(p.AppliedScaling-wantScale) > 1e-6 {
			t.Errorf("path[%d] AppliedScaling = %v, want %v", i, p.AppliedScaling, wantScale)
		}
		if math.Abs(p.X_Offset-10.0) > 1e-6 || math.Abs(p.Y_Offset-11.0) > 1e-6 {
			t.Errorf("path[%d] offsets = (%v, %v), want (10, 11)", i, p.X_Offset, p.Y_Offset)
		}
		if p.Color != "none" || p.FillOpacity != 0 {
			t.Errorf("path[%d] fill = (%q, %v), want (\"none\", 0)", i, p.Color, p.FillOpacity)
		}
		if p.Stroke != "#757575" || p.StrokeWidth != 2.0 || p.StrokeOpacity != 1.0 {
			t.Errorf("path[%d] stroke = (%q, width=%v, opacity=%v), want (\"#757575\", 2.0, 1.0)",
				i, p.Stroke, p.StrokeWidth, p.StrokeOpacity)
		}
		if p.Transform != "" {
			t.Errorf("path[%d] Transform = %q, want empty string for cross-renderer parity", i, p.Transform)
		}
		if p.RectAnchorType != svg.RECT_TOP_LEFT {
			t.Errorf("path[%d] RectAnchorType = %q, want %q", i, p.RectAnchorType, svg.RECT_TOP_LEFT)
		}
	}
}

func TestBuildSVGRectAnchoredPaths_MaterialSymbolsNegativeViewBox(t *testing.T) {
	raw := `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px">
		<path d="M160-160v-640h640v640H160Z"/>
	</svg>`

	paths := buildSVGRectAnchoredPaths("Participant", raw, svg.RECT_TOP_LEFT, 12.0, -42.0, 24.0, "#333333")
	if len(paths) != 1 {
		t.Fatalf("expected 1 path, got %d", len(paths))
	}

	p := paths[0]
	wantScale := 24.0 / 960.0 // 0.025
	if math.Abs(p.AppliedScaling-wantScale) > 1e-9 {
		t.Errorf("AppliedScaling = %v, want %v", p.AppliedScaling, wantScale)
	}
	// minY = -960 => yOffset = -42 - 0.025*(-960) = -18
	if math.Abs(p.X_Offset-12.0) > 1e-6 || math.Abs(p.Y_Offset-(-18.0)) > 1e-6 {
		t.Errorf("offsets = (%v, %v), want (12, -18)", p.X_Offset, p.Y_Offset)
	}
	if p.Color != "#333333" || p.FillOpacity != 1.0 {
		t.Errorf("fill = (%q, %v), want (\"#333333\", 1.0)", p.Color, p.FillOpacity)
	}
	if p.Stroke != "none" {
		t.Errorf("stroke = %q, want \"none\"", p.Stroke)
	}
}

func TestBuildSVGRectAnchoredPaths_MultiShapeGroupsAndHitboxSkip(t *testing.T) {
	raw := `<?xml version="1.0" encoding="UTF-8"?>
	<!-- Duotone icon with invisible hitbox -->
	<svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
		<path d="M0 0h24v24H0z" fill="none"/>
		<g fill="#1976D2" stroke="#0D47A1" stroke-width="1.5">
			<circle cx="12" cy="8" r="4" fill-opacity="0.2"/>
			<ellipse cx="12" cy="18" rx="7" ry="3" style="fill-opacity: 0.25; stroke-opacity: 0.9"/>
			<polygon points="10,2 14,2 12,5"/>
			<polyline points="8,18 11,20 16,16" fill="none"/>
		</g>
	</svg>`

	paths := buildSVGRectAnchoredPaths("Customer", raw, svg.RECT_LEFT, 10.0, -12.0, 24.0, "#333333")
	if len(paths) != 4 {
		t.Fatalf("expected 4 visible paths (invisible hitbox skipped), got %d", len(paths))
	}

	if paths[0].Color != "#1976D2" || math.Abs(paths[0].FillOpacity-0.2) > 1e-6 {
		t.Errorf("circle fill = (%q, %v), want (\"#1976D2\", 0.2)", paths[0].Color, paths[0].FillOpacity)
	}
	if paths[0].Stroke != "#0D47A1" || paths[0].StrokeOpacity != 1.0 {
		t.Errorf("circle stroke = (%q, %v), want (\"#0D47A1\", 1.0)", paths[0].Stroke, paths[0].StrokeOpacity)
	}
	if math.Abs(paths[1].FillOpacity-0.25) > 1e-6 || math.Abs(paths[1].StrokeOpacity-0.9) > 1e-6 {
		t.Errorf("ellipse inline style opacities = (fill=%v, stroke=%v), want (0.25, 0.9)",
			paths[1].FillOpacity, paths[1].StrokeOpacity)
	}
	if paths[3].Color != "none" || paths[3].Stroke != "#0D47A1" {
		t.Errorf("polyline fill/stroke = (%q, %q), want (\"none\", \"#0D47A1\")", paths[3].Color, paths[3].Stroke)
	}
}

func TestBuildSVGRectAnchoredPaths_BarePathFallback(t *testing.T) {
	paths := buildSVGRectAnchoredPaths("Data", "M 2 2 H 22 V 22 H 2 Z", svg.RECT_LEFT, 4.0, -8.0, 16.0, "#555555")
	if len(paths) != 1 {
		t.Fatalf("expected 1 path for bare path definition, got %d", len(paths))
	}
	wantScale := 16.0 / 24.0
	if math.Abs(paths[0].AppliedScaling-wantScale) > 1e-6 {
		t.Errorf("AppliedScaling = %v, want %v", paths[0].AppliedScaling, wantScale)
	}
}

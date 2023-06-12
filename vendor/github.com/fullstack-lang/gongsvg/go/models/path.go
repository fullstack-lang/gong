package models

type Path struct {
	Name string

	// Definition
	//
	//
	// d="M 10,30
	// A 20,20 0,0,1 50,30
	// A 20,20 0,0,1 90,30
	// Q 90,60 50,90
	// Q 10,60 10,30 z" />
	//
	// MoveTo: M, m
	// LineTo: L, l, H, h, V, v
	// Cubic Bézier Curve: C, c, S, s
	// Quadratic Bézier Curve: Q, q, T, t
	// Elliptical Arc Curve: A, a
	// ClosePath: Z, z
	Definition string
	Presentation
	Animates []*Animate
}

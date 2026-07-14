package models

type Triangle struct {
	Name string

	V1 int
	V2 int
	V3 int
}

type BufferGeometry struct {
	Name string

	Vertices []*Vector3
	Faces    []*Triangle
}

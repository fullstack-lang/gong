package models

import (
	"fmt"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) addQuad(geom *threejs.BufferGeometry, i, j int, isReverse bool, suffix string) {
	idx0 := i*4 + j
	idx1 := i*4 + j + 1
	idx2 := (i+1)*4 + j
	idx3 := (i+1)*4 + j + 1

	v1_t1, v2_t1, v3_t1 := idx0, idx1, idx2
	v1_t2, v2_t2, v3_t2 := idx1, idx3, idx2

	if isReverse {
		v2_t1, v3_t1 = v3_t1, v2_t1
		v2_t2, v3_t2 = v3_t2, v2_t2
	}

	t1 := (&threejs.Triangle{
		Name: fmt.Sprintf("T1 %d %d %s", i, j, suffix),
		V1:   v1_t1, V2: v2_t1, V3: v3_t1,
	}).Stage(stager.threejsStage)
	t2 := (&threejs.Triangle{
		Name: fmt.Sprintf("T2 %d %d %s", i, j, suffix),
		V1:   v1_t2, V2: v2_t2, V3: v3_t2,
	}).Stage(stager.threejsStage)

	geom.Faces = append(geom.Faces, t1, t2)
}

func (stager *Stager) addWallQuad(geomHoleWalls *threejs.BufferGeometry, v1_src, v2_src, v3_src, v4_src *threejs.Vector3, suffix string, reverseWinding bool) {
	p1 := (&threejs.Vector3{
		Name: fmt.Sprintf("%s %d %s", geomHoleWalls.Name, len(geomHoleWalls.Vertices), suffix),
		X:    v1_src.X, Y: v1_src.Y, Z: v1_src.Z,
	}).Stage(stager.threejsStage)
	p2 := (&threejs.Vector3{
		Name: fmt.Sprintf("%s %d %s", geomHoleWalls.Name, len(geomHoleWalls.Vertices)+1, suffix),
		X:    v2_src.X, Y: v2_src.Y, Z: v2_src.Z,
	}).Stage(stager.threejsStage)
	p3 := (&threejs.Vector3{
		Name: fmt.Sprintf("%s %d %s", geomHoleWalls.Name, len(geomHoleWalls.Vertices)+2, suffix),
		X:    v3_src.X, Y: v3_src.Y, Z: v3_src.Z,
	}).Stage(stager.threejsStage)
	p4 := (&threejs.Vector3{
		Name: fmt.Sprintf("%s %d %s", geomHoleWalls.Name, len(geomHoleWalls.Vertices)+3, suffix),
		X:    v4_src.X, Y: v4_src.Y, Z: v4_src.Z,
	}).Stage(stager.threejsStage)

	idx := len(geomHoleWalls.Vertices)
	geomHoleWalls.Vertices = append(geomHoleWalls.Vertices, p1, p2, p3, p4)

	v1_t1, v2_t1, v3_t1 := idx, idx+1, idx+2
	v1_t2, v2_t2, v3_t2 := idx+1, idx+3, idx+2

	if reverseWinding {
		v2_t1, v3_t1 = v3_t1, v2_t1
		v2_t2, v3_t2 = v3_t2, v2_t2
	}

	t1 := (&threejs.Triangle{
		Name: fmt.Sprintf("T1 %d %s", len(geomHoleWalls.Faces), suffix),
		V1:   v1_t1, V2: v2_t1, V3: v3_t1,
	}).Stage(stager.threejsStage)

	t2 := (&threejs.Triangle{
		Name: fmt.Sprintf("T2 %d %s", len(geomHoleWalls.Faces), suffix),
		V1:   v1_t2, V2: v2_t2, V3: v3_t2,
	}).Stage(stager.threejsStage)

	geomHoleWalls.Faces = append(geomHoleWalls.Faces, t1, t2)
}

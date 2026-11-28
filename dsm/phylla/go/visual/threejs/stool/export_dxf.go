package stool

import (
	"fmt"
	"math"
	"strings"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/cylinder"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

type dxfLayerSpec struct {
	Name  string
	Color int // AutoCAD Color Index (ACI)
}

func writeDXFHeaderAndLayers(sb *strings.Builder, layers []dxfLayerSpec) {
	sb.WriteString("0\nSECTION\n2\nHEADER\n")
	sb.WriteString("9\n$ACADVER\n1\nAC1009\n")
	sb.WriteString("9\n$INSUNITS\n70\n4\n") // 4 = millimeters
	sb.WriteString("0\nENDSEC\n")

	sb.WriteString("0\nSECTION\n2\nTABLES\n")
	sb.WriteString("0\nTABLE\n2\nLAYER\n")
	sb.WriteString(fmt.Sprintf("70\n%d\n", len(layers)))
	for _, l := range layers {
		sb.WriteString("0\nLAYER\n")
		sb.WriteString(fmt.Sprintf("2\n%s\n", l.Name))
		sb.WriteString("70\n0\n")
		sb.WriteString(fmt.Sprintf("62\n%d\n", l.Color))
		sb.WriteString("6\nCONTINUOUS\n")
	}
	sb.WriteString("0\nENDTAB\n")
	sb.WriteString("0\nENDSEC\n")
}

// writeDXFClosedPolylineZ0 writes a closed 2D POLYLINE at Z=0 in right-handed CAD XY coordinates
// (mapping Three.js (X, Z) -> CAD (X, -Z) so +Y_three -> +Z_cad preserves spiral chirality).
func writeDXFClosedPolylineZ0(sb *strings.Builder, layer string, pts []pt2D) {
	if len(pts) < 3 {
		return
	}
	sb.WriteString("0\nPOLYLINE\n")
	sb.WriteString(fmt.Sprintf("8\n%s\n", layer))
	sb.WriteString("66\n1\n")
	sb.WriteString("10\n0.0\n20\n0.0\n30\n0.0\n")
	sb.WriteString("70\n1\n") // flag 1 = closed polyline
	for _, p := range pts {
		sb.WriteString("0\nVERTEX\n")
		sb.WriteString(fmt.Sprintf("8\n%s\n", layer))
		sb.WriteString(fmt.Sprintf("10\n%.6f\n20\n%.6f\n30\n0.000000\n", p.X, -p.Z))
	}
	sb.WriteString("0\nSEQEND\n")
	sb.WriteString(fmt.Sprintf("8\n%s\n", layer))
}

func writeDXFGeom3DFaces(sb *strings.Builder, layer string, geom *threejs.BufferGeometry) {
	if geom == nil || len(geom.Vertices) == 0 || len(geom.Faces) == 0 {
		return
	}
	// Map Three.js right-handed (X, Y, Z) [Y up] -> CAD right-handed (X, -Z, Y) [Z up].
	toCAD := func(v *threejs.Vector3) (float64, float64, float64) {
		return v.X, -v.Z, v.Y
	}
	for _, f := range geom.Faces {
		if f.V1 >= len(geom.Vertices) || f.V2 >= len(geom.Vertices) || f.V3 >= len(geom.Vertices) {
			continue
		}
		x1, y1, z1 := toCAD(geom.Vertices[f.V1])
		x2, y2, z2 := toCAD(geom.Vertices[f.V2])
		x3, y3, z3 := toCAD(geom.Vertices[f.V3])

		sb.WriteString("0\n3DFACE\n")
		sb.WriteString(fmt.Sprintf("8\n%s\n", layer))
		sb.WriteString(fmt.Sprintf("10\n%.6f\n20\n%.6f\n30\n%.6f\n", x1, y1, z1))
		sb.WriteString(fmt.Sprintf("11\n%.6f\n21\n%.6f\n31\n%.6f\n", x2, y2, z2))
		sb.WriteString(fmt.Sprintf("12\n%.6f\n22\n%.6f\n32\n%.6f\n", x3, y3, z3))
		// Repeat 3rd vertex as 4th corner per DXF 3DFACE triangle convention.
		sb.WriteString(fmt.Sprintf("13\n%.6f\n23\n%.6f\n33\n%.6f\n", x3, y3, z3))
	}
}

// GenerateStool2DDXF exports the four planar contour sets projected onto Z=0 (as specified in
// Grasshopper stool_vt Panel #198 / Merge #295 for CNC/wire-cutting fabrication):
//   - SEAT_TOP:    closed outer contour of the seat top
//   - SEAT_BOTTOM: closed outer contour of the seat bottom
//   - FEET_TOP:    radial array of closed foot top contours
//   - FEET_BOTTOM: radial array of closed foot bottom contours
func GenerateStool2DDXF(stager *models.Stager, plant *models.PlantAbstract) string {
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		return ""
	}

	stage := threejs.NewStage("export_2d_dxf")
	params := cylinder.Cylinder3DParams{
		NamePrefix:            "Stool",
		CanvasName:            "Stool 3D Canvas",
		RadialRepetitions:     plant.StoolAbstract.RadialRepetitions,
		Transparency:          plant.StoolAbstract.Transparency,
		RelativeTubeDiameter:  plant.StoolAbstract.RelativeTubeDiameter,
		RelativeHeight3DTorus: plant.StoolAbstract.RelativeHeight3DTorus,
		VerticalScale:         plant.StoolAbstract.StoolTorusVerticalScale,
		RelativeHeight:        plant.StoolAbstract.RelativeHeight,
		ProjectionAngle:       plant.StoolAbstract.ProjectionAngle,
		HasRotatedShapes:      true,
	}

	base := cylinder.RenderCylinder3DBase(stage, stager, plant, params)
	if base == nil || base.ResampledBaseCurve == nil || plant.StartArcShapeGrid == nil || len(plant.StartArcShapeGrid.StartArcShapes) == 0 {
		return ""
	}

	var sb strings.Builder
	writeDXFHeaderAndLayers(&sb, []dxfLayerSpec{
		{Name: "SEAT_TOP", Color: 5},    // blue
		{Name: "SEAT_BOTTOM", Color: 3}, // green
		{Name: "FEET_TOP", Color: 1},    // red
		{Name: "FEET_BOTTOM", Color: 6}, // magenta
	})
	sb.WriteString("0\nSECTION\n2\nENTITIES\n")

	if plant.StoolAbstract.IsGrasshopperMode {
		ghRes := computeGrasshopperContours(
			stage,
			plant,
			base.ResampledBaseCurve,
			base.RadialRepetitions,
			base.GlobalR,
			base.TopHeight,
			base.TorusHeight,
			base.GrowthVectorX,
			base.GrowthVectorY,
			base.ProjAngleRad,
			base.ExpectedDegrees,
		)
		if ghRes == nil {
			return ""
		}

		writeDXFClosedPolylineZ0(&sb, "SEAT_TOP", ghRes.SeatTopRing)
		writeDXFClosedPolylineZ0(&sb, "SEAT_BOTTOM", ghRes.SeatBottomRing)

		rotateLoop := func(loop []pt2D, rot float64) []pt2D {
			c, s := math.Cos(rot), math.Sin(rot)
			out := make([]pt2D, len(loop))
			for i, p := range loop {
				out[i] = pt2D{X: p.X*c - p.Z*s, Z: p.X*s + p.Z*c}
			}
			return out
		}

		if ghRes.LegTopContour != nil {
			for k := range ghRes.RadialReps {
				rot := float64(k) * 2.0 * math.Pi / float64(ghRes.RadialReps)
				writeDXFClosedPolylineZ0(&sb, "FEET_TOP", rotateLoop(ghRes.LegTopContour.Loop, rot))
			}
		}
		if ghRes.LegBotContour != nil {
			for k := range ghRes.RadialReps {
				rot := float64(k) * 2.0 * math.Pi / float64(ghRes.RadialReps)
				writeDXFClosedPolylineZ0(&sb, "FEET_BOTTOM", rotateLoop(ghRes.LegBotContour.Loop, rot))
			}
		}
	} else {
		// Legacy mode fallback: export projected seat top/bottom and eye top/bottom contours.
		vecsTo2D := func(vecs []*threejs.Vector3) []pt2D {
			out := make([]pt2D, len(vecs))
			for i, v := range vecs {
				out[i] = pt2D{X: v.X, Z: v.Z}
			}
			return out
		}
		legacy := computeLegacyStoolCurves(stage, plant, base)
		if legacy == nil {
			return ""
		}
		writeDXFClosedPolylineZ0(&sb, "SEAT_TOP", vecsTo2D(base.RotTopPoints))
		writeDXFClosedPolylineZ0(&sb, "SEAT_BOTTOM", vecsTo2D(legacy.rotSeatBottomPoints))
		for k := range base.RadialRepetitions {
			rot := float64(k) * 2.0 * math.Pi / float64(base.RadialRepetitions)
			c, s := math.Cos(rot), math.Sin(rot)
			topLoop := make([]pt2D, len(legacy.projSeatBottomEyePoints))
			botLoop := make([]pt2D, len(legacy.projStoolBottomEyePoints))
			for i, p := range legacy.projSeatBottomEyePoints {
				topLoop[i] = pt2D{X: p.X*c - p.Z*s, Z: p.X*s + p.Z*c}
			}
			for i, p := range legacy.projStoolBottomEyePoints {
				botLoop[i] = pt2D{X: p.X*c - p.Z*s, Z: p.X*s + p.Z*c}
			}
			writeDXFClosedPolylineZ0(&sb, "FEET_TOP", topLoop)
			writeDXFClosedPolylineZ0(&sb, "FEET_BOTTOM", botLoop)
		}
	}

	sb.WriteString("0\nENDSEC\n0\nEOF\n")
	return sb.String()
}

// GenerateStool3DDXF exports the 3D stool seat and legs meshes as DXF 3DFACE entities
// on layers STOOL_SEAT and STOOL_LEGS in right-handed Z-up CAD coordinates.
func GenerateStool3DDXF(stager *models.Stager, plant *models.PlantAbstract) string {
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		return ""
	}

	stage := threejs.NewStage("export_3d_dxf")
	res := ComputeStoolGeometries(stage, stager, plant)
	if res == nil {
		return ""
	}

	var sb strings.Builder
	writeDXFHeaderAndLayers(&sb, []dxfLayerSpec{
		{Name: "STOOL_SEAT", Color: 40},
		{Name: "STOOL_LEGS", Color: 30},
	})
	sb.WriteString("0\nSECTION\n2\nENTITIES\n")

	writeDXFGeom3DFaces(&sb, "STOOL_SEAT", res.SeatGeom)
	writeDXFGeom3DFaces(&sb, "STOOL_LEGS", res.LegsGeom)

	sb.WriteString("0\nENDSEC\n0\nEOF\n")
	return sb.String()
}

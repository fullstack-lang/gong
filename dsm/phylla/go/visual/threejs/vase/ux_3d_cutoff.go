package vase

import (
	"math"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *ThreeJSStageUpdater) setupCutoffPlane(
	stager *models.Stager,
	canvas *threejs.Canvas,
	checkedDiagram *models.TubeVase3DDiagram,
	globalR float64,
) {
	if checkedDiagram == nil {
		return
	}
	threejsStage := stager.GetThreejsStage()
	if threejsStage == nil {
		return
	}

	minX, minY, minZ := math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
	maxX, maxY, maxZ := -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
	found := false

	expand := func(x, y, z, padX, padY, padZ float64) {
		found = true
		if x-padX < minX {
			minX = x - padX
		}
		if x+padX > maxX {
			maxX = x + padX
		}
		if y-padY < minY {
			minY = y - padY
		}
		if y+padY > maxY {
			maxY = y + padY
		}
		if z-padZ < minZ {
			minZ = z - padZ
		}
		if z+padZ > maxZ {
			maxZ = z + padZ
		}
	}

	for _, mesh := range canvas.Meshs {
		if mesh == nil || mesh.IsNotClipped {
			continue
		}
		mx, my, mz := mesh.X, mesh.Y, mesh.Z
		if bg := mesh.BufferGeometry; bg != nil {
			for _, v := range bg.Vertices {
				if v != nil {
					expand(mx+v.X, my+v.Y, mz+v.Z, 0, 0, 0)
				}
			}
		}
		if tg := mesh.TubeGeometry; tg != nil && tg.Path != nil {
			r := tg.Radius
			for _, p := range tg.Path.Points {
				if p != nil {
					expand(mx+p.X, my+p.Y, mz+p.Z, r, r, r)
				}
			}
		}
		if sg := mesh.SphereGeometry; sg != nil {
			r := sg.Radius
			expand(mx, my, mz, r, r, r)
		}
		if cg := mesh.CylinderGeometry; cg != nil {
			r := math.Max(cg.RadiusTop, cg.RadiusBottom)
			expand(mx, my, mz, r, 0.5*cg.Height, r)
		}
		if box := mesh.BoxGeometry; box != nil {
			expand(mx, my, mz, 0.5*box.Width, 0.5*box.Height, 0.5*box.Depth)
		}
		if tor := mesh.TorusGeometry; tor != nil {
			r := tor.Radius + tor.Tube
			expand(mx, my, mz, r, tor.Tube, r)
		}
	}

	if !found {
		r := globalR
		if r <= 0 {
			r = 50.0
		}
		minX, maxX = -r, r
		minY, maxY = 0, 2.0*r
		minZ, maxZ = -r, r
	} else {
		padX := math.Max(0.02*(maxX-minX), 0.5)
		padY := math.Max(0.02*(maxY-minY), 0.5)
		padZ := math.Max(0.02*(maxZ-minZ), 0.5)
		minX -= padX
		maxX += padX
		minY -= padY
		maxY += padY
		minZ -= padZ
		maxZ += padZ
	}

	u.boundsMinX, u.boundsMaxX = minX, maxX
	u.boundsMinY, u.boundsMaxY = minY, maxY
	u.boundsMinZ, u.boundsMaxZ = minZ, maxZ
	u.hasBounds = true

	cp := (&threejs.ClippingPlane{
		Name: "TubeVase Cutoff Plane",
	}).Stage(threejsStage)
	u.clippingPlane = cp
	canvas.ClippingPlanes = append(canvas.ClippingPlanes, cp)

	u.applyCutoffToPlane(checkedDiagram)
}

func (u *ThreeJSStageUpdater) applyCutoffToPlane(diagram *models.TubeVase3DDiagram) {
	if u.clippingPlane == nil || diagram == nil || !u.hasBounds {
		return
	}
	cp := u.clippingPlane
	cp.IsEnabled = diagram.IsCutoffEnabled
	cp.IsHelperVisible = diagram.IsCutoffHelperVisible

	centerX := 0.5 * (u.boundsMinX + u.boundsMaxX)
	centerY := 0.5 * (u.boundsMinY + u.boundsMaxY)
	centerZ := 0.5 * (u.boundsMinZ + u.boundsMaxZ)

	spanX := u.boundsMaxX - u.boundsMinX
	spanY := u.boundsMaxY - u.boundsMinY
	spanZ := u.boundsMaxZ - u.boundsMinZ
	maxSpan := math.Max(spanX, math.Max(spanY, spanZ))
	if maxSpan < 10.0 {
		maxSpan = 10.0
	}

	cp.HelperCenterX = centerX
	cp.HelperCenterY = centerY
	cp.HelperCenterZ = centerZ
	cp.HelperSize = maxSpan * 1.15

	axis := diagram.CutoffAxis
	if axis != models.CutoffAxisX && axis != models.CutoffAxisY && axis != models.CutoffAxisZ {
		axis = models.CutoffAxisX
	}

	// Default (!IsCutoffFlipped): normal = -1, keeps coord <= pos (cuts away +X/+Y/+Z facing camera).
	// Flipped (IsCutoffFlipped): normal = +1, keeps coord >= pos (cuts away -X/-Y/-Z).
	sign := -1.0
	if diagram.IsCutoffFlipped {
		sign = 1.0
	}

	offset := diagram.CutoffOffset
	if offset < -1.0 {
		offset = -1.0
	}
	if offset > 1.0 {
		offset = 1.0
	}

	cp.NormalX, cp.NormalY, cp.NormalZ = 0, 0, 0
	var pos float64
	switch axis {
	case models.CutoffAxisX:
		halfR := math.Max(math.Abs(u.boundsMinX), math.Abs(u.boundsMaxX))
		pos = offset * halfR
		cp.NormalX = sign
	case models.CutoffAxisY:
		pos = centerY + offset*(0.5*spanY)
		cp.NormalY = sign
	case models.CutoffAxisZ:
		halfR := math.Max(math.Abs(u.boundsMinZ), math.Abs(u.boundsMaxZ))
		pos = offset * halfR
		cp.NormalZ = sign
	}

	// Three.js keeps points where normal · p + Constant >= 0
	cp.Constant = -sign * pos
}

func (u *ThreeJSStageUpdater) updateCutoffPlane(stager *models.Stager) {
	threejsStage := stager.GetThreejsStage()
	plant := stager.GetCurrentPlant()
	if threejsStage == nil || plant == nil {
		return
	}

	var checkedDiagram *models.TubeVase3DDiagram
	for _, diagram := range plant.TubeVase3DDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}
	if checkedDiagram == nil && len(plant.TubeVase3DDiagrams) > 0 {
		checkedDiagram = plant.TubeVase3DDiagrams[0]
	}
	if checkedDiagram == nil {
		return
	}

	if u.clippingPlane == nil || !u.hasBounds {
		u.ux_3d_plant_diagram(stager)
		return
	}

	u.applyCutoffToPlane(checkedDiagram)
	threejsStage.Commit()
}

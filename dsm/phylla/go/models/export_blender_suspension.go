package models

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type blenderSuspensionLayer struct {
	index      int
	dyMM       float64
	thetaRad   float64
	colorHex   string
	isHorizCap bool
	horizMinY  float64
	horizMaxY  float64
}

var nonAlnumSlugRe = regexp.MustCompile(`[^a-zA-Z0-9_-]+`)

func sanitizeBlenderSlug(name string) string {
	s := strings.TrimSpace(name)
	s = nonAlnumSlugRe.ReplaceAllString(s, "_")
	s = strings.Trim(s, "_")
	if s == "" {
		return "tube_vase_suspension"
	}
	return s
}

// GenerateTubeVaseSuspensionBlenderPy generates a self-contained Blender Python (.py) script
// that builds an architectural room with the TubeVase hanging as a central ceiling suspension
// and sets up (and renders) two scenes:
//  1. 01_Pure_Suspension: pure sculptural suspension without lights in it (natural daylight room)
//  2. 02_Suspension_Inside_LEDs: suspension with warm architectural LEDs integrated along the inside border
func GenerateTubeVaseSuspensionBlenderPy(plant *PlantAbstract) string {
	if plant == nil || plant.TubeVaseAbstract == nil {
		return ""
	}

	vase := plant.TubeVaseAbstract
	radialRepetitions := max(vase.RadialRepetitions, 1)
	sideLength := plant.RhombusSideLength
	if sideLength <= 0 {
		sideLength = 200.0
	}
	relativeVerticalThickness := vase.RelativeVerticalThickness
	relativeRotatedTorusSeparation := vase.RelativeRotatedTorusSeparation
	relativeCuttedStackFloorHeight := vase.RelativeCuttedStackFloorHeight
	rotationRatio := vase.RotationRatio
	relativeHorizontalRingsHeight := vase.RelativeHorizontalRingsHeight

	thickness := 5.0
	if vase.RelativeRadialThickness*sideLength > 0 {
		thickness = vase.RelativeRadialThickness * sideLength
	}

	circumference := 10.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
		circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
	} else if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
		pGrid := plant.PerpendicularVectorGrid
		first := pGrid.PerpendicularVectors[0]
		last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
		circumference = last.StartX - first.StartX
	}
	if circumference <= 0 {
		circumference = 10.0
	}
	globalR := circumference * float64(radialRepetitions) / (2 * math.Pi)

	var resampledBaseBottom, resampledBaseTop []vector3
	if plant.GrowthCurve2D != nil && vase.TopGrowthCurve2D != nil &&
		plant.GrowthCurve2D.StartHalfwayArcShapeGrid != nil &&
		vase.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid != nil &&
		len(plant.GrowthCurve2D.StartHalfwayArcShapeGrid.StartHalfwayArcShapes) > 0 &&
		len(vase.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes) > 0 {

		gc := plant.GrowthCurve2D
		tgc := vase.TopGrowthCurve2D

		startArcs := gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes
		var endArcs []*EndHalfwayArcShape
		if gc.EndHalfwayArcShapeGrid != nil {
			endArcs = gc.EndHalfwayArcShapeGrid.EndHalfwayArcShapes
		}

		topStartArcs := tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes
		var topEndArcs []*TopEndHalfwayArcShape
		if tgc.TopEndHalfwayArcShapeGrid != nil {
			topEndArcs = tgc.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes
		}

		var curvePoints, topCurvePoints []vector3
		appendArcPts := func(targetCurve *[]vector3, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
			dx := (x1 - x2) / 2.0
			dy := (y1 - y2) / 2.0
			d2 := dx*dx + dy*dy
			var cx, cy float64
			if d2 == 0 || r*r < d2 {
				cx = (x1 + x2) / 2.0
				cy = (y1 + y2) / 2.0
				r = math.Sqrt(d2)
			} else {
				root := math.Sqrt(r*r/d2 - 1.0)
				if largeArcFlag == sweepFlag {
					root = -root
				}
				cx = (x1+x2)/2.0 + root*dy
				cy = (y1+y2)/2.0 - root*dx
			}

			startAngle := math.Atan2(y1-cy, x1-cx)
			endAngle := math.Atan2(y2-cy, x2-cx)
			if sweepFlag {
				for endAngle < startAngle {
					endAngle += 2 * math.Pi
				}
			} else {
				for endAngle > startAngle {
					endAngle -= 2 * math.Pi
				}
			}

			const steps = 50
			for i := 0; i <= steps; i++ {
				if i == 0 && len(*targetCurve) > 0 {
					continue
				}
				t := float64(i) / float64(steps)
				angle := startAngle + t*(endAngle-startAngle)
				x2d := cx + r*math.Cos(angle)
				y2d := cy + r*math.Sin(angle)

				theta := x2d / globalR
				*targetCurve = append(*targetCurve, vector3{
					X: globalR * math.Cos(theta),
					Y: y2d,
					Z: globalR * math.Sin(theta),
				})
			}
		}

		for i := range startArcs {
			sa := startArcs[i]
			appendArcPts(&curvePoints, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
			if i < len(endArcs) {
				ea := endArcs[i]
				appendArcPts(&curvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}
		for i := range topStartArcs {
			tsa := topStartArcs[i]
			appendArcPts(&topCurvePoints, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag)
			if i < len(topEndArcs) {
				ea := topEndArcs[i]
				appendArcPts(&topCurvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		if len(curvePoints) > 1 && len(topCurvePoints) > 1 {
			targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := getTargetAnglesSTL(curvePoints, topCurvePoints, 0.5, radialRepetitions)
			expectedDegrees := 360.0
			if radialRepetitions > 1 {
				expectedDegrees = 360.0 / float64(radialRepetitions)
			}
			resampledBaseBottom = resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
			resampledBaseTop = resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)
		}
	}

	if len(resampledBaseBottom) < 3 || len(resampledBaseTop) < 3 {
		return ""
	}

	var checkedDiagram *TubeVase3DDiagram
	for _, diagram := range plant.TubeVase3DDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}
	if checkedDiagram == nil && len(plant.TubeVase3DDiagrams) > 0 {
		checkedDiagram = plant.TubeVase3DDiagrams[0]
	}

	// Extract reference viewing azimuth from ThreeJS camera so the iconic profile faces the Blender room camera
	hasRefCamera := false
	refCamDX := -1.0
	refCamDZ := 0.0
	if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
		rs := checkedDiagram.Rendered3DShape
		dx := rs.ViewX - rs.TargetX
		dz := rs.ViewZ - rs.TargetZ
		if math.Hypot(dx, dz) > 1e-4 {
			refCamDX = dx
			refCamDZ = dz
			hasRefCamera = true
		}
	}

	stackHeight := max(plant.StackHeight, 1)
	var layerConfigs []stlLayerConfig

	computePartiallyRotatedStack := func() []stlLayerConfig {
		numSteps := stackHeight - 1
		dxs := make([]float64, stackHeight)
		dys := make([]float64, stackHeight)
		if numSteps > 0 {
			totalProgress := rotationRatio * float64(numSteps)
			var cumDX, cumDY float64
			for k := 1; k <= numSteps; k++ {
				var rK float64
				kFloat := float64(numSteps - k + 1)
				if totalProgress >= kFloat {
					rK = 1.0
				} else if totalProgress <= kFloat-1.0 {
					rK = 0.0
				} else {
					rK = totalProgress - (kFloat - 1.0)
				}
				stepDX, stepDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, rK)
				cumDX += stepDX
				cumDY += stepDY
				dxs[k] = cumDX
				dys[k] = cumDY
			}
		}
		out := make([]stlLayerConfig, 0, stackHeight)
		for h := range stackHeight {
			dx := dxs[h]
			dy := dys[h]
			out = append(out, stlLayerConfig{dx: dx, dy: dy, thetaOffset: dx / globalR})
		}
		return out
	}

	computeTorusStack := func() []stlLayerConfig {
		var growthVectorX, growthVectorY float64
		if plant.GrowthVectorShape != nil {
			growthVectorX = plant.GrowthVectorShape.X
			growthVectorY = plant.GrowthVectorShape.Y
		}
		var vx, vy float64
		if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
			vFirst := plant.PerpendicularVectorGrid.PerpendicularVectors[0]
			vx = vFirst.EndX - vFirst.StartX
			vy = vFirst.EndY - vFirst.StartY
			vLen := math.Hypot(vx, vy)
			if vLen == 0 {
				vLen = 1
			}
			vx, vy = vx/vLen, vy/vLen
		}
		verticalThickness := relativeVerticalThickness * sideLength
		rotatedSeparation := relativeRotatedTorusSeparation * sideLength

		out := make([]stlLayerConfig, 0, stackHeight)
		for h := range stackHeight {
			dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
			dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy + float64(h)*rotatedSeparation
			out = append(out, stlLayerConfig{dx: dx, dy: dy, thetaOffset: dx / globalR})
		}
		return out
	}

	if checkedDiagram != nil {
		if !checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape {
			layerConfigs = computePartiallyRotatedStack()
		} else if !checkedDiagram.IsHiddenTorusStackShape {
			layerConfigs = computeTorusStack()
		} else if !checkedDiagram.IsHiddenVerticalTorusStackShape {
			for h := range stackHeight {
				dy := float64(h) * relativeCuttedStackFloorHeight * sideLength
				layerConfigs = append(layerConfigs, stlLayerConfig{dx: 0, dy: dy, thetaOffset: 0})
			}
		} else if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
			dx, dy, _ := ComputePartiallyGrowthCurveDY(plant)
			layerConfigs = []stlLayerConfig{
				{dx: 0, dy: 0, thetaOffset: 0},
				{dx: dx, dy: dy, thetaOffset: dx / globalR},
			}
		}
	}
	if len(layerConfigs) == 0 {
		layerConfigs = computePartiallyRotatedStack()
	}

	var layers []blenderSuspensionLayer
	for h, cfg := range layerConfigs {
		layers = append(layers, blenderSuspensionLayer{
			index:    h,
			dyMM:     cfg.dy,
			thetaRad: cfg.thetaOffset,
			colorHex: GetVaseRingColor(vase.ColorCombination, h),
		})
	}

	hHoriz := relativeHorizontalRingsHeight * sideLength
	if hHoriz > 0 && len(layerConfigs) > 0 {
		massiveBot0, _ := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, layerConfigs[0].thetaOffset)
		minYBottom := math.MaxFloat64
		for _, p := range massiveBot0 {
			if p.Y+layerConfigs[0].dy < minYBottom {
				minYBottom = p.Y + layerConfigs[0].dy
			}
		}
		layers = append(layers, blenderSuspensionLayer{
			index:      0,
			dyMM:       0,
			thetaRad:   layerConfigs[0].thetaOffset,
			colorHex:   GetVaseRingColor(vase.ColorCombination, 0),
			isHorizCap: true,
			horizMinY:  minYBottom,
			horizMaxY:  minYBottom + hHoriz,
		})

		lastIdx := len(layerConfigs) - 1
		_, massiveTopLast := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, layerConfigs[lastIdx].thetaOffset)
		maxYTop := -math.MaxFloat64
		for _, p := range massiveTopLast {
			if p.Y+layerConfigs[lastIdx].dy > maxYTop {
				maxYTop = p.Y + layerConfigs[lastIdx].dy
			}
		}
		layers = append(layers, blenderSuspensionLayer{
			index:      lastIdx,
			dyMM:       0,
			thetaRad:   layerConfigs[lastIdx].thetaOffset,
			colorHex:   GetVaseRingColor(vase.ColorCombination, lastIdx),
			isHorizCap: true,
			horizMinY:  maxYTop - hHoriz,
			horizMaxY:  maxYTop,
		})
	}

	nBase := min(len(resampledBaseBottom), len(resampledBaseTop))
	slug := sanitizeBlenderSlug(plant.Name)
	transparency := math.Max(0.0, math.Min(0.95, vase.Transparency))

	var sb strings.Builder
	sb.Grow(64 * 1024)

	sb.WriteString(fmt.Sprintf(`# -*- coding: utf-8 -*-
"""
Blender Architectural Room Suspension Script — Generated by Phylla (Gong DSM)
Plant: %s (N=%d, M=%d, StackHeight=%d, Angle=%.1f°, Side=%.1f mm)

How to use in Blender (4.x / 5.x LTS):
  1. Open Blender -> switch to the 'Scripting' workspace tab.
  2. Click 'Open', select this .py file, and click 'Run Script' (▶).
  3. The script builds the full room + central hanging suspension and creates TWO scenes
     (selectable in the Scene dropdown at the top-right of Blender):
       • 01_Pure_Suspension       — Pure sculptural suspension without lights in it (natural daylight)
       • 02_Suspension_Inside_LEDs — Warm architectural LEDs integrated along the inside border of each ring
  4. When AUTO_RENDER_BOTH_IMAGES = True (default), running the script automatically renders
     both scenes to PNG files right next to this .py script (or ~/Downloads) and loads them
     into Blender's Image Editor ('Render_01_Pure_Suspension' and 'Render_02_Inside_LEDs').
"""

import math
import os
import sys
import bpy
from mathutils import Vector

# ==============================================================================
# USER CONFIGURATION
# ==============================================================================
PLANT_NAME = %q
PLANT_SLUG = %q

# Set to a target outer diameter in meters so the suspension has realistic chandelier
# proportions in the center of the room (set to None to use raw 1 mm = 0.001 m).
SUSPENSION_TARGET_DIAMETER_M = 0.86
SUSPENSION_MAX_HEIGHT_M = 1.32
SUSPENSION_CENTER_Z_M = 1.88

# Room dimensions in meters (floor at Z = 0, suspension centered at X = 0, Y = 0)
ROOM_WIDTH_M = 5.20
ROOM_DEPTH_M = 5.60
ROOM_HEIGHT_M = 3.10

# Ceiling canopy & ultra-fine tensioned aircraft cables
SHOW_SUSPENSION_HARDWARE = True
CABLE_COUNT = 3
CABLE_RADIUS_M = 0.00075

# Inside-border warm-white COB LED diffuser parameters (for Scene 2)
LED_COLOR_LINEAR = (1.00, 0.74, 0.44)   # ~2800K warm architectural white
LED_EMISSION_STRENGTH = 26.0

# Render output settings
AUTO_RENDER_BOTH_IMAGES = True
RENDER_RESOLUTION_X = 1600
RENDER_RESOLUTION_Y = 1600
CYCLES_SAMPLES_PURE = 160
CYCLES_SAMPLES_LEDS = 224

# ==============================================================================
# PHYLLA GEOMETRY DATA (millimeters & radians)
# ==============================================================================
RADIAL_REPETITIONS = %d
RADIAL_THICKNESS_MM = %.6f
GLOBAL_RADIUS_MM = %.6f
TRANSPARENCY = %.4f
HAS_REF_CAMERA = %s
REF_CAM_DX = %.6f
REF_CAM_DZ = %.6f

`,
		plant.Name, plant.N, plant.M, stackHeight, plant.RhombusInsideAngle, sideLength,
		plant.Name, slug,
		radialRepetitions, thickness, globalR, transparency,
		pyBool(hasRefCamera), refCamDX, refCamDZ,
	))

	// Emit BASE_BOTTOM: (theta_rad, r_mm, y_mm)
	sb.WriteString("BASE_BOTTOM = [\n")
	for i := range nBase {
		p := resampledBaseBottom[i]
		th := math.Atan2(p.Z, p.X)
		r := math.Hypot(p.X, p.Z)
		sb.WriteString(fmt.Sprintf("    (%.6f, %.4f, %.4f),\n", th, r, p.Y))
	}
	sb.WriteString("]\n\n")

	// Emit BASE_TOP: (theta_rad, r_mm, y_mm)
	sb.WriteString("BASE_TOP = [\n")
	for i := range nBase {
		p := resampledBaseTop[i]
		th := math.Atan2(p.Z, p.X)
		r := math.Hypot(p.X, p.Z)
		sb.WriteString(fmt.Sprintf("    (%.6f, %.4f, %.4f),\n", th, r, p.Y))
	}
	sb.WriteString("]\n\n")

	// Emit LAYERS
	sb.WriteString("LAYERS = [\n")
	for _, l := range layers {
		sb.WriteString(fmt.Sprintf(
			"    {\"index\": %d, \"dy_mm\": %.5f, \"theta_rad\": %.6f, \"color_hex\": %q, \"is_horiz\": %s, \"horiz_min_y\": %.5f, \"horiz_max_y\": %.5f},\n",
			l.index, l.dyMM, l.thetaRad, l.colorHex, pyBool(l.isHorizCap), l.horizMinY, l.horizMaxY,
		))
	}
	sb.WriteString("]\n\n")

	sb.WriteString(blenderSuspensionPythonEngine)

	return sb.String()
}

func pyBool(v bool) string {
	if v {
		return "True"
	}
	return "False"
}

const blenderSuspensionPythonEngine = `# ==============================================================================
# BLENDER SCENE BUILDER & DUAL RENDER ENGINE
# ==============================================================================

def ensure_node_tree(id_block):
    if getattr(id_block, "node_tree", None) is None and hasattr(id_block, "use_nodes"):
        id_block.use_nodes = True
    return id_block.node_tree


def srgb_to_linear(c):
    if c <= 0.04045:
        return c / 12.92
    return ((c + 0.055) / 1.055) ** 2.4


def hex_to_linear_rgba(hex_str, alpha=1.0):
    s = hex_str.strip().lstrip("#")
    if len(s) != 6:
        return (0.8, 0.8, 0.8, alpha)
    r = srgb_to_linear(int(s[0:2], 16) / 255.0)
    g = srgb_to_linear(int(s[2:4], 16) / 255.0)
    b = srgb_to_linear(int(s[4:6], 16) / 255.0)
    return (r, g, b, alpha)


def set_socket(node, names, value):
    if isinstance(names, str):
        names = (names,)
    for name in names:
        if name in node.inputs:
            node.inputs[name].default_value = value
            return True
    return False


def clear_all_datablocks():
    # Create a fresh temporary scene first so we can safely remove all old scenes
    fresh_scene = bpy.data.scenes.new("__phylla_init__")
    if bpy.context.window is not None:
        bpy.context.window.scene = fresh_scene
    for sc in list(bpy.data.scenes):
        if sc != fresh_scene:
            bpy.data.scenes.remove(sc, do_unlink=True)

    for block_collection in (
        bpy.data.objects,
        bpy.data.meshes,
        bpy.data.curves,
        bpy.data.materials,
        bpy.data.lights,
        bpy.data.cameras,
        bpy.data.collections,
        bpy.data.worlds,
    ):
        for item in list(block_collection):
            block_collection.remove(item, do_unlink=True)

    return fresh_scene


def make_collection(name, parent=None):
    col = bpy.data.collections.new(name)
    if parent is not None:
        parent.children.link(col)
    return col


def add_mesh_object(name, verts, faces, collection, material=None, smooth=True):
    mesh = bpy.data.meshes.new(name + "_Mesh")
    mesh.from_pydata(verts, [], faces)
    mesh.update(calc_edges=True)
    if smooth:
        if hasattr(mesh, "shade_smooth"):
            mesh.shade_smooth()
        else:
            for poly in mesh.polygons:
                poly.use_smooth = True
    else:
        if hasattr(mesh, "shade_flat"):
            mesh.shade_flat()

    obj = bpy.data.objects.new(name, mesh)
    collection.objects.link(obj)
    if material is not None:
        obj.data.materials.append(material)
    return obj


# ------------------------------------------------------------------------------
# MATERIALS
# ------------------------------------------------------------------------------

def create_ring_material(name, hex_color, transparency=0.0):
    mat = bpy.data.materials.new(name)
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (360, 0)

    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (40, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    rgba = hex_to_linear_rgba(hex_color, 1.0)
    set_socket(bsdf, "Base Color", rgba)
    set_socket(bsdf, "Metallic", 0.04)

    # Subtle micro-grain roughness variation for tactile ceramic / powder-coated finish
    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-560, -140)
    noise = nodes.new("ShaderNodeTexNoise")
    noise.location = (-360, -140)
    set_socket(noise, "Scale", 160.0)
    set_socket(noise, "Detail", 4.0)
    set_socket(noise, "Roughness", 0.55)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-180, -140)
    if transparency > 0.01:
        ramp.color_ramp.elements[0].color = (0.05, 0.05, 0.05, 1.0)
        ramp.color_ramp.elements[1].color = (0.11, 0.11, 0.11, 1.0)
        set_socket(bsdf, ("Transmission Weight", "Transmission"), transparency)
        set_socket(bsdf, "IOR", 1.52)
    else:
        ramp.color_ramp.elements[0].color = (0.28, 0.28, 0.28, 1.0)
        ramp.color_ramp.elements[1].color = (0.39, 0.39, 0.39, 1.0)
        set_socket(bsdf, ("Coat Weight", "Clearcoat"), 0.16)
        set_socket(bsdf, ("Coat Roughness", "Clearcoat Roughness"), 0.28)
        # Subtle subsurface warmth so strong internal LED grazing feels rich & organic
        set_socket(bsdf, ("Subsurface Weight", "Subsurface"), 0.04)

    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Roughness"])

    return mat


def create_led_diffuser_material():
    mat = bpy.data.materials.new("LED_Inside_Border_Diffuser")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (300, 0)

    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (0, 0)
    set_socket(bsdf, "Base Color", (1.0, 0.96, 0.90, 1.0))
    set_socket(bsdf, "Roughness", 0.22)
    set_socket(bsdf, ("Emission Color", "Emission"), (*LED_COLOR_LINEAR, 1.0))
    set_socket(bsdf, "Emission Strength", LED_EMISSION_STRENGTH)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])
    return mat


def create_plaster_wall_material(name, base_rgba=(0.76, 0.72, 0.66, 1.0), roughness_base=0.84):
    mat = bpy.data.materials.new(name)
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (420, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (140, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-620, 0)

    noise_macro = nodes.new("ShaderNodeTexNoise")
    noise_macro.location = (-400, 120)
    set_socket(noise_macro, "Scale", 2.6)
    set_socket(noise_macro, "Detail", 5.0)
    set_socket(noise_macro, "Roughness", 0.6)
    links.new(tex_coord.outputs["Object"], noise_macro.inputs["Vector"])

    mix_col = nodes.new("ShaderNodeValToRGB")
    mix_col.location = (-160, 140)
    c0 = tuple(max(0.0, v * 0.94) for v in base_rgba[:3]) + (1.0,)
    c1 = tuple(min(1.0, v * 1.04) for v in base_rgba[:3]) + (1.0,)
    mix_col.color_ramp.elements[0].position = 0.25
    mix_col.color_ramp.elements[0].color = c0
    mix_col.color_ramp.elements[1].position = 0.78
    mix_col.color_ramp.elements[1].color = c1
    links.new(noise_macro.outputs["Fac"], mix_col.inputs["Fac"])
    links.new(mix_col.outputs["Color"], bsdf.inputs["Base Color"])

    set_socket(bsdf, "Roughness", roughness_base)

    noise_micro = nodes.new("ShaderNodeTexNoise")
    noise_micro.location = (-400, -140)
    set_socket(noise_micro, "Scale", 55.0)
    set_socket(noise_micro, "Detail", 4.0)
    links.new(tex_coord.outputs["Object"], noise_micro.inputs["Vector"])

    bump = nodes.new("ShaderNodeBump")
    bump.location = (-120, -140)
    set_socket(bump, "Strength", 0.025)
    set_socket(bump, "Distance", 0.002)
    links.new(noise_micro.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])

    return mat


def create_oak_floor_material():
    mat = bpy.data.materials.new("Room_Oak_Floor")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (460, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (180, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-720, 0)
    mapping = nodes.new("ShaderNodeMapping")
    mapping.location = (-520, 0)
    set_socket(mapping, "Scale", (8.0, 1.4, 1.0))
    links.new(tex_coord.outputs["Object"], mapping.inputs["Vector"])

    wave = nodes.new("ShaderNodeTexWave")
    wave.location = (-300, 100)
    wave.wave_type = "BANDS"
    wave.bands_direction = "X"
    set_socket(wave, "Scale", 4.5)
    set_socket(wave, "Distortion", 4.2)
    set_socket(wave, "Detail", 6.0)
    links.new(mapping.outputs["Vector"], wave.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-70, 120)
    ramp.color_ramp.elements[0].position = 0.18
    ramp.color_ramp.elements[0].color = (0.33, 0.22, 0.13, 1.0)
    ramp.color_ramp.elements[1].position = 0.85
    ramp.color_ramp.elements[1].color = (0.48, 0.35, 0.22, 1.0)
    links.new(wave.outputs["Color"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])

    set_socket(bsdf, "Roughness", 0.38)
    set_socket(bsdf, ("Coat Weight", "Clearcoat"), 0.08)
    set_socket(bsdf, ("Coat Roughness", "Clearcoat Roughness"), 0.35)

    bump = nodes.new("ShaderNodeBump")
    bump.location = (-60, -120)
    set_socket(bump, "Strength", 0.03)
    set_socket(bump, "Distance", 0.002)
    links.new(wave.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])

    return mat


def create_metal_material(name, color_rgba=(0.08, 0.08, 0.09, 1.0), metallic=0.88, roughness=0.30):
    mat = bpy.data.materials.new(name)
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()
    out = nodes.new("ShaderNodeOutputMaterial")
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    set_socket(bsdf, "Base Color", color_rgba)
    set_socket(bsdf, "Metallic", metallic)
    set_socket(bsdf, "Roughness", roughness)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])
    return mat


def create_boucle_rug_material():
    mat = bpy.data.materials.new("Center_Wool_Rug")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (380, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (100, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-520, 0)
    noise = nodes.new("ShaderNodeTexNoise")
    noise.location = (-320, 0)
    set_socket(noise, "Scale", 220.0)
    set_socket(noise, "Detail", 8.0)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-110, 80)
    ramp.color_ramp.elements[0].color = (0.56, 0.52, 0.46, 1.0)
    ramp.color_ramp.elements[1].color = (0.69, 0.65, 0.58, 1.0)
    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])

    set_socket(bsdf, "Roughness", 0.94)
    set_socket(bsdf, ("Sheen Weight", "Sheen"), 0.35)

    bump = nodes.new("ShaderNodeBump")
    bump.location = (-110, -120)
    set_socket(bump, "Strength", 0.12)
    set_socket(bump, "Distance", 0.004)
    links.new(noise.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])

    return mat


# ------------------------------------------------------------------------------
# SUSPENSION MESH GENERATION (RINGS + INSIDE-BORDER LED DIFFUSERS + CABLES)
# ------------------------------------------------------------------------------

def build_repetition_loops(theta_offset):
    """Builds 1 closed loop (when continuous around 360°) or RADIAL_REPETITIONS open segments."""
    reps = max(1, RADIAL_REPETITIONS)
    raw_bot = []
    raw_top = []
    n_pts = min(len(BASE_BOTTOM), len(BASE_TOP))

    for k in range(reps):
        base_off = k * 2.0 * math.pi / reps
        tot_off = theta_offset + base_off
        seg_b = []
        seg_t = []
        for i in range(n_pts):
            th_b, r_b, y_b = BASE_BOTTOM[i]
            th_t, r_t, y_t = BASE_TOP[i]
            seg_b.append((th_b + tot_off, r_b, y_b))
            seg_t.append((th_t + tot_off, r_t, y_t))
        raw_bot.append(seg_b)
        raw_top.append(seg_t)

    def pt3(th, r, y):
        return (r * math.cos(th), y, r * math.sin(th))

    def dist3(a, b):
        return math.sqrt((a[0] - b[0]) ** 2 + (a[1] - b[1]) ** 2 + (a[2] - b[2]) ** 2)

    # Check if all repetitions stitch into a single closed 360-degree manifold ring
    all_stitch = True
    for k in range(reps):
        nxt = (k + 1) % reps
        end_b = pt3(*raw_bot[k][-1])
        start_b = pt3(*raw_bot[nxt][0])
        end_t = pt3(*raw_top[k][-1])
        start_t = pt3(*raw_top[nxt][0])
        if dist3(end_b, start_b) > 0.5 or dist3(end_t, start_t) > 0.5:
            all_stitch = False
            break

    if all_stitch:
        merged_b = []
        merged_t = []
        for k in range(reps):
            # Drop the last point of each repetition since it coincides with the start of the next
            merged_b.extend(raw_bot[k][:-1])
            merged_t.extend(raw_top[k][:-1])
        return [(merged_b, merged_t, True)]

    return [(raw_bot[k], raw_top[k], False) for k in range(reps)]


def build_suspension_geometry(col_rings, col_leds, col_hardware, cam_xy):
    # Compute raw bounding box in mm to determine uniform architectural scale & vertical centering
    min_y_mm = float("inf")
    max_y_mm = float("-inf")
    max_r_out_mm = GLOBAL_RADIUS_MM + RADIAL_THICKNESS_MM

    for _, r_b, _ in BASE_BOTTOM:
        max_r_out_mm = max(max_r_out_mm, r_b + RADIAL_THICKNESS_MM)
    for _, r_t, _ in BASE_TOP:
        max_r_out_mm = max(max_r_out_mm, r_t + RADIAL_THICKNESS_MM)

    for layer in LAYERS:
        if layer["is_horiz"]:
            min_y_mm = min(min_y_mm, layer["horiz_min_y"])
            max_y_mm = max(max_y_mm, layer["horiz_max_y"])
        else:
            dy = layer["dy_mm"]
            for _, _, yb in BASE_BOTTOM:
                min_y_mm = min(min_y_mm, yb + dy)
                max_y_mm = max(max_y_mm, yb + dy)
            for _, _, yt in BASE_TOP:
                min_y_mm = min(min_y_mm, yt + dy)
                max_y_mm = max(max_y_mm, yt + dy)

    raw_diam_mm = max(1e-3, 2.0 * max_r_out_mm)
    raw_height_mm = max(1e-3, max_y_mm - min_y_mm)

    if SUSPENSION_TARGET_DIAMETER_M is not None and SUSPENSION_TARGET_DIAMETER_M > 0:
        scale = SUSPENSION_TARGET_DIAMETER_M / raw_diam_mm
    else:
        scale = 0.001  # 1 mm = 0.001 m

    if raw_height_mm * scale > SUSPENSION_MAX_HEIGHT_M:
        scale = SUSPENSION_MAX_HEIGHT_M / raw_height_mm

    mid_y_mm = 0.5 * (min_y_mm + max_y_mm)
    z_offset_m = SUSPENSION_CENTER_Z_M - mid_y_mm * scale

    # Align the ThreeJS reference viewing angle directly toward our Blender room camera
    cam_azimuth = math.atan2(cam_xy[1], cam_xy[0])
    if HAS_REF_CAMERA:
        src_azimuth = math.atan2(-REF_CAM_DZ, REF_CAM_DX)
        delta_alpha = cam_azimuth - src_azimuth
    else:
        delta_alpha = 0.0

    material_cache = {}
    led_mat = create_led_diffuser_material()

    top_most_z = float("-inf")
    bottom_most_z = float("inf")

    for layer_idx, layer in enumerate(LAYERS):
        h = layer["index"]
        hex_col = layer["color_hex"]
        mat_key = (hex_col, round(TRANSPARENCY, 4))
        if mat_key not in material_cache:
            material_cache[mat_key] = create_ring_material(
                f"Suspension_Mat_{hex_col.lstrip('#')}", hex_col, TRANSPARENCY
            )
        ring_mat = material_cache[mat_key]

        loops = build_repetition_loops(layer["theta_rad"])
        dy_mm = layer["dy_mm"]

        ring_verts = []
        ring_faces = []
        led_verts = []
        led_faces = []

        for seg_b, seg_t, is_closed in loops:
            m_pts = len(seg_b)
            if m_pts < 3:
                continue

            ring_base_v = len(ring_verts)
            led_base_v = len(led_verts)

            # 8-point micro-beveled ring profile + 4-point recessed inside-border LED diffuser profile
            for i in range(m_pts):
                th_b, r_b, yb_raw = seg_b[i]
                th_t, r_t, yt_raw = seg_t[i]

                if layer["is_horiz"]:
                    z_bot = layer["horiz_min_y"] * scale + z_offset_m
                    z_top = layer["horiz_max_y"] * scale + z_offset_m
                else:
                    z_bot = (yb_raw + dy_mm) * scale + z_offset_m
                    z_top = (yt_raw + dy_mm) * scale + z_offset_m

                if z_top < z_bot:
                    z_bot, z_top = z_top, z_bot

                top_most_z = max(top_most_z, z_top)
                bottom_most_z = min(bottom_most_z, z_bot)

                phi_b = -th_b + delta_alpha
                phi_t = -th_t + delta_alpha
                # Handle seam wrap cleanly when interpolating phi_b and phi_t
                dphi = phi_t - phi_b
                while dphi > math.pi:
                    dphi -= 2.0 * math.pi
                while dphi < -math.pi:
                    dphi += 2.0 * math.pi

                r_in_b = r_b * scale
                r_in_t = r_t * scale
                w_rad = RADIAL_THICKNESS_MM * scale
                h_vert = max(1e-5, z_top - z_bot)
                chamfer = min(w_rad, h_vert) * 0.08

                def ring_pt(r_off, z_val):
                    t_fac = (z_val - z_bot) / h_vert if h_vert > 1e-6 else 0.5
                    phi = phi_b + t_fac * dphi
                    r_in = r_in_b + t_fac * (r_in_t - r_in_b)
                    r_tot = r_in + r_off
                    return (r_tot * math.cos(phi), r_tot * math.sin(phi), z_val)

                # Cross-section ordered around perimeter so outward normals face away from ring interior:
                # 0: bottom-inner, 1: bottom-outer, 2: outer-bottom, 3: outer-top,
                # 4: top-outer, 5: top-inner, 6: inner-top, 7: inner-bottom
                ring_verts.extend([
                    ring_pt(chamfer, z_bot),
                    ring_pt(w_rad - chamfer, z_bot),
                    ring_pt(w_rad, z_bot + chamfer),
                    ring_pt(w_rad, z_top - chamfer),
                    ring_pt(w_rad - chamfer, z_top),
                    ring_pt(chamfer, z_top),
                    ring_pt(0.0, z_top - chamfer),
                    ring_pt(0.0, z_bot + chamfer),
                ])

                # Inside-border LED opal diffuser profile along r_in (facing inward toward axis):
                z_l0 = z_bot + 0.20 * h_vert
                z_l1 = z_top - 0.20 * h_vert
                h_led = max(1e-5, z_l1 - z_l0)
                led_protrude = max(0.0012, min(0.0035, 0.16 * w_rad))
                led_tuck = 0.03 * w_rad

                led_verts.extend([
                    ring_pt(led_tuck, z_l0),
                    ring_pt(-led_protrude, z_l0 + 0.18 * h_led),
                    ring_pt(-led_protrude, z_l1 - 0.18 * h_led),
                    ring_pt(led_tuck, z_l1),
                ])

            n_steps = m_pts if is_closed else (m_pts - 1)
            for i in range(n_steps):
                nxt = (i + 1) % m_pts
                b0 = ring_base_v + 8 * i
                b1 = ring_base_v + 8 * nxt
                # As i increases, phi = -theta decreases (-Y tangent at +X), so (b0+p, b0+p_next, b1+p_next, b1+p)
                # points strictly OUTWARD for all 8 perimeter facets!
                for p in range(8):
                    pn = (p + 1) % 8
                    ring_faces.append((b0 + p, b0 + pn, b1 + pn, b1 + p))

                if not layer["is_horiz"]:
                    lb0 = led_base_v + 4 * i
                    lb1 = led_base_v + 4 * nxt
                    # For inward-facing LED bead (0: bot-tuck -> 1: bot-in -> 2: top-in -> 3: top-tuck):
                    # Perimeter outward toward center axis (-R) uses (lb0+p+1, lb0+p, lb1+p, lb1+p+1)
                    for p in range(3):
                        led_faces.append((lb0 + p + 1, lb0 + p, lb1 + p, lb1 + p + 1))

            if not is_closed:
                # End caps for open segments
                b_start = ring_base_v
                ring_faces.append(tuple(b_start + p for p in range(7, -1, -1)))
                b_end = ring_base_v + 8 * (m_pts - 1)
                ring_faces.append(tuple(b_end + p for p in range(8)))

        suffix = f"HorizCap_{layer_idx}" if layer["is_horiz"] else f"Ring_{h:02d}"
        add_mesh_object(f"Suspension_{suffix}", ring_verts, ring_faces, col_rings, ring_mat, smooth=True)

        if led_faces:
            led_obj = add_mesh_object(
                f"LED_InsideBorder_{suffix}", led_verts, led_faces, col_leds, led_mat, smooth=True
            )
            # Invisible to shadow rays so the diffuser never blocks its own bounce light
            if hasattr(led_obj, "visible_shadow"):
                led_obj.visible_shadow = False

    # Build ceiling canopy (rose) & ultra-fine stainless steel suspension wires
    if SHOW_SUSPENSION_HARDWARE:
        bronze_mat = create_metal_material("Canopy_Brushed_Bronze", (0.055, 0.050, 0.046, 1.0), metallic=0.85, roughness=0.32)
        wire_mat = create_metal_material("Suspension_Steel_Wire", (0.38, 0.39, 0.41, 1.0), metallic=0.92, roughness=0.22)

        r_mid_m = (GLOBAL_RADIUS_MM + 0.5 * RADIAL_THICKNESS_MM) * scale
        canopy_r = max(0.075, min(0.14, r_mid_m * 0.30))
        canopy_h = 0.026
        add_cylinder_mesh(
            "Ceiling_Canopy",
            center=(0.0, 0.0, ROOM_HEIGHT_M - 0.5 * canopy_h),
            radius=canopy_r,
            height=canopy_h,
            segments=64,
            collection=col_hardware,
            material=bronze_mat,
        )

        # Find exact bottom-most ring height at each cable angle so each wire terminates cleanly at Ring 0
        for c_idx in range(CABLE_COUNT):
            wire_phi = delta_alpha + (2.0 * math.pi * c_idx / max(1, CABLE_COUNT)) + 0.35
            # Attach wire from ceiling down to bottom ring at radius r_mid_m
            wx = r_mid_m * math.cos(wire_phi)
            wy = r_mid_m * math.sin(wire_phi)

            # Sample Ring 0 height at this angle so cable stops right inside the lowest ring
            z_wire_bottom = bottom_most_z + 0.04
            if LAYERS:
                th_target = -(wire_phi - delta_alpha)
                best_diff = float("inf")
                best_y = BASE_BOTTOM[0][2]
                for th_b, _, yb in BASE_BOTTOM:
                    diff = (th_b + LAYERS[0]["theta_rad"] - th_target + math.pi) % (2.0 * math.pi) - math.pi
                    if abs(diff) < best_diff:
                        best_diff = abs(diff)
                        best_y = yb
                z_wire_bottom = (best_y + LAYERS[0]["dy_mm"]) * scale + z_offset_m + 0.002

            wire_len = max(0.05, ROOM_HEIGHT_M - z_wire_bottom)
            wire_cz = z_wire_bottom + 0.5 * wire_len

            # Subtle top diverter pin on ceiling for architectural realism (vertical drop at r_mid_m)
            add_cylinder_mesh(
                f"Ceiling_Anchor_Pin_{c_idx + 1}",
                center=(wx, wy, ROOM_HEIGHT_M - 0.008),
                radius=0.0045,
                height=0.016,
                segments=20,
                collection=col_hardware,
                material=bronze_mat,
            )
            add_cylinder_mesh(
                f"Suspension_Cable_{c_idx + 1}",
                center=(wx, wy, wire_cz),
                radius=CABLE_RADIUS_M,
                height=wire_len,
                segments=12,
                collection=col_hardware,
                material=wire_mat,
            )

    return scale, bottom_most_z, top_most_z


def add_cylinder_mesh(name, center, radius, height, segments, collection, material, smooth=True):
    cx, cy, cz = center
    z0 = cz - 0.5 * height
    z1 = cz + 0.5 * height
    verts = []
    faces = []
    for i in range(segments):
        ang = 2.0 * math.pi * i / segments
        x = cx + radius * math.cos(ang)
        y = cy + radius * math.sin(ang)
        verts.append((x, y, z0))
        verts.append((x, y, z1))
    for i in range(segments):
        nxt = (i + 1) % segments
        b0 = 2 * i
        t0 = 2 * i + 1
        b1 = 2 * nxt
        t1 = 2 * nxt + 1
        faces.append((b0, b1, t1, t0))
    # Top and bottom ngon caps
    faces.append(tuple(2 * i + 1 for i in range(segments)))
    faces.append(tuple(2 * i for i in range(segments - 1, -1, -1)))
    return add_mesh_object(name, verts, faces, collection, material, smooth=smooth)


# ------------------------------------------------------------------------------
# ARCHITECTURAL ROOM BUILDER
# ------------------------------------------------------------------------------

def build_architectural_room(col_room):
    hw = 0.5 * ROOM_WIDTH_M
    hd = 0.5 * ROOM_DEPTH_M
    h = ROOM_HEIGHT_M

    wall_mat = create_plaster_wall_material("Room_Limewash_Plaster", (0.74, 0.70, 0.64, 1.0), roughness_base=0.85)
    ceil_mat = create_plaster_wall_material("Room_Matte_Ceiling", (0.85, 0.83, 0.79, 1.0), roughness_base=0.92)
    trim_mat = create_plaster_wall_material("Room_Baseboard_Trim", (0.80, 0.77, 0.72, 1.0), roughness_base=0.65)
    floor_mat = create_oak_floor_material()
    mullion_mat = create_metal_material("Window_Bronze_Mullion", (0.045, 0.042, 0.040, 1.0), metallic=0.80, roughness=0.35)
    rug_mat = create_boucle_rug_material()

    # 1. Floor (Z = 0, normal +Z)
    add_mesh_object(
        "Room_Floor",
        [(-hw, -hd, 0.0), (hw, -hd, 0.0), (hw, hd, 0.0), (-hw, hd, 0.0)],
        [(0, 1, 2, 3)],
        col_room,
        floor_mat,
        smooth=False,
    )

    # 2. Ceiling (Z = h, normal -Z)
    add_mesh_object(
        "Room_Ceiling",
        [(-hw, -hd, h), (-hw, hd, h), (hw, hd, h), (hw, -hd, h)],
        [(0, 1, 2, 3)],
        col_room,
        ceil_mat,
        smooth=False,
    )

    # 3. Back Wall (Y = +hd, normal -Y) & Left Wall (X = -hw, normal +X) & Front Wall behind camera (Y = -hd, normal +Y)
    add_mesh_object(
        "Room_Back_Wall",
        [(-hw, hd, 0.0), (hw, hd, 0.0), (hw, hd, h), (-hw, hd, h)],
        [(0, 1, 2, 3)],
        col_room,
        wall_mat,
        smooth=False,
    )
    add_mesh_object(
        "Room_Left_Wall",
        [(-hw, -hd, 0.0), (-hw, hd, 0.0), (-hw, hd, h), (-hw, -hd, h)],
        [(0, 1, 2, 3)],
        col_room,
        wall_mat,
        smooth=False,
    )
    add_mesh_object(
        "Room_Front_Wall",
        [(hw, -hd, 0.0), (-hw, -hd, 0.0), (-hw, -hd, h), (hw, -hd, h)],
        [(0, 1, 2, 3)],
        col_room,
        wall_mat,
        smooth=False,
    )

    # 4. Right Wall (X = +hw, normal -X) with recessed tall architectural window aperture
    wy0, wy1 = -0.85, 1.45
    wz0, wz1 = 0.48, 2.68
    reveal_x = hw + 0.24

    rw_verts = [
        # Sill band below window
        (hw, -hd, 0.0), (hw, -hd, wz0), (hw, hd, wz0), (hw, hd, 0.0),
        # Header band above window
        (hw, -hd, wz1), (hw, -hd, h), (hw, hd, h), (hw, hd, wz1),
        # Front pier
        (hw, -hd, wz0), (hw, -hd, wz1), (hw, wy0, wz1), (hw, wy0, wz0),
        # Back pier
        (hw, wy1, wz0), (hw, wy1, wz1), (hw, hd, wz1), (hw, hd, wz0),
        # Deep window reveal jambs (sill, head, front jamb, back jamb)
        (hw, wy0, wz0), (reveal_x, wy0, wz0), (reveal_x, wy1, wz0), (hw, wy1, wz0),
        (hw, wy1, wz1), (reveal_x, wy1, wz1), (reveal_x, wy0, wz1), (hw, wy0, wz1),
        (hw, wy0, wz0), (hw, wy0, wz1), (reveal_x, wy0, wz1), (reveal_x, wy0, wz0),
        (hw, wy1, wz1), (hw, wy1, wz0), (reveal_x, wy1, wz0), (reveal_x, wy1, wz1),
    ]
    rw_faces = [(4 * k, 4 * k + 1, 4 * k + 2, 4 * k + 3) for k in range(8)]
    add_mesh_object("Room_Right_Window_Wall", rw_verts, rw_faces, col_room, wall_mat, smooth=False)

    # Slim bronze window frame + architectural mullions
    def add_box(name, x0, x1, y0, y1, z0, z1, mat):
        v = [
            (x0, y0, z0), (x1, y0, z0), (x1, y1, z0), (x0, y1, z0),
            (x0, y0, z1), (x1, y0, z1), (x1, y1, z1), (x0, y1, z1),
        ]
        f = [
            (0, 3, 2, 1), (4, 5, 6, 7),
            (0, 1, 5, 4), (2, 3, 7, 6),
            (1, 2, 6, 5), (3, 0, 4, 7),
        ]
        return add_mesh_object(name, v, f, col_room, mat, smooth=False)

    fx = reveal_x - 0.025
    ft = 0.032
    add_box("Window_Frame_Bottom", fx - 0.02, fx + 0.02, wy0, wy1, wz0, wz0 + ft, mullion_mat)
    add_box("Window_Frame_Top", fx - 0.02, fx + 0.02, wy0, wy1, wz1 - ft, wz1, mullion_mat)
    add_box("Window_Frame_Front", fx - 0.02, fx + 0.02, wy0, wy0 + ft, wz0, wz1, mullion_mat)
    add_box("Window_Frame_Back", fx - 0.02, fx + 0.02, wy1 - ft, wy1, wz0, wz1, mullion_mat)
    ymid = 0.5 * (wy0 + wy1)
    add_box("Window_Mullion_Vertical", fx - 0.015, fx + 0.015, ymid - 0.014, ymid + 0.014, wz0 + ft, wz1 - ft, mullion_mat)
    z_transom = wz0 + 0.68 * (wz1 - wz0)
    add_box("Window_Mullion_Transom", fx - 0.015, fx + 0.015, wy0 + ft, wy1 - ft, z_transom - 0.012, z_transom + 0.012, mullion_mat)

    # 5. Subtle architectural skirting baseboards (8 cm high)
    bh = 0.085
    bt = 0.014
    add_box("Baseboard_Back", -hw, hw, hd - bt, hd, 0.0, bh, trim_mat)
    add_box("Baseboard_Left", -hw, -hw + bt, -hd, hd, 0.0, bh, trim_mat)
    add_box("Baseboard_Right", hw - bt, hw, -hd, hd, 0.0, bh, trim_mat)

    # 6. Center circular woven rug anchoring (0, 0) directly beneath the suspension
    add_cylinder_mesh(
        "Center_Floor_Rug",
        center=(0.0, 0.0, 0.003),
        radius=1.32,
        height=0.006,
        segments=96,
        collection=col_room,
        material=rug_mat,
        smooth=True,
    )

    return (reveal_x, 0.5 * (wy0 + wy1), 0.5 * (wz0 + wz1), wy1 - wy0, wz1 - wz0)


# ------------------------------------------------------------------------------
# CAMERA, WORLDS & SCENE LIGHTING
# ------------------------------------------------------------------------------

def create_shared_camera(col_cam):
    cam_loc = Vector((-0.52, -2.52, 1.66))
    cam_target = Vector((0.0, 0.0, SUSPENSION_CENTER_Z_M - 0.02))

    cam_data = bpy.data.cameras.new("Room_Suspension_Camera")
    cam_data.lens = 33.0
    cam_data.sensor_width = 36.0
    cam_data.clip_start = 0.05
    cam_data.clip_end = 50.0
    cam_data.dof.use_dof = True
    cam_data.dof.focus_distance = (cam_target - cam_loc).length
    cam_data.dof.aperture_fstop = 8.0

    cam_obj = bpy.data.objects.new("Room_Suspension_Camera", cam_data)
    cam_obj.location = cam_loc
    direction = cam_target - cam_loc
    cam_obj.rotation_euler = direction.to_track_quat("-Z", "Y").to_euler()
    col_cam.objects.link(cam_obj)
    return cam_obj


def create_world(name, color_rgba, strength):
    world = bpy.data.worlds.new(name)
    nt = ensure_node_tree(world)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()
    out = nodes.new("ShaderNodeOutputWorld")
    bg = nodes.new("ShaderNodeBackground")
    set_socket(bg, "Color", color_rgba)
    set_socket(bg, "Strength", strength)
    links.new(bg.outputs["Background"], out.inputs["Surface"])
    return world


def add_area_light(name, location, target, energy, color, size_x, size_y, collection, spread_deg=115.0):
    ldata = bpy.data.lights.new(name, type="AREA")
    ldata.shape = "RECTANGLE"
    ldata.size = size_x
    ldata.size_y = size_y
    ldata.energy = energy
    ldata.color = color
    if hasattr(ldata, "spread"):
        ldata.spread = math.radians(spread_deg)

    lobj = bpy.data.objects.new(name, ldata)
    lobj.location = Vector(location)
    direction = Vector(target) - Vector(location)
    lobj.rotation_euler = direction.to_track_quat("-Z", "Y").to_euler()
    collection.objects.link(lobj)
    return lobj


def populate_daylight_lights(col_day, win_info):
    wx, wy, wz, w_len, w_height = win_info
    # 1. Natural architectural daylight streaming through the right-wall window
    add_area_light(
        "Daylight_Window_Key",
        location=(wx + 0.18, wy, wz + 0.15),
        target=(-0.4, 0.15, SUSPENSION_CENTER_Z_M - 0.10),
        energy=260.0,
        color=(1.0, 0.965, 0.91),
        size_x=w_len * 0.92,
        size_y=w_height * 0.92,
        collection=col_day,
        spread_deg=105.0,
    )
    # 2. Soft cool sky fill from left-front of room so shadow troughs keep rich color nuance
    add_area_light(
        "Daylight_Room_Bounce_Fill",
        location=(-2.10, -1.80, 2.15),
        target=(0.0, 0.20, SUSPENSION_CENTER_Z_M),
        energy=65.0,
        color=(0.90, 0.95, 1.0),
        size_x=2.2,
        size_y=1.8,
        collection=col_day,
        spread_deg=140.0,
    )
    # 3. Gentle upward floor-bounce / ambient fill so ceiling canopy & upper room read airily
    add_area_light(
        "Daylight_Ambient_Ceiling_Fill",
        location=(-0.3, -1.4, 0.65),
        target=(0.0, 0.6, 2.75),
        energy=28.0,
        color=(0.98, 0.96, 0.93),
        size_x=2.4,
        size_y=2.4,
        collection=col_day,
        spread_deg=160.0,
    )


def populate_evening_lights(col_eve, win_info):
    wx, wy, wz, w_len, w_height = win_info
    # Quiet blue-hour twilight outside the window so the room architecture remains readable
    # while the warm 2800K inside-border LEDs act as the unmistakable hero luminaire!
    add_area_light(
        "Evening_Twilight_Window",
        location=(wx + 0.18, wy, wz),
        target=(-0.2, 0.0, 1.65),
        energy=24.0,
        color=(0.42, 0.56, 0.88),
        size_x=w_len * 0.92,
        size_y=w_height * 0.92,
        collection=col_eve,
        spread_deg=120.0,
    )
    add_area_light(
        "Evening_Subtle_Room_Fill",
        location=(-1.95, -2.05, 1.95),
        target=(0.0, 0.4, 1.75),
        energy=12.0,
        color=(0.62, 0.68, 0.84),
        size_x=2.2,
        size_y=1.8,
        collection=col_eve,
        spread_deg=155.0,
    )


# ------------------------------------------------------------------------------
# CYCLES & COLOR MANAGEMENT CONFIGURATION
# ------------------------------------------------------------------------------

def enable_cycles_gpu():
    try:
        cprefs = bpy.context.preferences.addons["cycles"].preferences
    except Exception:
        return "CPU"

    preferred_backends = ("METAL", "OPTIX", "CUDA", "HIP", "ONEAPI")
    available = []
    try:
        available = [item[0] for item in cprefs.get_device_types(bpy.context)]
    except Exception:
        available = list(preferred_backends)

    for backend in preferred_backends:
        if backend not in available:
            continue
        try:
            cprefs.compute_device_type = backend
            if hasattr(cprefs, "refresh_devices"):
                cprefs.refresh_devices()
            elif hasattr(cprefs, "get_devices"):
                cprefs.get_devices()

            found_gpu = False
            for dev in cprefs.devices:
                if dev.type == backend:
                    dev.use = True
                    found_gpu = True
                elif dev.type == "CPU":
                    dev.use = False
            if found_gpu:
                return backend
        except Exception:
            continue

    return "CPU"


def configure_scene_render(scene, camera_obj, world, samples, exposure=0.0):
    scene.camera = camera_obj
    scene.world = world

    rd = scene.render
    rd.engine = "CYCLES"
    rd.resolution_x = RENDER_RESOLUTION_X
    rd.resolution_y = RENDER_RESOLUTION_Y
    rd.resolution_percentage = 100
    rd.image_settings.file_format = "PNG"
    rd.image_settings.color_mode = "RGBA"
    rd.image_settings.color_depth = "8"

    # AgX tone-mapping preserves rich pigment saturation and smooth LED highlight roll-off
    vs = scene.view_settings
    for vt in ("AgX", "Filmic", "Standard"):
        try:
            vs.view_transform = vt
            break
        except Exception:
            continue
    for lk in ("AgX - Medium High Contrast", "Medium High Contrast", "None"):
        try:
            vs.look = lk
            break
        except Exception:
            continue
    vs.exposure = exposure

    cscene = scene.cycles
    cscene.samples = samples
    cscene.use_adaptive_sampling = True
    cscene.adaptive_threshold = 0.012
    cscene.use_denoising = True
    if hasattr(cscene, "use_preview_denoising"):
        cscene.use_preview_denoising = True
    if hasattr(cscene, "use_light_tree"):
        cscene.use_light_tree = True
    cscene.max_bounces = 8
    cscene.diffuse_bounces = 4
    cscene.glossy_bounces = 4
    cscene.transmission_bounces = 6
    cscene.sample_clamp_indirect = 10.0


def resolve_output_directory():
    # 1. If opened from a saved .py file in Blender's Text Editor, save right next to it
    try:
        space = getattr(bpy.context, "space_data", None)
        if space is not None and getattr(space, "text", None) is not None and space.text.filepath:
            p = bpy.path.abspath(space.text.filepath)
            if os.path.isdir(os.path.dirname(p)):
                return os.path.dirname(p)
    except Exception:
        pass

    # 2. If executed via --python <path/to/script.py>
    try:
        if "__file__" in globals() and __file__:
            d = os.path.dirname(os.path.abspath(__file__))
            if os.path.isdir(d) and os.access(d, os.W_OK):
                return d
    except Exception:
        pass

    # 3. Fallback to ~/Downloads or /tmp
    dl = os.path.expanduser("~/Downloads")
    if os.path.isdir(dl) and os.access(dl, os.W_OK):
        return dl
    return "/tmp"


def build_and_render():
    init_scene = clear_all_datablocks()
    gpu_backend = enable_cycles_gpu()

    # Shared collections
    col_room = make_collection("Room_Architecture")
    col_rings = make_collection("Suspension_Rings")
    col_hardware = make_collection("Suspension_Hardware")
    col_cam = make_collection("Cameras")

    # Mode-specific collections
    col_leds = make_collection("Suspension_Inside_LEDs")
    col_lights_day = make_collection("Lights_01_Pure_Daylight")
    col_lights_eve = make_collection("Lights_02_LED_Evening")

    cam_obj = create_shared_camera(col_cam)
    win_info = build_architectural_room(col_room)
    build_suspension_geometry(col_rings, col_leds, col_hardware, (cam_obj.location.x, cam_obj.location.y))

    populate_daylight_lights(col_lights_day, win_info)
    populate_evening_lights(col_lights_eve, win_info)

    world_day = create_world("World_01_Daylight", (0.85, 0.91, 1.0, 1.0), strength=1.8)
    world_eve = create_world("World_02_Evening", (0.10, 0.16, 0.32, 1.0), strength=0.35)

    # Scene 1: Pure sculptural suspension without lights in it
    scene_pure = init_scene
    scene_pure.name = "01_Pure_Suspension"
    for c in (col_room, col_rings, col_hardware, col_cam, col_lights_day):
        scene_pure.collection.children.link(c)
    configure_scene_render(scene_pure, cam_obj, world_day, CYCLES_SAMPLES_PURE, exposure=0.10)
    if gpu_backend != "CPU":
        scene_pure.cycles.device = "GPU"

    # Scene 2: Suspension with warm LEDs integrated in the inside border
    scene_leds = bpy.data.scenes.new("02_Suspension_Inside_LEDs")
    for c in (col_room, col_rings, col_hardware, col_leds, col_cam, col_lights_eve):
        scene_leds.collection.children.link(c)
    configure_scene_render(scene_leds, cam_obj, world_eve, CYCLES_SAMPLES_LEDS, exposure=0.25)
    if gpu_backend != "CPU":
        scene_leds.cycles.device = "GPU"

    out_dir = resolve_output_directory()
    path_pure = os.path.join(out_dir, f"{PLANT_SLUG}_01_pure_suspension.png")
    path_leds = os.path.join(out_dir, f"{PLANT_SLUG}_02_inside_leds.png")

    scene_pure.render.filepath = path_pure
    scene_leds.render.filepath = path_leds

    print(f"[Phylla Suspension] Built scenes using Cycles ({gpu_backend}).")
    print(f"  Scene 1 (Pure):       {path_pure}")
    print(f"  Scene 2 (Inside LED): {path_leds}")

    if AUTO_RENDER_BOTH_IMAGES:
        bpy.ops.render.render(write_still=True, scene=scene_pure.name)
        if os.path.isfile(path_pure):
            img1 = bpy.data.images.load(path_pure, check_existing=True)
            img1.name = "Render_01_Pure_Suspension"

        bpy.ops.render.render(write_still=True, scene=scene_leds.name)
        if os.path.isfile(path_leds):
            img2 = bpy.data.images.load(path_leds, check_existing=True)
            img2.name = "Render_02_Inside_LEDs"

        print("[Phylla Suspension] Finished rendering both images!")

    # Set active window scene to Scene 2 (LEDs) or Scene 1 and align 3D viewports to Camera view
    if bpy.context.window is not None:
        bpy.context.window.scene = scene_leds
        for screen in bpy.data.screens:
            for area in screen.areas:
                if area.type == "VIEW_3D":
                    for space in area.spaces:
                        if space.type == "VIEW_3D" and space.region_3d is not None:
                            space.region_3d.view_perspective = "CAMERA"


if __name__ == "__main__":
    build_and_render()
`

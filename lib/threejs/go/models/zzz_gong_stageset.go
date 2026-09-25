// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, ambiantlight := range __gong__sortStageSetInstances(stageSet.Stage.AmbiantLights, stageSet.Stage.AmbiantLight_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ambiantlightIdent := "__models" + ambiantlight.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AmbiantLight{Name: %s}).Stage(stageSet.Stage)", ambiantlightIdent, __gong__toRawStringLiteral(ambiantlight.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ambiantlightIdent, __gong__toRawStringLiteral(ambiantlight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Intensity = %f", ambiantlightIdent, ambiantlight.Intensity))
		}
	}
	if stageSet.Stage != nil {
		for _, boxgeometry := range __gong__sortStageSetInstances(stageSet.Stage.BoxGeometrys, stageSet.Stage.BoxGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			boxgeometryIdent := "__models" + boxgeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BoxGeometry{Name: %s}).Stage(stageSet.Stage)", boxgeometryIdent, __gong__toRawStringLiteral(boxgeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", boxgeometryIdent, __gong__toRawStringLiteral(boxgeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", boxgeometryIdent, boxgeometry.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", boxgeometryIdent, boxgeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %f", boxgeometryIdent, boxgeometry.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", boxgeometryIdent, boxgeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", boxgeometryIdent, boxgeometry.HeightSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.DepthSegments = %d", boxgeometryIdent, boxgeometry.DepthSegments))
		}
	}
	if stageSet.Stage != nil {
		for _, buffergeometry := range __gong__sortStageSetInstances(stageSet.Stage.BufferGeometrys, stageSet.Stage.BufferGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buffergeometryIdent := "__models" + buffergeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BufferGeometry{Name: %s}).Stage(stageSet.Stage)", buffergeometryIdent, __gong__toRawStringLiteral(buffergeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buffergeometryIdent, __gong__toRawStringLiteral(buffergeometry.Name)))
			for _, elem := range buffergeometry.Vertices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vertices = append(%s.Vertices, %s)", buffergeometryIdent, buffergeometryIdent, targetIdent))
			}
			for _, elem := range buffergeometry.Faces {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Faces = append(%s.Faces, %s)", buffergeometryIdent, buffergeometryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, camera := range __gong__sortStageSetInstances(stageSet.Stage.Cameras, stageSet.Stage.Camera_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cameraIdent := "__models" + camera.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Camera{Name: %s}).Stage(stageSet.Stage)", cameraIdent, __gong__toRawStringLiteral(camera.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cameraIdent, __gong__toRawStringLiteral(camera.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", cameraIdent, camera.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", cameraIdent, camera.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", cameraIdent, camera.Z))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetX = %f", cameraIdent, camera.TargetX))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetY = %f", cameraIdent, camera.TargetY))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetZ = %f", cameraIdent, camera.TargetZ))
			values.WriteString(fmt.Sprintf("\n\t%s.Fov = %f", cameraIdent, camera.Fov))
		}
	}
	if stageSet.Stage != nil {
		for _, canvas := range __gong__sortStageSetInstances(stageSet.Stage.Canvass, stageSet.Stage.Canvas_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			canvasIdent := "__models" + canvas.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Canvas{Name: %s}).Stage(stageSet.Stage)", canvasIdent, __gong__toRawStringLiteral(canvas.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", canvasIdent, __gong__toRawStringLiteral(canvas.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithLastRenderingUpdate = %t", canvasIdent, canvas.IsWithLastRenderingUpdate))
			values.WriteString(fmt.Sprintf("\n\t%s.LastRendering, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", canvasIdent, canvas.LastRendering.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame64BitsEncoded = %s", canvasIdent, __gong__toRawStringLiteral(canvas.Frame64BitsEncoded)))
			for _, elem := range canvas.DirectionalLights {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionalLights = append(%s.DirectionalLights, %s)", canvasIdent, canvasIdent, targetIdent))
			}
			if canvas.AmbiantLight != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + canvas.AmbiantLight.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AmbiantLight = %s", canvasIdent, targetIdent))
			}
			for _, elem := range canvas.Meshs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Meshs = append(%s.Meshs, %s)", canvasIdent, canvasIdent, targetIdent))
			}
			if canvas.Camera != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + canvas.Camera.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Camera = %s", canvasIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, curve := range __gong__sortStageSetInstances(stageSet.Stage.Curves, stageSet.Stage.Curve_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curveIdent := "__models" + curve.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Curve{Name: %s}).Stage(stageSet.Stage)", curveIdent, __gong__toRawStringLiteral(curve.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curveIdent, __gong__toRawStringLiteral(curve.Name)))
			for _, elem := range curve.Points {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Points = append(%s.Points, %s)", curveIdent, curveIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, cylindergeometry := range __gong__sortStageSetInstances(stageSet.Stage.CylinderGeometrys, stageSet.Stage.CylinderGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cylindergeometryIdent := "__models" + cylindergeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CylinderGeometry{Name: %s}).Stage(stageSet.Stage)", cylindergeometryIdent, __gong__toRawStringLiteral(cylindergeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cylindergeometryIdent, __gong__toRawStringLiteral(cylindergeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RadiusTop = %f", cylindergeometryIdent, cylindergeometry.RadiusTop))
			values.WriteString(fmt.Sprintf("\n\t%s.RadiusBottom = %f", cylindergeometryIdent, cylindergeometry.RadiusBottom))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", cylindergeometryIdent, cylindergeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", cylindergeometryIdent, cylindergeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", cylindergeometryIdent, cylindergeometry.HeightSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.OpenEnded = %t", cylindergeometryIdent, cylindergeometry.OpenEnded))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaStart = %f", cylindergeometryIdent, cylindergeometry.ThetaStart))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaLength = %f", cylindergeometryIdent, cylindergeometry.ThetaLength))
		}
	}
	if stageSet.Stage != nil {
		for _, directionallight := range __gong__sortStageSetInstances(stageSet.Stage.DirectionalLights, stageSet.Stage.DirectionalLight_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionallightIdent := "__models" + directionallight.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DirectionalLight{Name: %s}).Stage(stageSet.Stage)", directionallightIdent, __gong__toRawStringLiteral(directionallight.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionallightIdent, __gong__toRawStringLiteral(directionallight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", directionallightIdent, directionallight.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", directionallightIdent, directionallight.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", directionallightIdent, directionallight.Z))
			values.WriteString(fmt.Sprintf("\n\t%s.Intensity = %f", directionallightIdent, directionallight.Intensity))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCastShadow = %t", directionallightIdent, directionallight.IsWithCastShadow))
		}
	}
	if stageSet.Stage != nil {
		for _, extrudegeometry := range __gong__sortStageSetInstances(stageSet.Stage.ExtrudeGeometrys, stageSet.Stage.ExtrudeGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extrudegeometryIdent := "__models" + extrudegeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ExtrudeGeometry{Name: %s}).Stage(stageSet.Stage)", extrudegeometryIdent, __gong__toRawStringLiteral(extrudegeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extrudegeometryIdent, __gong__toRawStringLiteral(extrudegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steps = %d", extrudegeometryIdent, extrudegeometry.Steps))
			if extrudegeometry.Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + extrudegeometry.Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Shape = %s", extrudegeometryIdent, targetIdent))
			}
			if extrudegeometry.ExtrudePath != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + extrudegeometry.ExtrudePath.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtrudePath = %s", extrudegeometryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, mesh := range __gong__sortStageSetInstances(stageSet.Stage.Meshs, stageSet.Stage.Mesh_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meshIdent := "__models" + mesh.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Mesh{Name: %s}).Stage(stageSet.Stage)", meshIdent, __gong__toRawStringLiteral(mesh.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshIdent, __gong__toRawStringLiteral(mesh.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", meshIdent, mesh.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", meshIdent, mesh.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", meshIdent, mesh.Z))
			if mesh.MeshMaterialBasic != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.MeshMaterialBasic.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeshMaterialBasic = %s", meshIdent, targetIdent))
			}
			if mesh.MeshPhysicalMaterial != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.MeshPhysicalMaterial.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeshPhysicalMaterial = %s", meshIdent, targetIdent))
			}
			if mesh.CylinderGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.CylinderGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CylinderGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.BoxGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.BoxGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoxGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.SphereGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.SphereGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SphereGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.TorusGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.TorusGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TorusGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.PlaneGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.PlaneGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PlaneGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.TubeGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.TubeGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.ExtrudeGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.ExtrudeGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtrudeGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.BufferGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mesh.BufferGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BufferGeometry = %s", meshIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, meshmaterialbasic := range __gong__sortStageSetInstances(stageSet.Stage.MeshMaterialBasics, stageSet.Stage.MeshMaterialBasic_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meshmaterialbasicIdent := "__models" + meshmaterialbasic.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MeshMaterialBasic{Name: %s}).Stage(stageSet.Stage)", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, meshphysicalmaterial := range __gong__sortStageSetInstances(stageSet.Stage.MeshPhysicalMaterials, stageSet.Stage.MeshPhysicalMaterial_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meshphysicalmaterialIdent := "__models" + meshphysicalmaterial.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MeshPhysicalMaterial{Name: %s}).Stage(stageSet.Stage)", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Wireframe = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Wireframe))
			values.WriteString(fmt.Sprintf("\n\t%s.Opacity = %f", meshphysicalmaterialIdent, meshphysicalmaterial.Opacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparent = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Transparent))
			values.WriteString(fmt.Sprintf("\n\t%s.Visible = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Visible))
		}
	}
	if stageSet.Stage != nil {
		for _, planegeometry := range __gong__sortStageSetInstances(stageSet.Stage.PlaneGeometrys, stageSet.Stage.PlaneGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			planegeometryIdent := "__models" + planegeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PlaneGeometry{Name: %s}).Stage(stageSet.Stage)", planegeometryIdent, __gong__toRawStringLiteral(planegeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", planegeometryIdent, __gong__toRawStringLiteral(planegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", planegeometryIdent, planegeometry.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", planegeometryIdent, planegeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", planegeometryIdent, planegeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", planegeometryIdent, planegeometry.HeightSegments))
		}
	}
	if stageSet.Stage != nil {
		for _, shape := range __gong__sortStageSetInstances(stageSet.Stage.Shapes, stageSet.Stage.Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			shapeIdent := "__models" + shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Shape{Name: %s}).Stage(stageSet.Stage)", shapeIdent, __gong__toRawStringLiteral(shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", shapeIdent, __gong__toRawStringLiteral(shape.Name)))
			for _, elem := range shape.Points {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Points = append(%s.Points, %s)", shapeIdent, shapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, spheregeometry := range __gong__sortStageSetInstances(stageSet.Stage.SphereGeometrys, stageSet.Stage.SphereGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spheregeometryIdent := "__models" + spheregeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SphereGeometry{Name: %s}).Stage(stageSet.Stage)", spheregeometryIdent, __gong__toRawStringLiteral(spheregeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spheregeometryIdent, __gong__toRawStringLiteral(spheregeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", spheregeometryIdent, spheregeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", spheregeometryIdent, spheregeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", spheregeometryIdent, spheregeometry.HeightSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.PhiStart = %f", spheregeometryIdent, spheregeometry.PhiStart))
			values.WriteString(fmt.Sprintf("\n\t%s.PhiLength = %f", spheregeometryIdent, spheregeometry.PhiLength))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaStart = %f", spheregeometryIdent, spheregeometry.ThetaStart))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaLength = %f", spheregeometryIdent, spheregeometry.ThetaLength))
		}
	}
	if stageSet.Stage != nil {
		for _, torusgeometry := range __gong__sortStageSetInstances(stageSet.Stage.TorusGeometrys, stageSet.Stage.TorusGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			torusgeometryIdent := "__models" + torusgeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TorusGeometry{Name: %s}).Stage(stageSet.Stage)", torusgeometryIdent, __gong__toRawStringLiteral(torusgeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", torusgeometryIdent, __gong__toRawStringLiteral(torusgeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", torusgeometryIdent, torusgeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.Tube = %f", torusgeometryIdent, torusgeometry.Tube))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", torusgeometryIdent, torusgeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.TubularSegments = %d", torusgeometryIdent, torusgeometry.TubularSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Arc = %f", torusgeometryIdent, torusgeometry.Arc))
		}
	}
	if stageSet.Stage != nil {
		for _, triangle := range __gong__sortStageSetInstances(stageSet.Stage.Triangles, stageSet.Stage.Triangle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			triangleIdent := "__models" + triangle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Triangle{Name: %s}).Stage(stageSet.Stage)", triangleIdent, __gong__toRawStringLiteral(triangle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", triangleIdent, __gong__toRawStringLiteral(triangle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.V1 = %d", triangleIdent, triangle.V1))
			values.WriteString(fmt.Sprintf("\n\t%s.V2 = %d", triangleIdent, triangle.V2))
			values.WriteString(fmt.Sprintf("\n\t%s.V3 = %d", triangleIdent, triangle.V3))
		}
	}
	if stageSet.Stage != nil {
		for _, tubegeometry := range __gong__sortStageSetInstances(stageSet.Stage.TubeGeometrys, stageSet.Stage.TubeGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tubegeometryIdent := "__models" + tubegeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TubeGeometry{Name: %s}).Stage(stageSet.Stage)", tubegeometryIdent, __gong__toRawStringLiteral(tubegeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubegeometryIdent, __gong__toRawStringLiteral(tubegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.TubularSegments = %d", tubegeometryIdent, tubegeometry.TubularSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", tubegeometryIdent, tubegeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", tubegeometryIdent, tubegeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Closed = %t", tubegeometryIdent, tubegeometry.Closed))
			if tubegeometry.Path != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubegeometry.Path.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Path = %s", tubegeometryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, vector2 := range __gong__sortStageSetInstances(stageSet.Stage.Vector2s, stageSet.Stage.Vector2_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vector2Ident := "__models" + vector2.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Vector2{Name: %s}).Stage(stageSet.Stage)", vector2Ident, __gong__toRawStringLiteral(vector2.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vector2Ident, __gong__toRawStringLiteral(vector2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", vector2Ident, vector2.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", vector2Ident, vector2.Y))
		}
	}
	if stageSet.Stage != nil {
		for _, vector3 := range __gong__sortStageSetInstances(stageSet.Stage.Vector3s, stageSet.Stage.Vector3_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vector3Ident := "__models" + vector3.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Vector3{Name: %s}).Stage(stageSet.Stage)", vector3Ident, __gong__toRawStringLiteral(vector3.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vector3Ident, __gong__toRawStringLiteral(vector3.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", vector3Ident, vector3.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", vector3Ident, vector3.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", vector3Ident, vector3.Z))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/lib/threejs/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "AmbiantLight":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AmbiantLight), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BoxGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BoxGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BufferGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BufferGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Camera":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Camera), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Canvas":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Canvas), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Curve":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Curve), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CylinderGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CylinderGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DirectionalLight":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectionalLight), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ExtrudeGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ExtrudeGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Mesh":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Mesh), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeshMaterialBasic":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeshMaterialBasic), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeshPhysicalMaterial":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeshPhysicalMaterial), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PlaneGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PlaneGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SphereGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SphereGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TorusGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TorusGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Triangle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Triangle), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TubeGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TubeGeometry), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Vector2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Vector2), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Vector3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Vector3), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *AmbiantLight:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Intensity":
						inst.Intensity = GongExtractFloat(rhs)
					}
				case *BoxGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "Depth":
						inst.Depth = GongExtractFloat(rhs)
					case "WidthSegments":
						inst.WidthSegments = GongExtractInt(rhs)
					case "HeightSegments":
						inst.HeightSegments = GongExtractInt(rhs)
					case "DepthSegments":
						inst.DepthSegments = GongExtractInt(rhs)
					}
				case *BufferGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Vertices":
						__gong__assignSliceOfPointers(&inst.Vertices, rhs, identifierMap)
					case "Faces":
						__gong__assignSliceOfPointers(&inst.Faces, rhs, identifierMap)
					}
				case *Camera:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Z":
						inst.Z = GongExtractFloat(rhs)
					case "TargetX":
						inst.TargetX = GongExtractFloat(rhs)
					case "TargetY":
						inst.TargetY = GongExtractFloat(rhs)
					case "TargetZ":
						inst.TargetZ = GongExtractFloat(rhs)
					case "Fov":
						inst.Fov = GongExtractFloat(rhs)
					}
				case *Canvas:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DirectionalLights":
						__gong__assignSliceOfPointers(&inst.DirectionalLights, rhs, identifierMap)
					case "AmbiantLight":
						__gong__assignPointer(&inst.AmbiantLight, rhs, identifierMap)
					case "Meshs":
						__gong__assignSliceOfPointers(&inst.Meshs, rhs, identifierMap)
					case "Camera":
						__gong__assignPointer(&inst.Camera, rhs, identifierMap)
					case "IsWithLastRenderingUpdate":
						inst.IsWithLastRenderingUpdate = GongExtractBool(rhs)
					case "LastRendering":
						inst.LastRendering = GongExtractDate(rhs)
					case "Frame64BitsEncoded":
						inst.Frame64BitsEncoded = GongExtractString(rhs)
					}
				case *Curve:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						__gong__assignSliceOfPointers(&inst.Points, rhs, identifierMap)
					}
				case *CylinderGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RadiusTop":
						inst.RadiusTop = GongExtractFloat(rhs)
					case "RadiusBottom":
						inst.RadiusBottom = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RadialSegments":
						inst.RadialSegments = GongExtractInt(rhs)
					case "HeightSegments":
						inst.HeightSegments = GongExtractInt(rhs)
					case "OpenEnded":
						inst.OpenEnded = GongExtractBool(rhs)
					case "ThetaStart":
						inst.ThetaStart = GongExtractFloat(rhs)
					case "ThetaLength":
						inst.ThetaLength = GongExtractFloat(rhs)
					}
				case *DirectionalLight:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Z":
						inst.Z = GongExtractFloat(rhs)
					case "Intensity":
						inst.Intensity = GongExtractFloat(rhs)
					case "IsWithCastShadow":
						inst.IsWithCastShadow = GongExtractBool(rhs)
					}
				case *ExtrudeGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Shape":
						__gong__assignPointer(&inst.Shape, rhs, identifierMap)
					case "ExtrudePath":
						__gong__assignPointer(&inst.ExtrudePath, rhs, identifierMap)
					case "Steps":
						inst.Steps = GongExtractInt(rhs)
					}
				case *Mesh:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Z":
						inst.Z = GongExtractFloat(rhs)
					case "MeshMaterialBasic":
						__gong__assignPointer(&inst.MeshMaterialBasic, rhs, identifierMap)
					case "MeshPhysicalMaterial":
						__gong__assignPointer(&inst.MeshPhysicalMaterial, rhs, identifierMap)
					case "CylinderGeometry":
						__gong__assignPointer(&inst.CylinderGeometry, rhs, identifierMap)
					case "BoxGeometry":
						__gong__assignPointer(&inst.BoxGeometry, rhs, identifierMap)
					case "SphereGeometry":
						__gong__assignPointer(&inst.SphereGeometry, rhs, identifierMap)
					case "TorusGeometry":
						__gong__assignPointer(&inst.TorusGeometry, rhs, identifierMap)
					case "PlaneGeometry":
						__gong__assignPointer(&inst.PlaneGeometry, rhs, identifierMap)
					case "TubeGeometry":
						__gong__assignPointer(&inst.TubeGeometry, rhs, identifierMap)
					case "ExtrudeGeometry":
						__gong__assignPointer(&inst.ExtrudeGeometry, rhs, identifierMap)
					case "BufferGeometry":
						__gong__assignPointer(&inst.BufferGeometry, rhs, identifierMap)
					}
				case *MeshMaterialBasic:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *MeshPhysicalMaterial:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Wireframe":
						inst.Wireframe = GongExtractBool(rhs)
					case "Opacity":
						inst.Opacity = GongExtractFloat(rhs)
					case "Transparent":
						inst.Transparent = GongExtractBool(rhs)
					case "Visible":
						inst.Visible = GongExtractBool(rhs)
					}
				case *PlaneGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "WidthSegments":
						inst.WidthSegments = GongExtractInt(rhs)
					case "HeightSegments":
						inst.HeightSegments = GongExtractInt(rhs)
					}
				case *Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						__gong__assignSliceOfPointers(&inst.Points, rhs, identifierMap)
					}
				case *SphereGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Radius":
						inst.Radius = GongExtractFloat(rhs)
					case "WidthSegments":
						inst.WidthSegments = GongExtractInt(rhs)
					case "HeightSegments":
						inst.HeightSegments = GongExtractInt(rhs)
					case "PhiStart":
						inst.PhiStart = GongExtractFloat(rhs)
					case "PhiLength":
						inst.PhiLength = GongExtractFloat(rhs)
					case "ThetaStart":
						inst.ThetaStart = GongExtractFloat(rhs)
					case "ThetaLength":
						inst.ThetaLength = GongExtractFloat(rhs)
					}
				case *TorusGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Radius":
						inst.Radius = GongExtractFloat(rhs)
					case "Tube":
						inst.Tube = GongExtractFloat(rhs)
					case "RadialSegments":
						inst.RadialSegments = GongExtractInt(rhs)
					case "TubularSegments":
						inst.TubularSegments = GongExtractInt(rhs)
					case "Arc":
						inst.Arc = GongExtractFloat(rhs)
					}
				case *Triangle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "V1":
						inst.V1 = GongExtractInt(rhs)
					case "V2":
						inst.V2 = GongExtractInt(rhs)
					case "V3":
						inst.V3 = GongExtractInt(rhs)
					}
				case *TubeGeometry:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Path":
						__gong__assignPointer(&inst.Path, rhs, identifierMap)
					case "TubularSegments":
						inst.TubularSegments = GongExtractInt(rhs)
					case "Radius":
						inst.Radius = GongExtractFloat(rhs)
					case "RadialSegments":
						inst.RadialSegments = GongExtractInt(rhs)
					case "Closed":
						inst.Closed = GongExtractBool(rhs)
					}
				case *Vector2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					}
				case *Vector3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Z":
						inst.Z = GongExtractFloat(rhs)
					}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}

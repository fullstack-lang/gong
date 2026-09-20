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

	if stageSet.Stage != nil {
		ambiantlightOrdered := []*AmbiantLight{}
		for ambiantlight := range stageSet.Stage.AmbiantLights {
			ambiantlightOrdered = append(ambiantlightOrdered, ambiantlight)
		}
		sort.Slice(ambiantlightOrdered, func(i, j int) bool {
			return stageSet.Stage.AmbiantLight_stagedOrder[ambiantlightOrdered[i]] < stageSet.Stage.AmbiantLight_stagedOrder[ambiantlightOrdered[j]]
		})
		for _, ambiantlight := range ambiantlightOrdered {
			ambiantlightIdent := "__stage_0" + ambiantlight.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AmbiantLight{Name: %s}).Stage(stageSet.Stage)", ambiantlightIdent, __gong__toRawStringLiteral(ambiantlight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ambiantlightIdent, __gong__toRawStringLiteral(ambiantlight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Intensity = %f", ambiantlightIdent, ambiantlight.Intensity))
		}
	}
	if stageSet.Stage != nil {
		boxgeometryOrdered := []*BoxGeometry{}
		for boxgeometry := range stageSet.Stage.BoxGeometrys {
			boxgeometryOrdered = append(boxgeometryOrdered, boxgeometry)
		}
		sort.Slice(boxgeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.BoxGeometry_stagedOrder[boxgeometryOrdered[i]] < stageSet.Stage.BoxGeometry_stagedOrder[boxgeometryOrdered[j]]
		})
		for _, boxgeometry := range boxgeometryOrdered {
			boxgeometryIdent := "__stage_0" + boxgeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.BoxGeometry{Name: %s}).Stage(stageSet.Stage)", boxgeometryIdent, __gong__toRawStringLiteral(boxgeometry.Name)))
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
		buffergeometryOrdered := []*BufferGeometry{}
		for buffergeometry := range stageSet.Stage.BufferGeometrys {
			buffergeometryOrdered = append(buffergeometryOrdered, buffergeometry)
		}
		sort.Slice(buffergeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.BufferGeometry_stagedOrder[buffergeometryOrdered[i]] < stageSet.Stage.BufferGeometry_stagedOrder[buffergeometryOrdered[j]]
		})
		for _, buffergeometry := range buffergeometryOrdered {
			buffergeometryIdent := "__stage_0" + buffergeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.BufferGeometry{Name: %s}).Stage(stageSet.Stage)", buffergeometryIdent, __gong__toRawStringLiteral(buffergeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buffergeometryIdent, __gong__toRawStringLiteral(buffergeometry.Name)))
			for _, elem := range buffergeometry.Vertices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vertices = append(%s.Vertices, %s)", buffergeometryIdent, buffergeometryIdent, targetIdent))
			}
			for _, elem := range buffergeometry.Faces {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Faces = append(%s.Faces, %s)", buffergeometryIdent, buffergeometryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		cameraOrdered := []*Camera{}
		for camera := range stageSet.Stage.Cameras {
			cameraOrdered = append(cameraOrdered, camera)
		}
		sort.Slice(cameraOrdered, func(i, j int) bool {
			return stageSet.Stage.Camera_stagedOrder[cameraOrdered[i]] < stageSet.Stage.Camera_stagedOrder[cameraOrdered[j]]
		})
		for _, camera := range cameraOrdered {
			cameraIdent := "__stage_0" + camera.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Camera{Name: %s}).Stage(stageSet.Stage)", cameraIdent, __gong__toRawStringLiteral(camera.Name)))
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
		canvasOrdered := []*Canvas{}
		for canvas := range stageSet.Stage.Canvass {
			canvasOrdered = append(canvasOrdered, canvas)
		}
		sort.Slice(canvasOrdered, func(i, j int) bool {
			return stageSet.Stage.Canvas_stagedOrder[canvasOrdered[i]] < stageSet.Stage.Canvas_stagedOrder[canvasOrdered[j]]
		})
		for _, canvas := range canvasOrdered {
			canvasIdent := "__stage_0" + canvas.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Canvas{Name: %s}).Stage(stageSet.Stage)", canvasIdent, __gong__toRawStringLiteral(canvas.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", canvasIdent, __gong__toRawStringLiteral(canvas.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithLastRenderingUpdate = %t", canvasIdent, canvas.IsWithLastRenderingUpdate))
			values.WriteString(fmt.Sprintf("\n\t%s.LastRendering, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", canvasIdent, canvas.LastRendering.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame64BitsEncoded = %s", canvasIdent, __gong__toRawStringLiteral(canvas.Frame64BitsEncoded)))
			for _, elem := range canvas.DirectionalLights {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionalLights = append(%s.DirectionalLights, %s)", canvasIdent, canvasIdent, targetIdent))
			}
			if canvas.AmbiantLight != nil {
				targetIdent := "__stage_0" + canvas.AmbiantLight.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AmbiantLight = %s", canvasIdent, targetIdent))
			}
			for _, elem := range canvas.Meshs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Meshs = append(%s.Meshs, %s)", canvasIdent, canvasIdent, targetIdent))
			}
			if canvas.Camera != nil {
				targetIdent := "__stage_0" + canvas.Camera.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Camera = %s", canvasIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		curveOrdered := []*Curve{}
		for curve := range stageSet.Stage.Curves {
			curveOrdered = append(curveOrdered, curve)
		}
		sort.Slice(curveOrdered, func(i, j int) bool {
			return stageSet.Stage.Curve_stagedOrder[curveOrdered[i]] < stageSet.Stage.Curve_stagedOrder[curveOrdered[j]]
		})
		for _, curve := range curveOrdered {
			curveIdent := "__stage_0" + curve.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Curve{Name: %s}).Stage(stageSet.Stage)", curveIdent, __gong__toRawStringLiteral(curve.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curveIdent, __gong__toRawStringLiteral(curve.Name)))
			for _, elem := range curve.Points {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Points = append(%s.Points, %s)", curveIdent, curveIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		cylindergeometryOrdered := []*CylinderGeometry{}
		for cylindergeometry := range stageSet.Stage.CylinderGeometrys {
			cylindergeometryOrdered = append(cylindergeometryOrdered, cylindergeometry)
		}
		sort.Slice(cylindergeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.CylinderGeometry_stagedOrder[cylindergeometryOrdered[i]] < stageSet.Stage.CylinderGeometry_stagedOrder[cylindergeometryOrdered[j]]
		})
		for _, cylindergeometry := range cylindergeometryOrdered {
			cylindergeometryIdent := "__stage_0" + cylindergeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CylinderGeometry{Name: %s}).Stage(stageSet.Stage)", cylindergeometryIdent, __gong__toRawStringLiteral(cylindergeometry.Name)))
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
		directionallightOrdered := []*DirectionalLight{}
		for directionallight := range stageSet.Stage.DirectionalLights {
			directionallightOrdered = append(directionallightOrdered, directionallight)
		}
		sort.Slice(directionallightOrdered, func(i, j int) bool {
			return stageSet.Stage.DirectionalLight_stagedOrder[directionallightOrdered[i]] < stageSet.Stage.DirectionalLight_stagedOrder[directionallightOrdered[j]]
		})
		for _, directionallight := range directionallightOrdered {
			directionallightIdent := "__stage_0" + directionallight.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DirectionalLight{Name: %s}).Stage(stageSet.Stage)", directionallightIdent, __gong__toRawStringLiteral(directionallight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionallightIdent, __gong__toRawStringLiteral(directionallight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", directionallightIdent, directionallight.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", directionallightIdent, directionallight.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", directionallightIdent, directionallight.Z))
			values.WriteString(fmt.Sprintf("\n\t%s.Intensity = %f", directionallightIdent, directionallight.Intensity))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCastShadow = %t", directionallightIdent, directionallight.IsWithCastShadow))
		}
	}
	if stageSet.Stage != nil {
		extrudegeometryOrdered := []*ExtrudeGeometry{}
		for extrudegeometry := range stageSet.Stage.ExtrudeGeometrys {
			extrudegeometryOrdered = append(extrudegeometryOrdered, extrudegeometry)
		}
		sort.Slice(extrudegeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.ExtrudeGeometry_stagedOrder[extrudegeometryOrdered[i]] < stageSet.Stage.ExtrudeGeometry_stagedOrder[extrudegeometryOrdered[j]]
		})
		for _, extrudegeometry := range extrudegeometryOrdered {
			extrudegeometryIdent := "__stage_0" + extrudegeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ExtrudeGeometry{Name: %s}).Stage(stageSet.Stage)", extrudegeometryIdent, __gong__toRawStringLiteral(extrudegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extrudegeometryIdent, __gong__toRawStringLiteral(extrudegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steps = %d", extrudegeometryIdent, extrudegeometry.Steps))
			if extrudegeometry.Shape != nil {
				targetIdent := "__stage_0" + extrudegeometry.Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Shape = %s", extrudegeometryIdent, targetIdent))
			}
			if extrudegeometry.ExtrudePath != nil {
				targetIdent := "__stage_0" + extrudegeometry.ExtrudePath.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtrudePath = %s", extrudegeometryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		meshOrdered := []*Mesh{}
		for mesh := range stageSet.Stage.Meshs {
			meshOrdered = append(meshOrdered, mesh)
		}
		sort.Slice(meshOrdered, func(i, j int) bool {
			return stageSet.Stage.Mesh_stagedOrder[meshOrdered[i]] < stageSet.Stage.Mesh_stagedOrder[meshOrdered[j]]
		})
		for _, mesh := range meshOrdered {
			meshIdent := "__stage_0" + mesh.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Mesh{Name: %s}).Stage(stageSet.Stage)", meshIdent, __gong__toRawStringLiteral(mesh.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshIdent, __gong__toRawStringLiteral(mesh.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", meshIdent, mesh.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", meshIdent, mesh.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", meshIdent, mesh.Z))
			if mesh.MeshMaterialBasic != nil {
				targetIdent := "__stage_0" + mesh.MeshMaterialBasic.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeshMaterialBasic = %s", meshIdent, targetIdent))
			}
			if mesh.MeshPhysicalMaterial != nil {
				targetIdent := "__stage_0" + mesh.MeshPhysicalMaterial.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeshPhysicalMaterial = %s", meshIdent, targetIdent))
			}
			if mesh.CylinderGeometry != nil {
				targetIdent := "__stage_0" + mesh.CylinderGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CylinderGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.BoxGeometry != nil {
				targetIdent := "__stage_0" + mesh.BoxGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoxGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.SphereGeometry != nil {
				targetIdent := "__stage_0" + mesh.SphereGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SphereGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.TorusGeometry != nil {
				targetIdent := "__stage_0" + mesh.TorusGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TorusGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.PlaneGeometry != nil {
				targetIdent := "__stage_0" + mesh.PlaneGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PlaneGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.TubeGeometry != nil {
				targetIdent := "__stage_0" + mesh.TubeGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.ExtrudeGeometry != nil {
				targetIdent := "__stage_0" + mesh.ExtrudeGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtrudeGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.BufferGeometry != nil {
				targetIdent := "__stage_0" + mesh.BufferGeometry.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BufferGeometry = %s", meshIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		meshmaterialbasicOrdered := []*MeshMaterialBasic{}
		for meshmaterialbasic := range stageSet.Stage.MeshMaterialBasics {
			meshmaterialbasicOrdered = append(meshmaterialbasicOrdered, meshmaterialbasic)
		}
		sort.Slice(meshmaterialbasicOrdered, func(i, j int) bool {
			return stageSet.Stage.MeshMaterialBasic_stagedOrder[meshmaterialbasicOrdered[i]] < stageSet.Stage.MeshMaterialBasic_stagedOrder[meshmaterialbasicOrdered[j]]
		})
		for _, meshmaterialbasic := range meshmaterialbasicOrdered {
			meshmaterialbasicIdent := "__stage_0" + meshmaterialbasic.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MeshMaterialBasic{Name: %s}).Stage(stageSet.Stage)", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Color)))
		}
	}
	if stageSet.Stage != nil {
		meshphysicalmaterialOrdered := []*MeshPhysicalMaterial{}
		for meshphysicalmaterial := range stageSet.Stage.MeshPhysicalMaterials {
			meshphysicalmaterialOrdered = append(meshphysicalmaterialOrdered, meshphysicalmaterial)
		}
		sort.Slice(meshphysicalmaterialOrdered, func(i, j int) bool {
			return stageSet.Stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterialOrdered[i]] < stageSet.Stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterialOrdered[j]]
		})
		for _, meshphysicalmaterial := range meshphysicalmaterialOrdered {
			meshphysicalmaterialIdent := "__stage_0" + meshphysicalmaterial.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MeshPhysicalMaterial{Name: %s}).Stage(stageSet.Stage)", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Wireframe = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Wireframe))
			values.WriteString(fmt.Sprintf("\n\t%s.Opacity = %f", meshphysicalmaterialIdent, meshphysicalmaterial.Opacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparent = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Transparent))
			values.WriteString(fmt.Sprintf("\n\t%s.Visible = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Visible))
		}
	}
	if stageSet.Stage != nil {
		planegeometryOrdered := []*PlaneGeometry{}
		for planegeometry := range stageSet.Stage.PlaneGeometrys {
			planegeometryOrdered = append(planegeometryOrdered, planegeometry)
		}
		sort.Slice(planegeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.PlaneGeometry_stagedOrder[planegeometryOrdered[i]] < stageSet.Stage.PlaneGeometry_stagedOrder[planegeometryOrdered[j]]
		})
		for _, planegeometry := range planegeometryOrdered {
			planegeometryIdent := "__stage_0" + planegeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.PlaneGeometry{Name: %s}).Stage(stageSet.Stage)", planegeometryIdent, __gong__toRawStringLiteral(planegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", planegeometryIdent, __gong__toRawStringLiteral(planegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", planegeometryIdent, planegeometry.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", planegeometryIdent, planegeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", planegeometryIdent, planegeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", planegeometryIdent, planegeometry.HeightSegments))
		}
	}
	if stageSet.Stage != nil {
		shapeOrdered := []*Shape{}
		for shape := range stageSet.Stage.Shapes {
			shapeOrdered = append(shapeOrdered, shape)
		}
		sort.Slice(shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.Shape_stagedOrder[shapeOrdered[i]] < stageSet.Stage.Shape_stagedOrder[shapeOrdered[j]]
		})
		for _, shape := range shapeOrdered {
			shapeIdent := "__stage_0" + shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Shape{Name: %s}).Stage(stageSet.Stage)", shapeIdent, __gong__toRawStringLiteral(shape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", shapeIdent, __gong__toRawStringLiteral(shape.Name)))
			for _, elem := range shape.Points {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Points = append(%s.Points, %s)", shapeIdent, shapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		spheregeometryOrdered := []*SphereGeometry{}
		for spheregeometry := range stageSet.Stage.SphereGeometrys {
			spheregeometryOrdered = append(spheregeometryOrdered, spheregeometry)
		}
		sort.Slice(spheregeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.SphereGeometry_stagedOrder[spheregeometryOrdered[i]] < stageSet.Stage.SphereGeometry_stagedOrder[spheregeometryOrdered[j]]
		})
		for _, spheregeometry := range spheregeometryOrdered {
			spheregeometryIdent := "__stage_0" + spheregeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SphereGeometry{Name: %s}).Stage(stageSet.Stage)", spheregeometryIdent, __gong__toRawStringLiteral(spheregeometry.Name)))
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
		torusgeometryOrdered := []*TorusGeometry{}
		for torusgeometry := range stageSet.Stage.TorusGeometrys {
			torusgeometryOrdered = append(torusgeometryOrdered, torusgeometry)
		}
		sort.Slice(torusgeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.TorusGeometry_stagedOrder[torusgeometryOrdered[i]] < stageSet.Stage.TorusGeometry_stagedOrder[torusgeometryOrdered[j]]
		})
		for _, torusgeometry := range torusgeometryOrdered {
			torusgeometryIdent := "__stage_0" + torusgeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TorusGeometry{Name: %s}).Stage(stageSet.Stage)", torusgeometryIdent, __gong__toRawStringLiteral(torusgeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", torusgeometryIdent, __gong__toRawStringLiteral(torusgeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", torusgeometryIdent, torusgeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.Tube = %f", torusgeometryIdent, torusgeometry.Tube))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", torusgeometryIdent, torusgeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.TubularSegments = %d", torusgeometryIdent, torusgeometry.TubularSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Arc = %f", torusgeometryIdent, torusgeometry.Arc))
		}
	}
	if stageSet.Stage != nil {
		triangleOrdered := []*Triangle{}
		for triangle := range stageSet.Stage.Triangles {
			triangleOrdered = append(triangleOrdered, triangle)
		}
		sort.Slice(triangleOrdered, func(i, j int) bool {
			return stageSet.Stage.Triangle_stagedOrder[triangleOrdered[i]] < stageSet.Stage.Triangle_stagedOrder[triangleOrdered[j]]
		})
		for _, triangle := range triangleOrdered {
			triangleIdent := "__stage_0" + triangle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Triangle{Name: %s}).Stage(stageSet.Stage)", triangleIdent, __gong__toRawStringLiteral(triangle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", triangleIdent, __gong__toRawStringLiteral(triangle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.V1 = %d", triangleIdent, triangle.V1))
			values.WriteString(fmt.Sprintf("\n\t%s.V2 = %d", triangleIdent, triangle.V2))
			values.WriteString(fmt.Sprintf("\n\t%s.V3 = %d", triangleIdent, triangle.V3))
		}
	}
	if stageSet.Stage != nil {
		tubegeometryOrdered := []*TubeGeometry{}
		for tubegeometry := range stageSet.Stage.TubeGeometrys {
			tubegeometryOrdered = append(tubegeometryOrdered, tubegeometry)
		}
		sort.Slice(tubegeometryOrdered, func(i, j int) bool {
			return stageSet.Stage.TubeGeometry_stagedOrder[tubegeometryOrdered[i]] < stageSet.Stage.TubeGeometry_stagedOrder[tubegeometryOrdered[j]]
		})
		for _, tubegeometry := range tubegeometryOrdered {
			tubegeometryIdent := "__stage_0" + tubegeometry.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TubeGeometry{Name: %s}).Stage(stageSet.Stage)", tubegeometryIdent, __gong__toRawStringLiteral(tubegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubegeometryIdent, __gong__toRawStringLiteral(tubegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.TubularSegments = %d", tubegeometryIdent, tubegeometry.TubularSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", tubegeometryIdent, tubegeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", tubegeometryIdent, tubegeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Closed = %t", tubegeometryIdent, tubegeometry.Closed))
			if tubegeometry.Path != nil {
				targetIdent := "__stage_0" + tubegeometry.Path.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Path = %s", tubegeometryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		vector2Ordered := []*Vector2{}
		for vector2 := range stageSet.Stage.Vector2s {
			vector2Ordered = append(vector2Ordered, vector2)
		}
		sort.Slice(vector2Ordered, func(i, j int) bool {
			return stageSet.Stage.Vector2_stagedOrder[vector2Ordered[i]] < stageSet.Stage.Vector2_stagedOrder[vector2Ordered[j]]
		})
		for _, vector2 := range vector2Ordered {
			vector2Ident := "__stage_0" + vector2.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Vector2{Name: %s}).Stage(stageSet.Stage)", vector2Ident, __gong__toRawStringLiteral(vector2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vector2Ident, __gong__toRawStringLiteral(vector2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", vector2Ident, vector2.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", vector2Ident, vector2.Y))
		}
	}
	if stageSet.Stage != nil {
		vector3Ordered := []*Vector3{}
		for vector3 := range stageSet.Stage.Vector3s {
			vector3Ordered = append(vector3Ordered, vector3)
		}
		sort.Slice(vector3Ordered, func(i, j int) bool {
			return stageSet.Stage.Vector3_stagedOrder[vector3Ordered[i]] < stageSet.Stage.Vector3_stagedOrder[vector3Ordered[j]]
		})
		for _, vector3 := range vector3Ordered {
			vector3Ident := "__stage_0" + vector3.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Vector3{Name: %s}).Stage(stageSet.Stage)", vector3Ident, __gong__toRawStringLiteral(vector3.Name)))
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

	__stage_0__ "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

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

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "AmbiantLight":
					if !preserveOrder {
						inst := (&AmbiantLight{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AmbiantLight)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "BoxGeometry":
					if !preserveOrder {
						inst := (&BoxGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(BoxGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "BufferGeometry":
					if !preserveOrder {
						inst := (&BufferGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(BufferGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Camera":
					if !preserveOrder {
						inst := (&Camera{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Camera)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Canvas":
					if !preserveOrder {
						inst := (&Canvas{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Canvas)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Curve":
					if !preserveOrder {
						inst := (&Curve{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Curve)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CylinderGeometry":
					if !preserveOrder {
						inst := (&CylinderGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CylinderGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DirectionalLight":
					if !preserveOrder {
						inst := (&DirectionalLight{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DirectionalLight)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ExtrudeGeometry":
					if !preserveOrder {
						inst := (&ExtrudeGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ExtrudeGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Mesh":
					if !preserveOrder {
						inst := (&Mesh{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Mesh)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MeshMaterialBasic":
					if !preserveOrder {
						inst := (&MeshMaterialBasic{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MeshMaterialBasic)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MeshPhysicalMaterial":
					if !preserveOrder {
						inst := (&MeshPhysicalMaterial{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MeshPhysicalMaterial)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PlaneGeometry":
					if !preserveOrder {
						inst := (&PlaneGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PlaneGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Shape":
					if !preserveOrder {
						inst := (&Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SphereGeometry":
					if !preserveOrder {
						inst := (&SphereGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SphereGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TorusGeometry":
					if !preserveOrder {
						inst := (&TorusGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TorusGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Triangle":
					if !preserveOrder {
						inst := (&Triangle{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Triangle)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TubeGeometry":
					if !preserveOrder {
						inst := (&TubeGeometry{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TubeGeometry)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Vector2":
					if !preserveOrder {
						inst := (&Vector2{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Vector2)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Vector3":
					if !preserveOrder {
						inst := (&Vector3{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Vector3)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Vector3); ok {
										inst.Vertices = append(inst.Vertices, typedTarget)
									}
								}
							}
						}
					case "Faces":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Triangle); ok {
										inst.Faces = append(inst.Faces, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DirectionalLight); ok {
										inst.DirectionalLights = append(inst.DirectionalLights, typedTarget)
									}
								}
							}
						}
					case "AmbiantLight":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AmbiantLight); ok {
									inst.AmbiantLight = typedTarget
								}
							}
						}
					case "Meshs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Mesh); ok {
										inst.Meshs = append(inst.Meshs, typedTarget)
									}
								}
							}
						}
					case "Camera":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Camera); ok {
									inst.Camera = typedTarget
								}
							}
						}
					case "IsWithLastRenderingUpdate":
						inst.IsWithLastRenderingUpdate = GongExtractBool(rhs)
					case "LastRendering":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.LastRendering, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Frame64BitsEncoded":
						inst.Frame64BitsEncoded = GongExtractString(rhs)
					}
				case *Curve:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Vector3); ok {
										inst.Points = append(inst.Points, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Shape); ok {
									inst.Shape = typedTarget
								}
							}
						}
					case "ExtrudePath":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Curve); ok {
									inst.ExtrudePath = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MeshMaterialBasic); ok {
									inst.MeshMaterialBasic = typedTarget
								}
							}
						}
					case "MeshPhysicalMaterial":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MeshPhysicalMaterial); ok {
									inst.MeshPhysicalMaterial = typedTarget
								}
							}
						}
					case "CylinderGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CylinderGeometry); ok {
									inst.CylinderGeometry = typedTarget
								}
							}
						}
					case "BoxGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*BoxGeometry); ok {
									inst.BoxGeometry = typedTarget
								}
							}
						}
					case "SphereGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SphereGeometry); ok {
									inst.SphereGeometry = typedTarget
								}
							}
						}
					case "TorusGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TorusGeometry); ok {
									inst.TorusGeometry = typedTarget
								}
							}
						}
					case "PlaneGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PlaneGeometry); ok {
									inst.PlaneGeometry = typedTarget
								}
							}
						}
					case "TubeGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TubeGeometry); ok {
									inst.TubeGeometry = typedTarget
								}
							}
						}
					case "ExtrudeGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ExtrudeGeometry); ok {
									inst.ExtrudeGeometry = typedTarget
								}
							}
						}
					case "BufferGeometry":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*BufferGeometry); ok {
									inst.BufferGeometry = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Vector2); ok {
										inst.Points = append(inst.Points, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Curve); ok {
									inst.Path = typedTarget
								}
							}
						}
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

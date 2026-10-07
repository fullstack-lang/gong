// generated code - do not edit
package models

import (
	"embed"
	"go/ast"
	"go/token"
	"log"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var _ = time.Hour

// swagger:ignore
type GONG__ExpressionType = gong_runtime.GONG__ExpressionType

const (
	GONG__STRUCT_INSTANCE      = gong_runtime.GONG__STRUCT_INSTANCE
	GONG__FIELD_OR_CONST_VALUE = gong_runtime.GONG__FIELD_OR_CONST_VALUE
	GONG__FIELD_VALUE          = gong_runtime.GONG__FIELD_VALUE
	GONG__ENUM_CAST_INT        = gong_runtime.GONG__ENUM_CAST_INT
	GONG__ENUM_CAST_STRING     = gong_runtime.GONG__ENUM_CAST_STRING
	GONG__IDENTIFIER_CONST     = gong_runtime.GONG__IDENTIFIER_CONST
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// GongModelUnmarshaller abstracts the logic for setting fields on a staged instance
type GongModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

type ModelUnmarshaller = GongModelUnmarshaller

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func (stage *Stage) ParseAstFile(pathToFile string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstFile(pathToFile)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func (stage *Stage) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	inFile, fset, err := gong_runtime.ParseAstEmbeddedFile(directory, pathToFile, stage.GetName())
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string
func (stage *Stage) ParseAstString(blob string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstString(blob)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances using the Unmarshaller registry
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]GongstructIF)

	for _, instance := range stage.GetInstances() {
		identifierMap[instance.GongGetIdentifier(stage)] = instance
	}

	return gong_runtime.WalkAstFile(inFile,
		func(identName string, typeName string, instanceName string) {
			if typeName != "" {
				if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
					instance, err := unmarshaller.Initialize(stage, identName, instanceName, preserveOrder)
					if err == nil {
						identifierMap[identName] = instance
					}
				}
			}
		},
		func(identName string, fieldName string, valueExpr ast.Expr) {
			if instance, exists := identifierMap[identName]; exists {
				typeName := instance.GongGetGongstructName()
				if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
					unmarshaller.UnmarshallField(stage, instance, fieldName, valueExpr, identifierMap)
				}
			}
		},
		func(identName string) {
			if instance, ok := identifierMap[identName]; ok {
				instance.UnstageVoid(stage)
			}
		},
		func() {
			if stage.IsInDeltaMode() && stage.GetNavigationMode() != GongNavigationModeNavigating {
				stage.Commit()
			} else {
				stage.ComputeInstancesNb()
				stage.ComputeReferenceAndOrders()
				if stage.OnInitCommitCallback != nil {
					stage.OnInitCommitCallback.BeforeCommit(stage)
				}
				if stage.OnInitCommitFromBackCallback != nil {
					stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
				}
				// 1. Run all Before Commit hooks
				stage.RunBeforeCommitHooks()

				// 2. Run all After Commit hooks
				stage.RunAfterCommitHooks()
			}
		},
	)
}

// --- Generic Helpers for Unmarshallers (delegating to gong_runtime) ---

var GongExtractString = gong_runtime.ExtractString
var GongExtractInt = gong_runtime.ExtractInt
var GongExtractFloat = gong_runtime.ExtractFloat
var GongExtractBool = gong_runtime.ExtractBool
var GongExtractExpr = gong_runtime.ExtractExpr
var GongExtractDate = gong_runtime.ExtractDate

func __gong__extractMiddleUint(input string) (uint, error) {
	return gong_runtime.ExtractMiddleUint(input)
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T GongstructPtr](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {
	return gong_runtime.UnmarshallSliceOfPointers(slice, valueExpr, identifierMap)
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T GongstructPtr](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {
	gong_runtime.UnmarshallPointer(ptr, valueExpr, identifierMap)
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {
	gong_runtime.UnmarshallEnum(ptr, valueExpr)
}

// GongInitialize initializes a staged instance, sets its name, and stages it
func GongInitialize[P interface {
	GongstructIF
	StagePreserveOrder(stage *Stage, order uint)
}](instance P, stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if newOrder, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

// insertion point per named struct
type AmbiantLightUnmarshaller struct{}

func (u *AmbiantLightUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AmbiantLight), stage, identifier, instanceName, preserveOrder)
}

func (u *AmbiantLightUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AmbiantLight)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Intensity":
		instance.Intensity = GongExtractFloat(valueExpr)
	}
	return nil
}

type BoxGeometryUnmarshaller struct{}

func (u *BoxGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BoxGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *BoxGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BoxGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "Depth":
		instance.Depth = GongExtractFloat(valueExpr)
	case "WidthSegments":
		instance.WidthSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	case "DepthSegments":
		instance.DepthSegments = GongExtractInt(valueExpr)
	}
	return nil
}

type BufferGeometryUnmarshaller struct{}

func (u *BufferGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BufferGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *BufferGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BufferGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Vertices":
		GongUnmarshallSliceOfPointers(&instance.Vertices, valueExpr, identifierMap)
	case "Faces":
		GongUnmarshallSliceOfPointers(&instance.Faces, valueExpr, identifierMap)
	}
	return nil
}

type CameraUnmarshaller struct{}

func (u *CameraUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Camera), stage, identifier, instanceName, preserveOrder)
}

func (u *CameraUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Camera)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	case "TargetX":
		instance.TargetX = GongExtractFloat(valueExpr)
	case "TargetY":
		instance.TargetY = GongExtractFloat(valueExpr)
	case "TargetZ":
		instance.TargetZ = GongExtractFloat(valueExpr)
	case "Fov":
		instance.Fov = GongExtractFloat(valueExpr)
	}
	return nil
}

type CanvasUnmarshaller struct{}

func (u *CanvasUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Canvas), stage, identifier, instanceName, preserveOrder)
}

func (u *CanvasUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Canvas)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DirectionalLights":
		GongUnmarshallSliceOfPointers(&instance.DirectionalLights, valueExpr, identifierMap)
	case "AmbiantLight":
		GongUnmarshallPointer(&instance.AmbiantLight, valueExpr, identifierMap)
	case "Meshs":
		GongUnmarshallSliceOfPointers(&instance.Meshs, valueExpr, identifierMap)
	case "Camera":
		GongUnmarshallPointer(&instance.Camera, valueExpr, identifierMap)
	case "IsWithLastRenderingUpdate":
		instance.IsWithLastRenderingUpdate = GongExtractBool(valueExpr)
	case "LastRendering":
		instance.LastRendering = GongExtractDate(valueExpr)
	case "Frame64BitsEncoded":
		instance.Frame64BitsEncoded = GongExtractString(valueExpr)
	}
	return nil
}

type CurveUnmarshaller struct{}

func (u *CurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Curve), stage, identifier, instanceName, preserveOrder)
}

func (u *CurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Curve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		GongUnmarshallSliceOfPointers(&instance.Points, valueExpr, identifierMap)
	}
	return nil
}

type CylinderGeometryUnmarshaller struct{}

func (u *CylinderGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CylinderGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *CylinderGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CylinderGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RadiusTop":
		instance.RadiusTop = GongExtractFloat(valueExpr)
	case "RadiusBottom":
		instance.RadiusBottom = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RadialSegments":
		instance.RadialSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	case "OpenEnded":
		instance.OpenEnded = GongExtractBool(valueExpr)
	case "ThetaStart":
		instance.ThetaStart = GongExtractFloat(valueExpr)
	case "ThetaLength":
		instance.ThetaLength = GongExtractFloat(valueExpr)
	}
	return nil
}

type DirectionalLightUnmarshaller struct{}

func (u *DirectionalLightUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectionalLight), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionalLightUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectionalLight)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	case "Intensity":
		instance.Intensity = GongExtractFloat(valueExpr)
	case "IsWithCastShadow":
		instance.IsWithCastShadow = GongExtractBool(valueExpr)
	}
	return nil
}

type ExtrudeGeometryUnmarshaller struct{}

func (u *ExtrudeGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ExtrudeGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *ExtrudeGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExtrudeGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Shape":
		GongUnmarshallPointer(&instance.Shape, valueExpr, identifierMap)
	case "ExtrudePath":
		GongUnmarshallPointer(&instance.ExtrudePath, valueExpr, identifierMap)
	case "Steps":
		instance.Steps = GongExtractInt(valueExpr)
	}
	return nil
}

type MeshUnmarshaller struct{}

func (u *MeshUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Mesh), stage, identifier, instanceName, preserveOrder)
}

func (u *MeshUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Mesh)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	case "MeshMaterialBasic":
		GongUnmarshallPointer(&instance.MeshMaterialBasic, valueExpr, identifierMap)
	case "MeshPhysicalMaterial":
		GongUnmarshallPointer(&instance.MeshPhysicalMaterial, valueExpr, identifierMap)
	case "CylinderGeometry":
		GongUnmarshallPointer(&instance.CylinderGeometry, valueExpr, identifierMap)
	case "BoxGeometry":
		GongUnmarshallPointer(&instance.BoxGeometry, valueExpr, identifierMap)
	case "SphereGeometry":
		GongUnmarshallPointer(&instance.SphereGeometry, valueExpr, identifierMap)
	case "TorusGeometry":
		GongUnmarshallPointer(&instance.TorusGeometry, valueExpr, identifierMap)
	case "PlaneGeometry":
		GongUnmarshallPointer(&instance.PlaneGeometry, valueExpr, identifierMap)
	case "TubeGeometry":
		GongUnmarshallPointer(&instance.TubeGeometry, valueExpr, identifierMap)
	case "ExtrudeGeometry":
		GongUnmarshallPointer(&instance.ExtrudeGeometry, valueExpr, identifierMap)
	case "BufferGeometry":
		GongUnmarshallPointer(&instance.BufferGeometry, valueExpr, identifierMap)
	}
	return nil
}

type MeshMaterialBasicUnmarshaller struct{}

func (u *MeshMaterialBasicUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeshMaterialBasic), stage, identifier, instanceName, preserveOrder)
}

func (u *MeshMaterialBasicUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeshMaterialBasic)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type MeshPhysicalMaterialUnmarshaller struct{}

func (u *MeshPhysicalMaterialUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeshPhysicalMaterial), stage, identifier, instanceName, preserveOrder)
}

func (u *MeshPhysicalMaterialUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeshPhysicalMaterial)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Wireframe":
		instance.Wireframe = GongExtractBool(valueExpr)
	case "Opacity":
		instance.Opacity = GongExtractFloat(valueExpr)
	case "Transparent":
		instance.Transparent = GongExtractBool(valueExpr)
	case "Visible":
		instance.Visible = GongExtractBool(valueExpr)
	}
	return nil
}

type PlaneGeometryUnmarshaller struct{}

func (u *PlaneGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PlaneGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *PlaneGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlaneGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "WidthSegments":
		instance.WidthSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	}
	return nil
}

type ShapeUnmarshaller struct{}

func (u *ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		GongUnmarshallSliceOfPointers(&instance.Points, valueExpr, identifierMap)
	}
	return nil
}

type SphereGeometryUnmarshaller struct{}

func (u *SphereGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SphereGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *SphereGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SphereGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "WidthSegments":
		instance.WidthSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	case "PhiStart":
		instance.PhiStart = GongExtractFloat(valueExpr)
	case "PhiLength":
		instance.PhiLength = GongExtractFloat(valueExpr)
	case "ThetaStart":
		instance.ThetaStart = GongExtractFloat(valueExpr)
	case "ThetaLength":
		instance.ThetaLength = GongExtractFloat(valueExpr)
	}
	return nil
}

type TorusGeometryUnmarshaller struct{}

func (u *TorusGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TorusGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *TorusGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TorusGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "Tube":
		instance.Tube = GongExtractFloat(valueExpr)
	case "RadialSegments":
		instance.RadialSegments = GongExtractInt(valueExpr)
	case "TubularSegments":
		instance.TubularSegments = GongExtractInt(valueExpr)
	case "Arc":
		instance.Arc = GongExtractFloat(valueExpr)
	}
	return nil
}

type TriangleUnmarshaller struct{}

func (u *TriangleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Triangle), stage, identifier, instanceName, preserveOrder)
}

func (u *TriangleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Triangle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "V1":
		instance.V1 = GongExtractInt(valueExpr)
	case "V2":
		instance.V2 = GongExtractInt(valueExpr)
	case "V3":
		instance.V3 = GongExtractInt(valueExpr)
	}
	return nil
}

type TubeGeometryUnmarshaller struct{}

func (u *TubeGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TubeGeometry), stage, identifier, instanceName, preserveOrder)
}

func (u *TubeGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TubeGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Path":
		GongUnmarshallPointer(&instance.Path, valueExpr, identifierMap)
	case "TubularSegments":
		instance.TubularSegments = GongExtractInt(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "RadialSegments":
		instance.RadialSegments = GongExtractInt(valueExpr)
	case "Closed":
		instance.Closed = GongExtractBool(valueExpr)
	}
	return nil
}

type Vector2Unmarshaller struct{}

func (u *Vector2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Vector2), stage, identifier, instanceName, preserveOrder)
}

func (u *Vector2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Vector2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type Vector3Unmarshaller struct{}

func (u *Vector3Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Vector3), stage, identifier, instanceName, preserveOrder)
}

func (u *Vector3Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Vector3)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	}
	return nil
}

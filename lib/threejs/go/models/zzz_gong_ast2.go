// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"slices"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var (
	_ = errors.New
	_ = parser.ParseFile
	_ = filepath.Base
	_ = slices.Index[[]int, int]
	_ = strings.Trim
	_ = time.Hour
)

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
	var inFile *ast.File
	var fset *token.FileSet
	var err error
	if strings.HasPrefix(strings.TrimSpace(blob), "package ") || strings.Contains(blob, "\npackage ") {
		fset = token.NewFileSet()
		inFile, err = parser.ParseFile(fset, "", blob, parser.ParseComments)
		if err != nil {
			return errors.New("Unable to parse " + err.Error())
		}
	} else {
		inFile, fset, err = gong_runtime.ParseAstString(blob)
		if err != nil {
			return err
		}
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances across stage and its sub-stages
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	identifierMap := make(map[string]any)

	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
		}
	}

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
					} else if pkgAlias == "" || pkgAlias == "models" || len(inFile.Imports) == 0 {
						pkgAlias = "models"
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "AmbiantLight":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AmbiantLight), stage, ident.Name, instanceName, preserveOrder)
				case "BoxGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BoxGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "BufferGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BufferGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "Camera":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Camera), stage, ident.Name, instanceName, preserveOrder)
				case "Canvas":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Canvas), stage, ident.Name, instanceName, preserveOrder)
				case "Curve":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Curve), stage, ident.Name, instanceName, preserveOrder)
				case "CylinderGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CylinderGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "DirectionalLight":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectionalLight), stage, ident.Name, instanceName, preserveOrder)
				case "ExtrudeGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ExtrudeGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "Mesh":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Mesh), stage, ident.Name, instanceName, preserveOrder)
				case "MeshMaterialBasic":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeshMaterialBasic), stage, ident.Name, instanceName, preserveOrder)
				case "MeshPhysicalMaterial":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeshPhysicalMaterial), stage, ident.Name, instanceName, preserveOrder)
				case "PlaneGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PlaneGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Shape), stage, ident.Name, instanceName, preserveOrder)
				case "SphereGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SphereGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "TorusGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TorusGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "Triangle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Triangle), stage, ident.Name, instanceName, preserveOrder)
				case "TubeGeometry":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TubeGeometry), stage, ident.Name, instanceName, preserveOrder)
				case "Vector2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Vector2), stage, ident.Name, instanceName, preserveOrder)
				case "Vector3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Vector3), stage, ident.Name, instanceName, preserveOrder)
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
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								if inst, ok := instance.(GongstructIF); ok {
									inst.UnstageVoid(stage)
								}
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
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
								stage.RunBeforeCommitHooks()
								stage.RunAfterCommitHooks()
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
	if call, ok := rhs.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				if end <= len(*slice) {
					*slice = slices.Delete(*slice, start, end)
				}
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index <= len(*slice) {
					if rIdent, ok := call.Args[2].(*ast.Ident); ok {
						if target, ok := identifierMap[rIdent.Name]; ok {
							if typedTarget, ok := target.(*T); ok {
								*slice = slices.Insert(*slice, index, typedTarget)
							}
						}
					}
				}
			}
		} else if funcName == "append" && len(call.Args) >= 2 {
			for i := 1; i < len(call.Args); i++ {
				if rIdent, ok := call.Args[i].(*ast.Ident); ok {
					if target, ok := identifierMap[rIdent.Name]; ok {
						if typedTarget, ok := target.(*T); ok {
							*slice = append(*slice, typedTarget)
						}
					}
				}
			}
		}
	}
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

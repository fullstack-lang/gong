// generated code - do not edit
package aixmdatatypes

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
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
					} else if pkgAlias == "" || pkgAlias == "aixmdatatypes" || len(inFile.Imports) == 0 {
						pkgAlias = "aixmdatatypes"
					}

					switch pkgAlias {
			case "aixmdatatypes":
				switch typeName {
				case "CodeColourType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeColourType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDayType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeGroundLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeGroundLightingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLightIntensityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLightSourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLightSourceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNotePurposeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstacleAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAssessmentSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstacleAssessmentSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusConstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusConstructionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusOperationsType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeEventCombinationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeEventType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalDatumType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalStructureMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureMaterialType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalStructureMaterialType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalStructureType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeYesNoType), stage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DateMonthDayType), stage, ident.Name, instanceName, preserveOrder)
				case "TextDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "TextNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextNameType), stage, ident.Name, instanceName, preserveOrder)
				case "TextNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "TextPropertyNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextPropertyNameType), stage, ident.Name, instanceName, preserveOrder)
				case "TimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimeType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceSignedType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceVerticalType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDurationType), stage, ident.Name, instanceName, preserveOrder)
				case "ValLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValLightIntensityType), stage, ident.Name, instanceName, preserveOrder)
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
				case *CodeColourType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeGroundLightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLightIntensityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLightSourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNotePurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstacleAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstacleAssessmentSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusConstructionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusOperationsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeEventCombinationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeEventType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalDatumType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalStructureMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalStructureMaterialType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalStructureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeYesNoType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DateMonthDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextNoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextPropertyNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TimeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDistanceSignedType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDistanceVerticalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDurationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValLightIntensityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
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
type CodeColourTypeUnmarshaller struct{}

func (u *CodeColourTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeColourType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeColourTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeColourType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDayTypeUnmarshaller struct{}

func (u *CodeDayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDayType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeGroundLightingTypeUnmarshaller struct{}

func (u *CodeGroundLightingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeGroundLightingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeGroundLightingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeGroundLightingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLightIntensityTypeUnmarshaller struct{}

func (u *CodeLightIntensityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLightIntensityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLightIntensityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLightIntensityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLightSourceTypeUnmarshaller struct{}

func (u *CodeLightSourceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLightSourceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLightSourceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLightSourceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNotePurposeTypeUnmarshaller struct{}

func (u *CodeNotePurposeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNotePurposeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNotePurposeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNotePurposeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstacleAreaTypeUnmarshaller struct{}

func (u *CodeObstacleAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstacleAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstacleAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstacleAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstacleAssessmentSurfaceTypeUnmarshaller struct{}

func (u *CodeObstacleAssessmentSurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstacleAssessmentSurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstacleAssessmentSurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstacleAssessmentSurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusConstructionTypeUnmarshaller struct{}

func (u *CodeStatusConstructionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusConstructionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusConstructionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusConstructionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusOperationsTypeUnmarshaller struct{}

func (u *CodeStatusOperationsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusOperationsType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusOperationsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusOperationsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeEventCombinationTypeUnmarshaller struct{}

func (u *CodeTimeEventCombinationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeEventCombinationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeEventCombinationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeEventCombinationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeEventTypeUnmarshaller struct{}

func (u *CodeTimeEventTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeEventType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeEventTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeEventType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeReferenceTypeUnmarshaller struct{}

func (u *CodeTimeReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalDatumTypeUnmarshaller struct{}

func (u *CodeVerticalDatumTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalDatumType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalDatumTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalDatumType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalStructureMarkingTypeUnmarshaller struct{}

func (u *CodeVerticalStructureMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalStructureMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalStructureMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalStructureMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalStructureMaterialTypeUnmarshaller struct{}

func (u *CodeVerticalStructureMaterialTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalStructureMaterialType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalStructureMaterialTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalStructureMaterialType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalStructureTypeUnmarshaller struct{}

func (u *CodeVerticalStructureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalStructureType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalStructureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalStructureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeYesNoTypeUnmarshaller struct{}

func (u *CodeYesNoTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeYesNoType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeYesNoTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeYesNoType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DateMonthDayTypeUnmarshaller struct{}

func (u *DateMonthDayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DateMonthDayType), stage, identifier, instanceName, preserveOrder)
}

func (u *DateMonthDayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DateMonthDayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextDesignatorTypeUnmarshaller struct{}

func (u *TextDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextNameTypeUnmarshaller struct{}

func (u *TextNameTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextNameType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextNameTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextNameType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextNoteTypeUnmarshaller struct{}

func (u *TextNoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextNoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextNoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextNoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextPropertyNameTypeUnmarshaller struct{}

func (u *TextPropertyNameTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextPropertyNameType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextPropertyNameTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextPropertyNameType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TimeTypeUnmarshaller struct{}

func (u *TimeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimeType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceSignedTypeUnmarshaller struct{}

func (u *ValDistanceSignedTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceSignedType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceSignedTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceSignedType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceTypeUnmarshaller struct{}

func (u *ValDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceVerticalTypeUnmarshaller struct{}

func (u *ValDistanceVerticalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceVerticalType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceVerticalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceVerticalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDurationTypeUnmarshaller struct{}

func (u *ValDurationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDurationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDurationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDurationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValLightIntensityTypeUnmarshaller struct{}

func (u *ValLightIntensityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValLightIntensityType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValLightIntensityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValLightIntensityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

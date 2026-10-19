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
		case "github.com/fullstack-lang/gong/lib/split/go/models":
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
				case "AsSplit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AsSplit), stage, ident.Name, instanceName, preserveOrder)
				case "AsSplitArea":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AsSplitArea), stage, ident.Name, instanceName, preserveOrder)
				case "Button":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Button), stage, ident.Name, instanceName, preserveOrder)
				case "Cursor":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cursor), stage, ident.Name, instanceName, preserveOrder)
				case "FavIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FavIcon), stage, ident.Name, instanceName, preserveOrder)
				case "Form":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Form), stage, ident.Name, instanceName, preserveOrder)
				case "Load":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Load), stage, ident.Name, instanceName, preserveOrder)
				case "LogoOnTheLeft":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LogoOnTheLeft), stage, ident.Name, instanceName, preserveOrder)
				case "LogoOnTheRight":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LogoOnTheRight), stage, ident.Name, instanceName, preserveOrder)
				case "Markdown":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Markdown), stage, ident.Name, instanceName, preserveOrder)
				case "Slider":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slider), stage, ident.Name, instanceName, preserveOrder)
				case "Split":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Split), stage, ident.Name, instanceName, preserveOrder)
				case "Svg":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Svg), stage, ident.Name, instanceName, preserveOrder)
				case "Table":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Table), stage, ident.Name, instanceName, preserveOrder)
				case "Threejs":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Threejs), stage, ident.Name, instanceName, preserveOrder)
				case "Title":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Title), stage, ident.Name, instanceName, preserveOrder)
				case "Tone":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tone), stage, ident.Name, instanceName, preserveOrder)
				case "Tree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tree), stage, ident.Name, instanceName, preserveOrder)
				case "View":
					identifierMap[ident.Name] = __gong__stageSetInit(new(View), stage, ident.Name, instanceName, preserveOrder)
				case "Xlsx":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Xlsx), stage, ident.Name, instanceName, preserveOrder)
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
				case *AsSplit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
					case "AsSplitAreas":
						__gong__assignSliceOfPointers(&inst.AsSplitAreas, rhs, identifierMap)
					case "IsSizeInPixel":
						inst.IsSizeInPixel = GongExtractBool(rhs)
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
				case *AsSplitArea:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowNameInHeader":
						inst.ShowNameInHeader = GongExtractBool(rhs)
					case "Size":
						inst.Size = GongExtractFloat(rhs)
					case "IsAny":
						inst.IsAny = GongExtractBool(rhs)
					case "AsSplit":
						__gong__assignPointer(&inst.AsSplit, rhs, identifierMap)
					case "Button":
						__gong__assignPointer(&inst.Button, rhs, identifierMap)
					case "Cursor":
						__gong__assignPointer(&inst.Cursor, rhs, identifierMap)
					case "Form":
						__gong__assignPointer(&inst.Form, rhs, identifierMap)
					case "Load":
						__gong__assignPointer(&inst.Load, rhs, identifierMap)
					case "Markdown":
						__gong__assignPointer(&inst.Markdown, rhs, identifierMap)
					case "Slider":
						__gong__assignPointer(&inst.Slider, rhs, identifierMap)
					case "Split":
						__gong__assignPointer(&inst.Split, rhs, identifierMap)
					case "Svg":
						__gong__assignPointer(&inst.Svg, rhs, identifierMap)
					case "Table":
						__gong__assignPointer(&inst.Table, rhs, identifierMap)
					case "Tone":
						__gong__assignPointer(&inst.Tone, rhs, identifierMap)
					case "Tree":
						__gong__assignPointer(&inst.Tree, rhs, identifierMap)
					case "Threejs":
						__gong__assignPointer(&inst.Threejs, rhs, identifierMap)
					case "Xlsx":
						__gong__assignPointer(&inst.Xlsx, rhs, identifierMap)
					case "HasDiv":
						inst.HasDiv = GongExtractBool(rhs)
					case "DivStyle":
						inst.DivStyle = GongExtractString(rhs)
					}
				case *Button:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Cursor:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					case "Style":
						inst.Style = GongExtractString(rhs)
					}
				case *FavIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Form:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Load:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *LogoOnTheLeft:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *LogoOnTheRight:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Markdown:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Slider:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Split:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Svg:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					case "Style":
						inst.Style = GongExtractString(rhs)
					}
				case *Table:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Threejs:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Title:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Tone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Tree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *View:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowViewName":
						inst.ShowViewName = GongExtractBool(rhs)
					case "RootAsSplitAreas":
						__gong__assignSliceOfPointers(&inst.RootAsSplitAreas, rhs, identifierMap)
					case "IsSelectedView":
						inst.IsSelectedView = GongExtractBool(rhs)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
					case "IsSecondaryView":
						inst.IsSecondaryView = GongExtractBool(rhs)
					case "IsSizeInPixel":
						inst.IsSizeInPixel = GongExtractBool(rhs)
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
				case *Xlsx:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
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
type AsSplitUnmarshaller struct{}

func (u *AsSplitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AsSplit), stage, identifier, instanceName, preserveOrder)
}

func (u *AsSplitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AsSplit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "AsSplitAreas":
		GongUnmarshallSliceOfPointers(&instance.AsSplitAreas, valueExpr, identifierMap)
	case "IsSizeInPixel":
		instance.IsSizeInPixel = GongExtractBool(valueExpr)
	case "IsWithCustomGutterSize":
		instance.IsWithCustomGutterSize = GongExtractBool(valueExpr)
	case "GutterSize":
		instance.GutterSize = GongExtractFloat(valueExpr)
	}
	return nil
}

type AsSplitAreaUnmarshaller struct{}

func (u *AsSplitAreaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AsSplitArea), stage, identifier, instanceName, preserveOrder)
}

func (u *AsSplitAreaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AsSplitArea)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowNameInHeader":
		instance.ShowNameInHeader = GongExtractBool(valueExpr)
	case "Size":
		instance.Size = GongExtractFloat(valueExpr)
	case "IsAny":
		instance.IsAny = GongExtractBool(valueExpr)
	case "AsSplit":
		GongUnmarshallPointer(&instance.AsSplit, valueExpr, identifierMap)
	case "Button":
		GongUnmarshallPointer(&instance.Button, valueExpr, identifierMap)
	case "Cursor":
		GongUnmarshallPointer(&instance.Cursor, valueExpr, identifierMap)
	case "Form":
		GongUnmarshallPointer(&instance.Form, valueExpr, identifierMap)
	case "Load":
		GongUnmarshallPointer(&instance.Load, valueExpr, identifierMap)
	case "Markdown":
		GongUnmarshallPointer(&instance.Markdown, valueExpr, identifierMap)
	case "Slider":
		GongUnmarshallPointer(&instance.Slider, valueExpr, identifierMap)
	case "Split":
		GongUnmarshallPointer(&instance.Split, valueExpr, identifierMap)
	case "Svg":
		GongUnmarshallPointer(&instance.Svg, valueExpr, identifierMap)
	case "Table":
		GongUnmarshallPointer(&instance.Table, valueExpr, identifierMap)
	case "Tone":
		GongUnmarshallPointer(&instance.Tone, valueExpr, identifierMap)
	case "Tree":
		GongUnmarshallPointer(&instance.Tree, valueExpr, identifierMap)
	case "Threejs":
		GongUnmarshallPointer(&instance.Threejs, valueExpr, identifierMap)
	case "Xlsx":
		GongUnmarshallPointer(&instance.Xlsx, valueExpr, identifierMap)
	case "HasDiv":
		instance.HasDiv = GongExtractBool(valueExpr)
	case "DivStyle":
		instance.DivStyle = GongExtractString(valueExpr)
	}
	return nil
}

type ButtonUnmarshaller struct{}

func (u *ButtonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Button), stage, identifier, instanceName, preserveOrder)
}

func (u *ButtonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Button)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type CursorUnmarshaller struct{}

func (u *CursorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Cursor), stage, identifier, instanceName, preserveOrder)
}

func (u *CursorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Cursor)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	case "Style":
		instance.Style = GongExtractString(valueExpr)
	}
	return nil
}

type FavIconUnmarshaller struct{}

func (u *FavIconUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FavIcon), stage, identifier, instanceName, preserveOrder)
}

func (u *FavIconUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FavIcon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SVG":
		instance.SVG = GongExtractString(valueExpr)
	}
	return nil
}

type FormUnmarshaller struct{}

func (u *FormUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Form), stage, identifier, instanceName, preserveOrder)
}

func (u *FormUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Form)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type LoadUnmarshaller struct{}

func (u *LoadUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Load), stage, identifier, instanceName, preserveOrder)
}

func (u *LoadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Load)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type LogoOnTheLeftUnmarshaller struct{}

func (u *LogoOnTheLeftUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LogoOnTheLeft), stage, identifier, instanceName, preserveOrder)
}

func (u *LogoOnTheLeftUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LogoOnTheLeft)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractInt(valueExpr)
	case "Height":
		instance.Height = GongExtractInt(valueExpr)
	case "SVG":
		instance.SVG = GongExtractString(valueExpr)
	}
	return nil
}

type LogoOnTheRightUnmarshaller struct{}

func (u *LogoOnTheRightUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LogoOnTheRight), stage, identifier, instanceName, preserveOrder)
}

func (u *LogoOnTheRightUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LogoOnTheRight)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractInt(valueExpr)
	case "Height":
		instance.Height = GongExtractInt(valueExpr)
	case "SVG":
		instance.SVG = GongExtractString(valueExpr)
	}
	return nil
}

type MarkdownUnmarshaller struct{}

func (u *MarkdownUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Markdown), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkdownUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Markdown)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type SliderUnmarshaller struct{}

func (u *SliderUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slider), stage, identifier, instanceName, preserveOrder)
}

func (u *SliderUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slider)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type SplitUnmarshaller struct{}

func (u *SplitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Split), stage, identifier, instanceName, preserveOrder)
}

func (u *SplitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Split)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type SvgUnmarshaller struct{}

func (u *SvgUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Svg), stage, identifier, instanceName, preserveOrder)
}

func (u *SvgUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Svg)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	case "Style":
		instance.Style = GongExtractString(valueExpr)
	}
	return nil
}

type TableUnmarshaller struct{}

func (u *TableUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Table), stage, identifier, instanceName, preserveOrder)
}

func (u *TableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Table)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type ThreejsUnmarshaller struct{}

func (u *ThreejsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Threejs), stage, identifier, instanceName, preserveOrder)
}

func (u *ThreejsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Threejs)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type TitleUnmarshaller struct{}

func (u *TitleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Title), stage, identifier, instanceName, preserveOrder)
}

func (u *TitleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Title)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ToneUnmarshaller struct{}

func (u *ToneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tone), stage, identifier, instanceName, preserveOrder)
}

func (u *ToneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tone)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type TreeUnmarshaller struct{}

func (u *TreeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tree), stage, identifier, instanceName, preserveOrder)
}

func (u *TreeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tree)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

type ViewUnmarshaller struct{}

func (u *ViewUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(View), stage, identifier, instanceName, preserveOrder)
}

func (u *ViewUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*View)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowViewName":
		instance.ShowViewName = GongExtractBool(valueExpr)
	case "RootAsSplitAreas":
		GongUnmarshallSliceOfPointers(&instance.RootAsSplitAreas, valueExpr, identifierMap)
	case "IsSelectedView":
		instance.IsSelectedView = GongExtractBool(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "IsSecondaryView":
		instance.IsSecondaryView = GongExtractBool(valueExpr)
	case "IsSizeInPixel":
		instance.IsSizeInPixel = GongExtractBool(valueExpr)
	case "IsWithCustomGutterSize":
		instance.IsWithCustomGutterSize = GongExtractBool(valueExpr)
	case "GutterSize":
		instance.GutterSize = GongExtractFloat(valueExpr)
	}
	return nil
}

type XlsxUnmarshaller struct{}

func (u *XlsxUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Xlsx), stage, identifier, instanceName, preserveOrder)
}

func (u *XlsxUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Xlsx)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.StackName = GongExtractString(valueExpr)
	}
	return nil
}

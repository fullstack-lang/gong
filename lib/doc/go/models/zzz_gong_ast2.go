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
		case "github.com/fullstack-lang/gong/lib/doc/go/models":
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
				case "AttributeShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AttributeShape), stage, ident.Name, instanceName, preserveOrder)
				case "Classdiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Classdiagram), stage, ident.Name, instanceName, preserveOrder)
				case "DiagramPackage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramPackage), stage, ident.Name, instanceName, preserveOrder)
				case "GongEnumShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongEnumShape), stage, ident.Name, instanceName, preserveOrder)
				case "GongEnumValueShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongEnumValueShape), stage, ident.Name, instanceName, preserveOrder)
				case "GongNoteLinkShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongNoteLinkShape), stage, ident.Name, instanceName, preserveOrder)
				case "GongNoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongNoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "GongStructShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongStructShape), stage, ident.Name, instanceName, preserveOrder)
				case "LinkShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinkShape), stage, ident.Name, instanceName, preserveOrder)
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
				case *AttributeShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IdentifierMeta":
						inst.IdentifierMeta = GongExtractExpr(rhs)
					case "FieldTypeAsString":
						inst.FieldTypeAsString = GongExtractString(rhs)
					case "Structname":
						inst.Structname = GongExtractString(rhs)
					case "Fieldtypename":
						inst.Fieldtypename = GongExtractString(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *Classdiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "IsIncludedInStaticWebSite":
						inst.IsIncludedInStaticWebSite = GongExtractBool(rhs)
					case "GongStructShapes":
						__gong__assignSliceOfPointers(&inst.GongStructShapes, rhs, identifierMap)
					case "GongEnumShapes":
						__gong__assignSliceOfPointers(&inst.GongEnumShapes, rhs, identifierMap)
					case "GongNoteShapes":
						__gong__assignSliceOfPointers(&inst.GongNoteShapes, rhs, identifierMap)
					case "ShowNbInstances":
						inst.ShowNbInstances = GongExtractBool(rhs)
					case "ShowMultiplicity":
						inst.ShowMultiplicity = GongExtractBool(rhs)
					case "ShowLinkNames":
						inst.ShowLinkNames = GongExtractBool(rhs)
					case "IsInRenameMode":
						inst.IsInRenameMode = GongExtractBool(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "NodeGongStructsIsExpanded":
						inst.NodeGongStructsIsExpanded = GongExtractBool(rhs)
					case "NodeGongStructNodeExpansion":
						inst.NodeGongStructNodeExpansion = GongExtractString(rhs)
					case "NodeGongEnumsIsExpanded":
						inst.NodeGongEnumsIsExpanded = GongExtractBool(rhs)
					case "NodeGongEnumNodeExpansion":
						inst.NodeGongEnumNodeExpansion = GongExtractString(rhs)
					case "NodeGongNotesIsExpanded":
						inst.NodeGongNotesIsExpanded = GongExtractBool(rhs)
					case "NodeGongNoteNodeExpansion":
						inst.NodeGongNoteNodeExpansion = GongExtractString(rhs)
					}
				case *DiagramPackage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Path":
						inst.Path = GongExtractString(rhs)
					case "GongModelPath":
						inst.GongModelPath = GongExtractString(rhs)
					case "Classdiagrams":
						__gong__assignSliceOfPointers(&inst.Classdiagrams, rhs, identifierMap)
					case "SelectedClassdiagram":
						__gong__assignPointer(&inst.SelectedClassdiagram, rhs, identifierMap)
					case "AbsolutePathToDiagramPackage":
						inst.AbsolutePathToDiagramPackage = GongExtractString(rhs)
					}
				case *GongEnumShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "IdentifierMeta":
						inst.IdentifierMeta = GongExtractExpr(rhs)
					case "GongEnumValueShapes":
						__gong__assignSliceOfPointers(&inst.GongEnumValueShapes, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *GongEnumValueShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IdentifierMeta":
						inst.IdentifierMeta = GongExtractExpr(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *GongNoteLinkShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Identifier":
						inst.Identifier = GongExtractString(rhs)
					case "Type":
						GongUnmarshallEnum(&inst.Type, rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *GongNoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Identifier":
						inst.Identifier = GongExtractString(rhs)
					case "Body":
						inst.Body = GongExtractString(rhs)
					case "BodyHTML":
						inst.BodyHTML = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "Matched":
						inst.Matched = GongExtractBool(rhs)
					case "GongNoteLinkShapes":
						__gong__assignSliceOfPointers(&inst.GongNoteLinkShapes, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *GongStructShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "IdentifierMeta":
						inst.IdentifierMeta = GongExtractExpr(rhs)
					case "AttributeShapes":
						__gong__assignSliceOfPointers(&inst.AttributeShapes, rhs, identifierMap)
					case "LinkShapes":
						__gong__assignSliceOfPointers(&inst.LinkShapes, rhs, identifierMap)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					}
				case *LinkShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IdentifierMeta":
						inst.IdentifierMeta = GongExtractExpr(rhs)
					case "FieldTypeIdentifierMeta":
						inst.FieldTypeIdentifierMeta = GongExtractExpr(rhs)
					case "FieldOffsetX":
						inst.FieldOffsetX = GongExtractFloat(rhs)
					case "FieldOffsetY":
						inst.FieldOffsetY = GongExtractFloat(rhs)
					case "TargetMultiplicity":
						GongUnmarshallEnum(&inst.TargetMultiplicity, rhs)
					case "TargetMultiplicityOffsetX":
						inst.TargetMultiplicityOffsetX = GongExtractFloat(rhs)
					case "TargetMultiplicityOffsetY":
						inst.TargetMultiplicityOffsetY = GongExtractFloat(rhs)
					case "SourceMultiplicity":
						GongUnmarshallEnum(&inst.SourceMultiplicity, rhs)
					case "SourceMultiplicityOffsetX":
						inst.SourceMultiplicityOffsetX = GongExtractFloat(rhs)
					case "SourceMultiplicityOffsetY":
						inst.SourceMultiplicityOffsetY = GongExtractFloat(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "StartOrientation":
						GongUnmarshallEnum(&inst.StartOrientation, rhs)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndOrientation":
						GongUnmarshallEnum(&inst.EndOrientation, rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
type AttributeShapeUnmarshaller struct{}

func (u *AttributeShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AttributeShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributeShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributeShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "FieldTypeAsString":
		instance.FieldTypeAsString = GongExtractString(valueExpr)
	case "Structname":
		instance.Structname = GongExtractString(valueExpr)
	case "Fieldtypename":
		instance.Fieldtypename = GongExtractString(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type ClassdiagramUnmarshaller struct{}

func (u *ClassdiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Classdiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *ClassdiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Classdiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "IsIncludedInStaticWebSite":
		instance.IsIncludedInStaticWebSite = GongExtractBool(valueExpr)
	case "GongStructShapes":
		GongUnmarshallSliceOfPointers(&instance.GongStructShapes, valueExpr, identifierMap)
	case "GongEnumShapes":
		GongUnmarshallSliceOfPointers(&instance.GongEnumShapes, valueExpr, identifierMap)
	case "GongNoteShapes":
		GongUnmarshallSliceOfPointers(&instance.GongNoteShapes, valueExpr, identifierMap)
	case "ShowNbInstances":
		instance.ShowNbInstances = GongExtractBool(valueExpr)
	case "ShowMultiplicity":
		instance.ShowMultiplicity = GongExtractBool(valueExpr)
	case "ShowLinkNames":
		instance.ShowLinkNames = GongExtractBool(valueExpr)
	case "IsInRenameMode":
		instance.IsInRenameMode = GongExtractBool(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "NodeGongStructsIsExpanded":
		instance.NodeGongStructsIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongStructNodeExpansion":
		instance.NodeGongStructNodeExpansion = GongExtractString(valueExpr)
	case "NodeGongEnumsIsExpanded":
		instance.NodeGongEnumsIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongEnumNodeExpansion":
		instance.NodeGongEnumNodeExpansion = GongExtractString(valueExpr)
	case "NodeGongNotesIsExpanded":
		instance.NodeGongNotesIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongNoteNodeExpansion":
		instance.NodeGongNoteNodeExpansion = GongExtractString(valueExpr)
	}
	return nil
}

type DiagramPackageUnmarshaller struct{}

func (u *DiagramPackageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramPackage), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramPackageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramPackage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Path":
		instance.Path = GongExtractString(valueExpr)
	case "GongModelPath":
		instance.GongModelPath = GongExtractString(valueExpr)
	case "Classdiagrams":
		GongUnmarshallSliceOfPointers(&instance.Classdiagrams, valueExpr, identifierMap)
	case "SelectedClassdiagram":
		GongUnmarshallPointer(&instance.SelectedClassdiagram, valueExpr, identifierMap)
	case "AbsolutePathToDiagramPackage":
		instance.AbsolutePathToDiagramPackage = GongExtractString(valueExpr)
	}
	return nil
}

type GongEnumShapeUnmarshaller struct{}

func (u *GongEnumShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnumShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "GongEnumValueShapes":
		GongUnmarshallSliceOfPointers(&instance.GongEnumValueShapes, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GongEnumValueShapeUnmarshaller struct{}

func (u *GongEnumValueShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnumValueShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumValueShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumValueShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type GongNoteLinkShapeUnmarshaller struct{}

func (u *GongNoteLinkShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongNoteLinkShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongNoteLinkShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNoteLinkShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Identifier":
		instance.Identifier = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type GongNoteShapeUnmarshaller struct{}

func (u *GongNoteShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongNoteShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongNoteShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNoteShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Identifier":
		instance.Identifier = GongExtractString(valueExpr)
	case "Body":
		instance.Body = GongExtractString(valueExpr)
	case "BodyHTML":
		instance.BodyHTML = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "Matched":
		instance.Matched = GongExtractBool(valueExpr)
	case "GongNoteLinkShapes":
		GongUnmarshallSliceOfPointers(&instance.GongNoteLinkShapes, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GongStructShapeUnmarshaller struct{}

func (u *GongStructShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongStructShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongStructShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongStructShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "AttributeShapes":
		GongUnmarshallSliceOfPointers(&instance.AttributeShapes, valueExpr, identifierMap)
	case "LinkShapes":
		GongUnmarshallSliceOfPointers(&instance.LinkShapes, valueExpr, identifierMap)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	}
	return nil
}

type LinkShapeUnmarshaller struct{}

func (u *LinkShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinkShape), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "FieldTypeIdentifierMeta":
		instance.FieldTypeIdentifierMeta = GongExtractExpr(valueExpr)
	case "FieldOffsetX":
		instance.FieldOffsetX = GongExtractFloat(valueExpr)
	case "FieldOffsetY":
		instance.FieldOffsetY = GongExtractFloat(valueExpr)
	case "TargetMultiplicity":
		GongUnmarshallEnum(&instance.TargetMultiplicity, valueExpr)
	case "TargetMultiplicityOffsetX":
		instance.TargetMultiplicityOffsetX = GongExtractFloat(valueExpr)
	case "TargetMultiplicityOffsetY":
		instance.TargetMultiplicityOffsetY = GongExtractFloat(valueExpr)
	case "SourceMultiplicity":
		GongUnmarshallEnum(&instance.SourceMultiplicity, valueExpr)
	case "SourceMultiplicityOffsetX":
		instance.SourceMultiplicityOffsetX = GongExtractFloat(valueExpr)
	case "SourceMultiplicityOffsetY":
		instance.SourceMultiplicityOffsetY = GongExtractFloat(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

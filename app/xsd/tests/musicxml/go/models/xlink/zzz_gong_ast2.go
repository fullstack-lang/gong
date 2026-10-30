// generated code - do not edit
package xlink

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
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
					} else if pkgAlias == "" || pkgAlias == "xlink" || len(inFile.Imports) == 0 {
						pkgAlias = "xlink"
					}

					switch pkgAlias {
			case "xlink":
				switch typeName {
				case "ArcType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArcType), stage, ident.Name, instanceName, preserveOrder)
				case "Extended":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extended), stage, ident.Name, instanceName, preserveOrder)
				case "LocatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LocatorType), stage, ident.Name, instanceName, preserveOrder)
				case "ResourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceType), stage, ident.Name, instanceName, preserveOrder)
				case "Simple":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Simple), stage, ident.Name, instanceName, preserveOrder)
				case "TitleEltType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TitleEltType), stage, ident.Name, instanceName, preserveOrder)
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
				case *ArcType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "From":
						inst.From = GongExtractString(rhs)
					case "To":
						inst.To = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *Extended:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Title_1":
						__gong__assignPointer(&inst.Title_1, rhs, identifierMap)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
					case "Locator":
						__gong__assignPointer(&inst.Locator, rhs, identifierMap)
					case "Arc":
						__gong__assignPointer(&inst.Arc, rhs, identifierMap)
					}
				case *LocatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *ResourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					}
				case *Simple:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *TitleEltType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
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
type ArcTypeUnmarshaller struct{}

func (u *ArcTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArcType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArcTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	case "From":
		instance.From = GongExtractString(valueExpr)
	case "To":
		instance.To = GongExtractString(valueExpr)
	case "Title_1":
		GongUnmarshallSliceOfPointers(&instance.Title_1, valueExpr, identifierMap)
	}
	return nil
}

type ExtendedUnmarshaller struct{}

func (u *ExtendedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Extended), stage, identifier, instanceName, preserveOrder)
}

func (u *ExtendedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extended)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Title_1":
		GongUnmarshallPointer(&instance.Title_1, valueExpr, identifierMap)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
	case "Locator":
		GongUnmarshallPointer(&instance.Locator, valueExpr, identifierMap)
	case "Arc":
		GongUnmarshallPointer(&instance.Arc, valueExpr, identifierMap)
	}
	return nil
}

type LocatorTypeUnmarshaller struct{}

func (u *LocatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LocatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *LocatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LocatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "Title_1":
		GongUnmarshallSliceOfPointers(&instance.Title_1, valueExpr, identifierMap)
	}
	return nil
}

type ResourceTypeUnmarshaller struct{}

func (u *ResourceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ResourceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ResourceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ResourceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	}
	return nil
}

type SimpleUnmarshaller struct{}

func (u *SimpleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Simple), stage, identifier, instanceName, preserveOrder)
}

func (u *SimpleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Simple)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type TitleEltTypeUnmarshaller struct{}

func (u *TitleEltTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TitleEltType), stage, identifier, instanceName, preserveOrder)
}

func (u *TitleEltTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TitleEltType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

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
		case "github.com/fullstack-lang/gong/lib/table/go/models":
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
				case "Button":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Button), stage, ident.Name, instanceName, preserveOrder)
				case "Cell":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cell), stage, ident.Name, instanceName, preserveOrder)
				case "CellBoolean":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellBoolean), stage, ident.Name, instanceName, preserveOrder)
				case "CellFloat64":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellFloat64), stage, ident.Name, instanceName, preserveOrder)
				case "CellIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellIcon), stage, ident.Name, instanceName, preserveOrder)
				case "CellInt":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellInt), stage, ident.Name, instanceName, preserveOrder)
				case "CellString":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CellString), stage, ident.Name, instanceName, preserveOrder)
				case "DisplayedColumn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DisplayedColumn), stage, ident.Name, instanceName, preserveOrder)
				case "Row":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Row), stage, ident.Name, instanceName, preserveOrder)
				case "SVGIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SVGIcon), stage, ident.Name, instanceName, preserveOrder)
				case "Table":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Table), stage, ident.Name, instanceName, preserveOrder)
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
				case *Button:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "SVGIcon":
						__gong__assignPointer(&inst.SVGIcon, rhs, identifierMap)
					case "IsDisabled":
						inst.IsDisabled = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						GongUnmarshallEnum(&inst.ToolTipPosition, rhs)
					}
				case *Cell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CellString":
						__gong__assignPointer(&inst.CellString, rhs, identifierMap)
					case "CellFloat64":
						__gong__assignPointer(&inst.CellFloat64, rhs, identifierMap)
					case "CellInt":
						__gong__assignPointer(&inst.CellInt, rhs, identifierMap)
					case "CellBool":
						__gong__assignPointer(&inst.CellBool, rhs, identifierMap)
					case "CellIcon":
						__gong__assignPointer(&inst.CellIcon, rhs, identifierMap)
					}
				case *CellBoolean:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractBool(rhs)
					}
				case *CellFloat64:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractFloat(rhs)
					}
				case *CellIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Icon":
						inst.Icon = GongExtractString(rhs)
					case "NeedsConfirmation":
						inst.NeedsConfirmation = GongExtractBool(rhs)
					case "ConfirmationMessage":
						inst.ConfirmationMessage = GongExtractString(rhs)
					}
				case *CellInt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *CellString:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *DisplayedColumn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Row:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Cells":
						__gong__assignSliceOfPointers(&inst.Cells, rhs, identifierMap)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					}
				case *SVGIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Table:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DisplayedColumns":
						__gong__assignSliceOfPointers(&inst.DisplayedColumns, rhs, identifierMap)
					case "Rows":
						__gong__assignSliceOfPointers(&inst.Rows, rhs, identifierMap)
					case "HasFiltering":
						inst.HasFiltering = GongExtractBool(rhs)
					case "HasColumnSorting":
						inst.HasColumnSorting = GongExtractBool(rhs)
					case "HasPaginator":
						inst.HasPaginator = GongExtractBool(rhs)
					case "HasCheckableRows":
						inst.HasCheckableRows = GongExtractBool(rhs)
					case "HasSaveButton":
						inst.HasSaveButton = GongExtractBool(rhs)
					case "SaveButtonLabel":
						inst.SaveButtonLabel = GongExtractString(rhs)
					case "HasBulkDeleteButton":
						inst.HasBulkDeleteButton = GongExtractBool(rhs)
					case "BulkDeleteButtonTooltip":
						inst.BulkDeleteButtonTooltip = GongExtractString(rhs)
					case "RowsSelectedForBulkDelete":
						__gong__assignSliceOfPointers(&inst.RowsSelectedForBulkDelete, rhs, identifierMap)
					case "CanDragDropRows":
						inst.CanDragDropRows = GongExtractBool(rhs)
					case "HasCloseButton":
						inst.HasCloseButton = GongExtractBool(rhs)
					case "SavingInProgress":
						inst.SavingInProgress = GongExtractBool(rhs)
					case "NbOfStickyColumns":
						inst.NbOfStickyColumns = GongExtractInt(rhs)
					case "Buttons":
						__gong__assignSliceOfPointers(&inst.Buttons, rhs, identifierMap)
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
	case "Icon":
		instance.Icon = GongExtractString(valueExpr)
	case "SVGIcon":
		GongUnmarshallPointer(&instance.SVGIcon, valueExpr, identifierMap)
	case "IsDisabled":
		instance.IsDisabled = GongExtractBool(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "ToolTipPosition":
		GongUnmarshallEnum(&instance.ToolTipPosition, valueExpr)
	}
	return nil
}

type CellUnmarshaller struct{}

func (u *CellUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Cell), stage, identifier, instanceName, preserveOrder)
}

func (u *CellUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Cell)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CellString":
		GongUnmarshallPointer(&instance.CellString, valueExpr, identifierMap)
	case "CellFloat64":
		GongUnmarshallPointer(&instance.CellFloat64, valueExpr, identifierMap)
	case "CellInt":
		GongUnmarshallPointer(&instance.CellInt, valueExpr, identifierMap)
	case "CellBool":
		GongUnmarshallPointer(&instance.CellBool, valueExpr, identifierMap)
	case "CellIcon":
		GongUnmarshallPointer(&instance.CellIcon, valueExpr, identifierMap)
	}
	return nil
}

type CellBooleanUnmarshaller struct{}

func (u *CellBooleanUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CellBoolean), stage, identifier, instanceName, preserveOrder)
}

func (u *CellBooleanUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CellBoolean)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractBool(valueExpr)
	}
	return nil
}

type CellFloat64Unmarshaller struct{}

func (u *CellFloat64Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CellFloat64), stage, identifier, instanceName, preserveOrder)
}

func (u *CellFloat64Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CellFloat64)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractFloat(valueExpr)
	}
	return nil
}

type CellIconUnmarshaller struct{}

func (u *CellIconUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CellIcon), stage, identifier, instanceName, preserveOrder)
}

func (u *CellIconUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CellIcon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Icon":
		instance.Icon = GongExtractString(valueExpr)
	case "NeedsConfirmation":
		instance.NeedsConfirmation = GongExtractBool(valueExpr)
	case "ConfirmationMessage":
		instance.ConfirmationMessage = GongExtractString(valueExpr)
	}
	return nil
}

type CellIntUnmarshaller struct{}

func (u *CellIntUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CellInt), stage, identifier, instanceName, preserveOrder)
}

func (u *CellIntUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CellInt)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type CellStringUnmarshaller struct{}

func (u *CellStringUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CellString), stage, identifier, instanceName, preserveOrder)
}

func (u *CellStringUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CellString)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type DisplayedColumnUnmarshaller struct{}

func (u *DisplayedColumnUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DisplayedColumn), stage, identifier, instanceName, preserveOrder)
}

func (u *DisplayedColumnUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DisplayedColumn)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type RowUnmarshaller struct{}

func (u *RowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Row), stage, identifier, instanceName, preserveOrder)
}

func (u *RowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Row)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Cells":
		GongUnmarshallSliceOfPointers(&instance.Cells, valueExpr, identifierMap)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	}
	return nil
}

type SVGIconUnmarshaller struct{}

func (u *SVGIconUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SVGIcon), stage, identifier, instanceName, preserveOrder)
}

func (u *SVGIconUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SVGIcon)
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
	case "DisplayedColumns":
		GongUnmarshallSliceOfPointers(&instance.DisplayedColumns, valueExpr, identifierMap)
	case "Rows":
		GongUnmarshallSliceOfPointers(&instance.Rows, valueExpr, identifierMap)
	case "HasFiltering":
		instance.HasFiltering = GongExtractBool(valueExpr)
	case "HasColumnSorting":
		instance.HasColumnSorting = GongExtractBool(valueExpr)
	case "HasPaginator":
		instance.HasPaginator = GongExtractBool(valueExpr)
	case "HasCheckableRows":
		instance.HasCheckableRows = GongExtractBool(valueExpr)
	case "HasSaveButton":
		instance.HasSaveButton = GongExtractBool(valueExpr)
	case "SaveButtonLabel":
		instance.SaveButtonLabel = GongExtractString(valueExpr)
	case "HasBulkDeleteButton":
		instance.HasBulkDeleteButton = GongExtractBool(valueExpr)
	case "BulkDeleteButtonTooltip":
		instance.BulkDeleteButtonTooltip = GongExtractString(valueExpr)
	case "RowsSelectedForBulkDelete":
		GongUnmarshallSliceOfPointers(&instance.RowsSelectedForBulkDelete, valueExpr, identifierMap)
	case "CanDragDropRows":
		instance.CanDragDropRows = GongExtractBool(valueExpr)
	case "HasCloseButton":
		instance.HasCloseButton = GongExtractBool(valueExpr)
	case "SavingInProgress":
		instance.SavingInProgress = GongExtractBool(valueExpr)
	case "NbOfStickyColumns":
		instance.NbOfStickyColumns = GongExtractInt(valueExpr)
	case "Buttons":
		GongUnmarshallSliceOfPointers(&instance.Buttons, valueExpr, identifierMap)
	}
	return nil
}

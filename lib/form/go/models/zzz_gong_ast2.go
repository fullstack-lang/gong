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
		case "github.com/fullstack-lang/gong/lib/form/go/models":
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
				case "CheckBox":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckBox), stage, ident.Name, instanceName, preserveOrder)
				case "FormDiv":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormDiv), stage, ident.Name, instanceName, preserveOrder)
				case "FormEditAssocButton":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormEditAssocButton), stage, ident.Name, instanceName, preserveOrder)
				case "FormField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormField), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldDate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldDate), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldDateTime":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldDateTime), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldFloat64":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldFloat64), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldInt":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldInt), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldSelect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldSelect), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldString":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldString), stage, ident.Name, instanceName, preserveOrder)
				case "FormFieldTime":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormFieldTime), stage, ident.Name, instanceName, preserveOrder)
				case "FormGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormGroup), stage, ident.Name, instanceName, preserveOrder)
				case "FormSortAssocButton":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormSortAssocButton), stage, ident.Name, instanceName, preserveOrder)
				case "Option":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Option), stage, ident.Name, instanceName, preserveOrder)
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
				case *CheckBox:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractBool(rhs)
					}
				case *FormDiv:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FormFields":
						__gong__assignSliceOfPointers(&inst.FormFields, rhs, identifierMap)
					case "CheckBoxs":
						__gong__assignSliceOfPointers(&inst.CheckBoxs, rhs, identifierMap)
					case "FormEditAssocButton":
						__gong__assignPointer(&inst.FormEditAssocButton, rhs, identifierMap)
					case "FormSortAssocButton":
						__gong__assignPointer(&inst.FormSortAssocButton, rhs, identifierMap)
					case "IsADivider":
						inst.IsADivider = GongExtractBool(rhs)
					case "IsAStartAccordionGroup":
						inst.IsAStartAccordionGroup = GongExtractBool(rhs)
					case "AccordionGroupName":
						inst.AccordionGroupName = GongExtractString(rhs)
					case "IsAEndAccordionGroup":
						inst.IsAEndAccordionGroup = GongExtractBool(rhs)
					}
				case *FormEditAssocButton:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "AssociationStorage":
						inst.AssociationStorage = GongExtractString(rhs)
					case "HasChanged":
						inst.HasChanged = GongExtractBool(rhs)
					case "IsForSavePurpose":
						inst.IsForSavePurpose = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "MatTooltipShowDelay":
						inst.MatTooltipShowDelay = GongExtractString(rhs)
					}
				case *FormField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "InputTypeEnum":
						GongUnmarshallEnum(&inst.InputTypeEnum, rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Placeholder":
						inst.Placeholder = GongExtractString(rhs)
					case "FormFieldString":
						__gong__assignPointer(&inst.FormFieldString, rhs, identifierMap)
					case "FormFieldFloat64":
						__gong__assignPointer(&inst.FormFieldFloat64, rhs, identifierMap)
					case "FormFieldInt":
						__gong__assignPointer(&inst.FormFieldInt, rhs, identifierMap)
					case "FormFieldDate":
						__gong__assignPointer(&inst.FormFieldDate, rhs, identifierMap)
					case "FormFieldTime":
						__gong__assignPointer(&inst.FormFieldTime, rhs, identifierMap)
					case "FormFieldDateTime":
						__gong__assignPointer(&inst.FormFieldDateTime, rhs, identifierMap)
					case "FormFieldSelect":
						__gong__assignPointer(&inst.FormFieldSelect, rhs, identifierMap)
					case "HasBespokeWidth":
						inst.HasBespokeWidth = GongExtractBool(rhs)
					case "BespokeWidthPx":
						inst.BespokeWidthPx = GongExtractInt(rhs)
					case "HasBespokeHeight":
						inst.HasBespokeHeight = GongExtractBool(rhs)
					case "BespokeHeightPx":
						inst.BespokeHeightPx = GongExtractInt(rhs)
					}
				case *FormFieldDate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractDate(rhs)
					}
				case *FormFieldDateTime:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractDate(rhs)
					}
				case *FormFieldFloat64:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractFloat(rhs)
					case "HasMinValidator":
						inst.HasMinValidator = GongExtractBool(rhs)
					case "MinValue":
						inst.MinValue = GongExtractFloat(rhs)
					case "HasMaxValidator":
						inst.HasMaxValidator = GongExtractBool(rhs)
					case "MaxValue":
						inst.MaxValue = GongExtractFloat(rhs)
					}
				case *FormFieldInt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					case "HasMinValidator":
						inst.HasMinValidator = GongExtractBool(rhs)
					case "MinValue":
						inst.MinValue = GongExtractInt(rhs)
					case "HasMaxValidator":
						inst.HasMaxValidator = GongExtractBool(rhs)
					case "MaxValue":
						inst.MaxValue = GongExtractInt(rhs)
					}
				case *FormFieldSelect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						__gong__assignPointer(&inst.Value, rhs, identifierMap)
					case "Options":
						__gong__assignSliceOfPointers(&inst.Options, rhs, identifierMap)
					case "CanBeEmpty":
						inst.CanBeEmpty = GongExtractBool(rhs)
					case "PreserveInitialOrder":
						inst.PreserveInitialOrder = GongExtractBool(rhs)
					}
				case *FormFieldString:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					case "IsTextArea":
						inst.IsTextArea = GongExtractBool(rhs)
					}
				case *FormFieldTime:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractDate(rhs)
					case "Step":
						inst.Step = GongExtractFloat(rhs)
					}
				case *FormGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "TypeLabel":
						inst.TypeLabel = GongExtractString(rhs)
					case "FormDivs":
						__gong__assignSliceOfPointers(&inst.FormDivs, rhs, identifierMap)
					case "HasSuppressButton":
						inst.HasSuppressButton = GongExtractBool(rhs)
					case "HasSuppressButtonBeenPressed":
						inst.HasSuppressButtonBeenPressed = GongExtractBool(rhs)
					}
				case *FormSortAssocButton:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "MatTooltipShowDelay":
						inst.MatTooltipShowDelay = GongExtractString(rhs)
					case "FormEditAssocButton":
						__gong__assignPointer(&inst.FormEditAssocButton, rhs, identifierMap)
					}
				case *Option:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
type CheckBoxUnmarshaller struct{}

func (u *CheckBoxUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckBox), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckBoxUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckBox)
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

type FormDivUnmarshaller struct{}

func (u *FormDivUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormDiv), stage, identifier, instanceName, preserveOrder)
}

func (u *FormDivUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormDiv)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FormFields":
		GongUnmarshallSliceOfPointers(&instance.FormFields, valueExpr, identifierMap)
	case "CheckBoxs":
		GongUnmarshallSliceOfPointers(&instance.CheckBoxs, valueExpr, identifierMap)
	case "FormEditAssocButton":
		GongUnmarshallPointer(&instance.FormEditAssocButton, valueExpr, identifierMap)
	case "FormSortAssocButton":
		GongUnmarshallPointer(&instance.FormSortAssocButton, valueExpr, identifierMap)
	case "IsADivider":
		instance.IsADivider = GongExtractBool(valueExpr)
	case "IsAStartAccordionGroup":
		instance.IsAStartAccordionGroup = GongExtractBool(valueExpr)
	case "AccordionGroupName":
		instance.AccordionGroupName = GongExtractString(valueExpr)
	case "IsAEndAccordionGroup":
		instance.IsAEndAccordionGroup = GongExtractBool(valueExpr)
	}
	return nil
}

type FormEditAssocButtonUnmarshaller struct{}

func (u *FormEditAssocButtonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormEditAssocButton), stage, identifier, instanceName, preserveOrder)
}

func (u *FormEditAssocButtonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormEditAssocButton)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "AssociationStorage":
		instance.AssociationStorage = GongExtractString(valueExpr)
	case "HasChanged":
		instance.HasChanged = GongExtractBool(valueExpr)
	case "IsForSavePurpose":
		instance.IsForSavePurpose = GongExtractBool(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "MatTooltipShowDelay":
		instance.MatTooltipShowDelay = GongExtractString(valueExpr)
	}
	return nil
}

type FormFieldUnmarshaller struct{}

func (u *FormFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormField), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "InputTypeEnum":
		GongUnmarshallEnum(&instance.InputTypeEnum, valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "Placeholder":
		instance.Placeholder = GongExtractString(valueExpr)
	case "FormFieldString":
		GongUnmarshallPointer(&instance.FormFieldString, valueExpr, identifierMap)
	case "FormFieldFloat64":
		GongUnmarshallPointer(&instance.FormFieldFloat64, valueExpr, identifierMap)
	case "FormFieldInt":
		GongUnmarshallPointer(&instance.FormFieldInt, valueExpr, identifierMap)
	case "FormFieldDate":
		GongUnmarshallPointer(&instance.FormFieldDate, valueExpr, identifierMap)
	case "FormFieldTime":
		GongUnmarshallPointer(&instance.FormFieldTime, valueExpr, identifierMap)
	case "FormFieldDateTime":
		GongUnmarshallPointer(&instance.FormFieldDateTime, valueExpr, identifierMap)
	case "FormFieldSelect":
		GongUnmarshallPointer(&instance.FormFieldSelect, valueExpr, identifierMap)
	case "HasBespokeWidth":
		instance.HasBespokeWidth = GongExtractBool(valueExpr)
	case "BespokeWidthPx":
		instance.BespokeWidthPx = GongExtractInt(valueExpr)
	case "HasBespokeHeight":
		instance.HasBespokeHeight = GongExtractBool(valueExpr)
	case "BespokeHeightPx":
		instance.BespokeHeightPx = GongExtractInt(valueExpr)
	}
	return nil
}

type FormFieldDateUnmarshaller struct{}

func (u *FormFieldDateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldDate), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldDateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldDate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractDate(valueExpr)
	}
	return nil
}

type FormFieldDateTimeUnmarshaller struct{}

func (u *FormFieldDateTimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldDateTime), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldDateTimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldDateTime)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractDate(valueExpr)
	}
	return nil
}

type FormFieldFloat64Unmarshaller struct{}

func (u *FormFieldFloat64Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldFloat64), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldFloat64Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldFloat64)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractFloat(valueExpr)
	case "HasMinValidator":
		instance.HasMinValidator = GongExtractBool(valueExpr)
	case "MinValue":
		instance.MinValue = GongExtractFloat(valueExpr)
	case "HasMaxValidator":
		instance.HasMaxValidator = GongExtractBool(valueExpr)
	case "MaxValue":
		instance.MaxValue = GongExtractFloat(valueExpr)
	}
	return nil
}

type FormFieldIntUnmarshaller struct{}

func (u *FormFieldIntUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldInt), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldIntUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldInt)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	case "HasMinValidator":
		instance.HasMinValidator = GongExtractBool(valueExpr)
	case "MinValue":
		instance.MinValue = GongExtractInt(valueExpr)
	case "HasMaxValidator":
		instance.HasMaxValidator = GongExtractBool(valueExpr)
	case "MaxValue":
		instance.MaxValue = GongExtractInt(valueExpr)
	}
	return nil
}

type FormFieldSelectUnmarshaller struct{}

func (u *FormFieldSelectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldSelect), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldSelectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldSelect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallPointer(&instance.Value, valueExpr, identifierMap)
	case "Options":
		GongUnmarshallSliceOfPointers(&instance.Options, valueExpr, identifierMap)
	case "CanBeEmpty":
		instance.CanBeEmpty = GongExtractBool(valueExpr)
	case "PreserveInitialOrder":
		instance.PreserveInitialOrder = GongExtractBool(valueExpr)
	}
	return nil
}

type FormFieldStringUnmarshaller struct{}

func (u *FormFieldStringUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldString), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldStringUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldString)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	case "IsTextArea":
		instance.IsTextArea = GongExtractBool(valueExpr)
	}
	return nil
}

type FormFieldTimeUnmarshaller struct{}

func (u *FormFieldTimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldTime), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldTimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldTime)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractDate(valueExpr)
	case "Step":
		instance.Step = GongExtractFloat(valueExpr)
	}
	return nil
}

type FormGroupUnmarshaller struct{}

func (u *FormGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormGroup), stage, identifier, instanceName, preserveOrder)
}

func (u *FormGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "TypeLabel":
		instance.TypeLabel = GongExtractString(valueExpr)
	case "FormDivs":
		GongUnmarshallSliceOfPointers(&instance.FormDivs, valueExpr, identifierMap)
	case "HasSuppressButton":
		instance.HasSuppressButton = GongExtractBool(valueExpr)
	case "HasSuppressButtonBeenPressed":
		instance.HasSuppressButtonBeenPressed = GongExtractBool(valueExpr)
	}
	return nil
}

type FormSortAssocButtonUnmarshaller struct{}

func (u *FormSortAssocButtonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormSortAssocButton), stage, identifier, instanceName, preserveOrder)
}

func (u *FormSortAssocButtonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormSortAssocButton)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "MatTooltipShowDelay":
		instance.MatTooltipShowDelay = GongExtractString(valueExpr)
	case "FormEditAssocButton":
		GongUnmarshallPointer(&instance.FormEditAssocButton, valueExpr, identifierMap)
	}
	return nil
}

type OptionUnmarshaller struct{}

func (u *OptionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Option), stage, identifier, instanceName, preserveOrder)
}

func (u *OptionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Option)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

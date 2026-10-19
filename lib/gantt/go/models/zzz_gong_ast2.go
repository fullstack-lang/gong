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
		case "github.com/fullstack-lang/gong/lib/gantt/go/models":
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
				case "Arrow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arrow), stage, ident.Name, instanceName, preserveOrder)
				case "Bar":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bar), stage, ident.Name, instanceName, preserveOrder)
				case "Gantt":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Gantt), stage, ident.Name, instanceName, preserveOrder)
				case "Group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group), stage, ident.Name, instanceName, preserveOrder)
				case "Lane":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lane), stage, ident.Name, instanceName, preserveOrder)
				case "LaneUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LaneUse), stage, ident.Name, instanceName, preserveOrder)
				case "Milestone":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Milestone), stage, ident.Name, instanceName, preserveOrder)
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
				case *Arrow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "From":
						__gong__assignPointer(&inst.From, rhs, identifierMap)
					case "To":
						__gong__assignPointer(&inst.To, rhs, identifierMap)
					case "OptionnalColor":
						inst.OptionnalColor = GongExtractString(rhs)
					case "OptionnalStroke":
						inst.OptionnalStroke = GongExtractString(rhs)
					}
				case *Bar:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						inst.Start = GongExtractDate(rhs)
					case "End":
						inst.End = GongExtractDate(rhs)
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "OptionnalColor":
						inst.OptionnalColor = GongExtractString(rhs)
					case "OptionnalStroke":
						inst.OptionnalStroke = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					}
				case *Gantt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedStart":
						inst.ComputedStart = GongExtractDate(rhs)
					case "ComputedEnd":
						inst.ComputedEnd = GongExtractDate(rhs)
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "UseManualStartAndEndDates":
						inst.UseManualStartAndEndDates = GongExtractBool(rhs)
					case "ManualStart":
						inst.ManualStart = GongExtractDate(rhs)
					case "ManualEnd":
						inst.ManualEnd = GongExtractDate(rhs)
					case "LaneHeight":
						inst.LaneHeight = GongExtractFloat(rhs)
					case "RatioBarToLaneHeight":
						inst.RatioBarToLaneHeight = GongExtractFloat(rhs)
					case "YTopMargin":
						inst.YTopMargin = GongExtractFloat(rhs)
					case "XLeftText":
						inst.XLeftText = GongExtractFloat(rhs)
					case "TextHeight":
						inst.TextHeight = GongExtractFloat(rhs)
					case "XLeftLanes":
						inst.XLeftLanes = GongExtractFloat(rhs)
					case "XRightMargin":
						inst.XRightMargin = GongExtractFloat(rhs)
					case "ArrowLengthToTheRightOfStartBar":
						inst.ArrowLengthToTheRightOfStartBar = GongExtractFloat(rhs)
					case "ArrowTipLenght":
						inst.ArrowTipLenght = GongExtractFloat(rhs)
					case "TimeLine_Color":
						inst.TimeLine_Color = GongExtractString(rhs)
					case "TimeLine_FillOpacity":
						inst.TimeLine_FillOpacity = GongExtractFloat(rhs)
					case "TimeLine_Stroke":
						inst.TimeLine_Stroke = GongExtractString(rhs)
					case "TimeLine_StrokeWidth":
						inst.TimeLine_StrokeWidth = GongExtractFloat(rhs)
					case "Group_Stroke":
						inst.Group_Stroke = GongExtractString(rhs)
					case "Group_StrokeWidth":
						inst.Group_StrokeWidth = GongExtractFloat(rhs)
					case "Group_StrokeDashArray":
						inst.Group_StrokeDashArray = GongExtractString(rhs)
					case "DateYOffset":
						inst.DateYOffset = GongExtractFloat(rhs)
					case "AlignOnStartEndOnYearStart":
						inst.AlignOnStartEndOnYearStart = GongExtractBool(rhs)
					case "Lanes":
						__gong__assignSliceOfPointers(&inst.Lanes, rhs, identifierMap)
					case "Milestones":
						__gong__assignSliceOfPointers(&inst.Milestones, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Arrows":
						__gong__assignSliceOfPointers(&inst.Arrows, rhs, identifierMap)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GroupLanes":
						__gong__assignSliceOfPointers(&inst.GroupLanes, rhs, identifierMap)
					}
				case *Lane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Bars":
						__gong__assignSliceOfPointers(&inst.Bars, rhs, identifierMap)
					}
				case *LaneUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lane":
						__gong__assignPointer(&inst.Lane, rhs, identifierMap)
					}
				case *Milestone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						inst.Date = GongExtractDate(rhs)
					case "DisplayVerticalBar":
						inst.DisplayVerticalBar = GongExtractBool(rhs)
					case "LanesToDisplay":
						__gong__assignSliceOfPointers(&inst.LanesToDisplay, rhs, identifierMap)
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
type ArrowUnmarshaller struct{}

func (u *ArrowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Arrow), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Arrow)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "From":
		GongUnmarshallPointer(&instance.From, valueExpr, identifierMap)
	case "To":
		GongUnmarshallPointer(&instance.To, valueExpr, identifierMap)
	case "OptionnalColor":
		instance.OptionnalColor = GongExtractString(valueExpr)
	case "OptionnalStroke":
		instance.OptionnalStroke = GongExtractString(valueExpr)
	}
	return nil
}

type BarUnmarshaller struct{}

func (u *BarUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bar), stage, identifier, instanceName, preserveOrder)
}

func (u *BarUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bar)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Start":
		instance.Start = GongExtractDate(valueExpr)
	case "End":
		instance.End = GongExtractDate(valueExpr)
	case "ComputedDuration":
		instance.ComputedDuration = time.Duration(GongExtractInt(valueExpr))
	case "OptionnalColor":
		instance.OptionnalColor = GongExtractString(valueExpr)
	case "OptionnalStroke":
		instance.OptionnalStroke = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	}
	return nil
}

type GanttUnmarshaller struct{}

func (u *GanttUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Gantt), stage, identifier, instanceName, preserveOrder)
}

func (u *GanttUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Gantt)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedStart":
		instance.ComputedStart = GongExtractDate(valueExpr)
	case "ComputedEnd":
		instance.ComputedEnd = GongExtractDate(valueExpr)
	case "ComputedDuration":
		instance.ComputedDuration = time.Duration(GongExtractInt(valueExpr))
	case "UseManualStartAndEndDates":
		instance.UseManualStartAndEndDates = GongExtractBool(valueExpr)
	case "ManualStart":
		instance.ManualStart = GongExtractDate(valueExpr)
	case "ManualEnd":
		instance.ManualEnd = GongExtractDate(valueExpr)
	case "LaneHeight":
		instance.LaneHeight = GongExtractFloat(valueExpr)
	case "RatioBarToLaneHeight":
		instance.RatioBarToLaneHeight = GongExtractFloat(valueExpr)
	case "YTopMargin":
		instance.YTopMargin = GongExtractFloat(valueExpr)
	case "XLeftText":
		instance.XLeftText = GongExtractFloat(valueExpr)
	case "TextHeight":
		instance.TextHeight = GongExtractFloat(valueExpr)
	case "XLeftLanes":
		instance.XLeftLanes = GongExtractFloat(valueExpr)
	case "XRightMargin":
		instance.XRightMargin = GongExtractFloat(valueExpr)
	case "ArrowLengthToTheRightOfStartBar":
		instance.ArrowLengthToTheRightOfStartBar = GongExtractFloat(valueExpr)
	case "ArrowTipLenght":
		instance.ArrowTipLenght = GongExtractFloat(valueExpr)
	case "TimeLine_Color":
		instance.TimeLine_Color = GongExtractString(valueExpr)
	case "TimeLine_FillOpacity":
		instance.TimeLine_FillOpacity = GongExtractFloat(valueExpr)
	case "TimeLine_Stroke":
		instance.TimeLine_Stroke = GongExtractString(valueExpr)
	case "TimeLine_StrokeWidth":
		instance.TimeLine_StrokeWidth = GongExtractFloat(valueExpr)
	case "Group_Stroke":
		instance.Group_Stroke = GongExtractString(valueExpr)
	case "Group_StrokeWidth":
		instance.Group_StrokeWidth = GongExtractFloat(valueExpr)
	case "Group_StrokeDashArray":
		instance.Group_StrokeDashArray = GongExtractString(valueExpr)
	case "DateYOffset":
		instance.DateYOffset = GongExtractFloat(valueExpr)
	case "AlignOnStartEndOnYearStart":
		instance.AlignOnStartEndOnYearStart = GongExtractBool(valueExpr)
	case "Lanes":
		GongUnmarshallSliceOfPointers(&instance.Lanes, valueExpr, identifierMap)
	case "Milestones":
		GongUnmarshallSliceOfPointers(&instance.Milestones, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Arrows":
		GongUnmarshallSliceOfPointers(&instance.Arrows, valueExpr, identifierMap)
	}
	return nil
}

type GroupUnmarshaller struct{}

func (u *GroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GroupLanes":
		GongUnmarshallSliceOfPointers(&instance.GroupLanes, valueExpr, identifierMap)
	}
	return nil
}

type LaneUnmarshaller struct{}

func (u *LaneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Lane), stage, identifier, instanceName, preserveOrder)
}

func (u *LaneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lane)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Bars":
		GongUnmarshallSliceOfPointers(&instance.Bars, valueExpr, identifierMap)
	}
	return nil
}

type LaneUseUnmarshaller struct{}

func (u *LaneUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LaneUse), stage, identifier, instanceName, preserveOrder)
}

func (u *LaneUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LaneUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lane":
		GongUnmarshallPointer(&instance.Lane, valueExpr, identifierMap)
	}
	return nil
}

type MilestoneUnmarshaller struct{}

func (u *MilestoneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Milestone), stage, identifier, instanceName, preserveOrder)
}

func (u *MilestoneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Milestone)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	case "DisplayVerticalBar":
		instance.DisplayVerticalBar = GongExtractBool(valueExpr)
	case "LanesToDisplay":
		GongUnmarshallSliceOfPointers(&instance.LanesToDisplay, valueExpr, identifierMap)
	}
	return nil
}

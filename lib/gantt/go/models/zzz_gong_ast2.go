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

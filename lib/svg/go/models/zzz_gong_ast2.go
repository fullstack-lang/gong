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
type AnimateUnmarshaller struct{}

func (u *AnimateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Animate), stage, identifier, instanceName, preserveOrder)
}

func (u *AnimateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Animate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "AttributeName":
		instance.AttributeName = GongExtractString(valueExpr)
	case "Values":
		instance.Values = GongExtractString(valueExpr)
	case "From":
		instance.From = GongExtractString(valueExpr)
	case "To":
		instance.To = GongExtractString(valueExpr)
	case "Dur":
		instance.Dur = GongExtractString(valueExpr)
	case "RepeatCount":
		instance.RepeatCount = GongExtractString(valueExpr)
	}
	return nil
}

type CircleUnmarshaller struct{}

func (u *CircleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Circle), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Circle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CX":
		instance.CX = GongExtractFloat(valueExpr)
	case "CY":
		instance.CY = GongExtractFloat(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animations":
		GongUnmarshallSliceOfPointers(&instance.Animations, valueExpr, identifierMap)
	}
	return nil
}

type ConditionUnmarshaller struct{}

func (u *ConditionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Condition), stage, identifier, instanceName, preserveOrder)
}

func (u *ConditionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Condition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ControlPointUnmarshaller struct{}

func (u *ControlPointUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlPoint), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlPointUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPoint)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "ClosestRect":
		GongUnmarshallPointer(&instance.ClosestRect, valueExpr, identifierMap)
	}
	return nil
}

type EllipseUnmarshaller struct{}

func (u *EllipseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Ellipse), stage, identifier, instanceName, preserveOrder)
}

func (u *EllipseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ellipse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CX":
		instance.CX = GongExtractFloat(valueExpr)
	case "CY":
		instance.CY = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "RY":
		instance.RY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type FileToDownloadUnmarshaller struct{}

func (u *FileToDownloadUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FileToDownload), stage, identifier, instanceName, preserveOrder)
}

func (u *FileToDownloadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FileToDownload)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64EncodedContent":
		instance.Base64EncodedContent = GongExtractString(valueExpr)
	}
	return nil
}

type LayerUnmarshaller struct{}

func (u *LayerUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Layer), stage, identifier, instanceName, preserveOrder)
}

func (u *LayerUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Layer)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Rects":
		GongUnmarshallSliceOfPointers(&instance.Rects, valueExpr, identifierMap)
	case "Texts":
		GongUnmarshallSliceOfPointers(&instance.Texts, valueExpr, identifierMap)
	case "Circles":
		GongUnmarshallSliceOfPointers(&instance.Circles, valueExpr, identifierMap)
	case "Lines":
		GongUnmarshallSliceOfPointers(&instance.Lines, valueExpr, identifierMap)
	case "Ellipses":
		GongUnmarshallSliceOfPointers(&instance.Ellipses, valueExpr, identifierMap)
	case "Polylines":
		GongUnmarshallSliceOfPointers(&instance.Polylines, valueExpr, identifierMap)
	case "Polygones":
		GongUnmarshallSliceOfPointers(&instance.Polygones, valueExpr, identifierMap)
	case "Paths":
		GongUnmarshallSliceOfPointers(&instance.Paths, valueExpr, identifierMap)
	case "Links":
		GongUnmarshallSliceOfPointers(&instance.Links, valueExpr, identifierMap)
	case "RectLinkLinks":
		GongUnmarshallSliceOfPointers(&instance.RectLinkLinks, valueExpr, identifierMap)
	}
	return nil
}

type LineUnmarshaller struct{}

func (u *LineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Line), stage, identifier, instanceName, preserveOrder)
}

func (u *LineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X1":
		instance.X1 = GongExtractFloat(valueExpr)
	case "Y1":
		instance.Y1 = GongExtractFloat(valueExpr)
	case "X2":
		instance.X2 = GongExtractFloat(valueExpr)
	case "Y2":
		instance.Y2 = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	case "MouseClickX":
		instance.MouseClickX = GongExtractFloat(valueExpr)
	case "MouseClickY":
		instance.MouseClickY = GongExtractFloat(valueExpr)
	}
	return nil
}

type LinkUnmarshaller struct{}

func (u *LinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Link), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "IsBezierCurve":
		instance.IsBezierCurve = GongExtractBool(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "StartAnchorType":
		GongUnmarshallEnum(&instance.StartAnchorType, valueExpr)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "EndAnchorType":
		GongUnmarshallEnum(&instance.EndAnchorType, valueExpr)
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
	case "CornerRadius":
		instance.CornerRadius = GongExtractFloat(valueExpr)
	case "HasEndArrow":
		instance.HasEndArrow = GongExtractBool(valueExpr)
	case "EndArrowSize":
		instance.EndArrowSize = GongExtractFloat(valueExpr)
	case "EndArrowOffset":
		instance.EndArrowOffset = GongExtractFloat(valueExpr)
	case "HasStartArrow":
		instance.HasStartArrow = GongExtractBool(valueExpr)
	case "StartArrowSize":
		instance.StartArrowSize = GongExtractFloat(valueExpr)
	case "StartArrowOffset":
		instance.StartArrowOffset = GongExtractFloat(valueExpr)
	case "TextAtArrowStart":
		GongUnmarshallSliceOfPointers(&instance.TextAtArrowStart, valueExpr, identifierMap)
	case "TextAtArrowEnd":
		GongUnmarshallSliceOfPointers(&instance.TextAtArrowEnd, valueExpr, identifierMap)
	case "TextAtCorner":
		GongUnmarshallSliceOfPointers(&instance.TextAtCorner, valueExpr, identifierMap)
	case "PathAtArrowStart":
		GongUnmarshallSliceOfPointers(&instance.PathAtArrowStart, valueExpr, identifierMap)
	case "PathAtArrowEnd":
		GongUnmarshallSliceOfPointers(&instance.PathAtArrowEnd, valueExpr, identifierMap)
	case "PathAtCorner":
		GongUnmarshallSliceOfPointers(&instance.PathAtCorner, valueExpr, identifierMap)
	case "ControlPoints":
		GongUnmarshallSliceOfPointers(&instance.ControlPoints, valueExpr, identifierMap)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "MouseX":
		instance.MouseX = GongExtractFloat(valueExpr)
	case "MouseY":
		instance.MouseY = GongExtractFloat(valueExpr)
	case "MouseEventKey":
		GongUnmarshallEnum(&instance.MouseEventKey, valueExpr)
	}
	return nil
}

type LinkAnchoredPathUnmarshaller struct{}

func (u *LinkAnchoredPathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinkAnchoredPath), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkAnchoredPath)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "ScalePropotionnally":
		instance.ScalePropotionnally = GongExtractBool(valueExpr)
	case "AppliedScaling":
		instance.AppliedScaling = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type LinkAnchoredTextUnmarshaller struct{}

func (u *LinkAnchoredTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinkAnchoredText), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkAnchoredTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkAnchoredText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "AutomaticLayout":
		instance.AutomaticLayout = GongExtractBool(valueExpr)
	case "LinkAnchorType":
		GongUnmarshallEnum(&instance.LinkAnchorType, valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontStyle":
		instance.FontStyle = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PathUnmarshaller struct{}

func (u *PathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Path), stage, identifier, instanceName, preserveOrder)
}

func (u *PathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Path)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PointUnmarshaller struct{}

func (u *PointUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Point), stage, identifier, instanceName, preserveOrder)
}

func (u *PointUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Point)
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

type PolygoneUnmarshaller struct{}

func (u *PolygoneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Polygone), stage, identifier, instanceName, preserveOrder)
}

func (u *PolygoneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Polygone)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		instance.Points = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PolylineUnmarshaller struct{}

func (u *PolylineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Polyline), stage, identifier, instanceName, preserveOrder)
}

func (u *PolylineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Polyline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		instance.Points = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type RectUnmarshaller struct{}

func (u *RectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Rect), stage, identifier, instanceName, preserveOrder)
}

func (u *RectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rect)
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
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "Peers":
		GongUnmarshallSliceOfPointers(&instance.Peers, valueExpr, identifierMap)
	case "EnclosingRect":
		GongUnmarshallPointer(&instance.EnclosingRect, valueExpr, identifierMap)
	case "Obstacles":
		GongUnmarshallSliceOfPointers(&instance.Obstacles, valueExpr, identifierMap)
	case "AnchoredTo":
		GongUnmarshallPointer(&instance.AnchoredTo, valueExpr, identifierMap)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "HoveringTrigger":
		GongUnmarshallSliceOfPointers(&instance.HoveringTrigger, valueExpr, identifierMap)
	case "DisplayConditions":
		GongUnmarshallSliceOfPointers(&instance.DisplayConditions, valueExpr, identifierMap)
	case "Animations":
		GongUnmarshallSliceOfPointers(&instance.Animations, valueExpr, identifierMap)
	case "IsSelectable":
		instance.IsSelectable = GongExtractBool(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "CanHaveLeftHandle":
		instance.CanHaveLeftHandle = GongExtractBool(valueExpr)
	case "HasLeftHandle":
		instance.HasLeftHandle = GongExtractBool(valueExpr)
	case "CanHaveRightHandle":
		instance.CanHaveRightHandle = GongExtractBool(valueExpr)
	case "HasRightHandle":
		instance.HasRightHandle = GongExtractBool(valueExpr)
	case "CanHaveTopHandle":
		instance.CanHaveTopHandle = GongExtractBool(valueExpr)
	case "HasTopHandle":
		instance.HasTopHandle = GongExtractBool(valueExpr)
	case "IsScalingProportionally":
		instance.IsScalingProportionally = GongExtractBool(valueExpr)
	case "CanHaveBottomHandle":
		instance.CanHaveBottomHandle = GongExtractBool(valueExpr)
	case "HasBottomHandle":
		instance.HasBottomHandle = GongExtractBool(valueExpr)
	case "CanMoveHorizontaly":
		instance.CanMoveHorizontaly = GongExtractBool(valueExpr)
	case "CanMoveVerticaly":
		instance.CanMoveVerticaly = GongExtractBool(valueExpr)
	case "RectAnchoredTexts":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredTexts, valueExpr, identifierMap)
	case "RectAnchoredRects":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredRects, valueExpr, identifierMap)
	case "RectAnchoredPaths":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredPaths, valueExpr, identifierMap)
	case "RectAnchoredPngImages":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredPngImages, valueExpr, identifierMap)
	case "ChangeColorWhenHovered":
		instance.ChangeColorWhenHovered = GongExtractBool(valueExpr)
	case "ColorWhenHovered":
		instance.ColorWhenHovered = GongExtractString(valueExpr)
	case "OriginalColor":
		instance.OriginalColor = GongExtractString(valueExpr)
	case "FillOpacityWhenHovered":
		instance.FillOpacityWhenHovered = GongExtractFloat(valueExpr)
	case "OriginalFillOpacity":
		instance.OriginalFillOpacity = GongExtractFloat(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "ToolTipPosition":
		GongUnmarshallEnum(&instance.ToolTipPosition, valueExpr)
	case "MouseX":
		instance.MouseX = GongExtractFloat(valueExpr)
	case "MouseY":
		instance.MouseY = GongExtractFloat(valueExpr)
	case "MouseEventKey":
		GongUnmarshallEnum(&instance.MouseEventKey, valueExpr)
	case "URLPath":
		instance.URLPath = GongExtractString(valueExpr)
	case "URLTarget":
		GongUnmarshallEnum(&instance.URLTarget, valueExpr)
	}
	return nil
}

type RectAnchoredPathUnmarshaller struct{}

func (u *RectAnchoredPathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredPath), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredPath)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "ScalePropotionnally":
		instance.ScalePropotionnally = GongExtractBool(valueExpr)
	case "AppliedScaling":
		instance.AppliedScaling = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredPngImageUnmarshaller struct{}

func (u *RectAnchoredPngImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredPngImage), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredPngImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredPngImage)
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
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredRectUnmarshaller struct{}

func (u *RectAnchoredRectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredRect), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredRectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredRect)
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
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "WidthFollowRect":
		instance.WidthFollowRect = GongExtractBool(valueExpr)
	case "HeightFollowRect":
		instance.HeightFollowRect = GongExtractBool(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredTextUnmarshaller struct{}

func (u *RectAnchoredTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredText), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontStyle":
		instance.FontStyle = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "TextAnchorType":
		GongUnmarshallEnum(&instance.TextAnchorType, valueExpr)
	case "DominantBaseline":
		GongUnmarshallEnum(&instance.DominantBaseline, valueExpr)
	case "WritingMode":
		GongUnmarshallEnum(&instance.WritingMode, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	case "URLPath":
		instance.URLPath = GongExtractString(valueExpr)
	case "URLTarget":
		GongUnmarshallEnum(&instance.URLTarget, valueExpr)
	}
	return nil
}

type RectLinkLinkUnmarshaller struct{}

func (u *RectLinkLinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectLinkLink), stage, identifier, instanceName, preserveOrder)
}

func (u *RectLinkLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectLinkLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "TargetAnchorPosition":
		instance.TargetAnchorPosition = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type SVGUnmarshaller struct{}

func (u *SVGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SVG), stage, identifier, instanceName, preserveOrder)
}

func (u *SVGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SVG)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Layers":
		GongUnmarshallSliceOfPointers(&instance.Layers, valueExpr, identifierMap)
	case "DrawingState":
		GongUnmarshallEnum(&instance.DrawingState, valueExpr)
	case "StartRect":
		GongUnmarshallPointer(&instance.StartRect, valueExpr, identifierMap)
	case "EndRect":
		GongUnmarshallPointer(&instance.EndRect, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsSVGFrontEndFileGenerated":
		instance.IsSVGFrontEndFileGenerated = GongExtractBool(valueExpr)
	case "IsSVGBackEndFileGenerated":
		instance.IsSVGBackEndFileGenerated = GongExtractBool(valueExpr)
	case "DefaultDirectoryForGeneratedImages":
		instance.DefaultDirectoryForGeneratedImages = GongExtractString(valueExpr)
	case "IsControlBannerHidden":
		instance.IsControlBannerHidden = GongExtractBool(valueExpr)
	case "PanX":
		instance.PanX = GongExtractFloat(valueExpr)
	case "PanY":
		instance.PanY = GongExtractFloat(valueExpr)
	case "Zoom":
		instance.Zoom = GongExtractFloat(valueExpr)
	case "OverrideWidth":
		instance.OverrideWidth = GongExtractBool(valueExpr)
	case "OverriddenWidth":
		instance.OverriddenWidth = GongExtractFloat(valueExpr)
	case "OverrideHeight":
		instance.OverrideHeight = GongExtractBool(valueExpr)
	case "OverriddenHeight":
		instance.OverriddenHeight = GongExtractFloat(valueExpr)
	}
	return nil
}

type SvgTextUnmarshaller struct{}

func (u *SvgTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SvgText), stage, identifier, instanceName, preserveOrder)
}

func (u *SvgTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SvgText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	}
	return nil
}

type TextUnmarshaller struct{}

func (u *TextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Text), stage, identifier, instanceName, preserveOrder)
}

func (u *TextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontStyle":
		instance.FontStyle = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

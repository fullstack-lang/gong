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
type ArtefactTypeUnmarshaller struct{}

func (u *ArtefactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArtefactType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtefactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtefactType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ArtefactTypeShapeUnmarshaller struct{}

func (u *ArtefactTypeShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArtefactTypeShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtefactTypeShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtefactTypeShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ArtefactType":
		GongUnmarshallPointer(&instance.ArtefactType, valueExpr, identifierMap)
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
	}
	return nil
}

type ArtistUnmarshaller struct{}

func (u *ArtistUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Artist), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtistUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Artist)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsDead":
		instance.IsDead = GongExtractBool(valueExpr)
	case "DateOfDeath":
		instance.DateOfDeath = GongExtractDate(valueExpr)
	case "Place":
		GongUnmarshallPointer(&instance.Place, valueExpr, identifierMap)
	}
	return nil
}

type ArtistShapeUnmarshaller struct{}

func (u *ArtistShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArtistShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtistShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtistShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Artist":
		GongUnmarshallPointer(&instance.Artist, valueExpr, identifierMap)
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
	case "ImagePng_X":
		instance.ImagePng_X = GongExtractFloat(valueExpr)
	case "ImagePng_Y":
		instance.ImagePng_Y = GongExtractFloat(valueExpr)
	case "ImagePng_Width":
		instance.ImagePng_Width = GongExtractFloat(valueExpr)
	case "ImagePng_Height":
		instance.ImagePng_Height = GongExtractFloat(valueExpr)
	case "ImagePng_X_Offset":
		instance.ImagePng_X_Offset = GongExtractFloat(valueExpr)
	case "ImagePng_Y_Offset":
		instance.ImagePng_Y_Offset = GongExtractFloat(valueExpr)
	case "ImagePng_RectAnchorType":
		GongUnmarshallEnum(&instance.ImagePng_RectAnchorType, valueExpr)
	case "ImagePngBase64Content":
		instance.ImagePngBase64Content = GongExtractString(valueExpr)
	}
	return nil
}

type ControlPointShapeUnmarshaller struct{}

func (u *ControlPointShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlPointShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlPointShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPointShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "IsStartShapeTheClosestShape":
		instance.IsStartShapeTheClosestShape = GongExtractBool(valueExpr)
	}
	return nil
}

type DeskUnmarshaller struct{}

func (u *DeskUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Desk), stage, identifier, instanceName, preserveOrder)
}

func (u *DeskUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Desk)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SelectedDiagram":
		GongUnmarshallPointer(&instance.SelectedDiagram, valueExpr, identifierMap)
	}
	return nil
}

type DiagramUnmarshaller struct{}

func (u *DiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Diagram), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Diagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "MovementShapes":
		GongUnmarshallSliceOfPointers(&instance.MovementShapes, valueExpr, identifierMap)
	case "ArtefactTypeShapes":
		GongUnmarshallSliceOfPointers(&instance.ArtefactTypeShapes, valueExpr, identifierMap)
	case "ArtistShapes":
		GongUnmarshallSliceOfPointers(&instance.ArtistShapes, valueExpr, identifierMap)
	case "InfluenceShapes":
		GongUnmarshallSliceOfPointers(&instance.InfluenceShapes, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsNodeExpanded":
		instance.IsNodeExpanded = GongExtractBool(valueExpr)
	case "IsMovementCategoryNodeExpanded":
		instance.IsMovementCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsArtefactTypeCategoryNodeExpanded":
		instance.IsArtefactTypeCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsArtistCategoryNodeExpanded":
		instance.IsArtistCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryNodeExpanded":
		instance.IsInfluenceCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsMovementCategoryHidden":
		instance.IsMovementCategoryHidden = GongExtractBool(valueExpr)
	case "IsArtefactTypeCategoryHidden":
		instance.IsArtefactTypeCategoryHidden = GongExtractBool(valueExpr)
	case "IsArtistCategoryHidden":
		instance.IsArtistCategoryHidden = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryHidden":
		instance.IsInfluenceCategoryHidden = GongExtractBool(valueExpr)
	case "StartDate":
		instance.StartDate = GongExtractDate(valueExpr)
	case "EndDate":
		instance.EndDate = GongExtractDate(valueExpr)
	case "NbYearsForIntervals":
		instance.NbYearsForIntervals = GongExtractInt(valueExpr)
	case "XMargin":
		instance.XMargin = GongExtractFloat(valueExpr)
	case "YMargin":
		instance.YMargin = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "NextVerticalDateXMargin":
		instance.NextVerticalDateXMargin = GongExtractFloat(valueExpr)
	case "RedColorCode":
		instance.RedColorCode = GongExtractString(valueExpr)
	case "BackgroundGreyColorCode":
		instance.BackgroundGreyColorCode = GongExtractString(valueExpr)
	case "GrayColorCode":
		instance.GrayColorCode = GongExtractString(valueExpr)
	case "BottomBoxYOffset":
		instance.BottomBoxYOffset = GongExtractFloat(valueExpr)
	case "BottomBoxWidth":
		instance.BottomBoxWidth = GongExtractFloat(valueExpr)
	case "BottomBoxHeigth":
		instance.BottomBoxHeigth = GongExtractFloat(valueExpr)
	case "BottomBoxFontSize":
		instance.BottomBoxFontSize = GongExtractString(valueExpr)
	case "BottomBoxFontWeigth":
		instance.BottomBoxFontWeigth = GongExtractString(valueExpr)
	case "BottomBoxFontFamily":
		instance.BottomBoxFontFamily = GongExtractString(valueExpr)
	case "BottomBoxLetterSpacing":
		instance.BottomBoxLetterSpacing = GongExtractString(valueExpr)
	case "BottomBoxLetterColorCode":
		instance.BottomBoxLetterColorCode = GongExtractString(valueExpr)
	case "MovementRectAnchorType":
		GongUnmarshallEnum(&instance.MovementRectAnchorType, valueExpr)
	case "MovementTextAnchorType":
		GongUnmarshallEnum(&instance.MovementTextAnchorType, valueExpr)
	case "MovementDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementDominantBaselineType, valueExpr)
	case "MovementFontSize":
		instance.MovementFontSize = GongExtractString(valueExpr)
	case "MajorMovementFontSize":
		instance.MajorMovementFontSize = GongExtractString(valueExpr)
	case "MinorMovementFontSize":
		instance.MinorMovementFontSize = GongExtractString(valueExpr)
	case "MovementFontWeigth":
		instance.MovementFontWeigth = GongExtractString(valueExpr)
	case "MovementFontFamily":
		instance.MovementFontFamily = GongExtractString(valueExpr)
	case "MovementLetterSpacing":
		instance.MovementLetterSpacing = GongExtractString(valueExpr)
	case "AbstractMovementFontSize":
		instance.AbstractMovementFontSize = GongExtractString(valueExpr)
	case "AbstractMovementRectAnchorType":
		GongUnmarshallEnum(&instance.AbstractMovementRectAnchorType, valueExpr)
	case "AbstractMovementTextAnchorType":
		GongUnmarshallEnum(&instance.AbstractMovementTextAnchorType, valueExpr)
	case "AbstractDominantBaselineType":
		GongUnmarshallEnum(&instance.AbstractDominantBaselineType, valueExpr)
	case "MovementDateRectAnchorType":
		GongUnmarshallEnum(&instance.MovementDateRectAnchorType, valueExpr)
	case "MovementDateTextAnchorType":
		GongUnmarshallEnum(&instance.MovementDateTextAnchorType, valueExpr)
	case "MovementDateTextDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementDateTextDominantBaselineType, valueExpr)
	case "MovementDateAndPlacesFontSize":
		instance.MovementDateAndPlacesFontSize = GongExtractString(valueExpr)
	case "MovementDateAndPlacesFontWeigth":
		instance.MovementDateAndPlacesFontWeigth = GongExtractString(valueExpr)
	case "MovementDateAndPlacesFontFamily":
		instance.MovementDateAndPlacesFontFamily = GongExtractString(valueExpr)
	case "MovementDateAndPlacesLetterSpacing":
		instance.MovementDateAndPlacesLetterSpacing = GongExtractString(valueExpr)
	case "MovementBelowArcY_Offset":
		instance.MovementBelowArcY_Offset = GongExtractFloat(valueExpr)
	case "MovementBelowArcY_OffsetPerPlace":
		instance.MovementBelowArcY_OffsetPerPlace = GongExtractFloat(valueExpr)
	case "MovementPlacesRectAnchorType":
		GongUnmarshallEnum(&instance.MovementPlacesRectAnchorType, valueExpr)
	case "MovementPlacesTextAnchorType":
		GongUnmarshallEnum(&instance.MovementPlacesTextAnchorType, valueExpr)
	case "MovementPlacesDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementPlacesDominantBaselineType, valueExpr)
	case "ArtefactTypeFontSize":
		instance.ArtefactTypeFontSize = GongExtractString(valueExpr)
	case "ArtefactTypeFontWeigth":
		instance.ArtefactTypeFontWeigth = GongExtractString(valueExpr)
	case "ArtefactTypeFontFamily":
		instance.ArtefactTypeFontFamily = GongExtractString(valueExpr)
	case "ArtefactTypeLetterSpacing":
		instance.ArtefactTypeLetterSpacing = GongExtractString(valueExpr)
	case "ArtefactTypeRectAnchorType":
		GongUnmarshallEnum(&instance.ArtefactTypeRectAnchorType, valueExpr)
	case "ArtefactDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtefactDominantBaselineType, valueExpr)
	case "ArtefactTypeStrokeWidth":
		instance.ArtefactTypeStrokeWidth = GongExtractFloat(valueExpr)
	case "ArtistRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistRectAnchorType, valueExpr)
	case "ArtistTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistTextAnchorType, valueExpr)
	case "ArtistDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistDominantBaselineType, valueExpr)
	case "ArtistFontSize":
		instance.ArtistFontSize = GongExtractString(valueExpr)
	case "MajorArtistFontSize":
		instance.MajorArtistFontSize = GongExtractString(valueExpr)
	case "MinorArtistFontSize":
		instance.MinorArtistFontSize = GongExtractString(valueExpr)
	case "ArtistFontWeigth":
		instance.ArtistFontWeigth = GongExtractString(valueExpr)
	case "ArtistFontFamily":
		instance.ArtistFontFamily = GongExtractString(valueExpr)
	case "ArtistLetterSpacing":
		instance.ArtistLetterSpacing = GongExtractString(valueExpr)
	case "ArtistDateRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistDateRectAnchorType, valueExpr)
	case "ArtistDateTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistDateTextAnchorType, valueExpr)
	case "ArtistDateDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistDateDominantBaselineType, valueExpr)
	case "ArtistDateAndPlacesFontSize":
		instance.ArtistDateAndPlacesFontSize = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesFontWeigth":
		instance.ArtistDateAndPlacesFontWeigth = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesFontFamily":
		instance.ArtistDateAndPlacesFontFamily = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesLetterSpacing":
		instance.ArtistDateAndPlacesLetterSpacing = GongExtractString(valueExpr)
	case "ArtistPlacesRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistPlacesRectAnchorType, valueExpr)
	case "ArtistPlacesTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistPlacesTextAnchorType, valueExpr)
	case "ArtistPlacesDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistPlacesDominantBaselineType, valueExpr)
	case "InfluenceArrowSize":
		instance.InfluenceArrowSize = GongExtractFloat(valueExpr)
	case "InfluenceArrowStartOffset":
		instance.InfluenceArrowStartOffset = GongExtractFloat(valueExpr)
	case "InfluenceArrowEndOffset":
		instance.InfluenceArrowEndOffset = GongExtractFloat(valueExpr)
	case "InfluenceCornerRadius":
		instance.InfluenceCornerRadius = GongExtractFloat(valueExpr)
	case "InfluenceDashedLinePattern":
		instance.InfluenceDashedLinePattern = GongExtractString(valueExpr)
	}
	return nil
}

type InfluenceUnmarshaller struct{}

func (u *InfluenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Influence), stage, identifier, instanceName, preserveOrder)
}

func (u *InfluenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Influence)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SourceMovement":
		GongUnmarshallPointer(&instance.SourceMovement, valueExpr, identifierMap)
	case "SourceArtefactType":
		GongUnmarshallPointer(&instance.SourceArtefactType, valueExpr, identifierMap)
	case "SourceArtist":
		GongUnmarshallPointer(&instance.SourceArtist, valueExpr, identifierMap)
	case "TargetMovement":
		GongUnmarshallPointer(&instance.TargetMovement, valueExpr, identifierMap)
	case "TargetArtefactType":
		GongUnmarshallPointer(&instance.TargetArtefactType, valueExpr, identifierMap)
	case "TargetArtist":
		GongUnmarshallPointer(&instance.TargetArtist, valueExpr, identifierMap)
	case "IsHypothtical":
		instance.IsHypothtical = GongExtractBool(valueExpr)
	}
	return nil
}

type InfluenceShapeUnmarshaller struct{}

func (u *InfluenceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InfluenceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *InfluenceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InfluenceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Influence":
		GongUnmarshallPointer(&instance.Influence, valueExpr, identifierMap)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type LibraryUnmarshaller struct{}

func (u *LibraryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Library), stage, identifier, instanceName, preserveOrder)
}

func (u *LibraryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Library)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "IsSubLibrariesNodeExpanded":
		instance.IsSubLibrariesNodeExpanded = GongExtractBool(valueExpr)
	case "SubLibrariesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SubLibrariesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "IsExpandedTmp":
		instance.IsExpandedTmp = GongExtractBool(valueExpr)
	}
	return nil
}

type MovementUnmarshaller struct{}

func (u *MovementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Movement), stage, identifier, instanceName, preserveOrder)
}

func (u *MovementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Movement)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	case "HideDate":
		instance.HideDate = GongExtractBool(valueExpr)
	case "Places":
		GongUnmarshallSliceOfPointers(&instance.Places, valueExpr, identifierMap)
	case "HasTaxonomicFilter":
		instance.HasTaxonomicFilter = GongExtractBool(valueExpr)
	case "TaxonomicFilter":
		instance.TaxonomicFilter = GongExtractString(valueExpr)
	case "IsFeatured":
		instance.IsFeatured = GongExtractBool(valueExpr)
	case "FeaturePrefix":
		instance.FeaturePrefix = GongExtractString(valueExpr)
	case "IsMajor":
		instance.IsMajor = GongExtractBool(valueExpr)
	case "IsMinor":
		instance.IsMinor = GongExtractBool(valueExpr)
	case "AdditionnalName":
		instance.AdditionnalName = GongExtractString(valueExpr)
	}
	return nil
}

type MovementShapeUnmarshaller struct{}

func (u *MovementShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MovementShape), stage, identifier, instanceName, preserveOrder)
}

func (u *MovementShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MovementShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Movement":
		GongUnmarshallPointer(&instance.Movement, valueExpr, identifierMap)
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
	}
	return nil
}

type PlaceUnmarshaller struct{}

func (u *PlaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Place), stage, identifier, instanceName, preserveOrder)
}

func (u *PlaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Place)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

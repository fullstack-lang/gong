// generated code - do not edit
package score

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
type CreditUnmarshaller struct{}

func (u *CreditUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Credit), stage, identifier, instanceName, preserveOrder)
}

func (u *CreditUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Credit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Page":
		instance.Page = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "CreditType":
		instance.CreditType = GongExtractString(valueExpr)
	}
	return nil
}

type DefaultsUnmarshaller struct{}

func (u *DefaultsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Defaults), stage, identifier, instanceName, preserveOrder)
}

func (u *DefaultsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Defaults)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MusicFont":
		GongUnmarshallPointer(&instance.MusicFont, valueExpr, identifierMap)
	case "WordFont":
		GongUnmarshallPointer(&instance.WordFont, valueExpr, identifierMap)
	case "LyricFont":
		GongUnmarshallSliceOfPointers(&instance.LyricFont, valueExpr, identifierMap)
	case "LyricLanguage":
		GongUnmarshallSliceOfPointers(&instance.LyricLanguage, valueExpr, identifierMap)
	}
	return nil
}

type EmptyFontUnmarshaller struct{}

func (u *EmptyFontUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyFont), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyFontUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyFont)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	}
	return nil
}

type GroupBarlineUnmarshaller struct{}

func (u *GroupBarlineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroupBarline), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupBarlineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroupBarline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type GroupNameUnmarshaller struct{}

func (u *GroupNameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroupName), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupNameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroupName)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DefaultX":
		instance.DefaultX = GongExtractString(valueExpr)
	case "DefaultY":
		instance.DefaultY = GongExtractString(valueExpr)
	case "RelativeX":
		instance.RelativeX = GongExtractString(valueExpr)
	case "RelativeY":
		instance.RelativeY = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type GroupSymbolUnmarshaller struct{}

func (u *GroupSymbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroupSymbol), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupSymbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroupSymbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DefaultX":
		instance.DefaultX = GongExtractString(valueExpr)
	case "DefaultY":
		instance.DefaultY = GongExtractString(valueExpr)
	case "RelativeX":
		instance.RelativeX = GongExtractString(valueExpr)
	case "RelativeY":
		instance.RelativeY = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type InstrumentLinkUnmarshaller struct{}

func (u *InstrumentLinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InstrumentLink), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InstrumentLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RefScoreInstrument":
		GongUnmarshallPointer(&instance.RefScoreInstrument, valueExpr, identifierMap)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type LyricFontUnmarshaller struct{}

func (u *LyricFontUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LyricFont), stage, identifier, instanceName, preserveOrder)
}

func (u *LyricFontUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LyricFont)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameAttr":
		instance.NameAttr = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	}
	return nil
}

type LyricLanguageUnmarshaller struct{}

func (u *LyricLanguageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LyricLanguage), stage, identifier, instanceName, preserveOrder)
}

func (u *LyricLanguageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LyricLanguage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameAttr":
		instance.NameAttr = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

type MeasureUnmarshaller struct{}

func (u *MeasureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Measure), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "MusicData":
		GongUnmarshallSliceOfPointers(&instance.MusicData, valueExpr, identifierMap)
	}
	return nil
}

type MeasureMusicDataChoiceUnmarshaller struct{}

func (u *MeasureMusicDataChoiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeasureMusicDataChoice), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureMusicDataChoiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeasureMusicDataChoice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type OpusUnmarshaller struct{}

func (u *OpusUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Opus), stage, identifier, instanceName, preserveOrder)
}

func (u *OpusUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Opus)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	}
	return nil
}

type PartUnmarshaller struct{}

func (u *PartUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part), stage, identifier, instanceName, preserveOrder)
}

func (u *PartUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RefScorePart":
		GongUnmarshallPointer(&instance.RefScorePart, valueExpr, identifierMap)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Measure":
		GongUnmarshallSliceOfPointers(&instance.Measure, valueExpr, identifierMap)
	}
	return nil
}

type PartGroupUnmarshaller struct{}

func (u *PartGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartGroup), stage, identifier, instanceName, preserveOrder)
}

func (u *PartGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "GroupName":
		GongUnmarshallPointer(&instance.GroupName, valueExpr, identifierMap)
	case "GroupAbbreviation":
		GongUnmarshallPointer(&instance.GroupAbbreviation, valueExpr, identifierMap)
	case "GroupSymbol":
		GongUnmarshallPointer(&instance.GroupSymbol, valueExpr, identifierMap)
	case "GroupBarline":
		GongUnmarshallPointer(&instance.GroupBarline, valueExpr, identifierMap)
	}
	return nil
}

type PartLinkUnmarshaller struct{}

func (u *PartLinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartLink), stage, identifier, instanceName, preserveOrder)
}

func (u *PartLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "InstrumentLink":
		GongUnmarshallSliceOfPointers(&instance.InstrumentLink, valueExpr, identifierMap)
	case "GroupLink":
		instance.GroupLink = GongExtractString(valueExpr)
	}
	return nil
}

type PartListUnmarshaller struct{}

func (u *PartListUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartList), stage, identifier, instanceName, preserveOrder)
}

func (u *PartListUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartList)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PartGroup":
		GongUnmarshallSliceOfPointers(&instance.PartGroup, valueExpr, identifierMap)
	case "ScorePart":
		GongUnmarshallSliceOfPointers(&instance.ScorePart, valueExpr, identifierMap)
	}
	return nil
}

type PartNameUnmarshaller struct{}

func (u *PartNameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartName), stage, identifier, instanceName, preserveOrder)
}

func (u *PartNameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartName)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DefaultX":
		instance.DefaultX = GongExtractString(valueExpr)
	case "DefaultY":
		instance.DefaultY = GongExtractString(valueExpr)
	case "RelativeX":
		instance.RelativeX = GongExtractString(valueExpr)
	case "RelativeY":
		instance.RelativeY = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PlayerUnmarshaller struct{}

func (u *PlayerUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Player), stage, identifier, instanceName, preserveOrder)
}

func (u *PlayerUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Player)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "PlayerName":
		instance.PlayerName = GongExtractString(valueExpr)
	}
	return nil
}

type ScoreInstrumentUnmarshaller struct{}

func (u *ScoreInstrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ScoreInstrument), stage, identifier, instanceName, preserveOrder)
}

func (u *ScoreInstrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ScoreInstrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "InstrumentName":
		instance.InstrumentName = GongExtractString(valueExpr)
	case "InstrumentAbbreviation":
		instance.InstrumentAbbreviation = GongExtractString(valueExpr)
	case "InstrumentSound":
		instance.InstrumentSound = GongExtractString(valueExpr)
	case "Ensemble":
		instance.Ensemble = GongExtractString(valueExpr)
	}
	return nil
}

type ScorePartUnmarshaller struct{}

func (u *ScorePartUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ScorePart), stage, identifier, instanceName, preserveOrder)
}

func (u *ScorePartUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ScorePart)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "PartLink":
		GongUnmarshallSliceOfPointers(&instance.PartLink, valueExpr, identifierMap)
	case "PartName":
		GongUnmarshallPointer(&instance.PartName, valueExpr, identifierMap)
	case "PartAbbreviation":
		GongUnmarshallPointer(&instance.PartAbbreviation, valueExpr, identifierMap)
	case "Group":
		instance.Group = GongExtractString(valueExpr)
	case "ScoreInstrument":
		GongUnmarshallSliceOfPointers(&instance.ScoreInstrument, valueExpr, identifierMap)
	case "Player":
		GongUnmarshallSliceOfPointers(&instance.Player, valueExpr, identifierMap)
	}
	return nil
}

type ScorePartwiseUnmarshaller struct{}

func (u *ScorePartwiseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ScorePartwise), stage, identifier, instanceName, preserveOrder)
}

func (u *ScorePartwiseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ScorePartwise)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Version":
		instance.Version = GongExtractString(valueExpr)
	case "Work":
		GongUnmarshallPointer(&instance.Work, valueExpr, identifierMap)
	case "MovementNumber":
		instance.MovementNumber = GongExtractString(valueExpr)
	case "MovementTitle":
		instance.MovementTitle = GongExtractString(valueExpr)
	case "Defaults":
		GongUnmarshallPointer(&instance.Defaults, valueExpr, identifierMap)
	case "Credit":
		GongUnmarshallSliceOfPointers(&instance.Credit, valueExpr, identifierMap)
	case "PartList":
		GongUnmarshallPointer(&instance.PartList, valueExpr, identifierMap)
	case "Part":
		GongUnmarshallSliceOfPointers(&instance.Part, valueExpr, identifierMap)
	}
	return nil
}

type ScoreTimewiseUnmarshaller struct{}

func (u *ScoreTimewiseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ScoreTimewise), stage, identifier, instanceName, preserveOrder)
}

func (u *ScoreTimewiseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ScoreTimewise)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Version":
		instance.Version = GongExtractString(valueExpr)
	case "Work":
		GongUnmarshallPointer(&instance.Work, valueExpr, identifierMap)
	case "MovementNumber":
		instance.MovementNumber = GongExtractString(valueExpr)
	case "MovementTitle":
		instance.MovementTitle = GongExtractString(valueExpr)
	case "Defaults":
		GongUnmarshallPointer(&instance.Defaults, valueExpr, identifierMap)
	case "Credit":
		GongUnmarshallSliceOfPointers(&instance.Credit, valueExpr, identifierMap)
	case "PartList":
		GongUnmarshallPointer(&instance.PartList, valueExpr, identifierMap)
	case "Measure":
		GongUnmarshallSliceOfPointers(&instance.Measure, valueExpr, identifierMap)
	}
	return nil
}

type ScoreTimewiseMeasureUnmarshaller struct{}

func (u *ScoreTimewiseMeasureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ScoreTimewiseMeasure), stage, identifier, instanceName, preserveOrder)
}

func (u *ScoreTimewiseMeasureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ScoreTimewiseMeasure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallSliceOfPointers(&instance.Part, valueExpr, identifierMap)
	}
	return nil
}

type ScoreTimewiseMeasurePartUnmarshaller struct{}

func (u *ScoreTimewiseMeasurePartUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ScoreTimewiseMeasurePart), stage, identifier, instanceName, preserveOrder)
}

func (u *ScoreTimewiseMeasurePartUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ScoreTimewiseMeasurePart)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RefScorePart":
		GongUnmarshallPointer(&instance.RefScorePart, valueExpr, identifierMap)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type WorkUnmarshaller struct{}

func (u *WorkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Work), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Work)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "WorkNumber":
		instance.WorkNumber = GongExtractString(valueExpr)
	case "WorkTitle":
		instance.WorkTitle = GongExtractString(valueExpr)
	case "Opus":
		GongUnmarshallPointer(&instance.Opus, valueExpr, identifierMap)
	}
	return nil
}

// generated code - do not edit
package direction

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
type AccordUnmarshaller struct{}

func (u *AccordUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accord), stage, identifier, instanceName, preserveOrder)
}

func (u *AccordUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accord)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "String":
		instance.String = GongExtractInt(valueExpr)
	case "TuningAlter":
		instance.TuningAlter = GongExtractString(valueExpr)
	case "TuningOctave":
		instance.TuningOctave = GongExtractInt(valueExpr)
	}
	return nil
}

type AccordionRegistrationUnmarshaller struct{}

func (u *AccordionRegistrationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AccordionRegistration), stage, identifier, instanceName, preserveOrder)
}

func (u *AccordionRegistrationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AccordionRegistration)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "AccordionMiddle":
		instance.AccordionMiddle = GongExtractInt(valueExpr)
	}
	return nil
}

type BarreUnmarshaller struct{}

func (u *BarreUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Barre), stage, identifier, instanceName, preserveOrder)
}

func (u *BarreUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Barre)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type BassUnmarshaller struct{}

func (u *BassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bass), stage, identifier, instanceName, preserveOrder)
}

func (u *BassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Arrangement":
		GongUnmarshallEnum(&instance.Arrangement, valueExpr)
	case "BassSeparator":
		GongUnmarshallPointer(&instance.BassSeparator, valueExpr, identifierMap)
	case "BassStep":
		GongUnmarshallPointer(&instance.BassStep, valueExpr, identifierMap)
	case "BassAlter":
		GongUnmarshallPointer(&instance.BassAlter, valueExpr, identifierMap)
	}
	return nil
}

type BassStepUnmarshaller struct{}

func (u *BassStepUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BassStep), stage, identifier, instanceName, preserveOrder)
}

func (u *BassStepUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BassStep)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
	}
	return nil
}

type BeatUnitTiedUnmarshaller struct{}

func (u *BeatUnitTiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BeatUnitTied), stage, identifier, instanceName, preserveOrder)
}

func (u *BeatUnitTiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BeatUnitTied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type BeaterUnmarshaller struct{}

func (u *BeaterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Beater), stage, identifier, instanceName, preserveOrder)
}

func (u *BeaterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beater)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tip":
		GongUnmarshallEnum(&instance.Tip, valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type BracketUnmarshaller struct{}

func (u *BracketUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bracket), stage, identifier, instanceName, preserveOrder)
}

func (u *BracketUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bracket)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "LineEnd":
		GongUnmarshallEnum(&instance.LineEnd, valueExpr)
	case "EndLength":
		instance.EndLength = GongExtractString(valueExpr)
	case "DashLength":
		instance.DashLength = GongExtractString(valueExpr)
	case "SpaceLength":
		instance.SpaceLength = GongExtractString(valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type DashesUnmarshaller struct{}

func (u *DashesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Dashes), stage, identifier, instanceName, preserveOrder)
}

func (u *DashesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dashes)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "DashLength":
		instance.DashLength = GongExtractString(valueExpr)
	case "SpaceLength":
		instance.SpaceLength = GongExtractString(valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type DegreeUnmarshaller struct{}

func (u *DegreeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Degree), stage, identifier, instanceName, preserveOrder)
}

func (u *DegreeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DegreeValue":
		GongUnmarshallPointer(&instance.DegreeValue, valueExpr, identifierMap)
	case "DegreeAlter":
		GongUnmarshallPointer(&instance.DegreeAlter, valueExpr, identifierMap)
	case "DegreeType":
		GongUnmarshallPointer(&instance.DegreeType, valueExpr, identifierMap)
	}
	return nil
}

type DegreeAlterUnmarshaller struct{}

func (u *DegreeAlterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DegreeAlter), stage, identifier, instanceName, preserveOrder)
}

func (u *DegreeAlterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DegreeAlter)
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

type DegreeTypeUnmarshaller struct{}

func (u *DegreeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DegreeType), stage, identifier, instanceName, preserveOrder)
}

func (u *DegreeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DegreeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type DegreeValueUnmarshaller struct{}

func (u *DegreeValueUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DegreeValue), stage, identifier, instanceName, preserveOrder)
}

func (u *DegreeValueUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DegreeValue)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type DirectionUnmarshaller struct{}

func (u *DirectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Direction), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Direction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "DirectionType":
		GongUnmarshallSliceOfPointers(&instance.DirectionType, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Sound":
		GongUnmarshallPointer(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallPointer(&instance.Listening, valueExpr, identifierMap)
	}
	return nil
}

type DirectionTypeUnmarshaller struct{}

func (u *DirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Wedge":
		GongUnmarshallPointer(&instance.Wedge, valueExpr, identifierMap)
	case "Dashes":
		GongUnmarshallPointer(&instance.Dashes, valueExpr, identifierMap)
	case "Bracket":
		GongUnmarshallPointer(&instance.Bracket, valueExpr, identifierMap)
	case "Pedal":
		GongUnmarshallPointer(&instance.Pedal, valueExpr, identifierMap)
	case "Metronome":
		GongUnmarshallPointer(&instance.Metronome, valueExpr, identifierMap)
	case "OctaveShift":
		GongUnmarshallPointer(&instance.OctaveShift, valueExpr, identifierMap)
	case "HarpPedals":
		GongUnmarshallPointer(&instance.HarpPedals, valueExpr, identifierMap)
	case "StringMute":
		GongUnmarshallPointer(&instance.StringMute, valueExpr, identifierMap)
	case "Scordatura":
		GongUnmarshallPointer(&instance.Scordatura, valueExpr, identifierMap)
	case "Image":
		GongUnmarshallPointer(&instance.Image, valueExpr, identifierMap)
	case "PrincipalVoice":
		GongUnmarshallPointer(&instance.PrincipalVoice, valueExpr, identifierMap)
	case "Percussion":
		GongUnmarshallSliceOfPointers(&instance.Percussion, valueExpr, identifierMap)
	case "AccordionRegistration":
		GongUnmarshallPointer(&instance.AccordionRegistration, valueExpr, identifierMap)
	case "StaffDivide":
		GongUnmarshallPointer(&instance.StaffDivide, valueExpr, identifierMap)
	case "OtherDirection":
		GongUnmarshallPointer(&instance.OtherDirection, valueExpr, identifierMap)
	}
	return nil
}

type EffectUnmarshaller struct{}

func (u *EffectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Effect), stage, identifier, instanceName, preserveOrder)
}

func (u *EffectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Effect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type FeatureUnmarshaller struct{}

func (u *FeatureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Feature), stage, identifier, instanceName, preserveOrder)
}

func (u *FeatureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Feature)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type FirstFretUnmarshaller struct{}

func (u *FirstFretUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FirstFret), stage, identifier, instanceName, preserveOrder)
}

func (u *FirstFretUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FirstFret)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type FrameUnmarshaller struct{}

func (u *FrameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Frame), stage, identifier, instanceName, preserveOrder)
}

func (u *FrameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Frame)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Height":
		instance.Height = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Unplayed":
		instance.Unplayed = GongExtractString(valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "FrameStrings":
		instance.FrameStrings = GongExtractInt(valueExpr)
	case "FrameFrets":
		instance.FrameFrets = GongExtractInt(valueExpr)
	case "FirstFret":
		GongUnmarshallPointer(&instance.FirstFret, valueExpr, identifierMap)
	case "FrameNote":
		GongUnmarshallSliceOfPointers(&instance.FrameNote, valueExpr, identifierMap)
	}
	return nil
}

type FrameNoteUnmarshaller struct{}

func (u *FrameNoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FrameNote), stage, identifier, instanceName, preserveOrder)
}

func (u *FrameNoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FrameNote)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Barre":
		GongUnmarshallPointer(&instance.Barre, valueExpr, identifierMap)
	}
	return nil
}

type GlassUnmarshaller struct{}

func (u *GlassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Glass), stage, identifier, instanceName, preserveOrder)
}

func (u *GlassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type GroupingUnmarshaller struct{}

func (u *GroupingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Grouping), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Grouping)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "MemberOf":
		instance.MemberOf = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Feature":
		GongUnmarshallSliceOfPointers(&instance.Feature, valueExpr, identifierMap)
	}
	return nil
}

type HarmonyUnmarshaller struct{}

func (u *HarmonyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmony), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmony)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Arrangement":
		GongUnmarshallEnum(&instance.Arrangement, valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Root":
		GongUnmarshallSliceOfPointers(&instance.Root, valueExpr, identifierMap)
	case "Numeral":
		GongUnmarshallSliceOfPointers(&instance.Numeral, valueExpr, identifierMap)
	case "Function":
		GongUnmarshallSliceOfPointers(&instance.Function, valueExpr, identifierMap)
	case "Kind":
		GongUnmarshallSliceOfPointers(&instance.Kind, valueExpr, identifierMap)
	case "Inversion":
		GongUnmarshallSliceOfPointers(&instance.Inversion, valueExpr, identifierMap)
	case "Bass":
		GongUnmarshallSliceOfPointers(&instance.Bass, valueExpr, identifierMap)
	case "Degree":
		GongUnmarshallSliceOfPointers(&instance.Degree, valueExpr, identifierMap)
	case "Frame":
		GongUnmarshallPointer(&instance.Frame, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	}
	return nil
}

type HarmonyAlterUnmarshaller struct{}

func (u *HarmonyAlterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HarmonyAlter), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonyAlterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HarmonyAlter)
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

type HarpPedalsUnmarshaller struct{}

func (u *HarpPedalsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HarpPedals), stage, identifier, instanceName, preserveOrder)
}

func (u *HarpPedalsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HarpPedals)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "PedalTuning":
		GongUnmarshallSliceOfPointers(&instance.PedalTuning, valueExpr, identifierMap)
	}
	return nil
}

type ImageUnmarshaller struct{}

func (u *ImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Image), stage, identifier, instanceName, preserveOrder)
}

func (u *ImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Image)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Height":
		instance.Height = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "DefaultX":
		instance.DefaultX = GongExtractString(valueExpr)
	case "DefaultY":
		instance.DefaultY = GongExtractString(valueExpr)
	case "RelativeX":
		instance.RelativeX = GongExtractString(valueExpr)
	case "RelativeY":
		instance.RelativeY = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type InstrumentChangeUnmarshaller struct{}

func (u *InstrumentChangeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InstrumentChange), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentChangeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InstrumentChange)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "InstrumentSound":
		instance.InstrumentSound = GongExtractString(valueExpr)
	case "Ensemble":
		instance.Ensemble = GongExtractString(valueExpr)
	case "VirtualInstrument":
		GongUnmarshallPointer(&instance.VirtualInstrument, valueExpr, identifierMap)
	}
	return nil
}

type InversionUnmarshaller struct{}

func (u *InversionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Inversion), stage, identifier, instanceName, preserveOrder)
}

func (u *InversionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Inversion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type KindUnmarshaller struct{}

func (u *KindUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Kind), stage, identifier, instanceName, preserveOrder)
}

func (u *KindUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Kind)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type ListeningUnmarshaller struct{}

func (u *ListeningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Listening), stage, identifier, instanceName, preserveOrder)
}

func (u *ListeningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Listening)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sync":
		GongUnmarshallSliceOfPointers(&instance.Sync, valueExpr, identifierMap)
	case "OtherListening":
		GongUnmarshallSliceOfPointers(&instance.OtherListening, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	}
	return nil
}

type MeasureNumberingUnmarshaller struct{}

func (u *MeasureNumberingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeasureNumbering), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureNumberingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeasureNumbering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
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
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type MembraneUnmarshaller struct{}

func (u *MembraneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Membrane), stage, identifier, instanceName, preserveOrder)
}

func (u *MembraneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Membrane)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type MetalUnmarshaller struct{}

func (u *MetalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metal), stage, identifier, instanceName, preserveOrder)
}

func (u *MetalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metal)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type MetronomeUnmarshaller struct{}

func (u *MetronomeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metronome), stage, identifier, instanceName, preserveOrder)
}

func (u *MetronomeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "BeatUnitTied":
		GongUnmarshallSliceOfPointers(&instance.BeatUnitTied, valueExpr, identifierMap)
	case "PerMinute":
		GongUnmarshallPointer(&instance.PerMinute, valueExpr, identifierMap)
	case "MetronomeNote":
		GongUnmarshallSliceOfPointers(&instance.MetronomeNote, valueExpr, identifierMap)
	case "MetronomeRelation":
		instance.MetronomeRelation = GongExtractString(valueExpr)
	}
	return nil
}

type MetronomeBeamUnmarshaller struct{}

func (u *MetronomeBeamUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MetronomeBeam), stage, identifier, instanceName, preserveOrder)
}

func (u *MetronomeBeamUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MetronomeBeam)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type MetronomeNoteUnmarshaller struct{}

func (u *MetronomeNoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MetronomeNote), stage, identifier, instanceName, preserveOrder)
}

func (u *MetronomeNoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MetronomeNote)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MetronomeBeam":
		GongUnmarshallSliceOfPointers(&instance.MetronomeBeam, valueExpr, identifierMap)
	case "MetronomeTied":
		GongUnmarshallPointer(&instance.MetronomeTied, valueExpr, identifierMap)
	case "MetronomeTuplet":
		GongUnmarshallPointer(&instance.MetronomeTuplet, valueExpr, identifierMap)
	}
	return nil
}

type MetronomeTiedUnmarshaller struct{}

func (u *MetronomeTiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MetronomeTied), stage, identifier, instanceName, preserveOrder)
}

func (u *MetronomeTiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MetronomeTied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type MetronomeTupletUnmarshaller struct{}

func (u *MetronomeTupletUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MetronomeTuplet), stage, identifier, instanceName, preserveOrder)
}

func (u *MetronomeTupletUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MetronomeTuplet)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowNumber":
		GongUnmarshallEnum(&instance.ShowNumber, valueExpr)
	case "ActualNotes":
		instance.ActualNotes = GongExtractInt(valueExpr)
	case "NormalNotes":
		instance.NormalNotes = GongExtractInt(valueExpr)
	}
	return nil
}

type NumeralUnmarshaller struct{}

func (u *NumeralUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Numeral), stage, identifier, instanceName, preserveOrder)
}

func (u *NumeralUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumeralRoot":
		GongUnmarshallPointer(&instance.NumeralRoot, valueExpr, identifierMap)
	case "NumeralAlter":
		GongUnmarshallPointer(&instance.NumeralAlter, valueExpr, identifierMap)
	case "NumeralKey":
		GongUnmarshallPointer(&instance.NumeralKey, valueExpr, identifierMap)
	}
	return nil
}

type NumeralKeyUnmarshaller struct{}

func (u *NumeralKeyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NumeralKey), stage, identifier, instanceName, preserveOrder)
}

func (u *NumeralKeyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NumeralKey)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumeralFifths":
		instance.NumeralFifths = GongExtractInt(valueExpr)
	case "NumeralMode":
		GongUnmarshallEnum(&instance.NumeralMode, valueExpr)
	}
	return nil
}

type NumeralRootUnmarshaller struct{}

func (u *NumeralRootUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NumeralRoot), stage, identifier, instanceName, preserveOrder)
}

func (u *NumeralRootUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NumeralRoot)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type OctaveShiftUnmarshaller struct{}

func (u *OctaveShiftUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OctaveShift), stage, identifier, instanceName, preserveOrder)
}

func (u *OctaveShiftUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OctaveShift)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Size":
		instance.Size = GongExtractInt(valueExpr)
	case "DashLength":
		instance.DashLength = GongExtractString(valueExpr)
	case "SpaceLength":
		instance.SpaceLength = GongExtractString(valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type OffsetUnmarshaller struct{}

func (u *OffsetUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Offset), stage, identifier, instanceName, preserveOrder)
}

func (u *OffsetUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Offset)
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

type OtherDirectionUnmarshaller struct{}

func (u *OtherDirectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherDirection), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherDirectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherDirection)
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
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type OtherListeningUnmarshaller struct{}

func (u *OtherListeningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherListening), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherListeningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherListening)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PedalUnmarshaller struct{}

func (u *PedalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pedal), stage, identifier, instanceName, preserveOrder)
}

func (u *PedalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pedal)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type PedalTuningUnmarshaller struct{}

func (u *PedalTuningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PedalTuning), stage, identifier, instanceName, preserveOrder)
}

func (u *PedalTuningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PedalTuning)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PedalAlter":
		instance.PedalAlter = GongExtractString(valueExpr)
	}
	return nil
}

type PerMinuteUnmarshaller struct{}

func (u *PerMinuteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PerMinute), stage, identifier, instanceName, preserveOrder)
}

func (u *PerMinuteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerMinute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PercussionUnmarshaller struct{}

func (u *PercussionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Percussion), stage, identifier, instanceName, preserveOrder)
}

func (u *PercussionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Percussion)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Glass":
		GongUnmarshallPointer(&instance.Glass, valueExpr, identifierMap)
	case "Metal":
		GongUnmarshallPointer(&instance.Metal, valueExpr, identifierMap)
	case "Wood":
		GongUnmarshallPointer(&instance.Wood, valueExpr, identifierMap)
	case "Pitched":
		GongUnmarshallPointer(&instance.Pitched, valueExpr, identifierMap)
	case "Membrane":
		GongUnmarshallPointer(&instance.Membrane, valueExpr, identifierMap)
	case "Effect":
		GongUnmarshallPointer(&instance.Effect, valueExpr, identifierMap)
	case "Timpani":
		GongUnmarshallPointer(&instance.Timpani, valueExpr, identifierMap)
	case "Beater":
		GongUnmarshallPointer(&instance.Beater, valueExpr, identifierMap)
	case "Stick":
		GongUnmarshallPointer(&instance.Stick, valueExpr, identifierMap)
	case "StickLocation":
		GongUnmarshallEnum(&instance.StickLocation, valueExpr)
	}
	return nil
}

type PitchedUnmarshaller struct{}

func (u *PitchedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pitched), stage, identifier, instanceName, preserveOrder)
}

func (u *PitchedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pitched)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type PrincipalVoiceUnmarshaller struct{}

func (u *PrincipalVoiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrincipalVoice), stage, identifier, instanceName, preserveOrder)
}

func (u *PrincipalVoiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrincipalVoice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PrintUnmarshaller struct{}

func (u *PrintUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Print), stage, identifier, instanceName, preserveOrder)
}

func (u *PrintUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Print)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StaffSpacing":
		instance.StaffSpacing = GongExtractString(valueExpr)
	case "BlankPage":
		instance.BlankPage = GongExtractInt(valueExpr)
	case "PageNumber":
		instance.PageNumber = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "MeasureNumbering":
		GongUnmarshallPointer(&instance.MeasureNumbering, valueExpr, identifierMap)
	}
	return nil
}

type RootUnmarshaller struct{}

func (u *RootUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Root), stage, identifier, instanceName, preserveOrder)
}

func (u *RootUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Root)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RootStep":
		GongUnmarshallPointer(&instance.RootStep, valueExpr, identifierMap)
	case "RootAlter":
		GongUnmarshallPointer(&instance.RootAlter, valueExpr, identifierMap)
	}
	return nil
}

type RootStepUnmarshaller struct{}

func (u *RootStepUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RootStep), stage, identifier, instanceName, preserveOrder)
}

func (u *RootStepUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RootStep)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
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
	}
	return nil
}

type ScordaturaUnmarshaller struct{}

func (u *ScordaturaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Scordatura), stage, identifier, instanceName, preserveOrder)
}

func (u *ScordaturaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scordatura)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accord":
		GongUnmarshallSliceOfPointers(&instance.Accord, valueExpr, identifierMap)
	}
	return nil
}

type SoundUnmarshaller struct{}

func (u *SoundUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Sound), stage, identifier, instanceName, preserveOrder)
}

func (u *SoundUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sound)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tempo":
		instance.Tempo = GongExtractString(valueExpr)
	case "Dynamics":
		instance.Dynamics = GongExtractString(valueExpr)
	case "Segno":
		instance.Segno = GongExtractString(valueExpr)
	case "Dalsegno":
		instance.Dalsegno = GongExtractString(valueExpr)
	case "Coda":
		instance.Coda = GongExtractString(valueExpr)
	case "Tocoda":
		instance.Tocoda = GongExtractString(valueExpr)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Fine":
		instance.Fine = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
	case "Pan":
		instance.Pan = GongExtractString(valueExpr)
	case "Elevation":
		instance.Elevation = GongExtractString(valueExpr)
	case "DamperPedal":
		instance.DamperPedal = GongExtractString(valueExpr)
	case "SoftPedal":
		instance.SoftPedal = GongExtractString(valueExpr)
	case "SostenutoPedal":
		instance.SostenutoPedal = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "InstrumentChange":
		GongUnmarshallSliceOfPointers(&instance.InstrumentChange, valueExpr, identifierMap)
	case "Swing":
		GongUnmarshallPointer(&instance.Swing, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	}
	return nil
}

type StaffDivideUnmarshaller struct{}

func (u *StaffDivideUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaffDivide), stage, identifier, instanceName, preserveOrder)
}

func (u *StaffDivideUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaffDivide)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type StickUnmarshaller struct{}

func (u *StickUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stick), stage, identifier, instanceName, preserveOrder)
}

func (u *StickUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stick)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tip":
		GongUnmarshallEnum(&instance.Tip, valueExpr)
	case "StickType":
		GongUnmarshallEnum(&instance.StickType, valueExpr)
	case "StickMaterial":
		GongUnmarshallEnum(&instance.StickMaterial, valueExpr)
	}
	return nil
}

type StringMuteUnmarshaller struct{}

func (u *StringMuteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StringMute), stage, identifier, instanceName, preserveOrder)
}

func (u *StringMuteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StringMute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type StyleTextUnmarshaller struct{}

func (u *StyleTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StyleText), stage, identifier, instanceName, preserveOrder)
}

func (u *StyleTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StyleText)
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

type SwingUnmarshaller struct{}

func (u *SwingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Swing), stage, identifier, instanceName, preserveOrder)
}

func (u *SwingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Swing)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "First":
		instance.First = GongExtractInt(valueExpr)
	case "Second":
		instance.Second = GongExtractInt(valueExpr)
	case "SwingStyle":
		instance.SwingStyle = GongExtractString(valueExpr)
	}
	return nil
}

type SyncUnmarshaller struct{}

func (u *SyncUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Sync), stage, identifier, instanceName, preserveOrder)
}

func (u *SyncUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sync)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Latency":
		instance.Latency = GongExtractInt(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
	}
	return nil
}

type TimeModificationUnmarshaller struct{}

func (u *TimeModificationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimeModification), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeModificationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimeModification)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ActualNotes":
		instance.ActualNotes = GongExtractInt(valueExpr)
	case "NormalNotes":
		instance.NormalNotes = GongExtractInt(valueExpr)
	}
	return nil
}

type TimpaniUnmarshaller struct{}

func (u *TimpaniUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Timpani), stage, identifier, instanceName, preserveOrder)
}

func (u *TimpaniUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Timpani)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	}
	return nil
}

type VirtualInstrumentUnmarshaller struct{}

func (u *VirtualInstrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VirtualInstrument), stage, identifier, instanceName, preserveOrder)
}

func (u *VirtualInstrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VirtualInstrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "VirtualLibrary":
		instance.VirtualLibrary = GongExtractString(valueExpr)
	case "VirtualName":
		instance.VirtualName = GongExtractString(valueExpr)
	}
	return nil
}

type WedgeUnmarshaller struct{}

func (u *WedgeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wedge), stage, identifier, instanceName, preserveOrder)
}

func (u *WedgeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wedge)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Spread":
		instance.Spread = GongExtractString(valueExpr)
	case "DashLength":
		instance.DashLength = GongExtractString(valueExpr)
	case "SpaceLength":
		instance.SpaceLength = GongExtractString(valueExpr)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type WoodUnmarshaller struct{}

func (u *WoodUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wood), stage, identifier, instanceName, preserveOrder)
}

func (u *WoodUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wood)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

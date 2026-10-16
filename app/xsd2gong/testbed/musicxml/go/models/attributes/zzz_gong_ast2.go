// generated code - do not edit
package attributes

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
type AttributesUnmarshaller struct{}

func (u *AttributesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Attributes), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Attributes)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Key":
		GongUnmarshallSliceOfPointers(&instance.Key, valueExpr, identifierMap)
	case "Time":
		GongUnmarshallSliceOfPointers(&instance.Time, valueExpr, identifierMap)
	case "Staves":
		instance.Staves = GongExtractInt(valueExpr)
	case "PartSymbol":
		GongUnmarshallPointer(&instance.PartSymbol, valueExpr, identifierMap)
	case "Instruments":
		instance.Instruments = GongExtractInt(valueExpr)
	case "Clef":
		GongUnmarshallSliceOfPointers(&instance.Clef, valueExpr, identifierMap)
	case "StaffDetails":
		GongUnmarshallSliceOfPointers(&instance.StaffDetails, valueExpr, identifierMap)
	case "Transpose":
		GongUnmarshallSliceOfPointers(&instance.Transpose, valueExpr, identifierMap)
	case "ForPart":
		GongUnmarshallSliceOfPointers(&instance.ForPart, valueExpr, identifierMap)
	case "Directive":
		GongUnmarshallSliceOfPointers(&instance.Directive, valueExpr, identifierMap)
	case "MeasureStyle":
		GongUnmarshallSliceOfPointers(&instance.MeasureStyle, valueExpr, identifierMap)
	}
	return nil
}

type AttributesDirectiveUnmarshaller struct{}

func (u *AttributesDirectiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AttributesDirective), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributesDirectiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributesDirective)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
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

type BeatRepeatUnmarshaller struct{}

func (u *BeatRepeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BeatRepeat), stage, identifier, instanceName, preserveOrder)
}

func (u *BeatRepeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BeatRepeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Slashes":
		instance.Slashes = GongExtractInt(valueExpr)
	case "ExceptVoice":
		instance.ExceptVoice = GongExtractString(valueExpr)
	}
	return nil
}

type CancelUnmarshaller struct{}

func (u *CancelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Cancel), stage, identifier, instanceName, preserveOrder)
}

func (u *CancelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Cancel)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type ClefUnmarshaller struct{}

func (u *ClefUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Clef), stage, identifier, instanceName, preserveOrder)
}

func (u *ClefUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Clef)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "Sign":
		GongUnmarshallEnum(&instance.Sign, valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "ClefOctaveChange":
		instance.ClefOctaveChange = GongExtractInt(valueExpr)
	}
	return nil
}

type DoubleUnmarshaller struct{}

func (u *DoubleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Double), stage, identifier, instanceName, preserveOrder)
}

func (u *DoubleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Double)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ForPartUnmarshaller struct{}

func (u *ForPartUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ForPart), stage, identifier, instanceName, preserveOrder)
}

func (u *ForPartUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ForPart)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "PartClef":
		GongUnmarshallPointer(&instance.PartClef, valueExpr, identifierMap)
	case "PartTranspose":
		GongUnmarshallPointer(&instance.PartTranspose, valueExpr, identifierMap)
	}
	return nil
}

type InterchangeableUnmarshaller struct{}

func (u *InterchangeableUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Interchangeable), stage, identifier, instanceName, preserveOrder)
}

func (u *InterchangeableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Interchangeable)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Separator":
		GongUnmarshallEnum(&instance.Separator, valueExpr)
	case "TimeRelation":
		GongUnmarshallEnum(&instance.TimeRelation, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "BeatType":
		instance.BeatType = GongExtractString(valueExpr)
	}
	return nil
}

type KeyUnmarshaller struct{}

func (u *KeyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Key), stage, identifier, instanceName, preserveOrder)
}

func (u *KeyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "Cancel":
		GongUnmarshallPointer(&instance.Cancel, valueExpr, identifierMap)
	case "Fifths":
		instance.Fifths = GongExtractInt(valueExpr)
	case "Mode":
		instance.Mode = GongExtractString(valueExpr)
	case "KeyStep":
		GongUnmarshallEnum(&instance.KeyStep, valueExpr)
	case "KeyAlter":
		instance.KeyAlter = GongExtractString(valueExpr)
	case "KeyAccidental":
		GongUnmarshallSliceOfPointers(&instance.KeyAccidental, valueExpr, identifierMap)
	case "KeyOctave":
		GongUnmarshallSliceOfPointers(&instance.KeyOctave, valueExpr, identifierMap)
	}
	return nil
}

type KeyAccidentalUnmarshaller struct{}

func (u *KeyAccidentalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(KeyAccidental), stage, identifier, instanceName, preserveOrder)
}

func (u *KeyAccidentalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*KeyAccidental)
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

type KeyOctaveUnmarshaller struct{}

func (u *KeyOctaveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(KeyOctave), stage, identifier, instanceName, preserveOrder)
}

func (u *KeyOctaveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*KeyOctave)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type LineDetailUnmarshaller struct{}

func (u *LineDetailUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LineDetail), stage, identifier, instanceName, preserveOrder)
}

func (u *LineDetailUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LineDetail)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type MeasureRepeatUnmarshaller struct{}

func (u *MeasureRepeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeasureRepeat), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureRepeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeasureRepeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Slashes":
		instance.Slashes = GongExtractInt(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MeasureStyleUnmarshaller struct{}

func (u *MeasureStyleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeasureStyle), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureStyleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeasureStyle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "MultipleRest":
		GongUnmarshallPointer(&instance.MultipleRest, valueExpr, identifierMap)
	case "MeasureRepeat":
		GongUnmarshallPointer(&instance.MeasureRepeat, valueExpr, identifierMap)
	case "BeatRepeat":
		GongUnmarshallPointer(&instance.BeatRepeat, valueExpr, identifierMap)
	case "Slash":
		GongUnmarshallPointer(&instance.Slash, valueExpr, identifierMap)
	}
	return nil
}

type MultipleRestUnmarshaller struct{}

func (u *MultipleRestUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MultipleRest), stage, identifier, instanceName, preserveOrder)
}

func (u *MultipleRestUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MultipleRest)
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

type PartClefUnmarshaller struct{}

func (u *PartClefUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartClef), stage, identifier, instanceName, preserveOrder)
}

func (u *PartClefUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartClef)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sign":
		GongUnmarshallEnum(&instance.Sign, valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "ClefOctaveChange":
		instance.ClefOctaveChange = GongExtractInt(valueExpr)
	}
	return nil
}

type PartSymbolUnmarshaller struct{}

func (u *PartSymbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartSymbol), stage, identifier, instanceName, preserveOrder)
}

func (u *PartSymbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartSymbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStaff":
		instance.TopStaff = GongExtractInt(valueExpr)
	case "BottomStaff":
		instance.BottomStaff = GongExtractInt(valueExpr)
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
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type PartTransposeUnmarshaller struct{}

func (u *PartTransposeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartTranspose), stage, identifier, instanceName, preserveOrder)
}

func (u *PartTransposeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartTranspose)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Diatonic":
		instance.Diatonic = GongExtractInt(valueExpr)
	case "Chromatic":
		instance.Chromatic = GongExtractString(valueExpr)
	case "OctaveChange":
		instance.OctaveChange = GongExtractInt(valueExpr)
	case "Double":
		GongUnmarshallPointer(&instance.Double, valueExpr, identifierMap)
	}
	return nil
}

type SlashUnmarshaller struct{}

func (u *SlashUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slash), stage, identifier, instanceName, preserveOrder)
}

func (u *SlashUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slash)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ExceptVoice":
		instance.ExceptVoice = GongExtractString(valueExpr)
	}
	return nil
}

type StaffDetailsUnmarshaller struct{}

func (u *StaffDetailsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaffDetails), stage, identifier, instanceName, preserveOrder)
}

func (u *StaffDetailsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaffDetails)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "ShowFrets":
		GongUnmarshallEnum(&instance.ShowFrets, valueExpr)
	case "StaffType":
		GongUnmarshallEnum(&instance.StaffType, valueExpr)
	case "StaffLines":
		instance.StaffLines = GongExtractInt(valueExpr)
	case "LineDetail":
		GongUnmarshallSliceOfPointers(&instance.LineDetail, valueExpr, identifierMap)
	case "StaffTuning":
		GongUnmarshallSliceOfPointers(&instance.StaffTuning, valueExpr, identifierMap)
	case "Capo":
		instance.Capo = GongExtractInt(valueExpr)
	case "StaffSize":
		GongUnmarshallPointer(&instance.StaffSize, valueExpr, identifierMap)
	}
	return nil
}

type StaffSizeUnmarshaller struct{}

func (u *StaffSizeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaffSize), stage, identifier, instanceName, preserveOrder)
}

func (u *StaffSizeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaffSize)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Scaling":
		instance.Scaling = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type StaffTuningUnmarshaller struct{}

func (u *StaffTuningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaffTuning), stage, identifier, instanceName, preserveOrder)
}

func (u *StaffTuningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaffTuning)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "TuningStep":
		GongUnmarshallEnum(&instance.TuningStep, valueExpr)
	case "TuningAlter":
		instance.TuningAlter = GongExtractString(valueExpr)
	case "TuningOctave":
		instance.TuningOctave = GongExtractInt(valueExpr)
	}
	return nil
}

type TimeUnmarshaller struct{}

func (u *TimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Time), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Time)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Separator":
		GongUnmarshallEnum(&instance.Separator, valueExpr)
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
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "BeatType":
		instance.BeatType = GongExtractString(valueExpr)
	case "Interchangeable":
		GongUnmarshallPointer(&instance.Interchangeable, valueExpr, identifierMap)
	case "SenzaMisura":
		instance.SenzaMisura = GongExtractString(valueExpr)
	}
	return nil
}

type TransposeUnmarshaller struct{}

func (u *TransposeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Transpose), stage, identifier, instanceName, preserveOrder)
}

func (u *TransposeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Transpose)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Diatonic":
		instance.Diatonic = GongExtractInt(valueExpr)
	case "Chromatic":
		instance.Chromatic = GongExtractString(valueExpr)
	case "OctaveChange":
		instance.OctaveChange = GongExtractInt(valueExpr)
	case "Double":
		GongUnmarshallPointer(&instance.Double, valueExpr, identifierMap)
	}
	return nil
}

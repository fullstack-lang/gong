// generated code - do not edit
package note

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
type AccidentalUnmarshaller struct{}

func (u *AccidentalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accidental), stage, identifier, instanceName, preserveOrder)
}

func (u *AccidentalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
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

type AccidentalMarkUnmarshaller struct{}

func (u *AccidentalMarkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AccidentalMark), stage, identifier, instanceName, preserveOrder)
}

func (u *AccidentalMarkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AccidentalMark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
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

type ArpeggiateUnmarshaller struct{}

func (u *ArpeggiateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Arpeggiate), stage, identifier, instanceName, preserveOrder)
}

func (u *ArpeggiateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Arpeggiate)
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
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

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
	case "ArrowDirection":
		GongUnmarshallEnum(&instance.ArrowDirection, valueExpr)
	case "ArrowStyle":
		GongUnmarshallEnum(&instance.ArrowStyle, valueExpr)
	case "CircularArrow":
		GongUnmarshallEnum(&instance.CircularArrow, valueExpr)
	}
	return nil
}

type ArticulationsUnmarshaller struct{}

func (u *ArticulationsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Articulations), stage, identifier, instanceName, preserveOrder)
}

func (u *ArticulationsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Articulations)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "StrongAccent":
		GongUnmarshallSliceOfPointers(&instance.StrongAccent, valueExpr, identifierMap)
	case "Scoop":
		GongUnmarshallSliceOfPointers(&instance.Scoop, valueExpr, identifierMap)
	case "Plop":
		GongUnmarshallSliceOfPointers(&instance.Plop, valueExpr, identifierMap)
	case "Doit":
		GongUnmarshallSliceOfPointers(&instance.Doit, valueExpr, identifierMap)
	case "Falloff":
		GongUnmarshallSliceOfPointers(&instance.Falloff, valueExpr, identifierMap)
	case "BreathMark":
		GongUnmarshallSliceOfPointers(&instance.BreathMark, valueExpr, identifierMap)
	case "Caesura":
		GongUnmarshallSliceOfPointers(&instance.Caesura, valueExpr, identifierMap)
	case "OtherArticulation":
		GongUnmarshallSliceOfPointers(&instance.OtherArticulation, valueExpr, identifierMap)
	}
	return nil
}

type AssessUnmarshaller struct{}

func (u *AssessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Assess), stage, identifier, instanceName, preserveOrder)
}

func (u *AssessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Assess)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
	}
	return nil
}

type BackupUnmarshaller struct{}

func (u *BackupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Backup), stage, identifier, instanceName, preserveOrder)
}

func (u *BackupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Backup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	}
	return nil
}

type BeamUnmarshaller struct{}

func (u *BeamUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Beam), stage, identifier, instanceName, preserveOrder)
}

func (u *BeamUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beam)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Fan":
		GongUnmarshallEnum(&instance.Fan, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type BendUnmarshaller struct{}

func (u *BendUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bend), stage, identifier, instanceName, preserveOrder)
}

func (u *BendUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bend)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Shape":
		GongUnmarshallEnum(&instance.Shape, valueExpr)
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
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "FirstBeat":
		instance.FirstBeat = GongExtractString(valueExpr)
	case "LastBeat":
		instance.LastBeat = GongExtractString(valueExpr)
	case "BendAlter":
		instance.BendAlter = GongExtractString(valueExpr)
	case "Release":
		GongUnmarshallPointer(&instance.Release, valueExpr, identifierMap)
	case "WithBar":
		GongUnmarshallPointer(&instance.WithBar, valueExpr, identifierMap)
	}
	return nil
}

type BreathMarkUnmarshaller struct{}

func (u *BreathMarkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BreathMark), stage, identifier, instanceName, preserveOrder)
}

func (u *BreathMarkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BreathMark)
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
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type CaesuraUnmarshaller struct{}

func (u *CaesuraUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Caesura), stage, identifier, instanceName, preserveOrder)
}

func (u *CaesuraUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Caesura)
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
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type ElisionUnmarshaller struct{}

func (u *ElisionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Elision), stage, identifier, instanceName, preserveOrder)
}

func (u *ElisionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Elision)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
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

type EmptyLineUnmarshaller struct{}

func (u *EmptyLineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyLine), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyLineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyLine)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	}
	return nil
}

type ExtendUnmarshaller struct{}

func (u *ExtendUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Extend), stage, identifier, instanceName, preserveOrder)
}

func (u *ExtendUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extend)
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

type FigureUnmarshaller struct{}

func (u *FigureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Figure), stage, identifier, instanceName, preserveOrder)
}

func (u *FigureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Figure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Extend":
		GongUnmarshallPointer(&instance.Extend, valueExpr, identifierMap)
	}
	return nil
}

type FiguredBassUnmarshaller struct{}

func (u *FiguredBassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FiguredBass), stage, identifier, instanceName, preserveOrder)
}

func (u *FiguredBassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FiguredBass)
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
	case "Figure":
		GongUnmarshallSliceOfPointers(&instance.Figure, valueExpr, identifierMap)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	}
	return nil
}

type ForwardUnmarshaller struct{}

func (u *ForwardUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Forward), stage, identifier, instanceName, preserveOrder)
}

func (u *ForwardUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Forward)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	}
	return nil
}

type GlissandoUnmarshaller struct{}

func (u *GlissandoUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Glissando), stage, identifier, instanceName, preserveOrder)
}

func (u *GlissandoUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glissando)
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

type GraceUnmarshaller struct{}

func (u *GraceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Grace), stage, identifier, instanceName, preserveOrder)
}

func (u *GraceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Grace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StealTimePrevious":
		instance.StealTimePrevious = GongExtractString(valueExpr)
	case "StealTimeFollowing":
		instance.StealTimeFollowing = GongExtractString(valueExpr)
	case "MakeTime":
		instance.MakeTime = GongExtractString(valueExpr)
	}
	return nil
}

type HammerOnPullOffUnmarshaller struct{}

func (u *HammerOnPullOffUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HammerOnPullOff), stage, identifier, instanceName, preserveOrder)
}

func (u *HammerOnPullOffUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HammerOnPullOff)
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
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type HandbellUnmarshaller struct{}

func (u *HandbellUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Handbell), stage, identifier, instanceName, preserveOrder)
}

func (u *HandbellUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Handbell)
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
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type HarmonClosedUnmarshaller struct{}

func (u *HarmonClosedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HarmonClosed), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonClosedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HarmonClosed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type HarmonMuteUnmarshaller struct{}

func (u *HarmonMuteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HarmonMute), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonMuteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HarmonMute)
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
	case "HarmonClosed":
		GongUnmarshallPointer(&instance.HarmonClosed, valueExpr, identifierMap)
	}
	return nil
}

type HarmonicUnmarshaller struct{}

func (u *HarmonicUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmonic), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonicUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmonic)
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
	}
	return nil
}

type HeelToeUnmarshaller struct{}

func (u *HeelToeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HeelToe), stage, identifier, instanceName, preserveOrder)
}

func (u *HeelToeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HeelToe)
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
	}
	return nil
}

type HoleUnmarshaller struct{}

func (u *HoleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Hole), stage, identifier, instanceName, preserveOrder)
}

func (u *HoleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hole)
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
	case "HoleType":
		instance.HoleType = GongExtractString(valueExpr)
	case "HoleClosed":
		GongUnmarshallPointer(&instance.HoleClosed, valueExpr, identifierMap)
	case "HoleShape":
		instance.HoleShape = GongExtractString(valueExpr)
	}
	return nil
}

type HoleClosedUnmarshaller struct{}

func (u *HoleClosedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoleClosed), stage, identifier, instanceName, preserveOrder)
}

func (u *HoleClosedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoleClosed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type InstrumentUnmarshaller struct{}

func (u *InstrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Instrument), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type ListenUnmarshaller struct{}

func (u *ListenUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Listen), stage, identifier, instanceName, preserveOrder)
}

func (u *ListenUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Listen)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Assess":
		GongUnmarshallSliceOfPointers(&instance.Assess, valueExpr, identifierMap)
	case "Wait":
		GongUnmarshallSliceOfPointers(&instance.Wait, valueExpr, identifierMap)
	}
	return nil
}

type LyricUnmarshaller struct{}

func (u *LyricUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Lyric), stage, identifier, instanceName, preserveOrder)
}

func (u *LyricUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameAttr":
		instance.NameAttr = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
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
	case "Syllabic":
		GongUnmarshallEnum(&instance.Syllabic, valueExpr)
	case "Text":
		GongUnmarshallPointer(&instance.Text, valueExpr, identifierMap)
	case "Elision":
		GongUnmarshallSliceOfPointers(&instance.Elision, valueExpr, identifierMap)
	case "Extend":
		GongUnmarshallPointer(&instance.Extend, valueExpr, identifierMap)
	}
	return nil
}

type MordentUnmarshaller struct{}

func (u *MordentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Mordent), stage, identifier, instanceName, preserveOrder)
}

func (u *MordentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Mordent)
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
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "SecondBeat":
		instance.SecondBeat = GongExtractString(valueExpr)
	case "LastBeat":
		instance.LastBeat = GongExtractString(valueExpr)
	}
	return nil
}

type NonArpeggiateUnmarshaller struct{}

func (u *NonArpeggiateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NonArpeggiate), stage, identifier, instanceName, preserveOrder)
}

func (u *NonArpeggiateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NonArpeggiate)
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
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type NotationsUnmarshaller struct{}

func (u *NotationsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Notations), stage, identifier, instanceName, preserveOrder)
}

func (u *NotationsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notations)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Tied":
		GongUnmarshallSliceOfPointers(&instance.Tied, valueExpr, identifierMap)
	case "Slur":
		GongUnmarshallSliceOfPointers(&instance.Slur, valueExpr, identifierMap)
	case "Tuplet":
		GongUnmarshallSliceOfPointers(&instance.Tuplet, valueExpr, identifierMap)
	case "Glissando":
		GongUnmarshallSliceOfPointers(&instance.Glissando, valueExpr, identifierMap)
	case "Slide":
		GongUnmarshallSliceOfPointers(&instance.Slide, valueExpr, identifierMap)
	case "Ornaments":
		GongUnmarshallSliceOfPointers(&instance.Ornaments, valueExpr, identifierMap)
	case "Technical":
		GongUnmarshallSliceOfPointers(&instance.Technical, valueExpr, identifierMap)
	case "Articulations":
		GongUnmarshallSliceOfPointers(&instance.Articulations, valueExpr, identifierMap)
	case "Arpeggiate":
		GongUnmarshallSliceOfPointers(&instance.Arpeggiate, valueExpr, identifierMap)
	case "NonArpeggiate":
		GongUnmarshallSliceOfPointers(&instance.NonArpeggiate, valueExpr, identifierMap)
	case "AccidentalMark":
		GongUnmarshallSliceOfPointers(&instance.AccidentalMark, valueExpr, identifierMap)
	case "OtherNotation":
		GongUnmarshallSliceOfPointers(&instance.OtherNotation, valueExpr, identifierMap)
	}
	return nil
}

type NoteUnmarshaller struct{}

func (u *NoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Note), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Dynamics":
		instance.Dynamics = GongExtractString(valueExpr)
	case "EndDynamics":
		instance.EndDynamics = GongExtractString(valueExpr)
	case "Attack":
		instance.Attack = GongExtractString(valueExpr)
	case "Release":
		instance.Release = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
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
	case "Grace":
		GongUnmarshallPointer(&instance.Grace, valueExpr, identifierMap)
	case "Pitch":
		GongUnmarshallPointer(&instance.Pitch, valueExpr, identifierMap)
	case "Unpitched":
		GongUnmarshallPointer(&instance.Unpitched, valueExpr, identifierMap)
	case "Rest":
		GongUnmarshallPointer(&instance.Rest, valueExpr, identifierMap)
	case "Tie":
		GongUnmarshallSliceOfPointers(&instance.Tie, valueExpr, identifierMap)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	case "Instrument":
		GongUnmarshallSliceOfPointers(&instance.Instrument, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallPointer(&instance.Type, valueExpr, identifierMap)
	case "Accidental":
		GongUnmarshallPointer(&instance.Accidental, valueExpr, identifierMap)
	case "Stem":
		GongUnmarshallPointer(&instance.Stem, valueExpr, identifierMap)
	case "Notehead":
		GongUnmarshallPointer(&instance.Notehead, valueExpr, identifierMap)
	case "NoteheadText":
		GongUnmarshallPointer(&instance.NoteheadText, valueExpr, identifierMap)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Beam":
		GongUnmarshallSliceOfPointers(&instance.Beam, valueExpr, identifierMap)
	case "Notations":
		GongUnmarshallSliceOfPointers(&instance.Notations, valueExpr, identifierMap)
	case "Lyric":
		GongUnmarshallSliceOfPointers(&instance.Lyric, valueExpr, identifierMap)
	case "Listen":
		GongUnmarshallPointer(&instance.Listen, valueExpr, identifierMap)
	}
	return nil
}

type NoteTypeUnmarshaller struct{}

func (u *NoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type NoteheadUnmarshaller struct{}

func (u *NoteheadUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Notehead), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteheadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notehead)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type NoteheadTextUnmarshaller struct{}

func (u *NoteheadTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteheadText), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteheadTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteheadText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type OrnamentsUnmarshaller struct{}

func (u *OrnamentsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Ornaments), stage, identifier, instanceName, preserveOrder)
}

func (u *OrnamentsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ornaments)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Mordent":
		GongUnmarshallSliceOfPointers(&instance.Mordent, valueExpr, identifierMap)
	case "InvertedMordent":
		GongUnmarshallSliceOfPointers(&instance.InvertedMordent, valueExpr, identifierMap)
	case "Tremolo":
		GongUnmarshallSliceOfPointers(&instance.Tremolo, valueExpr, identifierMap)
	case "OtherOrnament":
		GongUnmarshallSliceOfPointers(&instance.OtherOrnament, valueExpr, identifierMap)
	case "AccidentalMark":
		GongUnmarshallSliceOfPointers(&instance.AccidentalMark, valueExpr, identifierMap)
	}
	return nil
}

type OtherNotationUnmarshaller struct{}

func (u *OtherNotationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherNotation), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherNotationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherNotation)
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
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type OtherPlacementTextUnmarshaller struct{}

func (u *OtherPlacementTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherPlacementText), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherPlacementTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherPlacementText)
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
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PitchUnmarshaller struct{}

func (u *PitchUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pitch), stage, identifier, instanceName, preserveOrder)
}

func (u *PitchUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pitch)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Alter":
		instance.Alter = GongExtractString(valueExpr)
	case "Octave":
		instance.Octave = GongExtractInt(valueExpr)
	}
	return nil
}

type PlacementTextUnmarshaller struct{}

func (u *PlacementTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PlacementText), stage, identifier, instanceName, preserveOrder)
}

func (u *PlacementTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlacementText)
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

type ReleaseUnmarshaller struct{}

func (u *ReleaseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Release), stage, identifier, instanceName, preserveOrder)
}

func (u *ReleaseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Release)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Offset":
		instance.Offset = GongExtractString(valueExpr)
	}
	return nil
}

type RestUnmarshaller struct{}

func (u *RestUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Rest), stage, identifier, instanceName, preserveOrder)
}

func (u *RestUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rest)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DisplayOctave":
		instance.DisplayOctave = GongExtractInt(valueExpr)
	}
	return nil
}

type SlideUnmarshaller struct{}

func (u *SlideUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slide), stage, identifier, instanceName, preserveOrder)
}

func (u *SlideUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slide)
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
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "FirstBeat":
		instance.FirstBeat = GongExtractString(valueExpr)
	case "LastBeat":
		instance.LastBeat = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type SlurUnmarshaller struct{}

func (u *SlurUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slur), stage, identifier, instanceName, preserveOrder)
}

func (u *SlurUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slur)
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
	case "BezierX":
		instance.BezierX = GongExtractString(valueExpr)
	case "BezierY":
		instance.BezierY = GongExtractString(valueExpr)
	case "BezierX2":
		instance.BezierX2 = GongExtractString(valueExpr)
	case "BezierY2":
		instance.BezierY2 = GongExtractString(valueExpr)
	case "BezierOffset":
		instance.BezierOffset = GongExtractString(valueExpr)
	case "BezierOffset2":
		instance.BezierOffset2 = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type StemUnmarshaller struct{}

func (u *StemUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stem), stage, identifier, instanceName, preserveOrder)
}

func (u *StemUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stem)
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
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type StrongAccentUnmarshaller struct{}

func (u *StrongAccentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StrongAccent), stage, identifier, instanceName, preserveOrder)
}

func (u *StrongAccentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StrongAccent)
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
	}
	return nil
}

type TapUnmarshaller struct{}

func (u *TapUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tap), stage, identifier, instanceName, preserveOrder)
}

func (u *TapUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tap)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Hand":
		GongUnmarshallEnum(&instance.Hand, valueExpr)
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

type TechnicalUnmarshaller struct{}

func (u *TechnicalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Technical), stage, identifier, instanceName, preserveOrder)
}

func (u *TechnicalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Technical)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Harmonic":
		GongUnmarshallSliceOfPointers(&instance.Harmonic, valueExpr, identifierMap)
	case "Pluck":
		GongUnmarshallSliceOfPointers(&instance.Pluck, valueExpr, identifierMap)
	case "HammerOn":
		GongUnmarshallSliceOfPointers(&instance.HammerOn, valueExpr, identifierMap)
	case "PullOff":
		GongUnmarshallSliceOfPointers(&instance.PullOff, valueExpr, identifierMap)
	case "Bend":
		GongUnmarshallSliceOfPointers(&instance.Bend, valueExpr, identifierMap)
	case "Tap":
		GongUnmarshallSliceOfPointers(&instance.Tap, valueExpr, identifierMap)
	case "Heel":
		GongUnmarshallSliceOfPointers(&instance.Heel, valueExpr, identifierMap)
	case "Toe":
		GongUnmarshallSliceOfPointers(&instance.Toe, valueExpr, identifierMap)
	case "Hole":
		GongUnmarshallSliceOfPointers(&instance.Hole, valueExpr, identifierMap)
	case "Arrow":
		GongUnmarshallSliceOfPointers(&instance.Arrow, valueExpr, identifierMap)
	case "Handbell":
		GongUnmarshallSliceOfPointers(&instance.Handbell, valueExpr, identifierMap)
	case "HarmonMute":
		GongUnmarshallSliceOfPointers(&instance.HarmonMute, valueExpr, identifierMap)
	case "OtherTechnical":
		GongUnmarshallSliceOfPointers(&instance.OtherTechnical, valueExpr, identifierMap)
	}
	return nil
}

type TextElementDataUnmarshaller struct{}

func (u *TextElementDataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextElementData), stage, identifier, instanceName, preserveOrder)
}

func (u *TextElementDataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextElementData)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "LineThrough":
		instance.LineThrough = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type TieUnmarshaller struct{}

func (u *TieUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tie), stage, identifier, instanceName, preserveOrder)
}

func (u *TieUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tie)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
	}
	return nil
}

type TiedUnmarshaller struct{}

func (u *TiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tied), stage, identifier, instanceName, preserveOrder)
}

func (u *TiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tied)
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
	case "BezierX":
		instance.BezierX = GongExtractString(valueExpr)
	case "BezierY":
		instance.BezierY = GongExtractString(valueExpr)
	case "BezierX2":
		instance.BezierX2 = GongExtractString(valueExpr)
	case "BezierY2":
		instance.BezierY2 = GongExtractString(valueExpr)
	case "BezierOffset":
		instance.BezierOffset = GongExtractString(valueExpr)
	case "BezierOffset2":
		instance.BezierOffset2 = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type TremoloUnmarshaller struct{}

func (u *TremoloUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tremolo), stage, identifier, instanceName, preserveOrder)
}

func (u *TremoloUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tremolo)
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
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type TupletUnmarshaller struct{}

func (u *TupletUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tuplet), stage, identifier, instanceName, preserveOrder)
}

func (u *TupletUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet)
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
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "TupletActual":
		GongUnmarshallPointer(&instance.TupletActual, valueExpr, identifierMap)
	case "TupletNormal":
		GongUnmarshallPointer(&instance.TupletNormal, valueExpr, identifierMap)
	}
	return nil
}

type TupletDotUnmarshaller struct{}

func (u *TupletDotUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TupletDot), stage, identifier, instanceName, preserveOrder)
}

func (u *TupletDotUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TupletDot)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type TupletNumberUnmarshaller struct{}

func (u *TupletNumberUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TupletNumber), stage, identifier, instanceName, preserveOrder)
}

func (u *TupletNumberUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TupletNumber)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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

type TupletPortionUnmarshaller struct{}

func (u *TupletPortionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TupletPortion), stage, identifier, instanceName, preserveOrder)
}

func (u *TupletPortionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TupletPortion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TupletNumber":
		GongUnmarshallPointer(&instance.TupletNumber, valueExpr, identifierMap)
	case "TupletType":
		GongUnmarshallPointer(&instance.TupletType, valueExpr, identifierMap)
	case "TupletDot":
		GongUnmarshallSliceOfPointers(&instance.TupletDot, valueExpr, identifierMap)
	}
	return nil
}

type TupletTypeUnmarshaller struct{}

func (u *TupletTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TupletType), stage, identifier, instanceName, preserveOrder)
}

func (u *TupletTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TupletType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type UnpitchedUnmarshaller struct{}

func (u *UnpitchedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Unpitched), stage, identifier, instanceName, preserveOrder)
}

func (u *UnpitchedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Unpitched)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DisplayOctave":
		instance.DisplayOctave = GongExtractInt(valueExpr)
	}
	return nil
}

type WaitUnmarshaller struct{}

func (u *WaitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wait), stage, identifier, instanceName, preserveOrder)
}

func (u *WaitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wait)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "TimeOnly":
		instance.TimeOnly = GongExtractString(valueExpr)
	}
	return nil
}

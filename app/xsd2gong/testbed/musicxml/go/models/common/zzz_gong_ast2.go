// generated code - do not edit
package common

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
type AccidentalTextUnmarshaller struct{}

func (u *AccidentalTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AccidentalText), stage, identifier, instanceName, preserveOrder)
}

func (u *AccidentalTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AccidentalText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
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
	case "LineHeight":
		instance.LineHeight = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type AppearanceUnmarshaller struct{}

func (u *AppearanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Appearance), stage, identifier, instanceName, preserveOrder)
}

func (u *AppearanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Appearance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LineWidth":
		GongUnmarshallSliceOfPointers(&instance.LineWidth, valueExpr, identifierMap)
	case "NoteSize":
		GongUnmarshallSliceOfPointers(&instance.NoteSize, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallSliceOfPointers(&instance.Distance, valueExpr, identifierMap)
	case "Glyph":
		GongUnmarshallSliceOfPointers(&instance.Glyph, valueExpr, identifierMap)
	case "OtherAppearance":
		GongUnmarshallSliceOfPointers(&instance.OtherAppearance, valueExpr, identifierMap)
	}
	return nil
}

type BookmarkUnmarshaller struct{}

func (u *BookmarkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bookmark), stage, identifier, instanceName, preserveOrder)
}

func (u *BookmarkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bookmark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "NameAttr":
		instance.NameAttr = GongExtractString(valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Position":
		instance.Position = GongExtractInt(valueExpr)
	}
	return nil
}

type CodaUnmarshaller struct{}

func (u *CodaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Coda), stage, identifier, instanceName, preserveOrder)
}

func (u *CodaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Coda)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type DistanceUnmarshaller struct{}

func (u *DistanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Distance), stage, identifier, instanceName, preserveOrder)
}

func (u *DistanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Distance)
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

type DynamicsUnmarshaller struct{}

func (u *DynamicsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Dynamics), stage, identifier, instanceName, preserveOrder)
}

func (u *DynamicsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dynamics)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "LineThrough":
		instance.LineThrough = GongExtractInt(valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "P":
		GongUnmarshallSliceOfPointers(&instance.P, valueExpr, identifierMap)
	case "Pp":
		GongUnmarshallSliceOfPointers(&instance.Pp, valueExpr, identifierMap)
	case "Ppp":
		GongUnmarshallSliceOfPointers(&instance.Ppp, valueExpr, identifierMap)
	case "Pppp":
		GongUnmarshallSliceOfPointers(&instance.Pppp, valueExpr, identifierMap)
	case "Ppppp":
		GongUnmarshallSliceOfPointers(&instance.Ppppp, valueExpr, identifierMap)
	case "Pppppp":
		GongUnmarshallSliceOfPointers(&instance.Pppppp, valueExpr, identifierMap)
	case "F":
		GongUnmarshallSliceOfPointers(&instance.F, valueExpr, identifierMap)
	case "Ff":
		GongUnmarshallSliceOfPointers(&instance.Ff, valueExpr, identifierMap)
	case "Fff":
		GongUnmarshallSliceOfPointers(&instance.Fff, valueExpr, identifierMap)
	case "Ffff":
		GongUnmarshallSliceOfPointers(&instance.Ffff, valueExpr, identifierMap)
	case "Fffff":
		GongUnmarshallSliceOfPointers(&instance.Fffff, valueExpr, identifierMap)
	case "Ffffff":
		GongUnmarshallSliceOfPointers(&instance.Ffffff, valueExpr, identifierMap)
	case "Mp":
		GongUnmarshallSliceOfPointers(&instance.Mp, valueExpr, identifierMap)
	case "Mf":
		GongUnmarshallSliceOfPointers(&instance.Mf, valueExpr, identifierMap)
	case "Sf":
		GongUnmarshallSliceOfPointers(&instance.Sf, valueExpr, identifierMap)
	case "Sfp":
		GongUnmarshallSliceOfPointers(&instance.Sfp, valueExpr, identifierMap)
	case "Sfpp":
		GongUnmarshallSliceOfPointers(&instance.Sfpp, valueExpr, identifierMap)
	case "Fp":
		GongUnmarshallSliceOfPointers(&instance.Fp, valueExpr, identifierMap)
	case "Rf":
		GongUnmarshallSliceOfPointers(&instance.Rf, valueExpr, identifierMap)
	case "Rfz":
		GongUnmarshallSliceOfPointers(&instance.Rfz, valueExpr, identifierMap)
	case "Sfz":
		GongUnmarshallSliceOfPointers(&instance.Sfz, valueExpr, identifierMap)
	case "Sffz":
		GongUnmarshallSliceOfPointers(&instance.Sffz, valueExpr, identifierMap)
	case "Fz":
		GongUnmarshallSliceOfPointers(&instance.Fz, valueExpr, identifierMap)
	case "N":
		GongUnmarshallSliceOfPointers(&instance.N, valueExpr, identifierMap)
	case "Pf":
		GongUnmarshallSliceOfPointers(&instance.Pf, valueExpr, identifierMap)
	case "Sfzp":
		GongUnmarshallSliceOfPointers(&instance.Sfzp, valueExpr, identifierMap)
	case "OtherDynamics":
		GongUnmarshallSliceOfPointers(&instance.OtherDynamics, valueExpr, identifierMap)
	}
	return nil
}

type EmptyUnmarshaller struct{}

func (u *EmptyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EmptyPlacementUnmarshaller struct{}

func (u *EmptyPlacementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyPlacement), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyPlacementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyPlacement)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	}
	return nil
}

type EmptyPlacementSmuflUnmarshaller struct{}

func (u *EmptyPlacementSmuflUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyPlacementSmufl), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyPlacementSmuflUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyPlacementSmufl)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	}
	return nil
}

type EmptyPrintObjectStyleAlignUnmarshaller struct{}

func (u *EmptyPrintObjectStyleAlignUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyPrintObjectStyleAlign), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyPrintObjectStyleAlignUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyPrintObjectStyleAlign)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PrintObject":
		GongUnmarshallEnum(&instance.PrintObject, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	}
	return nil
}

type EmptyPrintStyleUnmarshaller struct{}

func (u *EmptyPrintStyleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyPrintStyle), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyPrintStyleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyPrintStyle)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type EmptyPrintStyleAlignUnmarshaller struct{}

func (u *EmptyPrintStyleAlignUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyPrintStyleAlign), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyPrintStyleAlignUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyPrintStyleAlign)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	}
	return nil
}

type EmptyPrintStyleAlignIdUnmarshaller struct{}

func (u *EmptyPrintStyleAlignIdUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyPrintStyleAlignId), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyPrintStyleAlignIdUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyPrintStyleAlignId)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type EmptyTrillSoundUnmarshaller struct{}

func (u *EmptyTrillSoundUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmptyTrillSound), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyTrillSoundUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmptyTrillSound)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "StartNote":
		GongUnmarshallEnum(&instance.StartNote, valueExpr)
	case "TrillStep":
		GongUnmarshallEnum(&instance.TrillStep, valueExpr)
	case "TwoNoteTurn":
		GongUnmarshallEnum(&instance.TwoNoteTurn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "SecondBeat":
		instance.SecondBeat = GongExtractString(valueExpr)
	case "LastBeat":
		instance.LastBeat = GongExtractString(valueExpr)
	}
	return nil
}

type EncodingUnmarshaller struct{}

func (u *EncodingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Encoding), stage, identifier, instanceName, preserveOrder)
}

func (u *EncodingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Encoding)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EncodingDate":
		instance.EncodingDate = GongExtractString(valueExpr)
	case "Encoder":
		GongUnmarshallSliceOfPointers(&instance.Encoder, valueExpr, identifierMap)
	case "Software":
		instance.Software = GongExtractString(valueExpr)
	case "EncodingDescription":
		instance.EncodingDescription = GongExtractString(valueExpr)
	case "Supports":
		GongUnmarshallSliceOfPointers(&instance.Supports, valueExpr, identifierMap)
	}
	return nil
}

type FermataUnmarshaller struct{}

func (u *FermataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Fermata), stage, identifier, instanceName, preserveOrder)
}

func (u *FermataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fermata)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallEnum(&instance.Value, valueExpr)
	}
	return nil
}

type FingeringUnmarshaller struct{}

func (u *FingeringUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Fingering), stage, identifier, instanceName, preserveOrder)
}

func (u *FingeringUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fingering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Substitution":
		GongUnmarshallEnum(&instance.Substitution, valueExpr)
	case "Alternate":
		GongUnmarshallEnum(&instance.Alternate, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type FormattedSymbolUnmarshaller struct{}

func (u *FormattedSymbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormattedSymbol), stage, identifier, instanceName, preserveOrder)
}

func (u *FormattedSymbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormattedSymbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
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
	case "LineHeight":
		instance.LineHeight = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type FormattedSymbolIdUnmarshaller struct{}

func (u *FormattedSymbolIdUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormattedSymbolId), stage, identifier, instanceName, preserveOrder)
}

func (u *FormattedSymbolIdUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormattedSymbolId)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
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
	case "LineHeight":
		instance.LineHeight = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type FormattedTextUnmarshaller struct{}

func (u *FormattedTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormattedText), stage, identifier, instanceName, preserveOrder)
}

func (u *FormattedTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormattedText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
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
	case "LineHeight":
		instance.LineHeight = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type FormattedTextIdUnmarshaller struct{}

func (u *FormattedTextIdUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormattedTextId), stage, identifier, instanceName, preserveOrder)
}

func (u *FormattedTextIdUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormattedTextId)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
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
	case "LineHeight":
		instance.LineHeight = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type FretUnmarshaller struct{}

func (u *FretUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Fret), stage, identifier, instanceName, preserveOrder)
}

func (u *FretUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fret)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type GlyphUnmarshaller struct{}

func (u *GlyphUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Glyph), stage, identifier, instanceName, preserveOrder)
}

func (u *GlyphUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glyph)
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

type HorizontalTurnUnmarshaller struct{}

func (u *HorizontalTurnUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HorizontalTurn), stage, identifier, instanceName, preserveOrder)
}

func (u *HorizontalTurnUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HorizontalTurn)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Slash":
		GongUnmarshallEnum(&instance.Slash, valueExpr)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "StartNote":
		GongUnmarshallEnum(&instance.StartNote, valueExpr)
	case "TrillStep":
		GongUnmarshallEnum(&instance.TrillStep, valueExpr)
	case "TwoNoteTurn":
		GongUnmarshallEnum(&instance.TwoNoteTurn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "SecondBeat":
		instance.SecondBeat = GongExtractString(valueExpr)
	case "LastBeat":
		instance.LastBeat = GongExtractString(valueExpr)
	}
	return nil
}

type IdentificationUnmarshaller struct{}

func (u *IdentificationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Identification), stage, identifier, instanceName, preserveOrder)
}

func (u *IdentificationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Identification)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Creator":
		GongUnmarshallSliceOfPointers(&instance.Creator, valueExpr, identifierMap)
	case "Rights":
		GongUnmarshallSliceOfPointers(&instance.Rights, valueExpr, identifierMap)
	case "Encoding":
		GongUnmarshallPointer(&instance.Encoding, valueExpr, identifierMap)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Relation":
		GongUnmarshallSliceOfPointers(&instance.Relation, valueExpr, identifierMap)
	case "Miscellaneous":
		GongUnmarshallPointer(&instance.Miscellaneous, valueExpr, identifierMap)
	}
	return nil
}

type LevelUnmarshaller struct{}

func (u *LevelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Level), stage, identifier, instanceName, preserveOrder)
}

func (u *LevelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Level)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference":
		GongUnmarshallEnum(&instance.Reference, valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type LineWidthUnmarshaller struct{}

func (u *LineWidthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LineWidth), stage, identifier, instanceName, preserveOrder)
}

func (u *LineWidthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LineWidth)
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
	case "NameAttr":
		instance.NameAttr = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Position":
		instance.Position = GongExtractInt(valueExpr)
	case "DefaultX":
		instance.DefaultX = GongExtractString(valueExpr)
	case "DefaultY":
		instance.DefaultY = GongExtractString(valueExpr)
	case "RelativeX":
		instance.RelativeX = GongExtractString(valueExpr)
	case "RelativeY":
		instance.RelativeY = GongExtractString(valueExpr)
	}
	return nil
}

type MeasureLayoutUnmarshaller struct{}

func (u *MeasureLayoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeasureLayout), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureLayoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeasureLayout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MeasureDistance":
		instance.MeasureDistance = GongExtractString(valueExpr)
	}
	return nil
}

type MidiDeviceUnmarshaller struct{}

func (u *MidiDeviceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MidiDevice), stage, identifier, instanceName, preserveOrder)
}

func (u *MidiDeviceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MidiDevice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Port":
		instance.Port = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MidiInstrumentUnmarshaller struct{}

func (u *MidiInstrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MidiInstrument), stage, identifier, instanceName, preserveOrder)
}

func (u *MidiInstrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MidiInstrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "MidiChannel":
		instance.MidiChannel = GongExtractInt(valueExpr)
	case "MidiName":
		instance.MidiName = GongExtractString(valueExpr)
	case "MidiBank":
		instance.MidiBank = GongExtractInt(valueExpr)
	case "MidiProgram":
		instance.MidiProgram = GongExtractInt(valueExpr)
	case "MidiUnpitched":
		instance.MidiUnpitched = GongExtractInt(valueExpr)
	case "Volume":
		instance.Volume = GongExtractString(valueExpr)
	case "Pan":
		instance.Pan = GongExtractString(valueExpr)
	case "Elevation":
		instance.Elevation = GongExtractString(valueExpr)
	}
	return nil
}

type MiscellaneousUnmarshaller struct{}

func (u *MiscellaneousUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Miscellaneous), stage, identifier, instanceName, preserveOrder)
}

func (u *MiscellaneousUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Miscellaneous)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MiscellaneousField":
		GongUnmarshallSliceOfPointers(&instance.MiscellaneousField, valueExpr, identifierMap)
	}
	return nil
}

type MiscellaneousFieldUnmarshaller struct{}

func (u *MiscellaneousFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MiscellaneousField), stage, identifier, instanceName, preserveOrder)
}

func (u *MiscellaneousFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MiscellaneousField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameAttr":
		instance.NameAttr = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type NameDisplayUnmarshaller struct{}

func (u *NameDisplayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NameDisplay), stage, identifier, instanceName, preserveOrder)
}

func (u *NameDisplayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NameDisplay)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PrintObject":
		GongUnmarshallEnum(&instance.PrintObject, valueExpr)
	case "DisplayText":
		GongUnmarshallSliceOfPointers(&instance.DisplayText, valueExpr, identifierMap)
	case "AccidentalText":
		GongUnmarshallSliceOfPointers(&instance.AccidentalText, valueExpr, identifierMap)
	}
	return nil
}

type NoteSizeUnmarshaller struct{}

func (u *NoteSizeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteSize), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteSizeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteSize)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type OtherAppearanceUnmarshaller struct{}

func (u *OtherAppearanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherAppearance), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherAppearanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherAppearance)
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

type OtherPlayUnmarshaller struct{}

func (u *OtherPlayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherPlay), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherPlayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherPlay)
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

type OtherTextUnmarshaller struct{}

func (u *OtherTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OtherText), stage, identifier, instanceName, preserveOrder)
}

func (u *OtherTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OtherText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PageLayoutUnmarshaller struct{}

func (u *PageLayoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PageLayout), stage, identifier, instanceName, preserveOrder)
}

func (u *PageLayoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PageLayout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PageHeight":
		instance.PageHeight = GongExtractString(valueExpr)
	case "PageWidth":
		instance.PageWidth = GongExtractString(valueExpr)
	case "PageMargins":
		GongUnmarshallSliceOfPointers(&instance.PageMargins, valueExpr, identifierMap)
	}
	return nil
}

type PageMarginsUnmarshaller struct{}

func (u *PageMarginsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PageMargins), stage, identifier, instanceName, preserveOrder)
}

func (u *PageMarginsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PageMargins)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "LeftMargin":
		instance.LeftMargin = GongExtractString(valueExpr)
	case "RightMargin":
		instance.RightMargin = GongExtractString(valueExpr)
	case "TopMargin":
		instance.TopMargin = GongExtractString(valueExpr)
	case "BottomMargin":
		instance.BottomMargin = GongExtractString(valueExpr)
	}
	return nil
}

type PlayUnmarshaller struct{}

func (u *PlayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Play), stage, identifier, instanceName, preserveOrder)
}

func (u *PlayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Play)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Ipa":
		instance.Ipa = GongExtractString(valueExpr)
	case "Mute":
		GongUnmarshallEnum(&instance.Mute, valueExpr)
	case "SemiPitched":
		GongUnmarshallEnum(&instance.SemiPitched, valueExpr)
	case "OtherPlay":
		GongUnmarshallSliceOfPointers(&instance.OtherPlay, valueExpr, identifierMap)
	}
	return nil
}

type ScalingUnmarshaller struct{}

func (u *ScalingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Scaling), stage, identifier, instanceName, preserveOrder)
}

func (u *ScalingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scaling)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Millimeters":
		instance.Millimeters = GongExtractString(valueExpr)
	case "Tenths":
		instance.Tenths = GongExtractString(valueExpr)
	}
	return nil
}

type SegnoUnmarshaller struct{}

func (u *SegnoUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Segno), stage, identifier, instanceName, preserveOrder)
}

func (u *SegnoUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Segno)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type StaffLayoutUnmarshaller struct{}

func (u *StaffLayoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaffLayout), stage, identifier, instanceName, preserveOrder)
}

func (u *StaffLayoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaffLayout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "StaffDistance":
		instance.StaffDistance = GongExtractString(valueExpr)
	}
	return nil
}

type StringTypeUnmarshaller struct{}

func (u *StringTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StringType), stage, identifier, instanceName, preserveOrder)
}

func (u *StringTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StringType)
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
	case "FontStyle":
		GongUnmarshallEnum(&instance.FontStyle, valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		GongUnmarshallEnum(&instance.FontWeight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	}
	return nil
}

type SupportsUnmarshaller struct{}

func (u *SupportsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Supports), stage, identifier, instanceName, preserveOrder)
}

func (u *SupportsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Supports)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Attribute":
		instance.Attribute = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type SystemDividersUnmarshaller struct{}

func (u *SystemDividersUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SystemDividers), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemDividersUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SystemDividers)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LeftDivider":
		GongUnmarshallPointer(&instance.LeftDivider, valueExpr, identifierMap)
	case "RightDivider":
		GongUnmarshallPointer(&instance.RightDivider, valueExpr, identifierMap)
	}
	return nil
}

type SystemLayoutUnmarshaller struct{}

func (u *SystemLayoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SystemLayout), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemLayoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SystemLayout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SystemMargins":
		GongUnmarshallPointer(&instance.SystemMargins, valueExpr, identifierMap)
	case "SystemDistance":
		instance.SystemDistance = GongExtractString(valueExpr)
	case "TopSystemDistance":
		instance.TopSystemDistance = GongExtractString(valueExpr)
	case "SystemDividers":
		GongUnmarshallPointer(&instance.SystemDividers, valueExpr, identifierMap)
	}
	return nil
}

type SystemMarginsUnmarshaller struct{}

func (u *SystemMarginsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SystemMargins), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemMarginsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SystemMargins)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LeftMargin":
		instance.LeftMargin = GongExtractString(valueExpr)
	case "RightMargin":
		instance.RightMargin = GongExtractString(valueExpr)
	}
	return nil
}

type TypedTextUnmarshaller struct{}

func (u *TypedTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TypedText), stage, identifier, instanceName, preserveOrder)
}

func (u *TypedTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TypedText)
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

type WavyLineUnmarshaller struct{}

func (u *WavyLineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WavyLine), stage, identifier, instanceName, preserveOrder)
}

func (u *WavyLineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WavyLine)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
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
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "StartNote":
		GongUnmarshallEnum(&instance.StartNote, valueExpr)
	case "TrillStep":
		GongUnmarshallEnum(&instance.TrillStep, valueExpr)
	case "TwoNoteTurn":
		GongUnmarshallEnum(&instance.TwoNoteTurn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "SecondBeat":
		instance.SecondBeat = GongExtractString(valueExpr)
	case "LastBeat":
		instance.LastBeat = GongExtractString(valueExpr)
	}
	return nil
}

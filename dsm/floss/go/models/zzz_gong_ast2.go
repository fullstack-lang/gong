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
type CompareAnalysisUnmarshaller struct{}

func (u *CompareAnalysisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CompareAnalysis), stage, identifier, instanceName, preserveOrder)
}

func (u *CompareAnalysisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CompareAnalysis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FromSystem":
		GongUnmarshallPointer(&instance.FromSystem, valueExpr, identifierMap)
	case "ToSystem":
		GongUnmarshallPointer(&instance.ToSystem, valueExpr, identifierMap)
	case "Mu":
		instance.Mu = GongExtractFloat(valueExpr)
	case "Epsilon":
		instance.Epsilon = GongExtractFloat(valueExpr)
	case "DiagramFlossEquations":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquations, valueExpr, identifierMap)
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquationsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ComplexityUnmarshaller struct{}

func (u *ComplexityUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Complexity), stage, identifier, instanceName, preserveOrder)
}

func (u *ComplexityUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Complexity)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Strength":
		instance.Strength = GongExtractFloat(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type DiagramFlossEquationUnmarshaller struct{}

func (u *DiagramFlossEquationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramFlossEquation), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramFlossEquationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramFlossEquation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Scale":
		instance.Scale = GongExtractFloat(valueExpr)
	case "FontSize":
		GongUnmarshallEnum(&instance.FontSize, valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "IsEditable_":
		instance.IsEditable_ = GongExtractBool(valueExpr)
	case "IsInDelta3ColumnsMode":
		instance.IsInDelta3ColumnsMode = GongExtractBool(valueExpr)
	case "AreQuantitativeElementsVisible":
		instance.AreQuantitativeElementsVisible = GongExtractBool(valueExpr)
	case "AreSubsystemsVisible":
		instance.AreSubsystemsVisible = GongExtractBool(valueExpr)
	case "AreCommonElementsHidden":
		instance.AreCommonElementsHidden = GongExtractBool(valueExpr)
	case "AreCPEArrowsVisible":
		instance.AreCPEArrowsVisible = GongExtractBool(valueExpr)
	case "AreColumnTitlesVisible":
		instance.AreColumnTitlesVisible = GongExtractBool(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "DefaultBoxWidth":
		instance.DefaultBoxWidth = GongExtractFloat(valueExpr)
	case "DefaultBoxHeigth":
		instance.DefaultBoxHeigth = GongExtractFloat(valueExpr)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NoteComplexityShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteComplexityShapes, valueExpr, identifierMap)
	case "NotePerformanceShapes":
		GongUnmarshallSliceOfPointers(&instance.NotePerformanceShapes, valueExpr, identifierMap)
	case "NoteEffortShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteEffortShapes, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "ComplexitysWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ComplexitysWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "PerformancesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PerformancesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	case "EffortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EffortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type EffortUnmarshaller struct{}

func (u *EffortUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Effort), stage, identifier, instanceName, preserveOrder)
}

func (u *EffortUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Effort)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Strength":
		instance.Strength = GongExtractFloat(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "RootSystems":
		GongUnmarshallSliceOfPointers(&instance.RootSystems, valueExpr, identifierMap)
	case "RootComplexitys":
		GongUnmarshallSliceOfPointers(&instance.RootComplexitys, valueExpr, identifierMap)
	case "RootPerformances":
		GongUnmarshallSliceOfPointers(&instance.RootPerformances, valueExpr, identifierMap)
	case "RootEfforts":
		GongUnmarshallSliceOfPointers(&instance.RootEfforts, valueExpr, identifierMap)
	case "RootCompareAnalysis":
		GongUnmarshallSliceOfPointers(&instance.RootCompareAnalysis, valueExpr, identifierMap)
	case "RootNotes":
		GongUnmarshallSliceOfPointers(&instance.RootNotes, valueExpr, identifierMap)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "IsSubLibrariesNodeExpanded":
		instance.IsSubLibrariesNodeExpanded = GongExtractBool(valueExpr)
	case "SubLibrariesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SubLibrariesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "IsSystemsNodeExpanded":
		instance.IsSystemsNodeExpanded = GongExtractBool(valueExpr)
	case "SystemsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SystemsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "ComplexitysWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ComplexitysWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "PerformancesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PerformancesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	case "EffortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EffortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsCompareAnalysisNodeExpanded":
		instance.IsCompareAnalysisNodeExpanded = GongExtractBool(valueExpr)
	case "CompareAnalysisWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.CompareAnalysisWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsExpandedTmp":
		instance.IsExpandedTmp = GongExtractBool(valueExpr)
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
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Complexities":
		GongUnmarshallSliceOfPointers(&instance.Complexities, valueExpr, identifierMap)
	case "Performances":
		GongUnmarshallSliceOfPointers(&instance.Performances, valueExpr, identifierMap)
	case "Efforts":
		GongUnmarshallSliceOfPointers(&instance.Efforts, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type NoteComplexityShapeUnmarshaller struct{}

func (u *NoteComplexityShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteComplexityShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteComplexityShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteComplexityShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Complexity":
		GongUnmarshallPointer(&instance.Complexity, valueExpr, identifierMap)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type NoteEffortShapeUnmarshaller struct{}

func (u *NoteEffortShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteEffortShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteEffortShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteEffortShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Effort":
		GongUnmarshallPointer(&instance.Effort, valueExpr, identifierMap)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type NotePerformanceShapeUnmarshaller struct{}

func (u *NotePerformanceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePerformanceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePerformanceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePerformanceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Performance":
		GongUnmarshallPointer(&instance.Performance, valueExpr, identifierMap)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type NoteShapeUnmarshaller struct{}

func (u *NoteShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
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

type PerformanceUnmarshaller struct{}

func (u *PerformanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Performance), stage, identifier, instanceName, preserveOrder)
}

func (u *PerformanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Performance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Strength":
		instance.Strength = GongExtractFloat(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type SystemUnmarshaller struct{}

func (u *SystemUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(System), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Complexities":
		GongUnmarshallSliceOfPointers(&instance.Complexities, valueExpr, identifierMap)
	case "Performances":
		GongUnmarshallSliceOfPointers(&instance.Performances, valueExpr, identifierMap)
	case "Efforts":
		GongUnmarshallSliceOfPointers(&instance.Efforts, valueExpr, identifierMap)
	case "SubSystems":
		GongUnmarshallSliceOfPointers(&instance.SubSystems, valueExpr, identifierMap)
	case "AreCPEsCompoundedFromSubSystems":
		instance.AreCPEsCompoundedFromSubSystems = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	case "DiagramFlossEquations":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquations, valueExpr, identifierMap)
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquationsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubSystemNodeExpanded":
		instance.IsSubSystemNodeExpanded = GongExtractBool(valueExpr)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "ComplexitysWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ComplexitysWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "PerformancesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PerformancesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	case "EffortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EffortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

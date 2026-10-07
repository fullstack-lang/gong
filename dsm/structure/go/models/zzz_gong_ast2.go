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
type AllocatedResourceShapeUnmarshaller struct{}

func (u *AllocatedResourceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AllocatedResourceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AllocatedResourceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AllocatedResourceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
	}
	return nil
}

type AllocatedSystemShapeUnmarshaller struct{}

func (u *AllocatedSystemShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AllocatedSystemShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AllocatedSystemShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AllocatedSystemShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "System":
		GongUnmarshallPointer(&instance.System, valueExpr, identifierMap)
	}
	return nil
}

type ControlFlowUnmarshaller struct{}

func (u *ControlFlowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlFlow), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlFlowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlFlow)
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
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	}
	return nil
}

type ControlFlowShapeUnmarshaller struct{}

func (u *ControlFlowShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlFlowShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlFlowShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlFlowShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ControlFlow":
		GongUnmarshallPointer(&instance.ControlFlow, valueExpr, identifierMap)
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

type DataUnmarshaller struct{}

func (u *DataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Data), stage, identifier, instanceName, preserveOrder)
}

func (u *DataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Data)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Acronym":
		instance.Acronym = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	}
	return nil
}

type DataFlowUnmarshaller struct{}

func (u *DataFlowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DataFlow), stage, identifier, instanceName, preserveOrder)
}

func (u *DataFlowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DataFlow)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartPort":
		GongUnmarshallPointer(&instance.StartPort, valueExpr, identifierMap)
	case "EndPort":
		GongUnmarshallPointer(&instance.EndPort, valueExpr, identifierMap)
	case "StartExternalPart":
		GongUnmarshallPointer(&instance.StartExternalPart, valueExpr, identifierMap)
	case "EndExternalPart":
		GongUnmarshallPointer(&instance.EndExternalPart, valueExpr, identifierMap)
	case "Datas":
		GongUnmarshallSliceOfPointers(&instance.Datas, valueExpr, identifierMap)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "IsDatasNodeExpanded":
		instance.IsDatasNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	}
	return nil
}

type DataFlowShapeUnmarshaller struct{}

func (u *DataFlowShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DataFlowShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DataFlowShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DataFlowShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DataFlow":
		GongUnmarshallPointer(&instance.DataFlow, valueExpr, identifierMap)
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

type DataShapeUnmarshaller struct{}

func (u *DataShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DataShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DataShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DataShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Data":
		GongUnmarshallPointer(&instance.Data, valueExpr, identifierMap)
	case "DataFlow":
		GongUnmarshallPointer(&instance.DataFlow, valueExpr, identifierMap)
	}
	return nil
}

type DiagramLayerStateUnmarshaller struct{}

func (u *DiagramLayerStateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramLayerState), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramLayerStateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramLayerState)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DiagramStructure":
		GongUnmarshallPointer(&instance.DiagramStructure, valueExpr, identifierMap)
	case "LayerDefinition":
		GongUnmarshallPointer(&instance.LayerDefinition, valueExpr, identifierMap)
	}
	return nil
}

type DiagramStructureUnmarshaller struct{}

func (u *DiagramStructureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramStructure), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramStructureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramStructure)
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
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "IsEditable_":
		instance.IsEditable_ = GongExtractBool(valueExpr)
	case "IsShowPrefix":
		instance.IsShowPrefix = GongExtractBool(valueExpr)
	case "DefaultBoxWidth":
		instance.DefaultBoxWidth = GongExtractFloat(valueExpr)
	case "DefaultBoxHeigth":
		instance.DefaultBoxHeigth = GongExtractFloat(valueExpr)
	case "IsWithDiscretePorts":
		instance.IsWithDiscretePorts = GongExtractBool(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "System_Shapes":
		GongUnmarshallSliceOfPointers(&instance.System_Shapes, valueExpr, identifierMap)
	case "IsSystemsNodeExpanded":
		instance.IsSystemsNodeExpanded = GongExtractBool(valueExpr)
	case "SystemsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SystemsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Part_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Part_Shapes, valueExpr, identifierMap)
	case "IsPartsNodeExpanded":
		instance.IsPartsNodeExpanded = GongExtractBool(valueExpr)
	case "PartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PartWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalPart_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ExternalPart_Shapes, valueExpr, identifierMap)
	case "IsExternalPartsNodeExpanded":
		instance.IsExternalPartsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalPartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartsWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Port_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Port_Shapes, valueExpr, identifierMap)
	case "ControlFlowsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ControlFlowsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ControlFlow_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ControlFlow_Shapes, valueExpr, identifierMap)
	case "DataFlowsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DataFlowsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DataFlow_Shapes":
		GongUnmarshallSliceOfPointers(&instance.DataFlow_Shapes, valueExpr, identifierMap)
	case "DatasWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DatasWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Data_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Data_Shapes, valueExpr, identifierMap)
	case "DataFlowsWhoseDataNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DataFlowsWhoseDataNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedResourcesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.AllocatedResourcesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedResourceShapes":
		GongUnmarshallSliceOfPointers(&instance.AllocatedResourceShapes, valueExpr, identifierMap)
	case "AllocatedSystemesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.AllocatedSystemesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedSystemShapes":
		GongUnmarshallSliceOfPointers(&instance.AllocatedSystemShapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NotePortShapes":
		GongUnmarshallSliceOfPointers(&instance.NotePortShapes, valueExpr, identifierMap)
	case "NotePartShapes":
		GongUnmarshallSliceOfPointers(&instance.NotePartShapes, valueExpr, identifierMap)
	}
	return nil
}

type ExternalPartShapeUnmarshaller struct{}

func (u *ExternalPartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ExternalPartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ExternalPartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExternalPartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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
	case "TailHeigth":
		instance.TailHeigth = GongExtractFloat(valueExpr)
	}
	return nil
}

type LayerDefinitionUnmarshaller struct{}

func (u *LayerDefinitionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LayerDefinition), stage, identifier, instanceName, preserveOrder)
}

func (u *LayerDefinitionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LayerDefinition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Query":
		GongUnmarshallSliceOfPointers(&instance.Query, valueExpr, identifierMap)
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
	case "RootSystemes":
		GongUnmarshallSliceOfPointers(&instance.RootSystemes, valueExpr, identifierMap)
	case "IsSystemesNodeExpanded":
		instance.IsSystemesNodeExpanded = GongExtractBool(valueExpr)
	case "SystemsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SystemsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootDataFlows":
		GongUnmarshallSliceOfPointers(&instance.RootDataFlows, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "DataFlowsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DataFlowsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootDatas":
		GongUnmarshallSliceOfPointers(&instance.RootDatas, valueExpr, identifierMap)
	case "IsDatasNodeExpanded":
		instance.IsDatasNodeExpanded = GongExtractBool(valueExpr)
	case "DatasWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DatasWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootResources":
		GongUnmarshallSliceOfPointers(&instance.RootResources, valueExpr, identifierMap)
	case "IsResourcesNodeExpanded":
		instance.IsResourcesNodeExpanded = GongExtractBool(valueExpr)
	case "ResourcesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ResourcesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "PartsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PartsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootNotes":
		GongUnmarshallSliceOfPointers(&instance.RootNotes, valueExpr, identifierMap)
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
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsPartsNodeExpanded":
		instance.IsPartsNodeExpanded = GongExtractBool(valueExpr)
	case "Parts":
		GongUnmarshallSliceOfPointers(&instance.Parts, valueExpr, identifierMap)
	case "IsPortsNodeExpanded":
		instance.IsPortsNodeExpanded = GongExtractBool(valueExpr)
	case "Ports":
		GongUnmarshallSliceOfPointers(&instance.Ports, valueExpr, identifierMap)
	}
	return nil
}

type NotePartShapeUnmarshaller struct{}

func (u *NotePartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
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

type NotePortShapeUnmarshaller struct{}

func (u *NotePortShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePortShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePortShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePortShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Port":
		GongUnmarshallPointer(&instance.Port, valueExpr, identifierMap)
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
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Ports":
		GongUnmarshallSliceOfPointers(&instance.Ports, valueExpr, identifierMap)
	case "TypeOfPart":
		GongUnmarshallPointer(&instance.TypeOfPart, valueExpr, identifierMap)
	case "IsPartNameNotSystemName":
		instance.IsPartNameNotSystemName = GongExtractBool(valueExpr)
	case "IsControlFlowsNodeExpanded":
		instance.IsControlFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ControlFlows":
		GongUnmarshallSliceOfPointers(&instance.ControlFlows, valueExpr, identifierMap)
	case "PortWhoseOutControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseOutControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PortWhoseInControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseInControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "PortWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PortWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PartAnchoredPath":
		GongUnmarshallSliceOfPointers(&instance.PartAnchoredPath, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsPortsNodeExpanded":
		instance.IsPortsNodeExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type PartAnchoredPathUnmarshaller struct{}

func (u *PartAnchoredPathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartAnchoredPath), stage, identifier, instanceName, preserveOrder)
}

func (u *PartAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartAnchoredPath)
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

type PartShapeUnmarshaller struct{}

func (u *PartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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

type PortUnmarshaller struct{}

func (u *PortUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Port), stage, identifier, instanceName, preserveOrder)
}

func (u *PortUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Port)
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
	}
	return nil
}

type PortShapeUnmarshaller struct{}

func (u *PortShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PortShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PortShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PortShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Port":
		GongUnmarshallPointer(&instance.Port, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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

type ResourceUnmarshaller struct{}

func (u *ResourceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Resource), stage, identifier, instanceName, preserveOrder)
}

func (u *ResourceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Resource)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Acronym":
		instance.Acronym = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	}
	return nil
}

type SemanticTagUnmarshaller struct{}

func (u *SemanticTagUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SemanticTag), stage, identifier, instanceName, preserveOrder)
}

func (u *SemanticTagUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SemanticTag)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parts":
		GongUnmarshallSliceOfPointers(&instance.Parts, valueExpr, identifierMap)
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
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	case "DiagramStructures":
		GongUnmarshallSliceOfPointers(&instance.DiagramStructures, valueExpr, identifierMap)
	case "DiagramStructureWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramStructureWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubSystemNodeExpanded":
		instance.IsSubSystemNodeExpanded = GongExtractBool(valueExpr)
	case "SubSystemes":
		GongUnmarshallSliceOfPointers(&instance.SubSystemes, valueExpr, identifierMap)
	case "Parts":
		GongUnmarshallSliceOfPointers(&instance.Parts, valueExpr, identifierMap)
	case "PartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PartWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DataFlows":
		GongUnmarshallSliceOfPointers(&instance.DataFlows, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalParts":
		GongUnmarshallSliceOfPointers(&instance.ExternalParts, valueExpr, identifierMap)
	case "ExternalPartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type SystemShapeUnmarshaller struct{}

func (u *SystemShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SystemShape), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SystemShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "System":
		GongUnmarshallPointer(&instance.System, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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

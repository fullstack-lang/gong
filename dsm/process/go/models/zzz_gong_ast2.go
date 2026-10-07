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
type AllocatedProcessShapeUnmarshaller struct{}

func (u *AllocatedProcessShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AllocatedProcessShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AllocatedProcessShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AllocatedProcessShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
	case "Process":
		GongUnmarshallPointer(&instance.Process, valueExpr, identifierMap)
	}
	return nil
}

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
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
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
	case "Datas":
		GongUnmarshallSliceOfPointers(&instance.Datas, valueExpr, identifierMap)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "StartTask":
		GongUnmarshallPointer(&instance.StartTask, valueExpr, identifierMap)
	case "EndTask":
		GongUnmarshallPointer(&instance.EndTask, valueExpr, identifierMap)
	case "StartExternalParticipant":
		GongUnmarshallPointer(&instance.StartExternalParticipant, valueExpr, identifierMap)
	case "EndExternalParticipant":
		GongUnmarshallPointer(&instance.EndExternalParticipant, valueExpr, identifierMap)
	case "IsDatasNodeExpanded":
		instance.IsDatasNodeExpanded = GongExtractBool(valueExpr)
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

type DiagramProcessUnmarshaller struct{}

func (u *DiagramProcessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramProcess), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramProcessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramProcess)
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
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "Process_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Process_Shapes, valueExpr, identifierMap)
	case "IsProcesssNodeExpanded":
		instance.IsProcesssNodeExpanded = GongExtractBool(valueExpr)
	case "ProcesssWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ProcesssWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Participant_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Participant_Shapes, valueExpr, identifierMap)
	case "IsParticipantsNodeExpanded":
		instance.IsParticipantsNodeExpanded = GongExtractBool(valueExpr)
	case "ParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalParticipant_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipant_Shapes, valueExpr, identifierMap)
	case "IsExternalParticipantsNodeExpanded":
		instance.IsExternalParticipantsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "TasksWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TasksWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Task_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Task_Shapes, valueExpr, identifierMap)
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
	case "AllocatedProcessesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.AllocatedProcessesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedProcessShapes":
		GongUnmarshallSliceOfPointers(&instance.AllocatedProcessShapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NoteTaskShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteTaskShapes, valueExpr, identifierMap)
	}
	return nil
}

type ExternalParticipantShapeUnmarshaller struct{}

func (u *ExternalParticipantShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ExternalParticipantShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ExternalParticipantShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExternalParticipantShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
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
	case "RootProcesses":
		GongUnmarshallSliceOfPointers(&instance.RootProcesses, valueExpr, identifierMap)
	case "IsProcessesNodeExpanded":
		instance.IsProcessesNodeExpanded = GongExtractBool(valueExpr)
	case "ProcesssWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ProcesssWhoseNodeIsExpanded, valueExpr, identifierMap)
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
	case "ParticipantsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParticipantsWhoseNodeIsExpanded, valueExpr, identifierMap)
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
	case "IsTasksNodeExpanded":
		instance.IsTasksNodeExpanded = GongExtractBool(valueExpr)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
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

type NoteTaskShapeUnmarshaller struct{}

func (u *NoteTaskShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteTaskShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteTaskShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteTaskShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Task":
		GongUnmarshallPointer(&instance.Task, valueExpr, identifierMap)
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

type ParticipantUnmarshaller struct{}

func (u *ParticipantUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Participant), stage, identifier, instanceName, preserveOrder)
}

func (u *ParticipantUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Participant)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsProcessResource":
		instance.IsProcessResource = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Resources":
		GongUnmarshallSliceOfPointers(&instance.Resources, valueExpr, identifierMap)
	case "IsResourcesNodeExpanded":
		instance.IsResourcesNodeExpanded = GongExtractBool(valueExpr)
	case "Processes":
		GongUnmarshallSliceOfPointers(&instance.Processes, valueExpr, identifierMap)
	case "IsProcessesNodeExpanded":
		instance.IsProcessesNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsTasksNodeExpanded":
		instance.IsTasksNodeExpanded = GongExtractBool(valueExpr)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
	case "IsControlFlowsNodeExpanded":
		instance.IsControlFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ControlFlows":
		GongUnmarshallSliceOfPointers(&instance.ControlFlows, valueExpr, identifierMap)
	case "TaskWhoseOutControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseOutControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "TaskWhoseInControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseInControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "TaskWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "TaskWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type ParticipantShapeUnmarshaller struct{}

func (u *ParticipantShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParticipantShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParticipantShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParticipantShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
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
	case "WidthWeight":
		instance.WidthWeight = GongExtractFloat(valueExpr)
	}
	return nil
}

type ProcessUnmarshaller struct{}

func (u *ProcessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Process), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Process)
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
	case "DiagramProcesss":
		GongUnmarshallSliceOfPointers(&instance.DiagramProcesss, valueExpr, identifierMap)
	case "DiagramProcessWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramProcessWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubProcessNodeExpanded":
		instance.IsSubProcessNodeExpanded = GongExtractBool(valueExpr)
	case "SubProcesses":
		GongUnmarshallSliceOfPointers(&instance.SubProcesses, valueExpr, identifierMap)
	case "Participants":
		GongUnmarshallSliceOfPointers(&instance.Participants, valueExpr, identifierMap)
	case "ParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DataFlows":
		GongUnmarshallSliceOfPointers(&instance.DataFlows, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalParticipants":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipants, valueExpr, identifierMap)
	case "ExternalParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type ProcessShapeUnmarshaller struct{}

func (u *ProcessShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcessShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcessShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcessShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Process":
		GongUnmarshallPointer(&instance.Process, valueExpr, identifierMap)
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

type TaskUnmarshaller struct{}

func (u *TaskUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Task), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Task)
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
	case "IsStartTask":
		instance.IsStartTask = GongExtractBool(valueExpr)
	case "IsEndTask":
		instance.IsEndTask = GongExtractBool(valueExpr)
	case "Type":
		GongUnmarshallPointer(&instance.Type, valueExpr, identifierMap)
	case "IsTaskNameNotProcessName":
		instance.IsTaskNameNotProcessName = GongExtractBool(valueExpr)
	}
	return nil
}

type TaskShapeUnmarshaller struct{}

func (u *TaskShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Task":
		GongUnmarshallPointer(&instance.Task, valueExpr, identifierMap)
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

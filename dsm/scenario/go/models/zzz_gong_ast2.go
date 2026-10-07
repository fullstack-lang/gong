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
type ActorStateUnmarshaller struct{}

func (u *ActorStateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorState), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorState)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "IsWithProbaility":
		instance.IsWithProbaility = GongExtractBool(valueExpr)
	case "Probability":
		GongUnmarshallEnum(&instance.Probability, valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ActorStateShapeUnmarshaller struct{}

func (u *ActorStateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorStateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorStateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ActorState":
		GongUnmarshallPointer(&instance.ActorState, valueExpr, identifierMap)
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

type ActorStateTransitionUnmarshaller struct{}

func (u *ActorStateTransitionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorStateTransition), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateTransitionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorStateTransition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartState":
		GongUnmarshallPointer(&instance.StartState, valueExpr, identifierMap)
	case "EndState":
		GongUnmarshallPointer(&instance.EndState, valueExpr, identifierMap)
	case "Justifications":
		GongUnmarshallSliceOfPointers(&instance.Justifications, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ActorStateTransitionShapeUnmarshaller struct{}

func (u *ActorStateTransitionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorStateTransitionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateTransitionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorStateTransitionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ActorStateTransition":
		GongUnmarshallPointer(&instance.ActorStateTransition, valueExpr, identifierMap)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type AnalysisUnmarshaller struct{}

func (u *AnalysisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Analysis), stage, identifier, instanceName, preserveOrder)
}

func (u *AnalysisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Analysis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Scenarios":
		GongUnmarshallSliceOfPointers(&instance.Scenarios, valueExpr, identifierMap)
	case "IsScenariosNodeExpanded":
		instance.IsScenariosNodeExpanded = GongExtractBool(valueExpr)
	case "GroupUse":
		GongUnmarshallSliceOfPointers(&instance.GroupUse, valueExpr, identifierMap)
	case "IsGroupUseNodeExpanded":
		instance.IsGroupUseNodeExpanded = GongExtractBool(valueExpr)
	case "GeoObjectUse":
		GongUnmarshallSliceOfPointers(&instance.GeoObjectUse, valueExpr, identifierMap)
	case "IsGeoObjectUseNodeExpanded":
		instance.IsGeoObjectUseNodeExpanded = GongExtractBool(valueExpr)
	case "MapUse":
		GongUnmarshallSliceOfPointers(&instance.MapUse, valueExpr, identifierMap)
	case "IsMapUseNodeExpanded":
		instance.IsMapUseNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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
	case "IsShowPrefix":
		instance.IsShowPrefix = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "EvolutionDirectionShapes":
		GongUnmarshallSliceOfPointers(&instance.EvolutionDirectionShapes, valueExpr, identifierMap)
	case "EvolutionDirectionsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EvolutionDirectionsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEvolutionDirectionsNodeExpanded":
		instance.IsEvolutionDirectionsNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStateShapes":
		GongUnmarshallSliceOfPointers(&instance.ActorStateShapes, valueExpr, identifierMap)
	case "ActorStatesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ActorStatesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsActorStatesNodeExpanded":
		instance.IsActorStatesNodeExpanded = GongExtractBool(valueExpr)
	case "ParameterShapes":
		GongUnmarshallSliceOfPointers(&instance.ParameterShapes, valueExpr, identifierMap)
	case "ParametersWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParametersWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsParametersNodeExpanded":
		instance.IsParametersNodeExpanded = GongExtractBool(valueExpr)
	case "ScenarioParameterShapes":
		GongUnmarshallSliceOfPointers(&instance.ScenarioParameterShapes, valueExpr, identifierMap)
	case "ParametersAggregatesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParametersAggregatesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsParametersAggregatesNodeExpanded":
		instance.IsParametersAggregatesNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStateTransitionShapes":
		GongUnmarshallSliceOfPointers(&instance.ActorStateTransitionShapes, valueExpr, identifierMap)
	case "ActorStateTransitionsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ActorStateTransitionsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsActorStateTransitionsNodeExpanded":
		instance.IsActorStateTransitionsNodeExpanded = GongExtractBool(valueExpr)
	case "AxisOrign_X":
		instance.AxisOrign_X = GongExtractFloat(valueExpr)
	case "AxisOrign_Y":
		instance.AxisOrign_Y = GongExtractFloat(valueExpr)
	case "VerticalAxis_Top_Y":
		instance.VerticalAxis_Top_Y = GongExtractFloat(valueExpr)
	case "VerticalAxis_Bottom_Y":
		instance.VerticalAxis_Bottom_Y = GongExtractFloat(valueExpr)
	case "VerticalAxis_StrokeWidth":
		instance.VerticalAxis_StrokeWidth = GongExtractFloat(valueExpr)
	case "HorizontalAxis_Right_X":
		instance.HorizontalAxis_Right_X = GongExtractFloat(valueExpr)
	case "Start":
		instance.Start = GongExtractDate(valueExpr)
	case "End":
		instance.End = GongExtractDate(valueExpr)
	case "NumberOfYearsBetweenTicks":
		instance.NumberOfYearsBetweenTicks = GongExtractInt(valueExpr)
	}
	return nil
}

type DocumentUnmarshaller struct{}

func (u *DocumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Document), stage, identifier, instanceName, preserveOrder)
}

func (u *DocumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Document)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GeoObjectUse":
		GongUnmarshallSliceOfPointers(&instance.GeoObjectUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type DocumentUseUnmarshaller struct{}

func (u *DocumentUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DocumentUse), stage, identifier, instanceName, preserveOrder)
}

func (u *DocumentUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DocumentUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Document":
		GongUnmarshallPointer(&instance.Document, valueExpr, identifierMap)
	}
	return nil
}

type EvolutionDirectionUnmarshaller struct{}

func (u *EvolutionDirectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EvolutionDirection), stage, identifier, instanceName, preserveOrder)
}

func (u *EvolutionDirectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EvolutionDirection)
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

type EvolutionDirectionShapeUnmarshaller struct{}

func (u *EvolutionDirectionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EvolutionDirectionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EvolutionDirectionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EvolutionDirectionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EvolutionDirection":
		GongUnmarshallPointer(&instance.EvolutionDirection, valueExpr, identifierMap)
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

type FooUnmarshaller struct{}

func (u *FooUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Foo), stage, identifier, instanceName, preserveOrder)
}

func (u *FooUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Foo)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GeoObjectUnmarshaller struct{}

func (u *GeoObjectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoObject), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoObjectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoObject)
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

type GeoObjectUseUnmarshaller struct{}

func (u *GeoObjectUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoObjectUse), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoObjectUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoObjectUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GeoObject":
		GongUnmarshallPointer(&instance.GeoObject, valueExpr, identifierMap)
	}
	return nil
}

type GroupUnmarshaller struct{}

func (u *GroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "UserUse":
		GongUnmarshallSliceOfPointers(&instance.UserUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GroupUseUnmarshaller struct{}

func (u *GroupUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroupUse), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroupUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Group":
		GongUnmarshallPointer(&instance.Group, valueExpr, identifierMap)
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
	case "Analyses":
		GongUnmarshallSliceOfPointers(&instance.Analyses, valueExpr, identifierMap)
	case "IsAnalysesNodeExpanded":
		instance.IsAnalysesNodeExpanded = GongExtractBool(valueExpr)
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

type MapObjectUnmarshaller struct{}

func (u *MapObjectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MapObject), stage, identifier, instanceName, preserveOrder)
}

func (u *MapObjectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MapObject)
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

type MapObjectUseUnmarshaller struct{}

func (u *MapObjectUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MapObjectUse), stage, identifier, instanceName, preserveOrder)
}

func (u *MapObjectUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MapObjectUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Map":
		GongUnmarshallPointer(&instance.Map, valueExpr, identifierMap)
	}
	return nil
}

type ParameterUnmarshaller struct{}

func (u *ParameterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Parameter), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Parameter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "IsResponse":
		instance.IsResponse = GongExtractBool(valueExpr)
	case "Start":
		instance.Start = GongExtractDate(valueExpr)
	case "End":
		instance.End = GongExtractDate(valueExpr)
	case "Force":
		instance.Force = GongExtractFloat(valueExpr)
	case "GroupUse":
		GongUnmarshallSliceOfPointers(&instance.GroupUse, valueExpr, identifierMap)
	case "DocumentUse":
		GongUnmarshallSliceOfPointers(&instance.DocumentUse, valueExpr, identifierMap)
	case "GeoObjectUse":
		GongUnmarshallSliceOfPointers(&instance.GeoObjectUse, valueExpr, identifierMap)
	case "Tag":
		instance.Tag = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ParameterCategoryUnmarshaller struct{}

func (u *ParameterCategoryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParameterCategory), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterCategoryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParameterCategory)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ParameterUse":
		GongUnmarshallSliceOfPointers(&instance.ParameterUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ParameterCategoryUseUnmarshaller struct{}

func (u *ParameterCategoryUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParameterCategoryUse), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterCategoryUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParameterCategoryUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ParameterCategory":
		GongUnmarshallPointer(&instance.ParameterCategory, valueExpr, identifierMap)
	}
	return nil
}

type ParameterShapeUnmarshaller struct{}

func (u *ParameterShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParameterShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParameterShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parameter":
		GongUnmarshallPointer(&instance.Parameter, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "ShapeIsComputedFromModel":
		instance.ShapeIsComputedFromModel = GongExtractBool(valueExpr)
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

type ParametersAggregateUnmarshaller struct{}

func (u *ParametersAggregateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParametersAggregate), stage, identifier, instanceName, preserveOrder)
}

func (u *ParametersAggregateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParametersAggregate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tag":
		instance.Tag = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Parameters":
		GongUnmarshallSliceOfPointers(&instance.Parameters, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ParametersAggregateShapeUnmarshaller struct{}

func (u *ParametersAggregateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParametersAggregateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParametersAggregateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParametersAggregateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ScenarioParameter":
		GongUnmarshallPointer(&instance.ScenarioParameter, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
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

type PositionUnmarshaller struct{}

func (u *PositionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Position), stage, identifier, instanceName, preserveOrder)
}

func (u *PositionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Position)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	case "Ordinate":
		instance.Ordinate = GongExtractFloat(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type RepositoryUnmarshaller struct{}

func (u *RepositoryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Repository), stage, identifier, instanceName, preserveOrder)
}

func (u *RepositoryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Repository)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ParameterUse":
		GongUnmarshallSliceOfPointers(&instance.ParameterUse, valueExpr, identifierMap)
	case "GroupUse":
		GongUnmarshallSliceOfPointers(&instance.GroupUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ScenarioUnmarshaller struct{}

func (u *ScenarioUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Scenario), stage, identifier, instanceName, preserveOrder)
}

func (u *ScenarioUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scenario)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "IsDiagramsNodeExpanded":
		instance.IsDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStates":
		GongUnmarshallSliceOfPointers(&instance.ActorStates, valueExpr, identifierMap)
	case "IsActorStatesNodeExpanded":
		instance.IsActorStatesNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStateTransitions":
		GongUnmarshallSliceOfPointers(&instance.ActorStateTransitions, valueExpr, identifierMap)
	case "IsActorStateTransitionsNodeExpanded":
		instance.IsActorStateTransitionsNodeExpanded = GongExtractBool(valueExpr)
	case "EvolutionDirections":
		GongUnmarshallSliceOfPointers(&instance.EvolutionDirections, valueExpr, identifierMap)
	case "IsEvolutionDirectionsNodeExpanded":
		instance.IsEvolutionDirectionsNodeExpanded = GongExtractBool(valueExpr)
	case "Parameters":
		GongUnmarshallSliceOfPointers(&instance.Parameters, valueExpr, identifierMap)
	case "IsParametersNodeExpanded":
		instance.IsParametersNodeExpanded = GongExtractBool(valueExpr)
	case "ParametersAggretates":
		GongUnmarshallSliceOfPointers(&instance.ParametersAggretates, valueExpr, identifierMap)
	case "IsParametersAggretatesNodeExpanded":
		instance.IsParametersAggretatesNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type UserUnmarshaller struct{}

func (u *UserUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(User), stage, identifier, instanceName, preserveOrder)
}

func (u *UserUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*User)
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

type UserUseUnmarshaller struct{}

func (u *UserUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UserUse), stage, identifier, instanceName, preserveOrder)
}

func (u *UserUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UserUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "User":
		GongUnmarshallPointer(&instance.User, valueExpr, identifierMap)
	}
	return nil
}

type WorkspaceUnmarshaller struct{}

func (u *WorkspaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Workspace), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkspaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Workspace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SelectedDiagram":
		GongUnmarshallPointer(&instance.SelectedDiagram, valueExpr, identifierMap)
	case "Default_EvolutionDirectionShape":
		GongUnmarshallPointer(&instance.Default_EvolutionDirectionShape, valueExpr, identifierMap)
	case "Default_ParameterShape":
		GongUnmarshallPointer(&instance.Default_ParameterShape, valueExpr, identifierMap)
	case "Default_ScenarioParameterShape":
		GongUnmarshallPointer(&instance.Default_ScenarioParameterShape, valueExpr, identifierMap)
	case "Default_ActorStateShape":
		GongUnmarshallPointer(&instance.Default_ActorStateShape, valueExpr, identifierMap)
	case "Default_ActorStateTransitionShape":
		GongUnmarshallPointer(&instance.Default_ActorStateTransitionShape, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

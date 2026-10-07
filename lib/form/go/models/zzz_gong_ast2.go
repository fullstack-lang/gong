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
type CheckBoxUnmarshaller struct{}

func (u *CheckBoxUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckBox), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckBoxUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckBox)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractBool(valueExpr)
	}
	return nil
}

type FormDivUnmarshaller struct{}

func (u *FormDivUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormDiv), stage, identifier, instanceName, preserveOrder)
}

func (u *FormDivUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormDiv)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FormFields":
		GongUnmarshallSliceOfPointers(&instance.FormFields, valueExpr, identifierMap)
	case "CheckBoxs":
		GongUnmarshallSliceOfPointers(&instance.CheckBoxs, valueExpr, identifierMap)
	case "FormEditAssocButton":
		GongUnmarshallPointer(&instance.FormEditAssocButton, valueExpr, identifierMap)
	case "FormSortAssocButton":
		GongUnmarshallPointer(&instance.FormSortAssocButton, valueExpr, identifierMap)
	case "IsADivider":
		instance.IsADivider = GongExtractBool(valueExpr)
	case "IsAStartAccordionGroup":
		instance.IsAStartAccordionGroup = GongExtractBool(valueExpr)
	case "AccordionGroupName":
		instance.AccordionGroupName = GongExtractString(valueExpr)
	case "IsAEndAccordionGroup":
		instance.IsAEndAccordionGroup = GongExtractBool(valueExpr)
	}
	return nil
}

type FormEditAssocButtonUnmarshaller struct{}

func (u *FormEditAssocButtonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormEditAssocButton), stage, identifier, instanceName, preserveOrder)
}

func (u *FormEditAssocButtonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormEditAssocButton)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "AssociationStorage":
		instance.AssociationStorage = GongExtractString(valueExpr)
	case "HasChanged":
		instance.HasChanged = GongExtractBool(valueExpr)
	case "IsForSavePurpose":
		instance.IsForSavePurpose = GongExtractBool(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "MatTooltipShowDelay":
		instance.MatTooltipShowDelay = GongExtractString(valueExpr)
	}
	return nil
}

type FormFieldUnmarshaller struct{}

func (u *FormFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormField), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "InputTypeEnum":
		GongUnmarshallEnum(&instance.InputTypeEnum, valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "Placeholder":
		instance.Placeholder = GongExtractString(valueExpr)
	case "FormFieldString":
		GongUnmarshallPointer(&instance.FormFieldString, valueExpr, identifierMap)
	case "FormFieldFloat64":
		GongUnmarshallPointer(&instance.FormFieldFloat64, valueExpr, identifierMap)
	case "FormFieldInt":
		GongUnmarshallPointer(&instance.FormFieldInt, valueExpr, identifierMap)
	case "FormFieldDate":
		GongUnmarshallPointer(&instance.FormFieldDate, valueExpr, identifierMap)
	case "FormFieldTime":
		GongUnmarshallPointer(&instance.FormFieldTime, valueExpr, identifierMap)
	case "FormFieldDateTime":
		GongUnmarshallPointer(&instance.FormFieldDateTime, valueExpr, identifierMap)
	case "FormFieldSelect":
		GongUnmarshallPointer(&instance.FormFieldSelect, valueExpr, identifierMap)
	case "HasBespokeWidth":
		instance.HasBespokeWidth = GongExtractBool(valueExpr)
	case "BespokeWidthPx":
		instance.BespokeWidthPx = GongExtractInt(valueExpr)
	case "HasBespokeHeight":
		instance.HasBespokeHeight = GongExtractBool(valueExpr)
	case "BespokeHeightPx":
		instance.BespokeHeightPx = GongExtractInt(valueExpr)
	}
	return nil
}

type FormFieldDateUnmarshaller struct{}

func (u *FormFieldDateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldDate), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldDateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldDate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractDate(valueExpr)
	}
	return nil
}

type FormFieldDateTimeUnmarshaller struct{}

func (u *FormFieldDateTimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldDateTime), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldDateTimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldDateTime)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractDate(valueExpr)
	}
	return nil
}

type FormFieldFloat64Unmarshaller struct{}

func (u *FormFieldFloat64Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldFloat64), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldFloat64Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldFloat64)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractFloat(valueExpr)
	case "HasMinValidator":
		instance.HasMinValidator = GongExtractBool(valueExpr)
	case "MinValue":
		instance.MinValue = GongExtractFloat(valueExpr)
	case "HasMaxValidator":
		instance.HasMaxValidator = GongExtractBool(valueExpr)
	case "MaxValue":
		instance.MaxValue = GongExtractFloat(valueExpr)
	}
	return nil
}

type FormFieldIntUnmarshaller struct{}

func (u *FormFieldIntUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldInt), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldIntUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldInt)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractInt(valueExpr)
	case "HasMinValidator":
		instance.HasMinValidator = GongExtractBool(valueExpr)
	case "MinValue":
		instance.MinValue = GongExtractInt(valueExpr)
	case "HasMaxValidator":
		instance.HasMaxValidator = GongExtractBool(valueExpr)
	case "MaxValue":
		instance.MaxValue = GongExtractInt(valueExpr)
	}
	return nil
}

type FormFieldSelectUnmarshaller struct{}

func (u *FormFieldSelectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldSelect), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldSelectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldSelect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		GongUnmarshallPointer(&instance.Value, valueExpr, identifierMap)
	case "Options":
		GongUnmarshallSliceOfPointers(&instance.Options, valueExpr, identifierMap)
	case "CanBeEmpty":
		instance.CanBeEmpty = GongExtractBool(valueExpr)
	case "PreserveInitialOrder":
		instance.PreserveInitialOrder = GongExtractBool(valueExpr)
	}
	return nil
}

type FormFieldStringUnmarshaller struct{}

func (u *FormFieldStringUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldString), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldStringUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldString)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	case "IsTextArea":
		instance.IsTextArea = GongExtractBool(valueExpr)
	}
	return nil
}

type FormFieldTimeUnmarshaller struct{}

func (u *FormFieldTimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormFieldTime), stage, identifier, instanceName, preserveOrder)
}

func (u *FormFieldTimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormFieldTime)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractDate(valueExpr)
	case "Step":
		instance.Step = GongExtractFloat(valueExpr)
	}
	return nil
}

type FormGroupUnmarshaller struct{}

func (u *FormGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormGroup), stage, identifier, instanceName, preserveOrder)
}

func (u *FormGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "TypeLabel":
		instance.TypeLabel = GongExtractString(valueExpr)
	case "FormDivs":
		GongUnmarshallSliceOfPointers(&instance.FormDivs, valueExpr, identifierMap)
	case "HasSuppressButton":
		instance.HasSuppressButton = GongExtractBool(valueExpr)
	case "HasSuppressButtonBeenPressed":
		instance.HasSuppressButtonBeenPressed = GongExtractBool(valueExpr)
	}
	return nil
}

type FormSortAssocButtonUnmarshaller struct{}

func (u *FormSortAssocButtonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FormSortAssocButton), stage, identifier, instanceName, preserveOrder)
}

func (u *FormSortAssocButtonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FormSortAssocButton)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Label":
		instance.Label = GongExtractString(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "MatTooltipShowDelay":
		instance.MatTooltipShowDelay = GongExtractString(valueExpr)
	case "FormEditAssocButton":
		GongUnmarshallPointer(&instance.FormEditAssocButton, valueExpr, identifierMap)
	}
	return nil
}

type OptionUnmarshaller struct{}

func (u *OptionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Option), stage, identifier, instanceName, preserveOrder)
}

func (u *OptionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Option)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

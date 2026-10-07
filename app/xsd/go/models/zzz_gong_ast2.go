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
type AllUnmarshaller struct{}

func (u *AllUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(All), stage, identifier, instanceName, preserveOrder)
}

func (u *AllUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*All)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type AnnotationUnmarshaller struct{}

func (u *AnnotationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Annotation), stage, identifier, instanceName, preserveOrder)
}

func (u *AnnotationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Annotation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Documentations":
		GongUnmarshallSliceOfPointers(&instance.Documentations, valueExpr, identifierMap)
	}
	return nil
}

type AttributeUnmarshaller struct{}

func (u *AttributeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Attribute), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Attribute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "Default":
		instance.Default = GongExtractString(valueExpr)
	case "Use":
		instance.Use = GongExtractString(valueExpr)
	case "Form":
		instance.Form = GongExtractString(valueExpr)
	case "Fixed":
		instance.Fixed = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "TargetNamespace":
		instance.TargetNamespace = GongExtractString(valueExpr)
	case "SimpleType":
		instance.SimpleType = GongExtractString(valueExpr)
	case "IDXSD":
		instance.IDXSD = GongExtractString(valueExpr)
	}
	return nil
}

type AttributeGroupUnmarshaller struct{}

func (u *AttributeGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AttributeGroup), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributeGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributeGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type ChoiceUnmarshaller struct{}

func (u *ChoiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Choice), stage, identifier, instanceName, preserveOrder)
}

func (u *ChoiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Choice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type ComplexContentUnmarshaller struct{}

func (u *ComplexContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ComplexContent), stage, identifier, instanceName, preserveOrder)
}

func (u *ComplexContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ComplexContent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ComplexTypeUnmarshaller struct{}

func (u *ComplexTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ComplexType), stage, identifier, instanceName, preserveOrder)
}

func (u *ComplexTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ComplexType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "IsAnonymous":
		instance.IsAnonymous = GongExtractBool(valueExpr)
	case "OuterElement":
		GongUnmarshallPointer(&instance.OuterElement, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallPointer(&instance.Extension, valueExpr, identifierMap)
	case "SimpleContent":
		GongUnmarshallPointer(&instance.SimpleContent, valueExpr, identifierMap)
	case "ComplexContent":
		GongUnmarshallPointer(&instance.ComplexContent, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type DocumentationUnmarshaller struct{}

func (u *DocumentationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Documentation), stage, identifier, instanceName, preserveOrder)
}

func (u *DocumentationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Documentation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

type ElementUnmarshaller struct{}

func (u *ElementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Element), stage, identifier, instanceName, preserveOrder)
}

func (u *ElementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Element)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Default":
		instance.Default = GongExtractString(valueExpr)
	case "Fixed":
		instance.Fixed = GongExtractString(valueExpr)
	case "Nillable":
		instance.Nillable = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Abstract":
		instance.Abstract = GongExtractString(valueExpr)
	case "Form":
		instance.Form = GongExtractString(valueExpr)
	case "Block":
		instance.Block = GongExtractString(valueExpr)
	case "Final":
		instance.Final = GongExtractString(valueExpr)
	case "SimpleType":
		GongUnmarshallPointer(&instance.SimpleType, valueExpr, identifierMap)
	case "ComplexType":
		GongUnmarshallPointer(&instance.ComplexType, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type EnumerationUnmarshaller struct{}

func (u *EnumerationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Enumeration), stage, identifier, instanceName, preserveOrder)
}

func (u *EnumerationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Enumeration)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type ExtensionUnmarshaller struct{}

func (u *ExtensionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Extension), stage, identifier, instanceName, preserveOrder)
}

func (u *ExtensionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extension)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Base":
		instance.Base = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
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
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "IsAnonymous":
		instance.IsAnonymous = GongExtractBool(valueExpr)
	case "OuterElement":
		GongUnmarshallPointer(&instance.OuterElement, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type LengthUnmarshaller struct{}

func (u *LengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Length), stage, identifier, instanceName, preserveOrder)
}

func (u *LengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Length)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MaxInclusiveUnmarshaller struct{}

func (u *MaxInclusiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MaxInclusive), stage, identifier, instanceName, preserveOrder)
}

func (u *MaxInclusiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MaxInclusive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MaxLengthUnmarshaller struct{}

func (u *MaxLengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MaxLength), stage, identifier, instanceName, preserveOrder)
}

func (u *MaxLengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MaxLength)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MinInclusiveUnmarshaller struct{}

func (u *MinInclusiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MinInclusive), stage, identifier, instanceName, preserveOrder)
}

func (u *MinInclusiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinInclusive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MinLengthUnmarshaller struct{}

func (u *MinLengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MinLength), stage, identifier, instanceName, preserveOrder)
}

func (u *MinLengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinLength)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PatternUnmarshaller struct{}

func (u *PatternUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pattern), stage, identifier, instanceName, preserveOrder)
}

func (u *PatternUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pattern)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type RestrictionUnmarshaller struct{}

func (u *RestrictionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Restriction), stage, identifier, instanceName, preserveOrder)
}

func (u *RestrictionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Restriction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Base":
		instance.Base = GongExtractString(valueExpr)
	case "Enumerations":
		GongUnmarshallSliceOfPointers(&instance.Enumerations, valueExpr, identifierMap)
	case "MinInclusive":
		GongUnmarshallPointer(&instance.MinInclusive, valueExpr, identifierMap)
	case "MaxInclusive":
		GongUnmarshallPointer(&instance.MaxInclusive, valueExpr, identifierMap)
	case "Pattern":
		GongUnmarshallPointer(&instance.Pattern, valueExpr, identifierMap)
	case "WhiteSpace":
		GongUnmarshallPointer(&instance.WhiteSpace, valueExpr, identifierMap)
	case "MinLength":
		GongUnmarshallPointer(&instance.MinLength, valueExpr, identifierMap)
	case "MaxLength":
		GongUnmarshallPointer(&instance.MaxLength, valueExpr, identifierMap)
	case "Length":
		GongUnmarshallPointer(&instance.Length, valueExpr, identifierMap)
	case "TotalDigit":
		GongUnmarshallPointer(&instance.TotalDigit, valueExpr, identifierMap)
	}
	return nil
}

type SchemaUnmarshaller struct{}

func (u *SchemaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Schema), stage, identifier, instanceName, preserveOrder)
}

func (u *SchemaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Schema)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Xs":
		instance.Xs = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "SimpleTypes":
		GongUnmarshallSliceOfPointers(&instance.SimpleTypes, valueExpr, identifierMap)
	case "ComplexTypes":
		GongUnmarshallSliceOfPointers(&instance.ComplexTypes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type SequenceUnmarshaller struct{}

func (u *SequenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Sequence), stage, identifier, instanceName, preserveOrder)
}

func (u *SequenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sequence)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type SimpleContentUnmarshaller struct{}

func (u *SimpleContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SimpleContent), stage, identifier, instanceName, preserveOrder)
}

func (u *SimpleContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SimpleContent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallPointer(&instance.Extension, valueExpr, identifierMap)
	case "Restriction":
		GongUnmarshallPointer(&instance.Restriction, valueExpr, identifierMap)
	}
	return nil
}

type SimpleTypeUnmarshaller struct{}

func (u *SimpleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SimpleType), stage, identifier, instanceName, preserveOrder)
}

func (u *SimpleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SimpleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Restriction":
		GongUnmarshallPointer(&instance.Restriction, valueExpr, identifierMap)
	case "Union":
		GongUnmarshallPointer(&instance.Union, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type TotalDigitUnmarshaller struct{}

func (u *TotalDigitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TotalDigit), stage, identifier, instanceName, preserveOrder)
}

func (u *TotalDigitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TotalDigit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type UnionUnmarshaller struct{}

func (u *UnionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Union), stage, identifier, instanceName, preserveOrder)
}

func (u *UnionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Union)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "MemberTypes":
		instance.MemberTypes = GongExtractString(valueExpr)
	}
	return nil
}

type WhiteSpaceUnmarshaller struct{}

func (u *WhiteSpaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WhiteSpace), stage, identifier, instanceName, preserveOrder)
}

func (u *WhiteSpaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WhiteSpace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

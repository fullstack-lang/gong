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
type AttributeShapeUnmarshaller struct{}

func (u *AttributeShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AttributeShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributeShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributeShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "FieldTypeAsString":
		instance.FieldTypeAsString = GongExtractString(valueExpr)
	case "Structname":
		instance.Structname = GongExtractString(valueExpr)
	case "Fieldtypename":
		instance.Fieldtypename = GongExtractString(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type ClassdiagramUnmarshaller struct{}

func (u *ClassdiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Classdiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *ClassdiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Classdiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "IsIncludedInStaticWebSite":
		instance.IsIncludedInStaticWebSite = GongExtractBool(valueExpr)
	case "GongStructShapes":
		GongUnmarshallSliceOfPointers(&instance.GongStructShapes, valueExpr, identifierMap)
	case "GongEnumShapes":
		GongUnmarshallSliceOfPointers(&instance.GongEnumShapes, valueExpr, identifierMap)
	case "GongNoteShapes":
		GongUnmarshallSliceOfPointers(&instance.GongNoteShapes, valueExpr, identifierMap)
	case "ShowNbInstances":
		instance.ShowNbInstances = GongExtractBool(valueExpr)
	case "ShowMultiplicity":
		instance.ShowMultiplicity = GongExtractBool(valueExpr)
	case "ShowLinkNames":
		instance.ShowLinkNames = GongExtractBool(valueExpr)
	case "IsInRenameMode":
		instance.IsInRenameMode = GongExtractBool(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "NodeGongStructsIsExpanded":
		instance.NodeGongStructsIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongStructNodeExpansion":
		instance.NodeGongStructNodeExpansion = GongExtractString(valueExpr)
	case "NodeGongEnumsIsExpanded":
		instance.NodeGongEnumsIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongEnumNodeExpansion":
		instance.NodeGongEnumNodeExpansion = GongExtractString(valueExpr)
	case "NodeGongNotesIsExpanded":
		instance.NodeGongNotesIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongNoteNodeExpansion":
		instance.NodeGongNoteNodeExpansion = GongExtractString(valueExpr)
	}
	return nil
}

type DiagramPackageUnmarshaller struct{}

func (u *DiagramPackageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramPackage), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramPackageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramPackage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Path":
		instance.Path = GongExtractString(valueExpr)
	case "GongModelPath":
		instance.GongModelPath = GongExtractString(valueExpr)
	case "Classdiagrams":
		GongUnmarshallSliceOfPointers(&instance.Classdiagrams, valueExpr, identifierMap)
	case "SelectedClassdiagram":
		GongUnmarshallPointer(&instance.SelectedClassdiagram, valueExpr, identifierMap)
	case "AbsolutePathToDiagramPackage":
		instance.AbsolutePathToDiagramPackage = GongExtractString(valueExpr)
	}
	return nil
}

type GongEnumShapeUnmarshaller struct{}

func (u *GongEnumShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnumShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "GongEnumValueShapes":
		GongUnmarshallSliceOfPointers(&instance.GongEnumValueShapes, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GongEnumValueShapeUnmarshaller struct{}

func (u *GongEnumValueShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnumValueShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumValueShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumValueShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type GongNoteLinkShapeUnmarshaller struct{}

func (u *GongNoteLinkShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongNoteLinkShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongNoteLinkShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNoteLinkShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Identifier":
		instance.Identifier = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type GongNoteShapeUnmarshaller struct{}

func (u *GongNoteShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongNoteShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongNoteShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNoteShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Identifier":
		instance.Identifier = GongExtractString(valueExpr)
	case "Body":
		instance.Body = GongExtractString(valueExpr)
	case "BodyHTML":
		instance.BodyHTML = GongExtractString(valueExpr)
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
	case "Matched":
		instance.Matched = GongExtractBool(valueExpr)
	case "GongNoteLinkShapes":
		GongUnmarshallSliceOfPointers(&instance.GongNoteLinkShapes, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GongStructShapeUnmarshaller struct{}

func (u *GongStructShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongStructShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GongStructShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongStructShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "AttributeShapes":
		GongUnmarshallSliceOfPointers(&instance.AttributeShapes, valueExpr, identifierMap)
	case "LinkShapes":
		GongUnmarshallSliceOfPointers(&instance.LinkShapes, valueExpr, identifierMap)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	}
	return nil
}

type LinkShapeUnmarshaller struct{}

func (u *LinkShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinkShape), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IdentifierMeta":
		instance.IdentifierMeta = GongExtractExpr(valueExpr)
	case "FieldTypeIdentifierMeta":
		instance.FieldTypeIdentifierMeta = GongExtractExpr(valueExpr)
	case "FieldOffsetX":
		instance.FieldOffsetX = GongExtractFloat(valueExpr)
	case "FieldOffsetY":
		instance.FieldOffsetY = GongExtractFloat(valueExpr)
	case "TargetMultiplicity":
		GongUnmarshallEnum(&instance.TargetMultiplicity, valueExpr)
	case "TargetMultiplicityOffsetX":
		instance.TargetMultiplicityOffsetX = GongExtractFloat(valueExpr)
	case "TargetMultiplicityOffsetY":
		instance.TargetMultiplicityOffsetY = GongExtractFloat(valueExpr)
	case "SourceMultiplicity":
		GongUnmarshallEnum(&instance.SourceMultiplicity, valueExpr)
	case "SourceMultiplicityOffsetX":
		instance.SourceMultiplicityOffsetX = GongExtractFloat(valueExpr)
	case "SourceMultiplicityOffsetY":
		instance.SourceMultiplicityOffsetY = GongExtractFloat(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

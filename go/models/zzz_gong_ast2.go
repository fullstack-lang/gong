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
type GongBasicFieldUnmarshaller struct{}

func (u *GongBasicFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongBasicField), stage, identifier, instanceName, preserveOrder)
}

func (u *GongBasicFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongBasicField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BasicKindName":
		instance.BasicKindName = GongExtractString(valueExpr)
	case "GongEnum":
		GongUnmarshallPointer(&instance.GongEnum, valueExpr, identifierMap)
	case "DeclaredType":
		instance.DeclaredType = GongExtractString(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "IsTextArea":
		instance.IsTextArea = GongExtractBool(valueExpr)
	case "IsBespokeWidth":
		instance.IsBespokeWidth = GongExtractBool(valueExpr)
	case "BespokeWidth":
		instance.BespokeWidth = GongExtractInt(valueExpr)
	case "IsBespokeHeight":
		instance.IsBespokeHeight = GongExtractBool(valueExpr)
	case "BespokeHeight":
		instance.BespokeHeight = GongExtractInt(valueExpr)
	}
	return nil
}

type GongEnumUnmarshaller struct{}

func (u *GongEnumUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnum), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnum)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "GongEnumValues":
		GongUnmarshallSliceOfPointers(&instance.GongEnumValues, valueExpr, identifierMap)
	}
	return nil
}

type GongEnumValueUnmarshaller struct{}

func (u *GongEnumValueUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnumValue), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumValueUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumValue)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type GongLinkUnmarshaller struct{}

func (u *GongLinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongLink), stage, identifier, instanceName, preserveOrder)
}

func (u *GongLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Recv":
		instance.Recv = GongExtractString(valueExpr)
	case "ImportPath":
		instance.ImportPath = GongExtractString(valueExpr)
	}
	return nil
}

type GongNoteUnmarshaller struct{}

func (u *GongNoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongNote), stage, identifier, instanceName, preserveOrder)
}

func (u *GongNoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNote)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Body":
		instance.Body = GongExtractString(valueExpr)
	case "BodyHTML":
		instance.BodyHTML = GongExtractString(valueExpr)
	case "Links":
		GongUnmarshallSliceOfPointers(&instance.Links, valueExpr, identifierMap)
	}
	return nil
}

type GongStructUnmarshaller struct{}

func (u *GongStructUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongStruct), stage, identifier, instanceName, preserveOrder)
}

func (u *GongStructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongStruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GongBasicFields":
		GongUnmarshallSliceOfPointers(&instance.GongBasicFields, valueExpr, identifierMap)
	case "GongTimeFields":
		GongUnmarshallSliceOfPointers(&instance.GongTimeFields, valueExpr, identifierMap)
	case "PointerToGongStructFields":
		GongUnmarshallSliceOfPointers(&instance.PointerToGongStructFields, valueExpr, identifierMap)
	case "SliceOfPointerToGongStructFields":
		GongUnmarshallSliceOfPointers(&instance.SliceOfPointerToGongStructFields, valueExpr, identifierMap)
	case "HasOnAfterUpdateSignature":
		instance.HasOnAfterUpdateSignature = GongExtractBool(valueExpr)
	case "IsIgnoredForFront":
		instance.IsIgnoredForFront = GongExtractBool(valueExpr)
	case "IsOmittedForMarshalling":
		instance.IsOmittedForMarshalling = GongExtractBool(valueExpr)
	}
	return nil
}

type GongTimeFieldUnmarshaller struct{}

func (u *GongTimeFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongTimeField), stage, identifier, instanceName, preserveOrder)
}

func (u *GongTimeFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongTimeField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	case "BespokeTimeFormat":
		instance.BespokeTimeFormat = GongExtractString(valueExpr)
	case "TimeFormOnly":
		instance.TimeFormOnly = GongExtractBool(valueExpr)
	}
	return nil
}

type MetaReferenceUnmarshaller struct{}

func (u *MetaReferenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MetaReference), stage, identifier, instanceName, preserveOrder)
}

func (u *MetaReferenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MetaReference)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ModelPkgUnmarshaller struct{}

func (u *ModelPkgUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ModelPkg), stage, identifier, instanceName, preserveOrder)
}

func (u *ModelPkgUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ModelPkg)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PkgGoName":
		instance.PkgGoName = GongExtractString(valueExpr)
	case "PkgPath":
		instance.PkgPath = GongExtractString(valueExpr)
	case "PathToGoSubDirectory":
		instance.PathToGoSubDirectory = GongExtractString(valueExpr)
	case "OrmPkgGenPath":
		instance.OrmPkgGenPath = GongExtractString(valueExpr)
	case "DbOrmPkgGenPath":
		instance.DbOrmPkgGenPath = GongExtractString(valueExpr)
	case "DbLiteOrmPkgGenPath":
		instance.DbLiteOrmPkgGenPath = GongExtractString(valueExpr)
	case "DbPkgGenPath":
		instance.DbPkgGenPath = GongExtractString(valueExpr)
	case "ControllersPkgGenPath":
		instance.ControllersPkgGenPath = GongExtractString(valueExpr)
	case "FullstackPkgGenPath":
		instance.FullstackPkgGenPath = GongExtractString(valueExpr)
	case "StackPkgGenPath":
		instance.StackPkgGenPath = GongExtractString(valueExpr)
	case "Level1StackPkgGenPath":
		instance.Level1StackPkgGenPath = GongExtractString(valueExpr)
	case "StaticPkgGenPath":
		instance.StaticPkgGenPath = GongExtractString(valueExpr)
	case "ProbePkgGenPath":
		instance.ProbePkgGenPath = GongExtractString(valueExpr)
	case "NgWorkspacePath":
		instance.NgWorkspacePath = GongExtractString(valueExpr)
	case "NgWorkspaceName":
		instance.NgWorkspaceName = GongExtractString(valueExpr)
	case "NgDataLibrarySourceCodeDirectory":
		instance.NgDataLibrarySourceCodeDirectory = GongExtractString(valueExpr)
	case "NgSpecificLibrarySourceCodeDirectory":
		instance.NgSpecificLibrarySourceCodeDirectory = GongExtractString(valueExpr)
	case "MaterialLibDatamodelTargetPath":
		instance.MaterialLibDatamodelTargetPath = GongExtractString(valueExpr)
	}
	return nil
}

type PointerToGongStructFieldUnmarshaller struct{}

func (u *PointerToGongStructFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointerToGongStructField), stage, identifier, instanceName, preserveOrder)
}

func (u *PointerToGongStructFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointerToGongStructField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GongStruct":
		GongUnmarshallPointer(&instance.GongStruct, valueExpr, identifierMap)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	case "IsType":
		instance.IsType = GongExtractBool(valueExpr)
	}
	return nil
}

type SliceOfPointerToGongStructFieldUnmarshaller struct{}

func (u *SliceOfPointerToGongStructFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SliceOfPointerToGongStructField), stage, identifier, instanceName, preserveOrder)
}

func (u *SliceOfPointerToGongStructFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SliceOfPointerToGongStructField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GongStruct":
		GongUnmarshallPointer(&instance.GongStruct, valueExpr, identifierMap)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	}
	return nil
}

type StageSetFieldUnmarshaller struct{}

func (u *StageSetFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StageSetField), stage, identifier, instanceName, preserveOrder)
}

func (u *StageSetFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StageSetField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PackageName":
		instance.PackageName = GongExtractString(valueExpr)
	case "PackagePath":
		instance.PackagePath = GongExtractString(valueExpr)
	case "IsLocal":
		instance.IsLocal = GongExtractBool(valueExpr)
	case "ImportAlias":
		instance.ImportAlias = GongExtractString(valueExpr)
	}
	return nil
}

type StageSetModelUnmarshaller struct{}

func (u *StageSetModelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StageSetModel), stage, identifier, instanceName, preserveOrder)
}

func (u *StageSetModelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StageSetModel)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Fields":
		GongUnmarshallSliceOfPointers(&instance.Fields, valueExpr, identifierMap)
	case "IsManual":
		instance.IsManual = GongExtractBool(valueExpr)
	}
	return nil
}

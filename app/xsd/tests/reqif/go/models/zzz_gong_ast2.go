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
type ALTERNATIVE_IDUnmarshaller struct{}

func (u *ALTERNATIVE_IDUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ALTERNATIVE_ID), stage, identifier, instanceName, preserveOrder)
}

func (u *ALTERNATIVE_IDUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ALTERNATIVE_ID)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_BOOLEANUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_DATEUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_ENUMERATIONUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MULTI_VALUED":
		instance.MULTI_VALUED = GongExtractBool(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_INTEGERUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_REALUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_STRINGUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_XHTMLUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_BOOLEANUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractBool(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_DATEUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_INTEGERUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractInt(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_REALUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractFloat(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_STRINGUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_XHTMLUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IS_SIMPLIFIED":
		instance.IS_SIMPLIFIED = GongExtractBool(valueExpr)
	case "THE_VALUE":
		GongUnmarshallPointer(&instance.THE_VALUE, valueExpr, identifierMap)
	case "THE_ORIGINAL_VALUE":
		GongUnmarshallPointer(&instance.THE_ORIGINAL_VALUE, valueExpr, identifierMap)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	}
	return nil
}

type A_ALTERNATIVE_IDUnmarshaller struct{}

func (u *A_ALTERNATIVE_IDUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ALTERNATIVE_ID), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ALTERNATIVE_IDUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ALTERNATIVE_ID)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_BOOLEAN_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		instance.ATTRIBUTE_DEFINITION_BOOLEAN_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_DATE_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_DATE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_DATE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_DATE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_DATE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		instance.ATTRIBUTE_DEFINITION_DATE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_ENUMERATION_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		instance.ATTRIBUTE_DEFINITION_ENUMERATION_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_INTEGER_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_INTEGER_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_INTEGER_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_INTEGER_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		instance.ATTRIBUTE_DEFINITION_INTEGER_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_REAL_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_REAL_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_REAL_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_REAL_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_REAL_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		instance.ATTRIBUTE_DEFINITION_REAL_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_STRING_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_STRING_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_STRING_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_STRING_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_STRING_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		instance.ATTRIBUTE_DEFINITION_STRING_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_XHTML_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_XHTML_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_XHTML_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_XHTML_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		instance.ATTRIBUTE_DEFINITION_XHTML_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_BOOLEANUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_BOOLEAN, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_DATEUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_DATE":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_DATE, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_ENUMERATION, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_INTEGERUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_INTEGER, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_REALUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_REAL":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_REAL, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_STRINGUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_STRING":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_STRING, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_XHTMLUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_XHTML":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_XHTML_1Unmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_XHTML_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_XHTML_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_XHTML_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_XHTML_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_BOOLEAN, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_DATE":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_DATE, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_ENUMERATION, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_INTEGER, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_REAL":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_REAL, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_STRING":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_STRING, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_XHTML":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_CHILDRENUnmarshaller struct{}

func (u *A_CHILDRENUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_CHILDREN), stage, identifier, instanceName, preserveOrder)
}

func (u *A_CHILDRENUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_CHILDREN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_HIERARCHY":
		GongUnmarshallSliceOfPointers(&instance.SPEC_HIERARCHY, valueExpr, identifierMap)
	}
	return nil
}

type A_CORE_CONTENTUnmarshaller struct{}

func (u *A_CORE_CONTENTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_CORE_CONTENT), stage, identifier, instanceName, preserveOrder)
}

func (u *A_CORE_CONTENTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_CORE_CONTENT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "REQ_IF_CONTENT":
		GongUnmarshallPointer(&instance.REQ_IF_CONTENT, valueExpr, identifierMap)
	}
	return nil
}

type A_DATATYPESUnmarshaller struct{}

func (u *A_DATATYPESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_BOOLEAN, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_DATE":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_DATE, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_ENUMERATION, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_INTEGER, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_REAL":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_REAL, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_STRING":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_STRING, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_XHTML":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_DATATYPE_DEFINITION_BOOLEAN_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_BOOLEAN_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_BOOLEAN_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_BOOLEAN_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_BOOLEAN_REF":
		instance.DATATYPE_DEFINITION_BOOLEAN_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_DATE_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_DATE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_DATE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_DATE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_DATE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_DATE_REF":
		instance.DATATYPE_DEFINITION_DATE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_ENUMERATION_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_ENUMERATION_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_ENUMERATION_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_ENUMERATION_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_ENUMERATION_REF":
		instance.DATATYPE_DEFINITION_ENUMERATION_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_INTEGER_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_INTEGER_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_INTEGER_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_INTEGER_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_INTEGER_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_INTEGER_REF":
		instance.DATATYPE_DEFINITION_INTEGER_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_REAL_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_REAL_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_REAL_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_REAL_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_REAL_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_REAL_REF":
		instance.DATATYPE_DEFINITION_REAL_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_STRING_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_STRING_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_STRING_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_STRING_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_STRING_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_STRING_REF":
		instance.DATATYPE_DEFINITION_STRING_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_XHTML_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_XHTML_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_XHTML_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_XHTML_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_XHTML_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_XHTML_REF":
		instance.DATATYPE_DEFINITION_XHTML_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_EDITABLE_ATTSUnmarshaller struct{}

func (u *A_EDITABLE_ATTSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_EDITABLE_ATTS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_EDITABLE_ATTSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_EDITABLE_ATTS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		instance.ATTRIBUTE_DEFINITION_BOOLEAN_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		instance.ATTRIBUTE_DEFINITION_DATE_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		instance.ATTRIBUTE_DEFINITION_ENUMERATION_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		instance.ATTRIBUTE_DEFINITION_INTEGER_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		instance.ATTRIBUTE_DEFINITION_REAL_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		instance.ATTRIBUTE_DEFINITION_STRING_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		instance.ATTRIBUTE_DEFINITION_XHTML_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ENUM_VALUE_REFUnmarshaller struct{}

func (u *A_ENUM_VALUE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ENUM_VALUE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ENUM_VALUE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ENUM_VALUE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ENUM_VALUE_REF":
		instance.ENUM_VALUE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_OBJECTUnmarshaller struct{}

func (u *A_OBJECTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_OBJECT), stage, identifier, instanceName, preserveOrder)
}

func (u *A_OBJECTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_OBJECT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT_REF":
		instance.SPEC_OBJECT_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_PROPERTIESUnmarshaller struct{}

func (u *A_PROPERTIESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_PROPERTIES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_PROPERTIESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_PROPERTIES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EMBEDDED_VALUE":
		GongUnmarshallPointer(&instance.EMBEDDED_VALUE, valueExpr, identifierMap)
	}
	return nil
}

type A_RELATION_GROUP_TYPE_REFUnmarshaller struct{}

func (u *A_RELATION_GROUP_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_RELATION_GROUP_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_RELATION_GROUP_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_RELATION_GROUP_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RELATION_GROUP_TYPE_REF":
		instance.RELATION_GROUP_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SOURCE_1Unmarshaller struct{}

func (u *A_SOURCE_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SOURCE_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SOURCE_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SOURCE_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT_REF":
		instance.SPEC_OBJECT_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SOURCE_SPECIFICATION_1Unmarshaller struct{}

func (u *A_SOURCE_SPECIFICATION_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SOURCE_SPECIFICATION_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SOURCE_SPECIFICATION_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SOURCE_SPECIFICATION_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPECIFICATION_REF":
		GongUnmarshallEnum(&instance.SPECIFICATION_REF, valueExpr)
	}
	return nil
}

type A_SPECIFICATIONSUnmarshaller struct{}

func (u *A_SPECIFICATIONSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPECIFICATIONS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPECIFICATIONSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPECIFICATIONS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPECIFICATION":
		GongUnmarshallSliceOfPointers(&instance.SPECIFICATION, valueExpr, identifierMap)
	}
	return nil
}

type A_SPECIFICATION_TYPE_REFUnmarshaller struct{}

func (u *A_SPECIFICATION_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPECIFICATION_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPECIFICATION_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPECIFICATION_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPECIFICATION_TYPE_REF":
		instance.SPECIFICATION_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPECIFIED_VALUESUnmarshaller struct{}

func (u *A_SPECIFIED_VALUESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPECIFIED_VALUES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPECIFIED_VALUESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPECIFIED_VALUES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ENUM_VALUE":
		GongUnmarshallSliceOfPointers(&instance.ENUM_VALUE, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_ATTRIBUTESUnmarshaller struct{}

func (u *A_SPEC_ATTRIBUTESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_ATTRIBUTES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_ATTRIBUTESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_ATTRIBUTES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_BOOLEAN, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_DATE":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_DATE, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_ENUMERATION, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_INTEGER, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_REAL":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_REAL, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_STRING":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_STRING, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_XHTML":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_OBJECTSUnmarshaller struct{}

func (u *A_SPEC_OBJECTSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_OBJECTS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_OBJECTSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_OBJECTS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT":
		GongUnmarshallSliceOfPointers(&instance.SPEC_OBJECT, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_OBJECT_TYPE_REFUnmarshaller struct{}

func (u *A_SPEC_OBJECT_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_OBJECT_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_OBJECT_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_OBJECT_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT_TYPE_REF":
		instance.SPEC_OBJECT_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPEC_RELATIONSUnmarshaller struct{}

func (u *A_SPEC_RELATIONSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATIONS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATIONSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATIONS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_RELATION":
		GongUnmarshallSliceOfPointers(&instance.SPEC_RELATION, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_RELATION_GROUPSUnmarshaller struct{}

func (u *A_SPEC_RELATION_GROUPSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATION_GROUPS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATION_GROUPSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATION_GROUPS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RELATION_GROUP":
		GongUnmarshallSliceOfPointers(&instance.RELATION_GROUP, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_RELATION_REFUnmarshaller struct{}

func (u *A_SPEC_RELATION_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATION_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATION_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATION_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_RELATION_REF":
		instance.SPEC_RELATION_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPEC_RELATION_TYPE_REFUnmarshaller struct{}

func (u *A_SPEC_RELATION_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATION_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATION_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATION_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_RELATION_TYPE_REF":
		instance.SPEC_RELATION_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPEC_TYPESUnmarshaller struct{}

func (u *A_SPEC_TYPESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_TYPES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_TYPESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_TYPES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RELATION_GROUP_TYPE":
		GongUnmarshallSliceOfPointers(&instance.RELATION_GROUP_TYPE, valueExpr, identifierMap)
	case "SPEC_OBJECT_TYPE":
		GongUnmarshallSliceOfPointers(&instance.SPEC_OBJECT_TYPE, valueExpr, identifierMap)
	case "SPEC_RELATION_TYPE":
		GongUnmarshallSliceOfPointers(&instance.SPEC_RELATION_TYPE, valueExpr, identifierMap)
	case "SPECIFICATION_TYPE":
		GongUnmarshallSliceOfPointers(&instance.SPECIFICATION_TYPE, valueExpr, identifierMap)
	}
	return nil
}

type A_THE_HEADERUnmarshaller struct{}

func (u *A_THE_HEADERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_THE_HEADER), stage, identifier, instanceName, preserveOrder)
}

func (u *A_THE_HEADERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_THE_HEADER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "REQ_IF_HEADER":
		GongUnmarshallPointer(&instance.REQ_IF_HEADER, valueExpr, identifierMap)
	}
	return nil
}

type A_TOOL_EXTENSIONSUnmarshaller struct{}

func (u *A_TOOL_EXTENSIONSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_TOOL_EXTENSIONS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_TOOL_EXTENSIONSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_TOOL_EXTENSIONS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "REQ_IF_TOOL_EXTENSION":
		GongUnmarshallSliceOfPointers(&instance.REQ_IF_TOOL_EXTENSION, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_BOOLEANUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_DATEUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_ENUMERATIONUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPECIFIED_VALUES":
		GongUnmarshallPointer(&instance.SPECIFIED_VALUES, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_INTEGERUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MAX":
		instance.MAX = GongExtractInt(valueExpr)
	case "MIN":
		instance.MIN = GongExtractInt(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_REALUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ACCURACY":
		instance.ACCURACY = GongExtractInt(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MAX":
		instance.MAX = GongExtractFloat(valueExpr)
	case "MIN":
		instance.MIN = GongExtractFloat(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_STRINGUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MAX_LENGTH":
		instance.MAX_LENGTH = GongExtractInt(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_XHTMLUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type EMBEDDED_VALUEUnmarshaller struct{}

func (u *EMBEDDED_VALUEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EMBEDDED_VALUE), stage, identifier, instanceName, preserveOrder)
}

func (u *EMBEDDED_VALUEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EMBEDDED_VALUE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "KEY":
		instance.KEY = GongExtractInt(valueExpr)
	case "OTHER_CONTENT":
		instance.OTHER_CONTENT = GongExtractString(valueExpr)
	}
	return nil
}

type ENUM_VALUEUnmarshaller struct{}

func (u *ENUM_VALUEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ENUM_VALUE), stage, identifier, instanceName, preserveOrder)
}

func (u *ENUM_VALUEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ENUM_VALUE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "PROPERTIES":
		GongUnmarshallPointer(&instance.PROPERTIES, valueExpr, identifierMap)
	}
	return nil
}

type RELATION_GROUPUnmarshaller struct{}

func (u *RELATION_GROUPUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RELATION_GROUP), stage, identifier, instanceName, preserveOrder)
}

func (u *RELATION_GROUPUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RELATION_GROUP)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SOURCE_SPECIFICATION":
		GongUnmarshallPointer(&instance.SOURCE_SPECIFICATION, valueExpr, identifierMap)
	case "SPEC_RELATIONS":
		GongUnmarshallPointer(&instance.SPEC_RELATIONS, valueExpr, identifierMap)
	case "TARGET_SPECIFICATION":
		GongUnmarshallPointer(&instance.TARGET_SPECIFICATION, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type RELATION_GROUP_TYPEUnmarshaller struct{}

func (u *RELATION_GROUP_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RELATION_GROUP_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *RELATION_GROUP_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RELATION_GROUP_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type REQ_IFUnmarshaller struct{}

func (u *REQ_IFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "THE_HEADER":
		GongUnmarshallPointer(&instance.THE_HEADER, valueExpr, identifierMap)
	case "CORE_CONTENT":
		GongUnmarshallPointer(&instance.CORE_CONTENT, valueExpr, identifierMap)
	case "TOOL_EXTENSIONS":
		GongUnmarshallPointer(&instance.TOOL_EXTENSIONS, valueExpr, identifierMap)
	}
	return nil
}

type REQ_IF_CONTENTUnmarshaller struct{}

func (u *REQ_IF_CONTENTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF_CONTENT), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IF_CONTENTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF_CONTENT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPES":
		GongUnmarshallPointer(&instance.DATATYPES, valueExpr, identifierMap)
	case "SPEC_TYPES":
		GongUnmarshallPointer(&instance.SPEC_TYPES, valueExpr, identifierMap)
	case "SPEC_OBJECTS":
		GongUnmarshallPointer(&instance.SPEC_OBJECTS, valueExpr, identifierMap)
	case "SPEC_RELATIONS":
		GongUnmarshallPointer(&instance.SPEC_RELATIONS, valueExpr, identifierMap)
	case "SPECIFICATIONS":
		GongUnmarshallPointer(&instance.SPECIFICATIONS, valueExpr, identifierMap)
	case "SPEC_RELATION_GROUPS":
		GongUnmarshallPointer(&instance.SPEC_RELATION_GROUPS, valueExpr, identifierMap)
	}
	return nil
}

type REQ_IF_HEADERUnmarshaller struct{}

func (u *REQ_IF_HEADERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF_HEADER), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IF_HEADERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF_HEADER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "COMMENT":
		instance.COMMENT = GongExtractString(valueExpr)
	case "CREATION_TIME":
		instance.CREATION_TIME = GongExtractString(valueExpr)
	case "REPOSITORY_ID":
		instance.REPOSITORY_ID = GongExtractString(valueExpr)
	case "REQ_IF_TOOL_ID":
		instance.REQ_IF_TOOL_ID = GongExtractString(valueExpr)
	case "REQ_IF_VERSION":
		instance.REQ_IF_VERSION = GongExtractString(valueExpr)
	case "SOURCE_TOOL_ID":
		instance.SOURCE_TOOL_ID = GongExtractString(valueExpr)
	case "TITLE":
		instance.TITLE = GongExtractString(valueExpr)
	}
	return nil
}

type REQ_IF_TOOL_EXTENSIONUnmarshaller struct{}

func (u *REQ_IF_TOOL_EXTENSIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF_TOOL_EXTENSION), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IF_TOOL_EXTENSIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF_TOOL_EXTENSION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SPECIFICATIONUnmarshaller struct{}

func (u *SPECIFICATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPECIFICATION), stage, identifier, instanceName, preserveOrder)
}

func (u *SPECIFICATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPECIFICATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "CHILDREN":
		GongUnmarshallPointer(&instance.CHILDREN, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type SPECIFICATION_TYPEUnmarshaller struct{}

func (u *SPECIFICATION_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPECIFICATION_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *SPECIFICATION_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPECIFICATION_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_HIERARCHYUnmarshaller struct{}

func (u *SPEC_HIERARCHYUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_HIERARCHY), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_HIERARCHYUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_HIERARCHY)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "IS_TABLE_INTERNAL":
		instance.IS_TABLE_INTERNAL = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "CHILDREN":
		GongUnmarshallPointer(&instance.CHILDREN, valueExpr, identifierMap)
	case "EDITABLE_ATTS":
		GongUnmarshallPointer(&instance.EDITABLE_ATTS, valueExpr, identifierMap)
	case "OBJECT":
		GongUnmarshallPointer(&instance.OBJECT, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_OBJECTUnmarshaller struct{}

func (u *SPEC_OBJECTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_OBJECT), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_OBJECTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_OBJECT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_OBJECT_TYPEUnmarshaller struct{}

func (u *SPEC_OBJECT_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_OBJECT_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_OBJECT_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_OBJECT_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_RELATIONUnmarshaller struct{}

func (u *SPEC_RELATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_RELATION), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_RELATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_RELATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	case "SOURCE":
		GongUnmarshallPointer(&instance.SOURCE, valueExpr, identifierMap)
	case "TARGET":
		GongUnmarshallPointer(&instance.TARGET, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_RELATION_TYPEUnmarshaller struct{}

func (u *SPEC_RELATION_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_RELATION_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_RELATION_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_RELATION_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type XHTML_CONTENTUnmarshaller struct{}

func (u *XHTML_CONTENTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(XHTML_CONTENT), stage, identifier, instanceName, preserveOrder)
}

func (u *XHTML_CONTENTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*XHTML_CONTENT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

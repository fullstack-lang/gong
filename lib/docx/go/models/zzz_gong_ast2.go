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
type BodyUnmarshaller struct{}

func (u *BodyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Body), stage, identifier, instanceName, preserveOrder)
}

func (u *BodyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Body)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Paragraphs":
		GongUnmarshallSliceOfPointers(&instance.Paragraphs, valueExpr, identifierMap)
	case "Tables":
		GongUnmarshallSliceOfPointers(&instance.Tables, valueExpr, identifierMap)
	case "LastParagraph":
		GongUnmarshallPointer(&instance.LastParagraph, valueExpr, identifierMap)
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
	case "File":
		GongUnmarshallPointer(&instance.File, valueExpr, identifierMap)
	case "Root":
		GongUnmarshallPointer(&instance.Root, valueExpr, identifierMap)
	case "Body":
		GongUnmarshallPointer(&instance.Body, valueExpr, identifierMap)
	}
	return nil
}

type DocxUnmarshaller struct{}

func (u *DocxUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Docx), stage, identifier, instanceName, preserveOrder)
}

func (u *DocxUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Docx)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Files":
		GongUnmarshallSliceOfPointers(&instance.Files, valueExpr, identifierMap)
	case "Document":
		GongUnmarshallPointer(&instance.Document, valueExpr, identifierMap)
	}
	return nil
}

type FileUnmarshaller struct{}

func (u *FileUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(File), stage, identifier, instanceName, preserveOrder)
}

func (u *FileUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*File)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type NodeUnmarshaller struct{}

func (u *NodeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Node), stage, identifier, instanceName, preserveOrder)
}

func (u *NodeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Node)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Nodes":
		GongUnmarshallSliceOfPointers(&instance.Nodes, valueExpr, identifierMap)
	}
	return nil
}

type ParagraphUnmarshaller struct{}

func (u *ParagraphUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Paragraph), stage, identifier, instanceName, preserveOrder)
}

func (u *ParagraphUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Paragraph)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "ParagraphProperties":
		GongUnmarshallPointer(&instance.ParagraphProperties, valueExpr, identifierMap)
	case "Runes":
		GongUnmarshallSliceOfPointers(&instance.Runes, valueExpr, identifierMap)
	case "CollatedText":
		instance.CollatedText = GongExtractString(valueExpr)
	case "Next":
		GongUnmarshallPointer(&instance.Next, valueExpr, identifierMap)
	case "Previous":
		GongUnmarshallPointer(&instance.Previous, valueExpr, identifierMap)
	case "EnclosingBody":
		GongUnmarshallPointer(&instance.EnclosingBody, valueExpr, identifierMap)
	case "EnclosingTableColumn":
		GongUnmarshallPointer(&instance.EnclosingTableColumn, valueExpr, identifierMap)
	}
	return nil
}

type ParagraphPropertiesUnmarshaller struct{}

func (u *ParagraphPropertiesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParagraphProperties), stage, identifier, instanceName, preserveOrder)
}

func (u *ParagraphPropertiesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParagraphProperties)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "ParagraphStyle":
		GongUnmarshallPointer(&instance.ParagraphStyle, valueExpr, identifierMap)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	}
	return nil
}

type ParagraphStyleUnmarshaller struct{}

func (u *ParagraphStyleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParagraphStyle), stage, identifier, instanceName, preserveOrder)
}

func (u *ParagraphStyleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParagraphStyle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "ValAttr":
		instance.ValAttr = GongExtractString(valueExpr)
	}
	return nil
}

type RuneUnmarshaller struct{}

func (u *RuneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Rune), stage, identifier, instanceName, preserveOrder)
}

func (u *RuneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rune)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "Text":
		GongUnmarshallPointer(&instance.Text, valueExpr, identifierMap)
	case "RuneProperties":
		GongUnmarshallPointer(&instance.RuneProperties, valueExpr, identifierMap)
	case "EnclosingParagraph":
		GongUnmarshallPointer(&instance.EnclosingParagraph, valueExpr, identifierMap)
	}
	return nil
}

type RunePropertiesUnmarshaller struct{}

func (u *RunePropertiesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RuneProperties), stage, identifier, instanceName, preserveOrder)
}

func (u *RunePropertiesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RuneProperties)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "IsBold":
		instance.IsBold = GongExtractBool(valueExpr)
	case "IsStrike":
		instance.IsStrike = GongExtractBool(valueExpr)
	case "IsItalic":
		instance.IsItalic = GongExtractBool(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	}
	return nil
}

type TableUnmarshaller struct{}

func (u *TableUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Table), stage, identifier, instanceName, preserveOrder)
}

func (u *TableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Table)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "TableProperties":
		GongUnmarshallPointer(&instance.TableProperties, valueExpr, identifierMap)
	case "TableRows":
		GongUnmarshallSliceOfPointers(&instance.TableRows, valueExpr, identifierMap)
	}
	return nil
}

type TableColumnUnmarshaller struct{}

func (u *TableColumnUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TableColumn), stage, identifier, instanceName, preserveOrder)
}

func (u *TableColumnUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TableColumn)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "Paragraphs":
		GongUnmarshallSliceOfPointers(&instance.Paragraphs, valueExpr, identifierMap)
	}
	return nil
}

type TablePropertiesUnmarshaller struct{}

func (u *TablePropertiesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TableProperties), stage, identifier, instanceName, preserveOrder)
}

func (u *TablePropertiesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TableProperties)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "TableStyle":
		GongUnmarshallPointer(&instance.TableStyle, valueExpr, identifierMap)
	}
	return nil
}

type TableRowUnmarshaller struct{}

func (u *TableRowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TableRow), stage, identifier, instanceName, preserveOrder)
}

func (u *TableRowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TableRow)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "TableColumns":
		GongUnmarshallSliceOfPointers(&instance.TableColumns, valueExpr, identifierMap)
	}
	return nil
}

type TableStyleUnmarshaller struct{}

func (u *TableStyleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TableStyle), stage, identifier, instanceName, preserveOrder)
}

func (u *TableStyleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TableStyle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Val":
		instance.Val = GongExtractString(valueExpr)
	}
	return nil
}

type TextUnmarshaller struct{}

func (u *TextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Text), stage, identifier, instanceName, preserveOrder)
}

func (u *TextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Node":
		GongUnmarshallPointer(&instance.Node, valueExpr, identifierMap)
	case "PreserveWhiteSpace":
		instance.PreserveWhiteSpace = GongExtractBool(valueExpr)
	case "EnclosingRune":
		GongUnmarshallPointer(&instance.EnclosingRune, valueExpr, identifierMap)
	}
	return nil
}

// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

var _time__dummyDeclaration2 time.Duration
var _ = _time__dummyDeclaration2

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// ModelUnmarshaller abstracts the logic for setting fields on a staged instance
type ModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func ParseAstFile2(stage *Stage, pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func ParseAstEmbeddedFile2(stage *Stage, directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, false)
}

// ParseAstFileFromAst traverses the AST and stages instances using the Unmarshaller registry
func ParseAstFileFromAst2(stage *Stage, inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]GongstructIF)

	// 2. Visitor Pattern: Traverse the AST
	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var typeName string
					var instanceName string

					// Inspect RHS to find the Struct Type and Name
					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								typeName = selExpr.Sel.Name
								// Attempt to find Name field in literal
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					// Dispatch to specific Unmarshaller
					if typeName != "" {
						if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
							instance, err := unmarshaller.Initialize(stage, instanceName, preserveOrder)
							if err == nil {
								identifierMap[ident.Name] = instance
							}
						}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							// Extract TypeName from identifier convention (e.g. __A__...)
							// This avoids needing to store type info explicitly in the map
							parts := strings.Split(ident.Name, "__")
							if len(parts) >= 2 {
								typeName := parts[1]
								if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
									// 3. Strategy Pattern: Delegate to Handler
									unmarshaller.UnmarshallField(stage, instance, selExpr.Sel.Name, node.Rhs[0], identifierMap)
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// --- Generic Helpers for Unmarshallers ---

func GongExtractString(expr ast.Expr) string {
	if bl, ok := expr.(*ast.BasicLit); ok {
		return strings.Trim(bl.Value, "\"`")
	}
	return ""
}

func GongExtractInt(expr ast.Expr) int {
	if bl, ok := expr.(*ast.BasicLit); ok {
		val, _ := strconv.Atoi(bl.Value)
		return val
	}
	if ue, ok := expr.(*ast.UnaryExpr); ok && ue.Op == token.SUB {
		if bl, ok := ue.X.(*ast.BasicLit); ok {
			val, _ := strconv.Atoi(bl.Value)
			return -val
		}
	}
	return 0
}

func GongExtractFloat(expr ast.Expr) float64 {
	if bl, ok := expr.(*ast.BasicLit); ok {
		val, _ := strconv.ParseFloat(bl.Value, 64)
		return val
	}
	if ue, ok := expr.(*ast.UnaryExpr); ok && ue.Op == token.SUB {
		if bl, ok := ue.X.(*ast.BasicLit); ok {
			val, _ := strconv.ParseFloat(bl.Value, 64)
			return -val
		}
	}
	return 0.0
}

func GongExtractBool(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "true"
	}
	return false
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T PointerToGongstruct](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {

	if call, ok := valueExpr.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool

		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				*slice = slices.Delete(*slice, start, end)
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if ident, ok := call.Args[2].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = slices.Insert(*slice, index, val.(T))
					}
				}
			}
		} else if funcName == "append" {
			if len(call.Args) >= 2 {
				if ident, ok := call.Args[len(call.Args)-1].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = append(*slice, val.(T))
					}
				}
			}
		}
	}
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T PointerToGongstruct](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {

	if ident, ok := valueExpr.(*ast.Ident); ok {
		if val, ok := identifierMap[ident.Name]; ok {
			*ptr = val.(T)
		}
	}
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {

	// Case 1: Standard Enum usage (models.EnumType_Value)
	if sel, ok := valueExpr.(*ast.SelectorExpr); ok {
		if err := ptr.FromCodeString(sel.Sel.Name); err != nil {
			log.Printf("UnmarshallEnum: Error parsing code string '%s': %v", sel.Sel.Name, err)
		}
	} else {
		// Case 2: Fallback to string literal
		valStr := GongExtractString(valueExpr)
		if valStr != "" {
			if err := ptr.FromCodeString(valStr); err != nil {
				log.Printf("UnmarshallEnum: Error parsing string literal '%s': %v", valStr, err)
			}
		}
	}
}

// insertion point per named struct
type AttributeShapeUnmarshaller struct{}

func (u *AttributeShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AttributeShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AttributeShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributeShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FieldTypeAsString":
		instance.Name = GongExtractString(valueExpr)
	case "Structname":
		instance.Name = GongExtractString(valueExpr)
	case "Fieldtypename":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ClassdiagramUnmarshaller struct{}

func (u *ClassdiagramUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Classdiagram)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ClassdiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Classdiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Name = GongExtractString(valueExpr)
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
		instance.Name = GongExtractString(valueExpr)
	case "NodeGongEnumsIsExpanded":
		instance.NodeGongEnumsIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongEnumNodeExpansion":
		instance.Name = GongExtractString(valueExpr)
	case "NodeGongNotesIsExpanded":
		instance.NodeGongNotesIsExpanded = GongExtractBool(valueExpr)
	case "NodeGongNoteNodeExpansion":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type DiagramPackageUnmarshaller struct{}

func (u *DiagramPackageUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(DiagramPackage)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *DiagramPackageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramPackage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Path":
		instance.Name = GongExtractString(valueExpr)
	case "GongModelPath":
		instance.Name = GongExtractString(valueExpr)
	case "Classdiagrams":
		GongUnmarshallSliceOfPointers(&instance.Classdiagrams, valueExpr, identifierMap)
	case "SelectedClassdiagram":
		GongUnmarshallPointer(&instance.SelectedClassdiagram, valueExpr, identifierMap)
	case "AbsolutePathToDiagramPackage":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GongEnumShapeUnmarshaller struct{}

func (u *GongEnumShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongEnumShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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
	case "GongEnumValueShapes":
		GongUnmarshallSliceOfPointers(&instance.GongEnumValueShapes, valueExpr, identifierMap)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GongEnumValueShapeUnmarshaller struct{}

func (u *GongEnumValueShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongEnumValueShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongEnumValueShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumValueShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GongNoteLinkShapeUnmarshaller struct{}

func (u *GongNoteLinkShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongNoteLinkShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongNoteLinkShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNoteLinkShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Identifier":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	}
	return nil
}

type GongNoteShapeUnmarshaller struct{}

func (u *GongNoteShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongNoteShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongNoteShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNoteShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Identifier":
		instance.Name = GongExtractString(valueExpr)
	case "Body":
		instance.Name = GongExtractString(valueExpr)
	case "BodyHTML":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
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

func (u *GongStructShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongStructShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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
	case "AttributeShapes":
		GongUnmarshallSliceOfPointers(&instance.AttributeShapes, valueExpr, identifierMap)
	case "LinkShapes":
		GongUnmarshallSliceOfPointers(&instance.LinkShapes, valueExpr, identifierMap)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	}
	return nil
}

type LinkShapeUnmarshaller struct{}

func (u *LinkShapeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(LinkShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LinkShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	}
	return nil
}

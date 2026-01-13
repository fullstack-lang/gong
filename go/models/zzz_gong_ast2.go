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
type GongBasicFieldUnmarshaller struct{}

func (u *GongBasicFieldUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongBasicField)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongBasicFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongBasicField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BasicKindName":
		instance.Name = GongExtractString(valueExpr)
	case "GongEnum":
		GongUnmarshallPointer(&instance.GongEnum, valueExpr, identifierMap)
	case "DeclaredType":
		instance.Name = GongExtractString(valueExpr)
	case "CompositeStructName":
		instance.Name = GongExtractString(valueExpr)
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

func (u *GongEnumUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongEnum)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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

func (u *GongEnumValueUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongEnumValue)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongEnumValueUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumValue)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GongLinkUnmarshaller struct{}

func (u *GongLinkUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongLink)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Recv":
		instance.Name = GongExtractString(valueExpr)
	case "ImportPath":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GongNoteUnmarshaller struct{}

func (u *GongNoteUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongNote)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GongNoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNote)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Body":
		instance.Name = GongExtractString(valueExpr)
	case "BodyHTML":
		instance.Name = GongExtractString(valueExpr)
	case "Links":
		GongUnmarshallSliceOfPointers(&instance.Links, valueExpr, identifierMap)
	}
	return nil
}

type GongStructUnmarshaller struct{}

func (u *GongStructUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongStruct)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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
	}
	return nil
}

type GongTimeFieldUnmarshaller struct{}

func (u *GongTimeFieldUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GongTimeField)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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
		instance.Name = GongExtractString(valueExpr)
	case "BespokeTimeFormat":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type MetaReferenceUnmarshaller struct{}

func (u *MetaReferenceUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MetaReference)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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

func (u *ModelPkgUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ModelPkg)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ModelPkgUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ModelPkg)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PkgPath":
		instance.Name = GongExtractString(valueExpr)
	case "PathToGoSubDirectory":
		instance.Name = GongExtractString(valueExpr)
	case "OrmPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "DbOrmPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "DbLiteOrmPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "DbPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "ControllersPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "FullstackPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "StackPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "Level1StackPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "StaticPkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "ProbePkgGenPath":
		instance.Name = GongExtractString(valueExpr)
	case "NgWorkspacePath":
		instance.Name = GongExtractString(valueExpr)
	case "NgWorkspaceName":
		instance.Name = GongExtractString(valueExpr)
	case "NgDataLibrarySourceCodeDirectory":
		instance.Name = GongExtractString(valueExpr)
	case "NgSpecificLibrarySourceCodeDirectory":
		instance.Name = GongExtractString(valueExpr)
	case "MaterialLibDatamodelTargetPath":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PointerToGongStructFieldUnmarshaller struct{}

func (u *PointerToGongStructFieldUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PointerToGongStructField)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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
		instance.Name = GongExtractString(valueExpr)
	case "IsType":
		instance.IsType = GongExtractBool(valueExpr)
	}
	return nil
}

type SliceOfPointerToGongStructFieldUnmarshaller struct{}

func (u *SliceOfPointerToGongStructFieldUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SliceOfPointerToGongStructField)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
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
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

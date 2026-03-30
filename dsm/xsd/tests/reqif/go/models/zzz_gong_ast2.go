// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var _time__dummyDeclaration2 time.Duration
var _ = _time__dummyDeclaration2

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// ModelUnmarshaller abstracts the logic for setting fields on a staged instance
type ModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func ParseAstFile(stage *Stage, pathToFile string, preserveOrder bool) error {
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
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {
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

// GongParseAstString parses the Go source code from a string
func GongParseAstString(stage *Stage, blob string, preserveOrder bool) error {
	fileString := "package main\nfunc _() {\n" + blob + "\n}"
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", fileString, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances using the Unmarshaller registry
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]GongstructIF)

	for _, instance := range stage.GetInstances() {
		identifierMap[instance.GongGetIdentifier(stage)] = instance
	}

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
							instance, err := unmarshaller.Initialize(stage, ident.Name, instanceName, preserveOrder)
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
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								instance.UnstageVoid(stage)
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.navigationMode != GongNavigationModeNavigating {
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

// ExtractMiddleUint takes a formatted string and returns the extracted integer.
func ExtractMiddleUint(input string) (uint, error) {
	// Compile the Regex Pattern
	re := regexp.MustCompile(`__.*?__(\d+)_.*`)

	// Find the matches
	matches := re.FindStringSubmatch(input)

	// Validate that we found the pattern and the capture group
	if len(matches) < 2 {
		return 0, fmt.Errorf("pattern not found in string: %s", input)
	}

	// matches[0] is the whole string, matches[1] is the capture group (\d+)
	numberStr := matches[1]

	// Convert string to integer (handles leading zeros automatically)
	result, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to int: %v", numberStr, err)
	}

	return uint(result), nil
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

func GongExtractExpr(expr ast.Expr) any {
	switch v := expr.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.CompositeLit:
		// Reconstruct "Package.Struct{}"
		if sel, ok := v.Type.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				return id.Name + "." + sel.Sel.Name + "{}"
			}
		}
	case *ast.SelectorExpr:
		// Reconstruct "Package.Struct{}.Field"
		// X is likely a CompositeLit (Package.Struct{})
		if cl, ok := v.X.(*ast.CompositeLit); ok {
			if sel, ok := cl.Type.(*ast.SelectorExpr); ok {
				if id, ok := sel.X.(*ast.Ident); ok {
					return id.Name + "." + sel.Sel.Name + "{}." + v.Sel.Name
				}
			}
		}
		// Reconstruct "Package.Identifier"
		if id, ok := v.X.(*ast.Ident); ok {
			return id.Name + "." + v.Sel.Name
		}
	case *ast.CallExpr:
		// Reconstruct "new(Package.Struct)"
		if fun, ok := v.Fun.(*ast.Ident); ok && fun.Name == "new" {
			if len(v.Args) == 1 {
				if sel, ok := v.Args[0].(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok {
						return "new(" + id.Name + "." + sel.Sel.Name + ")"
					}
				}
			}
		}
	}
	return ""
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
					} else {
						log.Println("Ast2 Insert Unkown identifier", ident.Name)
					}
				}
			}
		} else if funcName == "append" {
			if len(call.Args) >= 2 {
				if ident, ok := call.Args[len(call.Args)-1].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = append(*slice, val.(T))
					} else {
						log.Println("Ast2 append Unkown identifier", ident.Name)
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
		if ident.Name == "nil" {
			var zero T
			*ptr = zero
			return
		}
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
type ALTERNATIVE_IDUnmarshaller struct{}

func (u *ALTERNATIVE_IDUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ALTERNATIVE_ID)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_BOOLEAN)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_DATE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_ENUMERATION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_INTEGER)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_REAL)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_STRING)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_DEFINITION_XHTML)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_BOOLEAN)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_DATE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_ENUMERATION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_INTEGER)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_REAL)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_STRING)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ATTRIBUTE_VALUE_XHTML)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ALTERNATIVE_ID)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_DATE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_REAL_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_STRING_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_DEFINITION_XHTML_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_BOOLEAN)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_DATE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_ENUMERATION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_INTEGER)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_REAL)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_STRING)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_XHTML)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ATTRIBUTE_VALUE_XHTML_1)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_CHILDREN)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_CORE_CONTENT)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPES)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_BOOLEAN_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_DATE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_ENUMERATION_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_INTEGER_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_REAL_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_STRING_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_DATATYPE_DEFINITION_XHTML_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_EDITABLE_ATTS)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_ENUM_VALUE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_OBJECT)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_PROPERTIES)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_RELATION_GROUP_TYPE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SOURCE_1)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SOURCE_SPECIFICATION_1)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPECIFICATIONS)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPECIFICATION_TYPE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPECIFIED_VALUES)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_ATTRIBUTES)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_OBJECTS)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_OBJECT_TYPE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_RELATIONS)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_RELATION_GROUPS)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_RELATION_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_RELATION_TYPE_REF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_SPEC_TYPES)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_THE_HEADER)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(A_TOOL_EXTENSIONS)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_BOOLEAN)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_DATE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_ENUMERATION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_INTEGER)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_REAL)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_STRING)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(DATATYPE_DEFINITION_XHTML)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(EMBEDDED_VALUE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(ENUM_VALUE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(RELATION_GROUP)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(RELATION_GROUP_TYPE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(REQ_IF)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(REQ_IF_CONTENT)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(REQ_IF_HEADER)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(REQ_IF_TOOL_EXTENSION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPECIFICATION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPECIFICATION_TYPE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPEC_HIERARCHY)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPEC_OBJECT)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPEC_OBJECT_TYPE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPEC_RELATION)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(SPEC_RELATION_TYPE)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	instance := new(XHTML_CONTENT)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
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
	case "PureText":
		instance.PureText = GongExtractString(valueExpr)
	}
	return nil
}

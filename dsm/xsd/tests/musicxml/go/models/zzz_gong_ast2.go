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

func (stage *Stage) UnmarshallFile(pathToFile string, preserveOrder bool) error {
	return ParseAstFile(stage, pathToFile, preserveOrder)
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
							typeName := instance.GongGetGongstructName()
							if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
								// 3. Strategy Pattern: Delegate to Handler
								unmarshaller.UnmarshallField(stage, instance, selExpr.Sel.Name, node.Rhs[0], identifierMap)
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
								// 1. Run all Before Commit hooks
								for _, hook := range stage.beforeCommitHooks {
									hook(stage)
								}

								// 2. Run all After Commit hooks
								for _, hook := range stage.afterCommitHooks {
									hook(stage)
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
	switch e := expr.(type) {

	// Case 1: It's a standard string literal
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			if unquoted, err := strconv.Unquote(e.Value); err == nil {
				return unquoted
			}
		}

	// Case 2: It's a concatenated string
	case *ast.BinaryExpr:
		// We only care if they are being added together
		if e.Op == token.ADD {
			// Recursively extract the left and right sides
			left := GongExtractString(e.X)
			right := GongExtractString(e.Y)

			// Join them back together
			return left + right
		}
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
	identifierMap map[string]GongstructIF) (err error) {

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
				if end > len(*slice) {
					// Handle error: log warning, resize slice, or return error
					// For example:
					return fmt.Errorf("index out of bounds: %d for len %d", end, len(*slice))
				}
				*slice = slices.Delete(*slice, start, end)
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index > len(*slice) {
					return fmt.Errorf("index out of bounds: %d for len %d", index, len(*slice))
				}
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
	return nil
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
type A_directiveUnmarshaller struct{}

func (u *A_directiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(A_directive)
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

func (u *A_directiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_directive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type A_measureUnmarshaller struct{}

func (u *A_measureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(A_measure)
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

func (u *A_measureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_measure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Implicit":
		GongUnmarshallEnum(&instance.Implicit, valueExpr)
	case "Non_controlling":
		GongUnmarshallEnum(&instance.Non_controlling, valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallSliceOfPointers(&instance.Note, valueExpr, identifierMap)
	case "Backup":
		GongUnmarshallSliceOfPointers(&instance.Backup, valueExpr, identifierMap)
	case "Forward":
		GongUnmarshallSliceOfPointers(&instance.Forward, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallSliceOfPointers(&instance.Direction, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Harmony":
		GongUnmarshallSliceOfPointers(&instance.Harmony, valueExpr, identifierMap)
	case "Figured_bass":
		GongUnmarshallSliceOfPointers(&instance.Figured_bass, valueExpr, identifierMap)
	case "Print":
		GongUnmarshallSliceOfPointers(&instance.Print, valueExpr, identifierMap)
	case "Sound":
		GongUnmarshallSliceOfPointers(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallSliceOfPointers(&instance.Listening, valueExpr, identifierMap)
	case "Barline":
		GongUnmarshallSliceOfPointers(&instance.Barline, valueExpr, identifierMap)
	case "Grouping":
		GongUnmarshallSliceOfPointers(&instance.Grouping, valueExpr, identifierMap)
	case "Link":
		GongUnmarshallSliceOfPointers(&instance.Link, valueExpr, identifierMap)
	case "Bookmark":
		GongUnmarshallSliceOfPointers(&instance.Bookmark, valueExpr, identifierMap)
	}
	return nil
}

type A_measure_1Unmarshaller struct{}

func (u *A_measure_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(A_measure_1)
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

func (u *A_measure_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_measure_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Implicit":
		GongUnmarshallEnum(&instance.Implicit, valueExpr)
	case "Non_controlling":
		GongUnmarshallEnum(&instance.Non_controlling, valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallSliceOfPointers(&instance.Part, valueExpr, identifierMap)
	}
	return nil
}

type A_partUnmarshaller struct{}

func (u *A_partUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(A_part)
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

func (u *A_partUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Measure":
		GongUnmarshallSliceOfPointers(&instance.Measure, valueExpr, identifierMap)
	}
	return nil
}

type A_part_1Unmarshaller struct{}

func (u *A_part_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(A_part_1)
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

func (u *A_part_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_part_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallSliceOfPointers(&instance.Note, valueExpr, identifierMap)
	case "Backup":
		GongUnmarshallSliceOfPointers(&instance.Backup, valueExpr, identifierMap)
	case "Forward":
		GongUnmarshallSliceOfPointers(&instance.Forward, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallSliceOfPointers(&instance.Direction, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Harmony":
		GongUnmarshallSliceOfPointers(&instance.Harmony, valueExpr, identifierMap)
	case "Figured_bass":
		GongUnmarshallSliceOfPointers(&instance.Figured_bass, valueExpr, identifierMap)
	case "Print":
		GongUnmarshallSliceOfPointers(&instance.Print, valueExpr, identifierMap)
	case "Sound":
		GongUnmarshallSliceOfPointers(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallSliceOfPointers(&instance.Listening, valueExpr, identifierMap)
	case "Barline":
		GongUnmarshallSliceOfPointers(&instance.Barline, valueExpr, identifierMap)
	case "Grouping":
		GongUnmarshallSliceOfPointers(&instance.Grouping, valueExpr, identifierMap)
	case "Link":
		GongUnmarshallSliceOfPointers(&instance.Link, valueExpr, identifierMap)
	case "Bookmark":
		GongUnmarshallSliceOfPointers(&instance.Bookmark, valueExpr, identifierMap)
	}
	return nil
}

type AccidentalUnmarshaller struct{}

func (u *AccidentalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Accidental)
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

func (u *AccidentalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Cautionary":
		instance.Cautionary = GongExtractString(valueExpr)
	case "Editorial":
		GongUnmarshallEnum(&instance.Editorial, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Accidental_markUnmarshaller struct{}

func (u *Accidental_markUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Accidental_mark)
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

func (u *Accidental_markUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental_mark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Accidental_textUnmarshaller struct{}

func (u *Accidental_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Accidental_text)
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

func (u *Accidental_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		instance.Dir = GongExtractString(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type AccordUnmarshaller struct{}

func (u *AccordUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Accord)
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

func (u *AccordUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accord)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "String":
		instance.String = GongExtractInt(valueExpr)
	case "Tuning_step":
		GongUnmarshallEnum(&instance.Tuning_step, valueExpr)
	case "Tuning_alter":
		instance.Tuning_alter = GongExtractString(valueExpr)
	case "Tuning_octave":
		instance.Tuning_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type Accordion_registrationUnmarshaller struct{}

func (u *Accordion_registrationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Accordion_registration)
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

func (u *Accordion_registrationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accordion_registration)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accordion_high":
		instance.Accordion_high = GongExtractString(valueExpr)
	case "Accordion_middle":
		instance.Accordion_middle = GongExtractInt(valueExpr)
	case "Accordion_low":
		instance.Accordion_low = GongExtractString(valueExpr)
	}
	return nil
}

type AppearanceUnmarshaller struct{}

func (u *AppearanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Appearance)
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

func (u *AppearanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Appearance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line_width":
		GongUnmarshallSliceOfPointers(&instance.Line_width, valueExpr, identifierMap)
	case "Note_size":
		GongUnmarshallSliceOfPointers(&instance.Note_size, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallSliceOfPointers(&instance.Distance, valueExpr, identifierMap)
	case "Glyph":
		GongUnmarshallSliceOfPointers(&instance.Glyph, valueExpr, identifierMap)
	case "Other_appearance":
		GongUnmarshallSliceOfPointers(&instance.Other_appearance, valueExpr, identifierMap)
	}
	return nil
}

type ArpeggiateUnmarshaller struct{}

func (u *ArpeggiateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Arpeggiate)
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

func (u *ArpeggiateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Arpeggiate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Direction":
		instance.Direction = GongExtractString(valueExpr)
	case "Unbroken":
		GongUnmarshallEnum(&instance.Unbroken, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type ArrowUnmarshaller struct{}

func (u *ArrowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Arrow)
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

func (u *ArrowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Arrow)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Arrow_direction":
		instance.Arrow_direction = GongExtractString(valueExpr)
	case "Arrow_style":
		instance.Arrow_style = GongExtractString(valueExpr)
	case "Arrowhead":
		instance.Arrowhead = GongExtractString(valueExpr)
	case "Circular_arrow":
		instance.Circular_arrow = GongExtractString(valueExpr)
	}
	return nil
}

type ArticulationsUnmarshaller struct{}

func (u *ArticulationsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Articulations)
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

func (u *ArticulationsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Articulations)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accent":
		GongUnmarshallSliceOfPointers(&instance.Accent, valueExpr, identifierMap)
	case "Strong_accent":
		GongUnmarshallSliceOfPointers(&instance.Strong_accent, valueExpr, identifierMap)
	case "Staccato":
		GongUnmarshallSliceOfPointers(&instance.Staccato, valueExpr, identifierMap)
	case "Tenuto":
		GongUnmarshallSliceOfPointers(&instance.Tenuto, valueExpr, identifierMap)
	case "Detached_legato":
		GongUnmarshallSliceOfPointers(&instance.Detached_legato, valueExpr, identifierMap)
	case "Staccatissimo":
		GongUnmarshallSliceOfPointers(&instance.Staccatissimo, valueExpr, identifierMap)
	case "Spiccato":
		GongUnmarshallSliceOfPointers(&instance.Spiccato, valueExpr, identifierMap)
	case "Scoop":
		GongUnmarshallSliceOfPointers(&instance.Scoop, valueExpr, identifierMap)
	case "Plop":
		GongUnmarshallSliceOfPointers(&instance.Plop, valueExpr, identifierMap)
	case "Doit":
		GongUnmarshallSliceOfPointers(&instance.Doit, valueExpr, identifierMap)
	case "Falloff":
		GongUnmarshallSliceOfPointers(&instance.Falloff, valueExpr, identifierMap)
	case "Breath_mark":
		GongUnmarshallSliceOfPointers(&instance.Breath_mark, valueExpr, identifierMap)
	case "Caesura":
		GongUnmarshallSliceOfPointers(&instance.Caesura, valueExpr, identifierMap)
	case "Stress":
		GongUnmarshallSliceOfPointers(&instance.Stress, valueExpr, identifierMap)
	case "Unstress":
		GongUnmarshallSliceOfPointers(&instance.Unstress, valueExpr, identifierMap)
	case "Soft_accent":
		GongUnmarshallSliceOfPointers(&instance.Soft_accent, valueExpr, identifierMap)
	case "Other_articulation":
		GongUnmarshallSliceOfPointers(&instance.Other_articulation, valueExpr, identifierMap)
	}
	return nil
}

type AssessUnmarshaller struct{}

func (u *AssessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Assess)
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

func (u *AssessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Assess)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	}
	return nil
}

type AttributesUnmarshaller struct{}

func (u *AttributesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Attributes)
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

func (u *AttributesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Attributes)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Key":
		GongUnmarshallSliceOfPointers(&instance.Key, valueExpr, identifierMap)
	case "Time":
		GongUnmarshallSliceOfPointers(&instance.Time, valueExpr, identifierMap)
	case "Staves":
		instance.Staves = GongExtractInt(valueExpr)
	case "Part_symbol":
		GongUnmarshallPointer(&instance.Part_symbol, valueExpr, identifierMap)
	case "Instruments":
		instance.Instruments = GongExtractInt(valueExpr)
	case "Clef":
		GongUnmarshallSliceOfPointers(&instance.Clef, valueExpr, identifierMap)
	case "Staff_details":
		GongUnmarshallSliceOfPointers(&instance.Staff_details, valueExpr, identifierMap)
	case "Transpose":
		GongUnmarshallSliceOfPointers(&instance.Transpose, valueExpr, identifierMap)
	case "For_part":
		GongUnmarshallSliceOfPointers(&instance.For_part, valueExpr, identifierMap)
	case "Directive":
		GongUnmarshallSliceOfPointers(&instance.Directive, valueExpr, identifierMap)
	case "Measure_style":
		GongUnmarshallSliceOfPointers(&instance.Measure_style, valueExpr, identifierMap)
	}
	return nil
}

type BackupUnmarshaller struct{}

func (u *BackupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Backup)
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

func (u *BackupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Backup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Bar_style_colorUnmarshaller struct{}

func (u *Bar_style_colorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bar_style_color)
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

func (u *Bar_style_colorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bar_style_color)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type BarlineUnmarshaller struct{}

func (u *BarlineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Barline)
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

func (u *BarlineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Barline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		instance.Location = GongExtractString(valueExpr)
	case "Segno":
		instance.Segno = GongExtractString(valueExpr)
	case "Coda":
		instance.Coda = GongExtractString(valueExpr)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Bar_style":
		GongUnmarshallPointer(&instance.Bar_style, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Wavy_line":
		GongUnmarshallPointer(&instance.Wavy_line, valueExpr, identifierMap)
	case "Segno_1":
		GongUnmarshallPointer(&instance.Segno_1, valueExpr, identifierMap)
	case "Coda_1":
		GongUnmarshallPointer(&instance.Coda_1, valueExpr, identifierMap)
	case "Fermata":
		GongUnmarshallPointer(&instance.Fermata, valueExpr, identifierMap)
	case "Ending":
		GongUnmarshallPointer(&instance.Ending, valueExpr, identifierMap)
	case "Repeat":
		GongUnmarshallPointer(&instance.Repeat, valueExpr, identifierMap)
	}
	return nil
}

type BarreUnmarshaller struct{}

func (u *BarreUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Barre)
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

func (u *BarreUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Barre)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type BassUnmarshaller struct{}

func (u *BassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bass)
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

func (u *BassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Arrangement":
		instance.Arrangement = GongExtractString(valueExpr)
	case "Bass_separator":
		GongUnmarshallPointer(&instance.Bass_separator, valueExpr, identifierMap)
	case "Bass_step":
		GongUnmarshallPointer(&instance.Bass_step, valueExpr, identifierMap)
	case "Bass_alter":
		GongUnmarshallPointer(&instance.Bass_alter, valueExpr, identifierMap)
	}
	return nil
}

type Bass_stepUnmarshaller struct{}

func (u *Bass_stepUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bass_step)
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

func (u *Bass_stepUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bass_step)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type BeamUnmarshaller struct{}

func (u *BeamUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Beam)
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

func (u *BeamUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beam)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Repeater":
		GongUnmarshallEnum(&instance.Repeater, valueExpr)
	case "Fan":
		instance.Fan = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Beat_repeatUnmarshaller struct{}

func (u *Beat_repeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Beat_repeat)
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

func (u *Beat_repeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beat_repeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Slashes":
		instance.Slashes = GongExtractInt(valueExpr)
	case "Use_dots":
		GongUnmarshallEnum(&instance.Use_dots, valueExpr)
	case "Slash_type":
		GongUnmarshallEnum(&instance.Slash_type, valueExpr)
	case "Slash_dot":
		instance.Slash_dot = GongExtractString(valueExpr)
	case "Except_voice":
		instance.Except_voice = GongExtractString(valueExpr)
	}
	return nil
}

type Beat_unit_tiedUnmarshaller struct{}

func (u *Beat_unit_tiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Beat_unit_tied)
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

func (u *Beat_unit_tiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beat_unit_tied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Beat_unit":
		GongUnmarshallEnum(&instance.Beat_unit, valueExpr)
	case "Beat_unit_dot":
		instance.Beat_unit_dot = GongExtractString(valueExpr)
	}
	return nil
}

type BeaterUnmarshaller struct{}

func (u *BeaterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Beater)
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

func (u *BeaterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beater)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tip":
		instance.Tip = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type BendUnmarshaller struct{}

func (u *BendUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bend)
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

func (u *BendUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bend)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Shape":
		instance.Shape = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "First_beat":
		instance.First_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	case "Bend_alter":
		instance.Bend_alter = GongExtractString(valueExpr)
	case "Pre_bend":
		instance.Pre_bend = GongExtractString(valueExpr)
	case "Release":
		GongUnmarshallPointer(&instance.Release, valueExpr, identifierMap)
	case "With_bar":
		GongUnmarshallPointer(&instance.With_bar, valueExpr, identifierMap)
	}
	return nil
}

type BookmarkUnmarshaller struct{}

func (u *BookmarkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bookmark)
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

func (u *BookmarkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bookmark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Position":
		instance.Position = GongExtractInt(valueExpr)
	}
	return nil
}

type BracketUnmarshaller struct{}

func (u *BracketUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bracket)
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

func (u *BracketUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bracket)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_end":
		instance.Line_end = GongExtractString(valueExpr)
	case "End_length":
		instance.End_length = GongExtractString(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Breath_markUnmarshaller struct{}

func (u *Breath_markUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Breath_mark)
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

func (u *Breath_markUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Breath_mark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CaesuraUnmarshaller struct{}

func (u *CaesuraUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Caesura)
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

func (u *CaesuraUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Caesura)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CancelUnmarshaller struct{}

func (u *CancelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Cancel)
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

func (u *CancelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Cancel)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		instance.Location = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type ClefUnmarshaller struct{}

func (u *ClefUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Clef)
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

func (u *ClefUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Clef)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Additional":
		GongUnmarshallEnum(&instance.Additional, valueExpr)
	case "Size":
		instance.Size = GongExtractString(valueExpr)
	case "After_barline":
		GongUnmarshallEnum(&instance.After_barline, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Sign":
		instance.Sign = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Clef_octave_change":
		instance.Clef_octave_change = GongExtractInt(valueExpr)
	}
	return nil
}

type CodaUnmarshaller struct{}

func (u *CodaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Coda)
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

func (u *CodaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Coda)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type CreditUnmarshaller struct{}

func (u *CreditUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Credit)
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

func (u *CreditUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Credit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Page":
		instance.Page = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Credit_type":
		instance.Credit_type = GongExtractString(valueExpr)
	case "Credit_image":
		GongUnmarshallPointer(&instance.Credit_image, valueExpr, identifierMap)
	case "Link":
		GongUnmarshallSliceOfPointers(&instance.Link, valueExpr, identifierMap)
	case "Bookmark":
		GongUnmarshallSliceOfPointers(&instance.Bookmark, valueExpr, identifierMap)
	case "Credit_words":
		GongUnmarshallSliceOfPointers(&instance.Credit_words, valueExpr, identifierMap)
	case "Credit_symbol":
		GongUnmarshallSliceOfPointers(&instance.Credit_symbol, valueExpr, identifierMap)
	}
	return nil
}

type DashesUnmarshaller struct{}

func (u *DashesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Dashes)
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

func (u *DashesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dashes)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type DefaultsUnmarshaller struct{}

func (u *DefaultsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Defaults)
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

func (u *DefaultsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Defaults)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Scaling":
		GongUnmarshallPointer(&instance.Scaling, valueExpr, identifierMap)
	case "Concert_score":
		instance.Concert_score = GongExtractString(valueExpr)
	case "Page_layout":
		GongUnmarshallPointer(&instance.Page_layout, valueExpr, identifierMap)
	case "System_layout":
		GongUnmarshallPointer(&instance.System_layout, valueExpr, identifierMap)
	case "Staff_layout":
		GongUnmarshallSliceOfPointers(&instance.Staff_layout, valueExpr, identifierMap)
	case "Appearance":
		GongUnmarshallPointer(&instance.Appearance, valueExpr, identifierMap)
	case "Music_font":
		GongUnmarshallPointer(&instance.Music_font, valueExpr, identifierMap)
	case "Word_font":
		GongUnmarshallPointer(&instance.Word_font, valueExpr, identifierMap)
	case "Lyric_font":
		GongUnmarshallSliceOfPointers(&instance.Lyric_font, valueExpr, identifierMap)
	case "Lyric_language":
		GongUnmarshallSliceOfPointers(&instance.Lyric_language, valueExpr, identifierMap)
	}
	return nil
}

type DegreeUnmarshaller struct{}

func (u *DegreeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Degree)
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

func (u *DegreeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Degree_value":
		GongUnmarshallPointer(&instance.Degree_value, valueExpr, identifierMap)
	case "Degree_alter":
		GongUnmarshallPointer(&instance.Degree_alter, valueExpr, identifierMap)
	case "Degree_type":
		GongUnmarshallPointer(&instance.Degree_type, valueExpr, identifierMap)
	}
	return nil
}

type Degree_alterUnmarshaller struct{}

func (u *Degree_alterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Degree_alter)
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

func (u *Degree_alterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree_alter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Plus_minus":
		GongUnmarshallEnum(&instance.Plus_minus, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Degree_typeUnmarshaller struct{}

func (u *Degree_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Degree_type)
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

func (u *Degree_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Degree_valueUnmarshaller struct{}

func (u *Degree_valueUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Degree_value)
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

func (u *Degree_valueUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree_value)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		instance.Symbol = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type DirectionUnmarshaller struct{}

func (u *DirectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Direction)
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

func (u *DirectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Direction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Directive":
		GongUnmarshallEnum(&instance.Directive, valueExpr)
	case "System":
		GongUnmarshallEnum(&instance.System, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Direction_type":
		GongUnmarshallSliceOfPointers(&instance.Direction_type, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Sound":
		GongUnmarshallPointer(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallPointer(&instance.Listening, valueExpr, identifierMap)
	}
	return nil
}

type Direction_typeUnmarshaller struct{}

func (u *Direction_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Direction_type)
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

func (u *Direction_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Direction_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Rehearsal":
		GongUnmarshallSliceOfPointers(&instance.Rehearsal, valueExpr, identifierMap)
	case "Segno":
		GongUnmarshallSliceOfPointers(&instance.Segno, valueExpr, identifierMap)
	case "Coda":
		GongUnmarshallSliceOfPointers(&instance.Coda, valueExpr, identifierMap)
	case "Words":
		GongUnmarshallSliceOfPointers(&instance.Words, valueExpr, identifierMap)
	case "Symbol":
		GongUnmarshallSliceOfPointers(&instance.Symbol, valueExpr, identifierMap)
	case "Wedge":
		GongUnmarshallPointer(&instance.Wedge, valueExpr, identifierMap)
	case "Dynamics":
		GongUnmarshallSliceOfPointers(&instance.Dynamics, valueExpr, identifierMap)
	case "Dashes":
		GongUnmarshallPointer(&instance.Dashes, valueExpr, identifierMap)
	case "Bracket":
		GongUnmarshallPointer(&instance.Bracket, valueExpr, identifierMap)
	case "Pedal":
		GongUnmarshallPointer(&instance.Pedal, valueExpr, identifierMap)
	case "Metronome":
		GongUnmarshallPointer(&instance.Metronome, valueExpr, identifierMap)
	case "Octave_shift":
		GongUnmarshallPointer(&instance.Octave_shift, valueExpr, identifierMap)
	case "Harp_pedals":
		GongUnmarshallPointer(&instance.Harp_pedals, valueExpr, identifierMap)
	case "Damp":
		GongUnmarshallPointer(&instance.Damp, valueExpr, identifierMap)
	case "Damp_all":
		GongUnmarshallPointer(&instance.Damp_all, valueExpr, identifierMap)
	case "Eyeglasses":
		GongUnmarshallPointer(&instance.Eyeglasses, valueExpr, identifierMap)
	case "String_mute":
		GongUnmarshallPointer(&instance.String_mute, valueExpr, identifierMap)
	case "Scordatura":
		GongUnmarshallPointer(&instance.Scordatura, valueExpr, identifierMap)
	case "Image":
		GongUnmarshallPointer(&instance.Image, valueExpr, identifierMap)
	case "Principal_voice":
		GongUnmarshallPointer(&instance.Principal_voice, valueExpr, identifierMap)
	case "Percussion":
		GongUnmarshallSliceOfPointers(&instance.Percussion, valueExpr, identifierMap)
	case "Accordion_registration":
		GongUnmarshallPointer(&instance.Accordion_registration, valueExpr, identifierMap)
	case "Staff_divide":
		GongUnmarshallPointer(&instance.Staff_divide, valueExpr, identifierMap)
	case "Other_direction":
		GongUnmarshallPointer(&instance.Other_direction, valueExpr, identifierMap)
	}
	return nil
}

type DistanceUnmarshaller struct{}

func (u *DistanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Distance)
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

func (u *DistanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Distance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DoubleUnmarshaller struct{}

func (u *DoubleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Double)
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

func (u *DoubleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Double)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Above":
		GongUnmarshallEnum(&instance.Above, valueExpr)
	}
	return nil
}

type DynamicsUnmarshaller struct{}

func (u *DynamicsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Dynamics)
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

func (u *DynamicsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dynamics)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "P":
		instance.P = GongExtractString(valueExpr)
	case "Pp":
		instance.Pp = GongExtractString(valueExpr)
	case "Ppp":
		instance.Ppp = GongExtractString(valueExpr)
	case "Pppp":
		instance.Pppp = GongExtractString(valueExpr)
	case "Ppppp":
		instance.Ppppp = GongExtractString(valueExpr)
	case "Pppppp":
		instance.Pppppp = GongExtractString(valueExpr)
	case "F":
		instance.F = GongExtractString(valueExpr)
	case "Ff":
		instance.Ff = GongExtractString(valueExpr)
	case "Fff":
		instance.Fff = GongExtractString(valueExpr)
	case "Ffff":
		instance.Ffff = GongExtractString(valueExpr)
	case "Fffff":
		instance.Fffff = GongExtractString(valueExpr)
	case "Ffffff":
		instance.Ffffff = GongExtractString(valueExpr)
	case "Mp":
		instance.Mp = GongExtractString(valueExpr)
	case "Mf":
		instance.Mf = GongExtractString(valueExpr)
	case "Sf":
		instance.Sf = GongExtractString(valueExpr)
	case "Sfp":
		instance.Sfp = GongExtractString(valueExpr)
	case "Sfpp":
		instance.Sfpp = GongExtractString(valueExpr)
	case "Fp":
		instance.Fp = GongExtractString(valueExpr)
	case "Rf":
		instance.Rf = GongExtractString(valueExpr)
	case "Rfz":
		instance.Rfz = GongExtractString(valueExpr)
	case "Sfz":
		instance.Sfz = GongExtractString(valueExpr)
	case "Sffz":
		instance.Sffz = GongExtractString(valueExpr)
	case "Fz":
		instance.Fz = GongExtractString(valueExpr)
	case "N":
		instance.N = GongExtractString(valueExpr)
	case "Pf":
		instance.Pf = GongExtractString(valueExpr)
	case "Sfzp":
		instance.Sfzp = GongExtractString(valueExpr)
	case "Other_dynamics":
		GongUnmarshallSliceOfPointers(&instance.Other_dynamics, valueExpr, identifierMap)
	}
	return nil
}

type EffectUnmarshaller struct{}

func (u *EffectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Effect)
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

func (u *EffectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Effect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ElisionUnmarshaller struct{}

func (u *ElisionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Elision)
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

func (u *ElisionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Elision)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type EmptyUnmarshaller struct{}

func (u *EmptyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty)
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

func (u *EmptyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_fontUnmarshaller struct{}

func (u *Empty_fontUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_font)
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

func (u *Empty_fontUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_font)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_lineUnmarshaller struct{}

func (u *Empty_lineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_line)
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

func (u *Empty_lineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line_shape":
		instance.Line_shape = GongExtractString(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Line_length":
		instance.Line_length = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_placementUnmarshaller struct{}

func (u *Empty_placementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_placement)
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

func (u *Empty_placementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_placement)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_placement_smuflUnmarshaller struct{}

func (u *Empty_placement_smuflUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_placement_smufl)
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

func (u *Empty_placement_smuflUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_placement_smufl)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_print_object_style_alignUnmarshaller struct{}

func (u *Empty_print_object_style_alignUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_print_object_style_align)
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

func (u *Empty_print_object_style_alignUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_object_style_align)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_print_styleUnmarshaller struct{}

func (u *Empty_print_styleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_print_style)
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

func (u *Empty_print_styleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_style)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_print_style_alignUnmarshaller struct{}

func (u *Empty_print_style_alignUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_print_style_align)
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

func (u *Empty_print_style_alignUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_style_align)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_print_style_align_idUnmarshaller struct{}

func (u *Empty_print_style_align_idUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_print_style_align_id)
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

func (u *Empty_print_style_align_idUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_style_align_id)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_trill_soundUnmarshaller struct{}

func (u *Empty_trill_soundUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Empty_trill_sound)
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

func (u *Empty_trill_soundUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_trill_sound)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Start_note":
		instance.Start_note = GongExtractString(valueExpr)
	case "Trill_step":
		instance.Trill_step = GongExtractString(valueExpr)
	case "Two_note_turn":
		instance.Two_note_turn = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	}
	return nil
}

type EncodingUnmarshaller struct{}

func (u *EncodingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Encoding)
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

func (u *EncodingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Encoding)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Encoder":
		GongUnmarshallSliceOfPointers(&instance.Encoder, valueExpr, identifierMap)
	case "Software":
		instance.Software = GongExtractString(valueExpr)
	case "Encoding_description":
		instance.Encoding_description = GongExtractString(valueExpr)
	case "Supports":
		GongUnmarshallSliceOfPointers(&instance.Supports, valueExpr, identifierMap)
	}
	return nil
}

type EndingUnmarshaller struct{}

func (u *EndingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Ending)
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

func (u *EndingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ending)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "End_length":
		instance.End_length = GongExtractString(valueExpr)
	case "Text_x":
		instance.Text_x = GongExtractString(valueExpr)
	case "Text_y":
		instance.Text_y = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "System":
		GongUnmarshallEnum(&instance.System, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ExtendUnmarshaller struct{}

func (u *ExtendUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Extend)
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

func (u *ExtendUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extend)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type FeatureUnmarshaller struct{}

func (u *FeatureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Feature)
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

func (u *FeatureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Feature)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type FermataUnmarshaller struct{}

func (u *FermataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Fermata)
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

func (u *FermataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fermata)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type FigureUnmarshaller struct{}

func (u *FigureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Figure)
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

func (u *FigureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Figure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Prefix":
		GongUnmarshallPointer(&instance.Prefix, valueExpr, identifierMap)
	case "Figure_number":
		GongUnmarshallPointer(&instance.Figure_number, valueExpr, identifierMap)
	case "Suffix":
		GongUnmarshallPointer(&instance.Suffix, valueExpr, identifierMap)
	case "Extend":
		GongUnmarshallPointer(&instance.Extend, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Figured_bassUnmarshaller struct{}

func (u *Figured_bassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Figured_bass)
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

func (u *Figured_bassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Figured_bass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Print_dot":
		GongUnmarshallEnum(&instance.Print_dot, valueExpr)
	case "Print_lyric":
		GongUnmarshallEnum(&instance.Print_lyric, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Print_spacing":
		GongUnmarshallEnum(&instance.Print_spacing, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Figure":
		GongUnmarshallSliceOfPointers(&instance.Figure, valueExpr, identifierMap)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type FingeringUnmarshaller struct{}

func (u *FingeringUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Fingering)
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

func (u *FingeringUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fingering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Substitution":
		GongUnmarshallEnum(&instance.Substitution, valueExpr)
	case "Alternate":
		GongUnmarshallEnum(&instance.Alternate, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type First_fretUnmarshaller struct{}

func (u *First_fretUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(First_fret)
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

func (u *First_fretUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*First_fret)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Location":
		instance.Location = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type For_partUnmarshaller struct{}

func (u *For_partUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(For_part)
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

func (u *For_partUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*For_part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Part_clef":
		GongUnmarshallPointer(&instance.Part_clef, valueExpr, identifierMap)
	case "Part_transpose":
		GongUnmarshallPointer(&instance.Part_transpose, valueExpr, identifierMap)
	}
	return nil
}

type Formatted_symbolUnmarshaller struct{}

func (u *Formatted_symbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Formatted_symbol)
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

func (u *Formatted_symbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_symbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		instance.Dir = GongExtractString(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Formatted_symbol_idUnmarshaller struct{}

func (u *Formatted_symbol_idUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Formatted_symbol_id)
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

func (u *Formatted_symbol_idUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_symbol_id)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		instance.Dir = GongExtractString(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Formatted_textUnmarshaller struct{}

func (u *Formatted_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Formatted_text)
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

func (u *Formatted_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		instance.Dir = GongExtractString(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Formatted_text_idUnmarshaller struct{}

func (u *Formatted_text_idUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Formatted_text_id)
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

func (u *Formatted_text_idUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_text_id)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		instance.Dir = GongExtractString(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ForwardUnmarshaller struct{}

func (u *ForwardUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Forward)
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

func (u *ForwardUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Forward)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	}
	return nil
}

type FrameUnmarshaller struct{}

func (u *FrameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Frame)
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

func (u *FrameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Frame)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Height":
		instance.Height = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Unplayed":
		instance.Unplayed = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Frame_strings":
		instance.Frame_strings = GongExtractInt(valueExpr)
	case "Frame_frets":
		instance.Frame_frets = GongExtractInt(valueExpr)
	case "First_fret":
		GongUnmarshallPointer(&instance.First_fret, valueExpr, identifierMap)
	case "Frame_note":
		GongUnmarshallSliceOfPointers(&instance.Frame_note, valueExpr, identifierMap)
	}
	return nil
}

type Frame_noteUnmarshaller struct{}

func (u *Frame_noteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Frame_note)
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

func (u *Frame_noteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Frame_note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "String":
		GongUnmarshallPointer(&instance.String, valueExpr, identifierMap)
	case "Fret":
		GongUnmarshallPointer(&instance.Fret, valueExpr, identifierMap)
	case "Fingering":
		GongUnmarshallPointer(&instance.Fingering, valueExpr, identifierMap)
	case "Barre":
		GongUnmarshallPointer(&instance.Barre, valueExpr, identifierMap)
	}
	return nil
}

type FretUnmarshaller struct{}

func (u *FretUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Fret)
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

func (u *FretUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fret)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type GlassUnmarshaller struct{}

func (u *GlassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Glass)
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

func (u *GlassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type GlissandoUnmarshaller struct{}

func (u *GlissandoUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Glissando)
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

func (u *GlissandoUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glissando)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type GlyphUnmarshaller struct{}

func (u *GlyphUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Glyph)
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

func (u *GlyphUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glyph)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type GraceUnmarshaller struct{}

func (u *GraceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Grace)
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

func (u *GraceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Grace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Steal_time_previous":
		instance.Steal_time_previous = GongExtractString(valueExpr)
	case "Steal_time_following":
		instance.Steal_time_following = GongExtractString(valueExpr)
	case "Make_time":
		instance.Make_time = GongExtractString(valueExpr)
	case "Slash":
		GongUnmarshallEnum(&instance.Slash, valueExpr)
	}
	return nil
}

type Group_barlineUnmarshaller struct{}

func (u *Group_barlineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Group_barline)
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

func (u *Group_barlineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group_barline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Group_nameUnmarshaller struct{}

func (u *Group_nameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Group_name)
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

func (u *Group_nameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group_name)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Group_symbolUnmarshaller struct{}

func (u *Group_symbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Group_symbol)
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

func (u *Group_symbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group_symbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type GroupingUnmarshaller struct{}

func (u *GroupingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Grouping)
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

func (u *GroupingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Grouping)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Member_of":
		instance.Member_of = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Feature":
		GongUnmarshallSliceOfPointers(&instance.Feature, valueExpr, identifierMap)
	}
	return nil
}

type Hammer_on_pull_offUnmarshaller struct{}

func (u *Hammer_on_pull_offUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Hammer_on_pull_off)
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

func (u *Hammer_on_pull_offUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hammer_on_pull_off)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type HandbellUnmarshaller struct{}

func (u *HandbellUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Handbell)
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

func (u *HandbellUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Handbell)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Harmon_closedUnmarshaller struct{}

func (u *Harmon_closedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Harmon_closed)
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

func (u *Harmon_closedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmon_closed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		instance.Location = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Harmon_muteUnmarshaller struct{}

func (u *Harmon_muteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Harmon_mute)
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

func (u *Harmon_muteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmon_mute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Harmon_closed":
		GongUnmarshallPointer(&instance.Harmon_closed, valueExpr, identifierMap)
	}
	return nil
}

type HarmonicUnmarshaller struct{}

func (u *HarmonicUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Harmonic)
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

func (u *HarmonicUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmonic)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Natural":
		instance.Natural = GongExtractString(valueExpr)
	case "Artificial":
		instance.Artificial = GongExtractString(valueExpr)
	case "Base_pitch":
		instance.Base_pitch = GongExtractString(valueExpr)
	case "Touching_pitch":
		instance.Touching_pitch = GongExtractString(valueExpr)
	case "Sounding_pitch":
		instance.Sounding_pitch = GongExtractString(valueExpr)
	}
	return nil
}

type HarmonyUnmarshaller struct{}

func (u *HarmonyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Harmony)
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

func (u *HarmonyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmony)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Print_frame":
		GongUnmarshallEnum(&instance.Print_frame, valueExpr)
	case "Arrangement":
		GongUnmarshallEnum(&instance.Arrangement, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "System":
		GongUnmarshallEnum(&instance.System, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Root":
		GongUnmarshallPointer(&instance.Root, valueExpr, identifierMap)
	case "Numeral":
		GongUnmarshallPointer(&instance.Numeral, valueExpr, identifierMap)
	case "Function":
		GongUnmarshallPointer(&instance.Function, valueExpr, identifierMap)
	case "Kind":
		GongUnmarshallPointer(&instance.Kind, valueExpr, identifierMap)
	case "Inversion":
		GongUnmarshallPointer(&instance.Inversion, valueExpr, identifierMap)
	case "Bass":
		GongUnmarshallPointer(&instance.Bass, valueExpr, identifierMap)
	case "Degree":
		GongUnmarshallSliceOfPointers(&instance.Degree, valueExpr, identifierMap)
	case "Frame":
		GongUnmarshallPointer(&instance.Frame, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	}
	return nil
}

type Harmony_alterUnmarshaller struct{}

func (u *Harmony_alterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Harmony_alter)
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

func (u *Harmony_alterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmony_alter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Harp_pedalsUnmarshaller struct{}

func (u *Harp_pedalsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Harp_pedals)
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

func (u *Harp_pedalsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harp_pedals)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Pedal_tuning":
		GongUnmarshallSliceOfPointers(&instance.Pedal_tuning, valueExpr, identifierMap)
	}
	return nil
}

type Heel_toeUnmarshaller struct{}

func (u *Heel_toeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Heel_toe)
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

func (u *Heel_toeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Heel_toe)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type HoleUnmarshaller struct{}

func (u *HoleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Hole)
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

func (u *HoleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hole)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Hole_type":
		instance.Hole_type = GongExtractString(valueExpr)
	case "Hole_closed":
		GongUnmarshallPointer(&instance.Hole_closed, valueExpr, identifierMap)
	case "Hole_shape":
		instance.Hole_shape = GongExtractString(valueExpr)
	}
	return nil
}

type Hole_closedUnmarshaller struct{}

func (u *Hole_closedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Hole_closed)
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

func (u *Hole_closedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hole_closed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		instance.Location = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Horizontal_turnUnmarshaller struct{}

func (u *Horizontal_turnUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Horizontal_turn)
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

func (u *Horizontal_turnUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Horizontal_turn)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Slash":
		GongUnmarshallEnum(&instance.Slash, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Start_note":
		instance.Start_note = GongExtractString(valueExpr)
	case "Trill_step":
		instance.Trill_step = GongExtractString(valueExpr)
	case "Two_note_turn":
		instance.Two_note_turn = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	}
	return nil
}

type IdentificationUnmarshaller struct{}

func (u *IdentificationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Identification)
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

func (u *IdentificationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Identification)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Creator":
		GongUnmarshallSliceOfPointers(&instance.Creator, valueExpr, identifierMap)
	case "Rights":
		GongUnmarshallSliceOfPointers(&instance.Rights, valueExpr, identifierMap)
	case "Encoding":
		GongUnmarshallPointer(&instance.Encoding, valueExpr, identifierMap)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Relation":
		GongUnmarshallSliceOfPointers(&instance.Relation, valueExpr, identifierMap)
	case "Miscellaneous":
		GongUnmarshallPointer(&instance.Miscellaneous, valueExpr, identifierMap)
	}
	return nil
}

type ImageUnmarshaller struct{}

func (u *ImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Image)
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

func (u *ImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Image)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Height":
		instance.Height = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type InstrumentUnmarshaller struct{}

func (u *InstrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Instrument)
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

func (u *InstrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Instrument_changeUnmarshaller struct{}

func (u *Instrument_changeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Instrument_change)
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

func (u *Instrument_changeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument_change)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Instrument_sound":
		instance.Instrument_sound = GongExtractString(valueExpr)
	case "Solo":
		instance.Solo = GongExtractString(valueExpr)
	case "Ensemble":
		instance.Ensemble = GongExtractString(valueExpr)
	case "Virtual_instrument":
		GongUnmarshallPointer(&instance.Virtual_instrument, valueExpr, identifierMap)
	}
	return nil
}

type Instrument_linkUnmarshaller struct{}

func (u *Instrument_linkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Instrument_link)
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

func (u *Instrument_linkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument_link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type InterchangeableUnmarshaller struct{}

func (u *InterchangeableUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Interchangeable)
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

func (u *InterchangeableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Interchangeable)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		instance.Symbol = GongExtractString(valueExpr)
	case "Separator":
		instance.Separator = GongExtractString(valueExpr)
	case "Time_relation":
		instance.Time_relation = GongExtractString(valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Beat_type":
		instance.Beat_type = GongExtractString(valueExpr)
	}
	return nil
}

type InversionUnmarshaller struct{}

func (u *InversionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Inversion)
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

func (u *InversionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Inversion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type KeyUnmarshaller struct{}

func (u *KeyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Key)
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

func (u *KeyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Cancel":
		GongUnmarshallPointer(&instance.Cancel, valueExpr, identifierMap)
	case "Fifths":
		instance.Fifths = GongExtractInt(valueExpr)
	case "Mode":
		instance.Mode = GongExtractString(valueExpr)
	case "Key_step":
		GongUnmarshallEnum(&instance.Key_step, valueExpr)
	case "Key_alter":
		instance.Key_alter = GongExtractString(valueExpr)
	case "Key_accidental":
		GongUnmarshallPointer(&instance.Key_accidental, valueExpr, identifierMap)
	case "Key_octave":
		GongUnmarshallSliceOfPointers(&instance.Key_octave, valueExpr, identifierMap)
	}
	return nil
}

type Key_accidentalUnmarshaller struct{}

func (u *Key_accidentalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Key_accidental)
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

func (u *Key_accidentalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key_accidental)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Key_octaveUnmarshaller struct{}

func (u *Key_octaveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Key_octave)
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

func (u *Key_octaveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key_octave)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Cancel":
		GongUnmarshallEnum(&instance.Cancel, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type KindUnmarshaller struct{}

func (u *KindUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Kind)
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

func (u *KindUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Kind)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Use_symbols":
		GongUnmarshallEnum(&instance.Use_symbols, valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Stack_degrees":
		GongUnmarshallEnum(&instance.Stack_degrees, valueExpr)
	case "Parentheses_degrees":
		GongUnmarshallEnum(&instance.Parentheses_degrees, valueExpr)
	case "Bracket_degrees":
		GongUnmarshallEnum(&instance.Bracket_degrees, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type LevelUnmarshaller struct{}

func (u *LevelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Level)
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

func (u *LevelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Level)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference":
		GongUnmarshallEnum(&instance.Reference, valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Line_detailUnmarshaller struct{}

func (u *Line_detailUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Line_detail)
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

func (u *Line_detailUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line_detail)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	}
	return nil
}

type Line_widthUnmarshaller struct{}

func (u *Line_widthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Line_width)
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

func (u *Line_widthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line_width)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type LinkUnmarshaller struct{}

func (u *LinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Link)
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

func (u *LinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Position":
		instance.Position = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	}
	return nil
}

type ListenUnmarshaller struct{}

func (u *ListenUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Listen)
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

func (u *ListenUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Listen)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Assess":
		GongUnmarshallSliceOfPointers(&instance.Assess, valueExpr, identifierMap)
	case "Wait":
		GongUnmarshallSliceOfPointers(&instance.Wait, valueExpr, identifierMap)
	case "Other_listen":
		GongUnmarshallSliceOfPointers(&instance.Other_listen, valueExpr, identifierMap)
	}
	return nil
}

type ListeningUnmarshaller struct{}

func (u *ListeningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Listening)
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

func (u *ListeningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Listening)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sync":
		GongUnmarshallSliceOfPointers(&instance.Sync, valueExpr, identifierMap)
	case "Other_listening":
		GongUnmarshallSliceOfPointers(&instance.Other_listening, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	}
	return nil
}

type LyricUnmarshaller struct{}

func (u *LyricUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Lyric)
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

func (u *LyricUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Elision":
		GongUnmarshallSliceOfPointers(&instance.Elision, valueExpr, identifierMap)
	case "Syllabic":
		instance.Syllabic = GongExtractString(valueExpr)
	case "Text":
		GongUnmarshallSliceOfPointers(&instance.Text, valueExpr, identifierMap)
	case "Extend":
		GongUnmarshallPointer(&instance.Extend, valueExpr, identifierMap)
	case "Laughing":
		instance.Laughing = GongExtractString(valueExpr)
	case "Humming":
		instance.Humming = GongExtractString(valueExpr)
	case "End_line":
		instance.End_line = GongExtractString(valueExpr)
	case "End_paragraph":
		instance.End_paragraph = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Lyric_fontUnmarshaller struct{}

func (u *Lyric_fontUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Lyric_font)
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

func (u *Lyric_fontUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric_font)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	}
	return nil
}

type Lyric_languageUnmarshaller struct{}

func (u *Lyric_languageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Lyric_language)
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

func (u *Lyric_languageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric_language)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_layoutUnmarshaller struct{}

func (u *Measure_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Measure_layout)
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

func (u *Measure_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Measure_distance":
		instance.Measure_distance = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_numberingUnmarshaller struct{}

func (u *Measure_numberingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Measure_numbering)
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

func (u *Measure_numberingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_numbering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "System":
		instance.System = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Multiple_rest_always":
		GongUnmarshallEnum(&instance.Multiple_rest_always, valueExpr)
	case "Multiple_rest_range":
		GongUnmarshallEnum(&instance.Multiple_rest_range, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_repeatUnmarshaller struct{}

func (u *Measure_repeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Measure_repeat)
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

func (u *Measure_repeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_repeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Slashes":
		instance.Slashes = GongExtractInt(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_styleUnmarshaller struct{}

func (u *Measure_styleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Measure_style)
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

func (u *Measure_styleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_style)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Multiple_rest":
		GongUnmarshallPointer(&instance.Multiple_rest, valueExpr, identifierMap)
	case "Measure_repeat":
		GongUnmarshallPointer(&instance.Measure_repeat, valueExpr, identifierMap)
	case "Beat_repeat":
		GongUnmarshallPointer(&instance.Beat_repeat, valueExpr, identifierMap)
	case "Slash":
		GongUnmarshallPointer(&instance.Slash, valueExpr, identifierMap)
	}
	return nil
}

type MembraneUnmarshaller struct{}

func (u *MembraneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Membrane)
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

func (u *MembraneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Membrane)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type MetalUnmarshaller struct{}

func (u *MetalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Metal)
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

func (u *MetalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metal)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type MetronomeUnmarshaller struct{}

func (u *MetronomeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Metronome)
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

func (u *MetronomeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Beat_unit":
		GongUnmarshallEnum(&instance.Beat_unit, valueExpr)
	case "Beat_unit_dot":
		instance.Beat_unit_dot = GongExtractString(valueExpr)
	case "Per_minute":
		GongUnmarshallPointer(&instance.Per_minute, valueExpr, identifierMap)
	case "Beat_unit_tied":
		GongUnmarshallSliceOfPointers(&instance.Beat_unit_tied, valueExpr, identifierMap)
	case "Metronome_arrows":
		instance.Metronome_arrows = GongExtractString(valueExpr)
	case "Metronome_relation":
		instance.Metronome_relation = GongExtractString(valueExpr)
	case "Metronome_note":
		GongUnmarshallSliceOfPointers(&instance.Metronome_note, valueExpr, identifierMap)
	}
	return nil
}

type Metronome_beamUnmarshaller struct{}

func (u *Metronome_beamUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Metronome_beam)
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

func (u *Metronome_beamUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_beam)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Metronome_noteUnmarshaller struct{}

func (u *Metronome_noteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Metronome_note)
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

func (u *Metronome_noteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Metronome_type":
		instance.Metronome_type = GongExtractString(valueExpr)
	case "Metronome_dot":
		instance.Metronome_dot = GongExtractString(valueExpr)
	case "Metronome_beam":
		GongUnmarshallSliceOfPointers(&instance.Metronome_beam, valueExpr, identifierMap)
	case "Metronome_tied":
		GongUnmarshallPointer(&instance.Metronome_tied, valueExpr, identifierMap)
	case "Metronome_tuplet":
		GongUnmarshallPointer(&instance.Metronome_tuplet, valueExpr, identifierMap)
	}
	return nil
}

type Metronome_tiedUnmarshaller struct{}

func (u *Metronome_tiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Metronome_tied)
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

func (u *Metronome_tiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_tied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	}
	return nil
}

type Metronome_tupletUnmarshaller struct{}

func (u *Metronome_tupletUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Metronome_tuplet)
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

func (u *Metronome_tupletUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_tuplet)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Midi_deviceUnmarshaller struct{}

func (u *Midi_deviceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Midi_device)
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

func (u *Midi_deviceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Midi_device)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Port":
		instance.Port = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Midi_instrumentUnmarshaller struct{}

func (u *Midi_instrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Midi_instrument)
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

func (u *Midi_instrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Midi_instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Midi_channel":
		instance.Midi_channel = GongExtractInt(valueExpr)
	case "Midi_name":
		instance.Midi_name = GongExtractString(valueExpr)
	case "Midi_bank":
		instance.Midi_bank = GongExtractInt(valueExpr)
	case "Midi_program":
		instance.Midi_program = GongExtractInt(valueExpr)
	case "Midi_unpitched":
		instance.Midi_unpitched = GongExtractInt(valueExpr)
	case "Volume":
		instance.Volume = GongExtractString(valueExpr)
	case "Pan":
		instance.Pan = GongExtractString(valueExpr)
	case "Elevation":
		instance.Elevation = GongExtractString(valueExpr)
	}
	return nil
}

type MiscellaneousUnmarshaller struct{}

func (u *MiscellaneousUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Miscellaneous)
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

func (u *MiscellaneousUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Miscellaneous)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Miscellaneous_field":
		GongUnmarshallSliceOfPointers(&instance.Miscellaneous_field, valueExpr, identifierMap)
	}
	return nil
}

type Miscellaneous_fieldUnmarshaller struct{}

func (u *Miscellaneous_fieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Miscellaneous_field)
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

func (u *Miscellaneous_fieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Miscellaneous_field)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type MordentUnmarshaller struct{}

func (u *MordentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Mordent)
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

func (u *MordentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Mordent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Multiple_restUnmarshaller struct{}

func (u *Multiple_restUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Multiple_rest)
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

func (u *Multiple_restUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Multiple_rest)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Use_symbols":
		GongUnmarshallEnum(&instance.Use_symbols, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Name_displayUnmarshaller struct{}

func (u *Name_displayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Name_display)
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

func (u *Name_displayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Name_display)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Display_text":
		GongUnmarshallSliceOfPointers(&instance.Display_text, valueExpr, identifierMap)
	case "Accidental_text":
		GongUnmarshallSliceOfPointers(&instance.Accidental_text, valueExpr, identifierMap)
	}
	return nil
}

type Non_arpeggiateUnmarshaller struct{}

func (u *Non_arpeggiateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Non_arpeggiate)
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

func (u *Non_arpeggiateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Non_arpeggiate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type NotationsUnmarshaller struct{}

func (u *NotationsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Notations)
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

func (u *NotationsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notations)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Tied":
		GongUnmarshallSliceOfPointers(&instance.Tied, valueExpr, identifierMap)
	case "Slur":
		GongUnmarshallSliceOfPointers(&instance.Slur, valueExpr, identifierMap)
	case "Tuplet":
		GongUnmarshallSliceOfPointers(&instance.Tuplet, valueExpr, identifierMap)
	case "Glissando":
		GongUnmarshallSliceOfPointers(&instance.Glissando, valueExpr, identifierMap)
	case "Slide":
		GongUnmarshallSliceOfPointers(&instance.Slide, valueExpr, identifierMap)
	case "Ornaments":
		GongUnmarshallSliceOfPointers(&instance.Ornaments, valueExpr, identifierMap)
	case "Technical":
		GongUnmarshallSliceOfPointers(&instance.Technical, valueExpr, identifierMap)
	case "Articulations":
		GongUnmarshallSliceOfPointers(&instance.Articulations, valueExpr, identifierMap)
	case "Dynamics":
		GongUnmarshallSliceOfPointers(&instance.Dynamics, valueExpr, identifierMap)
	case "Fermata":
		GongUnmarshallSliceOfPointers(&instance.Fermata, valueExpr, identifierMap)
	case "Arpeggiate":
		GongUnmarshallSliceOfPointers(&instance.Arpeggiate, valueExpr, identifierMap)
	case "Non_arpeggiate":
		GongUnmarshallSliceOfPointers(&instance.Non_arpeggiate, valueExpr, identifierMap)
	case "Accidental_mark":
		GongUnmarshallSliceOfPointers(&instance.Accidental_mark, valueExpr, identifierMap)
	case "Other_notation":
		GongUnmarshallSliceOfPointers(&instance.Other_notation, valueExpr, identifierMap)
	}
	return nil
}

type NoteUnmarshaller struct{}

func (u *NoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Note)
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

func (u *NoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_leger":
		GongUnmarshallEnum(&instance.Print_leger, valueExpr)
	case "Dynamics":
		instance.Dynamics = GongExtractString(valueExpr)
	case "End_dynamics":
		instance.End_dynamics = GongExtractString(valueExpr)
	case "Attack":
		instance.Attack = GongExtractString(valueExpr)
	case "Release":
		instance.Release = GongExtractString(valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	case "Pizzicato":
		GongUnmarshallEnum(&instance.Pizzicato, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_dot":
		GongUnmarshallEnum(&instance.Print_dot, valueExpr)
	case "Print_lyric":
		GongUnmarshallEnum(&instance.Print_lyric, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Print_spacing":
		GongUnmarshallEnum(&instance.Print_spacing, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Grace":
		GongUnmarshallPointer(&instance.Grace, valueExpr, identifierMap)
	case "Chord":
		instance.Chord = GongExtractString(valueExpr)
	case "Pitch":
		GongUnmarshallPointer(&instance.Pitch, valueExpr, identifierMap)
	case "Unpitched":
		GongUnmarshallPointer(&instance.Unpitched, valueExpr, identifierMap)
	case "Rest":
		GongUnmarshallPointer(&instance.Rest, valueExpr, identifierMap)
	case "Tie":
		GongUnmarshallPointer(&instance.Tie, valueExpr, identifierMap)
	case "Cue":
		instance.Cue = GongExtractString(valueExpr)
	case "Duration":
		instance.Duration = GongExtractString(valueExpr)
	case "Instrument":
		GongUnmarshallSliceOfPointers(&instance.Instrument, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallPointer(&instance.Type, valueExpr, identifierMap)
	case "Dot":
		GongUnmarshallSliceOfPointers(&instance.Dot, valueExpr, identifierMap)
	case "Accidental":
		GongUnmarshallPointer(&instance.Accidental, valueExpr, identifierMap)
	case "Time_modification":
		GongUnmarshallPointer(&instance.Time_modification, valueExpr, identifierMap)
	case "Stem":
		GongUnmarshallPointer(&instance.Stem, valueExpr, identifierMap)
	case "Notehead":
		GongUnmarshallPointer(&instance.Notehead, valueExpr, identifierMap)
	case "Notehead_text":
		GongUnmarshallPointer(&instance.Notehead_text, valueExpr, identifierMap)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Beam":
		GongUnmarshallPointer(&instance.Beam, valueExpr, identifierMap)
	case "Notations":
		GongUnmarshallSliceOfPointers(&instance.Notations, valueExpr, identifierMap)
	case "Lyric":
		GongUnmarshallSliceOfPointers(&instance.Lyric, valueExpr, identifierMap)
	case "Play":
		GongUnmarshallPointer(&instance.Play, valueExpr, identifierMap)
	case "Listen":
		GongUnmarshallPointer(&instance.Listen, valueExpr, identifierMap)
	}
	return nil
}

type Note_sizeUnmarshaller struct{}

func (u *Note_sizeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Note_size)
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

func (u *Note_sizeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note_size)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Note_typeUnmarshaller struct{}

func (u *Note_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Note_type)
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

func (u *Note_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type NoteheadUnmarshaller struct{}

func (u *NoteheadUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Notehead)
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

func (u *NoteheadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notehead)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Filled":
		GongUnmarshallEnum(&instance.Filled, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Notehead_textUnmarshaller struct{}

func (u *Notehead_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Notehead_text)
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

func (u *Notehead_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notehead_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Display_text":
		GongUnmarshallSliceOfPointers(&instance.Display_text, valueExpr, identifierMap)
	case "Accidental_text":
		GongUnmarshallSliceOfPointers(&instance.Accidental_text, valueExpr, identifierMap)
	}
	return nil
}

type NumeralUnmarshaller struct{}

func (u *NumeralUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Numeral)
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

func (u *NumeralUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Numeral_root":
		GongUnmarshallPointer(&instance.Numeral_root, valueExpr, identifierMap)
	case "Numeral_alter":
		GongUnmarshallPointer(&instance.Numeral_alter, valueExpr, identifierMap)
	case "Numeral_key":
		GongUnmarshallPointer(&instance.Numeral_key, valueExpr, identifierMap)
	}
	return nil
}

type Numeral_keyUnmarshaller struct{}

func (u *Numeral_keyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Numeral_key)
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

func (u *Numeral_keyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral_key)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Numeral_fifths":
		instance.Numeral_fifths = GongExtractInt(valueExpr)
	case "Numeral_mode":
		instance.Numeral_mode = GongExtractString(valueExpr)
	}
	return nil
}

type Numeral_rootUnmarshaller struct{}

func (u *Numeral_rootUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Numeral_root)
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

func (u *Numeral_rootUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral_root)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Octave_shiftUnmarshaller struct{}

func (u *Octave_shiftUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Octave_shift)
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

func (u *Octave_shiftUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Octave_shift)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Size":
		instance.Size = GongExtractInt(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type OffsetUnmarshaller struct{}

func (u *OffsetUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Offset)
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

func (u *OffsetUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Offset)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sound":
		GongUnmarshallEnum(&instance.Sound, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type OpusUnmarshaller struct{}

func (u *OpusUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Opus)
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

func (u *OpusUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Opus)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type OrnamentsUnmarshaller struct{}

func (u *OrnamentsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Ornaments)
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

func (u *OrnamentsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ornaments)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Trill_mark":
		GongUnmarshallSliceOfPointers(&instance.Trill_mark, valueExpr, identifierMap)
	case "Turn":
		GongUnmarshallSliceOfPointers(&instance.Turn, valueExpr, identifierMap)
	case "Delayed_turn":
		GongUnmarshallSliceOfPointers(&instance.Delayed_turn, valueExpr, identifierMap)
	case "Inverted_turn":
		GongUnmarshallSliceOfPointers(&instance.Inverted_turn, valueExpr, identifierMap)
	case "Delayed_inverted_turn":
		GongUnmarshallSliceOfPointers(&instance.Delayed_inverted_turn, valueExpr, identifierMap)
	case "Vertical_turn":
		GongUnmarshallSliceOfPointers(&instance.Vertical_turn, valueExpr, identifierMap)
	case "Inverted_vertical_turn":
		GongUnmarshallSliceOfPointers(&instance.Inverted_vertical_turn, valueExpr, identifierMap)
	case "Shake":
		GongUnmarshallSliceOfPointers(&instance.Shake, valueExpr, identifierMap)
	case "Wavy_line":
		GongUnmarshallSliceOfPointers(&instance.Wavy_line, valueExpr, identifierMap)
	case "Mordent":
		GongUnmarshallSliceOfPointers(&instance.Mordent, valueExpr, identifierMap)
	case "Inverted_mordent":
		GongUnmarshallSliceOfPointers(&instance.Inverted_mordent, valueExpr, identifierMap)
	case "Schleifer":
		GongUnmarshallSliceOfPointers(&instance.Schleifer, valueExpr, identifierMap)
	case "Tremolo":
		GongUnmarshallSliceOfPointers(&instance.Tremolo, valueExpr, identifierMap)
	case "Haydn":
		GongUnmarshallSliceOfPointers(&instance.Haydn, valueExpr, identifierMap)
	case "Other_ornament":
		GongUnmarshallSliceOfPointers(&instance.Other_ornament, valueExpr, identifierMap)
	case "Accidental_mark":
		GongUnmarshallSliceOfPointers(&instance.Accidental_mark, valueExpr, identifierMap)
	}
	return nil
}

type Other_appearanceUnmarshaller struct{}

func (u *Other_appearanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_appearance)
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

func (u *Other_appearanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_appearance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_directionUnmarshaller struct{}

func (u *Other_directionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_direction)
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

func (u *Other_directionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_direction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_listeningUnmarshaller struct{}

func (u *Other_listeningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_listening)
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

func (u *Other_listeningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_listening)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_notationUnmarshaller struct{}

func (u *Other_notationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_notation)
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

func (u *Other_notationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_notation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_placement_textUnmarshaller struct{}

func (u *Other_placement_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_placement_text)
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

func (u *Other_placement_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_placement_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_playUnmarshaller struct{}

func (u *Other_playUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_play)
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

func (u *Other_playUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_play)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_textUnmarshaller struct{}

func (u *Other_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Other_text)
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

func (u *Other_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Page_layoutUnmarshaller struct{}

func (u *Page_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Page_layout)
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

func (u *Page_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Page_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Page_height":
		instance.Page_height = GongExtractString(valueExpr)
	case "Page_width":
		instance.Page_width = GongExtractString(valueExpr)
	case "Page_margins":
		GongUnmarshallPointer(&instance.Page_margins, valueExpr, identifierMap)
	}
	return nil
}

type Page_marginsUnmarshaller struct{}

func (u *Page_marginsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Page_margins)
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

func (u *Page_marginsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Page_margins)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Left_margin":
		instance.Left_margin = GongExtractString(valueExpr)
	case "Right_margin":
		instance.Right_margin = GongExtractString(valueExpr)
	case "Top_margin":
		instance.Top_margin = GongExtractString(valueExpr)
	case "Bottom_margin":
		instance.Bottom_margin = GongExtractString(valueExpr)
	}
	return nil
}

type Part_clefUnmarshaller struct{}

func (u *Part_clefUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_clef)
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

func (u *Part_clefUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_clef)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sign":
		instance.Sign = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Clef_octave_change":
		instance.Clef_octave_change = GongExtractInt(valueExpr)
	}
	return nil
}

type Part_groupUnmarshaller struct{}

func (u *Part_groupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_group)
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

func (u *Part_groupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Group_name":
		GongUnmarshallPointer(&instance.Group_name, valueExpr, identifierMap)
	case "Group_name_display":
		GongUnmarshallPointer(&instance.Group_name_display, valueExpr, identifierMap)
	case "Group_abbreviation":
		GongUnmarshallPointer(&instance.Group_abbreviation, valueExpr, identifierMap)
	case "Group_abbreviation_display":
		GongUnmarshallPointer(&instance.Group_abbreviation_display, valueExpr, identifierMap)
	case "Group_symbol":
		GongUnmarshallPointer(&instance.Group_symbol, valueExpr, identifierMap)
	case "Group_barline":
		GongUnmarshallPointer(&instance.Group_barline, valueExpr, identifierMap)
	case "Group_time":
		instance.Group_time = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Part_linkUnmarshaller struct{}

func (u *Part_linkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_link)
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

func (u *Part_linkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	case "Instrument_link":
		GongUnmarshallSliceOfPointers(&instance.Instrument_link, valueExpr, identifierMap)
	case "Group_link":
		instance.Group_link = GongExtractString(valueExpr)
	}
	return nil
}

type Part_listUnmarshaller struct{}

func (u *Part_listUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_list)
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

func (u *Part_listUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_list)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part_group":
		GongUnmarshallPointer(&instance.Part_group, valueExpr, identifierMap)
	case "Score_part":
		GongUnmarshallPointer(&instance.Score_part, valueExpr, identifierMap)
	}
	return nil
}

type Part_nameUnmarshaller struct{}

func (u *Part_nameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_name)
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

func (u *Part_nameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_name)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Part_symbolUnmarshaller struct{}

func (u *Part_symbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_symbol)
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

func (u *Part_symbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_symbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Top_staff":
		instance.Top_staff = GongExtractInt(valueExpr)
	case "Bottom_staff":
		instance.Bottom_staff = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Part_transposeUnmarshaller struct{}

func (u *Part_transposeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Part_transpose)
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

func (u *Part_transposeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_transpose)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Diatonic":
		instance.Diatonic = GongExtractInt(valueExpr)
	case "Chromatic":
		instance.Chromatic = GongExtractString(valueExpr)
	case "Octave_change":
		instance.Octave_change = GongExtractInt(valueExpr)
	case "Double":
		instance.Double = GongExtractFloat(valueExpr)
	}
	return nil
}

type PedalUnmarshaller struct{}

func (u *PedalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Pedal)
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

func (u *PedalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pedal)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line":
		GongUnmarshallEnum(&instance.Line, valueExpr)
	case "Sign":
		GongUnmarshallEnum(&instance.Sign, valueExpr)
	case "Abbreviated":
		GongUnmarshallEnum(&instance.Abbreviated, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Pedal_tuningUnmarshaller struct{}

func (u *Pedal_tuningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Pedal_tuning)
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

func (u *Pedal_tuningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pedal_tuning)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Pedal_step":
		GongUnmarshallEnum(&instance.Pedal_step, valueExpr)
	case "Pedal_alter":
		instance.Pedal_alter = GongExtractString(valueExpr)
	}
	return nil
}

type Per_minuteUnmarshaller struct{}

func (u *Per_minuteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Per_minute)
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

func (u *Per_minuteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Per_minute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type PercussionUnmarshaller struct{}

func (u *PercussionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Percussion)
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

func (u *PercussionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Percussion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Enclosure":
		instance.Enclosure = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Glass":
		GongUnmarshallPointer(&instance.Glass, valueExpr, identifierMap)
	case "Metal":
		GongUnmarshallPointer(&instance.Metal, valueExpr, identifierMap)
	case "Wood":
		GongUnmarshallPointer(&instance.Wood, valueExpr, identifierMap)
	case "Pitched":
		GongUnmarshallPointer(&instance.Pitched, valueExpr, identifierMap)
	case "Membrane":
		GongUnmarshallPointer(&instance.Membrane, valueExpr, identifierMap)
	case "Effect":
		GongUnmarshallPointer(&instance.Effect, valueExpr, identifierMap)
	case "Timpani":
		GongUnmarshallPointer(&instance.Timpani, valueExpr, identifierMap)
	case "Beater":
		GongUnmarshallPointer(&instance.Beater, valueExpr, identifierMap)
	case "Stick":
		GongUnmarshallPointer(&instance.Stick, valueExpr, identifierMap)
	case "Stick_location":
		instance.Stick_location = GongExtractString(valueExpr)
	case "Other_percussion":
		GongUnmarshallPointer(&instance.Other_percussion, valueExpr, identifierMap)
	}
	return nil
}

type PitchUnmarshaller struct{}

func (u *PitchUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Pitch)
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

func (u *PitchUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pitch)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Step":
		GongUnmarshallEnum(&instance.Step, valueExpr)
	case "Alter":
		instance.Alter = GongExtractString(valueExpr)
	case "Octave":
		instance.Octave = GongExtractInt(valueExpr)
	}
	return nil
}

type PitchedUnmarshaller struct{}

func (u *PitchedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Pitched)
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

func (u *PitchedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pitched)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Placement_textUnmarshaller struct{}

func (u *Placement_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Placement_text)
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

func (u *Placement_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Placement_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type PlayUnmarshaller struct{}

func (u *PlayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Play)
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

func (u *PlayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Play)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Ipa":
		instance.Ipa = GongExtractString(valueExpr)
	case "Mute":
		instance.Mute = GongExtractString(valueExpr)
	case "Semi_pitched":
		instance.Semi_pitched = GongExtractString(valueExpr)
	case "Other_play":
		GongUnmarshallSliceOfPointers(&instance.Other_play, valueExpr, identifierMap)
	}
	return nil
}

type PlayerUnmarshaller struct{}

func (u *PlayerUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Player)
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

func (u *PlayerUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Player)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Player_name":
		instance.Player_name = GongExtractString(valueExpr)
	}
	return nil
}

type Principal_voiceUnmarshaller struct{}

func (u *Principal_voiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Principal_voice)
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

func (u *Principal_voiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Principal_voice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Symbol":
		instance.Symbol = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type PrintUnmarshaller struct{}

func (u *PrintUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Print)
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

func (u *PrintUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Print)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Staff_spacing":
		instance.Staff_spacing = GongExtractString(valueExpr)
	case "New_system":
		GongUnmarshallEnum(&instance.New_system, valueExpr)
	case "New_page":
		GongUnmarshallEnum(&instance.New_page, valueExpr)
	case "Blank_page":
		instance.Blank_page = GongExtractInt(valueExpr)
	case "Page_number":
		instance.Page_number = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Page_layout":
		GongUnmarshallPointer(&instance.Page_layout, valueExpr, identifierMap)
	case "System_layout":
		GongUnmarshallPointer(&instance.System_layout, valueExpr, identifierMap)
	case "Staff_layout":
		GongUnmarshallSliceOfPointers(&instance.Staff_layout, valueExpr, identifierMap)
	case "Measure_layout":
		GongUnmarshallPointer(&instance.Measure_layout, valueExpr, identifierMap)
	case "Measure_numbering":
		GongUnmarshallPointer(&instance.Measure_numbering, valueExpr, identifierMap)
	case "Part_name_display":
		GongUnmarshallPointer(&instance.Part_name_display, valueExpr, identifierMap)
	case "Part_abbreviation_display":
		GongUnmarshallPointer(&instance.Part_abbreviation_display, valueExpr, identifierMap)
	}
	return nil
}

type ReleaseUnmarshaller struct{}

func (u *ReleaseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Release)
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

func (u *ReleaseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Release)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type RepeatUnmarshaller struct{}

func (u *RepeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Repeat)
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

func (u *RepeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Repeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Direction":
		instance.Direction = GongExtractString(valueExpr)
	case "Times":
		instance.Times = GongExtractInt(valueExpr)
	case "After_jump":
		GongUnmarshallEnum(&instance.After_jump, valueExpr)
	case "Winged":
		instance.Winged = GongExtractString(valueExpr)
	}
	return nil
}

type RestUnmarshaller struct{}

func (u *RestUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Rest)
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

func (u *RestUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rest)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Measure":
		GongUnmarshallEnum(&instance.Measure, valueExpr)
	case "Display_step":
		GongUnmarshallEnum(&instance.Display_step, valueExpr)
	case "Display_octave":
		instance.Display_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type RootUnmarshaller struct{}

func (u *RootUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Root)
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

func (u *RootUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Root)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Root_step":
		GongUnmarshallPointer(&instance.Root_step, valueExpr, identifierMap)
	case "Root_alter":
		GongUnmarshallPointer(&instance.Root_alter, valueExpr, identifierMap)
	}
	return nil
}

type Root_stepUnmarshaller struct{}

func (u *Root_stepUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Root_step)
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

func (u *Root_stepUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Root_step)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type ScalingUnmarshaller struct{}

func (u *ScalingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Scaling)
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

func (u *ScalingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scaling)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Millimeters":
		instance.Millimeters = GongExtractString(valueExpr)
	case "Tenths":
		instance.Tenths = GongExtractString(valueExpr)
	}
	return nil
}

type ScordaturaUnmarshaller struct{}

func (u *ScordaturaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Scordatura)
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

func (u *ScordaturaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scordatura)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accord":
		GongUnmarshallSliceOfPointers(&instance.Accord, valueExpr, identifierMap)
	}
	return nil
}

type Score_instrumentUnmarshaller struct{}

func (u *Score_instrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Score_instrument)
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

func (u *Score_instrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Instrument_name":
		instance.Instrument_name = GongExtractString(valueExpr)
	case "Instrument_abbreviation":
		instance.Instrument_abbreviation = GongExtractString(valueExpr)
	case "Instrument_sound":
		instance.Instrument_sound = GongExtractString(valueExpr)
	case "Solo":
		instance.Solo = GongExtractString(valueExpr)
	case "Ensemble":
		instance.Ensemble = GongExtractString(valueExpr)
	case "Virtual_instrument":
		GongUnmarshallPointer(&instance.Virtual_instrument, valueExpr, identifierMap)
	}
	return nil
}

type Score_partUnmarshaller struct{}

func (u *Score_partUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Score_part)
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

func (u *Score_partUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Identification":
		GongUnmarshallPointer(&instance.Identification, valueExpr, identifierMap)
	case "Part_link":
		GongUnmarshallSliceOfPointers(&instance.Part_link, valueExpr, identifierMap)
	case "Part_name":
		GongUnmarshallPointer(&instance.Part_name, valueExpr, identifierMap)
	case "Part_name_display":
		GongUnmarshallPointer(&instance.Part_name_display, valueExpr, identifierMap)
	case "Part_abbreviation":
		GongUnmarshallPointer(&instance.Part_abbreviation, valueExpr, identifierMap)
	case "Part_abbreviation_display":
		GongUnmarshallPointer(&instance.Part_abbreviation_display, valueExpr, identifierMap)
	case "Group":
		instance.Group = GongExtractString(valueExpr)
	case "Score_instrument":
		GongUnmarshallSliceOfPointers(&instance.Score_instrument, valueExpr, identifierMap)
	case "Player":
		GongUnmarshallSliceOfPointers(&instance.Player, valueExpr, identifierMap)
	case "Midi_device":
		GongUnmarshallSliceOfPointers(&instance.Midi_device, valueExpr, identifierMap)
	case "Midi_instrument":
		GongUnmarshallSliceOfPointers(&instance.Midi_instrument, valueExpr, identifierMap)
	}
	return nil
}

type Score_partwiseUnmarshaller struct{}

func (u *Score_partwiseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Score_partwise)
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

func (u *Score_partwiseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_partwise)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Version":
		instance.Version = GongExtractString(valueExpr)
	case "Work":
		GongUnmarshallPointer(&instance.Work, valueExpr, identifierMap)
	case "Movement_number":
		instance.Movement_number = GongExtractString(valueExpr)
	case "Movement_title":
		instance.Movement_title = GongExtractString(valueExpr)
	case "Identification":
		GongUnmarshallPointer(&instance.Identification, valueExpr, identifierMap)
	case "Defaults":
		GongUnmarshallPointer(&instance.Defaults, valueExpr, identifierMap)
	case "Credit":
		GongUnmarshallSliceOfPointers(&instance.Credit, valueExpr, identifierMap)
	case "Part_list":
		GongUnmarshallPointer(&instance.Part_list, valueExpr, identifierMap)
	case "Part":
		GongUnmarshallSliceOfPointers(&instance.Part, valueExpr, identifierMap)
	}
	return nil
}

type Score_timewiseUnmarshaller struct{}

func (u *Score_timewiseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Score_timewise)
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

func (u *Score_timewiseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_timewise)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Version":
		instance.Version = GongExtractString(valueExpr)
	case "Work":
		GongUnmarshallPointer(&instance.Work, valueExpr, identifierMap)
	case "Movement_number":
		instance.Movement_number = GongExtractString(valueExpr)
	case "Movement_title":
		instance.Movement_title = GongExtractString(valueExpr)
	case "Identification":
		GongUnmarshallPointer(&instance.Identification, valueExpr, identifierMap)
	case "Defaults":
		GongUnmarshallPointer(&instance.Defaults, valueExpr, identifierMap)
	case "Credit":
		GongUnmarshallSliceOfPointers(&instance.Credit, valueExpr, identifierMap)
	case "Part_list":
		GongUnmarshallPointer(&instance.Part_list, valueExpr, identifierMap)
	case "Measure":
		GongUnmarshallSliceOfPointers(&instance.Measure, valueExpr, identifierMap)
	}
	return nil
}

type SegnoUnmarshaller struct{}

func (u *SegnoUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Segno)
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

func (u *SegnoUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Segno)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type SlashUnmarshaller struct{}

func (u *SlashUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Slash)
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

func (u *SlashUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slash)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Use_dots":
		GongUnmarshallEnum(&instance.Use_dots, valueExpr)
	case "Use_stems":
		GongUnmarshallEnum(&instance.Use_stems, valueExpr)
	case "Slash_type":
		GongUnmarshallEnum(&instance.Slash_type, valueExpr)
	case "Slash_dot":
		instance.Slash_dot = GongExtractString(valueExpr)
	case "Except_voice":
		instance.Except_voice = GongExtractString(valueExpr)
	}
	return nil
}

type SlideUnmarshaller struct{}

func (u *SlideUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Slide)
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

func (u *SlideUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slide)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "First_beat":
		instance.First_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type SlurUnmarshaller struct{}

func (u *SlurUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Slur)
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

func (u *SlurUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slur)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Orientation":
		instance.Orientation = GongExtractString(valueExpr)
	case "Bezier_x":
		instance.Bezier_x = GongExtractString(valueExpr)
	case "Bezier_y":
		instance.Bezier_y = GongExtractString(valueExpr)
	case "Bezier_x2":
		instance.Bezier_x2 = GongExtractString(valueExpr)
	case "Bezier_y2":
		instance.Bezier_y2 = GongExtractString(valueExpr)
	case "Bezier_offset":
		instance.Bezier_offset = GongExtractString(valueExpr)
	case "Bezier_offset2":
		instance.Bezier_offset2 = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type SoundUnmarshaller struct{}

func (u *SoundUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Sound)
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

func (u *SoundUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sound)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tempo":
		instance.Tempo = GongExtractString(valueExpr)
	case "Dynamics":
		instance.Dynamics = GongExtractString(valueExpr)
	case "Dacapo":
		GongUnmarshallEnum(&instance.Dacapo, valueExpr)
	case "Segno":
		instance.Segno = GongExtractString(valueExpr)
	case "Dalsegno":
		instance.Dalsegno = GongExtractString(valueExpr)
	case "Coda":
		instance.Coda = GongExtractString(valueExpr)
	case "Tocoda":
		instance.Tocoda = GongExtractString(valueExpr)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Forward_repeat":
		GongUnmarshallEnum(&instance.Forward_repeat, valueExpr)
	case "Fine":
		instance.Fine = GongExtractString(valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	case "Pizzicato":
		GongUnmarshallEnum(&instance.Pizzicato, valueExpr)
	case "Pan":
		instance.Pan = GongExtractString(valueExpr)
	case "Elevation":
		instance.Elevation = GongExtractString(valueExpr)
	case "Damper_pedal":
		instance.Damper_pedal = GongExtractString(valueExpr)
	case "Soft_pedal":
		instance.Soft_pedal = GongExtractString(valueExpr)
	case "Sostenuto_pedal":
		instance.Sostenuto_pedal = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Instrument_change":
		GongUnmarshallSliceOfPointers(&instance.Instrument_change, valueExpr, identifierMap)
	case "Midi_device":
		GongUnmarshallSliceOfPointers(&instance.Midi_device, valueExpr, identifierMap)
	case "Midi_instrument":
		GongUnmarshallSliceOfPointers(&instance.Midi_instrument, valueExpr, identifierMap)
	case "Play":
		GongUnmarshallSliceOfPointers(&instance.Play, valueExpr, identifierMap)
	case "Swing":
		GongUnmarshallPointer(&instance.Swing, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	}
	return nil
}

type Staff_detailsUnmarshaller struct{}

func (u *Staff_detailsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Staff_details)
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

func (u *Staff_detailsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_details)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Show_frets":
		instance.Show_frets = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Print_spacing":
		GongUnmarshallEnum(&instance.Print_spacing, valueExpr)
	case "Staff_type":
		instance.Staff_type = GongExtractString(valueExpr)
	case "Staff_lines":
		instance.Staff_lines = GongExtractInt(valueExpr)
	case "Line_detail":
		GongUnmarshallSliceOfPointers(&instance.Line_detail, valueExpr, identifierMap)
	case "Staff_tuning":
		GongUnmarshallSliceOfPointers(&instance.Staff_tuning, valueExpr, identifierMap)
	case "Capo":
		instance.Capo = GongExtractInt(valueExpr)
	case "Staff_size":
		GongUnmarshallPointer(&instance.Staff_size, valueExpr, identifierMap)
	}
	return nil
}

type Staff_divideUnmarshaller struct{}

func (u *Staff_divideUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Staff_divide)
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

func (u *Staff_divideUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_divide)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Staff_layoutUnmarshaller struct{}

func (u *Staff_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Staff_layout)
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

func (u *Staff_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Staff_distance":
		instance.Staff_distance = GongExtractString(valueExpr)
	}
	return nil
}

type Staff_sizeUnmarshaller struct{}

func (u *Staff_sizeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Staff_size)
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

func (u *Staff_sizeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_size)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Scaling":
		instance.Scaling = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Staff_tuningUnmarshaller struct{}

func (u *Staff_tuningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Staff_tuning)
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

func (u *Staff_tuningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_tuning)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Tuning_step":
		GongUnmarshallEnum(&instance.Tuning_step, valueExpr)
	case "Tuning_alter":
		instance.Tuning_alter = GongExtractString(valueExpr)
	case "Tuning_octave":
		instance.Tuning_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type StemUnmarshaller struct{}

func (u *StemUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Stem)
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

func (u *StemUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stem)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type StickUnmarshaller struct{}

func (u *StickUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Stick)
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

func (u *StickUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stick)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tip":
		GongUnmarshallEnum(&instance.Tip, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Dashed_circle":
		GongUnmarshallEnum(&instance.Dashed_circle, valueExpr)
	case "Stick_type":
		instance.Stick_type = GongExtractString(valueExpr)
	case "Stick_material":
		instance.Stick_material = GongExtractString(valueExpr)
	}
	return nil
}

type String_muteUnmarshaller struct{}

func (u *String_muteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(String_mute)
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

func (u *String_muteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*String_mute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type String_typeUnmarshaller struct{}

func (u *String_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(String_type)
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

func (u *String_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*String_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Strong_accentUnmarshaller struct{}

func (u *Strong_accentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Strong_accent)
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

func (u *Strong_accentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Strong_accent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Style_textUnmarshaller struct{}

func (u *Style_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Style_text)
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

func (u *Style_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Style_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type SupportsUnmarshaller struct{}

func (u *SupportsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Supports)
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

func (u *SupportsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Supports)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Attribute":
		instance.Attribute = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type SwingUnmarshaller struct{}

func (u *SwingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Swing)
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

func (u *SwingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Swing)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Straight":
		instance.Straight = GongExtractString(valueExpr)
	case "First":
		instance.First = GongExtractInt(valueExpr)
	case "Second":
		instance.Second = GongExtractInt(valueExpr)
	case "Swing_type":
		GongUnmarshallEnum(&instance.Swing_type, valueExpr)
	case "Swing_style":
		instance.Swing_style = GongExtractString(valueExpr)
	}
	return nil
}

type SyncUnmarshaller struct{}

func (u *SyncUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Sync)
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

func (u *SyncUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sync)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Latency":
		instance.Latency = GongExtractInt(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	}
	return nil
}

type System_dividersUnmarshaller struct{}

func (u *System_dividersUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(System_dividers)
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

func (u *System_dividersUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System_dividers)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Left_divider":
		GongUnmarshallPointer(&instance.Left_divider, valueExpr, identifierMap)
	case "Right_divider":
		GongUnmarshallPointer(&instance.Right_divider, valueExpr, identifierMap)
	}
	return nil
}

type System_layoutUnmarshaller struct{}

func (u *System_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(System_layout)
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

func (u *System_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "System_margins":
		GongUnmarshallPointer(&instance.System_margins, valueExpr, identifierMap)
	case "System_distance":
		instance.System_distance = GongExtractString(valueExpr)
	case "Top_system_distance":
		instance.Top_system_distance = GongExtractString(valueExpr)
	case "System_dividers":
		GongUnmarshallPointer(&instance.System_dividers, valueExpr, identifierMap)
	}
	return nil
}

type System_marginsUnmarshaller struct{}

func (u *System_marginsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(System_margins)
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

func (u *System_marginsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System_margins)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Left_margin":
		instance.Left_margin = GongExtractString(valueExpr)
	case "Right_margin":
		instance.Right_margin = GongExtractString(valueExpr)
	}
	return nil
}

type TapUnmarshaller struct{}

func (u *TapUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tap)
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

func (u *TapUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tap)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Hand":
		instance.Hand = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TechnicalUnmarshaller struct{}

func (u *TechnicalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Technical)
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

func (u *TechnicalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Technical)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Up_bow":
		GongUnmarshallSliceOfPointers(&instance.Up_bow, valueExpr, identifierMap)
	case "Down_bow":
		GongUnmarshallSliceOfPointers(&instance.Down_bow, valueExpr, identifierMap)
	case "Harmonic":
		GongUnmarshallSliceOfPointers(&instance.Harmonic, valueExpr, identifierMap)
	case "Open_string":
		GongUnmarshallSliceOfPointers(&instance.Open_string, valueExpr, identifierMap)
	case "Thumb_position":
		GongUnmarshallSliceOfPointers(&instance.Thumb_position, valueExpr, identifierMap)
	case "Fingering":
		GongUnmarshallSliceOfPointers(&instance.Fingering, valueExpr, identifierMap)
	case "Pluck":
		GongUnmarshallSliceOfPointers(&instance.Pluck, valueExpr, identifierMap)
	case "Double_tongue":
		GongUnmarshallSliceOfPointers(&instance.Double_tongue, valueExpr, identifierMap)
	case "Triple_tongue":
		GongUnmarshallSliceOfPointers(&instance.Triple_tongue, valueExpr, identifierMap)
	case "Stopped":
		GongUnmarshallSliceOfPointers(&instance.Stopped, valueExpr, identifierMap)
	case "Snap_pizzicato":
		GongUnmarshallSliceOfPointers(&instance.Snap_pizzicato, valueExpr, identifierMap)
	case "Fret":
		GongUnmarshallSliceOfPointers(&instance.Fret, valueExpr, identifierMap)
	case "String":
		GongUnmarshallSliceOfPointers(&instance.String, valueExpr, identifierMap)
	case "Hammer_on":
		GongUnmarshallSliceOfPointers(&instance.Hammer_on, valueExpr, identifierMap)
	case "Pull_off":
		GongUnmarshallSliceOfPointers(&instance.Pull_off, valueExpr, identifierMap)
	case "Bend":
		GongUnmarshallSliceOfPointers(&instance.Bend, valueExpr, identifierMap)
	case "Tap":
		GongUnmarshallSliceOfPointers(&instance.Tap, valueExpr, identifierMap)
	case "Heel":
		GongUnmarshallSliceOfPointers(&instance.Heel, valueExpr, identifierMap)
	case "Toe":
		GongUnmarshallSliceOfPointers(&instance.Toe, valueExpr, identifierMap)
	case "Fingernails":
		GongUnmarshallSliceOfPointers(&instance.Fingernails, valueExpr, identifierMap)
	case "Hole":
		GongUnmarshallSliceOfPointers(&instance.Hole, valueExpr, identifierMap)
	case "Arrow":
		GongUnmarshallSliceOfPointers(&instance.Arrow, valueExpr, identifierMap)
	case "Handbell":
		GongUnmarshallSliceOfPointers(&instance.Handbell, valueExpr, identifierMap)
	case "Brass_bend":
		GongUnmarshallSliceOfPointers(&instance.Brass_bend, valueExpr, identifierMap)
	case "Flip":
		GongUnmarshallSliceOfPointers(&instance.Flip, valueExpr, identifierMap)
	case "Smear":
		GongUnmarshallSliceOfPointers(&instance.Smear, valueExpr, identifierMap)
	case "Open":
		GongUnmarshallSliceOfPointers(&instance.Open, valueExpr, identifierMap)
	case "Half_muted":
		GongUnmarshallSliceOfPointers(&instance.Half_muted, valueExpr, identifierMap)
	case "Harmon_mute":
		GongUnmarshallSliceOfPointers(&instance.Harmon_mute, valueExpr, identifierMap)
	case "Golpe":
		GongUnmarshallSliceOfPointers(&instance.Golpe, valueExpr, identifierMap)
	case "Other_technical":
		GongUnmarshallSliceOfPointers(&instance.Other_technical, valueExpr, identifierMap)
	}
	return nil
}

type Text_element_dataUnmarshaller struct{}

func (u *Text_element_dataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Text_element_data)
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

func (u *Text_element_dataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Text_element_data)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Dir":
		instance.Dir = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TieUnmarshaller struct{}

func (u *TieUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tie)
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

func (u *TieUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tie)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	}
	return nil
}

type TiedUnmarshaller struct{}

func (u *TiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tied)
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

func (u *TiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Orientation":
		instance.Orientation = GongExtractString(valueExpr)
	case "Bezier_x":
		instance.Bezier_x = GongExtractString(valueExpr)
	case "Bezier_y":
		instance.Bezier_y = GongExtractString(valueExpr)
	case "Bezier_x2":
		instance.Bezier_x2 = GongExtractString(valueExpr)
	case "Bezier_y2":
		instance.Bezier_y2 = GongExtractString(valueExpr)
	case "Bezier_offset":
		instance.Bezier_offset = GongExtractString(valueExpr)
	case "Bezier_offset2":
		instance.Bezier_offset2 = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type TimeUnmarshaller struct{}

func (u *TimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Time)
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

func (u *TimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Time)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Separator":
		GongUnmarshallEnum(&instance.Separator, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		instance.Halign = GongExtractString(valueExpr)
	case "Valign":
		instance.Valign = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Beat_type":
		instance.Beat_type = GongExtractString(valueExpr)
	case "Interchangeable":
		GongUnmarshallPointer(&instance.Interchangeable, valueExpr, identifierMap)
	case "Senza_misura":
		instance.Senza_misura = GongExtractString(valueExpr)
	}
	return nil
}

type Time_modificationUnmarshaller struct{}

func (u *Time_modificationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Time_modification)
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

func (u *Time_modificationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Time_modification)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Actual_notes":
		instance.Actual_notes = GongExtractInt(valueExpr)
	case "Normal_notes":
		instance.Normal_notes = GongExtractInt(valueExpr)
	case "Normal_type":
		GongUnmarshallEnum(&instance.Normal_type, valueExpr)
	case "Normal_dot":
		instance.Normal_dot = GongExtractString(valueExpr)
	}
	return nil
}

type TimpaniUnmarshaller struct{}

func (u *TimpaniUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Timpani)
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

func (u *TimpaniUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Timpani)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	}
	return nil
}

type TransposeUnmarshaller struct{}

func (u *TransposeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Transpose)
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

func (u *TransposeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Transpose)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Diatonic":
		instance.Diatonic = GongExtractInt(valueExpr)
	case "Chromatic":
		instance.Chromatic = GongExtractString(valueExpr)
	case "Octave_change":
		instance.Octave_change = GongExtractInt(valueExpr)
	case "Double":
		instance.Double = GongExtractFloat(valueExpr)
	}
	return nil
}

type TremoloUnmarshaller struct{}

func (u *TremoloUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tremolo)
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

func (u *TremoloUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tremolo)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type TupletUnmarshaller struct{}

func (u *TupletUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tuplet)
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

func (u *TupletUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Show_number":
		instance.Show_number = GongExtractString(valueExpr)
	case "Show_type":
		GongUnmarshallEnum(&instance.Show_type, valueExpr)
	case "Line_shape":
		instance.Line_shape = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Tuplet_actual":
		GongUnmarshallPointer(&instance.Tuplet_actual, valueExpr, identifierMap)
	case "Tuplet_normal":
		GongUnmarshallPointer(&instance.Tuplet_normal, valueExpr, identifierMap)
	}
	return nil
}

type Tuplet_dotUnmarshaller struct{}

func (u *Tuplet_dotUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tuplet_dot)
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

func (u *Tuplet_dotUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_dot)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type Tuplet_numberUnmarshaller struct{}

func (u *Tuplet_numberUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tuplet_number)
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

func (u *Tuplet_numberUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_number)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Tuplet_portionUnmarshaller struct{}

func (u *Tuplet_portionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tuplet_portion)
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

func (u *Tuplet_portionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_portion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tuplet_number":
		GongUnmarshallPointer(&instance.Tuplet_number, valueExpr, identifierMap)
	case "Tuplet_type":
		GongUnmarshallPointer(&instance.Tuplet_type, valueExpr, identifierMap)
	case "Tuplet_dot":
		GongUnmarshallSliceOfPointers(&instance.Tuplet_dot, valueExpr, identifierMap)
	}
	return nil
}

type Tuplet_typeUnmarshaller struct{}

func (u *Tuplet_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tuplet_type)
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

func (u *Tuplet_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		instance.Font_style = GongExtractString(valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		instance.Font_weight = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Typed_textUnmarshaller struct{}

func (u *Typed_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Typed_text)
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

func (u *Typed_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Typed_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type UnpitchedUnmarshaller struct{}

func (u *UnpitchedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Unpitched)
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

func (u *UnpitchedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Unpitched)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Display_step":
		GongUnmarshallEnum(&instance.Display_step, valueExpr)
	case "Display_octave":
		instance.Display_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type Virtual_instrumentUnmarshaller struct{}

func (u *Virtual_instrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Virtual_instrument)
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

func (u *Virtual_instrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Virtual_instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Virtual_library":
		instance.Virtual_library = GongExtractString(valueExpr)
	case "Virtual_name":
		instance.Virtual_name = GongExtractString(valueExpr)
	}
	return nil
}

type WaitUnmarshaller struct{}

func (u *WaitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Wait)
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

func (u *WaitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wait)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		GongUnmarshallEnum(&instance.Time_only, valueExpr)
	}
	return nil
}

type Wavy_lineUnmarshaller struct{}

func (u *Wavy_lineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Wavy_line)
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

func (u *Wavy_lineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wavy_line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		instance.Placement = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Start_note":
		instance.Start_note = GongExtractString(valueExpr)
	case "Trill_step":
		instance.Trill_step = GongExtractString(valueExpr)
	case "Two_note_turn":
		instance.Two_note_turn = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	}
	return nil
}

type WedgeUnmarshaller struct{}

func (u *WedgeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Wedge)
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

func (u *WedgeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wedge)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Spread":
		instance.Spread = GongExtractString(valueExpr)
	case "Niente":
		GongUnmarshallEnum(&instance.Niente, valueExpr)
	case "Line_type":
		instance.Line_type = GongExtractString(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type WoodUnmarshaller struct{}

func (u *WoodUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Wood)
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

func (u *WoodUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wood)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type WorkUnmarshaller struct{}

func (u *WorkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Work)
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

func (u *WorkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Work)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Work_number":
		instance.Work_number = GongExtractString(valueExpr)
	case "Work_title":
		instance.Work_title = GongExtractString(valueExpr)
	case "Opus":
		GongUnmarshallPointer(&instance.Opus, valueExpr, identifierMap)
	}
	return nil
}

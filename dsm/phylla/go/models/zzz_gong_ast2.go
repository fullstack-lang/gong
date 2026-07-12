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
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	_time__dummyDeclaration2 time.Duration
	_                        = _time__dummyDeclaration2
)

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

	var fileModuleVersion string
	for _, commentGroup := range inFile.Comments {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "// go module version: ") {
				fileModuleVersion = strings.TrimPrefix(comment.Text, "// go module version: ")
			}
		}
	}

	var parserModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		parserModuleVersion = buildInfo.Main.Version
	}

	if fileModuleVersion != "" && parserModuleVersion != "" && fileModuleVersion != parserModuleVersion {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(time.Now(), fmt.Sprintf("Warning: file module version '%s' does not match parser module version '%s'", fileModuleVersion, parserModuleVersion))
		}
	}

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
				for i := 1; i < len(call.Args); i++ {
					if ident, ok := call.Args[i].(*ast.Ident); ok {
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
type AxesShapeUnmarshaller struct{}

func (u *AxesShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AxesShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AxesShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AxesShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LengthX":
		instance.LengthX = GongExtractFloat(valueExpr)
	case "LengthY":
		instance.LengthY = GongExtractFloat(valueExpr)
	case "IsWithHiddenHandle":
		instance.IsWithHiddenHandle = GongExtractBool(valueExpr)
	}
	return nil
}

type BaseVectorShapeUnmarshaller struct{}

func (u *BaseVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(BaseVectorShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *BaseVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BaseVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type BaseVectorShapeGridUnmarshaller struct{}

func (u *BaseVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(BaseVectorShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *BaseVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BaseVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BaseVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.BaseVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

type CircleGridShapeUnmarshaller struct{}

func (u *CircleGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(CircleGridShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CircleGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EndArcShapeUnmarshaller struct{}

func (u *EndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(EndArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type EndArcShapeGridUnmarshaller struct{}

func (u *EndArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(EndArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EndArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.EndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type EndArcShapeV2Unmarshaller struct{}

func (u *EndArcShapeV2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(EndArcShapeV2)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EndArcShapeV2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndArcShapeV2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type EndArcShapeV2GridUnmarshaller struct{}

func (u *EndArcShapeV2GridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(EndArcShapeV2Grid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EndArcShapeV2GridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndArcShapeV2Grid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EndArcShapesV2":
		GongUnmarshallSliceOfPointers(&instance.EndArcShapesV2, valueExpr, identifierMap)
	}
	return nil
}

type ExplanationTextShapeUnmarshaller struct{}

func (u *ExplanationTextShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ExplanationTextShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ExplanationTextShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExplanationTextShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GridPathShapeUnmarshaller struct{}

func (u *GridPathShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GridPathShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GridPathShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GridPathShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GrowthCurveBezierShapeUnmarshaller struct{}

func (u *GrowthCurveBezierShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GrowthCurveBezierShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GrowthCurveBezierShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurveBezierShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "ControlPointStartX":
		instance.ControlPointStartX = GongExtractFloat(valueExpr)
	case "ControlPointStartY":
		instance.ControlPointStartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "ControlPointEndX":
		instance.ControlPointEndX = GongExtractFloat(valueExpr)
	case "ControlPointEndY":
		instance.ControlPointEndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type GrowthCurveBezierShapeGridUnmarshaller struct{}

func (u *GrowthCurveBezierShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GrowthCurveBezierShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GrowthCurveBezierShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurveBezierShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GrowthCurveBezierShapes":
		GongUnmarshallSliceOfPointers(&instance.GrowthCurveBezierShapes, valueExpr, identifierMap)
	}
	return nil
}

type GrowthCurveRhombusGridShapeUnmarshaller struct{}

func (u *GrowthCurveRhombusGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GrowthCurveRhombusGridShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GrowthCurveRhombusGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurveRhombusGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GrowthCurveRhombusShapes":
		GongUnmarshallSliceOfPointers(&instance.GrowthCurveRhombusShapes, valueExpr, identifierMap)
	}
	return nil
}

type GrowthCurveRhombusShapeUnmarshaller struct{}

func (u *GrowthCurveRhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GrowthCurveRhombusShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GrowthCurveRhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurveRhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type GrowthVectorShapeUnmarshaller struct{}

func (u *GrowthVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GrowthVectorShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GrowthVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type InitialRhombusGridShapeUnmarshaller struct{}

func (u *InitialRhombusGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(InitialRhombusGridShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *InitialRhombusGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialRhombusGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "InitialRhombusShapes":
		GongUnmarshallSliceOfPointers(&instance.InitialRhombusShapes, valueExpr, identifierMap)
	}
	return nil
}

type InitialRhombusShapeUnmarshaller struct{}

func (u *InitialRhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(InitialRhombusShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *InitialRhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialRhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type LibraryUnmarshaller struct{}

func (u *LibraryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Library)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *LibraryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Library)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "Plants":
		GongUnmarshallSliceOfPointers(&instance.Plants, valueExpr, identifierMap)
	}
	return nil
}

type NextCircleShapeUnmarshaller struct{}

func (u *NextCircleShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(NextCircleShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *NextCircleShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NextCircleShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PerpendicularVectorUnmarshaller struct{}

func (u *PerpendicularVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PerpendicularVector)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PerpendicularVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVector)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PerpendicularVectorGridUnmarshaller struct{}

func (u *PerpendicularVectorGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PerpendicularVectorGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PerpendicularVectorGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVectorGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PerpendicularVectors":
		GongUnmarshallSliceOfPointers(&instance.PerpendicularVectors, valueExpr, identifierMap)
	}
	return nil
}

type PerpendicularVectorGridHalfwayUnmarshaller struct{}

func (u *PerpendicularVectorGridHalfwayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PerpendicularVectorGridHalfway)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PerpendicularVectorGridHalfwayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVectorGridHalfway)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PerpendicularVectorHalfways":
		GongUnmarshallSliceOfPointers(&instance.PerpendicularVectorHalfways, valueExpr, identifierMap)
	}
	return nil
}

type PerpendicularVectorHalfwayUnmarshaller struct{}

func (u *PerpendicularVectorHalfwayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PerpendicularVectorHalfway)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PerpendicularVectorHalfwayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVectorHalfway)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PlantUnmarshaller struct{}

func (u *PlantUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Plant)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PlantUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Plant)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "N":
		instance.N = GongExtractInt(valueExpr)
	case "M":
		instance.M = GongExtractInt(valueExpr)
	case "StackHeight":
		instance.StackHeight = GongExtractInt(valueExpr)
	case "RhombusInsideAngle":
		instance.RhombusInsideAngle = GongExtractFloat(valueExpr)
	case "Thickness":
		instance.Thickness = GongExtractFloat(valueExpr)
	case "RhombusSideLength":
		instance.RhombusSideLength = GongExtractFloat(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "IsPlantDiagramsNodeExpanded":
		instance.IsPlantDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "PlantDiagrams":
		GongUnmarshallSliceOfPointers(&instance.PlantDiagrams, valueExpr, identifierMap)
	case "AxesShape":
		GongUnmarshallPointer(&instance.AxesShape, valueExpr, identifierMap)
	case "ReferenceRhombus":
		GongUnmarshallPointer(&instance.ReferenceRhombus, valueExpr, identifierMap)
	case "PlantCircumferenceShape":
		GongUnmarshallPointer(&instance.PlantCircumferenceShape, valueExpr, identifierMap)
	case "GridPathShape":
		GongUnmarshallPointer(&instance.GridPathShape, valueExpr, identifierMap)
	case "InitialRhombusGridShape":
		GongUnmarshallPointer(&instance.InitialRhombusGridShape, valueExpr, identifierMap)
	case "ExplanationTextShape":
		GongUnmarshallPointer(&instance.ExplanationTextShape, valueExpr, identifierMap)
	case "RotatedReferenceRhombus":
		GongUnmarshallPointer(&instance.RotatedReferenceRhombus, valueExpr, identifierMap)
	case "RotatedPlantCircumferenceShape":
		GongUnmarshallPointer(&instance.RotatedPlantCircumferenceShape, valueExpr, identifierMap)
	case "RotatedGridPathShape":
		GongUnmarshallPointer(&instance.RotatedGridPathShape, valueExpr, identifierMap)
	case "RotatedRhombusGridShape2":
		GongUnmarshallPointer(&instance.RotatedRhombusGridShape2, valueExpr, identifierMap)
	case "GrowthCurveRhombusGridShape":
		GongUnmarshallPointer(&instance.GrowthCurveRhombusGridShape, valueExpr, identifierMap)
	case "GrowthVectorShape":
		GongUnmarshallPointer(&instance.GrowthVectorShape, valueExpr, identifierMap)
	case "PerpendicularVectorGrid":
		GongUnmarshallPointer(&instance.PerpendicularVectorGrid, valueExpr, identifierMap)
	case "PerpendicularVectorGridHalfway":
		GongUnmarshallPointer(&instance.PerpendicularVectorGridHalfway, valueExpr, identifierMap)
	case "BaseVectorShapeGrid":
		GongUnmarshallPointer(&instance.BaseVectorShapeGrid, valueExpr, identifierMap)
	case "StartArcShapeGrid":
		GongUnmarshallPointer(&instance.StartArcShapeGrid, valueExpr, identifierMap)
	case "StartArcShapeV2Grid":
		GongUnmarshallPointer(&instance.StartArcShapeV2Grid, valueExpr, identifierMap)
	case "EndArcShapeGrid":
		GongUnmarshallPointer(&instance.EndArcShapeGrid, valueExpr, identifierMap)
	case "EndArcShapeV2Grid":
		GongUnmarshallPointer(&instance.EndArcShapeV2Grid, valueExpr, identifierMap)
	case "GrowthCurveBezierShapeGrid":
		GongUnmarshallPointer(&instance.GrowthCurveBezierShapeGrid, valueExpr, identifierMap)
	case "StackOfGrowthCurve":
		GongUnmarshallPointer(&instance.StackOfGrowthCurve, valueExpr, identifierMap)
	}
	return nil
}

type PlantCircumferenceShapeUnmarshaller struct{}

func (u *PlantCircumferenceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PlantCircumferenceShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PlantCircumferenceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlantCircumferenceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "AngleDegree":
		instance.AngleDegree = GongExtractFloat(valueExpr)
	case "Length":
		instance.Length = GongExtractFloat(valueExpr)
	}
	return nil
}

type PlantDiagramUnmarshaller struct{}

func (u *PlantDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PlantDiagram)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PlantDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlantDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "OriginX":
		instance.OriginX = GongExtractFloat(valueExpr)
	case "OriginY":
		instance.OriginY = GongExtractFloat(valueExpr)
	case "IsHiddenAxesShape":
		instance.IsHiddenAxesShape = GongExtractBool(valueExpr)
	case "IsHiddenReferenceRhombus":
		instance.IsHiddenReferenceRhombus = GongExtractBool(valueExpr)
	case "IsHiddenPlantCircumferenceShape":
		instance.IsHiddenPlantCircumferenceShape = GongExtractBool(valueExpr)
	case "IsHiddenGridPathShape":
		instance.IsHiddenGridPathShape = GongExtractBool(valueExpr)
	case "IsHiddenRhombusGridShape":
		instance.IsHiddenRhombusGridShape = GongExtractBool(valueExpr)
	case "IsHiddenExplanationTextShape":
		instance.IsHiddenExplanationTextShape = GongExtractBool(valueExpr)
	case "IsHiddenRotatedReferenceRhombus":
		instance.IsHiddenRotatedReferenceRhombus = GongExtractBool(valueExpr)
	case "IsHiddenRotatedPlantCircumferenceShape":
		instance.IsHiddenRotatedPlantCircumferenceShape = GongExtractBool(valueExpr)
	case "IsHiddenRotatedGridPathShape":
		instance.IsHiddenRotatedGridPathShape = GongExtractBool(valueExpr)
	case "IsHiddenRotatedRhombusGridShape":
		instance.IsHiddenRotatedRhombusGridShape = GongExtractBool(valueExpr)
	case "IsHiddenGrowthPathRhombusGridShape":
		instance.IsHiddenGrowthPathRhombusGridShape = GongExtractBool(valueExpr)
	case "IsHiddenGrowthVectorShape":
		instance.IsHiddenGrowthVectorShape = GongExtractBool(valueExpr)
	case "IsHiddenPerpendicularVectorGrid":
		instance.IsHiddenPerpendicularVectorGrid = GongExtractBool(valueExpr)
	case "IsHiddenPerpendicularVectorGridHalfway":
		instance.IsHiddenPerpendicularVectorGridHalfway = GongExtractBool(valueExpr)
	case "IsHiddenBaseVectorShapeGrid":
		instance.IsHiddenBaseVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStartArcShapeGrid":
		instance.IsHiddenStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStartArcShapeV2Grid":
		instance.IsHiddenStartArcShapeV2Grid = GongExtractBool(valueExpr)
	case "IsHiddenEndArcShapeGrid":
		instance.IsHiddenEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenEndArcShapeV2Grid":
		instance.IsHiddenEndArcShapeV2Grid = GongExtractBool(valueExpr)
	case "IsHiddenGrowthCurveBezierShapeGrid":
		instance.IsHiddenGrowthCurveBezierShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStackOfGrowthCurve":
		instance.IsHiddenStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Rendered3DShape":
		GongUnmarshallPointer(&instance.Rendered3DShape, valueExpr, identifierMap)
	}
	return nil
}

type Rendered3DShapeUnmarshaller struct{}

func (u *Rendered3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Rendered3DShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Rendered3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rendered3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ViewX":
		instance.ViewX = GongExtractFloat(valueExpr)
	case "ViewY":
		instance.ViewY = GongExtractFloat(valueExpr)
	case "ViewZ":
		instance.ViewZ = GongExtractFloat(valueExpr)
	case "TargetX":
		instance.TargetX = GongExtractFloat(valueExpr)
	case "TargetY":
		instance.TargetY = GongExtractFloat(valueExpr)
	case "TargetZ":
		instance.TargetZ = GongExtractFloat(valueExpr)
	case "Fov":
		instance.Fov = GongExtractFloat(valueExpr)
	}
	return nil
}

type RhombusShapeUnmarshaller struct{}

func (u *RhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RhombusShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type RotatedRhombusGridShapeUnmarshaller struct{}

func (u *RotatedRhombusGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RotatedRhombusGridShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RotatedRhombusGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RotatedRhombusGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RotatedRhombusShapes":
		GongUnmarshallSliceOfPointers(&instance.RotatedRhombusShapes, valueExpr, identifierMap)
	}
	return nil
}

type RotatedRhombusShapeUnmarshaller struct{}

func (u *RotatedRhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RotatedRhombusShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RotatedRhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RotatedRhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type StackGrowthCurveBezierShapeUnmarshaller struct{}

func (u *StackGrowthCurveBezierShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StackGrowthCurveBezierShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StackGrowthCurveBezierShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurveBezierShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "ControlPointStartX":
		instance.ControlPointStartX = GongExtractFloat(valueExpr)
	case "ControlPointStartY":
		instance.ControlPointStartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "ControlPointEndX":
		instance.ControlPointEndX = GongExtractFloat(valueExpr)
	case "ControlPointEndY":
		instance.ControlPointEndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type StackOfGrowthCurveUnmarshaller struct{}

func (u *StackOfGrowthCurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StackOfGrowthCurve)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StackOfGrowthCurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfGrowthCurve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackGrowthCurveBezierShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurveBezierShapes, valueExpr, identifierMap)
	}
	return nil
}

type StartArcShapeUnmarshaller struct{}

func (u *StartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StartArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type StartArcShapeGridUnmarshaller struct{}

func (u *StartArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StartArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StartArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StartArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type StartArcShapeV2Unmarshaller struct{}

func (u *StartArcShapeV2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StartArcShapeV2)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StartArcShapeV2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartArcShapeV2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type StartArcShapeV2GridUnmarshaller struct{}

func (u *StartArcShapeV2GridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StartArcShapeV2Grid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StartArcShapeV2GridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartArcShapeV2Grid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartArcShapesV2":
		GongUnmarshallSliceOfPointers(&instance.StartArcShapesV2, valueExpr, identifierMap)
	}
	return nil
}

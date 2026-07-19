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
type ArcNormalVectorShapeUnmarshaller struct{}

func (u *ArcNormalVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ArcNormalVectorShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ArcNormalVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcNormalVectorShape)
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

type ArcNormalVectorShapeGridUnmarshaller struct{}

func (u *ArcNormalVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ArcNormalVectorShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ArcNormalVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcNormalVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ArcNormalVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.ArcNormalVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

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

type EndHalfwayArcShapeUnmarshaller struct{}

func (u *EndHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(EndHalfwayArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EndHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndHalfwayArcShape)
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
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type EndHalfwayArcShapeGridUnmarshaller struct{}

func (u *EndHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(EndHalfwayArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EndHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EndHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.EndHalfwayArcShapes, valueExpr, identifierMap)
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

type GrowthCurve2DUnmarshaller struct{}

func (u *GrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(GrowthCurve2D)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.StartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "EndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.EndHalfwayArcShapeGrid, valueExpr, identifierMap)
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

type MidArcVectorShapeUnmarshaller struct{}

func (u *MidArcVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MidArcVectorShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MidArcVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MidArcVectorShape)
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

type MidArcVectorShapeGridUnmarshaller struct{}

func (u *MidArcVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MidArcVectorShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MidArcVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MidArcVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MidArcVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.MidArcVectorShapes, valueExpr, identifierMap)
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
	case "RelativeVerticalThickness":
		instance.RelativeVerticalThickness = GongExtractFloat(valueExpr)
	case "RadialThickness":
		instance.RadialThickness = GongExtractFloat(valueExpr)
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
	case "RhombusStuff":
		GongUnmarshallPointer(&instance.RhombusStuff, valueExpr, identifierMap)
	case "GrowthVectorShape":
		GongUnmarshallPointer(&instance.GrowthVectorShape, valueExpr, identifierMap)
	case "PerpendicularVectorGrid":
		GongUnmarshallPointer(&instance.PerpendicularVectorGrid, valueExpr, identifierMap)
	case "PerpendicularVectorGridHalfway":
		GongUnmarshallPointer(&instance.PerpendicularVectorGridHalfway, valueExpr, identifierMap)
	case "BaseVectorShapeGrid":
		GongUnmarshallPointer(&instance.BaseVectorShapeGrid, valueExpr, identifierMap)
	case "ArcNormalVectorShapeGrid":
		GongUnmarshallPointer(&instance.ArcNormalVectorShapeGrid, valueExpr, identifierMap)
	case "StartArcShapeGrid":
		GongUnmarshallPointer(&instance.StartArcShapeGrid, valueExpr, identifierMap)
	case "TopStartArcShapeGrid":
		GongUnmarshallPointer(&instance.TopStartArcShapeGrid, valueExpr, identifierMap)
	case "EndArcShapeGrid":
		GongUnmarshallPointer(&instance.EndArcShapeGrid, valueExpr, identifierMap)
	case "TopEndArcShapeGrid":
		GongUnmarshallPointer(&instance.TopEndArcShapeGrid, valueExpr, identifierMap)
	case "ShiftedBottomTopStartArcShapeGrid":
		GongUnmarshallPointer(&instance.ShiftedBottomTopStartArcShapeGrid, valueExpr, identifierMap)
	case "MidArcVectorShapeGrid":
		GongUnmarshallPointer(&instance.MidArcVectorShapeGrid, valueExpr, identifierMap)
	case "TopMidArcVectorShapeGrid":
		GongUnmarshallPointer(&instance.TopMidArcVectorShapeGrid, valueExpr, identifierMap)
	case "StartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.StartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "TopStartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopStartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "EndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.EndHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "TopEndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopEndHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "StackOfGrowthCurve":
		GongUnmarshallPointer(&instance.StackOfGrowthCurve, valueExpr, identifierMap)
	case "TopStackOfGrowthCurve":
		GongUnmarshallPointer(&instance.TopStackOfGrowthCurve, valueExpr, identifierMap)
	case "ShiftedLeftStackOfGrowthCurve":
		GongUnmarshallPointer(&instance.ShiftedLeftStackOfGrowthCurve, valueExpr, identifierMap)
	case "ShiftedLeftStackOfNormalVector":
		GongUnmarshallPointer(&instance.ShiftedLeftStackOfNormalVector, valueExpr, identifierMap)
	case "GrowthCurve2D":
		GongUnmarshallPointer(&instance.GrowthCurve2D, valueExpr, identifierMap)
	case "TopGrowthCurve2D":
		GongUnmarshallPointer(&instance.TopGrowthCurve2D, valueExpr, identifierMap)
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
	case "IsRhombusNodesExpanded":
		instance.IsRhombusNodesExpanded = GongExtractBool(valueExpr)
	case "IsArcNodesExpanded":
		instance.IsArcNodesExpanded = GongExtractBool(valueExpr)
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
	case "IsHiddenArcNormalVectorShapeGrid":
		instance.IsHiddenArcNormalVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStartArcShapeGrid":
		instance.IsHiddenStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopStartArcShapeGrid":
		instance.IsHiddenTopStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenShiftedBottomTopStartArcShapeGrid":
		instance.IsHiddenShiftedBottomTopStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenMidArcVectorShapeGrid":
		instance.IsHiddenMidArcVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopMidArcVectorShapeGrid":
		instance.IsHiddenTopMidArcVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStartHalfwayArcShapeGrid":
		instance.IsHiddenStartHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopStartHalfwayArcShapeGrid":
		instance.IsHiddenTopStartHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenEndHalfwayArcShapeGrid":
		instance.IsHiddenEndHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopEndHalfwayArcShapeGrid":
		instance.IsHiddenTopEndHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenEndArcShapeGrid":
		instance.IsHiddenEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopEndArcShapeGrid":
		instance.IsHiddenTopEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenBottomStartArcShapeGrid":
		instance.IsHiddenBottomStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenBottomEndArcShapeGrid":
		instance.IsHiddenBottomEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStackOfGrowthCurve":
		instance.IsHiddenStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenTopStackOfGrowthCurve":
		instance.IsHiddenTopStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenBottomStackOfGrowthCurve":
		instance.IsHiddenBottomStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenShiftedLeftStackOfGrowthCurve":
		instance.IsHiddenShiftedLeftStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenShiftedLeftStackOfNormalVector":
		instance.IsHiddenShiftedLeftStackOfNormalVector = GongExtractBool(valueExpr)
	case "IsHiddenGrowthCurve2D":
		instance.IsHiddenGrowthCurve2D = GongExtractBool(valueExpr)
	case "IsHiddenTopGrowthCurve2D":
		instance.IsHiddenTopGrowthCurve2D = GongExtractBool(valueExpr)
	case "IsHiddenTorusStackShape":
		instance.IsHiddenTorusStackShape = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Rendered3DShape":
		GongUnmarshallPointer(&instance.Rendered3DShape, valueExpr, identifierMap)
	case "TorusStackShape":
		GongUnmarshallPointer(&instance.TorusStackShape, valueExpr, identifierMap)
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

type RhombusStuffUnmarshaller struct{}

func (u *RhombusStuffUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RhombusStuff)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RhombusStuffUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RhombusStuff)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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

type ShiftedBottomTopStartArcShapeUnmarshaller struct{}

func (u *ShiftedBottomTopStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedBottomTopStartArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedBottomTopStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedBottomTopStartArcShape)
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

type ShiftedBottomTopStartArcShapeGridUnmarshaller struct{}

func (u *ShiftedBottomTopStartArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedBottomTopStartArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedBottomTopStartArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedBottomTopStartArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedBottomTopStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedBottomTopStartArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller struct{}

func (u *ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedLeftStackGrowthCurveEndArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackGrowthCurveEndArcShape)
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

type ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller struct{}

func (u *ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedLeftStackGrowthCurveStartArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackGrowthCurveStartArcShape)
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

type ShiftedLeftStackNormalVectorUnmarshaller struct{}

func (u *ShiftedLeftStackNormalVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedLeftStackNormalVector)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedLeftStackNormalVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackNormalVector)
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

type ShiftedLeftStackOfGrowthCurveUnmarshaller struct{}

func (u *ShiftedLeftStackOfGrowthCurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedLeftStackOfGrowthCurve)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedLeftStackOfGrowthCurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackOfGrowthCurve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedLeftStackGrowthCurveStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftStackGrowthCurveStartArcShapes, valueExpr, identifierMap)
	case "ShiftedLeftStackGrowthCurveEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftStackGrowthCurveEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedLeftStackOfNormalVectorUnmarshaller struct{}

func (u *ShiftedLeftStackOfNormalVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ShiftedLeftStackOfNormalVector)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShiftedLeftStackOfNormalVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackOfNormalVector)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedLeftStackNormalVectors":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftStackNormalVectors, valueExpr, identifierMap)
	}
	return nil
}

type StackGrowthCurveEndArcShapeUnmarshaller struct{}

func (u *StackGrowthCurveEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StackGrowthCurveEndArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StackGrowthCurveEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurveEndArcShape)
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

type StackGrowthCurveStartArcShapeUnmarshaller struct{}

func (u *StackGrowthCurveStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StackGrowthCurveStartArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StackGrowthCurveStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurveStartArcShape)
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
	case "StackGrowthCurveStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurveStartArcShapes, valueExpr, identifierMap)
	case "StackGrowthCurveEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurveEndArcShapes, valueExpr, identifierMap)
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

type StartHalfwayArcShapeUnmarshaller struct{}

func (u *StartHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StartHalfwayArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StartHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartHalfwayArcShape)
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
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StartHalfwayArcShapeGridUnmarshaller struct{}

func (u *StartHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(StartHalfwayArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *StartHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StartHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopEndArcShapeUnmarshaller struct{}

func (u *TopEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopEndArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndArcShape)
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

type TopEndArcShapeGridUnmarshaller struct{}

func (u *TopEndArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopEndArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopEndArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopEndHalfwayArcShapeUnmarshaller struct{}

func (u *TopEndHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopEndHalfwayArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopEndHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndHalfwayArcShape)
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
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type TopEndHalfwayArcShapeGridUnmarshaller struct{}

func (u *TopEndHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopEndHalfwayArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopEndHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopEndHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopEndHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopGrowthCurve2DUnmarshaller struct{}

func (u *TopGrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopGrowthCurve2D)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopGrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopGrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopStartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "TopEndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopEndHalfwayArcShapeGrid, valueExpr, identifierMap)
	}
	return nil
}

type TopMidArcVectorShapeUnmarshaller struct{}

func (u *TopMidArcVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopMidArcVectorShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopMidArcVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopMidArcVectorShape)
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

type TopMidArcVectorShapeGridUnmarshaller struct{}

func (u *TopMidArcVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopMidArcVectorShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopMidArcVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopMidArcVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopMidArcVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.TopMidArcVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStackGrowthCurveEndArcShapeUnmarshaller struct{}

func (u *TopStackGrowthCurveEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStackGrowthCurveEndArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStackGrowthCurveEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackGrowthCurveEndArcShape)
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

type TopStackGrowthCurveStartArcShapeUnmarshaller struct{}

func (u *TopStackGrowthCurveStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStackGrowthCurveStartArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStackGrowthCurveStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackGrowthCurveStartArcShape)
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

type TopStackOfGrowthCurveUnmarshaller struct{}

func (u *TopStackOfGrowthCurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStackOfGrowthCurve)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStackOfGrowthCurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackOfGrowthCurve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStackGrowthCurveStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStackGrowthCurveStartArcShapes, valueExpr, identifierMap)
	case "TopStackGrowthCurveEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStackGrowthCurveEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStartArcShapeUnmarshaller struct{}

func (u *TopStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStartArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartArcShape)
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

type TopStartArcShapeGridUnmarshaller struct{}

func (u *TopStartArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStartArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStartArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStartArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStartHalfwayArcShapeUnmarshaller struct{}

func (u *TopStartHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStartHalfwayArcShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStartHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartHalfwayArcShape)
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
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type TopStartHalfwayArcShapeGridUnmarshaller struct{}

func (u *TopStartHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TopStartHalfwayArcShapeGrid)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TopStartHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStartHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStartHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TorusStackShapeUnmarshaller struct{}

func (u *TorusStackShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TorusStackShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TorusStackShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TorusStackShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

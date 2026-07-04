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
type AllUnmarshaller struct{}

func (u *AllUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(All)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AllUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*All)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type AnnotationUnmarshaller struct{}

func (u *AnnotationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Annotation)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AnnotationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Annotation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Documentations":
		GongUnmarshallSliceOfPointers(&instance.Documentations, valueExpr, identifierMap)
	}
	return nil
}

type AttributeUnmarshaller struct{}

func (u *AttributeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Attribute)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AttributeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Attribute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "Default":
		instance.Default = GongExtractString(valueExpr)
	case "Use":
		instance.Use = GongExtractString(valueExpr)
	case "Form":
		instance.Form = GongExtractString(valueExpr)
	case "Fixed":
		instance.Fixed = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "TargetNamespace":
		instance.TargetNamespace = GongExtractString(valueExpr)
	case "SimpleType":
		instance.SimpleType = GongExtractString(valueExpr)
	case "IDXSD":
		instance.IDXSD = GongExtractString(valueExpr)
	}
	return nil
}

type AttributeGroupUnmarshaller struct{}

func (u *AttributeGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AttributeGroup)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AttributeGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributeGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type ChoiceUnmarshaller struct{}

func (u *ChoiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Choice)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ChoiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Choice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type ComplexContentUnmarshaller struct{}

func (u *ComplexContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ComplexContent)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ComplexContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ComplexContent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ComplexTypeUnmarshaller struct{}

func (u *ComplexTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ComplexType)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ComplexTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ComplexType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "IsAnonymous":
		instance.IsAnonymous = GongExtractBool(valueExpr)
	case "OuterElement":
		GongUnmarshallPointer(&instance.OuterElement, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallPointer(&instance.Extension, valueExpr, identifierMap)
	case "SimpleContent":
		GongUnmarshallPointer(&instance.SimpleContent, valueExpr, identifierMap)
	case "ComplexContent":
		GongUnmarshallPointer(&instance.ComplexContent, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type DocumentationUnmarshaller struct{}

func (u *DocumentationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Documentation)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *DocumentationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Documentation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

type ElementUnmarshaller struct{}

func (u *ElementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Element)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ElementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Element)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Default":
		instance.Default = GongExtractString(valueExpr)
	case "Fixed":
		instance.Fixed = GongExtractString(valueExpr)
	case "Nillable":
		instance.Nillable = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Abstract":
		instance.Abstract = GongExtractString(valueExpr)
	case "Form":
		instance.Form = GongExtractString(valueExpr)
	case "Block":
		instance.Block = GongExtractString(valueExpr)
	case "Final":
		instance.Final = GongExtractString(valueExpr)
	case "SimpleType":
		GongUnmarshallPointer(&instance.SimpleType, valueExpr, identifierMap)
	case "ComplexType":
		GongUnmarshallPointer(&instance.ComplexType, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type EnumerationUnmarshaller struct{}

func (u *EnumerationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Enumeration)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *EnumerationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Enumeration)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type ExtensionUnmarshaller struct{}

func (u *ExtensionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Extension)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ExtensionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extension)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Base":
		instance.Base = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	}
	return nil
}

type GroupUnmarshaller struct{}

func (u *GroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Group)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *GroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "IsAnonymous":
		instance.IsAnonymous = GongExtractBool(valueExpr)
	case "OuterElement":
		GongUnmarshallPointer(&instance.OuterElement, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type LengthUnmarshaller struct{}

func (u *LengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Length)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *LengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Length)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MaxInclusiveUnmarshaller struct{}

func (u *MaxInclusiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MaxInclusive)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MaxInclusiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MaxInclusive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MaxLengthUnmarshaller struct{}

func (u *MaxLengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MaxLength)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MaxLengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MaxLength)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MinInclusiveUnmarshaller struct{}

func (u *MinInclusiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MinInclusive)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MinInclusiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinInclusive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MinLengthUnmarshaller struct{}

func (u *MinLengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MinLength)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MinLengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinLength)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PatternUnmarshaller struct{}

func (u *PatternUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Pattern)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PatternUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pattern)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type RestrictionUnmarshaller struct{}

func (u *RestrictionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Restriction)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *RestrictionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Restriction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Base":
		instance.Base = GongExtractString(valueExpr)
	case "Enumerations":
		GongUnmarshallSliceOfPointers(&instance.Enumerations, valueExpr, identifierMap)
	case "MinInclusive":
		GongUnmarshallPointer(&instance.MinInclusive, valueExpr, identifierMap)
	case "MaxInclusive":
		GongUnmarshallPointer(&instance.MaxInclusive, valueExpr, identifierMap)
	case "Pattern":
		GongUnmarshallPointer(&instance.Pattern, valueExpr, identifierMap)
	case "WhiteSpace":
		GongUnmarshallPointer(&instance.WhiteSpace, valueExpr, identifierMap)
	case "MinLength":
		GongUnmarshallPointer(&instance.MinLength, valueExpr, identifierMap)
	case "MaxLength":
		GongUnmarshallPointer(&instance.MaxLength, valueExpr, identifierMap)
	case "Length":
		GongUnmarshallPointer(&instance.Length, valueExpr, identifierMap)
	case "TotalDigit":
		GongUnmarshallPointer(&instance.TotalDigit, valueExpr, identifierMap)
	}
	return nil
}

type SchemaUnmarshaller struct{}

func (u *SchemaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Schema)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SchemaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Schema)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Xs":
		instance.Xs = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "SimpleTypes":
		GongUnmarshallSliceOfPointers(&instance.SimpleTypes, valueExpr, identifierMap)
	case "ComplexTypes":
		GongUnmarshallSliceOfPointers(&instance.ComplexTypes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type SequenceUnmarshaller struct{}

func (u *SequenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Sequence)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SequenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sequence)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type SimpleContentUnmarshaller struct{}

func (u *SimpleContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SimpleContent)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SimpleContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SimpleContent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallPointer(&instance.Extension, valueExpr, identifierMap)
	case "Restriction":
		GongUnmarshallPointer(&instance.Restriction, valueExpr, identifierMap)
	}
	return nil
}

type SimpleTypeUnmarshaller struct{}

func (u *SimpleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SimpleType)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SimpleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SimpleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Restriction":
		GongUnmarshallPointer(&instance.Restriction, valueExpr, identifierMap)
	case "Union":
		GongUnmarshallPointer(&instance.Union, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type TotalDigitUnmarshaller struct{}

func (u *TotalDigitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TotalDigit)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TotalDigitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TotalDigit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type UnionUnmarshaller struct{}

func (u *UnionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Union)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *UnionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Union)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "MemberTypes":
		instance.MemberTypes = GongExtractString(valueExpr)
	}
	return nil
}

type WhiteSpaceUnmarshaller struct{}

func (u *WhiteSpaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(WhiteSpace)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *WhiteSpaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WhiteSpace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

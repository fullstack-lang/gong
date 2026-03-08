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
type ArtefactTypeUnmarshaller struct{}

func (u *ArtefactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ArtefactType)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ArtefactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtefactType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ArtefactTypeShapeUnmarshaller struct{}

func (u *ArtefactTypeShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ArtefactTypeShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ArtefactTypeShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtefactTypeShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ArtefactType":
		GongUnmarshallPointer(&instance.ArtefactType, valueExpr, identifierMap)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	}
	return nil
}

type ArtistUnmarshaller struct{}

func (u *ArtistUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Artist)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ArtistUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Artist)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsDead":
		instance.IsDead = GongExtractBool(valueExpr)
	case "DateOfDeath":
		if call, ok := valueExpr.(*ast.CallExpr); ok {
			if len(call.Args) == 2 {
				if bl, ok := call.Args[1].(*ast.BasicLit); ok {
					instance.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
				}
			}
		}
	case "Place":
		GongUnmarshallPointer(&instance.Place, valueExpr, identifierMap)
	}
	return nil
}

type ArtistShapeUnmarshaller struct{}

func (u *ArtistShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ArtistShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ArtistShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtistShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Artist":
		GongUnmarshallPointer(&instance.Artist, valueExpr, identifierMap)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	}
	return nil
}

type ControlPointShapeUnmarshaller struct{}

func (u *ControlPointShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ControlPointShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ControlPointShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPointShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "IsStartShapeTheClosestShape":
		instance.IsStartShapeTheClosestShape = GongExtractBool(valueExpr)
	}
	return nil
}

type DeskUnmarshaller struct{}

func (u *DeskUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Desk)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *DeskUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Desk)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SelectedDiagram":
		GongUnmarshallPointer(&instance.SelectedDiagram, valueExpr, identifierMap)
	}
	return nil
}

type DiagramUnmarshaller struct{}

func (u *DiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Diagram)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *DiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Diagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MovementShapes":
		GongUnmarshallSliceOfPointers(&instance.MovementShapes, valueExpr, identifierMap)
	case "ArtefactTypeShapes":
		GongUnmarshallSliceOfPointers(&instance.ArtefactTypeShapes, valueExpr, identifierMap)
	case "ArtistShapes":
		GongUnmarshallSliceOfPointers(&instance.ArtistShapes, valueExpr, identifierMap)
	case "InfluenceShapes":
		GongUnmarshallSliceOfPointers(&instance.InfluenceShapes, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsNodeExpanded":
		instance.IsNodeExpanded = GongExtractBool(valueExpr)
	case "IsMovementCategoryNodeExpanded":
		instance.IsMovementCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsArtefactTypeCategoryNodeExpanded":
		instance.IsArtefactTypeCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsArtistCategoryNodeExpanded":
		instance.IsArtistCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryNodeExpanded":
		instance.IsInfluenceCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsMovementCategoryShown":
		instance.IsMovementCategoryShown = GongExtractBool(valueExpr)
	case "IsArtefactTypeCategoryShown":
		instance.IsArtefactTypeCategoryShown = GongExtractBool(valueExpr)
	case "IsArtistCategoryShown":
		instance.IsArtistCategoryShown = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryShown":
		instance.IsInfluenceCategoryShown = GongExtractBool(valueExpr)
	case "StartDate":
		if call, ok := valueExpr.(*ast.CallExpr); ok {
			if len(call.Args) == 2 {
				if bl, ok := call.Args[1].(*ast.BasicLit); ok {
					instance.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
				}
			}
		}
	case "EndDate":
		if call, ok := valueExpr.(*ast.CallExpr); ok {
			if len(call.Args) == 2 {
				if bl, ok := call.Args[1].(*ast.BasicLit); ok {
					instance.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
				}
			}
		}
	case "NbYearsForIntervals":
		instance.NbYearsForIntervals = GongExtractInt(valueExpr)
	case "XMargin":
		instance.XMargin = GongExtractFloat(valueExpr)
	case "YMargin":
		instance.YMargin = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "NextVerticalDateXMargin":
		instance.NextVerticalDateXMargin = GongExtractFloat(valueExpr)
	case "RedColorCode":
		instance.RedColorCode = GongExtractString(valueExpr)
	case "BackgroundGreyColorCode":
		instance.BackgroundGreyColorCode = GongExtractString(valueExpr)
	case "GrayColorCode":
		instance.GrayColorCode = GongExtractString(valueExpr)
	case "BottomBoxYOffset":
		instance.BottomBoxYOffset = GongExtractFloat(valueExpr)
	case "BottomBoxWidth":
		instance.BottomBoxWidth = GongExtractFloat(valueExpr)
	case "BottomBoxHeigth":
		instance.BottomBoxHeigth = GongExtractFloat(valueExpr)
	case "BottomBoxFontSize":
		instance.BottomBoxFontSize = GongExtractString(valueExpr)
	case "BottomBoxFontWeigth":
		instance.BottomBoxFontWeigth = GongExtractString(valueExpr)
	case "BottomBoxFontFamily":
		instance.BottomBoxFontFamily = GongExtractString(valueExpr)
	case "BottomBoxLetterSpacing":
		instance.BottomBoxLetterSpacing = GongExtractString(valueExpr)
	case "BottomBoxLetterColorCode":
		instance.BottomBoxLetterColorCode = GongExtractString(valueExpr)
	case "MovementRectAnchorType":
		GongUnmarshallEnum(&instance.MovementRectAnchorType, valueExpr)
	case "MovementTextAnchorType":
		GongUnmarshallEnum(&instance.MovementTextAnchorType, valueExpr)
	case "MovementDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementDominantBaselineType, valueExpr)
	case "MovementFontSize":
		instance.MovementFontSize = GongExtractString(valueExpr)
	case "MajorMovementFontSize":
		instance.MajorMovementFontSize = GongExtractString(valueExpr)
	case "MinorMovementFontSize":
		instance.MinorMovementFontSize = GongExtractString(valueExpr)
	case "MovementFontWeigth":
		instance.MovementFontWeigth = GongExtractString(valueExpr)
	case "MovementFontFamily":
		instance.MovementFontFamily = GongExtractString(valueExpr)
	case "MovementLetterSpacing":
		instance.MovementLetterSpacing = GongExtractString(valueExpr)
	case "AbstractMovementFontSize":
		instance.AbstractMovementFontSize = GongExtractString(valueExpr)
	case "AbstractMovementRectAnchorType":
		GongUnmarshallEnum(&instance.AbstractMovementRectAnchorType, valueExpr)
	case "AbstractMovementTextAnchorType":
		GongUnmarshallEnum(&instance.AbstractMovementTextAnchorType, valueExpr)
	case "AbstractDominantBaselineType":
		GongUnmarshallEnum(&instance.AbstractDominantBaselineType, valueExpr)
	case "MovementDateRectAnchorType":
		GongUnmarshallEnum(&instance.MovementDateRectAnchorType, valueExpr)
	case "MovementDateTextAnchorType":
		GongUnmarshallEnum(&instance.MovementDateTextAnchorType, valueExpr)
	case "MovementDateTextDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementDateTextDominantBaselineType, valueExpr)
	case "MovementDateAndPlacesFontSize":
		instance.MovementDateAndPlacesFontSize = GongExtractString(valueExpr)
	case "MovementDateAndPlacesFontWeigth":
		instance.MovementDateAndPlacesFontWeigth = GongExtractString(valueExpr)
	case "MovementDateAndPlacesFontFamily":
		instance.MovementDateAndPlacesFontFamily = GongExtractString(valueExpr)
	case "MovementDateAndPlacesLetterSpacing":
		instance.MovementDateAndPlacesLetterSpacing = GongExtractString(valueExpr)
	case "MovementBelowArcY_Offset":
		instance.MovementBelowArcY_Offset = GongExtractFloat(valueExpr)
	case "MovementBelowArcY_OffsetPerPlace":
		instance.MovementBelowArcY_OffsetPerPlace = GongExtractFloat(valueExpr)
	case "MovementPlacesRectAnchorType":
		GongUnmarshallEnum(&instance.MovementPlacesRectAnchorType, valueExpr)
	case "MovementPlacesTextAnchorType":
		GongUnmarshallEnum(&instance.MovementPlacesTextAnchorType, valueExpr)
	case "MovementPlacesDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementPlacesDominantBaselineType, valueExpr)
	case "ArtefactTypeFontSize":
		instance.ArtefactTypeFontSize = GongExtractString(valueExpr)
	case "ArtefactTypeFontWeigth":
		instance.ArtefactTypeFontWeigth = GongExtractString(valueExpr)
	case "ArtefactTypeFontFamily":
		instance.ArtefactTypeFontFamily = GongExtractString(valueExpr)
	case "ArtefactTypeLetterSpacing":
		instance.ArtefactTypeLetterSpacing = GongExtractString(valueExpr)
	case "ArtefactTypeRectAnchorType":
		GongUnmarshallEnum(&instance.ArtefactTypeRectAnchorType, valueExpr)
	case "ArtefactDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtefactDominantBaselineType, valueExpr)
	case "ArtefactTypeStrokeWidth":
		instance.ArtefactTypeStrokeWidth = GongExtractFloat(valueExpr)
	case "ArtistRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistRectAnchorType, valueExpr)
	case "ArtistTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistTextAnchorType, valueExpr)
	case "ArtistDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistDominantBaselineType, valueExpr)
	case "ArtistFontSize":
		instance.ArtistFontSize = GongExtractString(valueExpr)
	case "MajorArtistFontSize":
		instance.MajorArtistFontSize = GongExtractString(valueExpr)
	case "MinorArtistFontSize":
		instance.MinorArtistFontSize = GongExtractString(valueExpr)
	case "ArtistFontWeigth":
		instance.ArtistFontWeigth = GongExtractString(valueExpr)
	case "ArtistFontFamily":
		instance.ArtistFontFamily = GongExtractString(valueExpr)
	case "ArtistLetterSpacing":
		instance.ArtistLetterSpacing = GongExtractString(valueExpr)
	case "ArtistDateRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistDateRectAnchorType, valueExpr)
	case "ArtistDateTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistDateTextAnchorType, valueExpr)
	case "ArtistDateDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistDateDominantBaselineType, valueExpr)
	case "ArtistDateAndPlacesFontSize":
		instance.ArtistDateAndPlacesFontSize = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesFontWeigth":
		instance.ArtistDateAndPlacesFontWeigth = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesFontFamily":
		instance.ArtistDateAndPlacesFontFamily = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesLetterSpacing":
		instance.ArtistDateAndPlacesLetterSpacing = GongExtractString(valueExpr)
	case "ArtistPlacesRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistPlacesRectAnchorType, valueExpr)
	case "ArtistPlacesTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistPlacesTextAnchorType, valueExpr)
	case "ArtistPlacesDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistPlacesDominantBaselineType, valueExpr)
	case "InfluenceArrowSize":
		instance.InfluenceArrowSize = GongExtractFloat(valueExpr)
	case "InfluenceArrowStartOffset":
		instance.InfluenceArrowStartOffset = GongExtractFloat(valueExpr)
	case "InfluenceArrowEndOffset":
		instance.InfluenceArrowEndOffset = GongExtractFloat(valueExpr)
	case "InfluenceCornerRadius":
		instance.InfluenceCornerRadius = GongExtractFloat(valueExpr)
	case "InfluenceDashedLinePattern":
		instance.InfluenceDashedLinePattern = GongExtractString(valueExpr)
	}
	return nil
}

type InfluenceUnmarshaller struct{}

func (u *InfluenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Influence)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *InfluenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Influence)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SourceMovement":
		GongUnmarshallPointer(&instance.SourceMovement, valueExpr, identifierMap)
	case "SourceArtefactType":
		GongUnmarshallPointer(&instance.SourceArtefactType, valueExpr, identifierMap)
	case "SourceArtist":
		GongUnmarshallPointer(&instance.SourceArtist, valueExpr, identifierMap)
	case "TargetMovement":
		GongUnmarshallPointer(&instance.TargetMovement, valueExpr, identifierMap)
	case "TargetArtefactType":
		GongUnmarshallPointer(&instance.TargetArtefactType, valueExpr, identifierMap)
	case "TargetArtist":
		GongUnmarshallPointer(&instance.TargetArtist, valueExpr, identifierMap)
	case "IsHypothtical":
		instance.IsHypothtical = GongExtractBool(valueExpr)
	}
	return nil
}

type InfluenceShapeUnmarshaller struct{}

func (u *InfluenceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(InfluenceShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *InfluenceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InfluenceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Influence":
		GongUnmarshallPointer(&instance.Influence, valueExpr, identifierMap)
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type MovementUnmarshaller struct{}

func (u *MovementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Movement)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MovementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Movement)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Date":
		if call, ok := valueExpr.(*ast.CallExpr); ok {
			if len(call.Args) == 2 {
				if bl, ok := call.Args[1].(*ast.BasicLit); ok {
					instance.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
				}
			}
		}
	case "HideDate":
		instance.HideDate = GongExtractBool(valueExpr)
	case "Places":
		GongUnmarshallSliceOfPointers(&instance.Places, valueExpr, identifierMap)
	case "HasTaxonomicFilter":
		instance.HasTaxonomicFilter = GongExtractBool(valueExpr)
	case "TaxonomicFilter":
		instance.TaxonomicFilter = GongExtractString(valueExpr)
	case "IsFeatured":
		instance.IsFeatured = GongExtractBool(valueExpr)
	case "FeaturePrefix":
		instance.FeaturePrefix = GongExtractString(valueExpr)
	case "IsMajor":
		instance.IsMajor = GongExtractBool(valueExpr)
	case "IsMinor":
		instance.IsMinor = GongExtractBool(valueExpr)
	case "AdditionnalName":
		instance.AdditionnalName = GongExtractString(valueExpr)
	}
	return nil
}

type MovementShapeUnmarshaller struct{}

func (u *MovementShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MovementShape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MovementShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MovementShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Movement":
		GongUnmarshallPointer(&instance.Movement, valueExpr, identifierMap)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type PlaceUnmarshaller struct{}

func (u *PlaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Place)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PlaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Place)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

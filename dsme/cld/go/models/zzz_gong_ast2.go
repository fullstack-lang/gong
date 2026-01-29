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
type Category1Unmarshaller struct{}

func (u *Category1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Category1)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Category1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Category1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Category1ShapeUnmarshaller struct{}

func (u *Category1ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Category1Shape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Category1ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Category1Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Category1":
		GongUnmarshallPointer(&instance.Category1, valueExpr, identifierMap)
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

type Category2Unmarshaller struct{}

func (u *Category2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Category2)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Category2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Category2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Category2ShapeUnmarshaller struct{}

func (u *Category2ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Category2Shape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Category2ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Category2Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Category2":
		GongUnmarshallPointer(&instance.Category2, valueExpr, identifierMap)
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

type Category3Unmarshaller struct{}

func (u *Category3Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Category3)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Category3Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Category3)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Category3ShapeUnmarshaller struct{}

func (u *Category3ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Category3Shape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Category3ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Category3Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Category3":
		GongUnmarshallPointer(&instance.Category3, valueExpr, identifierMap)
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
	case "Category1Shapes":
		GongUnmarshallSliceOfPointers(&instance.Category1Shapes, valueExpr, identifierMap)
	case "Category2Shapes":
		GongUnmarshallSliceOfPointers(&instance.Category2Shapes, valueExpr, identifierMap)
	case "Category3Shapes":
		GongUnmarshallSliceOfPointers(&instance.Category3Shapes, valueExpr, identifierMap)
	case "InfluenceShapes":
		GongUnmarshallSliceOfPointers(&instance.InfluenceShapes, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsNodeExpanded":
		instance.IsNodeExpanded = GongExtractBool(valueExpr)
	case "IsCategory1NodeExpanded":
		instance.IsCategory1NodeExpanded = GongExtractBool(valueExpr)
	case "IsCategory2NodeExpanded":
		instance.IsCategory2NodeExpanded = GongExtractBool(valueExpr)
	case "IsCategory3NodeExpanded":
		instance.IsCategory3NodeExpanded = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryNodeExpanded":
		instance.IsInfluenceCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsCategory1Shown":
		instance.IsCategory1Shown = GongExtractBool(valueExpr)
	case "IsCategory2Shown":
		instance.IsCategory2Shown = GongExtractBool(valueExpr)
	case "IsCategory3Shown":
		instance.IsCategory3Shown = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryShown":
		instance.IsInfluenceCategoryShown = GongExtractBool(valueExpr)
	case "XMargin":
		instance.XMargin = GongExtractFloat(valueExpr)
	case "YMargin":
		instance.YMargin = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "RedColorCode":
		instance.RedColorCode = GongExtractString(valueExpr)
	case "BackgroundGreyColorCode":
		instance.BackgroundGreyColorCode = GongExtractString(valueExpr)
	case "GrayColorCode":
		instance.GrayColorCode = GongExtractString(valueExpr)
	case "Category1RectAnchorType":
		GongUnmarshallEnum(&instance.Category1RectAnchorType, valueExpr)
	case "Category1TextAnchorType":
		GongUnmarshallEnum(&instance.Category1TextAnchorType, valueExpr)
	case "Category1DominantBaselineType":
		GongUnmarshallEnum(&instance.Category1DominantBaselineType, valueExpr)
	case "Category1FontSize":
		instance.Category1FontSize = GongExtractString(valueExpr)
	case "Category1FontWeigth":
		instance.Category1FontWeigth = GongExtractString(valueExpr)
	case "Category1FontFamily":
		instance.Category1FontFamily = GongExtractString(valueExpr)
	case "Category1LetterSpacing":
		instance.Category1LetterSpacing = GongExtractString(valueExpr)
	case "Category2TypeFontSize":
		instance.Category2TypeFontSize = GongExtractString(valueExpr)
	case "Category2TypeFontWeigth":
		instance.Category2TypeFontWeigth = GongExtractString(valueExpr)
	case "Category2TypeFontFamily":
		instance.Category2TypeFontFamily = GongExtractString(valueExpr)
	case "Category2TypeLetterSpacing":
		instance.Category2TypeLetterSpacing = GongExtractString(valueExpr)
	case "Category2TypeRectAnchorType":
		GongUnmarshallEnum(&instance.Category2TypeRectAnchorType, valueExpr)
	case "Category2DominantBaselineType":
		GongUnmarshallEnum(&instance.Category2DominantBaselineType, valueExpr)
	case "Category2StrokeWidth":
		instance.Category2StrokeWidth = GongExtractFloat(valueExpr)
	case "Category3RectAnchorType":
		GongUnmarshallEnum(&instance.Category3RectAnchorType, valueExpr)
	case "Category3TextAnchorType":
		GongUnmarshallEnum(&instance.Category3TextAnchorType, valueExpr)
	case "Category3DominantBaselineType":
		GongUnmarshallEnum(&instance.Category3DominantBaselineType, valueExpr)
	case "Category3FontSize":
		instance.Category3FontSize = GongExtractString(valueExpr)
	case "Category3FontWeigth":
		instance.Category3FontWeigth = GongExtractString(valueExpr)
	case "Category3FontFamily":
		instance.Category3FontFamily = GongExtractString(valueExpr)
	case "Category3LetterSpacing":
		instance.Category3LetterSpacing = GongExtractString(valueExpr)
	case "InfluenceStrokeWidth":
		instance.InfluenceStrokeWidth = GongExtractFloat(valueExpr)
	case "InfluenceArrowSize":
		instance.InfluenceArrowSize = GongExtractFloat(valueExpr)
	case "InfluenceArrowStartOffset":
		instance.InfluenceArrowStartOffset = GongExtractFloat(valueExpr)
	case "InfluenceArrowEndOffset":
		instance.InfluenceArrowEndOffset = GongExtractFloat(valueExpr)
	case "InfluenceCornerRadius":
		instance.InfluenceCornerRadius = GongExtractFloat(valueExpr)
	case "InfluenceFontSize":
		instance.InfluenceFontSize = GongExtractString(valueExpr)
	case "InfluenceFontWeigth":
		instance.InfluenceFontWeigth = GongExtractString(valueExpr)
	case "InfluenceFontFamily":
		instance.InfluenceFontFamily = GongExtractString(valueExpr)
	case "InfluenceLetterSpacing":
		instance.InfluenceLetterSpacing = GongExtractString(valueExpr)
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
	case "SourceCategory1":
		GongUnmarshallPointer(&instance.SourceCategory1, valueExpr, identifierMap)
	case "SourceCategory2":
		GongUnmarshallPointer(&instance.SourceCategory2, valueExpr, identifierMap)
	case "SourceCategory3":
		GongUnmarshallPointer(&instance.SourceCategory3, valueExpr, identifierMap)
	case "TargetCategory1":
		GongUnmarshallPointer(&instance.TargetCategory1, valueExpr, identifierMap)
	case "TargetCategory2":
		GongUnmarshallPointer(&instance.TargetCategory2, valueExpr, identifierMap)
	case "TargetCategory3":
		GongUnmarshallPointer(&instance.TargetCategory3, valueExpr, identifierMap)
	case "IsHypothtical":
		instance.IsHypothtical = GongExtractBool(valueExpr)
	case "TextAtEndOfArrow":
		instance.TextAtEndOfArrow = GongExtractString(valueExpr)
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

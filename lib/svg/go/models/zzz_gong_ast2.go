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
type AnimateUnmarshaller struct{}

func (u *AnimateUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Animate)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AnimateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Animate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "AttributeName":
		instance.Name = GongExtractString(valueExpr)
	case "Values":
		instance.Name = GongExtractString(valueExpr)
	case "From":
		instance.Name = GongExtractString(valueExpr)
	case "To":
		instance.Name = GongExtractString(valueExpr)
	case "Dur":
		instance.Name = GongExtractString(valueExpr)
	case "RepeatCount":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type CircleUnmarshaller struct{}

func (u *CircleUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Circle)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *CircleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Circle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CX":
		instance.CX = GongExtractFloat(valueExpr)
	case "CY":
		instance.CY = GongExtractFloat(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animations":
		GongUnmarshallSliceOfPointers(&instance.Animations, valueExpr, identifierMap)
	}
	return nil
}

type ConditionUnmarshaller struct{}

func (u *ConditionUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Condition)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ConditionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Condition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ControlPointUnmarshaller struct{}

func (u *ControlPointUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ControlPoint)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ControlPointUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPoint)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "ClosestRect":
		GongUnmarshallPointer(&instance.ClosestRect, valueExpr, identifierMap)
	}
	return nil
}

type EllipseUnmarshaller struct{}

func (u *EllipseUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Ellipse)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *EllipseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ellipse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CX":
		instance.CX = GongExtractFloat(valueExpr)
	case "CY":
		instance.CY = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "RY":
		instance.RY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type LayerUnmarshaller struct{}

func (u *LayerUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Layer)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LayerUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Layer)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Rects":
		GongUnmarshallSliceOfPointers(&instance.Rects, valueExpr, identifierMap)
	case "Texts":
		GongUnmarshallSliceOfPointers(&instance.Texts, valueExpr, identifierMap)
	case "Circles":
		GongUnmarshallSliceOfPointers(&instance.Circles, valueExpr, identifierMap)
	case "Lines":
		GongUnmarshallSliceOfPointers(&instance.Lines, valueExpr, identifierMap)
	case "Ellipses":
		GongUnmarshallSliceOfPointers(&instance.Ellipses, valueExpr, identifierMap)
	case "Polylines":
		GongUnmarshallSliceOfPointers(&instance.Polylines, valueExpr, identifierMap)
	case "Polygones":
		GongUnmarshallSliceOfPointers(&instance.Polygones, valueExpr, identifierMap)
	case "Paths":
		GongUnmarshallSliceOfPointers(&instance.Paths, valueExpr, identifierMap)
	case "Links":
		GongUnmarshallSliceOfPointers(&instance.Links, valueExpr, identifierMap)
	case "RectLinkLinks":
		GongUnmarshallSliceOfPointers(&instance.RectLinkLinks, valueExpr, identifierMap)
	}
	return nil
}

type LineUnmarshaller struct{}

func (u *LineUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Line)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X1":
		instance.X1 = GongExtractFloat(valueExpr)
	case "Y1":
		instance.Y1 = GongExtractFloat(valueExpr)
	case "X2":
		instance.X2 = GongExtractFloat(valueExpr)
	case "Y2":
		instance.Y2 = GongExtractFloat(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	case "MouseClickX":
		instance.MouseClickX = GongExtractFloat(valueExpr)
	case "MouseClickY":
		instance.MouseClickY = GongExtractFloat(valueExpr)
	}
	return nil
}

type LinkUnmarshaller struct{}

func (u *LinkUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Link)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
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
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "IsBezierCurve":
		instance.IsBezierCurve = GongExtractBool(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "StartAnchorType":
		GongUnmarshallEnum(&instance.StartAnchorType, valueExpr)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "EndAnchorType":
		GongUnmarshallEnum(&instance.EndAnchorType, valueExpr)
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
	case "CornerRadius":
		instance.CornerRadius = GongExtractFloat(valueExpr)
	case "HasEndArrow":
		instance.HasEndArrow = GongExtractBool(valueExpr)
	case "EndArrowSize":
		instance.EndArrowSize = GongExtractFloat(valueExpr)
	case "EndArrowOffset":
		instance.EndArrowOffset = GongExtractFloat(valueExpr)
	case "HasStartArrow":
		instance.HasStartArrow = GongExtractBool(valueExpr)
	case "StartArrowSize":
		instance.StartArrowSize = GongExtractFloat(valueExpr)
	case "StartArrowOffset":
		instance.StartArrowOffset = GongExtractFloat(valueExpr)
	case "TextAtArrowStart":
		GongUnmarshallSliceOfPointers(&instance.TextAtArrowStart, valueExpr, identifierMap)
	case "TextAtArrowEnd":
		GongUnmarshallSliceOfPointers(&instance.TextAtArrowEnd, valueExpr, identifierMap)
	case "ControlPoints":
		GongUnmarshallSliceOfPointers(&instance.ControlPoints, valueExpr, identifierMap)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "MouseX":
		instance.MouseX = GongExtractFloat(valueExpr)
	case "MouseY":
		instance.MouseY = GongExtractFloat(valueExpr)
	case "MouseEventKey":
		GongUnmarshallEnum(&instance.MouseEventKey, valueExpr)
	}
	return nil
}

type LinkAnchoredTextUnmarshaller struct{}

func (u *LinkAnchoredTextUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(LinkAnchoredText)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LinkAnchoredTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkAnchoredText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Name = GongExtractString(valueExpr)
	case "AutomaticLayout":
		instance.AutomaticLayout = GongExtractBool(valueExpr)
	case "LinkAnchorType":
		GongUnmarshallEnum(&instance.LinkAnchorType, valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "FontWeight":
		instance.Name = GongExtractString(valueExpr)
	case "FontSize":
		instance.Name = GongExtractString(valueExpr)
	case "FontStyle":
		instance.Name = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.Name = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PathUnmarshaller struct{}

func (u *PathUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Path)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *PathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Path)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PointUnmarshaller struct{}

func (u *PointUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Point)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *PointUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Point)
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

type PolygoneUnmarshaller struct{}

func (u *PolygoneUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Polygone)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *PolygoneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Polygone)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PolylineUnmarshaller struct{}

func (u *PolylineUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Polyline)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *PolylineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Polyline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type RectUnmarshaller struct{}

func (u *RectUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Rect)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *RectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "HoveringTrigger":
		GongUnmarshallSliceOfPointers(&instance.HoveringTrigger, valueExpr, identifierMap)
	case "DisplayConditions":
		GongUnmarshallSliceOfPointers(&instance.DisplayConditions, valueExpr, identifierMap)
	case "Animations":
		GongUnmarshallSliceOfPointers(&instance.Animations, valueExpr, identifierMap)
	case "IsSelectable":
		instance.IsSelectable = GongExtractBool(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "CanHaveLeftHandle":
		instance.CanHaveLeftHandle = GongExtractBool(valueExpr)
	case "HasLeftHandle":
		instance.HasLeftHandle = GongExtractBool(valueExpr)
	case "CanHaveRightHandle":
		instance.CanHaveRightHandle = GongExtractBool(valueExpr)
	case "HasRightHandle":
		instance.HasRightHandle = GongExtractBool(valueExpr)
	case "CanHaveTopHandle":
		instance.CanHaveTopHandle = GongExtractBool(valueExpr)
	case "HasTopHandle":
		instance.HasTopHandle = GongExtractBool(valueExpr)
	case "IsScalingProportionally":
		instance.IsScalingProportionally = GongExtractBool(valueExpr)
	case "CanHaveBottomHandle":
		instance.CanHaveBottomHandle = GongExtractBool(valueExpr)
	case "HasBottomHandle":
		instance.HasBottomHandle = GongExtractBool(valueExpr)
	case "CanMoveHorizontaly":
		instance.CanMoveHorizontaly = GongExtractBool(valueExpr)
	case "CanMoveVerticaly":
		instance.CanMoveVerticaly = GongExtractBool(valueExpr)
	case "RectAnchoredTexts":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredTexts, valueExpr, identifierMap)
	case "RectAnchoredRects":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredRects, valueExpr, identifierMap)
	case "RectAnchoredPaths":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredPaths, valueExpr, identifierMap)
	case "ChangeColorWhenHovered":
		instance.ChangeColorWhenHovered = GongExtractBool(valueExpr)
	case "ColorWhenHovered":
		instance.Name = GongExtractString(valueExpr)
	case "OriginalColor":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacityWhenHovered":
		instance.FillOpacityWhenHovered = GongExtractFloat(valueExpr)
	case "OriginalFillOpacity":
		instance.OriginalFillOpacity = GongExtractFloat(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.Name = GongExtractString(valueExpr)
	case "ToolTipPosition":
		GongUnmarshallEnum(&instance.ToolTipPosition, valueExpr)
	case "MouseX":
		instance.MouseX = GongExtractFloat(valueExpr)
	case "MouseY":
		instance.MouseY = GongExtractFloat(valueExpr)
	case "MouseEventKey":
		GongUnmarshallEnum(&instance.MouseEventKey, valueExpr)
	}
	return nil
}

type RectAnchoredPathUnmarshaller struct{}

func (u *RectAnchoredPathUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RectAnchoredPath)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *RectAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredPath)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Name = GongExtractString(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "ScalePropotionnally":
		instance.ScalePropotionnally = GongExtractBool(valueExpr)
	case "AppliedScaling":
		instance.AppliedScaling = GongExtractFloat(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredRectUnmarshaller struct{}

func (u *RectAnchoredRectUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RectAnchoredRect)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *RectAnchoredRectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredRect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "WidthFollowRect":
		instance.WidthFollowRect = GongExtractBool(valueExpr)
	case "HeightFollowRect":
		instance.HeightFollowRect = GongExtractBool(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredTextUnmarshaller struct{}

func (u *RectAnchoredTextUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RectAnchoredText)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *RectAnchoredTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Name = GongExtractString(valueExpr)
	case "FontWeight":
		instance.Name = GongExtractString(valueExpr)
	case "FontSize":
		instance.Name = GongExtractString(valueExpr)
	case "FontStyle":
		instance.Name = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.Name = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "TextAnchorType":
		GongUnmarshallEnum(&instance.TextAnchorType, valueExpr)
	case "DominantBaseline":
		GongUnmarshallEnum(&instance.DominantBaseline, valueExpr)
	case "WritingMode":
		GongUnmarshallEnum(&instance.WritingMode, valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type RectLinkLinkUnmarshaller struct{}

func (u *RectLinkLinkUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(RectLinkLink)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *RectLinkLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectLinkLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "TargetAnchorPosition":
		instance.TargetAnchorPosition = GongExtractFloat(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SVGUnmarshaller struct{}

func (u *SVGUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SVG)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *SVGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SVG)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Layers":
		GongUnmarshallSliceOfPointers(&instance.Layers, valueExpr, identifierMap)
	case "DrawingState":
		GongUnmarshallEnum(&instance.DrawingState, valueExpr)
	case "StartRect":
		GongUnmarshallPointer(&instance.StartRect, valueExpr, identifierMap)
	case "EndRect":
		GongUnmarshallPointer(&instance.EndRect, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsSVGFrontEndFileGenerated":
		instance.IsSVGFrontEndFileGenerated = GongExtractBool(valueExpr)
	case "IsSVGBackEndFileGenerated":
		instance.IsSVGBackEndFileGenerated = GongExtractBool(valueExpr)
	case "DefaultDirectoryForGeneratedImages":
		instance.Name = GongExtractString(valueExpr)
	case "IsControlBannerHidden":
		instance.IsControlBannerHidden = GongExtractBool(valueExpr)
	case "OverrideWidth":
		instance.OverrideWidth = GongExtractBool(valueExpr)
	case "OverriddenWidth":
		instance.OverriddenWidth = GongExtractFloat(valueExpr)
	case "OverrideHeight":
		instance.OverrideHeight = GongExtractBool(valueExpr)
	case "OverriddenHeight":
		instance.OverriddenHeight = GongExtractFloat(valueExpr)
	}
	return nil
}

type SvgTextUnmarshaller struct{}

func (u *SvgTextUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SvgText)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *SvgTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SvgText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TextUnmarshaller struct{}

func (u *TextUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Text)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *TextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Content":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Name = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.Name = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.Name = GongExtractString(valueExpr)
	case "Transform":
		instance.Name = GongExtractString(valueExpr)
	case "FontWeight":
		instance.Name = GongExtractString(valueExpr)
	case "FontSize":
		instance.Name = GongExtractString(valueExpr)
	case "FontStyle":
		instance.Name = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.Name = GongExtractString(valueExpr)
	case "FontFamily":
		instance.Name = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

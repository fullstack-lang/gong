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
type AsSplitUnmarshaller struct{}

func (u *AsSplitUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AsSplit)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AsSplitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AsSplit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "AsSplitAreas":
		GongUnmarshallSliceOfPointers(&instance.AsSplitAreas, valueExpr, identifierMap)
	}
	return nil
}

type AsSplitAreaUnmarshaller struct{}

func (u *AsSplitAreaUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AsSplitArea)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AsSplitAreaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AsSplitArea)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowNameInHeader":
		instance.ShowNameInHeader = GongExtractBool(valueExpr)
	case "Size":
		instance.Size = GongExtractFloat(valueExpr)
	case "IsAny":
		instance.IsAny = GongExtractBool(valueExpr)
	case "AsSplit":
		GongUnmarshallPointer(&instance.AsSplit, valueExpr, identifierMap)
	case "Button":
		GongUnmarshallPointer(&instance.Button, valueExpr, identifierMap)
	case "Cursor":
		GongUnmarshallPointer(&instance.Cursor, valueExpr, identifierMap)
	case "Form":
		GongUnmarshallPointer(&instance.Form, valueExpr, identifierMap)
	case "Load":
		GongUnmarshallPointer(&instance.Load, valueExpr, identifierMap)
	case "Markdown":
		GongUnmarshallPointer(&instance.Markdown, valueExpr, identifierMap)
	case "Slider":
		GongUnmarshallPointer(&instance.Slider, valueExpr, identifierMap)
	case "Split":
		GongUnmarshallPointer(&instance.Split, valueExpr, identifierMap)
	case "Svg":
		GongUnmarshallPointer(&instance.Svg, valueExpr, identifierMap)
	case "Table":
		GongUnmarshallPointer(&instance.Table, valueExpr, identifierMap)
	case "Tone":
		GongUnmarshallPointer(&instance.Tone, valueExpr, identifierMap)
	case "Tree":
		GongUnmarshallPointer(&instance.Tree, valueExpr, identifierMap)
	case "Xlsx":
		GongUnmarshallPointer(&instance.Xlsx, valueExpr, identifierMap)
	case "HasDiv":
		instance.HasDiv = GongExtractBool(valueExpr)
	case "DivStyle":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ButtonUnmarshaller struct{}

func (u *ButtonUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Button)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ButtonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Button)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type CursorUnmarshaller struct{}

func (u *CursorUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Cursor)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *CursorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Cursor)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	case "Style":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type FavIconUnmarshaller struct{}

func (u *FavIconUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(FavIcon)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *FavIconUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FavIcon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SVG":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type FormUnmarshaller struct{}

func (u *FormUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Form)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *FormUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Form)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type LoadUnmarshaller struct{}

func (u *LoadUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Load)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LoadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Load)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type LogoOnTheLeftUnmarshaller struct{}

func (u *LogoOnTheLeftUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(LogoOnTheLeft)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LogoOnTheLeftUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LogoOnTheLeft)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractInt(valueExpr)
	case "Height":
		instance.Height = GongExtractInt(valueExpr)
	case "SVG":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type LogoOnTheRightUnmarshaller struct{}

func (u *LogoOnTheRightUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(LogoOnTheRight)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *LogoOnTheRightUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LogoOnTheRight)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractInt(valueExpr)
	case "Height":
		instance.Height = GongExtractInt(valueExpr)
	case "SVG":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type MarkdownUnmarshaller struct{}

func (u *MarkdownUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Markdown)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *MarkdownUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Markdown)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SliderUnmarshaller struct{}

func (u *SliderUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Slider)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *SliderUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slider)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SplitUnmarshaller struct{}

func (u *SplitUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Split)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *SplitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Split)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SvgUnmarshaller struct{}

func (u *SvgUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Svg)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *SvgUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Svg)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	case "Style":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TableUnmarshaller struct{}

func (u *TableUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Table)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *TableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Table)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TitleUnmarshaller struct{}

func (u *TitleUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Title)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *TitleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Title)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ToneUnmarshaller struct{}

func (u *ToneUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tone)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ToneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tone)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TreeUnmarshaller struct{}

func (u *TreeUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Tree)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *TreeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tree)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ViewUnmarshaller struct{}

func (u *ViewUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(View)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *ViewUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*View)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowViewName":
		instance.ShowViewName = GongExtractBool(valueExpr)
	case "RootAsSplitAreas":
		GongUnmarshallSliceOfPointers(&instance.RootAsSplitAreas, valueExpr, identifierMap)
	case "IsSelectedView":
		instance.IsSelectedView = GongExtractBool(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	}
	return nil
}

type XlsxUnmarshaller struct{}

func (u *XlsxUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Xlsx)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *XlsxUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Xlsx)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackName":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

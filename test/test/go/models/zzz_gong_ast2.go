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
type AstructUnmarshaller struct{}

func (u *AstructUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Astruct)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AstructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Astruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Associationtob":
		GongUnmarshallPointer(&instance.Associationtob, valueExpr, identifierMap)
	case "Anarrayofb":
		GongUnmarshallSliceOfPointers(&instance.Anarrayofb, valueExpr, identifierMap)
	case "Anotherassociationtob_2":
		GongUnmarshallPointer(&instance.Anotherassociationtob_2, valueExpr, identifierMap)
	case "Date":
		if bl, ok := valueExpr.(*ast.BasicLit); ok {
			instance.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
		}
	case "Date2":
		if bl, ok := valueExpr.(*ast.BasicLit); ok {
			instance.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
		}
	case "Booleanfield":
		instance.Booleanfield = GongExtractBool(valueExpr)
	case "Aenum":
		GongUnmarshallEnum(&instance.Aenum, valueExpr)
	case "Aenum_2":
		GongUnmarshallEnum(&instance.Aenum_2, valueExpr)
	case "Benum":
		GongUnmarshallEnum(&instance.Benum, valueExpr)
	case "CEnum":
		GongUnmarshallEnum(&instance.CEnum, valueExpr)
	case "CName":
		instance.Name = GongExtractString(valueExpr)
	case "CFloatfield":
		instance.CFloatfield = GongExtractFloat(valueExpr)
	case "Bstruct":
		GongUnmarshallPointer(&instance.Bstruct, valueExpr, identifierMap)
	case "Bstruct2":
		GongUnmarshallPointer(&instance.Bstruct2, valueExpr, identifierMap)
	case "Dstruct":
		GongUnmarshallPointer(&instance.Dstruct, valueExpr, identifierMap)
	case "Dstruct2":
		GongUnmarshallPointer(&instance.Dstruct2, valueExpr, identifierMap)
	case "Dstruct3":
		GongUnmarshallPointer(&instance.Dstruct3, valueExpr, identifierMap)
	case "Dstruct4":
		GongUnmarshallPointer(&instance.Dstruct4, valueExpr, identifierMap)
	case "Dstruct4s":
		GongUnmarshallSliceOfPointers(&instance.Dstruct4s, valueExpr, identifierMap)
	case "Floatfield":
		instance.Floatfield = GongExtractFloat(valueExpr)
	case "Intfield":
		instance.Intfield = GongExtractInt(valueExpr)
	case "Anotherbooleanfield":
		instance.Anotherbooleanfield = GongExtractBool(valueExpr)
	case "Duration1":
		instance.Duration1 = time.Duration(GongExtractInt(valueExpr))
	case "Anarrayofa":
		GongUnmarshallSliceOfPointers(&instance.Anarrayofa, valueExpr, identifierMap)
	case "Anotherarrayofb":
		GongUnmarshallSliceOfPointers(&instance.Anotherarrayofb, valueExpr, identifierMap)
	case "AnarrayofbUse":
		GongUnmarshallSliceOfPointers(&instance.AnarrayofbUse, valueExpr, identifierMap)
	case "Anarrayofb2Use":
		GongUnmarshallSliceOfPointers(&instance.Anarrayofb2Use, valueExpr, identifierMap)
	case "AnAstruct":
		GongUnmarshallPointer(&instance.AnAstruct, valueExpr, identifierMap)
	case "TextFieldBespokeSize":
		instance.Name = GongExtractString(valueExpr)
	case "TextArea":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type AstructBstruct2UseUnmarshaller struct{}

func (u *AstructBstruct2UseUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AstructBstruct2Use)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AstructBstruct2UseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AstructBstruct2Use)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Bstrcut2":
		GongUnmarshallPointer(&instance.Bstrcut2, valueExpr, identifierMap)
	}
	return nil
}

type AstructBstructUseUnmarshaller struct{}

func (u *AstructBstructUseUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AstructBstructUse)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AstructBstructUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AstructBstructUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Bstruct2":
		GongUnmarshallPointer(&instance.Bstruct2, valueExpr, identifierMap)
	}
	return nil
}

type BstructUnmarshaller struct{}

func (u *BstructUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Bstruct)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *BstructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bstruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Floatfield":
		instance.Floatfield = GongExtractFloat(valueExpr)
	case "Floatfield2":
		instance.Floatfield2 = GongExtractFloat(valueExpr)
	case "Intfield":
		instance.Intfield = GongExtractInt(valueExpr)
	}
	return nil
}

type DstructUnmarshaller struct{}

func (u *DstructUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Dstruct)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *DstructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dstruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Anarrayofb":
		GongUnmarshallSliceOfPointers(&instance.Anarrayofb, valueExpr, identifierMap)
	case "Gstruct":
		GongUnmarshallPointer(&instance.Gstruct, valueExpr, identifierMap)
	case "Gstructs":
		GongUnmarshallSliceOfPointers(&instance.Gstructs, valueExpr, identifierMap)
	}
	return nil
}

type F0123456789012345678901234567890Unmarshaller struct{}

func (u *F0123456789012345678901234567890Unmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(F0123456789012345678901234567890)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *F0123456789012345678901234567890Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*F0123456789012345678901234567890)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Date":
		if bl, ok := valueExpr.(*ast.BasicLit); ok {
			instance.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
		}
	}
	return nil
}

type GstructUnmarshaller struct{}

func (u *GstructUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Gstruct)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *GstructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Gstruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Floatfield":
		instance.Floatfield = GongExtractFloat(valueExpr)
	case "Floatfield2":
		instance.Floatfield2 = GongExtractFloat(valueExpr)
	case "Intfield":
		instance.Intfield = GongExtractInt(valueExpr)
	}
	return nil
}

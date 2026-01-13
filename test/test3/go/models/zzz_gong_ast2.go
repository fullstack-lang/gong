// package models
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

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// ModelUnmarshaller abstracts the logic for setting fields on a staged instance
type ModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, instanceName string, preserveOrder bool) (any, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance any, fieldName string, valueExpr ast.Expr, identifierMap map[string]any) error
}

// Unmarshallers is the registry of all model unmarshallers
var Unmarshallers = make(map[string]ModelUnmarshaller)

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
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]any)

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
						if unmarshaller, exists := Unmarshallers[typeName]; exists {
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
								if unmarshaller, exists := Unmarshallers[typeName]; exists {
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

func ExtractString(expr ast.Expr) string {
	if bl, ok := expr.(*ast.BasicLit); ok {
		return strings.Trim(bl.Value, "\"`")
	}
	return ""
}

func ExtractInt(expr ast.Expr) int {
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

func ExtractFloat(expr ast.Expr) float64 {
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

func ExtractBool(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "true"
	}
	return false
}

// ------------------------------------------------------------------------------------------------
// GENERATED MODULAR HANDLERS
// ------------------------------------------------------------------------------------------------

func init() {
	Unmarshallers["A"] = &AUnmarshaller{}
	Unmarshallers["B"] = &BUnmarshaller{}
}

// --- A Unmarshaller ---
type AUnmarshaller struct{}

func (u *AUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (any, error) {
	instance := new(A)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		instance.Stage(stage)
	}
	return instance, nil
}

func (u *AUnmarshaller) UnmarshallField(stage *Stage, i any, fieldName string, valueExpr ast.Expr, identifierMap map[string]any) error {
	instance := i.(*A)
	switch fieldName {
	case "Name":
		instance.Name = ExtractString(valueExpr)
	case "Date":
		if bl, ok := valueExpr.(*ast.BasicLit); ok {
			instance.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
		}
	case "FloatValue":
		instance.FloatValue = ExtractFloat(valueExpr)
	case "IntValue":
		instance.IntValue = ExtractInt(valueExpr)
	case "Duration":
		instance.Duration = time.Duration(ExtractInt(valueExpr))
	case "EnumString":
		// Handle Enum (Usually a SelectorExpr like models.Value1)
		if sel, ok := valueExpr.(*ast.SelectorExpr); ok {
			// In case the generated code requires FromCodeString (standard Gong pattern)
			// We can simulate it or direct cast if the method isn't generated yet.
			// Here we assume direct cast for simplicity or standard string extraction.
			instance.EnumString = EnumTypeString(sel.Sel.Name)
		} else {
			// Fallback if it is a string literal
			instance.EnumString = EnumTypeString(ExtractString(valueExpr))
		}
	case "EnumInt":
		if sel, ok := valueExpr.(*ast.SelectorExpr); ok {
			enumValue := sel.Sel.Name
			var val EnumTypeInt
			err := (&val).FromCodeString(enumValue)
			if err != nil {
				log.Fatalln(err)
			}
			instance.EnumInt = EnumTypeInt(val)
		}
	case "B":
		if ident, ok := valueExpr.(*ast.Ident); ok {
			if val, ok := identifierMap[ident.Name]; ok {
				instance.B = val.(*B)
			}
		}
	case "Bs":
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
					start := ExtractInt(call.Args[1])
					end := ExtractInt(call.Args[2])
					instance.Bs = slices.Delete(instance.Bs, start, end)
				} else if funcName == "Insert" && len(call.Args) == 3 {
					index := ExtractInt(call.Args[1])
					if ident, ok := call.Args[2].(*ast.Ident); ok {
						if val, ok := identifierMap[ident.Name]; ok {
							instance.Bs = slices.Insert(instance.Bs, index, val.(*B))
						}
					}
				}
			} else if funcName == "append" {
				if len(call.Args) >= 2 {
					if ident, ok := call.Args[len(call.Args)-1].(*ast.Ident); ok {
						if val, ok := identifierMap[ident.Name]; ok {
							instance.Bs = append(instance.Bs, val.(*B))
						}
					}
				}
			}
		}
	}
	return nil
}

// --- B Unmarshaller ---
type BUnmarshaller struct{}

func (u *BUnmarshaller) Initialize(stage *Stage, instanceName string, preserveOrder bool) (any, error) {
	instance := new(B)
	instance.Name = instanceName
	instance.Stage(stage)
	return instance, nil
}

func (u *BUnmarshaller) UnmarshallField(stage *Stage, i any, fieldName string, valueExpr ast.Expr, identifierMap map[string]any) error {
	instance := i.(*B)
	switch fieldName {
	case "Name":
		instance.Name = ExtractString(valueExpr)
	}
	return nil
}

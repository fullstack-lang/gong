package runtime

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

// GONG__ExpressionType represents the type of an expression in AST parsing
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

var middleUintRegex = regexp.MustCompile(`__.*?__(\d+)_.*`)

// NotificationProbeIF abstracts notifications from probes
type NotificationProbeIF interface {
	AddNotification(time.Time, string)
}

// ParseAstFile parses pathToFile into an AST file and token.FileSet
func ParseAstFile(pathToFile string) (*ast.File, *token.FileSet, error) {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return nil, nil, errors.New("Path does not exist " + fileOfInterest + " ;")
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return nil, nil, errors.New("Unable to parser " + errParser.Error())
	}
	return inFile, fset, nil
}

// ParseAstEmbeddedFile parses an embedded Go source file into an AST file and token.FileSet
func ParseAstEmbeddedFile(directory embed.FS, pathToFile string, stageName string) (*ast.File, *token.FileSet, error) {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return nil, nil, errors.New(stageName + "; Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return nil, nil, errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}
	return inFile, fset, nil
}

// ParseAstString parses a Go source string blob into an AST file and token.FileSet
func ParseAstString(blob string) (*ast.File, *token.FileSet, error) {
	fileString := "package main\nfunc _() {\n" + blob + "\n}"
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", fileString, parser.ParseComments)
	if errParser != nil {
		return nil, nil, errors.New("Unable to parser " + errParser.Error())
	}
	return inFile, fset, nil
}

// CheckModuleVersion validates the module version header against the runtime parser module version
func CheckModuleVersion(probe NotificationProbeIF, inFile *ast.File) {
	var fileModuleVersion string
	for _, commentGroup := range inFile.Comments {
		for _, comment := range commentGroup.List {
			if after, ok := strings.CutPrefix(comment.Text, "// go module version: "); ok {
				fileModuleVersion = after
			}
		}
	}

	var parserModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		parserModuleVersion = buildInfo.Main.Version
	}

	if fileModuleVersion != "" && parserModuleVersion != "" && fileModuleVersion != parserModuleVersion {
		if probe != nil {
			probe.AddNotification(time.Now(), fmt.Sprintf("Warning: file module version '%s' does not match parser module version '%s'", fileModuleVersion, parserModuleVersion))
		}
	}
}

// WalkAstFile traverses the AST and invokes callbacks for definitions, assignments, unstaging, and commits
func WalkAstFile(
	inFile *ast.File,
	onDefine func(identName string, typeName string, instanceName string),
	onAssign func(identName string, fieldName string, valueExpr ast.Expr),
	onUnstage func(identName string),
	onCommit func(),
) error {
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

					if onDefine != nil {
						onDefine(ident.Name, typeName, instanceName)
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if onAssign != nil {
							onAssign(ident.Name, selExpr.Sel.Name, node.Rhs[0])
						}
					}
				}
			}
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if onUnstage != nil {
								onUnstage(ident.Name)
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if onCommit != nil {
								onCommit()
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

// ExtractString extracts string literal or concatenated string from AST expression
func ExtractString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			if unquoted, err := strconv.Unquote(e.Value); err == nil {
				return unquoted
			}
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			left := ExtractString(e.X)
			right := ExtractString(e.Y)
			return left + right
		}
	}
	return ""
}

// ExtractInt extracts an integer from AST expression (including negative integers)
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

// ExtractMiddleUint extracts a middle uint identifier from a formatted string
func ExtractMiddleUint(input string) (uint, error) {
	matches := middleUintRegex.FindStringSubmatch(input)
	if len(matches) < 2 {
		return 0, fmt.Errorf("pattern not found in string: %s", input)
	}
	result, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to int: %v", matches[1], err)
	}
	return uint(result), nil
}

// ExtractFloat extracts float64 from AST expression (including negative floats)
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

// ExtractBool extracts boolean value from AST expression
func ExtractBool(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "true"
	}
	return false
}

// ExtractExpr extracts expression representation from AST expression
func ExtractExpr(expr ast.Expr) any {
	switch v := expr.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.CompositeLit:
		if sel, ok := v.Type.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				return id.Name + "." + sel.Sel.Name + "{}"
			}
		}
	case *ast.SelectorExpr:
		if cl, ok := v.X.(*ast.CompositeLit); ok {
			if sel, ok := cl.Type.(*ast.SelectorExpr); ok {
				if id, ok := sel.X.(*ast.Ident); ok {
					return id.Name + "." + sel.Sel.Name + "{}." + v.Sel.Name
				}
			}
		}
		if id, ok := v.X.(*ast.Ident); ok {
			return id.Name + "." + v.Sel.Name
		}
	case *ast.CallExpr:
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

// ExtractDate extracts time.Time from AST call expression
func ExtractDate(expr ast.Expr) time.Time {
	if call, ok := expr.(*ast.CallExpr); ok {
		if len(call.Args) == 2 {
			if bl, ok := call.Args[1].(*ast.BasicLit); ok {
				t, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
				return t
			}
		}
	}
	return time.Time{}
}

// UnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func UnmarshallSliceOfPointers[T any, M any](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]M,
) error {
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
				if end > len(*slice) {
					return fmt.Errorf("index out of bounds: %d for len %d", end, len(*slice))
				}
				*slice = slices.Delete(*slice, start, end)
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := ExtractInt(call.Args[1])
				if index > len(*slice) {
					return fmt.Errorf("index out of bounds: %d for len %d", index, len(*slice))
				}
				if ident, ok := call.Args[2].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = slices.Insert(*slice, index, any(val).(T))
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
							*slice = append(*slice, any(val).(T))
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

// UnmarshallPointer handles assignment of a single pointer field
func UnmarshallPointer[T any, M any](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]M,
) {
	if ident, ok := valueExpr.(*ast.Ident); ok {
		if ident.Name == "nil" {
			var zero T
			*ptr = zero
			return
		}
		if val, ok := identifierMap[ident.Name]; ok {
			*ptr = any(val).(T)
		}
	}
}

// UnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func UnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr,
) {
	if sel, ok := valueExpr.(*ast.SelectorExpr); ok {
		if err := ptr.FromCodeString(sel.Sel.Name); err != nil {
			log.Printf("UnmarshallEnum: Error parsing code string '%s': %v", sel.Sel.Name, err)
		}
	} else {
		valStr := ExtractString(valueExpr)
		if valStr != "" {
			if err := ptr.FromCodeString(valStr); err != nil {
				log.Printf("UnmarshallEnum: Error parsing string literal '%s': %v", valStr, err)
			}
		}
	}
}

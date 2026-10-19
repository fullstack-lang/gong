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
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var (
	_ = errors.New
	_ = parser.ParseFile
	_ = filepath.Base
	_ = slices.Index[[]int, int]
	_ = strings.Trim
	_ = time.Hour
)

// swagger:ignore
type GONG__ExpressionType = gong_runtime.GONG__ExpressionType

const (
	GONG__STRUCT_INSTANCE      = gong_runtime.GONG__STRUCT_INSTANCE
	GONG__FIELD_OR_CONST_VALUE = gong_runtime.GONG__FIELD_OR_CONST_VALUE
	GONG__FIELD_VALUE          = gong_runtime.GONG__FIELD_VALUE
	GONG__ENUM_CAST_INT        = gong_runtime.GONG__ENUM_CAST_INT
	GONG__ENUM_CAST_STRING     = gong_runtime.GONG__ENUM_CAST_STRING
	GONG__IDENTIFIER_CONST     = gong_runtime.GONG__IDENTIFIER_CONST
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// GongModelUnmarshaller abstracts the logic for setting fields on a staged instance
type GongModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

type ModelUnmarshaller = GongModelUnmarshaller

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func (stage *Stage) ParseAstFile(pathToFile string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstFile(pathToFile)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func (stage *Stage) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	inFile, fset, err := gong_runtime.ParseAstEmbeddedFile(directory, pathToFile, stage.GetName())
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string
func (stage *Stage) ParseAstString(blob string, preserveOrder bool) error {
	var inFile *ast.File
	var fset *token.FileSet
	var err error
	if strings.HasPrefix(strings.TrimSpace(blob), "package ") || strings.Contains(blob, "\npackage ") {
		fset = token.NewFileSet()
		inFile, err = parser.ParseFile(fset, "", blob, parser.ParseComments)
		if err != nil {
			return errors.New("Unable to parse " + err.Error())
		}
	} else {
		inFile, fset, err = gong_runtime.ParseAstString(blob)
		if err != nil {
			return err
		}
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances across stage and its sub-stages
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	identifierMap := make(map[string]any)

	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
			identifierMap["__models"+instance.GongGetIdentifier(stage)] = instance
		}
	}

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/test/test1/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
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

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					} else if pkgAlias == "" || pkgAlias == "models" || len(inFile.Imports) == 0 {
						pkgAlias = "models"
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "Astruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Astruct), stage, ident.Name, instanceName, preserveOrder)
				case "AstructBstruct2Use":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AstructBstruct2Use), stage, ident.Name, instanceName, preserveOrder)
				case "AstructBstructUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AstructBstructUse), stage, ident.Name, instanceName, preserveOrder)
				case "Bstruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bstruct), stage, ident.Name, instanceName, preserveOrder)
				case "Dstruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dstruct), stage, ident.Name, instanceName, preserveOrder)
				case "F0123456789012345678901234567890":
					identifierMap[ident.Name] = __gong__stageSetInit(new(F0123456789012345678901234567890), stage, ident.Name, instanceName, preserveOrder)
				case "Gstruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Gstruct), stage, ident.Name, instanceName, preserveOrder)
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
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *Astruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Field":
						inst.Field = GongExtractExpr(rhs)
					case "Associationtob":
						__gong__assignPointer(&inst.Associationtob, rhs, identifierMap)
					case "Anarrayofb":
						__gong__assignSliceOfPointers(&inst.Anarrayofb, rhs, identifierMap)
					case "Anotherassociationtob_2":
						__gong__assignPointer(&inst.Anotherassociationtob_2, rhs, identifierMap)
					case "Date":
						inst.Date = GongExtractDate(rhs)
					case "Date2":
						inst.Date2 = GongExtractDate(rhs)
					case "Booleanfield":
						inst.Booleanfield = GongExtractBool(rhs)
					case "Aenum":
						GongUnmarshallEnum(&inst.Aenum, rhs)
					case "Aenum_2":
						GongUnmarshallEnum(&inst.Aenum_2, rhs)
					case "Benum":
						GongUnmarshallEnum(&inst.Benum, rhs)
					case "CEnum":
						GongUnmarshallEnum(&inst.CEnum, rhs)
					case "CName":
						inst.CName = GongExtractString(rhs)
					case "CFloatfield":
						inst.CFloatfield = GongExtractFloat(rhs)
					case "Bstruct":
						__gong__assignPointer(&inst.Bstruct, rhs, identifierMap)
					case "Bstruct2":
						__gong__assignPointer(&inst.Bstruct2, rhs, identifierMap)
					case "Dstruct":
						__gong__assignPointer(&inst.Dstruct, rhs, identifierMap)
					case "Dstruct2":
						__gong__assignPointer(&inst.Dstruct2, rhs, identifierMap)
					case "Dstruct3":
						__gong__assignPointer(&inst.Dstruct3, rhs, identifierMap)
					case "Dstruct4":
						__gong__assignPointer(&inst.Dstruct4, rhs, identifierMap)
					case "Dstruct4s":
						__gong__assignSliceOfPointers(&inst.Dstruct4s, rhs, identifierMap)
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "Anotherbooleanfield":
						inst.Anotherbooleanfield = GongExtractBool(rhs)
					case "Duration1":
						inst.Duration1 = time.Duration(GongExtractInt(rhs))
					case "Anarrayofa":
						__gong__assignSliceOfPointers(&inst.Anarrayofa, rhs, identifierMap)
					case "Anotherarrayofb":
						__gong__assignSliceOfPointers(&inst.Anotherarrayofb, rhs, identifierMap)
					case "AnarrayofbUse":
						__gong__assignSliceOfPointers(&inst.AnarrayofbUse, rhs, identifierMap)
					case "Anarrayofb2Use":
						__gong__assignSliceOfPointers(&inst.Anarrayofb2Use, rhs, identifierMap)
					case "AnAstruct":
						__gong__assignPointer(&inst.AnAstruct, rhs, identifierMap)
					case "TextFieldBespokeSize":
						inst.TextFieldBespokeSize = GongExtractString(rhs)
					case "TextArea":
						inst.TextArea = GongExtractString(rhs)
					}
				case *AstructBstruct2Use:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Bstrcut2":
						__gong__assignPointer(&inst.Bstrcut2, rhs, identifierMap)
					}
				case *AstructBstructUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Bstruct2":
						__gong__assignPointer(&inst.Bstruct2, rhs, identifierMap)
					}
				case *Bstruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Floatfield2":
						inst.Floatfield2 = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "Struct":
						inst.Struct = GongExtractExpr(rhs)
					case "ToBeIgnored":
						inst.ToBeIgnored = GongExtractInt(rhs)
					}
				case *Dstruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Anarrayofb":
						__gong__assignSliceOfPointers(&inst.Anarrayofb, rhs, identifierMap)
					case "Gstruct":
						__gong__assignPointer(&inst.Gstruct, rhs, identifierMap)
					case "Gstructs":
						__gong__assignSliceOfPointers(&inst.Gstructs, rhs, identifierMap)
					}
				case *F0123456789012345678901234567890:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						inst.Date = GongExtractDate(rhs)
					}
				case *Gstruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Floatfield2":
						inst.Floatfield2 = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "Struct":
						inst.Struct = GongExtractExpr(rhs)
					case "ToBeIgnored":
						inst.ToBeIgnored = GongExtractInt(rhs)
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
								if inst, ok := instance.(GongstructIF); ok {
									inst.UnstageVoid(stage)
								}
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.GetNavigationMode() != GongNavigationModeNavigating {
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
								stage.RunBeforeCommitHooks()
								stage.RunAfterCommitHooks()
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

// --- Generic Helpers for Unmarshallers (delegating to gong_runtime) ---

var GongExtractString = gong_runtime.ExtractString
var GongExtractInt = gong_runtime.ExtractInt
var GongExtractFloat = gong_runtime.ExtractFloat
var GongExtractBool = gong_runtime.ExtractBool
var GongExtractExpr = gong_runtime.ExtractExpr
var GongExtractDate = gong_runtime.ExtractDate

func __gong__extractMiddleUint(input string) (uint, error) {
	return gong_runtime.ExtractMiddleUint(input)
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok {
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
				if end <= len(*slice) {
					*slice = slices.Delete(*slice, start, end)
				}
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index <= len(*slice) {
					if rIdent, ok := call.Args[2].(*ast.Ident); ok {
						if target, ok := identifierMap[rIdent.Name]; ok {
							if typedTarget, ok := target.(*T); ok {
								*slice = slices.Insert(*slice, index, typedTarget)
							}
						}
					}
				}
			}
		} else if funcName == "append" && len(call.Args) >= 2 {
			for i := 1; i < len(call.Args); i++ {
				if rIdent, ok := call.Args[i].(*ast.Ident); ok {
					if target, ok := identifierMap[rIdent.Name]; ok {
						if typedTarget, ok := target.(*T); ok {
							*slice = append(*slice, typedTarget)
						}
					}
				}
			}
		}
	}
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T GongstructPtr](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {
	return gong_runtime.UnmarshallSliceOfPointers(slice, valueExpr, identifierMap)
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T GongstructPtr](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {
	gong_runtime.UnmarshallPointer(ptr, valueExpr, identifierMap)
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {
	gong_runtime.UnmarshallEnum(ptr, valueExpr)
}

// GongInitialize initializes a staged instance, sets its name, and stages it
func GongInitialize[P interface {
	GongstructIF
	StagePreserveOrder(stage *Stage, order uint)
}](instance P, stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if newOrder, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

// insertion point per named struct
type AstructUnmarshaller struct{}

func (u *AstructUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Astruct), stage, identifier, instanceName, preserveOrder)
}

func (u *AstructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Astruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Field":
		instance.Field = GongExtractExpr(valueExpr)
	case "Associationtob":
		GongUnmarshallPointer(&instance.Associationtob, valueExpr, identifierMap)
	case "Anarrayofb":
		GongUnmarshallSliceOfPointers(&instance.Anarrayofb, valueExpr, identifierMap)
	case "Anotherassociationtob_2":
		GongUnmarshallPointer(&instance.Anotherassociationtob_2, valueExpr, identifierMap)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	case "Date2":
		instance.Date2 = GongExtractDate(valueExpr)
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
		instance.CName = GongExtractString(valueExpr)
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
		instance.TextFieldBespokeSize = GongExtractString(valueExpr)
	case "TextArea":
		instance.TextArea = GongExtractString(valueExpr)
	}
	return nil
}

type AstructBstruct2UseUnmarshaller struct{}

func (u *AstructBstruct2UseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AstructBstruct2Use), stage, identifier, instanceName, preserveOrder)
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

func (u *AstructBstructUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AstructBstructUse), stage, identifier, instanceName, preserveOrder)
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

func (u *BstructUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bstruct), stage, identifier, instanceName, preserveOrder)
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

func (u *DstructUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Dstruct), stage, identifier, instanceName, preserveOrder)
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

func (u *F0123456789012345678901234567890Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(F0123456789012345678901234567890), stage, identifier, instanceName, preserveOrder)
}

func (u *F0123456789012345678901234567890Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*F0123456789012345678901234567890)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	}
	return nil
}

type GstructUnmarshaller struct{}

func (u *GstructUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Gstruct), stage, identifier, instanceName, preserveOrder)
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

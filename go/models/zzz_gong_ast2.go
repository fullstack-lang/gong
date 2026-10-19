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
		case "github.com/fullstack-lang/gong/go/models":
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
				case "GongBasicField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongBasicField), stage, ident.Name, instanceName, preserveOrder)
				case "GongEnum":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongEnum), stage, ident.Name, instanceName, preserveOrder)
				case "GongEnumValue":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongEnumValue), stage, ident.Name, instanceName, preserveOrder)
				case "GongLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongLink), stage, ident.Name, instanceName, preserveOrder)
				case "GongNote":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongNote), stage, ident.Name, instanceName, preserveOrder)
				case "GongStruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongStruct), stage, ident.Name, instanceName, preserveOrder)
				case "GongTimeField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongTimeField), stage, ident.Name, instanceName, preserveOrder)
				case "MetaReference":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MetaReference), stage, ident.Name, instanceName, preserveOrder)
				case "ModelPkg":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ModelPkg), stage, ident.Name, instanceName, preserveOrder)
				case "PointerToGongStructField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointerToGongStructField), stage, ident.Name, instanceName, preserveOrder)
				case "SliceOfPointerToGongStructField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SliceOfPointerToGongStructField), stage, ident.Name, instanceName, preserveOrder)
				case "StageSetField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StageSetField), stage, ident.Name, instanceName, preserveOrder)
				case "StageSetModel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StageSetModel), stage, ident.Name, instanceName, preserveOrder)
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
				case *GongBasicField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "BasicKindName":
						inst.BasicKindName = GongExtractString(rhs)
					case "GongEnum":
						__gong__assignPointer(&inst.GongEnum, rhs, identifierMap)
					case "DeclaredType":
						inst.DeclaredType = GongExtractString(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "IsTextArea":
						inst.IsTextArea = GongExtractBool(rhs)
					case "IsBespokeWidth":
						inst.IsBespokeWidth = GongExtractBool(rhs)
					case "BespokeWidth":
						inst.BespokeWidth = GongExtractInt(rhs)
					case "IsBespokeHeight":
						inst.IsBespokeHeight = GongExtractBool(rhs)
					case "BespokeHeight":
						inst.BespokeHeight = GongExtractInt(rhs)
					}
				case *GongEnum:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						GongUnmarshallEnum(&inst.Type, rhs)
					case "GongEnumValues":
						__gong__assignSliceOfPointers(&inst.GongEnumValues, rhs, identifierMap)
					}
				case *GongEnumValue:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *GongLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Recv":
						inst.Recv = GongExtractString(rhs)
					case "ImportPath":
						inst.ImportPath = GongExtractString(rhs)
					}
				case *GongNote:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Body":
						inst.Body = GongExtractString(rhs)
					case "BodyHTML":
						inst.BodyHTML = GongExtractString(rhs)
					case "Links":
						__gong__assignSliceOfPointers(&inst.Links, rhs, identifierMap)
					}
				case *GongStruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongBasicFields":
						__gong__assignSliceOfPointers(&inst.GongBasicFields, rhs, identifierMap)
					case "GongTimeFields":
						__gong__assignSliceOfPointers(&inst.GongTimeFields, rhs, identifierMap)
					case "PointerToGongStructFields":
						__gong__assignSliceOfPointers(&inst.PointerToGongStructFields, rhs, identifierMap)
					case "SliceOfPointerToGongStructFields":
						__gong__assignSliceOfPointers(&inst.SliceOfPointerToGongStructFields, rhs, identifierMap)
					case "HasOnAfterUpdateSignature":
						inst.HasOnAfterUpdateSignature = GongExtractBool(rhs)
					case "IsIgnoredForFront":
						inst.IsIgnoredForFront = GongExtractBool(rhs)
					case "IsOmittedForMarshalling":
						inst.IsOmittedForMarshalling = GongExtractBool(rhs)
					case "ModelPkg":
						__gong__assignPointer(&inst.ModelPkg, rhs, identifierMap)
					}
				case *GongTimeField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					case "BespokeTimeFormat":
						inst.BespokeTimeFormat = GongExtractString(rhs)
					case "TimeFormOnly":
						inst.TimeFormOnly = GongExtractBool(rhs)
					}
				case *MetaReference:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ModelPkg:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PkgGoName":
						inst.PkgGoName = GongExtractString(rhs)
					case "PkgPath":
						inst.PkgPath = GongExtractString(rhs)
					case "StageSet":
						__gong__assignPointer(&inst.StageSet, rhs, identifierMap)
					case "PathToGoSubDirectory":
						inst.PathToGoSubDirectory = GongExtractString(rhs)
					case "OrmPkgGenPath":
						inst.OrmPkgGenPath = GongExtractString(rhs)
					case "DbOrmPkgGenPath":
						inst.DbOrmPkgGenPath = GongExtractString(rhs)
					case "DbLiteOrmPkgGenPath":
						inst.DbLiteOrmPkgGenPath = GongExtractString(rhs)
					case "DbPkgGenPath":
						inst.DbPkgGenPath = GongExtractString(rhs)
					case "ControllersPkgGenPath":
						inst.ControllersPkgGenPath = GongExtractString(rhs)
					case "FullstackPkgGenPath":
						inst.FullstackPkgGenPath = GongExtractString(rhs)
					case "StackPkgGenPath":
						inst.StackPkgGenPath = GongExtractString(rhs)
					case "Level1StackPkgGenPath":
						inst.Level1StackPkgGenPath = GongExtractString(rhs)
					case "StaticPkgGenPath":
						inst.StaticPkgGenPath = GongExtractString(rhs)
					case "ProbePkgGenPath":
						inst.ProbePkgGenPath = GongExtractString(rhs)
					case "NgWorkspacePath":
						inst.NgWorkspacePath = GongExtractString(rhs)
					case "NgWorkspaceName":
						inst.NgWorkspaceName = GongExtractString(rhs)
					case "NgDataLibrarySourceCodeDirectory":
						inst.NgDataLibrarySourceCodeDirectory = GongExtractString(rhs)
					case "NgSpecificLibrarySourceCodeDirectory":
						inst.NgSpecificLibrarySourceCodeDirectory = GongExtractString(rhs)
					case "MaterialLibDatamodelTargetPath":
						inst.MaterialLibDatamodelTargetPath = GongExtractString(rhs)
					}
				case *PointerToGongStructField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongStruct":
						__gong__assignPointer(&inst.GongStruct, rhs, identifierMap)
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					case "IsType":
						inst.IsType = GongExtractBool(rhs)
					}
				case *SliceOfPointerToGongStructField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongStruct":
						__gong__assignPointer(&inst.GongStruct, rhs, identifierMap)
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					}
				case *StageSetField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PackageName":
						inst.PackageName = GongExtractString(rhs)
					case "PackagePath":
						inst.PackagePath = GongExtractString(rhs)
					case "IsLocal":
						inst.IsLocal = GongExtractBool(rhs)
					case "ImportAlias":
						inst.ImportAlias = GongExtractString(rhs)
					}
				case *StageSetModel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Fields":
						__gong__assignSliceOfPointers(&inst.Fields, rhs, identifierMap)
					case "IsManual":
						inst.IsManual = GongExtractBool(rhs)
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
type GongBasicFieldUnmarshaller struct{}

func (u *GongBasicFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongBasicField), stage, identifier, instanceName, preserveOrder)
}

func (u *GongBasicFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongBasicField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BasicKindName":
		instance.BasicKindName = GongExtractString(valueExpr)
	case "GongEnum":
		GongUnmarshallPointer(&instance.GongEnum, valueExpr, identifierMap)
	case "DeclaredType":
		instance.DeclaredType = GongExtractString(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "IsTextArea":
		instance.IsTextArea = GongExtractBool(valueExpr)
	case "IsBespokeWidth":
		instance.IsBespokeWidth = GongExtractBool(valueExpr)
	case "BespokeWidth":
		instance.BespokeWidth = GongExtractInt(valueExpr)
	case "IsBespokeHeight":
		instance.IsBespokeHeight = GongExtractBool(valueExpr)
	case "BespokeHeight":
		instance.BespokeHeight = GongExtractInt(valueExpr)
	}
	return nil
}

type GongEnumUnmarshaller struct{}

func (u *GongEnumUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnum), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnum)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "GongEnumValues":
		GongUnmarshallSliceOfPointers(&instance.GongEnumValues, valueExpr, identifierMap)
	}
	return nil
}

type GongEnumValueUnmarshaller struct{}

func (u *GongEnumValueUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongEnumValue), stage, identifier, instanceName, preserveOrder)
}

func (u *GongEnumValueUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongEnumValue)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type GongLinkUnmarshaller struct{}

func (u *GongLinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongLink), stage, identifier, instanceName, preserveOrder)
}

func (u *GongLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Recv":
		instance.Recv = GongExtractString(valueExpr)
	case "ImportPath":
		instance.ImportPath = GongExtractString(valueExpr)
	}
	return nil
}

type GongNoteUnmarshaller struct{}

func (u *GongNoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongNote), stage, identifier, instanceName, preserveOrder)
}

func (u *GongNoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongNote)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Body":
		instance.Body = GongExtractString(valueExpr)
	case "BodyHTML":
		instance.BodyHTML = GongExtractString(valueExpr)
	case "Links":
		GongUnmarshallSliceOfPointers(&instance.Links, valueExpr, identifierMap)
	}
	return nil
}

type GongStructUnmarshaller struct{}

func (u *GongStructUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongStruct), stage, identifier, instanceName, preserveOrder)
}

func (u *GongStructUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongStruct)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GongBasicFields":
		GongUnmarshallSliceOfPointers(&instance.GongBasicFields, valueExpr, identifierMap)
	case "GongTimeFields":
		GongUnmarshallSliceOfPointers(&instance.GongTimeFields, valueExpr, identifierMap)
	case "PointerToGongStructFields":
		GongUnmarshallSliceOfPointers(&instance.PointerToGongStructFields, valueExpr, identifierMap)
	case "SliceOfPointerToGongStructFields":
		GongUnmarshallSliceOfPointers(&instance.SliceOfPointerToGongStructFields, valueExpr, identifierMap)
	case "HasOnAfterUpdateSignature":
		instance.HasOnAfterUpdateSignature = GongExtractBool(valueExpr)
	case "IsIgnoredForFront":
		instance.IsIgnoredForFront = GongExtractBool(valueExpr)
	case "IsOmittedForMarshalling":
		instance.IsOmittedForMarshalling = GongExtractBool(valueExpr)
	}
	return nil
}

type GongTimeFieldUnmarshaller struct{}

func (u *GongTimeFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GongTimeField), stage, identifier, instanceName, preserveOrder)
}

func (u *GongTimeFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GongTimeField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	case "BespokeTimeFormat":
		instance.BespokeTimeFormat = GongExtractString(valueExpr)
	case "TimeFormOnly":
		instance.TimeFormOnly = GongExtractBool(valueExpr)
	}
	return nil
}

type MetaReferenceUnmarshaller struct{}

func (u *MetaReferenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MetaReference), stage, identifier, instanceName, preserveOrder)
}

func (u *MetaReferenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MetaReference)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ModelPkgUnmarshaller struct{}

func (u *ModelPkgUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ModelPkg), stage, identifier, instanceName, preserveOrder)
}

func (u *ModelPkgUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ModelPkg)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PkgGoName":
		instance.PkgGoName = GongExtractString(valueExpr)
	case "PkgPath":
		instance.PkgPath = GongExtractString(valueExpr)
	case "PathToGoSubDirectory":
		instance.PathToGoSubDirectory = GongExtractString(valueExpr)
	case "OrmPkgGenPath":
		instance.OrmPkgGenPath = GongExtractString(valueExpr)
	case "DbOrmPkgGenPath":
		instance.DbOrmPkgGenPath = GongExtractString(valueExpr)
	case "DbLiteOrmPkgGenPath":
		instance.DbLiteOrmPkgGenPath = GongExtractString(valueExpr)
	case "DbPkgGenPath":
		instance.DbPkgGenPath = GongExtractString(valueExpr)
	case "ControllersPkgGenPath":
		instance.ControllersPkgGenPath = GongExtractString(valueExpr)
	case "FullstackPkgGenPath":
		instance.FullstackPkgGenPath = GongExtractString(valueExpr)
	case "StackPkgGenPath":
		instance.StackPkgGenPath = GongExtractString(valueExpr)
	case "Level1StackPkgGenPath":
		instance.Level1StackPkgGenPath = GongExtractString(valueExpr)
	case "StaticPkgGenPath":
		instance.StaticPkgGenPath = GongExtractString(valueExpr)
	case "ProbePkgGenPath":
		instance.ProbePkgGenPath = GongExtractString(valueExpr)
	case "NgWorkspacePath":
		instance.NgWorkspacePath = GongExtractString(valueExpr)
	case "NgWorkspaceName":
		instance.NgWorkspaceName = GongExtractString(valueExpr)
	case "NgDataLibrarySourceCodeDirectory":
		instance.NgDataLibrarySourceCodeDirectory = GongExtractString(valueExpr)
	case "NgSpecificLibrarySourceCodeDirectory":
		instance.NgSpecificLibrarySourceCodeDirectory = GongExtractString(valueExpr)
	case "MaterialLibDatamodelTargetPath":
		instance.MaterialLibDatamodelTargetPath = GongExtractString(valueExpr)
	}
	return nil
}

type PointerToGongStructFieldUnmarshaller struct{}

func (u *PointerToGongStructFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointerToGongStructField), stage, identifier, instanceName, preserveOrder)
}

func (u *PointerToGongStructFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointerToGongStructField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GongStruct":
		GongUnmarshallPointer(&instance.GongStruct, valueExpr, identifierMap)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	case "IsType":
		instance.IsType = GongExtractBool(valueExpr)
	}
	return nil
}

type SliceOfPointerToGongStructFieldUnmarshaller struct{}

func (u *SliceOfPointerToGongStructFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SliceOfPointerToGongStructField), stage, identifier, instanceName, preserveOrder)
}

func (u *SliceOfPointerToGongStructFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SliceOfPointerToGongStructField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GongStruct":
		GongUnmarshallPointer(&instance.GongStruct, valueExpr, identifierMap)
	case "Index":
		instance.Index = GongExtractInt(valueExpr)
	case "CompositeStructName":
		instance.CompositeStructName = GongExtractString(valueExpr)
	case "IsAccordionStart":
		instance.IsAccordionStart = GongExtractBool(valueExpr)
	case "AccordionName":
		instance.AccordionName = GongExtractString(valueExpr)
	case "IsAccordionEnd":
		instance.IsAccordionEnd = GongExtractBool(valueExpr)
	}
	return nil
}

type StageSetFieldUnmarshaller struct{}

func (u *StageSetFieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StageSetField), stage, identifier, instanceName, preserveOrder)
}

func (u *StageSetFieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StageSetField)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PackageName":
		instance.PackageName = GongExtractString(valueExpr)
	case "PackagePath":
		instance.PackagePath = GongExtractString(valueExpr)
	case "IsLocal":
		instance.IsLocal = GongExtractBool(valueExpr)
	case "ImportAlias":
		instance.ImportAlias = GongExtractString(valueExpr)
	}
	return nil
}

type StageSetModelUnmarshaller struct{}

func (u *StageSetModelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StageSetModel), stage, identifier, instanceName, preserveOrder)
}

func (u *StageSetModelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StageSetModel)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Fields":
		GongUnmarshallSliceOfPointers(&instance.Fields, valueExpr, identifierMap)
	case "IsManual":
		instance.IsManual = GongExtractBool(valueExpr)
	}
	return nil
}

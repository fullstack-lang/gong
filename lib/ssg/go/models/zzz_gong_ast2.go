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
		case "github.com/fullstack-lang/gong/lib/ssg/go/models":
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
				case "Chapter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Chapter), stage, ident.Name, instanceName, preserveOrder)
				case "Content":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Content), stage, ident.Name, instanceName, preserveOrder)
				case "DownloadableFile":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DownloadableFile), stage, ident.Name, instanceName, preserveOrder)
				case "JpgImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(JpgImage), stage, ident.Name, instanceName, preserveOrder)
				case "Page":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Page), stage, ident.Name, instanceName, preserveOrder)
				case "PngImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PngImage), stage, ident.Name, instanceName, preserveOrder)
				case "Section":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Section), stage, ident.Name, instanceName, preserveOrder)
				case "SvgImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SvgImage), stage, ident.Name, instanceName, preserveOrder)
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
				case *Chapter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "Sections":
						__gong__assignSliceOfPointers(&inst.Sections, rhs, identifierMap)
					case "Pages":
						__gong__assignSliceOfPointers(&inst.Pages, rhs, identifierMap)
					case "SubChapters":
						__gong__assignSliceOfPointers(&inst.SubChapters, rhs, identifierMap)
					}
				case *Content:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "ContentPath":
						inst.ContentPath = GongExtractString(rhs)
					case "OutputPath":
						inst.OutputPath = GongExtractString(rhs)
					case "StaticPath":
						inst.StaticPath = GongExtractString(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "IsBespokeLogoFileName":
						inst.IsBespokeLogoFileName = GongExtractBool(rhs)
					case "BespokeLogoFileName":
						inst.BespokeLogoFileName = GongExtractString(rhs)
					case "IsBespokePageTileLogoFileName":
						inst.IsBespokePageTileLogoFileName = GongExtractBool(rhs)
					case "BespokePageTileLogoFileName":
						inst.BespokePageTileLogoFileName = GongExtractString(rhs)
					case "Target":
						GongUnmarshallEnum(&inst.Target, rhs)
					case "Chapters":
						__gong__assignSliceOfPointers(&inst.Chapters, rhs, identifierMap)
					case "VersionInfo":
						inst.VersionInfo = GongExtractString(rhs)
					}
				case *DownloadableFile:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *JpgImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *Page:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "Sections":
						__gong__assignSliceOfPointers(&inst.Sections, rhs, identifierMap)
					}
				case *PngImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *Section:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MardownContent":
						inst.MardownContent = GongExtractString(rhs)
					case "IsImage":
						inst.IsImage = GongExtractBool(rhs)
					case "SvgImage":
						__gong__assignPointer(&inst.SvgImage, rhs, identifierMap)
					case "PngImage":
						__gong__assignPointer(&inst.PngImage, rhs, identifierMap)
					case "JpgImage":
						__gong__assignPointer(&inst.JpgImage, rhs, identifierMap)
					case "IsDownloadableFile":
						inst.IsDownloadableFile = GongExtractBool(rhs)
					case "DownloadableFile":
						__gong__assignPointer(&inst.DownloadableFile, rhs, identifierMap)
					}
				case *SvgImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
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
type ChapterUnmarshaller struct{}

func (u *ChapterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Chapter), stage, identifier, instanceName, preserveOrder)
}

func (u *ChapterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Chapter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MardownContent":
		instance.MardownContent = GongExtractString(valueExpr)
	case "Sections":
		GongUnmarshallSliceOfPointers(&instance.Sections, valueExpr, identifierMap)
	case "Pages":
		GongUnmarshallSliceOfPointers(&instance.Pages, valueExpr, identifierMap)
	case "SubChapters":
		GongUnmarshallSliceOfPointers(&instance.SubChapters, valueExpr, identifierMap)
	}
	return nil
}

type ContentUnmarshaller struct{}

func (u *ContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Content), stage, identifier, instanceName, preserveOrder)
}

func (u *ContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Content)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MardownContent":
		instance.MardownContent = GongExtractString(valueExpr)
	case "ContentPath":
		instance.ContentPath = GongExtractString(valueExpr)
	case "OutputPath":
		instance.OutputPath = GongExtractString(valueExpr)
	case "StaticPath":
		instance.StaticPath = GongExtractString(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "IsBespokeLogoFileName":
		instance.IsBespokeLogoFileName = GongExtractBool(valueExpr)
	case "BespokeLogoFileName":
		instance.BespokeLogoFileName = GongExtractString(valueExpr)
	case "IsBespokePageTileLogoFileName":
		instance.IsBespokePageTileLogoFileName = GongExtractBool(valueExpr)
	case "BespokePageTileLogoFileName":
		instance.BespokePageTileLogoFileName = GongExtractString(valueExpr)
	case "Target":
		GongUnmarshallEnum(&instance.Target, valueExpr)
	case "Chapters":
		GongUnmarshallSliceOfPointers(&instance.Chapters, valueExpr, identifierMap)
	case "VersionInfo":
		instance.VersionInfo = GongExtractString(valueExpr)
	}
	return nil
}

type DownloadableFileUnmarshaller struct{}

func (u *DownloadableFileUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DownloadableFile), stage, identifier, instanceName, preserveOrder)
}

func (u *DownloadableFileUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DownloadableFile)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type JpgImageUnmarshaller struct{}

func (u *JpgImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(JpgImage), stage, identifier, instanceName, preserveOrder)
}

func (u *JpgImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*JpgImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type PageUnmarshaller struct{}

func (u *PageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Page), stage, identifier, instanceName, preserveOrder)
}

func (u *PageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Page)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MardownContent":
		instance.MardownContent = GongExtractString(valueExpr)
	case "Sections":
		GongUnmarshallSliceOfPointers(&instance.Sections, valueExpr, identifierMap)
	}
	return nil
}

type PngImageUnmarshaller struct{}

func (u *PngImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PngImage), stage, identifier, instanceName, preserveOrder)
}

func (u *PngImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PngImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type SectionUnmarshaller struct{}

func (u *SectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Section), stage, identifier, instanceName, preserveOrder)
}

func (u *SectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Section)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MardownContent":
		instance.MardownContent = GongExtractString(valueExpr)
	case "IsImage":
		instance.IsImage = GongExtractBool(valueExpr)
	case "SvgImage":
		GongUnmarshallPointer(&instance.SvgImage, valueExpr, identifierMap)
	case "PngImage":
		GongUnmarshallPointer(&instance.PngImage, valueExpr, identifierMap)
	case "JpgImage":
		GongUnmarshallPointer(&instance.JpgImage, valueExpr, identifierMap)
	case "IsDownloadableFile":
		instance.IsDownloadableFile = GongExtractBool(valueExpr)
	case "DownloadableFile":
		GongUnmarshallPointer(&instance.DownloadableFile, valueExpr, identifierMap)
	}
	return nil
}

type SvgImageUnmarshaller struct{}

func (u *SvgImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SvgImage), stage, identifier, instanceName, preserveOrder)
}

func (u *SvgImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SvgImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	}
	return nil
}

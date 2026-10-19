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
		case "github.com/fullstack-lang/gong/dsm/floss/go/models":
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
				case "CompareAnalysis":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CompareAnalysis), stage, ident.Name, instanceName, preserveOrder)
				case "Complexity":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Complexity), stage, ident.Name, instanceName, preserveOrder)
				case "DiagramFlossEquation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramFlossEquation), stage, ident.Name, instanceName, preserveOrder)
				case "Effort":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Effort), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "NoteComplexityShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteComplexityShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteEffortShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteEffortShape), stage, ident.Name, instanceName, preserveOrder)
				case "NotePerformanceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePerformanceShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "Performance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Performance), stage, ident.Name, instanceName, preserveOrder)
				case "System":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System), stage, ident.Name, instanceName, preserveOrder)
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
				case *CompareAnalysis:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FromSystem":
						__gong__assignPointer(&inst.FromSystem, rhs, identifierMap)
					case "ToSystem":
						__gong__assignPointer(&inst.ToSystem, rhs, identifierMap)
					case "Mu":
						inst.Mu = GongExtractFloat(rhs)
					case "Epsilon":
						inst.Epsilon = GongExtractFloat(rhs)
					case "DiagramFlossEquations":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquations, rhs, identifierMap)
					case "DiagramFlossEquationsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquationsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Complexity:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Strength":
						inst.Strength = GongExtractFloat(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *DiagramFlossEquation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Scale":
						inst.Scale = GongExtractFloat(rhs)
					case "FontSize":
						GongUnmarshallEnum(&inst.FontSize, rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsEditable_":
						inst.IsEditable_ = GongExtractBool(rhs)
					case "IsInDelta3ColumnsMode":
						inst.IsInDelta3ColumnsMode = GongExtractBool(rhs)
					case "AreQuantitativeElementsVisible":
						inst.AreQuantitativeElementsVisible = GongExtractBool(rhs)
					case "AreSubsystemsVisible":
						inst.AreSubsystemsVisible = GongExtractBool(rhs)
					case "AreCommonElementsHidden":
						inst.AreCommonElementsHidden = GongExtractBool(rhs)
					case "AreCPEArrowsVisible":
						inst.AreCPEArrowsVisible = GongExtractBool(rhs)
					case "AreColumnTitlesVisible":
						inst.AreColumnTitlesVisible = GongExtractBool(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "DefaultBoxWidth":
						inst.DefaultBoxWidth = GongExtractFloat(rhs)
					case "DefaultBoxHeigth":
						inst.DefaultBoxHeigth = GongExtractFloat(rhs)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NoteComplexityShapes":
						__gong__assignSliceOfPointers(&inst.NoteComplexityShapes, rhs, identifierMap)
					case "NotePerformanceShapes":
						__gong__assignSliceOfPointers(&inst.NotePerformanceShapes, rhs, identifierMap)
					case "NoteEffortShapes":
						__gong__assignSliceOfPointers(&inst.NoteEffortShapes, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ComplexitysWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PerformancesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EffortsWhoseNodeIsExpanded, rhs, identifierMap)
					}
				case *Effort:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Strength":
						inst.Strength = GongExtractFloat(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "RootSystems":
						__gong__assignSliceOfPointers(&inst.RootSystems, rhs, identifierMap)
					case "RootComplexitys":
						__gong__assignSliceOfPointers(&inst.RootComplexitys, rhs, identifierMap)
					case "RootPerformances":
						__gong__assignSliceOfPointers(&inst.RootPerformances, rhs, identifierMap)
					case "RootEfforts":
						__gong__assignSliceOfPointers(&inst.RootEfforts, rhs, identifierMap)
					case "RootCompareAnalysis":
						__gong__assignSliceOfPointers(&inst.RootCompareAnalysis, rhs, identifierMap)
					case "RootNotes":
						__gong__assignSliceOfPointers(&inst.RootNotes, rhs, identifierMap)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SubLibrariesWhoseNodeIsExpanded, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "IsSystemsNodeExpanded":
						inst.IsSystemsNodeExpanded = GongExtractBool(rhs)
					case "SystemsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SystemsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ComplexitysWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PerformancesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EffortsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsCompareAnalysisNodeExpanded":
						inst.IsCompareAnalysisNodeExpanded = GongExtractBool(rhs)
					case "CompareAnalysisWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.CompareAnalysisWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsExpandedTmp":
						inst.IsExpandedTmp = GongExtractBool(rhs)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Complexities":
						__gong__assignSliceOfPointers(&inst.Complexities, rhs, identifierMap)
					case "Performances":
						__gong__assignSliceOfPointers(&inst.Performances, rhs, identifierMap)
					case "Efforts":
						__gong__assignSliceOfPointers(&inst.Efforts, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					}
				case *NoteComplexityShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Complexity":
						__gong__assignPointer(&inst.Complexity, rhs, identifierMap)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						GongUnmarshallEnum(&inst.StartOrientation, rhs)
					case "EndOrientation":
						GongUnmarshallEnum(&inst.EndOrientation, rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "IsPresentationOverridden":
						inst.IsPresentationOverridden = GongExtractBool(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontColor":
						inst.FontColor = GongExtractString(rhs)
					}
				case *NoteEffortShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Effort":
						__gong__assignPointer(&inst.Effort, rhs, identifierMap)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						GongUnmarshallEnum(&inst.StartOrientation, rhs)
					case "EndOrientation":
						GongUnmarshallEnum(&inst.EndOrientation, rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "IsPresentationOverridden":
						inst.IsPresentationOverridden = GongExtractBool(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontColor":
						inst.FontColor = GongExtractString(rhs)
					}
				case *NotePerformanceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Performance":
						__gong__assignPointer(&inst.Performance, rhs, identifierMap)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						GongUnmarshallEnum(&inst.StartOrientation, rhs)
					case "EndOrientation":
						GongUnmarshallEnum(&inst.EndOrientation, rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "IsPresentationOverridden":
						inst.IsPresentationOverridden = GongExtractBool(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontColor":
						inst.FontColor = GongExtractString(rhs)
					}
				case *NoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "IsPresentationOverridden":
						inst.IsPresentationOverridden = GongExtractBool(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontColor":
						inst.FontColor = GongExtractString(rhs)
					}
				case *Performance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Strength":
						inst.Strength = GongExtractFloat(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *System:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Complexities":
						__gong__assignSliceOfPointers(&inst.Complexities, rhs, identifierMap)
					case "Performances":
						__gong__assignSliceOfPointers(&inst.Performances, rhs, identifierMap)
					case "Efforts":
						__gong__assignSliceOfPointers(&inst.Efforts, rhs, identifierMap)
					case "SubSystems":
						__gong__assignSliceOfPointers(&inst.SubSystems, rhs, identifierMap)
					case "AreCPEsCompoundedFromSubSystems":
						inst.AreCPEsCompoundedFromSubSystems = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SVG_Path":
						inst.SVG_Path = GongExtractString(rhs)
					case "InverseAppliedScaling":
						inst.InverseAppliedScaling = GongExtractFloat(rhs)
					case "DiagramFlossEquations":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquations, rhs, identifierMap)
					case "DiagramFlossEquationsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquationsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsSubSystemNodeExpanded":
						inst.IsSubSystemNodeExpanded = GongExtractBool(rhs)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ComplexitysWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PerformancesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EffortsWhoseNodeIsExpanded, rhs, identifierMap)
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
type CompareAnalysisUnmarshaller struct{}

func (u *CompareAnalysisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CompareAnalysis), stage, identifier, instanceName, preserveOrder)
}

func (u *CompareAnalysisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CompareAnalysis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "FromSystem":
		GongUnmarshallPointer(&instance.FromSystem, valueExpr, identifierMap)
	case "ToSystem":
		GongUnmarshallPointer(&instance.ToSystem, valueExpr, identifierMap)
	case "Mu":
		instance.Mu = GongExtractFloat(valueExpr)
	case "Epsilon":
		instance.Epsilon = GongExtractFloat(valueExpr)
	case "DiagramFlossEquations":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquations, valueExpr, identifierMap)
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquationsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ComplexityUnmarshaller struct{}

func (u *ComplexityUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Complexity), stage, identifier, instanceName, preserveOrder)
}

func (u *ComplexityUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Complexity)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Strength":
		instance.Strength = GongExtractFloat(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type DiagramFlossEquationUnmarshaller struct{}

func (u *DiagramFlossEquationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramFlossEquation), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramFlossEquationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramFlossEquation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Scale":
		instance.Scale = GongExtractFloat(valueExpr)
	case "FontSize":
		GongUnmarshallEnum(&instance.FontSize, valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "IsEditable_":
		instance.IsEditable_ = GongExtractBool(valueExpr)
	case "IsInDelta3ColumnsMode":
		instance.IsInDelta3ColumnsMode = GongExtractBool(valueExpr)
	case "AreQuantitativeElementsVisible":
		instance.AreQuantitativeElementsVisible = GongExtractBool(valueExpr)
	case "AreSubsystemsVisible":
		instance.AreSubsystemsVisible = GongExtractBool(valueExpr)
	case "AreCommonElementsHidden":
		instance.AreCommonElementsHidden = GongExtractBool(valueExpr)
	case "AreCPEArrowsVisible":
		instance.AreCPEArrowsVisible = GongExtractBool(valueExpr)
	case "AreColumnTitlesVisible":
		instance.AreColumnTitlesVisible = GongExtractBool(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "DefaultBoxWidth":
		instance.DefaultBoxWidth = GongExtractFloat(valueExpr)
	case "DefaultBoxHeigth":
		instance.DefaultBoxHeigth = GongExtractFloat(valueExpr)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NoteComplexityShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteComplexityShapes, valueExpr, identifierMap)
	case "NotePerformanceShapes":
		GongUnmarshallSliceOfPointers(&instance.NotePerformanceShapes, valueExpr, identifierMap)
	case "NoteEffortShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteEffortShapes, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "ComplexitysWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ComplexitysWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "PerformancesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PerformancesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	case "EffortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EffortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type EffortUnmarshaller struct{}

func (u *EffortUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Effort), stage, identifier, instanceName, preserveOrder)
}

func (u *EffortUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Effort)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Strength":
		instance.Strength = GongExtractFloat(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type LibraryUnmarshaller struct{}

func (u *LibraryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Library), stage, identifier, instanceName, preserveOrder)
}

func (u *LibraryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Library)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "RootSystems":
		GongUnmarshallSliceOfPointers(&instance.RootSystems, valueExpr, identifierMap)
	case "RootComplexitys":
		GongUnmarshallSliceOfPointers(&instance.RootComplexitys, valueExpr, identifierMap)
	case "RootPerformances":
		GongUnmarshallSliceOfPointers(&instance.RootPerformances, valueExpr, identifierMap)
	case "RootEfforts":
		GongUnmarshallSliceOfPointers(&instance.RootEfforts, valueExpr, identifierMap)
	case "RootCompareAnalysis":
		GongUnmarshallSliceOfPointers(&instance.RootCompareAnalysis, valueExpr, identifierMap)
	case "RootNotes":
		GongUnmarshallSliceOfPointers(&instance.RootNotes, valueExpr, identifierMap)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "IsSubLibrariesNodeExpanded":
		instance.IsSubLibrariesNodeExpanded = GongExtractBool(valueExpr)
	case "SubLibrariesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SubLibrariesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "IsSystemsNodeExpanded":
		instance.IsSystemsNodeExpanded = GongExtractBool(valueExpr)
	case "SystemsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SystemsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "ComplexitysWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ComplexitysWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "PerformancesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PerformancesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	case "EffortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EffortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsCompareAnalysisNodeExpanded":
		instance.IsCompareAnalysisNodeExpanded = GongExtractBool(valueExpr)
	case "CompareAnalysisWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.CompareAnalysisWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsExpandedTmp":
		instance.IsExpandedTmp = GongExtractBool(valueExpr)
	}
	return nil
}

type NoteUnmarshaller struct{}

func (u *NoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Note), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Complexities":
		GongUnmarshallSliceOfPointers(&instance.Complexities, valueExpr, identifierMap)
	case "Performances":
		GongUnmarshallSliceOfPointers(&instance.Performances, valueExpr, identifierMap)
	case "Efforts":
		GongUnmarshallSliceOfPointers(&instance.Efforts, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type NoteComplexityShapeUnmarshaller struct{}

func (u *NoteComplexityShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteComplexityShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteComplexityShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteComplexityShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Complexity":
		GongUnmarshallPointer(&instance.Complexity, valueExpr, identifierMap)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "IsPresentationOverridden":
		instance.IsPresentationOverridden = GongExtractBool(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontColor":
		instance.FontColor = GongExtractString(valueExpr)
	}
	return nil
}

type NoteEffortShapeUnmarshaller struct{}

func (u *NoteEffortShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteEffortShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteEffortShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteEffortShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Effort":
		GongUnmarshallPointer(&instance.Effort, valueExpr, identifierMap)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "IsPresentationOverridden":
		instance.IsPresentationOverridden = GongExtractBool(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontColor":
		instance.FontColor = GongExtractString(valueExpr)
	}
	return nil
}

type NotePerformanceShapeUnmarshaller struct{}

func (u *NotePerformanceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePerformanceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePerformanceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePerformanceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Performance":
		GongUnmarshallPointer(&instance.Performance, valueExpr, identifierMap)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "IsPresentationOverridden":
		instance.IsPresentationOverridden = GongExtractBool(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontColor":
		instance.FontColor = GongExtractString(valueExpr)
	}
	return nil
}

type NoteShapeUnmarshaller struct{}

func (u *NoteShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
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
	case "IsPresentationOverridden":
		instance.IsPresentationOverridden = GongExtractBool(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontColor":
		instance.FontColor = GongExtractString(valueExpr)
	}
	return nil
}

type PerformanceUnmarshaller struct{}

func (u *PerformanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Performance), stage, identifier, instanceName, preserveOrder)
}

func (u *PerformanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Performance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Strength":
		instance.Strength = GongExtractFloat(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type SystemUnmarshaller struct{}

func (u *SystemUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(System), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Complexities":
		GongUnmarshallSliceOfPointers(&instance.Complexities, valueExpr, identifierMap)
	case "Performances":
		GongUnmarshallSliceOfPointers(&instance.Performances, valueExpr, identifierMap)
	case "Efforts":
		GongUnmarshallSliceOfPointers(&instance.Efforts, valueExpr, identifierMap)
	case "SubSystems":
		GongUnmarshallSliceOfPointers(&instance.SubSystems, valueExpr, identifierMap)
	case "AreCPEsCompoundedFromSubSystems":
		instance.AreCPEsCompoundedFromSubSystems = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	case "DiagramFlossEquations":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquations, valueExpr, identifierMap)
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramFlossEquationsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubSystemNodeExpanded":
		instance.IsSubSystemNodeExpanded = GongExtractBool(valueExpr)
	case "IsComplexitysNodeExpanded":
		instance.IsComplexitysNodeExpanded = GongExtractBool(valueExpr)
	case "ComplexitysWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ComplexitysWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPerformancesNodeExpanded":
		instance.IsPerformancesNodeExpanded = GongExtractBool(valueExpr)
	case "PerformancesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PerformancesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEffortsNodeExpanded":
		instance.IsEffortsNodeExpanded = GongExtractBool(valueExpr)
	case "EffortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EffortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

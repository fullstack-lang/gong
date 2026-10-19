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
		case "github.com/fullstack-lang/gong/dsm/structure/go/models":
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
				case "AllocatedResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AllocatedResourceShape), stage, ident.Name, instanceName, preserveOrder)
				case "AllocatedSystemShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AllocatedSystemShape), stage, ident.Name, instanceName, preserveOrder)
				case "ControlFlow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlFlow), stage, ident.Name, instanceName, preserveOrder)
				case "ControlFlowShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlFlowShape), stage, ident.Name, instanceName, preserveOrder)
				case "Data":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Data), stage, ident.Name, instanceName, preserveOrder)
				case "DataFlow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DataFlow), stage, ident.Name, instanceName, preserveOrder)
				case "DataFlowShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DataFlowShape), stage, ident.Name, instanceName, preserveOrder)
				case "DataShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DataShape), stage, ident.Name, instanceName, preserveOrder)
				case "DiagramLayerState":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramLayerState), stage, ident.Name, instanceName, preserveOrder)
				case "DiagramStructure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramStructure), stage, ident.Name, instanceName, preserveOrder)
				case "ExternalPartShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ExternalPartShape), stage, ident.Name, instanceName, preserveOrder)
				case "LayerDefinition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LayerDefinition), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "NotePartShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePartShape), stage, ident.Name, instanceName, preserveOrder)
				case "NotePortShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePortShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "Part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part), stage, ident.Name, instanceName, preserveOrder)
				case "PartAnchoredPath":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartAnchoredPath), stage, ident.Name, instanceName, preserveOrder)
				case "PartShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartShape), stage, ident.Name, instanceName, preserveOrder)
				case "Port":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Port), stage, ident.Name, instanceName, preserveOrder)
				case "PortShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PortShape), stage, ident.Name, instanceName, preserveOrder)
				case "Resource":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Resource), stage, ident.Name, instanceName, preserveOrder)
				case "SemanticTag":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SemanticTag), stage, ident.Name, instanceName, preserveOrder)
				case "System":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System), stage, ident.Name, instanceName, preserveOrder)
				case "SystemShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SystemShape), stage, ident.Name, instanceName, preserveOrder)
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
				case *AllocatedResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						__gong__assignPointer(&inst.Part, rhs, identifierMap)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
					}
				case *AllocatedSystemShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						__gong__assignPointer(&inst.Part, rhs, identifierMap)
					case "System":
						__gong__assignPointer(&inst.System, rhs, identifierMap)
					}
				case *ControlFlow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					}
				case *ControlFlowShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ControlFlow":
						__gong__assignPointer(&inst.ControlFlow, rhs, identifierMap)
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
				case *Data:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Acronym":
						inst.Acronym = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SVG_Path":
						inst.SVG_Path = GongExtractString(rhs)
					case "InverseAppliedScaling":
						inst.InverseAppliedScaling = GongExtractFloat(rhs)
					}
				case *DataFlow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StartPort":
						__gong__assignPointer(&inst.StartPort, rhs, identifierMap)
					case "EndPort":
						__gong__assignPointer(&inst.EndPort, rhs, identifierMap)
					case "StartExternalPart":
						__gong__assignPointer(&inst.StartExternalPart, rhs, identifierMap)
					case "EndExternalPart":
						__gong__assignPointer(&inst.EndExternalPart, rhs, identifierMap)
					case "Datas":
						__gong__assignSliceOfPointers(&inst.Datas, rhs, identifierMap)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
					case "IsDatasNodeExpanded":
						inst.IsDatasNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Type":
						GongUnmarshallEnum(&inst.Type, rhs)
					}
				case *DataFlowShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DataFlow":
						__gong__assignPointer(&inst.DataFlow, rhs, identifierMap)
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
				case *DataShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Data":
						__gong__assignPointer(&inst.Data, rhs, identifierMap)
					case "DataFlow":
						__gong__assignPointer(&inst.DataFlow, rhs, identifierMap)
					}
				case *DiagramLayerState:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DiagramStructure":
						__gong__assignPointer(&inst.DiagramStructure, rhs, identifierMap)
					case "LayerDefinition":
						__gong__assignPointer(&inst.LayerDefinition, rhs, identifierMap)
					}
				case *DiagramStructure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsEditable_":
						inst.IsEditable_ = GongExtractBool(rhs)
					case "IsShowPrefix":
						inst.IsShowPrefix = GongExtractBool(rhs)
					case "DefaultBoxWidth":
						inst.DefaultBoxWidth = GongExtractFloat(rhs)
					case "DefaultBoxHeigth":
						inst.DefaultBoxHeigth = GongExtractFloat(rhs)
					case "IsWithDiscretePorts":
						inst.IsWithDiscretePorts = GongExtractBool(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "System_Shapes":
						__gong__assignSliceOfPointers(&inst.System_Shapes, rhs, identifierMap)
					case "IsSystemsNodeExpanded":
						inst.IsSystemsNodeExpanded = GongExtractBool(rhs)
					case "SystemsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SystemsWhoseNodeIsExpanded, rhs, identifierMap)
					case "Part_Shapes":
						__gong__assignSliceOfPointers(&inst.Part_Shapes, rhs, identifierMap)
					case "IsPartsNodeExpanded":
						inst.IsPartsNodeExpanded = GongExtractBool(rhs)
					case "PartWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PartWhoseNodeIsExpanded, rhs, identifierMap)
					case "ExternalPart_Shapes":
						__gong__assignSliceOfPointers(&inst.ExternalPart_Shapes, rhs, identifierMap)
					case "IsExternalPartsNodeExpanded":
						inst.IsExternalPartsNodeExpanded = GongExtractBool(rhs)
					case "ExternalPartWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalPartWhoseNodeIsExpanded, rhs, identifierMap)
					case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalPartsWhoseInDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "PortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PortsWhoseNodeIsExpanded, rhs, identifierMap)
					case "Port_Shapes":
						__gong__assignSliceOfPointers(&inst.Port_Shapes, rhs, identifierMap)
					case "ControlFlowsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ControlFlowsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ControlFlow_Shapes":
						__gong__assignSliceOfPointers(&inst.ControlFlow_Shapes, rhs, identifierMap)
					case "DataFlowsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DataFlowsWhoseNodeIsExpanded, rhs, identifierMap)
					case "DataFlow_Shapes":
						__gong__assignSliceOfPointers(&inst.DataFlow_Shapes, rhs, identifierMap)
					case "DatasWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DatasWhoseNodeIsExpanded, rhs, identifierMap)
					case "Data_Shapes":
						__gong__assignSliceOfPointers(&inst.Data_Shapes, rhs, identifierMap)
					case "DataFlowsWhoseDataNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DataFlowsWhoseDataNodeIsExpanded, rhs, identifierMap)
					case "AllocatedResourcesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.AllocatedResourcesWhoseNodeIsExpanded, rhs, identifierMap)
					case "AllocatedResourceShapes":
						__gong__assignSliceOfPointers(&inst.AllocatedResourceShapes, rhs, identifierMap)
					case "AllocatedSystemesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.AllocatedSystemesWhoseNodeIsExpanded, rhs, identifierMap)
					case "AllocatedSystemShapes":
						__gong__assignSliceOfPointers(&inst.AllocatedSystemShapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NotePortShapes":
						__gong__assignSliceOfPointers(&inst.NotePortShapes, rhs, identifierMap)
					case "NotePartShapes":
						__gong__assignSliceOfPointers(&inst.NotePartShapes, rhs, identifierMap)
					}
				case *ExternalPartShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						__gong__assignPointer(&inst.Part, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
					case "TailHeigth":
						inst.TailHeigth = GongExtractFloat(rhs)
					}
				case *LayerDefinition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Query":
						__gong__assignSliceOfPointers(&inst.Query, rhs, identifierMap)
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
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SubLibrariesWhoseNodeIsExpanded, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "RootSystemes":
						__gong__assignSliceOfPointers(&inst.RootSystemes, rhs, identifierMap)
					case "IsSystemesNodeExpanded":
						inst.IsSystemesNodeExpanded = GongExtractBool(rhs)
					case "SystemsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SystemsWhoseNodeIsExpanded, rhs, identifierMap)
					case "RootDataFlows":
						__gong__assignSliceOfPointers(&inst.RootDataFlows, rhs, identifierMap)
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "DataFlowsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DataFlowsWhoseNodeIsExpanded, rhs, identifierMap)
					case "RootDatas":
						__gong__assignSliceOfPointers(&inst.RootDatas, rhs, identifierMap)
					case "IsDatasNodeExpanded":
						inst.IsDatasNodeExpanded = GongExtractBool(rhs)
					case "DatasWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DatasWhoseNodeIsExpanded, rhs, identifierMap)
					case "RootResources":
						__gong__assignSliceOfPointers(&inst.RootResources, rhs, identifierMap)
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "ResourcesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ResourcesWhoseNodeIsExpanded, rhs, identifierMap)
					case "PartsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PartsWhoseNodeIsExpanded, rhs, identifierMap)
					case "RootNotes":
						__gong__assignSliceOfPointers(&inst.RootNotes, rhs, identifierMap)
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
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsPartsNodeExpanded":
						inst.IsPartsNodeExpanded = GongExtractBool(rhs)
					case "Parts":
						__gong__assignSliceOfPointers(&inst.Parts, rhs, identifierMap)
					case "IsPortsNodeExpanded":
						inst.IsPortsNodeExpanded = GongExtractBool(rhs)
					case "Ports":
						__gong__assignSliceOfPointers(&inst.Ports, rhs, identifierMap)
					}
				case *NotePartShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Part":
						__gong__assignPointer(&inst.Part, rhs, identifierMap)
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
				case *NotePortShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Port":
						__gong__assignPointer(&inst.Port, rhs, identifierMap)
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
				case *Part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Ports":
						__gong__assignSliceOfPointers(&inst.Ports, rhs, identifierMap)
					case "TypeOfPart":
						__gong__assignPointer(&inst.TypeOfPart, rhs, identifierMap)
					case "IsPartNameNotSystemName":
						inst.IsPartNameNotSystemName = GongExtractBool(rhs)
					case "IsControlFlowsNodeExpanded":
						inst.IsControlFlowsNodeExpanded = GongExtractBool(rhs)
					case "ControlFlows":
						__gong__assignSliceOfPointers(&inst.ControlFlows, rhs, identifierMap)
					case "PortWhoseOutControlFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PortWhoseOutControlFlowsNodeIsExpanded, rhs, identifierMap)
					case "PortWhoseInControlFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PortWhoseInControlFlowsNodeIsExpanded, rhs, identifierMap)
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "PortWhoseOutDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PortWhoseOutDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "PortWhoseInDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PortWhoseInDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "PartAnchoredPath":
						__gong__assignSliceOfPointers(&inst.PartAnchoredPath, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsPortsNodeExpanded":
						inst.IsPortsNodeExpanded = GongExtractBool(rhs)
					}
				case *PartAnchoredPath:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						GongUnmarshallEnum(&inst.RectAnchorType, rhs)
					case "ScalePropotionnally":
						inst.ScalePropotionnally = GongExtractBool(rhs)
					case "AppliedScaling":
						inst.AppliedScaling = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *PartShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						__gong__assignPointer(&inst.Part, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
				case *Port:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *PortShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Port":
						__gong__assignPointer(&inst.Port, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
				case *Resource:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Acronym":
						inst.Acronym = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SVG_Path":
						inst.SVG_Path = GongExtractString(rhs)
					case "InverseAppliedScaling":
						inst.InverseAppliedScaling = GongExtractFloat(rhs)
					}
				case *SemanticTag:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parts":
						__gong__assignSliceOfPointers(&inst.Parts, rhs, identifierMap)
					}
				case *System:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SVG_Path":
						inst.SVG_Path = GongExtractString(rhs)
					case "InverseAppliedScaling":
						inst.InverseAppliedScaling = GongExtractFloat(rhs)
					case "DiagramStructures":
						__gong__assignSliceOfPointers(&inst.DiagramStructures, rhs, identifierMap)
					case "DiagramStructureWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramStructureWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsSubSystemNodeExpanded":
						inst.IsSubSystemNodeExpanded = GongExtractBool(rhs)
					case "SubSystemes":
						__gong__assignSliceOfPointers(&inst.SubSystemes, rhs, identifierMap)
					case "Parts":
						__gong__assignSliceOfPointers(&inst.Parts, rhs, identifierMap)
					case "PartWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PartWhoseNodeIsExpanded, rhs, identifierMap)
					case "DataFlows":
						__gong__assignSliceOfPointers(&inst.DataFlows, rhs, identifierMap)
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "ExternalParts":
						__gong__assignSliceOfPointers(&inst.ExternalParts, rhs, identifierMap)
					case "ExternalPartWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalPartWhoseNodeIsExpanded, rhs, identifierMap)
					}
				case *SystemShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System":
						__gong__assignPointer(&inst.System, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
type AllocatedResourceShapeUnmarshaller struct{}

func (u *AllocatedResourceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AllocatedResourceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AllocatedResourceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AllocatedResourceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
	}
	return nil
}

type AllocatedSystemShapeUnmarshaller struct{}

func (u *AllocatedSystemShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AllocatedSystemShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AllocatedSystemShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AllocatedSystemShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "System":
		GongUnmarshallPointer(&instance.System, valueExpr, identifierMap)
	}
	return nil
}

type ControlFlowUnmarshaller struct{}

func (u *ControlFlowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlFlow), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlFlowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlFlow)
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
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	}
	return nil
}

type ControlFlowShapeUnmarshaller struct{}

func (u *ControlFlowShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlFlowShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlFlowShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlFlowShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ControlFlow":
		GongUnmarshallPointer(&instance.ControlFlow, valueExpr, identifierMap)
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

type DataUnmarshaller struct{}

func (u *DataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Data), stage, identifier, instanceName, preserveOrder)
}

func (u *DataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Data)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Acronym":
		instance.Acronym = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	}
	return nil
}

type DataFlowUnmarshaller struct{}

func (u *DataFlowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DataFlow), stage, identifier, instanceName, preserveOrder)
}

func (u *DataFlowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DataFlow)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartPort":
		GongUnmarshallPointer(&instance.StartPort, valueExpr, identifierMap)
	case "EndPort":
		GongUnmarshallPointer(&instance.EndPort, valueExpr, identifierMap)
	case "StartExternalPart":
		GongUnmarshallPointer(&instance.StartExternalPart, valueExpr, identifierMap)
	case "EndExternalPart":
		GongUnmarshallPointer(&instance.EndExternalPart, valueExpr, identifierMap)
	case "Datas":
		GongUnmarshallSliceOfPointers(&instance.Datas, valueExpr, identifierMap)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "IsDatasNodeExpanded":
		instance.IsDatasNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	}
	return nil
}

type DataFlowShapeUnmarshaller struct{}

func (u *DataFlowShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DataFlowShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DataFlowShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DataFlowShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DataFlow":
		GongUnmarshallPointer(&instance.DataFlow, valueExpr, identifierMap)
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

type DataShapeUnmarshaller struct{}

func (u *DataShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DataShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DataShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DataShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Data":
		GongUnmarshallPointer(&instance.Data, valueExpr, identifierMap)
	case "DataFlow":
		GongUnmarshallPointer(&instance.DataFlow, valueExpr, identifierMap)
	}
	return nil
}

type DiagramLayerStateUnmarshaller struct{}

func (u *DiagramLayerStateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramLayerState), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramLayerStateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramLayerState)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DiagramStructure":
		GongUnmarshallPointer(&instance.DiagramStructure, valueExpr, identifierMap)
	case "LayerDefinition":
		GongUnmarshallPointer(&instance.LayerDefinition, valueExpr, identifierMap)
	}
	return nil
}

type DiagramStructureUnmarshaller struct{}

func (u *DiagramStructureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramStructure), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramStructureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramStructure)
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
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "IsEditable_":
		instance.IsEditable_ = GongExtractBool(valueExpr)
	case "IsShowPrefix":
		instance.IsShowPrefix = GongExtractBool(valueExpr)
	case "DefaultBoxWidth":
		instance.DefaultBoxWidth = GongExtractFloat(valueExpr)
	case "DefaultBoxHeigth":
		instance.DefaultBoxHeigth = GongExtractFloat(valueExpr)
	case "IsWithDiscretePorts":
		instance.IsWithDiscretePorts = GongExtractBool(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "System_Shapes":
		GongUnmarshallSliceOfPointers(&instance.System_Shapes, valueExpr, identifierMap)
	case "IsSystemsNodeExpanded":
		instance.IsSystemsNodeExpanded = GongExtractBool(valueExpr)
	case "SystemsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SystemsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Part_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Part_Shapes, valueExpr, identifierMap)
	case "IsPartsNodeExpanded":
		instance.IsPartsNodeExpanded = GongExtractBool(valueExpr)
	case "PartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PartWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalPart_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ExternalPart_Shapes, valueExpr, identifierMap)
	case "IsExternalPartsNodeExpanded":
		instance.IsExternalPartsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalPartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartsWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PortsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Port_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Port_Shapes, valueExpr, identifierMap)
	case "ControlFlowsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ControlFlowsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ControlFlow_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ControlFlow_Shapes, valueExpr, identifierMap)
	case "DataFlowsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DataFlowsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DataFlow_Shapes":
		GongUnmarshallSliceOfPointers(&instance.DataFlow_Shapes, valueExpr, identifierMap)
	case "DatasWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DatasWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Data_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Data_Shapes, valueExpr, identifierMap)
	case "DataFlowsWhoseDataNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DataFlowsWhoseDataNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedResourcesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.AllocatedResourcesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedResourceShapes":
		GongUnmarshallSliceOfPointers(&instance.AllocatedResourceShapes, valueExpr, identifierMap)
	case "AllocatedSystemesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.AllocatedSystemesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedSystemShapes":
		GongUnmarshallSliceOfPointers(&instance.AllocatedSystemShapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NotePortShapes":
		GongUnmarshallSliceOfPointers(&instance.NotePortShapes, valueExpr, identifierMap)
	case "NotePartShapes":
		GongUnmarshallSliceOfPointers(&instance.NotePartShapes, valueExpr, identifierMap)
	}
	return nil
}

type ExternalPartShapeUnmarshaller struct{}

func (u *ExternalPartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ExternalPartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ExternalPartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExternalPartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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
	case "TailHeigth":
		instance.TailHeigth = GongExtractFloat(valueExpr)
	}
	return nil
}

type LayerDefinitionUnmarshaller struct{}

func (u *LayerDefinitionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LayerDefinition), stage, identifier, instanceName, preserveOrder)
}

func (u *LayerDefinitionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LayerDefinition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Query":
		GongUnmarshallSliceOfPointers(&instance.Query, valueExpr, identifierMap)
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
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "IsSubLibrariesNodeExpanded":
		instance.IsSubLibrariesNodeExpanded = GongExtractBool(valueExpr)
	case "SubLibrariesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SubLibrariesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "RootSystemes":
		GongUnmarshallSliceOfPointers(&instance.RootSystemes, valueExpr, identifierMap)
	case "IsSystemesNodeExpanded":
		instance.IsSystemesNodeExpanded = GongExtractBool(valueExpr)
	case "SystemsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SystemsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootDataFlows":
		GongUnmarshallSliceOfPointers(&instance.RootDataFlows, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "DataFlowsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DataFlowsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootDatas":
		GongUnmarshallSliceOfPointers(&instance.RootDatas, valueExpr, identifierMap)
	case "IsDatasNodeExpanded":
		instance.IsDatasNodeExpanded = GongExtractBool(valueExpr)
	case "DatasWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DatasWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootResources":
		GongUnmarshallSliceOfPointers(&instance.RootResources, valueExpr, identifierMap)
	case "IsResourcesNodeExpanded":
		instance.IsResourcesNodeExpanded = GongExtractBool(valueExpr)
	case "ResourcesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ResourcesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "PartsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PartsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RootNotes":
		GongUnmarshallSliceOfPointers(&instance.RootNotes, valueExpr, identifierMap)
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
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsPartsNodeExpanded":
		instance.IsPartsNodeExpanded = GongExtractBool(valueExpr)
	case "Parts":
		GongUnmarshallSliceOfPointers(&instance.Parts, valueExpr, identifierMap)
	case "IsPortsNodeExpanded":
		instance.IsPortsNodeExpanded = GongExtractBool(valueExpr)
	case "Ports":
		GongUnmarshallSliceOfPointers(&instance.Ports, valueExpr, identifierMap)
	}
	return nil
}

type NotePartShapeUnmarshaller struct{}

func (u *NotePartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
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

type NotePortShapeUnmarshaller struct{}

func (u *NotePortShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePortShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePortShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePortShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Port":
		GongUnmarshallPointer(&instance.Port, valueExpr, identifierMap)
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

type PartUnmarshaller struct{}

func (u *PartUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part), stage, identifier, instanceName, preserveOrder)
}

func (u *PartUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Ports":
		GongUnmarshallSliceOfPointers(&instance.Ports, valueExpr, identifierMap)
	case "TypeOfPart":
		GongUnmarshallPointer(&instance.TypeOfPart, valueExpr, identifierMap)
	case "IsPartNameNotSystemName":
		instance.IsPartNameNotSystemName = GongExtractBool(valueExpr)
	case "IsControlFlowsNodeExpanded":
		instance.IsControlFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ControlFlows":
		GongUnmarshallSliceOfPointers(&instance.ControlFlows, valueExpr, identifierMap)
	case "PortWhoseOutControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseOutControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PortWhoseInControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseInControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "PortWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PortWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PortWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "PartAnchoredPath":
		GongUnmarshallSliceOfPointers(&instance.PartAnchoredPath, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsPortsNodeExpanded":
		instance.IsPortsNodeExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type PartAnchoredPathUnmarshaller struct{}

func (u *PartAnchoredPathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartAnchoredPath), stage, identifier, instanceName, preserveOrder)
}

func (u *PartAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartAnchoredPath)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
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
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type PartShapeUnmarshaller struct{}

func (u *PartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallPointer(&instance.Part, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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

type PortUnmarshaller struct{}

func (u *PortUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Port), stage, identifier, instanceName, preserveOrder)
}

func (u *PortUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Port)
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
	}
	return nil
}

type PortShapeUnmarshaller struct{}

func (u *PortShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PortShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PortShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PortShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Port":
		GongUnmarshallPointer(&instance.Port, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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

type ResourceUnmarshaller struct{}

func (u *ResourceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Resource), stage, identifier, instanceName, preserveOrder)
}

func (u *ResourceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Resource)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Acronym":
		instance.Acronym = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	}
	return nil
}

type SemanticTagUnmarshaller struct{}

func (u *SemanticTagUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SemanticTag), stage, identifier, instanceName, preserveOrder)
}

func (u *SemanticTagUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SemanticTag)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parts":
		GongUnmarshallSliceOfPointers(&instance.Parts, valueExpr, identifierMap)
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
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SVG_Path":
		instance.SVG_Path = GongExtractString(valueExpr)
	case "InverseAppliedScaling":
		instance.InverseAppliedScaling = GongExtractFloat(valueExpr)
	case "DiagramStructures":
		GongUnmarshallSliceOfPointers(&instance.DiagramStructures, valueExpr, identifierMap)
	case "DiagramStructureWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramStructureWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubSystemNodeExpanded":
		instance.IsSubSystemNodeExpanded = GongExtractBool(valueExpr)
	case "SubSystemes":
		GongUnmarshallSliceOfPointers(&instance.SubSystemes, valueExpr, identifierMap)
	case "Parts":
		GongUnmarshallSliceOfPointers(&instance.Parts, valueExpr, identifierMap)
	case "PartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.PartWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DataFlows":
		GongUnmarshallSliceOfPointers(&instance.DataFlows, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalParts":
		GongUnmarshallSliceOfPointers(&instance.ExternalParts, valueExpr, identifierMap)
	case "ExternalPartWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalPartWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type SystemShapeUnmarshaller struct{}

func (u *SystemShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SystemShape), stage, identifier, instanceName, preserveOrder)
}

func (u *SystemShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SystemShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "System":
		GongUnmarshallPointer(&instance.System, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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

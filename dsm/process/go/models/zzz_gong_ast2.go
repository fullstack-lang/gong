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
		case "github.com/fullstack-lang/gong/dsm/process/go/models":
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
				case "AllocatedProcessShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AllocatedProcessShape), stage, ident.Name, instanceName, preserveOrder)
				case "AllocatedResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AllocatedResourceShape), stage, ident.Name, instanceName, preserveOrder)
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
				case "DiagramProcess":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramProcess), stage, ident.Name, instanceName, preserveOrder)
				case "ExternalParticipantShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ExternalParticipantShape), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteTaskShape), stage, ident.Name, instanceName, preserveOrder)
				case "Participant":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Participant), stage, ident.Name, instanceName, preserveOrder)
				case "ParticipantShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParticipantShape), stage, ident.Name, instanceName, preserveOrder)
				case "Process":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Process), stage, ident.Name, instanceName, preserveOrder)
				case "ProcessShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcessShape), stage, ident.Name, instanceName, preserveOrder)
				case "Resource":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Resource), stage, ident.Name, instanceName, preserveOrder)
				case "Task":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Task), stage, ident.Name, instanceName, preserveOrder)
				case "TaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskShape), stage, ident.Name, instanceName, preserveOrder)
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
				case *AllocatedProcessShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						__gong__assignPointer(&inst.Participant, rhs, identifierMap)
					case "Process":
						__gong__assignPointer(&inst.Process, rhs, identifierMap)
					}
				case *AllocatedResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						__gong__assignPointer(&inst.Participant, rhs, identifierMap)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
					case "Datas":
						__gong__assignSliceOfPointers(&inst.Datas, rhs, identifierMap)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Type":
						GongUnmarshallEnum(&inst.Type, rhs)
					case "StartTask":
						__gong__assignPointer(&inst.StartTask, rhs, identifierMap)
					case "EndTask":
						__gong__assignPointer(&inst.EndTask, rhs, identifierMap)
					case "StartExternalParticipant":
						__gong__assignPointer(&inst.StartExternalParticipant, rhs, identifierMap)
					case "EndExternalParticipant":
						__gong__assignPointer(&inst.EndExternalParticipant, rhs, identifierMap)
					case "IsDatasNodeExpanded":
						inst.IsDatasNodeExpanded = GongExtractBool(rhs)
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
				case *DiagramProcess:
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
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "Process_Shapes":
						__gong__assignSliceOfPointers(&inst.Process_Shapes, rhs, identifierMap)
					case "IsProcesssNodeExpanded":
						inst.IsProcesssNodeExpanded = GongExtractBool(rhs)
					case "ProcesssWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ProcesssWhoseNodeIsExpanded, rhs, identifierMap)
					case "Participant_Shapes":
						__gong__assignSliceOfPointers(&inst.Participant_Shapes, rhs, identifierMap)
					case "IsParticipantsNodeExpanded":
						inst.IsParticipantsNodeExpanded = GongExtractBool(rhs)
					case "ParticipantWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ParticipantWhoseNodeIsExpanded, rhs, identifierMap)
					case "ExternalParticipant_Shapes":
						__gong__assignSliceOfPointers(&inst.ExternalParticipant_Shapes, rhs, identifierMap)
					case "IsExternalParticipantsNodeExpanded":
						inst.IsExternalParticipantsNodeExpanded = GongExtractBool(rhs)
					case "ExternalParticipantWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalParticipantWhoseNodeIsExpanded, rhs, identifierMap)
					case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "TasksWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseNodeIsExpanded, rhs, identifierMap)
					case "Task_Shapes":
						__gong__assignSliceOfPointers(&inst.Task_Shapes, rhs, identifierMap)
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
					case "AllocatedProcessesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.AllocatedProcessesWhoseNodeIsExpanded, rhs, identifierMap)
					case "AllocatedProcessShapes":
						__gong__assignSliceOfPointers(&inst.AllocatedProcessShapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NoteTaskShapes":
						__gong__assignSliceOfPointers(&inst.NoteTaskShapes, rhs, identifierMap)
					}
				case *ExternalParticipantShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						__gong__assignPointer(&inst.Participant, rhs, identifierMap)
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
					case "RootProcesses":
						__gong__assignSliceOfPointers(&inst.RootProcesses, rhs, identifierMap)
					case "IsProcessesNodeExpanded":
						inst.IsProcessesNodeExpanded = GongExtractBool(rhs)
					case "ProcesssWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ProcesssWhoseNodeIsExpanded, rhs, identifierMap)
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
					case "ParticipantsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ParticipantsWhoseNodeIsExpanded, rhs, identifierMap)
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
					case "IsTasksNodeExpanded":
						inst.IsTasksNodeExpanded = GongExtractBool(rhs)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
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
				case *NoteTaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
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
				case *Participant:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsProcessResource":
						inst.IsProcessResource = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Resources":
						__gong__assignSliceOfPointers(&inst.Resources, rhs, identifierMap)
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "Processes":
						__gong__assignSliceOfPointers(&inst.Processes, rhs, identifierMap)
					case "IsProcessesNodeExpanded":
						inst.IsProcessesNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsTasksNodeExpanded":
						inst.IsTasksNodeExpanded = GongExtractBool(rhs)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					case "IsControlFlowsNodeExpanded":
						inst.IsControlFlowsNodeExpanded = GongExtractBool(rhs)
					case "ControlFlows":
						__gong__assignSliceOfPointers(&inst.ControlFlows, rhs, identifierMap)
					case "TaskWhoseOutControlFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TaskWhoseOutControlFlowsNodeIsExpanded, rhs, identifierMap)
					case "TaskWhoseInControlFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TaskWhoseInControlFlowsNodeIsExpanded, rhs, identifierMap)
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "TaskWhoseOutDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TaskWhoseOutDataFlowsNodeIsExpanded, rhs, identifierMap)
					case "TaskWhoseInDataFlowsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TaskWhoseInDataFlowsNodeIsExpanded, rhs, identifierMap)
					}
				case *ParticipantShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						__gong__assignPointer(&inst.Participant, rhs, identifierMap)
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
					case "WidthWeight":
						inst.WidthWeight = GongExtractFloat(rhs)
					}
				case *Process:
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
					case "DiagramProcesss":
						__gong__assignSliceOfPointers(&inst.DiagramProcesss, rhs, identifierMap)
					case "DiagramProcessWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramProcessWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsSubProcessNodeExpanded":
						inst.IsSubProcessNodeExpanded = GongExtractBool(rhs)
					case "SubProcesses":
						__gong__assignSliceOfPointers(&inst.SubProcesses, rhs, identifierMap)
					case "Participants":
						__gong__assignSliceOfPointers(&inst.Participants, rhs, identifierMap)
					case "ParticipantWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ParticipantWhoseNodeIsExpanded, rhs, identifierMap)
					case "DataFlows":
						__gong__assignSliceOfPointers(&inst.DataFlows, rhs, identifierMap)
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "ExternalParticipants":
						__gong__assignSliceOfPointers(&inst.ExternalParticipants, rhs, identifierMap)
					case "ExternalParticipantWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ExternalParticipantWhoseNodeIsExpanded, rhs, identifierMap)
					}
				case *ProcessShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Process":
						__gong__assignPointer(&inst.Process, rhs, identifierMap)
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
				case *Task:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsStartTask":
						inst.IsStartTask = GongExtractBool(rhs)
					case "IsEndTask":
						inst.IsEndTask = GongExtractBool(rhs)
					case "Type":
						__gong__assignPointer(&inst.Type, rhs, identifierMap)
					case "IsTaskNameNotProcessName":
						inst.IsTaskNameNotProcessName = GongExtractBool(rhs)
					}
				case *TaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
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
type AllocatedProcessShapeUnmarshaller struct{}

func (u *AllocatedProcessShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AllocatedProcessShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AllocatedProcessShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AllocatedProcessShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
	case "Process":
		GongUnmarshallPointer(&instance.Process, valueExpr, identifierMap)
	}
	return nil
}

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
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
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
	case "Datas":
		GongUnmarshallSliceOfPointers(&instance.Datas, valueExpr, identifierMap)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "StartTask":
		GongUnmarshallPointer(&instance.StartTask, valueExpr, identifierMap)
	case "EndTask":
		GongUnmarshallPointer(&instance.EndTask, valueExpr, identifierMap)
	case "StartExternalParticipant":
		GongUnmarshallPointer(&instance.StartExternalParticipant, valueExpr, identifierMap)
	case "EndExternalParticipant":
		GongUnmarshallPointer(&instance.EndExternalParticipant, valueExpr, identifierMap)
	case "IsDatasNodeExpanded":
		instance.IsDatasNodeExpanded = GongExtractBool(valueExpr)
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

type DiagramProcessUnmarshaller struct{}

func (u *DiagramProcessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramProcess), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramProcessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramProcess)
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
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "Process_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Process_Shapes, valueExpr, identifierMap)
	case "IsProcesssNodeExpanded":
		instance.IsProcesssNodeExpanded = GongExtractBool(valueExpr)
	case "ProcesssWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ProcesssWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Participant_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Participant_Shapes, valueExpr, identifierMap)
	case "IsParticipantsNodeExpanded":
		instance.IsParticipantsNodeExpanded = GongExtractBool(valueExpr)
	case "ParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalParticipant_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipant_Shapes, valueExpr, identifierMap)
	case "IsExternalParticipantsNodeExpanded":
		instance.IsExternalParticipantsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "TasksWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TasksWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Task_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Task_Shapes, valueExpr, identifierMap)
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
	case "AllocatedProcessesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.AllocatedProcessesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "AllocatedProcessShapes":
		GongUnmarshallSliceOfPointers(&instance.AllocatedProcessShapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NoteTaskShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteTaskShapes, valueExpr, identifierMap)
	}
	return nil
}

type ExternalParticipantShapeUnmarshaller struct{}

func (u *ExternalParticipantShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ExternalParticipantShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ExternalParticipantShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExternalParticipantShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
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
	case "RootProcesses":
		GongUnmarshallSliceOfPointers(&instance.RootProcesses, valueExpr, identifierMap)
	case "IsProcessesNodeExpanded":
		instance.IsProcessesNodeExpanded = GongExtractBool(valueExpr)
	case "ProcesssWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ProcesssWhoseNodeIsExpanded, valueExpr, identifierMap)
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
	case "ParticipantsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParticipantsWhoseNodeIsExpanded, valueExpr, identifierMap)
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
	case "IsTasksNodeExpanded":
		instance.IsTasksNodeExpanded = GongExtractBool(valueExpr)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
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

type NoteTaskShapeUnmarshaller struct{}

func (u *NoteTaskShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteTaskShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteTaskShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteTaskShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Task":
		GongUnmarshallPointer(&instance.Task, valueExpr, identifierMap)
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

type ParticipantUnmarshaller struct{}

func (u *ParticipantUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Participant), stage, identifier, instanceName, preserveOrder)
}

func (u *ParticipantUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Participant)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsProcessResource":
		instance.IsProcessResource = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Resources":
		GongUnmarshallSliceOfPointers(&instance.Resources, valueExpr, identifierMap)
	case "IsResourcesNodeExpanded":
		instance.IsResourcesNodeExpanded = GongExtractBool(valueExpr)
	case "Processes":
		GongUnmarshallSliceOfPointers(&instance.Processes, valueExpr, identifierMap)
	case "IsProcessesNodeExpanded":
		instance.IsProcessesNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsTasksNodeExpanded":
		instance.IsTasksNodeExpanded = GongExtractBool(valueExpr)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
	case "IsControlFlowsNodeExpanded":
		instance.IsControlFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ControlFlows":
		GongUnmarshallSliceOfPointers(&instance.ControlFlows, valueExpr, identifierMap)
	case "TaskWhoseOutControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseOutControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "TaskWhoseInControlFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseInControlFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "TaskWhoseOutDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseOutDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	case "TaskWhoseInDataFlowsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskWhoseInDataFlowsNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type ParticipantShapeUnmarshaller struct{}

func (u *ParticipantShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParticipantShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParticipantShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParticipantShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Participant":
		GongUnmarshallPointer(&instance.Participant, valueExpr, identifierMap)
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
	case "WidthWeight":
		instance.WidthWeight = GongExtractFloat(valueExpr)
	}
	return nil
}

type ProcessUnmarshaller struct{}

func (u *ProcessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Process), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Process)
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
	case "DiagramProcesss":
		GongUnmarshallSliceOfPointers(&instance.DiagramProcesss, valueExpr, identifierMap)
	case "DiagramProcessWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramProcessWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubProcessNodeExpanded":
		instance.IsSubProcessNodeExpanded = GongExtractBool(valueExpr)
	case "SubProcesses":
		GongUnmarshallSliceOfPointers(&instance.SubProcesses, valueExpr, identifierMap)
	case "Participants":
		GongUnmarshallSliceOfPointers(&instance.Participants, valueExpr, identifierMap)
	case "ParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DataFlows":
		GongUnmarshallSliceOfPointers(&instance.DataFlows, valueExpr, identifierMap)
	case "IsDataFlowsNodeExpanded":
		instance.IsDataFlowsNodeExpanded = GongExtractBool(valueExpr)
	case "ExternalParticipants":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipants, valueExpr, identifierMap)
	case "ExternalParticipantWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ExternalParticipantWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type ProcessShapeUnmarshaller struct{}

func (u *ProcessShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcessShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcessShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcessShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Process":
		GongUnmarshallPointer(&instance.Process, valueExpr, identifierMap)
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

type TaskUnmarshaller struct{}

func (u *TaskUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Task), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Task)
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
	case "IsStartTask":
		instance.IsStartTask = GongExtractBool(valueExpr)
	case "IsEndTask":
		instance.IsEndTask = GongExtractBool(valueExpr)
	case "Type":
		GongUnmarshallPointer(&instance.Type, valueExpr, identifierMap)
	case "IsTaskNameNotProcessName":
		instance.IsTaskNameNotProcessName = GongExtractBool(valueExpr)
	}
	return nil
}

type TaskShapeUnmarshaller struct{}

func (u *TaskShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Task":
		GongUnmarshallPointer(&instance.Task, valueExpr, identifierMap)
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

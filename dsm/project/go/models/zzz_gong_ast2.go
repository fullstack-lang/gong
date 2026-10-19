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
		case "github.com/fullstack-lang/gong/dsm/project/go/models":
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
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "NoteProductShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteProductShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteResourceShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteTaskShape), stage, ident.Name, instanceName, preserveOrder)
				case "Product":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Product), stage, ident.Name, instanceName, preserveOrder)
				case "ProductCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProductCompositionShape), stage, ident.Name, instanceName, preserveOrder)
				case "ProductReferenceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProductReferenceShape), stage, ident.Name, instanceName, preserveOrder)
				case "ProductShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProductShape), stage, ident.Name, instanceName, preserveOrder)
				case "Resource":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Resource), stage, ident.Name, instanceName, preserveOrder)
				case "ResourceCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceCompositionShape), stage, ident.Name, instanceName, preserveOrder)
				case "ResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceShape), stage, ident.Name, instanceName, preserveOrder)
				case "ResourceTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceTaskShape), stage, ident.Name, instanceName, preserveOrder)
				case "Task":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Task), stage, ident.Name, instanceName, preserveOrder)
				case "TaskCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskCompositionShape), stage, ident.Name, instanceName, preserveOrder)
				case "TaskGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskGroup), stage, ident.Name, instanceName, preserveOrder)
				case "TaskGroupShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskGroupShape), stage, ident.Name, instanceName, preserveOrder)
				case "TaskInputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskInputShape), stage, ident.Name, instanceName, preserveOrder)
				case "TaskOutputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskOutputShape), stage, ident.Name, instanceName, preserveOrder)
				case "TaskPredecessorShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskPredecessorShape), stage, ident.Name, instanceName, preserveOrder)
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
				case *Diagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultBoxWidth":
						inst.DefaultBoxWidth = GongExtractFloat(rhs)
					case "DefaultBoxHeigth":
						inst.DefaultBoxHeigth = GongExtractFloat(rhs)
					case "DateFormat":
						inst.DateFormat = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsTimeDiagram":
						inst.IsTimeDiagram = GongExtractBool(rhs)
					case "ComputedStart":
						inst.ComputedStart = GongExtractDate(rhs)
					case "ComputedEnd":
						inst.ComputedEnd = GongExtractDate(rhs)
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "DrawVerticalTimeLines":
						inst.DrawVerticalTimeLines = GongExtractBool(rhs)
					case "DrawSecondaryVerticalTimeLines":
						inst.DrawSecondaryVerticalTimeLines = GongExtractBool(rhs)
					case "HideWeekendsPeriod":
						inst.HideWeekendsPeriod = GongExtractBool(rhs)
					case "UseManualStartAndEndDates":
						inst.UseManualStartAndEndDates = GongExtractBool(rhs)
					case "ManualStart":
						inst.ManualStart = GongExtractDate(rhs)
					case "ManualEnd":
						inst.ManualEnd = GongExtractDate(rhs)
					case "TimeStep":
						inst.TimeStep = GongExtractInt(rhs)
					case "TimeStepScale":
						GongUnmarshallEnum(&inst.TimeStepScale, rhs)
					case "SecondaryTimeStep":
						inst.SecondaryTimeStep = GongExtractInt(rhs)
					case "SecondaryTimeStepScale":
						GongUnmarshallEnum(&inst.SecondaryTimeStepScale, rhs)
					case "LaneHeight":
						inst.LaneHeight = GongExtractFloat(rhs)
					case "RatioBarToLaneHeight":
						inst.RatioBarToLaneHeight = GongExtractFloat(rhs)
					case "YTopMargin":
						inst.YTopMargin = GongExtractFloat(rhs)
					case "XLeftText":
						inst.XLeftText = GongExtractFloat(rhs)
					case "TextHeight":
						inst.TextHeight = GongExtractFloat(rhs)
					case "XLeftLanes":
						inst.XLeftLanes = GongExtractFloat(rhs)
					case "XRightMargin":
						inst.XRightMargin = GongExtractFloat(rhs)
					case "ArrowLengthToTheRightOfStartBar":
						inst.ArrowLengthToTheRightOfStartBar = GongExtractFloat(rhs)
					case "ArrowTipLenght":
						inst.ArrowTipLenght = GongExtractFloat(rhs)
					case "TimeLine_Color":
						inst.TimeLine_Color = GongExtractString(rhs)
					case "TimeLine_FillOpacity":
						inst.TimeLine_FillOpacity = GongExtractFloat(rhs)
					case "TimeLine_Stroke":
						inst.TimeLine_Stroke = GongExtractString(rhs)
					case "TimeLine_StrokeWidth":
						inst.TimeLine_StrokeWidth = GongExtractFloat(rhs)
					case "Group_Stroke":
						inst.Group_Stroke = GongExtractString(rhs)
					case "Group_StrokeWidth":
						inst.Group_StrokeWidth = GongExtractFloat(rhs)
					case "Group_StrokeDashArray":
						inst.Group_StrokeDashArray = GongExtractString(rhs)
					case "DateYOffset":
						inst.DateYOffset = GongExtractFloat(rhs)
					case "AlignOnBeginningOfTimeScale":
						inst.AlignOnBeginningOfTimeScale = GongExtractBool(rhs)
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
					case "IsInAutoLayoutMode":
						inst.IsInAutoLayoutMode = GongExtractBool(rhs)
					case "Product_Shapes":
						__gong__assignSliceOfPointers(&inst.Product_Shapes, rhs, identifierMap)
					case "ProductsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ProductsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPBSNodeExpanded":
						inst.IsPBSNodeExpanded = GongExtractBool(rhs)
					case "ProductComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.ProductComposition_Shapes, rhs, identifierMap)
					case "ProductReference_Shapes":
						__gong__assignSliceOfPointers(&inst.ProductReference_Shapes, rhs, identifierMap)
					case "IsWBSNodeExpanded":
						inst.IsWBSNodeExpanded = GongExtractBool(rhs)
					case "Task_Shapes":
						__gong__assignSliceOfPointers(&inst.Task_Shapes, rhs, identifierMap)
					case "TasksWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseNodeIsExpanded, rhs, identifierMap)
					case "TasksWhoseInputNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseInputNodeIsExpanded, rhs, identifierMap)
					case "TasksWhoseOutputNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseOutputNodeIsExpanded, rhs, identifierMap)
					case "TasksWhosePredecessorNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhosePredecessorNodeIsExpanded, rhs, identifierMap)
					case "IsTaskGroupsNodeExpanded":
						inst.IsTaskGroupsNodeExpanded = GongExtractBool(rhs)
					case "TaskGroupShapes":
						__gong__assignSliceOfPointers(&inst.TaskGroupShapes, rhs, identifierMap)
					case "TaskGroupsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TaskGroupsWhoseNodeIsExpanded, rhs, identifierMap)
					case "TaskComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.TaskComposition_Shapes, rhs, identifierMap)
					case "TaskInputShapes":
						__gong__assignSliceOfPointers(&inst.TaskInputShapes, rhs, identifierMap)
					case "TaskOutputShapes":
						__gong__assignSliceOfPointers(&inst.TaskOutputShapes, rhs, identifierMap)
					case "TaskPredecessorShapes":
						__gong__assignSliceOfPointers(&inst.TaskPredecessorShapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NoteProductShapes":
						__gong__assignSliceOfPointers(&inst.NoteProductShapes, rhs, identifierMap)
					case "NoteTaskShapes":
						__gong__assignSliceOfPointers(&inst.NoteTaskShapes, rhs, identifierMap)
					case "NoteResourceShapes":
						__gong__assignSliceOfPointers(&inst.NoteResourceShapes, rhs, identifierMap)
					case "Resource_Shapes":
						__gong__assignSliceOfPointers(&inst.Resource_Shapes, rhs, identifierMap)
					case "ResourcesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ResourcesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "ResourceComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.ResourceComposition_Shapes, rhs, identifierMap)
					case "ResourceTaskShapes":
						__gong__assignSliceOfPointers(&inst.ResourceTaskShapes, rhs, identifierMap)
					}
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "RootProducts":
						__gong__assignSliceOfPointers(&inst.RootProducts, rhs, identifierMap)
					case "RootTasks":
						__gong__assignSliceOfPointers(&inst.RootTasks, rhs, identifierMap)
					case "RootTaskGroups":
						__gong__assignSliceOfPointers(&inst.RootTaskGroups, rhs, identifierMap)
					case "RootResources":
						__gong__assignSliceOfPointers(&inst.RootResources, rhs, identifierMap)
					case "Notes":
						__gong__assignSliceOfPointers(&inst.Notes, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					case "Products":
						__gong__assignSliceOfPointers(&inst.Products, rhs, identifierMap)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					case "Resources":
						__gong__assignSliceOfPointers(&inst.Resources, rhs, identifierMap)
					}
				case *NoteProductShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *NoteResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
					case "IsLayoutDirectionDifferent":
						inst.IsLayoutDirectionDifferent = GongExtractBool(rhs)
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
				case *Product:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SubProducts":
						__gong__assignSliceOfPointers(&inst.SubProducts, rhs, identifierMap)
					case "IsProducersNodeExpanded":
						inst.IsProducersNodeExpanded = GongExtractBool(rhs)
					case "IsConsumersNodeExpanded":
						inst.IsConsumersNodeExpanded = GongExtractBool(rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedProduct":
						__gong__assignPointer(&inst.ReferencedProduct, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					}
				case *ProductCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *ProductReferenceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
					case "ReferencedProduct":
						__gong__assignPointer(&inst.ReferencedProduct, rhs, identifierMap)
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
				case *ProductShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
					case "IsShowType":
						inst.IsShowType = GongExtractBool(rhs)
					case "IsLayoutDirectionDifferent":
						inst.IsLayoutDirectionDifferent = GongExtractBool(rhs)
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
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					case "SubResources":
						__gong__assignSliceOfPointers(&inst.SubResources, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedResource":
						__gong__assignPointer(&inst.ReferencedResource, rhs, identifierMap)
					}
				case *ResourceCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
				case *ResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
					case "IsLayoutDirectionDifferent":
						inst.IsLayoutDirectionDifferent = GongExtractBool(rhs)
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
				case *ResourceTaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
				case *Task:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Start":
						inst.Start = GongExtractDate(rhs)
					case "End":
						inst.End = GongExtractDate(rhs)
					case "IsAllDay":
						inst.IsAllDay = GongExtractBool(rhs)
					case "IsMilestone":
						inst.IsMilestone = GongExtractBool(rhs)
					case "TaskGroups":
						__gong__assignSliceOfPointers(&inst.TaskGroups, rhs, identifierMap)
					case "Predecessors":
						__gong__assignSliceOfPointers(&inst.Predecessors, rhs, identifierMap)
					case "DependencyType":
						GongUnmarshallEnum(&inst.DependencyType, rhs)
					case "DependencyDurationYears":
						inst.DependencyDurationYears = GongExtractFloat(rhs)
					case "DependencyDurationMonths":
						inst.DependencyDurationMonths = GongExtractFloat(rhs)
					case "DependencyDurationWeeks":
						inst.DependencyDurationWeeks = GongExtractFloat(rhs)
					case "DependencyDurationDays":
						inst.DependencyDurationDays = GongExtractFloat(rhs)
					case "DependencyDurationHours":
						inst.DependencyDurationHours = GongExtractFloat(rhs)
					case "DurationYears":
						inst.DurationYears = GongExtractFloat(rhs)
					case "DurationMonths":
						inst.DurationMonths = GongExtractFloat(rhs)
					case "DurationWeeks":
						inst.DurationWeeks = GongExtractFloat(rhs)
					case "DurationDays":
						inst.DurationDays = GongExtractFloat(rhs)
					case "DurationHours":
						inst.DurationHours = GongExtractFloat(rhs)
					case "IsEndDateComputedFromDuration":
						inst.IsEndDateComputedFromDuration = GongExtractBool(rhs)
					case "Inputs":
						__gong__assignSliceOfPointers(&inst.Inputs, rhs, identifierMap)
					case "Outputs":
						__gong__assignSliceOfPointers(&inst.Outputs, rhs, identifierMap)
					case "SubTasks":
						__gong__assignSliceOfPointers(&inst.SubTasks, rhs, identifierMap)
					case "IsWithCompletion":
						inst.IsWithCompletion = GongExtractBool(rhs)
					case "Completion":
						GongUnmarshallEnum(&inst.Completion, rhs)
					case "TaskGroupsToDisplay":
						__gong__assignSliceOfPointers(&inst.TaskGroupsToDisplay, rhs, identifierMap)
					case "TextPosition":
						GongUnmarshallEnum(&inst.TextPosition, rhs)
					case "XOffset":
						inst.XOffset = GongExtractFloat(rhs)
					case "YOffset":
						inst.YOffset = GongExtractFloat(rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedTask":
						__gong__assignPointer(&inst.ReferencedTask, rhs, identifierMap)
					case "IsInputsNodeExpanded":
						inst.IsInputsNodeExpanded = GongExtractBool(rhs)
					case "IsOutputsNodeExpanded":
						inst.IsOutputsNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					}
				case *TaskCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
				case *TaskGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					}
				case *TaskGroupShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "TaskGroup":
						__gong__assignPointer(&inst.TaskGroup, rhs, identifierMap)
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
				case *TaskInputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *TaskOutputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *TaskPredecessorShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Predecessor":
						__gong__assignPointer(&inst.Predecessor, rhs, identifierMap)
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
				case *TaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
					case "IsShowDate":
						inst.IsShowDate = GongExtractBool(rhs)
					case "VerticalOffset":
						inst.VerticalOffset = GongExtractFloat(rhs)
					case "DisplayVerticalBar":
						inst.DisplayVerticalBar = GongExtractBool(rhs)
					case "IsLayoutDirectionDifferent":
						inst.IsLayoutDirectionDifferent = GongExtractBool(rhs)
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
type DiagramUnmarshaller struct{}

func (u *DiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Diagram), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Diagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DefaultBoxWidth":
		instance.DefaultBoxWidth = GongExtractFloat(valueExpr)
	case "DefaultBoxHeigth":
		instance.DefaultBoxHeigth = GongExtractFloat(valueExpr)
	case "DateFormat":
		instance.DateFormat = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsTimeDiagram":
		instance.IsTimeDiagram = GongExtractBool(valueExpr)
	case "ComputedStart":
		instance.ComputedStart = GongExtractDate(valueExpr)
	case "ComputedEnd":
		instance.ComputedEnd = GongExtractDate(valueExpr)
	case "ComputedDuration":
		instance.ComputedDuration = time.Duration(GongExtractInt(valueExpr))
	case "DrawVerticalTimeLines":
		instance.DrawVerticalTimeLines = GongExtractBool(valueExpr)
	case "DrawSecondaryVerticalTimeLines":
		instance.DrawSecondaryVerticalTimeLines = GongExtractBool(valueExpr)
	case "HideWeekendsPeriod":
		instance.HideWeekendsPeriod = GongExtractBool(valueExpr)
	case "UseManualStartAndEndDates":
		instance.UseManualStartAndEndDates = GongExtractBool(valueExpr)
	case "ManualStart":
		instance.ManualStart = GongExtractDate(valueExpr)
	case "ManualEnd":
		instance.ManualEnd = GongExtractDate(valueExpr)
	case "TimeStep":
		instance.TimeStep = GongExtractInt(valueExpr)
	case "TimeStepScale":
		GongUnmarshallEnum(&instance.TimeStepScale, valueExpr)
	case "SecondaryTimeStep":
		instance.SecondaryTimeStep = GongExtractInt(valueExpr)
	case "SecondaryTimeStepScale":
		GongUnmarshallEnum(&instance.SecondaryTimeStepScale, valueExpr)
	case "LaneHeight":
		instance.LaneHeight = GongExtractFloat(valueExpr)
	case "RatioBarToLaneHeight":
		instance.RatioBarToLaneHeight = GongExtractFloat(valueExpr)
	case "YTopMargin":
		instance.YTopMargin = GongExtractFloat(valueExpr)
	case "XLeftText":
		instance.XLeftText = GongExtractFloat(valueExpr)
	case "TextHeight":
		instance.TextHeight = GongExtractFloat(valueExpr)
	case "XLeftLanes":
		instance.XLeftLanes = GongExtractFloat(valueExpr)
	case "XRightMargin":
		instance.XRightMargin = GongExtractFloat(valueExpr)
	case "ArrowLengthToTheRightOfStartBar":
		instance.ArrowLengthToTheRightOfStartBar = GongExtractFloat(valueExpr)
	case "ArrowTipLenght":
		instance.ArrowTipLenght = GongExtractFloat(valueExpr)
	case "TimeLine_Color":
		instance.TimeLine_Color = GongExtractString(valueExpr)
	case "TimeLine_FillOpacity":
		instance.TimeLine_FillOpacity = GongExtractFloat(valueExpr)
	case "TimeLine_Stroke":
		instance.TimeLine_Stroke = GongExtractString(valueExpr)
	case "TimeLine_StrokeWidth":
		instance.TimeLine_StrokeWidth = GongExtractFloat(valueExpr)
	case "Group_Stroke":
		instance.Group_Stroke = GongExtractString(valueExpr)
	case "Group_StrokeWidth":
		instance.Group_StrokeWidth = GongExtractFloat(valueExpr)
	case "Group_StrokeDashArray":
		instance.Group_StrokeDashArray = GongExtractString(valueExpr)
	case "DateYOffset":
		instance.DateYOffset = GongExtractFloat(valueExpr)
	case "AlignOnBeginningOfTimeScale":
		instance.AlignOnBeginningOfTimeScale = GongExtractBool(valueExpr)
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
	case "IsInAutoLayoutMode":
		instance.IsInAutoLayoutMode = GongExtractBool(valueExpr)
	case "Product_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Product_Shapes, valueExpr, identifierMap)
	case "ProductsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ProductsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsPBSNodeExpanded":
		instance.IsPBSNodeExpanded = GongExtractBool(valueExpr)
	case "ProductComposition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ProductComposition_Shapes, valueExpr, identifierMap)
	case "ProductReference_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ProductReference_Shapes, valueExpr, identifierMap)
	case "IsWBSNodeExpanded":
		instance.IsWBSNodeExpanded = GongExtractBool(valueExpr)
	case "Task_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Task_Shapes, valueExpr, identifierMap)
	case "TasksWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TasksWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "TasksWhoseInputNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TasksWhoseInputNodeIsExpanded, valueExpr, identifierMap)
	case "TasksWhoseOutputNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TasksWhoseOutputNodeIsExpanded, valueExpr, identifierMap)
	case "TasksWhosePredecessorNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TasksWhosePredecessorNodeIsExpanded, valueExpr, identifierMap)
	case "IsTaskGroupsNodeExpanded":
		instance.IsTaskGroupsNodeExpanded = GongExtractBool(valueExpr)
	case "TaskGroupShapes":
		GongUnmarshallSliceOfPointers(&instance.TaskGroupShapes, valueExpr, identifierMap)
	case "TaskGroupsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.TaskGroupsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "TaskComposition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.TaskComposition_Shapes, valueExpr, identifierMap)
	case "TaskInputShapes":
		GongUnmarshallSliceOfPointers(&instance.TaskInputShapes, valueExpr, identifierMap)
	case "TaskOutputShapes":
		GongUnmarshallSliceOfPointers(&instance.TaskOutputShapes, valueExpr, identifierMap)
	case "TaskPredecessorShapes":
		GongUnmarshallSliceOfPointers(&instance.TaskPredecessorShapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NoteProductShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteProductShapes, valueExpr, identifierMap)
	case "NoteTaskShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteTaskShapes, valueExpr, identifierMap)
	case "NoteResourceShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteResourceShapes, valueExpr, identifierMap)
	case "Resource_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Resource_Shapes, valueExpr, identifierMap)
	case "ResourcesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ResourcesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsResourcesNodeExpanded":
		instance.IsResourcesNodeExpanded = GongExtractBool(valueExpr)
	case "ResourceComposition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ResourceComposition_Shapes, valueExpr, identifierMap)
	case "ResourceTaskShapes":
		GongUnmarshallSliceOfPointers(&instance.ResourceTaskShapes, valueExpr, identifierMap)
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
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "RootProducts":
		GongUnmarshallSliceOfPointers(&instance.RootProducts, valueExpr, identifierMap)
	case "RootTasks":
		GongUnmarshallSliceOfPointers(&instance.RootTasks, valueExpr, identifierMap)
	case "RootTaskGroups":
		GongUnmarshallSliceOfPointers(&instance.RootTaskGroups, valueExpr, identifierMap)
	case "RootResources":
		GongUnmarshallSliceOfPointers(&instance.RootResources, valueExpr, identifierMap)
	case "Notes":
		GongUnmarshallSliceOfPointers(&instance.Notes, valueExpr, identifierMap)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
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
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "LayoutDirection":
		GongUnmarshallEnum(&instance.LayoutDirection, valueExpr)
	case "Products":
		GongUnmarshallSliceOfPointers(&instance.Products, valueExpr, identifierMap)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
	case "Resources":
		GongUnmarshallSliceOfPointers(&instance.Resources, valueExpr, identifierMap)
	}
	return nil
}

type NoteProductShapeUnmarshaller struct{}

func (u *NoteProductShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteProductShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteProductShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteProductShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Product":
		GongUnmarshallPointer(&instance.Product, valueExpr, identifierMap)
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

type NoteResourceShapeUnmarshaller struct{}

func (u *NoteResourceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteResourceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteResourceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteResourceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
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
	case "IsLayoutDirectionDifferent":
		instance.IsLayoutDirectionDifferent = GongExtractBool(valueExpr)
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

type ProductUnmarshaller struct{}

func (u *ProductUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Product), stage, identifier, instanceName, preserveOrder)
}

func (u *ProductUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Product)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SubProducts":
		GongUnmarshallSliceOfPointers(&instance.SubProducts, valueExpr, identifierMap)
	case "IsProducersNodeExpanded":
		instance.IsProducersNodeExpanded = GongExtractBool(valueExpr)
	case "IsConsumersNodeExpanded":
		instance.IsConsumersNodeExpanded = GongExtractBool(valueExpr)
	case "IsImport":
		instance.IsImport = GongExtractBool(valueExpr)
	case "ReferencedProduct":
		GongUnmarshallPointer(&instance.ReferencedProduct, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "LayoutDirection":
		GongUnmarshallEnum(&instance.LayoutDirection, valueExpr)
	}
	return nil
}

type ProductCompositionShapeUnmarshaller struct{}

func (u *ProductCompositionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProductCompositionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ProductCompositionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProductCompositionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Product":
		GongUnmarshallPointer(&instance.Product, valueExpr, identifierMap)
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

type ProductReferenceShapeUnmarshaller struct{}

func (u *ProductReferenceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProductReferenceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ProductReferenceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProductReferenceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Product":
		GongUnmarshallPointer(&instance.Product, valueExpr, identifierMap)
	case "ReferencedProduct":
		GongUnmarshallPointer(&instance.ReferencedProduct, valueExpr, identifierMap)
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

type ProductShapeUnmarshaller struct{}

func (u *ProductShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProductShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ProductShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProductShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Product":
		GongUnmarshallPointer(&instance.Product, valueExpr, identifierMap)
	case "IsShowType":
		instance.IsShowType = GongExtractBool(valueExpr)
	case "IsLayoutDirectionDifferent":
		instance.IsLayoutDirectionDifferent = GongExtractBool(valueExpr)
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
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
	case "SubResources":
		GongUnmarshallSliceOfPointers(&instance.SubResources, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "LayoutDirection":
		GongUnmarshallEnum(&instance.LayoutDirection, valueExpr)
	case "IsImport":
		instance.IsImport = GongExtractBool(valueExpr)
	case "ReferencedResource":
		GongUnmarshallPointer(&instance.ReferencedResource, valueExpr, identifierMap)
	}
	return nil
}

type ResourceCompositionShapeUnmarshaller struct{}

func (u *ResourceCompositionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ResourceCompositionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ResourceCompositionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ResourceCompositionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
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

type ResourceShapeUnmarshaller struct{}

func (u *ResourceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ResourceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ResourceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ResourceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
	case "IsLayoutDirectionDifferent":
		instance.IsLayoutDirectionDifferent = GongExtractBool(valueExpr)
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

type ResourceTaskShapeUnmarshaller struct{}

func (u *ResourceTaskShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ResourceTaskShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ResourceTaskShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ResourceTaskShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Resource":
		GongUnmarshallPointer(&instance.Resource, valueExpr, identifierMap)
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
	case "Start":
		instance.Start = GongExtractDate(valueExpr)
	case "End":
		instance.End = GongExtractDate(valueExpr)
	case "IsAllDay":
		instance.IsAllDay = GongExtractBool(valueExpr)
	case "IsMilestone":
		instance.IsMilestone = GongExtractBool(valueExpr)
	case "TaskGroups":
		GongUnmarshallSliceOfPointers(&instance.TaskGroups, valueExpr, identifierMap)
	case "Predecessors":
		GongUnmarshallSliceOfPointers(&instance.Predecessors, valueExpr, identifierMap)
	case "DependencyType":
		GongUnmarshallEnum(&instance.DependencyType, valueExpr)
	case "DependencyDurationYears":
		instance.DependencyDurationYears = GongExtractFloat(valueExpr)
	case "DependencyDurationMonths":
		instance.DependencyDurationMonths = GongExtractFloat(valueExpr)
	case "DependencyDurationWeeks":
		instance.DependencyDurationWeeks = GongExtractFloat(valueExpr)
	case "DependencyDurationDays":
		instance.DependencyDurationDays = GongExtractFloat(valueExpr)
	case "DependencyDurationHours":
		instance.DependencyDurationHours = GongExtractFloat(valueExpr)
	case "DurationYears":
		instance.DurationYears = GongExtractFloat(valueExpr)
	case "DurationMonths":
		instance.DurationMonths = GongExtractFloat(valueExpr)
	case "DurationWeeks":
		instance.DurationWeeks = GongExtractFloat(valueExpr)
	case "DurationDays":
		instance.DurationDays = GongExtractFloat(valueExpr)
	case "DurationHours":
		instance.DurationHours = GongExtractFloat(valueExpr)
	case "IsEndDateComputedFromDuration":
		instance.IsEndDateComputedFromDuration = GongExtractBool(valueExpr)
	case "Inputs":
		GongUnmarshallSliceOfPointers(&instance.Inputs, valueExpr, identifierMap)
	case "Outputs":
		GongUnmarshallSliceOfPointers(&instance.Outputs, valueExpr, identifierMap)
	case "SubTasks":
		GongUnmarshallSliceOfPointers(&instance.SubTasks, valueExpr, identifierMap)
	case "IsWithCompletion":
		instance.IsWithCompletion = GongExtractBool(valueExpr)
	case "Completion":
		GongUnmarshallEnum(&instance.Completion, valueExpr)
	case "TaskGroupsToDisplay":
		GongUnmarshallSliceOfPointers(&instance.TaskGroupsToDisplay, valueExpr, identifierMap)
	case "TextPosition":
		GongUnmarshallEnum(&instance.TextPosition, valueExpr)
	case "XOffset":
		instance.XOffset = GongExtractFloat(valueExpr)
	case "YOffset":
		instance.YOffset = GongExtractFloat(valueExpr)
	case "IsImport":
		instance.IsImport = GongExtractBool(valueExpr)
	case "ReferencedTask":
		GongUnmarshallPointer(&instance.ReferencedTask, valueExpr, identifierMap)
	case "IsInputsNodeExpanded":
		instance.IsInputsNodeExpanded = GongExtractBool(valueExpr)
	case "IsOutputsNodeExpanded":
		instance.IsOutputsNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "LayoutDirection":
		GongUnmarshallEnum(&instance.LayoutDirection, valueExpr)
	}
	return nil
}

type TaskCompositionShapeUnmarshaller struct{}

func (u *TaskCompositionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskCompositionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskCompositionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskCompositionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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

type TaskGroupUnmarshaller struct{}

func (u *TaskGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskGroup), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
	}
	return nil
}

type TaskGroupShapeUnmarshaller struct{}

func (u *TaskGroupShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskGroupShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskGroupShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskGroupShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TaskGroup":
		GongUnmarshallPointer(&instance.TaskGroup, valueExpr, identifierMap)
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

type TaskInputShapeUnmarshaller struct{}

func (u *TaskInputShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskInputShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskInputShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskInputShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Product":
		GongUnmarshallPointer(&instance.Product, valueExpr, identifierMap)
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

type TaskOutputShapeUnmarshaller struct{}

func (u *TaskOutputShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskOutputShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskOutputShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskOutputShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Task":
		GongUnmarshallPointer(&instance.Task, valueExpr, identifierMap)
	case "Product":
		GongUnmarshallPointer(&instance.Product, valueExpr, identifierMap)
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

type TaskPredecessorShapeUnmarshaller struct{}

func (u *TaskPredecessorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaskPredecessorShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TaskPredecessorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaskPredecessorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Predecessor":
		GongUnmarshallPointer(&instance.Predecessor, valueExpr, identifierMap)
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
	case "IsShowDate":
		instance.IsShowDate = GongExtractBool(valueExpr)
	case "VerticalOffset":
		instance.VerticalOffset = GongExtractFloat(valueExpr)
	case "DisplayVerticalBar":
		instance.DisplayVerticalBar = GongExtractBool(valueExpr)
	case "IsLayoutDirectionDifferent":
		instance.IsLayoutDirectionDifferent = GongExtractBool(valueExpr)
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

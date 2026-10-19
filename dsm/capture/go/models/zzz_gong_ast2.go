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
		case "github.com/fullstack-lang/gong/dsm/capture/go/models":
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
				case "AnalysisNeed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AnalysisNeed), stage, ident.Name, instanceName, preserveOrder)
				case "Concept":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Concept), stage, ident.Name, instanceName, preserveOrder)
				case "ConceptShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConceptShape), stage, ident.Name, instanceName, preserveOrder)
				case "Concern":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Concern), stage, ident.Name, instanceName, preserveOrder)
				case "ConcernCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernCompositionShape), stage, ident.Name, instanceName, preserveOrder)
				case "ConcernInputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernInputShape), stage, ident.Name, instanceName, preserveOrder)
				case "ConcernOutputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernOutputShape), stage, ident.Name, instanceName, preserveOrder)
				case "ConcernShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernShape), stage, ident.Name, instanceName, preserveOrder)
				case "ControlPointShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPointShape), stage, ident.Name, instanceName, preserveOrder)
				case "Deliverable":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Deliverable), stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableCompositionShape), stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableConceptShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableConceptShape), stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableShape), stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableViewShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableViewShape), stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stage, ident.Name, instanceName, preserveOrder)
				case "DiagramShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramShape), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "NoteDeliverableShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteDeliverableShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteStakeholderShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteStakeholderShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteTaskShape), stage, ident.Name, instanceName, preserveOrder)
				case "Requirement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Requirement), stage, ident.Name, instanceName, preserveOrder)
				case "RequirementConceptShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RequirementConceptShape), stage, ident.Name, instanceName, preserveOrder)
				case "RequirementConcernShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RequirementConcernShape), stage, ident.Name, instanceName, preserveOrder)
				case "RequirementShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RequirementShape), stage, ident.Name, instanceName, preserveOrder)
				case "RequirementViewShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RequirementViewShape), stage, ident.Name, instanceName, preserveOrder)
				case "Stakeholder":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stakeholder), stage, ident.Name, instanceName, preserveOrder)
				case "StakeholderCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StakeholderCompositionShape), stage, ident.Name, instanceName, preserveOrder)
				case "StakeholderConcernShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StakeholderConcernShape), stage, ident.Name, instanceName, preserveOrder)
				case "StakeholderShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StakeholderShape), stage, ident.Name, instanceName, preserveOrder)
				case "SupportLevel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SupportLevel), stage, ident.Name, instanceName, preserveOrder)
				case "Tool":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tool), stage, ident.Name, instanceName, preserveOrder)
				case "View":
					identifierMap[ident.Name] = __gong__stageSetInit(new(View), stage, ident.Name, instanceName, preserveOrder)
				case "ViewConceptShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ViewConceptShape), stage, ident.Name, instanceName, preserveOrder)
				case "ViewShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ViewShape), stage, ident.Name, instanceName, preserveOrder)
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
				case *AnalysisNeed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Concept:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Tools":
						__gong__assignSliceOfPointers(&inst.Tools, rhs, identifierMap)
					}
				case *ConceptShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concept":
						__gong__assignPointer(&inst.Concept, rhs, identifierMap)
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
				case *Concern:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IDAirbus":
						inst.IDAirbus = GongExtractString(rhs)
					case "Priority":
						GongUnmarshallEnum(&inst.Priority, rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SubConcerns":
						__gong__assignSliceOfPointers(&inst.SubConcerns, rhs, identifierMap)
					case "Inputs":
						__gong__assignSliceOfPointers(&inst.Inputs, rhs, identifierMap)
					case "IsInputsNodeExpanded":
						inst.IsInputsNodeExpanded = GongExtractBool(rhs)
					case "Outputs":
						__gong__assignSliceOfPointers(&inst.Outputs, rhs, identifierMap)
					case "IsOutputsNodeExpanded":
						inst.IsOutputsNodeExpanded = GongExtractBool(rhs)
					case "IsWithCompletion":
						inst.IsWithCompletion = GongExtractBool(rhs)
					case "Completion":
						GongUnmarshallEnum(&inst.Completion, rhs)
					}
				case *ConcernCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concern":
						__gong__assignPointer(&inst.Concern, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *ConcernInputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
					case "Concern":
						__gong__assignPointer(&inst.Concern, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *ConcernOutputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concern":
						__gong__assignPointer(&inst.Concern, rhs, identifierMap)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *ConcernShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concern":
						__gong__assignPointer(&inst.Concern, rhs, identifierMap)
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
				case *ControlPointShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X_Relative":
						inst.X_Relative = GongExtractFloat(rhs)
					case "Y_Relative":
						inst.Y_Relative = GongExtractFloat(rhs)
					case "IsStartShapeTheClosestShape":
						inst.IsStartShapeTheClosestShape = GongExtractBool(rhs)
					}
				case *Deliverable:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SubDeliverables":
						__gong__assignSliceOfPointers(&inst.SubDeliverables, rhs, identifierMap)
					case "IsProducersNodeExpanded":
						inst.IsProducersNodeExpanded = GongExtractBool(rhs)
					case "IsConsumersNodeExpanded":
						inst.IsConsumersNodeExpanded = GongExtractBool(rhs)
					case "Views":
						__gong__assignSliceOfPointers(&inst.Views, rhs, identifierMap)
					case "Concepts":
						__gong__assignSliceOfPointers(&inst.Concepts, rhs, identifierMap)
					}
				case *DeliverableCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *DeliverableConceptShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
					case "Concept":
						__gong__assignPointer(&inst.Concept, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *DeliverableShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
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
				case *DeliverableViewShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
					case "View":
						__gong__assignPointer(&inst.View, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *Diagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsEditable_":
						inst.IsEditable_ = GongExtractBool(rhs)
					case "ShowPrefix":
						inst.ShowPrefix = GongExtractBool(rhs)
					case "DefaultBoxWidth":
						inst.DefaultBoxWidth = GongExtractFloat(rhs)
					case "DefaultBoxHeigth":
						inst.DefaultBoxHeigth = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "ConcernsWhoseRequirementsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConcernsWhoseRequirementsNodeIsExpanded, rhs, identifierMap)
					case "IsRequirementsNodeExpanded":
						inst.IsRequirementsNodeExpanded = GongExtractBool(rhs)
					case "IsViewsNodeExpanded":
						inst.IsViewsNodeExpanded = GongExtractBool(rhs)
					case "IsConceptsNodeExpanded":
						inst.IsConceptsNodeExpanded = GongExtractBool(rhs)
					case "Deliverable_Shapes":
						__gong__assignSliceOfPointers(&inst.Deliverable_Shapes, rhs, identifierMap)
					case "DeliverablesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DeliverablesWhoseNodeIsExpanded, rhs, identifierMap)
					case "DeliverablesWhoseViewsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DeliverablesWhoseViewsNodeIsExpanded, rhs, identifierMap)
					case "DeliverablesWhoseConceptsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DeliverablesWhoseConceptsNodeIsExpanded, rhs, identifierMap)
					case "IsPBSNodeExpanded":
						inst.IsPBSNodeExpanded = GongExtractBool(rhs)
					case "DeliverableComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.DeliverableComposition_Shapes, rhs, identifierMap)
					case "IsConcernsNodeExpanded":
						inst.IsConcernsNodeExpanded = GongExtractBool(rhs)
					case "Concern_Shapes":
						__gong__assignSliceOfPointers(&inst.Concern_Shapes, rhs, identifierMap)
					case "ConcernsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConcernsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ConcernsWhoseInputNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConcernsWhoseInputNodeIsExpanded, rhs, identifierMap)
					case "ConcernsWhoseStakeholderNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConcernsWhoseStakeholderNodeIsExpanded, rhs, identifierMap)
					case "ConcernssWhoseOutputNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConcernssWhoseOutputNodeIsExpanded, rhs, identifierMap)
					case "ConcernComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.ConcernComposition_Shapes, rhs, identifierMap)
					case "ConcernInputShapes":
						__gong__assignSliceOfPointers(&inst.ConcernInputShapes, rhs, identifierMap)
					case "ConcernOutputShapes":
						__gong__assignSliceOfPointers(&inst.ConcernOutputShapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NoteDeliverableShapes":
						__gong__assignSliceOfPointers(&inst.NoteDeliverableShapes, rhs, identifierMap)
					case "NoteTaskShapes":
						__gong__assignSliceOfPointers(&inst.NoteTaskShapes, rhs, identifierMap)
					case "NoteResourceShapes":
						__gong__assignSliceOfPointers(&inst.NoteResourceShapes, rhs, identifierMap)
					case "Stakeholder_Shapes":
						__gong__assignSliceOfPointers(&inst.Stakeholder_Shapes, rhs, identifierMap)
					case "ResourcesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ResourcesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsStakeholdersNodeExpanded":
						inst.IsStakeholdersNodeExpanded = GongExtractBool(rhs)
					case "ResourceComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.ResourceComposition_Shapes, rhs, identifierMap)
					case "StakeholderConcernShapes":
						__gong__assignSliceOfPointers(&inst.StakeholderConcernShapes, rhs, identifierMap)
					case "Requirement_Shapes":
						__gong__assignSliceOfPointers(&inst.Requirement_Shapes, rhs, identifierMap)
					case "RequirementsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.RequirementsWhoseNodeIsExpanded, rhs, identifierMap)
					case "RequirementsWhoseConcernsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.RequirementsWhoseConcernsNodeIsExpanded, rhs, identifierMap)
					case "RequirementsWhoseViewsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.RequirementsWhoseViewsNodeIsExpanded, rhs, identifierMap)
					case "RequirementsWhoseConceptsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.RequirementsWhoseConceptsNodeIsExpanded, rhs, identifierMap)
					case "RequirementConcernShapes":
						__gong__assignSliceOfPointers(&inst.RequirementConcernShapes, rhs, identifierMap)
					case "RequirementViewShapes":
						__gong__assignSliceOfPointers(&inst.RequirementViewShapes, rhs, identifierMap)
					case "RequirementConceptShapes":
						__gong__assignSliceOfPointers(&inst.RequirementConceptShapes, rhs, identifierMap)
					case "View_Shapes":
						__gong__assignSliceOfPointers(&inst.View_Shapes, rhs, identifierMap)
					case "ViewsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ViewsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ViewsWhoseDeliverablesNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ViewsWhoseDeliverablesNodeIsExpanded, rhs, identifierMap)
					case "ViewsWhoseConceptsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ViewsWhoseConceptsNodeIsExpanded, rhs, identifierMap)
					case "ViewsWhoseRequirementsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ViewsWhoseRequirementsNodeIsExpanded, rhs, identifierMap)
					case "DeliverableViewShapes":
						__gong__assignSliceOfPointers(&inst.DeliverableViewShapes, rhs, identifierMap)
					case "ViewConceptShapes":
						__gong__assignSliceOfPointers(&inst.ViewConceptShapes, rhs, identifierMap)
					case "Concept_Shapes":
						__gong__assignSliceOfPointers(&inst.Concept_Shapes, rhs, identifierMap)
					case "ConceptsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConceptsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ConceptsWhoseViewsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConceptsWhoseViewsNodeIsExpanded, rhs, identifierMap)
					case "ConceptsWhoseDeliverablesNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConceptsWhoseDeliverablesNodeIsExpanded, rhs, identifierMap)
					case "ConceptsWhoseRequirementsNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConceptsWhoseRequirementsNodeIsExpanded, rhs, identifierMap)
					case "DeliverableConceptShapes":
						__gong__assignSliceOfPointers(&inst.DeliverableConceptShapes, rhs, identifierMap)
					case "Diagram_Shapes":
						__gong__assignSliceOfPointers(&inst.Diagram_Shapes, rhs, identifierMap)
					case "IsDiagramsNodeExpanded":
						inst.IsDiagramsNodeExpanded = GongExtractBool(rhs)
					case "DiagramsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramsWhoseNodeIsExpanded, rhs, identifierMap)
					}
				case *DiagramShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Diagram":
						__gong__assignPointer(&inst.Diagram, rhs, identifierMap)
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
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "RootDeliverables":
						__gong__assignSliceOfPointers(&inst.RootDeliverables, rhs, identifierMap)
					case "RootConcerns":
						__gong__assignSliceOfPointers(&inst.RootConcerns, rhs, identifierMap)
					case "RootStakeholders":
						__gong__assignSliceOfPointers(&inst.RootStakeholders, rhs, identifierMap)
					case "RootRequirements":
						__gong__assignSliceOfPointers(&inst.RootRequirements, rhs, identifierMap)
					case "RootViews":
						__gong__assignSliceOfPointers(&inst.RootViews, rhs, identifierMap)
					case "RootConcepts":
						__gong__assignSliceOfPointers(&inst.RootConcepts, rhs, identifierMap)
					case "AnalysisNeeds":
						__gong__assignSliceOfPointers(&inst.AnalysisNeeds, rhs, identifierMap)
					case "Notes":
						__gong__assignSliceOfPointers(&inst.Notes, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Deliverables":
						__gong__assignSliceOfPointers(&inst.Deliverables, rhs, identifierMap)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					case "Resources":
						__gong__assignSliceOfPointers(&inst.Resources, rhs, identifierMap)
					}
				case *NoteDeliverableShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Deliverable":
						__gong__assignPointer(&inst.Deliverable, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *NoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
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
				case *NoteStakeholderShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Stakeholder":
						__gong__assignPointer(&inst.Stakeholder, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *Requirement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SupportLevels":
						__gong__assignSliceOfPointers(&inst.SupportLevels, rhs, identifierMap)
					case "Concerns":
						__gong__assignSliceOfPointers(&inst.Concerns, rhs, identifierMap)
					case "Views":
						__gong__assignSliceOfPointers(&inst.Views, rhs, identifierMap)
					case "Concepts":
						__gong__assignSliceOfPointers(&inst.Concepts, rhs, identifierMap)
					}
				case *RequirementConceptShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Requirement":
						__gong__assignPointer(&inst.Requirement, rhs, identifierMap)
					case "Concept":
						__gong__assignPointer(&inst.Concept, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *RequirementConcernShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Requirement":
						__gong__assignPointer(&inst.Requirement, rhs, identifierMap)
					case "Concern":
						__gong__assignPointer(&inst.Concern, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *RequirementShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Requirement":
						__gong__assignPointer(&inst.Requirement, rhs, identifierMap)
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
				case *RequirementViewShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Requirement":
						__gong__assignPointer(&inst.Requirement, rhs, identifierMap)
					case "View":
						__gong__assignPointer(&inst.View, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *Stakeholder:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IDAirbus":
						inst.IDAirbus = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Concerns":
						__gong__assignSliceOfPointers(&inst.Concerns, rhs, identifierMap)
					case "SubStakeholders":
						__gong__assignSliceOfPointers(&inst.SubStakeholders, rhs, identifierMap)
					}
				case *StakeholderCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Stakeholder":
						__gong__assignPointer(&inst.Stakeholder, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *StakeholderConcernShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Stakeholder":
						__gong__assignPointer(&inst.Stakeholder, rhs, identifierMap)
					case "Concern":
						__gong__assignPointer(&inst.Concern, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *StakeholderShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Stakeholder":
						__gong__assignPointer(&inst.Stakeholder, rhs, identifierMap)
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
				case *SupportLevel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tool":
						__gong__assignPointer(&inst.Tool, rhs, identifierMap)
					}
				case *Tool:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *View:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Concepts":
						__gong__assignSliceOfPointers(&inst.Concepts, rhs, identifierMap)
					}
				case *ViewConceptShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "View":
						__gong__assignPointer(&inst.View, rhs, identifierMap)
					case "Concept":
						__gong__assignPointer(&inst.Concept, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *ViewShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "View":
						__gong__assignPointer(&inst.View, rhs, identifierMap)
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
type AnalysisNeedUnmarshaller struct{}

func (u *AnalysisNeedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AnalysisNeed), stage, identifier, instanceName, preserveOrder)
}

func (u *AnalysisNeedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AnalysisNeed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ConceptUnmarshaller struct{}

func (u *ConceptUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Concept), stage, identifier, instanceName, preserveOrder)
}

func (u *ConceptUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Concept)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Tools":
		GongUnmarshallSliceOfPointers(&instance.Tools, valueExpr, identifierMap)
	}
	return nil
}

type ConceptShapeUnmarshaller struct{}

func (u *ConceptShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConceptShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ConceptShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConceptShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Concept":
		GongUnmarshallPointer(&instance.Concept, valueExpr, identifierMap)
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

type ConcernUnmarshaller struct{}

func (u *ConcernUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Concern), stage, identifier, instanceName, preserveOrder)
}

func (u *ConcernUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Concern)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IDAirbus":
		instance.IDAirbus = GongExtractString(valueExpr)
	case "Priority":
		GongUnmarshallEnum(&instance.Priority, valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SubConcerns":
		GongUnmarshallSliceOfPointers(&instance.SubConcerns, valueExpr, identifierMap)
	case "Inputs":
		GongUnmarshallSliceOfPointers(&instance.Inputs, valueExpr, identifierMap)
	case "IsInputsNodeExpanded":
		instance.IsInputsNodeExpanded = GongExtractBool(valueExpr)
	case "Outputs":
		GongUnmarshallSliceOfPointers(&instance.Outputs, valueExpr, identifierMap)
	case "IsOutputsNodeExpanded":
		instance.IsOutputsNodeExpanded = GongExtractBool(valueExpr)
	case "IsWithCompletion":
		instance.IsWithCompletion = GongExtractBool(valueExpr)
	case "Completion":
		GongUnmarshallEnum(&instance.Completion, valueExpr)
	}
	return nil
}

type ConcernCompositionShapeUnmarshaller struct{}

func (u *ConcernCompositionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConcernCompositionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ConcernCompositionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConcernCompositionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Concern":
		GongUnmarshallPointer(&instance.Concern, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type ConcernInputShapeUnmarshaller struct{}

func (u *ConcernInputShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConcernInputShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ConcernInputShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConcernInputShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
	case "Concern":
		GongUnmarshallPointer(&instance.Concern, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type ConcernOutputShapeUnmarshaller struct{}

func (u *ConcernOutputShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConcernOutputShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ConcernOutputShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConcernOutputShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Concern":
		GongUnmarshallPointer(&instance.Concern, valueExpr, identifierMap)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type ConcernShapeUnmarshaller struct{}

func (u *ConcernShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConcernShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ConcernShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConcernShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Concern":
		GongUnmarshallPointer(&instance.Concern, valueExpr, identifierMap)
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

type ControlPointShapeUnmarshaller struct{}

func (u *ControlPointShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlPointShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlPointShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPointShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "IsStartShapeTheClosestShape":
		instance.IsStartShapeTheClosestShape = GongExtractBool(valueExpr)
	}
	return nil
}

type DeliverableUnmarshaller struct{}

func (u *DeliverableUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Deliverable), stage, identifier, instanceName, preserveOrder)
}

func (u *DeliverableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Deliverable)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SubDeliverables":
		GongUnmarshallSliceOfPointers(&instance.SubDeliverables, valueExpr, identifierMap)
	case "IsProducersNodeExpanded":
		instance.IsProducersNodeExpanded = GongExtractBool(valueExpr)
	case "IsConsumersNodeExpanded":
		instance.IsConsumersNodeExpanded = GongExtractBool(valueExpr)
	case "Views":
		GongUnmarshallSliceOfPointers(&instance.Views, valueExpr, identifierMap)
	case "Concepts":
		GongUnmarshallSliceOfPointers(&instance.Concepts, valueExpr, identifierMap)
	}
	return nil
}

type DeliverableCompositionShapeUnmarshaller struct{}

func (u *DeliverableCompositionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeliverableCompositionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DeliverableCompositionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeliverableCompositionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type DeliverableConceptShapeUnmarshaller struct{}

func (u *DeliverableConceptShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeliverableConceptShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DeliverableConceptShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeliverableConceptShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
	case "Concept":
		GongUnmarshallPointer(&instance.Concept, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type DeliverableShapeUnmarshaller struct{}

func (u *DeliverableShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeliverableShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DeliverableShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeliverableShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
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

type DeliverableViewShapeUnmarshaller struct{}

func (u *DeliverableViewShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeliverableViewShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DeliverableViewShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeliverableViewShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
	case "View":
		GongUnmarshallPointer(&instance.View, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

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
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "IsEditable_":
		instance.IsEditable_ = GongExtractBool(valueExpr)
	case "ShowPrefix":
		instance.ShowPrefix = GongExtractBool(valueExpr)
	case "DefaultBoxWidth":
		instance.DefaultBoxWidth = GongExtractFloat(valueExpr)
	case "DefaultBoxHeigth":
		instance.DefaultBoxHeigth = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "ConcernsWhoseRequirementsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConcernsWhoseRequirementsNodeIsExpanded, valueExpr, identifierMap)
	case "IsRequirementsNodeExpanded":
		instance.IsRequirementsNodeExpanded = GongExtractBool(valueExpr)
	case "IsViewsNodeExpanded":
		instance.IsViewsNodeExpanded = GongExtractBool(valueExpr)
	case "IsConceptsNodeExpanded":
		instance.IsConceptsNodeExpanded = GongExtractBool(valueExpr)
	case "Deliverable_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Deliverable_Shapes, valueExpr, identifierMap)
	case "DeliverablesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DeliverablesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "DeliverablesWhoseViewsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DeliverablesWhoseViewsNodeIsExpanded, valueExpr, identifierMap)
	case "DeliverablesWhoseConceptsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DeliverablesWhoseConceptsNodeIsExpanded, valueExpr, identifierMap)
	case "IsPBSNodeExpanded":
		instance.IsPBSNodeExpanded = GongExtractBool(valueExpr)
	case "DeliverableComposition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.DeliverableComposition_Shapes, valueExpr, identifierMap)
	case "IsConcernsNodeExpanded":
		instance.IsConcernsNodeExpanded = GongExtractBool(valueExpr)
	case "Concern_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Concern_Shapes, valueExpr, identifierMap)
	case "ConcernsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConcernsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ConcernsWhoseInputNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConcernsWhoseInputNodeIsExpanded, valueExpr, identifierMap)
	case "ConcernsWhoseStakeholderNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConcernsWhoseStakeholderNodeIsExpanded, valueExpr, identifierMap)
	case "ConcernssWhoseOutputNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConcernssWhoseOutputNodeIsExpanded, valueExpr, identifierMap)
	case "ConcernComposition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ConcernComposition_Shapes, valueExpr, identifierMap)
	case "ConcernInputShapes":
		GongUnmarshallSliceOfPointers(&instance.ConcernInputShapes, valueExpr, identifierMap)
	case "ConcernOutputShapes":
		GongUnmarshallSliceOfPointers(&instance.ConcernOutputShapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NotesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.NotesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsNotesNodeExpanded":
		instance.IsNotesNodeExpanded = GongExtractBool(valueExpr)
	case "NoteDeliverableShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteDeliverableShapes, valueExpr, identifierMap)
	case "NoteTaskShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteTaskShapes, valueExpr, identifierMap)
	case "NoteResourceShapes":
		GongUnmarshallSliceOfPointers(&instance.NoteResourceShapes, valueExpr, identifierMap)
	case "Stakeholder_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Stakeholder_Shapes, valueExpr, identifierMap)
	case "ResourcesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ResourcesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsStakeholdersNodeExpanded":
		instance.IsStakeholdersNodeExpanded = GongExtractBool(valueExpr)
	case "ResourceComposition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.ResourceComposition_Shapes, valueExpr, identifierMap)
	case "StakeholderConcernShapes":
		GongUnmarshallSliceOfPointers(&instance.StakeholderConcernShapes, valueExpr, identifierMap)
	case "Requirement_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Requirement_Shapes, valueExpr, identifierMap)
	case "RequirementsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.RequirementsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "RequirementsWhoseConcernsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.RequirementsWhoseConcernsNodeIsExpanded, valueExpr, identifierMap)
	case "RequirementsWhoseViewsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.RequirementsWhoseViewsNodeIsExpanded, valueExpr, identifierMap)
	case "RequirementsWhoseConceptsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.RequirementsWhoseConceptsNodeIsExpanded, valueExpr, identifierMap)
	case "RequirementConcernShapes":
		GongUnmarshallSliceOfPointers(&instance.RequirementConcernShapes, valueExpr, identifierMap)
	case "RequirementViewShapes":
		GongUnmarshallSliceOfPointers(&instance.RequirementViewShapes, valueExpr, identifierMap)
	case "RequirementConceptShapes":
		GongUnmarshallSliceOfPointers(&instance.RequirementConceptShapes, valueExpr, identifierMap)
	case "View_Shapes":
		GongUnmarshallSliceOfPointers(&instance.View_Shapes, valueExpr, identifierMap)
	case "ViewsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ViewsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ViewsWhoseDeliverablesNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ViewsWhoseDeliverablesNodeIsExpanded, valueExpr, identifierMap)
	case "ViewsWhoseConceptsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ViewsWhoseConceptsNodeIsExpanded, valueExpr, identifierMap)
	case "ViewsWhoseRequirementsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ViewsWhoseRequirementsNodeIsExpanded, valueExpr, identifierMap)
	case "DeliverableViewShapes":
		GongUnmarshallSliceOfPointers(&instance.DeliverableViewShapes, valueExpr, identifierMap)
	case "ViewConceptShapes":
		GongUnmarshallSliceOfPointers(&instance.ViewConceptShapes, valueExpr, identifierMap)
	case "Concept_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Concept_Shapes, valueExpr, identifierMap)
	case "ConceptsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConceptsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "ConceptsWhoseViewsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConceptsWhoseViewsNodeIsExpanded, valueExpr, identifierMap)
	case "ConceptsWhoseDeliverablesNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConceptsWhoseDeliverablesNodeIsExpanded, valueExpr, identifierMap)
	case "ConceptsWhoseRequirementsNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ConceptsWhoseRequirementsNodeIsExpanded, valueExpr, identifierMap)
	case "DeliverableConceptShapes":
		GongUnmarshallSliceOfPointers(&instance.DeliverableConceptShapes, valueExpr, identifierMap)
	case "Diagram_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Diagram_Shapes, valueExpr, identifierMap)
	case "IsDiagramsNodeExpanded":
		instance.IsDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "DiagramsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.DiagramsWhoseNodeIsExpanded, valueExpr, identifierMap)
	}
	return nil
}

type DiagramShapeUnmarshaller struct{}

func (u *DiagramShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DiagramShape), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DiagramShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Diagram":
		GongUnmarshallPointer(&instance.Diagram, valueExpr, identifierMap)
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
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "RootDeliverables":
		GongUnmarshallSliceOfPointers(&instance.RootDeliverables, valueExpr, identifierMap)
	case "RootConcerns":
		GongUnmarshallSliceOfPointers(&instance.RootConcerns, valueExpr, identifierMap)
	case "RootStakeholders":
		GongUnmarshallSliceOfPointers(&instance.RootStakeholders, valueExpr, identifierMap)
	case "RootRequirements":
		GongUnmarshallSliceOfPointers(&instance.RootRequirements, valueExpr, identifierMap)
	case "RootViews":
		GongUnmarshallSliceOfPointers(&instance.RootViews, valueExpr, identifierMap)
	case "RootConcepts":
		GongUnmarshallSliceOfPointers(&instance.RootConcepts, valueExpr, identifierMap)
	case "AnalysisNeeds":
		GongUnmarshallSliceOfPointers(&instance.AnalysisNeeds, valueExpr, identifierMap)
	case "Notes":
		GongUnmarshallSliceOfPointers(&instance.Notes, valueExpr, identifierMap)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
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
	case "Deliverables":
		GongUnmarshallSliceOfPointers(&instance.Deliverables, valueExpr, identifierMap)
	case "Tasks":
		GongUnmarshallSliceOfPointers(&instance.Tasks, valueExpr, identifierMap)
	case "Resources":
		GongUnmarshallSliceOfPointers(&instance.Resources, valueExpr, identifierMap)
	}
	return nil
}

type NoteDeliverableShapeUnmarshaller struct{}

func (u *NoteDeliverableShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteDeliverableShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteDeliverableShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteDeliverableShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Deliverable":
		GongUnmarshallPointer(&instance.Deliverable, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
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

type NoteStakeholderShapeUnmarshaller struct{}

func (u *NoteStakeholderShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteStakeholderShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteStakeholderShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteStakeholderShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "Stakeholder":
		GongUnmarshallPointer(&instance.Stakeholder, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type RequirementUnmarshaller struct{}

func (u *RequirementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Requirement), stage, identifier, instanceName, preserveOrder)
}

func (u *RequirementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Requirement)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SupportLevels":
		GongUnmarshallSliceOfPointers(&instance.SupportLevels, valueExpr, identifierMap)
	case "Concerns":
		GongUnmarshallSliceOfPointers(&instance.Concerns, valueExpr, identifierMap)
	case "Views":
		GongUnmarshallSliceOfPointers(&instance.Views, valueExpr, identifierMap)
	case "Concepts":
		GongUnmarshallSliceOfPointers(&instance.Concepts, valueExpr, identifierMap)
	}
	return nil
}

type RequirementConceptShapeUnmarshaller struct{}

func (u *RequirementConceptShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RequirementConceptShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RequirementConceptShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RequirementConceptShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Requirement":
		GongUnmarshallPointer(&instance.Requirement, valueExpr, identifierMap)
	case "Concept":
		GongUnmarshallPointer(&instance.Concept, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type RequirementConcernShapeUnmarshaller struct{}

func (u *RequirementConcernShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RequirementConcernShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RequirementConcernShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RequirementConcernShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Requirement":
		GongUnmarshallPointer(&instance.Requirement, valueExpr, identifierMap)
	case "Concern":
		GongUnmarshallPointer(&instance.Concern, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type RequirementShapeUnmarshaller struct{}

func (u *RequirementShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RequirementShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RequirementShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RequirementShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Requirement":
		GongUnmarshallPointer(&instance.Requirement, valueExpr, identifierMap)
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

type RequirementViewShapeUnmarshaller struct{}

func (u *RequirementViewShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RequirementViewShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RequirementViewShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RequirementViewShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Requirement":
		GongUnmarshallPointer(&instance.Requirement, valueExpr, identifierMap)
	case "View":
		GongUnmarshallPointer(&instance.View, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type StakeholderUnmarshaller struct{}

func (u *StakeholderUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stakeholder), stage, identifier, instanceName, preserveOrder)
}

func (u *StakeholderUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stakeholder)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IDAirbus":
		instance.IDAirbus = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Concerns":
		GongUnmarshallSliceOfPointers(&instance.Concerns, valueExpr, identifierMap)
	case "SubStakeholders":
		GongUnmarshallSliceOfPointers(&instance.SubStakeholders, valueExpr, identifierMap)
	}
	return nil
}

type StakeholderCompositionShapeUnmarshaller struct{}

func (u *StakeholderCompositionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StakeholderCompositionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StakeholderCompositionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StakeholderCompositionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Stakeholder":
		GongUnmarshallPointer(&instance.Stakeholder, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type StakeholderConcernShapeUnmarshaller struct{}

func (u *StakeholderConcernShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StakeholderConcernShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StakeholderConcernShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StakeholderConcernShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Stakeholder":
		GongUnmarshallPointer(&instance.Stakeholder, valueExpr, identifierMap)
	case "Concern":
		GongUnmarshallPointer(&instance.Concern, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type StakeholderShapeUnmarshaller struct{}

func (u *StakeholderShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StakeholderShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StakeholderShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StakeholderShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Stakeholder":
		GongUnmarshallPointer(&instance.Stakeholder, valueExpr, identifierMap)
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

type SupportLevelUnmarshaller struct{}

func (u *SupportLevelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SupportLevel), stage, identifier, instanceName, preserveOrder)
}

func (u *SupportLevelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SupportLevel)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tool":
		GongUnmarshallPointer(&instance.Tool, valueExpr, identifierMap)
	}
	return nil
}

type ToolUnmarshaller struct{}

func (u *ToolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tool), stage, identifier, instanceName, preserveOrder)
}

func (u *ToolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tool)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ViewUnmarshaller struct{}

func (u *ViewUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(View), stage, identifier, instanceName, preserveOrder)
}

func (u *ViewUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*View)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Concepts":
		GongUnmarshallSliceOfPointers(&instance.Concepts, valueExpr, identifierMap)
	}
	return nil
}

type ViewConceptShapeUnmarshaller struct{}

func (u *ViewConceptShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ViewConceptShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ViewConceptShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ViewConceptShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "View":
		GongUnmarshallPointer(&instance.View, valueExpr, identifierMap)
	case "Concept":
		GongUnmarshallPointer(&instance.Concept, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type ViewShapeUnmarshaller struct{}

func (u *ViewShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ViewShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ViewShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ViewShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "View":
		GongUnmarshallPointer(&instance.View, valueExpr, identifierMap)
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

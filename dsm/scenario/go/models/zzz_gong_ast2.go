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
		case "github.com/fullstack-lang/gong/dsm/scenario/go/models":
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
				case "ActorState":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorState), stage, ident.Name, instanceName, preserveOrder)
				case "ActorStateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorStateShape), stage, ident.Name, instanceName, preserveOrder)
				case "ActorStateTransition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorStateTransition), stage, ident.Name, instanceName, preserveOrder)
				case "ActorStateTransitionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorStateTransitionShape), stage, ident.Name, instanceName, preserveOrder)
				case "Analysis":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Analysis), stage, ident.Name, instanceName, preserveOrder)
				case "ControlPointShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPointShape), stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stage, ident.Name, instanceName, preserveOrder)
				case "Document":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Document), stage, ident.Name, instanceName, preserveOrder)
				case "DocumentUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DocumentUse), stage, ident.Name, instanceName, preserveOrder)
				case "EvolutionDirection":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EvolutionDirection), stage, ident.Name, instanceName, preserveOrder)
				case "EvolutionDirectionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EvolutionDirectionShape), stage, ident.Name, instanceName, preserveOrder)
				case "Foo":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Foo), stage, ident.Name, instanceName, preserveOrder)
				case "GeoObject":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoObject), stage, ident.Name, instanceName, preserveOrder)
				case "GeoObjectUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoObjectUse), stage, ident.Name, instanceName, preserveOrder)
				case "Group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group), stage, ident.Name, instanceName, preserveOrder)
				case "GroupUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroupUse), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "MapObject":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MapObject), stage, ident.Name, instanceName, preserveOrder)
				case "MapObjectUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MapObjectUse), stage, ident.Name, instanceName, preserveOrder)
				case "Parameter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Parameter), stage, ident.Name, instanceName, preserveOrder)
				case "ParameterCategory":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParameterCategory), stage, ident.Name, instanceName, preserveOrder)
				case "ParameterCategoryUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParameterCategoryUse), stage, ident.Name, instanceName, preserveOrder)
				case "ParameterShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParameterShape), stage, ident.Name, instanceName, preserveOrder)
				case "ParametersAggregate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParametersAggregate), stage, ident.Name, instanceName, preserveOrder)
				case "ParametersAggregateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParametersAggregateShape), stage, ident.Name, instanceName, preserveOrder)
				case "Position":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Position), stage, ident.Name, instanceName, preserveOrder)
				case "Repository":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Repository), stage, ident.Name, instanceName, preserveOrder)
				case "Scenario":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scenario), stage, ident.Name, instanceName, preserveOrder)
				case "User":
					identifierMap[ident.Name] = __gong__stageSetInit(new(User), stage, ident.Name, instanceName, preserveOrder)
				case "UserUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UserUse), stage, ident.Name, instanceName, preserveOrder)
				case "Workspace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Workspace), stage, ident.Name, instanceName, preserveOrder)
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
				case *ActorState:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "IsWithProbaility":
						inst.IsWithProbaility = GongExtractBool(rhs)
					case "Probability":
						GongUnmarshallEnum(&inst.Probability, rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ActorStateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ActorState":
						__gong__assignPointer(&inst.ActorState, rhs, identifierMap)
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
				case *ActorStateTransition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StartState":
						__gong__assignPointer(&inst.StartState, rhs, identifierMap)
					case "EndState":
						__gong__assignPointer(&inst.EndState, rhs, identifierMap)
					case "Justifications":
						__gong__assignSliceOfPointers(&inst.Justifications, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ActorStateTransitionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ActorStateTransition":
						__gong__assignPointer(&inst.ActorStateTransition, rhs, identifierMap)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
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
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
					}
				case *Analysis:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Scenarios":
						__gong__assignSliceOfPointers(&inst.Scenarios, rhs, identifierMap)
					case "IsScenariosNodeExpanded":
						inst.IsScenariosNodeExpanded = GongExtractBool(rhs)
					case "GroupUse":
						__gong__assignSliceOfPointers(&inst.GroupUse, rhs, identifierMap)
					case "IsGroupUseNodeExpanded":
						inst.IsGroupUseNodeExpanded = GongExtractBool(rhs)
					case "GeoObjectUse":
						__gong__assignSliceOfPointers(&inst.GeoObjectUse, rhs, identifierMap)
					case "IsGeoObjectUseNodeExpanded":
						inst.IsGeoObjectUseNodeExpanded = GongExtractBool(rhs)
					case "MapUse":
						__gong__assignSliceOfPointers(&inst.MapUse, rhs, identifierMap)
					case "IsMapUseNodeExpanded":
						inst.IsMapUseNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
					case "IsShowPrefix":
						inst.IsShowPrefix = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "EvolutionDirectionShapes":
						__gong__assignSliceOfPointers(&inst.EvolutionDirectionShapes, rhs, identifierMap)
					case "EvolutionDirectionsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EvolutionDirectionsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEvolutionDirectionsNodeExpanded":
						inst.IsEvolutionDirectionsNodeExpanded = GongExtractBool(rhs)
					case "ActorStateShapes":
						__gong__assignSliceOfPointers(&inst.ActorStateShapes, rhs, identifierMap)
					case "ActorStatesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ActorStatesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsActorStatesNodeExpanded":
						inst.IsActorStatesNodeExpanded = GongExtractBool(rhs)
					case "ParameterShapes":
						__gong__assignSliceOfPointers(&inst.ParameterShapes, rhs, identifierMap)
					case "ParametersWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ParametersWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsParametersNodeExpanded":
						inst.IsParametersNodeExpanded = GongExtractBool(rhs)
					case "ScenarioParameterShapes":
						__gong__assignSliceOfPointers(&inst.ScenarioParameterShapes, rhs, identifierMap)
					case "ParametersAggregatesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ParametersAggregatesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsParametersAggregatesNodeExpanded":
						inst.IsParametersAggregatesNodeExpanded = GongExtractBool(rhs)
					case "ActorStateTransitionShapes":
						__gong__assignSliceOfPointers(&inst.ActorStateTransitionShapes, rhs, identifierMap)
					case "ActorStateTransitionsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ActorStateTransitionsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsActorStateTransitionsNodeExpanded":
						inst.IsActorStateTransitionsNodeExpanded = GongExtractBool(rhs)
					case "AxisOrign_X":
						inst.AxisOrign_X = GongExtractFloat(rhs)
					case "AxisOrign_Y":
						inst.AxisOrign_Y = GongExtractFloat(rhs)
					case "VerticalAxis_Top_Y":
						inst.VerticalAxis_Top_Y = GongExtractFloat(rhs)
					case "VerticalAxis_Bottom_Y":
						inst.VerticalAxis_Bottom_Y = GongExtractFloat(rhs)
					case "VerticalAxis_StrokeWidth":
						inst.VerticalAxis_StrokeWidth = GongExtractFloat(rhs)
					case "HorizontalAxis_Right_X":
						inst.HorizontalAxis_Right_X = GongExtractFloat(rhs)
					case "Start":
						inst.Start = GongExtractDate(rhs)
					case "End":
						inst.End = GongExtractDate(rhs)
					case "NumberOfYearsBetweenTicks":
						inst.NumberOfYearsBetweenTicks = GongExtractInt(rhs)
					case "IsInDrawMode":
						inst.IsInDrawMode = GongExtractBool(rhs)
					}
				case *Document:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GeoObjectUse":
						__gong__assignSliceOfPointers(&inst.GeoObjectUse, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *DocumentUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Document":
						__gong__assignPointer(&inst.Document, rhs, identifierMap)
					}
				case *EvolutionDirection:
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
				case *EvolutionDirectionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EvolutionDirection":
						__gong__assignPointer(&inst.EvolutionDirection, rhs, identifierMap)
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
				case *Foo:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *GeoObject:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *GeoObjectUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GeoObject":
						__gong__assignPointer(&inst.GeoObject, rhs, identifierMap)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "UserUse":
						__gong__assignSliceOfPointers(&inst.UserUse, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *GroupUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Group":
						__gong__assignPointer(&inst.Group, rhs, identifierMap)
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
					case "Analyses":
						__gong__assignSliceOfPointers(&inst.Analyses, rhs, identifierMap)
					case "IsAnalysesNodeExpanded":
						inst.IsAnalysesNodeExpanded = GongExtractBool(rhs)
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
					case "IsExpandedTmp":
						inst.IsExpandedTmp = GongExtractBool(rhs)
					}
				case *MapObject:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *MapObjectUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Map":
						__gong__assignPointer(&inst.Map, rhs, identifierMap)
					}
				case *Parameter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "IsResponse":
						inst.IsResponse = GongExtractBool(rhs)
					case "Start":
						inst.Start = GongExtractDate(rhs)
					case "End":
						inst.End = GongExtractDate(rhs)
					case "Force":
						inst.Force = GongExtractFloat(rhs)
					case "GroupUse":
						__gong__assignSliceOfPointers(&inst.GroupUse, rhs, identifierMap)
					case "DocumentUse":
						__gong__assignSliceOfPointers(&inst.DocumentUse, rhs, identifierMap)
					case "GeoObjectUse":
						__gong__assignSliceOfPointers(&inst.GeoObjectUse, rhs, identifierMap)
					case "Tag":
						inst.Tag = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ParameterCategory:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ParameterUse":
						__gong__assignSliceOfPointers(&inst.ParameterUse, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ParameterCategoryUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ParameterCategory":
						__gong__assignPointer(&inst.ParameterCategory, rhs, identifierMap)
					}
				case *ParameterShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parameter":
						__gong__assignPointer(&inst.Parameter, rhs, identifierMap)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
					case "ShapeIsComputedFromModel":
						inst.ShapeIsComputedFromModel = GongExtractBool(rhs)
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
				case *ParametersAggregate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tag":
						inst.Tag = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Parameters":
						__gong__assignSliceOfPointers(&inst.Parameters, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ParametersAggregateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ScenarioParameter":
						__gong__assignPointer(&inst.ScenarioParameter, rhs, identifierMap)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
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
				case *Position:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						inst.Date = GongExtractDate(rhs)
					case "Ordinate":
						inst.Ordinate = GongExtractFloat(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Repository:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ParameterUse":
						__gong__assignSliceOfPointers(&inst.ParameterUse, rhs, identifierMap)
					case "GroupUse":
						__gong__assignSliceOfPointers(&inst.GroupUse, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Scenario:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "IsDiagramsNodeExpanded":
						inst.IsDiagramsNodeExpanded = GongExtractBool(rhs)
					case "ActorStates":
						__gong__assignSliceOfPointers(&inst.ActorStates, rhs, identifierMap)
					case "IsActorStatesNodeExpanded":
						inst.IsActorStatesNodeExpanded = GongExtractBool(rhs)
					case "ActorStateTransitions":
						__gong__assignSliceOfPointers(&inst.ActorStateTransitions, rhs, identifierMap)
					case "IsActorStateTransitionsNodeExpanded":
						inst.IsActorStateTransitionsNodeExpanded = GongExtractBool(rhs)
					case "EvolutionDirections":
						__gong__assignSliceOfPointers(&inst.EvolutionDirections, rhs, identifierMap)
					case "IsEvolutionDirectionsNodeExpanded":
						inst.IsEvolutionDirectionsNodeExpanded = GongExtractBool(rhs)
					case "Parameters":
						__gong__assignSliceOfPointers(&inst.Parameters, rhs, identifierMap)
					case "IsParametersNodeExpanded":
						inst.IsParametersNodeExpanded = GongExtractBool(rhs)
					case "ParametersAggretates":
						__gong__assignSliceOfPointers(&inst.ParametersAggretates, rhs, identifierMap)
					case "IsParametersAggretatesNodeExpanded":
						inst.IsParametersAggretatesNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *User:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *UserUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "User":
						__gong__assignPointer(&inst.User, rhs, identifierMap)
					}
				case *Workspace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SelectedDiagram":
						__gong__assignPointer(&inst.SelectedDiagram, rhs, identifierMap)
					case "Default_EvolutionDirectionShape":
						__gong__assignPointer(&inst.Default_EvolutionDirectionShape, rhs, identifierMap)
					case "Default_ParameterShape":
						__gong__assignPointer(&inst.Default_ParameterShape, rhs, identifierMap)
					case "Default_ScenarioParameterShape":
						__gong__assignPointer(&inst.Default_ScenarioParameterShape, rhs, identifierMap)
					case "Default_ActorStateShape":
						__gong__assignPointer(&inst.Default_ActorStateShape, rhs, identifierMap)
					case "Default_ActorStateTransitionShape":
						__gong__assignPointer(&inst.Default_ActorStateTransitionShape, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
type ActorStateUnmarshaller struct{}

func (u *ActorStateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorState), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorState)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "IsWithProbaility":
		instance.IsWithProbaility = GongExtractBool(valueExpr)
	case "Probability":
		GongUnmarshallEnum(&instance.Probability, valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ActorStateShapeUnmarshaller struct{}

func (u *ActorStateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorStateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorStateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ActorState":
		GongUnmarshallPointer(&instance.ActorState, valueExpr, identifierMap)
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

type ActorStateTransitionUnmarshaller struct{}

func (u *ActorStateTransitionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorStateTransition), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateTransitionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorStateTransition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartState":
		GongUnmarshallPointer(&instance.StartState, valueExpr, identifierMap)
	case "EndState":
		GongUnmarshallPointer(&instance.EndState, valueExpr, identifierMap)
	case "Justifications":
		GongUnmarshallSliceOfPointers(&instance.Justifications, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ActorStateTransitionShapeUnmarshaller struct{}

func (u *ActorStateTransitionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ActorStateTransitionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ActorStateTransitionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ActorStateTransitionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ActorStateTransition":
		GongUnmarshallPointer(&instance.ActorStateTransition, valueExpr, identifierMap)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
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
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type AnalysisUnmarshaller struct{}

func (u *AnalysisUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Analysis), stage, identifier, instanceName, preserveOrder)
}

func (u *AnalysisUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Analysis)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Scenarios":
		GongUnmarshallSliceOfPointers(&instance.Scenarios, valueExpr, identifierMap)
	case "IsScenariosNodeExpanded":
		instance.IsScenariosNodeExpanded = GongExtractBool(valueExpr)
	case "GroupUse":
		GongUnmarshallSliceOfPointers(&instance.GroupUse, valueExpr, identifierMap)
	case "IsGroupUseNodeExpanded":
		instance.IsGroupUseNodeExpanded = GongExtractBool(valueExpr)
	case "GeoObjectUse":
		GongUnmarshallSliceOfPointers(&instance.GeoObjectUse, valueExpr, identifierMap)
	case "IsGeoObjectUseNodeExpanded":
		instance.IsGeoObjectUseNodeExpanded = GongExtractBool(valueExpr)
	case "MapUse":
		GongUnmarshallSliceOfPointers(&instance.MapUse, valueExpr, identifierMap)
	case "IsMapUseNodeExpanded":
		instance.IsMapUseNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
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
	case "IsShowPrefix":
		instance.IsShowPrefix = GongExtractBool(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "EvolutionDirectionShapes":
		GongUnmarshallSliceOfPointers(&instance.EvolutionDirectionShapes, valueExpr, identifierMap)
	case "EvolutionDirectionsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.EvolutionDirectionsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsEvolutionDirectionsNodeExpanded":
		instance.IsEvolutionDirectionsNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStateShapes":
		GongUnmarshallSliceOfPointers(&instance.ActorStateShapes, valueExpr, identifierMap)
	case "ActorStatesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ActorStatesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsActorStatesNodeExpanded":
		instance.IsActorStatesNodeExpanded = GongExtractBool(valueExpr)
	case "ParameterShapes":
		GongUnmarshallSliceOfPointers(&instance.ParameterShapes, valueExpr, identifierMap)
	case "ParametersWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParametersWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsParametersNodeExpanded":
		instance.IsParametersNodeExpanded = GongExtractBool(valueExpr)
	case "ScenarioParameterShapes":
		GongUnmarshallSliceOfPointers(&instance.ScenarioParameterShapes, valueExpr, identifierMap)
	case "ParametersAggregatesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ParametersAggregatesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsParametersAggregatesNodeExpanded":
		instance.IsParametersAggregatesNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStateTransitionShapes":
		GongUnmarshallSliceOfPointers(&instance.ActorStateTransitionShapes, valueExpr, identifierMap)
	case "ActorStateTransitionsWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.ActorStateTransitionsWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsActorStateTransitionsNodeExpanded":
		instance.IsActorStateTransitionsNodeExpanded = GongExtractBool(valueExpr)
	case "AxisOrign_X":
		instance.AxisOrign_X = GongExtractFloat(valueExpr)
	case "AxisOrign_Y":
		instance.AxisOrign_Y = GongExtractFloat(valueExpr)
	case "VerticalAxis_Top_Y":
		instance.VerticalAxis_Top_Y = GongExtractFloat(valueExpr)
	case "VerticalAxis_Bottom_Y":
		instance.VerticalAxis_Bottom_Y = GongExtractFloat(valueExpr)
	case "VerticalAxis_StrokeWidth":
		instance.VerticalAxis_StrokeWidth = GongExtractFloat(valueExpr)
	case "HorizontalAxis_Right_X":
		instance.HorizontalAxis_Right_X = GongExtractFloat(valueExpr)
	case "Start":
		instance.Start = GongExtractDate(valueExpr)
	case "End":
		instance.End = GongExtractDate(valueExpr)
	case "NumberOfYearsBetweenTicks":
		instance.NumberOfYearsBetweenTicks = GongExtractInt(valueExpr)
	}
	return nil
}

type DocumentUnmarshaller struct{}

func (u *DocumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Document), stage, identifier, instanceName, preserveOrder)
}

func (u *DocumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Document)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GeoObjectUse":
		GongUnmarshallSliceOfPointers(&instance.GeoObjectUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type DocumentUseUnmarshaller struct{}

func (u *DocumentUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DocumentUse), stage, identifier, instanceName, preserveOrder)
}

func (u *DocumentUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DocumentUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Document":
		GongUnmarshallPointer(&instance.Document, valueExpr, identifierMap)
	}
	return nil
}

type EvolutionDirectionUnmarshaller struct{}

func (u *EvolutionDirectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EvolutionDirection), stage, identifier, instanceName, preserveOrder)
}

func (u *EvolutionDirectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EvolutionDirection)
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

type EvolutionDirectionShapeUnmarshaller struct{}

func (u *EvolutionDirectionShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EvolutionDirectionShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EvolutionDirectionShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EvolutionDirectionShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EvolutionDirection":
		GongUnmarshallPointer(&instance.EvolutionDirection, valueExpr, identifierMap)
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

type FooUnmarshaller struct{}

func (u *FooUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Foo), stage, identifier, instanceName, preserveOrder)
}

func (u *FooUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Foo)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GeoObjectUnmarshaller struct{}

func (u *GeoObjectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoObject), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoObjectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoObject)
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

type GeoObjectUseUnmarshaller struct{}

func (u *GeoObjectUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoObjectUse), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoObjectUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoObjectUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GeoObject":
		GongUnmarshallPointer(&instance.GeoObject, valueExpr, identifierMap)
	}
	return nil
}

type GroupUnmarshaller struct{}

func (u *GroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "UserUse":
		GongUnmarshallSliceOfPointers(&instance.UserUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type GroupUseUnmarshaller struct{}

func (u *GroupUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroupUse), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroupUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Group":
		GongUnmarshallPointer(&instance.Group, valueExpr, identifierMap)
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
	case "Analyses":
		GongUnmarshallSliceOfPointers(&instance.Analyses, valueExpr, identifierMap)
	case "IsAnalysesNodeExpanded":
		instance.IsAnalysesNodeExpanded = GongExtractBool(valueExpr)
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
	case "IsExpandedTmp":
		instance.IsExpandedTmp = GongExtractBool(valueExpr)
	}
	return nil
}

type MapObjectUnmarshaller struct{}

func (u *MapObjectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MapObject), stage, identifier, instanceName, preserveOrder)
}

func (u *MapObjectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MapObject)
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

type MapObjectUseUnmarshaller struct{}

func (u *MapObjectUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MapObjectUse), stage, identifier, instanceName, preserveOrder)
}

func (u *MapObjectUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MapObjectUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Map":
		GongUnmarshallPointer(&instance.Map, valueExpr, identifierMap)
	}
	return nil
}

type ParameterUnmarshaller struct{}

func (u *ParameterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Parameter), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Parameter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "IsResponse":
		instance.IsResponse = GongExtractBool(valueExpr)
	case "Start":
		instance.Start = GongExtractDate(valueExpr)
	case "End":
		instance.End = GongExtractDate(valueExpr)
	case "Force":
		instance.Force = GongExtractFloat(valueExpr)
	case "GroupUse":
		GongUnmarshallSliceOfPointers(&instance.GroupUse, valueExpr, identifierMap)
	case "DocumentUse":
		GongUnmarshallSliceOfPointers(&instance.DocumentUse, valueExpr, identifierMap)
	case "GeoObjectUse":
		GongUnmarshallSliceOfPointers(&instance.GeoObjectUse, valueExpr, identifierMap)
	case "Tag":
		instance.Tag = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ParameterCategoryUnmarshaller struct{}

func (u *ParameterCategoryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParameterCategory), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterCategoryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParameterCategory)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ParameterUse":
		GongUnmarshallSliceOfPointers(&instance.ParameterUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ParameterCategoryUseUnmarshaller struct{}

func (u *ParameterCategoryUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParameterCategoryUse), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterCategoryUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParameterCategoryUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ParameterCategory":
		GongUnmarshallPointer(&instance.ParameterCategory, valueExpr, identifierMap)
	}
	return nil
}

type ParameterShapeUnmarshaller struct{}

func (u *ParameterShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParameterShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParameterShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParameterShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parameter":
		GongUnmarshallPointer(&instance.Parameter, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "ShapeIsComputedFromModel":
		instance.ShapeIsComputedFromModel = GongExtractBool(valueExpr)
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

type ParametersAggregateUnmarshaller struct{}

func (u *ParametersAggregateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParametersAggregate), stage, identifier, instanceName, preserveOrder)
}

func (u *ParametersAggregateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParametersAggregate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tag":
		instance.Tag = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Parameters":
		GongUnmarshallSliceOfPointers(&instance.Parameters, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ParametersAggregateShapeUnmarshaller struct{}

func (u *ParametersAggregateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParametersAggregateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParametersAggregateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParametersAggregateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ScenarioParameter":
		GongUnmarshallPointer(&instance.ScenarioParameter, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
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

type PositionUnmarshaller struct{}

func (u *PositionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Position), stage, identifier, instanceName, preserveOrder)
}

func (u *PositionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Position)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	case "Ordinate":
		instance.Ordinate = GongExtractFloat(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type RepositoryUnmarshaller struct{}

func (u *RepositoryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Repository), stage, identifier, instanceName, preserveOrder)
}

func (u *RepositoryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Repository)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ParameterUse":
		GongUnmarshallSliceOfPointers(&instance.ParameterUse, valueExpr, identifierMap)
	case "GroupUse":
		GongUnmarshallSliceOfPointers(&instance.GroupUse, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ScenarioUnmarshaller struct{}

func (u *ScenarioUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Scenario), stage, identifier, instanceName, preserveOrder)
}

func (u *ScenarioUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scenario)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "IsDiagramsNodeExpanded":
		instance.IsDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStates":
		GongUnmarshallSliceOfPointers(&instance.ActorStates, valueExpr, identifierMap)
	case "IsActorStatesNodeExpanded":
		instance.IsActorStatesNodeExpanded = GongExtractBool(valueExpr)
	case "ActorStateTransitions":
		GongUnmarshallSliceOfPointers(&instance.ActorStateTransitions, valueExpr, identifierMap)
	case "IsActorStateTransitionsNodeExpanded":
		instance.IsActorStateTransitionsNodeExpanded = GongExtractBool(valueExpr)
	case "EvolutionDirections":
		GongUnmarshallSliceOfPointers(&instance.EvolutionDirections, valueExpr, identifierMap)
	case "IsEvolutionDirectionsNodeExpanded":
		instance.IsEvolutionDirectionsNodeExpanded = GongExtractBool(valueExpr)
	case "Parameters":
		GongUnmarshallSliceOfPointers(&instance.Parameters, valueExpr, identifierMap)
	case "IsParametersNodeExpanded":
		instance.IsParametersNodeExpanded = GongExtractBool(valueExpr)
	case "ParametersAggretates":
		GongUnmarshallSliceOfPointers(&instance.ParametersAggretates, valueExpr, identifierMap)
	case "IsParametersAggretatesNodeExpanded":
		instance.IsParametersAggretatesNodeExpanded = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type UserUnmarshaller struct{}

func (u *UserUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(User), stage, identifier, instanceName, preserveOrder)
}

func (u *UserUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*User)
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

type UserUseUnmarshaller struct{}

func (u *UserUseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UserUse), stage, identifier, instanceName, preserveOrder)
}

func (u *UserUseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UserUse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "User":
		GongUnmarshallPointer(&instance.User, valueExpr, identifierMap)
	}
	return nil
}

type WorkspaceUnmarshaller struct{}

func (u *WorkspaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Workspace), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkspaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Workspace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SelectedDiagram":
		GongUnmarshallPointer(&instance.SelectedDiagram, valueExpr, identifierMap)
	case "Default_EvolutionDirectionShape":
		GongUnmarshallPointer(&instance.Default_EvolutionDirectionShape, valueExpr, identifierMap)
	case "Default_ParameterShape":
		GongUnmarshallPointer(&instance.Default_ParameterShape, valueExpr, identifierMap)
	case "Default_ScenarioParameterShape":
		GongUnmarshallPointer(&instance.Default_ScenarioParameterShape, valueExpr, identifierMap)
	case "Default_ActorStateShape":
		GongUnmarshallPointer(&instance.Default_ActorStateShape, valueExpr, identifierMap)
	case "Default_ActorStateTransitionShape":
		GongUnmarshallPointer(&instance.Default_ActorStateTransitionShape, valueExpr, identifierMap)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

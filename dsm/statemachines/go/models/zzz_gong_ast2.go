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
		case "github.com/fullstack-lang/gong/dsm/statemachines/go/models":
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
				case "Action":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Action), stage, ident.Name, instanceName, preserveOrder)
				case "Activities":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Activities), stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stage, ident.Name, instanceName, preserveOrder)
				case "Guard":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Guard), stage, ident.Name, instanceName, preserveOrder)
				case "Kill":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Kill), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Message":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Message), stage, ident.Name, instanceName, preserveOrder)
				case "MessageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MessageType), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stage, ident.Name, instanceName, preserveOrder)
				case "NoteStateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteStateShape), stage, ident.Name, instanceName, preserveOrder)
				case "Object":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Object), stage, ident.Name, instanceName, preserveOrder)
				case "Role":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Role), stage, ident.Name, instanceName, preserveOrder)
				case "State":
					identifierMap[ident.Name] = __gong__stageSetInit(new(State), stage, ident.Name, instanceName, preserveOrder)
				case "StateMachine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StateMachine), stage, ident.Name, instanceName, preserveOrder)
				case "StateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StateShape), stage, ident.Name, instanceName, preserveOrder)
				case "Transition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transition), stage, ident.Name, instanceName, preserveOrder)
				case "Transition_Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transition_Shape), stage, ident.Name, instanceName, preserveOrder)
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
				case *Action:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Criticality":
						GongUnmarshallEnum(&inst.Criticality, rhs)
					}
				case *Activities:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Criticality":
						GongUnmarshallEnum(&inst.Criticality, rhs)
					}
				case *Diagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsEditable_":
						inst.IsEditable_ = GongExtractBool(rhs)
					case "IsStatesNodeExpanded":
						inst.IsStatesNodeExpanded = GongExtractBool(rhs)
					case "State_Shapes":
						__gong__assignSliceOfPointers(&inst.State_Shapes, rhs, identifierMap)
					case "StatesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.StatesWhoseNodeIsExpanded, rhs, identifierMap)
					case "Transition_Shapes":
						__gong__assignSliceOfPointers(&inst.Transition_Shapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NoteState_Shapes":
						__gong__assignSliceOfPointers(&inst.NoteState_Shapes, rhs, identifierMap)
					case "ShowRoles":
						inst.ShowRoles = GongExtractBool(rhs)
					case "ShowMessages":
						inst.ShowMessages = GongExtractBool(rhs)
					}
				case *Guard:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Kill:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "RootStateMachines":
						__gong__assignSliceOfPointers(&inst.RootStateMachines, rhs, identifierMap)
					case "IsStateMachinesNodeExpanded":
						inst.IsStateMachinesNodeExpanded = GongExtractBool(rhs)
					case "StateMachinesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.StateMachinesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SubLibrariesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsExpandedTmp":
						inst.IsExpandedTmp = GongExtractBool(rhs)
					case "Roles":
						__gong__assignSliceOfPointers(&inst.Roles, rhs, identifierMap)
					case "IsRolesNodeExpanded":
						inst.IsRolesNodeExpanded = GongExtractBool(rhs)
					case "MessageTypes":
						__gong__assignSliceOfPointers(&inst.MessageTypes, rhs, identifierMap)
					case "IsMessageTypesNodeExpanded":
						inst.IsMessageTypesNodeExpanded = GongExtractBool(rhs)
					}
				case *Message:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "MessageType":
						__gong__assignPointer(&inst.MessageType, rhs, identifierMap)
					case "OriginTransition":
						__gong__assignPointer(&inst.OriginTransition, rhs, identifierMap)
					}
				case *MessageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
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
				case *NoteStateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
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
				case *Object:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					case "DOF":
						inst.DOF = GongExtractDate(rhs)
					case "Messages":
						__gong__assignSliceOfPointers(&inst.Messages, rhs, identifierMap)
					}
				case *Role:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Acronym":
						inst.Acronym = GongExtractString(rhs)
					case "RolesWithSamePermissions":
						__gong__assignSliceOfPointers(&inst.RolesWithSamePermissions, rhs, identifierMap)
					}
				case *State:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsEndState":
						inst.IsEndState = GongExtractBool(rhs)
					case "IsDecisionNode":
						inst.IsDecisionNode = GongExtractBool(rhs)
					case "SubStates":
						__gong__assignSliceOfPointers(&inst.SubStates, rhs, identifierMap)
					case "Entry":
						__gong__assignPointer(&inst.Entry, rhs, identifierMap)
					case "Activities":
						__gong__assignSliceOfPointers(&inst.Activities, rhs, identifierMap)
					case "Exit":
						__gong__assignPointer(&inst.Exit, rhs, identifierMap)
					case "Parent":
						__gong__assignPointer(&inst.Parent, rhs, identifierMap)
					case "IsFictious":
						inst.IsFictious = GongExtractBool(rhs)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "Notes":
						__gong__assignSliceOfPointers(&inst.Notes, rhs, identifierMap)
					}
				case *StateMachine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "InitialState":
						__gong__assignPointer(&inst.InitialState, rhs, identifierMap)
					case "States":
						__gong__assignSliceOfPointers(&inst.States, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "IsWithTransitionNameAutonamticalyGenerated":
						inst.IsWithTransitionNameAutonamticalyGenerated = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *StateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "State":
						__gong__assignPointer(&inst.State, rhs, identifierMap)
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
				case *Transition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "RolesWithPermissions":
						__gong__assignSliceOfPointers(&inst.RolesWithPermissions, rhs, identifierMap)
					case "GeneratedMessages":
						__gong__assignSliceOfPointers(&inst.GeneratedMessages, rhs, identifierMap)
					case "Guard":
						__gong__assignPointer(&inst.Guard, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRolesNodeExpanded":
						inst.IsRolesNodeExpanded = GongExtractBool(rhs)
					case "IsMessagesNodeExpanded":
						inst.IsMessagesNodeExpanded = GongExtractBool(rhs)
					}
				case *Transition_Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Transition":
						__gong__assignPointer(&inst.Transition, rhs, identifierMap)
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
type ActionUnmarshaller struct{}

func (u *ActionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Action), stage, identifier, instanceName, preserveOrder)
}

func (u *ActionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Action)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Criticality":
		GongUnmarshallEnum(&instance.Criticality, valueExpr)
	}
	return nil
}

type ActivitiesUnmarshaller struct{}

func (u *ActivitiesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Activities), stage, identifier, instanceName, preserveOrder)
}

func (u *ActivitiesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Activities)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Criticality":
		GongUnmarshallEnum(&instance.Criticality, valueExpr)
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
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsEditable_":
		instance.IsEditable_ = GongExtractBool(valueExpr)
	case "IsStatesNodeExpanded":
		instance.IsStatesNodeExpanded = GongExtractBool(valueExpr)
	case "State_Shapes":
		GongUnmarshallSliceOfPointers(&instance.State_Shapes, valueExpr, identifierMap)
	case "StatesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.StatesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "Transition_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Transition_Shapes, valueExpr, identifierMap)
	case "Note_Shapes":
		GongUnmarshallSliceOfPointers(&instance.Note_Shapes, valueExpr, identifierMap)
	case "NoteState_Shapes":
		GongUnmarshallSliceOfPointers(&instance.NoteState_Shapes, valueExpr, identifierMap)
	case "ShowRoles":
		instance.ShowRoles = GongExtractBool(valueExpr)
	case "ShowMessages":
		instance.ShowMessages = GongExtractBool(valueExpr)
	}
	return nil
}

type GuardUnmarshaller struct{}

func (u *GuardUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Guard), stage, identifier, instanceName, preserveOrder)
}

func (u *GuardUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Guard)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type KillUnmarshaller struct{}

func (u *KillUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Kill), stage, identifier, instanceName, preserveOrder)
}

func (u *KillUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Kill)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "RootStateMachines":
		GongUnmarshallSliceOfPointers(&instance.RootStateMachines, valueExpr, identifierMap)
	case "IsStateMachinesNodeExpanded":
		instance.IsStateMachinesNodeExpanded = GongExtractBool(valueExpr)
	case "StateMachinesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.StateMachinesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsSubLibrariesNodeExpanded":
		instance.IsSubLibrariesNodeExpanded = GongExtractBool(valueExpr)
	case "SubLibrariesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SubLibrariesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "IsExpandedTmp":
		instance.IsExpandedTmp = GongExtractBool(valueExpr)
	case "Roles":
		GongUnmarshallSliceOfPointers(&instance.Roles, valueExpr, identifierMap)
	case "IsRolesNodeExpanded":
		instance.IsRolesNodeExpanded = GongExtractBool(valueExpr)
	case "MessageTypes":
		GongUnmarshallSliceOfPointers(&instance.MessageTypes, valueExpr, identifierMap)
	case "IsMessageTypesNodeExpanded":
		instance.IsMessageTypesNodeExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type MessageUnmarshaller struct{}

func (u *MessageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Message), stage, identifier, instanceName, preserveOrder)
}

func (u *MessageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Message)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "MessageType":
		GongUnmarshallPointer(&instance.MessageType, valueExpr, identifierMap)
	case "OriginTransition":
		GongUnmarshallPointer(&instance.OriginTransition, valueExpr, identifierMap)
	}
	return nil
}

type MessageTypeUnmarshaller struct{}

func (u *MessageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MessageType), stage, identifier, instanceName, preserveOrder)
}

func (u *MessageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MessageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
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
	case "State":
		GongUnmarshallPointer(&instance.State, valueExpr, identifierMap)
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

type NoteStateShapeUnmarshaller struct{}

func (u *NoteStateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteStateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteStateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteStateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	case "State":
		GongUnmarshallPointer(&instance.State, valueExpr, identifierMap)
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

type ObjectUnmarshaller struct{}

func (u *ObjectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Object), stage, identifier, instanceName, preserveOrder)
}

func (u *ObjectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Object)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "State":
		GongUnmarshallPointer(&instance.State, valueExpr, identifierMap)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	case "DOF":
		instance.DOF = GongExtractDate(valueExpr)
	case "Messages":
		GongUnmarshallSliceOfPointers(&instance.Messages, valueExpr, identifierMap)
	}
	return nil
}

type RoleUnmarshaller struct{}

func (u *RoleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Role), stage, identifier, instanceName, preserveOrder)
}

func (u *RoleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Role)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Acronym":
		instance.Acronym = GongExtractString(valueExpr)
	case "RolesWithSamePermissions":
		GongUnmarshallSliceOfPointers(&instance.RolesWithSamePermissions, valueExpr, identifierMap)
	}
	return nil
}

type StateUnmarshaller struct{}

func (u *StateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(State), stage, identifier, instanceName, preserveOrder)
}

func (u *StateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*State)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsEndState":
		instance.IsEndState = GongExtractBool(valueExpr)
	case "IsDecisionNode":
		instance.IsDecisionNode = GongExtractBool(valueExpr)
	case "SubStates":
		GongUnmarshallSliceOfPointers(&instance.SubStates, valueExpr, identifierMap)
	case "Entry":
		GongUnmarshallPointer(&instance.Entry, valueExpr, identifierMap)
	case "Activities":
		GongUnmarshallSliceOfPointers(&instance.Activities, valueExpr, identifierMap)
	case "Exit":
		GongUnmarshallPointer(&instance.Exit, valueExpr, identifierMap)
	case "Parent":
		GongUnmarshallPointer(&instance.Parent, valueExpr, identifierMap)
	case "IsFictious":
		instance.IsFictious = GongExtractBool(valueExpr)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "Notes":
		GongUnmarshallSliceOfPointers(&instance.Notes, valueExpr, identifierMap)
	}
	return nil
}

type StateMachineUnmarshaller struct{}

func (u *StateMachineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StateMachine), stage, identifier, instanceName, preserveOrder)
}

func (u *StateMachineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StateMachine)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "InitialState":
		GongUnmarshallPointer(&instance.InitialState, valueExpr, identifierMap)
	case "States":
		GongUnmarshallSliceOfPointers(&instance.States, valueExpr, identifierMap)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "IsWithTransitionNameAutonamticalyGenerated":
		instance.IsWithTransitionNameAutonamticalyGenerated = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type StateShapeUnmarshaller struct{}

func (u *StateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "State":
		GongUnmarshallPointer(&instance.State, valueExpr, identifierMap)
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

type TransitionUnmarshaller struct{}

func (u *TransitionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Transition), stage, identifier, instanceName, preserveOrder)
}

func (u *TransitionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Transition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "RolesWithPermissions":
		GongUnmarshallSliceOfPointers(&instance.RolesWithPermissions, valueExpr, identifierMap)
	case "GeneratedMessages":
		GongUnmarshallSliceOfPointers(&instance.GeneratedMessages, valueExpr, identifierMap)
	case "Guard":
		GongUnmarshallPointer(&instance.Guard, valueExpr, identifierMap)
	case "Diagrams":
		GongUnmarshallSliceOfPointers(&instance.Diagrams, valueExpr, identifierMap)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsRolesNodeExpanded":
		instance.IsRolesNodeExpanded = GongExtractBool(valueExpr)
	case "IsMessagesNodeExpanded":
		instance.IsMessagesNodeExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type Transition_ShapeUnmarshaller struct{}

func (u *Transition_ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Transition_Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *Transition_ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Transition_Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Transition":
		GongUnmarshallPointer(&instance.Transition, valueExpr, identifierMap)
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

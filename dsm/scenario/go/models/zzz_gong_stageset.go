// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, actorstate := range __gong__sortStageSetInstances(stageSet.Stage.ActorStates, stageSet.Stage.ActorState_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstateIdent := "__models" + actorstate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorState{Name: %s}).Stage(stageSet.Stage)", actorstateIdent, __gong__toRawStringLiteral(actorstate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstateIdent, __gong__toRawStringLiteral(actorstate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", actorstateIdent, __gong__toRawStringLiteral(actorstate.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithProbaility = %t", actorstateIdent, actorstate.IsWithProbaility))
			values.WriteString(fmt.Sprintf("\n\t%s.Probability = %s", actorstateIdent, __gong__toRawStringLiteral(string(actorstate.Probability))))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", actorstateIdent, __gong__toRawStringLiteral(actorstate.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", actorstateIdent, actorstate.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, actorstateshape := range __gong__sortStageSetInstances(stageSet.Stage.ActorStateShapes, stageSet.Stage.ActorStateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstateshapeIdent := "__models" + actorstateshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorStateShape{Name: %s}).Stage(stageSet.Stage)", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", actorstateshapeIdent, actorstateshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", actorstateshapeIdent, actorstateshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", actorstateshapeIdent, actorstateshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", actorstateshapeIdent, actorstateshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", actorstateshapeIdent, actorstateshape.IsHidden))
			if actorstateshape.ActorState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + actorstateshape.ActorState.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorState = %s", actorstateshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, actorstatetransition := range __gong__sortStageSetInstances(stageSet.Stage.ActorStateTransitions, stageSet.Stage.ActorStateTransition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstatetransitionIdent := "__models" + actorstatetransition.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorStateTransition{Name: %s}).Stage(stageSet.Stage)", actorstatetransitionIdent, __gong__toRawStringLiteral(actorstatetransition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstatetransitionIdent, __gong__toRawStringLiteral(actorstatetransition.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", actorstatetransitionIdent, __gong__toRawStringLiteral(actorstatetransition.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", actorstatetransitionIdent, actorstatetransition.IsExpanded))
			if actorstatetransition.StartState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + actorstatetransition.StartState.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartState = %s", actorstatetransitionIdent, targetIdent))
			}
			if actorstatetransition.EndState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + actorstatetransition.EndState.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndState = %s", actorstatetransitionIdent, targetIdent))
			}
			for _, elem := range actorstatetransition.Justifications {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Justifications = append(%s.Justifications, %s)", actorstatetransitionIdent, actorstatetransitionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, actorstatetransitionshape := range __gong__sortStageSetInstances(stageSet.Stage.ActorStateTransitionShapes, stageSet.Stage.ActorStateTransitionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstatetransitionshapeIdent := "__models" + actorstatetransitionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorStateTransitionShape{Name: %s}).Stage(stageSet.Stage)", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", actorstatetransitionshapeIdent, actorstatetransitionshape.IsHidden))
			if actorstatetransitionshape.ActorStateTransition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + actorstatetransitionshape.ActorStateTransition.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransition = %s", actorstatetransitionshapeIdent, targetIdent))
			}
			if actorstatetransitionshape.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + actorstatetransitionshape.Start.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", actorstatetransitionshapeIdent, targetIdent))
			}
			if actorstatetransitionshape.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + actorstatetransitionshape.End.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", actorstatetransitionshapeIdent, targetIdent))
			}
			for _, elem := range actorstatetransitionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", actorstatetransitionshapeIdent, actorstatetransitionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, analysis := range __gong__sortStageSetInstances(stageSet.Stage.Analysiss, stageSet.Stage.Analysis_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			analysisIdent := "__models" + analysis.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Analysis{Name: %s}).Stage(stageSet.Stage)", analysisIdent, __gong__toRawStringLiteral(analysis.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", analysisIdent, __gong__toRawStringLiteral(analysis.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", analysisIdent, __gong__toRawStringLiteral(analysis.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsScenariosNodeExpanded = %t", analysisIdent, analysis.IsScenariosNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsGroupUseNodeExpanded = %t", analysisIdent, analysis.IsGroupUseNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsGeoObjectUseNodeExpanded = %t", analysisIdent, analysis.IsGeoObjectUseNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMapUseNodeExpanded = %t", analysisIdent, analysis.IsMapUseNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", analysisIdent, __gong__toRawStringLiteral(analysis.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", analysisIdent, analysis.IsExpanded))
			for _, elem := range analysis.Scenarios {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scenarios = append(%s.Scenarios, %s)", analysisIdent, analysisIdent, targetIdent))
			}
			for _, elem := range analysis.GroupUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupUse = append(%s.GroupUse, %s)", analysisIdent, analysisIdent, targetIdent))
			}
			for _, elem := range analysis.GeoObjectUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObjectUse = append(%s.GeoObjectUse, %s)", analysisIdent, analysisIdent, targetIdent))
			}
			for _, elem := range analysis.MapUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MapUse = append(%s.MapUse, %s)", analysisIdent, analysisIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, controlpointshape := range __gong__sortStageSetInstances(stageSet.Stage.ControlPointShapes, stageSet.Stage.ControlPointShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlpointshapeIdent := "__models" + controlpointshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlPointShape{Name: %s}).Stage(stageSet.Stage)", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointshapeIdent, controlpointshape.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointshapeIdent, controlpointshape.Y_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartShapeTheClosestShape = %t", controlpointshapeIdent, controlpointshape.IsStartShapeTheClosestShape))
		}
	}
	if stageSet.Stage != nil {
		for _, diagram := range __gong__sortStageSetInstances(stageSet.Stage.Diagrams, stageSet.Stage.Diagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramIdent := "__models" + diagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Diagram{Name: %s}).Stage(stageSet.Stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramIdent, diagram.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEvolutionDirectionsNodeExpanded = %t", diagramIdent, diagram.IsEvolutionDirectionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStatesNodeExpanded = %t", diagramIdent, diagram.IsActorStatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersNodeExpanded = %t", diagramIdent, diagram.IsParametersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersAggregatesNodeExpanded = %t", diagramIdent, diagram.IsParametersAggregatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStateTransitionsNodeExpanded = %t", diagramIdent, diagram.IsActorStateTransitionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.AxisOrign_X = %f", diagramIdent, diagram.AxisOrign_X))
			values.WriteString(fmt.Sprintf("\n\t%s.AxisOrign_Y = %f", diagramIdent, diagram.AxisOrign_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalAxis_Top_Y = %f", diagramIdent, diagram.VerticalAxis_Top_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalAxis_Bottom_Y = %f", diagramIdent, diagram.VerticalAxis_Bottom_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalAxis_StrokeWidth = %f", diagramIdent, diagram.VerticalAxis_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.HorizontalAxis_Right_X = %f", diagramIdent, diagram.HorizontalAxis_Right_X))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.NumberOfYearsBetweenTicks = %d", diagramIdent, diagram.NumberOfYearsBetweenTicks))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInDrawMode = %t", diagramIdent, diagram.IsInDrawMode))
			for _, elem := range diagram.EvolutionDirectionShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirectionShapes = append(%s.EvolutionDirectionShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirectionsWhoseNodeIsExpanded = append(%s.EvolutionDirectionsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStateShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateShapes = append(%s.ActorStateShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStatesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStatesWhoseNodeIsExpanded = append(%s.ActorStatesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ParameterShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterShapes = append(%s.ParameterShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ParametersWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParametersWhoseNodeIsExpanded = append(%s.ParametersWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ScenarioParameterShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScenarioParameterShapes = append(%s.ScenarioParameterShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParametersAggregatesWhoseNodeIsExpanded = append(%s.ParametersAggregatesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStateTransitionShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransitionShapes = append(%s.ActorStateTransitionShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransitionsWhoseNodeIsExpanded = append(%s.ActorStateTransitionsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, document := range __gong__sortStageSetInstances(stageSet.Stage.Documents, stageSet.Stage.Document_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentIdent := "__models" + document.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Document{Name: %s}).Stage(stageSet.Stage)", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentIdent, __gong__toRawStringLiteral(document.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", documentIdent, __gong__toRawStringLiteral(document.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", documentIdent, document.IsExpanded))
			for _, elem := range document.GeoObjectUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObjectUse = append(%s.GeoObjectUse, %s)", documentIdent, documentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, documentuse := range __gong__sortStageSetInstances(stageSet.Stage.DocumentUses, stageSet.Stage.DocumentUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentuseIdent := "__models" + documentuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DocumentUse{Name: %s}).Stage(stageSet.Stage)", documentuseIdent, __gong__toRawStringLiteral(documentuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentuseIdent, __gong__toRawStringLiteral(documentuse.Name)))
			if documentuse.Document != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + documentuse.Document.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Document = %s", documentuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, evolutiondirection := range __gong__sortStageSetInstances(stageSet.Stage.EvolutionDirections, stageSet.Stage.EvolutionDirection_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			evolutiondirectionIdent := "__models" + evolutiondirection.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EvolutionDirection{Name: %s}).Stage(stageSet.Stage)", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", evolutiondirectionIdent, evolutiondirection.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, evolutiondirectionshape := range __gong__sortStageSetInstances(stageSet.Stage.EvolutionDirectionShapes, stageSet.Stage.EvolutionDirectionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			evolutiondirectionshapeIdent := "__models" + evolutiondirectionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EvolutionDirectionShape{Name: %s}).Stage(stageSet.Stage)", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", evolutiondirectionshapeIdent, evolutiondirectionshape.IsHidden))
			if evolutiondirectionshape.EvolutionDirection != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + evolutiondirectionshape.EvolutionDirection.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirection = %s", evolutiondirectionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, foo := range __gong__sortStageSetInstances(stageSet.Stage.Foos, stageSet.Stage.Foo_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fooIdent := "__models" + foo.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Foo{Name: %s}).Stage(stageSet.Stage)", fooIdent, __gong__toRawStringLiteral(foo.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fooIdent, __gong__toRawStringLiteral(foo.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, geoobject := range __gong__sortStageSetInstances(stageSet.Stage.GeoObjects, stageSet.Stage.GeoObject_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geoobjectIdent := "__models" + geoobject.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GeoObject{Name: %s}).Stage(stageSet.Stage)", geoobjectIdent, __gong__toRawStringLiteral(geoobject.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geoobjectIdent, __gong__toRawStringLiteral(geoobject.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", geoobjectIdent, __gong__toRawStringLiteral(geoobject.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", geoobjectIdent, geoobject.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, geoobjectuse := range __gong__sortStageSetInstances(stageSet.Stage.GeoObjectUses, stageSet.Stage.GeoObjectUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geoobjectuseIdent := "__models" + geoobjectuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GeoObjectUse{Name: %s}).Stage(stageSet.Stage)", geoobjectuseIdent, __gong__toRawStringLiteral(geoobjectuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geoobjectuseIdent, __gong__toRawStringLiteral(geoobjectuse.Name)))
			if geoobjectuse.GeoObject != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + geoobjectuse.GeoObject.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObject = %s", geoobjectuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, group := range __gong__sortStageSetInstances(stageSet.Stage.Groups, stageSet.Stage.Group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := "__models" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", groupIdent, __gong__toRawStringLiteral(group.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", groupIdent, group.IsExpanded))
			for _, elem := range group.UserUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UserUse = append(%s.UserUse, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, groupuse := range __gong__sortStageSetInstances(stageSet.Stage.GroupUses, stageSet.Stage.GroupUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupuseIdent := "__models" + groupuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GroupUse{Name: %s}).Stage(stageSet.Stage)", groupuseIdent, __gong__toRawStringLiteral(groupuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupuseIdent, __gong__toRawStringLiteral(groupuse.Name)))
			if groupuse.Group != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + groupuse.Group.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group = %s", groupuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, library := range __gong__sortStageSetInstances(stageSet.Stage.Librarys, stageSet.Stage.Library_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			libraryIdent := "__models" + library.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Library{Name: %s}).Stage(stageSet.Stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", libraryIdent, __gong__toRawStringLiteral(library.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAnalysesNodeExpanded = %t", libraryIdent, library.IsAnalysesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			for _, elem := range library.Analyses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Analyses = append(%s.Analyses, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibraries {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, mapobject := range __gong__sortStageSetInstances(stageSet.Stage.MapObjects, stageSet.Stage.MapObject_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mapobjectIdent := "__models" + mapobject.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MapObject{Name: %s}).Stage(stageSet.Stage)", mapobjectIdent, __gong__toRawStringLiteral(mapobject.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mapobjectIdent, __gong__toRawStringLiteral(mapobject.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", mapobjectIdent, __gong__toRawStringLiteral(mapobject.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", mapobjectIdent, mapobject.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, mapobjectuse := range __gong__sortStageSetInstances(stageSet.Stage.MapObjectUses, stageSet.Stage.MapObjectUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mapobjectuseIdent := "__models" + mapobjectuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MapObjectUse{Name: %s}).Stage(stageSet.Stage)", mapobjectuseIdent, __gong__toRawStringLiteral(mapobjectuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mapobjectuseIdent, __gong__toRawStringLiteral(mapobjectuse.Name)))
			if mapobjectuse.Map != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mapobjectuse.Map.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Map = %s", mapobjectuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, parameter := range __gong__sortStageSetInstances(stageSet.Stage.Parameters, stageSet.Stage.Parameter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parameterIdent := "__models" + parameter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Parameter{Name: %s}).Stage(stageSet.Stage)", parameterIdent, __gong__toRawStringLiteral(parameter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parameterIdent, __gong__toRawStringLiteral(parameter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", parameterIdent, __gong__toRawStringLiteral(parameter.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResponse = %t", parameterIdent, parameter.IsResponse))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", parameterIdent, parameter.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", parameterIdent, parameter.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Force = %f", parameterIdent, parameter.Force))
			values.WriteString(fmt.Sprintf("\n\t%s.Tag = %s", parameterIdent, __gong__toRawStringLiteral(parameter.Tag)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", parameterIdent, __gong__toRawStringLiteral(parameter.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", parameterIdent, parameter.IsExpanded))
			for _, elem := range parameter.GroupUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupUse = append(%s.GroupUse, %s)", parameterIdent, parameterIdent, targetIdent))
			}
			for _, elem := range parameter.DocumentUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DocumentUse = append(%s.DocumentUse, %s)", parameterIdent, parameterIdent, targetIdent))
			}
			for _, elem := range parameter.GeoObjectUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObjectUse = append(%s.GeoObjectUse, %s)", parameterIdent, parameterIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, parametercategory := range __gong__sortStageSetInstances(stageSet.Stage.ParameterCategorys, stageSet.Stage.ParameterCategory_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametercategoryIdent := "__models" + parametercategory.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParameterCategory{Name: %s}).Stage(stageSet.Stage)", parametercategoryIdent, __gong__toRawStringLiteral(parametercategory.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametercategoryIdent, __gong__toRawStringLiteral(parametercategory.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", parametercategoryIdent, __gong__toRawStringLiteral(parametercategory.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", parametercategoryIdent, parametercategory.IsExpanded))
			for _, elem := range parametercategory.ParameterUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterUse = append(%s.ParameterUse, %s)", parametercategoryIdent, parametercategoryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, parametercategoryuse := range __gong__sortStageSetInstances(stageSet.Stage.ParameterCategoryUses, stageSet.Stage.ParameterCategoryUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametercategoryuseIdent := "__models" + parametercategoryuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParameterCategoryUse{Name: %s}).Stage(stageSet.Stage)", parametercategoryuseIdent, __gong__toRawStringLiteral(parametercategoryuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametercategoryuseIdent, __gong__toRawStringLiteral(parametercategoryuse.Name)))
			if parametercategoryuse.ParameterCategory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + parametercategoryuse.ParameterCategory.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterCategory = %s", parametercategoryuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, parametershape := range __gong__sortStageSetInstances(stageSet.Stage.ParameterShapes, stageSet.Stage.ParameterShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametershapeIdent := "__models" + parametershape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParameterShape{Name: %s}).Stage(stageSet.Stage)", parametershapeIdent, __gong__toRawStringLiteral(parametershape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", parametershapeIdent, __gong__toRawStringLiteral(string(parametershape.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.ShapeIsComputedFromModel = %t", parametershapeIdent, parametershape.ShapeIsComputedFromModel))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", parametershapeIdent, parametershape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", parametershapeIdent, parametershape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", parametershapeIdent, parametershape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", parametershapeIdent, parametershape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", parametershapeIdent, parametershape.IsHidden))
			if parametershape.Parameter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + parametershape.Parameter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parameter = %s", parametershapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, parametersaggregate := range __gong__sortStageSetInstances(stageSet.Stage.ParametersAggregates, stageSet.Stage.ParametersAggregate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametersaggregateIdent := "__models" + parametersaggregate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParametersAggregate{Name: %s}).Stage(stageSet.Stage)", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tag = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Tag)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", parametersaggregateIdent, parametersaggregate.IsExpanded))
			for _, elem := range parametersaggregate.Parameters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parameters = append(%s.Parameters, %s)", parametersaggregateIdent, parametersaggregateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, parametersaggregateshape := range __gong__sortStageSetInstances(stageSet.Stage.ParametersAggregateShapes, stageSet.Stage.ParametersAggregateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametersaggregateshapeIdent := "__models" + parametersaggregateshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParametersAggregateShape{Name: %s}).Stage(stageSet.Stage)", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(string(parametersaggregateshape.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", parametersaggregateshapeIdent, parametersaggregateshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", parametersaggregateshapeIdent, parametersaggregateshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", parametersaggregateshapeIdent, parametersaggregateshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", parametersaggregateshapeIdent, parametersaggregateshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", parametersaggregateshapeIdent, parametersaggregateshape.IsHidden))
			if parametersaggregateshape.ScenarioParameter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + parametersaggregateshape.ScenarioParameter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScenarioParameter = %s", parametersaggregateshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, position := range __gong__sortStageSetInstances(stageSet.Stage.Positions, stageSet.Stage.Position_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			positionIdent := "__models" + position.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Position{Name: %s}).Stage(stageSet.Stage)", positionIdent, __gong__toRawStringLiteral(position.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", positionIdent, __gong__toRawStringLiteral(position.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", positionIdent, position.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Ordinate = %f", positionIdent, position.Ordinate))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", positionIdent, __gong__toRawStringLiteral(position.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", positionIdent, position.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, repository := range __gong__sortStageSetInstances(stageSet.Stage.Repositorys, stageSet.Stage.Repository_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			repositoryIdent := "__models" + repository.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Repository{Name: %s}).Stage(stageSet.Stage)", repositoryIdent, __gong__toRawStringLiteral(repository.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", repositoryIdent, __gong__toRawStringLiteral(repository.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", repositoryIdent, __gong__toRawStringLiteral(repository.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", repositoryIdent, repository.IsExpanded))
			for _, elem := range repository.ParameterUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterUse = append(%s.ParameterUse, %s)", repositoryIdent, repositoryIdent, targetIdent))
			}
			for _, elem := range repository.GroupUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupUse = append(%s.GroupUse, %s)", repositoryIdent, repositoryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, scenario := range __gong__sortStageSetInstances(stageSet.Stage.Scenarios, stageSet.Stage.Scenario_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scenarioIdent := "__models" + scenario.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Scenario{Name: %s}).Stage(stageSet.Stage)", scenarioIdent, __gong__toRawStringLiteral(scenario.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scenarioIdent, __gong__toRawStringLiteral(scenario.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", scenarioIdent, __gong__toRawStringLiteral(scenario.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDiagramsNodeExpanded = %t", scenarioIdent, scenario.IsDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStatesNodeExpanded = %t", scenarioIdent, scenario.IsActorStatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStateTransitionsNodeExpanded = %t", scenarioIdent, scenario.IsActorStateTransitionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEvolutionDirectionsNodeExpanded = %t", scenarioIdent, scenario.IsEvolutionDirectionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersNodeExpanded = %t", scenarioIdent, scenario.IsParametersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersAggretatesNodeExpanded = %t", scenarioIdent, scenario.IsParametersAggretatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", scenarioIdent, __gong__toRawStringLiteral(scenario.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", scenarioIdent, scenario.IsExpanded))
			for _, elem := range scenario.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.ActorStates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStates = append(%s.ActorStates, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.ActorStateTransitions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransitions = append(%s.ActorStateTransitions, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.EvolutionDirections {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirections = append(%s.EvolutionDirections, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.Parameters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parameters = append(%s.Parameters, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.ParametersAggretates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParametersAggretates = append(%s.ParametersAggretates, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, user := range __gong__sortStageSetInstances(stageSet.Stage.Users, stageSet.Stage.User_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			userIdent := "__models" + user.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.User{Name: %s}).Stage(stageSet.Stage)", userIdent, __gong__toRawStringLiteral(user.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", userIdent, __gong__toRawStringLiteral(user.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", userIdent, __gong__toRawStringLiteral(user.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", userIdent, user.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, useruse := range __gong__sortStageSetInstances(stageSet.Stage.UserUses, stageSet.Stage.UserUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			useruseIdent := "__models" + useruse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.UserUse{Name: %s}).Stage(stageSet.Stage)", useruseIdent, __gong__toRawStringLiteral(useruse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", useruseIdent, __gong__toRawStringLiteral(useruse.Name)))
			if useruse.User != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + useruse.User.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.User = %s", useruseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, workspace := range __gong__sortStageSetInstances(stageSet.Stage.Workspaces, stageSet.Stage.Workspace_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			workspaceIdent := "__models" + workspace.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Workspace{Name: %s}).Stage(stageSet.Stage)", workspaceIdent, __gong__toRawStringLiteral(workspace.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", workspaceIdent, __gong__toRawStringLiteral(workspace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", workspaceIdent, __gong__toRawStringLiteral(workspace.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", workspaceIdent, workspace.IsExpanded))
			if workspace.SelectedDiagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + workspace.SelectedDiagram.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectedDiagram = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_EvolutionDirectionShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + workspace.Default_EvolutionDirectionShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_EvolutionDirectionShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ParameterShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + workspace.Default_ParameterShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ParameterShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ScenarioParameterShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + workspace.Default_ScenarioParameterShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ScenarioParameterShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ActorStateShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + workspace.Default_ActorStateShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ActorStateShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ActorStateTransitionShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + workspace.Default_ActorStateTransitionShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ActorStateTransitionShape = %s", workspaceIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

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
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "ActorState":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorState), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ActorStateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorStateShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ActorStateTransition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorStateTransition), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ActorStateTransitionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ActorStateTransitionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Analysis":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Analysis), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ControlPointShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPointShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Document":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Document), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DocumentUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DocumentUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EvolutionDirection":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EvolutionDirection), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EvolutionDirectionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EvolutionDirectionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Foo":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Foo), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GeoObject":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoObject), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GeoObjectUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoObjectUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GroupUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroupUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MapObject":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MapObject), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MapObjectUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MapObjectUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Parameter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Parameter), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParameterCategory":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParameterCategory), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParameterCategoryUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParameterCategoryUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParameterShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParameterShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParametersAggregate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParametersAggregate), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParametersAggregateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParametersAggregateShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Position":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Position), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Repository":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Repository), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Scenario":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scenario), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "User":
					identifierMap[ident.Name] = __gong__stageSetInit(new(User), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "UserUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UserUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Workspace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Workspace), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						inst.Probability = ProbabilityEnum(GongExtractString(rhs))
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
						inst.Direction = DirectionType(GongExtractString(rhs))
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
						inst.Direction = DirectionType(GongExtractString(rhs))
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
		}
		return true
	})

	return nil
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
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
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}

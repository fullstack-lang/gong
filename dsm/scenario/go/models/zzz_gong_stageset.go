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
		actorstateOrdered := []*ActorState{}
		for actorstate := range stageSet.Stage.ActorStates {
			actorstateOrdered = append(actorstateOrdered, actorstate)
		}
		sort.Slice(actorstateOrdered, func(i, j int) bool {
			return stageSet.Stage.ActorState_stagedOrder[actorstateOrdered[i]] < stageSet.Stage.ActorState_stagedOrder[actorstateOrdered[j]]
		})
		for _, actorstate := range actorstateOrdered {
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
		actorstateshapeOrdered := []*ActorStateShape{}
		for actorstateshape := range stageSet.Stage.ActorStateShapes {
			actorstateshapeOrdered = append(actorstateshapeOrdered, actorstateshape)
		}
		sort.Slice(actorstateshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ActorStateShape_stagedOrder[actorstateshapeOrdered[i]] < stageSet.Stage.ActorStateShape_stagedOrder[actorstateshapeOrdered[j]]
		})
		for _, actorstateshape := range actorstateshapeOrdered {
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
		actorstatetransitionOrdered := []*ActorStateTransition{}
		for actorstatetransition := range stageSet.Stage.ActorStateTransitions {
			actorstatetransitionOrdered = append(actorstatetransitionOrdered, actorstatetransition)
		}
		sort.Slice(actorstatetransitionOrdered, func(i, j int) bool {
			return stageSet.Stage.ActorStateTransition_stagedOrder[actorstatetransitionOrdered[i]] < stageSet.Stage.ActorStateTransition_stagedOrder[actorstatetransitionOrdered[j]]
		})
		for _, actorstatetransition := range actorstatetransitionOrdered {
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
		actorstatetransitionshapeOrdered := []*ActorStateTransitionShape{}
		for actorstatetransitionshape := range stageSet.Stage.ActorStateTransitionShapes {
			actorstatetransitionshapeOrdered = append(actorstatetransitionshapeOrdered, actorstatetransitionshape)
		}
		sort.Slice(actorstatetransitionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshapeOrdered[i]] < stageSet.Stage.ActorStateTransitionShape_stagedOrder[actorstatetransitionshapeOrdered[j]]
		})
		for _, actorstatetransitionshape := range actorstatetransitionshapeOrdered {
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
		analysisOrdered := []*Analysis{}
		for analysis := range stageSet.Stage.Analysiss {
			analysisOrdered = append(analysisOrdered, analysis)
		}
		sort.Slice(analysisOrdered, func(i, j int) bool {
			return stageSet.Stage.Analysis_stagedOrder[analysisOrdered[i]] < stageSet.Stage.Analysis_stagedOrder[analysisOrdered[j]]
		})
		for _, analysis := range analysisOrdered {
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
		controlpointshapeOrdered := []*ControlPointShape{}
		for controlpointshape := range stageSet.Stage.ControlPointShapes {
			controlpointshapeOrdered = append(controlpointshapeOrdered, controlpointshape)
		}
		sort.Slice(controlpointshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ControlPointShape_stagedOrder[controlpointshapeOrdered[i]] < stageSet.Stage.ControlPointShape_stagedOrder[controlpointshapeOrdered[j]]
		})
		for _, controlpointshape := range controlpointshapeOrdered {
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
		diagramOrdered := []*Diagram{}
		for diagram := range stageSet.Stage.Diagrams {
			diagramOrdered = append(diagramOrdered, diagram)
		}
		sort.Slice(diagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Diagram_stagedOrder[diagramOrdered[i]] < stageSet.Stage.Diagram_stagedOrder[diagramOrdered[j]]
		})
		for _, diagram := range diagramOrdered {
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
		documentOrdered := []*Document{}
		for document := range stageSet.Stage.Documents {
			documentOrdered = append(documentOrdered, document)
		}
		sort.Slice(documentOrdered, func(i, j int) bool {
			return stageSet.Stage.Document_stagedOrder[documentOrdered[i]] < stageSet.Stage.Document_stagedOrder[documentOrdered[j]]
		})
		for _, document := range documentOrdered {
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
		documentuseOrdered := []*DocumentUse{}
		for documentuse := range stageSet.Stage.DocumentUses {
			documentuseOrdered = append(documentuseOrdered, documentuse)
		}
		sort.Slice(documentuseOrdered, func(i, j int) bool {
			return stageSet.Stage.DocumentUse_stagedOrder[documentuseOrdered[i]] < stageSet.Stage.DocumentUse_stagedOrder[documentuseOrdered[j]]
		})
		for _, documentuse := range documentuseOrdered {
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
		evolutiondirectionOrdered := []*EvolutionDirection{}
		for evolutiondirection := range stageSet.Stage.EvolutionDirections {
			evolutiondirectionOrdered = append(evolutiondirectionOrdered, evolutiondirection)
		}
		sort.Slice(evolutiondirectionOrdered, func(i, j int) bool {
			return stageSet.Stage.EvolutionDirection_stagedOrder[evolutiondirectionOrdered[i]] < stageSet.Stage.EvolutionDirection_stagedOrder[evolutiondirectionOrdered[j]]
		})
		for _, evolutiondirection := range evolutiondirectionOrdered {
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
		evolutiondirectionshapeOrdered := []*EvolutionDirectionShape{}
		for evolutiondirectionshape := range stageSet.Stage.EvolutionDirectionShapes {
			evolutiondirectionshapeOrdered = append(evolutiondirectionshapeOrdered, evolutiondirectionshape)
		}
		sort.Slice(evolutiondirectionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshapeOrdered[i]] < stageSet.Stage.EvolutionDirectionShape_stagedOrder[evolutiondirectionshapeOrdered[j]]
		})
		for _, evolutiondirectionshape := range evolutiondirectionshapeOrdered {
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
		fooOrdered := []*Foo{}
		for foo := range stageSet.Stage.Foos {
			fooOrdered = append(fooOrdered, foo)
		}
		sort.Slice(fooOrdered, func(i, j int) bool {
			return stageSet.Stage.Foo_stagedOrder[fooOrdered[i]] < stageSet.Stage.Foo_stagedOrder[fooOrdered[j]]
		})
		for _, foo := range fooOrdered {
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
		geoobjectOrdered := []*GeoObject{}
		for geoobject := range stageSet.Stage.GeoObjects {
			geoobjectOrdered = append(geoobjectOrdered, geoobject)
		}
		sort.Slice(geoobjectOrdered, func(i, j int) bool {
			return stageSet.Stage.GeoObject_stagedOrder[geoobjectOrdered[i]] < stageSet.Stage.GeoObject_stagedOrder[geoobjectOrdered[j]]
		})
		for _, geoobject := range geoobjectOrdered {
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
		geoobjectuseOrdered := []*GeoObjectUse{}
		for geoobjectuse := range stageSet.Stage.GeoObjectUses {
			geoobjectuseOrdered = append(geoobjectuseOrdered, geoobjectuse)
		}
		sort.Slice(geoobjectuseOrdered, func(i, j int) bool {
			return stageSet.Stage.GeoObjectUse_stagedOrder[geoobjectuseOrdered[i]] < stageSet.Stage.GeoObjectUse_stagedOrder[geoobjectuseOrdered[j]]
		})
		for _, geoobjectuse := range geoobjectuseOrdered {
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
		groupOrdered := []*Group{}
		for group := range stageSet.Stage.Groups {
			groupOrdered = append(groupOrdered, group)
		}
		sort.Slice(groupOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_stagedOrder[groupOrdered[i]] < stageSet.Stage.Group_stagedOrder[groupOrdered[j]]
		})
		for _, group := range groupOrdered {
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
		groupuseOrdered := []*GroupUse{}
		for groupuse := range stageSet.Stage.GroupUses {
			groupuseOrdered = append(groupuseOrdered, groupuse)
		}
		sort.Slice(groupuseOrdered, func(i, j int) bool {
			return stageSet.Stage.GroupUse_stagedOrder[groupuseOrdered[i]] < stageSet.Stage.GroupUse_stagedOrder[groupuseOrdered[j]]
		})
		for _, groupuse := range groupuseOrdered {
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
		libraryOrdered := []*Library{}
		for library := range stageSet.Stage.Librarys {
			libraryOrdered = append(libraryOrdered, library)
		}
		sort.Slice(libraryOrdered, func(i, j int) bool {
			return stageSet.Stage.Library_stagedOrder[libraryOrdered[i]] < stageSet.Stage.Library_stagedOrder[libraryOrdered[j]]
		})
		for _, library := range libraryOrdered {
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
		mapobjectOrdered := []*MapObject{}
		for mapobject := range stageSet.Stage.MapObjects {
			mapobjectOrdered = append(mapobjectOrdered, mapobject)
		}
		sort.Slice(mapobjectOrdered, func(i, j int) bool {
			return stageSet.Stage.MapObject_stagedOrder[mapobjectOrdered[i]] < stageSet.Stage.MapObject_stagedOrder[mapobjectOrdered[j]]
		})
		for _, mapobject := range mapobjectOrdered {
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
		mapobjectuseOrdered := []*MapObjectUse{}
		for mapobjectuse := range stageSet.Stage.MapObjectUses {
			mapobjectuseOrdered = append(mapobjectuseOrdered, mapobjectuse)
		}
		sort.Slice(mapobjectuseOrdered, func(i, j int) bool {
			return stageSet.Stage.MapObjectUse_stagedOrder[mapobjectuseOrdered[i]] < stageSet.Stage.MapObjectUse_stagedOrder[mapobjectuseOrdered[j]]
		})
		for _, mapobjectuse := range mapobjectuseOrdered {
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
		parameterOrdered := []*Parameter{}
		for parameter := range stageSet.Stage.Parameters {
			parameterOrdered = append(parameterOrdered, parameter)
		}
		sort.Slice(parameterOrdered, func(i, j int) bool {
			return stageSet.Stage.Parameter_stagedOrder[parameterOrdered[i]] < stageSet.Stage.Parameter_stagedOrder[parameterOrdered[j]]
		})
		for _, parameter := range parameterOrdered {
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
		parametercategoryOrdered := []*ParameterCategory{}
		for parametercategory := range stageSet.Stage.ParameterCategorys {
			parametercategoryOrdered = append(parametercategoryOrdered, parametercategory)
		}
		sort.Slice(parametercategoryOrdered, func(i, j int) bool {
			return stageSet.Stage.ParameterCategory_stagedOrder[parametercategoryOrdered[i]] < stageSet.Stage.ParameterCategory_stagedOrder[parametercategoryOrdered[j]]
		})
		for _, parametercategory := range parametercategoryOrdered {
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
		parametercategoryuseOrdered := []*ParameterCategoryUse{}
		for parametercategoryuse := range stageSet.Stage.ParameterCategoryUses {
			parametercategoryuseOrdered = append(parametercategoryuseOrdered, parametercategoryuse)
		}
		sort.Slice(parametercategoryuseOrdered, func(i, j int) bool {
			return stageSet.Stage.ParameterCategoryUse_stagedOrder[parametercategoryuseOrdered[i]] < stageSet.Stage.ParameterCategoryUse_stagedOrder[parametercategoryuseOrdered[j]]
		})
		for _, parametercategoryuse := range parametercategoryuseOrdered {
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
		parametershapeOrdered := []*ParameterShape{}
		for parametershape := range stageSet.Stage.ParameterShapes {
			parametershapeOrdered = append(parametershapeOrdered, parametershape)
		}
		sort.Slice(parametershapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ParameterShape_stagedOrder[parametershapeOrdered[i]] < stageSet.Stage.ParameterShape_stagedOrder[parametershapeOrdered[j]]
		})
		for _, parametershape := range parametershapeOrdered {
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
		parametersaggregateOrdered := []*ParametersAggregate{}
		for parametersaggregate := range stageSet.Stage.ParametersAggregates {
			parametersaggregateOrdered = append(parametersaggregateOrdered, parametersaggregate)
		}
		sort.Slice(parametersaggregateOrdered, func(i, j int) bool {
			return stageSet.Stage.ParametersAggregate_stagedOrder[parametersaggregateOrdered[i]] < stageSet.Stage.ParametersAggregate_stagedOrder[parametersaggregateOrdered[j]]
		})
		for _, parametersaggregate := range parametersaggregateOrdered {
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
		parametersaggregateshapeOrdered := []*ParametersAggregateShape{}
		for parametersaggregateshape := range stageSet.Stage.ParametersAggregateShapes {
			parametersaggregateshapeOrdered = append(parametersaggregateshapeOrdered, parametersaggregateshape)
		}
		sort.Slice(parametersaggregateshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ParametersAggregateShape_stagedOrder[parametersaggregateshapeOrdered[i]] < stageSet.Stage.ParametersAggregateShape_stagedOrder[parametersaggregateshapeOrdered[j]]
		})
		for _, parametersaggregateshape := range parametersaggregateshapeOrdered {
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
		positionOrdered := []*Position{}
		for position := range stageSet.Stage.Positions {
			positionOrdered = append(positionOrdered, position)
		}
		sort.Slice(positionOrdered, func(i, j int) bool {
			return stageSet.Stage.Position_stagedOrder[positionOrdered[i]] < stageSet.Stage.Position_stagedOrder[positionOrdered[j]]
		})
		for _, position := range positionOrdered {
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
		repositoryOrdered := []*Repository{}
		for repository := range stageSet.Stage.Repositorys {
			repositoryOrdered = append(repositoryOrdered, repository)
		}
		sort.Slice(repositoryOrdered, func(i, j int) bool {
			return stageSet.Stage.Repository_stagedOrder[repositoryOrdered[i]] < stageSet.Stage.Repository_stagedOrder[repositoryOrdered[j]]
		})
		for _, repository := range repositoryOrdered {
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
		scenarioOrdered := []*Scenario{}
		for scenario := range stageSet.Stage.Scenarios {
			scenarioOrdered = append(scenarioOrdered, scenario)
		}
		sort.Slice(scenarioOrdered, func(i, j int) bool {
			return stageSet.Stage.Scenario_stagedOrder[scenarioOrdered[i]] < stageSet.Stage.Scenario_stagedOrder[scenarioOrdered[j]]
		})
		for _, scenario := range scenarioOrdered {
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
		userOrdered := []*User{}
		for user := range stageSet.Stage.Users {
			userOrdered = append(userOrdered, user)
		}
		sort.Slice(userOrdered, func(i, j int) bool {
			return stageSet.Stage.User_stagedOrder[userOrdered[i]] < stageSet.Stage.User_stagedOrder[userOrdered[j]]
		})
		for _, user := range userOrdered {
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
		useruseOrdered := []*UserUse{}
		for useruse := range stageSet.Stage.UserUses {
			useruseOrdered = append(useruseOrdered, useruse)
		}
		sort.Slice(useruseOrdered, func(i, j int) bool {
			return stageSet.Stage.UserUse_stagedOrder[useruseOrdered[i]] < stageSet.Stage.UserUse_stagedOrder[useruseOrdered[j]]
		})
		for _, useruse := range useruseOrdered {
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
		workspaceOrdered := []*Workspace{}
		for workspace := range stageSet.Stage.Workspaces {
			workspaceOrdered = append(workspaceOrdered, workspace)
		}
		sort.Slice(workspaceOrdered, func(i, j int) bool {
			return stageSet.Stage.Workspace_stagedOrder[workspaceOrdered[i]] < stageSet.Stage.Workspace_stagedOrder[workspaceOrdered[j]]
		})
		for _, workspace := range workspaceOrdered {
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
					if !preserveOrder {
						inst := (&ActorState{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ActorState)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ActorStateShape":
					if !preserveOrder {
						inst := (&ActorStateShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ActorStateShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ActorStateTransition":
					if !preserveOrder {
						inst := (&ActorStateTransition{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ActorStateTransition)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ActorStateTransitionShape":
					if !preserveOrder {
						inst := (&ActorStateTransitionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ActorStateTransitionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Analysis":
					if !preserveOrder {
						inst := (&Analysis{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Analysis)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ControlPointShape":
					if !preserveOrder {
						inst := (&ControlPointShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ControlPointShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Diagram":
					if !preserveOrder {
						inst := (&Diagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Diagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Document":
					if !preserveOrder {
						inst := (&Document{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Document)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DocumentUse":
					if !preserveOrder {
						inst := (&DocumentUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DocumentUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "EvolutionDirection":
					if !preserveOrder {
						inst := (&EvolutionDirection{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(EvolutionDirection)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "EvolutionDirectionShape":
					if !preserveOrder {
						inst := (&EvolutionDirectionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(EvolutionDirectionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Foo":
					if !preserveOrder {
						inst := (&Foo{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Foo)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GeoObject":
					if !preserveOrder {
						inst := (&GeoObject{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GeoObject)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GeoObjectUse":
					if !preserveOrder {
						inst := (&GeoObjectUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GeoObjectUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group":
					if !preserveOrder {
						inst := (&Group{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GroupUse":
					if !preserveOrder {
						inst := (&GroupUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GroupUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Library":
					if !preserveOrder {
						inst := (&Library{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Library)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MapObject":
					if !preserveOrder {
						inst := (&MapObject{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MapObject)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MapObjectUse":
					if !preserveOrder {
						inst := (&MapObjectUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MapObjectUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Parameter":
					if !preserveOrder {
						inst := (&Parameter{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Parameter)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParameterCategory":
					if !preserveOrder {
						inst := (&ParameterCategory{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParameterCategory)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParameterCategoryUse":
					if !preserveOrder {
						inst := (&ParameterCategoryUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParameterCategoryUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParameterShape":
					if !preserveOrder {
						inst := (&ParameterShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParameterShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParametersAggregate":
					if !preserveOrder {
						inst := (&ParametersAggregate{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParametersAggregate)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParametersAggregateShape":
					if !preserveOrder {
						inst := (&ParametersAggregateShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParametersAggregateShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Position":
					if !preserveOrder {
						inst := (&Position{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Position)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Repository":
					if !preserveOrder {
						inst := (&Repository{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Repository)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Scenario":
					if !preserveOrder {
						inst := (&Scenario{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Scenario)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "User":
					if !preserveOrder {
						inst := (&User{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(User)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "UserUse":
					if !preserveOrder {
						inst := (&UserUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(UserUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Workspace":
					if !preserveOrder {
						inst := (&Workspace{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Workspace)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorState); ok {
									inst.ActorState = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorState); ok {
									inst.StartState = typedTarget
								}
							}
						}
					case "EndState":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorState); ok {
									inst.EndState = typedTarget
								}
							}
						}
					case "Justifications":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Parameter); ok {
										inst.Justifications = append(inst.Justifications, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorStateTransition); ok {
									inst.ActorStateTransition = typedTarget
								}
							}
						}
					case "Start":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorStateShape); ok {
									inst.Start = typedTarget
								}
							}
						}
					case "End":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorStateShape); ok {
									inst.End = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ControlPointShape); ok {
										inst.ControlPointShapes = append(inst.ControlPointShapes, typedTarget)
									}
								}
							}
						}
					}
				case *Analysis:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Scenarios":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Scenario); ok {
										inst.Scenarios = append(inst.Scenarios, typedTarget)
									}
								}
							}
						}
					case "IsScenariosNodeExpanded":
						inst.IsScenariosNodeExpanded = GongExtractBool(rhs)
					case "GroupUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GroupUse); ok {
										inst.GroupUse = append(inst.GroupUse, typedTarget)
									}
								}
							}
						}
					case "IsGroupUseNodeExpanded":
						inst.IsGroupUseNodeExpanded = GongExtractBool(rhs)
					case "GeoObjectUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GeoObjectUse); ok {
										inst.GeoObjectUse = append(inst.GeoObjectUse, typedTarget)
									}
								}
							}
						}
					case "IsGeoObjectUseNodeExpanded":
						inst.IsGeoObjectUseNodeExpanded = GongExtractBool(rhs)
					case "MapUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*MapObjectUse); ok {
										inst.MapUse = append(inst.MapUse, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*EvolutionDirectionShape); ok {
										inst.EvolutionDirectionShapes = append(inst.EvolutionDirectionShapes, typedTarget)
									}
								}
							}
						}
					case "EvolutionDirectionsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*EvolutionDirection); ok {
										inst.EvolutionDirectionsWhoseNodeIsExpanded = append(inst.EvolutionDirectionsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsEvolutionDirectionsNodeExpanded":
						inst.IsEvolutionDirectionsNodeExpanded = GongExtractBool(rhs)
					case "ActorStateShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ActorStateShape); ok {
										inst.ActorStateShapes = append(inst.ActorStateShapes, typedTarget)
									}
								}
							}
						}
					case "ActorStatesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ActorState); ok {
										inst.ActorStatesWhoseNodeIsExpanded = append(inst.ActorStatesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsActorStatesNodeExpanded":
						inst.IsActorStatesNodeExpanded = GongExtractBool(rhs)
					case "ParameterShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParameterShape); ok {
										inst.ParameterShapes = append(inst.ParameterShapes, typedTarget)
									}
								}
							}
						}
					case "ParametersWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Parameter); ok {
										inst.ParametersWhoseNodeIsExpanded = append(inst.ParametersWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsParametersNodeExpanded":
						inst.IsParametersNodeExpanded = GongExtractBool(rhs)
					case "ScenarioParameterShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParametersAggregateShape); ok {
										inst.ScenarioParameterShapes = append(inst.ScenarioParameterShapes, typedTarget)
									}
								}
							}
						}
					case "ParametersAggregatesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParametersAggregate); ok {
										inst.ParametersAggregatesWhoseNodeIsExpanded = append(inst.ParametersAggregatesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsParametersAggregatesNodeExpanded":
						inst.IsParametersAggregatesNodeExpanded = GongExtractBool(rhs)
					case "ActorStateTransitionShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ActorStateTransitionShape); ok {
										inst.ActorStateTransitionShapes = append(inst.ActorStateTransitionShapes, typedTarget)
									}
								}
							}
						}
					case "ActorStateTransitionsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ActorStateTransition); ok {
										inst.ActorStateTransitionsWhoseNodeIsExpanded = append(inst.ActorStateTransitionsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "End":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GeoObjectUse); ok {
										inst.GeoObjectUse = append(inst.GeoObjectUse, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Document); ok {
									inst.Document = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EvolutionDirection); ok {
									inst.EvolutionDirection = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GeoObject); ok {
									inst.GeoObject = typedTarget
								}
							}
						}
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "UserUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*UserUse); ok {
										inst.UserUse = append(inst.UserUse, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
									inst.Group = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Analysis); ok {
										inst.Analyses = append(inst.Analyses, typedTarget)
									}
								}
							}
						}
					case "IsAnalysesNodeExpanded":
						inst.IsAnalysesNodeExpanded = GongExtractBool(rhs)
					case "SubLibraries":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Library); ok {
										inst.SubLibraries = append(inst.SubLibraries, typedTarget)
									}
								}
							}
						}
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Library); ok {
										inst.SubLibrariesWhoseNodeIsExpanded = append(inst.SubLibrariesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MapObject); ok {
									inst.Map = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "End":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Force":
						inst.Force = GongExtractFloat(rhs)
					case "GroupUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GroupUse); ok {
										inst.GroupUse = append(inst.GroupUse, typedTarget)
									}
								}
							}
						}
					case "DocumentUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DocumentUse); ok {
										inst.DocumentUse = append(inst.DocumentUse, typedTarget)
									}
								}
							}
						}
					case "GeoObjectUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GeoObjectUse); ok {
										inst.GeoObjectUse = append(inst.GeoObjectUse, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParameterShape); ok {
										inst.ParameterUse = append(inst.ParameterUse, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParameterCategory); ok {
									inst.ParameterCategory = typedTarget
								}
							}
						}
					}
				case *ParameterShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parameter":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Parameter); ok {
									inst.Parameter = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Parameter); ok {
										inst.Parameters = append(inst.Parameters, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParametersAggregate); ok {
									inst.ScenarioParameter = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParameterShape); ok {
										inst.ParameterUse = append(inst.ParameterUse, typedTarget)
									}
								}
							}
						}
					case "GroupUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GroupUse); ok {
										inst.GroupUse = append(inst.GroupUse, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Diagram); ok {
										inst.Diagrams = append(inst.Diagrams, typedTarget)
									}
								}
							}
						}
					case "IsDiagramsNodeExpanded":
						inst.IsDiagramsNodeExpanded = GongExtractBool(rhs)
					case "ActorStates":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ActorState); ok {
										inst.ActorStates = append(inst.ActorStates, typedTarget)
									}
								}
							}
						}
					case "IsActorStatesNodeExpanded":
						inst.IsActorStatesNodeExpanded = GongExtractBool(rhs)
					case "ActorStateTransitions":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ActorStateTransition); ok {
										inst.ActorStateTransitions = append(inst.ActorStateTransitions, typedTarget)
									}
								}
							}
						}
					case "IsActorStateTransitionsNodeExpanded":
						inst.IsActorStateTransitionsNodeExpanded = GongExtractBool(rhs)
					case "EvolutionDirections":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*EvolutionDirection); ok {
										inst.EvolutionDirections = append(inst.EvolutionDirections, typedTarget)
									}
								}
							}
						}
					case "IsEvolutionDirectionsNodeExpanded":
						inst.IsEvolutionDirectionsNodeExpanded = GongExtractBool(rhs)
					case "Parameters":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Parameter); ok {
										inst.Parameters = append(inst.Parameters, typedTarget)
									}
								}
							}
						}
					case "IsParametersNodeExpanded":
						inst.IsParametersNodeExpanded = GongExtractBool(rhs)
					case "ParametersAggretates":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParametersAggregate); ok {
										inst.ParametersAggretates = append(inst.ParametersAggretates, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*User); ok {
									inst.User = typedTarget
								}
							}
						}
					}
				case *Workspace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SelectedDiagram":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Diagram); ok {
									inst.SelectedDiagram = typedTarget
								}
							}
						}
					case "Default_EvolutionDirectionShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EvolutionDirectionShape); ok {
									inst.Default_EvolutionDirectionShape = typedTarget
								}
							}
						}
					case "Default_ParameterShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParameterShape); ok {
									inst.Default_ParameterShape = typedTarget
								}
							}
						}
					case "Default_ScenarioParameterShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParametersAggregateShape); ok {
									inst.Default_ScenarioParameterShape = typedTarget
								}
							}
						}
					case "Default_ActorStateShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorStateShape); ok {
									inst.Default_ActorStateShape = typedTarget
								}
							}
						}
					case "Default_ActorStateTransitionShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ActorStateTransitionShape); ok {
									inst.Default_ActorStateTransitionShape = typedTarget
								}
							}
						}
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

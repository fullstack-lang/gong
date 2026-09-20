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

	if stageSet.Stage != nil {
		analysisneedOrdered := []*AnalysisNeed{}
		for analysisneed := range stageSet.Stage.AnalysisNeeds {
			analysisneedOrdered = append(analysisneedOrdered, analysisneed)
		}
		sort.Slice(analysisneedOrdered, func(i, j int) bool {
			return stageSet.Stage.AnalysisNeed_stagedOrder[analysisneedOrdered[i]] < stageSet.Stage.AnalysisNeed_stagedOrder[analysisneedOrdered[j]]
		})
		for _, analysisneed := range analysisneedOrdered {
			analysisneedIdent := "__stage_0" + analysisneed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AnalysisNeed{Name: %s}).Stage(stageSet.Stage)", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", analysisneedIdent, analysisneed.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		conceptOrdered := []*Concept{}
		for concept := range stageSet.Stage.Concepts {
			conceptOrdered = append(conceptOrdered, concept)
		}
		sort.Slice(conceptOrdered, func(i, j int) bool {
			return stageSet.Stage.Concept_stagedOrder[conceptOrdered[i]] < stageSet.Stage.Concept_stagedOrder[conceptOrdered[j]]
		})
		for _, concept := range conceptOrdered {
			conceptIdent := "__stage_0" + concept.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Concept{Name: %s}).Stage(stageSet.Stage)", conceptIdent, __gong__toRawStringLiteral(concept.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conceptIdent, __gong__toRawStringLiteral(concept.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", conceptIdent, __gong__toRawStringLiteral(concept.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", conceptIdent, concept.IsExpanded))
			for _, elem := range concept.Tools {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tools = append(%s.Tools, %s)", conceptIdent, conceptIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		conceptshapeOrdered := []*ConceptShape{}
		for conceptshape := range stageSet.Stage.ConceptShapes {
			conceptshapeOrdered = append(conceptshapeOrdered, conceptshape)
		}
		sort.Slice(conceptshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ConceptShape_stagedOrder[conceptshapeOrdered[i]] < stageSet.Stage.ConceptShape_stagedOrder[conceptshapeOrdered[j]]
		})
		for _, conceptshape := range conceptshapeOrdered {
			conceptshapeIdent := "__stage_0" + conceptshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ConceptShape{Name: %s}).Stage(stageSet.Stage)", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", conceptshapeIdent, conceptshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", conceptshapeIdent, conceptshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", conceptshapeIdent, conceptshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", conceptshapeIdent, conceptshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", conceptshapeIdent, conceptshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", conceptshapeIdent, conceptshape.IsHidden))
			if conceptshape.Concept != nil {
				targetIdent := "__stage_0" + conceptshape.Concept.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", conceptshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		concernOrdered := []*Concern{}
		for concern := range stageSet.Stage.Concerns {
			concernOrdered = append(concernOrdered, concern)
		}
		sort.Slice(concernOrdered, func(i, j int) bool {
			return stageSet.Stage.Concern_stagedOrder[concernOrdered[i]] < stageSet.Stage.Concern_stagedOrder[concernOrdered[j]]
		})
		for _, concern := range concernOrdered {
			concernIdent := "__stage_0" + concern.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Concern{Name: %s}).Stage(stageSet.Stage)", concernIdent, __gong__toRawStringLiteral(concern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernIdent, __gong__toRawStringLiteral(concern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDAirbus = %s", concernIdent, __gong__toRawStringLiteral(concern.IDAirbus)))
			values.WriteString(fmt.Sprintf("\n\t%s.Priority = %s", concernIdent, __gong__toRawStringLiteral(string(concern.Priority))))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", concernIdent, __gong__toRawStringLiteral(concern.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", concernIdent, concern.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", concernIdent, __gong__toRawStringLiteral(concern.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInputsNodeExpanded = %t", concernIdent, concern.IsInputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOutputsNodeExpanded = %t", concernIdent, concern.IsOutputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCompletion = %t", concernIdent, concern.IsWithCompletion))
			values.WriteString(fmt.Sprintf("\n\t%s.Completion = %s", concernIdent, __gong__toRawStringLiteral(string(concern.Completion))))
			for _, elem := range concern.SubConcerns {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubConcerns = append(%s.SubConcerns, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Inputs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inputs = append(%s.Inputs, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Outputs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Outputs = append(%s.Outputs, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Requirements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirements = append(%s.Requirements, %s)", concernIdent, concernIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		concerncompositionshapeOrdered := []*ConcernCompositionShape{}
		for concerncompositionshape := range stageSet.Stage.ConcernCompositionShapes {
			concerncompositionshapeOrdered = append(concerncompositionshapeOrdered, concerncompositionshape)
		}
		sort.Slice(concerncompositionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ConcernCompositionShape_stagedOrder[concerncompositionshapeOrdered[i]] < stageSet.Stage.ConcernCompositionShape_stagedOrder[concerncompositionshapeOrdered[j]]
		})
		for _, concerncompositionshape := range concerncompositionshapeOrdered {
			concerncompositionshapeIdent := "__stage_0" + concerncompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ConcernCompositionShape{Name: %s}).Stage(stageSet.Stage)", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concerncompositionshapeIdent, concerncompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concerncompositionshapeIdent, concerncompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(string(concerncompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(string(concerncompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concerncompositionshapeIdent, concerncompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concerncompositionshapeIdent, concerncompositionshape.IsHidden))
			if concerncompositionshape.Concern != nil {
				targetIdent := "__stage_0" + concerncompositionshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concerncompositionshapeIdent, targetIdent))
			}
			for _, elem := range concerncompositionshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concerncompositionshapeIdent, concerncompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		concerninputshapeOrdered := []*ConcernInputShape{}
		for concerninputshape := range stageSet.Stage.ConcernInputShapes {
			concerninputshapeOrdered = append(concerninputshapeOrdered, concerninputshape)
		}
		sort.Slice(concerninputshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ConcernInputShape_stagedOrder[concerninputshapeOrdered[i]] < stageSet.Stage.ConcernInputShape_stagedOrder[concerninputshapeOrdered[j]]
		})
		for _, concerninputshape := range concerninputshapeOrdered {
			concerninputshapeIdent := "__stage_0" + concerninputshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ConcernInputShape{Name: %s}).Stage(stageSet.Stage)", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concerninputshapeIdent, concerninputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concerninputshapeIdent, concerninputshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", concerninputshapeIdent, __gong__toRawStringLiteral(string(concerninputshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", concerninputshapeIdent, __gong__toRawStringLiteral(string(concerninputshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concerninputshapeIdent, concerninputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concerninputshapeIdent, concerninputshape.IsHidden))
			if concerninputshape.Deliverable != nil {
				targetIdent := "__stage_0" + concerninputshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", concerninputshapeIdent, targetIdent))
			}
			if concerninputshape.Concern != nil {
				targetIdent := "__stage_0" + concerninputshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concerninputshapeIdent, targetIdent))
			}
			for _, elem := range concerninputshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concerninputshapeIdent, concerninputshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		concernoutputshapeOrdered := []*ConcernOutputShape{}
		for concernoutputshape := range stageSet.Stage.ConcernOutputShapes {
			concernoutputshapeOrdered = append(concernoutputshapeOrdered, concernoutputshape)
		}
		sort.Slice(concernoutputshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ConcernOutputShape_stagedOrder[concernoutputshapeOrdered[i]] < stageSet.Stage.ConcernOutputShape_stagedOrder[concernoutputshapeOrdered[j]]
		})
		for _, concernoutputshape := range concernoutputshapeOrdered {
			concernoutputshapeIdent := "__stage_0" + concernoutputshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ConcernOutputShape{Name: %s}).Stage(stageSet.Stage)", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concernoutputshapeIdent, concernoutputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concernoutputshapeIdent, concernoutputshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(string(concernoutputshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(string(concernoutputshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concernoutputshapeIdent, concernoutputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concernoutputshapeIdent, concernoutputshape.IsHidden))
			if concernoutputshape.Concern != nil {
				targetIdent := "__stage_0" + concernoutputshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concernoutputshapeIdent, targetIdent))
			}
			if concernoutputshape.Deliverable != nil {
				targetIdent := "__stage_0" + concernoutputshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", concernoutputshapeIdent, targetIdent))
			}
			for _, elem := range concernoutputshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concernoutputshapeIdent, concernoutputshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		concernshapeOrdered := []*ConcernShape{}
		for concernshape := range stageSet.Stage.ConcernShapes {
			concernshapeOrdered = append(concernshapeOrdered, concernshape)
		}
		sort.Slice(concernshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ConcernShape_stagedOrder[concernshapeOrdered[i]] < stageSet.Stage.ConcernShape_stagedOrder[concernshapeOrdered[j]]
		})
		for _, concernshape := range concernshapeOrdered {
			concernshapeIdent := "__stage_0" + concernshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ConcernShape{Name: %s}).Stage(stageSet.Stage)", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", concernshapeIdent, concernshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", concernshapeIdent, concernshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", concernshapeIdent, concernshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", concernshapeIdent, concernshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", concernshapeIdent, concernshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concernshapeIdent, concernshape.IsHidden))
			if concernshape.Concern != nil {
				targetIdent := "__stage_0" + concernshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concernshapeIdent, targetIdent))
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
			controlpointshapeIdent := "__stage_0" + controlpointshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ControlPointShape{Name: %s}).Stage(stageSet.Stage)", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointshapeIdent, controlpointshape.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointshapeIdent, controlpointshape.Y_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartShapeTheClosestShape = %t", controlpointshapeIdent, controlpointshape.IsStartShapeTheClosestShape))
		}
	}
	if stageSet.Stage != nil {
		deliverableOrdered := []*Deliverable{}
		for deliverable := range stageSet.Stage.Deliverables {
			deliverableOrdered = append(deliverableOrdered, deliverable)
		}
		sort.Slice(deliverableOrdered, func(i, j int) bool {
			return stageSet.Stage.Deliverable_stagedOrder[deliverableOrdered[i]] < stageSet.Stage.Deliverable_stagedOrder[deliverableOrdered[j]]
		})
		for _, deliverable := range deliverableOrdered {
			deliverableIdent := "__stage_0" + deliverable.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Deliverable{Name: %s}).Stage(stageSet.Stage)", deliverableIdent, __gong__toRawStringLiteral(deliverable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", deliverableIdent, deliverable.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProducersNodeExpanded = %t", deliverableIdent, deliverable.IsProducersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConsumersNodeExpanded = %t", deliverableIdent, deliverable.IsConsumersNodeExpanded))
			for _, elem := range deliverable.SubDeliverables {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubDeliverables = append(%s.SubDeliverables, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
			for _, elem := range deliverable.Concepts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		deliverablecompositionshapeOrdered := []*DeliverableCompositionShape{}
		for deliverablecompositionshape := range stageSet.Stage.DeliverableCompositionShapes {
			deliverablecompositionshapeOrdered = append(deliverablecompositionshapeOrdered, deliverablecompositionshape)
		}
		sort.Slice(deliverablecompositionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DeliverableCompositionShape_stagedOrder[deliverablecompositionshapeOrdered[i]] < stageSet.Stage.DeliverableCompositionShape_stagedOrder[deliverablecompositionshapeOrdered[j]]
		})
		for _, deliverablecompositionshape := range deliverablecompositionshapeOrdered {
			deliverablecompositionshapeIdent := "__stage_0" + deliverablecompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DeliverableCompositionShape{Name: %s}).Stage(stageSet.Stage)", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(string(deliverablecompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(string(deliverablecompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverablecompositionshapeIdent, deliverablecompositionshape.IsHidden))
			if deliverablecompositionshape.Deliverable != nil {
				targetIdent := "__stage_0" + deliverablecompositionshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverablecompositionshapeIdent, targetIdent))
			}
			for _, elem := range deliverablecompositionshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverablecompositionshapeIdent, deliverablecompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		deliverableconceptshapeOrdered := []*DeliverableConceptShape{}
		for deliverableconceptshape := range stageSet.Stage.DeliverableConceptShapes {
			deliverableconceptshapeOrdered = append(deliverableconceptshapeOrdered, deliverableconceptshape)
		}
		sort.Slice(deliverableconceptshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DeliverableConceptShape_stagedOrder[deliverableconceptshapeOrdered[i]] < stageSet.Stage.DeliverableConceptShape_stagedOrder[deliverableconceptshapeOrdered[j]]
		})
		for _, deliverableconceptshape := range deliverableconceptshapeOrdered {
			deliverableconceptshapeIdent := "__stage_0" + deliverableconceptshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DeliverableConceptShape{Name: %s}).Stage(stageSet.Stage)", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(string(deliverableconceptshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(string(deliverableconceptshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableconceptshapeIdent, deliverableconceptshape.IsHidden))
			if deliverableconceptshape.Deliverable != nil {
				targetIdent := "__stage_0" + deliverableconceptshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableconceptshapeIdent, targetIdent))
			}
			if deliverableconceptshape.Concept != nil {
				targetIdent := "__stage_0" + deliverableconceptshape.Concept.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", deliverableconceptshapeIdent, targetIdent))
			}
			for _, elem := range deliverableconceptshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverableconceptshapeIdent, deliverableconceptshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		deliverableshapeOrdered := []*DeliverableShape{}
		for deliverableshape := range stageSet.Stage.DeliverableShapes {
			deliverableshapeOrdered = append(deliverableshapeOrdered, deliverableshape)
		}
		sort.Slice(deliverableshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DeliverableShape_stagedOrder[deliverableshapeOrdered[i]] < stageSet.Stage.DeliverableShape_stagedOrder[deliverableshapeOrdered[j]]
		})
		for _, deliverableshape := range deliverableshapeOrdered {
			deliverableshapeIdent := "__stage_0" + deliverableshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DeliverableShape{Name: %s}).Stage(stageSet.Stage)", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", deliverableshapeIdent, deliverableshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", deliverableshapeIdent, deliverableshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", deliverableshapeIdent, deliverableshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", deliverableshapeIdent, deliverableshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", deliverableshapeIdent, deliverableshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableshapeIdent, deliverableshape.IsHidden))
			if deliverableshape.Deliverable != nil {
				targetIdent := "__stage_0" + deliverableshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableshapeIdent, targetIdent))
			}
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
			diagramIdent := "__stage_0" + diagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Diagram{Name: %s}).Stage(stageSet.Stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramIdent, diagram.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowPrefix = %t", diagramIdent, diagram.ShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramIdent, diagram.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramIdent, diagram.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramIdent, diagram.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRequirementsNodeExpanded = %t", diagramIdent, diagram.IsRequirementsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConceptsNodeExpanded = %t", diagramIdent, diagram.IsConceptsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPBSNodeExpanded = %t", diagramIdent, diagram.IsPBSNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConcernsNodeExpanded = %t", diagramIdent, diagram.IsConcernsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramIdent, diagram.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStakeholdersNodeExpanded = %t", diagramIdent, diagram.IsStakeholdersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDiagramsNodeExpanded = %t", diagramIdent, diagram.IsDiagramsNodeExpanded))
			for _, elem := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseRequirementsNodeIsExpanded = append(%s.ConcernsWhoseRequirementsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Deliverable_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable_Shapes = append(%s.Deliverable_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseNodeIsExpanded = append(%s.DeliverablesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseConceptsNodeIsExpanded = append(%s.DeliverablesWhoseConceptsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableComposition_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableComposition_Shapes = append(%s.DeliverableComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Concern_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern_Shapes = append(%s.Concern_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseNodeIsExpanded = append(%s.ConcernsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseInputNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseInputNodeIsExpanded = append(%s.ConcernsWhoseInputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseStakeholderNodeIsExpanded = append(%s.ConcernsWhoseStakeholderNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernssWhoseOutputNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernssWhoseOutputNodeIsExpanded = append(%s.ConcernssWhoseOutputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernComposition_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernComposition_Shapes = append(%s.ConcernComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernInputShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernInputShapes = append(%s.ConcernInputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernOutputShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernOutputShapes = append(%s.ConcernOutputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Note_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NotesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteDeliverableShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteDeliverableShapes = append(%s.NoteDeliverableShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteTaskShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteTaskShapes = append(%s.NoteTaskShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteResourceShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteResourceShapes = append(%s.NoteResourceShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Stakeholder_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder_Shapes = append(%s.Stakeholder_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourcesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourceComposition_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourceComposition_Shapes = append(%s.ResourceComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.StakeholderConcernShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StakeholderConcernShapes = append(%s.StakeholderConcernShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Requirement_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement_Shapes = append(%s.Requirement_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementsWhoseNodeIsExpanded = append(%s.RequirementsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Concept_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept_Shapes = append(%s.Concept_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseNodeIsExpanded = append(%s.ConceptsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseDeliverablesNodeIsExpanded = append(%s.ConceptsWhoseDeliverablesNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableConceptShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableConceptShapes = append(%s.DeliverableConceptShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Diagram_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagram_Shapes = append(%s.Diagram_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DiagramsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramsWhoseNodeIsExpanded = append(%s.DiagramsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		diagramshapeOrdered := []*DiagramShape{}
		for diagramshape := range stageSet.Stage.DiagramShapes {
			diagramshapeOrdered = append(diagramshapeOrdered, diagramshape)
		}
		sort.Slice(diagramshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DiagramShape_stagedOrder[diagramshapeOrdered[i]] < stageSet.Stage.DiagramShape_stagedOrder[diagramshapeOrdered[j]]
		})
		for _, diagramshape := range diagramshapeOrdered {
			diagramshapeIdent := "__stage_0" + diagramshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DiagramShape{Name: %s}).Stage(stageSet.Stage)", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramshapeIdent, diagramshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", diagramshapeIdent, diagramshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", diagramshapeIdent, diagramshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramshapeIdent, diagramshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramshapeIdent, diagramshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", diagramshapeIdent, diagramshape.IsHidden))
			if diagramshape.Diagram != nil {
				targetIdent := "__stage_0" + diagramshape.Diagram.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagram = %s", diagramshapeIdent, targetIdent))
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
			libraryIdent := "__stage_0" + library.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Library{Name: %s}).Stage(stageSet.Stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			for _, elem := range library.RootDeliverables {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDeliverables = append(%s.RootDeliverables, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootConcerns {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootConcerns = append(%s.RootConcerns, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootStakeholders {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStakeholders = append(%s.RootStakeholders, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootRequirements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootRequirements = append(%s.RootRequirements, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootConcepts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootConcepts = append(%s.RootConcepts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.AnalysisNeeds {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnalysisNeeds = append(%s.AnalysisNeeds, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Notes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notes = append(%s.Notes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Diagrams {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibraries {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteOrdered := []*Note{}
		for note := range stageSet.Stage.Notes {
			noteOrdered = append(noteOrdered, note)
		}
		sort.Slice(noteOrdered, func(i, j int) bool {
			return stageSet.Stage.Note_stagedOrder[noteOrdered[i]] < stageSet.Stage.Note_stagedOrder[noteOrdered[j]]
		})
		for _, note := range noteOrdered {
			noteIdent := "__stage_0" + note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Note{Name: %s}).Stage(stageSet.Stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			for _, elem := range note.Deliverables {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverables = append(%s.Deliverables, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Tasks {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Resources {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resources = append(%s.Resources, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		notedeliverableshapeOrdered := []*NoteDeliverableShape{}
		for notedeliverableshape := range stageSet.Stage.NoteDeliverableShapes {
			notedeliverableshapeOrdered = append(notedeliverableshapeOrdered, notedeliverableshape)
		}
		sort.Slice(notedeliverableshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteDeliverableShape_stagedOrder[notedeliverableshapeOrdered[i]] < stageSet.Stage.NoteDeliverableShape_stagedOrder[notedeliverableshapeOrdered[j]]
		})
		for _, notedeliverableshape := range notedeliverableshapeOrdered {
			notedeliverableshapeIdent := "__stage_0" + notedeliverableshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NoteDeliverableShape{Name: %s}).Stage(stageSet.Stage)", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notedeliverableshapeIdent, notedeliverableshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notedeliverableshapeIdent, notedeliverableshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(string(notedeliverableshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(string(notedeliverableshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notedeliverableshapeIdent, notedeliverableshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notedeliverableshapeIdent, notedeliverableshape.IsHidden))
			if notedeliverableshape.Note != nil {
				targetIdent := "__stage_0" + notedeliverableshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notedeliverableshapeIdent, targetIdent))
			}
			if notedeliverableshape.Deliverable != nil {
				targetIdent := "__stage_0" + notedeliverableshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", notedeliverableshapeIdent, targetIdent))
			}
			for _, elem := range notedeliverableshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notedeliverableshapeIdent, notedeliverableshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteshapeOrdered := []*NoteShape{}
		for noteshape := range stageSet.Stage.NoteShapes {
			noteshapeOrdered = append(noteshapeOrdered, noteshape)
		}
		sort.Slice(noteshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteShape_stagedOrder[noteshapeOrdered[i]] < stageSet.Stage.NoteShape_stagedOrder[noteshapeOrdered[j]]
		})
		for _, noteshape := range noteshapeOrdered {
			noteshapeIdent := "__stage_0" + noteshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NoteShape{Name: %s}).Stage(stageSet.Stage)", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteshapeIdent, noteshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", noteshapeIdent, noteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", noteshapeIdent, noteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", noteshapeIdent, noteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", noteshapeIdent, noteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteshapeIdent, noteshape.IsHidden))
			if noteshape.Note != nil {
				targetIdent := "__stage_0" + noteshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		notestakeholdershapeOrdered := []*NoteStakeholderShape{}
		for notestakeholdershape := range stageSet.Stage.NoteStakeholderShapes {
			notestakeholdershapeOrdered = append(notestakeholdershapeOrdered, notestakeholdershape)
		}
		sort.Slice(notestakeholdershapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteStakeholderShape_stagedOrder[notestakeholdershapeOrdered[i]] < stageSet.Stage.NoteStakeholderShape_stagedOrder[notestakeholdershapeOrdered[j]]
		})
		for _, notestakeholdershape := range notestakeholdershapeOrdered {
			notestakeholdershapeIdent := "__stage_0" + notestakeholdershape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NoteStakeholderShape{Name: %s}).Stage(stageSet.Stage)", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notestakeholdershapeIdent, notestakeholdershape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notestakeholdershapeIdent, notestakeholdershape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(string(notestakeholdershape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(string(notestakeholdershape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notestakeholdershapeIdent, notestakeholdershape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notestakeholdershapeIdent, notestakeholdershape.IsHidden))
			if notestakeholdershape.Note != nil {
				targetIdent := "__stage_0" + notestakeholdershape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notestakeholdershapeIdent, targetIdent))
			}
			if notestakeholdershape.Stakeholder != nil {
				targetIdent := "__stage_0" + notestakeholdershape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", notestakeholdershapeIdent, targetIdent))
			}
			for _, elem := range notestakeholdershape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notestakeholdershapeIdent, notestakeholdershapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		notetaskshapeOrdered := []*NoteTaskShape{}
		for notetaskshape := range stageSet.Stage.NoteTaskShapes {
			notetaskshapeOrdered = append(notetaskshapeOrdered, notetaskshape)
		}
		sort.Slice(notetaskshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteTaskShape_stagedOrder[notetaskshapeOrdered[i]] < stageSet.Stage.NoteTaskShape_stagedOrder[notetaskshapeOrdered[j]]
		})
		for _, notetaskshape := range notetaskshapeOrdered {
			notetaskshapeIdent := "__stage_0" + notetaskshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NoteTaskShape{Name: %s}).Stage(stageSet.Stage)", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notetaskshapeIdent, notetaskshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notetaskshapeIdent, notetaskshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notetaskshapeIdent, __gong__toRawStringLiteral(string(notetaskshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notetaskshapeIdent, __gong__toRawStringLiteral(string(notetaskshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notetaskshapeIdent, notetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notetaskshapeIdent, notetaskshape.IsHidden))
			if notetaskshape.Note != nil {
				targetIdent := "__stage_0" + notetaskshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notetaskshapeIdent, targetIdent))
			}
			if notetaskshape.Task != nil {
				targetIdent := "__stage_0" + notetaskshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", notetaskshapeIdent, targetIdent))
			}
			for _, elem := range notetaskshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notetaskshapeIdent, notetaskshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		requirementOrdered := []*Requirement{}
		for requirement := range stageSet.Stage.Requirements {
			requirementOrdered = append(requirementOrdered, requirement)
		}
		sort.Slice(requirementOrdered, func(i, j int) bool {
			return stageSet.Stage.Requirement_stagedOrder[requirementOrdered[i]] < stageSet.Stage.Requirement_stagedOrder[requirementOrdered[j]]
		})
		for _, requirement := range requirementOrdered {
			requirementIdent := "__stage_0" + requirement.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Requirement{Name: %s}).Stage(stageSet.Stage)", requirementIdent, __gong__toRawStringLiteral(requirement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementIdent, __gong__toRawStringLiteral(requirement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", requirementIdent, __gong__toRawStringLiteral(requirement.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", requirementIdent, requirement.IsExpanded))
			for _, elem := range requirement.SupportLevels {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportLevels = append(%s.SupportLevels, %s)", requirementIdent, requirementIdent, targetIdent))
			}
			for _, elem := range requirement.Concepts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", requirementIdent, requirementIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		requirementshapeOrdered := []*RequirementShape{}
		for requirementshape := range stageSet.Stage.RequirementShapes {
			requirementshapeOrdered = append(requirementshapeOrdered, requirementshape)
		}
		sort.Slice(requirementshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.RequirementShape_stagedOrder[requirementshapeOrdered[i]] < stageSet.Stage.RequirementShape_stagedOrder[requirementshapeOrdered[j]]
		})
		for _, requirementshape := range requirementshapeOrdered {
			requirementshapeIdent := "__stage_0" + requirementshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.RequirementShape{Name: %s}).Stage(stageSet.Stage)", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", requirementshapeIdent, requirementshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", requirementshapeIdent, requirementshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", requirementshapeIdent, requirementshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", requirementshapeIdent, requirementshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", requirementshapeIdent, requirementshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", requirementshapeIdent, requirementshape.IsHidden))
			if requirementshape.Requirement != nil {
				targetIdent := "__stage_0" + requirementshape.Requirement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement = %s", requirementshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		stakeholderOrdered := []*Stakeholder{}
		for stakeholder := range stageSet.Stage.Stakeholders {
			stakeholderOrdered = append(stakeholderOrdered, stakeholder)
		}
		sort.Slice(stakeholderOrdered, func(i, j int) bool {
			return stageSet.Stage.Stakeholder_stagedOrder[stakeholderOrdered[i]] < stageSet.Stage.Stakeholder_stagedOrder[stakeholderOrdered[j]]
		})
		for _, stakeholder := range stakeholderOrdered {
			stakeholderIdent := "__stage_0" + stakeholder.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Stakeholder{Name: %s}).Stage(stageSet.Stage)", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDAirbus = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.IDAirbus)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stakeholderIdent, stakeholder.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Description)))
			for _, elem := range stakeholder.Concerns {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concerns = append(%s.Concerns, %s)", stakeholderIdent, stakeholderIdent, targetIdent))
			}
			for _, elem := range stakeholder.SubStakeholders {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubStakeholders = append(%s.SubStakeholders, %s)", stakeholderIdent, stakeholderIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		stakeholdercompositionshapeOrdered := []*StakeholderCompositionShape{}
		for stakeholdercompositionshape := range stageSet.Stage.StakeholderCompositionShapes {
			stakeholdercompositionshapeOrdered = append(stakeholdercompositionshapeOrdered, stakeholdercompositionshape)
		}
		sort.Slice(stakeholdercompositionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshapeOrdered[i]] < stageSet.Stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshapeOrdered[j]]
		})
		for _, stakeholdercompositionshape := range stakeholdercompositionshapeOrdered {
			stakeholdercompositionshapeIdent := "__stage_0" + stakeholdercompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.StakeholderCompositionShape{Name: %s}).Stage(stageSet.Stage)", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(string(stakeholdercompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(string(stakeholdercompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholdercompositionshapeIdent, stakeholdercompositionshape.IsHidden))
			if stakeholdercompositionshape.Stakeholder != nil {
				targetIdent := "__stage_0" + stakeholdercompositionshape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholdercompositionshapeIdent, targetIdent))
			}
			for _, elem := range stakeholdercompositionshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", stakeholdercompositionshapeIdent, stakeholdercompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		stakeholderconcernshapeOrdered := []*StakeholderConcernShape{}
		for stakeholderconcernshape := range stageSet.Stage.StakeholderConcernShapes {
			stakeholderconcernshapeOrdered = append(stakeholderconcernshapeOrdered, stakeholderconcernshape)
		}
		sort.Slice(stakeholderconcernshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshapeOrdered[i]] < stageSet.Stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshapeOrdered[j]]
		})
		for _, stakeholderconcernshape := range stakeholderconcernshapeOrdered {
			stakeholderconcernshapeIdent := "__stage_0" + stakeholderconcernshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.StakeholderConcernShape{Name: %s}).Stage(stageSet.Stage)", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(string(stakeholderconcernshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(string(stakeholderconcernshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholderconcernshapeIdent, stakeholderconcernshape.IsHidden))
			if stakeholderconcernshape.Stakeholder != nil {
				targetIdent := "__stage_0" + stakeholderconcernshape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholderconcernshapeIdent, targetIdent))
			}
			if stakeholderconcernshape.Concern != nil {
				targetIdent := "__stage_0" + stakeholderconcernshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", stakeholderconcernshapeIdent, targetIdent))
			}
			for _, elem := range stakeholderconcernshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", stakeholderconcernshapeIdent, stakeholderconcernshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		stakeholdershapeOrdered := []*StakeholderShape{}
		for stakeholdershape := range stageSet.Stage.StakeholderShapes {
			stakeholdershapeOrdered = append(stakeholdershapeOrdered, stakeholdershape)
		}
		sort.Slice(stakeholdershapeOrdered, func(i, j int) bool {
			return stageSet.Stage.StakeholderShape_stagedOrder[stakeholdershapeOrdered[i]] < stageSet.Stage.StakeholderShape_stagedOrder[stakeholdershapeOrdered[j]]
		})
		for _, stakeholdershape := range stakeholdershapeOrdered {
			stakeholdershapeIdent := "__stage_0" + stakeholdershape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.StakeholderShape{Name: %s}).Stage(stageSet.Stage)", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stakeholdershapeIdent, stakeholdershape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", stakeholdershapeIdent, stakeholdershape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", stakeholdershapeIdent, stakeholdershape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", stakeholdershapeIdent, stakeholdershape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", stakeholdershapeIdent, stakeholdershape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholdershapeIdent, stakeholdershape.IsHidden))
			if stakeholdershape.Stakeholder != nil {
				targetIdent := "__stage_0" + stakeholdershape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholdershapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		supportlevelOrdered := []*SupportLevel{}
		for supportlevel := range stageSet.Stage.SupportLevels {
			supportlevelOrdered = append(supportlevelOrdered, supportlevel)
		}
		sort.Slice(supportlevelOrdered, func(i, j int) bool {
			return stageSet.Stage.SupportLevel_stagedOrder[supportlevelOrdered[i]] < stageSet.Stage.SupportLevel_stagedOrder[supportlevelOrdered[j]]
		})
		for _, supportlevel := range supportlevelOrdered {
			supportlevelIdent := "__stage_0" + supportlevel.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SupportLevel{Name: %s}).Stage(stageSet.Stage)", supportlevelIdent, __gong__toRawStringLiteral(supportlevel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", supportlevelIdent, __gong__toRawStringLiteral(supportlevel.Name)))
			if supportlevel.Tool != nil {
				targetIdent := "__stage_0" + supportlevel.Tool.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tool = %s", supportlevelIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		toolOrdered := []*Tool{}
		for tool := range stageSet.Stage.Tools {
			toolOrdered = append(toolOrdered, tool)
		}
		sort.Slice(toolOrdered, func(i, j int) bool {
			return stageSet.Stage.Tool_stagedOrder[toolOrdered[i]] < stageSet.Stage.Tool_stagedOrder[toolOrdered[j]]
		})
		for _, tool := range toolOrdered {
			toolIdent := "__stage_0" + tool.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tool{Name: %s}).Stage(stageSet.Stage)", toolIdent, __gong__toRawStringLiteral(tool.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", toolIdent, __gong__toRawStringLiteral(tool.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/dsm/capture/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

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

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "AnalysisNeed":
					if !preserveOrder {
						inst := (&AnalysisNeed{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AnalysisNeed)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Concept":
					if !preserveOrder {
						inst := (&Concept{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Concept)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ConceptShape":
					if !preserveOrder {
						inst := (&ConceptShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ConceptShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Concern":
					if !preserveOrder {
						inst := (&Concern{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Concern)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ConcernCompositionShape":
					if !preserveOrder {
						inst := (&ConcernCompositionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ConcernCompositionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ConcernInputShape":
					if !preserveOrder {
						inst := (&ConcernInputShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ConcernInputShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ConcernOutputShape":
					if !preserveOrder {
						inst := (&ConcernOutputShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ConcernOutputShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ConcernShape":
					if !preserveOrder {
						inst := (&ConcernShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ConcernShape)
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
				case "Deliverable":
					if !preserveOrder {
						inst := (&Deliverable{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Deliverable)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DeliverableCompositionShape":
					if !preserveOrder {
						inst := (&DeliverableCompositionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DeliverableCompositionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DeliverableConceptShape":
					if !preserveOrder {
						inst := (&DeliverableConceptShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DeliverableConceptShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DeliverableShape":
					if !preserveOrder {
						inst := (&DeliverableShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DeliverableShape)
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
				case "DiagramShape":
					if !preserveOrder {
						inst := (&DiagramShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DiagramShape)
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
				case "Note":
					if !preserveOrder {
						inst := (&Note{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Note)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteDeliverableShape":
					if !preserveOrder {
						inst := (&NoteDeliverableShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteDeliverableShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteShape":
					if !preserveOrder {
						inst := (&NoteShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteStakeholderShape":
					if !preserveOrder {
						inst := (&NoteStakeholderShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteStakeholderShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteTaskShape":
					if !preserveOrder {
						inst := (&NoteTaskShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteTaskShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Requirement":
					if !preserveOrder {
						inst := (&Requirement{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Requirement)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "RequirementShape":
					if !preserveOrder {
						inst := (&RequirementShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(RequirementShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Stakeholder":
					if !preserveOrder {
						inst := (&Stakeholder{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Stakeholder)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StakeholderCompositionShape":
					if !preserveOrder {
						inst := (&StakeholderCompositionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StakeholderCompositionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StakeholderConcernShape":
					if !preserveOrder {
						inst := (&StakeholderConcernShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StakeholderConcernShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StakeholderShape":
					if !preserveOrder {
						inst := (&StakeholderShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StakeholderShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SupportLevel":
					if !preserveOrder {
						inst := (&SupportLevel{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SupportLevel)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tool":
					if !preserveOrder {
						inst := (&Tool{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tool)
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tool); ok {
										inst.Tools = append(inst.Tools, typedTarget)
									}
								}
							}
						}
					}
				case *ConceptShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concept":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
									inst.Concept = typedTarget
								}
							}
						}
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
					}
				case *Concern:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IDAirbus":
						inst.IDAirbus = GongExtractString(rhs)
					case "Priority":
						inst.Priority = Priority(GongExtractString(rhs))
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SubConcerns":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.SubConcerns = append(inst.SubConcerns, typedTarget)
									}
								}
							}
						}
					case "Inputs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.Inputs = append(inst.Inputs, typedTarget)
									}
								}
							}
						}
					case "IsInputsNodeExpanded":
						inst.IsInputsNodeExpanded = GongExtractBool(rhs)
					case "Outputs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.Outputs = append(inst.Outputs, typedTarget)
									}
								}
							}
						}
					case "IsOutputsNodeExpanded":
						inst.IsOutputsNodeExpanded = GongExtractBool(rhs)
					case "IsWithCompletion":
						inst.IsWithCompletion = GongExtractBool(rhs)
					case "Completion":
						inst.Completion = CompletionEnum(GongExtractString(rhs))
					case "Requirements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Requirement); ok {
										inst.Requirements = append(inst.Requirements, typedTarget)
									}
								}
							}
						}
					}
				case *ConcernCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concern":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
									inst.Concern = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *ConcernInputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
									inst.Deliverable = typedTarget
								}
							}
						}
					case "Concern":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
									inst.Concern = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *ConcernOutputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concern":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
									inst.Concern = typedTarget
								}
							}
						}
					case "Deliverable":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
									inst.Deliverable = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *ConcernShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Concern":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
									inst.Concern = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.SubDeliverables = append(inst.SubDeliverables, typedTarget)
									}
								}
							}
						}
					case "IsProducersNodeExpanded":
						inst.IsProducersNodeExpanded = GongExtractBool(rhs)
					case "IsConsumersNodeExpanded":
						inst.IsConsumersNodeExpanded = GongExtractBool(rhs)
					case "Concepts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
										inst.Concepts = append(inst.Concepts, typedTarget)
									}
								}
							}
						}
					}
				case *DeliverableCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
									inst.Deliverable = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *DeliverableConceptShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
									inst.Deliverable = typedTarget
								}
							}
						}
					case "Concept":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
									inst.Concept = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *DeliverableShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Deliverable":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
									inst.Deliverable = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.ConcernsWhoseRequirementsNodeIsExpanded = append(inst.ConcernsWhoseRequirementsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsRequirementsNodeExpanded":
						inst.IsRequirementsNodeExpanded = GongExtractBool(rhs)
					case "IsConceptsNodeExpanded":
						inst.IsConceptsNodeExpanded = GongExtractBool(rhs)
					case "Deliverable_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DeliverableShape); ok {
										inst.Deliverable_Shapes = append(inst.Deliverable_Shapes, typedTarget)
									}
								}
							}
						}
					case "DeliverablesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.DeliverablesWhoseNodeIsExpanded = append(inst.DeliverablesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "DeliverablesWhoseConceptsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.DeliverablesWhoseConceptsNodeIsExpanded = append(inst.DeliverablesWhoseConceptsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsPBSNodeExpanded":
						inst.IsPBSNodeExpanded = GongExtractBool(rhs)
					case "DeliverableComposition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DeliverableCompositionShape); ok {
										inst.DeliverableComposition_Shapes = append(inst.DeliverableComposition_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsConcernsNodeExpanded":
						inst.IsConcernsNodeExpanded = GongExtractBool(rhs)
					case "Concern_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ConcernShape); ok {
										inst.Concern_Shapes = append(inst.Concern_Shapes, typedTarget)
									}
								}
							}
						}
					case "ConcernsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.ConcernsWhoseNodeIsExpanded = append(inst.ConcernsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ConcernsWhoseInputNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.ConcernsWhoseInputNodeIsExpanded = append(inst.ConcernsWhoseInputNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ConcernsWhoseStakeholderNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.ConcernsWhoseStakeholderNodeIsExpanded = append(inst.ConcernsWhoseStakeholderNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ConcernssWhoseOutputNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.ConcernssWhoseOutputNodeIsExpanded = append(inst.ConcernssWhoseOutputNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ConcernComposition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ConcernCompositionShape); ok {
										inst.ConcernComposition_Shapes = append(inst.ConcernComposition_Shapes, typedTarget)
									}
								}
							}
						}
					case "ConcernInputShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ConcernInputShape); ok {
										inst.ConcernInputShapes = append(inst.ConcernInputShapes, typedTarget)
									}
								}
							}
						}
					case "ConcernOutputShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ConcernOutputShape); ok {
										inst.ConcernOutputShapes = append(inst.ConcernOutputShapes, typedTarget)
									}
								}
							}
						}
					case "Note_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteShape); ok {
										inst.Note_Shapes = append(inst.Note_Shapes, typedTarget)
									}
								}
							}
						}
					case "NotesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
										inst.NotesWhoseNodeIsExpanded = append(inst.NotesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NoteDeliverableShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteDeliverableShape); ok {
										inst.NoteDeliverableShapes = append(inst.NoteDeliverableShapes, typedTarget)
									}
								}
							}
						}
					case "NoteTaskShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteTaskShape); ok {
										inst.NoteTaskShapes = append(inst.NoteTaskShapes, typedTarget)
									}
								}
							}
						}
					case "NoteResourceShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteStakeholderShape); ok {
										inst.NoteResourceShapes = append(inst.NoteResourceShapes, typedTarget)
									}
								}
							}
						}
					case "Stakeholder_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StakeholderShape); ok {
										inst.Stakeholder_Shapes = append(inst.Stakeholder_Shapes, typedTarget)
									}
								}
							}
						}
					case "ResourcesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
										inst.ResourcesWhoseNodeIsExpanded = append(inst.ResourcesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsStakeholdersNodeExpanded":
						inst.IsStakeholdersNodeExpanded = GongExtractBool(rhs)
					case "ResourceComposition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StakeholderCompositionShape); ok {
										inst.ResourceComposition_Shapes = append(inst.ResourceComposition_Shapes, typedTarget)
									}
								}
							}
						}
					case "StakeholderConcernShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StakeholderConcernShape); ok {
										inst.StakeholderConcernShapes = append(inst.StakeholderConcernShapes, typedTarget)
									}
								}
							}
						}
					case "Requirement_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*RequirementShape); ok {
										inst.Requirement_Shapes = append(inst.Requirement_Shapes, typedTarget)
									}
								}
							}
						}
					case "RequirementsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Requirement); ok {
										inst.RequirementsWhoseNodeIsExpanded = append(inst.RequirementsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Concept_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ConceptShape); ok {
										inst.Concept_Shapes = append(inst.Concept_Shapes, typedTarget)
									}
								}
							}
						}
					case "ConceptsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
										inst.ConceptsWhoseNodeIsExpanded = append(inst.ConceptsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ConceptsWhoseDeliverablesNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
										inst.ConceptsWhoseDeliverablesNodeIsExpanded = append(inst.ConceptsWhoseDeliverablesNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "DeliverableConceptShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DeliverableConceptShape); ok {
										inst.DeliverableConceptShapes = append(inst.DeliverableConceptShapes, typedTarget)
									}
								}
							}
						}
					case "Diagram_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramShape); ok {
										inst.Diagram_Shapes = append(inst.Diagram_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsDiagramsNodeExpanded":
						inst.IsDiagramsNodeExpanded = GongExtractBool(rhs)
					case "DiagramsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Diagram); ok {
										inst.DiagramsWhoseNodeIsExpanded = append(inst.DiagramsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					}
				case *DiagramShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Diagram":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Diagram); ok {
									inst.Diagram = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.RootDeliverables = append(inst.RootDeliverables, typedTarget)
									}
								}
							}
						}
					case "RootConcerns":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.RootConcerns = append(inst.RootConcerns, typedTarget)
									}
								}
							}
						}
					case "RootStakeholders":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
										inst.RootStakeholders = append(inst.RootStakeholders, typedTarget)
									}
								}
							}
						}
					case "RootRequirements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Requirement); ok {
										inst.RootRequirements = append(inst.RootRequirements, typedTarget)
									}
								}
							}
						}
					case "RootConcepts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
										inst.RootConcepts = append(inst.RootConcepts, typedTarget)
									}
								}
							}
						}
					case "AnalysisNeeds":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AnalysisNeed); ok {
										inst.AnalysisNeeds = append(inst.AnalysisNeeds, typedTarget)
									}
								}
							}
						}
					case "Notes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
										inst.Notes = append(inst.Notes, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
										inst.Deliverables = append(inst.Deliverables, typedTarget)
									}
								}
							}
						}
					case "Tasks":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.Tasks = append(inst.Tasks, typedTarget)
									}
								}
							}
						}
					case "Resources":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
										inst.Resources = append(inst.Resources, typedTarget)
									}
								}
							}
						}
					}
				case *NoteDeliverableShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
					case "Deliverable":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Deliverable); ok {
									inst.Deliverable = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *NoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
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
					}
				case *NoteStakeholderShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
					case "Stakeholder":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
									inst.Stakeholder = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *NoteTaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
					case "Task":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
									inst.Task = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *Requirement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SupportLevels":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SupportLevel); ok {
										inst.SupportLevels = append(inst.SupportLevels, typedTarget)
									}
								}
							}
						}
					case "Concepts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concept); ok {
										inst.Concepts = append(inst.Concepts, typedTarget)
									}
								}
							}
						}
					}
				case *RequirementShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Requirement":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Requirement); ok {
									inst.Requirement = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
										inst.Concerns = append(inst.Concerns, typedTarget)
									}
								}
							}
						}
					case "SubStakeholders":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
										inst.SubStakeholders = append(inst.SubStakeholders, typedTarget)
									}
								}
							}
						}
					}
				case *StakeholderCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Stakeholder":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
									inst.Stakeholder = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *StakeholderConcernShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Stakeholder":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
									inst.Stakeholder = typedTarget
								}
							}
						}
					case "Concern":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Concern); ok {
									inst.Concern = typedTarget
								}
							}
						}
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
				case *StakeholderShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Stakeholder":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stakeholder); ok {
									inst.Stakeholder = typedTarget
								}
							}
						}
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
					}
				case *SupportLevel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tool":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tool); ok {
									inst.Tool = typedTarget
								}
							}
						}
					}
				case *Tool:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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

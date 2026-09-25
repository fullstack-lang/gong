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
		for _, analysisneed := range __gong__sortStageSetInstances(stageSet.Stage.AnalysisNeeds, stageSet.Stage.AnalysisNeed_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			analysisneedIdent := "__models" + analysisneed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AnalysisNeed{Name: %s}).Stage(stageSet.Stage)", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", analysisneedIdent, analysisneed.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, concept := range __gong__sortStageSetInstances(stageSet.Stage.Concepts, stageSet.Stage.Concept_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			conceptIdent := "__models" + concept.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Concept{Name: %s}).Stage(stageSet.Stage)", conceptIdent, __gong__toRawStringLiteral(concept.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conceptIdent, __gong__toRawStringLiteral(concept.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", conceptIdent, __gong__toRawStringLiteral(concept.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", conceptIdent, concept.IsExpanded))
			for _, elem := range concept.Tools {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tools = append(%s.Tools, %s)", conceptIdent, conceptIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, conceptshape := range __gong__sortStageSetInstances(stageSet.Stage.ConceptShapes, stageSet.Stage.ConceptShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			conceptshapeIdent := "__models" + conceptshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConceptShape{Name: %s}).Stage(stageSet.Stage)", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", conceptshapeIdent, conceptshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", conceptshapeIdent, conceptshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", conceptshapeIdent, conceptshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", conceptshapeIdent, conceptshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", conceptshapeIdent, conceptshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", conceptshapeIdent, conceptshape.IsHidden))
			if conceptshape.Concept != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + conceptshape.Concept.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", conceptshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, concern := range __gong__sortStageSetInstances(stageSet.Stage.Concerns, stageSet.Stage.Concern_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concernIdent := "__models" + concern.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Concern{Name: %s}).Stage(stageSet.Stage)", concernIdent, __gong__toRawStringLiteral(concern.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubConcerns = append(%s.SubConcerns, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Inputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inputs = append(%s.Inputs, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Outputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Outputs = append(%s.Outputs, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Requirements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirements = append(%s.Requirements, %s)", concernIdent, concernIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, concerncompositionshape := range __gong__sortStageSetInstances(stageSet.Stage.ConcernCompositionShapes, stageSet.Stage.ConcernCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concerncompositionshapeIdent := "__models" + concerncompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernCompositionShape{Name: %s}).Stage(stageSet.Stage)", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concerncompositionshapeIdent, concerncompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concerncompositionshapeIdent, concerncompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(string(concerncompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(string(concerncompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concerncompositionshapeIdent, concerncompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concerncompositionshapeIdent, concerncompositionshape.IsHidden))
			if concerncompositionshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + concerncompositionshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concerncompositionshapeIdent, targetIdent))
			}
			for _, elem := range concerncompositionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concerncompositionshapeIdent, concerncompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, concerninputshape := range __gong__sortStageSetInstances(stageSet.Stage.ConcernInputShapes, stageSet.Stage.ConcernInputShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concerninputshapeIdent := "__models" + concerninputshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernInputShape{Name: %s}).Stage(stageSet.Stage)", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concerninputshapeIdent, concerninputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concerninputshapeIdent, concerninputshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", concerninputshapeIdent, __gong__toRawStringLiteral(string(concerninputshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", concerninputshapeIdent, __gong__toRawStringLiteral(string(concerninputshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concerninputshapeIdent, concerninputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concerninputshapeIdent, concerninputshape.IsHidden))
			if concerninputshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + concerninputshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", concerninputshapeIdent, targetIdent))
			}
			if concerninputshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + concerninputshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concerninputshapeIdent, targetIdent))
			}
			for _, elem := range concerninputshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concerninputshapeIdent, concerninputshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, concernoutputshape := range __gong__sortStageSetInstances(stageSet.Stage.ConcernOutputShapes, stageSet.Stage.ConcernOutputShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concernoutputshapeIdent := "__models" + concernoutputshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernOutputShape{Name: %s}).Stage(stageSet.Stage)", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concernoutputshapeIdent, concernoutputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concernoutputshapeIdent, concernoutputshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(string(concernoutputshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(string(concernoutputshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concernoutputshapeIdent, concernoutputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concernoutputshapeIdent, concernoutputshape.IsHidden))
			if concernoutputshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + concernoutputshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concernoutputshapeIdent, targetIdent))
			}
			if concernoutputshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + concernoutputshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", concernoutputshapeIdent, targetIdent))
			}
			for _, elem := range concernoutputshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concernoutputshapeIdent, concernoutputshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, concernshape := range __gong__sortStageSetInstances(stageSet.Stage.ConcernShapes, stageSet.Stage.ConcernShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concernshapeIdent := "__models" + concernshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernShape{Name: %s}).Stage(stageSet.Stage)", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", concernshapeIdent, concernshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", concernshapeIdent, concernshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", concernshapeIdent, concernshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", concernshapeIdent, concernshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", concernshapeIdent, concernshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concernshapeIdent, concernshape.IsHidden))
			if concernshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + concernshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concernshapeIdent, targetIdent))
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
		for _, deliverable := range __gong__sortStageSetInstances(stageSet.Stage.Deliverables, stageSet.Stage.Deliverable_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableIdent := "__models" + deliverable.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Deliverable{Name: %s}).Stage(stageSet.Stage)", deliverableIdent, __gong__toRawStringLiteral(deliverable.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", deliverableIdent, deliverable.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProducersNodeExpanded = %t", deliverableIdent, deliverable.IsProducersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConsumersNodeExpanded = %t", deliverableIdent, deliverable.IsConsumersNodeExpanded))
			for _, elem := range deliverable.SubDeliverables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubDeliverables = append(%s.SubDeliverables, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
			for _, elem := range deliverable.Concepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, deliverablecompositionshape := range __gong__sortStageSetInstances(stageSet.Stage.DeliverableCompositionShapes, stageSet.Stage.DeliverableCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverablecompositionshapeIdent := "__models" + deliverablecompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableCompositionShape{Name: %s}).Stage(stageSet.Stage)", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(string(deliverablecompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(string(deliverablecompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverablecompositionshapeIdent, deliverablecompositionshape.IsHidden))
			if deliverablecompositionshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + deliverablecompositionshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverablecompositionshapeIdent, targetIdent))
			}
			for _, elem := range deliverablecompositionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverablecompositionshapeIdent, deliverablecompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, deliverableconceptshape := range __gong__sortStageSetInstances(stageSet.Stage.DeliverableConceptShapes, stageSet.Stage.DeliverableConceptShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableconceptshapeIdent := "__models" + deliverableconceptshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableConceptShape{Name: %s}).Stage(stageSet.Stage)", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(string(deliverableconceptshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(string(deliverableconceptshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableconceptshapeIdent, deliverableconceptshape.IsHidden))
			if deliverableconceptshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + deliverableconceptshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableconceptshapeIdent, targetIdent))
			}
			if deliverableconceptshape.Concept != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + deliverableconceptshape.Concept.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", deliverableconceptshapeIdent, targetIdent))
			}
			for _, elem := range deliverableconceptshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverableconceptshapeIdent, deliverableconceptshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, deliverableshape := range __gong__sortStageSetInstances(stageSet.Stage.DeliverableShapes, stageSet.Stage.DeliverableShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableshapeIdent := "__models" + deliverableshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableShape{Name: %s}).Stage(stageSet.Stage)", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", deliverableshapeIdent, deliverableshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", deliverableshapeIdent, deliverableshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", deliverableshapeIdent, deliverableshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", deliverableshapeIdent, deliverableshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", deliverableshapeIdent, deliverableshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableshapeIdent, deliverableshape.IsHidden))
			if deliverableshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + deliverableshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableshapeIdent, targetIdent))
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseRequirementsNodeIsExpanded = append(%s.ConcernsWhoseRequirementsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Deliverable_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable_Shapes = append(%s.Deliverable_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseNodeIsExpanded = append(%s.DeliverablesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseConceptsNodeIsExpanded = append(%s.DeliverablesWhoseConceptsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableComposition_Shapes = append(%s.DeliverableComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Concern_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern_Shapes = append(%s.Concern_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseNodeIsExpanded = append(%s.ConcernsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseInputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseInputNodeIsExpanded = append(%s.ConcernsWhoseInputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseStakeholderNodeIsExpanded = append(%s.ConcernsWhoseStakeholderNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernssWhoseOutputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernssWhoseOutputNodeIsExpanded = append(%s.ConcernssWhoseOutputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernComposition_Shapes = append(%s.ConcernComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernInputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernInputShapes = append(%s.ConcernInputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernOutputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernOutputShapes = append(%s.ConcernOutputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteDeliverableShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteDeliverableShapes = append(%s.NoteDeliverableShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteTaskShapes = append(%s.NoteTaskShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteResourceShapes = append(%s.NoteResourceShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Stakeholder_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder_Shapes = append(%s.Stakeholder_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourceComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourceComposition_Shapes = append(%s.ResourceComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.StakeholderConcernShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StakeholderConcernShapes = append(%s.StakeholderConcernShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Requirement_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement_Shapes = append(%s.Requirement_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementsWhoseNodeIsExpanded = append(%s.RequirementsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Concept_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept_Shapes = append(%s.Concept_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseNodeIsExpanded = append(%s.ConceptsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseDeliverablesNodeIsExpanded = append(%s.ConceptsWhoseDeliverablesNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableConceptShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableConceptShapes = append(%s.DeliverableConceptShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Diagram_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagram_Shapes = append(%s.Diagram_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DiagramsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramsWhoseNodeIsExpanded = append(%s.DiagramsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, diagramshape := range __gong__sortStageSetInstances(stageSet.Stage.DiagramShapes, stageSet.Stage.DiagramShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramshapeIdent := "__models" + diagramshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramShape{Name: %s}).Stage(stageSet.Stage)", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramshapeIdent, diagramshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", diagramshapeIdent, diagramshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", diagramshapeIdent, diagramshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramshapeIdent, diagramshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramshapeIdent, diagramshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", diagramshapeIdent, diagramshape.IsHidden))
			if diagramshape.Diagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + diagramshape.Diagram.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagram = %s", diagramshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			for _, elem := range library.RootDeliverables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDeliverables = append(%s.RootDeliverables, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootConcerns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootConcerns = append(%s.RootConcerns, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootStakeholders {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStakeholders = append(%s.RootStakeholders, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootRequirements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootRequirements = append(%s.RootRequirements, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootConcepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootConcepts = append(%s.RootConcepts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.AnalysisNeeds {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnalysisNeeds = append(%s.AnalysisNeeds, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Notes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notes = append(%s.Notes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", libraryIdent, libraryIdent, targetIdent))
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
		}
	}
	if stageSet.Stage != nil {
		for _, note := range __gong__sortStageSetInstances(stageSet.Stage.Notes, stageSet.Stage.Note_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteIdent := "__models" + note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note{Name: %s}).Stage(stageSet.Stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			for _, elem := range note.Deliverables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverables = append(%s.Deliverables, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Resources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resources = append(%s.Resources, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notedeliverableshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteDeliverableShapes, stageSet.Stage.NoteDeliverableShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notedeliverableshapeIdent := "__models" + notedeliverableshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteDeliverableShape{Name: %s}).Stage(stageSet.Stage)", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notedeliverableshapeIdent, notedeliverableshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notedeliverableshapeIdent, notedeliverableshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(string(notedeliverableshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(string(notedeliverableshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notedeliverableshapeIdent, notedeliverableshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notedeliverableshapeIdent, notedeliverableshape.IsHidden))
			if notedeliverableshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notedeliverableshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notedeliverableshapeIdent, targetIdent))
			}
			if notedeliverableshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notedeliverableshape.Deliverable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", notedeliverableshapeIdent, targetIdent))
			}
			for _, elem := range notedeliverableshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notedeliverableshapeIdent, notedeliverableshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, noteshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteShapes, stageSet.Stage.NoteShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteshapeIdent := "__models" + noteshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteShape{Name: %s}).Stage(stageSet.Stage)", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteshapeIdent, noteshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", noteshapeIdent, noteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", noteshapeIdent, noteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", noteshapeIdent, noteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", noteshapeIdent, noteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteshapeIdent, noteshape.IsHidden))
			if noteshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notestakeholdershape := range __gong__sortStageSetInstances(stageSet.Stage.NoteStakeholderShapes, stageSet.Stage.NoteStakeholderShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notestakeholdershapeIdent := "__models" + notestakeholdershape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteStakeholderShape{Name: %s}).Stage(stageSet.Stage)", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notestakeholdershapeIdent, notestakeholdershape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notestakeholdershapeIdent, notestakeholdershape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(string(notestakeholdershape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(string(notestakeholdershape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notestakeholdershapeIdent, notestakeholdershape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notestakeholdershapeIdent, notestakeholdershape.IsHidden))
			if notestakeholdershape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notestakeholdershape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notestakeholdershapeIdent, targetIdent))
			}
			if notestakeholdershape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notestakeholdershape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", notestakeholdershapeIdent, targetIdent))
			}
			for _, elem := range notestakeholdershape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notestakeholdershapeIdent, notestakeholdershapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notetaskshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteTaskShapes, stageSet.Stage.NoteTaskShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notetaskshapeIdent := "__models" + notetaskshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteTaskShape{Name: %s}).Stage(stageSet.Stage)", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notetaskshapeIdent, notetaskshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notetaskshapeIdent, notetaskshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notetaskshapeIdent, __gong__toRawStringLiteral(string(notetaskshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notetaskshapeIdent, __gong__toRawStringLiteral(string(notetaskshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notetaskshapeIdent, notetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notetaskshapeIdent, notetaskshape.IsHidden))
			if notetaskshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notetaskshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notetaskshapeIdent, targetIdent))
			}
			if notetaskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notetaskshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", notetaskshapeIdent, targetIdent))
			}
			for _, elem := range notetaskshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notetaskshapeIdent, notetaskshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, requirement := range __gong__sortStageSetInstances(stageSet.Stage.Requirements, stageSet.Stage.Requirement_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementIdent := "__models" + requirement.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Requirement{Name: %s}).Stage(stageSet.Stage)", requirementIdent, __gong__toRawStringLiteral(requirement.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementIdent, __gong__toRawStringLiteral(requirement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", requirementIdent, __gong__toRawStringLiteral(requirement.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", requirementIdent, requirement.IsExpanded))
			for _, elem := range requirement.SupportLevels {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportLevels = append(%s.SupportLevels, %s)", requirementIdent, requirementIdent, targetIdent))
			}
			for _, elem := range requirement.Concepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", requirementIdent, requirementIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, requirementshape := range __gong__sortStageSetInstances(stageSet.Stage.RequirementShapes, stageSet.Stage.RequirementShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementshapeIdent := "__models" + requirementshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RequirementShape{Name: %s}).Stage(stageSet.Stage)", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", requirementshapeIdent, requirementshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", requirementshapeIdent, requirementshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", requirementshapeIdent, requirementshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", requirementshapeIdent, requirementshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", requirementshapeIdent, requirementshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", requirementshapeIdent, requirementshape.IsHidden))
			if requirementshape.Requirement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + requirementshape.Requirement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement = %s", requirementshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, stakeholder := range __gong__sortStageSetInstances(stageSet.Stage.Stakeholders, stageSet.Stage.Stakeholder_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholderIdent := "__models" + stakeholder.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stakeholder{Name: %s}).Stage(stageSet.Stage)", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDAirbus = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.IDAirbus)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stakeholderIdent, stakeholder.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Description)))
			for _, elem := range stakeholder.Concerns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concerns = append(%s.Concerns, %s)", stakeholderIdent, stakeholderIdent, targetIdent))
			}
			for _, elem := range stakeholder.SubStakeholders {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubStakeholders = append(%s.SubStakeholders, %s)", stakeholderIdent, stakeholderIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, stakeholdercompositionshape := range __gong__sortStageSetInstances(stageSet.Stage.StakeholderCompositionShapes, stageSet.Stage.StakeholderCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholdercompositionshapeIdent := "__models" + stakeholdercompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StakeholderCompositionShape{Name: %s}).Stage(stageSet.Stage)", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(string(stakeholdercompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(string(stakeholdercompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholdercompositionshapeIdent, stakeholdercompositionshape.IsHidden))
			if stakeholdercompositionshape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stakeholdercompositionshape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholdercompositionshapeIdent, targetIdent))
			}
			for _, elem := range stakeholdercompositionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", stakeholdercompositionshapeIdent, stakeholdercompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, stakeholderconcernshape := range __gong__sortStageSetInstances(stageSet.Stage.StakeholderConcernShapes, stageSet.Stage.StakeholderConcernShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholderconcernshapeIdent := "__models" + stakeholderconcernshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StakeholderConcernShape{Name: %s}).Stage(stageSet.Stage)", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(string(stakeholderconcernshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(string(stakeholderconcernshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholderconcernshapeIdent, stakeholderconcernshape.IsHidden))
			if stakeholderconcernshape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stakeholderconcernshape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholderconcernshapeIdent, targetIdent))
			}
			if stakeholderconcernshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stakeholderconcernshape.Concern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", stakeholderconcernshapeIdent, targetIdent))
			}
			for _, elem := range stakeholderconcernshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", stakeholderconcernshapeIdent, stakeholderconcernshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, stakeholdershape := range __gong__sortStageSetInstances(stageSet.Stage.StakeholderShapes, stageSet.Stage.StakeholderShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholdershapeIdent := "__models" + stakeholdershape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StakeholderShape{Name: %s}).Stage(stageSet.Stage)", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stakeholdershapeIdent, stakeholdershape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", stakeholdershapeIdent, stakeholdershape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", stakeholdershapeIdent, stakeholdershape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", stakeholdershapeIdent, stakeholdershape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", stakeholdershapeIdent, stakeholdershape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholdershapeIdent, stakeholdershape.IsHidden))
			if stakeholdershape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stakeholdershape.Stakeholder.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholdershapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, supportlevel := range __gong__sortStageSetInstances(stageSet.Stage.SupportLevels, stageSet.Stage.SupportLevel_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			supportlevelIdent := "__models" + supportlevel.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SupportLevel{Name: %s}).Stage(stageSet.Stage)", supportlevelIdent, __gong__toRawStringLiteral(supportlevel.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", supportlevelIdent, __gong__toRawStringLiteral(supportlevel.Name)))
			if supportlevel.Tool != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + supportlevel.Tool.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tool = %s", supportlevelIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tool := range __gong__sortStageSetInstances(stageSet.Stage.Tools, stageSet.Stage.Tool_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			toolIdent := "__models" + tool.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tool{Name: %s}).Stage(stageSet.Stage)", toolIdent, __gong__toRawStringLiteral(tool.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", toolIdent, __gong__toRawStringLiteral(tool.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
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
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "AnalysisNeed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AnalysisNeed), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Concept":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Concept), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ConceptShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConceptShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Concern":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Concern), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ConcernCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernCompositionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ConcernInputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernInputShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ConcernOutputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernOutputShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ConcernShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConcernShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ControlPointShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPointShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Deliverable":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Deliverable), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableCompositionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableConceptShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableConceptShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DeliverableShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeliverableShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DiagramShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteDeliverableShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteDeliverableShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteStakeholderShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteStakeholderShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteTaskShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Requirement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Requirement), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RequirementShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RequirementShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stakeholder":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stakeholder), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StakeholderCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StakeholderCompositionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StakeholderConcernShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StakeholderConcernShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StakeholderShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StakeholderShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SupportLevel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SupportLevel), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tool":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tool), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						inst.Completion = CompletionEnum(GongExtractString(rhs))
					case "Requirements":
						__gong__assignSliceOfPointers(&inst.Requirements, rhs, identifierMap)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
					case "IsConceptsNodeExpanded":
						inst.IsConceptsNodeExpanded = GongExtractBool(rhs)
					case "Deliverable_Shapes":
						__gong__assignSliceOfPointers(&inst.Deliverable_Shapes, rhs, identifierMap)
					case "DeliverablesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DeliverablesWhoseNodeIsExpanded, rhs, identifierMap)
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
					case "Concept_Shapes":
						__gong__assignSliceOfPointers(&inst.Concept_Shapes, rhs, identifierMap)
					case "ConceptsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConceptsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ConceptsWhoseDeliverablesNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ConceptsWhoseDeliverablesNodeIsExpanded, rhs, identifierMap)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
					case "Concepts":
						__gong__assignSliceOfPointers(&inst.Concepts, rhs, identifierMap)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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

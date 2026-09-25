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
		for _, allocatedresourceshape := range __gong__sortStageSetInstances(stageSet.Stage.AllocatedResourceShapes, stageSet.Stage.AllocatedResourceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedresourceshapeIdent := "__models" + allocatedresourceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedResourceShape{Name: %s}).Stage(stageSet.Stage)", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if allocatedresourceshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedresourceshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", allocatedresourceshapeIdent, targetIdent))
			}
			if allocatedresourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedresourceshape.Resource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", allocatedresourceshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, allocatedsystemshape := range __gong__sortStageSetInstances(stageSet.Stage.AllocatedSystemShapes, stageSet.Stage.AllocatedSystemShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedsystemshapeIdent := "__models" + allocatedsystemshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedSystemShape{Name: %s}).Stage(stageSet.Stage)", allocatedsystemshapeIdent, __gong__toRawStringLiteral(allocatedsystemshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedsystemshapeIdent, __gong__toRawStringLiteral(allocatedsystemshape.Name)))
			if allocatedsystemshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedsystemshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", allocatedsystemshapeIdent, targetIdent))
			}
			if allocatedsystemshape.System != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedsystemshape.System.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System = %s", allocatedsystemshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, controlflow := range __gong__sortStageSetInstances(stageSet.Stage.ControlFlows, stageSet.Stage.ControlFlow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlflowIdent := "__models" + controlflow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlFlow{Name: %s}).Stage(stageSet.Stage)", controlflowIdent, __gong__toRawStringLiteral(controlflow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", controlflowIdent, controlflow.IsExpanded))
			if controlflow.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + controlflow.Start.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", controlflowIdent, targetIdent))
			}
			if controlflow.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + controlflow.End.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", controlflowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, controlflowshape := range __gong__sortStageSetInstances(stageSet.Stage.ControlFlowShapes, stageSet.Stage.ControlFlowShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlflowshapeIdent := "__models" + controlflowshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlFlowShape{Name: %s}).Stage(stageSet.Stage)", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", controlflowshapeIdent, controlflowshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", controlflowshapeIdent, controlflowshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", controlflowshapeIdent, __gong__toRawStringLiteral(string(controlflowshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", controlflowshapeIdent, __gong__toRawStringLiteral(string(controlflowshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", controlflowshapeIdent, controlflowshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", controlflowshapeIdent, controlflowshape.IsHidden))
			if controlflowshape.ControlFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + controlflowshape.ControlFlow.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow = %s", controlflowshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, data := range __gong__sortStageSetInstances(stageSet.Stage.Datas, stageSet.Stage.Data_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dataIdent := "__models" + data.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Data{Name: %s}).Stage(stageSet.Stage)", dataIdent, __gong__toRawStringLiteral(data.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataIdent, __gong__toRawStringLiteral(data.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", dataIdent, __gong__toRawStringLiteral(data.Acronym)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dataIdent, __gong__toRawStringLiteral(data.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataIdent, __gong__toRawStringLiteral(data.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataIdent, data.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", dataIdent, __gong__toRawStringLiteral(data.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", dataIdent, data.InverseAppliedScaling))
		}
	}
	if stageSet.Stage != nil {
		for _, dataflow := range __gong__sortStageSetInstances(stageSet.Stage.DataFlows, stageSet.Stage.DataFlow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dataflowIdent := "__models" + dataflow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataFlow{Name: %s}).Stage(stageSet.Stage)", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", dataflowIdent, __gong__toRawStringLiteral(string(dataflow.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", dataflowIdent, dataflow.IsDatasNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataflowIdent, dataflow.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", dataflowIdent, __gong__toRawStringLiteral(string(dataflow.Type))))
			if dataflow.StartPort != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.StartPort.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPort = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndPort != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.EndPort.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPort = %s", dataflowIdent, targetIdent))
			}
			if dataflow.StartExternalPart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.StartExternalPart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartExternalPart = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndExternalPart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.EndExternalPart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndExternalPart = %s", dataflowIdent, targetIdent))
			}
			for _, elem := range dataflow.Datas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Datas = append(%s.Datas, %s)", dataflowIdent, dataflowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, dataflowshape := range __gong__sortStageSetInstances(stageSet.Stage.DataFlowShapes, stageSet.Stage.DataFlowShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dataflowshapeIdent := "__models" + dataflowshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataFlowShape{Name: %s}).Stage(stageSet.Stage)", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", dataflowshapeIdent, dataflowshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", dataflowshapeIdent, dataflowshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", dataflowshapeIdent, __gong__toRawStringLiteral(string(dataflowshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", dataflowshapeIdent, __gong__toRawStringLiteral(string(dataflowshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", dataflowshapeIdent, dataflowshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", dataflowshapeIdent, dataflowshape.IsHidden))
			if dataflowshape.DataFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflowshape.DataFlow.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", dataflowshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datashape := range __gong__sortStageSetInstances(stageSet.Stage.DataShapes, stageSet.Stage.DataShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datashapeIdent := "__models" + datashape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataShape{Name: %s}).Stage(stageSet.Stage)", datashapeIdent, __gong__toRawStringLiteral(datashape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datashapeIdent, __gong__toRawStringLiteral(datashape.Name)))
			if datashape.Data != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datashape.Data.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data = %s", datashapeIdent, targetIdent))
			}
			if datashape.DataFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datashape.DataFlow.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", datashapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, diagramlayerstate := range __gong__sortStageSetInstances(stageSet.Stage.DiagramLayerStates, stageSet.Stage.DiagramLayerState_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramlayerstateIdent := "__models" + diagramlayerstate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramLayerState{Name: %s}).Stage(stageSet.Stage)", diagramlayerstateIdent, __gong__toRawStringLiteral(diagramlayerstate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramlayerstateIdent, __gong__toRawStringLiteral(diagramlayerstate.Name)))
			if diagramlayerstate.DiagramStructure != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + diagramlayerstate.DiagramStructure.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructure = %s", diagramlayerstateIdent, targetIdent))
			}
			if diagramlayerstate.LayerDefinition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + diagramlayerstate.LayerDefinition.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LayerDefinition = %s", diagramlayerstateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, diagramstructure := range __gong__sortStageSetInstances(stageSet.Stage.DiagramStructures, stageSet.Stage.DiagramStructure_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramstructureIdent := "__models" + diagramstructure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramStructure{Name: %s}).Stage(stageSet.Stage)", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramstructureIdent, diagramstructure.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramstructureIdent, diagramstructure.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramstructureIdent, diagramstructure.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramstructureIdent, diagramstructure.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramstructureIdent, diagramstructure.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramstructureIdent, diagramstructure.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithDiscretePorts = %t", diagramstructureIdent, diagramstructure.IsWithDiscretePorts))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramstructureIdent, diagramstructure.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramstructureIdent, diagramstructure.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemsNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsSystemsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartsNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExternalPartsNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsExternalPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsNotesNodeExpanded))
			for _, elem := range diagramstructure.System_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_Shapes = append(%s.System_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.SystemsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Part_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_Shapes = append(%s.Part_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.PartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartWhoseNodeIsExpanded = append(%s.PartWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPart_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPart_Shapes = append(%s.ExternalPart_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartWhoseNodeIsExpanded = append(%s.ExternalPartWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(%s.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(%s.ExternalPartsWhoseInDataFlowsNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.PortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortsWhoseNodeIsExpanded = append(%s.PortsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Port_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port_Shapes = append(%s.Port_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlowsWhoseNodeIsExpanded = append(%s.ControlFlowsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ControlFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow_Shapes = append(%s.ControlFlow_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow_Shapes = append(%s.DataFlow_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Data_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data_Shapes = append(%s.Data_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseDataNodeIsExpanded = append(%s.DataFlowsWhoseDataNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourcesWhoseNodeIsExpanded = append(%s.AllocatedResourcesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourceShapes = append(%s.AllocatedResourceShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedSystemesWhoseNodeIsExpanded = append(%s.AllocatedSystemesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedSystemShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedSystemShapes = append(%s.AllocatedSystemShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotePortShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePortShapes = append(%s.NotePortShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotePartShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePartShapes = append(%s.NotePartShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, externalpartshape := range __gong__sortStageSetInstances(stageSet.Stage.ExternalPartShapes, stageSet.Stage.ExternalPartShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			externalpartshapeIdent := "__models" + externalpartshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ExternalPartShape{Name: %s}).Stage(stageSet.Stage)", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", externalpartshapeIdent, externalpartshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", externalpartshapeIdent, externalpartshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", externalpartshapeIdent, externalpartshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", externalpartshapeIdent, externalpartshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", externalpartshapeIdent, externalpartshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", externalpartshapeIdent, externalpartshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.TailHeigth = %f", externalpartshapeIdent, externalpartshape.TailHeigth))
			if externalpartshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + externalpartshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", externalpartshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, layerdefinition := range __gong__sortStageSetInstances(stageSet.Stage.LayerDefinitions, stageSet.Stage.LayerDefinition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layerdefinitionIdent := "__models" + layerdefinition.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LayerDefinition{Name: %s}).Stage(stageSet.Stage)", layerdefinitionIdent, __gong__toRawStringLiteral(layerdefinition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layerdefinitionIdent, __gong__toRawStringLiteral(layerdefinition.Name)))
			for _, elem := range layerdefinition.Query {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Query = append(%s.Query, %s)", layerdefinitionIdent, layerdefinitionIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemesNodeExpanded = %t", libraryIdent, library.IsSystemesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", libraryIdent, library.IsDataFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", libraryIdent, library.IsDatasNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResourcesNodeExpanded = %t", libraryIdent, library.IsResourcesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", libraryIdent, library.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
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
			for _, elem := range library.RootSystemes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootSystemes = append(%s.RootSystemes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SystemsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDataFlows = append(%s.RootDataFlows, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDatas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDatas = append(%s.RootDatas, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootResources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootResources = append(%s.RootResources, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.PartsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartsWhoseNodeIsExpanded = append(%s.PartsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootNotes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNotes = append(%s.RootNotes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", noteIdent, __gong__toRawStringLiteral(note.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartsNodeExpanded = %t", noteIdent, note.IsPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPortsNodeExpanded = %t", noteIdent, note.IsPortsNodeExpanded))
			for _, elem := range note.Parts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Ports {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ports = append(%s.Ports, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notepartshape := range __gong__sortStageSetInstances(stageSet.Stage.NotePartShapes, stageSet.Stage.NotePartShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notepartshapeIdent := "__models" + notepartshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NotePartShape{Name: %s}).Stage(stageSet.Stage)", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notepartshapeIdent, notepartshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notepartshapeIdent, notepartshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notepartshapeIdent, __gong__toRawStringLiteral(string(notepartshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notepartshapeIdent, __gong__toRawStringLiteral(string(notepartshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notepartshapeIdent, notepartshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notepartshapeIdent, notepartshape.IsHidden))
			if notepartshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notepartshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notepartshapeIdent, targetIdent))
			}
			if notepartshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notepartshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", notepartshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, noteportshape := range __gong__sortStageSetInstances(stageSet.Stage.NotePortShapes, stageSet.Stage.NotePortShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteportshapeIdent := "__models" + noteportshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NotePortShape{Name: %s}).Stage(stageSet.Stage)", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteportshapeIdent, noteportshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteportshapeIdent, noteportshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteportshapeIdent, __gong__toRawStringLiteral(string(noteportshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteportshapeIdent, __gong__toRawStringLiteral(string(noteportshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteportshapeIdent, noteportshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteportshapeIdent, noteportshape.IsHidden))
			if noteportshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteportshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteportshapeIdent, targetIdent))
			}
			if noteportshape.Port != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteportshape.Port.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port = %s", noteportshapeIdent, targetIdent))
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
		for _, part := range __gong__sortStageSetInstances(stageSet.Stage.Parts, stageSet.Stage.Part_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partIdent := "__models" + part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part{Name: %s}).Stage(stageSet.Stage)", partIdent, __gong__toRawStringLiteral(part.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partIdent, __gong__toRawStringLiteral(part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", partIdent, __gong__toRawStringLiteral(part.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartNameNotSystemName = %t", partIdent, part.IsPartNameNotSystemName))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlFlowsNodeExpanded = %t", partIdent, part.IsControlFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", partIdent, part.IsDataFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", partIdent, __gong__toRawStringLiteral(part.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", partIdent, part.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPortsNodeExpanded = %t", partIdent, part.IsPortsNodeExpanded))
			for _, elem := range part.Ports {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ports = append(%s.Ports, %s)", partIdent, partIdent, targetIdent))
			}
			if part.TypeOfPart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part.TypeOfPart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TypeOfPart = %s", partIdent, targetIdent))
			}
			for _, elem := range part.ControlFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlows = append(%s.ControlFlows, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseOutControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseOutControlFlowsNodeIsExpanded = append(%s.PortWhoseOutControlFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseInControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseInControlFlowsNodeIsExpanded = append(%s.PortWhoseInControlFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseOutDataFlowsNodeIsExpanded = append(%s.PortWhoseOutDataFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseInDataFlowsNodeIsExpanded = append(%s.PortWhoseInDataFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PartAnchoredPath {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAnchoredPath = append(%s.PartAnchoredPath, %s)", partIdent, partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, partanchoredpath := range __gong__sortStageSetInstances(stageSet.Stage.PartAnchoredPaths, stageSet.Stage.PartAnchoredPath_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partanchoredpathIdent := "__models" + partanchoredpath.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PartAnchoredPath{Name: %s}).Stage(stageSet.Stage)", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", partanchoredpathIdent, partanchoredpath.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", partanchoredpathIdent, partanchoredpath.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = %s", partanchoredpathIdent, __gong__toRawStringLiteral(string(partanchoredpath.RectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ScalePropotionnally = %t", partanchoredpathIdent, partanchoredpath.ScalePropotionnally))
			values.WriteString(fmt.Sprintf("\n\t%s.AppliedScaling = %f", partanchoredpathIdent, partanchoredpath.AppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", partanchoredpathIdent, partanchoredpath.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", partanchoredpathIdent, partanchoredpath.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", partanchoredpathIdent, partanchoredpath.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Transform)))
		}
	}
	if stageSet.Stage != nil {
		for _, partshape := range __gong__sortStageSetInstances(stageSet.Stage.PartShapes, stageSet.Stage.PartShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partshapeIdent := "__models" + partshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PartShape{Name: %s}).Stage(stageSet.Stage)", partshapeIdent, __gong__toRawStringLiteral(partshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", partshapeIdent, partshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", partshapeIdent, partshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", partshapeIdent, partshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", partshapeIdent, partshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", partshapeIdent, partshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", partshapeIdent, partshape.IsHidden))
			if partshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + partshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", partshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, port := range __gong__sortStageSetInstances(stageSet.Stage.Ports, stageSet.Stage.Port_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			portIdent := "__models" + port.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Port{Name: %s}).Stage(stageSet.Stage)", portIdent, __gong__toRawStringLiteral(port.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", portIdent, __gong__toRawStringLiteral(port.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", portIdent, __gong__toRawStringLiteral(port.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", portIdent, __gong__toRawStringLiteral(port.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", portIdent, port.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, portshape := range __gong__sortStageSetInstances(stageSet.Stage.PortShapes, stageSet.Stage.PortShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			portshapeIdent := "__models" + portshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PortShape{Name: %s}).Stage(stageSet.Stage)", portshapeIdent, __gong__toRawStringLiteral(portshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", portshapeIdent, portshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", portshapeIdent, portshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", portshapeIdent, portshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", portshapeIdent, portshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", portshapeIdent, portshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", portshapeIdent, portshape.IsHidden))
			if portshape.Port != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + portshape.Port.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port = %s", portshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, resource := range __gong__sortStageSetInstances(stageSet.Stage.Resources, stageSet.Stage.Resource_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourceIdent := "__models" + resource.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Resource{Name: %s}).Stage(stageSet.Stage)", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", resourceIdent, __gong__toRawStringLiteral(resource.Acronym)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", resourceIdent, __gong__toRawStringLiteral(resource.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", resourceIdent, __gong__toRawStringLiteral(resource.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", resourceIdent, resource.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", resourceIdent, __gong__toRawStringLiteral(resource.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", resourceIdent, resource.InverseAppliedScaling))
		}
	}
	if stageSet.Stage != nil {
		for _, semantictag := range __gong__sortStageSetInstances(stageSet.Stage.SemanticTags, stageSet.Stage.SemanticTag_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			semantictagIdent := "__models" + semantictag.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SemanticTag{Name: %s}).Stage(stageSet.Stage)", semantictagIdent, __gong__toRawStringLiteral(semantictag.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", semantictagIdent, __gong__toRawStringLiteral(semantictag.Name)))
			for _, elem := range semantictag.Parts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", semantictagIdent, semantictagIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, system := range __gong__sortStageSetInstances(stageSet.Stage.Systems, stageSet.Stage.System_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemIdent := "__models" + system.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.System{Name: %s}).Stage(stageSet.Stage)", systemIdent, __gong__toRawStringLiteral(system.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemIdent, __gong__toRawStringLiteral(system.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", systemIdent, __gong__toRawStringLiteral(system.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", systemIdent, __gong__toRawStringLiteral(system.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemIdent, system.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", systemIdent, __gong__toRawStringLiteral(system.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", systemIdent, system.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubSystemNodeExpanded = %t", systemIdent, system.IsSubSystemNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", systemIdent, system.IsDataFlowsNodeExpanded))
			for _, elem := range system.DiagramStructures {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructures = append(%s.DiagramStructures, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramStructureWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructureWhoseNodeIsExpanded = append(%s.DiagramStructureWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.SubSystemes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubSystemes = append(%s.SubSystemes, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Parts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.PartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartWhoseNodeIsExpanded = append(%s.PartWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlows = append(%s.DataFlows, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ExternalParts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParts = append(%s.ExternalParts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ExternalPartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartWhoseNodeIsExpanded = append(%s.ExternalPartWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, systemshape := range __gong__sortStageSetInstances(stageSet.Stage.SystemShapes, stageSet.Stage.SystemShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemshapeIdent := "__models" + systemshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SystemShape{Name: %s}).Stage(stageSet.Stage)", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemshapeIdent, systemshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", systemshapeIdent, systemshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", systemshapeIdent, systemshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", systemshapeIdent, systemshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", systemshapeIdent, systemshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", systemshapeIdent, systemshape.IsHidden))
			if systemshape.System != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + systemshape.System.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System = %s", systemshapeIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
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
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "AllocatedResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AllocatedResourceShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AllocatedSystemShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AllocatedSystemShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ControlFlow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlFlow), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ControlFlowShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlFlowShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Data":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Data), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DataFlow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DataFlow), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DataFlowShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DataFlowShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DataShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DataShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DiagramLayerState":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramLayerState), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DiagramStructure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramStructure), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ExternalPartShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ExternalPartShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LayerDefinition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LayerDefinition), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NotePartShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePartShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NotePortShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePortShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartAnchoredPath":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartAnchoredPath), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Port":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Port), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PortShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PortShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Resource":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Resource), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SemanticTag":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SemanticTag), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "System":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SystemShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SystemShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.Direction = DataFlowDirection(GongExtractString(rhs))
					case "IsDatasNodeExpanded":
						inst.IsDatasNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Type":
						inst.Type = DataFlowType(GongExtractString(rhs))
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
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
						inst.RectAnchorType = RectAnchorType(GongExtractString(rhs))
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

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
		allocatedresourceshapeOrdered := []*AllocatedResourceShape{}
		for allocatedresourceshape := range stageSet.Stage.AllocatedResourceShapes {
			allocatedresourceshapeOrdered = append(allocatedresourceshapeOrdered, allocatedresourceshape)
		}
		sort.Slice(allocatedresourceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.AllocatedResourceShape_stagedOrder[allocatedresourceshapeOrdered[i]] < stageSet.Stage.AllocatedResourceShape_stagedOrder[allocatedresourceshapeOrdered[j]]
		})
		for _, allocatedresourceshape := range allocatedresourceshapeOrdered {
			allocatedresourceshapeIdent := "__stage_0" + allocatedresourceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AllocatedResourceShape{Name: %s}).Stage(stageSet.Stage)", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if allocatedresourceshape.Part != nil {
				targetIdent := "__stage_0" + allocatedresourceshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", allocatedresourceshapeIdent, targetIdent))
			}
			if allocatedresourceshape.Resource != nil {
				targetIdent := "__stage_0" + allocatedresourceshape.Resource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", allocatedresourceshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		allocatedsystemshapeOrdered := []*AllocatedSystemShape{}
		for allocatedsystemshape := range stageSet.Stage.AllocatedSystemShapes {
			allocatedsystemshapeOrdered = append(allocatedsystemshapeOrdered, allocatedsystemshape)
		}
		sort.Slice(allocatedsystemshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.AllocatedSystemShape_stagedOrder[allocatedsystemshapeOrdered[i]] < stageSet.Stage.AllocatedSystemShape_stagedOrder[allocatedsystemshapeOrdered[j]]
		})
		for _, allocatedsystemshape := range allocatedsystemshapeOrdered {
			allocatedsystemshapeIdent := "__stage_0" + allocatedsystemshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AllocatedSystemShape{Name: %s}).Stage(stageSet.Stage)", allocatedsystemshapeIdent, __gong__toRawStringLiteral(allocatedsystemshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedsystemshapeIdent, __gong__toRawStringLiteral(allocatedsystemshape.Name)))
			if allocatedsystemshape.Part != nil {
				targetIdent := "__stage_0" + allocatedsystemshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", allocatedsystemshapeIdent, targetIdent))
			}
			if allocatedsystemshape.System != nil {
				targetIdent := "__stage_0" + allocatedsystemshape.System.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System = %s", allocatedsystemshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		controlflowOrdered := []*ControlFlow{}
		for controlflow := range stageSet.Stage.ControlFlows {
			controlflowOrdered = append(controlflowOrdered, controlflow)
		}
		sort.Slice(controlflowOrdered, func(i, j int) bool {
			return stageSet.Stage.ControlFlow_stagedOrder[controlflowOrdered[i]] < stageSet.Stage.ControlFlow_stagedOrder[controlflowOrdered[j]]
		})
		for _, controlflow := range controlflowOrdered {
			controlflowIdent := "__stage_0" + controlflow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ControlFlow{Name: %s}).Stage(stageSet.Stage)", controlflowIdent, __gong__toRawStringLiteral(controlflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", controlflowIdent, controlflow.IsExpanded))
			if controlflow.Start != nil {
				targetIdent := "__stage_0" + controlflow.Start.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", controlflowIdent, targetIdent))
			}
			if controlflow.End != nil {
				targetIdent := "__stage_0" + controlflow.End.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", controlflowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		controlflowshapeOrdered := []*ControlFlowShape{}
		for controlflowshape := range stageSet.Stage.ControlFlowShapes {
			controlflowshapeOrdered = append(controlflowshapeOrdered, controlflowshape)
		}
		sort.Slice(controlflowshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ControlFlowShape_stagedOrder[controlflowshapeOrdered[i]] < stageSet.Stage.ControlFlowShape_stagedOrder[controlflowshapeOrdered[j]]
		})
		for _, controlflowshape := range controlflowshapeOrdered {
			controlflowshapeIdent := "__stage_0" + controlflowshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ControlFlowShape{Name: %s}).Stage(stageSet.Stage)", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", controlflowshapeIdent, controlflowshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", controlflowshapeIdent, controlflowshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", controlflowshapeIdent, __gong__toRawStringLiteral(string(controlflowshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", controlflowshapeIdent, __gong__toRawStringLiteral(string(controlflowshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", controlflowshapeIdent, controlflowshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", controlflowshapeIdent, controlflowshape.IsHidden))
			if controlflowshape.ControlFlow != nil {
				targetIdent := "__stage_0" + controlflowshape.ControlFlow.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow = %s", controlflowshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		dataOrdered := []*Data{}
		for data := range stageSet.Stage.Datas {
			dataOrdered = append(dataOrdered, data)
		}
		sort.Slice(dataOrdered, func(i, j int) bool {
			return stageSet.Stage.Data_stagedOrder[dataOrdered[i]] < stageSet.Stage.Data_stagedOrder[dataOrdered[j]]
		})
		for _, data := range dataOrdered {
			dataIdent := "__stage_0" + data.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Data{Name: %s}).Stage(stageSet.Stage)", dataIdent, __gong__toRawStringLiteral(data.Name)))
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
		dataflowOrdered := []*DataFlow{}
		for dataflow := range stageSet.Stage.DataFlows {
			dataflowOrdered = append(dataflowOrdered, dataflow)
		}
		sort.Slice(dataflowOrdered, func(i, j int) bool {
			return stageSet.Stage.DataFlow_stagedOrder[dataflowOrdered[i]] < stageSet.Stage.DataFlow_stagedOrder[dataflowOrdered[j]]
		})
		for _, dataflow := range dataflowOrdered {
			dataflowIdent := "__stage_0" + dataflow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DataFlow{Name: %s}).Stage(stageSet.Stage)", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", dataflowIdent, __gong__toRawStringLiteral(string(dataflow.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", dataflowIdent, dataflow.IsDatasNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataflowIdent, dataflow.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", dataflowIdent, __gong__toRawStringLiteral(string(dataflow.Type))))
			if dataflow.StartPort != nil {
				targetIdent := "__stage_0" + dataflow.StartPort.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPort = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndPort != nil {
				targetIdent := "__stage_0" + dataflow.EndPort.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPort = %s", dataflowIdent, targetIdent))
			}
			if dataflow.StartExternalPart != nil {
				targetIdent := "__stage_0" + dataflow.StartExternalPart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartExternalPart = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndExternalPart != nil {
				targetIdent := "__stage_0" + dataflow.EndExternalPart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndExternalPart = %s", dataflowIdent, targetIdent))
			}
			for _, elem := range dataflow.Datas {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Datas = append(%s.Datas, %s)", dataflowIdent, dataflowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		dataflowshapeOrdered := []*DataFlowShape{}
		for dataflowshape := range stageSet.Stage.DataFlowShapes {
			dataflowshapeOrdered = append(dataflowshapeOrdered, dataflowshape)
		}
		sort.Slice(dataflowshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DataFlowShape_stagedOrder[dataflowshapeOrdered[i]] < stageSet.Stage.DataFlowShape_stagedOrder[dataflowshapeOrdered[j]]
		})
		for _, dataflowshape := range dataflowshapeOrdered {
			dataflowshapeIdent := "__stage_0" + dataflowshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DataFlowShape{Name: %s}).Stage(stageSet.Stage)", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", dataflowshapeIdent, dataflowshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", dataflowshapeIdent, dataflowshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", dataflowshapeIdent, __gong__toRawStringLiteral(string(dataflowshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", dataflowshapeIdent, __gong__toRawStringLiteral(string(dataflowshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", dataflowshapeIdent, dataflowshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", dataflowshapeIdent, dataflowshape.IsHidden))
			if dataflowshape.DataFlow != nil {
				targetIdent := "__stage_0" + dataflowshape.DataFlow.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", dataflowshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		datashapeOrdered := []*DataShape{}
		for datashape := range stageSet.Stage.DataShapes {
			datashapeOrdered = append(datashapeOrdered, datashape)
		}
		sort.Slice(datashapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DataShape_stagedOrder[datashapeOrdered[i]] < stageSet.Stage.DataShape_stagedOrder[datashapeOrdered[j]]
		})
		for _, datashape := range datashapeOrdered {
			datashapeIdent := "__stage_0" + datashape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DataShape{Name: %s}).Stage(stageSet.Stage)", datashapeIdent, __gong__toRawStringLiteral(datashape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datashapeIdent, __gong__toRawStringLiteral(datashape.Name)))
			if datashape.Data != nil {
				targetIdent := "__stage_0" + datashape.Data.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data = %s", datashapeIdent, targetIdent))
			}
			if datashape.DataFlow != nil {
				targetIdent := "__stage_0" + datashape.DataFlow.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", datashapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		diagramlayerstateOrdered := []*DiagramLayerState{}
		for diagramlayerstate := range stageSet.Stage.DiagramLayerStates {
			diagramlayerstateOrdered = append(diagramlayerstateOrdered, diagramlayerstate)
		}
		sort.Slice(diagramlayerstateOrdered, func(i, j int) bool {
			return stageSet.Stage.DiagramLayerState_stagedOrder[diagramlayerstateOrdered[i]] < stageSet.Stage.DiagramLayerState_stagedOrder[diagramlayerstateOrdered[j]]
		})
		for _, diagramlayerstate := range diagramlayerstateOrdered {
			diagramlayerstateIdent := "__stage_0" + diagramlayerstate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DiagramLayerState{Name: %s}).Stage(stageSet.Stage)", diagramlayerstateIdent, __gong__toRawStringLiteral(diagramlayerstate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramlayerstateIdent, __gong__toRawStringLiteral(diagramlayerstate.Name)))
			if diagramlayerstate.DiagramStructure != nil {
				targetIdent := "__stage_0" + diagramlayerstate.DiagramStructure.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructure = %s", diagramlayerstateIdent, targetIdent))
			}
			if diagramlayerstate.LayerDefinition != nil {
				targetIdent := "__stage_0" + diagramlayerstate.LayerDefinition.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LayerDefinition = %s", diagramlayerstateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		diagramstructureOrdered := []*DiagramStructure{}
		for diagramstructure := range stageSet.Stage.DiagramStructures {
			diagramstructureOrdered = append(diagramstructureOrdered, diagramstructure)
		}
		sort.Slice(diagramstructureOrdered, func(i, j int) bool {
			return stageSet.Stage.DiagramStructure_stagedOrder[diagramstructureOrdered[i]] < stageSet.Stage.DiagramStructure_stagedOrder[diagramstructureOrdered[j]]
		})
		for _, diagramstructure := range diagramstructureOrdered {
			diagramstructureIdent := "__stage_0" + diagramstructure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DiagramStructure{Name: %s}).Stage(stageSet.Stage)", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_Shapes = append(%s.System_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.SystemsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Part_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_Shapes = append(%s.Part_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.PartWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartWhoseNodeIsExpanded = append(%s.PartWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPart_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPart_Shapes = append(%s.ExternalPart_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartWhoseNodeIsExpanded = append(%s.ExternalPartWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(%s.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(%s.ExternalPartsWhoseInDataFlowsNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.PortsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortsWhoseNodeIsExpanded = append(%s.PortsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Port_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port_Shapes = append(%s.Port_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlowsWhoseNodeIsExpanded = append(%s.ControlFlowsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ControlFlow_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow_Shapes = append(%s.ControlFlow_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlow_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow_Shapes = append(%s.DataFlow_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DatasWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Data_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data_Shapes = append(%s.Data_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseDataNodeIsExpanded = append(%s.DataFlowsWhoseDataNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourcesWhoseNodeIsExpanded = append(%s.AllocatedResourcesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedResourceShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourceShapes = append(%s.AllocatedResourceShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedSystemesWhoseNodeIsExpanded = append(%s.AllocatedSystemesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedSystemShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedSystemShapes = append(%s.AllocatedSystemShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Note_Shapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotePortShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePortShapes = append(%s.NotePortShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotePartShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePartShapes = append(%s.NotePartShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		externalpartshapeOrdered := []*ExternalPartShape{}
		for externalpartshape := range stageSet.Stage.ExternalPartShapes {
			externalpartshapeOrdered = append(externalpartshapeOrdered, externalpartshape)
		}
		sort.Slice(externalpartshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ExternalPartShape_stagedOrder[externalpartshapeOrdered[i]] < stageSet.Stage.ExternalPartShape_stagedOrder[externalpartshapeOrdered[j]]
		})
		for _, externalpartshape := range externalpartshapeOrdered {
			externalpartshapeIdent := "__stage_0" + externalpartshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ExternalPartShape{Name: %s}).Stage(stageSet.Stage)", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", externalpartshapeIdent, externalpartshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", externalpartshapeIdent, externalpartshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", externalpartshapeIdent, externalpartshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", externalpartshapeIdent, externalpartshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", externalpartshapeIdent, externalpartshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", externalpartshapeIdent, externalpartshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.TailHeigth = %f", externalpartshapeIdent, externalpartshape.TailHeigth))
			if externalpartshape.Part != nil {
				targetIdent := "__stage_0" + externalpartshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", externalpartshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		layerdefinitionOrdered := []*LayerDefinition{}
		for layerdefinition := range stageSet.Stage.LayerDefinitions {
			layerdefinitionOrdered = append(layerdefinitionOrdered, layerdefinition)
		}
		sort.Slice(layerdefinitionOrdered, func(i, j int) bool {
			return stageSet.Stage.LayerDefinition_stagedOrder[layerdefinitionOrdered[i]] < stageSet.Stage.LayerDefinition_stagedOrder[layerdefinitionOrdered[j]]
		})
		for _, layerdefinition := range layerdefinitionOrdered {
			layerdefinitionIdent := "__stage_0" + layerdefinition.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.LayerDefinition{Name: %s}).Stage(stageSet.Stage)", layerdefinitionIdent, __gong__toRawStringLiteral(layerdefinition.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layerdefinitionIdent, __gong__toRawStringLiteral(layerdefinition.Name)))
			for _, elem := range layerdefinition.Query {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Query = append(%s.Query, %s)", layerdefinitionIdent, layerdefinitionIdent, targetIdent))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootSystemes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootSystemes = append(%s.RootSystemes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SystemsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDataFlows {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDataFlows = append(%s.RootDataFlows, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DataFlowsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDatas {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDatas = append(%s.RootDatas, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DatasWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootResources {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootResources = append(%s.RootResources, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ResourcesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.PartsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartsWhoseNodeIsExpanded = append(%s.PartsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootNotes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNotes = append(%s.RootNotes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.NotesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", noteIdent, __gong__toRawStringLiteral(note.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartsNodeExpanded = %t", noteIdent, note.IsPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPortsNodeExpanded = %t", noteIdent, note.IsPortsNodeExpanded))
			for _, elem := range note.Parts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Ports {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ports = append(%s.Ports, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		notepartshapeOrdered := []*NotePartShape{}
		for notepartshape := range stageSet.Stage.NotePartShapes {
			notepartshapeOrdered = append(notepartshapeOrdered, notepartshape)
		}
		sort.Slice(notepartshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NotePartShape_stagedOrder[notepartshapeOrdered[i]] < stageSet.Stage.NotePartShape_stagedOrder[notepartshapeOrdered[j]]
		})
		for _, notepartshape := range notepartshapeOrdered {
			notepartshapeIdent := "__stage_0" + notepartshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NotePartShape{Name: %s}).Stage(stageSet.Stage)", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notepartshapeIdent, notepartshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notepartshapeIdent, notepartshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notepartshapeIdent, __gong__toRawStringLiteral(string(notepartshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notepartshapeIdent, __gong__toRawStringLiteral(string(notepartshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notepartshapeIdent, notepartshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notepartshapeIdent, notepartshape.IsHidden))
			if notepartshape.Note != nil {
				targetIdent := "__stage_0" + notepartshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notepartshapeIdent, targetIdent))
			}
			if notepartshape.Part != nil {
				targetIdent := "__stage_0" + notepartshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", notepartshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteportshapeOrdered := []*NotePortShape{}
		for noteportshape := range stageSet.Stage.NotePortShapes {
			noteportshapeOrdered = append(noteportshapeOrdered, noteportshape)
		}
		sort.Slice(noteportshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NotePortShape_stagedOrder[noteportshapeOrdered[i]] < stageSet.Stage.NotePortShape_stagedOrder[noteportshapeOrdered[j]]
		})
		for _, noteportshape := range noteportshapeOrdered {
			noteportshapeIdent := "__stage_0" + noteportshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NotePortShape{Name: %s}).Stage(stageSet.Stage)", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteportshapeIdent, noteportshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteportshapeIdent, noteportshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteportshapeIdent, __gong__toRawStringLiteral(string(noteportshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteportshapeIdent, __gong__toRawStringLiteral(string(noteportshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteportshapeIdent, noteportshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteportshapeIdent, noteportshape.IsHidden))
			if noteportshape.Note != nil {
				targetIdent := "__stage_0" + noteportshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteportshapeIdent, targetIdent))
			}
			if noteportshape.Port != nil {
				targetIdent := "__stage_0" + noteportshape.Port.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port = %s", noteportshapeIdent, targetIdent))
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
		partOrdered := []*Part{}
		for part := range stageSet.Stage.Parts {
			partOrdered = append(partOrdered, part)
		}
		sort.Slice(partOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_stagedOrder[partOrdered[i]] < stageSet.Stage.Part_stagedOrder[partOrdered[j]]
		})
		for _, part := range partOrdered {
			partIdent := "__stage_0" + part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part{Name: %s}).Stage(stageSet.Stage)", partIdent, __gong__toRawStringLiteral(part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partIdent, __gong__toRawStringLiteral(part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", partIdent, __gong__toRawStringLiteral(part.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartNameNotSystemName = %t", partIdent, part.IsPartNameNotSystemName))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlFlowsNodeExpanded = %t", partIdent, part.IsControlFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", partIdent, part.IsDataFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", partIdent, __gong__toRawStringLiteral(part.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", partIdent, part.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPortsNodeExpanded = %t", partIdent, part.IsPortsNodeExpanded))
			for _, elem := range part.Ports {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ports = append(%s.Ports, %s)", partIdent, partIdent, targetIdent))
			}
			if part.TypeOfPart != nil {
				targetIdent := "__stage_0" + part.TypeOfPart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TypeOfPart = %s", partIdent, targetIdent))
			}
			for _, elem := range part.ControlFlows {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlows = append(%s.ControlFlows, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseOutControlFlowsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseOutControlFlowsNodeIsExpanded = append(%s.PortWhoseOutControlFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseInControlFlowsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseInControlFlowsNodeIsExpanded = append(%s.PortWhoseInControlFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseOutDataFlowsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseOutDataFlowsNodeIsExpanded = append(%s.PortWhoseOutDataFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseInDataFlowsNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseInDataFlowsNodeIsExpanded = append(%s.PortWhoseInDataFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PartAnchoredPath {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAnchoredPath = append(%s.PartAnchoredPath, %s)", partIdent, partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		partanchoredpathOrdered := []*PartAnchoredPath{}
		for partanchoredpath := range stageSet.Stage.PartAnchoredPaths {
			partanchoredpathOrdered = append(partanchoredpathOrdered, partanchoredpath)
		}
		sort.Slice(partanchoredpathOrdered, func(i, j int) bool {
			return stageSet.Stage.PartAnchoredPath_stagedOrder[partanchoredpathOrdered[i]] < stageSet.Stage.PartAnchoredPath_stagedOrder[partanchoredpathOrdered[j]]
		})
		for _, partanchoredpath := range partanchoredpathOrdered {
			partanchoredpathIdent := "__stage_0" + partanchoredpath.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.PartAnchoredPath{Name: %s}).Stage(stageSet.Stage)", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Name)))
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
		partshapeOrdered := []*PartShape{}
		for partshape := range stageSet.Stage.PartShapes {
			partshapeOrdered = append(partshapeOrdered, partshape)
		}
		sort.Slice(partshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.PartShape_stagedOrder[partshapeOrdered[i]] < stageSet.Stage.PartShape_stagedOrder[partshapeOrdered[j]]
		})
		for _, partshape := range partshapeOrdered {
			partshapeIdent := "__stage_0" + partshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.PartShape{Name: %s}).Stage(stageSet.Stage)", partshapeIdent, __gong__toRawStringLiteral(partshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", partshapeIdent, partshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", partshapeIdent, partshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", partshapeIdent, partshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", partshapeIdent, partshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", partshapeIdent, partshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", partshapeIdent, partshape.IsHidden))
			if partshape.Part != nil {
				targetIdent := "__stage_0" + partshape.Part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", partshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		portOrdered := []*Port{}
		for port := range stageSet.Stage.Ports {
			portOrdered = append(portOrdered, port)
		}
		sort.Slice(portOrdered, func(i, j int) bool {
			return stageSet.Stage.Port_stagedOrder[portOrdered[i]] < stageSet.Stage.Port_stagedOrder[portOrdered[j]]
		})
		for _, port := range portOrdered {
			portIdent := "__stage_0" + port.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Port{Name: %s}).Stage(stageSet.Stage)", portIdent, __gong__toRawStringLiteral(port.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", portIdent, __gong__toRawStringLiteral(port.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", portIdent, __gong__toRawStringLiteral(port.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", portIdent, __gong__toRawStringLiteral(port.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", portIdent, port.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		portshapeOrdered := []*PortShape{}
		for portshape := range stageSet.Stage.PortShapes {
			portshapeOrdered = append(portshapeOrdered, portshape)
		}
		sort.Slice(portshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.PortShape_stagedOrder[portshapeOrdered[i]] < stageSet.Stage.PortShape_stagedOrder[portshapeOrdered[j]]
		})
		for _, portshape := range portshapeOrdered {
			portshapeIdent := "__stage_0" + portshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.PortShape{Name: %s}).Stage(stageSet.Stage)", portshapeIdent, __gong__toRawStringLiteral(portshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", portshapeIdent, portshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", portshapeIdent, portshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", portshapeIdent, portshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", portshapeIdent, portshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", portshapeIdent, portshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", portshapeIdent, portshape.IsHidden))
			if portshape.Port != nil {
				targetIdent := "__stage_0" + portshape.Port.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port = %s", portshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		resourceOrdered := []*Resource{}
		for resource := range stageSet.Stage.Resources {
			resourceOrdered = append(resourceOrdered, resource)
		}
		sort.Slice(resourceOrdered, func(i, j int) bool {
			return stageSet.Stage.Resource_stagedOrder[resourceOrdered[i]] < stageSet.Stage.Resource_stagedOrder[resourceOrdered[j]]
		})
		for _, resource := range resourceOrdered {
			resourceIdent := "__stage_0" + resource.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Resource{Name: %s}).Stage(stageSet.Stage)", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
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
		semantictagOrdered := []*SemanticTag{}
		for semantictag := range stageSet.Stage.SemanticTags {
			semantictagOrdered = append(semantictagOrdered, semantictag)
		}
		sort.Slice(semantictagOrdered, func(i, j int) bool {
			return stageSet.Stage.SemanticTag_stagedOrder[semantictagOrdered[i]] < stageSet.Stage.SemanticTag_stagedOrder[semantictagOrdered[j]]
		})
		for _, semantictag := range semantictagOrdered {
			semantictagIdent := "__stage_0" + semantictag.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SemanticTag{Name: %s}).Stage(stageSet.Stage)", semantictagIdent, __gong__toRawStringLiteral(semantictag.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", semantictagIdent, __gong__toRawStringLiteral(semantictag.Name)))
			for _, elem := range semantictag.Parts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", semantictagIdent, semantictagIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		systemOrdered := []*System{}
		for system := range stageSet.Stage.Systems {
			systemOrdered = append(systemOrdered, system)
		}
		sort.Slice(systemOrdered, func(i, j int) bool {
			return stageSet.Stage.System_stagedOrder[systemOrdered[i]] < stageSet.Stage.System_stagedOrder[systemOrdered[j]]
		})
		for _, system := range systemOrdered {
			systemIdent := "__stage_0" + system.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.System{Name: %s}).Stage(stageSet.Stage)", systemIdent, __gong__toRawStringLiteral(system.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemIdent, __gong__toRawStringLiteral(system.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", systemIdent, __gong__toRawStringLiteral(system.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", systemIdent, __gong__toRawStringLiteral(system.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemIdent, system.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", systemIdent, __gong__toRawStringLiteral(system.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", systemIdent, system.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubSystemNodeExpanded = %t", systemIdent, system.IsSubSystemNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", systemIdent, system.IsDataFlowsNodeExpanded))
			for _, elem := range system.DiagramStructures {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructures = append(%s.DiagramStructures, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramStructureWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructureWhoseNodeIsExpanded = append(%s.DiagramStructureWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.SubSystemes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubSystemes = append(%s.SubSystemes, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Parts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.PartWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartWhoseNodeIsExpanded = append(%s.PartWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DataFlows {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlows = append(%s.DataFlows, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ExternalParts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParts = append(%s.ExternalParts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ExternalPartWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartWhoseNodeIsExpanded = append(%s.ExternalPartWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		systemshapeOrdered := []*SystemShape{}
		for systemshape := range stageSet.Stage.SystemShapes {
			systemshapeOrdered = append(systemshapeOrdered, systemshape)
		}
		sort.Slice(systemshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.SystemShape_stagedOrder[systemshapeOrdered[i]] < stageSet.Stage.SystemShape_stagedOrder[systemshapeOrdered[j]]
		})
		for _, systemshape := range systemshapeOrdered {
			systemshapeIdent := "__stage_0" + systemshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SystemShape{Name: %s}).Stage(stageSet.Stage)", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemshapeIdent, systemshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", systemshapeIdent, systemshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", systemshapeIdent, systemshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", systemshapeIdent, systemshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", systemshapeIdent, systemshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", systemshapeIdent, systemshape.IsHidden))
			if systemshape.System != nil {
				targetIdent := "__stage_0" + systemshape.System.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System = %s", systemshapeIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/dsm/structure/go/models"
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
				case "AllocatedResourceShape":
					if !preserveOrder {
						inst := (&AllocatedResourceShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AllocatedResourceShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "AllocatedSystemShape":
					if !preserveOrder {
						inst := (&AllocatedSystemShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AllocatedSystemShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ControlFlow":
					if !preserveOrder {
						inst := (&ControlFlow{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ControlFlow)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ControlFlowShape":
					if !preserveOrder {
						inst := (&ControlFlowShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ControlFlowShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Data":
					if !preserveOrder {
						inst := (&Data{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Data)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DataFlow":
					if !preserveOrder {
						inst := (&DataFlow{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DataFlow)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DataFlowShape":
					if !preserveOrder {
						inst := (&DataFlowShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DataFlowShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DataShape":
					if !preserveOrder {
						inst := (&DataShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DataShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DiagramLayerState":
					if !preserveOrder {
						inst := (&DiagramLayerState{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DiagramLayerState)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DiagramStructure":
					if !preserveOrder {
						inst := (&DiagramStructure{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DiagramStructure)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ExternalPartShape":
					if !preserveOrder {
						inst := (&ExternalPartShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ExternalPartShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "LayerDefinition":
					if !preserveOrder {
						inst := (&LayerDefinition{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(LayerDefinition)
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
				case "NotePartShape":
					if !preserveOrder {
						inst := (&NotePartShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NotePartShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NotePortShape":
					if !preserveOrder {
						inst := (&NotePortShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NotePortShape)
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
				case "Part":
					if !preserveOrder {
						inst := (&Part{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PartAnchoredPath":
					if !preserveOrder {
						inst := (&PartAnchoredPath{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PartAnchoredPath)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PartShape":
					if !preserveOrder {
						inst := (&PartShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PartShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Port":
					if !preserveOrder {
						inst := (&Port{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Port)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PortShape":
					if !preserveOrder {
						inst := (&PortShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PortShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Resource":
					if !preserveOrder {
						inst := (&Resource{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Resource)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SemanticTag":
					if !preserveOrder {
						inst := (&SemanticTag{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SemanticTag)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "System":
					if !preserveOrder {
						inst := (&System{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(System)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SystemShape":
					if !preserveOrder {
						inst := (&SystemShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SystemShape)
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
				case *AllocatedResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.Part = typedTarget
								}
							}
						}
					case "Resource":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
									inst.Resource = typedTarget
								}
							}
						}
					}
				case *AllocatedSystemShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.Part = typedTarget
								}
							}
						}
					case "System":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
									inst.System = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
									inst.Start = typedTarget
								}
							}
						}
					case "End":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
									inst.End = typedTarget
								}
							}
						}
					}
				case *ControlFlowShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ControlFlow":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ControlFlow); ok {
									inst.ControlFlow = typedTarget
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
									inst.StartPort = typedTarget
								}
							}
						}
					case "EndPort":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
									inst.EndPort = typedTarget
								}
							}
						}
					case "StartExternalPart":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.StartExternalPart = typedTarget
								}
							}
						}
					case "EndExternalPart":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.EndExternalPart = typedTarget
								}
							}
						}
					case "Datas":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Data); ok {
										inst.Datas = append(inst.Datas, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
									inst.DataFlow = typedTarget
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
					}
				case *DataShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Data":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Data); ok {
									inst.Data = typedTarget
								}
							}
						}
					case "DataFlow":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
									inst.DataFlow = typedTarget
								}
							}
						}
					}
				case *DiagramLayerState:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DiagramStructure":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramStructure); ok {
									inst.DiagramStructure = typedTarget
								}
							}
						}
					case "LayerDefinition":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*LayerDefinition); ok {
									inst.LayerDefinition = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SystemShape); ok {
										inst.System_Shapes = append(inst.System_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsSystemsNodeExpanded":
						inst.IsSystemsNodeExpanded = GongExtractBool(rhs)
					case "SystemsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.SystemsWhoseNodeIsExpanded = append(inst.SystemsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Part_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartShape); ok {
										inst.Part_Shapes = append(inst.Part_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsPartsNodeExpanded":
						inst.IsPartsNodeExpanded = GongExtractBool(rhs)
					case "PartWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.PartWhoseNodeIsExpanded = append(inst.PartWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ExternalPart_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ExternalPartShape); ok {
										inst.ExternalPart_Shapes = append(inst.ExternalPart_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsExternalPartsNodeExpanded":
						inst.IsExternalPartsNodeExpanded = GongExtractBool(rhs)
					case "ExternalPartWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.ExternalPartWhoseNodeIsExpanded = append(inst.ExternalPartWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(inst.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(inst.ExternalPartsWhoseInDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "PortsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.PortsWhoseNodeIsExpanded = append(inst.PortsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Port_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PortShape); ok {
										inst.Port_Shapes = append(inst.Port_Shapes, typedTarget)
									}
								}
							}
						}
					case "ControlFlowsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ControlFlow); ok {
										inst.ControlFlowsWhoseNodeIsExpanded = append(inst.ControlFlowsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ControlFlow_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ControlFlowShape); ok {
										inst.ControlFlow_Shapes = append(inst.ControlFlow_Shapes, typedTarget)
									}
								}
							}
						}
					case "DataFlowsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
										inst.DataFlowsWhoseNodeIsExpanded = append(inst.DataFlowsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "DataFlow_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlowShape); ok {
										inst.DataFlow_Shapes = append(inst.DataFlow_Shapes, typedTarget)
									}
								}
							}
						}
					case "DatasWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Data); ok {
										inst.DatasWhoseNodeIsExpanded = append(inst.DatasWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Data_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataShape); ok {
										inst.Data_Shapes = append(inst.Data_Shapes, typedTarget)
									}
								}
							}
						}
					case "DataFlowsWhoseDataNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
										inst.DataFlowsWhoseDataNodeIsExpanded = append(inst.DataFlowsWhoseDataNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "AllocatedResourcesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
										inst.AllocatedResourcesWhoseNodeIsExpanded = append(inst.AllocatedResourcesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "AllocatedResourceShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AllocatedResourceShape); ok {
										inst.AllocatedResourceShapes = append(inst.AllocatedResourceShapes, typedTarget)
									}
								}
							}
						}
					case "AllocatedSystemesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.AllocatedSystemesWhoseNodeIsExpanded = append(inst.AllocatedSystemesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "AllocatedSystemShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AllocatedSystemShape); ok {
										inst.AllocatedSystemShapes = append(inst.AllocatedSystemShapes, typedTarget)
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
					case "NotePortShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NotePortShape); ok {
										inst.NotePortShapes = append(inst.NotePortShapes, typedTarget)
									}
								}
							}
						}
					case "NotePartShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NotePartShape); ok {
										inst.NotePartShapes = append(inst.NotePartShapes, typedTarget)
									}
								}
							}
						}
					}
				case *ExternalPartShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.Part = typedTarget
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
					case "TailHeigth":
						inst.TailHeigth = GongExtractFloat(rhs)
					}
				case *LayerDefinition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Query":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SemanticTag); ok {
										inst.Query = append(inst.Query, typedTarget)
									}
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
					case "RootSystemes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.RootSystemes = append(inst.RootSystemes, typedTarget)
									}
								}
							}
						}
					case "IsSystemesNodeExpanded":
						inst.IsSystemesNodeExpanded = GongExtractBool(rhs)
					case "SystemsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.SystemsWhoseNodeIsExpanded = append(inst.SystemsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "RootDataFlows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
										inst.RootDataFlows = append(inst.RootDataFlows, typedTarget)
									}
								}
							}
						}
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "DataFlowsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
										inst.DataFlowsWhoseNodeIsExpanded = append(inst.DataFlowsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "RootDatas":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Data); ok {
										inst.RootDatas = append(inst.RootDatas, typedTarget)
									}
								}
							}
						}
					case "IsDatasNodeExpanded":
						inst.IsDatasNodeExpanded = GongExtractBool(rhs)
					case "DatasWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Data); ok {
										inst.DatasWhoseNodeIsExpanded = append(inst.DatasWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "RootResources":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
										inst.RootResources = append(inst.RootResources, typedTarget)
									}
								}
							}
						}
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "ResourcesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
										inst.ResourcesWhoseNodeIsExpanded = append(inst.ResourcesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "PartsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.PartsWhoseNodeIsExpanded = append(inst.PartsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "RootNotes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
										inst.RootNotes = append(inst.RootNotes, typedTarget)
									}
								}
							}
						}
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.Parts = append(inst.Parts, typedTarget)
									}
								}
							}
						}
					case "IsPortsNodeExpanded":
						inst.IsPortsNodeExpanded = GongExtractBool(rhs)
					case "Ports":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.Ports = append(inst.Ports, typedTarget)
									}
								}
							}
						}
					}
				case *NotePartShape:
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
					case "Part":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.Part = typedTarget
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
					}
				case *NotePortShape:
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
					case "Port":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
									inst.Port = typedTarget
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.Ports = append(inst.Ports, typedTarget)
									}
								}
							}
						}
					case "TypeOfPart":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
									inst.TypeOfPart = typedTarget
								}
							}
						}
					case "IsPartNameNotSystemName":
						inst.IsPartNameNotSystemName = GongExtractBool(rhs)
					case "IsControlFlowsNodeExpanded":
						inst.IsControlFlowsNodeExpanded = GongExtractBool(rhs)
					case "ControlFlows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ControlFlow); ok {
										inst.ControlFlows = append(inst.ControlFlows, typedTarget)
									}
								}
							}
						}
					case "PortWhoseOutControlFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.PortWhoseOutControlFlowsNodeIsExpanded = append(inst.PortWhoseOutControlFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "PortWhoseInControlFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.PortWhoseInControlFlowsNodeIsExpanded = append(inst.PortWhoseInControlFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "PortWhoseOutDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.PortWhoseOutDataFlowsNodeIsExpanded = append(inst.PortWhoseOutDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "PortWhoseInDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
										inst.PortWhoseInDataFlowsNodeIsExpanded = append(inst.PortWhoseInDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "PartAnchoredPath":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartAnchoredPath); ok {
										inst.PartAnchoredPath = append(inst.PartAnchoredPath, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
									inst.Part = typedTarget
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Port); ok {
									inst.Port = typedTarget
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.Parts = append(inst.Parts, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramStructure); ok {
										inst.DiagramStructures = append(inst.DiagramStructures, typedTarget)
									}
								}
							}
						}
					case "DiagramStructureWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramStructure); ok {
										inst.DiagramStructureWhoseNodeIsExpanded = append(inst.DiagramStructureWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsSubSystemNodeExpanded":
						inst.IsSubSystemNodeExpanded = GongExtractBool(rhs)
					case "SubSystemes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.SubSystemes = append(inst.SubSystemes, typedTarget)
									}
								}
							}
						}
					case "Parts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.Parts = append(inst.Parts, typedTarget)
									}
								}
							}
						}
					case "PartWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.PartWhoseNodeIsExpanded = append(inst.PartWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "DataFlows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DataFlow); ok {
										inst.DataFlows = append(inst.DataFlows, typedTarget)
									}
								}
							}
						}
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "ExternalParts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.ExternalParts = append(inst.ExternalParts, typedTarget)
									}
								}
							}
						}
					case "ExternalPartWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part); ok {
										inst.ExternalPartWhoseNodeIsExpanded = append(inst.ExternalPartWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					}
				case *SystemShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
									inst.System = typedTarget
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

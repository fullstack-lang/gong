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
		allocatedprocessshapeOrdered := []*AllocatedProcessShape{}
		for allocatedprocessshape := range stageSet.Stage.AllocatedProcessShapes {
			allocatedprocessshapeOrdered = append(allocatedprocessshapeOrdered, allocatedprocessshape)
		}
		sort.Slice(allocatedprocessshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.AllocatedProcessShape_stagedOrder[allocatedprocessshapeOrdered[i]] < stageSet.Stage.AllocatedProcessShape_stagedOrder[allocatedprocessshapeOrdered[j]]
		})
		for _, allocatedprocessshape := range allocatedprocessshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedprocessshapeIdent := "__models" + allocatedprocessshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedProcessShape{Name: %s}).Stage(stageSet.Stage)", allocatedprocessshapeIdent, __gong__toRawStringLiteral(allocatedprocessshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedprocessshapeIdent, __gong__toRawStringLiteral(allocatedprocessshape.Name)))
			if allocatedprocessshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedprocessshape.Participant.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", allocatedprocessshapeIdent, targetIdent))
			}
			if allocatedprocessshape.Process != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedprocessshape.Process.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Process = %s", allocatedprocessshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		allocatedresourceshapeOrdered := []*AllocatedResourceShape{}
		for allocatedresourceshape := range stageSet.Stage.AllocatedResourceShapes {
			allocatedresourceshapeOrdered = append(allocatedresourceshapeOrdered, allocatedresourceshape)
		}
		sort.Slice(allocatedresourceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.AllocatedResourceShape_stagedOrder[allocatedresourceshapeOrdered[i]] < stageSet.Stage.AllocatedResourceShape_stagedOrder[allocatedresourceshapeOrdered[j]]
		})
		for _, allocatedresourceshape := range allocatedresourceshapeOrdered {
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
			if allocatedresourceshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + allocatedresourceshape.Participant.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", allocatedresourceshapeIdent, targetIdent))
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
		controlflowOrdered := []*ControlFlow{}
		for controlflow := range stageSet.Stage.ControlFlows {
			controlflowOrdered = append(controlflowOrdered, controlflow)
		}
		sort.Slice(controlflowOrdered, func(i, j int) bool {
			return stageSet.Stage.ControlFlow_stagedOrder[controlflowOrdered[i]] < stageSet.Stage.ControlFlow_stagedOrder[controlflowOrdered[j]]
		})
		for _, controlflow := range controlflowOrdered {
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
		controlflowshapeOrdered := []*ControlFlowShape{}
		for controlflowshape := range stageSet.Stage.ControlFlowShapes {
			controlflowshapeOrdered = append(controlflowshapeOrdered, controlflowshape)
		}
		sort.Slice(controlflowshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ControlFlowShape_stagedOrder[controlflowshapeOrdered[i]] < stageSet.Stage.ControlFlowShape_stagedOrder[controlflowshapeOrdered[j]]
		})
		for _, controlflowshape := range controlflowshapeOrdered {
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
		dataOrdered := []*Data{}
		for data := range stageSet.Stage.Datas {
			dataOrdered = append(dataOrdered, data)
		}
		sort.Slice(dataOrdered, func(i, j int) bool {
			return stageSet.Stage.Data_stagedOrder[dataOrdered[i]] < stageSet.Stage.Data_stagedOrder[dataOrdered[j]]
		})
		for _, data := range dataOrdered {
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
		dataflowOrdered := []*DataFlow{}
		for dataflow := range stageSet.Stage.DataFlows {
			dataflowOrdered = append(dataflowOrdered, dataflow)
		}
		sort.Slice(dataflowOrdered, func(i, j int) bool {
			return stageSet.Stage.DataFlow_stagedOrder[dataflowOrdered[i]] < stageSet.Stage.DataFlow_stagedOrder[dataflowOrdered[j]]
		})
		for _, dataflow := range dataflowOrdered {
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
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataflowIdent, dataflow.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", dataflowIdent, __gong__toRawStringLiteral(string(dataflow.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", dataflowIdent, dataflow.IsDatasNodeExpanded))
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
			if dataflow.StartTask != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.StartTask.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartTask = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndTask != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.EndTask.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndTask = %s", dataflowIdent, targetIdent))
			}
			if dataflow.StartExternalParticipant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.StartExternalParticipant.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartExternalParticipant = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndExternalParticipant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dataflow.EndExternalParticipant.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndExternalParticipant = %s", dataflowIdent, targetIdent))
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
		datashapeOrdered := []*DataShape{}
		for datashape := range stageSet.Stage.DataShapes {
			datashapeOrdered = append(datashapeOrdered, datashape)
		}
		sort.Slice(datashapeOrdered, func(i, j int) bool {
			return stageSet.Stage.DataShape_stagedOrder[datashapeOrdered[i]] < stageSet.Stage.DataShape_stagedOrder[datashapeOrdered[j]]
		})
		for _, datashape := range datashapeOrdered {
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
		diagramprocessOrdered := []*DiagramProcess{}
		for diagramprocess := range stageSet.Stage.DiagramProcesss {
			diagramprocessOrdered = append(diagramprocessOrdered, diagramprocess)
		}
		sort.Slice(diagramprocessOrdered, func(i, j int) bool {
			return stageSet.Stage.DiagramProcess_stagedOrder[diagramprocessOrdered[i]] < stageSet.Stage.DiagramProcess_stagedOrder[diagramprocessOrdered[j]]
		})
		for _, diagramprocess := range diagramprocessOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramprocessIdent := "__models" + diagramprocess.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramProcess{Name: %s}).Stage(stageSet.Stage)", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramprocessIdent, diagramprocess.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramprocessIdent, diagramprocess.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramprocessIdent, diagramprocess.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramprocessIdent, diagramprocess.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramprocessIdent, diagramprocess.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramprocessIdent, diagramprocess.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramprocessIdent, diagramprocess.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramprocessIdent, diagramprocess.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcesssNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsProcesssNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParticipantsNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsParticipantsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExternalParticipantsNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsExternalParticipantsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsNotesNodeExpanded))
			for _, elem := range diagramprocess.Process_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Process_Shapes = append(%s.Process_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ProcesssWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcesssWhoseNodeIsExpanded = append(%s.ProcesssWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Participant_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant_Shapes = append(%s.Participant_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParticipantWhoseNodeIsExpanded = append(%s.ParticipantWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipant_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipant_Shapes = append(%s.ExternalParticipant_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantWhoseNodeIsExpanded = append(%s.ExternalParticipantWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(%s.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(%s.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.TasksWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseNodeIsExpanded = append(%s.TasksWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Task_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task_Shapes = append(%s.Task_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlowsWhoseNodeIsExpanded = append(%s.ControlFlowsWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ControlFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow_Shapes = append(%s.ControlFlow_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DataFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow_Shapes = append(%s.DataFlow_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Data_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data_Shapes = append(%s.Data_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseDataNodeIsExpanded = append(%s.DataFlowsWhoseDataNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourcesWhoseNodeIsExpanded = append(%s.AllocatedResourcesWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourceShapes = append(%s.AllocatedResourceShapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedProcessesWhoseNodeIsExpanded = append(%s.AllocatedProcessesWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedProcessShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedProcessShapes = append(%s.AllocatedProcessShapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.NoteTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteTaskShapes = append(%s.NoteTaskShapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		externalparticipantshapeOrdered := []*ExternalParticipantShape{}
		for externalparticipantshape := range stageSet.Stage.ExternalParticipantShapes {
			externalparticipantshapeOrdered = append(externalparticipantshapeOrdered, externalparticipantshape)
		}
		sort.Slice(externalparticipantshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ExternalParticipantShape_stagedOrder[externalparticipantshapeOrdered[i]] < stageSet.Stage.ExternalParticipantShape_stagedOrder[externalparticipantshapeOrdered[j]]
		})
		for _, externalparticipantshape := range externalparticipantshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			externalparticipantshapeIdent := "__models" + externalparticipantshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ExternalParticipantShape{Name: %s}).Stage(stageSet.Stage)", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", externalparticipantshapeIdent, externalparticipantshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", externalparticipantshapeIdent, externalparticipantshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", externalparticipantshapeIdent, externalparticipantshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", externalparticipantshapeIdent, externalparticipantshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", externalparticipantshapeIdent, externalparticipantshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", externalparticipantshapeIdent, externalparticipantshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.TailHeigth = %f", externalparticipantshapeIdent, externalparticipantshape.TailHeigth))
			if externalparticipantshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + externalparticipantshape.Participant.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", externalparticipantshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcessesNodeExpanded = %t", libraryIdent, library.IsProcessesNodeExpanded))
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
			for _, elem := range library.RootProcesses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootProcesses = append(%s.RootProcesses, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ProcesssWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcesssWhoseNodeIsExpanded = append(%s.ProcesssWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			for _, elem := range library.ParticipantsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParticipantsWhoseNodeIsExpanded = append(%s.ParticipantsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
		noteOrdered := []*Note{}
		for note := range stageSet.Stage.Notes {
			noteOrdered = append(noteOrdered, note)
		}
		sort.Slice(noteOrdered, func(i, j int) bool {
			return stageSet.Stage.Note_stagedOrder[noteOrdered[i]] < stageSet.Stage.Note_stagedOrder[noteOrdered[j]]
		})
		for _, note := range noteOrdered {
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsTasksNodeExpanded = %t", noteIdent, note.IsTasksNodeExpanded))
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
		notetaskshapeOrdered := []*NoteTaskShape{}
		for notetaskshape := range stageSet.Stage.NoteTaskShapes {
			notetaskshapeOrdered = append(notetaskshapeOrdered, notetaskshape)
		}
		sort.Slice(notetaskshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteTaskShape_stagedOrder[notetaskshapeOrdered[i]] < stageSet.Stage.NoteTaskShape_stagedOrder[notetaskshapeOrdered[j]]
		})
		for _, notetaskshape := range notetaskshapeOrdered {
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
		}
	}
	if stageSet.Stage != nil {
		participantOrdered := []*Participant{}
		for participant := range stageSet.Stage.Participants {
			participantOrdered = append(participantOrdered, participant)
		}
		sort.Slice(participantOrdered, func(i, j int) bool {
			return stageSet.Stage.Participant_stagedOrder[participantOrdered[i]] < stageSet.Stage.Participant_stagedOrder[participantOrdered[j]]
		})
		for _, participant := range participantOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			participantIdent := "__models" + participant.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Participant{Name: %s}).Stage(stageSet.Stage)", participantIdent, __gong__toRawStringLiteral(participant.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", participantIdent, __gong__toRawStringLiteral(participant.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcessResource = %t", participantIdent, participant.IsProcessResource))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", participantIdent, __gong__toRawStringLiteral(participant.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResourcesNodeExpanded = %t", participantIdent, participant.IsResourcesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcessesNodeExpanded = %t", participantIdent, participant.IsProcessesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", participantIdent, __gong__toRawStringLiteral(participant.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", participantIdent, participant.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTasksNodeExpanded = %t", participantIdent, participant.IsTasksNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlFlowsNodeExpanded = %t", participantIdent, participant.IsControlFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", participantIdent, participant.IsDataFlowsNodeExpanded))
			for _, elem := range participant.Resources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resources = append(%s.Resources, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.Processes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Processes = append(%s.Processes, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.ControlFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlows = append(%s.ControlFlows, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseOutControlFlowsNodeIsExpanded = append(%s.TaskWhoseOutControlFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseInControlFlowsNodeIsExpanded = append(%s.TaskWhoseInControlFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseOutDataFlowsNodeIsExpanded = append(%s.TaskWhoseOutDataFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseInDataFlowsNodeIsExpanded = append(%s.TaskWhoseInDataFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		participantshapeOrdered := []*ParticipantShape{}
		for participantshape := range stageSet.Stage.ParticipantShapes {
			participantshapeOrdered = append(participantshapeOrdered, participantshape)
		}
		sort.Slice(participantshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ParticipantShape_stagedOrder[participantshapeOrdered[i]] < stageSet.Stage.ParticipantShape_stagedOrder[participantshapeOrdered[j]]
		})
		for _, participantshape := range participantshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			participantshapeIdent := "__models" + participantshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParticipantShape{Name: %s}).Stage(stageSet.Stage)", participantshapeIdent, __gong__toRawStringLiteral(participantshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", participantshapeIdent, participantshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", participantshapeIdent, participantshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", participantshapeIdent, participantshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", participantshapeIdent, participantshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", participantshapeIdent, participantshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", participantshapeIdent, participantshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthWeight = %f", participantshapeIdent, participantshape.WidthWeight))
			if participantshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + participantshape.Participant.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", participantshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		processOrdered := []*Process{}
		for process := range stageSet.Stage.Processs {
			processOrdered = append(processOrdered, process)
		}
		sort.Slice(processOrdered, func(i, j int) bool {
			return stageSet.Stage.Process_stagedOrder[processOrdered[i]] < stageSet.Stage.Process_stagedOrder[processOrdered[j]]
		})
		for _, process := range processOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			processIdent := "__models" + process.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Process{Name: %s}).Stage(stageSet.Stage)", processIdent, __gong__toRawStringLiteral(process.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", processIdent, __gong__toRawStringLiteral(process.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", processIdent, __gong__toRawStringLiteral(process.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", processIdent, __gong__toRawStringLiteral(process.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", processIdent, process.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", processIdent, __gong__toRawStringLiteral(process.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", processIdent, process.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubProcessNodeExpanded = %t", processIdent, process.IsSubProcessNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", processIdent, process.IsDataFlowsNodeExpanded))
			for _, elem := range process.DiagramProcesss {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramProcesss = append(%s.DiagramProcesss, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.DiagramProcessWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramProcessWhoseNodeIsExpanded = append(%s.DiagramProcessWhoseNodeIsExpanded, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.SubProcesses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubProcesses = append(%s.SubProcesses, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.Participants {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participants = append(%s.Participants, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.ParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParticipantWhoseNodeIsExpanded = append(%s.ParticipantWhoseNodeIsExpanded, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.DataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlows = append(%s.DataFlows, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.ExternalParticipants {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipants = append(%s.ExternalParticipants, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.ExternalParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantWhoseNodeIsExpanded = append(%s.ExternalParticipantWhoseNodeIsExpanded, %s)", processIdent, processIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		processshapeOrdered := []*ProcessShape{}
		for processshape := range stageSet.Stage.ProcessShapes {
			processshapeOrdered = append(processshapeOrdered, processshape)
		}
		sort.Slice(processshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ProcessShape_stagedOrder[processshapeOrdered[i]] < stageSet.Stage.ProcessShape_stagedOrder[processshapeOrdered[j]]
		})
		for _, processshape := range processshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			processshapeIdent := "__models" + processshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProcessShape{Name: %s}).Stage(stageSet.Stage)", processshapeIdent, __gong__toRawStringLiteral(processshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", processshapeIdent, processshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", processshapeIdent, processshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", processshapeIdent, processshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", processshapeIdent, processshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", processshapeIdent, processshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", processshapeIdent, processshape.IsHidden))
			if processshape.Process != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + processshape.Process.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Process = %s", processshapeIdent, targetIdent))
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
		taskOrdered := []*Task{}
		for task := range stageSet.Stage.Tasks {
			taskOrdered = append(taskOrdered, task)
		}
		sort.Slice(taskOrdered, func(i, j int) bool {
			return stageSet.Stage.Task_stagedOrder[taskOrdered[i]] < stageSet.Stage.Task_stagedOrder[taskOrdered[j]]
		})
		for _, task := range taskOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskIdent := "__models" + task.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Task{Name: %s}).Stage(stageSet.Stage)", taskIdent, __gong__toRawStringLiteral(task.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskIdent, __gong__toRawStringLiteral(task.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", taskIdent, __gong__toRawStringLiteral(task.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskIdent, __gong__toRawStringLiteral(task.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskIdent, task.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartTask = %t", taskIdent, task.IsStartTask))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndTask = %t", taskIdent, task.IsEndTask))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTaskNameNotProcessName = %t", taskIdent, task.IsTaskNameNotProcessName))
			if task.Type != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + task.Type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type = %s", taskIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskshapeOrdered := []*TaskShape{}
		for taskshape := range stageSet.Stage.TaskShapes {
			taskshapeOrdered = append(taskshapeOrdered, taskshape)
		}
		sort.Slice(taskshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskShape_stagedOrder[taskshapeOrdered[i]] < stageSet.Stage.TaskShape_stagedOrder[taskshapeOrdered[j]]
		})
		for _, taskshape := range taskshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskshapeIdent := "__models" + taskshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskShape{Name: %s}).Stage(stageSet.Stage)", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskshapeIdent, taskshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", taskshapeIdent, taskshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", taskshapeIdent, taskshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", taskshapeIdent, taskshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", taskshapeIdent, taskshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskshapeIdent, taskshape.IsHidden))
			if taskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskshapeIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
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
		case "github.com/fullstack-lang/gong/dsm/process/go/models":
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
				case "AllocatedProcessShape":
					if !preserveOrder {
						inst := (&AllocatedProcessShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AllocatedProcessShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
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
				case "DiagramProcess":
					if !preserveOrder {
						inst := (&DiagramProcess{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DiagramProcess)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ExternalParticipantShape":
					if !preserveOrder {
						inst := (&ExternalParticipantShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ExternalParticipantShape)
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
				case "Participant":
					if !preserveOrder {
						inst := (&Participant{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Participant)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParticipantShape":
					if !preserveOrder {
						inst := (&ParticipantShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParticipantShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Process":
					if !preserveOrder {
						inst := (&Process{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Process)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ProcessShape":
					if !preserveOrder {
						inst := (&ProcessShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ProcessShape)
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
				case "Task":
					if !preserveOrder {
						inst := (&Task{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Task)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TaskShape":
					if !preserveOrder {
						inst := (&TaskShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskShape)
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
				case *AllocatedProcessShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Participant); ok {
									inst.Participant = typedTarget
								}
							}
						}
					case "Process":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Process); ok {
									inst.Process = typedTarget
								}
							}
						}
					}
				case *AllocatedResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Participant); ok {
									inst.Participant = typedTarget
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
								if typedTarget, ok := target.(*Task); ok {
									inst.Start = typedTarget
								}
							}
						}
					case "End":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Task); ok {
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
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Type":
						inst.Type = DataFlowType(GongExtractString(rhs))
					case "StartTask":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Task); ok {
									inst.StartTask = typedTarget
								}
							}
						}
					case "EndTask":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Task); ok {
									inst.EndTask = typedTarget
								}
							}
						}
					case "StartExternalParticipant":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Participant); ok {
									inst.StartExternalParticipant = typedTarget
								}
							}
						}
					case "EndExternalParticipant":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Participant); ok {
									inst.EndExternalParticipant = typedTarget
								}
							}
						}
					case "IsDatasNodeExpanded":
						inst.IsDatasNodeExpanded = GongExtractBool(rhs)
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
				case *DiagramProcess:
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
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "Process_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ProcessShape); ok {
										inst.Process_Shapes = append(inst.Process_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsProcesssNodeExpanded":
						inst.IsProcesssNodeExpanded = GongExtractBool(rhs)
					case "ProcesssWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Process); ok {
										inst.ProcesssWhoseNodeIsExpanded = append(inst.ProcesssWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Participant_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ParticipantShape); ok {
										inst.Participant_Shapes = append(inst.Participant_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsParticipantsNodeExpanded":
						inst.IsParticipantsNodeExpanded = GongExtractBool(rhs)
					case "ParticipantWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ParticipantWhoseNodeIsExpanded = append(inst.ParticipantWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ExternalParticipant_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ExternalParticipantShape); ok {
										inst.ExternalParticipant_Shapes = append(inst.ExternalParticipant_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsExternalParticipantsNodeExpanded":
						inst.IsExternalParticipantsNodeExpanded = GongExtractBool(rhs)
					case "ExternalParticipantWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ExternalParticipantWhoseNodeIsExpanded = append(inst.ExternalParticipantWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(inst.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(inst.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "TasksWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TasksWhoseNodeIsExpanded = append(inst.TasksWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "Task_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskShape); ok {
										inst.Task_Shapes = append(inst.Task_Shapes, typedTarget)
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
					case "AllocatedProcessesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Process); ok {
										inst.AllocatedProcessesWhoseNodeIsExpanded = append(inst.AllocatedProcessesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "AllocatedProcessShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*AllocatedProcessShape); ok {
										inst.AllocatedProcessShapes = append(inst.AllocatedProcessShapes, typedTarget)
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
					}
				case *ExternalParticipantShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Participant); ok {
									inst.Participant = typedTarget
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
					case "RootProcesses":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Process); ok {
										inst.RootProcesses = append(inst.RootProcesses, typedTarget)
									}
								}
							}
						}
					case "IsProcessesNodeExpanded":
						inst.IsProcessesNodeExpanded = GongExtractBool(rhs)
					case "ProcesssWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Process); ok {
										inst.ProcesssWhoseNodeIsExpanded = append(inst.ProcesssWhoseNodeIsExpanded, typedTarget)
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
					case "ParticipantsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ParticipantsWhoseNodeIsExpanded = append(inst.ParticipantsWhoseNodeIsExpanded, typedTarget)
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
					case "IsTasksNodeExpanded":
						inst.IsTasksNodeExpanded = GongExtractBool(rhs)
					case "Tasks":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.Tasks = append(inst.Tasks, typedTarget)
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
								if typedTarget, ok := target.(*Task); ok {
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
					}
				case *Participant:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsProcessResource":
						inst.IsProcessResource = GongExtractBool(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Resources":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Resource); ok {
										inst.Resources = append(inst.Resources, typedTarget)
									}
								}
							}
						}
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "Processes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Process); ok {
										inst.Processes = append(inst.Processes, typedTarget)
									}
								}
							}
						}
					case "IsProcessesNodeExpanded":
						inst.IsProcessesNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsTasksNodeExpanded":
						inst.IsTasksNodeExpanded = GongExtractBool(rhs)
					case "Tasks":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.Tasks = append(inst.Tasks, typedTarget)
									}
								}
							}
						}
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
					case "TaskWhoseOutControlFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TaskWhoseOutControlFlowsNodeIsExpanded = append(inst.TaskWhoseOutControlFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "TaskWhoseInControlFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TaskWhoseInControlFlowsNodeIsExpanded = append(inst.TaskWhoseInControlFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsDataFlowsNodeExpanded":
						inst.IsDataFlowsNodeExpanded = GongExtractBool(rhs)
					case "TaskWhoseOutDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TaskWhoseOutDataFlowsNodeIsExpanded = append(inst.TaskWhoseOutDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "TaskWhoseInDataFlowsNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TaskWhoseInDataFlowsNodeIsExpanded = append(inst.TaskWhoseInDataFlowsNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					}
				case *ParticipantShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Participant":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Participant); ok {
									inst.Participant = typedTarget
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
					case "WidthWeight":
						inst.WidthWeight = GongExtractFloat(rhs)
					}
				case *Process:
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
					case "DiagramProcesss":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DiagramProcess); ok {
										inst.DiagramProcesss = append(inst.DiagramProcesss, typedTarget)
									}
								}
							}
						}
					case "DiagramProcessWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DiagramProcess); ok {
										inst.DiagramProcessWhoseNodeIsExpanded = append(inst.DiagramProcessWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsSubProcessNodeExpanded":
						inst.IsSubProcessNodeExpanded = GongExtractBool(rhs)
					case "SubProcesses":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Process); ok {
										inst.SubProcesses = append(inst.SubProcesses, typedTarget)
									}
								}
							}
						}
					case "Participants":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.Participants = append(inst.Participants, typedTarget)
									}
								}
							}
						}
					case "ParticipantWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ParticipantWhoseNodeIsExpanded = append(inst.ParticipantWhoseNodeIsExpanded, typedTarget)
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
					case "ExternalParticipants":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ExternalParticipants = append(inst.ExternalParticipants, typedTarget)
									}
								}
							}
						}
					case "ExternalParticipantWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Participant); ok {
										inst.ExternalParticipantWhoseNodeIsExpanded = append(inst.ExternalParticipantWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					}
				case *ProcessShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Process":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Process); ok {
									inst.Process = typedTarget
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
				case *Task:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsStartTask":
						inst.IsStartTask = GongExtractBool(rhs)
					case "IsEndTask":
						inst.IsEndTask = GongExtractBool(rhs)
					case "Type":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Process); ok {
									inst.Type = typedTarget
								}
							}
						}
					case "IsTaskNameNotProcessName":
						inst.IsTaskNameNotProcessName = GongExtractBool(rhs)
					}
				case *TaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Task); ok {
									inst.Task = typedTarget
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

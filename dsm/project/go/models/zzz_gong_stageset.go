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
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramIdent, diagram.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramIdent, diagram.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.DateFormat = %s", diagramIdent, __gong__toRawStringLiteral(diagram.DateFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramIdent, diagram.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTimeDiagram = %t", diagramIdent, diagram.IsTimeDiagram))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ComputedStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ComputedEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", diagramIdent, int64(diagram.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseManualStartAndEndDates = %t", diagramIdent, diagram.UseManualStartAndEndDates))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ManualStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ManualEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeStep = %d", diagramIdent, diagram.TimeStep))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeStepScale = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.TimeStepScale))))
			values.WriteString(fmt.Sprintf("\n\t%s.LaneHeight = %f", diagramIdent, diagram.LaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RatioBarToLaneHeight = %f", diagramIdent, diagram.RatioBarToLaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.YTopMargin = %f", diagramIdent, diagram.YTopMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftText = %f", diagramIdent, diagram.XLeftText))
			values.WriteString(fmt.Sprintf("\n\t%s.TextHeight = %f", diagramIdent, diagram.TextHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftLanes = %f", diagramIdent, diagram.XLeftLanes))
			values.WriteString(fmt.Sprintf("\n\t%s.XRightMargin = %f", diagramIdent, diagram.XRightMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowLengthToTheRightOfStartBar = %f", diagramIdent, diagram.ArrowLengthToTheRightOfStartBar))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowTipLenght = %f", diagramIdent, diagram.ArrowTipLenght))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Color = %s", diagramIdent, __gong__toRawStringLiteral(diagram.TimeLine_Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_FillOpacity = %f", diagramIdent, diagram.TimeLine_FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Stroke = %s", diagramIdent, __gong__toRawStringLiteral(diagram.TimeLine_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_StrokeWidth = %f", diagramIdent, diagram.TimeLine_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DrawVerticalTimeLines = %t", diagramIdent, diagram.DrawVerticalTimeLines))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_Stroke = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Group_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeWidth = %f", diagramIdent, diagram.Group_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeDashArray = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Group_StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.DateYOffset = %f", diagramIdent, diagram.DateYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.AlignOnStartEndOnYearStart = %t", diagramIdent, diagram.AlignOnStartEndOnYearStart))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramIdent, diagram.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramIdent, diagram.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInAutoLayoutMode = %t", diagramIdent, diagram.IsInAutoLayoutMode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPBSNodeExpanded = %t", diagramIdent, diagram.IsPBSNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWBSNodeExpanded = %t", diagramIdent, diagram.IsWBSNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTaskGroupsNodeExpanded = %t", diagramIdent, diagram.IsTaskGroupsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramIdent, diagram.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResourcesNodeExpanded = %t", diagramIdent, diagram.IsResourcesNodeExpanded))
			for _, elem := range diagram.Product_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product_Shapes = append(%s.Product_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ProductsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProductsWhoseNodeIsExpanded = append(%s.ProductsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ProductComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProductComposition_Shapes = append(%s.ProductComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ProductReference_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProductReference_Shapes = append(%s.ProductReference_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Task_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task_Shapes = append(%s.Task_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseNodeIsExpanded = append(%s.TasksWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhoseInputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseInputNodeIsExpanded = append(%s.TasksWhoseInputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhoseOutputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseOutputNodeIsExpanded = append(%s.TasksWhoseOutputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhosePredecessorNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhosePredecessorNodeIsExpanded = append(%s.TasksWhosePredecessorNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskGroupShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroupShapes = append(%s.TaskGroupShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskGroupsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroupsWhoseNodeIsExpanded = append(%s.TaskGroupsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskComposition_Shapes = append(%s.TaskComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskInputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskInputShapes = append(%s.TaskInputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskOutputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskOutputShapes = append(%s.TaskOutputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskPredecessorShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskPredecessorShapes = append(%s.TaskPredecessorShapes, %s)", diagramIdent, diagramIdent, targetIdent))
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
			for _, elem := range diagram.NoteProductShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteProductShapes = append(%s.NoteProductShapes, %s)", diagramIdent, diagramIdent, targetIdent))
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
			for _, elem := range diagram.Resource_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource_Shapes = append(%s.Resource_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
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
			for _, elem := range diagram.ResourceTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourceTaskShapes = append(%s.ResourceTaskShapes, %s)", diagramIdent, diagramIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
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
			for _, elem := range library.RootProducts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootProducts = append(%s.RootProducts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootTasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootTasks = append(%s.RootTasks, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootTaskGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootTaskGroups = append(%s.RootTaskGroups, %s)", libraryIdent, libraryIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", noteIdent, int(note.LayoutDirection)))
			for _, elem := range note.Products {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Products = append(%s.Products, %s)", noteIdent, noteIdent, targetIdent))
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
		noteproductshapeOrdered := []*NoteProductShape{}
		for noteproductshape := range stageSet.Stage.NoteProductShapes {
			noteproductshapeOrdered = append(noteproductshapeOrdered, noteproductshape)
		}
		sort.Slice(noteproductshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteProductShape_stagedOrder[noteproductshapeOrdered[i]] < stageSet.Stage.NoteProductShape_stagedOrder[noteproductshapeOrdered[j]]
		})
		for _, noteproductshape := range noteproductshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteproductshapeIdent := "__models" + noteproductshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteProductShape{Name: %s}).Stage(stageSet.Stage)", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteproductshapeIdent, noteproductshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteproductshapeIdent, noteproductshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteproductshapeIdent, __gong__toRawStringLiteral(string(noteproductshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteproductshapeIdent, __gong__toRawStringLiteral(string(noteproductshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteproductshapeIdent, noteproductshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteproductshapeIdent, noteproductshape.IsHidden))
			if noteproductshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteproductshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteproductshapeIdent, targetIdent))
			}
			if noteproductshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteproductshape.Product.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", noteproductshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteresourceshapeOrdered := []*NoteResourceShape{}
		for noteresourceshape := range stageSet.Stage.NoteResourceShapes {
			noteresourceshapeOrdered = append(noteresourceshapeOrdered, noteresourceshape)
		}
		sort.Slice(noteresourceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteResourceShape_stagedOrder[noteresourceshapeOrdered[i]] < stageSet.Stage.NoteResourceShape_stagedOrder[noteresourceshapeOrdered[j]]
		})
		for _, noteresourceshape := range noteresourceshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteresourceshapeIdent := "__models" + noteresourceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteResourceShape{Name: %s}).Stage(stageSet.Stage)", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteresourceshapeIdent, noteresourceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteresourceshapeIdent, noteresourceshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(string(noteresourceshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(string(noteresourceshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteresourceshapeIdent, noteresourceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteresourceshapeIdent, noteresourceshape.IsHidden))
			if noteresourceshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteresourceshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteresourceshapeIdent, targetIdent))
			}
			if noteresourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteresourceshape.Resource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", noteresourceshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.OverideLayoutDirection = %t", noteshapeIdent, noteshape.OverideLayoutDirection))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", noteshapeIdent, int(noteshape.LayoutDirection)))
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
		productOrdered := []*Product{}
		for product := range stageSet.Stage.Products {
			productOrdered = append(productOrdered, product)
		}
		sort.Slice(productOrdered, func(i, j int) bool {
			return stageSet.Stage.Product_stagedOrder[productOrdered[i]] < stageSet.Stage.Product_stagedOrder[productOrdered[j]]
		})
		for _, product := range productOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productIdent := "__models" + product.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Product{Name: %s}).Stage(stageSet.Stage)", productIdent, __gong__toRawStringLiteral(product.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productIdent, __gong__toRawStringLiteral(product.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", productIdent, __gong__toRawStringLiteral(product.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProducersNodeExpanded = %t", productIdent, product.IsProducersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConsumersNodeExpanded = %t", productIdent, product.IsConsumersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", productIdent, product.IsImport))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", productIdent, __gong__toRawStringLiteral(product.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", productIdent, product.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", productIdent, int(product.LayoutDirection)))
			for _, elem := range product.SubProducts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubProducts = append(%s.SubProducts, %s)", productIdent, productIdent, targetIdent))
			}
			if product.ReferencedProduct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + product.ReferencedProduct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedProduct = %s", productIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		productcompositionshapeOrdered := []*ProductCompositionShape{}
		for productcompositionshape := range stageSet.Stage.ProductCompositionShapes {
			productcompositionshapeOrdered = append(productcompositionshapeOrdered, productcompositionshape)
		}
		sort.Slice(productcompositionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ProductCompositionShape_stagedOrder[productcompositionshapeOrdered[i]] < stageSet.Stage.ProductCompositionShape_stagedOrder[productcompositionshapeOrdered[j]]
		})
		for _, productcompositionshape := range productcompositionshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productcompositionshapeIdent := "__models" + productcompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProductCompositionShape{Name: %s}).Stage(stageSet.Stage)", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", productcompositionshapeIdent, productcompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", productcompositionshapeIdent, productcompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(string(productcompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(string(productcompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", productcompositionshapeIdent, productcompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productcompositionshapeIdent, productcompositionshape.IsHidden))
			if productcompositionshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + productcompositionshape.Product.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", productcompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		productreferenceshapeOrdered := []*ProductReferenceShape{}
		for productreferenceshape := range stageSet.Stage.ProductReferenceShapes {
			productreferenceshapeOrdered = append(productreferenceshapeOrdered, productreferenceshape)
		}
		sort.Slice(productreferenceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ProductReferenceShape_stagedOrder[productreferenceshapeOrdered[i]] < stageSet.Stage.ProductReferenceShape_stagedOrder[productreferenceshapeOrdered[j]]
		})
		for _, productreferenceshape := range productreferenceshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productreferenceshapeIdent := "__models" + productreferenceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProductReferenceShape{Name: %s}).Stage(stageSet.Stage)", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", productreferenceshapeIdent, productreferenceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", productreferenceshapeIdent, productreferenceshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(string(productreferenceshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(string(productreferenceshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", productreferenceshapeIdent, productreferenceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productreferenceshapeIdent, productreferenceshape.IsHidden))
			if productreferenceshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + productreferenceshape.Product.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", productreferenceshapeIdent, targetIdent))
			}
			if productreferenceshape.ReferencedProduct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + productreferenceshape.ReferencedProduct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedProduct = %s", productreferenceshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		productshapeOrdered := []*ProductShape{}
		for productshape := range stageSet.Stage.ProductShapes {
			productshapeOrdered = append(productshapeOrdered, productshape)
		}
		sort.Slice(productshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ProductShape_stagedOrder[productshapeOrdered[i]] < stageSet.Stage.ProductShape_stagedOrder[productshapeOrdered[j]]
		})
		for _, productshape := range productshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productshapeIdent := "__models" + productshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProductShape{Name: %s}).Stage(stageSet.Stage)", productshapeIdent, __gong__toRawStringLiteral(productshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowType = %t", productshapeIdent, productshape.IsShowType))
			values.WriteString(fmt.Sprintf("\n\t%s.OverideLayoutDirection = %t", productshapeIdent, productshape.OverideLayoutDirection))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", productshapeIdent, int(productshape.LayoutDirection)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", productshapeIdent, productshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", productshapeIdent, productshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", productshapeIdent, productshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", productshapeIdent, productshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productshapeIdent, productshape.IsHidden))
			if productshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + productshape.Product.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", productshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", resourceIdent, __gong__toRawStringLiteral(resource.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", resourceIdent, __gong__toRawStringLiteral(resource.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", resourceIdent, resource.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", resourceIdent, int(resource.LayoutDirection)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", resourceIdent, resource.IsImport))
			for _, elem := range resource.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", resourceIdent, resourceIdent, targetIdent))
			}
			for _, elem := range resource.SubResources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubResources = append(%s.SubResources, %s)", resourceIdent, resourceIdent, targetIdent))
			}
			if resource.ReferencedResource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + resource.ReferencedResource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedResource = %s", resourceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		resourcecompositionshapeOrdered := []*ResourceCompositionShape{}
		for resourcecompositionshape := range stageSet.Stage.ResourceCompositionShapes {
			resourcecompositionshapeOrdered = append(resourcecompositionshapeOrdered, resourcecompositionshape)
		}
		sort.Slice(resourcecompositionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ResourceCompositionShape_stagedOrder[resourcecompositionshapeOrdered[i]] < stageSet.Stage.ResourceCompositionShape_stagedOrder[resourcecompositionshapeOrdered[j]]
		})
		for _, resourcecompositionshape := range resourcecompositionshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourcecompositionshapeIdent := "__models" + resourcecompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ResourceCompositionShape{Name: %s}).Stage(stageSet.Stage)", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(string(resourcecompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(string(resourcecompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourcecompositionshapeIdent, resourcecompositionshape.IsHidden))
			if resourcecompositionshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + resourcecompositionshape.Resource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", resourcecompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		resourceshapeOrdered := []*ResourceShape{}
		for resourceshape := range stageSet.Stage.ResourceShapes {
			resourceshapeOrdered = append(resourceshapeOrdered, resourceshape)
		}
		sort.Slice(resourceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ResourceShape_stagedOrder[resourceshapeOrdered[i]] < stageSet.Stage.ResourceShape_stagedOrder[resourceshapeOrdered[j]]
		})
		for _, resourceshape := range resourceshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourceshapeIdent := "__models" + resourceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ResourceShape{Name: %s}).Stage(stageSet.Stage)", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OverideLayoutDirection = %t", resourceshapeIdent, resourceshape.OverideLayoutDirection))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", resourceshapeIdent, int(resourceshape.LayoutDirection)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", resourceshapeIdent, resourceshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", resourceshapeIdent, resourceshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", resourceshapeIdent, resourceshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", resourceshapeIdent, resourceshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourceshapeIdent, resourceshape.IsHidden))
			if resourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + resourceshape.Resource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", resourceshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		resourcetaskshapeOrdered := []*ResourceTaskShape{}
		for resourcetaskshape := range stageSet.Stage.ResourceTaskShapes {
			resourcetaskshapeOrdered = append(resourcetaskshapeOrdered, resourcetaskshape)
		}
		sort.Slice(resourcetaskshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ResourceTaskShape_stagedOrder[resourcetaskshapeOrdered[i]] < stageSet.Stage.ResourceTaskShape_stagedOrder[resourcetaskshapeOrdered[j]]
		})
		for _, resourcetaskshape := range resourcetaskshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourcetaskshapeIdent := "__models" + resourcetaskshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ResourceTaskShape{Name: %s}).Stage(stageSet.Stage)", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", resourcetaskshapeIdent, resourcetaskshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", resourcetaskshapeIdent, resourcetaskshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(string(resourcetaskshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(string(resourcetaskshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", resourcetaskshapeIdent, resourcetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourcetaskshapeIdent, resourcetaskshape.IsHidden))
			if resourcetaskshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + resourcetaskshape.Resource.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", resourcetaskshapeIdent, targetIdent))
			}
			if resourcetaskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + resourcetaskshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", resourcetaskshapeIdent, targetIdent))
			}
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
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", taskIdent, task.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", taskIdent, task.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartDateComputedFromPredecessors = %t", taskIdent, task.IsStartDateComputedFromPredecessors))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationYears = %f", taskIdent, task.DurationYears))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationMonths = %f", taskIdent, task.DurationMonths))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationWeeks = %f", taskIdent, task.DurationWeeks))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationDays = %f", taskIdent, task.DurationDays))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationHours = %f", taskIdent, task.DurationHours))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndDateComputedFromDuration = %t", taskIdent, task.IsEndDateComputedFromDuration))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMilestone = %t", taskIdent, task.IsMilestone))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCompletion = %t", taskIdent, task.IsWithCompletion))
			values.WriteString(fmt.Sprintf("\n\t%s.Completion = %s", taskIdent, __gong__toRawStringLiteral(string(task.Completion))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayVerticalBar = %t", taskIdent, task.DisplayVerticalBar))
			values.WriteString(fmt.Sprintf("\n\t%s.TextPosition = %s", taskIdent, __gong__toRawStringLiteral(string(task.TextPosition))))
			values.WriteString(fmt.Sprintf("\n\t%s.XOffset = %f", taskIdent, task.XOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.YOffset = %f", taskIdent, task.YOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", taskIdent, task.IsImport))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInputsNodeExpanded = %t", taskIdent, task.IsInputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOutputsNodeExpanded = %t", taskIdent, task.IsOutputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskIdent, __gong__toRawStringLiteral(task.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskIdent, task.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", taskIdent, int(task.LayoutDirection)))
			for _, elem := range task.Predecessors {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Predecessors = append(%s.Predecessors, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.Inputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inputs = append(%s.Inputs, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.Outputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Outputs = append(%s.Outputs, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.SubTasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubTasks = append(%s.SubTasks, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.TaskGroupsToDisplay {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroupsToDisplay = append(%s.TaskGroupsToDisplay, %s)", taskIdent, taskIdent, targetIdent))
			}
			if task.ReferencedTask != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + task.ReferencedTask.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedTask = %s", taskIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskcompositionshapeOrdered := []*TaskCompositionShape{}
		for taskcompositionshape := range stageSet.Stage.TaskCompositionShapes {
			taskcompositionshapeOrdered = append(taskcompositionshapeOrdered, taskcompositionshape)
		}
		sort.Slice(taskcompositionshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskCompositionShape_stagedOrder[taskcompositionshapeOrdered[i]] < stageSet.Stage.TaskCompositionShape_stagedOrder[taskcompositionshapeOrdered[j]]
		})
		for _, taskcompositionshape := range taskcompositionshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskcompositionshapeIdent := "__models" + taskcompositionshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskCompositionShape{Name: %s}).Stage(stageSet.Stage)", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskcompositionshapeIdent, taskcompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskcompositionshapeIdent, taskcompositionshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(string(taskcompositionshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(string(taskcompositionshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskcompositionshapeIdent, taskcompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskcompositionshapeIdent, taskcompositionshape.IsHidden))
			if taskcompositionshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskcompositionshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskcompositionshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskgroupOrdered := []*TaskGroup{}
		for taskgroup := range stageSet.Stage.TaskGroups {
			taskgroupOrdered = append(taskgroupOrdered, taskgroup)
		}
		sort.Slice(taskgroupOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskGroup_stagedOrder[taskgroupOrdered[i]] < stageSet.Stage.TaskGroup_stagedOrder[taskgroupOrdered[j]]
		})
		for _, taskgroup := range taskgroupOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskgroupIdent := "__models" + taskgroup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskGroup{Name: %s}).Stage(stageSet.Stage)", taskgroupIdent, __gong__toRawStringLiteral(taskgroup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskgroupIdent, __gong__toRawStringLiteral(taskgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskgroupIdent, __gong__toRawStringLiteral(taskgroup.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskgroupIdent, taskgroup.IsExpanded))
			for _, elem := range taskgroup.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", taskgroupIdent, taskgroupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskgroupshapeOrdered := []*TaskGroupShape{}
		for taskgroupshape := range stageSet.Stage.TaskGroupShapes {
			taskgroupshapeOrdered = append(taskgroupshapeOrdered, taskgroupshape)
		}
		sort.Slice(taskgroupshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskGroupShape_stagedOrder[taskgroupshapeOrdered[i]] < stageSet.Stage.TaskGroupShape_stagedOrder[taskgroupshapeOrdered[j]]
		})
		for _, taskgroupshape := range taskgroupshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskgroupshapeIdent := "__models" + taskgroupshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskGroupShape{Name: %s}).Stage(stageSet.Stage)", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", taskgroupshapeIdent, taskgroupshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", taskgroupshapeIdent, taskgroupshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", taskgroupshapeIdent, taskgroupshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", taskgroupshapeIdent, taskgroupshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskgroupshapeIdent, taskgroupshape.IsHidden))
			if taskgroupshape.TaskGroup != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskgroupshape.TaskGroup.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroup = %s", taskgroupshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskinputshapeOrdered := []*TaskInputShape{}
		for taskinputshape := range stageSet.Stage.TaskInputShapes {
			taskinputshapeOrdered = append(taskinputshapeOrdered, taskinputshape)
		}
		sort.Slice(taskinputshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskInputShape_stagedOrder[taskinputshapeOrdered[i]] < stageSet.Stage.TaskInputShape_stagedOrder[taskinputshapeOrdered[j]]
		})
		for _, taskinputshape := range taskinputshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskinputshapeIdent := "__models" + taskinputshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskInputShape{Name: %s}).Stage(stageSet.Stage)", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskinputshapeIdent, taskinputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskinputshapeIdent, taskinputshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", taskinputshapeIdent, __gong__toRawStringLiteral(string(taskinputshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", taskinputshapeIdent, __gong__toRawStringLiteral(string(taskinputshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskinputshapeIdent, taskinputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskinputshapeIdent, taskinputshape.IsHidden))
			if taskinputshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskinputshape.Product.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", taskinputshapeIdent, targetIdent))
			}
			if taskinputshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskinputshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskinputshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskoutputshapeOrdered := []*TaskOutputShape{}
		for taskoutputshape := range stageSet.Stage.TaskOutputShapes {
			taskoutputshapeOrdered = append(taskoutputshapeOrdered, taskoutputshape)
		}
		sort.Slice(taskoutputshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskOutputShape_stagedOrder[taskoutputshapeOrdered[i]] < stageSet.Stage.TaskOutputShape_stagedOrder[taskoutputshapeOrdered[j]]
		})
		for _, taskoutputshape := range taskoutputshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskoutputshapeIdent := "__models" + taskoutputshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskOutputShape{Name: %s}).Stage(stageSet.Stage)", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskoutputshapeIdent, taskoutputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskoutputshapeIdent, taskoutputshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(string(taskoutputshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(string(taskoutputshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskoutputshapeIdent, taskoutputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskoutputshapeIdent, taskoutputshape.IsHidden))
			if taskoutputshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskoutputshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskoutputshapeIdent, targetIdent))
			}
			if taskoutputshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskoutputshape.Product.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", taskoutputshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		taskpredecessorshapeOrdered := []*TaskPredecessorShape{}
		for taskpredecessorshape := range stageSet.Stage.TaskPredecessorShapes {
			taskpredecessorshapeOrdered = append(taskpredecessorshapeOrdered, taskpredecessorshape)
		}
		sort.Slice(taskpredecessorshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TaskPredecessorShape_stagedOrder[taskpredecessorshapeOrdered[i]] < stageSet.Stage.TaskPredecessorShape_stagedOrder[taskpredecessorshapeOrdered[j]]
		})
		for _, taskpredecessorshape := range taskpredecessorshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskpredecessorshapeIdent := "__models" + taskpredecessorshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskPredecessorShape{Name: %s}).Stage(stageSet.Stage)", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(string(taskpredecessorshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(string(taskpredecessorshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskpredecessorshapeIdent, taskpredecessorshape.IsHidden))
			if taskpredecessorshape.Predecessor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskpredecessorshape.Predecessor.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Predecessor = %s", taskpredecessorshapeIdent, targetIdent))
			}
			if taskpredecessorshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + taskpredecessorshape.Task.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskpredecessorshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowDate = %t", taskshapeIdent, taskshape.IsShowDate))
			values.WriteString(fmt.Sprintf("\n\t%s.OverideLayoutDirection = %t", taskshapeIdent, taskshape.OverideLayoutDirection))
			values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = %d", taskshapeIdent, int(taskshape.LayoutDirection)))
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

	"github.com/fullstack-lang/gong/dsm/project/go/models"
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
		case "github.com/fullstack-lang/gong/dsm/project/go/models":
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
				case "NoteProductShape":
					if !preserveOrder {
						inst := (&NoteProductShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteProductShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteResourceShape":
					if !preserveOrder {
						inst := (&NoteResourceShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteResourceShape)
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
				case "Product":
					if !preserveOrder {
						inst := (&Product{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Product)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ProductCompositionShape":
					if !preserveOrder {
						inst := (&ProductCompositionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ProductCompositionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ProductReferenceShape":
					if !preserveOrder {
						inst := (&ProductReferenceShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ProductReferenceShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ProductShape":
					if !preserveOrder {
						inst := (&ProductShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ProductShape)
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
				case "ResourceCompositionShape":
					if !preserveOrder {
						inst := (&ResourceCompositionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ResourceCompositionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ResourceShape":
					if !preserveOrder {
						inst := (&ResourceShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ResourceShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ResourceTaskShape":
					if !preserveOrder {
						inst := (&ResourceTaskShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ResourceTaskShape)
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
				case "TaskCompositionShape":
					if !preserveOrder {
						inst := (&TaskCompositionShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskCompositionShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TaskGroup":
					if !preserveOrder {
						inst := (&TaskGroup{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskGroup)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TaskGroupShape":
					if !preserveOrder {
						inst := (&TaskGroupShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskGroupShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TaskInputShape":
					if !preserveOrder {
						inst := (&TaskInputShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskInputShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TaskOutputShape":
					if !preserveOrder {
						inst := (&TaskOutputShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskOutputShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TaskPredecessorShape":
					if !preserveOrder {
						inst := (&TaskPredecessorShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TaskPredecessorShape)
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
				case *Diagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultBoxWidth":
						inst.DefaultBoxWidth = GongExtractFloat(rhs)
					case "DefaultBoxHeigth":
						inst.DefaultBoxHeigth = GongExtractFloat(rhs)
					case "DateFormat":
						inst.DateFormat = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsTimeDiagram":
						inst.IsTimeDiagram = GongExtractBool(rhs)
					case "ComputedStart":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ComputedEnd":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "UseManualStartAndEndDates":
						inst.UseManualStartAndEndDates = GongExtractBool(rhs)
					case "ManualStart":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ManualEnd":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "TimeStep":
						inst.TimeStep = GongExtractInt(rhs)
					case "TimeStepScale":
						inst.TimeStepScale = TimeStepScaleEnum(GongExtractString(rhs))
					case "LaneHeight":
						inst.LaneHeight = GongExtractFloat(rhs)
					case "RatioBarToLaneHeight":
						inst.RatioBarToLaneHeight = GongExtractFloat(rhs)
					case "YTopMargin":
						inst.YTopMargin = GongExtractFloat(rhs)
					case "XLeftText":
						inst.XLeftText = GongExtractFloat(rhs)
					case "TextHeight":
						inst.TextHeight = GongExtractFloat(rhs)
					case "XLeftLanes":
						inst.XLeftLanes = GongExtractFloat(rhs)
					case "XRightMargin":
						inst.XRightMargin = GongExtractFloat(rhs)
					case "ArrowLengthToTheRightOfStartBar":
						inst.ArrowLengthToTheRightOfStartBar = GongExtractFloat(rhs)
					case "ArrowTipLenght":
						inst.ArrowTipLenght = GongExtractFloat(rhs)
					case "TimeLine_Color":
						inst.TimeLine_Color = GongExtractString(rhs)
					case "TimeLine_FillOpacity":
						inst.TimeLine_FillOpacity = GongExtractFloat(rhs)
					case "TimeLine_Stroke":
						inst.TimeLine_Stroke = GongExtractString(rhs)
					case "TimeLine_StrokeWidth":
						inst.TimeLine_StrokeWidth = GongExtractFloat(rhs)
					case "DrawVerticalTimeLines":
						inst.DrawVerticalTimeLines = GongExtractBool(rhs)
					case "Group_Stroke":
						inst.Group_Stroke = GongExtractString(rhs)
					case "Group_StrokeWidth":
						inst.Group_StrokeWidth = GongExtractFloat(rhs)
					case "Group_StrokeDashArray":
						inst.Group_StrokeDashArray = GongExtractString(rhs)
					case "DateYOffset":
						inst.DateYOffset = GongExtractFloat(rhs)
					case "AlignOnStartEndOnYearStart":
						inst.AlignOnStartEndOnYearStart = GongExtractBool(rhs)
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
					case "IsInAutoLayoutMode":
						inst.IsInAutoLayoutMode = GongExtractBool(rhs)
					case "Product_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ProductShape); ok {
										inst.Product_Shapes = append(inst.Product_Shapes, typedTarget)
									}
								}
							}
						}
					case "ProductsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Product); ok {
										inst.ProductsWhoseNodeIsExpanded = append(inst.ProductsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsPBSNodeExpanded":
						inst.IsPBSNodeExpanded = GongExtractBool(rhs)
					case "ProductComposition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ProductCompositionShape); ok {
										inst.ProductComposition_Shapes = append(inst.ProductComposition_Shapes, typedTarget)
									}
								}
							}
						}
					case "ProductReference_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ProductReferenceShape); ok {
										inst.ProductReference_Shapes = append(inst.ProductReference_Shapes, typedTarget)
									}
								}
							}
						}
					case "IsWBSNodeExpanded":
						inst.IsWBSNodeExpanded = GongExtractBool(rhs)
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
					case "TasksWhoseInputNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TasksWhoseInputNodeIsExpanded = append(inst.TasksWhoseInputNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "TasksWhoseOutputNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TasksWhoseOutputNodeIsExpanded = append(inst.TasksWhoseOutputNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "TasksWhosePredecessorNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.TasksWhosePredecessorNodeIsExpanded = append(inst.TasksWhosePredecessorNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsTaskGroupsNodeExpanded":
						inst.IsTaskGroupsNodeExpanded = GongExtractBool(rhs)
					case "TaskGroupShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskGroupShape); ok {
										inst.TaskGroupShapes = append(inst.TaskGroupShapes, typedTarget)
									}
								}
							}
						}
					case "TaskGroupsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskGroup); ok {
										inst.TaskGroupsWhoseNodeIsExpanded = append(inst.TaskGroupsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "TaskComposition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskCompositionShape); ok {
										inst.TaskComposition_Shapes = append(inst.TaskComposition_Shapes, typedTarget)
									}
								}
							}
						}
					case "TaskInputShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskInputShape); ok {
										inst.TaskInputShapes = append(inst.TaskInputShapes, typedTarget)
									}
								}
							}
						}
					case "TaskOutputShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskOutputShape); ok {
										inst.TaskOutputShapes = append(inst.TaskOutputShapes, typedTarget)
									}
								}
							}
						}
					case "TaskPredecessorShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskPredecessorShape); ok {
										inst.TaskPredecessorShapes = append(inst.TaskPredecessorShapes, typedTarget)
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
					case "NoteProductShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*NoteProductShape); ok {
										inst.NoteProductShapes = append(inst.NoteProductShapes, typedTarget)
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
									if typedTarget, ok := target.(*NoteResourceShape); ok {
										inst.NoteResourceShapes = append(inst.NoteResourceShapes, typedTarget)
									}
								}
							}
						}
					case "Resource_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ResourceShape); ok {
										inst.Resource_Shapes = append(inst.Resource_Shapes, typedTarget)
									}
								}
							}
						}
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
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "ResourceComposition_Shapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ResourceCompositionShape); ok {
										inst.ResourceComposition_Shapes = append(inst.ResourceComposition_Shapes, typedTarget)
									}
								}
							}
						}
					case "ResourceTaskShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ResourceTaskShape); ok {
										inst.ResourceTaskShapes = append(inst.ResourceTaskShapes, typedTarget)
									}
								}
							}
						}
					}
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "RootProducts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Product); ok {
										inst.RootProducts = append(inst.RootProducts, typedTarget)
									}
								}
							}
						}
					case "RootTasks":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.RootTasks = append(inst.RootTasks, typedTarget)
									}
								}
							}
						}
					case "RootTaskGroups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskGroup); ok {
										inst.RootTaskGroups = append(inst.RootTaskGroups, typedTarget)
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
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
					case "Products":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Product); ok {
										inst.Products = append(inst.Products, typedTarget)
									}
								}
							}
						}
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
					}
				case *NoteProductShape:
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
					case "Product":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.Product = typedTarget
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
				case *NoteResourceShape:
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
					case "Resource":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
									inst.Resource = typedTarget
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
					case "OverideLayoutDirection":
						inst.OverideLayoutDirection = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
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
				case *Product:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SubProducts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Product); ok {
										inst.SubProducts = append(inst.SubProducts, typedTarget)
									}
								}
							}
						}
					case "IsProducersNodeExpanded":
						inst.IsProducersNodeExpanded = GongExtractBool(rhs)
					case "IsConsumersNodeExpanded":
						inst.IsConsumersNodeExpanded = GongExtractBool(rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedProduct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.ReferencedProduct = typedTarget
								}
							}
						}
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
					}
				case *ProductCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.Product = typedTarget
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
				case *ProductReferenceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.Product = typedTarget
								}
							}
						}
					case "ReferencedProduct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.ReferencedProduct = typedTarget
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
				case *ProductShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.Product = typedTarget
								}
							}
						}
					case "IsShowType":
						inst.IsShowType = GongExtractBool(rhs)
					case "OverideLayoutDirection":
						inst.OverideLayoutDirection = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
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
					case "Description":
						inst.Description = GongExtractString(rhs)
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
					case "SubResources":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Resource); ok {
										inst.SubResources = append(inst.SubResources, typedTarget)
									}
								}
							}
						}
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedResource":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
									inst.ReferencedResource = typedTarget
								}
							}
						}
					}
				case *ResourceCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
									inst.Resource = typedTarget
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
				case *ResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
									inst.Resource = typedTarget
								}
							}
						}
					case "OverideLayoutDirection":
						inst.OverideLayoutDirection = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
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
				case *ResourceTaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Resource); ok {
									inst.Resource = typedTarget
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
				case *Task:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
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
					case "Predecessors":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.Predecessors = append(inst.Predecessors, typedTarget)
									}
								}
							}
						}
					case "IsStartDateComputedFromPredecessors":
						inst.IsStartDateComputedFromPredecessors = GongExtractBool(rhs)
					case "DurationYears":
						inst.DurationYears = GongExtractFloat(rhs)
					case "DurationMonths":
						inst.DurationMonths = GongExtractFloat(rhs)
					case "DurationWeeks":
						inst.DurationWeeks = GongExtractFloat(rhs)
					case "DurationDays":
						inst.DurationDays = GongExtractFloat(rhs)
					case "DurationHours":
						inst.DurationHours = GongExtractFloat(rhs)
					case "IsEndDateComputedFromDuration":
						inst.IsEndDateComputedFromDuration = GongExtractBool(rhs)
					case "IsMilestone":
						inst.IsMilestone = GongExtractBool(rhs)
					case "Inputs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Product); ok {
										inst.Inputs = append(inst.Inputs, typedTarget)
									}
								}
							}
						}
					case "Outputs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Product); ok {
										inst.Outputs = append(inst.Outputs, typedTarget)
									}
								}
							}
						}
					case "SubTasks":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Task); ok {
										inst.SubTasks = append(inst.SubTasks, typedTarget)
									}
								}
							}
						}
					case "IsWithCompletion":
						inst.IsWithCompletion = GongExtractBool(rhs)
					case "Completion":
						inst.Completion = CompletionEnum(GongExtractString(rhs))
					case "DisplayVerticalBar":
						inst.DisplayVerticalBar = GongExtractBool(rhs)
					case "TaskGroupsToDisplay":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TaskGroup); ok {
										inst.TaskGroupsToDisplay = append(inst.TaskGroupsToDisplay, typedTarget)
									}
								}
							}
						}
					case "TextPosition":
						inst.TextPosition = TextPositionEnum(GongExtractString(rhs))
					case "XOffset":
						inst.XOffset = GongExtractFloat(rhs)
					case "YOffset":
						inst.YOffset = GongExtractFloat(rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedTask":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Task); ok {
									inst.ReferencedTask = typedTarget
								}
							}
						}
					case "IsInputsNodeExpanded":
						inst.IsInputsNodeExpanded = GongExtractBool(rhs)
					case "IsOutputsNodeExpanded":
						inst.IsOutputsNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
					}
				case *TaskCompositionShape:
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
				case *TaskGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
				case *TaskGroupShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "TaskGroup":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TaskGroup); ok {
									inst.TaskGroup = typedTarget
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
				case *TaskInputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.Product = typedTarget
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
				case *TaskOutputShape:
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
					case "Product":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Product); ok {
									inst.Product = typedTarget
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
				case *TaskPredecessorShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Predecessor":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Task); ok {
									inst.Predecessor = typedTarget
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
					case "IsShowDate":
						inst.IsShowDate = GongExtractBool(rhs)
					case "OverideLayoutDirection":
						inst.OverideLayoutDirection = GongExtractBool(rhs)
					case "LayoutDirection":
						inst.LayoutDirection = LayoutDirection(GongExtractInt(rhs))
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

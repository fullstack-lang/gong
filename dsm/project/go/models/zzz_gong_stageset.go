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
	"runtime/debug"
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
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramIdent, diagram.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramIdent, diagram.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.DateFormat = %s", diagramIdent, __gong__toRawStringLiteral(diagram.DateFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramIdent, diagram.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTimeDiagram = %t", diagramIdent, diagram.IsTimeDiagram))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ComputedStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ComputedEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", diagramIdent, int64(diagram.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.DrawVerticalTimeLines = %t", diagramIdent, diagram.DrawVerticalTimeLines))
			values.WriteString(fmt.Sprintf("\n\t%s.DrawSecondaryVerticalTimeLines = %t", diagramIdent, diagram.DrawSecondaryVerticalTimeLines))
			values.WriteString(fmt.Sprintf("\n\t%s.HideWeekendsPeriod = %t", diagramIdent, diagram.HideWeekendsPeriod))
			values.WriteString(fmt.Sprintf("\n\t%s.UseManualStartAndEndDates = %t", diagramIdent, diagram.UseManualStartAndEndDates))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ManualStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ManualEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeStep = %d", diagramIdent, diagram.TimeStep))
			if diagram.TimeStepScale.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.TimeStepScale = models.%s", diagramIdent, diagram.TimeStepScale.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.TimeStepScale = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.SecondaryTimeStep = %d", diagramIdent, diagram.SecondaryTimeStep))
			if diagram.SecondaryTimeStepScale.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.SecondaryTimeStepScale = models.%s", diagramIdent, diagram.SecondaryTimeStepScale.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.SecondaryTimeStepScale = \"\"", diagramIdent))
			}
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
			values.WriteString(fmt.Sprintf("\n\t%s.Group_Stroke = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Group_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeWidth = %f", diagramIdent, diagram.Group_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeDashArray = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Group_StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.DateYOffset = %f", diagramIdent, diagram.DateYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.AlignOnBeginningOfTimeScale = %t", diagramIdent, diagram.AlignOnBeginningOfTimeScale))
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
			if note.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", noteIdent, note.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", noteIdent))
			}
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
		for _, noteproductshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteProductShapes, stageSet.Stage.NoteProductShape_stagedOrder) {
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
			if noteproductshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteproductshapeIdent, noteproductshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteproductshapeIdent))
			}
			if noteproductshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteproductshapeIdent, noteproductshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteproductshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteproductshapeIdent, noteproductshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteproductshapeIdent, noteproductshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteproductshapeIdent, noteproductshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteproductshapeIdent, noteproductshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteproductshapeIdent, noteproductshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.FontColor)))
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
		for _, noteresourceshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteResourceShapes, stageSet.Stage.NoteResourceShape_stagedOrder) {
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
			if noteresourceshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteresourceshapeIdent, noteresourceshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteresourceshapeIdent))
			}
			if noteresourceshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteresourceshapeIdent, noteresourceshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteresourceshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteresourceshapeIdent, noteresourceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteresourceshapeIdent, noteresourceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteresourceshapeIdent, noteresourceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteresourceshapeIdent, noteresourceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteresourceshapeIdent, noteresourceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.FontColor)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", noteshapeIdent, noteshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", noteshapeIdent, noteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", noteshapeIdent, noteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", noteshapeIdent, noteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", noteshapeIdent, noteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteshapeIdent, noteshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteshapeIdent, noteshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", noteshapeIdent, noteshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteshapeIdent, noteshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteshapeIdent, noteshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", noteshapeIdent, noteshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.FontColor)))
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
			if notetaskshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notetaskshapeIdent, notetaskshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notetaskshapeIdent))
			}
			if notetaskshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notetaskshapeIdent, notetaskshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notetaskshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notetaskshapeIdent, notetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notetaskshapeIdent, notetaskshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notetaskshapeIdent, notetaskshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notetaskshapeIdent, notetaskshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notetaskshapeIdent, notetaskshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.FontColor)))
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
		for _, product := range __gong__sortStageSetInstances(stageSet.Stage.Products, stageSet.Stage.Product_stagedOrder) {
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
			if product.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", productIdent, product.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", productIdent))
			}
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
		for _, productcompositionshape := range __gong__sortStageSetInstances(stageSet.Stage.ProductCompositionShapes, stageSet.Stage.ProductCompositionShape_stagedOrder) {
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
			if productcompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", productcompositionshapeIdent, productcompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", productcompositionshapeIdent))
			}
			if productcompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", productcompositionshapeIdent, productcompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", productcompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", productcompositionshapeIdent, productcompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productcompositionshapeIdent, productcompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", productcompositionshapeIdent, productcompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", productcompositionshapeIdent, productcompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", productcompositionshapeIdent, productcompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.FontColor)))
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
		for _, productreferenceshape := range __gong__sortStageSetInstances(stageSet.Stage.ProductReferenceShapes, stageSet.Stage.ProductReferenceShape_stagedOrder) {
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
			if productreferenceshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", productreferenceshapeIdent, productreferenceshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", productreferenceshapeIdent))
			}
			if productreferenceshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", productreferenceshapeIdent, productreferenceshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", productreferenceshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", productreferenceshapeIdent, productreferenceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productreferenceshapeIdent, productreferenceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", productreferenceshapeIdent, productreferenceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", productreferenceshapeIdent, productreferenceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", productreferenceshapeIdent, productreferenceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.FontColor)))
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
		for _, productshape := range __gong__sortStageSetInstances(stageSet.Stage.ProductShapes, stageSet.Stage.ProductShape_stagedOrder) {
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", productshapeIdent, productshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", productshapeIdent, productshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", productshapeIdent, productshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", productshapeIdent, productshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", productshapeIdent, productshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productshapeIdent, productshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", productshapeIdent, productshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", productshapeIdent, productshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", productshapeIdent, productshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", productshapeIdent, productshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", productshapeIdent, productshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.FontColor)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", resourceIdent, __gong__toRawStringLiteral(resource.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", resourceIdent, __gong__toRawStringLiteral(resource.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", resourceIdent, resource.IsExpanded))
			if resource.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", resourceIdent, resource.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", resourceIdent))
			}
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
		for _, resourcecompositionshape := range __gong__sortStageSetInstances(stageSet.Stage.ResourceCompositionShapes, stageSet.Stage.ResourceCompositionShape_stagedOrder) {
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
			if resourcecompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", resourcecompositionshapeIdent, resourcecompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", resourcecompositionshapeIdent))
			}
			if resourcecompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", resourcecompositionshapeIdent, resourcecompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", resourcecompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourcecompositionshapeIdent, resourcecompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", resourcecompositionshapeIdent, resourcecompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", resourcecompositionshapeIdent, resourcecompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", resourcecompositionshapeIdent, resourcecompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.FontColor)))
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
		for _, resourceshape := range __gong__sortStageSetInstances(stageSet.Stage.ResourceShapes, stageSet.Stage.ResourceShape_stagedOrder) {
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", resourceshapeIdent, resourceshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", resourceshapeIdent, resourceshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", resourceshapeIdent, resourceshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", resourceshapeIdent, resourceshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", resourceshapeIdent, resourceshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourceshapeIdent, resourceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", resourceshapeIdent, resourceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", resourceshapeIdent, resourceshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", resourceshapeIdent, resourceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", resourceshapeIdent, resourceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", resourceshapeIdent, resourceshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.FontColor)))
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
		for _, resourcetaskshape := range __gong__sortStageSetInstances(stageSet.Stage.ResourceTaskShapes, stageSet.Stage.ResourceTaskShape_stagedOrder) {
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
			if resourcetaskshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", resourcetaskshapeIdent, resourcetaskshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", resourcetaskshapeIdent))
			}
			if resourcetaskshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", resourcetaskshapeIdent, resourcetaskshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", resourcetaskshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", resourcetaskshapeIdent, resourcetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourcetaskshapeIdent, resourcetaskshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", resourcetaskshapeIdent, resourcetaskshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", resourcetaskshapeIdent, resourcetaskshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", resourcetaskshapeIdent, resourcetaskshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.FontColor)))
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
		for _, task := range __gong__sortStageSetInstances(stageSet.Stage.Tasks, stageSet.Stage.Task_stagedOrder) {
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsAllDay = %t", taskIdent, task.IsAllDay))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMilestone = %t", taskIdent, task.IsMilestone))
			if task.DependencyType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.DependencyType = models.%s", taskIdent, task.DependencyType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.DependencyType = \"\"", taskIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationYears = %f", taskIdent, task.DependencyDurationYears))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationMonths = %f", taskIdent, task.DependencyDurationMonths))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationWeeks = %f", taskIdent, task.DependencyDurationWeeks))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationDays = %f", taskIdent, task.DependencyDurationDays))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationHours = %f", taskIdent, task.DependencyDurationHours))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationYears = %f", taskIdent, task.DurationYears))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationMonths = %f", taskIdent, task.DurationMonths))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationWeeks = %f", taskIdent, task.DurationWeeks))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationDays = %f", taskIdent, task.DurationDays))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationHours = %f", taskIdent, task.DurationHours))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndDateComputedFromDuration = %t", taskIdent, task.IsEndDateComputedFromDuration))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCompletion = %t", taskIdent, task.IsWithCompletion))
			if task.Completion.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Completion = models.%s", taskIdent, task.Completion.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Completion = \"\"", taskIdent))
			}
			if task.TextPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.TextPosition = models.%s", taskIdent, task.TextPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.TextPosition = \"\"", taskIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.XOffset = %f", taskIdent, task.XOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.YOffset = %f", taskIdent, task.YOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", taskIdent, task.IsImport))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInputsNodeExpanded = %t", taskIdent, task.IsInputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOutputsNodeExpanded = %t", taskIdent, task.IsOutputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskIdent, __gong__toRawStringLiteral(task.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskIdent, task.IsExpanded))
			if task.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", taskIdent, task.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", taskIdent))
			}
			for _, elem := range task.TaskGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroups = append(%s.TaskGroups, %s)", taskIdent, taskIdent, targetIdent))
			}
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
		for _, taskcompositionshape := range __gong__sortStageSetInstances(stageSet.Stage.TaskCompositionShapes, stageSet.Stage.TaskCompositionShape_stagedOrder) {
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
			if taskcompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskcompositionshapeIdent, taskcompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskcompositionshapeIdent))
			}
			if taskcompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskcompositionshapeIdent, taskcompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskcompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskcompositionshapeIdent, taskcompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskcompositionshapeIdent, taskcompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskcompositionshapeIdent, taskcompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskcompositionshapeIdent, taskcompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskcompositionshapeIdent, taskcompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.FontColor)))
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
		for _, taskgroup := range __gong__sortStageSetInstances(stageSet.Stage.TaskGroups, stageSet.Stage.TaskGroup_stagedOrder) {
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
		for _, taskgroupshape := range __gong__sortStageSetInstances(stageSet.Stage.TaskGroupShapes, stageSet.Stage.TaskGroupShape_stagedOrder) {
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskgroupshapeIdent, taskgroupshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", taskgroupshapeIdent, taskgroupshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskgroupshapeIdent, taskgroupshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskgroupshapeIdent, taskgroupshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", taskgroupshapeIdent, taskgroupshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.FontColor)))
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
		for _, taskinputshape := range __gong__sortStageSetInstances(stageSet.Stage.TaskInputShapes, stageSet.Stage.TaskInputShape_stagedOrder) {
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
			if taskinputshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskinputshapeIdent, taskinputshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskinputshapeIdent))
			}
			if taskinputshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskinputshapeIdent, taskinputshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskinputshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskinputshapeIdent, taskinputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskinputshapeIdent, taskinputshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskinputshapeIdent, taskinputshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskinputshapeIdent, taskinputshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskinputshapeIdent, taskinputshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.FontColor)))
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
		for _, taskoutputshape := range __gong__sortStageSetInstances(stageSet.Stage.TaskOutputShapes, stageSet.Stage.TaskOutputShape_stagedOrder) {
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
			if taskoutputshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskoutputshapeIdent, taskoutputshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskoutputshapeIdent))
			}
			if taskoutputshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskoutputshapeIdent, taskoutputshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskoutputshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskoutputshapeIdent, taskoutputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskoutputshapeIdent, taskoutputshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskoutputshapeIdent, taskoutputshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskoutputshapeIdent, taskoutputshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskoutputshapeIdent, taskoutputshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.FontColor)))
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
		for _, taskpredecessorshape := range __gong__sortStageSetInstances(stageSet.Stage.TaskPredecessorShapes, stageSet.Stage.TaskPredecessorShape_stagedOrder) {
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
			if taskpredecessorshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskpredecessorshapeIdent, taskpredecessorshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskpredecessorshapeIdent))
			}
			if taskpredecessorshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskpredecessorshapeIdent, taskpredecessorshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskpredecessorshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskpredecessorshapeIdent, taskpredecessorshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskpredecessorshapeIdent, taskpredecessorshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskpredecessorshapeIdent, taskpredecessorshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskpredecessorshapeIdent, taskpredecessorshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.FontColor)))
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
		for _, taskshape := range __gong__sortStageSetInstances(stageSet.Stage.TaskShapes, stageSet.Stage.TaskShape_stagedOrder) {
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
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalOffset = %f", taskshapeIdent, taskshape.VerticalOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayVerticalBar = %t", taskshapeIdent, taskshape.DisplayVerticalBar))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", taskshapeIdent, taskshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", taskshapeIdent, taskshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", taskshapeIdent, taskshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", taskshapeIdent, taskshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", taskshapeIdent, taskshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskshapeIdent, taskshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskshapeIdent, taskshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", taskshapeIdent, taskshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskshapeIdent, taskshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskshapeIdent, taskshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", taskshapeIdent, taskshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.FontColor)))
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


	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/dsm/project/go/models/probe/cmd/probe@%s stage.go
*/
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
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, declarations.String(), values.String(), pointers.String())

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
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteProductShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteProductShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteResourceShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteTaskShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Product":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Product), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ProductCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProductCompositionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ProductReferenceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProductReferenceShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ProductShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProductShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Resource":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Resource), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ResourceCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceCompositionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ResourceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ResourceTaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceTaskShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Task":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Task), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskCompositionShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskCompositionShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskGroup), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskGroupShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskGroupShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskInputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskInputShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskOutputShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskOutputShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskPredecessorShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskPredecessorShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TaskShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaskShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						inst.ComputedStart = GongExtractDate(rhs)
					case "ComputedEnd":
						inst.ComputedEnd = GongExtractDate(rhs)
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "DrawVerticalTimeLines":
						inst.DrawVerticalTimeLines = GongExtractBool(rhs)
					case "DrawSecondaryVerticalTimeLines":
						inst.DrawSecondaryVerticalTimeLines = GongExtractBool(rhs)
					case "HideWeekendsPeriod":
						inst.HideWeekendsPeriod = GongExtractBool(rhs)
					case "UseManualStartAndEndDates":
						inst.UseManualStartAndEndDates = GongExtractBool(rhs)
					case "ManualStart":
						inst.ManualStart = GongExtractDate(rhs)
					case "ManualEnd":
						inst.ManualEnd = GongExtractDate(rhs)
					case "TimeStep":
						inst.TimeStep = GongExtractInt(rhs)
					case "TimeStepScale":
						GongUnmarshallEnum(&inst.TimeStepScale, rhs)
					case "SecondaryTimeStep":
						inst.SecondaryTimeStep = GongExtractInt(rhs)
					case "SecondaryTimeStepScale":
						GongUnmarshallEnum(&inst.SecondaryTimeStepScale, rhs)
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
					case "Group_Stroke":
						inst.Group_Stroke = GongExtractString(rhs)
					case "Group_StrokeWidth":
						inst.Group_StrokeWidth = GongExtractFloat(rhs)
					case "Group_StrokeDashArray":
						inst.Group_StrokeDashArray = GongExtractString(rhs)
					case "DateYOffset":
						inst.DateYOffset = GongExtractFloat(rhs)
					case "AlignOnBeginningOfTimeScale":
						inst.AlignOnBeginningOfTimeScale = GongExtractBool(rhs)
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
						__gong__assignSliceOfPointers(&inst.Product_Shapes, rhs, identifierMap)
					case "ProductsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ProductsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPBSNodeExpanded":
						inst.IsPBSNodeExpanded = GongExtractBool(rhs)
					case "ProductComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.ProductComposition_Shapes, rhs, identifierMap)
					case "ProductReference_Shapes":
						__gong__assignSliceOfPointers(&inst.ProductReference_Shapes, rhs, identifierMap)
					case "IsWBSNodeExpanded":
						inst.IsWBSNodeExpanded = GongExtractBool(rhs)
					case "Task_Shapes":
						__gong__assignSliceOfPointers(&inst.Task_Shapes, rhs, identifierMap)
					case "TasksWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseNodeIsExpanded, rhs, identifierMap)
					case "TasksWhoseInputNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseInputNodeIsExpanded, rhs, identifierMap)
					case "TasksWhoseOutputNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhoseOutputNodeIsExpanded, rhs, identifierMap)
					case "TasksWhosePredecessorNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TasksWhosePredecessorNodeIsExpanded, rhs, identifierMap)
					case "IsTaskGroupsNodeExpanded":
						inst.IsTaskGroupsNodeExpanded = GongExtractBool(rhs)
					case "TaskGroupShapes":
						__gong__assignSliceOfPointers(&inst.TaskGroupShapes, rhs, identifierMap)
					case "TaskGroupsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.TaskGroupsWhoseNodeIsExpanded, rhs, identifierMap)
					case "TaskComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.TaskComposition_Shapes, rhs, identifierMap)
					case "TaskInputShapes":
						__gong__assignSliceOfPointers(&inst.TaskInputShapes, rhs, identifierMap)
					case "TaskOutputShapes":
						__gong__assignSliceOfPointers(&inst.TaskOutputShapes, rhs, identifierMap)
					case "TaskPredecessorShapes":
						__gong__assignSliceOfPointers(&inst.TaskPredecessorShapes, rhs, identifierMap)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NoteProductShapes":
						__gong__assignSliceOfPointers(&inst.NoteProductShapes, rhs, identifierMap)
					case "NoteTaskShapes":
						__gong__assignSliceOfPointers(&inst.NoteTaskShapes, rhs, identifierMap)
					case "NoteResourceShapes":
						__gong__assignSliceOfPointers(&inst.NoteResourceShapes, rhs, identifierMap)
					case "Resource_Shapes":
						__gong__assignSliceOfPointers(&inst.Resource_Shapes, rhs, identifierMap)
					case "ResourcesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ResourcesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsResourcesNodeExpanded":
						inst.IsResourcesNodeExpanded = GongExtractBool(rhs)
					case "ResourceComposition_Shapes":
						__gong__assignSliceOfPointers(&inst.ResourceComposition_Shapes, rhs, identifierMap)
					case "ResourceTaskShapes":
						__gong__assignSliceOfPointers(&inst.ResourceTaskShapes, rhs, identifierMap)
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
					case "RootProducts":
						__gong__assignSliceOfPointers(&inst.RootProducts, rhs, identifierMap)
					case "RootTasks":
						__gong__assignSliceOfPointers(&inst.RootTasks, rhs, identifierMap)
					case "RootTaskGroups":
						__gong__assignSliceOfPointers(&inst.RootTaskGroups, rhs, identifierMap)
					case "RootResources":
						__gong__assignSliceOfPointers(&inst.RootResources, rhs, identifierMap)
					case "Notes":
						__gong__assignSliceOfPointers(&inst.Notes, rhs, identifierMap)
					case "Diagrams":
						__gong__assignSliceOfPointers(&inst.Diagrams, rhs, identifierMap)
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
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					case "Products":
						__gong__assignSliceOfPointers(&inst.Products, rhs, identifierMap)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					case "Resources":
						__gong__assignSliceOfPointers(&inst.Resources, rhs, identifierMap)
					}
				case *NoteProductShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *NoteResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
				case *Product:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SubProducts":
						__gong__assignSliceOfPointers(&inst.SubProducts, rhs, identifierMap)
					case "IsProducersNodeExpanded":
						inst.IsProducersNodeExpanded = GongExtractBool(rhs)
					case "IsConsumersNodeExpanded":
						inst.IsConsumersNodeExpanded = GongExtractBool(rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedProduct":
						__gong__assignPointer(&inst.ReferencedProduct, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					}
				case *ProductCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *ProductReferenceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
					case "ReferencedProduct":
						__gong__assignPointer(&inst.ReferencedProduct, rhs, identifierMap)
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
				case *ProductShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
					case "IsShowType":
						inst.IsShowType = GongExtractBool(rhs)
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
				case *Resource:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					case "SubResources":
						__gong__assignSliceOfPointers(&inst.SubResources, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedResource":
						__gong__assignPointer(&inst.ReferencedResource, rhs, identifierMap)
					}
				case *ResourceCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
				case *ResourceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
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
				case *ResourceTaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
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
				case *Task:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Start":
						inst.Start = GongExtractDate(rhs)
					case "End":
						inst.End = GongExtractDate(rhs)
					case "IsAllDay":
						inst.IsAllDay = GongExtractBool(rhs)
					case "IsMilestone":
						inst.IsMilestone = GongExtractBool(rhs)
					case "TaskGroups":
						__gong__assignSliceOfPointers(&inst.TaskGroups, rhs, identifierMap)
					case "Predecessors":
						__gong__assignSliceOfPointers(&inst.Predecessors, rhs, identifierMap)
					case "DependencyType":
						GongUnmarshallEnum(&inst.DependencyType, rhs)
					case "DependencyDurationYears":
						inst.DependencyDurationYears = GongExtractFloat(rhs)
					case "DependencyDurationMonths":
						inst.DependencyDurationMonths = GongExtractFloat(rhs)
					case "DependencyDurationWeeks":
						inst.DependencyDurationWeeks = GongExtractFloat(rhs)
					case "DependencyDurationDays":
						inst.DependencyDurationDays = GongExtractFloat(rhs)
					case "DependencyDurationHours":
						inst.DependencyDurationHours = GongExtractFloat(rhs)
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
					case "Inputs":
						__gong__assignSliceOfPointers(&inst.Inputs, rhs, identifierMap)
					case "Outputs":
						__gong__assignSliceOfPointers(&inst.Outputs, rhs, identifierMap)
					case "SubTasks":
						__gong__assignSliceOfPointers(&inst.SubTasks, rhs, identifierMap)
					case "IsWithCompletion":
						inst.IsWithCompletion = GongExtractBool(rhs)
					case "Completion":
						GongUnmarshallEnum(&inst.Completion, rhs)
					case "TaskGroupsToDisplay":
						__gong__assignSliceOfPointers(&inst.TaskGroupsToDisplay, rhs, identifierMap)
					case "TextPosition":
						GongUnmarshallEnum(&inst.TextPosition, rhs)
					case "XOffset":
						inst.XOffset = GongExtractFloat(rhs)
					case "YOffset":
						inst.YOffset = GongExtractFloat(rhs)
					case "IsImport":
						inst.IsImport = GongExtractBool(rhs)
					case "ReferencedTask":
						__gong__assignPointer(&inst.ReferencedTask, rhs, identifierMap)
					case "IsInputsNodeExpanded":
						inst.IsInputsNodeExpanded = GongExtractBool(rhs)
					case "IsOutputsNodeExpanded":
						inst.IsOutputsNodeExpanded = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "LayoutDirection":
						GongUnmarshallEnum(&inst.LayoutDirection, rhs)
					}
				case *TaskCompositionShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
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
				case *TaskGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Tasks":
						__gong__assignSliceOfPointers(&inst.Tasks, rhs, identifierMap)
					}
				case *TaskGroupShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "TaskGroup":
						__gong__assignPointer(&inst.TaskGroup, rhs, identifierMap)
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
				case *TaskInputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
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
				case *TaskOutputShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
					case "Product":
						__gong__assignPointer(&inst.Product, rhs, identifierMap)
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
				case *TaskPredecessorShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Predecessor":
						__gong__assignPointer(&inst.Predecessor, rhs, identifierMap)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
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
				case *TaskShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Task":
						__gong__assignPointer(&inst.Task, rhs, identifierMap)
					case "IsShowDate":
						inst.IsShowDate = GongExtractBool(rhs)
					case "VerticalOffset":
						inst.VerticalOffset = GongExtractFloat(rhs)
					case "DisplayVerticalBar":
						inst.DisplayVerticalBar = GongExtractBool(rhs)
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

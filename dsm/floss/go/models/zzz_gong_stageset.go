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
		for _, compareanalysis := range __gong__sortStageSetInstances(stageSet.Stage.CompareAnalysiss, stageSet.Stage.CompareAnalysis_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			compareanalysisIdent := "__models" + compareanalysis.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CompareAnalysis{Name: %s}).Stage(stageSet.Stage)", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mu = %f", compareanalysisIdent, compareanalysis.Mu))
			values.WriteString(fmt.Sprintf("\n\t%s.Epsilon = %f", compareanalysisIdent, compareanalysis.Epsilon))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", compareanalysisIdent, compareanalysis.IsExpanded))
			if compareanalysis.FromSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + compareanalysis.FromSystem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FromSystem = %s", compareanalysisIdent, targetIdent))
			}
			if compareanalysis.ToSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + compareanalysis.ToSystem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ToSystem = %s", compareanalysisIdent, targetIdent))
			}
			for _, elem := range compareanalysis.DiagramFlossEquations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquations = append(%s.DiagramFlossEquations, %s)", compareanalysisIdent, compareanalysisIdent, targetIdent))
			}
			for _, elem := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquationsWhoseNodeIsExpanded = append(%s.DiagramFlossEquationsWhoseNodeIsExpanded, %s)", compareanalysisIdent, compareanalysisIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, complexity := range __gong__sortStageSetInstances(stageSet.Stage.Complexitys, stageSet.Stage.Complexity_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			complexityIdent := "__models" + complexity.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Complexity{Name: %s}).Stage(stageSet.Stage)", complexityIdent, __gong__toRawStringLiteral(complexity.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complexityIdent, __gong__toRawStringLiteral(complexity.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", complexityIdent, complexity.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", complexityIdent, __gong__toRawStringLiteral(complexity.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", complexityIdent, __gong__toRawStringLiteral(complexity.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", complexityIdent, complexity.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, diagramflossequation := range __gong__sortStageSetInstances(stageSet.Stage.DiagramFlossEquations, stageSet.Stage.DiagramFlossEquation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramflossequationIdent := "__models" + diagramflossequation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramFlossEquation{Name: %s}).Stage(stageSet.Stage)", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scale = %f", diagramflossequationIdent, diagramflossequation.Scale))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", diagramflossequationIdent, __gong__toRawStringLiteral(string(diagramflossequation.FontSize))))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramflossequationIdent, diagramflossequation.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramflossequationIdent, diagramflossequation.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramflossequationIdent, diagramflossequation.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInDelta3ColumnsMode = %t", diagramflossequationIdent, diagramflossequation.IsInDelta3ColumnsMode))
			values.WriteString(fmt.Sprintf("\n\t%s.AreQuantitativeElementsVisible = %t", diagramflossequationIdent, diagramflossequation.AreQuantitativeElementsVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.AreSubsystemsVisible = %t", diagramflossequationIdent, diagramflossequation.AreSubsystemsVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.AreCommonElementsHidden = %t", diagramflossequationIdent, diagramflossequation.AreCommonElementsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.AreCPEArrowsVisible = %t", diagramflossequationIdent, diagramflossequation.AreCPEArrowsVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.AreColumnTitlesVisible = %t", diagramflossequationIdent, diagramflossequation.AreColumnTitlesVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramflossequationIdent, diagramflossequation.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramflossequationIdent, diagramflossequation.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramflossequationIdent, diagramflossequation.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramflossequationIdent, diagramflossequation.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsEffortsNodeExpanded))
			for _, elem := range diagramflossequation.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NoteComplexityShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteComplexityShapes = append(%s.NoteComplexityShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NotePerformanceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePerformanceShapes = append(%s.NotePerformanceShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NoteEffortShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteEffortShapes = append(%s.NoteEffortShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.EffortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, effort := range __gong__sortStageSetInstances(stageSet.Stage.Efforts, stageSet.Stage.Effort_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			effortIdent := "__models" + effort.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Effort{Name: %s}).Stage(stageSet.Stage)", effortIdent, __gong__toRawStringLiteral(effort.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effortIdent, __gong__toRawStringLiteral(effort.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", effortIdent, effort.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", effortIdent, __gong__toRawStringLiteral(effort.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", effortIdent, __gong__toRawStringLiteral(effort.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", effortIdent, effort.IsExpanded))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemsNodeExpanded = %t", libraryIdent, library.IsSystemsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", libraryIdent, library.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", libraryIdent, library.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", libraryIdent, library.IsEffortsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsCompareAnalysisNodeExpanded = %t", libraryIdent, library.IsCompareAnalysisNodeExpanded))
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
			for _, elem := range library.RootSystems {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootSystems = append(%s.RootSystems, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootComplexitys {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootComplexitys = append(%s.RootComplexitys, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootPerformances {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootPerformances = append(%s.RootPerformances, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootEfforts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootEfforts = append(%s.RootEfforts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootCompareAnalysis {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootCompareAnalysis = append(%s.RootCompareAnalysis, %s)", libraryIdent, libraryIdent, targetIdent))
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
			for _, elem := range library.ComplexitysWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.PerformancesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.EffortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.CompareAnalysisWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompareAnalysisWhoseNodeIsExpanded = append(%s.CompareAnalysisWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", noteIdent, note.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", noteIdent, note.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", noteIdent, note.IsEffortsNodeExpanded))
			for _, elem := range note.Complexities {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexities = append(%s.Complexities, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Performances {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performances = append(%s.Performances, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Efforts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Efforts = append(%s.Efforts, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notecomplexityshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteComplexityShapes, stageSet.Stage.NoteComplexityShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notecomplexityshapeIdent := "__models" + notecomplexityshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteComplexityShape{Name: %s}).Stage(stageSet.Stage)", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notecomplexityshapeIdent, notecomplexityshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notecomplexityshapeIdent, notecomplexityshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(string(notecomplexityshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(string(notecomplexityshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notecomplexityshapeIdent, notecomplexityshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notecomplexityshapeIdent, notecomplexityshape.IsHidden))
			if notecomplexityshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notecomplexityshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notecomplexityshapeIdent, targetIdent))
			}
			if notecomplexityshape.Complexity != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notecomplexityshape.Complexity.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexity = %s", notecomplexityshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, noteeffortshape := range __gong__sortStageSetInstances(stageSet.Stage.NoteEffortShapes, stageSet.Stage.NoteEffortShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteeffortshapeIdent := "__models" + noteeffortshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteEffortShape{Name: %s}).Stage(stageSet.Stage)", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteeffortshapeIdent, noteeffortshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteeffortshapeIdent, noteeffortshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(string(noteeffortshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(string(noteeffortshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteeffortshapeIdent, noteeffortshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteeffortshapeIdent, noteeffortshape.IsHidden))
			if noteeffortshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteeffortshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteeffortshapeIdent, targetIdent))
			}
			if noteeffortshape.Effort != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteeffortshape.Effort.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effort = %s", noteeffortshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, noteperformanceshape := range __gong__sortStageSetInstances(stageSet.Stage.NotePerformanceShapes, stageSet.Stage.NotePerformanceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteperformanceshapeIdent := "__models" + noteperformanceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NotePerformanceShape{Name: %s}).Stage(stageSet.Stage)", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteperformanceshapeIdent, noteperformanceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteperformanceshapeIdent, noteperformanceshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(string(noteperformanceshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(string(noteperformanceshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteperformanceshapeIdent, noteperformanceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteperformanceshapeIdent, noteperformanceshape.IsHidden))
			if noteperformanceshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteperformanceshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteperformanceshapeIdent, targetIdent))
			}
			if noteperformanceshape.Performance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + noteperformanceshape.Performance.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performance = %s", noteperformanceshapeIdent, targetIdent))
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
		for _, performance := range __gong__sortStageSetInstances(stageSet.Stage.Performances, stageSet.Stage.Performance_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			performanceIdent := "__models" + performance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Performance{Name: %s}).Stage(stageSet.Stage)", performanceIdent, __gong__toRawStringLiteral(performance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", performanceIdent, __gong__toRawStringLiteral(performance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", performanceIdent, performance.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", performanceIdent, __gong__toRawStringLiteral(performance.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", performanceIdent, __gong__toRawStringLiteral(performance.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", performanceIdent, performance.IsExpanded))
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
			values.WriteString(fmt.Sprintf("\n\t%s.AreCPEsCompoundedFromSubSystems = %t", systemIdent, system.AreCPEsCompoundedFromSubSystems))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", systemIdent, __gong__toRawStringLiteral(system.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemIdent, system.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", systemIdent, __gong__toRawStringLiteral(system.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", systemIdent, system.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubSystemNodeExpanded = %t", systemIdent, system.IsSubSystemNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", systemIdent, system.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", systemIdent, system.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", systemIdent, system.IsEffortsNodeExpanded))
			for _, elem := range system.Complexities {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexities = append(%s.Complexities, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Performances {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performances = append(%s.Performances, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Efforts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Efforts = append(%s.Efforts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.SubSystems {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubSystems = append(%s.SubSystems, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramFlossEquations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquations = append(%s.DiagramFlossEquations, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquationsWhoseNodeIsExpanded = append(%s.DiagramFlossEquationsWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ComplexitysWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.PerformancesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.EffortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
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
		case "github.com/fullstack-lang/gong/dsm/floss/go/models":
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
				case "CompareAnalysis":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CompareAnalysis), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Complexity":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Complexity), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DiagramFlossEquation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DiagramFlossEquation), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Effort":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Effort), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteComplexityShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteComplexityShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteEffortShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteEffortShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NotePerformanceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePerformanceShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Performance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Performance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "System":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *CompareAnalysis:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FromSystem":
						__gong__assignPointer(&inst.FromSystem, rhs, identifierMap)
					case "ToSystem":
						__gong__assignPointer(&inst.ToSystem, rhs, identifierMap)
					case "Mu":
						inst.Mu = GongExtractFloat(rhs)
					case "Epsilon":
						inst.Epsilon = GongExtractFloat(rhs)
					case "DiagramFlossEquations":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquations, rhs, identifierMap)
					case "DiagramFlossEquationsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquationsWhoseNodeIsExpanded, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Complexity:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Strength":
						inst.Strength = GongExtractFloat(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *DiagramFlossEquation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Scale":
						inst.Scale = GongExtractFloat(rhs)
					case "FontSize":
						inst.FontSize = FontSize(GongExtractString(rhs))
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "IsEditable_":
						inst.IsEditable_ = GongExtractBool(rhs)
					case "IsInDelta3ColumnsMode":
						inst.IsInDelta3ColumnsMode = GongExtractBool(rhs)
					case "AreQuantitativeElementsVisible":
						inst.AreQuantitativeElementsVisible = GongExtractBool(rhs)
					case "AreSubsystemsVisible":
						inst.AreSubsystemsVisible = GongExtractBool(rhs)
					case "AreCommonElementsHidden":
						inst.AreCommonElementsHidden = GongExtractBool(rhs)
					case "AreCPEArrowsVisible":
						inst.AreCPEArrowsVisible = GongExtractBool(rhs)
					case "AreColumnTitlesVisible":
						inst.AreColumnTitlesVisible = GongExtractBool(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "DefaultBoxWidth":
						inst.DefaultBoxWidth = GongExtractFloat(rhs)
					case "DefaultBoxHeigth":
						inst.DefaultBoxHeigth = GongExtractFloat(rhs)
					case "Note_Shapes":
						__gong__assignSliceOfPointers(&inst.Note_Shapes, rhs, identifierMap)
					case "NoteComplexityShapes":
						__gong__assignSliceOfPointers(&inst.NoteComplexityShapes, rhs, identifierMap)
					case "NotePerformanceShapes":
						__gong__assignSliceOfPointers(&inst.NotePerformanceShapes, rhs, identifierMap)
					case "NoteEffortShapes":
						__gong__assignSliceOfPointers(&inst.NoteEffortShapes, rhs, identifierMap)
					case "IsNotesNodeExpanded":
						inst.IsNotesNodeExpanded = GongExtractBool(rhs)
					case "NotesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.NotesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ComplexitysWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PerformancesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EffortsWhoseNodeIsExpanded, rhs, identifierMap)
					}
				case *Effort:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Strength":
						inst.Strength = GongExtractFloat(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
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
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "RootSystems":
						__gong__assignSliceOfPointers(&inst.RootSystems, rhs, identifierMap)
					case "RootComplexitys":
						__gong__assignSliceOfPointers(&inst.RootComplexitys, rhs, identifierMap)
					case "RootPerformances":
						__gong__assignSliceOfPointers(&inst.RootPerformances, rhs, identifierMap)
					case "RootEfforts":
						__gong__assignSliceOfPointers(&inst.RootEfforts, rhs, identifierMap)
					case "RootCompareAnalysis":
						__gong__assignSliceOfPointers(&inst.RootCompareAnalysis, rhs, identifierMap)
					case "RootNotes":
						__gong__assignSliceOfPointers(&inst.RootNotes, rhs, identifierMap)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SubLibrariesWhoseNodeIsExpanded, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "IsSystemsNodeExpanded":
						inst.IsSystemsNodeExpanded = GongExtractBool(rhs)
					case "SystemsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SystemsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ComplexitysWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PerformancesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EffortsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsCompareAnalysisNodeExpanded":
						inst.IsCompareAnalysisNodeExpanded = GongExtractBool(rhs)
					case "CompareAnalysisWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.CompareAnalysisWhoseNodeIsExpanded, rhs, identifierMap)
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
					case "Complexities":
						__gong__assignSliceOfPointers(&inst.Complexities, rhs, identifierMap)
					case "Performances":
						__gong__assignSliceOfPointers(&inst.Performances, rhs, identifierMap)
					case "Efforts":
						__gong__assignSliceOfPointers(&inst.Efforts, rhs, identifierMap)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					}
				case *NoteComplexityShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Complexity":
						__gong__assignPointer(&inst.Complexity, rhs, identifierMap)
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
				case *NoteEffortShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Effort":
						__gong__assignPointer(&inst.Effort, rhs, identifierMap)
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
				case *NotePerformanceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Performance":
						__gong__assignPointer(&inst.Performance, rhs, identifierMap)
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
				case *Performance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Strength":
						inst.Strength = GongExtractFloat(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *System:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Complexities":
						__gong__assignSliceOfPointers(&inst.Complexities, rhs, identifierMap)
					case "Performances":
						__gong__assignSliceOfPointers(&inst.Performances, rhs, identifierMap)
					case "Efforts":
						__gong__assignSliceOfPointers(&inst.Efforts, rhs, identifierMap)
					case "SubSystems":
						__gong__assignSliceOfPointers(&inst.SubSystems, rhs, identifierMap)
					case "AreCPEsCompoundedFromSubSystems":
						inst.AreCPEsCompoundedFromSubSystems = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SVG_Path":
						inst.SVG_Path = GongExtractString(rhs)
					case "InverseAppliedScaling":
						inst.InverseAppliedScaling = GongExtractFloat(rhs)
					case "DiagramFlossEquations":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquations, rhs, identifierMap)
					case "DiagramFlossEquationsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.DiagramFlossEquationsWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsSubSystemNodeExpanded":
						inst.IsSubSystemNodeExpanded = GongExtractBool(rhs)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.ComplexitysWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.PerformancesWhoseNodeIsExpanded, rhs, identifierMap)
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.EffortsWhoseNodeIsExpanded, rhs, identifierMap)
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

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
		compareanalysisOrdered := []*CompareAnalysis{}
		for compareanalysis := range stageSet.Stage.CompareAnalysiss {
			compareanalysisOrdered = append(compareanalysisOrdered, compareanalysis)
		}
		sort.Slice(compareanalysisOrdered, func(i, j int) bool {
			return stageSet.Stage.CompareAnalysis_stagedOrder[compareanalysisOrdered[i]] < stageSet.Stage.CompareAnalysis_stagedOrder[compareanalysisOrdered[j]]
		})
		for _, compareanalysis := range compareanalysisOrdered {
			compareanalysisIdent := "__stage_0" + compareanalysis.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CompareAnalysis{Name: %s}).Stage(stageSet.Stage)", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mu = %f", compareanalysisIdent, compareanalysis.Mu))
			values.WriteString(fmt.Sprintf("\n\t%s.Epsilon = %f", compareanalysisIdent, compareanalysis.Epsilon))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", compareanalysisIdent, compareanalysis.IsExpanded))
			if compareanalysis.FromSystem != nil {
				targetIdent := "__stage_0" + compareanalysis.FromSystem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FromSystem = %s", compareanalysisIdent, targetIdent))
			}
			if compareanalysis.ToSystem != nil {
				targetIdent := "__stage_0" + compareanalysis.ToSystem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ToSystem = %s", compareanalysisIdent, targetIdent))
			}
			for _, elem := range compareanalysis.DiagramFlossEquations {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquations = append(%s.DiagramFlossEquations, %s)", compareanalysisIdent, compareanalysisIdent, targetIdent))
			}
			for _, elem := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquationsWhoseNodeIsExpanded = append(%s.DiagramFlossEquationsWhoseNodeIsExpanded, %s)", compareanalysisIdent, compareanalysisIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		complexityOrdered := []*Complexity{}
		for complexity := range stageSet.Stage.Complexitys {
			complexityOrdered = append(complexityOrdered, complexity)
		}
		sort.Slice(complexityOrdered, func(i, j int) bool {
			return stageSet.Stage.Complexity_stagedOrder[complexityOrdered[i]] < stageSet.Stage.Complexity_stagedOrder[complexityOrdered[j]]
		})
		for _, complexity := range complexityOrdered {
			complexityIdent := "__stage_0" + complexity.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Complexity{Name: %s}).Stage(stageSet.Stage)", complexityIdent, __gong__toRawStringLiteral(complexity.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complexityIdent, __gong__toRawStringLiteral(complexity.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", complexityIdent, complexity.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", complexityIdent, __gong__toRawStringLiteral(complexity.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", complexityIdent, __gong__toRawStringLiteral(complexity.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", complexityIdent, complexity.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		diagramflossequationOrdered := []*DiagramFlossEquation{}
		for diagramflossequation := range stageSet.Stage.DiagramFlossEquations {
			diagramflossequationOrdered = append(diagramflossequationOrdered, diagramflossequation)
		}
		sort.Slice(diagramflossequationOrdered, func(i, j int) bool {
			return stageSet.Stage.DiagramFlossEquation_stagedOrder[diagramflossequationOrdered[i]] < stageSet.Stage.DiagramFlossEquation_stagedOrder[diagramflossequationOrdered[j]]
		})
		for _, diagramflossequation := range diagramflossequationOrdered {
			diagramflossequationIdent := "__stage_0" + diagramflossequation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.DiagramFlossEquation{Name: %s}).Stage(stageSet.Stage)", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NoteComplexityShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteComplexityShapes = append(%s.NoteComplexityShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NotePerformanceShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePerformanceShapes = append(%s.NotePerformanceShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NoteEffortShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteEffortShapes = append(%s.NoteEffortShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NotesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.EffortsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		effortOrdered := []*Effort{}
		for effort := range stageSet.Stage.Efforts {
			effortOrdered = append(effortOrdered, effort)
		}
		sort.Slice(effortOrdered, func(i, j int) bool {
			return stageSet.Stage.Effort_stagedOrder[effortOrdered[i]] < stageSet.Stage.Effort_stagedOrder[effortOrdered[j]]
		})
		for _, effort := range effortOrdered {
			effortIdent := "__stage_0" + effort.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Effort{Name: %s}).Stage(stageSet.Stage)", effortIdent, __gong__toRawStringLiteral(effort.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effortIdent, __gong__toRawStringLiteral(effort.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", effortIdent, effort.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", effortIdent, __gong__toRawStringLiteral(effort.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", effortIdent, __gong__toRawStringLiteral(effort.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", effortIdent, effort.IsExpanded))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemsNodeExpanded = %t", libraryIdent, library.IsSystemsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", libraryIdent, library.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", libraryIdent, library.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", libraryIdent, library.IsEffortsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsCompareAnalysisNodeExpanded = %t", libraryIdent, library.IsCompareAnalysisNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", libraryIdent, library.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			for _, elem := range library.SubLibraries {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootSystems {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootSystems = append(%s.RootSystems, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootComplexitys {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootComplexitys = append(%s.RootComplexitys, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootPerformances {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootPerformances = append(%s.RootPerformances, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootEfforts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootEfforts = append(%s.RootEfforts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootCompareAnalysis {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootCompareAnalysis = append(%s.RootCompareAnalysis, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootNotes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNotes = append(%s.RootNotes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SystemsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ComplexitysWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.PerformancesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.EffortsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.CompareAnalysisWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompareAnalysisWhoseNodeIsExpanded = append(%s.CompareAnalysisWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", noteIdent, note.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", noteIdent, note.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", noteIdent, note.IsEffortsNodeExpanded))
			for _, elem := range note.Complexities {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexities = append(%s.Complexities, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Performances {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performances = append(%s.Performances, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Efforts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Efforts = append(%s.Efforts, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		notecomplexityshapeOrdered := []*NoteComplexityShape{}
		for notecomplexityshape := range stageSet.Stage.NoteComplexityShapes {
			notecomplexityshapeOrdered = append(notecomplexityshapeOrdered, notecomplexityshape)
		}
		sort.Slice(notecomplexityshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteComplexityShape_stagedOrder[notecomplexityshapeOrdered[i]] < stageSet.Stage.NoteComplexityShape_stagedOrder[notecomplexityshapeOrdered[j]]
		})
		for _, notecomplexityshape := range notecomplexityshapeOrdered {
			notecomplexityshapeIdent := "__stage_0" + notecomplexityshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NoteComplexityShape{Name: %s}).Stage(stageSet.Stage)", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notecomplexityshapeIdent, notecomplexityshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notecomplexityshapeIdent, notecomplexityshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(string(notecomplexityshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(string(notecomplexityshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notecomplexityshapeIdent, notecomplexityshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notecomplexityshapeIdent, notecomplexityshape.IsHidden))
			if notecomplexityshape.Note != nil {
				targetIdent := "__stage_0" + notecomplexityshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notecomplexityshapeIdent, targetIdent))
			}
			if notecomplexityshape.Complexity != nil {
				targetIdent := "__stage_0" + notecomplexityshape.Complexity.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexity = %s", notecomplexityshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteeffortshapeOrdered := []*NoteEffortShape{}
		for noteeffortshape := range stageSet.Stage.NoteEffortShapes {
			noteeffortshapeOrdered = append(noteeffortshapeOrdered, noteeffortshape)
		}
		sort.Slice(noteeffortshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NoteEffortShape_stagedOrder[noteeffortshapeOrdered[i]] < stageSet.Stage.NoteEffortShape_stagedOrder[noteeffortshapeOrdered[j]]
		})
		for _, noteeffortshape := range noteeffortshapeOrdered {
			noteeffortshapeIdent := "__stage_0" + noteeffortshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NoteEffortShape{Name: %s}).Stage(stageSet.Stage)", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteeffortshapeIdent, noteeffortshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteeffortshapeIdent, noteeffortshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(string(noteeffortshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(string(noteeffortshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteeffortshapeIdent, noteeffortshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteeffortshapeIdent, noteeffortshape.IsHidden))
			if noteeffortshape.Note != nil {
				targetIdent := "__stage_0" + noteeffortshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteeffortshapeIdent, targetIdent))
			}
			if noteeffortshape.Effort != nil {
				targetIdent := "__stage_0" + noteeffortshape.Effort.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effort = %s", noteeffortshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteperformanceshapeOrdered := []*NotePerformanceShape{}
		for noteperformanceshape := range stageSet.Stage.NotePerformanceShapes {
			noteperformanceshapeOrdered = append(noteperformanceshapeOrdered, noteperformanceshape)
		}
		sort.Slice(noteperformanceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.NotePerformanceShape_stagedOrder[noteperformanceshapeOrdered[i]] < stageSet.Stage.NotePerformanceShape_stagedOrder[noteperformanceshapeOrdered[j]]
		})
		for _, noteperformanceshape := range noteperformanceshapeOrdered {
			noteperformanceshapeIdent := "__stage_0" + noteperformanceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.NotePerformanceShape{Name: %s}).Stage(stageSet.Stage)", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteperformanceshapeIdent, noteperformanceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteperformanceshapeIdent, noteperformanceshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(string(noteperformanceshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(string(noteperformanceshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteperformanceshapeIdent, noteperformanceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteperformanceshapeIdent, noteperformanceshape.IsHidden))
			if noteperformanceshape.Note != nil {
				targetIdent := "__stage_0" + noteperformanceshape.Note.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteperformanceshapeIdent, targetIdent))
			}
			if noteperformanceshape.Performance != nil {
				targetIdent := "__stage_0" + noteperformanceshape.Performance.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performance = %s", noteperformanceshapeIdent, targetIdent))
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
		performanceOrdered := []*Performance{}
		for performance := range stageSet.Stage.Performances {
			performanceOrdered = append(performanceOrdered, performance)
		}
		sort.Slice(performanceOrdered, func(i, j int) bool {
			return stageSet.Stage.Performance_stagedOrder[performanceOrdered[i]] < stageSet.Stage.Performance_stagedOrder[performanceOrdered[j]]
		})
		for _, performance := range performanceOrdered {
			performanceIdent := "__stage_0" + performance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Performance{Name: %s}).Stage(stageSet.Stage)", performanceIdent, __gong__toRawStringLiteral(performance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", performanceIdent, __gong__toRawStringLiteral(performance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", performanceIdent, performance.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", performanceIdent, __gong__toRawStringLiteral(performance.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", performanceIdent, __gong__toRawStringLiteral(performance.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", performanceIdent, performance.IsExpanded))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexities = append(%s.Complexities, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Performances {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performances = append(%s.Performances, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Efforts {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Efforts = append(%s.Efforts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.SubSystems {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubSystems = append(%s.SubSystems, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramFlossEquations {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquations = append(%s.DiagramFlossEquations, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquationsWhoseNodeIsExpanded = append(%s.DiagramFlossEquationsWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ComplexitysWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.PerformancesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.EffortsWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/dsm/floss/go/models"
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
				case "CompareAnalysis":
					if !preserveOrder {
						inst := (&CompareAnalysis{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CompareAnalysis)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Complexity":
					if !preserveOrder {
						inst := (&Complexity{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Complexity)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DiagramFlossEquation":
					if !preserveOrder {
						inst := (&DiagramFlossEquation{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DiagramFlossEquation)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Effort":
					if !preserveOrder {
						inst := (&Effort{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Effort)
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
				case "NoteComplexityShape":
					if !preserveOrder {
						inst := (&NoteComplexityShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteComplexityShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NoteEffortShape":
					if !preserveOrder {
						inst := (&NoteEffortShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NoteEffortShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "NotePerformanceShape":
					if !preserveOrder {
						inst := (&NotePerformanceShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(NotePerformanceShape)
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
				case "Performance":
					if !preserveOrder {
						inst := (&Performance{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Performance)
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
									inst.FromSystem = typedTarget
								}
							}
						}
					case "ToSystem":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
									inst.ToSystem = typedTarget
								}
							}
						}
					case "Mu":
						inst.Mu = GongExtractFloat(rhs)
					case "Epsilon":
						inst.Epsilon = GongExtractFloat(rhs)
					case "DiagramFlossEquations":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramFlossEquation); ok {
										inst.DiagramFlossEquations = append(inst.DiagramFlossEquations, typedTarget)
									}
								}
							}
						}
					case "DiagramFlossEquationsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramFlossEquation); ok {
										inst.DiagramFlossEquationsWhoseNodeIsExpanded = append(inst.DiagramFlossEquationsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteShape); ok {
										inst.Note_Shapes = append(inst.Note_Shapes, typedTarget)
									}
								}
							}
						}
					case "NoteComplexityShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteComplexityShape); ok {
										inst.NoteComplexityShapes = append(inst.NoteComplexityShapes, typedTarget)
									}
								}
							}
						}
					case "NotePerformanceShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NotePerformanceShape); ok {
										inst.NotePerformanceShapes = append(inst.NotePerformanceShapes, typedTarget)
									}
								}
							}
						}
					case "NoteEffortShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*NoteEffortShape); ok {
										inst.NoteEffortShapes = append(inst.NoteEffortShapes, typedTarget)
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
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
										inst.ComplexitysWhoseNodeIsExpanded = append(inst.ComplexitysWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
										inst.PerformancesWhoseNodeIsExpanded = append(inst.PerformancesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
										inst.EffortsWhoseNodeIsExpanded = append(inst.EffortsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Library); ok {
										inst.SubLibraries = append(inst.SubLibraries, typedTarget)
									}
								}
							}
						}
					case "RootSystems":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.RootSystems = append(inst.RootSystems, typedTarget)
									}
								}
							}
						}
					case "RootComplexitys":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
										inst.RootComplexitys = append(inst.RootComplexitys, typedTarget)
									}
								}
							}
						}
					case "RootPerformances":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
										inst.RootPerformances = append(inst.RootPerformances, typedTarget)
									}
								}
							}
						}
					case "RootEfforts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
										inst.RootEfforts = append(inst.RootEfforts, typedTarget)
									}
								}
							}
						}
					case "RootCompareAnalysis":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CompareAnalysis); ok {
										inst.RootCompareAnalysis = append(inst.RootCompareAnalysis, typedTarget)
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
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
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
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
										inst.ComplexitysWhoseNodeIsExpanded = append(inst.ComplexitysWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
										inst.PerformancesWhoseNodeIsExpanded = append(inst.PerformancesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
										inst.EffortsWhoseNodeIsExpanded = append(inst.EffortsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsCompareAnalysisNodeExpanded":
						inst.IsCompareAnalysisNodeExpanded = GongExtractBool(rhs)
					case "CompareAnalysisWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CompareAnalysis); ok {
										inst.CompareAnalysisWhoseNodeIsExpanded = append(inst.CompareAnalysisWhoseNodeIsExpanded, typedTarget)
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
					case "Complexities":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
										inst.Complexities = append(inst.Complexities, typedTarget)
									}
								}
							}
						}
					case "Performances":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
										inst.Performances = append(inst.Performances, typedTarget)
									}
								}
							}
						}
					case "Efforts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
										inst.Efforts = append(inst.Efforts, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
									inst.Note = typedTarget
								}
							}
						}
					case "Complexity":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
									inst.Complexity = typedTarget
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
				case *NoteEffortShape:
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
					case "Effort":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
									inst.Effort = typedTarget
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
				case *NotePerformanceShape:
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
					case "Performance":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
									inst.Performance = typedTarget
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
										inst.Complexities = append(inst.Complexities, typedTarget)
									}
								}
							}
						}
					case "Performances":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
										inst.Performances = append(inst.Performances, typedTarget)
									}
								}
							}
						}
					case "Efforts":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
										inst.Efforts = append(inst.Efforts, typedTarget)
									}
								}
							}
						}
					case "SubSystems":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System); ok {
										inst.SubSystems = append(inst.SubSystems, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramFlossEquation); ok {
										inst.DiagramFlossEquations = append(inst.DiagramFlossEquations, typedTarget)
									}
								}
							}
						}
					case "DiagramFlossEquationsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*DiagramFlossEquation); ok {
										inst.DiagramFlossEquationsWhoseNodeIsExpanded = append(inst.DiagramFlossEquationsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsSubSystemNodeExpanded":
						inst.IsSubSystemNodeExpanded = GongExtractBool(rhs)
					case "IsComplexitysNodeExpanded":
						inst.IsComplexitysNodeExpanded = GongExtractBool(rhs)
					case "ComplexitysWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Complexity); ok {
										inst.ComplexitysWhoseNodeIsExpanded = append(inst.ComplexitysWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsPerformancesNodeExpanded":
						inst.IsPerformancesNodeExpanded = GongExtractBool(rhs)
					case "PerformancesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Performance); ok {
										inst.PerformancesWhoseNodeIsExpanded = append(inst.PerformancesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "IsEffortsNodeExpanded":
						inst.IsEffortsNodeExpanded = GongExtractBool(rhs)
					case "EffortsWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effort); ok {
										inst.EffortsWhoseNodeIsExpanded = append(inst.EffortsWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
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

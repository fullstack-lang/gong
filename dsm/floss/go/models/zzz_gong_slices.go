// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct CompareAnalysis
	// insertion point per field
	stage.CompareAnalysis_DiagramFlossEquations_reverseMap = make(map[*DiagramFlossEquation]*CompareAnalysis)
	for compareanalysis := range stage.CompareAnalysiss {
		_ = compareanalysis
		for _, _diagramflossequation := range compareanalysis.DiagramFlossEquations {
			stage.CompareAnalysis_DiagramFlossEquations_reverseMap[_diagramflossequation] = compareanalysis
		}
	}
	stage.CompareAnalysis_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap = make(map[*DiagramFlossEquation]*CompareAnalysis)
	for compareanalysis := range stage.CompareAnalysiss {
		_ = compareanalysis
		for _, _diagramflossequation := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
			stage.CompareAnalysis_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap[_diagramflossequation] = compareanalysis
		}
	}

	// Compute reverse map for named struct DiagramFlossEquation
	// insertion point per field
	stage.DiagramFlossEquation_Note_Shapes_reverseMap = make(map[*NoteShape]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _noteshape := range diagramflossequation.Note_Shapes {
			stage.DiagramFlossEquation_Note_Shapes_reverseMap[_noteshape] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_NoteComplexityShapes_reverseMap = make(map[*NoteComplexityShape]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _notecomplexityshape := range diagramflossequation.NoteComplexityShapes {
			stage.DiagramFlossEquation_NoteComplexityShapes_reverseMap[_notecomplexityshape] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_NotePerformanceShapes_reverseMap = make(map[*NotePerformanceShape]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _noteperformanceshape := range diagramflossequation.NotePerformanceShapes {
			stage.DiagramFlossEquation_NotePerformanceShapes_reverseMap[_noteperformanceshape] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_NoteEffortShapes_reverseMap = make(map[*NoteEffortShape]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _noteeffortshape := range diagramflossequation.NoteEffortShapes {
			stage.DiagramFlossEquation_NoteEffortShapes_reverseMap[_noteeffortshape] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _note := range diagramflossequation.NotesWhoseNodeIsExpanded {
			stage.DiagramFlossEquation_NotesWhoseNodeIsExpanded_reverseMap[_note] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_ComplexitysWhoseNodeIsExpanded_reverseMap = make(map[*Complexity]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _complexity := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
			stage.DiagramFlossEquation_ComplexitysWhoseNodeIsExpanded_reverseMap[_complexity] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_PerformancesWhoseNodeIsExpanded_reverseMap = make(map[*Performance]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _performance := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
			stage.DiagramFlossEquation_PerformancesWhoseNodeIsExpanded_reverseMap[_performance] = diagramflossequation
		}
	}
	stage.DiagramFlossEquation_EffortsWhoseNodeIsExpanded_reverseMap = make(map[*Effort]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		_ = diagramflossequation
		for _, _effort := range diagramflossequation.EffortsWhoseNodeIsExpanded {
			stage.DiagramFlossEquation_EffortsWhoseNodeIsExpanded_reverseMap[_effort] = diagramflossequation
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_RootSystems_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.RootSystems {
			stage.Library_RootSystems_reverseMap[_system] = library
		}
	}
	stage.Library_RootComplexitys_reverseMap = make(map[*Complexity]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _complexity := range library.RootComplexitys {
			stage.Library_RootComplexitys_reverseMap[_complexity] = library
		}
	}
	stage.Library_RootPerformances_reverseMap = make(map[*Performance]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _performance := range library.RootPerformances {
			stage.Library_RootPerformances_reverseMap[_performance] = library
		}
	}
	stage.Library_RootEfforts_reverseMap = make(map[*Effort]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _effort := range library.RootEfforts {
			stage.Library_RootEfforts_reverseMap[_effort] = library
		}
	}
	stage.Library_RootCompareAnalysis_reverseMap = make(map[*CompareAnalysis]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _compareanalysis := range library.RootCompareAnalysis {
			stage.Library_RootCompareAnalysis_reverseMap[_compareanalysis] = library
		}
	}
	stage.Library_RootNotes_reverseMap = make(map[*Note]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _note := range library.RootNotes {
			stage.Library_RootNotes_reverseMap[_note] = library
		}
	}
	stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[_library] = library
		}
	}
	stage.Library_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[_system] = library
		}
	}
	stage.Library_ComplexitysWhoseNodeIsExpanded_reverseMap = make(map[*Complexity]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _complexity := range library.ComplexitysWhoseNodeIsExpanded {
			stage.Library_ComplexitysWhoseNodeIsExpanded_reverseMap[_complexity] = library
		}
	}
	stage.Library_PerformancesWhoseNodeIsExpanded_reverseMap = make(map[*Performance]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _performance := range library.PerformancesWhoseNodeIsExpanded {
			stage.Library_PerformancesWhoseNodeIsExpanded_reverseMap[_performance] = library
		}
	}
	stage.Library_EffortsWhoseNodeIsExpanded_reverseMap = make(map[*Effort]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _effort := range library.EffortsWhoseNodeIsExpanded {
			stage.Library_EffortsWhoseNodeIsExpanded_reverseMap[_effort] = library
		}
	}
	stage.Library_CompareAnalysisWhoseNodeIsExpanded_reverseMap = make(map[*CompareAnalysis]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _compareanalysis := range library.CompareAnalysisWhoseNodeIsExpanded {
			stage.Library_CompareAnalysisWhoseNodeIsExpanded_reverseMap[_compareanalysis] = library
		}
	}
	stage.Library_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _note := range library.NotesWhoseNodeIsExpanded {
			stage.Library_NotesWhoseNodeIsExpanded_reverseMap[_note] = library
		}
	}

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Complexities_reverseMap = make(map[*Complexity]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _complexity := range note.Complexities {
			stage.Note_Complexities_reverseMap[_complexity] = note
		}
	}
	stage.Note_Performances_reverseMap = make(map[*Performance]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _performance := range note.Performances {
			stage.Note_Performances_reverseMap[_performance] = note
		}
	}
	stage.Note_Efforts_reverseMap = make(map[*Effort]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _effort := range note.Efforts {
			stage.Note_Efforts_reverseMap[_effort] = note
		}
	}

	// Compute reverse map for named struct System
	// insertion point per field
	stage.System_Complexities_reverseMap = make(map[*Complexity]*System)
	for system := range stage.Systems {
		_ = system
		for _, _complexity := range system.Complexities {
			stage.System_Complexities_reverseMap[_complexity] = system
		}
	}
	stage.System_Performances_reverseMap = make(map[*Performance]*System)
	for system := range stage.Systems {
		_ = system
		for _, _performance := range system.Performances {
			stage.System_Performances_reverseMap[_performance] = system
		}
	}
	stage.System_Efforts_reverseMap = make(map[*Effort]*System)
	for system := range stage.Systems {
		_ = system
		for _, _effort := range system.Efforts {
			stage.System_Efforts_reverseMap[_effort] = system
		}
	}
	stage.System_SubSystems_reverseMap = make(map[*System]*System)
	for system := range stage.Systems {
		_ = system
		for _, _system := range system.SubSystems {
			stage.System_SubSystems_reverseMap[_system] = system
		}
	}
	stage.System_DiagramFlossEquations_reverseMap = make(map[*DiagramFlossEquation]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramflossequation := range system.DiagramFlossEquations {
			stage.System_DiagramFlossEquations_reverseMap[_diagramflossequation] = system
		}
	}
	stage.System_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap = make(map[*DiagramFlossEquation]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramflossequation := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
			stage.System_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap[_diagramflossequation] = system
		}
	}
	stage.System_ComplexitysWhoseNodeIsExpanded_reverseMap = make(map[*Complexity]*System)
	for system := range stage.Systems {
		_ = system
		for _, _complexity := range system.ComplexitysWhoseNodeIsExpanded {
			stage.System_ComplexitysWhoseNodeIsExpanded_reverseMap[_complexity] = system
		}
	}
	stage.System_PerformancesWhoseNodeIsExpanded_reverseMap = make(map[*Performance]*System)
	for system := range stage.Systems {
		_ = system
		for _, _performance := range system.PerformancesWhoseNodeIsExpanded {
			stage.System_PerformancesWhoseNodeIsExpanded_reverseMap[_performance] = system
		}
	}
	stage.System_EffortsWhoseNodeIsExpanded_reverseMap = make(map[*Effort]*System)
	for system := range stage.Systems {
		_ = system
		for _, _effort := range system.EffortsWhoseNodeIsExpanded {
			stage.System_EffortsWhoseNodeIsExpanded_reverseMap[_effort] = system
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.CompareAnalysiss)

	res = __gong__appendInstances(res, stage.Complexitys)

	res = __gong__appendInstances(res, stage.DiagramFlossEquations)

	res = __gong__appendInstances(res, stage.Efforts)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NoteComplexityShapes)

	res = __gong__appendInstances(res, stage.NoteEffortShapes)

	res = __gong__appendInstances(res, stage.NotePerformanceShapes)

	res = __gong__appendInstances(res, stage.NoteShapes)

	res = __gong__appendInstances(res, stage.Performances)

	res = __gong__appendInstances(res, stage.Systems)

	return
}

// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongCopy() GongstructIF {
	newInstance := new(CompareAnalysis)
	compareanalysis.GongCopyBasicFields(newInstance)
	return newInstance
}

func (complexity *Complexity) GongCopy() GongstructIF {
	newInstance := new(Complexity)
	complexity.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagramflossequation *DiagramFlossEquation) GongCopy() GongstructIF {
	newInstance := new(DiagramFlossEquation)
	diagramflossequation.GongCopyBasicFields(newInstance)
	return newInstance
}

func (effort *Effort) GongCopy() GongstructIF {
	newInstance := new(Effort)
	effort.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notecomplexityshape *NoteComplexityShape) GongCopy() GongstructIF {
	newInstance := new(NoteComplexityShape)
	notecomplexityshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteeffortshape *NoteEffortShape) GongCopy() GongstructIF {
	newInstance := new(NoteEffortShape)
	noteeffortshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteperformanceshape *NotePerformanceShape) GongCopy() GongstructIF {
	newInstance := new(NotePerformanceShape)
	noteperformanceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (performance *Performance) GongCopy() GongstructIF {
	newInstance := new(Performance)
	performance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (system *System) GongCopy() GongstructIF {
	newInstance := new(System)
	system.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, compareanalysis)
}

func (complexity *Complexity) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, complexity)
}

func (diagramflossequation *DiagramFlossEquation) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagramflossequation)
}

func (effort *Effort) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, effort)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (notecomplexityshape *NoteComplexityShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notecomplexityshape)
}

func (noteeffortshape *NoteEffortShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteeffortshape)
}

func (noteperformanceshape *NotePerformanceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteperformanceshape)
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteshape)
}

func (performance *Performance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, performance)
}

func (system *System) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, system)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.CompareAnalysiss,
		stage.CompareAnalysis_stagedOrder,
		stage.CompareAnalysiss_reference,
		&stage.CompareAnalysiss_referenceOrder,
		stage.CompareAnalysiss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Complexitys,
		stage.Complexity_stagedOrder,
		stage.Complexitys_reference,
		&stage.Complexitys_referenceOrder,
		stage.Complexitys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DiagramFlossEquations,
		stage.DiagramFlossEquation_stagedOrder,
		stage.DiagramFlossEquations_reference,
		&stage.DiagramFlossEquations_referenceOrder,
		stage.DiagramFlossEquations_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Efforts,
		stage.Effort_stagedOrder,
		stage.Efforts_reference,
		&stage.Efforts_referenceOrder,
		stage.Efforts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notes,
		stage.Note_stagedOrder,
		stage.Notes_reference,
		&stage.Notes_referenceOrder,
		stage.Notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteComplexityShapes,
		stage.NoteComplexityShape_stagedOrder,
		stage.NoteComplexityShapes_reference,
		&stage.NoteComplexityShapes_referenceOrder,
		stage.NoteComplexityShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteEffortShapes,
		stage.NoteEffortShape_stagedOrder,
		stage.NoteEffortShapes_reference,
		&stage.NoteEffortShapes_referenceOrder,
		stage.NoteEffortShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NotePerformanceShapes,
		stage.NotePerformanceShape_stagedOrder,
		stage.NotePerformanceShapes_reference,
		&stage.NotePerformanceShapes_referenceOrder,
		stage.NotePerformanceShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteShapes,
		stage.NoteShape_stagedOrder,
		stage.NoteShapes_reference,
		&stage.NoteShapes_referenceOrder,
		stage.NoteShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Performances,
		stage.Performance_stagedOrder,
		stage.Performances_reference,
		&stage.Performances_referenceOrder,
		stage.Performances_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Systems,
		stage.System_stagedOrder,
		stage.Systems_reference,
		&stage.Systems_referenceOrder,
		stage.Systems_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.CompareAnalysiss, &stage.CompareAnalysiss_reference, &stage.CompareAnalysiss_referenceOrder, &stage.CompareAnalysiss_instance)

	__gong__computeReferencePass1(stage, stage.Complexitys, &stage.Complexitys_reference, &stage.Complexitys_referenceOrder, &stage.Complexitys_instance)

	__gong__computeReferencePass1(stage, stage.DiagramFlossEquations, &stage.DiagramFlossEquations_reference, &stage.DiagramFlossEquations_referenceOrder, &stage.DiagramFlossEquations_instance)

	__gong__computeReferencePass1(stage, stage.Efforts, &stage.Efforts_reference, &stage.Efforts_referenceOrder, &stage.Efforts_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NoteComplexityShapes, &stage.NoteComplexityShapes_reference, &stage.NoteComplexityShapes_referenceOrder, &stage.NoteComplexityShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteEffortShapes, &stage.NoteEffortShapes_reference, &stage.NoteEffortShapes_referenceOrder, &stage.NoteEffortShapes_instance)

	__gong__computeReferencePass1(stage, stage.NotePerformanceShapes, &stage.NotePerformanceShapes_reference, &stage.NotePerformanceShapes_referenceOrder, &stage.NotePerformanceShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteShapes, &stage.NoteShapes_reference, &stage.NoteShapes_referenceOrder, &stage.NoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.Performances, &stage.Performances_reference, &stage.Performances_referenceOrder, &stage.Performances_instance)

	__gong__computeReferencePass1(stage, stage.Systems, &stage.Systems_reference, &stage.Systems_referenceOrder, &stage.Systems_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CompareAnalysiss, stage.CompareAnalysiss_reference, stage)

	__gong__computeReferencePass2(stage.Complexitys, stage.Complexitys_reference, stage)

	__gong__computeReferencePass2(stage.DiagramFlossEquations, stage.DiagramFlossEquations_reference, stage)

	__gong__computeReferencePass2(stage.Efforts, stage.Efforts_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NoteComplexityShapes, stage.NoteComplexityShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteEffortShapes, stage.NoteEffortShapes_reference, stage)

	__gong__computeReferencePass2(stage.NotePerformanceShapes, stage.NotePerformanceShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteShapes, stage.NoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.Performances, stage.Performances_reference, stage)

	__gong__computeReferencePass2(stage.Systems, stage.Systems_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CompareAnalysis_stagedOrder, stage.CompareAnalysiss_referenceOrder, compareanalysis, "CompareAnalysis")
}

func (complexity *Complexity) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Complexity_stagedOrder, stage.Complexitys_referenceOrder, complexity, "Complexity")
}

func (diagramflossequation *DiagramFlossEquation) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DiagramFlossEquation_stagedOrder, stage.DiagramFlossEquations_referenceOrder, diagramflossequation, "DiagramFlossEquation")
}

func (effort *Effort) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Effort_stagedOrder, stage.Efforts_referenceOrder, effort, "Effort")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (notecomplexityshape *NoteComplexityShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteComplexityShape_stagedOrder, stage.NoteComplexityShapes_referenceOrder, notecomplexityshape, "NoteComplexityShape")
}

func (noteeffortshape *NoteEffortShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteEffortShape_stagedOrder, stage.NoteEffortShapes_referenceOrder, noteeffortshape, "NoteEffortShape")
}

func (noteperformanceshape *NotePerformanceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePerformanceShape_stagedOrder, stage.NotePerformanceShapes_referenceOrder, noteperformanceshape, "NotePerformanceShape")
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteShape_stagedOrder, stage.NoteShapes_referenceOrder, noteshape, "NoteShape")
}

func (performance *Performance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Performance_stagedOrder, stage.Performances_referenceOrder, performance, "Performance")
}

func (system *System) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.System_stagedOrder, stage.Systems_referenceOrder, system, "System")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(compareanalysis, compareanalysis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (compareanalysis *CompareAnalysis) GongGetReferenceIdentifier(stage *Stage) string {
	return compareanalysis.GongGetIdentifier(stage)
}

func (complexity *Complexity) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(complexity, complexity.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexity *Complexity) GongGetReferenceIdentifier(stage *Stage) string {
	return complexity.GongGetIdentifier(stage)
}

func (diagramflossequation *DiagramFlossEquation) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagramflossequation, diagramflossequation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramflossequation *DiagramFlossEquation) GongGetReferenceIdentifier(stage *Stage) string {
	return diagramflossequation.GongGetIdentifier(stage)
}

func (effort *Effort) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(effort, effort.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effort *Effort) GongGetReferenceIdentifier(stage *Stage) string {
	return effort.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note, note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return note.GongGetIdentifier(stage)
}

func (notecomplexityshape *NoteComplexityShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notecomplexityshape, notecomplexityshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notecomplexityshape *NoteComplexityShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notecomplexityshape.GongGetIdentifier(stage)
}

func (noteeffortshape *NoteEffortShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteeffortshape, noteeffortshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteeffortshape *NoteEffortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteeffortshape.GongGetIdentifier(stage)
}

func (noteperformanceshape *NotePerformanceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteperformanceshape, noteperformanceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteperformanceshape *NotePerformanceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteperformanceshape.GongGetIdentifier(stage)
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteshape, noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteshape.GongGetIdentifier(stage)
}

func (performance *Performance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(performance, performance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (performance *Performance) GongGetReferenceIdentifier(stage *Stage) string {
	return performance.GongGetIdentifier(stage)
}

func (system *System) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(system, system.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system *System) GongGetReferenceIdentifier(stage *Stage) string {
	return system.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(compareanalysis.GongGetIdentifier(stage), "CompareAnalysis", compareanalysis.Name)
}

func (complexity *Complexity) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(complexity.GongGetIdentifier(stage), "Complexity", complexity.Name)
}

func (diagramflossequation *DiagramFlossEquation) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagramflossequation.GongGetIdentifier(stage), "DiagramFlossEquation", diagramflossequation.Name)
}

func (effort *Effort) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(effort.GongGetIdentifier(stage), "Effort", effort.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (notecomplexityshape *NoteComplexityShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notecomplexityshape.GongGetIdentifier(stage), "NoteComplexityShape", notecomplexityshape.Name)
}

func (noteeffortshape *NoteEffortShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteeffortshape.GongGetIdentifier(stage), "NoteEffortShape", noteeffortshape.Name)
}

func (noteperformanceshape *NotePerformanceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteperformanceshape.GongGetIdentifier(stage), "NotePerformanceShape", noteperformanceshape.Name)
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteshape.GongGetIdentifier(stage), "NoteShape", noteshape.Name)
}

func (performance *Performance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(performance.GongGetIdentifier(stage), "Performance", performance.Name)
}

func (system *System) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(system.GongGetIdentifier(stage), "System", system.Name)
}

// insertion point for unstaging
func (compareanalysis *CompareAnalysis) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(compareanalysis.GongGetReferenceIdentifier(stage))
}

func (complexity *Complexity) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(complexity.GongGetReferenceIdentifier(stage))
}

func (diagramflossequation *DiagramFlossEquation) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagramflossequation.GongGetReferenceIdentifier(stage))
}

func (effort *Effort) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(effort.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (notecomplexityshape *NoteComplexityShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notecomplexityshape.GongGetReferenceIdentifier(stage))
}

func (noteeffortshape *NoteEffortShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteeffortshape.GongGetReferenceIdentifier(stage))
}

func (noteperformanceshape *NotePerformanceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteperformanceshape.GongGetReferenceIdentifier(stage))
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteshape.GongGetReferenceIdentifier(stage))
}

func (performance *Performance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(performance.GongGetReferenceIdentifier(stage))
}

func (system *System) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(system.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template

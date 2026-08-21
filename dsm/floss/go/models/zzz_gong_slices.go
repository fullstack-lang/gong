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

	// Compute reverse map for named struct Complexity
	// insertion point per field

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

	// Compute reverse map for named struct Effort
	// insertion point per field

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

	// Compute reverse map for named struct NoteComplexityShape
	// insertion point per field

	// Compute reverse map for named struct NoteEffortShape
	// insertion point per field

	// Compute reverse map for named struct NotePerformanceShape
	// insertion point per field

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct Performance
	// insertion point per field

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
	stage.System_SubSystemes_reverseMap = make(map[*System]*System)
	for system := range stage.Systems {
		_ = system
		for _, _system := range system.SubSystemes {
			stage.System_SubSystemes_reverseMap[_system] = system
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
	for instance := range stage.CompareAnalysiss {
		res = append(res, instance)
	}

	for instance := range stage.Complexitys {
		res = append(res, instance)
	}

	for instance := range stage.DiagramFlossEquations {
		res = append(res, instance)
	}

	for instance := range stage.Efforts {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NoteComplexityShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteEffortShapes {
		res = append(res, instance)
	}

	for instance := range stage.NotePerformanceShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.Performances {
		res = append(res, instance)
	}

	for instance := range stage.Systems {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongCopy() GongstructIF {
	newInstance := new(CompareAnalysis)
	compareanalysis.CopyBasicFields(newInstance)
	return newInstance
}

func (complexity *Complexity) GongCopy() GongstructIF {
	newInstance := new(Complexity)
	complexity.CopyBasicFields(newInstance)
	return newInstance
}

func (diagramflossequation *DiagramFlossEquation) GongCopy() GongstructIF {
	newInstance := new(DiagramFlossEquation)
	diagramflossequation.CopyBasicFields(newInstance)
	return newInstance
}

func (effort *Effort) GongCopy() GongstructIF {
	newInstance := new(Effort)
	effort.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.CopyBasicFields(newInstance)
	return newInstance
}

func (notecomplexityshape *NoteComplexityShape) GongCopy() GongstructIF {
	newInstance := new(NoteComplexityShape)
	notecomplexityshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteeffortshape *NoteEffortShape) GongCopy() GongstructIF {
	newInstance := new(NoteEffortShape)
	noteeffortshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteperformanceshape *NotePerformanceShape) GongCopy() GongstructIF {
	newInstance := new(NotePerformanceShape)
	noteperformanceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.CopyBasicFields(newInstance)
	return newInstance
}

func (performance *Performance) GongCopy() GongstructIF {
	newInstance := new(Performance)
	performance.CopyBasicFields(newInstance)
	return newInstance
}

func (system *System) GongCopy() GongstructIF {
	newInstance := new(System)
	system.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(compareanalysis).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(compareanalysis), uint64(GetOrderPointerGongstruct(stage, compareanalysis)))
	return
}

func (complexity *Complexity) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complexity).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(complexity), uint64(GetOrderPointerGongstruct(stage, complexity)))
	return
}

func (diagramflossequation *DiagramFlossEquation) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramflossequation).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramflossequation), uint64(GetOrderPointerGongstruct(stage, diagramflossequation)))
	return
}

func (effort *Effort) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(effort).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(effort), uint64(GetOrderPointerGongstruct(stage, effort)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(note), uint64(GetOrderPointerGongstruct(stage, note)))
	return
}

func (notecomplexityshape *NoteComplexityShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notecomplexityshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(notecomplexityshape), uint64(GetOrderPointerGongstruct(stage, notecomplexityshape)))
	return
}

func (noteeffortshape *NoteEffortShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteeffortshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteeffortshape), uint64(GetOrderPointerGongstruct(stage, noteeffortshape)))
	return
}

func (noteperformanceshape *NotePerformanceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteperformanceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteperformanceshape), uint64(GetOrderPointerGongstruct(stage, noteperformanceshape)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteshape), uint64(GetOrderPointerGongstruct(stage, noteshape)))
	return
}

func (performance *Performance) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(performance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(performance), uint64(GetOrderPointerGongstruct(stage, performance)))
	return
}

func (system *System) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(system), uint64(GetOrderPointerGongstruct(stage, system)))
	return
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
	var compareanalysiss_newInstances []*CompareAnalysis
	var compareanalysiss_deletedInstances []*CompareAnalysis

	// parse all staged instances and check if they have a reference
	for compareanalysis := range stage.CompareAnalysiss {
		if ref, ok := stage.CompareAnalysiss_reference[compareanalysis]; !ok {
			compareanalysiss_newInstances = append(compareanalysiss_newInstances, compareanalysis)
			newInstancesSlice = append(newInstancesSlice, compareanalysis.GongMarshallIdentifier(stage))
			if stage.CompareAnalysiss_referenceOrder == nil {
				stage.CompareAnalysiss_referenceOrder = make(map[*CompareAnalysis]uint)
			}
			stage.CompareAnalysiss_referenceOrder[compareanalysis] = stage.CompareAnalysis_stagedOrder[compareanalysis]
			newInstancesReverseSlice = append(newInstancesReverseSlice, compareanalysis.GongMarshallUnstaging(stage))
			// delete(stage.CompareAnalysiss_referenceOrder, compareanalysis)
			fieldInitializers, pointersInitializations := compareanalysis.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CompareAnalysis_stagedOrder[ref] = stage.CompareAnalysis_stagedOrder[compareanalysis]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := compareanalysis.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, compareanalysis)
			// delete(stage.CompareAnalysis_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if compareanalysis.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", compareanalysis.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.CompareAnalysiss_reference {
		instance := stage.CompareAnalysiss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CompareAnalysiss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			compareanalysiss_deletedInstances = append(compareanalysiss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(compareanalysiss_newInstances)
	lenDeletedInstances += len(compareanalysiss_deletedInstances)
	var complexitys_newInstances []*Complexity
	var complexitys_deletedInstances []*Complexity

	// parse all staged instances and check if they have a reference
	for complexity := range stage.Complexitys {
		if ref, ok := stage.Complexitys_reference[complexity]; !ok {
			complexitys_newInstances = append(complexitys_newInstances, complexity)
			newInstancesSlice = append(newInstancesSlice, complexity.GongMarshallIdentifier(stage))
			if stage.Complexitys_referenceOrder == nil {
				stage.Complexitys_referenceOrder = make(map[*Complexity]uint)
			}
			stage.Complexitys_referenceOrder[complexity] = stage.Complexity_stagedOrder[complexity]
			newInstancesReverseSlice = append(newInstancesReverseSlice, complexity.GongMarshallUnstaging(stage))
			// delete(stage.Complexitys_referenceOrder, complexity)
			fieldInitializers, pointersInitializations := complexity.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Complexity_stagedOrder[ref] = stage.Complexity_stagedOrder[complexity]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := complexity.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, complexity)
			// delete(stage.Complexity_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if complexity.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", complexity.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Complexitys_reference {
		instance := stage.Complexitys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Complexitys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			complexitys_deletedInstances = append(complexitys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(complexitys_newInstances)
	lenDeletedInstances += len(complexitys_deletedInstances)
	var diagramflossequations_newInstances []*DiagramFlossEquation
	var diagramflossequations_deletedInstances []*DiagramFlossEquation

	// parse all staged instances and check if they have a reference
	for diagramflossequation := range stage.DiagramFlossEquations {
		if ref, ok := stage.DiagramFlossEquations_reference[diagramflossequation]; !ok {
			diagramflossequations_newInstances = append(diagramflossequations_newInstances, diagramflossequation)
			newInstancesSlice = append(newInstancesSlice, diagramflossequation.GongMarshallIdentifier(stage))
			if stage.DiagramFlossEquations_referenceOrder == nil {
				stage.DiagramFlossEquations_referenceOrder = make(map[*DiagramFlossEquation]uint)
			}
			stage.DiagramFlossEquations_referenceOrder[diagramflossequation] = stage.DiagramFlossEquation_stagedOrder[diagramflossequation]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagramflossequation.GongMarshallUnstaging(stage))
			// delete(stage.DiagramFlossEquations_referenceOrder, diagramflossequation)
			fieldInitializers, pointersInitializations := diagramflossequation.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramFlossEquation_stagedOrder[ref] = stage.DiagramFlossEquation_stagedOrder[diagramflossequation]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagramflossequation.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagramflossequation)
			// delete(stage.DiagramFlossEquation_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagramflossequation.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagramflossequation.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DiagramFlossEquations_reference {
		instance := stage.DiagramFlossEquations_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DiagramFlossEquations[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagramflossequations_deletedInstances = append(diagramflossequations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagramflossequations_newInstances)
	lenDeletedInstances += len(diagramflossequations_deletedInstances)
	var efforts_newInstances []*Effort
	var efforts_deletedInstances []*Effort

	// parse all staged instances and check if they have a reference
	for effort := range stage.Efforts {
		if ref, ok := stage.Efforts_reference[effort]; !ok {
			efforts_newInstances = append(efforts_newInstances, effort)
			newInstancesSlice = append(newInstancesSlice, effort.GongMarshallIdentifier(stage))
			if stage.Efforts_referenceOrder == nil {
				stage.Efforts_referenceOrder = make(map[*Effort]uint)
			}
			stage.Efforts_referenceOrder[effort] = stage.Effort_stagedOrder[effort]
			newInstancesReverseSlice = append(newInstancesReverseSlice, effort.GongMarshallUnstaging(stage))
			// delete(stage.Efforts_referenceOrder, effort)
			fieldInitializers, pointersInitializations := effort.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Effort_stagedOrder[ref] = stage.Effort_stagedOrder[effort]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := effort.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, effort)
			// delete(stage.Effort_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if effort.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", effort.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Efforts_reference {
		instance := stage.Efforts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Efforts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			efforts_deletedInstances = append(efforts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(efforts_newInstances)
	lenDeletedInstances += len(efforts_deletedInstances)
	var librarys_newInstances []*Library
	var librarys_deletedInstances []*Library

	// parse all staged instances and check if they have a reference
	for library := range stage.Librarys {
		if ref, ok := stage.Librarys_reference[library]; !ok {
			librarys_newInstances = append(librarys_newInstances, library)
			newInstancesSlice = append(newInstancesSlice, library.GongMarshallIdentifier(stage))
			if stage.Librarys_referenceOrder == nil {
				stage.Librarys_referenceOrder = make(map[*Library]uint)
			}
			stage.Librarys_referenceOrder[library] = stage.Library_stagedOrder[library]
			newInstancesReverseSlice = append(newInstancesReverseSlice, library.GongMarshallUnstaging(stage))
			// delete(stage.Librarys_referenceOrder, library)
			fieldInitializers, pointersInitializations := library.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Library_stagedOrder[ref] = stage.Library_stagedOrder[library]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := library.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, library)
			// delete(stage.Library_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if library.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Librarys_reference {
		instance := stage.Librarys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Librarys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			librarys_deletedInstances = append(librarys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(librarys_newInstances)
	lenDeletedInstances += len(librarys_deletedInstances)
	var notes_newInstances []*Note
	var notes_deletedInstances []*Note

	// parse all staged instances and check if they have a reference
	for note := range stage.Notes {
		if ref, ok := stage.Notes_reference[note]; !ok {
			notes_newInstances = append(notes_newInstances, note)
			newInstancesSlice = append(newInstancesSlice, note.GongMarshallIdentifier(stage))
			if stage.Notes_referenceOrder == nil {
				stage.Notes_referenceOrder = make(map[*Note]uint)
			}
			stage.Notes_referenceOrder[note] = stage.Note_stagedOrder[note]
			newInstancesReverseSlice = append(newInstancesReverseSlice, note.GongMarshallUnstaging(stage))
			// delete(stage.Notes_referenceOrder, note)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Note_stagedOrder[ref] = stage.Note_stagedOrder[note]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := note.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, note)
			// delete(stage.Note_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if note.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", note.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Notes_reference {
		instance := stage.Notes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Notes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notes_deletedInstances = append(notes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notes_newInstances)
	lenDeletedInstances += len(notes_deletedInstances)
	var notecomplexityshapes_newInstances []*NoteComplexityShape
	var notecomplexityshapes_deletedInstances []*NoteComplexityShape

	// parse all staged instances and check if they have a reference
	for notecomplexityshape := range stage.NoteComplexityShapes {
		if ref, ok := stage.NoteComplexityShapes_reference[notecomplexityshape]; !ok {
			notecomplexityshapes_newInstances = append(notecomplexityshapes_newInstances, notecomplexityshape)
			newInstancesSlice = append(newInstancesSlice, notecomplexityshape.GongMarshallIdentifier(stage))
			if stage.NoteComplexityShapes_referenceOrder == nil {
				stage.NoteComplexityShapes_referenceOrder = make(map[*NoteComplexityShape]uint)
			}
			stage.NoteComplexityShapes_referenceOrder[notecomplexityshape] = stage.NoteComplexityShape_stagedOrder[notecomplexityshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, notecomplexityshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteComplexityShapes_referenceOrder, notecomplexityshape)
			fieldInitializers, pointersInitializations := notecomplexityshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteComplexityShape_stagedOrder[ref] = stage.NoteComplexityShape_stagedOrder[notecomplexityshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := notecomplexityshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, notecomplexityshape)
			// delete(stage.NoteComplexityShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if notecomplexityshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", notecomplexityshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteComplexityShapes_reference {
		instance := stage.NoteComplexityShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteComplexityShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notecomplexityshapes_deletedInstances = append(notecomplexityshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notecomplexityshapes_newInstances)
	lenDeletedInstances += len(notecomplexityshapes_deletedInstances)
	var noteeffortshapes_newInstances []*NoteEffortShape
	var noteeffortshapes_deletedInstances []*NoteEffortShape

	// parse all staged instances and check if they have a reference
	for noteeffortshape := range stage.NoteEffortShapes {
		if ref, ok := stage.NoteEffortShapes_reference[noteeffortshape]; !ok {
			noteeffortshapes_newInstances = append(noteeffortshapes_newInstances, noteeffortshape)
			newInstancesSlice = append(newInstancesSlice, noteeffortshape.GongMarshallIdentifier(stage))
			if stage.NoteEffortShapes_referenceOrder == nil {
				stage.NoteEffortShapes_referenceOrder = make(map[*NoteEffortShape]uint)
			}
			stage.NoteEffortShapes_referenceOrder[noteeffortshape] = stage.NoteEffortShape_stagedOrder[noteeffortshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteeffortshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteEffortShapes_referenceOrder, noteeffortshape)
			fieldInitializers, pointersInitializations := noteeffortshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteEffortShape_stagedOrder[ref] = stage.NoteEffortShape_stagedOrder[noteeffortshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteeffortshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteeffortshape)
			// delete(stage.NoteEffortShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if noteeffortshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", noteeffortshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteEffortShapes_reference {
		instance := stage.NoteEffortShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteEffortShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteeffortshapes_deletedInstances = append(noteeffortshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteeffortshapes_newInstances)
	lenDeletedInstances += len(noteeffortshapes_deletedInstances)
	var noteperformanceshapes_newInstances []*NotePerformanceShape
	var noteperformanceshapes_deletedInstances []*NotePerformanceShape

	// parse all staged instances and check if they have a reference
	for noteperformanceshape := range stage.NotePerformanceShapes {
		if ref, ok := stage.NotePerformanceShapes_reference[noteperformanceshape]; !ok {
			noteperformanceshapes_newInstances = append(noteperformanceshapes_newInstances, noteperformanceshape)
			newInstancesSlice = append(newInstancesSlice, noteperformanceshape.GongMarshallIdentifier(stage))
			if stage.NotePerformanceShapes_referenceOrder == nil {
				stage.NotePerformanceShapes_referenceOrder = make(map[*NotePerformanceShape]uint)
			}
			stage.NotePerformanceShapes_referenceOrder[noteperformanceshape] = stage.NotePerformanceShape_stagedOrder[noteperformanceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteperformanceshape.GongMarshallUnstaging(stage))
			// delete(stage.NotePerformanceShapes_referenceOrder, noteperformanceshape)
			fieldInitializers, pointersInitializations := noteperformanceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NotePerformanceShape_stagedOrder[ref] = stage.NotePerformanceShape_stagedOrder[noteperformanceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteperformanceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteperformanceshape)
			// delete(stage.NotePerformanceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if noteperformanceshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", noteperformanceshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NotePerformanceShapes_reference {
		instance := stage.NotePerformanceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NotePerformanceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteperformanceshapes_deletedInstances = append(noteperformanceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteperformanceshapes_newInstances)
	lenDeletedInstances += len(noteperformanceshapes_deletedInstances)
	var noteshapes_newInstances []*NoteShape
	var noteshapes_deletedInstances []*NoteShape

	// parse all staged instances and check if they have a reference
	for noteshape := range stage.NoteShapes {
		if ref, ok := stage.NoteShapes_reference[noteshape]; !ok {
			noteshapes_newInstances = append(noteshapes_newInstances, noteshape)
			newInstancesSlice = append(newInstancesSlice, noteshape.GongMarshallIdentifier(stage))
			if stage.NoteShapes_referenceOrder == nil {
				stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)
			}
			stage.NoteShapes_referenceOrder[noteshape] = stage.NoteShape_stagedOrder[noteshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteShapes_referenceOrder, noteshape)
			fieldInitializers, pointersInitializations := noteshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteShape_stagedOrder[ref] = stage.NoteShape_stagedOrder[noteshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteshape)
			// delete(stage.NoteShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if noteshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", noteshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteShapes_reference {
		instance := stage.NoteShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteshapes_deletedInstances = append(noteshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteshapes_newInstances)
	lenDeletedInstances += len(noteshapes_deletedInstances)
	var performances_newInstances []*Performance
	var performances_deletedInstances []*Performance

	// parse all staged instances and check if they have a reference
	for performance := range stage.Performances {
		if ref, ok := stage.Performances_reference[performance]; !ok {
			performances_newInstances = append(performances_newInstances, performance)
			newInstancesSlice = append(newInstancesSlice, performance.GongMarshallIdentifier(stage))
			if stage.Performances_referenceOrder == nil {
				stage.Performances_referenceOrder = make(map[*Performance]uint)
			}
			stage.Performances_referenceOrder[performance] = stage.Performance_stagedOrder[performance]
			newInstancesReverseSlice = append(newInstancesReverseSlice, performance.GongMarshallUnstaging(stage))
			// delete(stage.Performances_referenceOrder, performance)
			fieldInitializers, pointersInitializations := performance.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Performance_stagedOrder[ref] = stage.Performance_stagedOrder[performance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := performance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, performance)
			// delete(stage.Performance_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if performance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", performance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Performances_reference {
		instance := stage.Performances_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Performances[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			performances_deletedInstances = append(performances_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(performances_newInstances)
	lenDeletedInstances += len(performances_deletedInstances)
	var systems_newInstances []*System
	var systems_deletedInstances []*System

	// parse all staged instances and check if they have a reference
	for system := range stage.Systems {
		if ref, ok := stage.Systems_reference[system]; !ok {
			systems_newInstances = append(systems_newInstances, system)
			newInstancesSlice = append(newInstancesSlice, system.GongMarshallIdentifier(stage))
			if stage.Systems_referenceOrder == nil {
				stage.Systems_referenceOrder = make(map[*System]uint)
			}
			stage.Systems_referenceOrder[system] = stage.System_stagedOrder[system]
			newInstancesReverseSlice = append(newInstancesReverseSlice, system.GongMarshallUnstaging(stage))
			// delete(stage.Systems_referenceOrder, system)
			fieldInitializers, pointersInitializations := system.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.System_stagedOrder[ref] = stage.System_stagedOrder[system]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := system.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, system)
			// delete(stage.System_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if system.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", system.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Systems_reference {
		instance := stage.Systems_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Systems[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			systems_deletedInstances = append(systems_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(systems_newInstances)
	lenDeletedInstances += len(systems_deletedInstances)

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
	stage.CompareAnalysiss_reference = make(map[*CompareAnalysis]*CompareAnalysis)
	stage.CompareAnalysiss_referenceOrder = make(map[*CompareAnalysis]uint) // diff Unstage needs the reference order
	stage.CompareAnalysiss_instance = make(map[*CompareAnalysis]*CompareAnalysis)
	for instance := range stage.CompareAnalysiss {
		_copy := instance.GongCopy().(*CompareAnalysis)
		stage.CompareAnalysiss_reference[instance] = _copy
		stage.CompareAnalysiss_instance[_copy] = instance
		stage.CompareAnalysiss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Complexitys_reference = make(map[*Complexity]*Complexity)
	stage.Complexitys_referenceOrder = make(map[*Complexity]uint) // diff Unstage needs the reference order
	stage.Complexitys_instance = make(map[*Complexity]*Complexity)
	for instance := range stage.Complexitys {
		_copy := instance.GongCopy().(*Complexity)
		stage.Complexitys_reference[instance] = _copy
		stage.Complexitys_instance[_copy] = instance
		stage.Complexitys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DiagramFlossEquations_reference = make(map[*DiagramFlossEquation]*DiagramFlossEquation)
	stage.DiagramFlossEquations_referenceOrder = make(map[*DiagramFlossEquation]uint) // diff Unstage needs the reference order
	stage.DiagramFlossEquations_instance = make(map[*DiagramFlossEquation]*DiagramFlossEquation)
	for instance := range stage.DiagramFlossEquations {
		_copy := instance.GongCopy().(*DiagramFlossEquation)
		stage.DiagramFlossEquations_reference[instance] = _copy
		stage.DiagramFlossEquations_instance[_copy] = instance
		stage.DiagramFlossEquations_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Efforts_reference = make(map[*Effort]*Effort)
	stage.Efforts_referenceOrder = make(map[*Effort]uint) // diff Unstage needs the reference order
	stage.Efforts_instance = make(map[*Effort]*Effort)
	for instance := range stage.Efforts {
		_copy := instance.GongCopy().(*Effort)
		stage.Efforts_reference[instance] = _copy
		stage.Efforts_instance[_copy] = instance
		stage.Efforts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	stage.Notes_instance = make(map[*Note]*Note)
	for instance := range stage.Notes {
		_copy := instance.GongCopy().(*Note)
		stage.Notes_reference[instance] = _copy
		stage.Notes_instance[_copy] = instance
		stage.Notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteComplexityShapes_reference = make(map[*NoteComplexityShape]*NoteComplexityShape)
	stage.NoteComplexityShapes_referenceOrder = make(map[*NoteComplexityShape]uint) // diff Unstage needs the reference order
	stage.NoteComplexityShapes_instance = make(map[*NoteComplexityShape]*NoteComplexityShape)
	for instance := range stage.NoteComplexityShapes {
		_copy := instance.GongCopy().(*NoteComplexityShape)
		stage.NoteComplexityShapes_reference[instance] = _copy
		stage.NoteComplexityShapes_instance[_copy] = instance
		stage.NoteComplexityShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteEffortShapes_reference = make(map[*NoteEffortShape]*NoteEffortShape)
	stage.NoteEffortShapes_referenceOrder = make(map[*NoteEffortShape]uint) // diff Unstage needs the reference order
	stage.NoteEffortShapes_instance = make(map[*NoteEffortShape]*NoteEffortShape)
	for instance := range stage.NoteEffortShapes {
		_copy := instance.GongCopy().(*NoteEffortShape)
		stage.NoteEffortShapes_reference[instance] = _copy
		stage.NoteEffortShapes_instance[_copy] = instance
		stage.NoteEffortShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NotePerformanceShapes_reference = make(map[*NotePerformanceShape]*NotePerformanceShape)
	stage.NotePerformanceShapes_referenceOrder = make(map[*NotePerformanceShape]uint) // diff Unstage needs the reference order
	stage.NotePerformanceShapes_instance = make(map[*NotePerformanceShape]*NotePerformanceShape)
	for instance := range stage.NotePerformanceShapes {
		_copy := instance.GongCopy().(*NotePerformanceShape)
		stage.NotePerformanceShapes_reference[instance] = _copy
		stage.NotePerformanceShapes_instance[_copy] = instance
		stage.NotePerformanceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint) // diff Unstage needs the reference order
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	for instance := range stage.NoteShapes {
		_copy := instance.GongCopy().(*NoteShape)
		stage.NoteShapes_reference[instance] = _copy
		stage.NoteShapes_instance[_copy] = instance
		stage.NoteShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Performances_reference = make(map[*Performance]*Performance)
	stage.Performances_referenceOrder = make(map[*Performance]uint) // diff Unstage needs the reference order
	stage.Performances_instance = make(map[*Performance]*Performance)
	for instance := range stage.Performances {
		_copy := instance.GongCopy().(*Performance)
		stage.Performances_reference[instance] = _copy
		stage.Performances_instance[_copy] = instance
		stage.Performances_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint) // diff Unstage needs the reference order
	stage.Systems_instance = make(map[*System]*System)
	for instance := range stage.Systems {
		_copy := instance.GongCopy().(*System)
		stage.Systems_reference[instance] = _copy
		stage.Systems_instance[_copy] = instance
		stage.Systems_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.CompareAnalysiss {
		reference := stage.CompareAnalysiss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Complexitys {
		reference := stage.Complexitys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DiagramFlossEquations {
		reference := stage.DiagramFlossEquations_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Efforts {
		reference := stage.Efforts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notes {
		reference := stage.Notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteComplexityShapes {
		reference := stage.NoteComplexityShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteEffortShapes {
		reference := stage.NoteEffortShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NotePerformanceShapes {
		reference := stage.NotePerformanceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteShapes {
		reference := stage.NoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Performances {
		reference := stage.Performances_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Systems {
		reference := stage.Systems_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

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
	if order, ok := stage.CompareAnalysis_stagedOrder[compareanalysis]; ok {
		return order
	}
	if order, ok := stage.CompareAnalysiss_referenceOrder[compareanalysis]; ok {
		return order
	} else {
		log.Printf("instance %p of type CompareAnalysis was not staged and does not have a reference order", compareanalysis)
		return 0
	}
}

func (complexity *Complexity) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Complexity_stagedOrder[complexity]; ok {
		return order
	}
	if order, ok := stage.Complexitys_referenceOrder[complexity]; ok {
		return order
	} else {
		log.Printf("instance %p of type Complexity was not staged and does not have a reference order", complexity)
		return 0
	}
}

func (diagramflossequation *DiagramFlossEquation) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramFlossEquation_stagedOrder[diagramflossequation]; ok {
		return order
	}
	if order, ok := stage.DiagramFlossEquations_referenceOrder[diagramflossequation]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramFlossEquation was not staged and does not have a reference order", diagramflossequation)
		return 0
	}
}

func (effort *Effort) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Effort_stagedOrder[effort]; ok {
		return order
	}
	if order, ok := stage.Efforts_referenceOrder[effort]; ok {
		return order
	} else {
		log.Printf("instance %p of type Effort was not staged and does not have a reference order", effort)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_stagedOrder[note]; ok {
		return order
	}
	if order, ok := stage.Notes_referenceOrder[note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note was not staged and does not have a reference order", note)
		return 0
	}
}

func (notecomplexityshape *NoteComplexityShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteComplexityShape_stagedOrder[notecomplexityshape]; ok {
		return order
	}
	if order, ok := stage.NoteComplexityShapes_referenceOrder[notecomplexityshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteComplexityShape was not staged and does not have a reference order", notecomplexityshape)
		return 0
	}
}

func (noteeffortshape *NoteEffortShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteEffortShape_stagedOrder[noteeffortshape]; ok {
		return order
	}
	if order, ok := stage.NoteEffortShapes_referenceOrder[noteeffortshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteEffortShape was not staged and does not have a reference order", noteeffortshape)
		return 0
	}
}

func (noteperformanceshape *NotePerformanceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NotePerformanceShape_stagedOrder[noteperformanceshape]; ok {
		return order
	}
	if order, ok := stage.NotePerformanceShapes_referenceOrder[noteperformanceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NotePerformanceShape was not staged and does not have a reference order", noteperformanceshape)
		return 0
	}
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteShape_stagedOrder[noteshape]; ok {
		return order
	}
	if order, ok := stage.NoteShapes_referenceOrder[noteshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteShape was not staged and does not have a reference order", noteshape)
		return 0
	}
}

func (performance *Performance) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Performance_stagedOrder[performance]; ok {
		return order
	}
	if order, ok := stage.Performances_referenceOrder[performance]; ok {
		return order
	} else {
		log.Printf("instance %p of type Performance was not staged and does not have a reference order", performance)
		return 0
	}
}

func (system *System) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.System_stagedOrder[system]; ok {
		return order
	}
	if order, ok := stage.Systems_referenceOrder[system]; ok {
		return order
	} else {
		log.Printf("instance %p of type System was not staged and does not have a reference order", system)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", compareanalysis.GongGetGongstructName(), compareanalysis.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (compareanalysis *CompareAnalysis) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", compareanalysis.GongGetGongstructName(), compareanalysis.GongGetOrder(stage))
}

func (complexity *Complexity) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexity.GongGetGongstructName(), complexity.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexity *Complexity) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexity.GongGetGongstructName(), complexity.GongGetOrder(stage))
}

func (diagramflossequation *DiagramFlossEquation) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramflossequation.GongGetGongstructName(), diagramflossequation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramflossequation *DiagramFlossEquation) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramflossequation.GongGetGongstructName(), diagramflossequation.GongGetOrder(stage))
}

func (effort *Effort) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effort.GongGetGongstructName(), effort.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effort *Effort) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effort.GongGetGongstructName(), effort.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (notecomplexityshape *NoteComplexityShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notecomplexityshape.GongGetGongstructName(), notecomplexityshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notecomplexityshape *NoteComplexityShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notecomplexityshape.GongGetGongstructName(), notecomplexityshape.GongGetOrder(stage))
}

func (noteeffortshape *NoteEffortShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteeffortshape.GongGetGongstructName(), noteeffortshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteeffortshape *NoteEffortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteeffortshape.GongGetGongstructName(), noteeffortshape.GongGetOrder(stage))
}

func (noteperformanceshape *NotePerformanceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteperformanceshape.GongGetGongstructName(), noteperformanceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteperformanceshape *NotePerformanceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteperformanceshape.GongGetGongstructName(), noteperformanceshape.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (performance *Performance) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", performance.GongGetGongstructName(), performance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (performance *Performance) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", performance.GongGetGongstructName(), performance.GongGetOrder(stage))
}

func (system *System) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system *System) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (compareanalysis *CompareAnalysis) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", compareanalysis.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CompareAnalysis")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(compareanalysis.Name))
	return
}

func (complexity *Complexity) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexity.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Complexity")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(complexity.Name))
	return
}

func (diagramflossequation *DiagramFlossEquation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramflossequation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramFlossEquation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramflossequation.Name))
	return
}

func (effort *Effort) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effort.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Effort")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(effort.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.Name))
	return
}

func (notecomplexityshape *NoteComplexityShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notecomplexityshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteComplexityShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(notecomplexityshape.Name))
	return
}

func (noteeffortshape *NoteEffortShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteeffortshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteEffortShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteeffortshape.Name))
	return
}

func (noteperformanceshape *NotePerformanceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteperformanceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NotePerformanceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteperformanceshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteshape.Name))
	return
}

func (performance *Performance) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", performance.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Performance")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(performance.Name))
	return
}

func (system *System) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Name))
	return
}

// insertion point for unstaging
func (compareanalysis *CompareAnalysis) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", compareanalysis.GongGetReferenceIdentifier(stage))
	return
}

func (complexity *Complexity) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexity.GongGetReferenceIdentifier(stage))
	return
}

func (diagramflossequation *DiagramFlossEquation) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramflossequation.GongGetReferenceIdentifier(stage))
	return
}

func (effort *Effort) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effort.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetReferenceIdentifier(stage))
	return
}

func (notecomplexityshape *NoteComplexityShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notecomplexityshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteeffortshape *NoteEffortShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteeffortshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteperformanceshape *NotePerformanceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteperformanceshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetReferenceIdentifier(stage))
	return
}

func (performance *Performance) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", performance.GongGetReferenceIdentifier(stage))
	return
}

func (system *System) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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

// end of template

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
	// Compute reverse map for named struct Complexity
	// insertion point per field

	// Compute reverse map for named struct ComplexityShape
	// insertion point per field

	// Compute reverse map for named struct DiagramFloss
	// insertion point per field
	stage.DiagramFloss_System_Shapes_reverseMap = make(map[*SystemShape]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _systemshape := range diagramfloss.System_Shapes {
			stage.DiagramFloss_System_Shapes_reverseMap[_systemshape] = diagramfloss
		}
	}
	stage.DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _system := range diagramfloss.SystemsWhoseNodeIsExpanded {
			stage.DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap[_system] = diagramfloss
		}
	}
	stage.DiagramFloss_Complexity_Shapes_reverseMap = make(map[*ComplexityShape]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _complexityshape := range diagramfloss.Complexity_Shapes {
			stage.DiagramFloss_Complexity_Shapes_reverseMap[_complexityshape] = diagramfloss
		}
	}
	stage.DiagramFloss_ComplexitysWhoseNodeIsExpanded_reverseMap = make(map[*Complexity]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _complexity := range diagramfloss.ComplexitysWhoseNodeIsExpanded {
			stage.DiagramFloss_ComplexitysWhoseNodeIsExpanded_reverseMap[_complexity] = diagramfloss
		}
	}
	stage.DiagramFloss_Performance_Shapes_reverseMap = make(map[*PerformanceShape]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _performanceshape := range diagramfloss.Performance_Shapes {
			stage.DiagramFloss_Performance_Shapes_reverseMap[_performanceshape] = diagramfloss
		}
	}
	stage.DiagramFloss_PerformancesWhoseNodeIsExpanded_reverseMap = make(map[*Performance]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _performance := range diagramfloss.PerformancesWhoseNodeIsExpanded {
			stage.DiagramFloss_PerformancesWhoseNodeIsExpanded_reverseMap[_performance] = diagramfloss
		}
	}
	stage.DiagramFloss_Effort_Shapes_reverseMap = make(map[*EffortShape]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _effortshape := range diagramfloss.Effort_Shapes {
			stage.DiagramFloss_Effort_Shapes_reverseMap[_effortshape] = diagramfloss
		}
	}
	stage.DiagramFloss_EffortsWhoseNodeIsExpanded_reverseMap = make(map[*Effort]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _effort := range diagramfloss.EffortsWhoseNodeIsExpanded {
			stage.DiagramFloss_EffortsWhoseNodeIsExpanded_reverseMap[_effort] = diagramfloss
		}
	}

	// Compute reverse map for named struct Effort
	// insertion point per field

	// Compute reverse map for named struct EffortShape
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

	// Compute reverse map for named struct Performance
	// insertion point per field

	// Compute reverse map for named struct PerformanceShape
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
	stage.System_DiagramFlosses_reverseMap = make(map[*DiagramFloss]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramfloss := range system.DiagramFlosses {
			stage.System_DiagramFlosses_reverseMap[_diagramfloss] = system
		}
	}
	stage.System_DiagramFlossWhoseNodeIsExpanded_reverseMap = make(map[*DiagramFloss]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramfloss := range system.DiagramFlossWhoseNodeIsExpanded {
			stage.System_DiagramFlossWhoseNodeIsExpanded_reverseMap[_diagramfloss] = system
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

	// Compute reverse map for named struct SystemShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Complexitys {
		res = append(res, instance)
	}

	for instance := range stage.ComplexityShapes {
		res = append(res, instance)
	}

	for instance := range stage.DiagramFlosss {
		res = append(res, instance)
	}

	for instance := range stage.Efforts {
		res = append(res, instance)
	}

	for instance := range stage.EffortShapes {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Performances {
		res = append(res, instance)
	}

	for instance := range stage.PerformanceShapes {
		res = append(res, instance)
	}

	for instance := range stage.Systems {
		res = append(res, instance)
	}

	for instance := range stage.SystemShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (complexity *Complexity) GongCopy() GongstructIF {
	newInstance := new(Complexity)
	complexity.CopyBasicFields(newInstance)
	return newInstance
}

func (complexityshape *ComplexityShape) GongCopy() GongstructIF {
	newInstance := new(ComplexityShape)
	complexityshape.CopyBasicFields(newInstance)
	return newInstance
}

func (diagramfloss *DiagramFloss) GongCopy() GongstructIF {
	newInstance := new(DiagramFloss)
	diagramfloss.CopyBasicFields(newInstance)
	return newInstance
}

func (effort *Effort) GongCopy() GongstructIF {
	newInstance := new(Effort)
	effort.CopyBasicFields(newInstance)
	return newInstance
}

func (effortshape *EffortShape) GongCopy() GongstructIF {
	newInstance := new(EffortShape)
	effortshape.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (performance *Performance) GongCopy() GongstructIF {
	newInstance := new(Performance)
	performance.CopyBasicFields(newInstance)
	return newInstance
}

func (performanceshape *PerformanceShape) GongCopy() GongstructIF {
	newInstance := new(PerformanceShape)
	performanceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (system *System) GongCopy() GongstructIF {
	newInstance := new(System)
	system.CopyBasicFields(newInstance)
	return newInstance
}

func (systemshape *SystemShape) GongCopy() GongstructIF {
	newInstance := new(SystemShape)
	systemshape.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (complexity *Complexity) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complexity).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(complexity), uint64(GetOrderPointerGongstruct(stage, complexity)))
	return
}

func (complexityshape *ComplexityShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complexityshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(complexityshape), uint64(GetOrderPointerGongstruct(stage, complexityshape)))
	return
}

func (diagramfloss *DiagramFloss) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramfloss).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramfloss), uint64(GetOrderPointerGongstruct(stage, diagramfloss)))
	return
}

func (effort *Effort) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(effort).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(effort), uint64(GetOrderPointerGongstruct(stage, effort)))
	return
}

func (effortshape *EffortShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(effortshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(effortshape), uint64(GetOrderPointerGongstruct(stage, effortshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (performance *Performance) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(performance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(performance), uint64(GetOrderPointerGongstruct(stage, performance)))
	return
}

func (performanceshape *PerformanceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(performanceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(performanceshape), uint64(GetOrderPointerGongstruct(stage, performanceshape)))
	return
}

func (system *System) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(system), uint64(GetOrderPointerGongstruct(stage, system)))
	return
}

func (systemshape *SystemShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(systemshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(systemshape), uint64(GetOrderPointerGongstruct(stage, systemshape)))
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
	var complexityshapes_newInstances []*ComplexityShape
	var complexityshapes_deletedInstances []*ComplexityShape

	// parse all staged instances and check if they have a reference
	for complexityshape := range stage.ComplexityShapes {
		if ref, ok := stage.ComplexityShapes_reference[complexityshape]; !ok {
			complexityshapes_newInstances = append(complexityshapes_newInstances, complexityshape)
			newInstancesSlice = append(newInstancesSlice, complexityshape.GongMarshallIdentifier(stage))
			if stage.ComplexityShapes_referenceOrder == nil {
				stage.ComplexityShapes_referenceOrder = make(map[*ComplexityShape]uint)
			}
			stage.ComplexityShapes_referenceOrder[complexityshape] = stage.ComplexityShape_stagedOrder[complexityshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, complexityshape.GongMarshallUnstaging(stage))
			// delete(stage.ComplexityShapes_referenceOrder, complexityshape)
			fieldInitializers, pointersInitializations := complexityshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ComplexityShape_stagedOrder[ref] = stage.ComplexityShape_stagedOrder[complexityshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := complexityshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, complexityshape)
			// delete(stage.ComplexityShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if complexityshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", complexityshape.GetName())
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
	for _, ref := range stage.ComplexityShapes_reference {
		instance := stage.ComplexityShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ComplexityShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			complexityshapes_deletedInstances = append(complexityshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(complexityshapes_newInstances)
	lenDeletedInstances += len(complexityshapes_deletedInstances)
	var diagramflosss_newInstances []*DiagramFloss
	var diagramflosss_deletedInstances []*DiagramFloss

	// parse all staged instances and check if they have a reference
	for diagramfloss := range stage.DiagramFlosss {
		if ref, ok := stage.DiagramFlosss_reference[diagramfloss]; !ok {
			diagramflosss_newInstances = append(diagramflosss_newInstances, diagramfloss)
			newInstancesSlice = append(newInstancesSlice, diagramfloss.GongMarshallIdentifier(stage))
			if stage.DiagramFlosss_referenceOrder == nil {
				stage.DiagramFlosss_referenceOrder = make(map[*DiagramFloss]uint)
			}
			stage.DiagramFlosss_referenceOrder[diagramfloss] = stage.DiagramFloss_stagedOrder[diagramfloss]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagramfloss.GongMarshallUnstaging(stage))
			// delete(stage.DiagramFlosss_referenceOrder, diagramfloss)
			fieldInitializers, pointersInitializations := diagramfloss.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramFloss_stagedOrder[ref] = stage.DiagramFloss_stagedOrder[diagramfloss]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagramfloss.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagramfloss)
			// delete(stage.DiagramFloss_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagramfloss.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagramfloss.GetName())
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
	for _, ref := range stage.DiagramFlosss_reference {
		instance := stage.DiagramFlosss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DiagramFlosss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagramflosss_deletedInstances = append(diagramflosss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagramflosss_newInstances)
	lenDeletedInstances += len(diagramflosss_deletedInstances)
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
	var effortshapes_newInstances []*EffortShape
	var effortshapes_deletedInstances []*EffortShape

	// parse all staged instances and check if they have a reference
	for effortshape := range stage.EffortShapes {
		if ref, ok := stage.EffortShapes_reference[effortshape]; !ok {
			effortshapes_newInstances = append(effortshapes_newInstances, effortshape)
			newInstancesSlice = append(newInstancesSlice, effortshape.GongMarshallIdentifier(stage))
			if stage.EffortShapes_referenceOrder == nil {
				stage.EffortShapes_referenceOrder = make(map[*EffortShape]uint)
			}
			stage.EffortShapes_referenceOrder[effortshape] = stage.EffortShape_stagedOrder[effortshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, effortshape.GongMarshallUnstaging(stage))
			// delete(stage.EffortShapes_referenceOrder, effortshape)
			fieldInitializers, pointersInitializations := effortshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EffortShape_stagedOrder[ref] = stage.EffortShape_stagedOrder[effortshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := effortshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, effortshape)
			// delete(stage.EffortShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if effortshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", effortshape.GetName())
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
	for _, ref := range stage.EffortShapes_reference {
		instance := stage.EffortShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EffortShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			effortshapes_deletedInstances = append(effortshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(effortshapes_newInstances)
	lenDeletedInstances += len(effortshapes_deletedInstances)
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
	var performanceshapes_newInstances []*PerformanceShape
	var performanceshapes_deletedInstances []*PerformanceShape

	// parse all staged instances and check if they have a reference
	for performanceshape := range stage.PerformanceShapes {
		if ref, ok := stage.PerformanceShapes_reference[performanceshape]; !ok {
			performanceshapes_newInstances = append(performanceshapes_newInstances, performanceshape)
			newInstancesSlice = append(newInstancesSlice, performanceshape.GongMarshallIdentifier(stage))
			if stage.PerformanceShapes_referenceOrder == nil {
				stage.PerformanceShapes_referenceOrder = make(map[*PerformanceShape]uint)
			}
			stage.PerformanceShapes_referenceOrder[performanceshape] = stage.PerformanceShape_stagedOrder[performanceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, performanceshape.GongMarshallUnstaging(stage))
			// delete(stage.PerformanceShapes_referenceOrder, performanceshape)
			fieldInitializers, pointersInitializations := performanceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PerformanceShape_stagedOrder[ref] = stage.PerformanceShape_stagedOrder[performanceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := performanceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, performanceshape)
			// delete(stage.PerformanceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if performanceshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", performanceshape.GetName())
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
	for _, ref := range stage.PerformanceShapes_reference {
		instance := stage.PerformanceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PerformanceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			performanceshapes_deletedInstances = append(performanceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(performanceshapes_newInstances)
	lenDeletedInstances += len(performanceshapes_deletedInstances)
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
	var systemshapes_newInstances []*SystemShape
	var systemshapes_deletedInstances []*SystemShape

	// parse all staged instances and check if they have a reference
	for systemshape := range stage.SystemShapes {
		if ref, ok := stage.SystemShapes_reference[systemshape]; !ok {
			systemshapes_newInstances = append(systemshapes_newInstances, systemshape)
			newInstancesSlice = append(newInstancesSlice, systemshape.GongMarshallIdentifier(stage))
			if stage.SystemShapes_referenceOrder == nil {
				stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint)
			}
			stage.SystemShapes_referenceOrder[systemshape] = stage.SystemShape_stagedOrder[systemshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, systemshape.GongMarshallUnstaging(stage))
			// delete(stage.SystemShapes_referenceOrder, systemshape)
			fieldInitializers, pointersInitializations := systemshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SystemShape_stagedOrder[ref] = stage.SystemShape_stagedOrder[systemshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := systemshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, systemshape)
			// delete(stage.SystemShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if systemshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", systemshape.GetName())
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
	for _, ref := range stage.SystemShapes_reference {
		instance := stage.SystemShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SystemShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			systemshapes_deletedInstances = append(systemshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(systemshapes_newInstances)
	lenDeletedInstances += len(systemshapes_deletedInstances)

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
	stage.Complexitys_reference = make(map[*Complexity]*Complexity)
	stage.Complexitys_referenceOrder = make(map[*Complexity]uint) // diff Unstage needs the reference order
	stage.Complexitys_instance = make(map[*Complexity]*Complexity)
	for instance := range stage.Complexitys {
		_copy := instance.GongCopy().(*Complexity)
		stage.Complexitys_reference[instance] = _copy
		stage.Complexitys_instance[_copy] = instance
		stage.Complexitys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ComplexityShapes_reference = make(map[*ComplexityShape]*ComplexityShape)
	stage.ComplexityShapes_referenceOrder = make(map[*ComplexityShape]uint) // diff Unstage needs the reference order
	stage.ComplexityShapes_instance = make(map[*ComplexityShape]*ComplexityShape)
	for instance := range stage.ComplexityShapes {
		_copy := instance.GongCopy().(*ComplexityShape)
		stage.ComplexityShapes_reference[instance] = _copy
		stage.ComplexityShapes_instance[_copy] = instance
		stage.ComplexityShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DiagramFlosss_reference = make(map[*DiagramFloss]*DiagramFloss)
	stage.DiagramFlosss_referenceOrder = make(map[*DiagramFloss]uint) // diff Unstage needs the reference order
	stage.DiagramFlosss_instance = make(map[*DiagramFloss]*DiagramFloss)
	for instance := range stage.DiagramFlosss {
		_copy := instance.GongCopy().(*DiagramFloss)
		stage.DiagramFlosss_reference[instance] = _copy
		stage.DiagramFlosss_instance[_copy] = instance
		stage.DiagramFlosss_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.EffortShapes_reference = make(map[*EffortShape]*EffortShape)
	stage.EffortShapes_referenceOrder = make(map[*EffortShape]uint) // diff Unstage needs the reference order
	stage.EffortShapes_instance = make(map[*EffortShape]*EffortShape)
	for instance := range stage.EffortShapes {
		_copy := instance.GongCopy().(*EffortShape)
		stage.EffortShapes_reference[instance] = _copy
		stage.EffortShapes_instance[_copy] = instance
		stage.EffortShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Performances_reference = make(map[*Performance]*Performance)
	stage.Performances_referenceOrder = make(map[*Performance]uint) // diff Unstage needs the reference order
	stage.Performances_instance = make(map[*Performance]*Performance)
	for instance := range stage.Performances {
		_copy := instance.GongCopy().(*Performance)
		stage.Performances_reference[instance] = _copy
		stage.Performances_instance[_copy] = instance
		stage.Performances_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PerformanceShapes_reference = make(map[*PerformanceShape]*PerformanceShape)
	stage.PerformanceShapes_referenceOrder = make(map[*PerformanceShape]uint) // diff Unstage needs the reference order
	stage.PerformanceShapes_instance = make(map[*PerformanceShape]*PerformanceShape)
	for instance := range stage.PerformanceShapes {
		_copy := instance.GongCopy().(*PerformanceShape)
		stage.PerformanceShapes_reference[instance] = _copy
		stage.PerformanceShapes_instance[_copy] = instance
		stage.PerformanceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.SystemShapes_reference = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint) // diff Unstage needs the reference order
	stage.SystemShapes_instance = make(map[*SystemShape]*SystemShape)
	for instance := range stage.SystemShapes {
		_copy := instance.GongCopy().(*SystemShape)
		stage.SystemShapes_reference[instance] = _copy
		stage.SystemShapes_instance[_copy] = instance
		stage.SystemShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Complexitys {
		reference := stage.Complexitys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ComplexityShapes {
		reference := stage.ComplexityShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DiagramFlosss {
		reference := stage.DiagramFlosss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Efforts {
		reference := stage.Efforts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EffortShapes {
		reference := stage.EffortShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Performances {
		reference := stage.Performances_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PerformanceShapes {
		reference := stage.PerformanceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Systems {
		reference := stage.Systems_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SystemShapes {
		reference := stage.SystemShapes_reference[instance]
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

func (complexityshape *ComplexityShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ComplexityShape_stagedOrder[complexityshape]; ok {
		return order
	}
	if order, ok := stage.ComplexityShapes_referenceOrder[complexityshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ComplexityShape was not staged and does not have a reference order", complexityshape)
		return 0
	}
}

func (diagramfloss *DiagramFloss) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramFloss_stagedOrder[diagramfloss]; ok {
		return order
	}
	if order, ok := stage.DiagramFlosss_referenceOrder[diagramfloss]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramFloss was not staged and does not have a reference order", diagramfloss)
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

func (effortshape *EffortShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EffortShape_stagedOrder[effortshape]; ok {
		return order
	}
	if order, ok := stage.EffortShapes_referenceOrder[effortshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EffortShape was not staged and does not have a reference order", effortshape)
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

func (performanceshape *PerformanceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PerformanceShape_stagedOrder[performanceshape]; ok {
		return order
	}
	if order, ok := stage.PerformanceShapes_referenceOrder[performanceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PerformanceShape was not staged and does not have a reference order", performanceshape)
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

func (systemshape *SystemShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SystemShape_stagedOrder[systemshape]; ok {
		return order
	}
	if order, ok := stage.SystemShapes_referenceOrder[systemshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SystemShape was not staged and does not have a reference order", systemshape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (complexity *Complexity) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexity.GongGetGongstructName(), complexity.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexity *Complexity) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexity.GongGetGongstructName(), complexity.GongGetOrder(stage))
}

func (complexityshape *ComplexityShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexityshape.GongGetGongstructName(), complexityshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexityshape *ComplexityShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexityshape.GongGetGongstructName(), complexityshape.GongGetOrder(stage))
}

func (diagramfloss *DiagramFloss) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramfloss.GongGetGongstructName(), diagramfloss.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramfloss *DiagramFloss) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramfloss.GongGetGongstructName(), diagramfloss.GongGetOrder(stage))
}

func (effort *Effort) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effort.GongGetGongstructName(), effort.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effort *Effort) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effort.GongGetGongstructName(), effort.GongGetOrder(stage))
}

func (effortshape *EffortShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effortshape.GongGetGongstructName(), effortshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effortshape *EffortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effortshape.GongGetGongstructName(), effortshape.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (performance *Performance) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", performance.GongGetGongstructName(), performance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (performance *Performance) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", performance.GongGetGongstructName(), performance.GongGetOrder(stage))
}

func (performanceshape *PerformanceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", performanceshape.GongGetGongstructName(), performanceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (performanceshape *PerformanceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", performanceshape.GongGetGongstructName(), performanceshape.GongGetOrder(stage))
}

func (system *System) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system *System) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

func (systemshape *SystemShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", systemshape.GongGetGongstructName(), systemshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemshape *SystemShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", systemshape.GongGetGongstructName(), systemshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (complexity *Complexity) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexity.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Complexity")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(complexity.Name))
	return
}

func (complexityshape *ComplexityShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexityshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexityShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(complexityshape.Name))
	return
}

func (diagramfloss *DiagramFloss) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramfloss.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramFloss")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramfloss.Name))
	return
}

func (effort *Effort) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effort.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Effort")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(effort.Name))
	return
}

func (effortshape *EffortShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effortshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EffortShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(effortshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (performance *Performance) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", performance.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Performance")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(performance.Name))
	return
}

func (performanceshape *PerformanceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", performanceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerformanceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(performanceshape.Name))
	return
}

func (system *System) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Name))
	return
}

func (systemshape *SystemShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SystemShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(systemshape.Name))
	return
}

// insertion point for unstaging
func (complexity *Complexity) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexity.GongGetReferenceIdentifier(stage))
	return
}

func (complexityshape *ComplexityShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexityshape.GongGetReferenceIdentifier(stage))
	return
}

func (diagramfloss *DiagramFloss) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramfloss.GongGetReferenceIdentifier(stage))
	return
}

func (effort *Effort) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effort.GongGetReferenceIdentifier(stage))
	return
}

func (effortshape *EffortShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effortshape.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (performance *Performance) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", performance.GongGetReferenceIdentifier(stage))
	return
}

func (performanceshape *PerformanceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", performanceshape.GongGetReferenceIdentifier(stage))
	return
}

func (system *System) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetReferenceIdentifier(stage))
	return
}

func (systemshape *SystemShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetReferenceIdentifier(stage))
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

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
	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_Product_Shapes_reverseMap = make(map[*ProductShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _productshape := range diagram.Product_Shapes {
			stage.Diagram_Product_Shapes_reverseMap[_productshape] = diagram
		}
	}
	stage.Diagram_ProductsWhoseNodeIsExpanded_reverseMap = make(map[*Product]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
			stage.Diagram_ProductsWhoseNodeIsExpanded_reverseMap[_product] = diagram
		}
	}
	stage.Diagram_ProductComposition_Shapes_reverseMap = make(map[*ProductCompositionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
			stage.Diagram_ProductComposition_Shapes_reverseMap[_productcompositionshape] = diagram
		}
	}
	stage.Diagram_ProductReference_Shapes_reverseMap = make(map[*ProductReferenceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _productreferenceshape := range diagram.ProductReference_Shapes {
			stage.Diagram_ProductReference_Shapes_reverseMap[_productreferenceshape] = diagram
		}
	}
	stage.Diagram_Task_Shapes_reverseMap = make(map[*TaskShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskshape := range diagram.Task_Shapes {
			stage.Diagram_Task_Shapes_reverseMap[_taskshape] = diagram
		}
	}
	stage.Diagram_TasksWhoseNodeIsExpanded_reverseMap = make(map[*Task]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _task := range diagram.TasksWhoseNodeIsExpanded {
			stage.Diagram_TasksWhoseNodeIsExpanded_reverseMap[_task] = diagram
		}
	}
	stage.Diagram_TasksWhoseInputNodeIsExpanded_reverseMap = make(map[*Task]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
			stage.Diagram_TasksWhoseInputNodeIsExpanded_reverseMap[_task] = diagram
		}
	}
	stage.Diagram_TasksWhoseOutputNodeIsExpanded_reverseMap = make(map[*Task]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
			stage.Diagram_TasksWhoseOutputNodeIsExpanded_reverseMap[_task] = diagram
		}
	}
	stage.Diagram_TasksWhosePredecessorNodeIsExpanded_reverseMap = make(map[*Task]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _task := range diagram.TasksWhosePredecessorNodeIsExpanded {
			stage.Diagram_TasksWhosePredecessorNodeIsExpanded_reverseMap[_task] = diagram
		}
	}
	stage.Diagram_TaskGroupShapes_reverseMap = make(map[*TaskGroupShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskgroupshape := range diagram.TaskGroupShapes {
			stage.Diagram_TaskGroupShapes_reverseMap[_taskgroupshape] = diagram
		}
	}
	stage.Diagram_TaskGroupsWhoseNodeIsExpanded_reverseMap = make(map[*TaskGroup]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskgroup := range diagram.TaskGroupsWhoseNodeIsExpanded {
			stage.Diagram_TaskGroupsWhoseNodeIsExpanded_reverseMap[_taskgroup] = diagram
		}
	}
	stage.Diagram_TaskComposition_Shapes_reverseMap = make(map[*TaskCompositionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
			stage.Diagram_TaskComposition_Shapes_reverseMap[_taskcompositionshape] = diagram
		}
	}
	stage.Diagram_TaskInputShapes_reverseMap = make(map[*TaskInputShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskinputshape := range diagram.TaskInputShapes {
			stage.Diagram_TaskInputShapes_reverseMap[_taskinputshape] = diagram
		}
	}
	stage.Diagram_TaskOutputShapes_reverseMap = make(map[*TaskOutputShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskoutputshape := range diagram.TaskOutputShapes {
			stage.Diagram_TaskOutputShapes_reverseMap[_taskoutputshape] = diagram
		}
	}
	stage.Diagram_TaskPredecessorShapes_reverseMap = make(map[*TaskPredecessorShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _taskpredecessorshape := range diagram.TaskPredecessorShapes {
			stage.Diagram_TaskPredecessorShapes_reverseMap[_taskpredecessorshape] = diagram
		}
	}
	stage.Diagram_Note_Shapes_reverseMap = make(map[*NoteShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _noteshape := range diagram.Note_Shapes {
			stage.Diagram_Note_Shapes_reverseMap[_noteshape] = diagram
		}
	}
	stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _note := range diagram.NotesWhoseNodeIsExpanded {
			stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap[_note] = diagram
		}
	}
	stage.Diagram_NoteProductShapes_reverseMap = make(map[*NoteProductShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _noteproductshape := range diagram.NoteProductShapes {
			stage.Diagram_NoteProductShapes_reverseMap[_noteproductshape] = diagram
		}
	}
	stage.Diagram_NoteTaskShapes_reverseMap = make(map[*NoteTaskShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _notetaskshape := range diagram.NoteTaskShapes {
			stage.Diagram_NoteTaskShapes_reverseMap[_notetaskshape] = diagram
		}
	}
	stage.Diagram_NoteResourceShapes_reverseMap = make(map[*NoteResourceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _noteresourceshape := range diagram.NoteResourceShapes {
			stage.Diagram_NoteResourceShapes_reverseMap[_noteresourceshape] = diagram
		}
	}
	stage.Diagram_Resource_Shapes_reverseMap = make(map[*ResourceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _resourceshape := range diagram.Resource_Shapes {
			stage.Diagram_Resource_Shapes_reverseMap[_resourceshape] = diagram
		}
	}
	stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap = make(map[*Resource]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _resource := range diagram.ResourcesWhoseNodeIsExpanded {
			stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap[_resource] = diagram
		}
	}
	stage.Diagram_ResourceComposition_Shapes_reverseMap = make(map[*ResourceCompositionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _resourcecompositionshape := range diagram.ResourceComposition_Shapes {
			stage.Diagram_ResourceComposition_Shapes_reverseMap[_resourcecompositionshape] = diagram
		}
	}
	stage.Diagram_ResourceTaskShapes_reverseMap = make(map[*ResourceTaskShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _resourcetaskshape := range diagram.ResourceTaskShapes {
			stage.Diagram_ResourceTaskShapes_reverseMap[_resourcetaskshape] = diagram
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
	stage.Library_RootProducts_reverseMap = make(map[*Product]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _product := range library.RootProducts {
			stage.Library_RootProducts_reverseMap[_product] = library
		}
	}
	stage.Library_RootTasks_reverseMap = make(map[*Task]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _task := range library.RootTasks {
			stage.Library_RootTasks_reverseMap[_task] = library
		}
	}
	stage.Library_RootTaskGroups_reverseMap = make(map[*TaskGroup]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _taskgroup := range library.RootTaskGroups {
			stage.Library_RootTaskGroups_reverseMap[_taskgroup] = library
		}
	}
	stage.Library_RootResources_reverseMap = make(map[*Resource]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _resource := range library.RootResources {
			stage.Library_RootResources_reverseMap[_resource] = library
		}
	}
	stage.Library_Notes_reverseMap = make(map[*Note]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _note := range library.Notes {
			stage.Library_Notes_reverseMap[_note] = library
		}
	}
	stage.Library_Diagrams_reverseMap = make(map[*Diagram]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _diagram := range library.Diagrams {
			stage.Library_Diagrams_reverseMap[_diagram] = library
		}
	}

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Products_reverseMap = make(map[*Product]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _product := range note.Products {
			stage.Note_Products_reverseMap[_product] = note
		}
	}
	stage.Note_Tasks_reverseMap = make(map[*Task]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _task := range note.Tasks {
			stage.Note_Tasks_reverseMap[_task] = note
		}
	}
	stage.Note_Resources_reverseMap = make(map[*Resource]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _resource := range note.Resources {
			stage.Note_Resources_reverseMap[_resource] = note
		}
	}

	// Compute reverse map for named struct Product
	// insertion point per field
	stage.Product_SubProducts_reverseMap = make(map[*Product]*Product)
	for product := range stage.Products {
		_ = product
		for _, _product := range product.SubProducts {
			stage.Product_SubProducts_reverseMap[_product] = product
		}
	}

	// Compute reverse map for named struct Resource
	// insertion point per field
	stage.Resource_Tasks_reverseMap = make(map[*Task]*Resource)
	for resource := range stage.Resources {
		_ = resource
		for _, _task := range resource.Tasks {
			stage.Resource_Tasks_reverseMap[_task] = resource
		}
	}
	stage.Resource_SubResources_reverseMap = make(map[*Resource]*Resource)
	for resource := range stage.Resources {
		_ = resource
		for _, _resource := range resource.SubResources {
			stage.Resource_SubResources_reverseMap[_resource] = resource
		}
	}

	// Compute reverse map for named struct Task
	// insertion point per field
	stage.Task_Predecessors_reverseMap = make(map[*Task]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _task := range task.Predecessors {
			stage.Task_Predecessors_reverseMap[_task] = task
		}
	}
	stage.Task_Inputs_reverseMap = make(map[*Product]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _product := range task.Inputs {
			stage.Task_Inputs_reverseMap[_product] = task
		}
	}
	stage.Task_Outputs_reverseMap = make(map[*Product]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _product := range task.Outputs {
			stage.Task_Outputs_reverseMap[_product] = task
		}
	}
	stage.Task_SubTasks_reverseMap = make(map[*Task]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _task := range task.SubTasks {
			stage.Task_SubTasks_reverseMap[_task] = task
		}
	}
	stage.Task_TaskGroupsToDisplay_reverseMap = make(map[*TaskGroup]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _taskgroup := range task.TaskGroupsToDisplay {
			stage.Task_TaskGroupsToDisplay_reverseMap[_taskgroup] = task
		}
	}

	// Compute reverse map for named struct TaskGroup
	// insertion point per field
	stage.TaskGroup_Tasks_reverseMap = make(map[*Task]*TaskGroup)
	for taskgroup := range stage.TaskGroups {
		_ = taskgroup
		for _, _task := range taskgroup.Tasks {
			stage.TaskGroup_Tasks_reverseMap[_task] = taskgroup
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Diagrams)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NoteProductShapes)

	res = __gong__appendInstances(res, stage.NoteResourceShapes)

	res = __gong__appendInstances(res, stage.NoteShapes)

	res = __gong__appendInstances(res, stage.NoteTaskShapes)

	res = __gong__appendInstances(res, stage.Products)

	res = __gong__appendInstances(res, stage.ProductCompositionShapes)

	res = __gong__appendInstances(res, stage.ProductReferenceShapes)

	res = __gong__appendInstances(res, stage.ProductShapes)

	res = __gong__appendInstances(res, stage.Resources)

	res = __gong__appendInstances(res, stage.ResourceCompositionShapes)

	res = __gong__appendInstances(res, stage.ResourceShapes)

	res = __gong__appendInstances(res, stage.ResourceTaskShapes)

	res = __gong__appendInstances(res, stage.Tasks)

	res = __gong__appendInstances(res, stage.TaskCompositionShapes)

	res = __gong__appendInstances(res, stage.TaskGroups)

	res = __gong__appendInstances(res, stage.TaskGroupShapes)

	res = __gong__appendInstances(res, stage.TaskInputShapes)

	res = __gong__appendInstances(res, stage.TaskOutputShapes)

	res = __gong__appendInstances(res, stage.TaskPredecessorShapes)

	res = __gong__appendInstances(res, stage.TaskShapes)

	return
}

// insertion point per named struct
func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.GongCopyBasicFields(newInstance)
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

func (noteproductshape *NoteProductShape) GongCopy() GongstructIF {
	newInstance := new(NoteProductShape)
	noteproductshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteresourceshape *NoteResourceShape) GongCopy() GongstructIF {
	newInstance := new(NoteResourceShape)
	noteresourceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notetaskshape *NoteTaskShape) GongCopy() GongstructIF {
	newInstance := new(NoteTaskShape)
	notetaskshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (product *Product) GongCopy() GongstructIF {
	newInstance := new(Product)
	product.GongCopyBasicFields(newInstance)
	return newInstance
}

func (productcompositionshape *ProductCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ProductCompositionShape)
	productcompositionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (productreferenceshape *ProductReferenceShape) GongCopy() GongstructIF {
	newInstance := new(ProductReferenceShape)
	productreferenceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (productshape *ProductShape) GongCopy() GongstructIF {
	newInstance := new(ProductShape)
	productshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (resource *Resource) GongCopy() GongstructIF {
	newInstance := new(Resource)
	resource.GongCopyBasicFields(newInstance)
	return newInstance
}

func (resourcecompositionshape *ResourceCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ResourceCompositionShape)
	resourcecompositionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (resourceshape *ResourceShape) GongCopy() GongstructIF {
	newInstance := new(ResourceShape)
	resourceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (resourcetaskshape *ResourceTaskShape) GongCopy() GongstructIF {
	newInstance := new(ResourceTaskShape)
	resourcetaskshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (task *Task) GongCopy() GongstructIF {
	newInstance := new(Task)
	task.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskcompositionshape *TaskCompositionShape) GongCopy() GongstructIF {
	newInstance := new(TaskCompositionShape)
	taskcompositionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskgroup *TaskGroup) GongCopy() GongstructIF {
	newInstance := new(TaskGroup)
	taskgroup.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskgroupshape *TaskGroupShape) GongCopy() GongstructIF {
	newInstance := new(TaskGroupShape)
	taskgroupshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskinputshape *TaskInputShape) GongCopy() GongstructIF {
	newInstance := new(TaskInputShape)
	taskinputshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskoutputshape *TaskOutputShape) GongCopy() GongstructIF {
	newInstance := new(TaskOutputShape)
	taskoutputshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskpredecessorshape *TaskPredecessorShape) GongCopy() GongstructIF {
	newInstance := new(TaskPredecessorShape)
	taskpredecessorshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskshape *TaskShape) GongCopy() GongstructIF {
	newInstance := new(TaskShape)
	taskshape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (diagram *Diagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagram)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (noteproductshape *NoteProductShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteproductshape)
}

func (noteresourceshape *NoteResourceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteresourceshape)
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteshape)
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetaskshape)
}

func (product *Product) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, product)
}

func (productcompositionshape *ProductCompositionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, productcompositionshape)
}

func (productreferenceshape *ProductReferenceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, productreferenceshape)
}

func (productshape *ProductShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, productshape)
}

func (resource *Resource) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, resource)
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, resourcecompositionshape)
}

func (resourceshape *ResourceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, resourceshape)
}

func (resourcetaskshape *ResourceTaskShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, resourcetaskshape)
}

func (task *Task) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, task)
}

func (taskcompositionshape *TaskCompositionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskcompositionshape)
}

func (taskgroup *TaskGroup) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskgroup)
}

func (taskgroupshape *TaskGroupShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskgroupshape)
}

func (taskinputshape *TaskInputShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskinputshape)
}

func (taskoutputshape *TaskOutputShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskoutputshape)
}

func (taskpredecessorshape *TaskPredecessorShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskpredecessorshape)
}

func (taskshape *TaskShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskshape)
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
		stage.Diagrams,
		stage.Diagram_stagedOrder,
		stage.Diagrams_reference,
		&stage.Diagrams_referenceOrder,
		stage.Diagrams_instance,
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
		stage.NoteProductShapes,
		stage.NoteProductShape_stagedOrder,
		stage.NoteProductShapes_reference,
		&stage.NoteProductShapes_referenceOrder,
		stage.NoteProductShapes_instance,
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
		stage.NoteResourceShapes,
		stage.NoteResourceShape_stagedOrder,
		stage.NoteResourceShapes_reference,
		&stage.NoteResourceShapes_referenceOrder,
		stage.NoteResourceShapes_instance,
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
		stage.NoteTaskShapes,
		stage.NoteTaskShape_stagedOrder,
		stage.NoteTaskShapes_reference,
		&stage.NoteTaskShapes_referenceOrder,
		stage.NoteTaskShapes_instance,
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
		stage.Products,
		stage.Product_stagedOrder,
		stage.Products_reference,
		&stage.Products_referenceOrder,
		stage.Products_instance,
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
		stage.ProductCompositionShapes,
		stage.ProductCompositionShape_stagedOrder,
		stage.ProductCompositionShapes_reference,
		&stage.ProductCompositionShapes_referenceOrder,
		stage.ProductCompositionShapes_instance,
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
		stage.ProductReferenceShapes,
		stage.ProductReferenceShape_stagedOrder,
		stage.ProductReferenceShapes_reference,
		&stage.ProductReferenceShapes_referenceOrder,
		stage.ProductReferenceShapes_instance,
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
		stage.ProductShapes,
		stage.ProductShape_stagedOrder,
		stage.ProductShapes_reference,
		&stage.ProductShapes_referenceOrder,
		stage.ProductShapes_instance,
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
		stage.Resources,
		stage.Resource_stagedOrder,
		stage.Resources_reference,
		&stage.Resources_referenceOrder,
		stage.Resources_instance,
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
		stage.ResourceCompositionShapes,
		stage.ResourceCompositionShape_stagedOrder,
		stage.ResourceCompositionShapes_reference,
		&stage.ResourceCompositionShapes_referenceOrder,
		stage.ResourceCompositionShapes_instance,
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
		stage.ResourceShapes,
		stage.ResourceShape_stagedOrder,
		stage.ResourceShapes_reference,
		&stage.ResourceShapes_referenceOrder,
		stage.ResourceShapes_instance,
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
		stage.ResourceTaskShapes,
		stage.ResourceTaskShape_stagedOrder,
		stage.ResourceTaskShapes_reference,
		&stage.ResourceTaskShapes_referenceOrder,
		stage.ResourceTaskShapes_instance,
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
		stage.Tasks,
		stage.Task_stagedOrder,
		stage.Tasks_reference,
		&stage.Tasks_referenceOrder,
		stage.Tasks_instance,
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
		stage.TaskCompositionShapes,
		stage.TaskCompositionShape_stagedOrder,
		stage.TaskCompositionShapes_reference,
		&stage.TaskCompositionShapes_referenceOrder,
		stage.TaskCompositionShapes_instance,
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
		stage.TaskGroups,
		stage.TaskGroup_stagedOrder,
		stage.TaskGroups_reference,
		&stage.TaskGroups_referenceOrder,
		stage.TaskGroups_instance,
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
		stage.TaskGroupShapes,
		stage.TaskGroupShape_stagedOrder,
		stage.TaskGroupShapes_reference,
		&stage.TaskGroupShapes_referenceOrder,
		stage.TaskGroupShapes_instance,
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
		stage.TaskInputShapes,
		stage.TaskInputShape_stagedOrder,
		stage.TaskInputShapes_reference,
		&stage.TaskInputShapes_referenceOrder,
		stage.TaskInputShapes_instance,
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
		stage.TaskOutputShapes,
		stage.TaskOutputShape_stagedOrder,
		stage.TaskOutputShapes_reference,
		&stage.TaskOutputShapes_referenceOrder,
		stage.TaskOutputShapes_instance,
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
		stage.TaskPredecessorShapes,
		stage.TaskPredecessorShape_stagedOrder,
		stage.TaskPredecessorShapes_reference,
		&stage.TaskPredecessorShapes_referenceOrder,
		stage.TaskPredecessorShapes_instance,
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
		stage.TaskShapes,
		stage.TaskShape_stagedOrder,
		stage.TaskShapes_reference,
		&stage.TaskShapes_referenceOrder,
		stage.TaskShapes_instance,
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
	__gong__computeReferencePass1(stage, stage.Diagrams, &stage.Diagrams_reference, &stage.Diagrams_referenceOrder, &stage.Diagrams_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NoteProductShapes, &stage.NoteProductShapes_reference, &stage.NoteProductShapes_referenceOrder, &stage.NoteProductShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteResourceShapes, &stage.NoteResourceShapes_reference, &stage.NoteResourceShapes_referenceOrder, &stage.NoteResourceShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteShapes, &stage.NoteShapes_reference, &stage.NoteShapes_referenceOrder, &stage.NoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTaskShapes, &stage.NoteTaskShapes_reference, &stage.NoteTaskShapes_referenceOrder, &stage.NoteTaskShapes_instance)

	__gong__computeReferencePass1(stage, stage.Products, &stage.Products_reference, &stage.Products_referenceOrder, &stage.Products_instance)

	__gong__computeReferencePass1(stage, stage.ProductCompositionShapes, &stage.ProductCompositionShapes_reference, &stage.ProductCompositionShapes_referenceOrder, &stage.ProductCompositionShapes_instance)

	__gong__computeReferencePass1(stage, stage.ProductReferenceShapes, &stage.ProductReferenceShapes_reference, &stage.ProductReferenceShapes_referenceOrder, &stage.ProductReferenceShapes_instance)

	__gong__computeReferencePass1(stage, stage.ProductShapes, &stage.ProductShapes_reference, &stage.ProductShapes_referenceOrder, &stage.ProductShapes_instance)

	__gong__computeReferencePass1(stage, stage.Resources, &stage.Resources_reference, &stage.Resources_referenceOrder, &stage.Resources_instance)

	__gong__computeReferencePass1(stage, stage.ResourceCompositionShapes, &stage.ResourceCompositionShapes_reference, &stage.ResourceCompositionShapes_referenceOrder, &stage.ResourceCompositionShapes_instance)

	__gong__computeReferencePass1(stage, stage.ResourceShapes, &stage.ResourceShapes_reference, &stage.ResourceShapes_referenceOrder, &stage.ResourceShapes_instance)

	__gong__computeReferencePass1(stage, stage.ResourceTaskShapes, &stage.ResourceTaskShapes_reference, &stage.ResourceTaskShapes_referenceOrder, &stage.ResourceTaskShapes_instance)

	__gong__computeReferencePass1(stage, stage.Tasks, &stage.Tasks_reference, &stage.Tasks_referenceOrder, &stage.Tasks_instance)

	__gong__computeReferencePass1(stage, stage.TaskCompositionShapes, &stage.TaskCompositionShapes_reference, &stage.TaskCompositionShapes_referenceOrder, &stage.TaskCompositionShapes_instance)

	__gong__computeReferencePass1(stage, stage.TaskGroups, &stage.TaskGroups_reference, &stage.TaskGroups_referenceOrder, &stage.TaskGroups_instance)

	__gong__computeReferencePass1(stage, stage.TaskGroupShapes, &stage.TaskGroupShapes_reference, &stage.TaskGroupShapes_referenceOrder, &stage.TaskGroupShapes_instance)

	__gong__computeReferencePass1(stage, stage.TaskInputShapes, &stage.TaskInputShapes_reference, &stage.TaskInputShapes_referenceOrder, &stage.TaskInputShapes_instance)

	__gong__computeReferencePass1(stage, stage.TaskOutputShapes, &stage.TaskOutputShapes_reference, &stage.TaskOutputShapes_referenceOrder, &stage.TaskOutputShapes_instance)

	__gong__computeReferencePass1(stage, stage.TaskPredecessorShapes, &stage.TaskPredecessorShapes_reference, &stage.TaskPredecessorShapes_referenceOrder, &stage.TaskPredecessorShapes_instance)

	__gong__computeReferencePass1(stage, stage.TaskShapes, &stage.TaskShapes_reference, &stage.TaskShapes_referenceOrder, &stage.TaskShapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Diagrams, stage.Diagrams_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NoteProductShapes, stage.NoteProductShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteResourceShapes, stage.NoteResourceShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteShapes, stage.NoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTaskShapes, stage.NoteTaskShapes_reference, stage)

	__gong__computeReferencePass2(stage.Products, stage.Products_reference, stage)

	__gong__computeReferencePass2(stage.ProductCompositionShapes, stage.ProductCompositionShapes_reference, stage)

	__gong__computeReferencePass2(stage.ProductReferenceShapes, stage.ProductReferenceShapes_reference, stage)

	__gong__computeReferencePass2(stage.ProductShapes, stage.ProductShapes_reference, stage)

	__gong__computeReferencePass2(stage.Resources, stage.Resources_reference, stage)

	__gong__computeReferencePass2(stage.ResourceCompositionShapes, stage.ResourceCompositionShapes_reference, stage)

	__gong__computeReferencePass2(stage.ResourceShapes, stage.ResourceShapes_reference, stage)

	__gong__computeReferencePass2(stage.ResourceTaskShapes, stage.ResourceTaskShapes_reference, stage)

	__gong__computeReferencePass2(stage.Tasks, stage.Tasks_reference, stage)

	__gong__computeReferencePass2(stage.TaskCompositionShapes, stage.TaskCompositionShapes_reference, stage)

	__gong__computeReferencePass2(stage.TaskGroups, stage.TaskGroups_reference, stage)

	__gong__computeReferencePass2(stage.TaskGroupShapes, stage.TaskGroupShapes_reference, stage)

	__gong__computeReferencePass2(stage.TaskInputShapes, stage.TaskInputShapes_reference, stage)

	__gong__computeReferencePass2(stage.TaskOutputShapes, stage.TaskOutputShapes_reference, stage)

	__gong__computeReferencePass2(stage.TaskPredecessorShapes, stage.TaskPredecessorShapes_reference, stage)

	__gong__computeReferencePass2(stage.TaskShapes, stage.TaskShapes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Diagram_stagedOrder, stage.Diagrams_referenceOrder, diagram, "Diagram")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (noteproductshape *NoteProductShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteProductShape_stagedOrder, stage.NoteProductShapes_referenceOrder, noteproductshape, "NoteProductShape")
}

func (noteresourceshape *NoteResourceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteResourceShape_stagedOrder, stage.NoteResourceShapes_referenceOrder, noteresourceshape, "NoteResourceShape")
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteShape_stagedOrder, stage.NoteShapes_referenceOrder, noteshape, "NoteShape")
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteTaskShape_stagedOrder, stage.NoteTaskShapes_referenceOrder, notetaskshape, "NoteTaskShape")
}

func (product *Product) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Product_stagedOrder, stage.Products_referenceOrder, product, "Product")
}

func (productcompositionshape *ProductCompositionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProductCompositionShape_stagedOrder, stage.ProductCompositionShapes_referenceOrder, productcompositionshape, "ProductCompositionShape")
}

func (productreferenceshape *ProductReferenceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProductReferenceShape_stagedOrder, stage.ProductReferenceShapes_referenceOrder, productreferenceshape, "ProductReferenceShape")
}

func (productshape *ProductShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProductShape_stagedOrder, stage.ProductShapes_referenceOrder, productshape, "ProductShape")
}

func (resource *Resource) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Resource_stagedOrder, stage.Resources_referenceOrder, resource, "Resource")
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ResourceCompositionShape_stagedOrder, stage.ResourceCompositionShapes_referenceOrder, resourcecompositionshape, "ResourceCompositionShape")
}

func (resourceshape *ResourceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ResourceShape_stagedOrder, stage.ResourceShapes_referenceOrder, resourceshape, "ResourceShape")
}

func (resourcetaskshape *ResourceTaskShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ResourceTaskShape_stagedOrder, stage.ResourceTaskShapes_referenceOrder, resourcetaskshape, "ResourceTaskShape")
}

func (task *Task) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Task_stagedOrder, stage.Tasks_referenceOrder, task, "Task")
}

func (taskcompositionshape *TaskCompositionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskCompositionShape_stagedOrder, stage.TaskCompositionShapes_referenceOrder, taskcompositionshape, "TaskCompositionShape")
}

func (taskgroup *TaskGroup) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskGroup_stagedOrder, stage.TaskGroups_referenceOrder, taskgroup, "TaskGroup")
}

func (taskgroupshape *TaskGroupShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskGroupShape_stagedOrder, stage.TaskGroupShapes_referenceOrder, taskgroupshape, "TaskGroupShape")
}

func (taskinputshape *TaskInputShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskInputShape_stagedOrder, stage.TaskInputShapes_referenceOrder, taskinputshape, "TaskInputShape")
}

func (taskoutputshape *TaskOutputShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskOutputShape_stagedOrder, stage.TaskOutputShapes_referenceOrder, taskoutputshape, "TaskOutputShape")
}

func (taskpredecessorshape *TaskPredecessorShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskPredecessorShape_stagedOrder, stage.TaskPredecessorShapes_referenceOrder, taskpredecessorshape, "TaskPredecessorShape")
}

func (taskshape *TaskShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskShape_stagedOrder, stage.TaskShapes_referenceOrder, taskshape, "TaskShape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagram, diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return diagram.GongGetIdentifier(stage)
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

func (noteproductshape *NoteProductShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteproductshape, noteproductshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteproductshape *NoteProductShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteproductshape.GongGetIdentifier(stage)
}

func (noteresourceshape *NoteResourceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteresourceshape, noteresourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteresourceshape *NoteResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteresourceshape.GongGetIdentifier(stage)
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteshape, noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteshape.GongGetIdentifier(stage)
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notetaskshape, notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notetaskshape.GongGetIdentifier(stage)
}

func (product *Product) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(product, product.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (product *Product) GongGetReferenceIdentifier(stage *Stage) string {
	return product.GongGetIdentifier(stage)
}

func (productcompositionshape *ProductCompositionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(productcompositionshape, productcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productcompositionshape *ProductCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return productcompositionshape.GongGetIdentifier(stage)
}

func (productreferenceshape *ProductReferenceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(productreferenceshape, productreferenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productreferenceshape *ProductReferenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return productreferenceshape.GongGetIdentifier(stage)
}

func (productshape *ProductShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(productshape, productshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productshape *ProductShape) GongGetReferenceIdentifier(stage *Stage) string {
	return productshape.GongGetIdentifier(stage)
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(resource, resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return resource.GongGetIdentifier(stage)
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(resourcecompositionshape, resourcecompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourcecompositionshape *ResourceCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return resourcecompositionshape.GongGetIdentifier(stage)
}

func (resourceshape *ResourceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(resourceshape, resourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourceshape *ResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return resourceshape.GongGetIdentifier(stage)
}

func (resourcetaskshape *ResourceTaskShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(resourcetaskshape, resourcetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourcetaskshape *ResourceTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return resourcetaskshape.GongGetIdentifier(stage)
}

func (task *Task) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(task, task.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (task *Task) GongGetReferenceIdentifier(stage *Stage) string {
	return task.GongGetIdentifier(stage)
}

func (taskcompositionshape *TaskCompositionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskcompositionshape, taskcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskcompositionshape *TaskCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskcompositionshape.GongGetIdentifier(stage)
}

func (taskgroup *TaskGroup) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskgroup, taskgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskgroup *TaskGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return taskgroup.GongGetIdentifier(stage)
}

func (taskgroupshape *TaskGroupShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskgroupshape, taskgroupshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskgroupshape *TaskGroupShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskgroupshape.GongGetIdentifier(stage)
}

func (taskinputshape *TaskInputShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskinputshape, taskinputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskinputshape *TaskInputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskinputshape.GongGetIdentifier(stage)
}

func (taskoutputshape *TaskOutputShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskoutputshape, taskoutputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskoutputshape *TaskOutputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskoutputshape.GongGetIdentifier(stage)
}

func (taskpredecessorshape *TaskPredecessorShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskpredecessorshape, taskpredecessorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskpredecessorshape *TaskPredecessorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskpredecessorshape.GongGetIdentifier(stage)
}

func (taskshape *TaskShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskshape, taskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskshape *TaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskshape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagram.GongGetIdentifier(stage), "Diagram", diagram.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (noteproductshape *NoteProductShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteproductshape.GongGetIdentifier(stage), "NoteProductShape", noteproductshape.Name)
}

func (noteresourceshape *NoteResourceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteresourceshape.GongGetIdentifier(stage), "NoteResourceShape", noteresourceshape.Name)
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteshape.GongGetIdentifier(stage), "NoteShape", noteshape.Name)
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetaskshape.GongGetIdentifier(stage), "NoteTaskShape", notetaskshape.Name)
}

func (product *Product) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(product.GongGetIdentifier(stage), "Product", product.Name)
}

func (productcompositionshape *ProductCompositionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(productcompositionshape.GongGetIdentifier(stage), "ProductCompositionShape", productcompositionshape.Name)
}

func (productreferenceshape *ProductReferenceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(productreferenceshape.GongGetIdentifier(stage), "ProductReferenceShape", productreferenceshape.Name)
}

func (productshape *ProductShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(productshape.GongGetIdentifier(stage), "ProductShape", productshape.Name)
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(resource.GongGetIdentifier(stage), "Resource", resource.Name)
}

func (resourcecompositionshape *ResourceCompositionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(resourcecompositionshape.GongGetIdentifier(stage), "ResourceCompositionShape", resourcecompositionshape.Name)
}

func (resourceshape *ResourceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(resourceshape.GongGetIdentifier(stage), "ResourceShape", resourceshape.Name)
}

func (resourcetaskshape *ResourceTaskShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(resourcetaskshape.GongGetIdentifier(stage), "ResourceTaskShape", resourcetaskshape.Name)
}

func (task *Task) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(task.GongGetIdentifier(stage), "Task", task.Name)
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskcompositionshape.GongGetIdentifier(stage), "TaskCompositionShape", taskcompositionshape.Name)
}

func (taskgroup *TaskGroup) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskgroup.GongGetIdentifier(stage), "TaskGroup", taskgroup.Name)
}

func (taskgroupshape *TaskGroupShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskgroupshape.GongGetIdentifier(stage), "TaskGroupShape", taskgroupshape.Name)
}

func (taskinputshape *TaskInputShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskinputshape.GongGetIdentifier(stage), "TaskInputShape", taskinputshape.Name)
}

func (taskoutputshape *TaskOutputShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskoutputshape.GongGetIdentifier(stage), "TaskOutputShape", taskoutputshape.Name)
}

func (taskpredecessorshape *TaskPredecessorShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskpredecessorshape.GongGetIdentifier(stage), "TaskPredecessorShape", taskpredecessorshape.Name)
}

func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskshape.GongGetIdentifier(stage), "TaskShape", taskshape.Name)
}

// insertion point for unstaging
func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagram.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (noteproductshape *NoteProductShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteproductshape.GongGetReferenceIdentifier(stage))
}

func (noteresourceshape *NoteResourceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteresourceshape.GongGetReferenceIdentifier(stage))
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteshape.GongGetReferenceIdentifier(stage))
}

func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetaskshape.GongGetReferenceIdentifier(stage))
}

func (product *Product) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(product.GongGetReferenceIdentifier(stage))
}

func (productcompositionshape *ProductCompositionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(productcompositionshape.GongGetReferenceIdentifier(stage))
}

func (productreferenceshape *ProductReferenceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(productreferenceshape.GongGetReferenceIdentifier(stage))
}

func (productshape *ProductShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(productshape.GongGetReferenceIdentifier(stage))
}

func (resource *Resource) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(resource.GongGetReferenceIdentifier(stage))
}

func (resourcecompositionshape *ResourceCompositionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(resourcecompositionshape.GongGetReferenceIdentifier(stage))
}

func (resourceshape *ResourceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(resourceshape.GongGetReferenceIdentifier(stage))
}

func (resourcetaskshape *ResourceTaskShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(resourcetaskshape.GongGetReferenceIdentifier(stage))
}

func (task *Task) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(task.GongGetReferenceIdentifier(stage))
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskcompositionshape.GongGetReferenceIdentifier(stage))
}

func (taskgroup *TaskGroup) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskgroup.GongGetReferenceIdentifier(stage))
}

func (taskgroupshape *TaskGroupShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskgroupshape.GongGetReferenceIdentifier(stage))
}

func (taskinputshape *TaskInputShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskinputshape.GongGetReferenceIdentifier(stage))
}

func (taskoutputshape *TaskOutputShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskoutputshape.GongGetReferenceIdentifier(stage))
}

func (taskpredecessorshape *TaskPredecessorShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskpredecessorshape.GongGetReferenceIdentifier(stage))
}

func (taskshape *TaskShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskshape.GongGetReferenceIdentifier(stage))
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

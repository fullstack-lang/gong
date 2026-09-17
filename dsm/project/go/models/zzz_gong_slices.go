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

	// Compute reverse map for named struct NoteProductShape
	// insertion point per field

	// Compute reverse map for named struct NoteResourceShape
	// insertion point per field

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct NoteTaskShape
	// insertion point per field

	// Compute reverse map for named struct Product
	// insertion point per field
	stage.Product_SubProducts_reverseMap = make(map[*Product]*Product)
	for product := range stage.Products {
		_ = product
		for _, _product := range product.SubProducts {
			stage.Product_SubProducts_reverseMap[_product] = product
		}
	}

	// Compute reverse map for named struct ProductCompositionShape
	// insertion point per field

	// Compute reverse map for named struct ProductShape
	// insertion point per field

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

	// Compute reverse map for named struct ResourceCompositionShape
	// insertion point per field

	// Compute reverse map for named struct ResourceShape
	// insertion point per field

	// Compute reverse map for named struct ResourceTaskShape
	// insertion point per field

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

	// Compute reverse map for named struct TaskCompositionShape
	// insertion point per field

	// Compute reverse map for named struct TaskGroup
	// insertion point per field
	stage.TaskGroup_Tasks_reverseMap = make(map[*Task]*TaskGroup)
	for taskgroup := range stage.TaskGroups {
		_ = taskgroup
		for _, _task := range taskgroup.Tasks {
			stage.TaskGroup_Tasks_reverseMap[_task] = taskgroup
		}
	}

	// Compute reverse map for named struct TaskGroupShape
	// insertion point per field

	// Compute reverse map for named struct TaskInputShape
	// insertion point per field

	// Compute reverse map for named struct TaskOutputShape
	// insertion point per field

	// Compute reverse map for named struct TaskPredecessorShape
	// insertion point per field

	// Compute reverse map for named struct TaskShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NoteProductShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteResourceShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteTaskShapes {
		res = append(res, instance)
	}

	for instance := range stage.Products {
		res = append(res, instance)
	}

	for instance := range stage.ProductCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.ProductShapes {
		res = append(res, instance)
	}

	for instance := range stage.Resources {
		res = append(res, instance)
	}

	for instance := range stage.ResourceCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.ResourceShapes {
		res = append(res, instance)
	}

	for instance := range stage.ResourceTaskShapes {
		res = append(res, instance)
	}

	for instance := range stage.Tasks {
		res = append(res, instance)
	}

	for instance := range stage.TaskCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskGroups {
		res = append(res, instance)
	}

	for instance := range stage.TaskGroupShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskInputShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskOutputShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskPredecessorShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskShapes {
		res = append(res, instance)
	}

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
func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagram), uint64(stage.GetOrder(diagram)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note), uint64(stage.GetOrder(note)))
	return
}

func (noteproductshape *NoteProductShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteproductshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteproductshape), uint64(stage.GetOrder(noteproductshape)))
	return
}

func (noteresourceshape *NoteResourceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteresourceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteresourceshape), uint64(stage.GetOrder(noteresourceshape)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteshape), uint64(stage.GetOrder(noteshape)))
	return
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notetaskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notetaskshape), uint64(stage.GetOrder(notetaskshape)))
	return
}

func (product *Product) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(product).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(product), uint64(stage.GetOrder(product)))
	return
}

func (productcompositionshape *ProductCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(productcompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(productcompositionshape), uint64(stage.GetOrder(productcompositionshape)))
	return
}

func (productshape *ProductShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(productshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(productshape), uint64(stage.GetOrder(productshape)))
	return
}

func (resource *Resource) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resource).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(resource), uint64(stage.GetOrder(resource)))
	return
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resourcecompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(resourcecompositionshape), uint64(stage.GetOrder(resourcecompositionshape)))
	return
}

func (resourceshape *ResourceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resourceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(resourceshape), uint64(stage.GetOrder(resourceshape)))
	return
}

func (resourcetaskshape *ResourceTaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resourcetaskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(resourcetaskshape), uint64(stage.GetOrder(resourcetaskshape)))
	return
}

func (task *Task) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(task).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(task), uint64(stage.GetOrder(task)))
	return
}

func (taskcompositionshape *TaskCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskcompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskcompositionshape), uint64(stage.GetOrder(taskcompositionshape)))
	return
}

func (taskgroup *TaskGroup) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskgroup).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskgroup), uint64(stage.GetOrder(taskgroup)))
	return
}

func (taskgroupshape *TaskGroupShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskgroupshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskgroupshape), uint64(stage.GetOrder(taskgroupshape)))
	return
}

func (taskinputshape *TaskInputShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskinputshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskinputshape), uint64(stage.GetOrder(taskinputshape)))
	return
}

func (taskoutputshape *TaskOutputShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskoutputshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskoutputshape), uint64(stage.GetOrder(taskoutputshape)))
	return
}

func (taskpredecessorshape *TaskPredecessorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskpredecessorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskpredecessorshape), uint64(stage.GetOrder(taskpredecessorshape)))
	return
}

func (taskshape *TaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskshape), uint64(stage.GetOrder(taskshape)))
	return
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
	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		_copy := instance.GongCopy().(*Diagram)
		stage.Diagrams_reference[instance] = _copy
		stage.Diagrams_instance[_copy] = instance
		stage.Diagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.NoteProductShapes_reference = make(map[*NoteProductShape]*NoteProductShape)
	stage.NoteProductShapes_referenceOrder = make(map[*NoteProductShape]uint) // diff Unstage needs the reference order
	stage.NoteProductShapes_instance = make(map[*NoteProductShape]*NoteProductShape)
	for instance := range stage.NoteProductShapes {
		_copy := instance.GongCopy().(*NoteProductShape)
		stage.NoteProductShapes_reference[instance] = _copy
		stage.NoteProductShapes_instance[_copy] = instance
		stage.NoteProductShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteResourceShapes_reference = make(map[*NoteResourceShape]*NoteResourceShape)
	stage.NoteResourceShapes_referenceOrder = make(map[*NoteResourceShape]uint) // diff Unstage needs the reference order
	stage.NoteResourceShapes_instance = make(map[*NoteResourceShape]*NoteResourceShape)
	for instance := range stage.NoteResourceShapes {
		_copy := instance.GongCopy().(*NoteResourceShape)
		stage.NoteResourceShapes_reference[instance] = _copy
		stage.NoteResourceShapes_instance[_copy] = instance
		stage.NoteResourceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint) // diff Unstage needs the reference order
	stage.NoteTaskShapes_instance = make(map[*NoteTaskShape]*NoteTaskShape)
	for instance := range stage.NoteTaskShapes {
		_copy := instance.GongCopy().(*NoteTaskShape)
		stage.NoteTaskShapes_reference[instance] = _copy
		stage.NoteTaskShapes_instance[_copy] = instance
		stage.NoteTaskShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Products_reference = make(map[*Product]*Product)
	stage.Products_referenceOrder = make(map[*Product]uint) // diff Unstage needs the reference order
	stage.Products_instance = make(map[*Product]*Product)
	for instance := range stage.Products {
		_copy := instance.GongCopy().(*Product)
		stage.Products_reference[instance] = _copy
		stage.Products_instance[_copy] = instance
		stage.Products_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ProductCompositionShapes_reference = make(map[*ProductCompositionShape]*ProductCompositionShape)
	stage.ProductCompositionShapes_referenceOrder = make(map[*ProductCompositionShape]uint) // diff Unstage needs the reference order
	stage.ProductCompositionShapes_instance = make(map[*ProductCompositionShape]*ProductCompositionShape)
	for instance := range stage.ProductCompositionShapes {
		_copy := instance.GongCopy().(*ProductCompositionShape)
		stage.ProductCompositionShapes_reference[instance] = _copy
		stage.ProductCompositionShapes_instance[_copy] = instance
		stage.ProductCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ProductShapes_reference = make(map[*ProductShape]*ProductShape)
	stage.ProductShapes_referenceOrder = make(map[*ProductShape]uint) // diff Unstage needs the reference order
	stage.ProductShapes_instance = make(map[*ProductShape]*ProductShape)
	for instance := range stage.ProductShapes {
		_copy := instance.GongCopy().(*ProductShape)
		stage.ProductShapes_reference[instance] = _copy
		stage.ProductShapes_instance[_copy] = instance
		stage.ProductShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Resources_reference = make(map[*Resource]*Resource)
	stage.Resources_referenceOrder = make(map[*Resource]uint) // diff Unstage needs the reference order
	stage.Resources_instance = make(map[*Resource]*Resource)
	for instance := range stage.Resources {
		_copy := instance.GongCopy().(*Resource)
		stage.Resources_reference[instance] = _copy
		stage.Resources_instance[_copy] = instance
		stage.Resources_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ResourceCompositionShapes_reference = make(map[*ResourceCompositionShape]*ResourceCompositionShape)
	stage.ResourceCompositionShapes_referenceOrder = make(map[*ResourceCompositionShape]uint) // diff Unstage needs the reference order
	stage.ResourceCompositionShapes_instance = make(map[*ResourceCompositionShape]*ResourceCompositionShape)
	for instance := range stage.ResourceCompositionShapes {
		_copy := instance.GongCopy().(*ResourceCompositionShape)
		stage.ResourceCompositionShapes_reference[instance] = _copy
		stage.ResourceCompositionShapes_instance[_copy] = instance
		stage.ResourceCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ResourceShapes_reference = make(map[*ResourceShape]*ResourceShape)
	stage.ResourceShapes_referenceOrder = make(map[*ResourceShape]uint) // diff Unstage needs the reference order
	stage.ResourceShapes_instance = make(map[*ResourceShape]*ResourceShape)
	for instance := range stage.ResourceShapes {
		_copy := instance.GongCopy().(*ResourceShape)
		stage.ResourceShapes_reference[instance] = _copy
		stage.ResourceShapes_instance[_copy] = instance
		stage.ResourceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ResourceTaskShapes_reference = make(map[*ResourceTaskShape]*ResourceTaskShape)
	stage.ResourceTaskShapes_referenceOrder = make(map[*ResourceTaskShape]uint) // diff Unstage needs the reference order
	stage.ResourceTaskShapes_instance = make(map[*ResourceTaskShape]*ResourceTaskShape)
	for instance := range stage.ResourceTaskShapes {
		_copy := instance.GongCopy().(*ResourceTaskShape)
		stage.ResourceTaskShapes_reference[instance] = _copy
		stage.ResourceTaskShapes_instance[_copy] = instance
		stage.ResourceTaskShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tasks_reference = make(map[*Task]*Task)
	stage.Tasks_referenceOrder = make(map[*Task]uint) // diff Unstage needs the reference order
	stage.Tasks_instance = make(map[*Task]*Task)
	for instance := range stage.Tasks {
		_copy := instance.GongCopy().(*Task)
		stage.Tasks_reference[instance] = _copy
		stage.Tasks_instance[_copy] = instance
		stage.Tasks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskCompositionShapes_reference = make(map[*TaskCompositionShape]*TaskCompositionShape)
	stage.TaskCompositionShapes_referenceOrder = make(map[*TaskCompositionShape]uint) // diff Unstage needs the reference order
	stage.TaskCompositionShapes_instance = make(map[*TaskCompositionShape]*TaskCompositionShape)
	for instance := range stage.TaskCompositionShapes {
		_copy := instance.GongCopy().(*TaskCompositionShape)
		stage.TaskCompositionShapes_reference[instance] = _copy
		stage.TaskCompositionShapes_instance[_copy] = instance
		stage.TaskCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskGroups_reference = make(map[*TaskGroup]*TaskGroup)
	stage.TaskGroups_referenceOrder = make(map[*TaskGroup]uint) // diff Unstage needs the reference order
	stage.TaskGroups_instance = make(map[*TaskGroup]*TaskGroup)
	for instance := range stage.TaskGroups {
		_copy := instance.GongCopy().(*TaskGroup)
		stage.TaskGroups_reference[instance] = _copy
		stage.TaskGroups_instance[_copy] = instance
		stage.TaskGroups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskGroupShapes_reference = make(map[*TaskGroupShape]*TaskGroupShape)
	stage.TaskGroupShapes_referenceOrder = make(map[*TaskGroupShape]uint) // diff Unstage needs the reference order
	stage.TaskGroupShapes_instance = make(map[*TaskGroupShape]*TaskGroupShape)
	for instance := range stage.TaskGroupShapes {
		_copy := instance.GongCopy().(*TaskGroupShape)
		stage.TaskGroupShapes_reference[instance] = _copy
		stage.TaskGroupShapes_instance[_copy] = instance
		stage.TaskGroupShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskInputShapes_reference = make(map[*TaskInputShape]*TaskInputShape)
	stage.TaskInputShapes_referenceOrder = make(map[*TaskInputShape]uint) // diff Unstage needs the reference order
	stage.TaskInputShapes_instance = make(map[*TaskInputShape]*TaskInputShape)
	for instance := range stage.TaskInputShapes {
		_copy := instance.GongCopy().(*TaskInputShape)
		stage.TaskInputShapes_reference[instance] = _copy
		stage.TaskInputShapes_instance[_copy] = instance
		stage.TaskInputShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskOutputShapes_reference = make(map[*TaskOutputShape]*TaskOutputShape)
	stage.TaskOutputShapes_referenceOrder = make(map[*TaskOutputShape]uint) // diff Unstage needs the reference order
	stage.TaskOutputShapes_instance = make(map[*TaskOutputShape]*TaskOutputShape)
	for instance := range stage.TaskOutputShapes {
		_copy := instance.GongCopy().(*TaskOutputShape)
		stage.TaskOutputShapes_reference[instance] = _copy
		stage.TaskOutputShapes_instance[_copy] = instance
		stage.TaskOutputShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskPredecessorShapes_reference = make(map[*TaskPredecessorShape]*TaskPredecessorShape)
	stage.TaskPredecessorShapes_referenceOrder = make(map[*TaskPredecessorShape]uint) // diff Unstage needs the reference order
	stage.TaskPredecessorShapes_instance = make(map[*TaskPredecessorShape]*TaskPredecessorShape)
	for instance := range stage.TaskPredecessorShapes {
		_copy := instance.GongCopy().(*TaskPredecessorShape)
		stage.TaskPredecessorShapes_reference[instance] = _copy
		stage.TaskPredecessorShapes_instance[_copy] = instance
		stage.TaskPredecessorShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskShapes_reference = make(map[*TaskShape]*TaskShape)
	stage.TaskShapes_referenceOrder = make(map[*TaskShape]uint) // diff Unstage needs the reference order
	stage.TaskShapes_instance = make(map[*TaskShape]*TaskShape)
	for instance := range stage.TaskShapes {
		_copy := instance.GongCopy().(*TaskShape)
		stage.TaskShapes_reference[instance] = _copy
		stage.TaskShapes_instance[_copy] = instance
		stage.TaskShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Diagrams {
		reference := stage.Diagrams_reference[instance]
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

	for instance := range stage.NoteProductShapes {
		reference := stage.NoteProductShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteResourceShapes {
		reference := stage.NoteResourceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteShapes {
		reference := stage.NoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteTaskShapes {
		reference := stage.NoteTaskShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Products {
		reference := stage.Products_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ProductCompositionShapes {
		reference := stage.ProductCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ProductShapes {
		reference := stage.ProductShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Resources {
		reference := stage.Resources_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ResourceCompositionShapes {
		reference := stage.ResourceCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ResourceShapes {
		reference := stage.ResourceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ResourceTaskShapes {
		reference := stage.ResourceTaskShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tasks {
		reference := stage.Tasks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskCompositionShapes {
		reference := stage.TaskCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskGroups {
		reference := stage.TaskGroups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskGroupShapes {
		reference := stage.TaskGroupShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskInputShapes {
		reference := stage.TaskInputShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskOutputShapes {
		reference := stage.TaskOutputShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskPredecessorShapes {
		reference := stage.TaskPredecessorShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskShapes {
		reference := stage.TaskShapes_reference[instance]
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
func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Diagram_stagedOrder[diagram]; ok {
		return order
	}
	if order, ok := stage.Diagrams_referenceOrder[diagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Diagram was not staged and does not have a reference order", diagram)
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

func (noteproductshape *NoteProductShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteProductShape_stagedOrder[noteproductshape]; ok {
		return order
	}
	if order, ok := stage.NoteProductShapes_referenceOrder[noteproductshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteProductShape was not staged and does not have a reference order", noteproductshape)
		return 0
	}
}

func (noteresourceshape *NoteResourceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteResourceShape_stagedOrder[noteresourceshape]; ok {
		return order
	}
	if order, ok := stage.NoteResourceShapes_referenceOrder[noteresourceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteResourceShape was not staged and does not have a reference order", noteresourceshape)
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

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteTaskShape_stagedOrder[notetaskshape]; ok {
		return order
	}
	if order, ok := stage.NoteTaskShapes_referenceOrder[notetaskshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteTaskShape was not staged and does not have a reference order", notetaskshape)
		return 0
	}
}

func (product *Product) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Product_stagedOrder[product]; ok {
		return order
	}
	if order, ok := stage.Products_referenceOrder[product]; ok {
		return order
	} else {
		log.Printf("instance %p of type Product was not staged and does not have a reference order", product)
		return 0
	}
}

func (productcompositionshape *ProductCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProductCompositionShape_stagedOrder[productcompositionshape]; ok {
		return order
	}
	if order, ok := stage.ProductCompositionShapes_referenceOrder[productcompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ProductCompositionShape was not staged and does not have a reference order", productcompositionshape)
		return 0
	}
}

func (productshape *ProductShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProductShape_stagedOrder[productshape]; ok {
		return order
	}
	if order, ok := stage.ProductShapes_referenceOrder[productshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ProductShape was not staged and does not have a reference order", productshape)
		return 0
	}
}

func (resource *Resource) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Resource_stagedOrder[resource]; ok {
		return order
	}
	if order, ok := stage.Resources_referenceOrder[resource]; ok {
		return order
	} else {
		log.Printf("instance %p of type Resource was not staged and does not have a reference order", resource)
		return 0
	}
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ResourceCompositionShape_stagedOrder[resourcecompositionshape]; ok {
		return order
	}
	if order, ok := stage.ResourceCompositionShapes_referenceOrder[resourcecompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ResourceCompositionShape was not staged and does not have a reference order", resourcecompositionshape)
		return 0
	}
}

func (resourceshape *ResourceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ResourceShape_stagedOrder[resourceshape]; ok {
		return order
	}
	if order, ok := stage.ResourceShapes_referenceOrder[resourceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ResourceShape was not staged and does not have a reference order", resourceshape)
		return 0
	}
}

func (resourcetaskshape *ResourceTaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ResourceTaskShape_stagedOrder[resourcetaskshape]; ok {
		return order
	}
	if order, ok := stage.ResourceTaskShapes_referenceOrder[resourcetaskshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ResourceTaskShape was not staged and does not have a reference order", resourcetaskshape)
		return 0
	}
}

func (task *Task) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Task_stagedOrder[task]; ok {
		return order
	}
	if order, ok := stage.Tasks_referenceOrder[task]; ok {
		return order
	} else {
		log.Printf("instance %p of type Task was not staged and does not have a reference order", task)
		return 0
	}
}

func (taskcompositionshape *TaskCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskCompositionShape_stagedOrder[taskcompositionshape]; ok {
		return order
	}
	if order, ok := stage.TaskCompositionShapes_referenceOrder[taskcompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskCompositionShape was not staged and does not have a reference order", taskcompositionshape)
		return 0
	}
}

func (taskgroup *TaskGroup) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskGroup_stagedOrder[taskgroup]; ok {
		return order
	}
	if order, ok := stage.TaskGroups_referenceOrder[taskgroup]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskGroup was not staged and does not have a reference order", taskgroup)
		return 0
	}
}

func (taskgroupshape *TaskGroupShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskGroupShape_stagedOrder[taskgroupshape]; ok {
		return order
	}
	if order, ok := stage.TaskGroupShapes_referenceOrder[taskgroupshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskGroupShape was not staged and does not have a reference order", taskgroupshape)
		return 0
	}
}

func (taskinputshape *TaskInputShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskInputShape_stagedOrder[taskinputshape]; ok {
		return order
	}
	if order, ok := stage.TaskInputShapes_referenceOrder[taskinputshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskInputShape was not staged and does not have a reference order", taskinputshape)
		return 0
	}
}

func (taskoutputshape *TaskOutputShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskOutputShape_stagedOrder[taskoutputshape]; ok {
		return order
	}
	if order, ok := stage.TaskOutputShapes_referenceOrder[taskoutputshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskOutputShape was not staged and does not have a reference order", taskoutputshape)
		return 0
	}
}

func (taskpredecessorshape *TaskPredecessorShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskPredecessorShape_stagedOrder[taskpredecessorshape]; ok {
		return order
	}
	if order, ok := stage.TaskPredecessorShapes_referenceOrder[taskpredecessorshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskPredecessorShape was not staged and does not have a reference order", taskpredecessorshape)
		return 0
	}
}

func (taskshape *TaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskShape_stagedOrder[taskshape]; ok {
		return order
	}
	if order, ok := stage.TaskShapes_referenceOrder[taskshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskShape was not staged and does not have a reference order", taskshape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
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

func (noteproductshape *NoteProductShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteproductshape.GongGetGongstructName(), noteproductshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteproductshape *NoteProductShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteproductshape.GongGetGongstructName(), noteproductshape.GongGetOrder(stage))
}

func (noteresourceshape *NoteResourceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteresourceshape.GongGetGongstructName(), noteresourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteresourceshape *NoteResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteresourceshape.GongGetGongstructName(), noteresourceshape.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

func (product *Product) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", product.GongGetGongstructName(), product.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (product *Product) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", product.GongGetGongstructName(), product.GongGetOrder(stage))
}

func (productcompositionshape *ProductCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productcompositionshape.GongGetGongstructName(), productcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productcompositionshape *ProductCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productcompositionshape.GongGetGongstructName(), productcompositionshape.GongGetOrder(stage))
}

func (productshape *ProductShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productshape.GongGetGongstructName(), productshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productshape *ProductShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productshape.GongGetGongstructName(), productshape.GongGetOrder(stage))
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcecompositionshape.GongGetGongstructName(), resourcecompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourcecompositionshape *ResourceCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcecompositionshape.GongGetGongstructName(), resourcecompositionshape.GongGetOrder(stage))
}

func (resourceshape *ResourceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourceshape.GongGetGongstructName(), resourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourceshape *ResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourceshape.GongGetGongstructName(), resourceshape.GongGetOrder(stage))
}

func (resourcetaskshape *ResourceTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcetaskshape.GongGetGongstructName(), resourcetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourcetaskshape *ResourceTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcetaskshape.GongGetGongstructName(), resourcetaskshape.GongGetOrder(stage))
}

func (task *Task) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (task *Task) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetOrder(stage))
}

func (taskcompositionshape *TaskCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskcompositionshape.GongGetGongstructName(), taskcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskcompositionshape *TaskCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskcompositionshape.GongGetGongstructName(), taskcompositionshape.GongGetOrder(stage))
}

func (taskgroup *TaskGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskgroup.GongGetGongstructName(), taskgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskgroup *TaskGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskgroup.GongGetGongstructName(), taskgroup.GongGetOrder(stage))
}

func (taskgroupshape *TaskGroupShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskgroupshape.GongGetGongstructName(), taskgroupshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskgroupshape *TaskGroupShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskgroupshape.GongGetGongstructName(), taskgroupshape.GongGetOrder(stage))
}

func (taskinputshape *TaskInputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskinputshape.GongGetGongstructName(), taskinputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskinputshape *TaskInputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskinputshape.GongGetGongstructName(), taskinputshape.GongGetOrder(stage))
}

func (taskoutputshape *TaskOutputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskoutputshape.GongGetGongstructName(), taskoutputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskoutputshape *TaskOutputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskoutputshape.GongGetGongstructName(), taskoutputshape.GongGetOrder(stage))
}

func (taskpredecessorshape *TaskPredecessorShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskpredecessorshape.GongGetGongstructName(), taskpredecessorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskpredecessorshape *TaskPredecessorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskpredecessorshape.GongGetGongstructName(), taskpredecessorshape.GongGetOrder(stage))
}

func (taskshape *TaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskshape *TaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagram.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note.Name))
	return
}

func (noteproductshape *NoteProductShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteProductShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteproductshape.Name))
	return
}

func (noteresourceshape *NoteResourceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteresourceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteResourceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteresourceshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteshape.Name))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notetaskshape.Name))
	return
}

func (product *Product) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", product.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Product")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(product.Name))
	return
}

func (productcompositionshape *ProductCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(productcompositionshape.Name))
	return
}

func (productshape *ProductShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(productshape.Name))
	return
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Resource")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(resource.Name))
	return
}

func (resourcecompositionshape *ResourceCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourcecompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ResourceCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(resourcecompositionshape.Name))
	return
}

func (resourceshape *ResourceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ResourceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(resourceshape.Name))
	return
}

func (resourcetaskshape *ResourceTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourcetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ResourceTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(resourcetaskshape.Name))
	return
}

func (task *Task) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Task")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(task.Name))
	return
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskcompositionshape.Name))
	return
}

func (taskgroup *TaskGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskgroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskgroup.Name))
	return
}

func (taskgroupshape *TaskGroupShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskgroupshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskGroupShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskgroupshape.Name))
	return
}

func (taskinputshape *TaskInputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskInputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskinputshape.Name))
	return
}

func (taskoutputshape *TaskOutputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskOutputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskoutputshape.Name))
	return
}

func (taskpredecessorshape *TaskPredecessorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskpredecessorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskPredecessorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskpredecessorshape.Name))
	return
}

func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskshape.Name))
	return
}

// insertion point for unstaging
func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
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

func (noteproductshape *NoteProductShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteproductshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteresourceshape *NoteResourceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteresourceshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetReferenceIdentifier(stage))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetReferenceIdentifier(stage))
	return
}

func (product *Product) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", product.GongGetReferenceIdentifier(stage))
	return
}

func (productcompositionshape *ProductCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productcompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (productshape *ProductShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productshape.GongGetReferenceIdentifier(stage))
	return
}

func (resource *Resource) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetReferenceIdentifier(stage))
	return
}

func (resourcecompositionshape *ResourceCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourcecompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (resourceshape *ResourceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourceshape.GongGetReferenceIdentifier(stage))
	return
}

func (resourcetaskshape *ResourceTaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourcetaskshape.GongGetReferenceIdentifier(stage))
	return
}

func (task *Task) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetReferenceIdentifier(stage))
	return
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskcompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (taskgroup *TaskGroup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskgroup.GongGetReferenceIdentifier(stage))
	return
}

func (taskgroupshape *TaskGroupShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskgroupshape.GongGetReferenceIdentifier(stage))
	return
}

func (taskinputshape *TaskInputShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskinputshape.GongGetReferenceIdentifier(stage))
	return
}

func (taskoutputshape *TaskOutputShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskoutputshape.GongGetReferenceIdentifier(stage))
	return
}

func (taskpredecessorshape *TaskPredecessorShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskpredecessorshape.GongGetReferenceIdentifier(stage))
	return
}

func (taskshape *TaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetReferenceIdentifier(stage))
	return
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

// end of template

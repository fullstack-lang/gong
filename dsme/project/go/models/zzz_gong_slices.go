// generated code - do not edit
package models

import (
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

	// Compute reverse map for named struct Project
	// insertion point per field
	stage.Project_RootProducts_reverseMap = make(map[*Product]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _product := range project.RootProducts {
			stage.Project_RootProducts_reverseMap[_product] = project
		}
	}
	stage.Project_RootTasks_reverseMap = make(map[*Task]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _task := range project.RootTasks {
			stage.Project_RootTasks_reverseMap[_task] = project
		}
	}
	stage.Project_RootResources_reverseMap = make(map[*Resource]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _resource := range project.RootResources {
			stage.Project_RootResources_reverseMap[_resource] = project
		}
	}
	stage.Project_Notes_reverseMap = make(map[*Note]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _note := range project.Notes {
			stage.Project_Notes_reverseMap[_note] = project
		}
	}
	stage.Project_Diagrams_reverseMap = make(map[*Diagram]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _diagram := range project.Diagrams {
			stage.Project_Diagrams_reverseMap[_diagram] = project
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

	// Compute reverse map for named struct ResourceCompositionShape
	// insertion point per field

	// Compute reverse map for named struct ResourceShape
	// insertion point per field

	// Compute reverse map for named struct ResourceTaskShape
	// insertion point per field

	// Compute reverse map for named struct Root
	// insertion point per field
	stage.Root_Projects_reverseMap = make(map[*Project]*Root)
	for root := range stage.Roots {
		_ = root
		for _, _project := range root.Projects {
			stage.Root_Projects_reverseMap[_project] = root
		}
	}

	// Compute reverse map for named struct Task
	// insertion point per field
	stage.Task_SubTasks_reverseMap = make(map[*Task]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _task := range task.SubTasks {
			stage.Task_SubTasks_reverseMap[_task] = task
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

	// Compute reverse map for named struct TaskCompositionShape
	// insertion point per field

	// Compute reverse map for named struct TaskInputShape
	// insertion point per field

	// Compute reverse map for named struct TaskOutputShape
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

	for instance := range stage.Projects {
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

	for instance := range stage.Roots {
		res = append(res, instance)
	}

	for instance := range stage.Tasks {
		res = append(res, instance)
	}

	for instance := range stage.TaskCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskInputShapes {
		res = append(res, instance)
	}

	for instance := range stage.TaskOutputShapes {
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
	diagram.CopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.CopyBasicFields(newInstance)
	return newInstance
}

func (noteproductshape *NoteProductShape) GongCopy() GongstructIF {
	newInstance := new(NoteProductShape)
	noteproductshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteresourceshape *NoteResourceShape) GongCopy() GongstructIF {
	newInstance := new(NoteResourceShape)
	noteresourceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.CopyBasicFields(newInstance)
	return newInstance
}

func (notetaskshape *NoteTaskShape) GongCopy() GongstructIF {
	newInstance := new(NoteTaskShape)
	notetaskshape.CopyBasicFields(newInstance)
	return newInstance
}

func (product *Product) GongCopy() GongstructIF {
	newInstance := new(Product)
	product.CopyBasicFields(newInstance)
	return newInstance
}

func (productcompositionshape *ProductCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ProductCompositionShape)
	productcompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (productshape *ProductShape) GongCopy() GongstructIF {
	newInstance := new(ProductShape)
	productshape.CopyBasicFields(newInstance)
	return newInstance
}

func (project *Project) GongCopy() GongstructIF {
	newInstance := new(Project)
	project.CopyBasicFields(newInstance)
	return newInstance
}

func (resource *Resource) GongCopy() GongstructIF {
	newInstance := new(Resource)
	resource.CopyBasicFields(newInstance)
	return newInstance
}

func (resourcecompositionshape *ResourceCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ResourceCompositionShape)
	resourcecompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (resourceshape *ResourceShape) GongCopy() GongstructIF {
	newInstance := new(ResourceShape)
	resourceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (resourcetaskshape *ResourceTaskShape) GongCopy() GongstructIF {
	newInstance := new(ResourceTaskShape)
	resourcetaskshape.CopyBasicFields(newInstance)
	return newInstance
}

func (root *Root) GongCopy() GongstructIF {
	newInstance := new(Root)
	root.CopyBasicFields(newInstance)
	return newInstance
}

func (task *Task) GongCopy() GongstructIF {
	newInstance := new(Task)
	task.CopyBasicFields(newInstance)
	return newInstance
}

func (taskcompositionshape *TaskCompositionShape) GongCopy() GongstructIF {
	newInstance := new(TaskCompositionShape)
	taskcompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (taskinputshape *TaskInputShape) GongCopy() GongstructIF {
	newInstance := new(TaskInputShape)
	taskinputshape.CopyBasicFields(newInstance)
	return newInstance
}

func (taskoutputshape *TaskOutputShape) GongCopy() GongstructIF {
	newInstance := new(TaskOutputShape)
	taskoutputshape.CopyBasicFields(newInstance)
	return newInstance
}

func (taskshape *TaskShape) GongCopy() GongstructIF {
	newInstance := new(TaskShape)
	taskshape.CopyBasicFields(newInstance)
	return newInstance
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
	var diagrams_newInstances []*Diagram
	var diagrams_deletedInstances []*Diagram

	// parse all staged instances and check if they have a reference
	for diagram := range stage.Diagrams {
		if ref, ok := stage.Diagrams_reference[diagram]; !ok {
			diagrams_newInstances = append(diagrams_newInstances, diagram)
			newInstancesSlice = append(newInstancesSlice, diagram.GongMarshallIdentifier(stage))
			if stage.Diagrams_referenceOrder == nil {
				stage.Diagrams_referenceOrder = make(map[*Diagram]uint)
			}
			stage.Diagrams_referenceOrder[diagram] = stage.Diagram_stagedOrder[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			// delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Diagram_stagedOrder[ref] = stage.Diagram_stagedOrder[diagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			// delete(stage.Diagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", diagram.GetName())
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
	for _, ref := range stage.Diagrams_reference {
		instance := stage.Diagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Diagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagrams_deletedInstances = append(diagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrams_newInstances)
	lenDeletedInstances += len(diagrams_deletedInstances)
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", note.GetName())
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
	var noteproductshapes_newInstances []*NoteProductShape
	var noteproductshapes_deletedInstances []*NoteProductShape

	// parse all staged instances and check if they have a reference
	for noteproductshape := range stage.NoteProductShapes {
		if ref, ok := stage.NoteProductShapes_reference[noteproductshape]; !ok {
			noteproductshapes_newInstances = append(noteproductshapes_newInstances, noteproductshape)
			newInstancesSlice = append(newInstancesSlice, noteproductshape.GongMarshallIdentifier(stage))
			if stage.NoteProductShapes_referenceOrder == nil {
				stage.NoteProductShapes_referenceOrder = make(map[*NoteProductShape]uint)
			}
			stage.NoteProductShapes_referenceOrder[noteproductshape] = stage.NoteProductShape_stagedOrder[noteproductshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteproductshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteProductShapes_referenceOrder, noteproductshape)
			fieldInitializers, pointersInitializations := noteproductshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteProductShape_stagedOrder[ref] = stage.NoteProductShape_stagedOrder[noteproductshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteproductshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteproductshape)
			// delete(stage.NoteProductShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", noteproductshape.GetName())
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
	for _, ref := range stage.NoteProductShapes_reference {
		instance := stage.NoteProductShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteProductShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteproductshapes_deletedInstances = append(noteproductshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteproductshapes_newInstances)
	lenDeletedInstances += len(noteproductshapes_deletedInstances)
	var noteresourceshapes_newInstances []*NoteResourceShape
	var noteresourceshapes_deletedInstances []*NoteResourceShape

	// parse all staged instances and check if they have a reference
	for noteresourceshape := range stage.NoteResourceShapes {
		if ref, ok := stage.NoteResourceShapes_reference[noteresourceshape]; !ok {
			noteresourceshapes_newInstances = append(noteresourceshapes_newInstances, noteresourceshape)
			newInstancesSlice = append(newInstancesSlice, noteresourceshape.GongMarshallIdentifier(stage))
			if stage.NoteResourceShapes_referenceOrder == nil {
				stage.NoteResourceShapes_referenceOrder = make(map[*NoteResourceShape]uint)
			}
			stage.NoteResourceShapes_referenceOrder[noteresourceshape] = stage.NoteResourceShape_stagedOrder[noteresourceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteresourceshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteResourceShapes_referenceOrder, noteresourceshape)
			fieldInitializers, pointersInitializations := noteresourceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteResourceShape_stagedOrder[ref] = stage.NoteResourceShape_stagedOrder[noteresourceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteresourceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteresourceshape)
			// delete(stage.NoteResourceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", noteresourceshape.GetName())
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
	for _, ref := range stage.NoteResourceShapes_reference {
		instance := stage.NoteResourceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteResourceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteresourceshapes_deletedInstances = append(noteresourceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteresourceshapes_newInstances)
	lenDeletedInstances += len(noteresourceshapes_deletedInstances)
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", noteshape.GetName())
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
	var notetaskshapes_newInstances []*NoteTaskShape
	var notetaskshapes_deletedInstances []*NoteTaskShape

	// parse all staged instances and check if they have a reference
	for notetaskshape := range stage.NoteTaskShapes {
		if ref, ok := stage.NoteTaskShapes_reference[notetaskshape]; !ok {
			notetaskshapes_newInstances = append(notetaskshapes_newInstances, notetaskshape)
			newInstancesSlice = append(newInstancesSlice, notetaskshape.GongMarshallIdentifier(stage))
			if stage.NoteTaskShapes_referenceOrder == nil {
				stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint)
			}
			stage.NoteTaskShapes_referenceOrder[notetaskshape] = stage.NoteTaskShape_stagedOrder[notetaskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, notetaskshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteTaskShapes_referenceOrder, notetaskshape)
			fieldInitializers, pointersInitializations := notetaskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteTaskShape_stagedOrder[ref] = stage.NoteTaskShape_stagedOrder[notetaskshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := notetaskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, notetaskshape)
			// delete(stage.NoteTaskShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", notetaskshape.GetName())
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
	for _, ref := range stage.NoteTaskShapes_reference {
		instance := stage.NoteTaskShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteTaskShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notetaskshapes_deletedInstances = append(notetaskshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notetaskshapes_newInstances)
	lenDeletedInstances += len(notetaskshapes_deletedInstances)
	var products_newInstances []*Product
	var products_deletedInstances []*Product

	// parse all staged instances and check if they have a reference
	for product := range stage.Products {
		if ref, ok := stage.Products_reference[product]; !ok {
			products_newInstances = append(products_newInstances, product)
			newInstancesSlice = append(newInstancesSlice, product.GongMarshallIdentifier(stage))
			if stage.Products_referenceOrder == nil {
				stage.Products_referenceOrder = make(map[*Product]uint)
			}
			stage.Products_referenceOrder[product] = stage.Product_stagedOrder[product]
			newInstancesReverseSlice = append(newInstancesReverseSlice, product.GongMarshallUnstaging(stage))
			// delete(stage.Products_referenceOrder, product)
			fieldInitializers, pointersInitializations := product.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Product_stagedOrder[ref] = stage.Product_stagedOrder[product]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := product.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, product)
			// delete(stage.Product_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", product.GetName())
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
	for _, ref := range stage.Products_reference {
		instance := stage.Products_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Products[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			products_deletedInstances = append(products_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(products_newInstances)
	lenDeletedInstances += len(products_deletedInstances)
	var productcompositionshapes_newInstances []*ProductCompositionShape
	var productcompositionshapes_deletedInstances []*ProductCompositionShape

	// parse all staged instances and check if they have a reference
	for productcompositionshape := range stage.ProductCompositionShapes {
		if ref, ok := stage.ProductCompositionShapes_reference[productcompositionshape]; !ok {
			productcompositionshapes_newInstances = append(productcompositionshapes_newInstances, productcompositionshape)
			newInstancesSlice = append(newInstancesSlice, productcompositionshape.GongMarshallIdentifier(stage))
			if stage.ProductCompositionShapes_referenceOrder == nil {
				stage.ProductCompositionShapes_referenceOrder = make(map[*ProductCompositionShape]uint)
			}
			stage.ProductCompositionShapes_referenceOrder[productcompositionshape] = stage.ProductCompositionShape_stagedOrder[productcompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, productcompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.ProductCompositionShapes_referenceOrder, productcompositionshape)
			fieldInitializers, pointersInitializations := productcompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProductCompositionShape_stagedOrder[ref] = stage.ProductCompositionShape_stagedOrder[productcompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := productcompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, productcompositionshape)
			// delete(stage.ProductCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", productcompositionshape.GetName())
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
	for _, ref := range stage.ProductCompositionShapes_reference {
		instance := stage.ProductCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ProductCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			productcompositionshapes_deletedInstances = append(productcompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(productcompositionshapes_newInstances)
	lenDeletedInstances += len(productcompositionshapes_deletedInstances)
	var productshapes_newInstances []*ProductShape
	var productshapes_deletedInstances []*ProductShape

	// parse all staged instances and check if they have a reference
	for productshape := range stage.ProductShapes {
		if ref, ok := stage.ProductShapes_reference[productshape]; !ok {
			productshapes_newInstances = append(productshapes_newInstances, productshape)
			newInstancesSlice = append(newInstancesSlice, productshape.GongMarshallIdentifier(stage))
			if stage.ProductShapes_referenceOrder == nil {
				stage.ProductShapes_referenceOrder = make(map[*ProductShape]uint)
			}
			stage.ProductShapes_referenceOrder[productshape] = stage.ProductShape_stagedOrder[productshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, productshape.GongMarshallUnstaging(stage))
			// delete(stage.ProductShapes_referenceOrder, productshape)
			fieldInitializers, pointersInitializations := productshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProductShape_stagedOrder[ref] = stage.ProductShape_stagedOrder[productshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := productshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, productshape)
			// delete(stage.ProductShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", productshape.GetName())
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
	for _, ref := range stage.ProductShapes_reference {
		instance := stage.ProductShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ProductShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			productshapes_deletedInstances = append(productshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(productshapes_newInstances)
	lenDeletedInstances += len(productshapes_deletedInstances)
	var projects_newInstances []*Project
	var projects_deletedInstances []*Project

	// parse all staged instances and check if they have a reference
	for project := range stage.Projects {
		if ref, ok := stage.Projects_reference[project]; !ok {
			projects_newInstances = append(projects_newInstances, project)
			newInstancesSlice = append(newInstancesSlice, project.GongMarshallIdentifier(stage))
			if stage.Projects_referenceOrder == nil {
				stage.Projects_referenceOrder = make(map[*Project]uint)
			}
			stage.Projects_referenceOrder[project] = stage.Project_stagedOrder[project]
			newInstancesReverseSlice = append(newInstancesReverseSlice, project.GongMarshallUnstaging(stage))
			// delete(stage.Projects_referenceOrder, project)
			fieldInitializers, pointersInitializations := project.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Project_stagedOrder[ref] = stage.Project_stagedOrder[project]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := project.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, project)
			// delete(stage.Project_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", project.GetName())
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
	for _, ref := range stage.Projects_reference {
		instance := stage.Projects_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Projects[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			projects_deletedInstances = append(projects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(projects_newInstances)
	lenDeletedInstances += len(projects_deletedInstances)
	var resources_newInstances []*Resource
	var resources_deletedInstances []*Resource

	// parse all staged instances and check if they have a reference
	for resource := range stage.Resources {
		if ref, ok := stage.Resources_reference[resource]; !ok {
			resources_newInstances = append(resources_newInstances, resource)
			newInstancesSlice = append(newInstancesSlice, resource.GongMarshallIdentifier(stage))
			if stage.Resources_referenceOrder == nil {
				stage.Resources_referenceOrder = make(map[*Resource]uint)
			}
			stage.Resources_referenceOrder[resource] = stage.Resource_stagedOrder[resource]
			newInstancesReverseSlice = append(newInstancesReverseSlice, resource.GongMarshallUnstaging(stage))
			// delete(stage.Resources_referenceOrder, resource)
			fieldInitializers, pointersInitializations := resource.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Resource_stagedOrder[ref] = stage.Resource_stagedOrder[resource]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := resource.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, resource)
			// delete(stage.Resource_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", resource.GetName())
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
	for _, ref := range stage.Resources_reference {
		instance := stage.Resources_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Resources[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			resources_deletedInstances = append(resources_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(resources_newInstances)
	lenDeletedInstances += len(resources_deletedInstances)
	var resourcecompositionshapes_newInstances []*ResourceCompositionShape
	var resourcecompositionshapes_deletedInstances []*ResourceCompositionShape

	// parse all staged instances and check if they have a reference
	for resourcecompositionshape := range stage.ResourceCompositionShapes {
		if ref, ok := stage.ResourceCompositionShapes_reference[resourcecompositionshape]; !ok {
			resourcecompositionshapes_newInstances = append(resourcecompositionshapes_newInstances, resourcecompositionshape)
			newInstancesSlice = append(newInstancesSlice, resourcecompositionshape.GongMarshallIdentifier(stage))
			if stage.ResourceCompositionShapes_referenceOrder == nil {
				stage.ResourceCompositionShapes_referenceOrder = make(map[*ResourceCompositionShape]uint)
			}
			stage.ResourceCompositionShapes_referenceOrder[resourcecompositionshape] = stage.ResourceCompositionShape_stagedOrder[resourcecompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, resourcecompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.ResourceCompositionShapes_referenceOrder, resourcecompositionshape)
			fieldInitializers, pointersInitializations := resourcecompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ResourceCompositionShape_stagedOrder[ref] = stage.ResourceCompositionShape_stagedOrder[resourcecompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := resourcecompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, resourcecompositionshape)
			// delete(stage.ResourceCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", resourcecompositionshape.GetName())
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
	for _, ref := range stage.ResourceCompositionShapes_reference {
		instance := stage.ResourceCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ResourceCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			resourcecompositionshapes_deletedInstances = append(resourcecompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(resourcecompositionshapes_newInstances)
	lenDeletedInstances += len(resourcecompositionshapes_deletedInstances)
	var resourceshapes_newInstances []*ResourceShape
	var resourceshapes_deletedInstances []*ResourceShape

	// parse all staged instances and check if they have a reference
	for resourceshape := range stage.ResourceShapes {
		if ref, ok := stage.ResourceShapes_reference[resourceshape]; !ok {
			resourceshapes_newInstances = append(resourceshapes_newInstances, resourceshape)
			newInstancesSlice = append(newInstancesSlice, resourceshape.GongMarshallIdentifier(stage))
			if stage.ResourceShapes_referenceOrder == nil {
				stage.ResourceShapes_referenceOrder = make(map[*ResourceShape]uint)
			}
			stage.ResourceShapes_referenceOrder[resourceshape] = stage.ResourceShape_stagedOrder[resourceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, resourceshape.GongMarshallUnstaging(stage))
			// delete(stage.ResourceShapes_referenceOrder, resourceshape)
			fieldInitializers, pointersInitializations := resourceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ResourceShape_stagedOrder[ref] = stage.ResourceShape_stagedOrder[resourceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := resourceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, resourceshape)
			// delete(stage.ResourceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", resourceshape.GetName())
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
	for _, ref := range stage.ResourceShapes_reference {
		instance := stage.ResourceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ResourceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			resourceshapes_deletedInstances = append(resourceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(resourceshapes_newInstances)
	lenDeletedInstances += len(resourceshapes_deletedInstances)
	var resourcetaskshapes_newInstances []*ResourceTaskShape
	var resourcetaskshapes_deletedInstances []*ResourceTaskShape

	// parse all staged instances and check if they have a reference
	for resourcetaskshape := range stage.ResourceTaskShapes {
		if ref, ok := stage.ResourceTaskShapes_reference[resourcetaskshape]; !ok {
			resourcetaskshapes_newInstances = append(resourcetaskshapes_newInstances, resourcetaskshape)
			newInstancesSlice = append(newInstancesSlice, resourcetaskshape.GongMarshallIdentifier(stage))
			if stage.ResourceTaskShapes_referenceOrder == nil {
				stage.ResourceTaskShapes_referenceOrder = make(map[*ResourceTaskShape]uint)
			}
			stage.ResourceTaskShapes_referenceOrder[resourcetaskshape] = stage.ResourceTaskShape_stagedOrder[resourcetaskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, resourcetaskshape.GongMarshallUnstaging(stage))
			// delete(stage.ResourceTaskShapes_referenceOrder, resourcetaskshape)
			fieldInitializers, pointersInitializations := resourcetaskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ResourceTaskShape_stagedOrder[ref] = stage.ResourceTaskShape_stagedOrder[resourcetaskshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := resourcetaskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, resourcetaskshape)
			// delete(stage.ResourceTaskShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", resourcetaskshape.GetName())
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
	for _, ref := range stage.ResourceTaskShapes_reference {
		instance := stage.ResourceTaskShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ResourceTaskShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			resourcetaskshapes_deletedInstances = append(resourcetaskshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(resourcetaskshapes_newInstances)
	lenDeletedInstances += len(resourcetaskshapes_deletedInstances)
	var roots_newInstances []*Root
	var roots_deletedInstances []*Root

	// parse all staged instances and check if they have a reference
	for root := range stage.Roots {
		if ref, ok := stage.Roots_reference[root]; !ok {
			roots_newInstances = append(roots_newInstances, root)
			newInstancesSlice = append(newInstancesSlice, root.GongMarshallIdentifier(stage))
			if stage.Roots_referenceOrder == nil {
				stage.Roots_referenceOrder = make(map[*Root]uint)
			}
			stage.Roots_referenceOrder[root] = stage.Root_stagedOrder[root]
			newInstancesReverseSlice = append(newInstancesReverseSlice, root.GongMarshallUnstaging(stage))
			// delete(stage.Roots_referenceOrder, root)
			fieldInitializers, pointersInitializations := root.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Root_stagedOrder[ref] = stage.Root_stagedOrder[root]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := root.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, root)
			// delete(stage.Root_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", root.GetName())
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
	for _, ref := range stage.Roots_reference {
		instance := stage.Roots_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Roots[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			roots_deletedInstances = append(roots_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(roots_newInstances)
	lenDeletedInstances += len(roots_deletedInstances)
	var tasks_newInstances []*Task
	var tasks_deletedInstances []*Task

	// parse all staged instances and check if they have a reference
	for task := range stage.Tasks {
		if ref, ok := stage.Tasks_reference[task]; !ok {
			tasks_newInstances = append(tasks_newInstances, task)
			newInstancesSlice = append(newInstancesSlice, task.GongMarshallIdentifier(stage))
			if stage.Tasks_referenceOrder == nil {
				stage.Tasks_referenceOrder = make(map[*Task]uint)
			}
			stage.Tasks_referenceOrder[task] = stage.Task_stagedOrder[task]
			newInstancesReverseSlice = append(newInstancesReverseSlice, task.GongMarshallUnstaging(stage))
			// delete(stage.Tasks_referenceOrder, task)
			fieldInitializers, pointersInitializations := task.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Task_stagedOrder[ref] = stage.Task_stagedOrder[task]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := task.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, task)
			// delete(stage.Task_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", task.GetName())
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
	for _, ref := range stage.Tasks_reference {
		instance := stage.Tasks_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Tasks[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			tasks_deletedInstances = append(tasks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tasks_newInstances)
	lenDeletedInstances += len(tasks_deletedInstances)
	var taskcompositionshapes_newInstances []*TaskCompositionShape
	var taskcompositionshapes_deletedInstances []*TaskCompositionShape

	// parse all staged instances and check if they have a reference
	for taskcompositionshape := range stage.TaskCompositionShapes {
		if ref, ok := stage.TaskCompositionShapes_reference[taskcompositionshape]; !ok {
			taskcompositionshapes_newInstances = append(taskcompositionshapes_newInstances, taskcompositionshape)
			newInstancesSlice = append(newInstancesSlice, taskcompositionshape.GongMarshallIdentifier(stage))
			if stage.TaskCompositionShapes_referenceOrder == nil {
				stage.TaskCompositionShapes_referenceOrder = make(map[*TaskCompositionShape]uint)
			}
			stage.TaskCompositionShapes_referenceOrder[taskcompositionshape] = stage.TaskCompositionShape_stagedOrder[taskcompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskcompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.TaskCompositionShapes_referenceOrder, taskcompositionshape)
			fieldInitializers, pointersInitializations := taskcompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskCompositionShape_stagedOrder[ref] = stage.TaskCompositionShape_stagedOrder[taskcompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := taskcompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskcompositionshape)
			// delete(stage.TaskCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", taskcompositionshape.GetName())
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
	for _, ref := range stage.TaskCompositionShapes_reference {
		instance := stage.TaskCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TaskCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			taskcompositionshapes_deletedInstances = append(taskcompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(taskcompositionshapes_newInstances)
	lenDeletedInstances += len(taskcompositionshapes_deletedInstances)
	var taskinputshapes_newInstances []*TaskInputShape
	var taskinputshapes_deletedInstances []*TaskInputShape

	// parse all staged instances and check if they have a reference
	for taskinputshape := range stage.TaskInputShapes {
		if ref, ok := stage.TaskInputShapes_reference[taskinputshape]; !ok {
			taskinputshapes_newInstances = append(taskinputshapes_newInstances, taskinputshape)
			newInstancesSlice = append(newInstancesSlice, taskinputshape.GongMarshallIdentifier(stage))
			if stage.TaskInputShapes_referenceOrder == nil {
				stage.TaskInputShapes_referenceOrder = make(map[*TaskInputShape]uint)
			}
			stage.TaskInputShapes_referenceOrder[taskinputshape] = stage.TaskInputShape_stagedOrder[taskinputshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskinputshape.GongMarshallUnstaging(stage))
			// delete(stage.TaskInputShapes_referenceOrder, taskinputshape)
			fieldInitializers, pointersInitializations := taskinputshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskInputShape_stagedOrder[ref] = stage.TaskInputShape_stagedOrder[taskinputshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := taskinputshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskinputshape)
			// delete(stage.TaskInputShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", taskinputshape.GetName())
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
	for _, ref := range stage.TaskInputShapes_reference {
		instance := stage.TaskInputShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TaskInputShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			taskinputshapes_deletedInstances = append(taskinputshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(taskinputshapes_newInstances)
	lenDeletedInstances += len(taskinputshapes_deletedInstances)
	var taskoutputshapes_newInstances []*TaskOutputShape
	var taskoutputshapes_deletedInstances []*TaskOutputShape

	// parse all staged instances and check if they have a reference
	for taskoutputshape := range stage.TaskOutputShapes {
		if ref, ok := stage.TaskOutputShapes_reference[taskoutputshape]; !ok {
			taskoutputshapes_newInstances = append(taskoutputshapes_newInstances, taskoutputshape)
			newInstancesSlice = append(newInstancesSlice, taskoutputshape.GongMarshallIdentifier(stage))
			if stage.TaskOutputShapes_referenceOrder == nil {
				stage.TaskOutputShapes_referenceOrder = make(map[*TaskOutputShape]uint)
			}
			stage.TaskOutputShapes_referenceOrder[taskoutputshape] = stage.TaskOutputShape_stagedOrder[taskoutputshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskoutputshape.GongMarshallUnstaging(stage))
			// delete(stage.TaskOutputShapes_referenceOrder, taskoutputshape)
			fieldInitializers, pointersInitializations := taskoutputshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskOutputShape_stagedOrder[ref] = stage.TaskOutputShape_stagedOrder[taskoutputshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := taskoutputshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskoutputshape)
			// delete(stage.TaskOutputShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", taskoutputshape.GetName())
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
	for _, ref := range stage.TaskOutputShapes_reference {
		instance := stage.TaskOutputShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TaskOutputShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			taskoutputshapes_deletedInstances = append(taskoutputshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(taskoutputshapes_newInstances)
	lenDeletedInstances += len(taskoutputshapes_deletedInstances)
	var taskshapes_newInstances []*TaskShape
	var taskshapes_deletedInstances []*TaskShape

	// parse all staged instances and check if they have a reference
	for taskshape := range stage.TaskShapes {
		if ref, ok := stage.TaskShapes_reference[taskshape]; !ok {
			taskshapes_newInstances = append(taskshapes_newInstances, taskshape)
			newInstancesSlice = append(newInstancesSlice, taskshape.GongMarshallIdentifier(stage))
			if stage.TaskShapes_referenceOrder == nil {
				stage.TaskShapes_referenceOrder = make(map[*TaskShape]uint)
			}
			stage.TaskShapes_referenceOrder[taskshape] = stage.TaskShape_stagedOrder[taskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskshape.GongMarshallUnstaging(stage))
			// delete(stage.TaskShapes_referenceOrder, taskshape)
			fieldInitializers, pointersInitializations := taskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskShape_stagedOrder[ref] = stage.TaskShape_stagedOrder[taskshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := taskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskshape)
			// delete(stage.TaskShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", taskshape.GetName())
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
	for _, ref := range stage.TaskShapes_reference {
		instance := stage.TaskShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TaskShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			taskshapes_deletedInstances = append(taskshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(taskshapes_newInstances)
	lenDeletedInstances += len(taskshapes_deletedInstances)

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

	stage.Projects_reference = make(map[*Project]*Project)
	stage.Projects_referenceOrder = make(map[*Project]uint) // diff Unstage needs the reference order
	stage.Projects_instance = make(map[*Project]*Project)
	for instance := range stage.Projects {
		_copy := instance.GongCopy().(*Project)
		stage.Projects_reference[instance] = _copy
		stage.Projects_instance[_copy] = instance
		stage.Projects_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Roots_reference = make(map[*Root]*Root)
	stage.Roots_referenceOrder = make(map[*Root]uint) // diff Unstage needs the reference order
	stage.Roots_instance = make(map[*Root]*Root)
	for instance := range stage.Roots {
		_copy := instance.GongCopy().(*Root)
		stage.Roots_reference[instance] = _copy
		stage.Roots_instance[_copy] = instance
		stage.Roots_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	for instance := range stage.Projects {
		reference := stage.Projects_reference[instance]
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

	for instance := range stage.Roots {
		reference := stage.Roots_reference[instance]
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

	for instance := range stage.TaskInputShapes {
		reference := stage.TaskInputShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskOutputShapes {
		reference := stage.TaskOutputShapes_reference[instance]
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

func (diagram *Diagram) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Diagrams_referenceOrder[diagram]
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

func (note *Note) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Notes_referenceOrder[note]
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

func (noteproductshape *NoteProductShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteProductShapes_referenceOrder[noteproductshape]
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

func (noteresourceshape *NoteResourceShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteResourceShapes_referenceOrder[noteresourceshape]
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

func (noteshape *NoteShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteShapes_referenceOrder[noteshape]
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

func (notetaskshape *NoteTaskShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteTaskShapes_referenceOrder[notetaskshape]
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

func (product *Product) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Products_referenceOrder[product]
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

func (productcompositionshape *ProductCompositionShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ProductCompositionShapes_referenceOrder[productcompositionshape]
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

func (productshape *ProductShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ProductShapes_referenceOrder[productshape]
}

func (project *Project) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Project_stagedOrder[project]; ok {
		return order
	}
	if order, ok := stage.Projects_referenceOrder[project]; ok {
		return order
	} else {
		log.Printf("instance %p of type Project was not staged and does not have a reference order", project)
		return 0
	}
}

func (project *Project) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Projects_referenceOrder[project]
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

func (resource *Resource) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Resources_referenceOrder[resource]
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

func (resourcecompositionshape *ResourceCompositionShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ResourceCompositionShapes_referenceOrder[resourcecompositionshape]
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

func (resourceshape *ResourceShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ResourceShapes_referenceOrder[resourceshape]
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

func (resourcetaskshape *ResourceTaskShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ResourceTaskShapes_referenceOrder[resourcetaskshape]
}

func (root *Root) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Root_stagedOrder[root]; ok {
		return order
	}
	if order, ok := stage.Roots_referenceOrder[root]; ok {
		return order
	} else {
		log.Printf("instance %p of type Root was not staged and does not have a reference order", root)
		return 0
	}
}

func (root *Root) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Roots_referenceOrder[root]
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

func (task *Task) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Tasks_referenceOrder[task]
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

func (taskcompositionshape *TaskCompositionShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskCompositionShapes_referenceOrder[taskcompositionshape]
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

func (taskinputshape *TaskInputShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskInputShapes_referenceOrder[taskinputshape]
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

func (taskoutputshape *TaskOutputShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskOutputShapes_referenceOrder[taskoutputshape]
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

func (taskshape *TaskShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskShapes_referenceOrder[taskshape]
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
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetReferenceOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetReferenceOrder(stage))
}

func (noteproductshape *NoteProductShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteproductshape.GongGetGongstructName(), noteproductshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteproductshape *NoteProductShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteproductshape.GongGetGongstructName(), noteproductshape.GongGetReferenceOrder(stage))
}

func (noteresourceshape *NoteResourceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteresourceshape.GongGetGongstructName(), noteresourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteresourceshape *NoteResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteresourceshape.GongGetGongstructName(), noteresourceshape.GongGetReferenceOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetReferenceOrder(stage))
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetReferenceOrder(stage))
}

func (product *Product) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", product.GongGetGongstructName(), product.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (product *Product) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", product.GongGetGongstructName(), product.GongGetReferenceOrder(stage))
}

func (productcompositionshape *ProductCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productcompositionshape.GongGetGongstructName(), productcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productcompositionshape *ProductCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productcompositionshape.GongGetGongstructName(), productcompositionshape.GongGetReferenceOrder(stage))
}

func (productshape *ProductShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productshape.GongGetGongstructName(), productshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (productshape *ProductShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productshape.GongGetGongstructName(), productshape.GongGetReferenceOrder(stage))
}

func (project *Project) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", project.GongGetGongstructName(), project.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (project *Project) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", project.GongGetGongstructName(), project.GongGetReferenceOrder(stage))
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetReferenceOrder(stage))
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcecompositionshape.GongGetGongstructName(), resourcecompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourcecompositionshape *ResourceCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcecompositionshape.GongGetGongstructName(), resourcecompositionshape.GongGetReferenceOrder(stage))
}

func (resourceshape *ResourceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourceshape.GongGetGongstructName(), resourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourceshape *ResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourceshape.GongGetGongstructName(), resourceshape.GongGetReferenceOrder(stage))
}

func (resourcetaskshape *ResourceTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcetaskshape.GongGetGongstructName(), resourcetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resourcetaskshape *ResourceTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resourcetaskshape.GongGetGongstructName(), resourcetaskshape.GongGetReferenceOrder(stage))
}

func (root *Root) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root.GongGetGongstructName(), root.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (root *Root) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root.GongGetGongstructName(), root.GongGetReferenceOrder(stage))
}

func (task *Task) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (task *Task) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetReferenceOrder(stage))
}

func (taskcompositionshape *TaskCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskcompositionshape.GongGetGongstructName(), taskcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskcompositionshape *TaskCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskcompositionshape.GongGetGongstructName(), taskcompositionshape.GongGetReferenceOrder(stage))
}

func (taskinputshape *TaskInputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskinputshape.GongGetGongstructName(), taskinputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskinputshape *TaskInputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskinputshape.GongGetGongstructName(), taskinputshape.GongGetReferenceOrder(stage))
}

func (taskoutputshape *TaskOutputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskoutputshape.GongGetGongstructName(), taskoutputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskoutputshape *TaskOutputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskoutputshape.GongGetGongstructName(), taskoutputshape.GongGetReferenceOrder(stage))
}

func (taskshape *TaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskshape *TaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.Name))
	return
}

func (noteproductshape *NoteProductShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteProductShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteproductshape.Name))
	return
}

func (noteresourceshape *NoteResourceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteresourceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteResourceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteresourceshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteshape.Name))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(notetaskshape.Name))
	return
}

func (product *Product) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", product.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Product")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(product.Name))
	return
}

func (productcompositionshape *ProductCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(productcompositionshape.Name))
	return
}

func (productshape *ProductShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(productshape.Name))
	return
}

func (project *Project) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", project.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Project")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(project.Name))
	return
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Resource")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.Name))
	return
}

func (resourcecompositionshape *ResourceCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourcecompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ResourceCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resourcecompositionshape.Name))
	return
}

func (resourceshape *ResourceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ResourceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resourceshape.Name))
	return
}

func (resourcetaskshape *ResourceTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resourcetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ResourceTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resourcetaskshape.Name))
	return
}

func (root *Root) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(root.Name))
	return
}

func (task *Task) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Task")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(task.Name))
	return
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(taskcompositionshape.Name))
	return
}

func (taskinputshape *TaskInputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskInputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(taskinputshape.Name))
	return
}

func (taskoutputshape *TaskOutputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskOutputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(taskoutputshape.Name))
	return
}

func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(taskshape.Name))
	return
}

// insertion point for unstaging
func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
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

func (project *Project) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", project.GongGetReferenceIdentifier(stage))
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

func (root *Root) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root.GongGetReferenceIdentifier(stage))
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

func (taskshape *TaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetReferenceIdentifier(stage))
	return
}

// end of template

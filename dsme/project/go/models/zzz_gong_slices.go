// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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

	// Compute reverse map for named struct NoteProductShape
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
	newInstance := *diagram
	return &newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := *note
	return &newInstance
}

func (noteproductshape *NoteProductShape) GongCopy() GongstructIF {
	newInstance := *noteproductshape
	return &newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := *noteshape
	return &newInstance
}

func (notetaskshape *NoteTaskShape) GongCopy() GongstructIF {
	newInstance := *notetaskshape
	return &newInstance
}

func (product *Product) GongCopy() GongstructIF {
	newInstance := *product
	return &newInstance
}

func (productcompositionshape *ProductCompositionShape) GongCopy() GongstructIF {
	newInstance := *productcompositionshape
	return &newInstance
}

func (productshape *ProductShape) GongCopy() GongstructIF {
	newInstance := *productshape
	return &newInstance
}

func (project *Project) GongCopy() GongstructIF {
	newInstance := *project
	return &newInstance
}

func (root *Root) GongCopy() GongstructIF {
	newInstance := *root
	return &newInstance
}

func (task *Task) GongCopy() GongstructIF {
	newInstance := *task
	return &newInstance
}

func (taskcompositionshape *TaskCompositionShape) GongCopy() GongstructIF {
	newInstance := *taskcompositionshape
	return &newInstance
}

func (taskinputshape *TaskInputShape) GongCopy() GongstructIF {
	newInstance := *taskinputshape
	return &newInstance
}

func (taskoutputshape *TaskOutputShape) GongCopy() GongstructIF {
	newInstance := *taskoutputshape
	return &newInstance
}

func (taskshape *TaskShape) GongCopy() GongstructIF {
	newInstance := *taskshape
	return &newInstance
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
			stage.Diagrams_referenceOrder[diagram] = stage.DiagramMap_Staged_Order[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramMap_Staged_Order[ref] = stage.DiagramMap_Staged_Order[diagram]
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			delete(stage.DiagramMap_Staged_Order, ref)
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
	for ref := range stage.Diagrams_reference {
		if _, ok := stage.Diagrams[ref]; !ok {
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
			stage.Notes_referenceOrder[note] = stage.NoteMap_Staged_Order[note]
			newInstancesReverseSlice = append(newInstancesReverseSlice, note.GongMarshallUnstaging(stage))
			delete(stage.Notes_referenceOrder, note)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteMap_Staged_Order[ref] = stage.NoteMap_Staged_Order[note]
			diffs := note.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, note)
			delete(stage.NoteMap_Staged_Order, ref)
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
	for ref := range stage.Notes_reference {
		if _, ok := stage.Notes[ref]; !ok {
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
			stage.NoteProductShapes_referenceOrder[noteproductshape] = stage.NoteProductShapeMap_Staged_Order[noteproductshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteproductshape.GongMarshallUnstaging(stage))
			delete(stage.NoteProductShapes_referenceOrder, noteproductshape)
			fieldInitializers, pointersInitializations := noteproductshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteProductShapeMap_Staged_Order[ref] = stage.NoteProductShapeMap_Staged_Order[noteproductshape]
			diffs := noteproductshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteproductshape)
			delete(stage.NoteProductShapeMap_Staged_Order, ref)
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
	for ref := range stage.NoteProductShapes_reference {
		if _, ok := stage.NoteProductShapes[ref]; !ok {
			noteproductshapes_deletedInstances = append(noteproductshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteproductshapes_newInstances)
	lenDeletedInstances += len(noteproductshapes_deletedInstances)
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
			stage.NoteShapes_referenceOrder[noteshape] = stage.NoteShapeMap_Staged_Order[noteshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteshape.GongMarshallUnstaging(stage))
			delete(stage.NoteShapes_referenceOrder, noteshape)
			fieldInitializers, pointersInitializations := noteshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteShapeMap_Staged_Order[ref] = stage.NoteShapeMap_Staged_Order[noteshape]
			diffs := noteshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteshape)
			delete(stage.NoteShapeMap_Staged_Order, ref)
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
	for ref := range stage.NoteShapes_reference {
		if _, ok := stage.NoteShapes[ref]; !ok {
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
			stage.NoteTaskShapes_referenceOrder[notetaskshape] = stage.NoteTaskShapeMap_Staged_Order[notetaskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, notetaskshape.GongMarshallUnstaging(stage))
			delete(stage.NoteTaskShapes_referenceOrder, notetaskshape)
			fieldInitializers, pointersInitializations := notetaskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteTaskShapeMap_Staged_Order[ref] = stage.NoteTaskShapeMap_Staged_Order[notetaskshape]
			diffs := notetaskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, notetaskshape)
			delete(stage.NoteTaskShapeMap_Staged_Order, ref)
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
	for ref := range stage.NoteTaskShapes_reference {
		if _, ok := stage.NoteTaskShapes[ref]; !ok {
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
			stage.Products_referenceOrder[product] = stage.ProductMap_Staged_Order[product]
			newInstancesReverseSlice = append(newInstancesReverseSlice, product.GongMarshallUnstaging(stage))
			delete(stage.Products_referenceOrder, product)
			fieldInitializers, pointersInitializations := product.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProductMap_Staged_Order[ref] = stage.ProductMap_Staged_Order[product]
			diffs := product.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, product)
			delete(stage.ProductMap_Staged_Order, ref)
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
	for ref := range stage.Products_reference {
		if _, ok := stage.Products[ref]; !ok {
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
			stage.ProductCompositionShapes_referenceOrder[productcompositionshape] = stage.ProductCompositionShapeMap_Staged_Order[productcompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, productcompositionshape.GongMarshallUnstaging(stage))
			delete(stage.ProductCompositionShapes_referenceOrder, productcompositionshape)
			fieldInitializers, pointersInitializations := productcompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProductCompositionShapeMap_Staged_Order[ref] = stage.ProductCompositionShapeMap_Staged_Order[productcompositionshape]
			diffs := productcompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, productcompositionshape)
			delete(stage.ProductCompositionShapeMap_Staged_Order, ref)
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
	for ref := range stage.ProductCompositionShapes_reference {
		if _, ok := stage.ProductCompositionShapes[ref]; !ok {
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
			stage.ProductShapes_referenceOrder[productshape] = stage.ProductShapeMap_Staged_Order[productshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, productshape.GongMarshallUnstaging(stage))
			delete(stage.ProductShapes_referenceOrder, productshape)
			fieldInitializers, pointersInitializations := productshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProductShapeMap_Staged_Order[ref] = stage.ProductShapeMap_Staged_Order[productshape]
			diffs := productshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, productshape)
			delete(stage.ProductShapeMap_Staged_Order, ref)
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
	for ref := range stage.ProductShapes_reference {
		if _, ok := stage.ProductShapes[ref]; !ok {
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
			stage.Projects_referenceOrder[project] = stage.ProjectMap_Staged_Order[project]
			newInstancesReverseSlice = append(newInstancesReverseSlice, project.GongMarshallUnstaging(stage))
			delete(stage.Projects_referenceOrder, project)
			fieldInitializers, pointersInitializations := project.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProjectMap_Staged_Order[ref] = stage.ProjectMap_Staged_Order[project]
			diffs := project.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, project)
			delete(stage.ProjectMap_Staged_Order, ref)
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
	for ref := range stage.Projects_reference {
		if _, ok := stage.Projects[ref]; !ok {
			projects_deletedInstances = append(projects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(projects_newInstances)
	lenDeletedInstances += len(projects_deletedInstances)
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
			stage.Roots_referenceOrder[root] = stage.RootMap_Staged_Order[root]
			newInstancesReverseSlice = append(newInstancesReverseSlice, root.GongMarshallUnstaging(stage))
			delete(stage.Roots_referenceOrder, root)
			fieldInitializers, pointersInitializations := root.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RootMap_Staged_Order[ref] = stage.RootMap_Staged_Order[root]
			diffs := root.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, root)
			delete(stage.RootMap_Staged_Order, ref)
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
	for ref := range stage.Roots_reference {
		if _, ok := stage.Roots[ref]; !ok {
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
			stage.Tasks_referenceOrder[task] = stage.TaskMap_Staged_Order[task]
			newInstancesReverseSlice = append(newInstancesReverseSlice, task.GongMarshallUnstaging(stage))
			delete(stage.Tasks_referenceOrder, task)
			fieldInitializers, pointersInitializations := task.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskMap_Staged_Order[ref] = stage.TaskMap_Staged_Order[task]
			diffs := task.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, task)
			delete(stage.TaskMap_Staged_Order, ref)
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
	for ref := range stage.Tasks_reference {
		if _, ok := stage.Tasks[ref]; !ok {
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
			stage.TaskCompositionShapes_referenceOrder[taskcompositionshape] = stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskcompositionshape.GongMarshallUnstaging(stage))
			delete(stage.TaskCompositionShapes_referenceOrder, taskcompositionshape)
			fieldInitializers, pointersInitializations := taskcompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskCompositionShapeMap_Staged_Order[ref] = stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshape]
			diffs := taskcompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskcompositionshape)
			delete(stage.TaskCompositionShapeMap_Staged_Order, ref)
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
	for ref := range stage.TaskCompositionShapes_reference {
		if _, ok := stage.TaskCompositionShapes[ref]; !ok {
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
			stage.TaskInputShapes_referenceOrder[taskinputshape] = stage.TaskInputShapeMap_Staged_Order[taskinputshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskinputshape.GongMarshallUnstaging(stage))
			delete(stage.TaskInputShapes_referenceOrder, taskinputshape)
			fieldInitializers, pointersInitializations := taskinputshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskInputShapeMap_Staged_Order[ref] = stage.TaskInputShapeMap_Staged_Order[taskinputshape]
			diffs := taskinputshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskinputshape)
			delete(stage.TaskInputShapeMap_Staged_Order, ref)
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
	for ref := range stage.TaskInputShapes_reference {
		if _, ok := stage.TaskInputShapes[ref]; !ok {
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
			stage.TaskOutputShapes_referenceOrder[taskoutputshape] = stage.TaskOutputShapeMap_Staged_Order[taskoutputshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskoutputshape.GongMarshallUnstaging(stage))
			delete(stage.TaskOutputShapes_referenceOrder, taskoutputshape)
			fieldInitializers, pointersInitializations := taskoutputshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskOutputShapeMap_Staged_Order[ref] = stage.TaskOutputShapeMap_Staged_Order[taskoutputshape]
			diffs := taskoutputshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskoutputshape)
			delete(stage.TaskOutputShapeMap_Staged_Order, ref)
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
	for ref := range stage.TaskOutputShapes_reference {
		if _, ok := stage.TaskOutputShapes[ref]; !ok {
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
			stage.TaskShapes_referenceOrder[taskshape] = stage.TaskShapeMap_Staged_Order[taskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskshape.GongMarshallUnstaging(stage))
			delete(stage.TaskShapes_referenceOrder, taskshape)
			fieldInitializers, pointersInitializations := taskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskShapeMap_Staged_Order[ref] = stage.TaskShapeMap_Staged_Order[taskshape]
			diffs := taskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskshape)
			delete(stage.TaskShapeMap_Staged_Order, ref)
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
	for ref := range stage.TaskShapes_reference {
		if _, ok := stage.TaskShapes[ref]; !ok {
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

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	for instance := range stage.Diagrams {
		stage.Diagrams_reference[instance] = instance.GongCopy().(*Diagram)
		stage.Diagrams_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	for instance := range stage.Notes {
		stage.Notes_reference[instance] = instance.GongCopy().(*Note)
		stage.Notes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.NoteProductShapes_reference = make(map[*NoteProductShape]*NoteProductShape)
	stage.NoteProductShapes_referenceOrder = make(map[*NoteProductShape]uint) // diff Unstage needs the reference order
	for instance := range stage.NoteProductShapes {
		stage.NoteProductShapes_reference[instance] = instance.GongCopy().(*NoteProductShape)
		stage.NoteProductShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint) // diff Unstage needs the reference order
	for instance := range stage.NoteShapes {
		stage.NoteShapes_reference[instance] = instance.GongCopy().(*NoteShape)
		stage.NoteShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint) // diff Unstage needs the reference order
	for instance := range stage.NoteTaskShapes {
		stage.NoteTaskShapes_reference[instance] = instance.GongCopy().(*NoteTaskShape)
		stage.NoteTaskShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Products_reference = make(map[*Product]*Product)
	stage.Products_referenceOrder = make(map[*Product]uint) // diff Unstage needs the reference order
	for instance := range stage.Products {
		stage.Products_reference[instance] = instance.GongCopy().(*Product)
		stage.Products_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ProductCompositionShapes_reference = make(map[*ProductCompositionShape]*ProductCompositionShape)
	stage.ProductCompositionShapes_referenceOrder = make(map[*ProductCompositionShape]uint) // diff Unstage needs the reference order
	for instance := range stage.ProductCompositionShapes {
		stage.ProductCompositionShapes_reference[instance] = instance.GongCopy().(*ProductCompositionShape)
		stage.ProductCompositionShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ProductShapes_reference = make(map[*ProductShape]*ProductShape)
	stage.ProductShapes_referenceOrder = make(map[*ProductShape]uint) // diff Unstage needs the reference order
	for instance := range stage.ProductShapes {
		stage.ProductShapes_reference[instance] = instance.GongCopy().(*ProductShape)
		stage.ProductShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Projects_reference = make(map[*Project]*Project)
	stage.Projects_referenceOrder = make(map[*Project]uint) // diff Unstage needs the reference order
	for instance := range stage.Projects {
		stage.Projects_reference[instance] = instance.GongCopy().(*Project)
		stage.Projects_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Roots_reference = make(map[*Root]*Root)
	stage.Roots_referenceOrder = make(map[*Root]uint) // diff Unstage needs the reference order
	for instance := range stage.Roots {
		stage.Roots_reference[instance] = instance.GongCopy().(*Root)
		stage.Roots_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Tasks_reference = make(map[*Task]*Task)
	stage.Tasks_referenceOrder = make(map[*Task]uint) // diff Unstage needs the reference order
	for instance := range stage.Tasks {
		stage.Tasks_reference[instance] = instance.GongCopy().(*Task)
		stage.Tasks_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TaskCompositionShapes_reference = make(map[*TaskCompositionShape]*TaskCompositionShape)
	stage.TaskCompositionShapes_referenceOrder = make(map[*TaskCompositionShape]uint) // diff Unstage needs the reference order
	for instance := range stage.TaskCompositionShapes {
		stage.TaskCompositionShapes_reference[instance] = instance.GongCopy().(*TaskCompositionShape)
		stage.TaskCompositionShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TaskInputShapes_reference = make(map[*TaskInputShape]*TaskInputShape)
	stage.TaskInputShapes_referenceOrder = make(map[*TaskInputShape]uint) // diff Unstage needs the reference order
	for instance := range stage.TaskInputShapes {
		stage.TaskInputShapes_reference[instance] = instance.GongCopy().(*TaskInputShape)
		stage.TaskInputShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TaskOutputShapes_reference = make(map[*TaskOutputShape]*TaskOutputShape)
	stage.TaskOutputShapes_referenceOrder = make(map[*TaskOutputShape]uint) // diff Unstage needs the reference order
	for instance := range stage.TaskOutputShapes {
		stage.TaskOutputShapes_reference[instance] = instance.GongCopy().(*TaskOutputShape)
		stage.TaskOutputShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TaskShapes_reference = make(map[*TaskShape]*TaskShape)
	stage.TaskShapes_referenceOrder = make(map[*TaskShape]uint) // diff Unstage needs the reference order
	for instance := range stage.TaskShapes {
		stage.TaskShapes_reference[instance] = instance.GongCopy().(*TaskShape)
		stage.TaskShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramMap_Staged_Order[diagram]; ok {
		return order
	}
	return stage.Diagrams_referenceOrder[diagram]
}

func (diagram *Diagram) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Diagrams_referenceOrder[diagram]
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteMap_Staged_Order[note]; ok {
		return order
	}
	return stage.Notes_referenceOrder[note]
}

func (note *Note) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Notes_referenceOrder[note]
}

func (noteproductshape *NoteProductShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteProductShapeMap_Staged_Order[noteproductshape]; ok {
		return order
	}
	return stage.NoteProductShapes_referenceOrder[noteproductshape]
}

func (noteproductshape *NoteProductShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteProductShapes_referenceOrder[noteproductshape]
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteShapeMap_Staged_Order[noteshape]; ok {
		return order
	}
	return stage.NoteShapes_referenceOrder[noteshape]
}

func (noteshape *NoteShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteShapes_referenceOrder[noteshape]
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteTaskShapeMap_Staged_Order[notetaskshape]; ok {
		return order
	}
	return stage.NoteTaskShapes_referenceOrder[notetaskshape]
}

func (notetaskshape *NoteTaskShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.NoteTaskShapes_referenceOrder[notetaskshape]
}

func (product *Product) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProductMap_Staged_Order[product]; ok {
		return order
	}
	return stage.Products_referenceOrder[product]
}

func (product *Product) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Products_referenceOrder[product]
}

func (productcompositionshape *ProductCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProductCompositionShapeMap_Staged_Order[productcompositionshape]; ok {
		return order
	}
	return stage.ProductCompositionShapes_referenceOrder[productcompositionshape]
}

func (productcompositionshape *ProductCompositionShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ProductCompositionShapes_referenceOrder[productcompositionshape]
}

func (productshape *ProductShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProductShapeMap_Staged_Order[productshape]; ok {
		return order
	}
	return stage.ProductShapes_referenceOrder[productshape]
}

func (productshape *ProductShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ProductShapes_referenceOrder[productshape]
}

func (project *Project) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProjectMap_Staged_Order[project]; ok {
		return order
	}
	return stage.Projects_referenceOrder[project]
}

func (project *Project) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Projects_referenceOrder[project]
}

func (root *Root) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RootMap_Staged_Order[root]; ok {
		return order
	}
	return stage.Roots_referenceOrder[root]
}

func (root *Root) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Roots_referenceOrder[root]
}

func (task *Task) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskMap_Staged_Order[task]; ok {
		return order
	}
	return stage.Tasks_referenceOrder[task]
}

func (task *Task) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Tasks_referenceOrder[task]
}

func (taskcompositionshape *TaskCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshape]; ok {
		return order
	}
	return stage.TaskCompositionShapes_referenceOrder[taskcompositionshape]
}

func (taskcompositionshape *TaskCompositionShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskCompositionShapes_referenceOrder[taskcompositionshape]
}

func (taskinputshape *TaskInputShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskInputShapeMap_Staged_Order[taskinputshape]; ok {
		return order
	}
	return stage.TaskInputShapes_referenceOrder[taskinputshape]
}

func (taskinputshape *TaskInputShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskInputShapes_referenceOrder[taskinputshape]
}

func (taskoutputshape *TaskOutputShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskOutputShapeMap_Staged_Order[taskoutputshape]; ok {
		return order
	}
	return stage.TaskOutputShapes_referenceOrder[taskoutputshape]
}

func (taskoutputshape *TaskOutputShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TaskOutputShapes_referenceOrder[taskoutputshape]
}

func (taskshape *TaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskShapeMap_Staged_Order[taskshape]; ok {
		return order
	}
	return stage.TaskShapes_referenceOrder[taskshape]
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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagram.Name)
	return
}
func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", note.Name)
	return
}
func (noteproductshape *NoteProductShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteProductShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", noteproductshape.Name)
	return
}
func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", noteshape.Name)
	return
}
func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", notetaskshape.Name)
	return
}
func (product *Product) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", product.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Product")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", product.Name)
	return
}
func (productcompositionshape *ProductCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", productcompositionshape.Name)
	return
}
func (productshape *ProductShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", productshape.Name)
	return
}
func (project *Project) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", project.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Project")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", project.Name)
	return
}
func (root *Root) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", root.Name)
	return
}
func (task *Task) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Task")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", task.Name)
	return
}
func (taskcompositionshape *TaskCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", taskcompositionshape.Name)
	return
}
func (taskinputshape *TaskInputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskInputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", taskinputshape.Name)
	return
}
func (taskoutputshape *TaskOutputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskOutputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", taskoutputshape.Name)
	return
}
func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", taskshape.Name)
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

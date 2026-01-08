// generated code - do not edit
package models

import (
	"fmt"
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
	stage.Project_Diagrams_reverseMap = make(map[*Diagram]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _diagram := range project.Diagrams {
			stage.Project_Diagrams_reverseMap[_diagram] = project
		}
	}
	stage.Project_Notes_reverseMap = make(map[*Note]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _note := range project.Notes {
			stage.Project_Notes_reverseMap[_note] = project
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
	stage.Root_OrphanedProducts_reverseMap = make(map[*Product]*Root)
	for root := range stage.Roots {
		_ = root
		for _, _product := range root.OrphanedProducts {
			stage.Root_OrphanedProducts_reverseMap[_product] = root
		}
	}
	stage.Root_OrphanedTasks_reverseMap = make(map[*Task]*Root)
	for root := range stage.Roots {
		_ = root
		for _, _task := range root.OrphanedTasks {
			stage.Root_OrphanedTasks_reverseMap[_task] = root
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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += diagram.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := diagram.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", diagram.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for diagram := range stage.Diagrams_reference {
		if _, ok := stage.Diagrams[diagram]; !ok {
			diagrams_deletedInstances = append(diagrams_deletedInstances, diagram)
			deletedInstancesStmt += diagram.GongMarshallUnstaging(stage)
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
			newInstancesStmt += note.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := note.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", note.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for note := range stage.Notes_reference {
		if _, ok := stage.Notes[note]; !ok {
			notes_deletedInstances = append(notes_deletedInstances, note)
			deletedInstancesStmt += note.GongMarshallUnstaging(stage)
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
			newInstancesStmt += noteproductshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := noteproductshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := noteproductshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", noteproductshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for noteproductshape := range stage.NoteProductShapes_reference {
		if _, ok := stage.NoteProductShapes[noteproductshape]; !ok {
			noteproductshapes_deletedInstances = append(noteproductshapes_deletedInstances, noteproductshape)
			deletedInstancesStmt += noteproductshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += noteshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := noteshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := noteshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", noteshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for noteshape := range stage.NoteShapes_reference {
		if _, ok := stage.NoteShapes[noteshape]; !ok {
			noteshapes_deletedInstances = append(noteshapes_deletedInstances, noteshape)
			deletedInstancesStmt += noteshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += notetaskshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := notetaskshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := notetaskshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", notetaskshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for notetaskshape := range stage.NoteTaskShapes_reference {
		if _, ok := stage.NoteTaskShapes[notetaskshape]; !ok {
			notetaskshapes_deletedInstances = append(notetaskshapes_deletedInstances, notetaskshape)
			deletedInstancesStmt += notetaskshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += product.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := product.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := product.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", product.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for product := range stage.Products_reference {
		if _, ok := stage.Products[product]; !ok {
			products_deletedInstances = append(products_deletedInstances, product)
			deletedInstancesStmt += product.GongMarshallUnstaging(stage)
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
			newInstancesStmt += productcompositionshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := productcompositionshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := productcompositionshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", productcompositionshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for productcompositionshape := range stage.ProductCompositionShapes_reference {
		if _, ok := stage.ProductCompositionShapes[productcompositionshape]; !ok {
			productcompositionshapes_deletedInstances = append(productcompositionshapes_deletedInstances, productcompositionshape)
			deletedInstancesStmt += productcompositionshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += productshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := productshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := productshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", productshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for productshape := range stage.ProductShapes_reference {
		if _, ok := stage.ProductShapes[productshape]; !ok {
			productshapes_deletedInstances = append(productshapes_deletedInstances, productshape)
			deletedInstancesStmt += productshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += project.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := project.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := project.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", project.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for project := range stage.Projects_reference {
		if _, ok := stage.Projects[project]; !ok {
			projects_deletedInstances = append(projects_deletedInstances, project)
			deletedInstancesStmt += project.GongMarshallUnstaging(stage)
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
			newInstancesStmt += root.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := root.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := root.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", root.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for root := range stage.Roots_reference {
		if _, ok := stage.Roots[root]; !ok {
			roots_deletedInstances = append(roots_deletedInstances, root)
			deletedInstancesStmt += root.GongMarshallUnstaging(stage)
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
			newInstancesStmt += task.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := task.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := task.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", task.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for task := range stage.Tasks_reference {
		if _, ok := stage.Tasks[task]; !ok {
			tasks_deletedInstances = append(tasks_deletedInstances, task)
			deletedInstancesStmt += task.GongMarshallUnstaging(stage)
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
			newInstancesStmt += taskcompositionshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := taskcompositionshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := taskcompositionshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", taskcompositionshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for taskcompositionshape := range stage.TaskCompositionShapes_reference {
		if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; !ok {
			taskcompositionshapes_deletedInstances = append(taskcompositionshapes_deletedInstances, taskcompositionshape)
			deletedInstancesStmt += taskcompositionshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += taskinputshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := taskinputshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := taskinputshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", taskinputshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for taskinputshape := range stage.TaskInputShapes_reference {
		if _, ok := stage.TaskInputShapes[taskinputshape]; !ok {
			taskinputshapes_deletedInstances = append(taskinputshapes_deletedInstances, taskinputshape)
			deletedInstancesStmt += taskinputshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += taskoutputshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := taskoutputshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := taskoutputshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", taskoutputshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for taskoutputshape := range stage.TaskOutputShapes_reference {
		if _, ok := stage.TaskOutputShapes[taskoutputshape]; !ok {
			taskoutputshapes_deletedInstances = append(taskoutputshapes_deletedInstances, taskoutputshape)
			deletedInstancesStmt += taskoutputshape.GongMarshallUnstaging(stage)
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
			newInstancesStmt += taskshape.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := taskshape.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := taskshape.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", taskshape.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for taskshape := range stage.TaskShapes_reference {
		if _, ok := stage.TaskShapes[taskshape]; !ok {
			taskshapes_deletedInstances = append(taskshapes_deletedInstances, taskshape)
			deletedInstancesStmt += taskshape.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(taskshapes_newInstances)
	lenDeletedInstances += len(taskshapes_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt+fieldsEditStmt+deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += fmt.Sprintf("\n\tstage.Commit()")
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		stage.Diagrams_reference[instance] = instance.GongCopy().(*Diagram)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	for instance := range stage.Notes {
		stage.Notes_reference[instance] = instance.GongCopy().(*Note)
	}

	stage.NoteProductShapes_reference = make(map[*NoteProductShape]*NoteProductShape)
	for instance := range stage.NoteProductShapes {
		stage.NoteProductShapes_reference[instance] = instance.GongCopy().(*NoteProductShape)
	}

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	for instance := range stage.NoteShapes {
		stage.NoteShapes_reference[instance] = instance.GongCopy().(*NoteShape)
	}

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	for instance := range stage.NoteTaskShapes {
		stage.NoteTaskShapes_reference[instance] = instance.GongCopy().(*NoteTaskShape)
	}

	stage.Products_reference = make(map[*Product]*Product)
	for instance := range stage.Products {
		stage.Products_reference[instance] = instance.GongCopy().(*Product)
	}

	stage.ProductCompositionShapes_reference = make(map[*ProductCompositionShape]*ProductCompositionShape)
	for instance := range stage.ProductCompositionShapes {
		stage.ProductCompositionShapes_reference[instance] = instance.GongCopy().(*ProductCompositionShape)
	}

	stage.ProductShapes_reference = make(map[*ProductShape]*ProductShape)
	for instance := range stage.ProductShapes {
		stage.ProductShapes_reference[instance] = instance.GongCopy().(*ProductShape)
	}

	stage.Projects_reference = make(map[*Project]*Project)
	for instance := range stage.Projects {
		stage.Projects_reference[instance] = instance.GongCopy().(*Project)
	}

	stage.Roots_reference = make(map[*Root]*Root)
	for instance := range stage.Roots {
		stage.Roots_reference[instance] = instance.GongCopy().(*Root)
	}

	stage.Tasks_reference = make(map[*Task]*Task)
	for instance := range stage.Tasks {
		stage.Tasks_reference[instance] = instance.GongCopy().(*Task)
	}

	stage.TaskCompositionShapes_reference = make(map[*TaskCompositionShape]*TaskCompositionShape)
	for instance := range stage.TaskCompositionShapes {
		stage.TaskCompositionShapes_reference[instance] = instance.GongCopy().(*TaskCompositionShape)
	}

	stage.TaskInputShapes_reference = make(map[*TaskInputShape]*TaskInputShape)
	for instance := range stage.TaskInputShapes {
		stage.TaskInputShapes_reference[instance] = instance.GongCopy().(*TaskInputShape)
	}

	stage.TaskOutputShapes_reference = make(map[*TaskOutputShape]*TaskOutputShape)
	for instance := range stage.TaskOutputShapes {
		stage.TaskOutputShapes_reference[instance] = instance.GongCopy().(*TaskOutputShape)
	}

	stage.TaskShapes_reference = make(map[*TaskShape]*TaskShape)
	for instance := range stage.TaskShapes {
		stage.TaskShapes_reference[instance] = instance.GongCopy().(*TaskShape)
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
	return stage.DiagramMap_Staged_Order[diagram]
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return stage.NoteMap_Staged_Order[note]
}

func (noteproductshape *NoteProductShape) GongGetOrder(stage *Stage) uint {
	return stage.NoteProductShapeMap_Staged_Order[noteproductshape]
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return stage.NoteShapeMap_Staged_Order[noteshape]
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	return stage.NoteTaskShapeMap_Staged_Order[notetaskshape]
}

func (product *Product) GongGetOrder(stage *Stage) uint {
	return stage.ProductMap_Staged_Order[product]
}

func (productcompositionshape *ProductCompositionShape) GongGetOrder(stage *Stage) uint {
	return stage.ProductCompositionShapeMap_Staged_Order[productcompositionshape]
}

func (productshape *ProductShape) GongGetOrder(stage *Stage) uint {
	return stage.ProductShapeMap_Staged_Order[productshape]
}

func (project *Project) GongGetOrder(stage *Stage) uint {
	return stage.ProjectMap_Staged_Order[project]
}

func (root *Root) GongGetOrder(stage *Stage) uint {
	return stage.RootMap_Staged_Order[root]
}

func (task *Task) GongGetOrder(stage *Stage) uint {
	return stage.TaskMap_Staged_Order[task]
}

func (taskcompositionshape *TaskCompositionShape) GongGetOrder(stage *Stage) uint {
	return stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshape]
}

func (taskinputshape *TaskInputShape) GongGetOrder(stage *Stage) uint {
	return stage.TaskInputShapeMap_Staged_Order[taskinputshape]
}

func (taskoutputshape *TaskOutputShape) GongGetOrder(stage *Stage) uint {
	return stage.TaskOutputShapeMap_Staged_Order[taskoutputshape]
}

func (taskshape *TaskShape) GongGetOrder(stage *Stage) uint {
	return stage.TaskShapeMap_Staged_Order[taskshape]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (noteproductshape *NoteProductShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteproductshape.GongGetGongstructName(), noteproductshape.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

func (product *Product) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", product.GongGetGongstructName(), product.GongGetOrder(stage))
}

func (productcompositionshape *ProductCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productcompositionshape.GongGetGongstructName(), productcompositionshape.GongGetOrder(stage))
}

func (productshape *ProductShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", productshape.GongGetGongstructName(), productshape.GongGetOrder(stage))
}

func (project *Project) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", project.GongGetGongstructName(), project.GongGetOrder(stage))
}

func (root *Root) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root.GongGetGongstructName(), root.GongGetOrder(stage))
}

func (task *Task) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetOrder(stage))
}

func (taskcompositionshape *TaskCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskcompositionshape.GongGetGongstructName(), taskcompositionshape.GongGetOrder(stage))
}

func (taskinputshape *TaskInputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskinputshape.GongGetGongstructName(), taskinputshape.GongGetOrder(stage))
}

func (taskoutputshape *TaskOutputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskoutputshape.GongGetGongstructName(), taskoutputshape.GongGetOrder(stage))
}

func (taskshape *TaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	return
}
func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	return
}
func (noteproductshape *NoteProductShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
	return
}
func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	return
}
func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	return
}
func (product *Product) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", product.GongGetIdentifier(stage))
	return
}
func (productcompositionshape *ProductCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
	return
}
func (productshape *ProductShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", productshape.GongGetIdentifier(stage))
	return
}
func (project *Project) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", project.GongGetIdentifier(stage))
	return
}
func (root *Root) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root.GongGetIdentifier(stage))
	return
}
func (task *Task) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetIdentifier(stage))
	return
}
func (taskcompositionshape *TaskCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
	return
}
func (taskinputshape *TaskInputShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
	return
}
func (taskoutputshape *TaskOutputShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
	return
}
func (taskshape *TaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
	return
}

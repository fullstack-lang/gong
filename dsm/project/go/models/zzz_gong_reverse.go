// generated code - do not edit
package models

// insertion point
func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Diagrams":
			if _library, ok := stage.Library_Diagrams_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "Notes":
			if _library, ok := stage.Library_Notes_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *NoteProductShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteProductShapes":
			if _diagram, ok := stage.Diagram_NoteProductShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteResourceShapes":
			if _diagram, ok := stage.Diagram_NoteResourceShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			if _diagram, ok := stage.Diagram_Note_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteTaskShapes":
			if _diagram, ok := stage.Diagram_NoteTaskShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Product) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ProductsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ProductsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootProducts":
			if _library, ok := stage.Library_RootProducts_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Products":
			if _note, ok := stage.Note_Products_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Product":
		switch reverseField.Fieldname {
		case "SubProducts":
			if _product, ok := stage.Product_SubProducts_reverseMap[inst]; ok {
				res = _product.Name
			}
		}
	case "Task":
		switch reverseField.Fieldname {
		case "Inputs":
			if _task, ok := stage.Task_Inputs_reverseMap[inst]; ok {
				res = _task.Name
			}
		case "Outputs":
			if _task, ok := stage.Task_Outputs_reverseMap[inst]; ok {
				res = _task.Name
			}
		}
	}
	return
}

func (inst *ProductCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ProductComposition_Shapes":
			if _diagram, ok := stage.Diagram_ProductComposition_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ProductReferenceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ProductReference_Shapes":
			if _diagram, ok := stage.Diagram_ProductReference_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ProductShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Product_Shapes":
			if _diagram, ok := stage.Diagram_Product_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Resource) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourcesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootResources":
			if _library, ok := stage.Library_RootResources_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Resources":
			if _note, ok := stage.Note_Resources_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Resource":
		switch reverseField.Fieldname {
		case "SubResources":
			if _resource, ok := stage.Resource_SubResources_reverseMap[inst]; ok {
				res = _resource.Name
			}
		}
	}
	return
}

func (inst *ResourceCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourceComposition_Shapes":
			if _diagram, ok := stage.Diagram_ResourceComposition_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Resource_Shapes":
			if _diagram, ok := stage.Diagram_Resource_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *ResourceTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourceTaskShapes":
			if _diagram, ok := stage.Diagram_ResourceTaskShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Task) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TasksWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_TasksWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "TasksWhoseInputNodeIsExpanded":
			if _diagram, ok := stage.Diagram_TasksWhoseInputNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "TasksWhoseOutputNodeIsExpanded":
			if _diagram, ok := stage.Diagram_TasksWhoseOutputNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		case "TasksWhosePredecessorNodeIsExpanded":
			if _diagram, ok := stage.Diagram_TasksWhosePredecessorNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootTasks":
			if _library, ok := stage.Library_RootTasks_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Tasks":
			if _note, ok := stage.Note_Tasks_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Resource":
		switch reverseField.Fieldname {
		case "Tasks":
			if _resource, ok := stage.Resource_Tasks_reverseMap[inst]; ok {
				res = _resource.Name
			}
		}
	case "Task":
		switch reverseField.Fieldname {
		case "Predecessors":
			if _task, ok := stage.Task_Predecessors_reverseMap[inst]; ok {
				res = _task.Name
			}
		case "SubTasks":
			if _task, ok := stage.Task_SubTasks_reverseMap[inst]; ok {
				res = _task.Name
			}
		}
	case "TaskGroup":
		switch reverseField.Fieldname {
		case "Tasks":
			if _taskgroup, ok := stage.TaskGroup_Tasks_reverseMap[inst]; ok {
				res = _taskgroup.Name
			}
		}
	}
	return
}

func (inst *TaskCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskComposition_Shapes":
			if _diagram, ok := stage.Diagram_TaskComposition_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *TaskGroup) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskGroupsWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_TaskGroupsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootTaskGroups":
			if _library, ok := stage.Library_RootTaskGroups_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Task":
		switch reverseField.Fieldname {
		case "TaskGroupsToDisplay":
			if _task, ok := stage.Task_TaskGroupsToDisplay_reverseMap[inst]; ok {
				res = _task.Name
			}
		}
	}
	return
}

func (inst *TaskGroupShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskGroupShapes":
			if _diagram, ok := stage.Diagram_TaskGroupShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *TaskInputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskInputShapes":
			if _diagram, ok := stage.Diagram_TaskInputShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *TaskOutputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskOutputShapes":
			if _diagram, ok := stage.Diagram_TaskOutputShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *TaskPredecessorShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskPredecessorShapes":
			if _diagram, ok := stage.Diagram_TaskPredecessorShapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *TaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Task_Shapes":
			if _diagram, ok := stage.Diagram_Task_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

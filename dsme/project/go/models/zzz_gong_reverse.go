// generated code - do not edit
package models

// insertion point
func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Project":
		switch reverseField.Fieldname {
		case "Diagrams":
			if _project, ok := stage.Project_Diagrams_reverseMap[inst]; ok {
				res = _project.Name
			}
		}
	}
	return
}

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
	case "Project":
		switch reverseField.Fieldname {
		case "Notes":
			if _project, ok := stage.Project_Notes_reverseMap[inst]; ok {
				res = _project.Name
			}
		}
	}
	return
}

func (inst *NoteProductShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *NoteTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Product) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
	case "Project":
		switch reverseField.Fieldname {
		case "RootProducts":
			if _project, ok := stage.Project_RootProducts_reverseMap[inst]; ok {
				res = _project.Name
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

func (inst *ProductCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ProductShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Project) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Root":
		switch reverseField.Fieldname {
		case "Projects":
			if _root, ok := stage.Root_Projects_reverseMap[inst]; ok {
				res = _root.Name
			}
		}
	}
	return
}

func (inst *Resource) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
	case "Note":
		switch reverseField.Fieldname {
		case "Resources":
			if _note, ok := stage.Note_Resources_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Project":
		switch reverseField.Fieldname {
		case "RootResources":
			if _project, ok := stage.Project_RootResources_reverseMap[inst]; ok {
				res = _project.Name
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

func (inst *ResourceCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ResourceTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Root) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Task) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Tasks":
			if _note, ok := stage.Note_Tasks_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Project":
		switch reverseField.Fieldname {
		case "RootTasks":
			if _project, ok := stage.Project_RootTasks_reverseMap[inst]; ok {
				res = _project.Name
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
		case "SubTasks":
			if _task, ok := stage.Task_SubTasks_reverseMap[inst]; ok {
				res = _task.Name
			}
		}
	}
	return
}

func (inst *TaskCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *TaskInputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *TaskOutputShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *TaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

// insertion point
func (inst *Diagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Project":
		switch reverseField.Fieldname {
		case "Diagrams":
			res = stage.Project_Diagrams_reverseMap[inst]
		}
	}
	return res
}

func (inst *Note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			res = stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Project":
		switch reverseField.Fieldname {
		case "Notes":
			res = stage.Project_Notes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteProductShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteProductShapes":
			res = stage.Diagram_NoteProductShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteResourceShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteResourceShapes":
			res = stage.Diagram_NoteResourceShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			res = stage.Diagram_Note_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteTaskShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteTaskShapes":
			res = stage.Diagram_NoteTaskShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Product) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ProductsWhoseNodeIsExpanded":
			res = stage.Diagram_ProductsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Products":
			res = stage.Note_Products_reverseMap[inst]
		}
	case "Product":
		switch reverseField.Fieldname {
		case "SubProducts":
			res = stage.Product_SubProducts_reverseMap[inst]
		}
	case "Project":
		switch reverseField.Fieldname {
		case "RootProducts":
			res = stage.Project_RootProducts_reverseMap[inst]
		}
	case "Task":
		switch reverseField.Fieldname {
		case "Inputs":
			res = stage.Task_Inputs_reverseMap[inst]
		case "Outputs":
			res = stage.Task_Outputs_reverseMap[inst]
		}
	}
	return res
}

func (inst *ProductCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ProductComposition_Shapes":
			res = stage.Diagram_ProductComposition_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ProductShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Product_Shapes":
			res = stage.Diagram_Product_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Project) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Root":
		switch reverseField.Fieldname {
		case "Projects":
			res = stage.Root_Projects_reverseMap[inst]
		}
	}
	return res
}

func (inst *Resource) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourcesWhoseNodeIsExpanded":
			res = stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Resources":
			res = stage.Note_Resources_reverseMap[inst]
		}
	case "Project":
		switch reverseField.Fieldname {
		case "RootResources":
			res = stage.Project_RootResources_reverseMap[inst]
		}
	case "Resource":
		switch reverseField.Fieldname {
		case "SubResources":
			res = stage.Resource_SubResources_reverseMap[inst]
		}
	}
	return res
}

func (inst *ResourceCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourceComposition_Shapes":
			res = stage.Diagram_ResourceComposition_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ResourceShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Resource_Shapes":
			res = stage.Diagram_Resource_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ResourceTaskShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "ResourceTaskShapes":
			res = stage.Diagram_ResourceTaskShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Root) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Task) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TasksWhoseNodeIsExpanded":
			res = stage.Diagram_TasksWhoseNodeIsExpanded_reverseMap[inst]
		case "TasksWhoseInputNodeIsExpanded":
			res = stage.Diagram_TasksWhoseInputNodeIsExpanded_reverseMap[inst]
		case "TasksWhoseOutputNodeIsExpanded":
			res = stage.Diagram_TasksWhoseOutputNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Tasks":
			res = stage.Note_Tasks_reverseMap[inst]
		}
	case "Project":
		switch reverseField.Fieldname {
		case "RootTasks":
			res = stage.Project_RootTasks_reverseMap[inst]
		}
	case "Resource":
		switch reverseField.Fieldname {
		case "Tasks":
			res = stage.Resource_Tasks_reverseMap[inst]
		}
	case "Task":
		switch reverseField.Fieldname {
		case "SubTasks":
			res = stage.Task_SubTasks_reverseMap[inst]
		}
	}
	return res
}

func (inst *TaskCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskComposition_Shapes":
			res = stage.Diagram_TaskComposition_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *TaskInputShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskInputShapes":
			res = stage.Diagram_TaskInputShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *TaskOutputShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "TaskOutputShapes":
			res = stage.Diagram_TaskOutputShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *TaskShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Task_Shapes":
			res = stage.Diagram_Task_Shapes_reverseMap[inst]
		}
	}
	return res
}

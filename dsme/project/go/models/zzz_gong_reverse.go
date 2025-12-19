// generated code - do not edit
package models

// insertion point
func (inst *Product) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
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
		case "Root":
			switch reverseField.Fieldname {
			case "OrphanedProducts":
				if _root, ok := stage.Root_OrphanedProducts_reverseMap[inst]; ok {
					res = _root.Name
				}
			}
		case "Task":
			switch reverseField.Fieldname {
			case "InputProducts":
				if _task, ok := stage.Task_InputProducts_reverseMap[inst]; ok {
					res = _task.Name
				}
			case "OutputProducts":
				if _task, ok := stage.Task_OutputProducts_reverseMap[inst]; ok {
					res = _task.Name
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
		case "Project":
			switch reverseField.Fieldname {
			case "RootTasks":
				if _project, ok := stage.Project_RootTasks_reverseMap[inst]; ok {
					res = _project.Name
				}
			}
		case "Root":
			switch reverseField.Fieldname {
			case "OrphanedTasks":
				if _root, ok := stage.Root_OrphanedTasks_reverseMap[inst]; ok {
					res = _root.Name
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


// insertion point
func (inst *Product) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
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
		case "Root":
			switch reverseField.Fieldname {
			case "OrphanedProducts":
				res = stage.Root_OrphanedProducts_reverseMap[inst]
			}
		case "Task":
			switch reverseField.Fieldname {
			case "InputProducts":
				res = stage.Task_InputProducts_reverseMap[inst]
			case "OutputProducts":
				res = stage.Task_OutputProducts_reverseMap[inst]
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
		case "Project":
			switch reverseField.Fieldname {
			case "RootTasks":
				res = stage.Project_RootTasks_reverseMap[inst]
			}
		case "Root":
			switch reverseField.Fieldname {
			case "OrphanedTasks":
				res = stage.Root_OrphanedTasks_reverseMap[inst]
			}
		case "Task":
			switch reverseField.Fieldname {
			case "SubTasks":
				res = stage.Task_SubTasks_reverseMap[inst]
			}
	}
	return res
}


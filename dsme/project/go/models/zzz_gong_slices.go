// generated code - do not edit
package models

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

	// Compute reverse map for named struct Product
	// insertion point per field
	stage.Product_SubProducts_reverseMap = make(map[*Product]*Product)
	for product := range stage.Products {
		_ = product
		for _, _product := range product.SubProducts {
			stage.Product_SubProducts_reverseMap[_product] = product
		}
	}

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

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Products {
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

	return
}

// insertion point per named struct
func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := *diagram
	return &newInstance
}

func (product *Product) GongCopy() GongstructIF {
	newInstance := *product
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

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
}

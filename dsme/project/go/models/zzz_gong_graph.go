// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CompositionShape:
		ok = stage.IsStagedCompositionShape(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Product:
		ok = stage.IsStagedProduct(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Project:
		ok = stage.IsStagedProject(target)

	case *Root:
		ok = stage.IsStagedRoot(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CompositionShape:
		ok = stage.IsStagedCompositionShape(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Product:
		ok = stage.IsStagedProduct(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Project:
		ok = stage.IsStagedProject(target)

	case *Root:
		ok = stage.IsStagedRoot(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCompositionShape(compositionshape *CompositionShape) (ok bool) {

	_, ok = stage.CompositionShapes[compositionshape]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedProduct(product *Product) (ok bool) {

	_, ok = stage.Products[product]

	return
}

func (stage *Stage) IsStagedProductShape(productshape *ProductShape) (ok bool) {

	_, ok = stage.ProductShapes[productshape]

	return
}

func (stage *Stage) IsStagedProject(project *Project) (ok bool) {

	_, ok = stage.Projects[project]

	return
}

func (stage *Stage) IsStagedRoot(root *Root) (ok bool) {

	_, ok = stage.Roots[root]

	return
}

func (stage *Stage) IsStagedTask(task *Task) (ok bool) {

	_, ok = stage.Tasks[task]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *CompositionShape:
		stage.StageBranchCompositionShape(target)

	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Product:
		stage.StageBranchProduct(target)

	case *ProductShape:
		stage.StageBranchProductShape(target)

	case *Project:
		stage.StageBranchProject(target)

	case *Root:
		stage.StageBranchRoot(target)

	case *Task:
		stage.StageBranchTask(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCompositionShape(compositionshape *CompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, compositionshape) {
		return
	}

	compositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compositionshape.Product != nil {
		StageBranch(stage, compositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagram.Product_Shapes {
		StageBranch(stage, _productshape)
	}
	for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
		StageBranch(stage, _product)
	}
	for _, _compositionshape := range diagram.Composition_Shapes {
		StageBranch(stage, _compositionshape)
	}

}

func (stage *Stage) StageBranchProduct(product *Product) {

	// check if instance is already staged
	if IsStaged(stage, product) {
		return
	}

	product.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range product.SubProducts {
		StageBranch(stage, _product)
	}

}

func (stage *Stage) StageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if IsStaged(stage, productshape) {
		return
	}

	productshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		StageBranch(stage, productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProject(project *Project) {

	// check if instance is already staged
	if IsStaged(stage, project) {
		return
	}

	project.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range project.RootProducts {
		StageBranch(stage, _product)
	}
	for _, _task := range project.RootTasks {
		StageBranch(stage, _task)
	}
	for _, _diagram := range project.Diagrams {
		StageBranch(stage, _diagram)
	}

}

func (stage *Stage) StageBranchRoot(root *Root) {

	// check if instance is already staged
	if IsStaged(stage, root) {
		return
	}

	root.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _project := range root.Projects {
		StageBranch(stage, _project)
	}
	for _, _product := range root.OrphanedProducts {
		StageBranch(stage, _product)
	}
	for _, _task := range root.OrphanedTasks {
		StageBranch(stage, _task)
	}

}

func (stage *Stage) StageBranchTask(task *Task) {

	// check if instance is already staged
	if IsStaged(stage, task) {
		return
	}

	task.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range task.SubTasks {
		StageBranch(stage, _task)
	}
	for _, _product := range task.Inputs {
		StageBranch(stage, _product)
	}
	for _, _product := range task.Outputs {
		StageBranch(stage, _product)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *CompositionShape:
		toT := CopyBranchCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Product:
		toT := CopyBranchProduct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductShape:
		toT := CopyBranchProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Project:
		toT := CopyBranchProject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Root:
		toT := CopyBranchRoot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Task:
		toT := CopyBranchTask(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCompositionShape(mapOrigCopy map[any]any, compositionshapeFrom *CompositionShape) (compositionshapeTo *CompositionShape) {

	// compositionshapeFrom has already been copied
	if _compositionshapeTo, ok := mapOrigCopy[compositionshapeFrom]; ok {
		compositionshapeTo = _compositionshapeTo.(*CompositionShape)
		return
	}

	compositionshapeTo = new(CompositionShape)
	mapOrigCopy[compositionshapeFrom] = compositionshapeTo
	compositionshapeFrom.CopyBasicFields(compositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if compositionshapeFrom.Product != nil {
		compositionshapeTo.Product = CopyBranchProduct(mapOrigCopy, compositionshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagramFrom.Product_Shapes {
		diagramTo.Product_Shapes = append(diagramTo.Product_Shapes, CopyBranchProductShape(mapOrigCopy, _productshape))
	}
	for _, _product := range diagramFrom.ProductsWhoseNodeIsExpanded {
		diagramTo.ProductsWhoseNodeIsExpanded = append(diagramTo.ProductsWhoseNodeIsExpanded, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _compositionshape := range diagramFrom.Composition_Shapes {
		diagramTo.Composition_Shapes = append(diagramTo.Composition_Shapes, CopyBranchCompositionShape(mapOrigCopy, _compositionshape))
	}

	return
}

func CopyBranchProduct(mapOrigCopy map[any]any, productFrom *Product) (productTo *Product) {

	// productFrom has already been copied
	if _productTo, ok := mapOrigCopy[productFrom]; ok {
		productTo = _productTo.(*Product)
		return
	}

	productTo = new(Product)
	mapOrigCopy[productFrom] = productTo
	productFrom.CopyBasicFields(productTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range productFrom.SubProducts {
		productTo.SubProducts = append(productTo.SubProducts, CopyBranchProduct(mapOrigCopy, _product))
	}

	return
}

func CopyBranchProductShape(mapOrigCopy map[any]any, productshapeFrom *ProductShape) (productshapeTo *ProductShape) {

	// productshapeFrom has already been copied
	if _productshapeTo, ok := mapOrigCopy[productshapeFrom]; ok {
		productshapeTo = _productshapeTo.(*ProductShape)
		return
	}

	productshapeTo = new(ProductShape)
	mapOrigCopy[productshapeFrom] = productshapeTo
	productshapeFrom.CopyBasicFields(productshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productshapeFrom.Product != nil {
		productshapeTo.Product = CopyBranchProduct(mapOrigCopy, productshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProject(mapOrigCopy map[any]any, projectFrom *Project) (projectTo *Project) {

	// projectFrom has already been copied
	if _projectTo, ok := mapOrigCopy[projectFrom]; ok {
		projectTo = _projectTo.(*Project)
		return
	}

	projectTo = new(Project)
	mapOrigCopy[projectFrom] = projectTo
	projectFrom.CopyBasicFields(projectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range projectFrom.RootProducts {
		projectTo.RootProducts = append(projectTo.RootProducts, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range projectFrom.RootTasks {
		projectTo.RootTasks = append(projectTo.RootTasks, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _diagram := range projectFrom.Diagrams {
		projectTo.Diagrams = append(projectTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func CopyBranchRoot(mapOrigCopy map[any]any, rootFrom *Root) (rootTo *Root) {

	// rootFrom has already been copied
	if _rootTo, ok := mapOrigCopy[rootFrom]; ok {
		rootTo = _rootTo.(*Root)
		return
	}

	rootTo = new(Root)
	mapOrigCopy[rootFrom] = rootTo
	rootFrom.CopyBasicFields(rootTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _project := range rootFrom.Projects {
		rootTo.Projects = append(rootTo.Projects, CopyBranchProject(mapOrigCopy, _project))
	}
	for _, _product := range rootFrom.OrphanedProducts {
		rootTo.OrphanedProducts = append(rootTo.OrphanedProducts, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range rootFrom.OrphanedTasks {
		rootTo.OrphanedTasks = append(rootTo.OrphanedTasks, CopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func CopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {

	// taskFrom has already been copied
	if _taskTo, ok := mapOrigCopy[taskFrom]; ok {
		taskTo = _taskTo.(*Task)
		return
	}

	taskTo = new(Task)
	mapOrigCopy[taskFrom] = taskTo
	taskFrom.CopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskFrom.SubTasks {
		taskTo.SubTasks = append(taskTo.SubTasks, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _product := range taskFrom.Inputs {
		taskTo.Inputs = append(taskTo.Inputs, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _product := range taskFrom.Outputs {
		taskTo.Outputs = append(taskTo.Outputs, CopyBranchProduct(mapOrigCopy, _product))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *CompositionShape:
		stage.UnstageBranchCompositionShape(target)

	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Product:
		stage.UnstageBranchProduct(target)

	case *ProductShape:
		stage.UnstageBranchProductShape(target)

	case *Project:
		stage.UnstageBranchProject(target)

	case *Root:
		stage.UnstageBranchRoot(target)

	case *Task:
		stage.UnstageBranchTask(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCompositionShape(compositionshape *CompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, compositionshape) {
		return
	}

	compositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compositionshape.Product != nil {
		UnstageBranch(stage, compositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagram.Product_Shapes {
		UnstageBranch(stage, _productshape)
	}
	for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
		UnstageBranch(stage, _product)
	}
	for _, _compositionshape := range diagram.Composition_Shapes {
		UnstageBranch(stage, _compositionshape)
	}

}

func (stage *Stage) UnstageBranchProduct(product *Product) {

	// check if instance is already staged
	if !IsStaged(stage, product) {
		return
	}

	product.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range product.SubProducts {
		UnstageBranch(stage, _product)
	}

}

func (stage *Stage) UnstageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if !IsStaged(stage, productshape) {
		return
	}

	productshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		UnstageBranch(stage, productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProject(project *Project) {

	// check if instance is already staged
	if !IsStaged(stage, project) {
		return
	}

	project.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range project.RootProducts {
		UnstageBranch(stage, _product)
	}
	for _, _task := range project.RootTasks {
		UnstageBranch(stage, _task)
	}
	for _, _diagram := range project.Diagrams {
		UnstageBranch(stage, _diagram)
	}

}

func (stage *Stage) UnstageBranchRoot(root *Root) {

	// check if instance is already staged
	if !IsStaged(stage, root) {
		return
	}

	root.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _project := range root.Projects {
		UnstageBranch(stage, _project)
	}
	for _, _product := range root.OrphanedProducts {
		UnstageBranch(stage, _product)
	}
	for _, _task := range root.OrphanedTasks {
		UnstageBranch(stage, _task)
	}

}

func (stage *Stage) UnstageBranchTask(task *Task) {

	// check if instance is already staged
	if !IsStaged(stage, task) {
		return
	}

	task.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range task.SubTasks {
		UnstageBranch(stage, _task)
	}
	for _, _product := range task.Inputs {
		UnstageBranch(stage, _product)
	}
	for _, _product := range task.Outputs {
		UnstageBranch(stage, _product)
	}

}

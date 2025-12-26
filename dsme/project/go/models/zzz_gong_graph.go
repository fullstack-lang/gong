// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *Product:
		ok = stage.IsStagedProduct(target)

	case *ProductCompositionShape:
		ok = stage.IsStagedProductCompositionShape(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Project:
		ok = stage.IsStagedProject(target)

	case *Root:
		ok = stage.IsStagedRoot(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	case *TaskCompositionShape:
		ok = stage.IsStagedTaskCompositionShape(target)

	case *TaskInputShape:
		ok = stage.IsStagedTaskInputShape(target)

	case *TaskOutputShape:
		ok = stage.IsStagedTaskOutputShape(target)

	case *TaskShape:
		ok = stage.IsStagedTaskShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *Product:
		ok = stage.IsStagedProduct(target)

	case *ProductCompositionShape:
		ok = stage.IsStagedProductCompositionShape(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Project:
		ok = stage.IsStagedProject(target)

	case *Root:
		ok = stage.IsStagedRoot(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	case *TaskCompositionShape:
		ok = stage.IsStagedTaskCompositionShape(target)

	case *TaskInputShape:
		ok = stage.IsStagedTaskInputShape(target)

	case *TaskOutputShape:
		ok = stage.IsStagedTaskOutputShape(target)

	case *TaskShape:
		ok = stage.IsStagedTaskShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedProduct(product *Product) (ok bool) {

	_, ok = stage.Products[product]

	return
}

func (stage *Stage) IsStagedProductCompositionShape(productcompositionshape *ProductCompositionShape) (ok bool) {

	_, ok = stage.ProductCompositionShapes[productcompositionshape]

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

func (stage *Stage) IsStagedTaskCompositionShape(taskcompositionshape *TaskCompositionShape) (ok bool) {

	_, ok = stage.TaskCompositionShapes[taskcompositionshape]

	return
}

func (stage *Stage) IsStagedTaskInputShape(taskinputshape *TaskInputShape) (ok bool) {

	_, ok = stage.TaskInputShapes[taskinputshape]

	return
}

func (stage *Stage) IsStagedTaskOutputShape(taskoutputshape *TaskOutputShape) (ok bool) {

	_, ok = stage.TaskOutputShapes[taskoutputshape]

	return
}

func (stage *Stage) IsStagedTaskShape(taskshape *TaskShape) (ok bool) {

	_, ok = stage.TaskShapes[taskshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Note:
		stage.StageBranchNote(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *Product:
		stage.StageBranchProduct(target)

	case *ProductCompositionShape:
		stage.StageBranchProductCompositionShape(target)

	case *ProductShape:
		stage.StageBranchProductShape(target)

	case *Project:
		stage.StageBranchProject(target)

	case *Root:
		stage.StageBranchRoot(target)

	case *Task:
		stage.StageBranchTask(target)

	case *TaskCompositionShape:
		stage.StageBranchTaskCompositionShape(target)

	case *TaskInputShape:
		stage.StageBranchTaskInputShape(target)

	case *TaskOutputShape:
		stage.StageBranchTaskOutputShape(target)

	case *TaskShape:
		stage.StageBranchTaskShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
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
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		StageBranch(stage, _productcompositionshape)
	}
	for _, _taskshape := range diagram.Task_Shapes {
		StageBranch(stage, _taskshape)
	}
	for _, _task := range diagram.TasksWhoseNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
		StageBranch(stage, _taskcompositionshape)
	}
	for _, _taskinputshape := range diagram.TaskInputShapes {
		StageBranch(stage, _taskinputshape)
	}
	for _, _taskoutputshape := range diagram.TaskOutputShapes {
		StageBranch(stage, _taskoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range note.Products {
		StageBranch(stage, _product)
	}

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		StageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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

func (stage *Stage) StageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, productcompositionshape) {
		return
	}

	productcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		StageBranch(stage, productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	for _, _note := range project.Notes {
		StageBranch(stage, _note)
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

func (stage *Stage) StageBranchTaskCompositionShape(taskcompositionshape *TaskCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, taskcompositionshape) {
		return
	}

	taskcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshape.Task != nil {
		StageBranch(stage, taskcompositionshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskInputShape(taskinputshape *TaskInputShape) {

	// check if instance is already staged
	if IsStaged(stage, taskinputshape) {
		return
	}

	taskinputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshape.Task != nil {
		StageBranch(stage, taskinputshape.Task)
	}
	if taskinputshape.Product != nil {
		StageBranch(stage, taskinputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskOutputShape(taskoutputshape *TaskOutputShape) {

	// check if instance is already staged
	if IsStaged(stage, taskoutputshape) {
		return
	}

	taskoutputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshape.Task != nil {
		StageBranch(stage, taskoutputshape.Task)
	}
	if taskoutputshape.Product != nil {
		StageBranch(stage, taskoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if IsStaged(stage, taskshape) {
		return
	}

	taskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		StageBranch(stage, taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Product:
		toT := CopyBranchProduct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductCompositionShape:
		toT := CopyBranchProductCompositionShape(mapOrigCopy, fromT)
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

	case *TaskCompositionShape:
		toT := CopyBranchTaskCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskInputShape:
		toT := CopyBranchTaskInputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskOutputShape:
		toT := CopyBranchTaskOutputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskShape:
		toT := CopyBranchTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
	for _, _productcompositionshape := range diagramFrom.ProductComposition_Shapes {
		diagramTo.ProductComposition_Shapes = append(diagramTo.ProductComposition_Shapes, CopyBranchProductCompositionShape(mapOrigCopy, _productcompositionshape))
	}
	for _, _taskshape := range diagramFrom.Task_Shapes {
		diagramTo.Task_Shapes = append(diagramTo.Task_Shapes, CopyBranchTaskShape(mapOrigCopy, _taskshape))
	}
	for _, _task := range diagramFrom.TasksWhoseNodeIsExpanded {
		diagramTo.TasksWhoseNodeIsExpanded = append(diagramTo.TasksWhoseNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhoseInputNodeIsExpanded {
		diagramTo.TasksWhoseInputNodeIsExpanded = append(diagramTo.TasksWhoseInputNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhoseOutputNodeIsExpanded {
		diagramTo.TasksWhoseOutputNodeIsExpanded = append(diagramTo.TasksWhoseOutputNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskcompositionshape := range diagramFrom.TaskComposition_Shapes {
		diagramTo.TaskComposition_Shapes = append(diagramTo.TaskComposition_Shapes, CopyBranchTaskCompositionShape(mapOrigCopy, _taskcompositionshape))
	}
	for _, _taskinputshape := range diagramFrom.TaskInputShapes {
		diagramTo.TaskInputShapes = append(diagramTo.TaskInputShapes, CopyBranchTaskInputShape(mapOrigCopy, _taskinputshape))
	}
	for _, _taskoutputshape := range diagramFrom.TaskOutputShapes {
		diagramTo.TaskOutputShapes = append(diagramTo.TaskOutputShapes, CopyBranchTaskOutputShape(mapOrigCopy, _taskoutputshape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramFrom.NotesWhoseNodeIsExpanded {
		diagramTo.NotesWhoseNodeIsExpanded = append(diagramTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range noteFrom.Products {
		noteTo.Products = append(noteTo.Products, CopyBranchProduct(mapOrigCopy, _product))
	}

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = CopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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

func CopyBranchProductCompositionShape(mapOrigCopy map[any]any, productcompositionshapeFrom *ProductCompositionShape) (productcompositionshapeTo *ProductCompositionShape) {

	// productcompositionshapeFrom has already been copied
	if _productcompositionshapeTo, ok := mapOrigCopy[productcompositionshapeFrom]; ok {
		productcompositionshapeTo = _productcompositionshapeTo.(*ProductCompositionShape)
		return
	}

	productcompositionshapeTo = new(ProductCompositionShape)
	mapOrigCopy[productcompositionshapeFrom] = productcompositionshapeTo
	productcompositionshapeFrom.CopyBasicFields(productcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshapeFrom.Product != nil {
		productcompositionshapeTo.Product = CopyBranchProduct(mapOrigCopy, productcompositionshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	for _, _note := range projectFrom.Notes {
		projectTo.Notes = append(projectTo.Notes, CopyBranchNote(mapOrigCopy, _note))
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

func CopyBranchTaskCompositionShape(mapOrigCopy map[any]any, taskcompositionshapeFrom *TaskCompositionShape) (taskcompositionshapeTo *TaskCompositionShape) {

	// taskcompositionshapeFrom has already been copied
	if _taskcompositionshapeTo, ok := mapOrigCopy[taskcompositionshapeFrom]; ok {
		taskcompositionshapeTo = _taskcompositionshapeTo.(*TaskCompositionShape)
		return
	}

	taskcompositionshapeTo = new(TaskCompositionShape)
	mapOrigCopy[taskcompositionshapeFrom] = taskcompositionshapeTo
	taskcompositionshapeFrom.CopyBasicFields(taskcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshapeFrom.Task != nil {
		taskcompositionshapeTo.Task = CopyBranchTask(mapOrigCopy, taskcompositionshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskInputShape(mapOrigCopy map[any]any, taskinputshapeFrom *TaskInputShape) (taskinputshapeTo *TaskInputShape) {

	// taskinputshapeFrom has already been copied
	if _taskinputshapeTo, ok := mapOrigCopy[taskinputshapeFrom]; ok {
		taskinputshapeTo = _taskinputshapeTo.(*TaskInputShape)
		return
	}

	taskinputshapeTo = new(TaskInputShape)
	mapOrigCopy[taskinputshapeFrom] = taskinputshapeTo
	taskinputshapeFrom.CopyBasicFields(taskinputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshapeFrom.Task != nil {
		taskinputshapeTo.Task = CopyBranchTask(mapOrigCopy, taskinputshapeFrom.Task)
	}
	if taskinputshapeFrom.Product != nil {
		taskinputshapeTo.Product = CopyBranchProduct(mapOrigCopy, taskinputshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskOutputShape(mapOrigCopy map[any]any, taskoutputshapeFrom *TaskOutputShape) (taskoutputshapeTo *TaskOutputShape) {

	// taskoutputshapeFrom has already been copied
	if _taskoutputshapeTo, ok := mapOrigCopy[taskoutputshapeFrom]; ok {
		taskoutputshapeTo = _taskoutputshapeTo.(*TaskOutputShape)
		return
	}

	taskoutputshapeTo = new(TaskOutputShape)
	mapOrigCopy[taskoutputshapeFrom] = taskoutputshapeTo
	taskoutputshapeFrom.CopyBasicFields(taskoutputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshapeFrom.Task != nil {
		taskoutputshapeTo.Task = CopyBranchTask(mapOrigCopy, taskoutputshapeFrom.Task)
	}
	if taskoutputshapeFrom.Product != nil {
		taskoutputshapeTo.Product = CopyBranchProduct(mapOrigCopy, taskoutputshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskShape(mapOrigCopy map[any]any, taskshapeFrom *TaskShape) (taskshapeTo *TaskShape) {

	// taskshapeFrom has already been copied
	if _taskshapeTo, ok := mapOrigCopy[taskshapeFrom]; ok {
		taskshapeTo = _taskshapeTo.(*TaskShape)
		return
	}

	taskshapeTo = new(TaskShape)
	mapOrigCopy[taskshapeFrom] = taskshapeTo
	taskshapeFrom.CopyBasicFields(taskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskshapeFrom.Task != nil {
		taskshapeTo.Task = CopyBranchTask(mapOrigCopy, taskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *Product:
		stage.UnstageBranchProduct(target)

	case *ProductCompositionShape:
		stage.UnstageBranchProductCompositionShape(target)

	case *ProductShape:
		stage.UnstageBranchProductShape(target)

	case *Project:
		stage.UnstageBranchProject(target)

	case *Root:
		stage.UnstageBranchRoot(target)

	case *Task:
		stage.UnstageBranchTask(target)

	case *TaskCompositionShape:
		stage.UnstageBranchTaskCompositionShape(target)

	case *TaskInputShape:
		stage.UnstageBranchTaskInputShape(target)

	case *TaskOutputShape:
		stage.UnstageBranchTaskOutputShape(target)

	case *TaskShape:
		stage.UnstageBranchTaskShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
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
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		UnstageBranch(stage, _productcompositionshape)
	}
	for _, _taskshape := range diagram.Task_Shapes {
		UnstageBranch(stage, _taskshape)
	}
	for _, _task := range diagram.TasksWhoseNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
		UnstageBranch(stage, _taskcompositionshape)
	}
	for _, _taskinputshape := range diagram.TaskInputShapes {
		UnstageBranch(stage, _taskinputshape)
	}
	for _, _taskoutputshape := range diagram.TaskOutputShapes {
		UnstageBranch(stage, _taskoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range note.Products {
		UnstageBranch(stage, _product)
	}

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		UnstageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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

func (stage *Stage) UnstageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, productcompositionshape) {
		return
	}

	productcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		UnstageBranch(stage, productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	for _, _note := range project.Notes {
		UnstageBranch(stage, _note)
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

func (stage *Stage) UnstageBranchTaskCompositionShape(taskcompositionshape *TaskCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskcompositionshape) {
		return
	}

	taskcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshape.Task != nil {
		UnstageBranch(stage, taskcompositionshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskInputShape(taskinputshape *TaskInputShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskinputshape) {
		return
	}

	taskinputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshape.Task != nil {
		UnstageBranch(stage, taskinputshape.Task)
	}
	if taskinputshape.Product != nil {
		UnstageBranch(stage, taskinputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskOutputShape(taskoutputshape *TaskOutputShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskoutputshape) {
		return
	}

	taskoutputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshape.Task != nil {
		UnstageBranch(stage, taskoutputshape.Task)
	}
	if taskoutputshape.Product != nil {
		UnstageBranch(stage, taskoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskshape) {
		return
	}

	taskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		UnstageBranch(stage, taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

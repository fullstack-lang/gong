// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteProductShape:
		ok = stage.IsStagedNoteProductShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteTaskShape:
		ok = stage.IsStagedNoteTaskShape(target)

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

	case *NoteProductShape:
		ok = stage.IsStagedNoteProductShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteTaskShape:
		ok = stage.IsStagedNoteTaskShape(target)

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

func (stage *Stage) IsStagedNoteProductShape(noteproductshape *NoteProductShape) (ok bool) {

	_, ok = stage.NoteProductShapes[noteproductshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteTaskShape(notetaskshape *NoteTaskShape) (ok bool) {

	_, ok = stage.NoteTaskShapes[notetaskshape]

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

	case *NoteProductShape:
		stage.StageBranchNoteProductShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteTaskShape:
		stage.StageBranchNoteTaskShape(target)

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
	for _, _noteproductshape := range diagram.NoteProductShapes {
		StageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		StageBranch(stage, _notetaskshape)
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
	for _, _task := range note.Tasks {
		StageBranch(stage, _task)
	}

}

func (stage *Stage) StageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if IsStaged(stage, noteproductshape) {
		return
	}

	noteproductshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		StageBranch(stage, noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		StageBranch(stage, noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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

func (stage *Stage) StageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if IsStaged(stage, notetaskshape) {
		return
	}

	notetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		StageBranch(stage, notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		StageBranch(stage, notetaskshape.Task)
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

	case *NoteProductShape:
		toT := CopyBranchNoteProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTaskShape:
		toT := CopyBranchNoteTaskShape(mapOrigCopy, fromT)
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
	for _, _noteproductshape := range diagramFrom.NoteProductShapes {
		diagramTo.NoteProductShapes = append(diagramTo.NoteProductShapes, CopyBranchNoteProductShape(mapOrigCopy, _noteproductshape))
	}
	for _, _notetaskshape := range diagramFrom.NoteTaskShapes {
		diagramTo.NoteTaskShapes = append(diagramTo.NoteTaskShapes, CopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
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
	for _, _task := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, CopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func CopyBranchNoteProductShape(mapOrigCopy map[any]any, noteproductshapeFrom *NoteProductShape) (noteproductshapeTo *NoteProductShape) {

	// noteproductshapeFrom has already been copied
	if _noteproductshapeTo, ok := mapOrigCopy[noteproductshapeFrom]; ok {
		noteproductshapeTo = _noteproductshapeTo.(*NoteProductShape)
		return
	}

	noteproductshapeTo = new(NoteProductShape)
	mapOrigCopy[noteproductshapeFrom] = noteproductshapeTo
	noteproductshapeFrom.CopyBasicFields(noteproductshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshapeFrom.Note != nil {
		noteproductshapeTo.Note = CopyBranchNote(mapOrigCopy, noteproductshapeFrom.Note)
	}
	if noteproductshapeFrom.Product != nil {
		noteproductshapeTo.Product = CopyBranchProduct(mapOrigCopy, noteproductshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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

func CopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {

	// notetaskshapeFrom has already been copied
	if _notetaskshapeTo, ok := mapOrigCopy[notetaskshapeFrom]; ok {
		notetaskshapeTo = _notetaskshapeTo.(*NoteTaskShape)
		return
	}

	notetaskshapeTo = new(NoteTaskShape)
	mapOrigCopy[notetaskshapeFrom] = notetaskshapeTo
	notetaskshapeFrom.CopyBasicFields(notetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshapeFrom.Note != nil {
		notetaskshapeTo.Note = CopyBranchNote(mapOrigCopy, notetaskshapeFrom.Note)
	}
	if notetaskshapeFrom.Task != nil {
		notetaskshapeTo.Task = CopyBranchTask(mapOrigCopy, notetaskshapeFrom.Task)
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

	case *NoteProductShape:
		stage.UnstageBranchNoteProductShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteTaskShape:
		stage.UnstageBranchNoteTaskShape(target)

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
	for _, _noteproductshape := range diagram.NoteProductShapes {
		UnstageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		UnstageBranch(stage, _notetaskshape)
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
	for _, _task := range note.Tasks {
		UnstageBranch(stage, _task)
	}

}

func (stage *Stage) UnstageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteproductshape) {
		return
	}

	noteproductshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		UnstageBranch(stage, noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		UnstageBranch(stage, noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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

func (stage *Stage) UnstageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, notetaskshape) {
		return
	}

	notetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		UnstageBranch(stage, notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		UnstageBranch(stage, notetaskshape.Task)
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

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, "Name")
	}
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, "IsChecked")
	}
	if diagram.IsEditable_ != diagramOther.IsEditable_ {
		diffs = append(diffs, "IsEditable_")
	}
	if diagram.IsInRenameMode != diagramOther.IsInRenameMode {
		diffs = append(diffs, "IsInRenameMode")
	}
	if diagram.ShowPrefix != diagramOther.ShowPrefix {
		diffs = append(diffs, "ShowPrefix")
	}
	if diagram.DefaultBoxWidth != diagramOther.DefaultBoxWidth {
		diffs = append(diffs, "DefaultBoxWidth")
	}
	if diagram.DefaultBoxHeigth != diagramOther.DefaultBoxHeigth {
		diffs = append(diffs, "DefaultBoxHeigth")
	}
	if diagram.IsExpanded != diagramOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if diagram.ComputedPrefix != diagramOther.ComputedPrefix {
		diffs = append(diffs, "ComputedPrefix")
	}
	Product_ShapesDifferent := false
	if len(diagram.Product_Shapes) != len(diagramOther.Product_Shapes) {
		Product_ShapesDifferent = true
	} else {
		for i := range diagram.Product_Shapes {
			if (diagram.Product_Shapes[i] == nil) != (diagramOther.Product_Shapes[i] == nil) {
				Product_ShapesDifferent = true
				break
			} else if diagram.Product_Shapes[i] != nil && diagramOther.Product_Shapes[i] != nil {
				if len(diagram.Product_Shapes[i].GongDiff(diagramOther.Product_Shapes[i])) > 0 {
					Product_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Product_ShapesDifferent {
		diffs = append(diffs, "Product_Shapes")
	}
	ProductsWhoseNodeIsExpandedDifferent := false
	if len(diagram.ProductsWhoseNodeIsExpanded) != len(diagramOther.ProductsWhoseNodeIsExpanded) {
		ProductsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ProductsWhoseNodeIsExpanded {
			if (diagram.ProductsWhoseNodeIsExpanded[i] == nil) != (diagramOther.ProductsWhoseNodeIsExpanded[i] == nil) {
				ProductsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ProductsWhoseNodeIsExpanded[i] != nil && diagramOther.ProductsWhoseNodeIsExpanded[i] != nil {
				if len(diagram.ProductsWhoseNodeIsExpanded[i].GongDiff(diagramOther.ProductsWhoseNodeIsExpanded[i])) > 0 {
					ProductsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ProductsWhoseNodeIsExpandedDifferent {
		diffs = append(diffs, "ProductsWhoseNodeIsExpanded")
	}
	if diagram.IsPBSNodeExpanded != diagramOther.IsPBSNodeExpanded {
		diffs = append(diffs, "IsPBSNodeExpanded")
	}
	ProductComposition_ShapesDifferent := false
	if len(diagram.ProductComposition_Shapes) != len(diagramOther.ProductComposition_Shapes) {
		ProductComposition_ShapesDifferent = true
	} else {
		for i := range diagram.ProductComposition_Shapes {
			if (diagram.ProductComposition_Shapes[i] == nil) != (diagramOther.ProductComposition_Shapes[i] == nil) {
				ProductComposition_ShapesDifferent = true
				break
			} else if diagram.ProductComposition_Shapes[i] != nil && diagramOther.ProductComposition_Shapes[i] != nil {
				if len(diagram.ProductComposition_Shapes[i].GongDiff(diagramOther.ProductComposition_Shapes[i])) > 0 {
					ProductComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ProductComposition_ShapesDifferent {
		diffs = append(diffs, "ProductComposition_Shapes")
	}
	if diagram.IsWBSNodeExpanded != diagramOther.IsWBSNodeExpanded {
		diffs = append(diffs, "IsWBSNodeExpanded")
	}
	Task_ShapesDifferent := false
	if len(diagram.Task_Shapes) != len(diagramOther.Task_Shapes) {
		Task_ShapesDifferent = true
	} else {
		for i := range diagram.Task_Shapes {
			if (diagram.Task_Shapes[i] == nil) != (diagramOther.Task_Shapes[i] == nil) {
				Task_ShapesDifferent = true
				break
			} else if diagram.Task_Shapes[i] != nil && diagramOther.Task_Shapes[i] != nil {
				if len(diagram.Task_Shapes[i].GongDiff(diagramOther.Task_Shapes[i])) > 0 {
					Task_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Task_ShapesDifferent {
		diffs = append(diffs, "Task_Shapes")
	}
	TasksWhoseNodeIsExpandedDifferent := false
	if len(diagram.TasksWhoseNodeIsExpanded) != len(diagramOther.TasksWhoseNodeIsExpanded) {
		TasksWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhoseNodeIsExpanded {
			if (diagram.TasksWhoseNodeIsExpanded[i] == nil) != (diagramOther.TasksWhoseNodeIsExpanded[i] == nil) {
				TasksWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhoseNodeIsExpanded[i] != nil && diagramOther.TasksWhoseNodeIsExpanded[i] != nil {
				if len(diagram.TasksWhoseNodeIsExpanded[i].GongDiff(diagramOther.TasksWhoseNodeIsExpanded[i])) > 0 {
					TasksWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseNodeIsExpandedDifferent {
		diffs = append(diffs, "TasksWhoseNodeIsExpanded")
	}
	TasksWhoseInputNodeIsExpandedDifferent := false
	if len(diagram.TasksWhoseInputNodeIsExpanded) != len(diagramOther.TasksWhoseInputNodeIsExpanded) {
		TasksWhoseInputNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhoseInputNodeIsExpanded {
			if (diagram.TasksWhoseInputNodeIsExpanded[i] == nil) != (diagramOther.TasksWhoseInputNodeIsExpanded[i] == nil) {
				TasksWhoseInputNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhoseInputNodeIsExpanded[i] != nil && diagramOther.TasksWhoseInputNodeIsExpanded[i] != nil {
				if len(diagram.TasksWhoseInputNodeIsExpanded[i].GongDiff(diagramOther.TasksWhoseInputNodeIsExpanded[i])) > 0 {
					TasksWhoseInputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseInputNodeIsExpandedDifferent {
		diffs = append(diffs, "TasksWhoseInputNodeIsExpanded")
	}
	TasksWhoseOutputNodeIsExpandedDifferent := false
	if len(diagram.TasksWhoseOutputNodeIsExpanded) != len(diagramOther.TasksWhoseOutputNodeIsExpanded) {
		TasksWhoseOutputNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhoseOutputNodeIsExpanded {
			if (diagram.TasksWhoseOutputNodeIsExpanded[i] == nil) != (diagramOther.TasksWhoseOutputNodeIsExpanded[i] == nil) {
				TasksWhoseOutputNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhoseOutputNodeIsExpanded[i] != nil && diagramOther.TasksWhoseOutputNodeIsExpanded[i] != nil {
				if len(diagram.TasksWhoseOutputNodeIsExpanded[i].GongDiff(diagramOther.TasksWhoseOutputNodeIsExpanded[i])) > 0 {
					TasksWhoseOutputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseOutputNodeIsExpandedDifferent {
		diffs = append(diffs, "TasksWhoseOutputNodeIsExpanded")
	}
	TaskComposition_ShapesDifferent := false
	if len(diagram.TaskComposition_Shapes) != len(diagramOther.TaskComposition_Shapes) {
		TaskComposition_ShapesDifferent = true
	} else {
		for i := range diagram.TaskComposition_Shapes {
			if (diagram.TaskComposition_Shapes[i] == nil) != (diagramOther.TaskComposition_Shapes[i] == nil) {
				TaskComposition_ShapesDifferent = true
				break
			} else if diagram.TaskComposition_Shapes[i] != nil && diagramOther.TaskComposition_Shapes[i] != nil {
				if len(diagram.TaskComposition_Shapes[i].GongDiff(diagramOther.TaskComposition_Shapes[i])) > 0 {
					TaskComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskComposition_ShapesDifferent {
		diffs = append(diffs, "TaskComposition_Shapes")
	}
	TaskInputShapesDifferent := false
	if len(diagram.TaskInputShapes) != len(diagramOther.TaskInputShapes) {
		TaskInputShapesDifferent = true
	} else {
		for i := range diagram.TaskInputShapes {
			if (diagram.TaskInputShapes[i] == nil) != (diagramOther.TaskInputShapes[i] == nil) {
				TaskInputShapesDifferent = true
				break
			} else if diagram.TaskInputShapes[i] != nil && diagramOther.TaskInputShapes[i] != nil {
				if len(diagram.TaskInputShapes[i].GongDiff(diagramOther.TaskInputShapes[i])) > 0 {
					TaskInputShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskInputShapesDifferent {
		diffs = append(diffs, "TaskInputShapes")
	}
	TaskOutputShapesDifferent := false
	if len(diagram.TaskOutputShapes) != len(diagramOther.TaskOutputShapes) {
		TaskOutputShapesDifferent = true
	} else {
		for i := range diagram.TaskOutputShapes {
			if (diagram.TaskOutputShapes[i] == nil) != (diagramOther.TaskOutputShapes[i] == nil) {
				TaskOutputShapesDifferent = true
				break
			} else if diagram.TaskOutputShapes[i] != nil && diagramOther.TaskOutputShapes[i] != nil {
				if len(diagram.TaskOutputShapes[i].GongDiff(diagramOther.TaskOutputShapes[i])) > 0 {
					TaskOutputShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskOutputShapesDifferent {
		diffs = append(diffs, "TaskOutputShapes")
	}
	Note_ShapesDifferent := false
	if len(diagram.Note_Shapes) != len(diagramOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagram.Note_Shapes {
			if (diagram.Note_Shapes[i] == nil) != (diagramOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagram.Note_Shapes[i] != nil && diagramOther.Note_Shapes[i] != nil {
				if len(diagram.Note_Shapes[i].GongDiff(diagramOther.Note_Shapes[i])) > 0 {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		diffs = append(diffs, "Note_Shapes")
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagram.NotesWhoseNodeIsExpanded) != len(diagramOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.NotesWhoseNodeIsExpanded {
			if (diagram.NotesWhoseNodeIsExpanded[i] == nil) != (diagramOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.NotesWhoseNodeIsExpanded[i] != nil && diagramOther.NotesWhoseNodeIsExpanded[i] != nil {
				if len(diagram.NotesWhoseNodeIsExpanded[i].GongDiff(diagramOther.NotesWhoseNodeIsExpanded[i])) > 0 {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		diffs = append(diffs, "NotesWhoseNodeIsExpanded")
	}
	if diagram.IsNotesNodeExpanded != diagramOther.IsNotesNodeExpanded {
		diffs = append(diffs, "IsNotesNodeExpanded")
	}
	NoteProductShapesDifferent := false
	if len(diagram.NoteProductShapes) != len(diagramOther.NoteProductShapes) {
		NoteProductShapesDifferent = true
	} else {
		for i := range diagram.NoteProductShapes {
			if (diagram.NoteProductShapes[i] == nil) != (diagramOther.NoteProductShapes[i] == nil) {
				NoteProductShapesDifferent = true
				break
			} else if diagram.NoteProductShapes[i] != nil && diagramOther.NoteProductShapes[i] != nil {
				if len(diagram.NoteProductShapes[i].GongDiff(diagramOther.NoteProductShapes[i])) > 0 {
					NoteProductShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteProductShapesDifferent {
		diffs = append(diffs, "NoteProductShapes")
	}
	NoteTaskShapesDifferent := false
	if len(diagram.NoteTaskShapes) != len(diagramOther.NoteTaskShapes) {
		NoteTaskShapesDifferent = true
	} else {
		for i := range diagram.NoteTaskShapes {
			if (diagram.NoteTaskShapes[i] == nil) != (diagramOther.NoteTaskShapes[i] == nil) {
				NoteTaskShapesDifferent = true
				break
			} else if diagram.NoteTaskShapes[i] != nil && diagramOther.NoteTaskShapes[i] != nil {
				if len(diagram.NoteTaskShapes[i].GongDiff(diagramOther.NoteTaskShapes[i])) > 0 {
					NoteTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteTaskShapesDifferent {
		diffs = append(diffs, "NoteTaskShapes")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, "Name")
	}
	ProductsDifferent := false
	if len(note.Products) != len(noteOther.Products) {
		ProductsDifferent = true
	} else {
		for i := range note.Products {
			if (note.Products[i] == nil) != (noteOther.Products[i] == nil) {
				ProductsDifferent = true
				break
			} else if note.Products[i] != nil && noteOther.Products[i] != nil {
				if len(note.Products[i].GongDiff(noteOther.Products[i])) > 0 {
					ProductsDifferent = true
					break
				}
			}
		}
	}
	if ProductsDifferent {
		diffs = append(diffs, "Products")
	}
	TasksDifferent := false
	if len(note.Tasks) != len(noteOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range note.Tasks {
			if (note.Tasks[i] == nil) != (noteOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if note.Tasks[i] != nil && noteOther.Tasks[i] != nil {
				if len(note.Tasks[i].GongDiff(noteOther.Tasks[i])) > 0 {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		diffs = append(diffs, "Tasks")
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteproductshape *NoteProductShape) GongDiff(noteproductshapeOther *NoteProductShape) (diffs []string) {
	// insertion point for field diffs
	if noteproductshape.Name != noteproductshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (noteproductshape.Note == nil) != (noteproductshapeOther.Note == nil) {
		diffs = append(diffs, "Note")
	} else if noteproductshape.Note != nil && noteproductshapeOther.Note != nil {
		if noteproductshape.Note != noteproductshapeOther.Note {
			diffs = append(diffs, "Note")
		}
	}
	if (noteproductshape.Product == nil) != (noteproductshapeOther.Product == nil) {
		diffs = append(diffs, "Product")
	} else if noteproductshape.Product != nil && noteproductshapeOther.Product != nil {
		if noteproductshape.Product != noteproductshapeOther.Product {
			diffs = append(diffs, "Product")
		}
	}
	if noteproductshape.StartRatio != noteproductshapeOther.StartRatio {
		diffs = append(diffs, "StartRatio")
	}
	if noteproductshape.EndRatio != noteproductshapeOther.EndRatio {
		diffs = append(diffs, "EndRatio")
	}
	if noteproductshape.StartOrientation != noteproductshapeOther.StartOrientation {
		diffs = append(diffs, "StartOrientation")
	}
	if noteproductshape.EndOrientation != noteproductshapeOther.EndOrientation {
		diffs = append(diffs, "EndOrientation")
	}
	if noteproductshape.CornerOffsetRatio != noteproductshapeOther.CornerOffsetRatio {
		diffs = append(diffs, "CornerOffsetRatio")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, "Note")
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, "Note")
		}
	}
	if noteshape.IsExpanded != noteshapeOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, "X")
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, "Y")
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, "Width")
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, "Height")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetaskshape *NoteTaskShape) GongDiff(notetaskshapeOther *NoteTaskShape) (diffs []string) {
	// insertion point for field diffs
	if notetaskshape.Name != notetaskshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (notetaskshape.Note == nil) != (notetaskshapeOther.Note == nil) {
		diffs = append(diffs, "Note")
	} else if notetaskshape.Note != nil && notetaskshapeOther.Note != nil {
		if notetaskshape.Note != notetaskshapeOther.Note {
			diffs = append(diffs, "Note")
		}
	}
	if (notetaskshape.Task == nil) != (notetaskshapeOther.Task == nil) {
		diffs = append(diffs, "Task")
	} else if notetaskshape.Task != nil && notetaskshapeOther.Task != nil {
		if notetaskshape.Task != notetaskshapeOther.Task {
			diffs = append(diffs, "Task")
		}
	}
	if notetaskshape.StartRatio != notetaskshapeOther.StartRatio {
		diffs = append(diffs, "StartRatio")
	}
	if notetaskshape.EndRatio != notetaskshapeOther.EndRatio {
		diffs = append(diffs, "EndRatio")
	}
	if notetaskshape.StartOrientation != notetaskshapeOther.StartOrientation {
		diffs = append(diffs, "StartOrientation")
	}
	if notetaskshape.EndOrientation != notetaskshapeOther.EndOrientation {
		diffs = append(diffs, "EndOrientation")
	}
	if notetaskshape.CornerOffsetRatio != notetaskshapeOther.CornerOffsetRatio {
		diffs = append(diffs, "CornerOffsetRatio")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (product *Product) GongDiff(productOther *Product) (diffs []string) {
	// insertion point for field diffs
	if product.Name != productOther.Name {
		diffs = append(diffs, "Name")
	}
	if product.Description != productOther.Description {
		diffs = append(diffs, "Description")
	}
	SubProductsDifferent := false
	if len(product.SubProducts) != len(productOther.SubProducts) {
		SubProductsDifferent = true
	} else {
		for i := range product.SubProducts {
			if (product.SubProducts[i] == nil) != (productOther.SubProducts[i] == nil) {
				SubProductsDifferent = true
				break
			} else if product.SubProducts[i] != nil && productOther.SubProducts[i] != nil {
				if len(product.SubProducts[i].GongDiff(productOther.SubProducts[i])) > 0 {
					SubProductsDifferent = true
					break
				}
			}
		}
	}
	if SubProductsDifferent {
		diffs = append(diffs, "SubProducts")
	}
	if product.IsExpanded != productOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if product.ComputedPrefix != productOther.ComputedPrefix {
		diffs = append(diffs, "ComputedPrefix")
	}
	if product.IsProducersNodeExpanded != productOther.IsProducersNodeExpanded {
		diffs = append(diffs, "IsProducersNodeExpanded")
	}
	if product.IsConsumersNodeExpanded != productOther.IsConsumersNodeExpanded {
		diffs = append(diffs, "IsConsumersNodeExpanded")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productcompositionshape *ProductCompositionShape) GongDiff(productcompositionshapeOther *ProductCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if productcompositionshape.Name != productcompositionshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (productcompositionshape.Product == nil) != (productcompositionshapeOther.Product == nil) {
		diffs = append(diffs, "Product")
	} else if productcompositionshape.Product != nil && productcompositionshapeOther.Product != nil {
		if productcompositionshape.Product != productcompositionshapeOther.Product {
			diffs = append(diffs, "Product")
		}
	}
	if productcompositionshape.StartRatio != productcompositionshapeOther.StartRatio {
		diffs = append(diffs, "StartRatio")
	}
	if productcompositionshape.EndRatio != productcompositionshapeOther.EndRatio {
		diffs = append(diffs, "EndRatio")
	}
	if productcompositionshape.StartOrientation != productcompositionshapeOther.StartOrientation {
		diffs = append(diffs, "StartOrientation")
	}
	if productcompositionshape.EndOrientation != productcompositionshapeOther.EndOrientation {
		diffs = append(diffs, "EndOrientation")
	}
	if productcompositionshape.CornerOffsetRatio != productcompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, "CornerOffsetRatio")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productshape *ProductShape) GongDiff(productshapeOther *ProductShape) (diffs []string) {
	// insertion point for field diffs
	if productshape.Name != productshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (productshape.Product == nil) != (productshapeOther.Product == nil) {
		diffs = append(diffs, "Product")
	} else if productshape.Product != nil && productshapeOther.Product != nil {
		if productshape.Product != productshapeOther.Product {
			diffs = append(diffs, "Product")
		}
	}
	if productshape.IsExpanded != productshapeOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if productshape.X != productshapeOther.X {
		diffs = append(diffs, "X")
	}
	if productshape.Y != productshapeOther.Y {
		diffs = append(diffs, "Y")
	}
	if productshape.Width != productshapeOther.Width {
		diffs = append(diffs, "Width")
	}
	if productshape.Height != productshapeOther.Height {
		diffs = append(diffs, "Height")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (project *Project) GongDiff(projectOther *Project) (diffs []string) {
	// insertion point for field diffs
	if project.Name != projectOther.Name {
		diffs = append(diffs, "Name")
	}
	RootProductsDifferent := false
	if len(project.RootProducts) != len(projectOther.RootProducts) {
		RootProductsDifferent = true
	} else {
		for i := range project.RootProducts {
			if (project.RootProducts[i] == nil) != (projectOther.RootProducts[i] == nil) {
				RootProductsDifferent = true
				break
			} else if project.RootProducts[i] != nil && projectOther.RootProducts[i] != nil {
				if len(project.RootProducts[i].GongDiff(projectOther.RootProducts[i])) > 0 {
					RootProductsDifferent = true
					break
				}
			}
		}
	}
	if RootProductsDifferent {
		diffs = append(diffs, "RootProducts")
	}
	if project.IsPBSNodeExpanded != projectOther.IsPBSNodeExpanded {
		diffs = append(diffs, "IsPBSNodeExpanded")
	}
	RootTasksDifferent := false
	if len(project.RootTasks) != len(projectOther.RootTasks) {
		RootTasksDifferent = true
	} else {
		for i := range project.RootTasks {
			if (project.RootTasks[i] == nil) != (projectOther.RootTasks[i] == nil) {
				RootTasksDifferent = true
				break
			} else if project.RootTasks[i] != nil && projectOther.RootTasks[i] != nil {
				if len(project.RootTasks[i].GongDiff(projectOther.RootTasks[i])) > 0 {
					RootTasksDifferent = true
					break
				}
			}
		}
	}
	if RootTasksDifferent {
		diffs = append(diffs, "RootTasks")
	}
	if project.IsWBSNodeExpanded != projectOther.IsWBSNodeExpanded {
		diffs = append(diffs, "IsWBSNodeExpanded")
	}
	DiagramsDifferent := false
	if len(project.Diagrams) != len(projectOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range project.Diagrams {
			if (project.Diagrams[i] == nil) != (projectOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if project.Diagrams[i] != nil && projectOther.Diagrams[i] != nil {
				if len(project.Diagrams[i].GongDiff(projectOther.Diagrams[i])) > 0 {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		diffs = append(diffs, "Diagrams")
	}
	if project.IsDiagramsNodeExpanded != projectOther.IsDiagramsNodeExpanded {
		diffs = append(diffs, "IsDiagramsNodeExpanded")
	}
	NotesDifferent := false
	if len(project.Notes) != len(projectOther.Notes) {
		NotesDifferent = true
	} else {
		for i := range project.Notes {
			if (project.Notes[i] == nil) != (projectOther.Notes[i] == nil) {
				NotesDifferent = true
				break
			} else if project.Notes[i] != nil && projectOther.Notes[i] != nil {
				if len(project.Notes[i].GongDiff(projectOther.Notes[i])) > 0 {
					NotesDifferent = true
					break
				}
			}
		}
	}
	if NotesDifferent {
		diffs = append(diffs, "Notes")
	}
	if project.IsNotesNodeExpanded != projectOther.IsNotesNodeExpanded {
		diffs = append(diffs, "IsNotesNodeExpanded")
	}
	if project.IsExpanded != projectOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if project.ComputedPrefix != projectOther.ComputedPrefix {
		diffs = append(diffs, "ComputedPrefix")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (root *Root) GongDiff(rootOther *Root) (diffs []string) {
	// insertion point for field diffs
	if root.Name != rootOther.Name {
		diffs = append(diffs, "Name")
	}
	ProjectsDifferent := false
	if len(root.Projects) != len(rootOther.Projects) {
		ProjectsDifferent = true
	} else {
		for i := range root.Projects {
			if (root.Projects[i] == nil) != (rootOther.Projects[i] == nil) {
				ProjectsDifferent = true
				break
			} else if root.Projects[i] != nil && rootOther.Projects[i] != nil {
				if len(root.Projects[i].GongDiff(rootOther.Projects[i])) > 0 {
					ProjectsDifferent = true
					break
				}
			}
		}
	}
	if ProjectsDifferent {
		diffs = append(diffs, "Projects")
	}
	OrphanedProductsDifferent := false
	if len(root.OrphanedProducts) != len(rootOther.OrphanedProducts) {
		OrphanedProductsDifferent = true
	} else {
		for i := range root.OrphanedProducts {
			if (root.OrphanedProducts[i] == nil) != (rootOther.OrphanedProducts[i] == nil) {
				OrphanedProductsDifferent = true
				break
			} else if root.OrphanedProducts[i] != nil && rootOther.OrphanedProducts[i] != nil {
				if len(root.OrphanedProducts[i].GongDiff(rootOther.OrphanedProducts[i])) > 0 {
					OrphanedProductsDifferent = true
					break
				}
			}
		}
	}
	if OrphanedProductsDifferent {
		diffs = append(diffs, "OrphanedProducts")
	}
	OrphanedTasksDifferent := false
	if len(root.OrphanedTasks) != len(rootOther.OrphanedTasks) {
		OrphanedTasksDifferent = true
	} else {
		for i := range root.OrphanedTasks {
			if (root.OrphanedTasks[i] == nil) != (rootOther.OrphanedTasks[i] == nil) {
				OrphanedTasksDifferent = true
				break
			} else if root.OrphanedTasks[i] != nil && rootOther.OrphanedTasks[i] != nil {
				if len(root.OrphanedTasks[i].GongDiff(rootOther.OrphanedTasks[i])) > 0 {
					OrphanedTasksDifferent = true
					break
				}
			}
		}
	}
	if OrphanedTasksDifferent {
		diffs = append(diffs, "OrphanedTasks")
	}
	if root.NbPixPerCharacter != rootOther.NbPixPerCharacter {
		diffs = append(diffs, "NbPixPerCharacter")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (task *Task) GongDiff(taskOther *Task) (diffs []string) {
	// insertion point for field diffs
	if task.Name != taskOther.Name {
		diffs = append(diffs, "Name")
	}
	if task.Description != taskOther.Description {
		diffs = append(diffs, "Description")
	}
	SubTasksDifferent := false
	if len(task.SubTasks) != len(taskOther.SubTasks) {
		SubTasksDifferent = true
	} else {
		for i := range task.SubTasks {
			if (task.SubTasks[i] == nil) != (taskOther.SubTasks[i] == nil) {
				SubTasksDifferent = true
				break
			} else if task.SubTasks[i] != nil && taskOther.SubTasks[i] != nil {
				if len(task.SubTasks[i].GongDiff(taskOther.SubTasks[i])) > 0 {
					SubTasksDifferent = true
					break
				}
			}
		}
	}
	if SubTasksDifferent {
		diffs = append(diffs, "SubTasks")
	}
	if task.IsExpanded != taskOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if task.ComputedPrefix != taskOther.ComputedPrefix {
		diffs = append(diffs, "ComputedPrefix")
	}
	InputsDifferent := false
	if len(task.Inputs) != len(taskOther.Inputs) {
		InputsDifferent = true
	} else {
		for i := range task.Inputs {
			if (task.Inputs[i] == nil) != (taskOther.Inputs[i] == nil) {
				InputsDifferent = true
				break
			} else if task.Inputs[i] != nil && taskOther.Inputs[i] != nil {
				if len(task.Inputs[i].GongDiff(taskOther.Inputs[i])) > 0 {
					InputsDifferent = true
					break
				}
			}
		}
	}
	if InputsDifferent {
		diffs = append(diffs, "Inputs")
	}
	if task.IsInputsNodeExpanded != taskOther.IsInputsNodeExpanded {
		diffs = append(diffs, "IsInputsNodeExpanded")
	}
	OutputsDifferent := false
	if len(task.Outputs) != len(taskOther.Outputs) {
		OutputsDifferent = true
	} else {
		for i := range task.Outputs {
			if (task.Outputs[i] == nil) != (taskOther.Outputs[i] == nil) {
				OutputsDifferent = true
				break
			} else if task.Outputs[i] != nil && taskOther.Outputs[i] != nil {
				if len(task.Outputs[i].GongDiff(taskOther.Outputs[i])) > 0 {
					OutputsDifferent = true
					break
				}
			}
		}
	}
	if OutputsDifferent {
		diffs = append(diffs, "Outputs")
	}
	if task.IsOutputsNodeExpanded != taskOther.IsOutputsNodeExpanded {
		diffs = append(diffs, "IsOutputsNodeExpanded")
	}
	if task.IsWithCompletion != taskOther.IsWithCompletion {
		diffs = append(diffs, "IsWithCompletion")
	}
	if task.Completion != taskOther.Completion {
		diffs = append(diffs, "Completion")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskcompositionshape *TaskCompositionShape) GongDiff(taskcompositionshapeOther *TaskCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if taskcompositionshape.Name != taskcompositionshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (taskcompositionshape.Task == nil) != (taskcompositionshapeOther.Task == nil) {
		diffs = append(diffs, "Task")
	} else if taskcompositionshape.Task != nil && taskcompositionshapeOther.Task != nil {
		if taskcompositionshape.Task != taskcompositionshapeOther.Task {
			diffs = append(diffs, "Task")
		}
	}
	if taskcompositionshape.StartRatio != taskcompositionshapeOther.StartRatio {
		diffs = append(diffs, "StartRatio")
	}
	if taskcompositionshape.EndRatio != taskcompositionshapeOther.EndRatio {
		diffs = append(diffs, "EndRatio")
	}
	if taskcompositionshape.StartOrientation != taskcompositionshapeOther.StartOrientation {
		diffs = append(diffs, "StartOrientation")
	}
	if taskcompositionshape.EndOrientation != taskcompositionshapeOther.EndOrientation {
		diffs = append(diffs, "EndOrientation")
	}
	if taskcompositionshape.CornerOffsetRatio != taskcompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, "CornerOffsetRatio")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskinputshape *TaskInputShape) GongDiff(taskinputshapeOther *TaskInputShape) (diffs []string) {
	// insertion point for field diffs
	if taskinputshape.Name != taskinputshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (taskinputshape.Task == nil) != (taskinputshapeOther.Task == nil) {
		diffs = append(diffs, "Task")
	} else if taskinputshape.Task != nil && taskinputshapeOther.Task != nil {
		if taskinputshape.Task != taskinputshapeOther.Task {
			diffs = append(diffs, "Task")
		}
	}
	if (taskinputshape.Product == nil) != (taskinputshapeOther.Product == nil) {
		diffs = append(diffs, "Product")
	} else if taskinputshape.Product != nil && taskinputshapeOther.Product != nil {
		if taskinputshape.Product != taskinputshapeOther.Product {
			diffs = append(diffs, "Product")
		}
	}
	if taskinputshape.StartRatio != taskinputshapeOther.StartRatio {
		diffs = append(diffs, "StartRatio")
	}
	if taskinputshape.EndRatio != taskinputshapeOther.EndRatio {
		diffs = append(diffs, "EndRatio")
	}
	if taskinputshape.StartOrientation != taskinputshapeOther.StartOrientation {
		diffs = append(diffs, "StartOrientation")
	}
	if taskinputshape.EndOrientation != taskinputshapeOther.EndOrientation {
		diffs = append(diffs, "EndOrientation")
	}
	if taskinputshape.CornerOffsetRatio != taskinputshapeOther.CornerOffsetRatio {
		diffs = append(diffs, "CornerOffsetRatio")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskoutputshape *TaskOutputShape) GongDiff(taskoutputshapeOther *TaskOutputShape) (diffs []string) {
	// insertion point for field diffs
	if taskoutputshape.Name != taskoutputshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (taskoutputshape.Task == nil) != (taskoutputshapeOther.Task == nil) {
		diffs = append(diffs, "Task")
	} else if taskoutputshape.Task != nil && taskoutputshapeOther.Task != nil {
		if taskoutputshape.Task != taskoutputshapeOther.Task {
			diffs = append(diffs, "Task")
		}
	}
	if (taskoutputshape.Product == nil) != (taskoutputshapeOther.Product == nil) {
		diffs = append(diffs, "Product")
	} else if taskoutputshape.Product != nil && taskoutputshapeOther.Product != nil {
		if taskoutputshape.Product != taskoutputshapeOther.Product {
			diffs = append(diffs, "Product")
		}
	}
	if taskoutputshape.StartRatio != taskoutputshapeOther.StartRatio {
		diffs = append(diffs, "StartRatio")
	}
	if taskoutputshape.EndRatio != taskoutputshapeOther.EndRatio {
		diffs = append(diffs, "EndRatio")
	}
	if taskoutputshape.StartOrientation != taskoutputshapeOther.StartOrientation {
		diffs = append(diffs, "StartOrientation")
	}
	if taskoutputshape.EndOrientation != taskoutputshapeOther.EndOrientation {
		diffs = append(diffs, "EndOrientation")
	}
	if taskoutputshape.CornerOffsetRatio != taskoutputshapeOther.CornerOffsetRatio {
		diffs = append(diffs, "CornerOffsetRatio")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskshape *TaskShape) GongDiff(taskshapeOther *TaskShape) (diffs []string) {
	// insertion point for field diffs
	if taskshape.Name != taskshapeOther.Name {
		diffs = append(diffs, "Name")
	}
	if (taskshape.Task == nil) != (taskshapeOther.Task == nil) {
		diffs = append(diffs, "Task")
	} else if taskshape.Task != nil && taskshapeOther.Task != nil {
		if taskshape.Task != taskshapeOther.Task {
			diffs = append(diffs, "Task")
		}
	}
	if taskshape.IsExpanded != taskshapeOther.IsExpanded {
		diffs = append(diffs, "IsExpanded")
	}
	if taskshape.X != taskshapeOther.X {
		diffs = append(diffs, "X")
	}
	if taskshape.Y != taskshapeOther.Y {
		diffs = append(diffs, "Y")
	}
	if taskshape.Width != taskshapeOther.Width {
		diffs = append(diffs, "Width")
	}
	if taskshape.Height != taskshapeOther.Height {
		diffs = append(diffs, "Height")
	}

	return
}

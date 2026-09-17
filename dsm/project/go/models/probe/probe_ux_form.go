// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Diagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagram := new(models.Diagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagram, formGroup, probe)
	case "Library":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			probe,
			formGroup,
		)
		library := new(models.Library)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(library, formGroup, probe)
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "NoteProductShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteProductShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteproductshape := new(models.NoteProductShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteproductshape, formGroup, probe)
	case "NoteResourceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteResourceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteResourceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteresourceshape := new(models.NoteResourceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteresourceshape, formGroup, probe)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteshape := new(models.NoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshape, formGroup, probe)
	case "NoteTaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteTaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		notetaskshape := new(models.NoteTaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notetaskshape, formGroup, probe)
	case "Product":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Product Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductFormCallback(
			nil,
			probe,
			formGroup,
		)
		product := new(models.Product)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(product, formGroup, probe)
	case "ProductCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProductCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		productcompositionshape := new(models.ProductCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(productcompositionshape, formGroup, probe)
	case "ProductShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		productshape := new(models.ProductShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(productshape, formGroup, probe)
	case "Resource":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Resource Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceFormCallback(
			nil,
			probe,
			formGroup,
		)
		resource := new(models.Resource)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resource, formGroup, probe)
	case "ResourceCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ResourceCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		resourcecompositionshape := new(models.ResourceCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resourcecompositionshape, formGroup, probe)
	case "ResourceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ResourceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		resourceshape := new(models.ResourceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resourceshape, formGroup, probe)
	case "ResourceTaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ResourceTaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceTaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		resourcetaskshape := new(models.ResourceTaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resourcetaskshape, formGroup, probe)
	case "Task":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Task Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			probe,
			formGroup,
		)
		task := new(models.Task)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(task, formGroup, probe)
	case "TaskCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskcompositionshape := new(models.TaskCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskcompositionshape, formGroup, probe)
	case "TaskGroup":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskGroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskgroup := new(models.TaskGroup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskgroup, formGroup, probe)
	case "TaskGroupShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskGroupShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskGroupShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskgroupshape := new(models.TaskGroupShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskgroupshape, formGroup, probe)
	case "TaskInputShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskInputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskInputShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskinputshape := new(models.TaskInputShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskinputshape, formGroup, probe)
	case "TaskOutputShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskOutputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskOutputShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskoutputshape := new(models.TaskOutputShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskoutputshape, formGroup, probe)
	case "TaskPredecessorShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskPredecessorShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskPredecessorShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskpredecessorshape := new(models.TaskPredecessorShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskpredecessorshape, formGroup, probe)
	case "TaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskshape := new(models.TaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskshape, formGroup, probe)
	}
	formStage.Commit()
}

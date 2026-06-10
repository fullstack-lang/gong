// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Diagram:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteProductShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteProductShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteResourceShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteResourceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteResourceShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteTaskShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteTaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Product:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Product Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ProductCompositionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ProductCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductCompositionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ProductShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ProductShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProductShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Resource:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Resource Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ResourceCompositionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ResourceCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceCompositionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ResourceShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ResourceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ResourceTaskShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ResourceTaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceTaskShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Task:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Task Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskCompositionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskCompositionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskGroup:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskGroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskGroupShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskGroupShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskGroupShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskInputShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskInputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskInputShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskOutputShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskOutputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskOutputShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}

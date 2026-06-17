// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AnalysisNeed:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AnalysisNeed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnalysisNeedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Concept:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Concept Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConceptFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Concern:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Concern Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ConcernCompositionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ConcernCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernCompositionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ConcernInputShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ConcernInputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernInputShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ConcernOutputShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ConcernOutputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernOutputShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ConcernShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ConcernShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Deliverable:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Deliverable Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
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
	case *models.NoteStakeholderShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteStakeholderShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteStakeholderShapeFormCallback(
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
	case *models.Requirement:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Requirement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Stakeholder:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Stakeholder Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StakeholderCompositionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "StakeholderCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderCompositionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StakeholderConcernShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "StakeholderConcernShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderConcernShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StakeholderShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "StakeholderShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SupportLevel:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SupportLevel Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SupportLevelFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tool:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tool Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ToolFormCallback(
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

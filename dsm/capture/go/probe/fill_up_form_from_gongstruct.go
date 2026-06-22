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
			Label: "AnalysisNeed",
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
			Label: "Concept",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConceptFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ConceptShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ConceptShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConceptShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Concern:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Concern",
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
			Label: "ConcernCompositionShape",
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
			Label: "ConcernInputShape",
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
			Label: "ConcernOutputShape",
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
			Label: "ConcernShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ControlPointShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ControlPointShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Deliverable:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Deliverable",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DeliverableCompositionShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DeliverableCompositionShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableCompositionShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DeliverableConceptShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DeliverableConceptShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableConceptShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DeliverableShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DeliverableShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Diagram:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Diagram",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DiagramShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DiagramShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Library",
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
			Label: "Note",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteDeliverableShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteDeliverableShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteDeliverableShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteShape",
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
			Label: "NoteStakeholderShape",
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
			Label: "NoteTaskShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Requirement:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Requirement",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RequirementShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "RequirementShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Stakeholder:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Stakeholder",
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
			Label: "StakeholderCompositionShape",
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
			Label: "StakeholderConcernShape",
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
			Label: "StakeholderShape",
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
			Label: "SupportLevel",
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
			Label: "Tool",
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

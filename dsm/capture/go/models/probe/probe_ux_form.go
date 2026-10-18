// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
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
	case "AnalysisNeed":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AnalysisNeed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnalysisNeedFormCallback(
			nil,
			probe,
			formGroup,
		)
		analysisneed := new(models.AnalysisNeed)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(analysisneed, formGroup, probe)
	case "Concept":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Concept Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConceptFormCallback(
			nil,
			probe,
			formGroup,
		)
		concept := new(models.Concept)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(concept, formGroup, probe)
	case "ConceptShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ConceptShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConceptShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		conceptshape := new(models.ConceptShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(conceptshape, formGroup, probe)
	case "Concern":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Concern Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernFormCallback(
			nil,
			probe,
			formGroup,
		)
		concern := new(models.Concern)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(concern, formGroup, probe)
	case "ConcernCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ConcernCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		concerncompositionshape := new(models.ConcernCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(concerncompositionshape, formGroup, probe)
	case "ConcernInputShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ConcernInputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernInputShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		concerninputshape := new(models.ConcernInputShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(concerninputshape, formGroup, probe)
	case "ConcernOutputShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ConcernOutputShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernOutputShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		concernoutputshape := new(models.ConcernOutputShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(concernoutputshape, formGroup, probe)
	case "ConcernShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ConcernShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ConcernShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		concernshape := new(models.ConcernShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(concernshape, formGroup, probe)
	case "ControlPointShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlPointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlpointshape, formGroup, probe)
	case "Deliverable":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Deliverable Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableFormCallback(
			nil,
			probe,
			formGroup,
		)
		deliverable := new(models.Deliverable)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(deliverable, formGroup, probe)
	case "DeliverableCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DeliverableCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		deliverablecompositionshape := new(models.DeliverableCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(deliverablecompositionshape, formGroup, probe)
	case "DeliverableConceptShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DeliverableConceptShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableConceptShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		deliverableconceptshape := new(models.DeliverableConceptShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(deliverableconceptshape, formGroup, probe)
	case "DeliverableShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DeliverableShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		deliverableshape := new(models.DeliverableShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(deliverableshape, formGroup, probe)
	case "DeliverableViewShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DeliverableViewShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeliverableViewShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		deliverableviewshape := new(models.DeliverableViewShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(deliverableviewshape, formGroup, probe)
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
	case "DiagramShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagramshape := new(models.DiagramShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagramshape, formGroup, probe)
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
	case "NoteDeliverableShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteDeliverableShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteDeliverableShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		notedeliverableshape := new(models.NoteDeliverableShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notedeliverableshape, formGroup, probe)
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
	case "NoteStakeholderShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteStakeholderShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteStakeholderShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		notestakeholdershape := new(models.NoteStakeholderShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notestakeholdershape, formGroup, probe)
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
	case "Requirement":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Requirement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementFormCallback(
			nil,
			probe,
			formGroup,
		)
		requirement := new(models.Requirement)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(requirement, formGroup, probe)
	case "RequirementConceptShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RequirementConceptShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementConceptShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		requirementconceptshape := new(models.RequirementConceptShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(requirementconceptshape, formGroup, probe)
	case "RequirementConcernShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RequirementConcernShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementConcernShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		requirementconcernshape := new(models.RequirementConcernShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(requirementconcernshape, formGroup, probe)
	case "RequirementShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RequirementShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		requirementshape := new(models.RequirementShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(requirementshape, formGroup, probe)
	case "RequirementViewShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RequirementViewShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RequirementViewShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		requirementviewshape := new(models.RequirementViewShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(requirementviewshape, formGroup, probe)
	case "Stakeholder":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Stakeholder Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderFormCallback(
			nil,
			probe,
			formGroup,
		)
		stakeholder := new(models.Stakeholder)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stakeholder, formGroup, probe)
	case "StakeholderCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StakeholderCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stakeholdercompositionshape := new(models.StakeholderCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stakeholdercompositionshape, formGroup, probe)
	case "StakeholderConcernShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StakeholderConcernShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderConcernShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stakeholderconcernshape := new(models.StakeholderConcernShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stakeholderconcernshape, formGroup, probe)
	case "StakeholderShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "StakeholderShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StakeholderShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		stakeholdershape := new(models.StakeholderShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stakeholdershape, formGroup, probe)
	case "SupportLevel":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SupportLevel Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SupportLevelFormCallback(
			nil,
			probe,
			formGroup,
		)
		supportlevel := new(models.SupportLevel)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(supportlevel, formGroup, probe)
	case "Tool":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tool Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ToolFormCallback(
			nil,
			probe,
			formGroup,
		)
		tool := new(models.Tool)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tool, formGroup, probe)
	case "View":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "View Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ViewFormCallback(
			nil,
			probe,
			formGroup,
		)
		view := new(models.View)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(view, formGroup, probe)
	case "ViewConceptShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ViewConceptShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ViewConceptShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		viewconceptshape := new(models.ViewConceptShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(viewconceptshape, formGroup, probe)
	case "ViewShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ViewShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ViewShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		viewshape := new(models.ViewShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(viewshape, formGroup, probe)
	}
	formStage.Commit()
}

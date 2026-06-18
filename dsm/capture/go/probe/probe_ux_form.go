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
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AnalysisNeedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AnalysisNeed", true)
			} else {
				FillUpFormFromGongstruct(onSave.analysisneed, probe)
			}
		case *ConceptFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Concept", true)
			} else {
				FillUpFormFromGongstruct(onSave.concept, probe)
			}
		case *ConceptShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ConceptShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.conceptshape, probe)
			}
		case *ConcernFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Concern", true)
			} else {
				FillUpFormFromGongstruct(onSave.concern, probe)
			}
		case *ConcernCompositionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ConcernCompositionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.concerncompositionshape, probe)
			}
		case *ConcernInputShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ConcernInputShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.concerninputshape, probe)
			}
		case *ConcernOutputShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ConcernOutputShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.concernoutputshape, probe)
			}
		case *ConcernShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ConcernShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.concernshape, probe)
			}
		case *ControlPointShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ControlPointShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.controlpointshape, probe)
			}
		case *DeliverableFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Deliverable", true)
			} else {
				FillUpFormFromGongstruct(onSave.deliverable, probe)
			}
		case *DeliverableCompositionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DeliverableCompositionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.deliverablecompositionshape, probe)
			}
		case *DeliverableConceptShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DeliverableConceptShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.deliverableconceptshape, probe)
			}
		case *DeliverableShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DeliverableShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.deliverableshape, probe)
			}
		case *DiagramFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Diagram", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagram, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *NoteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note", true)
			} else {
				FillUpFormFromGongstruct(onSave.note, probe)
			}
		case *NoteDeliverableShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteDeliverableShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.notedeliverableshape, probe)
			}
		case *NoteShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteshape, probe)
			}
		case *NoteStakeholderShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteStakeholderShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.notestakeholdershape, probe)
			}
		case *NoteTaskShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteTaskShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.notetaskshape, probe)
			}
		case *RequirementFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Requirement", true)
			} else {
				FillUpFormFromGongstruct(onSave.requirement, probe)
			}
		case *RequirementShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "RequirementShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.requirementshape, probe)
			}
		case *StakeholderFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Stakeholder", true)
			} else {
				FillUpFormFromGongstruct(onSave.stakeholder, probe)
			}
		case *StakeholderCompositionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StakeholderCompositionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stakeholdercompositionshape, probe)
			}
		case *StakeholderConcernShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StakeholderConcernShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stakeholderconcernshape, probe)
			}
		case *StakeholderShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "StakeholderShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.stakeholdershape, probe)
			}
		case *SupportLevelFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SupportLevel", true)
			} else {
				FillUpFormFromGongstruct(onSave.supportlevel, probe)
			}
		case *ToolFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tool", true)
			} else {
				FillUpFormFromGongstruct(onSave.tool, probe)
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
	}
	formStage.Commit()
}

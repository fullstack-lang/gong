// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func StageSetFillUpFormFromGongstruct(
	instance any,
	probe *StageSetProbe,
) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name: "Form",
	}).Stage(probe.formStage)

	switch inst := instance.(type) {
	case *models.AnalysisNeed:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AnalysisNeed_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_AnalysisNeed_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Concept:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Concept_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Concept_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ConceptShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConceptShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ConceptShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Concern:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Concern_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Concern_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ConcernCompositionShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConcernCompositionShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ConcernCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ConcernInputShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConcernInputShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ConcernInputShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ConcernOutputShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConcernOutputShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ConcernOutputShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ConcernShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConcernShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ConcernShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ControlPointShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ControlPointShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ControlPointShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Deliverable:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Deliverable_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Deliverable_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DeliverableCompositionShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeliverableCompositionShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DeliverableCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DeliverableConceptShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeliverableConceptShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DeliverableConceptShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DeliverableShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeliverableShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DeliverableShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Diagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Diagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Diagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DiagramShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DiagramShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DiagramShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Library:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Library_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Library_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Note:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Note_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Note_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteDeliverableShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteDeliverableShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteDeliverableShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteStakeholderShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteStakeholderShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteStakeholderShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteTaskShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteTaskShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteTaskShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Requirement:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Requirement_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Requirement_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RequirementShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RequirementShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RequirementShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Stakeholder:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stakeholder_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Stakeholder_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StakeholderCompositionShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StakeholderCompositionShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StakeholderCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StakeholderConcernShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StakeholderConcernShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StakeholderConcernShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StakeholderShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StakeholderShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StakeholderShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SupportLevel:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SupportLevel_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SupportLevel_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Tool:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tool_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Tool_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	default:
		_ = inst
	}

	probe.formStage.Commit()
}

func saveStageSet_AnalysisNeed_Stage(
	inst *models.AnalysisNeed,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_Concept_Stage(
	inst *models.Concept,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_ConceptShape_Stage(
	inst *models.ConceptShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Concept":
			StageSetFormDivSelectFieldToField(&inst.Concept, probe.stageSet.Stage.GetInstancesSet[*models.Concept](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Concern_Stage(
	inst *models.Concern,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IDAirbus":
			FormDivBasicFieldToField(&inst.IDAirbus, formDiv)
		case "Priority":
			StageSetFormDivEnumStringFieldToField(&inst.Priority, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "IsInputsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsInputsNodeExpanded, formDiv)
		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsOutputsNodeExpanded, formDiv)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&inst.IsWithCompletion, formDiv)
		case "Completion":
			StageSetFormDivEnumStringFieldToField(&inst.Completion, formDiv)
		}
	}
}

func saveStageSet_ConcernCompositionShape_Stage(
	inst *models.ConcernCompositionShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Concern":
			StageSetFormDivSelectFieldToField(&inst.Concern, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ConcernInputShape_Stage(
	inst *models.ConcernInputShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Deliverable":
			StageSetFormDivSelectFieldToField(&inst.Deliverable, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), formDiv)
		case "Concern":
			StageSetFormDivSelectFieldToField(&inst.Concern, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ConcernOutputShape_Stage(
	inst *models.ConcernOutputShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Concern":
			StageSetFormDivSelectFieldToField(&inst.Concern, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), formDiv)
		case "Deliverable":
			StageSetFormDivSelectFieldToField(&inst.Deliverable, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ConcernShape_Stage(
	inst *models.ConcernShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Concern":
			StageSetFormDivSelectFieldToField(&inst.Concern, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ControlPointShape_Stage(
	inst *models.ControlPointShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&inst.X_Relative, formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&inst.Y_Relative, formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&inst.IsStartShapeTheClosestShape, formDiv)
		}
	}
}

func saveStageSet_Deliverable_Stage(
	inst *models.Deliverable,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&inst.IsProducersNodeExpanded, formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&inst.IsConsumersNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_DeliverableCompositionShape_Stage(
	inst *models.DeliverableCompositionShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Deliverable":
			StageSetFormDivSelectFieldToField(&inst.Deliverable, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_DeliverableConceptShape_Stage(
	inst *models.DeliverableConceptShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Deliverable":
			StageSetFormDivSelectFieldToField(&inst.Deliverable, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), formDiv)
		case "Concept":
			StageSetFormDivSelectFieldToField(&inst.Concept, probe.stageSet.Stage.GetInstancesSet[*models.Concept](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_DeliverableShape_Stage(
	inst *models.DeliverableShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Deliverable":
			StageSetFormDivSelectFieldToField(&inst.Deliverable, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Diagram_Stage(
	inst *models.Diagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&inst.IsEditable_, formDiv)
		case "ShowPrefix":
			FormDivBasicFieldToField(&inst.ShowPrefix, formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&inst.DefaultBoxWidth, formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&inst.DefaultBoxHeigth, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsRequirementsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsRequirementsNodeExpanded, formDiv)
		case "IsConceptsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsConceptsNodeExpanded, formDiv)
		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&inst.IsPBSNodeExpanded, formDiv)
		case "IsConcernsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsConcernsNodeExpanded, formDiv)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNotesNodeExpanded, formDiv)
		case "IsStakeholdersNodeExpanded":
			FormDivBasicFieldToField(&inst.IsStakeholdersNodeExpanded, formDiv)
		case "IsDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsDiagramsNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_DiagramShape_Stage(
	inst *models.DiagramShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Diagram":
			StageSetFormDivSelectFieldToField(&inst.Diagram, probe.stageSet.Stage.GetInstancesSet[*models.Diagram](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Library_Stage(
	inst *models.Library,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&inst.IsRootLibrary, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&inst.NbPixPerCharacter, formDiv)
		}
	}
}

func saveStageSet_Note_Stage(
	inst *models.Note,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_NoteDeliverableShape_Stage(
	inst *models.NoteDeliverableShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "Deliverable":
			StageSetFormDivSelectFieldToField(&inst.Deliverable, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_NoteShape_Stage(
	inst *models.NoteShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_NoteStakeholderShape_Stage(
	inst *models.NoteStakeholderShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "Stakeholder":
			StageSetFormDivSelectFieldToField(&inst.Stakeholder, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_NoteTaskShape_Stage(
	inst *models.NoteTaskShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Requirement_Stage(
	inst *models.Requirement,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_RequirementShape_Stage(
	inst *models.RequirementShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Requirement":
			StageSetFormDivSelectFieldToField(&inst.Requirement, probe.stageSet.Stage.GetInstancesSet[*models.Requirement](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Stakeholder_Stage(
	inst *models.Stakeholder,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IDAirbus":
			FormDivBasicFieldToField(&inst.IDAirbus, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		}
	}
}

func saveStageSet_StakeholderCompositionShape_Stage(
	inst *models.StakeholderCompositionShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Stakeholder":
			StageSetFormDivSelectFieldToField(&inst.Stakeholder, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_StakeholderConcernShape_Stage(
	inst *models.StakeholderConcernShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Stakeholder":
			StageSetFormDivSelectFieldToField(&inst.Stakeholder, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), formDiv)
		case "Concern":
			StageSetFormDivSelectFieldToField(&inst.Concern, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_StakeholderShape_Stage(
	inst *models.StakeholderShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Stakeholder":
			StageSetFormDivSelectFieldToField(&inst.Stakeholder, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_SupportLevel_Stage(
	inst *models.SupportLevel,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tool":
			StageSetFormDivSelectFieldToField(&inst.Tool, probe.stageSet.Stage.GetInstancesSet[*models.Tool](), formDiv)
		}
	}
}

func saveStageSet_Tool_Stage(
	inst *models.Tool,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}


func StageSetNewInstance_AnalysisNeed_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AnalysisNeed",
	}).Stage(probe.formStage)
	inst := new(models.AnalysisNeed)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_AnalysisNeed_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_AnalysisNeed_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Concept_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Concept",
	}).Stage(probe.formStage)
	inst := new(models.Concept)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Concept_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Concept_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ConceptShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ConceptShape",
	}).Stage(probe.formStage)
	inst := new(models.ConceptShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ConceptShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ConceptShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Concern_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Concern",
	}).Stage(probe.formStage)
	inst := new(models.Concern)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Concern_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Concern_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ConcernCompositionShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ConcernCompositionShape",
	}).Stage(probe.formStage)
	inst := new(models.ConcernCompositionShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ConcernCompositionShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ConcernCompositionShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ConcernInputShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ConcernInputShape",
	}).Stage(probe.formStage)
	inst := new(models.ConcernInputShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ConcernInputShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ConcernInputShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ConcernOutputShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ConcernOutputShape",
	}).Stage(probe.formStage)
	inst := new(models.ConcernOutputShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ConcernOutputShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ConcernOutputShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ConcernShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ConcernShape",
	}).Stage(probe.formStage)
	inst := new(models.ConcernShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ConcernShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ConcernShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ControlPointShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ControlPointShape",
	}).Stage(probe.formStage)
	inst := new(models.ControlPointShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ControlPointShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ControlPointShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Deliverable_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Deliverable",
	}).Stage(probe.formStage)
	inst := new(models.Deliverable)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Deliverable_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Deliverable_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DeliverableCompositionShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DeliverableCompositionShape",
	}).Stage(probe.formStage)
	inst := new(models.DeliverableCompositionShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DeliverableCompositionShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DeliverableCompositionShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DeliverableConceptShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DeliverableConceptShape",
	}).Stage(probe.formStage)
	inst := new(models.DeliverableConceptShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DeliverableConceptShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DeliverableConceptShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DeliverableShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DeliverableShape",
	}).Stage(probe.formStage)
	inst := new(models.DeliverableShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DeliverableShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DeliverableShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Diagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Diagram",
	}).Stage(probe.formStage)
	inst := new(models.Diagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Diagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Diagram_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DiagramShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DiagramShape",
	}).Stage(probe.formStage)
	inst := new(models.DiagramShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DiagramShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DiagramShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Library_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Library",
	}).Stage(probe.formStage)
	inst := new(models.Library)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Library_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Library_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Note_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Note",
	}).Stage(probe.formStage)
	inst := new(models.Note)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Note_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Note_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteDeliverableShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteDeliverableShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteDeliverableShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteDeliverableShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteDeliverableShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteStakeholderShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteStakeholderShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteStakeholderShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteStakeholderShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteStakeholderShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteTaskShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteTaskShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteTaskShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteTaskShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteTaskShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Requirement_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Requirement",
	}).Stage(probe.formStage)
	inst := new(models.Requirement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Requirement_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Requirement_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RequirementShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RequirementShape",
	}).Stage(probe.formStage)
	inst := new(models.RequirementShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RequirementShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RequirementShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Stakeholder_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Stakeholder",
	}).Stage(probe.formStage)
	inst := new(models.Stakeholder)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Stakeholder_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Stakeholder_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StakeholderCompositionShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StakeholderCompositionShape",
	}).Stage(probe.formStage)
	inst := new(models.StakeholderCompositionShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StakeholderCompositionShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StakeholderCompositionShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StakeholderConcernShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StakeholderConcernShape",
	}).Stage(probe.formStage)
	inst := new(models.StakeholderConcernShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StakeholderConcernShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StakeholderConcernShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StakeholderShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StakeholderShape",
	}).Stage(probe.formStage)
	inst := new(models.StakeholderShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StakeholderShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StakeholderShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SupportLevel_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SupportLevel",
	}).Stage(probe.formStage)
	inst := new(models.SupportLevel)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SupportLevel_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SupportLevel_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tool_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tool",
	}).Stage(probe.formStage)
	inst := new(models.Tool)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Tool_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Tool_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}


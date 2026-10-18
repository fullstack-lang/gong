// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
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
	case *models.ALTERNATIVE_ID:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ALTERNATIVE_ID_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ALTERNATIVE_ID_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_DATE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_DATE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_DATE_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_DATE_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_DATE_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_REAL_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_REAL_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_REAL_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_REAL_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_REAL_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_STRING_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_STRING_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_STRING_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_STRING_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_STRING_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_XHTML_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_DEFINITION_XHTML_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_BOOLEAN_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_DATE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_DATE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_DATE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_ENUMERATION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_INTEGER_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_REAL:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_REAL_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_REAL_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_STRING:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_STRING_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_STRING_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_XHTML:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ATTRIBUTE_VALUE_XHTML_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ATTRIBUTE_VALUE_XHTML_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ALTERNATIVE_ID:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ALTERNATIVE_ID_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ALTERNATIVE_ID_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_DATE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_DATE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_INTEGER_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_REAL_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_REAL_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_STRING_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_STRING_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_XHTML_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ATTRIBUTE_VALUE_XHTML_1_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_CHILDREN:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_CHILDREN_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_CHILDREN_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_CORE_CONTENT:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_CORE_CONTENT_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_CORE_CONTENT_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPES:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPES_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPES_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_DATE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_REAL_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_STRING_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_DATATYPE_DEFINITION_XHTML_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_EDITABLE_ATTS:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_EDITABLE_ATTS_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_EDITABLE_ATTS_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_ENUM_VALUE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_ENUM_VALUE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_ENUM_VALUE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_OBJECT:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_OBJECT_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_OBJECT_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_PROPERTIES:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_PROPERTIES_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_PROPERTIES_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_RELATION_GROUP_TYPE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_RELATION_GROUP_TYPE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_RELATION_GROUP_TYPE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SOURCE_1:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SOURCE_1_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SOURCE_1_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SOURCE_SPECIFICATION_1:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SOURCE_SPECIFICATION_1_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SOURCE_SPECIFICATION_1_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPECIFICATIONS:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPECIFICATIONS_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPECIFICATIONS_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPECIFICATION_TYPE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPECIFICATION_TYPE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPECIFICATION_TYPE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPECIFIED_VALUES:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPECIFIED_VALUES_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPECIFIED_VALUES_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_ATTRIBUTES:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_ATTRIBUTES_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_ATTRIBUTES_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_OBJECTS:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_OBJECTS_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_OBJECTS_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_OBJECT_TYPE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_OBJECT_TYPE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_RELATIONS:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_RELATIONS_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_RELATIONS_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_RELATION_GROUPS:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_RELATION_GROUPS_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_RELATION_GROUPS_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_RELATION_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_RELATION_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_RELATION_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_RELATION_TYPE_REF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_RELATION_TYPE_REF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_RELATION_TYPE_REF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_SPEC_TYPES:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_SPEC_TYPES_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_SPEC_TYPES_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_THE_HEADER:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_THE_HEADER_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_THE_HEADER_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.A_TOOL_EXTENSIONS:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_TOOL_EXTENSIONS_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_TOOL_EXTENSIONS_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_BOOLEAN_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_BOOLEAN_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_DATE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_DATE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_DATE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_ENUMERATION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_ENUMERATION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_INTEGER:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_INTEGER_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_INTEGER_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_REAL:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_REAL_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_REAL_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_STRING:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_STRING_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_STRING_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DATATYPE_DEFINITION_XHTML:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DATATYPE_DEFINITION_XHTML_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DATATYPE_DEFINITION_XHTML_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.EMBEDDED_VALUE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EMBEDDED_VALUE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_EMBEDDED_VALUE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ENUM_VALUE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ENUM_VALUE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ENUM_VALUE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.EmbeddedJpgImage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmbeddedJpgImage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_EmbeddedJpgImage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.EmbeddedPngImage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmbeddedPngImage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_EmbeddedPngImage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.EmbeddedSvgImage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmbeddedSvgImage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_EmbeddedSvgImage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Kill:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Kill_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Kill_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Map_identifier_bool:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Map_identifier_bool_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Map_identifier_bool_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RELATION_GROUP:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RELATION_GROUP_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RELATION_GROUP_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.RELATION_GROUP_TYPE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RELATION_GROUP_TYPE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_RELATION_GROUP_TYPE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.REQ_IF:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_REQ_IF_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_REQ_IF_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.REQ_IF_CONTENT:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_REQ_IF_CONTENT_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_REQ_IF_CONTENT_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.REQ_IF_HEADER:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_REQ_IF_HEADER_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_REQ_IF_HEADER_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.REQ_IF_TOOL_EXTENSION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_REQ_IF_TOOL_EXTENSION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_REQ_IF_TOOL_EXTENSION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPECIFICATION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPECIFICATION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPECIFICATION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPECIFICATION_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPECIFICATION_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPECIFICATION_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPECIFICATION_TYPE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPECIFICATION_TYPE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPECIFICATION_TYPE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPEC_HIERARCHY:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPEC_HIERARCHY_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPEC_HIERARCHY_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPEC_OBJECT:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPEC_OBJECT_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPEC_OBJECT_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPEC_OBJECT_TYPE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPEC_OBJECT_TYPE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPEC_OBJECT_TYPE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPEC_OBJECT_TYPE_Rendering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPEC_OBJECT_TYPE_Rendering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPEC_OBJECT_TYPE_Rendering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPEC_RELATION:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPEC_RELATION_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPEC_RELATION_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SPEC_RELATION_TYPE:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SPEC_RELATION_TYPE_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SPEC_RELATION_TYPE_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StaticWebSite:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaticWebSite_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StaticWebSite_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StaticWebSiteChapter:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaticWebSiteChapter_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StaticWebSiteChapter_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StaticWebSiteGeneratedImage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaticWebSiteGeneratedImage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StaticWebSiteGeneratedImage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StaticWebSiteImage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaticWebSiteImage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StaticWebSiteImage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StaticWebSiteParagraph:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaticWebSiteParagraph_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StaticWebSiteParagraph_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.XHTML_CONTENT:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_XHTML_CONTENT_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_XHTML_CONTENT_Stage(probe)
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

func saveStageSet_ALTERNATIVE_ID_Stage(
	inst *models.ALTERNATIVE_ID,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(
	inst *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_BOOLEAN](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_DATE_Stage(
	inst *models.ATTRIBUTE_DEFINITION_DATE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_DATE](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_DATE_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_DATE_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_DATE_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(
	inst *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MULTI_VALUED":
			FormDivBasicFieldToField(&inst.MULTI_VALUED, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_ENUMERATION](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Stage(
	inst *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_INTEGER](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_INTEGER_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_INTEGER_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_REAL_Stage(
	inst *models.ATTRIBUTE_DEFINITION_REAL,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_REAL](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_REAL_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_REAL_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_REAL_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_STRING_Stage(
	inst *models.ATTRIBUTE_DEFINITION_STRING,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_STRING](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_STRING_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_STRING_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_STRING_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Stage(
	inst *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_XHTML_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Rendering_Stage(
	inst *models.ATTRIBUTE_DEFINITION_XHTML_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ShowInTable":
			FormDivBasicFieldToField(&inst.ShowInTable, formDiv)
		case "ShowInTitle":
			FormDivBasicFieldToField(&inst.ShowInTitle, formDiv)
		case "ShowInSubject":
			FormDivBasicFieldToField(&inst.ShowInSubject, formDiv)
		case "Rank":
			FormDivBasicFieldToField(&inst.Rank, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_BOOLEAN_Stage(
	inst *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_DATE_Stage(
	inst *models.ATTRIBUTE_VALUE_DATE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_ENUMERATION_Stage(
	inst *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageSet.Stage.GetInstancesSet[*models.A_ENUM_VALUE_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_INTEGER_Stage(
	inst *models.ATTRIBUTE_VALUE_INTEGER,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_REAL_Stage(
	inst *models.ATTRIBUTE_VALUE_REAL,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_STRING_Stage(
	inst *models.ATTRIBUTE_VALUE_STRING,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_XHTML_Stage(
	inst *models.ATTRIBUTE_VALUE_XHTML,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](), formDiv)
		case "IS_SIMPLIFIED":
			FormDivBasicFieldToField(&inst.IS_SIMPLIFIED, formDiv)
		case "THE_VALUE":
			StageSetFormDivSelectFieldToField(&inst.THE_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.XHTML_CONTENT](), formDiv)
		case "THE_ORIGINAL_VALUE":
			StageSetFormDivSelectFieldToField(&inst.THE_ORIGINAL_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.XHTML_CONTENT](), formDiv)
		}
	}
}

func saveStageSet_A_ALTERNATIVE_ID_Stage(
	inst *models.A_ALTERNATIVE_ID,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_DATE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_DATE_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_INTEGER_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_REAL_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_REAL_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_STRING_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_STRING_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_XHTML_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_XHTML_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(
	inst *models.A_ATTRIBUTE_VALUE_BOOLEAN,
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

func saveStageSet_A_ATTRIBUTE_VALUE_DATE_Stage(
	inst *models.A_ATTRIBUTE_VALUE_DATE,
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

func saveStageSet_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(
	inst *models.A_ATTRIBUTE_VALUE_ENUMERATION,
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

func saveStageSet_A_ATTRIBUTE_VALUE_INTEGER_Stage(
	inst *models.A_ATTRIBUTE_VALUE_INTEGER,
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

func saveStageSet_A_ATTRIBUTE_VALUE_REAL_Stage(
	inst *models.A_ATTRIBUTE_VALUE_REAL,
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

func saveStageSet_A_ATTRIBUTE_VALUE_STRING_Stage(
	inst *models.A_ATTRIBUTE_VALUE_STRING,
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

func saveStageSet_A_ATTRIBUTE_VALUE_XHTML_Stage(
	inst *models.A_ATTRIBUTE_VALUE_XHTML,
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

func saveStageSet_A_ATTRIBUTE_VALUE_XHTML_1_Stage(
	inst *models.A_ATTRIBUTE_VALUE_XHTML_1,
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

func saveStageSet_A_CHILDREN_Stage(
	inst *models.A_CHILDREN,
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

func saveStageSet_A_CORE_CONTENT_Stage(
	inst *models.A_CORE_CONTENT,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "REQ_IF_CONTENT":
			StageSetFormDivSelectFieldToField(&inst.REQ_IF_CONTENT, probe.stageSet.Stage.GetInstancesSet[*models.REQ_IF_CONTENT](), formDiv)
		}
	}
}

func saveStageSet_A_DATATYPES_Stage(
	inst *models.A_DATATYPES,
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

func saveStageSet_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_BOOLEAN_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_BOOLEAN_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_DATE_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_DATE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_DATE_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_ENUMERATION_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_ENUMERATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_INTEGER_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_INTEGER_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_REAL_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_REAL_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_REAL_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_STRING_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_STRING_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_STRING_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_XHTML_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_XHTML_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_XHTML_REF, formDiv)
		}
	}
}

func saveStageSet_A_EDITABLE_ATTS_Stage(
	inst *models.A_EDITABLE_ATTS,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_DATE_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_INTEGER_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_REAL_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_STRING_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_XHTML_REF, formDiv)
		}
	}
}

func saveStageSet_A_ENUM_VALUE_REF_Stage(
	inst *models.A_ENUM_VALUE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ENUM_VALUE_REF":
			FormDivBasicFieldToField(&inst.ENUM_VALUE_REF, formDiv)
		}
	}
}

func saveStageSet_A_OBJECT_Stage(
	inst *models.A_OBJECT,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&inst.SPEC_OBJECT_REF, formDiv)
		}
	}
}

func saveStageSet_A_PROPERTIES_Stage(
	inst *models.A_PROPERTIES,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EMBEDDED_VALUE":
			StageSetFormDivSelectFieldToField(&inst.EMBEDDED_VALUE, probe.stageSet.Stage.GetInstancesSet[*models.EMBEDDED_VALUE](), formDiv)
		}
	}
}

func saveStageSet_A_RELATION_GROUP_TYPE_REF_Stage(
	inst *models.A_RELATION_GROUP_TYPE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RELATION_GROUP_TYPE_REF":
			FormDivBasicFieldToField(&inst.RELATION_GROUP_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SOURCE_1_Stage(
	inst *models.A_SOURCE_1,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&inst.SPEC_OBJECT_REF, formDiv)
		}
	}
}

func saveStageSet_A_SOURCE_SPECIFICATION_1_Stage(
	inst *models.A_SOURCE_SPECIFICATION_1,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPECIFICATION_REF":
			StageSetFormDivEnumStringFieldToField(&inst.SPECIFICATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPECIFICATIONS_Stage(
	inst *models.A_SPECIFICATIONS,
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

func saveStageSet_A_SPECIFICATION_TYPE_REF_Stage(
	inst *models.A_SPECIFICATION_TYPE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPECIFICATION_TYPE_REF":
			FormDivBasicFieldToField(&inst.SPECIFICATION_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPECIFIED_VALUES_Stage(
	inst *models.A_SPECIFIED_VALUES,
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

func saveStageSet_A_SPEC_ATTRIBUTES_Stage(
	inst *models.A_SPEC_ATTRIBUTES,
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

func saveStageSet_A_SPEC_OBJECTS_Stage(
	inst *models.A_SPEC_OBJECTS,
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

func saveStageSet_A_SPEC_OBJECT_TYPE_REF_Stage(
	inst *models.A_SPEC_OBJECT_TYPE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT_TYPE_REF":
			FormDivBasicFieldToField(&inst.SPEC_OBJECT_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPEC_RELATIONS_Stage(
	inst *models.A_SPEC_RELATIONS,
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

func saveStageSet_A_SPEC_RELATION_GROUPS_Stage(
	inst *models.A_SPEC_RELATION_GROUPS,
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

func saveStageSet_A_SPEC_RELATION_REF_Stage(
	inst *models.A_SPEC_RELATION_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_RELATION_REF":
			FormDivBasicFieldToField(&inst.SPEC_RELATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPEC_RELATION_TYPE_REF_Stage(
	inst *models.A_SPEC_RELATION_TYPE_REF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_RELATION_TYPE_REF":
			FormDivBasicFieldToField(&inst.SPEC_RELATION_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPEC_TYPES_Stage(
	inst *models.A_SPEC_TYPES,
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

func saveStageSet_A_THE_HEADER_Stage(
	inst *models.A_THE_HEADER,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "REQ_IF_HEADER":
			StageSetFormDivSelectFieldToField(&inst.REQ_IF_HEADER, probe.stageSet.Stage.GetInstancesSet[*models.REQ_IF_HEADER](), formDiv)
		}
	}
}

func saveStageSet_A_TOOL_EXTENSIONS_Stage(
	inst *models.A_TOOL_EXTENSIONS,
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

func saveStageSet_DATATYPE_DEFINITION_BOOLEAN_Stage(
	inst *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_DATE_Stage(
	inst *models.DATATYPE_DEFINITION_DATE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_ENUMERATION_Stage(
	inst *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPECIFIED_VALUES":
			StageSetFormDivSelectFieldToField(&inst.SPECIFIED_VALUES, probe.stageSet.Stage.GetInstancesSet[*models.A_SPECIFIED_VALUES](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_INTEGER_Stage(
	inst *models.DATATYPE_DEFINITION_INTEGER,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MAX":
			FormDivBasicFieldToField(&inst.MAX, formDiv)
		case "MIN":
			FormDivBasicFieldToField(&inst.MIN, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_REAL_Stage(
	inst *models.DATATYPE_DEFINITION_REAL,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ACCURACY":
			FormDivBasicFieldToField(&inst.ACCURACY, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MAX":
			FormDivBasicFieldToField(&inst.MAX, formDiv)
		case "MIN":
			FormDivBasicFieldToField(&inst.MIN, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_STRING_Stage(
	inst *models.DATATYPE_DEFINITION_STRING,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MAX_LENGTH":
			FormDivBasicFieldToField(&inst.MAX_LENGTH, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_XHTML_Stage(
	inst *models.DATATYPE_DEFINITION_XHTML,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_EMBEDDED_VALUE_Stage(
	inst *models.EMBEDDED_VALUE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "KEY":
			FormDivBasicFieldToField(&inst.KEY, formDiv)
		case "OTHER_CONTENT":
			FormDivBasicFieldToField(&inst.OTHER_CONTENT, formDiv)
		}
	}
}

func saveStageSet_ENUM_VALUE_Stage(
	inst *models.ENUM_VALUE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "PROPERTIES":
			StageSetFormDivSelectFieldToField(&inst.PROPERTIES, probe.stageSet.Stage.GetInstancesSet[*models.A_PROPERTIES](), formDiv)
		}
	}
}

func saveStageSet_EmbeddedJpgImage_Stage(
	inst *models.EmbeddedJpgImage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&inst.Base64Content, formDiv)
		}
	}
}

func saveStageSet_EmbeddedPngImage_Stage(
	inst *models.EmbeddedPngImage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&inst.Base64Content, formDiv)
		}
	}
}

func saveStageSet_EmbeddedSvgImage_Stage(
	inst *models.EmbeddedSvgImage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Content":
			FormDivBasicFieldToField(&inst.Content, formDiv)
		}
	}
}

func saveStageSet_Kill_Stage(
	inst *models.Kill,
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

func saveStageSet_Map_identifier_bool_Stage(
	inst *models.Map_identifier_bool,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_RELATION_GROUP_Stage(
	inst *models.RELATION_GROUP,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SOURCE_SPECIFICATION":
			StageSetFormDivSelectFieldToField(&inst.SOURCE_SPECIFICATION, probe.stageSet.Stage.GetInstancesSet[*models.A_SOURCE_SPECIFICATION_1](), formDiv)
		case "SPEC_RELATIONS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_RELATIONS, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_RELATION_REF](), formDiv)
		case "TARGET_SPECIFICATION":
			StageSetFormDivSelectFieldToField(&inst.TARGET_SPECIFICATION, probe.stageSet.Stage.GetInstancesSet[*models.A_SOURCE_SPECIFICATION_1](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_RELATION_GROUP_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_RELATION_GROUP_TYPE_Stage(
	inst *models.RELATION_GROUP_TYPE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_REQ_IF_Stage(
	inst *models.REQ_IF,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "THE_HEADER":
			StageSetFormDivSelectFieldToField(&inst.THE_HEADER, probe.stageSet.Stage.GetInstancesSet[*models.A_THE_HEADER](), formDiv)
		case "CORE_CONTENT":
			StageSetFormDivSelectFieldToField(&inst.CORE_CONTENT, probe.stageSet.Stage.GetInstancesSet[*models.A_CORE_CONTENT](), formDiv)
		case "TOOL_EXTENSIONS":
			StageSetFormDivSelectFieldToField(&inst.TOOL_EXTENSIONS, probe.stageSet.Stage.GetInstancesSet[*models.A_TOOL_EXTENSIONS](), formDiv)
		}
	}
}

func saveStageSet_REQ_IF_CONTENT_Stage(
	inst *models.REQ_IF_CONTENT,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPES":
			StageSetFormDivSelectFieldToField(&inst.DATATYPES, probe.stageSet.Stage.GetInstancesSet[*models.A_DATATYPES](), formDiv)
		case "SPEC_TYPES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_TYPES, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_TYPES](), formDiv)
		case "SPEC_OBJECTS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_OBJECTS, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_OBJECTS](), formDiv)
		case "SPEC_RELATIONS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_RELATIONS, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_RELATIONS](), formDiv)
		case "SPECIFICATIONS":
			StageSetFormDivSelectFieldToField(&inst.SPECIFICATIONS, probe.stageSet.Stage.GetInstancesSet[*models.A_SPECIFICATIONS](), formDiv)
		case "SPEC_RELATION_GROUPS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_RELATION_GROUPS, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_RELATION_GROUPS](), formDiv)
		}
	}
}

func saveStageSet_REQ_IF_HEADER_Stage(
	inst *models.REQ_IF_HEADER,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&inst.COMMENT, formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&inst.CREATION_TIME, formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&inst.REPOSITORY_ID, formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&inst.REQ_IF_TOOL_ID, formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&inst.REQ_IF_VERSION, formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&inst.SOURCE_TOOL_ID, formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&inst.TITLE, formDiv)
		}
	}
}

func saveStageSet_REQ_IF_TOOL_EXTENSION_Stage(
	inst *models.REQ_IF_TOOL_EXTENSION,
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

func saveStageSet_SPECIFICATION_Stage(
	inst *models.SPECIFICATION,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_SPECIFICATION_TYPE_REF](), formDiv)
		case "CHILDREN":
			StageSetFormDivSelectFieldToField(&inst.CHILDREN, probe.stageSet.Stage.GetInstancesSet[*models.A_CHILDREN](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), formDiv)
		}
	}
}

func saveStageSet_SPECIFICATION_Rendering_Stage(
	inst *models.SPECIFICATION_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNodeExpanded, formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&inst.IsSelected, formDiv)
		case "IsWithHeadingNumbering":
			FormDivBasicFieldToField(&inst.IsWithHeadingNumbering, formDiv)
		}
	}
}

func saveStageSet_SPECIFICATION_TYPE_Stage(
	inst *models.SPECIFICATION_TYPE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_SPEC_HIERARCHY_Stage(
	inst *models.SPEC_HIERARCHY,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&inst.IS_TABLE_INTERNAL, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "OBJECT":
			StageSetFormDivSelectFieldToField(&inst.OBJECT, probe.stageSet.Stage.GetInstancesSet[*models.A_OBJECT](), formDiv)
		case "CHILDREN":
			StageSetFormDivSelectFieldToField(&inst.CHILDREN, probe.stageSet.Stage.GetInstancesSet[*models.A_CHILDREN](), formDiv)
		case "EDITABLE_ATTS":
			StageSetFormDivSelectFieldToField(&inst.EDITABLE_ATTS, probe.stageSet.Stage.GetInstancesSet[*models.A_EDITABLE_ATTS](), formDiv)
		}
	}
}

func saveStageSet_SPEC_OBJECT_Stage(
	inst *models.SPEC_OBJECT,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_OBJECT_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_SPEC_OBJECT_TYPE_Stage(
	inst *models.SPEC_OBJECT_TYPE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_SPEC_OBJECT_TYPE_Rendering_Stage(
	inst *models.SPEC_OBJECT_TYPE_Rendering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNodeExpanded, formDiv)
		case "ShowIdentifier":
			FormDivBasicFieldToField(&inst.ShowIdentifier, formDiv)
		case "ShowName":
			FormDivBasicFieldToField(&inst.ShowName, formDiv)
		case "ShowRelations":
			FormDivBasicFieldToField(&inst.ShowRelations, formDiv)
		case "IsHeading":
			FormDivBasicFieldToField(&inst.IsHeading, formDiv)
		}
	}
}

func saveStageSet_SPEC_RELATION_Stage(
	inst *models.SPEC_RELATION,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageSet.Stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), formDiv)
		case "SOURCE":
			StageSetFormDivSelectFieldToField(&inst.SOURCE, probe.stageSet.Stage.GetInstancesSet[*models.A_SOURCE_1](), formDiv)
		case "TARGET":
			StageSetFormDivSelectFieldToField(&inst.TARGET, probe.stageSet.Stage.GetInstancesSet[*models.A_SOURCE_1](), formDiv)
		case "TYPE":
			StageSetFormDivSelectFieldToField(&inst.TYPE, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_RELATION_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_SPEC_RELATION_TYPE_Stage(
	inst *models.SPEC_RELATION_TYPE,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageSet.Stage.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageSet.Stage.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_StaticWebSite_Stage(
	inst *models.StaticWebSite,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "MarkdownContent":
			FormDivBasicFieldToField(&inst.MarkdownContent, formDiv)
		case "InputImagesDir":
			FormDivBasicFieldToField(&inst.InputImagesDir, formDiv)
		case "OutputStaticWebDir":
			FormDivBasicFieldToField(&inst.OutputStaticWebDir, formDiv)
		case "VersionInfo":
			FormDivBasicFieldToField(&inst.VersionInfo, formDiv)
		}
	}
}

func saveStageSet_StaticWebSiteChapter_Stage(
	inst *models.StaticWebSiteChapter,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "MarkdownContent":
			FormDivBasicFieldToField(&inst.MarkdownContent, formDiv)
		}
	}
}

func saveStageSet_StaticWebSiteGeneratedImage_Stage(
	inst *models.StaticWebSiteGeneratedImage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SourceDirectoryPath":
			FormDivBasicFieldToField(&inst.SourceDirectoryPath, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		}
	}
}

func saveStageSet_StaticWebSiteImage_Stage(
	inst *models.StaticWebSiteImage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SourceDirectoryPath":
			FormDivBasicFieldToField(&inst.SourceDirectoryPath, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		}
	}
}

func saveStageSet_StaticWebSiteParagraph_Stage(
	inst *models.StaticWebSiteParagraph,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "LegendMarkdownContent":
			FormDivBasicFieldToField(&inst.LegendMarkdownContent, formDiv)
		case "Image":
			StageSetFormDivSelectFieldToField(&inst.Image, probe.stageSet.Stage.GetInstancesSet[*models.StaticWebSiteImage](), formDiv)
		}
	}
}

func saveStageSet_XHTML_CONTENT_Stage(
	inst *models.XHTML_CONTENT,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		case "PureText":
			FormDivBasicFieldToField(&inst.PureText, formDiv)
		}
	}
}


func StageSetNewInstance_ALTERNATIVE_ID_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ALTERNATIVE_ID",
	}).Stage(probe.formStage)
	inst := new(models.ALTERNATIVE_ID)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ALTERNATIVE_ID_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ALTERNATIVE_ID_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_BOOLEAN_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_DATE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_DATE",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_DATE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_DATE_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_DATE_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_DATE_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_DATE_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_DATE_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_DATE_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_ENUMERATION_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_INTEGER_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_INTEGER_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_INTEGER_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_REAL_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_REAL",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_REAL_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_REAL_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_REAL_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_REAL_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_REAL_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_REAL_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_REAL_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_STRING_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_STRING",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_STRING_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_STRING_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_STRING_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_STRING_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_STRING_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_STRING_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_STRING_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_XHTML_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_XHTML_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_DEFINITION_XHTML_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_DEFINITION_XHTML_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_XHTML_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_XHTML_Rendering_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_DATE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_DATE",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_DATE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_DATE_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_INTEGER_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_INTEGER_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_REAL_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_REAL",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_REAL_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_REAL_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_STRING_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_STRING",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_STRING_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_STRING_Stage(probe)
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

func StageSetNewInstance_ATTRIBUTE_VALUE_XHTML_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ATTRIBUTE_VALUE_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ATTRIBUTE_VALUE_XHTML_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_XHTML_Stage(probe)
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

func StageSetNewInstance_A_ALTERNATIVE_ID_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ALTERNATIVE_ID",
	}).Stage(probe.formStage)
	inst := new(models.A_ALTERNATIVE_ID)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ALTERNATIVE_ID_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ALTERNATIVE_ID_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_BOOLEAN_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_DATE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_DATE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_ENUMERATION_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_INTEGER_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_REAL_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_REAL_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_STRING_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_STRING_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_DEFINITION_XHTML_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_XHTML_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_DATE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_DATE",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_DATE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_DATE_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_INTEGER_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_REAL_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_REAL",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_REAL_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_REAL_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_STRING_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_STRING",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_STRING_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_STRING_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_XHTML_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_XHTML_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_Stage(probe)
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

func StageSetNewInstance_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ATTRIBUTE_VALUE_XHTML_1",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_XHTML_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ATTRIBUTE_VALUE_XHTML_1_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe)
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

func StageSetNewInstance_A_CHILDREN_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_CHILDREN",
	}).Stage(probe.formStage)
	inst := new(models.A_CHILDREN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_CHILDREN_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_CHILDREN_Stage(probe)
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

func StageSetNewInstance_A_CORE_CONTENT_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_CORE_CONTENT",
	}).Stage(probe.formStage)
	inst := new(models.A_CORE_CONTENT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_CORE_CONTENT_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_CORE_CONTENT_Stage(probe)
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

func StageSetNewInstance_A_DATATYPES_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPES",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPES_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPES_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_BOOLEAN_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_DATE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_DATE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_DATE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_ENUMERATION_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_ENUMERATION_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_INTEGER_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_INTEGER_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_REAL_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_REAL_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_REAL_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_STRING_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_STRING_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_STRING_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe)
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

func StageSetNewInstance_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_DATATYPE_DEFINITION_XHTML_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_XHTML_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_DATATYPE_DEFINITION_XHTML_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe)
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

func StageSetNewInstance_A_EDITABLE_ATTS_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_EDITABLE_ATTS",
	}).Stage(probe.formStage)
	inst := new(models.A_EDITABLE_ATTS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_EDITABLE_ATTS_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_EDITABLE_ATTS_Stage(probe)
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

func StageSetNewInstance_A_ENUM_VALUE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_ENUM_VALUE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ENUM_VALUE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_ENUM_VALUE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_ENUM_VALUE_REF_Stage(probe)
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

func StageSetNewInstance_A_OBJECT_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_OBJECT",
	}).Stage(probe.formStage)
	inst := new(models.A_OBJECT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_OBJECT_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_OBJECT_Stage(probe)
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

func StageSetNewInstance_A_PROPERTIES_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_PROPERTIES",
	}).Stage(probe.formStage)
	inst := new(models.A_PROPERTIES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_PROPERTIES_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_PROPERTIES_Stage(probe)
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

func StageSetNewInstance_A_RELATION_GROUP_TYPE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_RELATION_GROUP_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_RELATION_GROUP_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_RELATION_GROUP_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_RELATION_GROUP_TYPE_REF_Stage(probe)
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

func StageSetNewInstance_A_SOURCE_1_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SOURCE_1",
	}).Stage(probe.formStage)
	inst := new(models.A_SOURCE_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SOURCE_1_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SOURCE_1_Stage(probe)
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

func StageSetNewInstance_A_SOURCE_SPECIFICATION_1_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SOURCE_SPECIFICATION_1",
	}).Stage(probe.formStage)
	inst := new(models.A_SOURCE_SPECIFICATION_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SOURCE_SPECIFICATION_1_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SOURCE_SPECIFICATION_1_Stage(probe)
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

func StageSetNewInstance_A_SPECIFICATIONS_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPECIFICATIONS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPECIFICATIONS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPECIFICATIONS_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPECIFICATIONS_Stage(probe)
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

func StageSetNewInstance_A_SPECIFICATION_TYPE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPECIFICATION_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPECIFICATION_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPECIFICATION_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPECIFICATION_TYPE_REF_Stage(probe)
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

func StageSetNewInstance_A_SPECIFIED_VALUES_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPECIFIED_VALUES",
	}).Stage(probe.formStage)
	inst := new(models.A_SPECIFIED_VALUES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPECIFIED_VALUES_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPECIFIED_VALUES_Stage(probe)
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

func StageSetNewInstance_A_SPEC_ATTRIBUTES_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_ATTRIBUTES",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_ATTRIBUTES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_ATTRIBUTES_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_ATTRIBUTES_Stage(probe)
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

func StageSetNewInstance_A_SPEC_OBJECTS_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_OBJECTS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_OBJECTS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_OBJECTS_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_OBJECTS_Stage(probe)
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

func StageSetNewInstance_A_SPEC_OBJECT_TYPE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_OBJECT_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_OBJECT_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_OBJECT_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_OBJECT_TYPE_REF_Stage(probe)
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

func StageSetNewInstance_A_SPEC_RELATIONS_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_RELATIONS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATIONS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_RELATIONS_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_RELATIONS_Stage(probe)
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

func StageSetNewInstance_A_SPEC_RELATION_GROUPS_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_RELATION_GROUPS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATION_GROUPS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_RELATION_GROUPS_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_RELATION_GROUPS_Stage(probe)
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

func StageSetNewInstance_A_SPEC_RELATION_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_RELATION_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATION_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_RELATION_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_RELATION_REF_Stage(probe)
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

func StageSetNewInstance_A_SPEC_RELATION_TYPE_REF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_RELATION_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATION_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_RELATION_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_RELATION_TYPE_REF_Stage(probe)
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

func StageSetNewInstance_A_SPEC_TYPES_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_SPEC_TYPES",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_TYPES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_SPEC_TYPES_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_SPEC_TYPES_Stage(probe)
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

func StageSetNewInstance_A_THE_HEADER_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_THE_HEADER",
	}).Stage(probe.formStage)
	inst := new(models.A_THE_HEADER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_THE_HEADER_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_THE_HEADER_Stage(probe)
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

func StageSetNewInstance_A_TOOL_EXTENSIONS_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A_TOOL_EXTENSIONS",
	}).Stage(probe.formStage)
	inst := new(models.A_TOOL_EXTENSIONS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_TOOL_EXTENSIONS_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_TOOL_EXTENSIONS_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_BOOLEAN_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_BOOLEAN_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_DATE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_DATE",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_DATE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_DATE_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_ENUMERATION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_ENUMERATION_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_INTEGER_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_INTEGER_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_INTEGER_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_REAL_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_REAL",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_REAL_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_REAL_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_STRING_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_STRING",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_STRING_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_STRING_Stage(probe)
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

func StageSetNewInstance_DATATYPE_DEFINITION_XHTML_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DATATYPE_DEFINITION_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DATATYPE_DEFINITION_XHTML_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_XHTML_Stage(probe)
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

func StageSetNewInstance_EMBEDDED_VALUE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EMBEDDED_VALUE",
	}).Stage(probe.formStage)
	inst := new(models.EMBEDDED_VALUE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_EMBEDDED_VALUE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_EMBEDDED_VALUE_Stage(probe)
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

func StageSetNewInstance_ENUM_VALUE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ENUM_VALUE",
	}).Stage(probe.formStage)
	inst := new(models.ENUM_VALUE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ENUM_VALUE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ENUM_VALUE_Stage(probe)
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

func StageSetNewInstance_EmbeddedJpgImage_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmbeddedJpgImage",
	}).Stage(probe.formStage)
	inst := new(models.EmbeddedJpgImage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_EmbeddedJpgImage_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_EmbeddedJpgImage_Stage(probe)
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

func StageSetNewInstance_EmbeddedPngImage_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmbeddedPngImage",
	}).Stage(probe.formStage)
	inst := new(models.EmbeddedPngImage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_EmbeddedPngImage_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_EmbeddedPngImage_Stage(probe)
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

func StageSetNewInstance_EmbeddedSvgImage_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmbeddedSvgImage",
	}).Stage(probe.formStage)
	inst := new(models.EmbeddedSvgImage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_EmbeddedSvgImage_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_EmbeddedSvgImage_Stage(probe)
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

func StageSetNewInstance_Kill_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Kill",
	}).Stage(probe.formStage)
	inst := new(models.Kill)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Kill_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Kill_Stage(probe)
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

func StageSetNewInstance_Map_identifier_bool_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Map_identifier_bool",
	}).Stage(probe.formStage)
	inst := new(models.Map_identifier_bool)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Map_identifier_bool_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Map_identifier_bool_Stage(probe)
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

func StageSetNewInstance_RELATION_GROUP_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RELATION_GROUP",
	}).Stage(probe.formStage)
	inst := new(models.RELATION_GROUP)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RELATION_GROUP_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RELATION_GROUP_Stage(probe)
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

func StageSetNewInstance_RELATION_GROUP_TYPE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RELATION_GROUP_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.RELATION_GROUP_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_RELATION_GROUP_TYPE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_RELATION_GROUP_TYPE_Stage(probe)
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

func StageSetNewInstance_REQ_IF_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New REQ_IF",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_REQ_IF_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_REQ_IF_Stage(probe)
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

func StageSetNewInstance_REQ_IF_CONTENT_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New REQ_IF_CONTENT",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF_CONTENT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_REQ_IF_CONTENT_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_REQ_IF_CONTENT_Stage(probe)
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

func StageSetNewInstance_REQ_IF_HEADER_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New REQ_IF_HEADER",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF_HEADER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_REQ_IF_HEADER_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_REQ_IF_HEADER_Stage(probe)
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

func StageSetNewInstance_REQ_IF_TOOL_EXTENSION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New REQ_IF_TOOL_EXTENSION",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF_TOOL_EXTENSION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_REQ_IF_TOOL_EXTENSION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_REQ_IF_TOOL_EXTENSION_Stage(probe)
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

func StageSetNewInstance_SPECIFICATION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPECIFICATION",
	}).Stage(probe.formStage)
	inst := new(models.SPECIFICATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPECIFICATION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPECIFICATION_Stage(probe)
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

func StageSetNewInstance_SPECIFICATION_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPECIFICATION_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.SPECIFICATION_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPECIFICATION_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPECIFICATION_Rendering_Stage(probe)
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

func StageSetNewInstance_SPECIFICATION_TYPE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPECIFICATION_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.SPECIFICATION_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPECIFICATION_TYPE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPECIFICATION_TYPE_Stage(probe)
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

func StageSetNewInstance_SPEC_HIERARCHY_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPEC_HIERARCHY",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_HIERARCHY)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPEC_HIERARCHY_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPEC_HIERARCHY_Stage(probe)
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

func StageSetNewInstance_SPEC_OBJECT_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPEC_OBJECT",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_OBJECT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPEC_OBJECT_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPEC_OBJECT_Stage(probe)
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

func StageSetNewInstance_SPEC_OBJECT_TYPE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPEC_OBJECT_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_OBJECT_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPEC_OBJECT_TYPE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPEC_OBJECT_TYPE_Stage(probe)
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

func StageSetNewInstance_SPEC_OBJECT_TYPE_Rendering_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPEC_OBJECT_TYPE_Rendering",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_OBJECT_TYPE_Rendering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPEC_OBJECT_TYPE_Rendering_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPEC_OBJECT_TYPE_Rendering_Stage(probe)
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

func StageSetNewInstance_SPEC_RELATION_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPEC_RELATION",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_RELATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPEC_RELATION_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPEC_RELATION_Stage(probe)
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

func StageSetNewInstance_SPEC_RELATION_TYPE_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SPEC_RELATION_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_RELATION_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SPEC_RELATION_TYPE_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SPEC_RELATION_TYPE_Stage(probe)
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

func StageSetNewInstance_StaticWebSite_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaticWebSite",
	}).Stage(probe.formStage)
	inst := new(models.StaticWebSite)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StaticWebSite_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StaticWebSite_Stage(probe)
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

func StageSetNewInstance_StaticWebSiteChapter_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaticWebSiteChapter",
	}).Stage(probe.formStage)
	inst := new(models.StaticWebSiteChapter)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StaticWebSiteChapter_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StaticWebSiteChapter_Stage(probe)
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

func StageSetNewInstance_StaticWebSiteGeneratedImage_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaticWebSiteGeneratedImage",
	}).Stage(probe.formStage)
	inst := new(models.StaticWebSiteGeneratedImage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StaticWebSiteGeneratedImage_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StaticWebSiteGeneratedImage_Stage(probe)
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

func StageSetNewInstance_StaticWebSiteImage_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaticWebSiteImage",
	}).Stage(probe.formStage)
	inst := new(models.StaticWebSiteImage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StaticWebSiteImage_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StaticWebSiteImage_Stage(probe)
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

func StageSetNewInstance_StaticWebSiteParagraph_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaticWebSiteParagraph",
	}).Stage(probe.formStage)
	inst := new(models.StaticWebSiteParagraph)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StaticWebSiteParagraph_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StaticWebSiteParagraph_Stage(probe)
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

func StageSetNewInstance_XHTML_CONTENT_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New XHTML_CONTENT",
	}).Stage(probe.formStage)
	inst := new(models.XHTML_CONTENT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_XHTML_CONTENT_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_XHTML_CONTENT_Stage(probe)
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


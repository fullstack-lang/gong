// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/cursor/go/models"
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
	case *models.Cursor:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Cursor_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Cursor_Stage(probe)
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

func saveStageSet_Cursor_Stage(
	inst *models.Cursor,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "StartX":
			FormDivBasicFieldToField(&inst.StartX, formDiv)
		case "EndX":
			FormDivBasicFieldToField(&inst.EndX, formDiv)
		case "Y1":
			FormDivBasicFieldToField(&inst.Y1, formDiv)
		case "Y2":
			FormDivBasicFieldToField(&inst.Y2, formDiv)
		case "DurationSeconds":
			FormDivBasicFieldToField(&inst.DurationSeconds, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&inst.FillOpacity, formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&inst.Stroke, formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&inst.StrokeOpacity, formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&inst.StrokeWidth, formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&inst.StrokeDashArray, formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&inst.StrokeDashArrayWhenSelected, formDiv)
		case "Transform":
			FormDivBasicFieldToField(&inst.Transform, formDiv)
		case "IsPlaying":
			FormDivBasicFieldToField(&inst.IsPlaying, formDiv)
		}
	}
}


func StageSetNewInstance_Cursor_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Cursor",
	}).Stage(probe.formStage)
	inst := new(models.Cursor)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Cursor_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Cursor_Stage(probe)
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


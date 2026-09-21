// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
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
	case *models.A:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_A_Stage(probe)
				probe.ux_tree()
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.B:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_B_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_B_Stage(probe)
				probe.ux_tree()
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *x.X:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XStage.Lock()
				defer probe.stageSet.XStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_X_XStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XStage)
				}
				probe.stageSet.XStage.Commit()
				updateStageSetTable_X_XStage(probe)
				probe.ux_tree()
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *y.Y:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.YStage.Lock()
				defer probe.stageSet.YStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Y_YStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.YStage)
				}
				probe.stageSet.YStage.Commit()
				updateStageSetTable_Y_YStage(probe)
				probe.ux_tree()
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *model.SubModel:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ModelStage.Lock()
				defer probe.stageSet.ModelStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SubModel_ModelStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ModelStage)
				}
				probe.stageSet.ModelStage.Commit()
				updateStageSetTable_SubModel_ModelStage(probe)
				probe.ux_tree()
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	default:
		_ = inst
	}

	probe.formStage.Commit()
}

func saveStageSet_A_Stage(
	inst *models.A,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumberField":
			FormDivBasicFieldToField(&inst.NumberField, formDiv)
		case "B":
			StageSetFormDivSelectFieldToField(&inst.B, probe.stageSet.Stage.GetInstancesSet[*models.B](), formDiv)
		case "X":
			StageSetFormDivSelectFieldToField(&inst.X, probe.stageSet.XStage.GetInstancesSet[*x.X](), formDiv)
		case "Foo":
			FormDivBasicFieldToField(&inst.Foo, formDiv)
		case "Bar":
			FormDivBasicFieldToField(&inst.Bar, formDiv)
		case "Zorgh":
			FormDivBasicFieldToField(&inst.Zorgh, formDiv)
		}
	}
}

func saveStageSet_B_Stage(
	inst *models.B,
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

func saveStageSet_X_XStage(
	inst *x.X,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Y":
			StageSetFormDivSelectFieldToField(&inst.Y, probe.stageSet.YStage.GetInstancesSet[*y.Y](), formDiv)
		}
	}
}

func saveStageSet_Y_YStage(
	inst *y.Y,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SubModel":
			StageSetFormDivSelectFieldToField(&inst.SubModel, probe.stageSet.ModelStage.GetInstancesSet[*model.SubModel](), formDiv)
		}
	}
}

func saveStageSet_SubModel_ModelStage(
	inst *model.SubModel,
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


func StageSetNewInstance_A_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New A",
	}).Stage(probe.formStage)
	inst := new(models.A)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_A_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_A_Stage(probe)
			probe.ux_tree()
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_B_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New B",
	}).Stage(probe.formStage)
	inst := new(models.B)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_B_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_B_Stage(probe)
			probe.ux_tree()
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_X_XStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New X",
	}).Stage(probe.formStage)
	inst := new(x.X)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XStage.Lock()
			defer probe.stageSet.XStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XStage)
			saveStageSet_X_XStage(inst, probe, formGroup)
			probe.stageSet.XStage.Commit()
			updateStageSetTable_X_XStage(probe)
			probe.ux_tree()
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Y_YStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Y",
	}).Stage(probe.formStage)
	inst := new(y.Y)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.YStage.Lock()
			defer probe.stageSet.YStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.YStage)
			saveStageSet_Y_YStage(inst, probe, formGroup)
			probe.stageSet.YStage.Commit()
			updateStageSetTable_Y_YStage(probe)
			probe.ux_tree()
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SubModel_ModelStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SubModel",
	}).Stage(probe.formStage)
	inst := new(model.SubModel)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ModelStage.Lock()
			defer probe.stageSet.ModelStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ModelStage)
			saveStageSet_SubModel_ModelStage(inst, probe, formGroup)
			probe.stageSet.ModelStage.Commit()
			updateStageSetTable_SubModel_ModelStage(probe)
			probe.ux_tree()
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}


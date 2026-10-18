// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test1/go/models"
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
	case *models.Astruct:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Astruct_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Astruct_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.AstructBstruct2Use:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AstructBstruct2Use_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_AstructBstruct2Use_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.AstructBstructUse:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AstructBstructUse_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_AstructBstructUse_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Bstruct:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bstruct_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Bstruct_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Dstruct:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Dstruct_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Dstruct_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.F0123456789012345678901234567890:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_F0123456789012345678901234567890_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_F0123456789012345678901234567890_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Gstruct:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Gstruct_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Gstruct_Stage(probe)
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

func saveStageSet_Astruct_Stage(
	inst *models.Astruct,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Associationtob":
			StageSetFormDivSelectFieldToField(&inst.Associationtob, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), formDiv)
		case "Anotherassociationtob_2":
			StageSetFormDivSelectFieldToField(&inst.Anotherassociationtob_2, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), formDiv)
		case "Date":
			FormDivTimeFieldToField(&inst.Date, formDiv, false)
		case "Date2":
			FormDivTimeFieldToField(&inst.Date2, formDiv, false)
		case "Booleanfield":
			FormDivBasicFieldToField(&inst.Booleanfield, formDiv)
		case "Aenum":
			StageSetFormDivEnumStringFieldToField(&inst.Aenum, formDiv)
		case "Aenum_2":
			StageSetFormDivEnumStringFieldToField(&inst.Aenum_2, formDiv)
		case "Benum":
			StageSetFormDivEnumStringFieldToField(&inst.Benum, formDiv)
		case "CEnum":
			StageSetFormDivEnumIntFieldToField(&inst.CEnum, formDiv)
		case "CName":
			FormDivBasicFieldToField(&inst.CName, formDiv)
		case "CFloatfield":
			FormDivBasicFieldToField(&inst.CFloatfield, formDiv)
		case "Bstruct":
			StageSetFormDivSelectFieldToField(&inst.Bstruct, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), formDiv)
		case "Bstruct2":
			StageSetFormDivSelectFieldToField(&inst.Bstruct2, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), formDiv)
		case "Dstruct":
			StageSetFormDivSelectFieldToField(&inst.Dstruct, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), formDiv)
		case "Dstruct2":
			StageSetFormDivSelectFieldToField(&inst.Dstruct2, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), formDiv)
		case "Dstruct3":
			StageSetFormDivSelectFieldToField(&inst.Dstruct3, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), formDiv)
		case "Dstruct4":
			StageSetFormDivSelectFieldToField(&inst.Dstruct4, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&inst.Floatfield, formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&inst.Intfield, formDiv)
		case "Anotherbooleanfield":
			FormDivBasicFieldToField(&inst.Anotherbooleanfield, formDiv)
		case "Duration1":
			FormDivBasicFieldToField(&inst.Duration1, formDiv)
		case "AnAstruct":
			StageSetFormDivSelectFieldToField(&inst.AnAstruct, probe.stageSet.Stage.GetInstancesSet[*models.Astruct](), formDiv)
		case "TextFieldBespokeSize":
			FormDivBasicFieldToField(&inst.TextFieldBespokeSize, formDiv)
		case "TextArea":
			FormDivBasicFieldToField(&inst.TextArea, formDiv)
		}
	}
}

func saveStageSet_AstructBstruct2Use_Stage(
	inst *models.AstructBstruct2Use,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Bstrcut2":
			StageSetFormDivSelectFieldToField(&inst.Bstrcut2, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), formDiv)
		}
	}
}

func saveStageSet_AstructBstructUse_Stage(
	inst *models.AstructBstructUse,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Bstruct2":
			StageSetFormDivSelectFieldToField(&inst.Bstruct2, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), formDiv)
		}
	}
}

func saveStageSet_Bstruct_Stage(
	inst *models.Bstruct,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&inst.Floatfield, formDiv)
		case "Floatfield2":
			FormDivBasicFieldToField(&inst.Floatfield2, formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&inst.Intfield, formDiv)
		case "ToBeIgnored":
			FormDivBasicFieldToField(&inst.ToBeIgnored, formDiv)
		}
	}
}

func saveStageSet_Dstruct_Stage(
	inst *models.Dstruct,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Gstruct":
			StageSetFormDivSelectFieldToField(&inst.Gstruct, probe.stageSet.Stage.GetInstancesSet[*models.Gstruct](), formDiv)
		}
	}
}

func saveStageSet_F0123456789012345678901234567890_Stage(
	inst *models.F0123456789012345678901234567890,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Date":
			FormDivTimeFieldToField(&inst.Date, formDiv, false)
		}
	}
}

func saveStageSet_Gstruct_Stage(
	inst *models.Gstruct,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Floatfield":
			FormDivBasicFieldToField(&inst.Floatfield, formDiv)
		case "Floatfield2":
			FormDivBasicFieldToField(&inst.Floatfield2, formDiv)
		case "Intfield":
			FormDivBasicFieldToField(&inst.Intfield, formDiv)
		case "ToBeIgnored":
			FormDivBasicFieldToField(&inst.ToBeIgnored, formDiv)
		}
	}
}


func StageSetNewInstance_Astruct_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Astruct",
	}).Stage(probe.formStage)
	inst := new(models.Astruct)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Astruct_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Astruct_Stage(probe)
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

func StageSetNewInstance_AstructBstruct2Use_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AstructBstruct2Use",
	}).Stage(probe.formStage)
	inst := new(models.AstructBstruct2Use)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_AstructBstruct2Use_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_AstructBstruct2Use_Stage(probe)
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

func StageSetNewInstance_AstructBstructUse_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AstructBstructUse",
	}).Stage(probe.formStage)
	inst := new(models.AstructBstructUse)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_AstructBstructUse_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_AstructBstructUse_Stage(probe)
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

func StageSetNewInstance_Bstruct_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Bstruct",
	}).Stage(probe.formStage)
	inst := new(models.Bstruct)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Bstruct_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Bstruct_Stage(probe)
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

func StageSetNewInstance_Dstruct_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Dstruct",
	}).Stage(probe.formStage)
	inst := new(models.Dstruct)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Dstruct_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Dstruct_Stage(probe)
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

func StageSetNewInstance_F0123456789012345678901234567890_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New F0123456789012345678901234567890",
	}).Stage(probe.formStage)
	inst := new(models.F0123456789012345678901234567890)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_F0123456789012345678901234567890_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_F0123456789012345678901234567890_Stage(probe)
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

func StageSetNewInstance_Gstruct_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Gstruct",
	}).Stage(probe.formStage)
	inst := new(models.Gstruct)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Gstruct_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Gstruct_Stage(probe)
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


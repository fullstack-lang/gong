// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
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
	case *models.Button:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Button_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Button_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Menu:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Menu_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Menu_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Node:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Node_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Node_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SVGIcon:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SVGIcon_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SVGIcon_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Tree:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tree_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Tree_Stage(probe)
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

func saveStageSet_Button_Stage(
	inst *models.Button,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Icon":
			FormDivBasicFieldToField(&inst.Icon, formDiv)
		case "SVGIcon":
			StageSetFormDivSelectFieldToField(&inst.SVGIcon, probe.stageSet.Stage.GetInstancesSet[*models.SVGIcon](), formDiv)
		case "IsDisabled":
			FormDivBasicFieldToField(&inst.IsDisabled, formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&inst.HasToolTip, formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&inst.ToolTipText, formDiv)
		case "ToolTipPosition":
			StageSetFormDivEnumStringFieldToField(&inst.ToolTipPosition, formDiv)
		case "ClientOnX":
			FormDivBasicFieldToField(&inst.ClientOnX, formDiv)
		case "ClientOnY":
			FormDivBasicFieldToField(&inst.ClientOnY, formDiv)
		}
	}
}

func saveStageSet_Menu_Stage(
	inst *models.Menu,
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

func saveStageSet_Node_Stage(
	inst *models.Node,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsWithPrefix":
			FormDivBasicFieldToField(&inst.IsWithPrefix, formDiv)
		case "Prefix":
			FormDivBasicFieldToField(&inst.Prefix, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "BackgroundColor":
			FormDivBasicFieldToField(&inst.BackgroundColor, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "HasCheckboxButton":
			FormDivBasicFieldToField(&inst.HasCheckboxButton, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "IsCheckboxDisabled":
			FormDivBasicFieldToField(&inst.IsCheckboxDisabled, formDiv)
		case "CheckboxHasToolTip":
			FormDivBasicFieldToField(&inst.CheckboxHasToolTip, formDiv)
		case "CheckboxToolTipText":
			FormDivBasicFieldToField(&inst.CheckboxToolTipText, formDiv)
		case "CheckboxToolTipPosition":
			StageSetFormDivEnumStringFieldToField(&inst.CheckboxToolTipPosition, formDiv)
		case "HasSecondCheckboxButton":
			FormDivBasicFieldToField(&inst.HasSecondCheckboxButton, formDiv)
		case "IsSecondCheckboxChecked":
			FormDivBasicFieldToField(&inst.IsSecondCheckboxChecked, formDiv)
		case "IsSecondCheckboxDisabled":
			FormDivBasicFieldToField(&inst.IsSecondCheckboxDisabled, formDiv)
		case "SecondCheckboxHasToolTip":
			FormDivBasicFieldToField(&inst.SecondCheckboxHasToolTip, formDiv)
		case "SecondCheckboxToolTipText":
			FormDivBasicFieldToField(&inst.SecondCheckboxToolTipText, formDiv)
		case "SecondCheckboxToolTipPosition":
			StageSetFormDivEnumStringFieldToField(&inst.SecondCheckboxToolTipPosition, formDiv)
		case "TextAfterSecondCheckbox":
			FormDivBasicFieldToField(&inst.TextAfterSecondCheckbox, formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&inst.HasToolTip, formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&inst.ToolTipText, formDiv)
		case "ToolTipPosition":
			StageSetFormDivEnumStringFieldToField(&inst.ToolTipPosition, formDiv)
		case "ClientOnY":
			FormDivBasicFieldToField(&inst.ClientOnY, formDiv)
		case "IsInEditMode":
			FormDivBasicFieldToField(&inst.IsInEditMode, formDiv)
		case "IsNodeClickable":
			FormDivBasicFieldToField(&inst.IsNodeClickable, formDiv)
		case "IsWithPreceedingIcon":
			FormDivBasicFieldToField(&inst.IsWithPreceedingIcon, formDiv)
		case "PreceedingIcon":
			FormDivBasicFieldToField(&inst.PreceedingIcon, formDiv)
		case "PreceedingSVGIcon":
			StageSetFormDivSelectFieldToField(&inst.PreceedingSVGIcon, probe.stageSet.Stage.GetInstancesSet[*models.SVGIcon](), formDiv)
		case "Menu":
			StageSetFormDivSelectFieldToField(&inst.Menu, probe.stageSet.Stage.GetInstancesSet[*models.Menu](), formDiv)
		}
	}
}

func saveStageSet_SVGIcon_Stage(
	inst *models.SVGIcon,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SVG":
			FormDivBasicFieldToField(&inst.SVG, formDiv)
		}
	}
}

func saveStageSet_Tree_Stage(
	inst *models.Tree,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "HaveSearch":
			FormDivBasicFieldToField(&inst.HaveSearch, formDiv)
		}
	}
}


func StageSetNewInstance_Button_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Button",
	}).Stage(probe.formStage)
	inst := new(models.Button)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Button_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Button_Stage(probe)
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

func StageSetNewInstance_Menu_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Menu",
	}).Stage(probe.formStage)
	inst := new(models.Menu)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Menu_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Menu_Stage(probe)
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

func StageSetNewInstance_Node_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Node",
	}).Stage(probe.formStage)
	inst := new(models.Node)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Node_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Node_Stage(probe)
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

func StageSetNewInstance_SVGIcon_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SVGIcon",
	}).Stage(probe.formStage)
	inst := new(models.SVGIcon)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SVGIcon_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SVGIcon_Stage(probe)
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

func StageSetNewInstance_Tree_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tree",
	}).Stage(probe.formStage)
	inst := new(models.Tree)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Tree_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Tree_Stage(probe)
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


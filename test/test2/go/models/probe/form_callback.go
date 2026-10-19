// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
)

type FormCallbackIF = gongprobe.FormCallbackIF

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func saveStageSet_A_Stage(
	inst *models.A,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumberField":
			FormDivBasicFieldToField(&inst.NumberField, formDiv)
		case "B":
			StageSetFormDivSelectFieldToField(&inst.B, probe.stageOfInterest.GetInstancesSet[*models.B](), formDiv)
		case "Bs":
			FormDivSliceOfPointersToField(inst, "Bs", &inst.Bs, formDiv, probe)
		case "X":
			StageSetFormDivSelectFieldToField(&inst.X, probe.stageOfInterest.XStage.GetInstancesSet[*x.X](), formDiv)
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
	probe *Probe,
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
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Y":
			StageSetFormDivSelectFieldToField(&inst.Y, probe.stageOfInterest.YStage.GetInstancesSet[*y.Y](), formDiv)
		}
	}
}

func saveStageSet_Y_YStage(
	inst *y.Y,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SubModel":
			StageSetFormDivSelectFieldToField(&inst.SubModel, probe.stageOfInterest.ModelStage.GetInstancesSet[*model.SubModel](), formDiv)
		}
	}
}

func saveStageSet_SubModel_ModelStage(
	inst *model.SubModel,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}


func StageSetNewInstance_A_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A",
	}).Stage(probe.formStage)
	inst := new(models.A)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_B_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New B",
	}).Stage(probe.formStage)
	inst := new(models.B)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_B_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_B_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_X_XStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New X",
	}).Stage(probe.formStage)
	inst := new(x.X)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XStage.Lock()
			defer probe.stageOfInterest.XStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XStage)
			saveStageSet_X_XStage(inst, probe, formGroup)
			probe.stageOfInterest.XStage.Commit()
			updateStageSetTable_X_XStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Y_YStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Y",
	}).Stage(probe.formStage)
	inst := new(y.Y)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.YStage.Lock()
			defer probe.stageOfInterest.YStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.YStage)
			saveStageSet_Y_YStage(inst, probe, formGroup)
			probe.stageOfInterest.YStage.Commit()
			updateStageSetTable_Y_YStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SubModel_ModelStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SubModel",
	}).Stage(probe.formStage)
	inst := new(model.SubModel)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.ModelStage.Lock()
			defer probe.stageOfInterest.ModelStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.ModelStage)
			saveStageSet_SubModel_ModelStage(inst, probe, formGroup)
			probe.stageOfInterest.ModelStage.Commit()
			updateStageSetTable_SubModel_ModelStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}


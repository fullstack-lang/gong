// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/refs"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/common"
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

func saveStageSet_Catalog_Stage(
	inst *models.Catalog,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Item":
			FormDivSliceOfPointersToField(inst, "Item", &inst.Item, formDiv, probe)
		}
	}
}

func saveStageSet_DerivedItemType_Stage(
	inst *models.DerivedItemType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Code":
			FormDivBasicFieldToField(&inst.Code, formDiv)
		case "Score":
			FormDivBasicFieldToField(&inst.Score, formDiv)
		case "TargetRef":
			FormDivBasicFieldToField(&inst.TargetRef, formDiv)
		case "OptionA":
			FormDivBasicFieldToField(&inst.OptionA, formDiv)
		case "OptionB":
			FormDivBasicFieldToField(&inst.OptionB, formDiv)
		}
	}
}

func saveStageSet_RefTargetType_RefsStage(
	inst *refs.RefTargetType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Status":
			StageSetFormDivEnumStringFieldToField(&inst.Status, formDiv)
		case "CreatedBy":
			FormDivBasicFieldToField(&inst.CreatedBy, formDiv)
		case "Label":
			FormDivBasicFieldToField(&inst.Label, formDiv)
		}
	}
}

func saveStageSet_BaseItemType_CommonStage(
	inst *common.BaseItemType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Code":
			FormDivBasicFieldToField(&inst.Code, formDiv)
		case "Score":
			FormDivBasicFieldToField(&inst.Score, formDiv)
		}
	}
}


func StageSetNewInstance_Catalog_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Catalog",
	}).Stage(probe.formStage)
	inst := new(models.Catalog)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_Catalog_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_Catalog_Stage(probe)
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

func StageSetNewInstance_DerivedItemType_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DerivedItemType",
	}).Stage(probe.formStage)
	inst := new(models.DerivedItemType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DerivedItemType_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DerivedItemType_Stage(probe)
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

func StageSetNewInstance_RefTargetType_RefsStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RefTargetType",
	}).Stage(probe.formStage)
	inst := new(refs.RefTargetType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.RefsStage.Lock()
			defer probe.stageOfInterest.RefsStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.RefsStage)
			saveStageSet_RefTargetType_RefsStage(inst, probe, formGroup)
			probe.stageOfInterest.RefsStage.Commit()
			updateStageSetTable_RefTargetType_RefsStage(probe)
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

func StageSetNewInstance_BaseItemType_CommonStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New BaseItemType",
	}).Stage(probe.formStage)
	inst := new(common.BaseItemType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.CommonStage.Lock()
			defer probe.stageOfInterest.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.CommonStage)
			saveStageSet_BaseItemType_CommonStage(inst, probe, formGroup)
			probe.stageOfInterest.CommonStage.Commit()
			updateStageSetTable_BaseItemType_CommonStage(probe)
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


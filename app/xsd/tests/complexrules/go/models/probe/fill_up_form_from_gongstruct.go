// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/refs"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/common"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)
}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {
	switch inst := instance.(type) {
	case *models.Catalog:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Catalog",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Catalog_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Catalog_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.DerivedItemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DerivedItemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DerivedItemType_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_DerivedItemType_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *refs.RefTargetType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RefTargetType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.RefsStage.Lock()
				defer probe.stageOfInterest.RefsStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RefTargetType_RefsStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.RefsStage)
				}
				probe.stageOfInterest.RefsStage.Commit()
				updateStageSetTable_RefTargetType_RefsStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *common.BaseItemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "BaseItemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.CommonStage.Lock()
				defer probe.stageOfInterest.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BaseItemType_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.CommonStage)
				}
				probe.stageOfInterest.CommonStage.Commit()
				updateStageSetTable_BaseItemType_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	default:
		_ = inst
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}

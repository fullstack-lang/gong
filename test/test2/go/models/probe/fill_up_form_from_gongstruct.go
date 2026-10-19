// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)
}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {
	switch inst := instance.(type) {
	case *models.A:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.B:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "B",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_B_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_B_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *x.X:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "X",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XStage.Lock()
				defer probe.stageOfInterest.XStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_X_XStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XStage)
				}
				probe.stageOfInterest.XStage.Commit()
				updateStageSetTable_X_XStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *y.Y:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Y",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.YStage.Lock()
				defer probe.stageOfInterest.YStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Y_YStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.YStage)
				}
				probe.stageOfInterest.YStage.Commit()
				updateStageSetTable_Y_YStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *model.SubModel:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SubModel",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.ModelStage.Lock()
				defer probe.stageOfInterest.ModelStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SubModel_ModelStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.ModelStage)
				}
				probe.stageOfInterest.ModelStage.Commit()
				updateStageSetTable_SubModel_ModelStage(probe)
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

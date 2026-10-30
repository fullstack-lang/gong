// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)
}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {
	switch inst := instance.(type) {
	case *models.AIXMBasicMessage:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AIXMBasicMessage",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AIXMBasicMessage_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_AIXMBasicMessage_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.AIXMBasicMessageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AIXMBasicMessageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AIXMBasicMessageType_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_AIXMBasicMessageType_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.BasicMessageMemberAIXMPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "BasicMessageMemberAIXMPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BasicMessageMemberAIXMPropertyType_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_BasicMessageMemberAIXMPropertyType_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_1:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_1",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_1_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_1_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_10:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_10",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_10_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_10_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_11:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_11",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_11_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_11_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_12:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_12",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_12_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_12_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_13:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_13",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_13_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_13_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_14:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_14",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_14_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_14_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_15:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_15",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_15_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_15_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_16:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_16",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_16_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_16_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_17:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_17",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_17_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_17_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_18:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_18",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_18_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_18_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_19:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_19",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_19_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_19_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_2:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_2",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_2_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_2_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_20:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_20",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_20_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_20_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_21:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_21",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_21_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_21_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_22:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_22",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_22_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_22_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_23:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_23",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_23_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_23_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_24:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_24",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_24_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_24_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_25:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_25",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_25_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_25_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_26:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_26",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_26_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_26_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_27:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_27",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_27_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_27_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_28:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_28",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_28_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_28_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_29:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_29",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_29_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_29_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_3:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_3",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_3_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_3_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_30:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_30",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_30_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_30_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_31:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_31",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_31_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_31_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_32:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_32",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_32_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_32_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_33:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_33",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_33_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_33_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_34:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_34",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_34_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_34_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_35:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_35",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_35_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_35_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_36:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_36",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_36_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_36_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_37:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_37",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_37_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_37_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_38:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_38",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_38_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_38_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_4:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_4",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_4_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_4_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_5:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_5",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_5_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_5_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_6:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_6",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_6_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_6_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_7:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_7",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_7_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_7_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_8:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_8",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_8_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_8_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_9:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_9",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_9_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_9_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractAirportGroundServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAirportGroundServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAirportGroundServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractAirportGroundServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractPropertiesWithScheduleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractPropertiesWithScheduleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractPropertiesWithScheduleType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractTrafficSeparationServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractTrafficSeparationServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractTrafficSeparationServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractTrafficSeparationServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AerialRefuellingPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AerialRefuellingPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AerialRefuellingPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AerialRefuellingPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirTrafficControlServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirTrafficControlServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirTrafficControlServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirTrafficControlServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirTrafficControlServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirTrafficControlServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirTrafficControlServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirTrafficControlServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirTrafficControlServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirTrafficControlServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirTrafficControlServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirTrafficManagementServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirTrafficManagementServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirTrafficManagementServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirTrafficManagementServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirTrafficManagementServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirTrafficManagementServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirTrafficManagementServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirTrafficManagementServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirTrafficManagementServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirTrafficManagementServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirTrafficManagementServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftGroundServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftGroundServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftGroundServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftGroundServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftGroundServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftGroundServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftGroundServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftGroundServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftGroundServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftGroundServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftGroundServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftGroundServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportClearanceServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportClearanceServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportClearanceServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportClearanceServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportClearanceServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportClearanceServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportClearanceServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportClearanceServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportClearanceServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportClearanceServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportClearanceServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportClearanceServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportSuppliesServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportSuppliesServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportSuppliesServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportSuppliesServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportSuppliesServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportSuppliesServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportSuppliesServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportSuppliesServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportSuppliesServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportSuppliesServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportSuppliesServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspacePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspacePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspacePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspacePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CallsignDetailPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CallsignDetailPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CallsignDetailPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CallsignDetailPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CallsignDetailType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CallsignDetailType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CallsignDetailType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CallsignDetailType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ContactInformationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ContactInformationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ContactInformationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ContactInformationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ContactInformationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ContactInformationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ContactInformationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ContactInformationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DesignatedPointPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DesignatedPointPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DesignatedPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DesignatedPointPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DirectionFinderPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DirectionFinderPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectionFinderPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DirectionFinderPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevatedPointPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevatedPointPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevatedPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevatedPointPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevatedPointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevatedPointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevatedPointType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevatedPointType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FireFightingServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FireFightingServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FireFightingServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FireFightingServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FireFightingServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FireFightingServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FireFightingServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FireFightingServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FireFightingServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FireFightingServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FireFightingServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FireFightingServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FuelPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FuelPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FuelPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FuelPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FuelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FuelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FuelType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FuelType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GroundLightSystemPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GroundLightSystemPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroundLightSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GroundLightSystemPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GroundTrafficControlServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroundTrafficControlServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GroundTrafficControlServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GroundTrafficControlServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GroundTrafficControlServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroundTrafficControlServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GroundTrafficControlServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GroundTrafficControlServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GroundTrafficControlServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroundTrafficControlServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GroundTrafficControlServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.HoldingPatternPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "HoldingPatternPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HoldingPatternPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_HoldingPatternPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.InformationServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "InformationServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InformationServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_InformationServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.InformationServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "InformationServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InformationServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_InformationServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.InformationServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "InformationServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InformationServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_InformationServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LightActivationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LightActivationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LightActivationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LightActivationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LightActivationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LightActivationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LightActivationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LightActivationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LinguisticNotePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LinguisticNotePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LinguisticNotePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LinguisticNotePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LinguisticNoteType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LinguisticNoteType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LinguisticNoteType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LinguisticNoteType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NitrogenPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NitrogenPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NitrogenPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NitrogenPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NitrogenType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NitrogenType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NitrogenType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NitrogenType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NotePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NotePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NotePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NotePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NoteType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NoteType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NoteType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OilPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OilPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OilPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OilPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OilType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OilType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OilType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OilType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OnlineContactPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OnlineContactPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OnlineContactPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OnlineContactPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OnlineContactType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OnlineContactType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OnlineContactType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OnlineContactType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OrganisationAuthorityAssociationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrganisationAuthorityAssociationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrganisationAuthorityAssociationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OrganisationAuthorityAssociationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OrganisationAuthorityAssociationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrganisationAuthorityAssociationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrganisationAuthorityAssociationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OrganisationAuthorityAssociationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OrganisationAuthorityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrganisationAuthorityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrganisationAuthorityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OrganisationAuthorityTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrganisationAuthorityTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrganisationAuthorityTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OrganisationAuthorityTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OrganisationAuthorityTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrganisationAuthorityTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrganisationAuthorityTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OrganisationAuthorityTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OrganisationAuthorityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrganisationAuthorityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrganisationAuthorityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OrganisationAuthorityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OxygenPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OxygenPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OxygenPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OxygenPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.OxygenType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OxygenType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OxygenType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_OxygenType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PassengerServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PassengerServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PassengerServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PilotControlledLightingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PilotControlledLightingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PilotControlledLightingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PilotControlledLightingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PilotControlledLightingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PilotControlledLightingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PilotControlledLightingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PilotControlledLightingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PilotControlledLightingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PilotControlledLightingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PilotControlledLightingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PilotControlledLightingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PointPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PointPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PointPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PointType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PointType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PostalAddressPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PostalAddressPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PostalAddressPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PostalAddressPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PostalAddressType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PostalAddressType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PostalAddressType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PostalAddressType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ProcedurePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ProcedurePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ProcedurePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ProcedurePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadioCommunicationChannelPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadioCommunicationChannelPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadioCommunicationChannelPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadioCommunicationChannelPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadioCommunicationChannelTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadioCommunicationChannelTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadioCommunicationChannelTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadioCommunicationChannelTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadioCommunicationChannelTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadioCommunicationChannelTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadioCommunicationChannelTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadioCommunicationChannelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadioCommunicationChannelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadioCommunicationChannelType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadioCommunicationChannelType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadioCommunicationOperationalStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadioCommunicationOperationalStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadioCommunicationOperationalStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadioCommunicationOperationalStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadioCommunicationOperationalStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadioCommunicationOperationalStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadioCommunicationOperationalStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadioCommunicationOperationalStatusType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RoutePortionPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RoutePortionPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RoutePortionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RoutePortionPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RoutePortionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RoutePortionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RoutePortionType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RoutePortionType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RoutePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RoutePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RoutePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RoutePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayCentrelinePointPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayCentrelinePointPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayCentrelinePointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayCentrelinePointPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SearchRescueServiceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SearchRescueServiceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SearchRescueServiceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SearchRescueServiceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SearchRescueServiceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SearchRescueServiceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SearchRescueServiceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SearchRescueServiceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SearchRescueServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SearchRescueServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SearchRescueServiceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SearchRescueServiceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ServiceOperationalStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ServiceOperationalStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ServiceOperationalStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ServiceOperationalStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ServiceOperationalStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ServiceOperationalStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ServiceOperationalStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ServiceOperationalStatusType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialDateTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialDateTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialDateTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialDateTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialDateTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialDateTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialDateTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialDateTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialDateType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialDateType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialDateType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialDateType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelColumnPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelColumnPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelColumnPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelColumnPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelColumnTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelColumnTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelColumnTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelColumnTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelColumnTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelColumnTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelColumnTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelColumnTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelColumnType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelColumnType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelColumnType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelColumnType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelSectorTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelSectorTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelSectorTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelSectorTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelSectorTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelSectorTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelSectorTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelSectorTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelSectorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelSectorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelSectorType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelSectorType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelTablePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelTablePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelTablePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelTablePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelTableTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelTableTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelTableTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelTableTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelTableTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelTableTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelTableTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelTableTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelTableType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelTableType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelTableType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelTableType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandardLevelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandardLevelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandardLevelType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandardLevelType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TelephoneContactPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TelephoneContactPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TelephoneContactPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TelephoneContactPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TelephoneContactType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TelephoneContactType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TelephoneContactType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TelephoneContactType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TimesheetPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimesheetPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimesheetPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TimesheetPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TimesheetType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimesheetType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimesheetType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TimesheetType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitAvailabilityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitAvailabilityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitAvailabilityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitAvailabilityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitAvailabilityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitAvailabilityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitDependencyPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitDependencyPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitDependencyPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitDependencyPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitDependencyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitDependencyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitDependencyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitDependencyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.UnitType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "UnitType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_UnitType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_UnitType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VORPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VORPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VORPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VORPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftGroundServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftGroundServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftGroundServiceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftGroundServiceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAviationStandardsType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAviationStandardsType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAviationStandardsType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAviationStandardsType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeCommunicationChannelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeCommunicationChannelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeCommunicationChannelType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeCommunicationChannelType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeCommunicationDirectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeCommunicationDirectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeCommunicationDirectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeCommunicationDirectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeCommunicationModeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeCommunicationModeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeCommunicationModeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeCommunicationModeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDayType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDayType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDayType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDayType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFacilityRankingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFacilityRankingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFacilityRankingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFacilityRankingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFireFightingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFireFightingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFireFightingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFireFightingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFlightDestinationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFlightDestinationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFlightDestinationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFlightDestinationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFlightRuleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFlightRuleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFlightRuleType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFlightRuleType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFuelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFuelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFuelType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFuelType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeIntensityStandByType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeIntensityStandByType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeIntensityStandByType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeIntensityStandByType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLanguageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLanguageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLanguageType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLanguageType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLevelSeriesType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLevelSeriesType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLevelSeriesType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLevelSeriesType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLevelTableDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLevelTableDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLevelTableDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLevelTableDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLightIntensityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLightIntensityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLightIntensityType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLightIntensityType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMilitaryOperationsType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMilitaryOperationsType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMilitaryOperationsType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMilitaryOperationsType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNitrogenType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNitrogenType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNitrogenType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNitrogenType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNorthReferenceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNorthReferenceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNorthReferenceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNorthReferenceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNotePurposeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNotePurposeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNotePurposeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNotePurposeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeOilType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOilType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOilType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOilType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeOrganisationDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOrganisationDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOrganisationDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOrganisationDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeOrganisationHierarchyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOrganisationHierarchyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOrganisationHierarchyType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOrganisationHierarchyType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeOrganisationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOrganisationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOrganisationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOrganisationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeOxygenType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOxygenType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOxygenType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOxygenType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePassengerServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePassengerServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePassengerServiceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePassengerServiceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePilotControlledLightingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePilotControlledLightingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePilotControlledLightingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePilotControlledLightingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRVSMType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRVSMType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRVSMType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRVSMType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRadioEmissionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRadioEmissionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRadioEmissionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRadioEmissionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeServiceATCType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeServiceATCType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeServiceATCType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeServiceATCType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeServiceATFMType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeServiceATFMType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeServiceATFMType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeServiceATFMType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeServiceGroundControlType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeServiceGroundControlType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeServiceGroundControlType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeServiceGroundControlType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeServiceInformationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeServiceInformationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeServiceInformationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeServiceInformationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeServiceSARType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeServiceSARType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeServiceSARType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeServiceSARType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSpecialDateType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSpecialDateType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSpecialDateType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSpecialDateType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeStatusOperationsType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStatusOperationsType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStatusOperationsType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStatusOperationsType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeStatusServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStatusServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStatusServiceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStatusServiceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSystemActivationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSystemActivationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSystemActivationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSystemActivationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTelecomNetworkType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTelecomNetworkType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTelecomNetworkType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTelecomNetworkType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTimeEventCombinationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTimeEventCombinationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTimeEventCombinationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTimeEventCombinationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTimeEventType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTimeEventType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTimeEventType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTimeEventType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTimeReferenceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTimeReferenceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTimeReferenceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTimeReferenceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeUnitDependencyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeUnitDependencyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeUnitDependencyType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeUnitDependencyType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeUnitType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeUnitType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeUnitType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeUnitType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeVerticalDatumType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVerticalDatumType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVerticalDatumType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVerticalDatumType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeYesNoType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeYesNoType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeYesNoType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeYesNoType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.DateMonthDayType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DateMonthDayType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DateMonthDayType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_DateMonthDayType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.DateYearType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DateYearType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DateYearType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_DateYearType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.NoNumberType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NoNumberType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoNumberType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_NoNumberType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextAddressType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextAddressType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextAddressType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextAddressType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextInstructionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextInstructionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextInstructionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextInstructionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextNameType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextNameType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextNameType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextNameType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextNoteType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextNoteType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextNoteType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextNoteType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextPhoneType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextPhoneType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextPhoneType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextPhoneType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TextPropertyNameType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TextPropertyNameType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextPropertyNameType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TextPropertyNameType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.TimeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_TimeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValBearingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValBearingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValBearingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValBearingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValDistanceSignedType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValDistanceSignedType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValDistanceSignedType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValDistanceSignedType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValDistanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValDistanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValDistanceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValDistanceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValDistanceVerticalType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValDistanceVerticalType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValDistanceVerticalType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValDistanceVerticalType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValDurationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValDurationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValDurationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValDurationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValFrequencyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValFrequencyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValFrequencyType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValFrequencyType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMFeatureBaseType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMFeatureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMFeatureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMMessageBaseType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMMessageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMMessageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMObjectType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMObjectType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMTimeSliceBaseType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAIXMTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.AbstractExtensionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractExtensionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractExtensionType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FeatureMetadataPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FeatureTimeSliceMetadataPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmabstractgmlobjecttypes.MessageMetadataPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MessageMetadataPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
				defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				}
				probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
				updateStageSetTable_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractFeatureMemberType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractFeatureMemberType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractFeatureMemberType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractFeatureMemberType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractFeatureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractFeatureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractFeatureType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractFeatureType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractGMLType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractGMLType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractGMLType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractGMLType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractGeometricPrimitiveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractGeometricPrimitiveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractGeometricPrimitiveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractGeometricPrimitiveType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractGeometryType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractGeometryType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractGeometryType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractGeometryType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractMetadataPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractMetadataPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractMetadataPropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractMetadataPropertyType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractTimeGeometricPrimitiveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractTimeGeometricPrimitiveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractTimeGeometricPrimitiveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractTimeGeometricPrimitiveType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractTimeObjectType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractTimeObjectType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractTimeObjectType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractTimeObjectType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractTimePrimitiveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractTimePrimitiveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractTimePrimitiveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractTimePrimitiveType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractTimeSliceType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractTimeSliceType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.BoundingShapeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "BoundingShapeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BoundingShapeType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_BoundingShapeType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.CodeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CodeType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.CodeWithAuthorityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeWithAuthorityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeWithAuthorityType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CodeWithAuthorityType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.DirectPositionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DirectPositionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectPositionType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_DirectPositionType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.DynamicFeatureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DynamicFeatureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DynamicFeatureType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_DynamicFeatureType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.EnvelopeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "EnvelopeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EnvelopeType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_EnvelopeType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.PointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PointType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_PointType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.TimeInstantType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimeInstantType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimeInstantType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_TimeInstantType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.TimePeriodType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimePeriodType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimePeriodType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_TimePeriodType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.TimePositionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimePositionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimePositionType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_TimePositionType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.TimePrimitivePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TimePrimitivePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimePrimitivePropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_TimePrimitivePropertyType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.Xlink_SchemaElement:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Xlink_SchemaElement",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Xlink_SchemaElement_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_Xlink_SchemaElement_XlinkStage(probe)
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

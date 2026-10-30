// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/xlink"
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
	case *aixmfeatures.A_extension_39:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_39",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_39_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_39_AixmfeaturesStage(probe)
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
	case *aixmfeatures.A_extension_40:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_40",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_40_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_40_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_41:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_41",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_41_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_41_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AbstractNavaidEquipmentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractNavaidEquipmentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractNavaidEquipmentType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractNavaidEquipmentType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractNavigationSystemCheckpointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractNavigationSystemCheckpointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AbstractRadarEquipmentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractRadarEquipmentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractRadarEquipmentType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractRadarEquipmentType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractSurveillanceRadarType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractSurveillanceRadarType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractSurveillanceRadarType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractSurveillanceRadarType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AngleIndicationTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AngleIndicationTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AngleIndicationTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AngleIndicationTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AngleIndicationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AngleIndicationTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AngleIndicationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AngleIndicationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AngleIndicationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AngleIndicationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForNavaidEquipmentPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForNavaidEquipmentPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForNavaidEquipmentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForNavaidEquipmentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForNavaidEquipmentType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForNavaidEquipmentType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForSpecialNavigationStationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForSpecialNavigationStationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForSpecialNavigationStationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForSpecialNavigationStationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForSpecialNavigationSystemPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForSpecialNavigationSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForSpecialNavigationSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AzimuthTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AzimuthTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AzimuthTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AzimuthTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AzimuthTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AzimuthTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AzimuthTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AzimuthTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AzimuthType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AzimuthType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AzimuthType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AzimuthType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CheckpointINSTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CheckpointINSTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CheckpointINSTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CheckpointINSTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CheckpointINSTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CheckpointINSTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CheckpointINSType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CheckpointINSType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CheckpointINSType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CheckpointINSType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CheckpointVORTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CheckpointVORTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CheckpointVORTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CheckpointVORTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CheckpointVORTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CheckpointVORTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CheckpointVORType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CheckpointVORType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CheckpointVORType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CheckpointVORType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.DMETimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DMETimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DMETimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DMETimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DMETimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DMETimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DMETimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DMETimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DMEType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DMEType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DMEType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DMEType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.DesignatedPointTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DesignatedPointTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DesignatedPointTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DesignatedPointTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DesignatedPointTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DesignatedPointTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DesignatedPointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DesignatedPointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DesignatedPointType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DesignatedPointType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DirectionFinderTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DirectionFinderTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DirectionFinderTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DirectionFinderTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectionFinderTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DirectionFinderTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DirectionFinderType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DirectionFinderType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectionFinderType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DirectionFinderType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DistanceIndicationTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DistanceIndicationTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DistanceIndicationTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DistanceIndicationTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DistanceIndicationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DistanceIndicationTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DistanceIndicationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DistanceIndicationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DistanceIndicationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DistanceIndicationType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.ElevationTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevationTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevationTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevationTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevationTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevationTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GlidepathTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GlidepathTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GlidepathTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GlidepathTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GlidepathTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GlidepathTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GlidepathTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GlidepathTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GlidepathType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GlidepathType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GlidepathType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GlidepathType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.InformationServicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "InformationServicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InformationServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_InformationServicePropertyType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.LocalizerTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LocalizerTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LocalizerTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LocalizerTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LocalizerTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LocalizerTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LocalizerTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LocalizerTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LocalizerType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LocalizerType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LocalizerType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LocalizerType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkerBeaconTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkerBeaconTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkerBeaconTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkerBeaconTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkerBeaconTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkerBeaconTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkerBeaconType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkerBeaconType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkerBeaconType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkerBeaconType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NDBTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NDBTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NDBTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NDBTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NDBTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NDBTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NDBTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NDBTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NDBType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NDBType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NDBType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NDBType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidComponentPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidComponentPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidComponentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidComponentPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidComponentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidComponentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidComponentType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidComponentType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidEquipmentMonitoringPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidEquipmentMonitoringPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidEquipmentMonitoringType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidEquipmentMonitoringType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidEquipmentMonitoringType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidEquipmentMonitoringType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidEquipmentPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidEquipmentPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidEquipmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidEquipmentPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidOperationalStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidOperationalStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidOperationalStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidOperationalStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidOperationalStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidOperationalStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidOperationalStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidOperationalStatusType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.NavaidTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PrecisionApproachRadarTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PrecisionApproachRadarTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PrecisionApproachRadarTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PrecisionApproachRadarType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PrecisionApproachRadarType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrecisionApproachRadarType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PrecisionApproachRadarType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PrimarySurveillanceRadarTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PrimarySurveillanceRadarTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PrimarySurveillanceRadarTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PrimarySurveillanceRadarType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PrimarySurveillanceRadarType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrimarySurveillanceRadarType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PrimarySurveillanceRadarType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadarComponentPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadarComponentPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadarComponentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadarComponentPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadarComponentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadarComponentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadarComponentType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadarComponentType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadarEquipmentPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadarEquipmentPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadarEquipmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadarEquipmentPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadarSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadarSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadarSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadarSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadarSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadarSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RadarSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RadarSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RadarSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RadarSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ReflectorPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ReflectorPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ReflectorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ReflectorPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ReflectorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ReflectorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ReflectorType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ReflectorType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.RunwayDirectionPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SDFTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SDFTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SDFTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SDFTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SDFTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SDFTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SDFTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SDFTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SDFType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SDFType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SDFType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SDFType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SecondarySurveillanceRadarTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SecondarySurveillanceRadarTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SecondarySurveillanceRadarTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SecondarySurveillanceRadarType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SecondarySurveillanceRadarType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SecondarySurveillanceRadarType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SecondarySurveillanceRadarType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SignificantPointInAirspaceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SignificantPointInAirspaceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SignificantPointInAirspaceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SignificantPointInAirspaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SignificantPointInAirspaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SignificantPointInAirspaceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SignificantPointInAirspaceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationStationStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationStationStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationStationStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationStationStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationStationStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationStationStatusType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationStationTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationStationTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationStationTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationStationTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationStationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationStationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationStationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationStationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationSystemPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationSystemPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurveillanceGroundStationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurveillanceGroundStationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurveillanceGroundStationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurveillanceGroundStationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurveillanceGroundStationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurveillanceGroundStationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurveillanceGroundStationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurveillanceGroundStationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TACANTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TACANTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TACANTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TACANTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TACANTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TACANTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TACANTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TACANTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TACANType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TACANType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TACANType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TACANType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.VORTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VORTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VORTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VORTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VORTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VORTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VORTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VORTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VORType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VORType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VORType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VORType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspacePointPositionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspacePointPositionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspacePointPositionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspacePointPositionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspacePointRoleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspacePointRoleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspacePointRoleType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspacePointRoleType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAltitudeUseType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAltitudeUseType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAltitudeUseType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAltitudeUseType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAuralMorseType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAuralMorseType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAuralMorseType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAuralMorseType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAuthorityRoleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAuthorityRoleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAuthorityRoleType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAuthorityRoleType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeBearingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeBearingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeBearingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeBearingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeCardinalDirectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeCardinalDirectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeCardinalDirectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeCardinalDirectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeCheckpointCategoryType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeCheckpointCategoryType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeCheckpointCategoryType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeCheckpointCategoryType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeCourseQualityILSType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeCourseQualityILSType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeCourseQualityILSType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeCourseQualityILSType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDMEChannelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDMEChannelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDMEChannelType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDMEChannelType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDMEType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDMEType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDMEType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDMEType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeDesignatedPointDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDesignatedPointDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDesignatedPointDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDesignatedPointDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDesignatedPointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDesignatedPointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDesignatedPointType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDesignatedPointType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDirectionReferenceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDirectionReferenceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDirectionReferenceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDirectionReferenceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDistanceIndicationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDistanceIndicationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDistanceIndicationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDistanceIndicationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeEmissionBandType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeEmissionBandType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeEmissionBandType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeEmissionBandType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeILSBackCourseType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeILSBackCourseType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeILSBackCourseType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeILSBackCourseType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeIntegrityLevelILSType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeIntegrityLevelILSType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeIntegrityLevelILSType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeIntegrityLevelILSType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMLSAzimuthType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMLSAzimuthType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMLSAzimuthType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMLSAzimuthType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMLSChannelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMLSChannelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMLSChannelType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMLSChannelType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMarkerBeaconSignalType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMarkerBeaconSignalType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMarkerBeaconSignalType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMarkerBeaconSignalType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNDBUsageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNDBUsageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNDBUsageType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNDBUsageType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNavaidDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNavaidDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNavaidDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNavaidDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNavaidPurposeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNavaidPurposeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNavaidPurposeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNavaidPurposeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNavaidServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNavaidServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNavaidServiceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNavaidServiceType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodePARType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePARType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePARType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePARType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePositionInILSType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePositionInILSType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePositionInILSType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePositionInILSType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePrimaryRadarType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePrimaryRadarType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePrimaryRadarType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePrimaryRadarType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRadarServiceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRadarServiceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRadarServiceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRadarServiceType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeRadioSignalType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRadioSignalType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRadioSignalType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRadioSignalType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeReflectorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeReflectorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeReflectorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeReflectorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSignalPerformanceILSType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSignalPerformanceILSType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSignalPerformanceILSType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSignalPerformanceILSType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSpecialNavigationChainDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSpecialNavigationChainDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSpecialNavigationStationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSpecialNavigationStationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSpecialNavigationStationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSpecialNavigationStationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSpecialNavigationSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSpecialNavigationSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSpecialNavigationSystemType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSpecialNavigationSystemType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeStandbyPowerType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStandbyPowerType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStandbyPowerType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStandbyPowerType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeStatusNavaidType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStatusNavaidType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStatusNavaidType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStatusNavaidType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTACANChannelType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTACANChannelType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTACANChannelType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTACANChannelType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeTransponderType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTransponderType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTransponderType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTransponderType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeVORType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVORType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVORType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVORType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeVerticalReferenceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVerticalReferenceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVerticalReferenceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVerticalReferenceType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.NoSequenceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NoSequenceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoSequenceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_NoSequenceType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.ValAngleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValAngleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValAngleType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValAngleType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.ValMagneticVariationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValMagneticVariationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValMagneticVariationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValMagneticVariationType_AixmdatatypesStage(probe)
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

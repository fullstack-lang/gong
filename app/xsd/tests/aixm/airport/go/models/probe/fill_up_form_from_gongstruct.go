// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/xlink"
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
	case *aixmfeatures.A_extension_42:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_42",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_42_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_42_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_43:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_43",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_43_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_43_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_44:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_44",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_44_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_44_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_45:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_45",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_45_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_45_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_46:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_46",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_46_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_46_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_47:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_47",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_47_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_47_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_48:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_48",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_48_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_48_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_49:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_49",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_49_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_49_AixmfeaturesStage(probe)
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
	case *aixmfeatures.A_extension_50:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_50",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_50_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_50_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_51:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_51",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_51_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_51_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_52:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_52",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_52_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_52_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_53:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_53",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_53_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_53_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_54:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_54",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_54_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_54_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_55:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_55",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_55_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_55_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_56:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_56",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_56_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_56_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_57:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_57",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_57_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_57_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_58:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_58",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_58_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_58_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_59:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_59",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_59_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_59_AixmfeaturesStage(probe)
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
	case *aixmfeatures.A_extension_60:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_60",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_60_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_60_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_61:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_61",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_61_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_61_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_62:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_62",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_62_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_62_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_63:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_63",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_63_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_63_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_64:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_64",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_64_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_64_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_65:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_65",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_65_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_65_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_66:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_66",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_66_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_66_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_67:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_67",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_67_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_67_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_68:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_68",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_68_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_68_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_69:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_69",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_69_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_69_AixmfeaturesStage(probe)
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
	case *aixmfeatures.A_extension_70:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_70",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_70_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_70_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_71:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_71",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_71_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_71_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_72:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_72",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_72_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_72_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_73:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_73",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_73_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_73_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_74:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_74",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_74_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_74_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_75:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_75",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_75_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_75_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_76:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_76",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_76_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_76_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_77:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_77",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_77_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_77_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_78:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_78",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_78_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_78_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_79:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_79",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_79_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_79_AixmfeaturesStage(probe)
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
	case *aixmfeatures.A_extension_80:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_80",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_80_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_80_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_81:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_81",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_81_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_81_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_82:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_82",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_82_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_82_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_83:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_83",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_83_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_83_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_84:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_84",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_84_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_84_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_85:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_85",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_85_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_85_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_86:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_86",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_86_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_86_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_87:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_87",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_87_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_87_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_88:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_88",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_88_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_88_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.A_extension_89:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_extension_89",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_extension_89_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_A_extension_89_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AbstractAirportHeliportProtectionAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractAirportHeliportProtectionAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractAirportHeliportProtectionAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractAirportHeliportProtectionAreaType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractGroundLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractGroundLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractGroundLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractGroundLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractMarkingType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AbstractSurfaceContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractSurfaceContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractSurfaceContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractSurfaceContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AbstractUsageConditionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractUsageConditionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractUsageConditionType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AbstractUsageConditionType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftCharacteristicPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftCharacteristicPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftCharacteristicPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftCharacteristicPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftCharacteristicType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftCharacteristicType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftCharacteristicType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftCharacteristicType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftStandContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftStandContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftStandContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftStandContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftStandContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftStandContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftStandContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftStandContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftStandPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftStandPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftStandPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftStandPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftStandTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftStandTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftStandTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftStandTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftStandTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftStandTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftStandTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftStandTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AircraftStandType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AircraftStandType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AircraftStandType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AircraftStandType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportAvailabilityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportAvailabilityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportAvailabilityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportAvailabilityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportAvailabilityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportAvailabilityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportCollocationTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportCollocationTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportCollocationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportCollocationTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportCollocationTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportCollocationTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportCollocationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportCollocationTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportCollocationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportCollocationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportCollocationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportCollocationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportContaminationType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AirportHeliportProtectionAreaPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportProtectionAreaPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportProtectionAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportProtectionAreaPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportResponsibilityOrganisationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportResponsibilityOrganisationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportResponsibilityOrganisationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportResponsibilityOrganisationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportResponsibilityOrganisationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportResponsibilityOrganisationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportResponsibilityOrganisationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportResponsibilityOrganisationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportUsagePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportUsagePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportUsagePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportUsagePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHeliportUsageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHeliportUsageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHeliportUsageType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHeliportUsageType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHotSpotTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHotSpotTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHotSpotTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHotSpotTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHotSpotTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHotSpotTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHotSpotTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHotSpotTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportHotSpotType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportHotSpotType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportHotSpotType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportHotSpotType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportProtectionAreaMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportProtectionAreaMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportProtectionAreaMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportProtectionAreaMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportProtectionAreaMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportProtectionAreaMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportProtectionAreaMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportProtectionAreaMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportProtectionAreaMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportProtectionAreaMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportProtectionAreaMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportProtectionAreaMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirportSuppliesServicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirportSuppliesServicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirportSuppliesServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirportSuppliesServicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AltimeterSourcePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AltimeterSourcePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AltimeterSourcePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AltimeterSourcePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AltimeterSourceStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AltimeterSourceStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AltimeterSourceStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AltimeterSourceStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AltimeterSourceStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AltimeterSourceStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AltimeterSourceStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AltimeterSourceStatusType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AltimeterSourceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AltimeterSourceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AltimeterSourceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AltimeterSourceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AltimeterSourceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AltimeterSourceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AltimeterSourceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AltimeterSourceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AltimeterSourceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AltimeterSourceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AltimeterSourceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AltimeterSourceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApproachLightingSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApproachLightingSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApproachLightingSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApproachLightingSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApproachLightingSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApproachLightingSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApproachLightingSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApproachLightingSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApproachLightingSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApproachLightingSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApproachLightingSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApproachLightingSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronAreaAvailabilityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronAreaAvailabilityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronAreaAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronAreaAvailabilityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronAreaAvailabilityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronAreaAvailabilityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronAreaAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronAreaAvailabilityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronAreaUsagePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronAreaUsagePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronAreaUsagePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronAreaUsagePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronAreaUsageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronAreaUsageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronAreaUsageType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronAreaUsageType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronElementPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronElementPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronElementPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronElementTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronElementTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronElementTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronElementTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronElementTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronElementTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronElementTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronElementTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronElementType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronElementType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ApronType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ApronType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ApronType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ApronType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ArrestingGearTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArrestingGearTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArrestingGearTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ArrestingGearTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ArrestingGearTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArrestingGearTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArrestingGearTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ArrestingGearTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ArrestingGearType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArrestingGearType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArrestingGearType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ArrestingGearType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ConditionCombinationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ConditionCombinationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConditionCombinationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ConditionCombinationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ConditionCombinationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ConditionCombinationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ConditionCombinationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ConditionCombinationType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.CurveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CurveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CurveType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CurveType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.DeicingAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DeicingAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DeicingAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_DeicingAreaType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevatedCurvePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevatedCurvePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevatedCurvePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevatedCurvePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevatedCurveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevatedCurveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevatedCurveType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevatedCurveType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.ElevatedSurfacePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevatedSurfacePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevatedSurfacePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevatedSurfacePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ElevatedSurfaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ElevatedSurfaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ElevatedSurfaceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ElevatedSurfaceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FlightCharacteristicPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FlightCharacteristicPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FlightCharacteristicPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FlightCharacteristicPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FlightCharacteristicType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FlightCharacteristicType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FlightCharacteristicType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FlightCharacteristicType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FloatingDockSitePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FloatingDockSitePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FloatingDockSitePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FloatingDockSitePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FloatingDockSiteTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FloatingDockSiteTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FloatingDockSiteTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FloatingDockSiteTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FloatingDockSiteTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FloatingDockSiteTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FloatingDockSiteTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FloatingDockSiteTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.FloatingDockSiteType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "FloatingDockSiteType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FloatingDockSiteType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_FloatingDockSiteType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GroundLightingAvailabilityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GroundLightingAvailabilityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroundLightingAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GroundLightingAvailabilityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GroundLightingAvailabilityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GroundLightingAvailabilityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroundLightingAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GroundLightingAvailabilityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLinePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLinePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLinePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLinePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GuidanceLineType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GuidanceLineType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GuidanceLineType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GuidanceLineType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LightElementPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LightElementPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LightElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LightElementPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LightElementStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LightElementStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LightElementStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LightElementStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LightElementStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LightElementStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LightElementStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LightElementStatusType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.LightElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LightElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LightElementType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_LightElementType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.ManoeuvringAreaAvailabilityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ManoeuvringAreaAvailabilityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ManoeuvringAreaAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ManoeuvringAreaAvailabilityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ManoeuvringAreaAvailabilityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ManoeuvringAreaAvailabilityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ManoeuvringAreaAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ManoeuvringAreaAvailabilityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ManoeuvringAreaUsagePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ManoeuvringAreaUsagePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ManoeuvringAreaUsagePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ManoeuvringAreaUsagePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ManoeuvringAreaUsageType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ManoeuvringAreaUsageType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ManoeuvringAreaUsageType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ManoeuvringAreaUsageType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkingBuoyTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkingBuoyTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkingBuoyTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkingBuoyTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkingBuoyTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkingBuoyTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkingBuoyTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkingBuoyTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkingBuoyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkingBuoyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkingBuoyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkingBuoyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkingElementPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkingElementPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkingElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkingElementPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MarkingElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkingElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkingElementType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkingElementType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MeteorologyPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MeteorologyPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeteorologyPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MeteorologyPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.MeteorologyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MeteorologyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeteorologyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MeteorologyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidEquipmentDistancePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidEquipmentDistancePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidEquipmentDistancePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidEquipmentDistancePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NavaidEquipmentDistanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NavaidEquipmentDistanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NavaidEquipmentDistanceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NavaidEquipmentDistanceType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.NonMovementAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NonMovementAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NonMovementAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NonMovementAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NonMovementAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NonMovementAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NonMovementAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NonMovementAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.NonMovementAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "NonMovementAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NonMovementAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_NonMovementAreaType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.PassengerLoadingBridgeTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerLoadingBridgeTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerLoadingBridgeTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerLoadingBridgeTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PassengerLoadingBridgeTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerLoadingBridgeTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerLoadingBridgeTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerLoadingBridgeTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.PassengerLoadingBridgeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerLoadingBridgeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerLoadingBridgeType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerLoadingBridgeType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.RidgePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RidgePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RidgePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RidgePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RidgeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RidgeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RidgeType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RidgeType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RoadTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RoadTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RoadTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RoadTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RoadTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RoadTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RoadTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RoadTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RoadType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RoadType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RoadType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RoadType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayBlastPadTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayBlastPadTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayBlastPadTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayBlastPadTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayBlastPadTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayBlastPadTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayBlastPadTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayBlastPadTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayBlastPadType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayBlastPadType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayBlastPadType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayBlastPadType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.RunwayCentrelinePointTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayCentrelinePointTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayCentrelinePointTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayCentrelinePointTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayCentrelinePointTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayCentrelinePointTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayCentrelinePointTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayCentrelinePointTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayCentrelinePointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayCentrelinePointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayCentrelinePointType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayCentrelinePointType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDeclaredDistancePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDeclaredDistancePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDeclaredDistancePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDeclaredDistancePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDeclaredDistanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDeclaredDistanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDeclaredDistanceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDeclaredDistanceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDeclaredDistanceValuePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDeclaredDistanceValuePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDeclaredDistanceValuePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDeclaredDistanceValuePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDeclaredDistanceValueType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDeclaredDistanceValueType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDeclaredDistanceValueType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDeclaredDistanceValueType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDirectionLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDirectionLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDirectionLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionLightSystemType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.RunwayDirectionTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDirectionTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayDirectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayDirectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayDirectionType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayDirectionType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayElementPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayElementPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayElementPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayElementTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayElementTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayElementTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayElementTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayElementTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayElementTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayElementTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayElementTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayElementType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayElementType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayMarkingType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.RunwayProtectAreaLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayProtectAreaLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayProtectAreaLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayProtectAreaPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayProtectAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayProtectAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayProtectAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayProtectAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayProtectAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayProtectAreaType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwaySectionContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwaySectionContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwaySectionContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwaySectionContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwaySectionContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwaySectionContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwaySectionContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwaySectionContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayVisualRangeTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayVisualRangeTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayVisualRangeTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayVisualRangeTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayVisualRangeTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayVisualRangeTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayVisualRangeTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayVisualRangeTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.RunwayVisualRangeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RunwayVisualRangeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RunwayVisualRangeType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_RunwayVisualRangeType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneLandingAreaPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneLandingAreaPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneLandingAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneLandingAreaPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneLandingAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneLandingAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneLandingAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneLandingAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneLandingAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneLandingAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneLandingAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneLandingAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneLandingAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneLandingAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneLandingAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneLandingAreaType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneRampSitePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneRampSitePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneRampSitePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneRampSitePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneRampSiteTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneRampSiteTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneRampSiteTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneRampSiteTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneRampSiteTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneRampSiteTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneRampSiteTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneRampSiteTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SeaplaneRampSiteType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SeaplaneRampSiteType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SeaplaneRampSiteType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SeaplaneRampSiteType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.StandMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "StandMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StandMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_StandMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurfaceCharacteristicsPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfaceCharacteristicsPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfaceCharacteristicsPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurfaceCharacteristicsPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurfaceCharacteristicsType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfaceCharacteristicsType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfaceCharacteristicsType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurfaceCharacteristicsType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurfaceContaminationLayerPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfaceContaminationLayerPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfaceContaminationLayerPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurfaceContaminationLayerPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurfaceContaminationLayerType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfaceContaminationLayerType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfaceContaminationLayerType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurfaceContaminationLayerType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurfaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfaceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurfaceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurveyControlPointTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurveyControlPointTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurveyControlPointTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurveyControlPointTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurveyControlPointTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurveyControlPointTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurveyControlPointTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurveyControlPointTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SurveyControlPointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurveyControlPointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurveyControlPointType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurveyControlPointType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiHoldingPositionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiHoldingPositionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiHoldingPositionType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiHoldingPositionType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayElementPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayElementPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayElementPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayElementTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayElementTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayElementTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayElementTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayElementTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayElementTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayElementTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayElementTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayElementType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayElementType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayMarkingType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TaxiwayType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TaxiwayType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaxiwayType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TaxiwayType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.TouchDownLiftOffContaminationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffContaminationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffContaminationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffContaminationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffContaminationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffContaminationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffLightSystemTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffLightSystemTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffLightSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffLightSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffLightSystemTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffLightSystemTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffLightSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffLightSystemTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffLightSystemType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffLightSystemType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffLightSystemType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffLightSystemType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffMarkingTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffMarkingTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffMarkingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffMarkingTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffMarkingTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffMarkingTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffMarkingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffMarkingTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffMarkingType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffMarkingType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.TouchDownLiftOffSafeAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffSafeAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffSafeAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffSafeAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffSafeAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffSafeAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffSafeAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffSafeAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffSafeAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffSafeAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffSafeAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffSafeAreaType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.TouchDownLiftOffType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TouchDownLiftOffType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TouchDownLiftOffType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_TouchDownLiftOffType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VisualGlideSlopeIndicatorTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VisualGlideSlopeIndicatorTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VisualGlideSlopeIndicatorTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VisualGlideSlopeIndicatorTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VisualGlideSlopeIndicatorTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VisualGlideSlopeIndicatorTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VisualGlideSlopeIndicatorTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VisualGlideSlopeIndicatorTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VisualGlideSlopeIndicatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VisualGlideSlopeIndicatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VisualGlideSlopeIndicatorType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VisualGlideSlopeIndicatorType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.WorkAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "WorkAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_WorkAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_WorkAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.WorkAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "WorkAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_WorkAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_WorkAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.WorkAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "WorkAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_WorkAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_WorkAreaType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.WorkareaActivityPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "WorkareaActivityPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_WorkareaActivityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_WorkareaActivityPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.WorkareaActivityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "WorkareaActivityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_WorkareaActivityType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_WorkareaActivityType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftCategoryType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftCategoryType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftCategoryType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftCategoryType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftEngineNumberType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftEngineNumberType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftEngineNumberType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftEngineNumberType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftEngineType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftEngineType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftEngineType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftEngineType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftICAOType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftICAOType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftICAOType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftICAOType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftStandType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftStandType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftStandType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftStandType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAircraftWingspanClassType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAircraftWingspanClassType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAircraftWingspanClassType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAircraftWingspanClassType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirportHeliportCollocationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirportHeliportCollocationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirportHeliportCollocationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirportHeliportCollocationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirportHeliportDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirportHeliportDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirportHeliportDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirportHeliportDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirportHeliportType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirportHeliportType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirportHeliportType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirportHeliportType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirportWarningType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirportWarningType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirportWarningType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirportWarningType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeApproachGuidanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeApproachGuidanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeApproachGuidanceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeApproachGuidanceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeApproachLightingICAOType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeApproachLightingICAOType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeApproachLightingICAOType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeApproachLightingICAOType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeApproachLightingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeApproachLightingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeApproachLightingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeApproachLightingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeApronElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeApronElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeApronElementType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeApronElementType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeApronSectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeApronSectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeApronSectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeApronSectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeArrestingGearEnergyAbsorbType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeArrestingGearEnergyAbsorbType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeArrestingGearEnergyAbsorbType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeArrestingGearEnergyAbsorbType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeArrestingGearEngageDeviceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeArrestingGearEngageDeviceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeArrestingGearEngageDeviceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeArrestingGearEngageDeviceType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeBuoyDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeBuoyDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeBuoyDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeBuoyDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeBuoyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeBuoyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeBuoyType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeBuoyType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeColourType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeColourType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeColourType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeColourType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeContaminationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeContaminationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeContaminationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeContaminationType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeDeclaredDistanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDeclaredDistanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDeclaredDistanceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDeclaredDistanceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDirectionTurnType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDirectionTurnType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDirectionTurnType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDirectionTurnType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeDirectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeDirectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeDirectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeDirectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeEquipmentAntiCollisionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeEquipmentAntiCollisionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeEquipmentAntiCollisionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeEquipmentAntiCollisionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFlightOriginType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFlightOriginType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFlightOriginType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFlightOriginType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFlightPurposeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFlightPurposeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFlightPurposeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFlightPurposeType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeFlightStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFlightStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFlightStatusType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFlightStatusType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFlightType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFlightType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFlightType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFlightType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFrictionDeviceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFrictionDeviceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFrictionDeviceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFrictionDeviceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeFrictionEstimateType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeFrictionEstimateType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeFrictionEstimateType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeFrictionEstimateType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeGradeSeparationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeGradeSeparationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeGradeSeparationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeGradeSeparationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeGuidanceLineType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeGuidanceLineType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeGuidanceLineType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeGuidanceLineType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeHelicopterPerformanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeHelicopterPerformanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeHelicopterPerformanceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeHelicopterPerformanceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeHoldingCategoryType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeHoldingCategoryType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeHoldingCategoryType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeHoldingCategoryType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeIATAType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeIATAType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeIATAType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeIATAType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeICAOType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeICAOType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeICAOType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeICAOType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLightHoldingPositionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLightHoldingPositionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLightHoldingPositionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLightHoldingPositionType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeLightSourceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLightSourceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLightSourceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLightSourceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLightingJARType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLightingJARType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLightingJARType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLightingJARType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLoadingBridgeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLoadingBridgeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLoadingBridgeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLoadingBridgeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeLogicalOperatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeLogicalOperatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeLogicalOperatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeLogicalOperatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMarkingConditionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMarkingConditionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMarkingConditionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMarkingConditionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMarkingStyleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMarkingStyleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMarkingStyleType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMarkingStyleType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeMeteoConditionsType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMeteoConditionsType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMeteoConditionsType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMeteoConditionsType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeMilitaryStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeMilitaryStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeMilitaryStatusType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeMilitaryStatusType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNavigationEquipmentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNavigationEquipmentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNavigationEquipmentType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNavigationEquipmentType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeNavigationSpecificationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeNavigationSpecificationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeNavigationSpecificationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeNavigationSpecificationType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeOperationAirportHeliportType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOperationAirportHeliportType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOperationAirportHeliportType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOperationAirportHeliportType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeOperationManoeuvringAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeOperationManoeuvringAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeOperationManoeuvringAreaType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeOperationManoeuvringAreaType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePCNMethodType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePCNMethodType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePCNMethodType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePCNMethodType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePCNPavementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePCNPavementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePCNPavementType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePCNPavementType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePCNSubgradeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePCNSubgradeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePCNSubgradeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePCNSubgradeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodePCNTyrePressureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodePCNTyrePressureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodePCNTyrePressureType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodePCNTyrePressureType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeProtectAreaSectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeProtectAreaSectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeProtectAreaSectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeProtectAreaSectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRVRReadingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRVRReadingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRVRReadingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRVRReadingType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeRoadType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRoadType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRoadType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRoadType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRunwayElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRunwayElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRunwayElementType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRunwayElementType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRunwayMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRunwayMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRunwayMarkingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRunwayMarkingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRunwayPointRoleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRunwayPointRoleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRunwayPointRoleType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRunwayPointRoleType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRunwayProtectionAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRunwayProtectionAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRunwayProtectionAreaType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRunwayProtectionAreaType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRunwaySectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRunwaySectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRunwaySectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRunwaySectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeRunwayType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeRunwayType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeRunwayType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeRunwayType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSideType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSideType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSideType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSideType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeStatusAirportType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStatusAirportType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStatusAirportType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStatusAirportType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeSurfaceCompositionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSurfaceCompositionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSurfaceCompositionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSurfaceCompositionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSurfaceConditionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSurfaceConditionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSurfaceConditionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSurfaceConditionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeSurfacePreparationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeSurfacePreparationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeSurfacePreparationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeSurfacePreparationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTLOFSectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTLOFSectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTLOFSectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTLOFSectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTaxiwayElementType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTaxiwayElementType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTaxiwayElementType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTaxiwayElementType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTaxiwaySectionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTaxiwaySectionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTaxiwaySectionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTaxiwaySectionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeTaxiwayType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeTaxiwayType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeTaxiwayType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeTaxiwayType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeUsageLimitationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeUsageLimitationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeUsageLimitationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeUsageLimitationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeVASISType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVASISType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVASISType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVASISType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeValueInterpretationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeValueInterpretationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeValueInterpretationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeValueInterpretationType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeVisualDockingGuidanceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVisualDockingGuidanceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVisualDockingGuidanceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVisualDockingGuidanceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeWakeTurbulenceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeWakeTurbulenceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeWakeTurbulenceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeWakeTurbulenceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeWorkAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeWorkAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeWorkAreaType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeWorkAreaType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.DateTimeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DateTimeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DateTimeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_DateTimeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.DateType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DateType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DateType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_DateType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.ValDepthType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValDepthType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValDepthType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValDepthType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.ValFLType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValFLType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValFLType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValFLType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValFrictionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValFrictionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValFrictionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValFrictionType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValLCNType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValLCNType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValLCNType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValLCNType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValLightIntensityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValLightIntensityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValLightIntensityType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValLightIntensityType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValMagneticVariationChangeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValMagneticVariationChangeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValMagneticVariationChangeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValMagneticVariationChangeType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.ValPCNType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValPCNType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValPCNType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValPCNType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValPercentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValPercentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValPercentType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValPercentType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValPressureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValPressureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValPressureType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValPressureType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValSlopeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValSlopeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValSlopeType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValSlopeType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValSpeedType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValSpeedType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValSpeedType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValSpeedType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValTemperatureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValTemperatureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValTemperatureType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValTemperatureType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.ValWeightType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ValWeightType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ValWeightType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_ValWeightType_AixmdatatypesStage(probe)
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
	case *gml.AbstractCurveSegmentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractCurveSegmentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractCurveSegmentType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractCurveSegmentType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractCurveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractCurveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractCurveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractCurveType_GmlStage(probe)
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
	case *gml.AbstractRingPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractRingPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractRingPropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractRingPropertyType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractRingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractRingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractRingType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractRingType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractSurfacePatchType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractSurfacePatchType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractSurfacePatchType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractSurfacePatchType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.AbstractSurfaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AbstractSurfaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AbstractSurfaceType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AbstractSurfaceType_GmlStage(probe)
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
	case *gml.AngleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AngleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AngleType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_AngleType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.ArcByCenterPointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArcByCenterPointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArcByCenterPointType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_ArcByCenterPointType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.ArcStringType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArcStringType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArcStringType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_ArcStringType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.ArcType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArcType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArcType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_ArcType_GmlStage(probe)
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
	case *gml.CircleByCenterPointType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CircleByCenterPointType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CircleByCenterPointType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CircleByCenterPointType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.CircleType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CircleType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CircleType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CircleType_GmlStage(probe)
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
	case *gml.CompositeCurveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CompositeCurveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CompositeCurveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CompositeCurveType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.CurvePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CurvePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CurvePropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CurvePropertyType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.CurveSegmentArrayPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CurveSegmentArrayPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CurveSegmentArrayPropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CurveSegmentArrayPropertyType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.CurveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CurveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CurveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_CurveType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.DirectPositionListType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "DirectPositionListType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectPositionListType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_DirectPositionListType_GmlStage(probe)
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
	case *gml.GeodesicStringType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GeodesicStringType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GeodesicStringType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_GeodesicStringType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.GeodesicType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GeodesicType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GeodesicType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_GeodesicType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.LengthType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LengthType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LengthType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_LengthType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.LineStringSegmentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LineStringSegmentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LineStringSegmentType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_LineStringSegmentType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.MeasureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MeasureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeasureType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_MeasureType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.OrientableCurveType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "OrientableCurveType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OrientableCurveType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_OrientableCurveType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.PointPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PointPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PointPropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_PointPropertyType_GmlStage(probe)
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
	case *gml.PolygonPatchType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PolygonPatchType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PolygonPatchType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_PolygonPatchType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.RingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "RingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RingType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_RingType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.SurfacePatchArrayPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfacePatchArrayPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfacePatchArrayPropertyType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_SurfacePatchArrayPropertyType_GmlStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *gml.SurfaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.GmlStage.Lock()
				defer probe.stageOfInterest.GmlStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfaceType_GmlStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.GmlStage)
				}
				probe.stageOfInterest.GmlStage.Commit()
				updateStageSetTable_SurfaceType_GmlStage(probe)
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

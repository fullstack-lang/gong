// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/xlink"
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
	case *aixmfeatures.AirspaceActivationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceActivationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceActivationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceActivationPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceActivationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceActivationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceActivationType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceActivationType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceGeometryComponentPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceGeometryComponentPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceGeometryComponentType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceGeometryComponentType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceGeometryComponentType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceGeometryComponentType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceLayerClassPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceLayerClassPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceLayerClassPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceLayerClassPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceLayerClassType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceLayerClassType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceLayerClassType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceLayerClassType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceLayerPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceLayerPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceLayerPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceLayerPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceLayerType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceLayerType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceLayerType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceLayerType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.AirspaceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceVolumeDependencyPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceVolumeDependencyPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceVolumeDependencyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceVolumeDependencyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceVolumeDependencyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceVolumeDependencyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceVolumePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceVolumePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceVolumePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceVolumePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AirspaceVolumeType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AirspaceVolumeType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AirspaceVolumeType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AirspaceVolumeType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForAirspaceTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForAirspaceTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForAirspaceTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AuthorityForAirspaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AuthorityForAirspaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AuthorityForAirspaceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AuthorityForAirspaceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.CurvePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CurvePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CurvePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_CurvePropertyType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.GeoBorderTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GeoBorderTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GeoBorderTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GeoBorderTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GeoBorderTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GeoBorderTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.GeoBorderType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "GeoBorderType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GeoBorderType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_GeoBorderType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.SurfacePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SurfacePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SurfacePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SurfacePropertyType_AixmfeaturesStage(probe)
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
	case *aixmdatatypes.CodeAirspaceActivityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspaceActivityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspaceActivityType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspaceActivityType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspaceAggregationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspaceAggregationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspaceAggregationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspaceAggregationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspaceClassificationType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspaceClassificationType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspaceClassificationType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspaceClassificationType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspaceDependencyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspaceDependencyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspaceDependencyType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspaceDependencyType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspaceDesignatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspaceDesignatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspaceDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspaceDesignatorType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeAirspaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAirspaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAirspaceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAirspaceType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeAuthorityType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeAuthorityType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeAuthorityType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeAuthorityType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeGeoBorderType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeGeoBorderType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeGeoBorderType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeGeoBorderType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeStatusAirspaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStatusAirspaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStatusAirspaceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStatusAirspaceType_AixmdatatypesStage(probe)
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

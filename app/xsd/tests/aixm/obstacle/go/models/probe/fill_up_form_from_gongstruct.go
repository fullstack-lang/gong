// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
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
	case *aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AeronauticalGroundLightTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AeronauticalGroundLightTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AeronauticalGroundLightTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.AeronauticalGroundLightType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "AeronauticalGroundLightType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AeronauticalGroundLightType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_AeronauticalGroundLightType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.MarkerBeaconPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "MarkerBeaconPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MarkerBeaconPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_MarkerBeaconPropertyType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.ObstacleAreaTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ObstacleAreaTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ObstacleAreaTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ObstacleAreaTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ObstacleAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ObstacleAreaTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.ObstacleAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ObstacleAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ObstacleAreaType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ObstacleAreaType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.PassengerServicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "PassengerServicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PassengerServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_PassengerServicePropertyType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.ServicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ServicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_ServicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.SpecialNavigationStationPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "SpecialNavigationStationPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SpecialNavigationStationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_SpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
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
	case *aixmfeatures.VerticalStructureLightingStatusPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructureLightingStatusPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructureLightingStatusType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructureLightingStatusType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructureLightingStatusType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructureLightingStatusType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructurePartPropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructurePartPropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructurePartPropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructurePartPropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructurePartType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructurePartType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructurePartType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructurePartType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructurePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructurePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructurePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructurePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructureTimeSlicePropertyType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructureTimeSlicePropertyType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructureTimeSliceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructureTimeSliceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructureTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructureTimeSliceType_AixmfeaturesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmfeatures.VerticalStructureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "VerticalStructureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmfeaturesStage.Lock()
				defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VerticalStructureType_AixmfeaturesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				}
				probe.stageOfInterest.AixmfeaturesStage.Commit()
				updateStageSetTable_VerticalStructureType_AixmfeaturesStage(probe)
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
	case *aixmdatatypes.CodeGroundLightingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeGroundLightingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeGroundLightingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeGroundLightingType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeObstacleAreaType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeObstacleAreaType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeObstacleAreaType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeObstacleAreaType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeObstacleAssessmentSurfaceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeObstacleAssessmentSurfaceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeStatusConstructionType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeStatusConstructionType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeStatusConstructionType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeStatusConstructionType_AixmdatatypesStage(probe)
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
	case *aixmdatatypes.CodeVerticalStructureMarkingType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVerticalStructureMarkingType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVerticalStructureMarkingType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVerticalStructureMarkingType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeVerticalStructureMaterialType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVerticalStructureMaterialType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVerticalStructureMaterialType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVerticalStructureMaterialType_AixmdatatypesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *aixmdatatypes.CodeVerticalStructureType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "CodeVerticalStructureType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.AixmdatatypesStage.Lock()
				defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CodeVerticalStructureType_AixmdatatypesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				}
				probe.stageOfInterest.AixmdatatypesStage.Commit()
				updateStageSetTable_CodeVerticalStructureType_AixmdatatypesStage(probe)
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

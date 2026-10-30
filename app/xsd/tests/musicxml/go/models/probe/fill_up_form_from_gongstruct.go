// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xmlxsd"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)
}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {
	switch inst := instance.(type) {
	case *models.A_directive:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_directive",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_directive_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_directive_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.A_measure:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_measure",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_measure_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_measure_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.A_measure_1:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_measure_1",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_measure_1_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_measure_1_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.A_part:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_part",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_part_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_part_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.A_part_1:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "A_part_1",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_A_part_1_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_part_1_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Accidental:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Accidental",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accidental_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Accidental_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Accidental_mark:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Accidental_mark",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accidental_mark_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Accidental_mark_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Accidental_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Accidental_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accidental_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Accidental_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Accord:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Accord",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accord_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Accord_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Accordion_registration:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Accordion_registration",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accordion_registration_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Accordion_registration_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Appearance:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Appearance",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Appearance_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Appearance_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Arpeggiate:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Arpeggiate",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Arpeggiate_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Arpeggiate_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Arrow:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Arrow",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Arrow_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Arrow_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Articulations:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Articulations",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Articulations_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Articulations_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Assess:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Assess",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Assess_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Assess_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Attributes:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Attributes",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Attributes_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Attributes_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Backup:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Backup",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Backup_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Backup_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Bar_style_color:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Bar_style_color",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bar_style_color_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Bar_style_color_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Barline:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Barline",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Barline_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Barline_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Barre:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Barre",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Barre_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Barre_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Bass:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Bass",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bass_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Bass_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Bass_step:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Bass_step",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bass_step_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Bass_step_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Beam:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Beam",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Beam_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Beam_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Beat_repeat:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Beat_repeat",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Beat_repeat_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Beat_repeat_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Beat_unit_tied:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Beat_unit_tied",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Beat_unit_tied_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Beat_unit_tied_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Beater:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Beater",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Beater_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Beater_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Bend:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Bend",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bend_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Bend_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Bookmark:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Bookmark",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bookmark_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Bookmark_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Bracket:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Bracket",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bracket_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Bracket_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Breath_mark:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Breath_mark",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Breath_mark_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Breath_mark_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Caesura:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Caesura",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Caesura_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Caesura_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Cancel:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Cancel",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Cancel_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Cancel_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Clef:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Clef",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Clef_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Clef_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Coda:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Coda",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Coda_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Coda_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Credit:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Credit",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Credit_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Credit_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Dashes:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Dashes",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Dashes_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Dashes_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Defaults:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Defaults",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Defaults_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Defaults_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Degree:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Degree",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Degree_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Degree_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Degree_alter:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Degree_alter",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Degree_alter_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Degree_alter_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Degree_type:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Degree_type",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Degree_type_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Degree_type_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Degree_value:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Degree_value",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Degree_value_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Degree_value_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Direction:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Direction",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Direction_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Direction_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Direction_type:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Direction_type",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Direction_type_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Direction_type_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Distance:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Distance",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Distance_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Distance_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Double:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Double",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Double_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Double_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Dynamics:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Dynamics",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Dynamics_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Dynamics_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Effect:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Effect",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Effect_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Effect_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Elision:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Elision",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Elision_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Elision_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_font:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_font",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_font_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_font_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_line:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_line",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_line_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_line_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_placement:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_placement",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_placement_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_placement_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_placement_smufl:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_placement_smufl",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_placement_smufl_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_placement_smufl_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_print_object_style_align:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_print_object_style_align",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_print_object_style_align_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_print_object_style_align_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_print_style:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_print_style",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_print_style_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_print_style_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_print_style_align:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_print_style_align",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_print_style_align_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_print_style_align_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_print_style_align_id:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_print_style_align_id",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_print_style_align_id_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_print_style_align_id_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Empty_trill_sound:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Empty_trill_sound",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_trill_sound_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Empty_trill_sound_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Encoding:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Encoding",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Encoding_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Encoding_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Ending:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Ending",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Ending_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Ending_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Extend:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Extend",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Extend_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Extend_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Feature:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Feature",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Feature_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Feature_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Fermata:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Fermata",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Fermata_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Fermata_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Figure:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Figure",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Figure_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Figure_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Figured_bass:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Figured_bass",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Figured_bass_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Figured_bass_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Fingering:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Fingering",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Fingering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Fingering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.First_fret:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "First_fret",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_First_fret_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_First_fret_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.For_part:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "For_part",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_For_part_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_For_part_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Formatted_symbol:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Formatted_symbol",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Formatted_symbol_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Formatted_symbol_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Formatted_symbol_id:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Formatted_symbol_id",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Formatted_symbol_id_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Formatted_symbol_id_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Formatted_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Formatted_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Formatted_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Formatted_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Formatted_text_id:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Formatted_text_id",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Formatted_text_id_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Formatted_text_id_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Forward:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Forward",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Forward_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Forward_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Frame:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Frame",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Frame_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Frame_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Frame_note:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Frame_note",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Frame_note_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Frame_note_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Fret:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Fret",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Fret_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Fret_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Glass:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Glass",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Glass_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Glass_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Glissando:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Glissando",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Glissando_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Glissando_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Glyph:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Glyph",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Glyph_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Glyph_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Grace:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Grace",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Grace_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Grace_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Group_barline:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Group_barline",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Group_barline_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Group_barline_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Group_name:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Group_name",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Group_name_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Group_name_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Group_symbol:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Group_symbol",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Group_symbol_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Group_symbol_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Grouping:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Grouping",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Grouping_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Grouping_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Hammer_on_pull_off:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Hammer_on_pull_off",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Hammer_on_pull_off_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Hammer_on_pull_off_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Handbell:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Handbell",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Handbell_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Handbell_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Harmon_closed:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Harmon_closed",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmon_closed_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Harmon_closed_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Harmon_mute:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Harmon_mute",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmon_mute_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Harmon_mute_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Harmonic:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Harmonic",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmonic_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Harmonic_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Harmony:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Harmony",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmony_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Harmony_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Harmony_alter:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Harmony_alter",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmony_alter_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Harmony_alter_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Harp_pedals:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Harp_pedals",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harp_pedals_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Harp_pedals_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Heel_toe:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Heel_toe",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Heel_toe_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Heel_toe_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Hole:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Hole",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Hole_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Hole_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Hole_closed:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Hole_closed",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Hole_closed_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Hole_closed_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Horizontal_turn:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Horizontal_turn",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Horizontal_turn_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Horizontal_turn_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Identification:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Identification",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Identification_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Identification_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Image:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Image",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Image_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Image_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Instrument:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Instrument",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Instrument_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Instrument_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Instrument_change:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Instrument_change",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Instrument_change_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Instrument_change_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Instrument_link:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Instrument_link",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Instrument_link_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Instrument_link_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Interchangeable:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Interchangeable",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Interchangeable_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Interchangeable_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Inversion:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Inversion",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Inversion_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Inversion_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Key:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Key",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Key_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Key_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Key_accidental:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Key_accidental",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Key_accidental_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Key_accidental_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Key_octave:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Key_octave",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Key_octave_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Key_octave_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Kind:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Kind",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Kind_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Kind_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Level:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Level",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Level_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Level_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Line_detail:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Line_detail",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Line_detail_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Line_detail_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Line_width:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Line_width",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Line_width_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Line_width_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Link:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Link",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Link_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Link_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Listen:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Listen",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Listen_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Listen_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Listening:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Listening",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Listening_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Listening_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Lyric:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Lyric",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Lyric_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Lyric_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Lyric_font:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Lyric_font",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Lyric_font_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Lyric_font_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Lyric_language:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Lyric_language",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Lyric_language_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Lyric_language_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Measure_layout:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Measure_layout",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Measure_layout_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Measure_layout_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Measure_numbering:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Measure_numbering",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Measure_numbering_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Measure_numbering_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Measure_repeat:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Measure_repeat",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Measure_repeat_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Measure_repeat_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Measure_style:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Measure_style",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Measure_style_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Measure_style_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Membrane:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Membrane",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Membrane_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Membrane_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Metal:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Metal",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metal_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Metal_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Metronome:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Metronome",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metronome_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Metronome_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Metronome_beam:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Metronome_beam",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metronome_beam_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Metronome_beam_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Metronome_note:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Metronome_note",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metronome_note_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Metronome_note_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Metronome_tied:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Metronome_tied",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metronome_tied_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Metronome_tied_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Metronome_tuplet:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Metronome_tuplet",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metronome_tuplet_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Metronome_tuplet_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Midi_device:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Midi_device",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Midi_device_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Midi_device_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Midi_instrument:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Midi_instrument",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Midi_instrument_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Midi_instrument_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Miscellaneous:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Miscellaneous",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Miscellaneous_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Miscellaneous_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Miscellaneous_field:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Miscellaneous_field",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Miscellaneous_field_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Miscellaneous_field_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Mordent:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Mordent",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Mordent_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Mordent_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Multiple_rest:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Multiple_rest",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Multiple_rest_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Multiple_rest_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Name_display:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Name_display",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Name_display_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Name_display_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Non_arpeggiate:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Non_arpeggiate",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Non_arpeggiate_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Non_arpeggiate_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Notations:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Notations",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Notations_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Notations_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Note:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Note",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Note_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Note_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Note_size:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Note_size",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Note_size_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Note_size_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Note_type:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Note_type",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Note_type_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Note_type_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Notehead:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Notehead",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Notehead_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Notehead_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Notehead_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Notehead_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Notehead_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Notehead_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Numeral:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Numeral",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Numeral_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Numeral_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Numeral_key:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Numeral_key",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Numeral_key_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Numeral_key_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Numeral_root:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Numeral_root",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Numeral_root_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Numeral_root_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Octave_shift:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Octave_shift",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Octave_shift_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Octave_shift_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Offset:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Offset",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Offset_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Offset_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Opus:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Opus",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Opus_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Opus_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Ornaments:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Ornaments",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Ornaments_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Ornaments_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_appearance:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_appearance",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_appearance_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_appearance_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_direction:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_direction",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_direction_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_direction_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_listening:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_listening",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_listening_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_listening_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_notation:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_notation",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_notation_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_notation_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_placement_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_placement_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_placement_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_placement_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_play:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_play",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_play_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_play_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Other_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Other_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Other_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Other_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Page_layout:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Page_layout",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Page_layout_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Page_layout_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Page_margins:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Page_margins",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Page_margins_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Page_margins_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_clef:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_clef",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_clef_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_clef_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_group:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_group",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_group_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_group_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_link:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_link",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_link_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_link_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_list:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_list",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_list_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_list_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_name:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_name",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_name_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_name_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_symbol:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_symbol",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_symbol_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_symbol_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Part_transpose:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Part_transpose",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_transpose_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Part_transpose_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Pedal:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Pedal",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pedal_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Pedal_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Pedal_tuning:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Pedal_tuning",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pedal_tuning_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Pedal_tuning_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Per_minute:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Per_minute",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Per_minute_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Per_minute_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Percussion:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Percussion",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Percussion_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Percussion_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Pitch:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Pitch",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pitch_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Pitch_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Pitched:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Pitched",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pitched_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Pitched_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Placement_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Placement_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Placement_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Placement_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Play:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Play",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Play_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Play_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Player:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Player",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Player_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Player_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Principal_voice:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Principal_voice",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Principal_voice_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Principal_voice_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Print_XSD:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Print_XSD",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Print_XSD_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Print_XSD_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Release:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Release",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Release_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Release_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Repeat:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Repeat",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Repeat_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Repeat_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Rest:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Rest",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Rest_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Rest_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Root:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Root",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Root_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Root_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Root_step:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Root_step",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Root_step_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Root_step_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Scaling:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Scaling",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Scaling_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Scaling_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Scordatura:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Scordatura",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Scordatura_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Scordatura_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Score_instrument:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Score_instrument",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Score_instrument_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Score_instrument_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Score_part:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Score_part",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Score_part_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Score_part_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Score_partwise:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Score_partwise",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Score_partwise_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Score_partwise_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Score_timewise:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Score_timewise",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Score_timewise_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Score_timewise_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Segno:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Segno",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Segno_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Segno_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Slash:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Slash",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Slash_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Slash_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Slide:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Slide",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Slide_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Slide_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Slur:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Slur",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Slur_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Slur_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Sound:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Sound",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Sound_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Sound_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Staff_details:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Staff_details",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Staff_details_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Staff_details_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Staff_divide:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Staff_divide",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Staff_divide_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Staff_divide_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Staff_layout:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Staff_layout",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Staff_layout_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Staff_layout_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Staff_size:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Staff_size",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Staff_size_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Staff_size_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Staff_tuning:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Staff_tuning",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Staff_tuning_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Staff_tuning_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Stem:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Stem",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stem_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Stem_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Stick:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Stick",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stick_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Stick_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.String_mute:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "String_mute",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_String_mute_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_String_mute_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.String_type:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "String_type",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_String_type_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_String_type_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Strong_accent:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Strong_accent",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Strong_accent_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Strong_accent_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Style_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Style_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Style_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Style_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Supports:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Supports",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Supports_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Supports_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Swing:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Swing",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Swing_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Swing_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Sync:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Sync",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Sync_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Sync_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.System_dividers:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "System_dividers",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_System_dividers_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_System_dividers_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.System_layout:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "System_layout",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_System_layout_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_System_layout_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.System_margins:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "System_margins",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_System_margins_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_System_margins_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tap:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tap",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tap_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tap_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Technical:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Technical",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Technical_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Technical_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Text_element_data:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Text_element_data",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Text_element_data_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Text_element_data_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tie:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tie",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tie_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tie_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tied:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tied",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tied_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tied_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Time:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Time",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Time_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Time_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Time_modification:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Time_modification",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Time_modification_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Time_modification_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Timpani:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Timpani",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Timpani_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Timpani_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Transpose:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Transpose",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Transpose_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Transpose_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tremolo:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tremolo",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tremolo_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tremolo_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tuplet:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tuplet",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tuplet_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tuplet_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tuplet_dot:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tuplet_dot",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tuplet_dot_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tuplet_dot_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tuplet_number:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tuplet_number",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tuplet_number_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tuplet_number_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tuplet_portion:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tuplet_portion",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tuplet_portion_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tuplet_portion_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Tuplet_type:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Tuplet_type",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tuplet_type_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Tuplet_type_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Typed_text:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Typed_text",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Typed_text_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Typed_text_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Unpitched:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Unpitched",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Unpitched_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Unpitched_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Virtual_instrument:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Virtual_instrument",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Virtual_instrument_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Virtual_instrument_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Wait:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Wait",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wait_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Wait_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Wavy_line:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Wavy_line",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wavy_line_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Wavy_line_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Wedge:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Wedge",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wedge_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Wedge_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Wood:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Wood",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wood_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Wood_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *models.Work:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Work",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.Lock()
				defer probe.stageOfInterest.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Work_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest)
				}
				probe.stageOfInterest.Commit()
				updateStageSetTable_Work_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.ArcType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ArcType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArcType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_ArcType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.Extended:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Extended",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Extended_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_Extended_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.LocatorType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "LocatorType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LocatorType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_LocatorType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.ResourceType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "ResourceType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ResourceType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_ResourceType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.Simple:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Simple",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Simple_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_Simple_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xlink.TitleEltType:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "TitleEltType",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XlinkStage.Lock()
				defer probe.stageOfInterest.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TitleEltType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XlinkStage)
				}
				probe.stageOfInterest.XlinkStage.Commit()
				updateStageSetTable_TitleEltType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
	case *xmlxsd.Xmlxsd_SchemaElement:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "Xmlxsd_SchemaElement",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageOfInterest.XmlxsdStage.Lock()
				defer probe.stageOfInterest.XmlxsdStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Xmlxsd_SchemaElement_XmlxsdStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageOfInterest.XmlxsdStage)
				}
				probe.stageOfInterest.XmlxsdStage.Commit()
				updateStageSetTable_Xmlxsd_SchemaElement_XmlxsdStage(probe)
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

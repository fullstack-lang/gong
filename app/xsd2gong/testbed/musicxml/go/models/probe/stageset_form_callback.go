// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
)

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func StageSetFillUpFormFromGongstruct(
	instance any,
	probe *StageSetProbe,
) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name: "Form",
	}).Stage(probe.formStage)

	switch inst := instance.(type) {
	case *models.MusicXMLDocument:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MusicXMLDocument_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_MusicXMLDocument_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Credit:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Credit_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Credit_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Defaults:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Defaults_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Defaults_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.EmptyFont:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyFont_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_EmptyFont_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.GroupBarline:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroupBarline_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_GroupBarline_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.GroupName:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroupName_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_GroupName_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.GroupSymbol:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_GroupSymbol_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_GroupSymbol_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.InstrumentLink:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InstrumentLink_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_InstrumentLink_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.LyricFont:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LyricFont_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_LyricFont_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.LyricLanguage:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LyricLanguage_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_LyricLanguage_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Measure:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Measure_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Measure_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.MeasureMusicDataChoice:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeasureMusicDataChoice_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_MeasureMusicDataChoice_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Opus:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Opus_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Opus_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Part:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Part_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Part_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.PartGroup:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartGroup_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_PartGroup_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.PartLink:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartLink_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_PartLink_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.PartList:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartList_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_PartList_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.PartName:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartName_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_PartName_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Player:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Player_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Player_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.ScoreInstrument:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ScoreInstrument_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_ScoreInstrument_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.ScorePart:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ScorePart_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_ScorePart_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.ScorePartwise:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ScorePartwise_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_ScorePartwise_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.ScoreTimewise:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ScoreTimewise_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_ScoreTimewise_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.ScoreTimewiseMeasure:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ScoreTimewiseMeasure_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_ScoreTimewiseMeasure_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.ScoreTimewiseMeasurePart:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ScoreTimewiseMeasurePart_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_ScoreTimewiseMeasurePart_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *score.Work:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ScoreStage.Lock()
				defer probe.stageSet.ScoreStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Work_ScoreStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ScoreStage)
				}
				probe.stageSet.ScoreStage.Commit()
				updateStageSetTable_Work_ScoreStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Accidental:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accidental_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Accidental_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.AccidentalMark:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AccidentalMark_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_AccidentalMark_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Arpeggiate:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Arpeggiate_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Arpeggiate_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Arrow:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Arrow_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Arrow_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Articulations:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Articulations_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Articulations_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Assess:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Assess_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Assess_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Backup:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Backup_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Backup_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Beam:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Beam_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Beam_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Bend:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bend_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Bend_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.BreathMark:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BreathMark_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_BreathMark_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Caesura:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Caesura_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Caesura_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Elision:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Elision_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Elision_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.EmptyLine:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyLine_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_EmptyLine_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Extend:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Extend_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Extend_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Figure:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Figure_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Figure_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.FiguredBass:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FiguredBass_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_FiguredBass_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Forward:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Forward_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Forward_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Glissando:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Glissando_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Glissando_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Grace:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Grace_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Grace_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.HammerOnPullOff:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HammerOnPullOff_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_HammerOnPullOff_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Handbell:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Handbell_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Handbell_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.HarmonClosed:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HarmonClosed_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_HarmonClosed_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.HarmonMute:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HarmonMute_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_HarmonMute_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Harmonic:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmonic_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Harmonic_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.HeelToe:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HeelToe_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_HeelToe_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Hole:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Hole_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Hole_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.HoleClosed:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HoleClosed_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_HoleClosed_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Instrument:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Instrument_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Instrument_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Listen:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Listen_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Listen_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Lyric:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Lyric_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Lyric_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Mordent:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Mordent_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Mordent_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.NonArpeggiate:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NonArpeggiate_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_NonArpeggiate_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Notations:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Notations_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Notations_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Note:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Note_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Note_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.NoteType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteType_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_NoteType_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Notehead:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Notehead_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Notehead_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.NoteheadText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteheadText_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_NoteheadText_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Ornaments:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Ornaments_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Ornaments_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.OtherNotation:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherNotation_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_OtherNotation_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.OtherPlacementText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherPlacementText_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_OtherPlacementText_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Pitch:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pitch_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Pitch_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.PlacementText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PlacementText_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_PlacementText_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Release:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Release_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Release_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Rest:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Rest_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Rest_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Slide:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Slide_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Slide_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Slur:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Slur_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Slur_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Stem:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stem_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Stem_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.StrongAccent:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StrongAccent_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_StrongAccent_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Tap:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tap_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Tap_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Technical:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Technical_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Technical_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.TextElementData:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TextElementData_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_TextElementData_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Tie:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tie_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Tie_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Tied:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tied_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Tied_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Tremolo:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tremolo_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Tremolo_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Tuplet:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Tuplet_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Tuplet_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.TupletDot:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TupletDot_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_TupletDot_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.TupletNumber:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TupletNumber_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_TupletNumber_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.TupletPortion:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TupletPortion_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_TupletPortion_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.TupletType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TupletType_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_TupletType_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Unpitched:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Unpitched_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Unpitched_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *note.Wait:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.NoteStage.Lock()
				defer probe.stageSet.NoteStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wait_NoteStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.NoteStage)
				}
				probe.stageSet.NoteStage.Commit()
				updateStageSetTable_Wait_NoteStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Accord:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Accord_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Accord_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.AccordionRegistration:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AccordionRegistration_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_AccordionRegistration_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Barre:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Barre_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Barre_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Bass:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bass_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Bass_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.BassStep:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BassStep_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_BassStep_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.BeatUnitTied:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BeatUnitTied_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_BeatUnitTied_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Beater:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Beater_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Beater_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Bracket:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bracket_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Bracket_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Dashes:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Dashes_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Dashes_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Degree:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Degree_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Degree_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.DegreeAlter:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DegreeAlter_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_DegreeAlter_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.DegreeType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DegreeType_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_DegreeType_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.DegreeValue:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DegreeValue_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_DegreeValue_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Direction:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Direction_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Direction_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.DirectionType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DirectionType_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_DirectionType_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Effect:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Effect_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Effect_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Feature:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Feature_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Feature_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.FirstFret:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FirstFret_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_FirstFret_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Frame:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Frame_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Frame_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.FrameNote:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FrameNote_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_FrameNote_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Glass:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Glass_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Glass_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Grouping:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Grouping_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Grouping_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Harmony:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Harmony_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Harmony_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.HarmonyAlter:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HarmonyAlter_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_HarmonyAlter_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.HarpPedals:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HarpPedals_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_HarpPedals_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Image:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Image_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Image_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.InstrumentChange:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InstrumentChange_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_InstrumentChange_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Inversion:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Inversion_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Inversion_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Kind:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Kind_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Kind_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Listening:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Listening_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Listening_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.MeasureNumbering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeasureNumbering_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_MeasureNumbering_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Membrane:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Membrane_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Membrane_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Metal:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metal_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Metal_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Metronome:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Metronome_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Metronome_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.MetronomeBeam:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MetronomeBeam_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_MetronomeBeam_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.MetronomeNote:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MetronomeNote_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_MetronomeNote_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.MetronomeTied:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MetronomeTied_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_MetronomeTied_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.MetronomeTuplet:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MetronomeTuplet_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_MetronomeTuplet_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Numeral:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Numeral_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Numeral_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.NumeralKey:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NumeralKey_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_NumeralKey_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.NumeralRoot:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NumeralRoot_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_NumeralRoot_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.OctaveShift:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OctaveShift_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_OctaveShift_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Offset:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Offset_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Offset_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.OtherDirection:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherDirection_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_OtherDirection_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.OtherListening:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherListening_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_OtherListening_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Pedal:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pedal_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Pedal_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.PedalTuning:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PedalTuning_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_PedalTuning_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.PerMinute:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PerMinute_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_PerMinute_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Percussion:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Percussion_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Percussion_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Pitched:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Pitched_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Pitched_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.PrincipalVoice:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PrincipalVoice_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_PrincipalVoice_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Print:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Print_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Print_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Root:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Root_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Root_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.RootStep:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_RootStep_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_RootStep_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Scordatura:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Scordatura_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Scordatura_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Sound:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Sound_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Sound_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.StaffDivide:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaffDivide_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_StaffDivide_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Stick:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stick_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Stick_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.StringMute:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StringMute_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_StringMute_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.StyleText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StyleText_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_StyleText_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Swing:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Swing_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Swing_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Sync:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Sync_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Sync_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.TimeModification:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TimeModification_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_TimeModification_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Timpani:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Timpani_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Timpani_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.VirtualInstrument:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VirtualInstrument_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_VirtualInstrument_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Wedge:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wedge_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Wedge_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *direction.Wood:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.DirectionStage.Lock()
				defer probe.stageSet.DirectionStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Wood_DirectionStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.DirectionStage)
				}
				probe.stageSet.DirectionStage.Commit()
				updateStageSetTable_Wood_DirectionStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *barline.BarStyleColor:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.BarlineStage.Lock()
				defer probe.stageSet.BarlineStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BarStyleColor_BarlineStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.BarlineStage)
				}
				probe.stageSet.BarlineStage.Commit()
				updateStageSetTable_BarStyleColor_BarlineStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *barline.Barline:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.BarlineStage.Lock()
				defer probe.stageSet.BarlineStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Barline_BarlineStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.BarlineStage)
				}
				probe.stageSet.BarlineStage.Commit()
				updateStageSetTable_Barline_BarlineStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *barline.Ending:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.BarlineStage.Lock()
				defer probe.stageSet.BarlineStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Ending_BarlineStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.BarlineStage)
				}
				probe.stageSet.BarlineStage.Commit()
				updateStageSetTable_Ending_BarlineStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *barline.Repeat:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.BarlineStage.Lock()
				defer probe.stageSet.BarlineStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Repeat_BarlineStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.BarlineStage)
				}
				probe.stageSet.BarlineStage.Commit()
				updateStageSetTable_Repeat_BarlineStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Attributes:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Attributes_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Attributes_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.AttributesDirective:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AttributesDirective_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_AttributesDirective_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.BeatRepeat:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BeatRepeat_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_BeatRepeat_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Cancel:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Cancel_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Cancel_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Clef:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Clef_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Clef_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Double:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Double_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Double_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.ForPart:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ForPart_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_ForPart_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Interchangeable:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Interchangeable_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Interchangeable_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Key:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Key_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Key_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.KeyAccidental:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_KeyAccidental_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_KeyAccidental_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.KeyOctave:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_KeyOctave_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_KeyOctave_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.LineDetail:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LineDetail_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_LineDetail_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.MeasureRepeat:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeasureRepeat_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_MeasureRepeat_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.MeasureStyle:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeasureStyle_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_MeasureStyle_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.MultipleRest:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MultipleRest_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_MultipleRest_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.PartClef:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartClef_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_PartClef_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.PartSymbol:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartSymbol_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_PartSymbol_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.PartTranspose:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PartTranspose_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_PartTranspose_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Slash:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Slash_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Slash_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.StaffDetails:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaffDetails_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_StaffDetails_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.StaffSize:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaffSize_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_StaffSize_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.StaffTuning:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaffTuning_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_StaffTuning_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Time:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Time_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Time_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *attributes.Transpose:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.AttributesStage.Lock()
				defer probe.stageSet.AttributesStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Transpose_AttributesStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.AttributesStage)
				}
				probe.stageSet.AttributesStage.Commit()
				updateStageSetTable_Transpose_AttributesStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.AccidentalText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AccidentalText_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_AccidentalText_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Appearance:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Appearance_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Appearance_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Bookmark:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Bookmark_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Bookmark_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Coda:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Coda_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Coda_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Distance:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Distance_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Distance_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Dynamics:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Dynamics_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Dynamics_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Empty:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Empty_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Empty_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyPlacement:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyPlacement_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyPlacement_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyPlacementSmufl:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyPlacementSmufl_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyPlacementSmufl_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyPrintObjectStyleAlign:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyPrintObjectStyleAlign_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyPrintObjectStyleAlign_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyPrintStyle:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyPrintStyle_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyPrintStyle_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyPrintStyleAlign:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyPrintStyleAlign_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyPrintStyleAlign_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyPrintStyleAlignId:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyPrintStyleAlignId_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyPrintStyleAlignId_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.EmptyTrillSound:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_EmptyTrillSound_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_EmptyTrillSound_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Encoding:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Encoding_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Encoding_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Fermata:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Fermata_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Fermata_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Fingering:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Fingering_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Fingering_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.FormattedSymbol:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FormattedSymbol_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_FormattedSymbol_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.FormattedSymbolId:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FormattedSymbolId_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_FormattedSymbolId_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.FormattedText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FormattedText_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_FormattedText_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.FormattedTextId:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_FormattedTextId_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_FormattedTextId_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Fret:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Fret_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Fret_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Glyph:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Glyph_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Glyph_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.HorizontalTurn:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_HorizontalTurn_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_HorizontalTurn_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Identification:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Identification_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Identification_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Level:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Level_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Level_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.LineWidth:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LineWidth_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_LineWidth_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Link:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Link_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Link_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.MeasureLayout:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MeasureLayout_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_MeasureLayout_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.MidiDevice:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MidiDevice_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_MidiDevice_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.MidiInstrument:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MidiInstrument_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_MidiInstrument_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Miscellaneous:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Miscellaneous_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Miscellaneous_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.MiscellaneousField:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MiscellaneousField_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_MiscellaneousField_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.NameDisplay:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NameDisplay_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_NameDisplay_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.NoteSize:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteSize_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_NoteSize_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.OtherAppearance:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherAppearance_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_OtherAppearance_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.OtherPlay:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherPlay_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_OtherPlay_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.OtherText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OtherText_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_OtherText_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.PageLayout:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PageLayout_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_PageLayout_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.PageMargins:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PageMargins_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_PageMargins_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Play:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Play_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Play_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Scaling:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Scaling_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Scaling_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Segno:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Segno_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Segno_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.StaffLayout:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StaffLayout_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_StaffLayout_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.StringType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StringType_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_StringType_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.Supports:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Supports_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_Supports_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.SystemDividers:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SystemDividers_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_SystemDividers_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.SystemLayout:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SystemLayout_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_SystemLayout_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.SystemMargins:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SystemMargins_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_SystemMargins_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.TypedText:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TypedText_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_TypedText_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *common.WavyLine:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.CommonStage.Lock()
				defer probe.stageSet.CommonStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_WavyLine_CommonStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.CommonStage)
				}
				probe.stageSet.CommonStage.Commit()
				updateStageSetTable_WavyLine_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *xlink.ArcType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XlinkStage.Lock()
				defer probe.stageSet.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArcType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XlinkStage)
				}
				probe.stageSet.XlinkStage.Commit()
				updateStageSetTable_ArcType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *xlink.Extended:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XlinkStage.Lock()
				defer probe.stageSet.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Extended_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XlinkStage)
				}
				probe.stageSet.XlinkStage.Commit()
				updateStageSetTable_Extended_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *xlink.LocatorType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XlinkStage.Lock()
				defer probe.stageSet.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_LocatorType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XlinkStage)
				}
				probe.stageSet.XlinkStage.Commit()
				updateStageSetTable_LocatorType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *xlink.ResourceType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XlinkStage.Lock()
				defer probe.stageSet.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ResourceType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XlinkStage)
				}
				probe.stageSet.XlinkStage.Commit()
				updateStageSetTable_ResourceType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *xlink.Simple:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XlinkStage.Lock()
				defer probe.stageSet.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Simple_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XlinkStage)
				}
				probe.stageSet.XlinkStage.Commit()
				updateStageSetTable_Simple_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *xlink.TitleEltType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.XlinkStage.Lock()
				defer probe.stageSet.XlinkStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TitleEltType_XlinkStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.XlinkStage)
				}
				probe.stageSet.XlinkStage.Commit()
				updateStageSetTable_TitleEltType_XlinkStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	default:
		_ = inst
	}

	probe.formStage.Commit()
}

func saveStageSet_MusicXMLDocument_Stage(
	inst *models.MusicXMLDocument,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ScorePartwise":
			StageSetFormDivSelectFieldToField(&inst.ScorePartwise, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScorePartwise](), formDiv)
		case "ScoreTimewise":
			StageSetFormDivSelectFieldToField(&inst.ScoreTimewise, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScoreTimewise](), formDiv)
		}
	}
}

func saveStageSet_Credit_ScoreStage(
	inst *score.Credit,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Page":
			FormDivBasicFieldToField(&inst.Page, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "CreditType":
			FormDivBasicFieldToField(&inst.CreditType, formDiv)
		case "CreditImage":
			StageSetFormDivSelectFieldToField(&inst.CreditImage, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Image](), formDiv)
		}
	}
}

func saveStageSet_Defaults_ScoreStage(
	inst *score.Defaults,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Scaling":
			StageSetFormDivSelectFieldToField(&inst.Scaling, probe.stageSet.CommonStage.GetInstancesSet[*common.Scaling](), formDiv)
		case "ConcertScore":
			StageSetFormDivSelectFieldToField(&inst.ConcertScore, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "PageLayout":
			StageSetFormDivSelectFieldToField(&inst.PageLayout, probe.stageSet.CommonStage.GetInstancesSet[*common.PageLayout](), formDiv)
		case "SystemLayout":
			StageSetFormDivSelectFieldToField(&inst.SystemLayout, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemLayout](), formDiv)
		case "Appearance":
			StageSetFormDivSelectFieldToField(&inst.Appearance, probe.stageSet.CommonStage.GetInstancesSet[*common.Appearance](), formDiv)
		case "MusicFont":
			StageSetFormDivSelectFieldToField(&inst.MusicFont, probe.stageSet.ScoreStage.GetInstancesSet[*score.EmptyFont](), formDiv)
		case "WordFont":
			StageSetFormDivSelectFieldToField(&inst.WordFont, probe.stageSet.ScoreStage.GetInstancesSet[*score.EmptyFont](), formDiv)
		}
	}
}

func saveStageSet_EmptyFont_ScoreStage(
	inst *score.EmptyFont,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		}
	}
}

func saveStageSet_GroupBarline_ScoreStage(
	inst *score.GroupBarline,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_GroupName_ScoreStage(
	inst *score.GroupName,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_GroupSymbol_ScoreStage(
	inst *score.GroupSymbol,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_InstrumentLink_ScoreStage(
	inst *score.InstrumentLink,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RefScoreInstrument":
			StageSetFormDivSelectFieldToField(&inst.RefScoreInstrument, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScoreInstrument](), formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_LyricFont_ScoreStage(
	inst *score.LyricFont,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "NameAttr":
			FormDivBasicFieldToField(&inst.NameAttr, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		}
	}
}

func saveStageSet_LyricLanguage_ScoreStage(
	inst *score.LyricLanguage,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "NameAttr":
			FormDivBasicFieldToField(&inst.NameAttr, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		}
	}
}

func saveStageSet_Measure_ScoreStage(
	inst *score.Measure,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Implicit":
			StageSetFormDivEnumStringFieldToField(&inst.Implicit, formDiv)
		case "NonControlling":
			StageSetFormDivEnumStringFieldToField(&inst.NonControlling, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_MeasureMusicDataChoice_ScoreStage(
	inst *score.MeasureMusicDataChoice,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.NoteStage.GetInstancesSet[*note.Note](), formDiv)
		case "Backup":
			StageSetFormDivSelectFieldToField(&inst.Backup, probe.stageSet.NoteStage.GetInstancesSet[*note.Backup](), formDiv)
		case "Forward":
			StageSetFormDivSelectFieldToField(&inst.Forward, probe.stageSet.NoteStage.GetInstancesSet[*note.Forward](), formDiv)
		case "Direction":
			StageSetFormDivSelectFieldToField(&inst.Direction, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Direction](), formDiv)
		case "Attributes":
			StageSetFormDivSelectFieldToField(&inst.Attributes, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Attributes](), formDiv)
		case "Harmony":
			StageSetFormDivSelectFieldToField(&inst.Harmony, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Harmony](), formDiv)
		case "FiguredBass":
			StageSetFormDivSelectFieldToField(&inst.FiguredBass, probe.stageSet.NoteStage.GetInstancesSet[*note.FiguredBass](), formDiv)
		case "Print":
			StageSetFormDivSelectFieldToField(&inst.Print, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Print](), formDiv)
		case "Sound":
			StageSetFormDivSelectFieldToField(&inst.Sound, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Sound](), formDiv)
		case "Listening":
			StageSetFormDivSelectFieldToField(&inst.Listening, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Listening](), formDiv)
		case "Barline":
			StageSetFormDivSelectFieldToField(&inst.Barline, probe.stageSet.BarlineStage.GetInstancesSet[*barline.Barline](), formDiv)
		case "Grouping":
			StageSetFormDivSelectFieldToField(&inst.Grouping, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Grouping](), formDiv)
		case "Link":
			StageSetFormDivSelectFieldToField(&inst.Link, probe.stageSet.CommonStage.GetInstancesSet[*common.Link](), formDiv)
		case "Bookmark":
			StageSetFormDivSelectFieldToField(&inst.Bookmark, probe.stageSet.CommonStage.GetInstancesSet[*common.Bookmark](), formDiv)
		}
	}
}

func saveStageSet_Opus_ScoreStage(
	inst *score.Opus,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			StageSetFormDivEnumStringFieldToField(&inst.Show, formDiv)
		case "Actuate":
			StageSetFormDivEnumStringFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_Part_ScoreStage(
	inst *score.Part,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RefScorePart":
			StageSetFormDivSelectFieldToField(&inst.RefScorePart, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScorePart](), formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_PartGroup_ScoreStage(
	inst *score.PartGroup,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "GroupName":
			StageSetFormDivSelectFieldToField(&inst.GroupName, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupName](), formDiv)
		case "GroupNameDisplay":
			StageSetFormDivSelectFieldToField(&inst.GroupNameDisplay, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), formDiv)
		case "GroupAbbreviation":
			StageSetFormDivSelectFieldToField(&inst.GroupAbbreviation, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupName](), formDiv)
		case "GroupAbbreviationDisplay":
			StageSetFormDivSelectFieldToField(&inst.GroupAbbreviationDisplay, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), formDiv)
		case "GroupSymbol":
			StageSetFormDivSelectFieldToField(&inst.GroupSymbol, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupSymbol](), formDiv)
		case "GroupBarline":
			StageSetFormDivSelectFieldToField(&inst.GroupBarline, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupBarline](), formDiv)
		case "GroupTime":
			StageSetFormDivSelectFieldToField(&inst.GroupTime, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		}
	}
}

func saveStageSet_PartLink_ScoreStage(
	inst *score.PartLink,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			StageSetFormDivEnumStringFieldToField(&inst.Show, formDiv)
		case "Actuate":
			StageSetFormDivEnumStringFieldToField(&inst.Actuate, formDiv)
		case "GroupLink":
			FormDivBasicFieldToField(&inst.GroupLink, formDiv)
		}
	}
}

func saveStageSet_PartList_ScoreStage(
	inst *score.PartList,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_PartName_ScoreStage(
	inst *score.PartName,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Player_ScoreStage(
	inst *score.Player,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "PlayerName":
			FormDivBasicFieldToField(&inst.PlayerName, formDiv)
		}
	}
}

func saveStageSet_ScoreInstrument_ScoreStage(
	inst *score.ScoreInstrument,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "InstrumentName":
			FormDivBasicFieldToField(&inst.InstrumentName, formDiv)
		case "InstrumentAbbreviation":
			FormDivBasicFieldToField(&inst.InstrumentAbbreviation, formDiv)
		case "InstrumentSound":
			FormDivBasicFieldToField(&inst.InstrumentSound, formDiv)
		case "Solo":
			StageSetFormDivSelectFieldToField(&inst.Solo, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Ensemble":
			FormDivBasicFieldToField(&inst.Ensemble, formDiv)
		case "VirtualInstrument":
			StageSetFormDivSelectFieldToField(&inst.VirtualInstrument, probe.stageSet.DirectionStage.GetInstancesSet[*direction.VirtualInstrument](), formDiv)
		}
	}
}

func saveStageSet_ScorePart_ScoreStage(
	inst *score.ScorePart,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Identification":
			StageSetFormDivSelectFieldToField(&inst.Identification, probe.stageSet.CommonStage.GetInstancesSet[*common.Identification](), formDiv)
		case "PartName":
			StageSetFormDivSelectFieldToField(&inst.PartName, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartName](), formDiv)
		case "PartNameDisplay":
			StageSetFormDivSelectFieldToField(&inst.PartNameDisplay, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), formDiv)
		case "PartAbbreviation":
			StageSetFormDivSelectFieldToField(&inst.PartAbbreviation, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartName](), formDiv)
		case "PartAbbreviationDisplay":
			StageSetFormDivSelectFieldToField(&inst.PartAbbreviationDisplay, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), formDiv)
		case "Group":
			FormDivBasicFieldToField(&inst.Group, formDiv)
		}
	}
}

func saveStageSet_ScorePartwise_ScoreStage(
	inst *score.ScorePartwise,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Version":
			FormDivBasicFieldToField(&inst.Version, formDiv)
		case "Work":
			StageSetFormDivSelectFieldToField(&inst.Work, probe.stageSet.ScoreStage.GetInstancesSet[*score.Work](), formDiv)
		case "MovementNumber":
			FormDivBasicFieldToField(&inst.MovementNumber, formDiv)
		case "MovementTitle":
			FormDivBasicFieldToField(&inst.MovementTitle, formDiv)
		case "Identification":
			StageSetFormDivSelectFieldToField(&inst.Identification, probe.stageSet.CommonStage.GetInstancesSet[*common.Identification](), formDiv)
		case "Defaults":
			StageSetFormDivSelectFieldToField(&inst.Defaults, probe.stageSet.ScoreStage.GetInstancesSet[*score.Defaults](), formDiv)
		case "PartList":
			StageSetFormDivSelectFieldToField(&inst.PartList, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartList](), formDiv)
		}
	}
}

func saveStageSet_ScoreTimewise_ScoreStage(
	inst *score.ScoreTimewise,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Version":
			FormDivBasicFieldToField(&inst.Version, formDiv)
		case "Work":
			StageSetFormDivSelectFieldToField(&inst.Work, probe.stageSet.ScoreStage.GetInstancesSet[*score.Work](), formDiv)
		case "MovementNumber":
			FormDivBasicFieldToField(&inst.MovementNumber, formDiv)
		case "MovementTitle":
			FormDivBasicFieldToField(&inst.MovementTitle, formDiv)
		case "Identification":
			StageSetFormDivSelectFieldToField(&inst.Identification, probe.stageSet.CommonStage.GetInstancesSet[*common.Identification](), formDiv)
		case "Defaults":
			StageSetFormDivSelectFieldToField(&inst.Defaults, probe.stageSet.ScoreStage.GetInstancesSet[*score.Defaults](), formDiv)
		case "PartList":
			StageSetFormDivSelectFieldToField(&inst.PartList, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartList](), formDiv)
		}
	}
}

func saveStageSet_ScoreTimewiseMeasure_ScoreStage(
	inst *score.ScoreTimewiseMeasure,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Implicit":
			StageSetFormDivEnumStringFieldToField(&inst.Implicit, formDiv)
		case "NonControlling":
			StageSetFormDivEnumStringFieldToField(&inst.NonControlling, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_ScoreTimewiseMeasurePart_ScoreStage(
	inst *score.ScoreTimewiseMeasurePart,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RefScorePart":
			StageSetFormDivSelectFieldToField(&inst.RefScorePart, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScorePart](), formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Work_ScoreStage(
	inst *score.Work,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "WorkNumber":
			FormDivBasicFieldToField(&inst.WorkNumber, formDiv)
		case "WorkTitle":
			FormDivBasicFieldToField(&inst.WorkTitle, formDiv)
		case "Opus":
			StageSetFormDivSelectFieldToField(&inst.Opus, probe.stageSet.ScoreStage.GetInstancesSet[*score.Opus](), formDiv)
		}
	}
}

func saveStageSet_Accidental_NoteStage(
	inst *note.Accidental,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Cautionary":
			StageSetFormDivEnumStringFieldToField(&inst.Cautionary, formDiv)
		case "Editorial":
			StageSetFormDivEnumStringFieldToField(&inst.Editorial, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_AccidentalMark_NoteStage(
	inst *note.AccidentalMark,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Arpeggiate_NoteStage(
	inst *note.Arpeggiate,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Direction":
			StageSetFormDivEnumStringFieldToField(&inst.Direction, formDiv)
		case "Unbroken":
			StageSetFormDivEnumStringFieldToField(&inst.Unbroken, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Arrow_NoteStage(
	inst *note.Arrow,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "ArrowDirection":
			StageSetFormDivEnumStringFieldToField(&inst.ArrowDirection, formDiv)
		case "ArrowStyle":
			StageSetFormDivEnumStringFieldToField(&inst.ArrowStyle, formDiv)
		case "Arrowhead":
			StageSetFormDivSelectFieldToField(&inst.Arrowhead, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "CircularArrow":
			StageSetFormDivEnumStringFieldToField(&inst.CircularArrow, formDiv)
		}
	}
}

func saveStageSet_Articulations_NoteStage(
	inst *note.Articulations,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Assess_NoteStage(
	inst *note.Assess,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		}
	}
}

func saveStageSet_Backup_NoteStage(
	inst *note.Backup,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Duration":
			FormDivBasicFieldToField(&inst.Duration, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		}
	}
}

func saveStageSet_Beam_NoteStage(
	inst *note.Beam,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Repeater":
			StageSetFormDivEnumStringFieldToField(&inst.Repeater, formDiv)
		case "Fan":
			StageSetFormDivEnumStringFieldToField(&inst.Fan, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Bend_NoteStage(
	inst *note.Bend,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Shape":
			StageSetFormDivEnumStringFieldToField(&inst.Shape, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "FirstBeat":
			FormDivBasicFieldToField(&inst.FirstBeat, formDiv)
		case "LastBeat":
			FormDivBasicFieldToField(&inst.LastBeat, formDiv)
		case "BendAlter":
			FormDivBasicFieldToField(&inst.BendAlter, formDiv)
		case "PreBend":
			StageSetFormDivSelectFieldToField(&inst.PreBend, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Release":
			StageSetFormDivSelectFieldToField(&inst.Release, probe.stageSet.NoteStage.GetInstancesSet[*note.Release](), formDiv)
		case "WithBar":
			StageSetFormDivSelectFieldToField(&inst.WithBar, probe.stageSet.NoteStage.GetInstancesSet[*note.PlacementText](), formDiv)
		}
	}
}

func saveStageSet_BreathMark_NoteStage(
	inst *note.BreathMark,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Caesura_NoteStage(
	inst *note.Caesura,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Elision_NoteStage(
	inst *note.Elision,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_EmptyLine_NoteStage(
	inst *note.EmptyLine,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "LineShape":
			StageSetFormDivEnumStringFieldToField(&inst.LineShape, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "LineLength":
			StageSetFormDivEnumStringFieldToField(&inst.LineLength, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		}
	}
}

func saveStageSet_Extend_NoteStage(
	inst *note.Extend,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_Figure_NoteStage(
	inst *note.Figure,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Prefix":
			StageSetFormDivSelectFieldToField(&inst.Prefix, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), formDiv)
		case "FigureNumber":
			StageSetFormDivSelectFieldToField(&inst.FigureNumber, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), formDiv)
		case "Suffix":
			StageSetFormDivSelectFieldToField(&inst.Suffix, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), formDiv)
		case "Extend":
			StageSetFormDivSelectFieldToField(&inst.Extend, probe.stageSet.NoteStage.GetInstancesSet[*note.Extend](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		}
	}
}

func saveStageSet_FiguredBass_NoteStage(
	inst *note.FiguredBass,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "PrintDot":
			StageSetFormDivEnumStringFieldToField(&inst.PrintDot, formDiv)
		case "PrintLyric":
			StageSetFormDivEnumStringFieldToField(&inst.PrintLyric, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "PrintSpacing":
			StageSetFormDivEnumStringFieldToField(&inst.PrintSpacing, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Duration":
			FormDivBasicFieldToField(&inst.Duration, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		}
	}
}

func saveStageSet_Forward_NoteStage(
	inst *note.Forward,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Duration":
			FormDivBasicFieldToField(&inst.Duration, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		case "Voice":
			FormDivBasicFieldToField(&inst.Voice, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		}
	}
}

func saveStageSet_Glissando_NoteStage(
	inst *note.Glissando,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Grace_NoteStage(
	inst *note.Grace,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "StealTimePrevious":
			FormDivBasicFieldToField(&inst.StealTimePrevious, formDiv)
		case "StealTimeFollowing":
			FormDivBasicFieldToField(&inst.StealTimeFollowing, formDiv)
		case "MakeTime":
			FormDivBasicFieldToField(&inst.MakeTime, formDiv)
		case "Slash":
			StageSetFormDivEnumStringFieldToField(&inst.Slash, formDiv)
		}
	}
}

func saveStageSet_HammerOnPullOff_NoteStage(
	inst *note.HammerOnPullOff,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Handbell_NoteStage(
	inst *note.Handbell,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_HarmonClosed_NoteStage(
	inst *note.HarmonClosed,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_HarmonMute_NoteStage(
	inst *note.HarmonMute,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "HarmonClosed":
			StageSetFormDivSelectFieldToField(&inst.HarmonClosed, probe.stageSet.NoteStage.GetInstancesSet[*note.HarmonClosed](), formDiv)
		}
	}
}

func saveStageSet_Harmonic_NoteStage(
	inst *note.Harmonic,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Natural":
			StageSetFormDivSelectFieldToField(&inst.Natural, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Artificial":
			StageSetFormDivSelectFieldToField(&inst.Artificial, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "BasePitch":
			StageSetFormDivSelectFieldToField(&inst.BasePitch, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "TouchingPitch":
			StageSetFormDivSelectFieldToField(&inst.TouchingPitch, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "SoundingPitch":
			StageSetFormDivSelectFieldToField(&inst.SoundingPitch, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		}
	}
}

func saveStageSet_HeelToe_NoteStage(
	inst *note.HeelToe,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Substitution":
			StageSetFormDivEnumStringFieldToField(&inst.Substitution, formDiv)
		}
	}
}

func saveStageSet_Hole_NoteStage(
	inst *note.Hole,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "HoleType":
			FormDivBasicFieldToField(&inst.HoleType, formDiv)
		case "HoleClosed":
			StageSetFormDivSelectFieldToField(&inst.HoleClosed, probe.stageSet.NoteStage.GetInstancesSet[*note.HoleClosed](), formDiv)
		case "HoleShape":
			FormDivBasicFieldToField(&inst.HoleShape, formDiv)
		}
	}
}

func saveStageSet_HoleClosed_NoteStage(
	inst *note.HoleClosed,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Instrument_NoteStage(
	inst *note.Instrument,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Listen_NoteStage(
	inst *note.Listen,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_Lyric_NoteStage(
	inst *note.Lyric,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "NameAttr":
			FormDivBasicFieldToField(&inst.NameAttr, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Syllabic":
			StageSetFormDivEnumStringFieldToField(&inst.Syllabic, formDiv)
		case "Text":
			StageSetFormDivSelectFieldToField(&inst.Text, probe.stageSet.NoteStage.GetInstancesSet[*note.TextElementData](), formDiv)
		case "Extend":
			StageSetFormDivSelectFieldToField(&inst.Extend, probe.stageSet.NoteStage.GetInstancesSet[*note.Extend](), formDiv)
		case "Laughing":
			StageSetFormDivSelectFieldToField(&inst.Laughing, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Humming":
			StageSetFormDivSelectFieldToField(&inst.Humming, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "EndLine":
			StageSetFormDivSelectFieldToField(&inst.EndLine, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "EndParagraph":
			StageSetFormDivSelectFieldToField(&inst.EndParagraph, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		}
	}
}

func saveStageSet_Mordent_NoteStage(
	inst *note.Mordent,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "StartNote":
			StageSetFormDivEnumStringFieldToField(&inst.StartNote, formDiv)
		case "TrillStep":
			StageSetFormDivEnumStringFieldToField(&inst.TrillStep, formDiv)
		case "TwoNoteTurn":
			StageSetFormDivEnumStringFieldToField(&inst.TwoNoteTurn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "SecondBeat":
			FormDivBasicFieldToField(&inst.SecondBeat, formDiv)
		case "LastBeat":
			FormDivBasicFieldToField(&inst.LastBeat, formDiv)
		case "Long":
			StageSetFormDivEnumStringFieldToField(&inst.Long, formDiv)
		case "Approach":
			StageSetFormDivEnumStringFieldToField(&inst.Approach, formDiv)
		case "Departure":
			StageSetFormDivEnumStringFieldToField(&inst.Departure, formDiv)
		}
	}
}

func saveStageSet_NonArpeggiate_NoteStage(
	inst *note.NonArpeggiate,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Notations_NoteStage(
	inst *note.Notations,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		}
	}
}

func saveStageSet_Note_NoteStage(
	inst *note.Note,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintLeger":
			StageSetFormDivEnumStringFieldToField(&inst.PrintLeger, formDiv)
		case "Dynamics":
			FormDivBasicFieldToField(&inst.Dynamics, formDiv)
		case "EndDynamics":
			FormDivBasicFieldToField(&inst.EndDynamics, formDiv)
		case "Attack":
			FormDivBasicFieldToField(&inst.Attack, formDiv)
		case "Release":
			FormDivBasicFieldToField(&inst.Release, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		case "Pizzicato":
			StageSetFormDivEnumStringFieldToField(&inst.Pizzicato, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "PrintDot":
			StageSetFormDivEnumStringFieldToField(&inst.PrintDot, formDiv)
		case "PrintLyric":
			StageSetFormDivEnumStringFieldToField(&inst.PrintLyric, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "PrintSpacing":
			StageSetFormDivEnumStringFieldToField(&inst.PrintSpacing, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Grace":
			StageSetFormDivSelectFieldToField(&inst.Grace, probe.stageSet.NoteStage.GetInstancesSet[*note.Grace](), formDiv)
		case "Chord":
			StageSetFormDivSelectFieldToField(&inst.Chord, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Pitch":
			StageSetFormDivSelectFieldToField(&inst.Pitch, probe.stageSet.NoteStage.GetInstancesSet[*note.Pitch](), formDiv)
		case "Unpitched":
			StageSetFormDivSelectFieldToField(&inst.Unpitched, probe.stageSet.NoteStage.GetInstancesSet[*note.Unpitched](), formDiv)
		case "Rest":
			StageSetFormDivSelectFieldToField(&inst.Rest, probe.stageSet.NoteStage.GetInstancesSet[*note.Rest](), formDiv)
		case "Cue":
			StageSetFormDivSelectFieldToField(&inst.Cue, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Duration":
			FormDivBasicFieldToField(&inst.Duration, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		case "Voice":
			FormDivBasicFieldToField(&inst.Voice, formDiv)
		case "Type":
			StageSetFormDivSelectFieldToField(&inst.Type, probe.stageSet.NoteStage.GetInstancesSet[*note.NoteType](), formDiv)
		case "Accidental":
			StageSetFormDivSelectFieldToField(&inst.Accidental, probe.stageSet.NoteStage.GetInstancesSet[*note.Accidental](), formDiv)
		case "TimeModification":
			StageSetFormDivSelectFieldToField(&inst.TimeModification, probe.stageSet.DirectionStage.GetInstancesSet[*direction.TimeModification](), formDiv)
		case "Stem":
			StageSetFormDivSelectFieldToField(&inst.Stem, probe.stageSet.NoteStage.GetInstancesSet[*note.Stem](), formDiv)
		case "Notehead":
			StageSetFormDivSelectFieldToField(&inst.Notehead, probe.stageSet.NoteStage.GetInstancesSet[*note.Notehead](), formDiv)
		case "NoteheadText":
			StageSetFormDivSelectFieldToField(&inst.NoteheadText, probe.stageSet.NoteStage.GetInstancesSet[*note.NoteheadText](), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		case "Play":
			StageSetFormDivSelectFieldToField(&inst.Play, probe.stageSet.CommonStage.GetInstancesSet[*common.Play](), formDiv)
		case "Listen":
			StageSetFormDivSelectFieldToField(&inst.Listen, probe.stageSet.NoteStage.GetInstancesSet[*note.Listen](), formDiv)
		}
	}
}

func saveStageSet_NoteType_NoteStage(
	inst *note.NoteType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Notehead_NoteStage(
	inst *note.Notehead,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Filled":
			StageSetFormDivEnumStringFieldToField(&inst.Filled, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_NoteheadText_NoteStage(
	inst *note.NoteheadText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_Ornaments_NoteStage(
	inst *note.Ornaments,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_OtherNotation_NoteStage(
	inst *note.OtherNotation,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OtherPlacementText_NoteStage(
	inst *note.OtherPlacementText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Pitch_NoteStage(
	inst *note.Pitch,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Step":
			StageSetFormDivEnumStringFieldToField(&inst.Step, formDiv)
		case "Alter":
			FormDivBasicFieldToField(&inst.Alter, formDiv)
		case "Octave":
			FormDivBasicFieldToField(&inst.Octave, formDiv)
		}
	}
}

func saveStageSet_PlacementText_NoteStage(
	inst *note.PlacementText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Release_NoteStage(
	inst *note.Release,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Offset":
			FormDivBasicFieldToField(&inst.Offset, formDiv)
		}
	}
}

func saveStageSet_Rest_NoteStage(
	inst *note.Rest,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Measure":
			StageSetFormDivEnumStringFieldToField(&inst.Measure, formDiv)
		case "DisplayStep":
			StageSetFormDivEnumStringFieldToField(&inst.DisplayStep, formDiv)
		case "DisplayOctave":
			FormDivBasicFieldToField(&inst.DisplayOctave, formDiv)
		}
	}
}

func saveStageSet_Slide_NoteStage(
	inst *note.Slide,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "FirstBeat":
			FormDivBasicFieldToField(&inst.FirstBeat, formDiv)
		case "LastBeat":
			FormDivBasicFieldToField(&inst.LastBeat, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Slur_NoteStage(
	inst *note.Slur,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Orientation":
			StageSetFormDivEnumStringFieldToField(&inst.Orientation, formDiv)
		case "BezierX":
			FormDivBasicFieldToField(&inst.BezierX, formDiv)
		case "BezierY":
			FormDivBasicFieldToField(&inst.BezierY, formDiv)
		case "BezierX2":
			FormDivBasicFieldToField(&inst.BezierX2, formDiv)
		case "BezierY2":
			FormDivBasicFieldToField(&inst.BezierY2, formDiv)
		case "BezierOffset":
			FormDivBasicFieldToField(&inst.BezierOffset, formDiv)
		case "BezierOffset2":
			FormDivBasicFieldToField(&inst.BezierOffset2, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Stem_NoteStage(
	inst *note.Stem,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_StrongAccent_NoteStage(
	inst *note.StrongAccent,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		}
	}
}

func saveStageSet_Tap_NoteStage(
	inst *note.Tap,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Hand":
			StageSetFormDivEnumStringFieldToField(&inst.Hand, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Technical_NoteStage(
	inst *note.Technical,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_TextElementData_NoteStage(
	inst *note.TextElementData,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Tie_NoteStage(
	inst *note.Tie,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		}
	}
}

func saveStageSet_Tied_NoteStage(
	inst *note.Tied,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Orientation":
			StageSetFormDivEnumStringFieldToField(&inst.Orientation, formDiv)
		case "BezierX":
			FormDivBasicFieldToField(&inst.BezierX, formDiv)
		case "BezierY":
			FormDivBasicFieldToField(&inst.BezierY, formDiv)
		case "BezierX2":
			FormDivBasicFieldToField(&inst.BezierX2, formDiv)
		case "BezierY2":
			FormDivBasicFieldToField(&inst.BezierY2, formDiv)
		case "BezierOffset":
			FormDivBasicFieldToField(&inst.BezierOffset, formDiv)
		case "BezierOffset2":
			FormDivBasicFieldToField(&inst.BezierOffset2, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Tremolo_NoteStage(
	inst *note.Tremolo,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Tuplet_NoteStage(
	inst *note.Tuplet,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "ShowNumber":
			StageSetFormDivEnumStringFieldToField(&inst.ShowNumber, formDiv)
		case "ShowType":
			StageSetFormDivEnumStringFieldToField(&inst.ShowType, formDiv)
		case "LineShape":
			StageSetFormDivEnumStringFieldToField(&inst.LineShape, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "TupletActual":
			StageSetFormDivSelectFieldToField(&inst.TupletActual, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletPortion](), formDiv)
		case "TupletNormal":
			StageSetFormDivSelectFieldToField(&inst.TupletNormal, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletPortion](), formDiv)
		}
	}
}

func saveStageSet_TupletDot_NoteStage(
	inst *note.TupletDot,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_TupletNumber_NoteStage(
	inst *note.TupletNumber,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_TupletPortion_NoteStage(
	inst *note.TupletPortion,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "TupletNumber":
			StageSetFormDivSelectFieldToField(&inst.TupletNumber, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletNumber](), formDiv)
		case "TupletType":
			StageSetFormDivSelectFieldToField(&inst.TupletType, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletType](), formDiv)
		}
	}
}

func saveStageSet_TupletType_NoteStage(
	inst *note.TupletType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Unpitched_NoteStage(
	inst *note.Unpitched,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DisplayStep":
			StageSetFormDivEnumStringFieldToField(&inst.DisplayStep, formDiv)
		case "DisplayOctave":
			FormDivBasicFieldToField(&inst.DisplayOctave, formDiv)
		}
	}
}

func saveStageSet_Wait_NoteStage(
	inst *note.Wait,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		}
	}
}

func saveStageSet_Accord_DirectionStage(
	inst *direction.Accord,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "String":
			FormDivBasicFieldToField(&inst.String, formDiv)
		case "TuningStep":
			StageSetFormDivEnumStringFieldToField(&inst.TuningStep, formDiv)
		case "TuningAlter":
			FormDivBasicFieldToField(&inst.TuningAlter, formDiv)
		case "TuningOctave":
			FormDivBasicFieldToField(&inst.TuningOctave, formDiv)
		}
	}
}

func saveStageSet_AccordionRegistration_DirectionStage(
	inst *direction.AccordionRegistration,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "AccordionHigh":
			StageSetFormDivSelectFieldToField(&inst.AccordionHigh, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "AccordionMiddle":
			FormDivBasicFieldToField(&inst.AccordionMiddle, formDiv)
		case "AccordionLow":
			StageSetFormDivSelectFieldToField(&inst.AccordionLow, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		}
	}
}

func saveStageSet_Barre_DirectionStage(
	inst *direction.Barre,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_Bass_DirectionStage(
	inst *direction.Bass,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Arrangement":
			StageSetFormDivEnumStringFieldToField(&inst.Arrangement, formDiv)
		case "BassSeparator":
			StageSetFormDivSelectFieldToField(&inst.BassSeparator, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), formDiv)
		case "BassStep":
			StageSetFormDivSelectFieldToField(&inst.BassStep, probe.stageSet.DirectionStage.GetInstancesSet[*direction.BassStep](), formDiv)
		case "BassAlter":
			StageSetFormDivSelectFieldToField(&inst.BassAlter, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarmonyAlter](), formDiv)
		}
	}
}

func saveStageSet_BassStep_DirectionStage(
	inst *direction.BassStep,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_BeatUnitTied_DirectionStage(
	inst *direction.BeatUnitTied,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "BeatUnit":
			StageSetFormDivEnumStringFieldToField(&inst.BeatUnit, formDiv)
		}
	}
}

func saveStageSet_Beater_DirectionStage(
	inst *direction.Beater,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tip":
			StageSetFormDivEnumStringFieldToField(&inst.Tip, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Bracket_DirectionStage(
	inst *direction.Bracket,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "LineEnd":
			StageSetFormDivEnumStringFieldToField(&inst.LineEnd, formDiv)
		case "EndLength":
			FormDivBasicFieldToField(&inst.EndLength, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Dashes_DirectionStage(
	inst *direction.Dashes,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Degree_DirectionStage(
	inst *direction.Degree,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DegreeValue":
			StageSetFormDivSelectFieldToField(&inst.DegreeValue, probe.stageSet.DirectionStage.GetInstancesSet[*direction.DegreeValue](), formDiv)
		case "DegreeAlter":
			StageSetFormDivSelectFieldToField(&inst.DegreeAlter, probe.stageSet.DirectionStage.GetInstancesSet[*direction.DegreeAlter](), formDiv)
		case "DegreeType":
			StageSetFormDivSelectFieldToField(&inst.DegreeType, probe.stageSet.DirectionStage.GetInstancesSet[*direction.DegreeType](), formDiv)
		}
	}
}

func saveStageSet_DegreeAlter_DirectionStage(
	inst *direction.DegreeAlter,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PlusMinus":
			StageSetFormDivEnumStringFieldToField(&inst.PlusMinus, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_DegreeType_DirectionStage(
	inst *direction.DegreeType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_DegreeValue_DirectionStage(
	inst *direction.DegreeValue,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Direction_DirectionStage(
	inst *direction.Direction,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Directive":
			StageSetFormDivEnumStringFieldToField(&inst.Directive, formDiv)
		case "System":
			StageSetFormDivEnumStringFieldToField(&inst.System, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		case "Voice":
			FormDivBasicFieldToField(&inst.Voice, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		case "Sound":
			StageSetFormDivSelectFieldToField(&inst.Sound, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Sound](), formDiv)
		case "Listening":
			StageSetFormDivSelectFieldToField(&inst.Listening, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Listening](), formDiv)
		}
	}
}

func saveStageSet_DirectionType_DirectionStage(
	inst *direction.DirectionType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Wedge":
			StageSetFormDivSelectFieldToField(&inst.Wedge, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Wedge](), formDiv)
		case "Dashes":
			StageSetFormDivSelectFieldToField(&inst.Dashes, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Dashes](), formDiv)
		case "Bracket":
			StageSetFormDivSelectFieldToField(&inst.Bracket, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Bracket](), formDiv)
		case "Pedal":
			StageSetFormDivSelectFieldToField(&inst.Pedal, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Pedal](), formDiv)
		case "Metronome":
			StageSetFormDivSelectFieldToField(&inst.Metronome, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Metronome](), formDiv)
		case "OctaveShift":
			StageSetFormDivSelectFieldToField(&inst.OctaveShift, probe.stageSet.DirectionStage.GetInstancesSet[*direction.OctaveShift](), formDiv)
		case "HarpPedals":
			StageSetFormDivSelectFieldToField(&inst.HarpPedals, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarpPedals](), formDiv)
		case "Damp":
			StageSetFormDivSelectFieldToField(&inst.Damp, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintStyleAlignId](), formDiv)
		case "DampAll":
			StageSetFormDivSelectFieldToField(&inst.DampAll, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintStyleAlignId](), formDiv)
		case "Eyeglasses":
			StageSetFormDivSelectFieldToField(&inst.Eyeglasses, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintStyleAlignId](), formDiv)
		case "StringMute":
			StageSetFormDivSelectFieldToField(&inst.StringMute, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StringMute](), formDiv)
		case "Scordatura":
			StageSetFormDivSelectFieldToField(&inst.Scordatura, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Scordatura](), formDiv)
		case "Image":
			StageSetFormDivSelectFieldToField(&inst.Image, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Image](), formDiv)
		case "PrincipalVoice":
			StageSetFormDivSelectFieldToField(&inst.PrincipalVoice, probe.stageSet.DirectionStage.GetInstancesSet[*direction.PrincipalVoice](), formDiv)
		case "AccordionRegistration":
			StageSetFormDivSelectFieldToField(&inst.AccordionRegistration, probe.stageSet.DirectionStage.GetInstancesSet[*direction.AccordionRegistration](), formDiv)
		case "StaffDivide":
			StageSetFormDivSelectFieldToField(&inst.StaffDivide, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StaffDivide](), formDiv)
		case "OtherDirection":
			StageSetFormDivSelectFieldToField(&inst.OtherDirection, probe.stageSet.DirectionStage.GetInstancesSet[*direction.OtherDirection](), formDiv)
		}
	}
}

func saveStageSet_Effect_DirectionStage(
	inst *direction.Effect,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Feature_DirectionStage(
	inst *direction.Feature,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_FirstFret_DirectionStage(
	inst *direction.FirstFret,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Frame_DirectionStage(
	inst *direction.Frame,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Unplayed":
			FormDivBasicFieldToField(&inst.Unplayed, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "FrameStrings":
			FormDivBasicFieldToField(&inst.FrameStrings, formDiv)
		case "FrameFrets":
			FormDivBasicFieldToField(&inst.FrameFrets, formDiv)
		case "FirstFret":
			StageSetFormDivSelectFieldToField(&inst.FirstFret, probe.stageSet.DirectionStage.GetInstancesSet[*direction.FirstFret](), formDiv)
		}
	}
}

func saveStageSet_FrameNote_DirectionStage(
	inst *direction.FrameNote,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "String":
			StageSetFormDivSelectFieldToField(&inst.String, probe.stageSet.CommonStage.GetInstancesSet[*common.StringType](), formDiv)
		case "Fret":
			StageSetFormDivSelectFieldToField(&inst.Fret, probe.stageSet.CommonStage.GetInstancesSet[*common.Fret](), formDiv)
		case "Fingering":
			StageSetFormDivSelectFieldToField(&inst.Fingering, probe.stageSet.CommonStage.GetInstancesSet[*common.Fingering](), formDiv)
		case "Barre":
			StageSetFormDivSelectFieldToField(&inst.Barre, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Barre](), formDiv)
		}
	}
}

func saveStageSet_Glass_DirectionStage(
	inst *direction.Glass,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Grouping_DirectionStage(
	inst *direction.Grouping,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "MemberOf":
			FormDivBasicFieldToField(&inst.MemberOf, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Harmony_DirectionStage(
	inst *direction.Harmony,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "PrintFrame":
			StageSetFormDivEnumStringFieldToField(&inst.PrintFrame, formDiv)
		case "Arrangement":
			StageSetFormDivEnumStringFieldToField(&inst.Arrangement, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "System":
			StageSetFormDivEnumStringFieldToField(&inst.System, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Frame":
			StageSetFormDivSelectFieldToField(&inst.Frame, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Frame](), formDiv)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		}
	}
}

func saveStageSet_HarmonyAlter_DirectionStage(
	inst *direction.HarmonyAlter,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_HarpPedals_DirectionStage(
	inst *direction.HarpPedals,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Image_DirectionStage(
	inst *direction.Image,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Source":
			FormDivBasicFieldToField(&inst.Source, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_InstrumentChange_DirectionStage(
	inst *direction.InstrumentChange,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "InstrumentSound":
			FormDivBasicFieldToField(&inst.InstrumentSound, formDiv)
		case "Solo":
			StageSetFormDivSelectFieldToField(&inst.Solo, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "Ensemble":
			FormDivBasicFieldToField(&inst.Ensemble, formDiv)
		case "VirtualInstrument":
			StageSetFormDivSelectFieldToField(&inst.VirtualInstrument, probe.stageSet.DirectionStage.GetInstancesSet[*direction.VirtualInstrument](), formDiv)
		}
	}
}

func saveStageSet_Inversion_DirectionStage(
	inst *direction.Inversion,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Kind_DirectionStage(
	inst *direction.Kind,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "UseSymbols":
			StageSetFormDivEnumStringFieldToField(&inst.UseSymbols, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "StackDegrees":
			StageSetFormDivEnumStringFieldToField(&inst.StackDegrees, formDiv)
		case "ParenthesesDegrees":
			StageSetFormDivEnumStringFieldToField(&inst.ParenthesesDegrees, formDiv)
		case "BracketDegrees":
			StageSetFormDivEnumStringFieldToField(&inst.BracketDegrees, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Listening_DirectionStage(
	inst *direction.Listening,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), formDiv)
		}
	}
}

func saveStageSet_MeasureNumbering_DirectionStage(
	inst *direction.MeasureNumbering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "System":
			StageSetFormDivEnumStringFieldToField(&inst.System, formDiv)
		case "Staff":
			FormDivBasicFieldToField(&inst.Staff, formDiv)
		case "MultipleRestAlways":
			StageSetFormDivEnumStringFieldToField(&inst.MultipleRestAlways, formDiv)
		case "MultipleRestRange":
			StageSetFormDivEnumStringFieldToField(&inst.MultipleRestRange, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Membrane_DirectionStage(
	inst *direction.Membrane,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Metal_DirectionStage(
	inst *direction.Metal,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Metronome_DirectionStage(
	inst *direction.Metronome,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "PerMinute":
			StageSetFormDivSelectFieldToField(&inst.PerMinute, probe.stageSet.DirectionStage.GetInstancesSet[*direction.PerMinute](), formDiv)
		case "MetronomeArrows":
			StageSetFormDivSelectFieldToField(&inst.MetronomeArrows, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "MetronomeRelation":
			FormDivBasicFieldToField(&inst.MetronomeRelation, formDiv)
		}
	}
}

func saveStageSet_MetronomeBeam_DirectionStage(
	inst *direction.MetronomeBeam,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_MetronomeNote_DirectionStage(
	inst *direction.MetronomeNote,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "MetronomeType":
			StageSetFormDivEnumStringFieldToField(&inst.MetronomeType, formDiv)
		case "MetronomeTied":
			StageSetFormDivSelectFieldToField(&inst.MetronomeTied, probe.stageSet.DirectionStage.GetInstancesSet[*direction.MetronomeTied](), formDiv)
		case "MetronomeTuplet":
			StageSetFormDivSelectFieldToField(&inst.MetronomeTuplet, probe.stageSet.DirectionStage.GetInstancesSet[*direction.MetronomeTuplet](), formDiv)
		}
	}
}

func saveStageSet_MetronomeTied_DirectionStage(
	inst *direction.MetronomeTied,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		}
	}
}

func saveStageSet_MetronomeTuplet_DirectionStage(
	inst *direction.MetronomeTuplet,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "ShowNumber":
			StageSetFormDivEnumStringFieldToField(&inst.ShowNumber, formDiv)
		case "ActualNotes":
			FormDivBasicFieldToField(&inst.ActualNotes, formDiv)
		case "NormalNotes":
			FormDivBasicFieldToField(&inst.NormalNotes, formDiv)
		case "NormalType":
			StageSetFormDivEnumStringFieldToField(&inst.NormalType, formDiv)
		}
	}
}

func saveStageSet_Numeral_DirectionStage(
	inst *direction.Numeral,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumeralRoot":
			StageSetFormDivSelectFieldToField(&inst.NumeralRoot, probe.stageSet.DirectionStage.GetInstancesSet[*direction.NumeralRoot](), formDiv)
		case "NumeralAlter":
			StageSetFormDivSelectFieldToField(&inst.NumeralAlter, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarmonyAlter](), formDiv)
		case "NumeralKey":
			StageSetFormDivSelectFieldToField(&inst.NumeralKey, probe.stageSet.DirectionStage.GetInstancesSet[*direction.NumeralKey](), formDiv)
		}
	}
}

func saveStageSet_NumeralKey_DirectionStage(
	inst *direction.NumeralKey,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "NumeralFifths":
			FormDivBasicFieldToField(&inst.NumeralFifths, formDiv)
		case "NumeralMode":
			StageSetFormDivEnumStringFieldToField(&inst.NumeralMode, formDiv)
		}
	}
}

func saveStageSet_NumeralRoot_DirectionStage(
	inst *direction.NumeralRoot,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OctaveShift_DirectionStage(
	inst *direction.OctaveShift,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Size":
			FormDivBasicFieldToField(&inst.Size, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Offset_DirectionStage(
	inst *direction.Offset,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Sound":
			StageSetFormDivEnumStringFieldToField(&inst.Sound, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OtherDirection_DirectionStage(
	inst *direction.OtherDirection,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OtherListening_DirectionStage(
	inst *direction.OtherListening,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Pedal_DirectionStage(
	inst *direction.Pedal,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Line":
			StageSetFormDivEnumStringFieldToField(&inst.Line, formDiv)
		case "Sign":
			StageSetFormDivEnumStringFieldToField(&inst.Sign, formDiv)
		case "Abbreviated":
			StageSetFormDivEnumStringFieldToField(&inst.Abbreviated, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_PedalTuning_DirectionStage(
	inst *direction.PedalTuning,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PedalStep":
			StageSetFormDivEnumStringFieldToField(&inst.PedalStep, formDiv)
		case "PedalAlter":
			FormDivBasicFieldToField(&inst.PedalAlter, formDiv)
		}
	}
}

func saveStageSet_PerMinute_DirectionStage(
	inst *direction.PerMinute,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Percussion_DirectionStage(
	inst *direction.Percussion,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Glass":
			StageSetFormDivSelectFieldToField(&inst.Glass, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Glass](), formDiv)
		case "Metal":
			StageSetFormDivSelectFieldToField(&inst.Metal, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Metal](), formDiv)
		case "Wood":
			StageSetFormDivSelectFieldToField(&inst.Wood, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Wood](), formDiv)
		case "Pitched":
			StageSetFormDivSelectFieldToField(&inst.Pitched, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Pitched](), formDiv)
		case "Membrane":
			StageSetFormDivSelectFieldToField(&inst.Membrane, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Membrane](), formDiv)
		case "Effect":
			StageSetFormDivSelectFieldToField(&inst.Effect, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Effect](), formDiv)
		case "Timpani":
			StageSetFormDivSelectFieldToField(&inst.Timpani, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Timpani](), formDiv)
		case "Beater":
			StageSetFormDivSelectFieldToField(&inst.Beater, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Beater](), formDiv)
		case "Stick":
			StageSetFormDivSelectFieldToField(&inst.Stick, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Stick](), formDiv)
		case "StickLocation":
			StageSetFormDivEnumStringFieldToField(&inst.StickLocation, formDiv)
		case "OtherPercussion":
			StageSetFormDivSelectFieldToField(&inst.OtherPercussion, probe.stageSet.CommonStage.GetInstancesSet[*common.OtherText](), formDiv)
		}
	}
}

func saveStageSet_Pitched_DirectionStage(
	inst *direction.Pitched,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_PrincipalVoice_DirectionStage(
	inst *direction.PrincipalVoice,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Print_DirectionStage(
	inst *direction.Print,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "StaffSpacing":
			FormDivBasicFieldToField(&inst.StaffSpacing, formDiv)
		case "NewSystem":
			StageSetFormDivEnumStringFieldToField(&inst.NewSystem, formDiv)
		case "NewPage":
			StageSetFormDivEnumStringFieldToField(&inst.NewPage, formDiv)
		case "BlankPage":
			FormDivBasicFieldToField(&inst.BlankPage, formDiv)
		case "PageNumber":
			FormDivBasicFieldToField(&inst.PageNumber, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "PageLayout":
			StageSetFormDivSelectFieldToField(&inst.PageLayout, probe.stageSet.CommonStage.GetInstancesSet[*common.PageLayout](), formDiv)
		case "SystemLayout":
			StageSetFormDivSelectFieldToField(&inst.SystemLayout, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemLayout](), formDiv)
		case "MeasureLayout":
			StageSetFormDivSelectFieldToField(&inst.MeasureLayout, probe.stageSet.CommonStage.GetInstancesSet[*common.MeasureLayout](), formDiv)
		case "MeasureNumbering":
			StageSetFormDivSelectFieldToField(&inst.MeasureNumbering, probe.stageSet.DirectionStage.GetInstancesSet[*direction.MeasureNumbering](), formDiv)
		case "PartNameDisplay":
			StageSetFormDivSelectFieldToField(&inst.PartNameDisplay, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), formDiv)
		case "PartAbbreviationDisplay":
			StageSetFormDivSelectFieldToField(&inst.PartAbbreviationDisplay, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), formDiv)
		}
	}
}

func saveStageSet_Root_DirectionStage(
	inst *direction.Root,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RootStep":
			StageSetFormDivSelectFieldToField(&inst.RootStep, probe.stageSet.DirectionStage.GetInstancesSet[*direction.RootStep](), formDiv)
		case "RootAlter":
			StageSetFormDivSelectFieldToField(&inst.RootAlter, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarmonyAlter](), formDiv)
		}
	}
}

func saveStageSet_RootStep_DirectionStage(
	inst *direction.RootStep,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Text":
			FormDivBasicFieldToField(&inst.Text, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Scordatura_DirectionStage(
	inst *direction.Scordatura,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Sound_DirectionStage(
	inst *direction.Sound,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tempo":
			FormDivBasicFieldToField(&inst.Tempo, formDiv)
		case "Dynamics":
			FormDivBasicFieldToField(&inst.Dynamics, formDiv)
		case "Dacapo":
			StageSetFormDivEnumStringFieldToField(&inst.Dacapo, formDiv)
		case "Segno":
			FormDivBasicFieldToField(&inst.Segno, formDiv)
		case "Dalsegno":
			FormDivBasicFieldToField(&inst.Dalsegno, formDiv)
		case "Coda":
			FormDivBasicFieldToField(&inst.Coda, formDiv)
		case "Tocoda":
			FormDivBasicFieldToField(&inst.Tocoda, formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&inst.Divisions, formDiv)
		case "ForwardRepeat":
			StageSetFormDivEnumStringFieldToField(&inst.ForwardRepeat, formDiv)
		case "Fine":
			FormDivBasicFieldToField(&inst.Fine, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		case "Pizzicato":
			StageSetFormDivEnumStringFieldToField(&inst.Pizzicato, formDiv)
		case "Pan":
			FormDivBasicFieldToField(&inst.Pan, formDiv)
		case "Elevation":
			FormDivBasicFieldToField(&inst.Elevation, formDiv)
		case "DamperPedal":
			FormDivBasicFieldToField(&inst.DamperPedal, formDiv)
		case "SoftPedal":
			FormDivBasicFieldToField(&inst.SoftPedal, formDiv)
		case "SostenutoPedal":
			FormDivBasicFieldToField(&inst.SostenutoPedal, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Swing":
			StageSetFormDivSelectFieldToField(&inst.Swing, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Swing](), formDiv)
		case "Offset":
			StageSetFormDivSelectFieldToField(&inst.Offset, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), formDiv)
		}
	}
}

func saveStageSet_StaffDivide_DirectionStage(
	inst *direction.StaffDivide,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Stick_DirectionStage(
	inst *direction.Stick,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Tip":
			StageSetFormDivEnumStringFieldToField(&inst.Tip, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "DashedCircle":
			StageSetFormDivEnumStringFieldToField(&inst.DashedCircle, formDiv)
		case "StickType":
			StageSetFormDivEnumStringFieldToField(&inst.StickType, formDiv)
		case "StickMaterial":
			StageSetFormDivEnumStringFieldToField(&inst.StickMaterial, formDiv)
		}
	}
}

func saveStageSet_StringMute_DirectionStage(
	inst *direction.StringMute,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_StyleText_DirectionStage(
	inst *direction.StyleText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Swing_DirectionStage(
	inst *direction.Swing,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Straight":
			StageSetFormDivSelectFieldToField(&inst.Straight, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), formDiv)
		case "First":
			FormDivBasicFieldToField(&inst.First, formDiv)
		case "Second":
			FormDivBasicFieldToField(&inst.Second, formDiv)
		case "SwingType":
			StageSetFormDivEnumStringFieldToField(&inst.SwingType, formDiv)
		case "SwingStyle":
			FormDivBasicFieldToField(&inst.SwingStyle, formDiv)
		}
	}
}

func saveStageSet_Sync_DirectionStage(
	inst *direction.Sync,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Latency":
			FormDivBasicFieldToField(&inst.Latency, formDiv)
		case "Player":
			FormDivBasicFieldToField(&inst.Player, formDiv)
		case "TimeOnly":
			FormDivBasicFieldToField(&inst.TimeOnly, formDiv)
		}
	}
}

func saveStageSet_TimeModification_DirectionStage(
	inst *direction.TimeModification,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ActualNotes":
			FormDivBasicFieldToField(&inst.ActualNotes, formDiv)
		case "NormalNotes":
			FormDivBasicFieldToField(&inst.NormalNotes, formDiv)
		case "NormalType":
			StageSetFormDivEnumStringFieldToField(&inst.NormalType, formDiv)
		}
	}
}

func saveStageSet_Timpani_DirectionStage(
	inst *direction.Timpani,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		}
	}
}

func saveStageSet_VirtualInstrument_DirectionStage(
	inst *direction.VirtualInstrument,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "VirtualLibrary":
			FormDivBasicFieldToField(&inst.VirtualLibrary, formDiv)
		case "VirtualName":
			FormDivBasicFieldToField(&inst.VirtualName, formDiv)
		}
	}
}

func saveStageSet_Wedge_DirectionStage(
	inst *direction.Wedge,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Spread":
			FormDivBasicFieldToField(&inst.Spread, formDiv)
		case "Niente":
			StageSetFormDivEnumStringFieldToField(&inst.Niente, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "DashLength":
			FormDivBasicFieldToField(&inst.DashLength, formDiv)
		case "SpaceLength":
			FormDivBasicFieldToField(&inst.SpaceLength, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Wood_DirectionStage(
	inst *direction.Wood,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_BarStyleColor_BarlineStage(
	inst *barline.BarStyleColor,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Barline_BarlineStage(
	inst *barline.Barline,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Segno":
			FormDivBasicFieldToField(&inst.Segno, formDiv)
		case "Coda":
			FormDivBasicFieldToField(&inst.Coda, formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&inst.Divisions, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "BarStyle":
			StageSetFormDivSelectFieldToField(&inst.BarStyle, probe.stageSet.BarlineStage.GetInstancesSet[*barline.BarStyleColor](), formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		case "WavyLine":
			StageSetFormDivSelectFieldToField(&inst.WavyLine, probe.stageSet.CommonStage.GetInstancesSet[*common.WavyLine](), formDiv)
		case "Segno_1":
			StageSetFormDivSelectFieldToField(&inst.Segno_1, probe.stageSet.CommonStage.GetInstancesSet[*common.Segno](), formDiv)
		case "Coda_1":
			StageSetFormDivSelectFieldToField(&inst.Coda_1, probe.stageSet.CommonStage.GetInstancesSet[*common.Coda](), formDiv)
		case "Ending":
			StageSetFormDivSelectFieldToField(&inst.Ending, probe.stageSet.BarlineStage.GetInstancesSet[*barline.Ending](), formDiv)
		case "Repeat":
			StageSetFormDivSelectFieldToField(&inst.Repeat, probe.stageSet.BarlineStage.GetInstancesSet[*barline.Repeat](), formDiv)
		}
	}
}

func saveStageSet_Ending_BarlineStage(
	inst *barline.Ending,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "EndLength":
			FormDivBasicFieldToField(&inst.EndLength, formDiv)
		case "TextX":
			FormDivBasicFieldToField(&inst.TextX, formDiv)
		case "TextY":
			FormDivBasicFieldToField(&inst.TextY, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "System":
			StageSetFormDivEnumStringFieldToField(&inst.System, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Repeat_BarlineStage(
	inst *barline.Repeat,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Direction":
			StageSetFormDivEnumStringFieldToField(&inst.Direction, formDiv)
		case "Times":
			FormDivBasicFieldToField(&inst.Times, formDiv)
		case "AfterJump":
			StageSetFormDivEnumStringFieldToField(&inst.AfterJump, formDiv)
		case "Winged":
			StageSetFormDivEnumStringFieldToField(&inst.Winged, formDiv)
		}
	}
}

func saveStageSet_Attributes_AttributesStage(
	inst *attributes.Attributes,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Footnote":
			StageSetFormDivSelectFieldToField(&inst.Footnote, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), formDiv)
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), formDiv)
		case "Divisions":
			FormDivBasicFieldToField(&inst.Divisions, formDiv)
		case "Staves":
			FormDivBasicFieldToField(&inst.Staves, formDiv)
		case "PartSymbol":
			StageSetFormDivSelectFieldToField(&inst.PartSymbol, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.PartSymbol](), formDiv)
		case "Instruments":
			FormDivBasicFieldToField(&inst.Instruments, formDiv)
		}
	}
}

func saveStageSet_AttributesDirective_AttributesStage(
	inst *attributes.AttributesDirective,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_BeatRepeat_AttributesStage(
	inst *attributes.BeatRepeat,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Slashes":
			FormDivBasicFieldToField(&inst.Slashes, formDiv)
		case "UseDots":
			StageSetFormDivEnumStringFieldToField(&inst.UseDots, formDiv)
		case "SlashType":
			StageSetFormDivEnumStringFieldToField(&inst.SlashType, formDiv)
		case "ExceptVoice":
			FormDivBasicFieldToField(&inst.ExceptVoice, formDiv)
		}
	}
}

func saveStageSet_Cancel_AttributesStage(
	inst *attributes.Cancel,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Location":
			StageSetFormDivEnumStringFieldToField(&inst.Location, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Clef_AttributesStage(
	inst *attributes.Clef,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Additional":
			StageSetFormDivEnumStringFieldToField(&inst.Additional, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "AfterBarline":
			StageSetFormDivEnumStringFieldToField(&inst.AfterBarline, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Sign":
			StageSetFormDivEnumStringFieldToField(&inst.Sign, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "ClefOctaveChange":
			FormDivBasicFieldToField(&inst.ClefOctaveChange, formDiv)
		}
	}
}

func saveStageSet_Double_AttributesStage(
	inst *attributes.Double,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Above":
			StageSetFormDivEnumStringFieldToField(&inst.Above, formDiv)
		}
	}
}

func saveStageSet_ForPart_AttributesStage(
	inst *attributes.ForPart,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "PartClef":
			StageSetFormDivSelectFieldToField(&inst.PartClef, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.PartClef](), formDiv)
		case "PartTranspose":
			StageSetFormDivSelectFieldToField(&inst.PartTranspose, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.PartTranspose](), formDiv)
		}
	}
}

func saveStageSet_Interchangeable_AttributesStage(
	inst *attributes.Interchangeable,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Separator":
			StageSetFormDivEnumStringFieldToField(&inst.Separator, formDiv)
		case "TimeRelation":
			StageSetFormDivEnumStringFieldToField(&inst.TimeRelation, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "BeatType":
			FormDivBasicFieldToField(&inst.BeatType, formDiv)
		}
	}
}

func saveStageSet_Key_AttributesStage(
	inst *attributes.Key,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Cancel":
			StageSetFormDivSelectFieldToField(&inst.Cancel, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Cancel](), formDiv)
		case "Fifths":
			FormDivBasicFieldToField(&inst.Fifths, formDiv)
		case "Mode":
			FormDivBasicFieldToField(&inst.Mode, formDiv)
		case "KeyStep":
			StageSetFormDivEnumStringFieldToField(&inst.KeyStep, formDiv)
		case "KeyAlter":
			FormDivBasicFieldToField(&inst.KeyAlter, formDiv)
		}
	}
}

func saveStageSet_KeyAccidental_AttributesStage(
	inst *attributes.KeyAccidental,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_KeyOctave_AttributesStage(
	inst *attributes.KeyOctave,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Cancel":
			StageSetFormDivEnumStringFieldToField(&inst.Cancel, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_LineDetail_AttributesStage(
	inst *attributes.LineDetail,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "LineType":
			StageSetFormDivEnumStringFieldToField(&inst.LineType, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		}
	}
}

func saveStageSet_MeasureRepeat_AttributesStage(
	inst *attributes.MeasureRepeat,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Slashes":
			FormDivBasicFieldToField(&inst.Slashes, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_MeasureStyle_AttributesStage(
	inst *attributes.MeasureStyle,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "MultipleRest":
			StageSetFormDivSelectFieldToField(&inst.MultipleRest, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.MultipleRest](), formDiv)
		case "MeasureRepeat":
			StageSetFormDivSelectFieldToField(&inst.MeasureRepeat, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.MeasureRepeat](), formDiv)
		case "BeatRepeat":
			StageSetFormDivSelectFieldToField(&inst.BeatRepeat, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.BeatRepeat](), formDiv)
		case "Slash":
			StageSetFormDivSelectFieldToField(&inst.Slash, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Slash](), formDiv)
		}
	}
}

func saveStageSet_MultipleRest_AttributesStage(
	inst *attributes.MultipleRest,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "UseSymbols":
			StageSetFormDivEnumStringFieldToField(&inst.UseSymbols, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_PartClef_AttributesStage(
	inst *attributes.PartClef,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Sign":
			StageSetFormDivEnumStringFieldToField(&inst.Sign, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "ClefOctaveChange":
			FormDivBasicFieldToField(&inst.ClefOctaveChange, formDiv)
		}
	}
}

func saveStageSet_PartSymbol_AttributesStage(
	inst *attributes.PartSymbol,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "TopStaff":
			FormDivBasicFieldToField(&inst.TopStaff, formDiv)
		case "BottomStaff":
			FormDivBasicFieldToField(&inst.BottomStaff, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_PartTranspose_AttributesStage(
	inst *attributes.PartTranspose,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Diatonic":
			FormDivBasicFieldToField(&inst.Diatonic, formDiv)
		case "Chromatic":
			FormDivBasicFieldToField(&inst.Chromatic, formDiv)
		case "OctaveChange":
			FormDivBasicFieldToField(&inst.OctaveChange, formDiv)
		case "Double":
			StageSetFormDivSelectFieldToField(&inst.Double, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Double](), formDiv)
		}
	}
}

func saveStageSet_Slash_AttributesStage(
	inst *attributes.Slash,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "UseDots":
			StageSetFormDivEnumStringFieldToField(&inst.UseDots, formDiv)
		case "UseStems":
			StageSetFormDivEnumStringFieldToField(&inst.UseStems, formDiv)
		case "SlashType":
			StageSetFormDivEnumStringFieldToField(&inst.SlashType, formDiv)
		case "ExceptVoice":
			FormDivBasicFieldToField(&inst.ExceptVoice, formDiv)
		}
	}
}

func saveStageSet_StaffDetails_AttributesStage(
	inst *attributes.StaffDetails,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "ShowFrets":
			StageSetFormDivEnumStringFieldToField(&inst.ShowFrets, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "PrintSpacing":
			StageSetFormDivEnumStringFieldToField(&inst.PrintSpacing, formDiv)
		case "StaffType":
			StageSetFormDivEnumStringFieldToField(&inst.StaffType, formDiv)
		case "StaffLines":
			FormDivBasicFieldToField(&inst.StaffLines, formDiv)
		case "Capo":
			FormDivBasicFieldToField(&inst.Capo, formDiv)
		case "StaffSize":
			StageSetFormDivSelectFieldToField(&inst.StaffSize, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.StaffSize](), formDiv)
		}
	}
}

func saveStageSet_StaffSize_AttributesStage(
	inst *attributes.StaffSize,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Scaling":
			FormDivBasicFieldToField(&inst.Scaling, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_StaffTuning_AttributesStage(
	inst *attributes.StaffTuning,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Line":
			FormDivBasicFieldToField(&inst.Line, formDiv)
		case "TuningStep":
			StageSetFormDivEnumStringFieldToField(&inst.TuningStep, formDiv)
		case "TuningAlter":
			FormDivBasicFieldToField(&inst.TuningAlter, formDiv)
		case "TuningOctave":
			FormDivBasicFieldToField(&inst.TuningOctave, formDiv)
		}
	}
}

func saveStageSet_Time_AttributesStage(
	inst *attributes.Time,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Symbol":
			StageSetFormDivEnumStringFieldToField(&inst.Symbol, formDiv)
		case "Separator":
			StageSetFormDivEnumStringFieldToField(&inst.Separator, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "BeatType":
			FormDivBasicFieldToField(&inst.BeatType, formDiv)
		case "Interchangeable":
			StageSetFormDivSelectFieldToField(&inst.Interchangeable, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Interchangeable](), formDiv)
		case "SenzaMisura":
			FormDivBasicFieldToField(&inst.SenzaMisura, formDiv)
		}
	}
}

func saveStageSet_Transpose_AttributesStage(
	inst *attributes.Transpose,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Diatonic":
			FormDivBasicFieldToField(&inst.Diatonic, formDiv)
		case "Chromatic":
			FormDivBasicFieldToField(&inst.Chromatic, formDiv)
		case "OctaveChange":
			FormDivBasicFieldToField(&inst.OctaveChange, formDiv)
		case "Double":
			StageSetFormDivSelectFieldToField(&inst.Double, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Double](), formDiv)
		}
	}
}

func saveStageSet_AccidentalText_CommonStage(
	inst *common.AccidentalText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Space":
			FormDivBasicFieldToField(&inst.Space, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "LineHeight":
			FormDivBasicFieldToField(&inst.LineHeight, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Appearance_CommonStage(
	inst *common.Appearance,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_Bookmark_CommonStage(
	inst *common.Bookmark,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "NameAttr":
			FormDivBasicFieldToField(&inst.NameAttr, formDiv)
		case "Element":
			FormDivBasicFieldToField(&inst.Element, formDiv)
		case "Position":
			FormDivBasicFieldToField(&inst.Position, formDiv)
		}
	}
}

func saveStageSet_Coda_CommonStage(
	inst *common.Coda,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Distance_CommonStage(
	inst *common.Distance,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Dynamics_CommonStage(
	inst *common.Dynamics,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_Empty_CommonStage(
	inst *common.Empty,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_EmptyPlacement_CommonStage(
	inst *common.EmptyPlacement,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		}
	}
}

func saveStageSet_EmptyPlacementSmufl_CommonStage(
	inst *common.EmptyPlacementSmufl,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		}
	}
}

func saveStageSet_EmptyPrintObjectStyleAlign_CommonStage(
	inst *common.EmptyPrintObjectStyleAlign,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		}
	}
}

func saveStageSet_EmptyPrintStyle_CommonStage(
	inst *common.EmptyPrintStyle,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		}
	}
}

func saveStageSet_EmptyPrintStyleAlign_CommonStage(
	inst *common.EmptyPrintStyleAlign,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		}
	}
}

func saveStageSet_EmptyPrintStyleAlignId_CommonStage(
	inst *common.EmptyPrintStyleAlignId,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_EmptyTrillSound_CommonStage(
	inst *common.EmptyTrillSound,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "StartNote":
			StageSetFormDivEnumStringFieldToField(&inst.StartNote, formDiv)
		case "TrillStep":
			StageSetFormDivEnumStringFieldToField(&inst.TrillStep, formDiv)
		case "TwoNoteTurn":
			StageSetFormDivEnumStringFieldToField(&inst.TwoNoteTurn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "SecondBeat":
			FormDivBasicFieldToField(&inst.SecondBeat, formDiv)
		case "LastBeat":
			FormDivBasicFieldToField(&inst.LastBeat, formDiv)
		}
	}
}

func saveStageSet_Encoding_CommonStage(
	inst *common.Encoding,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EncodingDate":
			FormDivBasicFieldToField(&inst.EncodingDate, formDiv)
		case "Software":
			FormDivBasicFieldToField(&inst.Software, formDiv)
		case "EncodingDescription":
			FormDivBasicFieldToField(&inst.EncodingDescription, formDiv)
		}
	}
}

func saveStageSet_Fermata_CommonStage(
	inst *common.Fermata,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			StageSetFormDivEnumStringFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Fingering_CommonStage(
	inst *common.Fingering,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Substitution":
			StageSetFormDivEnumStringFieldToField(&inst.Substitution, formDiv)
		case "Alternate":
			StageSetFormDivEnumStringFieldToField(&inst.Alternate, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_FormattedSymbol_CommonStage(
	inst *common.FormattedSymbol,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "LineHeight":
			FormDivBasicFieldToField(&inst.LineHeight, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_FormattedSymbolId_CommonStage(
	inst *common.FormattedSymbolId,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "LineHeight":
			FormDivBasicFieldToField(&inst.LineHeight, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_FormattedText_CommonStage(
	inst *common.FormattedText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Space":
			FormDivBasicFieldToField(&inst.Space, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "LineHeight":
			FormDivBasicFieldToField(&inst.LineHeight, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_FormattedTextId_CommonStage(
	inst *common.FormattedTextId,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Space":
			FormDivBasicFieldToField(&inst.Space, formDiv)
		case "Justify":
			StageSetFormDivEnumStringFieldToField(&inst.Justify, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Underline":
			FormDivBasicFieldToField(&inst.Underline, formDiv)
		case "Overline":
			FormDivBasicFieldToField(&inst.Overline, formDiv)
		case "LineThrough":
			FormDivBasicFieldToField(&inst.LineThrough, formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&inst.Rotation, formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&inst.LetterSpacing, formDiv)
		case "LineHeight":
			FormDivBasicFieldToField(&inst.LineHeight, formDiv)
		case "Dir":
			StageSetFormDivEnumStringFieldToField(&inst.Dir, formDiv)
		case "Enclosure":
			StageSetFormDivEnumStringFieldToField(&inst.Enclosure, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Fret_CommonStage(
	inst *common.Fret,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Glyph_CommonStage(
	inst *common.Glyph,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_HorizontalTurn_CommonStage(
	inst *common.HorizontalTurn,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Slash":
			StageSetFormDivEnumStringFieldToField(&inst.Slash, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "StartNote":
			StageSetFormDivEnumStringFieldToField(&inst.StartNote, formDiv)
		case "TrillStep":
			StageSetFormDivEnumStringFieldToField(&inst.TrillStep, formDiv)
		case "TwoNoteTurn":
			StageSetFormDivEnumStringFieldToField(&inst.TwoNoteTurn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "SecondBeat":
			FormDivBasicFieldToField(&inst.SecondBeat, formDiv)
		case "LastBeat":
			FormDivBasicFieldToField(&inst.LastBeat, formDiv)
		}
	}
}

func saveStageSet_Identification_CommonStage(
	inst *common.Identification,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Encoding":
			StageSetFormDivSelectFieldToField(&inst.Encoding, probe.stageSet.CommonStage.GetInstancesSet[*common.Encoding](), formDiv)
		case "Source":
			FormDivBasicFieldToField(&inst.Source, formDiv)
		case "Miscellaneous":
			StageSetFormDivSelectFieldToField(&inst.Miscellaneous, probe.stageSet.CommonStage.GetInstancesSet[*common.Miscellaneous](), formDiv)
		}
	}
}

func saveStageSet_Level_CommonStage(
	inst *common.Level,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Reference":
			StageSetFormDivEnumStringFieldToField(&inst.Reference, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Parentheses":
			StageSetFormDivEnumStringFieldToField(&inst.Parentheses, formDiv)
		case "Bracket":
			StageSetFormDivEnumStringFieldToField(&inst.Bracket, formDiv)
		case "Size":
			StageSetFormDivEnumStringFieldToField(&inst.Size, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_LineWidth_CommonStage(
	inst *common.LineWidth,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Link_CommonStage(
	inst *common.Link,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NameAttr":
			FormDivBasicFieldToField(&inst.NameAttr, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			StageSetFormDivEnumStringFieldToField(&inst.Show, formDiv)
		case "Actuate":
			StageSetFormDivEnumStringFieldToField(&inst.Actuate, formDiv)
		case "Element":
			FormDivBasicFieldToField(&inst.Element, formDiv)
		case "Position":
			FormDivBasicFieldToField(&inst.Position, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		}
	}
}

func saveStageSet_MeasureLayout_CommonStage(
	inst *common.MeasureLayout,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "MeasureDistance":
			FormDivBasicFieldToField(&inst.MeasureDistance, formDiv)
		}
	}
}

func saveStageSet_MidiDevice_CommonStage(
	inst *common.MidiDevice,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Port":
			FormDivBasicFieldToField(&inst.Port, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_MidiInstrument_CommonStage(
	inst *common.MidiInstrument,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "MidiChannel":
			FormDivBasicFieldToField(&inst.MidiChannel, formDiv)
		case "MidiName":
			FormDivBasicFieldToField(&inst.MidiName, formDiv)
		case "MidiBank":
			FormDivBasicFieldToField(&inst.MidiBank, formDiv)
		case "MidiProgram":
			FormDivBasicFieldToField(&inst.MidiProgram, formDiv)
		case "MidiUnpitched":
			FormDivBasicFieldToField(&inst.MidiUnpitched, formDiv)
		case "Volume":
			FormDivBasicFieldToField(&inst.Volume, formDiv)
		case "Pan":
			FormDivBasicFieldToField(&inst.Pan, formDiv)
		case "Elevation":
			FormDivBasicFieldToField(&inst.Elevation, formDiv)
		}
	}
}

func saveStageSet_Miscellaneous_CommonStage(
	inst *common.Miscellaneous,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_MiscellaneousField_CommonStage(
	inst *common.MiscellaneousField,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NameAttr":
			FormDivBasicFieldToField(&inst.NameAttr, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_NameDisplay_CommonStage(
	inst *common.NameDisplay,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PrintObject":
			StageSetFormDivEnumStringFieldToField(&inst.PrintObject, formDiv)
		}
	}
}

func saveStageSet_NoteSize_CommonStage(
	inst *common.NoteSize,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OtherAppearance_CommonStage(
	inst *common.OtherAppearance,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OtherPlay_CommonStage(
	inst *common.OtherPlay,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_OtherText_CommonStage(
	inst *common.OtherText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_PageLayout_CommonStage(
	inst *common.PageLayout,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "PageHeight":
			FormDivBasicFieldToField(&inst.PageHeight, formDiv)
		case "PageWidth":
			FormDivBasicFieldToField(&inst.PageWidth, formDiv)
		}
	}
}

func saveStageSet_PageMargins_CommonStage(
	inst *common.PageMargins,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "LeftMargin":
			FormDivBasicFieldToField(&inst.LeftMargin, formDiv)
		case "RightMargin":
			FormDivBasicFieldToField(&inst.RightMargin, formDiv)
		case "TopMargin":
			FormDivBasicFieldToField(&inst.TopMargin, formDiv)
		case "BottomMargin":
			FormDivBasicFieldToField(&inst.BottomMargin, formDiv)
		}
	}
}

func saveStageSet_Play_CommonStage(
	inst *common.Play,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Ipa":
			FormDivBasicFieldToField(&inst.Ipa, formDiv)
		case "Mute":
			StageSetFormDivEnumStringFieldToField(&inst.Mute, formDiv)
		case "SemiPitched":
			StageSetFormDivEnumStringFieldToField(&inst.SemiPitched, formDiv)
		}
	}
}

func saveStageSet_Scaling_CommonStage(
	inst *common.Scaling,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Millimeters":
			FormDivBasicFieldToField(&inst.Millimeters, formDiv)
		case "Tenths":
			FormDivBasicFieldToField(&inst.Tenths, formDiv)
		}
	}
}

func saveStageSet_Segno_CommonStage(
	inst *common.Segno,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Halign":
			StageSetFormDivEnumStringFieldToField(&inst.Halign, formDiv)
		case "Valign":
			StageSetFormDivEnumStringFieldToField(&inst.Valign, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_StaffLayout_CommonStage(
	inst *common.StaffLayout,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "StaffDistance":
			FormDivBasicFieldToField(&inst.StaffDistance, formDiv)
		}
	}
}

func saveStageSet_StringType_CommonStage(
	inst *common.StringType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&inst.FontFamily, formDiv)
		case "FontStyle":
			StageSetFormDivEnumStringFieldToField(&inst.FontStyle, formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&inst.FontSize, formDiv)
		case "FontWeight":
			StageSetFormDivEnumStringFieldToField(&inst.FontWeight, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Supports_CommonStage(
	inst *common.Supports,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Element":
			FormDivBasicFieldToField(&inst.Element, formDiv)
		case "Attribute":
			FormDivBasicFieldToField(&inst.Attribute, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_SystemDividers_CommonStage(
	inst *common.SystemDividers,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "LeftDivider":
			StageSetFormDivSelectFieldToField(&inst.LeftDivider, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintObjectStyleAlign](), formDiv)
		case "RightDivider":
			StageSetFormDivSelectFieldToField(&inst.RightDivider, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintObjectStyleAlign](), formDiv)
		}
	}
}

func saveStageSet_SystemLayout_CommonStage(
	inst *common.SystemLayout,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SystemMargins":
			StageSetFormDivSelectFieldToField(&inst.SystemMargins, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemMargins](), formDiv)
		case "SystemDistance":
			FormDivBasicFieldToField(&inst.SystemDistance, formDiv)
		case "TopSystemDistance":
			FormDivBasicFieldToField(&inst.TopSystemDistance, formDiv)
		case "SystemDividers":
			StageSetFormDivSelectFieldToField(&inst.SystemDividers, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemDividers](), formDiv)
		}
	}
}

func saveStageSet_SystemMargins_CommonStage(
	inst *common.SystemMargins,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "LeftMargin":
			FormDivBasicFieldToField(&inst.LeftMargin, formDiv)
		case "RightMargin":
			FormDivBasicFieldToField(&inst.RightMargin, formDiv)
		}
	}
}

func saveStageSet_TypedText_CommonStage(
	inst *common.TypedText,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			FormDivBasicFieldToField(&inst.Type, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_WavyLine_CommonStage(
	inst *common.WavyLine,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Number":
			FormDivBasicFieldToField(&inst.Number, formDiv)
		case "Smufl":
			FormDivBasicFieldToField(&inst.Smufl, formDiv)
		case "DefaultX":
			FormDivBasicFieldToField(&inst.DefaultX, formDiv)
		case "DefaultY":
			FormDivBasicFieldToField(&inst.DefaultY, formDiv)
		case "RelativeX":
			FormDivBasicFieldToField(&inst.RelativeX, formDiv)
		case "RelativeY":
			FormDivBasicFieldToField(&inst.RelativeY, formDiv)
		case "Placement":
			StageSetFormDivEnumStringFieldToField(&inst.Placement, formDiv)
		case "Color":
			FormDivBasicFieldToField(&inst.Color, formDiv)
		case "StartNote":
			StageSetFormDivEnumStringFieldToField(&inst.StartNote, formDiv)
		case "TrillStep":
			StageSetFormDivEnumStringFieldToField(&inst.TrillStep, formDiv)
		case "TwoNoteTurn":
			StageSetFormDivEnumStringFieldToField(&inst.TwoNoteTurn, formDiv)
		case "Accelerate":
			StageSetFormDivEnumStringFieldToField(&inst.Accelerate, formDiv)
		case "Beats":
			FormDivBasicFieldToField(&inst.Beats, formDiv)
		case "SecondBeat":
			FormDivBasicFieldToField(&inst.SecondBeat, formDiv)
		case "LastBeat":
			FormDivBasicFieldToField(&inst.LastBeat, formDiv)
		}
	}
}

func saveStageSet_ArcType_XlinkStage(
	inst *xlink.ArcType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			StageSetFormDivEnumStringFieldToField(&inst.Show, formDiv)
		case "Actuate":
			StageSetFormDivEnumStringFieldToField(&inst.Actuate, formDiv)
		case "From":
			FormDivBasicFieldToField(&inst.From, formDiv)
		case "To":
			FormDivBasicFieldToField(&inst.To, formDiv)
		}
	}
}

func saveStageSet_Extended_XlinkStage(
	inst *xlink.Extended,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		}
	}
}

func saveStageSet_LocatorType_XlinkStage(
	inst *xlink.LocatorType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Label":
			FormDivBasicFieldToField(&inst.Label, formDiv)
		}
	}
}

func saveStageSet_ResourceType_XlinkStage(
	inst *xlink.ResourceType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Label":
			FormDivBasicFieldToField(&inst.Label, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_Simple_XlinkStage(
	inst *xlink.Simple,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			StageSetFormDivEnumStringFieldToField(&inst.Show, formDiv)
		case "Actuate":
			StageSetFormDivEnumStringFieldToField(&inst.Actuate, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}

func saveStageSet_TitleEltType_XlinkStage(
	inst *xlink.TitleEltType,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "Value":
			FormDivBasicFieldToField(&inst.Value, formDiv)
		}
	}
}


func StageSetNewInstance_MusicXMLDocument_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MusicXMLDocument",
	}).Stage(probe.formStage)
	inst := new(models.MusicXMLDocument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_MusicXMLDocument_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_MusicXMLDocument_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Credit_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Credit",
	}).Stage(probe.formStage)
	inst := new(score.Credit)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Credit_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Credit_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Defaults_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Defaults",
	}).Stage(probe.formStage)
	inst := new(score.Defaults)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Defaults_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Defaults_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyFont_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyFont",
	}).Stage(probe.formStage)
	inst := new(score.EmptyFont)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_EmptyFont_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_EmptyFont_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GroupBarline_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New GroupBarline",
	}).Stage(probe.formStage)
	inst := new(score.GroupBarline)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_GroupBarline_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_GroupBarline_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GroupName_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New GroupName",
	}).Stage(probe.formStage)
	inst := new(score.GroupName)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_GroupName_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_GroupName_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GroupSymbol_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New GroupSymbol",
	}).Stage(probe.formStage)
	inst := new(score.GroupSymbol)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_GroupSymbol_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_GroupSymbol_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InstrumentLink_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New InstrumentLink",
	}).Stage(probe.formStage)
	inst := new(score.InstrumentLink)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_InstrumentLink_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_InstrumentLink_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LyricFont_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LyricFont",
	}).Stage(probe.formStage)
	inst := new(score.LyricFont)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_LyricFont_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_LyricFont_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LyricLanguage_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LyricLanguage",
	}).Stage(probe.formStage)
	inst := new(score.LyricLanguage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_LyricLanguage_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_LyricLanguage_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Measure_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Measure",
	}).Stage(probe.formStage)
	inst := new(score.Measure)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Measure_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Measure_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeasureMusicDataChoice_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MeasureMusicDataChoice",
	}).Stage(probe.formStage)
	inst := new(score.MeasureMusicDataChoice)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_MeasureMusicDataChoice_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_MeasureMusicDataChoice_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Opus_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Opus",
	}).Stage(probe.formStage)
	inst := new(score.Opus)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Opus_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Opus_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Part_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Part",
	}).Stage(probe.formStage)
	inst := new(score.Part)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Part_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Part_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartGroup_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartGroup",
	}).Stage(probe.formStage)
	inst := new(score.PartGroup)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_PartGroup_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_PartGroup_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartLink_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartLink",
	}).Stage(probe.formStage)
	inst := new(score.PartLink)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_PartLink_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_PartLink_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartList_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartList",
	}).Stage(probe.formStage)
	inst := new(score.PartList)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_PartList_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_PartList_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartName_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartName",
	}).Stage(probe.formStage)
	inst := new(score.PartName)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_PartName_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_PartName_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Player_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Player",
	}).Stage(probe.formStage)
	inst := new(score.Player)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Player_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Player_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ScoreInstrument_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ScoreInstrument",
	}).Stage(probe.formStage)
	inst := new(score.ScoreInstrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_ScoreInstrument_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_ScoreInstrument_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ScorePart_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ScorePart",
	}).Stage(probe.formStage)
	inst := new(score.ScorePart)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_ScorePart_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_ScorePart_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ScorePartwise_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ScorePartwise",
	}).Stage(probe.formStage)
	inst := new(score.ScorePartwise)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_ScorePartwise_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_ScorePartwise_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ScoreTimewise_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ScoreTimewise",
	}).Stage(probe.formStage)
	inst := new(score.ScoreTimewise)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_ScoreTimewise_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_ScoreTimewise_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ScoreTimewiseMeasure_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ScoreTimewiseMeasure",
	}).Stage(probe.formStage)
	inst := new(score.ScoreTimewiseMeasure)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_ScoreTimewiseMeasure_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_ScoreTimewiseMeasure_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ScoreTimewiseMeasurePart_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ScoreTimewiseMeasurePart",
	}).Stage(probe.formStage)
	inst := new(score.ScoreTimewiseMeasurePart)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_ScoreTimewiseMeasurePart_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_ScoreTimewiseMeasurePart_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Work_ScoreStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Work",
	}).Stage(probe.formStage)
	inst := new(score.Work)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ScoreStage.Lock()
			defer probe.stageSet.ScoreStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ScoreStage)
			saveStageSet_Work_ScoreStage(inst, probe, formGroup)
			probe.stageSet.ScoreStage.Commit()
			updateStageSetTable_Work_ScoreStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accidental_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Accidental",
	}).Stage(probe.formStage)
	inst := new(note.Accidental)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Accidental_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Accidental_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AccidentalMark_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AccidentalMark",
	}).Stage(probe.formStage)
	inst := new(note.AccidentalMark)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_AccidentalMark_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_AccidentalMark_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Arpeggiate_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Arpeggiate",
	}).Stage(probe.formStage)
	inst := new(note.Arpeggiate)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Arpeggiate_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Arpeggiate_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Arrow_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Arrow",
	}).Stage(probe.formStage)
	inst := new(note.Arrow)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Arrow_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Arrow_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Articulations_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Articulations",
	}).Stage(probe.formStage)
	inst := new(note.Articulations)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Articulations_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Articulations_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Assess_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Assess",
	}).Stage(probe.formStage)
	inst := new(note.Assess)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Assess_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Assess_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Backup_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Backup",
	}).Stage(probe.formStage)
	inst := new(note.Backup)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Backup_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Backup_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Beam_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Beam",
	}).Stage(probe.formStage)
	inst := new(note.Beam)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Beam_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Beam_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bend_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Bend",
	}).Stage(probe.formStage)
	inst := new(note.Bend)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Bend_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Bend_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BreathMark_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BreathMark",
	}).Stage(probe.formStage)
	inst := new(note.BreathMark)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_BreathMark_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_BreathMark_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Caesura_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Caesura",
	}).Stage(probe.formStage)
	inst := new(note.Caesura)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Caesura_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Caesura_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Elision_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Elision",
	}).Stage(probe.formStage)
	inst := new(note.Elision)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Elision_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Elision_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyLine_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyLine",
	}).Stage(probe.formStage)
	inst := new(note.EmptyLine)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_EmptyLine_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_EmptyLine_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Extend_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Extend",
	}).Stage(probe.formStage)
	inst := new(note.Extend)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Extend_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Extend_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Figure_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Figure",
	}).Stage(probe.formStage)
	inst := new(note.Figure)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Figure_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Figure_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FiguredBass_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FiguredBass",
	}).Stage(probe.formStage)
	inst := new(note.FiguredBass)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_FiguredBass_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_FiguredBass_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Forward_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Forward",
	}).Stage(probe.formStage)
	inst := new(note.Forward)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Forward_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Forward_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Glissando_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Glissando",
	}).Stage(probe.formStage)
	inst := new(note.Glissando)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Glissando_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Glissando_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Grace_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Grace",
	}).Stage(probe.formStage)
	inst := new(note.Grace)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Grace_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Grace_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HammerOnPullOff_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HammerOnPullOff",
	}).Stage(probe.formStage)
	inst := new(note.HammerOnPullOff)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_HammerOnPullOff_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_HammerOnPullOff_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Handbell_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Handbell",
	}).Stage(probe.formStage)
	inst := new(note.Handbell)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Handbell_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Handbell_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HarmonClosed_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HarmonClosed",
	}).Stage(probe.formStage)
	inst := new(note.HarmonClosed)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_HarmonClosed_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_HarmonClosed_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HarmonMute_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HarmonMute",
	}).Stage(probe.formStage)
	inst := new(note.HarmonMute)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_HarmonMute_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_HarmonMute_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmonic_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Harmonic",
	}).Stage(probe.formStage)
	inst := new(note.Harmonic)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Harmonic_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Harmonic_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HeelToe_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HeelToe",
	}).Stage(probe.formStage)
	inst := new(note.HeelToe)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_HeelToe_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_HeelToe_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Hole_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Hole",
	}).Stage(probe.formStage)
	inst := new(note.Hole)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Hole_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Hole_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoleClosed_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HoleClosed",
	}).Stage(probe.formStage)
	inst := new(note.HoleClosed)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_HoleClosed_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_HoleClosed_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Instrument_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Instrument",
	}).Stage(probe.formStage)
	inst := new(note.Instrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Instrument_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Instrument_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Listen_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Listen",
	}).Stage(probe.formStage)
	inst := new(note.Listen)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Listen_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Listen_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Lyric_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Lyric",
	}).Stage(probe.formStage)
	inst := new(note.Lyric)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Lyric_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Lyric_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Mordent_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Mordent",
	}).Stage(probe.formStage)
	inst := new(note.Mordent)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Mordent_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Mordent_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NonArpeggiate_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NonArpeggiate",
	}).Stage(probe.formStage)
	inst := new(note.NonArpeggiate)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_NonArpeggiate_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_NonArpeggiate_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Notations_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Notations",
	}).Stage(probe.formStage)
	inst := new(note.Notations)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Notations_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Notations_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Note_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Note",
	}).Stage(probe.formStage)
	inst := new(note.Note)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Note_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Note_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteType_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteType",
	}).Stage(probe.formStage)
	inst := new(note.NoteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_NoteType_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_NoteType_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Notehead_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Notehead",
	}).Stage(probe.formStage)
	inst := new(note.Notehead)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Notehead_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Notehead_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteheadText_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteheadText",
	}).Stage(probe.formStage)
	inst := new(note.NoteheadText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_NoteheadText_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_NoteheadText_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Ornaments_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Ornaments",
	}).Stage(probe.formStage)
	inst := new(note.Ornaments)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Ornaments_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Ornaments_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherNotation_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherNotation",
	}).Stage(probe.formStage)
	inst := new(note.OtherNotation)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_OtherNotation_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_OtherNotation_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherPlacementText_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherPlacementText",
	}).Stage(probe.formStage)
	inst := new(note.OtherPlacementText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_OtherPlacementText_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_OtherPlacementText_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pitch_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Pitch",
	}).Stage(probe.formStage)
	inst := new(note.Pitch)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Pitch_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Pitch_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PlacementText_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PlacementText",
	}).Stage(probe.formStage)
	inst := new(note.PlacementText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_PlacementText_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_PlacementText_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Release_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Release",
	}).Stage(probe.formStage)
	inst := new(note.Release)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Release_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Release_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Rest_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Rest",
	}).Stage(probe.formStage)
	inst := new(note.Rest)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Rest_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Rest_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Slide_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Slide",
	}).Stage(probe.formStage)
	inst := new(note.Slide)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Slide_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Slide_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Slur_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Slur",
	}).Stage(probe.formStage)
	inst := new(note.Slur)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Slur_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Slur_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Stem_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Stem",
	}).Stage(probe.formStage)
	inst := new(note.Stem)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Stem_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Stem_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StrongAccent_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StrongAccent",
	}).Stage(probe.formStage)
	inst := new(note.StrongAccent)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_StrongAccent_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_StrongAccent_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tap_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tap",
	}).Stage(probe.formStage)
	inst := new(note.Tap)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Tap_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Tap_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Technical_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Technical",
	}).Stage(probe.formStage)
	inst := new(note.Technical)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Technical_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Technical_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextElementData_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TextElementData",
	}).Stage(probe.formStage)
	inst := new(note.TextElementData)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_TextElementData_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_TextElementData_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tie_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tie",
	}).Stage(probe.formStage)
	inst := new(note.Tie)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Tie_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Tie_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tied_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tied",
	}).Stage(probe.formStage)
	inst := new(note.Tied)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Tied_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Tied_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tremolo_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tremolo",
	}).Stage(probe.formStage)
	inst := new(note.Tremolo)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Tremolo_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Tremolo_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Tuplet_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Tuplet",
	}).Stage(probe.formStage)
	inst := new(note.Tuplet)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Tuplet_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Tuplet_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TupletDot_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TupletDot",
	}).Stage(probe.formStage)
	inst := new(note.TupletDot)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_TupletDot_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_TupletDot_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TupletNumber_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TupletNumber",
	}).Stage(probe.formStage)
	inst := new(note.TupletNumber)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_TupletNumber_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_TupletNumber_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TupletPortion_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TupletPortion",
	}).Stage(probe.formStage)
	inst := new(note.TupletPortion)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_TupletPortion_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_TupletPortion_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TupletType_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TupletType",
	}).Stage(probe.formStage)
	inst := new(note.TupletType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_TupletType_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_TupletType_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Unpitched_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Unpitched",
	}).Stage(probe.formStage)
	inst := new(note.Unpitched)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Unpitched_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Unpitched_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wait_NoteStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Wait",
	}).Stage(probe.formStage)
	inst := new(note.Wait)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.NoteStage.Lock()
			defer probe.stageSet.NoteStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.NoteStage)
			saveStageSet_Wait_NoteStage(inst, probe, formGroup)
			probe.stageSet.NoteStage.Commit()
			updateStageSetTable_Wait_NoteStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Accord_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Accord",
	}).Stage(probe.formStage)
	inst := new(direction.Accord)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Accord_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Accord_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AccordionRegistration_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AccordionRegistration",
	}).Stage(probe.formStage)
	inst := new(direction.AccordionRegistration)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_AccordionRegistration_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_AccordionRegistration_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Barre_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Barre",
	}).Stage(probe.formStage)
	inst := new(direction.Barre)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Barre_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Barre_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bass_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Bass",
	}).Stage(probe.formStage)
	inst := new(direction.Bass)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Bass_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Bass_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BassStep_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BassStep",
	}).Stage(probe.formStage)
	inst := new(direction.BassStep)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_BassStep_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_BassStep_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BeatUnitTied_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BeatUnitTied",
	}).Stage(probe.formStage)
	inst := new(direction.BeatUnitTied)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_BeatUnitTied_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_BeatUnitTied_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Beater_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Beater",
	}).Stage(probe.formStage)
	inst := new(direction.Beater)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Beater_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Beater_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bracket_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Bracket",
	}).Stage(probe.formStage)
	inst := new(direction.Bracket)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Bracket_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Bracket_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Dashes_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Dashes",
	}).Stage(probe.formStage)
	inst := new(direction.Dashes)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Dashes_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Dashes_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Degree_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Degree",
	}).Stage(probe.formStage)
	inst := new(direction.Degree)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Degree_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Degree_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DegreeAlter_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DegreeAlter",
	}).Stage(probe.formStage)
	inst := new(direction.DegreeAlter)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_DegreeAlter_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_DegreeAlter_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DegreeType_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DegreeType",
	}).Stage(probe.formStage)
	inst := new(direction.DegreeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_DegreeType_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_DegreeType_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DegreeValue_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DegreeValue",
	}).Stage(probe.formStage)
	inst := new(direction.DegreeValue)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_DegreeValue_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_DegreeValue_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Direction_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Direction",
	}).Stage(probe.formStage)
	inst := new(direction.Direction)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Direction_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Direction_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectionType_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DirectionType",
	}).Stage(probe.formStage)
	inst := new(direction.DirectionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_DirectionType_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_DirectionType_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Effect_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Effect",
	}).Stage(probe.formStage)
	inst := new(direction.Effect)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Effect_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Effect_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Feature_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Feature",
	}).Stage(probe.formStage)
	inst := new(direction.Feature)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Feature_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Feature_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FirstFret_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FirstFret",
	}).Stage(probe.formStage)
	inst := new(direction.FirstFret)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_FirstFret_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_FirstFret_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Frame_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Frame",
	}).Stage(probe.formStage)
	inst := new(direction.Frame)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Frame_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Frame_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FrameNote_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FrameNote",
	}).Stage(probe.formStage)
	inst := new(direction.FrameNote)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_FrameNote_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_FrameNote_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Glass_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Glass",
	}).Stage(probe.formStage)
	inst := new(direction.Glass)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Glass_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Glass_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Grouping_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Grouping",
	}).Stage(probe.formStage)
	inst := new(direction.Grouping)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Grouping_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Grouping_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Harmony_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Harmony",
	}).Stage(probe.formStage)
	inst := new(direction.Harmony)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Harmony_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Harmony_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HarmonyAlter_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HarmonyAlter",
	}).Stage(probe.formStage)
	inst := new(direction.HarmonyAlter)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_HarmonyAlter_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_HarmonyAlter_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HarpPedals_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HarpPedals",
	}).Stage(probe.formStage)
	inst := new(direction.HarpPedals)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_HarpPedals_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_HarpPedals_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Image_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Image",
	}).Stage(probe.formStage)
	inst := new(direction.Image)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Image_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Image_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InstrumentChange_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New InstrumentChange",
	}).Stage(probe.formStage)
	inst := new(direction.InstrumentChange)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_InstrumentChange_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_InstrumentChange_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Inversion_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Inversion",
	}).Stage(probe.formStage)
	inst := new(direction.Inversion)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Inversion_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Inversion_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Kind_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Kind",
	}).Stage(probe.formStage)
	inst := new(direction.Kind)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Kind_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Kind_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Listening_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Listening",
	}).Stage(probe.formStage)
	inst := new(direction.Listening)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Listening_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Listening_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeasureNumbering_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MeasureNumbering",
	}).Stage(probe.formStage)
	inst := new(direction.MeasureNumbering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_MeasureNumbering_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_MeasureNumbering_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Membrane_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Membrane",
	}).Stage(probe.formStage)
	inst := new(direction.Membrane)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Membrane_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Membrane_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metal_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Metal",
	}).Stage(probe.formStage)
	inst := new(direction.Metal)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Metal_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Metal_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Metronome_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Metronome",
	}).Stage(probe.formStage)
	inst := new(direction.Metronome)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Metronome_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Metronome_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MetronomeBeam_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MetronomeBeam",
	}).Stage(probe.formStage)
	inst := new(direction.MetronomeBeam)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_MetronomeBeam_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_MetronomeBeam_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MetronomeNote_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MetronomeNote",
	}).Stage(probe.formStage)
	inst := new(direction.MetronomeNote)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_MetronomeNote_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_MetronomeNote_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MetronomeTied_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MetronomeTied",
	}).Stage(probe.formStage)
	inst := new(direction.MetronomeTied)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_MetronomeTied_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_MetronomeTied_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MetronomeTuplet_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MetronomeTuplet",
	}).Stage(probe.formStage)
	inst := new(direction.MetronomeTuplet)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_MetronomeTuplet_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_MetronomeTuplet_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Numeral_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Numeral",
	}).Stage(probe.formStage)
	inst := new(direction.Numeral)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Numeral_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Numeral_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NumeralKey_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NumeralKey",
	}).Stage(probe.formStage)
	inst := new(direction.NumeralKey)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_NumeralKey_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_NumeralKey_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NumeralRoot_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NumeralRoot",
	}).Stage(probe.formStage)
	inst := new(direction.NumeralRoot)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_NumeralRoot_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_NumeralRoot_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OctaveShift_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OctaveShift",
	}).Stage(probe.formStage)
	inst := new(direction.OctaveShift)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_OctaveShift_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_OctaveShift_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Offset_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Offset",
	}).Stage(probe.formStage)
	inst := new(direction.Offset)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Offset_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Offset_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherDirection_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherDirection",
	}).Stage(probe.formStage)
	inst := new(direction.OtherDirection)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_OtherDirection_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_OtherDirection_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherListening_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherListening",
	}).Stage(probe.formStage)
	inst := new(direction.OtherListening)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_OtherListening_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_OtherListening_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pedal_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Pedal",
	}).Stage(probe.formStage)
	inst := new(direction.Pedal)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Pedal_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Pedal_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PedalTuning_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PedalTuning",
	}).Stage(probe.formStage)
	inst := new(direction.PedalTuning)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_PedalTuning_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_PedalTuning_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PerMinute_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PerMinute",
	}).Stage(probe.formStage)
	inst := new(direction.PerMinute)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_PerMinute_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_PerMinute_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Percussion_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Percussion",
	}).Stage(probe.formStage)
	inst := new(direction.Percussion)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Percussion_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Percussion_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Pitched_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Pitched",
	}).Stage(probe.formStage)
	inst := new(direction.Pitched)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Pitched_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Pitched_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrincipalVoice_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PrincipalVoice",
	}).Stage(probe.formStage)
	inst := new(direction.PrincipalVoice)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_PrincipalVoice_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_PrincipalVoice_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Print_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Print",
	}).Stage(probe.formStage)
	inst := new(direction.Print)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Print_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Print_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Root_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Root",
	}).Stage(probe.formStage)
	inst := new(direction.Root)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Root_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Root_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RootStep_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New RootStep",
	}).Stage(probe.formStage)
	inst := new(direction.RootStep)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_RootStep_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_RootStep_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Scordatura_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Scordatura",
	}).Stage(probe.formStage)
	inst := new(direction.Scordatura)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Scordatura_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Scordatura_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Sound_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Sound",
	}).Stage(probe.formStage)
	inst := new(direction.Sound)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Sound_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Sound_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StaffDivide_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaffDivide",
	}).Stage(probe.formStage)
	inst := new(direction.StaffDivide)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_StaffDivide_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_StaffDivide_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Stick_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Stick",
	}).Stage(probe.formStage)
	inst := new(direction.Stick)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Stick_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Stick_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StringMute_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StringMute",
	}).Stage(probe.formStage)
	inst := new(direction.StringMute)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_StringMute_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_StringMute_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StyleText_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StyleText",
	}).Stage(probe.formStage)
	inst := new(direction.StyleText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_StyleText_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_StyleText_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Swing_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Swing",
	}).Stage(probe.formStage)
	inst := new(direction.Swing)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Swing_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Swing_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Sync_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Sync",
	}).Stage(probe.formStage)
	inst := new(direction.Sync)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Sync_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Sync_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimeModification_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TimeModification",
	}).Stage(probe.formStage)
	inst := new(direction.TimeModification)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_TimeModification_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_TimeModification_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Timpani_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Timpani",
	}).Stage(probe.formStage)
	inst := new(direction.Timpani)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Timpani_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Timpani_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VirtualInstrument_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New VirtualInstrument",
	}).Stage(probe.formStage)
	inst := new(direction.VirtualInstrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_VirtualInstrument_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_VirtualInstrument_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wedge_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Wedge",
	}).Stage(probe.formStage)
	inst := new(direction.Wedge)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Wedge_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Wedge_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Wood_DirectionStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Wood",
	}).Stage(probe.formStage)
	inst := new(direction.Wood)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.DirectionStage.Lock()
			defer probe.stageSet.DirectionStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.DirectionStage)
			saveStageSet_Wood_DirectionStage(inst, probe, formGroup)
			probe.stageSet.DirectionStage.Commit()
			updateStageSetTable_Wood_DirectionStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BarStyleColor_BarlineStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BarStyleColor",
	}).Stage(probe.formStage)
	inst := new(barline.BarStyleColor)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.BarlineStage.Lock()
			defer probe.stageSet.BarlineStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.BarlineStage)
			saveStageSet_BarStyleColor_BarlineStage(inst, probe, formGroup)
			probe.stageSet.BarlineStage.Commit()
			updateStageSetTable_BarStyleColor_BarlineStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Barline_BarlineStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Barline",
	}).Stage(probe.formStage)
	inst := new(barline.Barline)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.BarlineStage.Lock()
			defer probe.stageSet.BarlineStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.BarlineStage)
			saveStageSet_Barline_BarlineStage(inst, probe, formGroup)
			probe.stageSet.BarlineStage.Commit()
			updateStageSetTable_Barline_BarlineStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Ending_BarlineStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Ending",
	}).Stage(probe.formStage)
	inst := new(barline.Ending)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.BarlineStage.Lock()
			defer probe.stageSet.BarlineStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.BarlineStage)
			saveStageSet_Ending_BarlineStage(inst, probe, formGroup)
			probe.stageSet.BarlineStage.Commit()
			updateStageSetTable_Ending_BarlineStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Repeat_BarlineStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Repeat",
	}).Stage(probe.formStage)
	inst := new(barline.Repeat)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.BarlineStage.Lock()
			defer probe.stageSet.BarlineStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.BarlineStage)
			saveStageSet_Repeat_BarlineStage(inst, probe, formGroup)
			probe.stageSet.BarlineStage.Commit()
			updateStageSetTable_Repeat_BarlineStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Attributes_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Attributes",
	}).Stage(probe.formStage)
	inst := new(attributes.Attributes)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Attributes_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Attributes_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AttributesDirective_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AttributesDirective",
	}).Stage(probe.formStage)
	inst := new(attributes.AttributesDirective)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_AttributesDirective_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_AttributesDirective_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BeatRepeat_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BeatRepeat",
	}).Stage(probe.formStage)
	inst := new(attributes.BeatRepeat)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_BeatRepeat_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_BeatRepeat_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Cancel_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Cancel",
	}).Stage(probe.formStage)
	inst := new(attributes.Cancel)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Cancel_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Cancel_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Clef_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Clef",
	}).Stage(probe.formStage)
	inst := new(attributes.Clef)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Clef_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Clef_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Double_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Double",
	}).Stage(probe.formStage)
	inst := new(attributes.Double)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Double_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Double_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ForPart_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ForPart",
	}).Stage(probe.formStage)
	inst := new(attributes.ForPart)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_ForPart_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_ForPart_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Interchangeable_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Interchangeable",
	}).Stage(probe.formStage)
	inst := new(attributes.Interchangeable)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Interchangeable_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Interchangeable_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Key_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Key",
	}).Stage(probe.formStage)
	inst := new(attributes.Key)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Key_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Key_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_KeyAccidental_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New KeyAccidental",
	}).Stage(probe.formStage)
	inst := new(attributes.KeyAccidental)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_KeyAccidental_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_KeyAccidental_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_KeyOctave_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New KeyOctave",
	}).Stage(probe.formStage)
	inst := new(attributes.KeyOctave)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_KeyOctave_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_KeyOctave_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LineDetail_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LineDetail",
	}).Stage(probe.formStage)
	inst := new(attributes.LineDetail)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_LineDetail_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_LineDetail_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeasureRepeat_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MeasureRepeat",
	}).Stage(probe.formStage)
	inst := new(attributes.MeasureRepeat)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_MeasureRepeat_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_MeasureRepeat_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeasureStyle_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MeasureStyle",
	}).Stage(probe.formStage)
	inst := new(attributes.MeasureStyle)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_MeasureStyle_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_MeasureStyle_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MultipleRest_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MultipleRest",
	}).Stage(probe.formStage)
	inst := new(attributes.MultipleRest)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_MultipleRest_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_MultipleRest_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartClef_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartClef",
	}).Stage(probe.formStage)
	inst := new(attributes.PartClef)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_PartClef_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_PartClef_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartSymbol_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartSymbol",
	}).Stage(probe.formStage)
	inst := new(attributes.PartSymbol)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_PartSymbol_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_PartSymbol_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PartTranspose_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PartTranspose",
	}).Stage(probe.formStage)
	inst := new(attributes.PartTranspose)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_PartTranspose_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_PartTranspose_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Slash_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Slash",
	}).Stage(probe.formStage)
	inst := new(attributes.Slash)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Slash_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Slash_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StaffDetails_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaffDetails",
	}).Stage(probe.formStage)
	inst := new(attributes.StaffDetails)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_StaffDetails_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_StaffDetails_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StaffSize_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaffSize",
	}).Stage(probe.formStage)
	inst := new(attributes.StaffSize)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_StaffSize_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_StaffSize_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StaffTuning_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaffTuning",
	}).Stage(probe.formStage)
	inst := new(attributes.StaffTuning)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_StaffTuning_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_StaffTuning_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Time_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Time",
	}).Stage(probe.formStage)
	inst := new(attributes.Time)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Time_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Time_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Transpose_AttributesStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Transpose",
	}).Stage(probe.formStage)
	inst := new(attributes.Transpose)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.AttributesStage.Lock()
			defer probe.stageSet.AttributesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.AttributesStage)
			saveStageSet_Transpose_AttributesStage(inst, probe, formGroup)
			probe.stageSet.AttributesStage.Commit()
			updateStageSetTable_Transpose_AttributesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AccidentalText_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AccidentalText",
	}).Stage(probe.formStage)
	inst := new(common.AccidentalText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_AccidentalText_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_AccidentalText_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Appearance_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Appearance",
	}).Stage(probe.formStage)
	inst := new(common.Appearance)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Appearance_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Appearance_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Bookmark_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Bookmark",
	}).Stage(probe.formStage)
	inst := new(common.Bookmark)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Bookmark_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Bookmark_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Coda_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Coda",
	}).Stage(probe.formStage)
	inst := new(common.Coda)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Coda_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Coda_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Distance_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Distance",
	}).Stage(probe.formStage)
	inst := new(common.Distance)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Distance_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Distance_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Dynamics_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Dynamics",
	}).Stage(probe.formStage)
	inst := new(common.Dynamics)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Dynamics_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Dynamics_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Empty_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Empty",
	}).Stage(probe.formStage)
	inst := new(common.Empty)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Empty_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Empty_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyPlacement_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyPlacement",
	}).Stage(probe.formStage)
	inst := new(common.EmptyPlacement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyPlacement_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyPlacement_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyPlacementSmufl_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyPlacementSmufl",
	}).Stage(probe.formStage)
	inst := new(common.EmptyPlacementSmufl)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyPlacementSmufl_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyPlacementSmufl_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyPrintObjectStyleAlign_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyPrintObjectStyleAlign",
	}).Stage(probe.formStage)
	inst := new(common.EmptyPrintObjectStyleAlign)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyPrintObjectStyleAlign_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyPrintObjectStyleAlign_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyPrintStyle_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyPrintStyle",
	}).Stage(probe.formStage)
	inst := new(common.EmptyPrintStyle)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyPrintStyle_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyPrintStyle_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyPrintStyleAlign_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyPrintStyleAlign",
	}).Stage(probe.formStage)
	inst := new(common.EmptyPrintStyleAlign)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyPrintStyleAlign_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyPrintStyleAlign_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyPrintStyleAlignId_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyPrintStyleAlignId",
	}).Stage(probe.formStage)
	inst := new(common.EmptyPrintStyleAlignId)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyPrintStyleAlignId_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyPrintStyleAlignId_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EmptyTrillSound_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New EmptyTrillSound",
	}).Stage(probe.formStage)
	inst := new(common.EmptyTrillSound)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_EmptyTrillSound_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_EmptyTrillSound_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Encoding_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Encoding",
	}).Stage(probe.formStage)
	inst := new(common.Encoding)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Encoding_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Encoding_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Fermata_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Fermata",
	}).Stage(probe.formStage)
	inst := new(common.Fermata)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Fermata_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Fermata_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Fingering_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Fingering",
	}).Stage(probe.formStage)
	inst := new(common.Fingering)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Fingering_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Fingering_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FormattedSymbol_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FormattedSymbol",
	}).Stage(probe.formStage)
	inst := new(common.FormattedSymbol)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_FormattedSymbol_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_FormattedSymbol_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FormattedSymbolId_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FormattedSymbolId",
	}).Stage(probe.formStage)
	inst := new(common.FormattedSymbolId)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_FormattedSymbolId_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_FormattedSymbolId_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FormattedText_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FormattedText",
	}).Stage(probe.formStage)
	inst := new(common.FormattedText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_FormattedText_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_FormattedText_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FormattedTextId_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New FormattedTextId",
	}).Stage(probe.formStage)
	inst := new(common.FormattedTextId)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_FormattedTextId_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_FormattedTextId_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Fret_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Fret",
	}).Stage(probe.formStage)
	inst := new(common.Fret)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Fret_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Fret_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Glyph_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Glyph",
	}).Stage(probe.formStage)
	inst := new(common.Glyph)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Glyph_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Glyph_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HorizontalTurn_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New HorizontalTurn",
	}).Stage(probe.formStage)
	inst := new(common.HorizontalTurn)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_HorizontalTurn_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_HorizontalTurn_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Identification_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Identification",
	}).Stage(probe.formStage)
	inst := new(common.Identification)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Identification_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Identification_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Level_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Level",
	}).Stage(probe.formStage)
	inst := new(common.Level)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Level_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Level_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LineWidth_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LineWidth",
	}).Stage(probe.formStage)
	inst := new(common.LineWidth)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_LineWidth_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_LineWidth_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Link_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Link",
	}).Stage(probe.formStage)
	inst := new(common.Link)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Link_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Link_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeasureLayout_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MeasureLayout",
	}).Stage(probe.formStage)
	inst := new(common.MeasureLayout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_MeasureLayout_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_MeasureLayout_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MidiDevice_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MidiDevice",
	}).Stage(probe.formStage)
	inst := new(common.MidiDevice)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_MidiDevice_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_MidiDevice_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MidiInstrument_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MidiInstrument",
	}).Stage(probe.formStage)
	inst := new(common.MidiInstrument)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_MidiInstrument_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_MidiInstrument_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Miscellaneous_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Miscellaneous",
	}).Stage(probe.formStage)
	inst := new(common.Miscellaneous)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Miscellaneous_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Miscellaneous_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MiscellaneousField_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MiscellaneousField",
	}).Stage(probe.formStage)
	inst := new(common.MiscellaneousField)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_MiscellaneousField_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_MiscellaneousField_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NameDisplay_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NameDisplay",
	}).Stage(probe.formStage)
	inst := new(common.NameDisplay)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_NameDisplay_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_NameDisplay_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteSize_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteSize",
	}).Stage(probe.formStage)
	inst := new(common.NoteSize)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_NoteSize_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_NoteSize_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherAppearance_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherAppearance",
	}).Stage(probe.formStage)
	inst := new(common.OtherAppearance)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_OtherAppearance_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_OtherAppearance_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherPlay_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherPlay",
	}).Stage(probe.formStage)
	inst := new(common.OtherPlay)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_OtherPlay_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_OtherPlay_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OtherText_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OtherText",
	}).Stage(probe.formStage)
	inst := new(common.OtherText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_OtherText_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_OtherText_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PageLayout_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PageLayout",
	}).Stage(probe.formStage)
	inst := new(common.PageLayout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_PageLayout_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_PageLayout_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PageMargins_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PageMargins",
	}).Stage(probe.formStage)
	inst := new(common.PageMargins)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_PageMargins_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_PageMargins_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Play_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Play",
	}).Stage(probe.formStage)
	inst := new(common.Play)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Play_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Play_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Scaling_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Scaling",
	}).Stage(probe.formStage)
	inst := new(common.Scaling)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Scaling_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Scaling_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Segno_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Segno",
	}).Stage(probe.formStage)
	inst := new(common.Segno)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Segno_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Segno_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StaffLayout_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StaffLayout",
	}).Stage(probe.formStage)
	inst := new(common.StaffLayout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_StaffLayout_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_StaffLayout_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StringType_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StringType",
	}).Stage(probe.formStage)
	inst := new(common.StringType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_StringType_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_StringType_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Supports_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Supports",
	}).Stage(probe.formStage)
	inst := new(common.Supports)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_Supports_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_Supports_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SystemDividers_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SystemDividers",
	}).Stage(probe.formStage)
	inst := new(common.SystemDividers)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_SystemDividers_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_SystemDividers_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SystemLayout_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SystemLayout",
	}).Stage(probe.formStage)
	inst := new(common.SystemLayout)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_SystemLayout_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_SystemLayout_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SystemMargins_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SystemMargins",
	}).Stage(probe.formStage)
	inst := new(common.SystemMargins)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_SystemMargins_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_SystemMargins_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TypedText_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TypedText",
	}).Stage(probe.formStage)
	inst := new(common.TypedText)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_TypedText_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_TypedText_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_WavyLine_CommonStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New WavyLine",
	}).Stage(probe.formStage)
	inst := new(common.WavyLine)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.CommonStage.Lock()
			defer probe.stageSet.CommonStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.CommonStage)
			saveStageSet_WavyLine_CommonStage(inst, probe, formGroup)
			probe.stageSet.CommonStage.Commit()
			updateStageSetTable_WavyLine_CommonStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArcType_XlinkStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ArcType",
	}).Stage(probe.formStage)
	inst := new(xlink.ArcType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XlinkStage.Lock()
			defer probe.stageSet.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XlinkStage)
			saveStageSet_ArcType_XlinkStage(inst, probe, formGroup)
			probe.stageSet.XlinkStage.Commit()
			updateStageSetTable_ArcType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Extended_XlinkStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Extended",
	}).Stage(probe.formStage)
	inst := new(xlink.Extended)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XlinkStage.Lock()
			defer probe.stageSet.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XlinkStage)
			saveStageSet_Extended_XlinkStage(inst, probe, formGroup)
			probe.stageSet.XlinkStage.Commit()
			updateStageSetTable_Extended_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LocatorType_XlinkStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New LocatorType",
	}).Stage(probe.formStage)
	inst := new(xlink.LocatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XlinkStage.Lock()
			defer probe.stageSet.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XlinkStage)
			saveStageSet_LocatorType_XlinkStage(inst, probe, formGroup)
			probe.stageSet.XlinkStage.Commit()
			updateStageSetTable_LocatorType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ResourceType_XlinkStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ResourceType",
	}).Stage(probe.formStage)
	inst := new(xlink.ResourceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XlinkStage.Lock()
			defer probe.stageSet.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XlinkStage)
			saveStageSet_ResourceType_XlinkStage(inst, probe, formGroup)
			probe.stageSet.XlinkStage.Commit()
			updateStageSetTable_ResourceType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Simple_XlinkStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Simple",
	}).Stage(probe.formStage)
	inst := new(xlink.Simple)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XlinkStage.Lock()
			defer probe.stageSet.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XlinkStage)
			saveStageSet_Simple_XlinkStage(inst, probe, formGroup)
			probe.stageSet.XlinkStage.Commit()
			updateStageSetTable_Simple_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TitleEltType_XlinkStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TitleEltType",
	}).Stage(probe.formStage)
	inst := new(xlink.TitleEltType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.XlinkStage.Lock()
			defer probe.stageSet.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.XlinkStage)
			saveStageSet_TitleEltType_XlinkStage(inst, probe, formGroup)
			probe.stageSet.XlinkStage.Commit()
			updateStageSetTable_TitleEltType_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}


// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"
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
	case *models.Angle0Shape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Angle0Shape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Angle0Shape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.BottomCurvePlane1Shape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BottomCurvePlane1Shape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_BottomCurvePlane1Shape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.BottomCurvePlane2Shape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_BottomCurvePlane2Shape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_BottomCurvePlane2Shape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Circumference3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Circumference3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Circumference3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Clock2DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Clock2DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Clock2DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Clock3DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Clock3DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Clock3DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.CutLine3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_CutLine3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_CutLine3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Leaves3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Leaves3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Leaves3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Library:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Library_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Library_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.OriginalPoints3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_OriginalPoints3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_OriginalPoints3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ParastichyMCurves3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ParastichyMCurves3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ParastichyMCurves3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ParastichyNCurves3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ParastichyNCurves3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ParastichyNCurves3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Plant2DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Plant2DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Plant2DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Plant3DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Plant3DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Plant3DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.PlantAbstract:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_PlantAbstract_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_PlantAbstract_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Rendered3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Rendered3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Rendered3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.SampledPoints3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_SampledPoints3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_SampledPoints3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StackOfRotatedVaseTrapezeRingsShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StackOfRotatedVaseTrapezeRingsShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StackOfRotatedVaseTrapezeRingsShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StackOfVaseTrapezeRingsShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StackOfVaseTrapezeRingsShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StackOfVaseTrapezeRingsShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.StemCylinder3DShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StemCylinder3DShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_StemCylinder3DShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Stool2DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stool2DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Stool2DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Stool3DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Stool3DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Stool3DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TopCurvePlane1Shape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TopCurvePlane1Shape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TopCurvePlane1Shape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TopCurvePlane2Shape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TopCurvePlane2Shape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TopCurvePlane2Shape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TubeVase3DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TubeVase3DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TubeVase3DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TubeVaseAbstract:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TubeVaseAbstract_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TubeVaseAbstract_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Vase2DDiagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Vase2DDiagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Vase2DDiagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.VaseTrapezeRingShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_VaseTrapezeRingShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_VaseTrapezeRingShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *stool.StoolAbstract:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.StoolStage.Lock()
				defer probe.stageSet.StoolStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_StoolAbstract_StoolStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.StoolStage)
				}
				probe.stageSet.StoolStage.Commit()
				updateStageSetTable_StoolAbstract_StoolStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *music.MusicAbstract:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.MusicStage.Lock()
				defer probe.stageSet.MusicStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MusicAbstract_MusicStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.MusicStage)
				}
				probe.stageSet.MusicStage.Commit()
				updateStageSetTable_MusicAbstract_MusicStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *clock.ClockAbstract:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.ClockStage.Lock()
				defer probe.stageSet.ClockStage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ClockAbstract_ClockStage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.ClockStage)
				}
				probe.stageSet.ClockStage.Commit()
				updateStageSetTable_ClockAbstract_ClockStage(probe)
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

func saveStageSet_Angle0Shape_Stage(
	inst *models.Angle0Shape,
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

func saveStageSet_BottomCurvePlane1Shape_Stage(
	inst *models.BottomCurvePlane1Shape,
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

func saveStageSet_BottomCurvePlane2Shape_Stage(
	inst *models.BottomCurvePlane2Shape,
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

func saveStageSet_Circumference3DShape_Stage(
	inst *models.Circumference3DShape,
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

func saveStageSet_Clock2DDiagram_Stage(
	inst *models.Clock2DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&inst.Zoom, formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&inst.IsHiddenAxesShape, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_Clock3DDiagram_Stage(
	inst *models.Clock3DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsHiddenClockTopCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenClockTopCurveShape, formDiv)
		case "ClockTopCurveShape":
			StageSetFormDivSelectFieldToField(&inst.ClockTopCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.ClockTopCurveShape](), formDiv)
		case "IsHiddenTorus3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTorus3DShape, formDiv)
		case "Torus3DShape":
			StageSetFormDivSelectFieldToField(&inst.Torus3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Torus3DShape](), formDiv)
		case "IsHiddenSampledPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenSampledPoints3DShape, formDiv)
		case "SampledPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.SampledPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.SampledPoints3DShape](), formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTiledFloor3DShape, formDiv)
		case "TiledFloor3DShape":
			StageSetFormDivSelectFieldToField(&inst.TiledFloor3DShape, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), formDiv)
		case "Rendered3DShape":
			StageSetFormDivSelectFieldToField(&inst.Rendered3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_CutLine3DShape_Stage(
	inst *models.CutLine3DShape,
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

func saveStageSet_Leaves3DShape_Stage(
	inst *models.Leaves3DShape,
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

func saveStageSet_Library_Stage(
	inst *models.Library,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&inst.NbPixPerCharacter, formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&inst.LogoSVGFile, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&inst.IsRootLibrary, formDiv)
		}
	}
}

func saveStageSet_OriginalPoints3DShape_Stage(
	inst *models.OriginalPoints3DShape,
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

func saveStageSet_ParastichyMCurves3DShape_Stage(
	inst *models.ParastichyMCurves3DShape,
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

func saveStageSet_ParastichyNCurves3DShape_Stage(
	inst *models.ParastichyNCurves3DShape,
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

func saveStageSet_Plant2DDiagram_Stage(
	inst *models.Plant2DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&inst.OriginX, formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&inst.OriginY, formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&inst.Zoom, formDiv)
		case "IsRhombusNodesExpanded":
			FormDivBasicFieldToField(&inst.IsRhombusNodesExpanded, formDiv)
		case "IsArcNodesExpanded":
			FormDivBasicFieldToField(&inst.IsArcNodesExpanded, formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&inst.IsHiddenAxesShape, formDiv)
		case "IsHiddenReferenceRhombus":
			FormDivBasicFieldToField(&inst.IsHiddenReferenceRhombus, formDiv)
		case "IsHiddenPlantCircumferenceShape":
			FormDivBasicFieldToField(&inst.IsHiddenPlantCircumferenceShape, formDiv)
		case "IsHiddenGridPathShape":
			FormDivBasicFieldToField(&inst.IsHiddenGridPathShape, formDiv)
		case "IsHiddenRhombusGridShape":
			FormDivBasicFieldToField(&inst.IsHiddenRhombusGridShape, formDiv)
		case "IsHiddenExplanationTextShape":
			FormDivBasicFieldToField(&inst.IsHiddenExplanationTextShape, formDiv)
		case "IsHiddenRotatedReferenceRhombus":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedReferenceRhombus, formDiv)
		case "IsHiddenRotatedPlantCircumferenceShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedPlantCircumferenceShape, formDiv)
		case "IsHiddenRotatedGridPathShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedGridPathShape, formDiv)
		case "IsHiddenRotatedRhombusGridShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedRhombusGridShape, formDiv)
		case "IsHiddenGrowthPathRhombusGridShape":
			FormDivBasicFieldToField(&inst.IsHiddenGrowthPathRhombusGridShape, formDiv)
		case "IsHiddenGrowthVectorShape":
			FormDivBasicFieldToField(&inst.IsHiddenGrowthVectorShape, formDiv)
		case "IsHiddenPerpendicularVectorGrid":
			FormDivBasicFieldToField(&inst.IsHiddenPerpendicularVectorGrid, formDiv)
		case "IsHiddenBaseVectorShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenBaseVectorShapeGrid, formDiv)
		case "IsHiddenArcNormalVectorShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenArcNormalVectorShapeGrid, formDiv)
		case "IsHiddenStartArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenStartArcShapeGrid, formDiv)
		case "IsHiddenMidArcVectorShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenMidArcVectorShapeGrid, formDiv)
		case "IsHiddenEndArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenEndArcShapeGrid, formDiv)
		case "IsHiddenGrowthCurve2D":
			FormDivBasicFieldToField(&inst.IsHiddenGrowthCurve2D, formDiv)
		case "IsHiddenStackOfGrowthCurve2DByGrowthVector":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfGrowthCurve2DByGrowthVector, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_Plant3DDiagram_Stage(
	inst *models.Plant3DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsHiddenStemCylinder3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenStemCylinder3DShape, formDiv)
		case "StemCylinder3DShape":
			StageSetFormDivSelectFieldToField(&inst.StemCylinder3DShape, probe.stageSet.Stage.GetInstancesSet[*models.StemCylinder3DShape](), formDiv)
		case "IsHiddenParastichyNCurves3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenParastichyNCurves3DShape, formDiv)
		case "ParastichyNCurves3DShape":
			StageSetFormDivSelectFieldToField(&inst.ParastichyNCurves3DShape, probe.stageSet.Stage.GetInstancesSet[*models.ParastichyNCurves3DShape](), formDiv)
		case "IsHiddenParastichyMCurves3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenParastichyMCurves3DShape, formDiv)
		case "ParastichyMCurves3DShape":
			StageSetFormDivSelectFieldToField(&inst.ParastichyMCurves3DShape, probe.stageSet.Stage.GetInstancesSet[*models.ParastichyMCurves3DShape](), formDiv)
		case "IsHiddenCutLine3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenCutLine3DShape, formDiv)
		case "CutLine3DShape":
			StageSetFormDivSelectFieldToField(&inst.CutLine3DShape, probe.stageSet.Stage.GetInstancesSet[*models.CutLine3DShape](), formDiv)
		case "IsHiddenCircumference3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenCircumference3DShape, formDiv)
		case "Circumference3DShape":
			StageSetFormDivSelectFieldToField(&inst.Circumference3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Circumference3DShape](), formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTiledFloor3DShape, formDiv)
		case "TiledFloor3DShape":
			StageSetFormDivSelectFieldToField(&inst.TiledFloor3DShape, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), formDiv)
		case "IsHiddenLeaves3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenLeaves3DShape, formDiv)
		case "Leaves3DShape":
			StageSetFormDivSelectFieldToField(&inst.Leaves3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Leaves3DShape](), formDiv)
		case "Rendered3DShape":
			StageSetFormDivSelectFieldToField(&inst.Rendered3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_PlantAbstract_Stage(
	inst *models.PlantAbstract,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "N":
			FormDivBasicFieldToField(&inst.N, formDiv)
		case "M":
			FormDivBasicFieldToField(&inst.M, formDiv)
		case "StackHeight":
			FormDivBasicFieldToField(&inst.StackHeight, formDiv)
		case "RhombusInsideAngle":
			FormDivBasicFieldToField(&inst.RhombusInsideAngle, formDiv)
		case "RhombusSideLength":
			FormDivBasicFieldToField(&inst.RhombusSideLength, formDiv)
		case "PlantType":
			StageSetFormDivEnumStringFieldToField(&inst.PlantType, formDiv)
		case "TubeVaseAbstract":
			StageSetFormDivSelectFieldToField(&inst.TubeVaseAbstract, probe.stageSet.Stage.GetInstancesSet[*models.TubeVaseAbstract](), formDiv)
		case "StoolAbstract":
			StageSetFormDivSelectFieldToField(&inst.StoolAbstract, probe.stageSet.StoolStage.GetInstancesSet[*stool.StoolAbstract](), formDiv)
		case "ClockAbstract":
			StageSetFormDivSelectFieldToField(&inst.ClockAbstract, probe.stageSet.ClockStage.GetInstancesSet[*clock.ClockAbstract](), formDiv)
		case "MusicAbstract":
			StageSetFormDivSelectFieldToField(&inst.MusicAbstract, probe.stageSet.MusicStage.GetInstancesSet[*music.MusicAbstract](), formDiv)
		case "CurrentView":
			StageSetFormDivEnumStringFieldToField(&inst.CurrentView, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&inst.IsSelected, formDiv)
		case "IsPlant2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsPlant2DDiagramsNodeExpanded, formDiv)
		case "IsPlant3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsPlant3DDiagramsNodeExpanded, formDiv)
		case "IsVase2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsVase2DDiagramsNodeExpanded, formDiv)
		case "IsTubeVase3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsTubeVase3DDiagramsNodeExpanded, formDiv)
		case "IsStool2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsStool2DDiagramsNodeExpanded, formDiv)
		case "IsStool3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsStool3DDiagramsNodeExpanded, formDiv)
		case "IsClock2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsClock2DDiagramsNodeExpanded, formDiv)
		case "IsClock3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsClock3DDiagramsNodeExpanded, formDiv)
		case "AxesShape":
			StageSetFormDivSelectFieldToField(&inst.AxesShape, probe.stageSet.Stage.GetInstancesSet[*models.AxesShape](), formDiv)
		case "RhombusStuff":
			StageSetFormDivSelectFieldToField(&inst.RhombusStuff, probe.stageSet.Stage.GetInstancesSet[*models.RhombusStuff](), formDiv)
		case "GrowthVectorShape":
			StageSetFormDivSelectFieldToField(&inst.GrowthVectorShape, probe.stageSet.Stage.GetInstancesSet[*models.GrowthVectorShape](), formDiv)
		case "PerpendicularVectorGrid":
			StageSetFormDivSelectFieldToField(&inst.PerpendicularVectorGrid, probe.stageSet.Stage.GetInstancesSet[*models.PerpendicularVectorGrid](), formDiv)
		case "BaseVectorShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.BaseVectorShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.BaseVectorShapeGrid](), formDiv)
		case "ArcNormalVectorShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.ArcNormalVectorShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.ArcNormalVectorShapeGrid](), formDiv)
		case "StartArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.StartArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.StartArcShapeGrid](), formDiv)
		case "MidArcVectorShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.MidArcVectorShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.MidArcVectorShapeGrid](), formDiv)
		case "EndArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.EndArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.EndArcShapeGrid](), formDiv)
		case "GrowthCurve2D":
			StageSetFormDivSelectFieldToField(&inst.GrowthCurve2D, probe.stageSet.Stage.GetInstancesSet[*models.GrowthCurve2D](), formDiv)
		case "StackOfGrowthCurve2DByGrowthVector":
			StageSetFormDivSelectFieldToField(&inst.StackOfGrowthCurve2DByGrowthVector, probe.stageSet.Stage.GetInstancesSet[*models.StackOfGrowthCurve2DByGrowthVector](), formDiv)
		}
	}
}

func saveStageSet_Rendered3DShape_Stage(
	inst *models.Rendered3DShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ViewX":
			FormDivBasicFieldToField(&inst.ViewX, formDiv)
		case "ViewY":
			FormDivBasicFieldToField(&inst.ViewY, formDiv)
		case "ViewZ":
			FormDivBasicFieldToField(&inst.ViewZ, formDiv)
		case "TargetX":
			FormDivBasicFieldToField(&inst.TargetX, formDiv)
		case "TargetY":
			FormDivBasicFieldToField(&inst.TargetY, formDiv)
		case "TargetZ":
			FormDivBasicFieldToField(&inst.TargetZ, formDiv)
		case "Fov":
			FormDivBasicFieldToField(&inst.Fov, formDiv)
		}
	}
}

func saveStageSet_SampledPoints3DShape_Stage(
	inst *models.SampledPoints3DShape,
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

func saveStageSet_StackOfRotatedVaseTrapezeRingsShape_Stage(
	inst *models.StackOfRotatedVaseTrapezeRingsShape,
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

func saveStageSet_StackOfVaseTrapezeRingsShape_Stage(
	inst *models.StackOfVaseTrapezeRingsShape,
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

func saveStageSet_StemCylinder3DShape_Stage(
	inst *models.StemCylinder3DShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&inst.Transparency, formDiv)
		}
	}
}

func saveStageSet_Stool2DDiagram_Stage(
	inst *models.Stool2DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&inst.Zoom, formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&inst.IsHiddenAxesShape, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_Stool3DDiagram_Stage(
	inst *models.Stool3DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsHiddenSeatTopCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenSeatTopCurveShape, formDiv)
		case "SeatTopCurveShape":
			StageSetFormDivSelectFieldToField(&inst.SeatTopCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.SeatTopCurveShape](), formDiv)
		case "IsHiddenRotatedSeatTopCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedSeatTopCurveShape, formDiv)
		case "RotatedSeatTopCurveShape":
			StageSetFormDivSelectFieldToField(&inst.RotatedSeatTopCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedSeatTopCurveShape](), formDiv)
		case "IsHiddenSeatBottomCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenSeatBottomCurveShape, formDiv)
		case "SeatBottomCurveShape":
			StageSetFormDivSelectFieldToField(&inst.SeatBottomCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.SeatBottomCurveShape](), formDiv)
		case "IsHiddenRotatedSeatBottomCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedSeatBottomCurveShape, formDiv)
		case "RotatedSeatBottomCurveShape":
			StageSetFormDivSelectFieldToField(&inst.RotatedSeatBottomCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedSeatBottomCurveShape](), formDiv)
		case "IsHiddenTorus3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTorus3DShape, formDiv)
		case "Torus3DShape":
			StageSetFormDivSelectFieldToField(&inst.Torus3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Torus3DShape](), formDiv)
		case "IsHiddenRotatedTorusShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedTorusShape, formDiv)
		case "RotatedTorusShape":
			StageSetFormDivSelectFieldToField(&inst.RotatedTorusShape, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedTorusShape](), formDiv)
		case "IsHiddenSampledPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenSampledPoints3DShape, formDiv)
		case "SampledPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.SampledPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.SampledPoints3DShape](), formDiv)
		case "IsHiddenRotatedSampledPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedSampledPoints3DShape, formDiv)
		case "RotatedSampledPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.RotatedSampledPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.RotatedSampledPoints3DShape](), formDiv)
		case "IsHiddenEyeSampledPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenEyeSampledPoints3DShape, formDiv)
		case "EyeSampledPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.EyeSampledPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.EyeSampledPoints3DShape](), formDiv)
		case "IsHiddenEyeCornersSampledPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenEyeCornersSampledPoints3DShape, formDiv)
		case "EyeCornersSampledPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.EyeCornersSampledPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.EyeCornersSampledPoints3DShape](), formDiv)
		case "IsHiddenEye3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenEye3DShape, formDiv)
		case "Eye3DShape":
			StageSetFormDivSelectFieldToField(&inst.Eye3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Eye3DShape](), formDiv)
		case "IsHiddenEyeSeatBottomCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenEyeSeatBottomCurveShape, formDiv)
		case "EyeSeatBottomCurveShape":
			StageSetFormDivSelectFieldToField(&inst.EyeSeatBottomCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.EyeSeatBottomCurveShape](), formDiv)
		case "IsHiddenEyeStoolBottomCurveShape":
			FormDivBasicFieldToField(&inst.IsHiddenEyeStoolBottomCurveShape, formDiv)
		case "EyeStoolBottomCurveShape":
			StageSetFormDivSelectFieldToField(&inst.EyeStoolBottomCurveShape, probe.stageSet.Stage.GetInstancesSet[*models.EyeStoolBottomCurveShape](), formDiv)
		case "IsHiddenSeat3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenSeat3DShape, formDiv)
		case "Seat3DShape":
			StageSetFormDivSelectFieldToField(&inst.Seat3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Seat3DShape](), formDiv)
		case "IsHiddenEyeVolume3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenEyeVolume3DShape, formDiv)
		case "EyeVolume3DShape":
			StageSetFormDivSelectFieldToField(&inst.EyeVolume3DShape, probe.stageSet.Stage.GetInstancesSet[*models.EyeVolume3DShape](), formDiv)
		case "IsHiddenSeatAndLegs3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenSeatAndLegs3DShape, formDiv)
		case "SeatAndLegs3DShape":
			StageSetFormDivSelectFieldToField(&inst.SeatAndLegs3DShape, probe.stageSet.Stage.GetInstancesSet[*models.SeatAndLegs3DShape](), formDiv)
		case "IsHiddenRotatedSeatAndLegs3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenRotatedSeatAndLegs3DShape, formDiv)
		case "RotatedSeatAndLegs3DShape":
			StageSetFormDivSelectFieldToField(&inst.RotatedSeatAndLegs3DShape, probe.stageSet.Stage.GetInstancesSet[*models.RotatedSeatAndLegs3DShape](), formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTiledFloor3DShape, formDiv)
		case "TiledFloor3DShape":
			StageSetFormDivSelectFieldToField(&inst.TiledFloor3DShape, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), formDiv)
		case "Rendered3DShape":
			StageSetFormDivSelectFieldToField(&inst.Rendered3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_TopCurvePlane1Shape_Stage(
	inst *models.TopCurvePlane1Shape,
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

func saveStageSet_TopCurvePlane2Shape_Stage(
	inst *models.TopCurvePlane2Shape,
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

func saveStageSet_TubeVase3DDiagram_Stage(
	inst *models.TubeVase3DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon, formDiv)
		case "IsHiddenTorusStackShape":
			FormDivBasicFieldToField(&inst.IsHiddenTorusStackShape, formDiv)
		case "IsHiddenVerticalTorusStackShape":
			FormDivBasicFieldToField(&inst.IsHiddenVerticalTorusStackShape, formDiv)
		case "IsHiddenPartiallyRotatedTorusShape":
			FormDivBasicFieldToField(&inst.IsHiddenPartiallyRotatedTorusShape, formDiv)
		case "IsHiddenStackOfPartiallyRotatedTorusShape":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfPartiallyRotatedTorusShape, formDiv)
		case "IsHiddenPointsAndLines3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenPointsAndLines3DShape, formDiv)
		case "IsHiddenKeyHole3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenKeyHole3DShape, formDiv)
		case "IsHiddenKey3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenKey3DShape, formDiv)
		case "IsHiddenVolumeKey3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenVolumeKey3DShape, formDiv)
		case "IsHiddenTorusEdge3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTorusEdge3DShape, formDiv)
		case "IsHiddenSampledPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenSampledPoints3DShape, formDiv)
		case "IsHiddenOriginalPoints3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenOriginalPoints3DShape, formDiv)
		case "IsHiddenAngle0Shape":
			FormDivBasicFieldToField(&inst.IsHiddenAngle0Shape, formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&inst.IsHiddenTiledFloor3DShape, formDiv)
		case "IsHiddenTopCurvePlane1Shape":
			FormDivBasicFieldToField(&inst.IsHiddenTopCurvePlane1Shape, formDiv)
		case "IsHiddenBottomCurvePlane1Shape":
			FormDivBasicFieldToField(&inst.IsHiddenBottomCurvePlane1Shape, formDiv)
		case "IsHiddenTopCurvePlane2Shape":
			FormDivBasicFieldToField(&inst.IsHiddenTopCurvePlane2Shape, formDiv)
		case "IsHiddenBottomCurvePlane2Shape":
			FormDivBasicFieldToField(&inst.IsHiddenBottomCurvePlane2Shape, formDiv)
		case "IsHiddenVaseTrapezeRingShape":
			FormDivBasicFieldToField(&inst.IsHiddenVaseTrapezeRingShape, formDiv)
		case "IsHiddenStackOfVaseTrapezeRingsShape":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfVaseTrapezeRingsShape, formDiv)
		case "IsHiddenStackOfRotatedVaseTrapezeRingsShape":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfRotatedVaseTrapezeRingsShape, formDiv)
		case "Rendered3DShape":
			StageSetFormDivSelectFieldToField(&inst.Rendered3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), formDiv)
		case "TorusStackShape":
			StageSetFormDivSelectFieldToField(&inst.TorusStackShape, probe.stageSet.Stage.GetInstancesSet[*models.TorusStackShape](), formDiv)
		case "VerticalTorusStackShape":
			StageSetFormDivSelectFieldToField(&inst.VerticalTorusStackShape, probe.stageSet.Stage.GetInstancesSet[*models.VerticalTorusStackShape](), formDiv)
		case "PartiallyRotatedTorusShape":
			StageSetFormDivSelectFieldToField(&inst.PartiallyRotatedTorusShape, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedTorusShape](), formDiv)
		case "StackOfPartiallyRotatedTorusShape":
			StageSetFormDivSelectFieldToField(&inst.StackOfPartiallyRotatedTorusShape, probe.stageSet.Stage.GetInstancesSet[*models.StackOfPartiallyRotatedTorusShape](), formDiv)
		case "PointsAndLines3DShape":
			StageSetFormDivSelectFieldToField(&inst.PointsAndLines3DShape, probe.stageSet.Stage.GetInstancesSet[*models.PointsAndLines3DShape](), formDiv)
		case "SampledPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.SampledPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.SampledPoints3DShape](), formDiv)
		case "OriginalPoints3DShape":
			StageSetFormDivSelectFieldToField(&inst.OriginalPoints3DShape, probe.stageSet.Stage.GetInstancesSet[*models.OriginalPoints3DShape](), formDiv)
		case "Angle0Shape":
			StageSetFormDivSelectFieldToField(&inst.Angle0Shape, probe.stageSet.Stage.GetInstancesSet[*models.Angle0Shape](), formDiv)
		case "KeyHole3DShape":
			StageSetFormDivSelectFieldToField(&inst.KeyHole3DShape, probe.stageSet.Stage.GetInstancesSet[*models.KeyHole3DShape](), formDiv)
		case "Key3DShape":
			StageSetFormDivSelectFieldToField(&inst.Key3DShape, probe.stageSet.Stage.GetInstancesSet[*models.Key3DShape](), formDiv)
		case "VolumeKey3DShape":
			StageSetFormDivSelectFieldToField(&inst.VolumeKey3DShape, probe.stageSet.Stage.GetInstancesSet[*models.VolumeKey3DShape](), formDiv)
		case "TorusEdge3DShape":
			StageSetFormDivSelectFieldToField(&inst.TorusEdge3DShape, probe.stageSet.Stage.GetInstancesSet[*models.TorusEdge3DShape](), formDiv)
		case "TiledFloor3DShape":
			StageSetFormDivSelectFieldToField(&inst.TiledFloor3DShape, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), formDiv)
		case "TopCurvePlane1Shape":
			StageSetFormDivSelectFieldToField(&inst.TopCurvePlane1Shape, probe.stageSet.Stage.GetInstancesSet[*models.TopCurvePlane1Shape](), formDiv)
		case "BottomCurvePlane1Shape":
			StageSetFormDivSelectFieldToField(&inst.BottomCurvePlane1Shape, probe.stageSet.Stage.GetInstancesSet[*models.BottomCurvePlane1Shape](), formDiv)
		case "TopCurvePlane2Shape":
			StageSetFormDivSelectFieldToField(&inst.TopCurvePlane2Shape, probe.stageSet.Stage.GetInstancesSet[*models.TopCurvePlane2Shape](), formDiv)
		case "BottomCurvePlane2Shape":
			StageSetFormDivSelectFieldToField(&inst.BottomCurvePlane2Shape, probe.stageSet.Stage.GetInstancesSet[*models.BottomCurvePlane2Shape](), formDiv)
		case "VaseTrapezeRingShape":
			StageSetFormDivSelectFieldToField(&inst.VaseTrapezeRingShape, probe.stageSet.Stage.GetInstancesSet[*models.VaseTrapezeRingShape](), formDiv)
		case "StackOfVaseTrapezeRingsShape":
			StageSetFormDivSelectFieldToField(&inst.StackOfVaseTrapezeRingsShape, probe.stageSet.Stage.GetInstancesSet[*models.StackOfVaseTrapezeRingsShape](), formDiv)
		case "StackOfRotatedVaseTrapezeRingsShape":
			StageSetFormDivSelectFieldToField(&inst.StackOfRotatedVaseTrapezeRingsShape, probe.stageSet.Stage.GetInstancesSet[*models.StackOfRotatedVaseTrapezeRingsShape](), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_TubeVaseAbstract_Stage(
	inst *models.TubeVaseAbstract,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Z_Ribbon":
			FormDivBasicFieldToField(&inst.Z_Ribbon, formDiv)
		case "RibbonVerticalScale":
			FormDivBasicFieldToField(&inst.RibbonVerticalScale, formDiv)
		case "Plane1Height":
			FormDivBasicFieldToField(&inst.Plane1Height, formDiv)
		case "Plane2Height":
			FormDivBasicFieldToField(&inst.Plane2Height, formDiv)
		case "ProjectionAngle":
			FormDivBasicFieldToField(&inst.ProjectionAngle, formDiv)
		case "RelativeVerticalThickness":
			FormDivBasicFieldToField(&inst.RelativeVerticalThickness, formDiv)
		case "RelativeRadialThickness":
			FormDivBasicFieldToField(&inst.RelativeRadialThickness, formDiv)
		case "RelativeCuttedStackFloorHeight":
			FormDivBasicFieldToField(&inst.RelativeCuttedStackFloorHeight, formDiv)
		case "RelativeRotatedTorusSeparation":
			FormDivBasicFieldToField(&inst.RelativeRotatedTorusSeparation, formDiv)
		case "RotationRatio":
			FormDivBasicFieldToField(&inst.RotationRatio, formDiv)
		case "RadialRepetitions":
			FormDivBasicFieldToField(&inst.RadialRepetitions, formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&inst.Transparency, formDiv)
		case "HasAlternatingRingColors":
			FormDivBasicFieldToField(&inst.HasAlternatingRingColors, formDiv)
		case "RelativeTrajectoryOffsetX":
			FormDivBasicFieldToField(&inst.RelativeTrajectoryOffsetX, formDiv)
		case "RelativeTrajectoryOffsetY":
			FormDivBasicFieldToField(&inst.RelativeTrajectoryOffsetY, formDiv)
		case "NbStepP1P2":
			FormDivBasicFieldToField(&inst.NbStepP1P2, formDiv)
		case "ChosenStep":
			FormDivBasicFieldToField(&inst.ChosenStep, formDiv)
		case "RelativeHorizontalRingsHeight":
			FormDivBasicFieldToField(&inst.RelativeHorizontalRingsHeight, formDiv)
		case "OffsetKeyX":
			FormDivBasicFieldToField(&inst.OffsetKeyX, formDiv)
		case "OffsetKeyY":
			FormDivBasicFieldToField(&inst.OffsetKeyY, formDiv)
		case "HeightKey":
			FormDivBasicFieldToField(&inst.HeightKey, formDiv)
		case "WidthKey":
			FormDivBasicFieldToField(&inst.WidthKey, formDiv)
		case "RelativeKeySize":
			FormDivBasicFieldToField(&inst.RelativeKeySize, formDiv)
		case "MovieNbFrames":
			FormDivBasicFieldToField(&inst.MovieNbFrames, formDiv)
		case "PerpendicularVectorGridHalfway":
			StageSetFormDivSelectFieldToField(&inst.PerpendicularVectorGridHalfway, probe.stageSet.Stage.GetInstancesSet[*models.PerpendicularVectorGridHalfway](), formDiv)
		case "TopStartArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.TopStartArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.TopStartArcShapeGrid](), formDiv)
		case "TopEndArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.TopEndArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.TopEndArcShapeGrid](), formDiv)
		case "ShiftedBottomTopStartArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.ShiftedBottomTopStartArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedBottomTopStartArcShapeGrid](), formDiv)
		case "TopMidArcVectorShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.TopMidArcVectorShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.TopMidArcVectorShapeGrid](), formDiv)
		case "StartHalfwayArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.StartHalfwayArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.StartHalfwayArcShapeGrid](), formDiv)
		case "TopStartHalfwayArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.TopStartHalfwayArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.TopStartHalfwayArcShapeGrid](), formDiv)
		case "EndHalfwayArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.EndHalfwayArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.EndHalfwayArcShapeGrid](), formDiv)
		case "TopEndHalfwayArcShapeGrid":
			StageSetFormDivSelectFieldToField(&inst.TopEndHalfwayArcShapeGrid, probe.stageSet.Stage.GetInstancesSet[*models.TopEndHalfwayArcShapeGrid](), formDiv)
		case "StackOfRotatedGrowthCurve2D":
			StageSetFormDivSelectFieldToField(&inst.StackOfRotatedGrowthCurve2D, probe.stageSet.Stage.GetInstancesSet[*models.StackOfRotatedGrowthCurve2D](), formDiv)
		case "TopStackOfRotatedGrowthCurve2D":
			StageSetFormDivSelectFieldToField(&inst.TopStackOfRotatedGrowthCurve2D, probe.stageSet.Stage.GetInstancesSet[*models.TopStackOfRotatedGrowthCurve2D](), formDiv)
		case "TopGrowthCurve2D":
			StageSetFormDivSelectFieldToField(&inst.TopGrowthCurve2D, probe.stageSet.Stage.GetInstancesSet[*models.TopGrowthCurve2D](), formDiv)
		case "StackOfGrowthCurve2D":
			StageSetFormDivSelectFieldToField(&inst.StackOfGrowthCurve2D, probe.stageSet.Stage.GetInstancesSet[*models.StackOfGrowthCurve2D](), formDiv)
		case "TopStackOfGrowthCurve2D":
			StageSetFormDivSelectFieldToField(&inst.TopStackOfGrowthCurve2D, probe.stageSet.Stage.GetInstancesSet[*models.TopStackOfGrowthCurve2D](), formDiv)
		case "StackOfGrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.StackOfGrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.StackOfGrowthCurve2DRibbon](), formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.StackOfRotatedGrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.StackOfRotatedGrowthCurve2DRibbon](), formDiv)
		case "GrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.GrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.GrowthCurve2DRibbon](), formDiv)
		case "ShiftedRightGrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.ShiftedRightGrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedRightGrowthCurve2DRibbon](), formDiv)
		case "ShiftedLeftGrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.ShiftedLeftGrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedLeftGrowthCurve2DRibbon](), formDiv)
		case "PartiallyGrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.PartiallyGrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyGrowthCurve2DRibbon](), formDiv)
		case "ShiftedLeftPartiallyGrowthCurve2DRibbon":
			StageSetFormDivSelectFieldToField(&inst.ShiftedLeftPartiallyGrowthCurve2DRibbon, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedLeftPartiallyGrowthCurve2DRibbon](), formDiv)
		case "PartiallyGrowthCurve2DTrajectory":
			StageSetFormDivSelectFieldToField(&inst.PartiallyGrowthCurve2DTrajectory, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyGrowthCurve2DTrajectory](), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2":
			StageSetFormDivSelectFieldToField(&inst.PartiallyGrowthCurve2DTrajectoryP1P2, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyGrowthCurve2DTrajectoryP1P2](), formDiv)
		case "PxShape":
			StageSetFormDivSelectFieldToField(&inst.PxShape, probe.stageSet.Stage.GetInstancesSet[*models.PxShape](), formDiv)
		case "ChosenP1P2PairShape":
			StageSetFormDivSelectFieldToField(&inst.ChosenP1P2PairShape, probe.stageSet.Stage.GetInstancesSet[*models.ChosenP1P2PairShape](), formDiv)
		case "KeyHoleShape":
			StageSetFormDivSelectFieldToField(&inst.KeyHoleShape, probe.stageSet.Stage.GetInstancesSet[*models.KeyHoleShape](), formDiv)
		}
	}
}

func saveStageSet_Vase2DDiagram_Stage(
	inst *models.Vase2DDiagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&inst.Zoom, formDiv)
		case "IsVaseArcNodesExpanded":
			FormDivBasicFieldToField(&inst.IsVaseArcNodesExpanded, formDiv)
		case "IsVaseClampingNodesExpanded":
			FormDivBasicFieldToField(&inst.IsVaseClampingNodesExpanded, formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&inst.IsHiddenAxesShape, formDiv)
		case "IsHiddenBottomStartArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenBottomStartArcShapeGrid, formDiv)
		case "IsHiddenBottomEndArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenBottomEndArcShapeGrid, formDiv)
		case "IsHiddenBottomStackOfGrowthCurve":
			FormDivBasicFieldToField(&inst.IsHiddenBottomStackOfGrowthCurve, formDiv)
		case "IsHiddenShiftedLeftStackOfGrowthCurve":
			FormDivBasicFieldToField(&inst.IsHiddenShiftedLeftStackOfGrowthCurve, formDiv)
		case "IsHiddenShiftedLeftStackOfNormalVector":
			FormDivBasicFieldToField(&inst.IsHiddenShiftedLeftStackOfNormalVector, formDiv)
		case "IsHiddenPerpendicularVectorGridHalfway":
			FormDivBasicFieldToField(&inst.IsHiddenPerpendicularVectorGridHalfway, formDiv)
		case "IsHiddenTopStartArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenTopStartArcShapeGrid, formDiv)
		case "IsHiddenShiftedBottomTopStartArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenShiftedBottomTopStartArcShapeGrid, formDiv)
		case "IsHiddenTopMidArcVectorShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenTopMidArcVectorShapeGrid, formDiv)
		case "IsHiddenStartHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenStartHalfwayArcShapeGrid, formDiv)
		case "IsHiddenTopStartHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenTopStartHalfwayArcShapeGrid, formDiv)
		case "IsHiddenEndHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenEndHalfwayArcShapeGrid, formDiv)
		case "IsHiddenTopEndHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenTopEndHalfwayArcShapeGrid, formDiv)
		case "IsHiddenTopEndArcShapeGrid":
			FormDivBasicFieldToField(&inst.IsHiddenTopEndArcShapeGrid, formDiv)
		case "IsHiddenStackOfGrowthCurve":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfGrowthCurve, formDiv)
		case "IsHiddenTopStackOfGrowthCurve":
			FormDivBasicFieldToField(&inst.IsHiddenTopStackOfGrowthCurve, formDiv)
		case "IsHiddenTopGrowthCurve2D":
			FormDivBasicFieldToField(&inst.IsHiddenTopGrowthCurve2D, formDiv)
		case "IsHiddenStackOfGrowthCurve2D":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfGrowthCurve2D, formDiv)
		case "IsHiddenTopStackOfGrowthCurve2D":
			FormDivBasicFieldToField(&inst.IsHiddenTopStackOfGrowthCurve2D, formDiv)
		case "IsHiddenGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenGrowthCurve2DRibbon, formDiv)
		case "IsHiddenShiftedRightGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenShiftedRightGrowthCurve2DRibbon, formDiv)
		case "IsHiddenShiftedLeftGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenShiftedLeftGrowthCurve2DRibbon, formDiv)
		case "IsHiddenStackOfGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfGrowthCurve2DRibbon, formDiv)
		case "IsHiddenStackOfRotatedGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenStackOfRotatedGrowthCurve2DRibbon, formDiv)
		case "IsHiddenPartiallyGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenPartiallyGrowthCurve2DRibbon, formDiv)
		case "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&inst.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon, formDiv)
		case "IsHiddenPartiallyGrowthCurve2DTrajectory":
			FormDivBasicFieldToField(&inst.IsHiddenPartiallyGrowthCurve2DTrajectory, formDiv)
		case "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2":
			FormDivBasicFieldToField(&inst.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2, formDiv)
		case "IsHiddenPxShape":
			FormDivBasicFieldToField(&inst.IsHiddenPxShape, formDiv)
		case "IsHiddenChosenP1P2PairShape":
			FormDivBasicFieldToField(&inst.IsHiddenChosenP1P2PairShape, formDiv)
		case "IsHiddenKeyHoleShape":
			FormDivBasicFieldToField(&inst.IsHiddenKeyHoleShape, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_VaseTrapezeRingShape_Stage(
	inst *models.VaseTrapezeRingShape,
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

func saveStageSet_StoolAbstract_StoolStage(
	inst *stool.StoolAbstract,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RadialRepetitions":
			FormDivBasicFieldToField(&inst.RadialRepetitions, formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&inst.Transparency, formDiv)
		case "RelativeTubeDiameter":
			FormDivBasicFieldToField(&inst.RelativeTubeDiameter, formDiv)
		case "RelativeHeight3DTorus":
			FormDivBasicFieldToField(&inst.RelativeHeight3DTorus, formDiv)
		case "StoolTorusVerticalScale":
			FormDivBasicFieldToField(&inst.StoolTorusVerticalScale, formDiv)
		case "RelativeHeight":
			FormDivBasicFieldToField(&inst.RelativeHeight, formDiv)
		case "RelativeSeatThickness":
			FormDivBasicFieldToField(&inst.RelativeSeatThickness, formDiv)
		case "ProjectionAngle":
			FormDivBasicFieldToField(&inst.ProjectionAngle, formDiv)
		case "RelativeEyeSeparationCriteria":
			FormDivBasicFieldToField(&inst.RelativeEyeSeparationCriteria, formDiv)
		case "RelativeEyeCornerControlVectorStrength":
			FormDivBasicFieldToField(&inst.RelativeEyeCornerControlVectorStrength, formDiv)
		}
	}
}

func saveStageSet_MusicAbstract_MusicStage(
	inst *music.MusicAbstract,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "PitchHeight":
			FormDivBasicFieldToField(&inst.PitchHeight, formDiv)
		case "NbOfBeatsInTheme":
			FormDivBasicFieldToField(&inst.NbOfBeatsInTheme, formDiv)
		case "BeatsPerSecond":
			FormDivBasicFieldToField(&inst.BeatsPerSecond, formDiv)
		case "FirstVoiceShiftX":
			FormDivBasicFieldToField(&inst.FirstVoiceShiftX, formDiv)
		case "FirstVoiceShiftY":
			FormDivBasicFieldToField(&inst.FirstVoiceShiftY, formDiv)
		case "PitchDifference":
			FormDivBasicFieldToField(&inst.PitchDifference, formDiv)
		case "Level":
			FormDivBasicFieldToField(&inst.Level, formDiv)
		case "ActualBeatsTemporalShift":
			FormDivBasicFieldToField(&inst.ActualBeatsTemporalShift, formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&inst.IsMinor, formDiv)
		case "ThemeBinaryEncoding":
			FormDivBasicFieldToField(&inst.ThemeBinaryEncoding, formDiv)
		case "BezierControlLengthRatio":
			FormDivBasicFieldToField(&inst.BezierControlLengthRatio, formDiv)
		case "NbPitchLines":
			FormDivBasicFieldToField(&inst.NbPitchLines, formDiv)
		case "NbBeatLines":
			FormDivBasicFieldToField(&inst.NbBeatLines, formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&inst.OriginX, formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&inst.OriginY, formDiv)
		case "ScoreScale":
			FormDivBasicFieldToField(&inst.ScoreScale, formDiv)
		case "ShowFirstVoice":
			FormDivBasicFieldToField(&inst.ShowFirstVoice, formDiv)
		case "ShowFirstVoiceShiftRight":
			FormDivBasicFieldToField(&inst.ShowFirstVoiceShiftRight, formDiv)
		case "ShowSecondVoice":
			FormDivBasicFieldToField(&inst.ShowSecondVoice, formDiv)
		case "ShowSecondVoiceShiftRight":
			FormDivBasicFieldToField(&inst.ShowSecondVoiceShiftRight, formDiv)
		case "ShowFirstVoiceNotes":
			FormDivBasicFieldToField(&inst.ShowFirstVoiceNotes, formDiv)
		case "ShowFirstVoiceNotesShiftRight":
			FormDivBasicFieldToField(&inst.ShowFirstVoiceNotesShiftRight, formDiv)
		case "ShowSecondVoiceNotes":
			FormDivBasicFieldToField(&inst.ShowSecondVoiceNotes, formDiv)
		case "ShowSecondVoiceNotesShiftRight":
			FormDivBasicFieldToField(&inst.ShowSecondVoiceNotesShiftRight, formDiv)
		case "IsComposerNodeExpanded":
			FormDivBasicFieldToField(&inst.IsComposerNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_ClockAbstract_ClockStage(
	inst *clock.ClockAbstract,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RadialRepetitions":
			FormDivBasicFieldToField(&inst.RadialRepetitions, formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&inst.Transparency, formDiv)
		case "RelativeTubeDiameter":
			FormDivBasicFieldToField(&inst.RelativeTubeDiameter, formDiv)
		case "RelativeHeight3DTorus":
			FormDivBasicFieldToField(&inst.RelativeHeight3DTorus, formDiv)
		case "ClockTorusVerticalScale":
			FormDivBasicFieldToField(&inst.ClockTorusVerticalScale, formDiv)
		case "RelativeHeight":
			FormDivBasicFieldToField(&inst.RelativeHeight, formDiv)
		case "ProjectionAngle":
			FormDivBasicFieldToField(&inst.ProjectionAngle, formDiv)
		}
	}
}


func StageSetNewInstance_Angle0Shape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Angle0Shape",
	}).Stage(probe.formStage)
	inst := new(models.Angle0Shape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Angle0Shape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Angle0Shape_Stage(probe)
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

func StageSetNewInstance_BottomCurvePlane1Shape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BottomCurvePlane1Shape",
	}).Stage(probe.formStage)
	inst := new(models.BottomCurvePlane1Shape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_BottomCurvePlane1Shape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_BottomCurvePlane1Shape_Stage(probe)
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

func StageSetNewInstance_BottomCurvePlane2Shape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New BottomCurvePlane2Shape",
	}).Stage(probe.formStage)
	inst := new(models.BottomCurvePlane2Shape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_BottomCurvePlane2Shape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_BottomCurvePlane2Shape_Stage(probe)
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

func StageSetNewInstance_Circumference3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Circumference3DShape",
	}).Stage(probe.formStage)
	inst := new(models.Circumference3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Circumference3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Circumference3DShape_Stage(probe)
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

func StageSetNewInstance_Clock2DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Clock2DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Clock2DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Clock2DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Clock2DDiagram_Stage(probe)
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

func StageSetNewInstance_Clock3DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Clock3DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Clock3DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Clock3DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Clock3DDiagram_Stage(probe)
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

func StageSetNewInstance_CutLine3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New CutLine3DShape",
	}).Stage(probe.formStage)
	inst := new(models.CutLine3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_CutLine3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_CutLine3DShape_Stage(probe)
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

func StageSetNewInstance_Leaves3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Leaves3DShape",
	}).Stage(probe.formStage)
	inst := new(models.Leaves3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Leaves3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Leaves3DShape_Stage(probe)
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

func StageSetNewInstance_Library_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Library",
	}).Stage(probe.formStage)
	inst := new(models.Library)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Library_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Library_Stage(probe)
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

func StageSetNewInstance_OriginalPoints3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New OriginalPoints3DShape",
	}).Stage(probe.formStage)
	inst := new(models.OriginalPoints3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_OriginalPoints3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_OriginalPoints3DShape_Stage(probe)
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

func StageSetNewInstance_ParastichyMCurves3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ParastichyMCurves3DShape",
	}).Stage(probe.formStage)
	inst := new(models.ParastichyMCurves3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ParastichyMCurves3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ParastichyMCurves3DShape_Stage(probe)
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

func StageSetNewInstance_ParastichyNCurves3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ParastichyNCurves3DShape",
	}).Stage(probe.formStage)
	inst := new(models.ParastichyNCurves3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ParastichyNCurves3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ParastichyNCurves3DShape_Stage(probe)
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

func StageSetNewInstance_Plant2DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Plant2DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Plant2DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Plant2DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Plant2DDiagram_Stage(probe)
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

func StageSetNewInstance_Plant3DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Plant3DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Plant3DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Plant3DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Plant3DDiagram_Stage(probe)
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

func StageSetNewInstance_PlantAbstract_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New PlantAbstract",
	}).Stage(probe.formStage)
	inst := new(models.PlantAbstract)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_PlantAbstract_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_PlantAbstract_Stage(probe)
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

func StageSetNewInstance_Rendered3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Rendered3DShape",
	}).Stage(probe.formStage)
	inst := new(models.Rendered3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Rendered3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Rendered3DShape_Stage(probe)
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

func StageSetNewInstance_SampledPoints3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New SampledPoints3DShape",
	}).Stage(probe.formStage)
	inst := new(models.SampledPoints3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_SampledPoints3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_SampledPoints3DShape_Stage(probe)
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

func StageSetNewInstance_StackOfRotatedVaseTrapezeRingsShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StackOfRotatedVaseTrapezeRingsShape",
	}).Stage(probe.formStage)
	inst := new(models.StackOfRotatedVaseTrapezeRingsShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StackOfRotatedVaseTrapezeRingsShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StackOfRotatedVaseTrapezeRingsShape_Stage(probe)
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

func StageSetNewInstance_StackOfVaseTrapezeRingsShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StackOfVaseTrapezeRingsShape",
	}).Stage(probe.formStage)
	inst := new(models.StackOfVaseTrapezeRingsShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StackOfVaseTrapezeRingsShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StackOfVaseTrapezeRingsShape_Stage(probe)
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

func StageSetNewInstance_StemCylinder3DShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StemCylinder3DShape",
	}).Stage(probe.formStage)
	inst := new(models.StemCylinder3DShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_StemCylinder3DShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_StemCylinder3DShape_Stage(probe)
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

func StageSetNewInstance_Stool2DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Stool2DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Stool2DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Stool2DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Stool2DDiagram_Stage(probe)
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

func StageSetNewInstance_Stool3DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Stool3DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Stool3DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Stool3DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Stool3DDiagram_Stage(probe)
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

func StageSetNewInstance_TopCurvePlane1Shape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TopCurvePlane1Shape",
	}).Stage(probe.formStage)
	inst := new(models.TopCurvePlane1Shape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TopCurvePlane1Shape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TopCurvePlane1Shape_Stage(probe)
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

func StageSetNewInstance_TopCurvePlane2Shape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TopCurvePlane2Shape",
	}).Stage(probe.formStage)
	inst := new(models.TopCurvePlane2Shape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TopCurvePlane2Shape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TopCurvePlane2Shape_Stage(probe)
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

func StageSetNewInstance_TubeVase3DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TubeVase3DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.TubeVase3DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TubeVase3DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TubeVase3DDiagram_Stage(probe)
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

func StageSetNewInstance_TubeVaseAbstract_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TubeVaseAbstract",
	}).Stage(probe.formStage)
	inst := new(models.TubeVaseAbstract)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TubeVaseAbstract_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TubeVaseAbstract_Stage(probe)
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

func StageSetNewInstance_Vase2DDiagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Vase2DDiagram",
	}).Stage(probe.formStage)
	inst := new(models.Vase2DDiagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Vase2DDiagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Vase2DDiagram_Stage(probe)
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

func StageSetNewInstance_VaseTrapezeRingShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New VaseTrapezeRingShape",
	}).Stage(probe.formStage)
	inst := new(models.VaseTrapezeRingShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_VaseTrapezeRingShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_VaseTrapezeRingShape_Stage(probe)
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

func StageSetNewInstance_StoolAbstract_StoolStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New StoolAbstract",
	}).Stage(probe.formStage)
	inst := new(stool.StoolAbstract)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.StoolStage.Lock()
			defer probe.stageSet.StoolStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.StoolStage)
			saveStageSet_StoolAbstract_StoolStage(inst, probe, formGroup)
			probe.stageSet.StoolStage.Commit()
			updateStageSetTable_StoolAbstract_StoolStage(probe)
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

func StageSetNewInstance_MusicAbstract_MusicStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MusicAbstract",
	}).Stage(probe.formStage)
	inst := new(music.MusicAbstract)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.MusicStage.Lock()
			defer probe.stageSet.MusicStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.MusicStage)
			saveStageSet_MusicAbstract_MusicStage(inst, probe, formGroup)
			probe.stageSet.MusicStage.Commit()
			updateStageSetTable_MusicAbstract_MusicStage(probe)
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

func StageSetNewInstance_ClockAbstract_ClockStage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ClockAbstract",
	}).Stage(probe.formStage)
	inst := new(clock.ClockAbstract)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.ClockStage.Lock()
			defer probe.stageSet.ClockStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.ClockStage)
			saveStageSet_ClockAbstract_ClockStage(inst, probe, formGroup)
			probe.stageSet.ClockStage.Commit()
			updateStageSetTable_ClockAbstract_ClockStage(probe)
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


// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
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
	case *models.Diagram:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Diagram_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Diagram_Stage(probe)
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
	case *models.Note:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Note_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Note_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteProductShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteProductShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteProductShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteResourceShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteResourceShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteResourceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.NoteTaskShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_NoteTaskShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_NoteTaskShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Product:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Product_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Product_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ProductCompositionShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ProductCompositionShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ProductCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ProductReferenceShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ProductReferenceShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ProductReferenceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ProductShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ProductShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ProductShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Resource:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Resource_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Resource_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ResourceCompositionShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ResourceCompositionShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ResourceCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ResourceShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ResourceShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ResourceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ResourceTaskShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ResourceTaskShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ResourceTaskShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Task:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Task_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Task_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskCompositionShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskCompositionShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskGroup:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskGroup_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskGroup_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskGroupShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskGroupShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskGroupShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskInputShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskInputShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskInputShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskOutputShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskOutputShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskOutputShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskPredecessorShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskPredecessorShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskPredecessorShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.TaskShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_TaskShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_TaskShape_Stage(probe)
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

func saveStageSet_Diagram_Stage(
	inst *models.Diagram,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&inst.DefaultBoxWidth, formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&inst.DefaultBoxHeigth, formDiv)
		case "DateFormat":
			FormDivBasicFieldToField(&inst.DateFormat, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsTimeDiagram":
			FormDivBasicFieldToField(&inst.IsTimeDiagram, formDiv)
		case "ComputedStart":
			FormDivTimeFieldToField(&inst.ComputedStart, formDiv, false)
		case "ComputedEnd":
			FormDivTimeFieldToField(&inst.ComputedEnd, formDiv, false)
		case "ComputedDuration":
			FormDivBasicFieldToField(&inst.ComputedDuration, formDiv)
		case "DrawVerticalTimeLines":
			FormDivBasicFieldToField(&inst.DrawVerticalTimeLines, formDiv)
		case "DrawSecondaryVerticalTimeLines":
			FormDivBasicFieldToField(&inst.DrawSecondaryVerticalTimeLines, formDiv)
		case "HideWeekendsPeriod":
			FormDivBasicFieldToField(&inst.HideWeekendsPeriod, formDiv)
		case "UseManualStartAndEndDates":
			FormDivBasicFieldToField(&inst.UseManualStartAndEndDates, formDiv)
		case "ManualStart":
			FormDivTimeFieldToField(&inst.ManualStart, formDiv, false)
		case "ManualEnd":
			FormDivTimeFieldToField(&inst.ManualEnd, formDiv, false)
		case "TimeStep":
			FormDivBasicFieldToField(&inst.TimeStep, formDiv)
		case "TimeStepScale":
			StageSetFormDivEnumStringFieldToField(&inst.TimeStepScale, formDiv)
		case "SecondaryTimeStep":
			FormDivBasicFieldToField(&inst.SecondaryTimeStep, formDiv)
		case "SecondaryTimeStepScale":
			StageSetFormDivEnumStringFieldToField(&inst.SecondaryTimeStepScale, formDiv)
		case "LaneHeight":
			FormDivBasicFieldToField(&inst.LaneHeight, formDiv)
		case "RatioBarToLaneHeight":
			FormDivBasicFieldToField(&inst.RatioBarToLaneHeight, formDiv)
		case "YTopMargin":
			FormDivBasicFieldToField(&inst.YTopMargin, formDiv)
		case "XLeftText":
			FormDivBasicFieldToField(&inst.XLeftText, formDiv)
		case "TextHeight":
			FormDivBasicFieldToField(&inst.TextHeight, formDiv)
		case "XLeftLanes":
			FormDivBasicFieldToField(&inst.XLeftLanes, formDiv)
		case "XRightMargin":
			FormDivBasicFieldToField(&inst.XRightMargin, formDiv)
		case "ArrowLengthToTheRightOfStartBar":
			FormDivBasicFieldToField(&inst.ArrowLengthToTheRightOfStartBar, formDiv)
		case "ArrowTipLenght":
			FormDivBasicFieldToField(&inst.ArrowTipLenght, formDiv)
		case "TimeLine_Color":
			FormDivBasicFieldToField(&inst.TimeLine_Color, formDiv)
		case "TimeLine_FillOpacity":
			FormDivBasicFieldToField(&inst.TimeLine_FillOpacity, formDiv)
		case "TimeLine_Stroke":
			FormDivBasicFieldToField(&inst.TimeLine_Stroke, formDiv)
		case "TimeLine_StrokeWidth":
			FormDivBasicFieldToField(&inst.TimeLine_StrokeWidth, formDiv)
		case "Group_Stroke":
			FormDivBasicFieldToField(&inst.Group_Stroke, formDiv)
		case "Group_StrokeWidth":
			FormDivBasicFieldToField(&inst.Group_StrokeWidth, formDiv)
		case "Group_StrokeDashArray":
			FormDivBasicFieldToField(&inst.Group_StrokeDashArray, formDiv)
		case "DateYOffset":
			FormDivBasicFieldToField(&inst.DateYOffset, formDiv)
		case "AlignOnBeginningOfTimeScale":
			FormDivBasicFieldToField(&inst.AlignOnBeginningOfTimeScale, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&inst.IsEditable_, formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&inst.IsShowPrefix, formDiv)
		case "IsInAutoLayoutMode":
			FormDivBasicFieldToField(&inst.IsInAutoLayoutMode, formDiv)
		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&inst.IsPBSNodeExpanded, formDiv)
		case "IsWBSNodeExpanded":
			FormDivBasicFieldToField(&inst.IsWBSNodeExpanded, formDiv)
		case "IsTaskGroupsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsTaskGroupsNodeExpanded, formDiv)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNotesNodeExpanded, formDiv)
		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsResourcesNodeExpanded, formDiv)
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

func saveStageSet_Note_Stage(
	inst *models.Note,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "LayoutDirection":
			StageSetFormDivEnumIntFieldToField(&inst.LayoutDirection, formDiv)
		}
	}
}

func saveStageSet_NoteProductShape_Stage(
	inst *models.NoteProductShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "Product":
			StageSetFormDivSelectFieldToField(&inst.Product, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_NoteResourceShape_Stage(
	inst *models.NoteResourceShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "Resource":
			StageSetFormDivSelectFieldToField(&inst.Resource, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_NoteShape_Stage(
	inst *models.NoteShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "IsLayoutDirectionDifferent":
			FormDivBasicFieldToField(&inst.IsLayoutDirectionDifferent, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_NoteTaskShape_Stage(
	inst *models.NoteTaskShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageSet.Stage.GetInstancesSet[*models.Note](), formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Product_Stage(
	inst *models.Product,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&inst.IsProducersNodeExpanded, formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&inst.IsConsumersNodeExpanded, formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&inst.IsImport, formDiv)
		case "ReferencedProduct":
			StageSetFormDivSelectFieldToField(&inst.ReferencedProduct, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "LayoutDirection":
			StageSetFormDivEnumIntFieldToField(&inst.LayoutDirection, formDiv)
		}
	}
}

func saveStageSet_ProductCompositionShape_Stage(
	inst *models.ProductCompositionShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Product":
			StageSetFormDivSelectFieldToField(&inst.Product, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ProductReferenceShape_Stage(
	inst *models.ProductReferenceShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Product":
			StageSetFormDivSelectFieldToField(&inst.Product, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "ReferencedProduct":
			StageSetFormDivSelectFieldToField(&inst.ReferencedProduct, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ProductShape_Stage(
	inst *models.ProductShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Product":
			StageSetFormDivSelectFieldToField(&inst.Product, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "IsShowType":
			FormDivBasicFieldToField(&inst.IsShowType, formDiv)
		case "IsLayoutDirectionDifferent":
			FormDivBasicFieldToField(&inst.IsLayoutDirectionDifferent, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Resource_Stage(
	inst *models.Resource,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "LayoutDirection":
			StageSetFormDivEnumIntFieldToField(&inst.LayoutDirection, formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&inst.IsImport, formDiv)
		case "ReferencedResource":
			StageSetFormDivSelectFieldToField(&inst.ReferencedResource, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), formDiv)
		}
	}
}

func saveStageSet_ResourceCompositionShape_Stage(
	inst *models.ResourceCompositionShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Resource":
			StageSetFormDivSelectFieldToField(&inst.Resource, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ResourceShape_Stage(
	inst *models.ResourceShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Resource":
			StageSetFormDivSelectFieldToField(&inst.Resource, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), formDiv)
		case "IsLayoutDirectionDifferent":
			FormDivBasicFieldToField(&inst.IsLayoutDirectionDifferent, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_ResourceTaskShape_Stage(
	inst *models.ResourceTaskShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Resource":
			StageSetFormDivSelectFieldToField(&inst.Resource, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_Task_Stage(
	inst *models.Task,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Start":
			FormDivTimeFieldToField(&inst.Start, formDiv, false)
		case "End":
			FormDivTimeFieldToField(&inst.End, formDiv, false)
		case "IsAllDay":
			FormDivBasicFieldToField(&inst.IsAllDay, formDiv)
		case "IsMilestone":
			FormDivBasicFieldToField(&inst.IsMilestone, formDiv)
		case "DependencyType":
			StageSetFormDivEnumStringFieldToField(&inst.DependencyType, formDiv)
		case "DependencyDurationYears":
			FormDivBasicFieldToField(&inst.DependencyDurationYears, formDiv)
		case "DependencyDurationMonths":
			FormDivBasicFieldToField(&inst.DependencyDurationMonths, formDiv)
		case "DependencyDurationWeeks":
			FormDivBasicFieldToField(&inst.DependencyDurationWeeks, formDiv)
		case "DependencyDurationDays":
			FormDivBasicFieldToField(&inst.DependencyDurationDays, formDiv)
		case "DependencyDurationHours":
			FormDivBasicFieldToField(&inst.DependencyDurationHours, formDiv)
		case "DurationYears":
			FormDivBasicFieldToField(&inst.DurationYears, formDiv)
		case "DurationMonths":
			FormDivBasicFieldToField(&inst.DurationMonths, formDiv)
		case "DurationWeeks":
			FormDivBasicFieldToField(&inst.DurationWeeks, formDiv)
		case "DurationDays":
			FormDivBasicFieldToField(&inst.DurationDays, formDiv)
		case "DurationHours":
			FormDivBasicFieldToField(&inst.DurationHours, formDiv)
		case "IsEndDateComputedFromDuration":
			FormDivBasicFieldToField(&inst.IsEndDateComputedFromDuration, formDiv)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&inst.IsWithCompletion, formDiv)
		case "Completion":
			StageSetFormDivEnumStringFieldToField(&inst.Completion, formDiv)
		case "TextPosition":
			StageSetFormDivEnumStringFieldToField(&inst.TextPosition, formDiv)
		case "XOffset":
			FormDivBasicFieldToField(&inst.XOffset, formDiv)
		case "YOffset":
			FormDivBasicFieldToField(&inst.YOffset, formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&inst.IsImport, formDiv)
		case "ReferencedTask":
			StageSetFormDivSelectFieldToField(&inst.ReferencedTask, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "IsInputsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsInputsNodeExpanded, formDiv)
		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsOutputsNodeExpanded, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "LayoutDirection":
			StageSetFormDivEnumIntFieldToField(&inst.LayoutDirection, formDiv)
		}
	}
}

func saveStageSet_TaskCompositionShape_Stage(
	inst *models.TaskCompositionShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_TaskGroup_Stage(
	inst *models.TaskGroup,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		}
	}
}

func saveStageSet_TaskGroupShape_Stage(
	inst *models.TaskGroupShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "TaskGroup":
			StageSetFormDivSelectFieldToField(&inst.TaskGroup, probe.stageSet.Stage.GetInstancesSet[*models.TaskGroup](), formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_TaskInputShape_Stage(
	inst *models.TaskInputShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Product":
			StageSetFormDivSelectFieldToField(&inst.Product, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_TaskOutputShape_Stage(
	inst *models.TaskOutputShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "Product":
			StageSetFormDivSelectFieldToField(&inst.Product, probe.stageSet.Stage.GetInstancesSet[*models.Product](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_TaskPredecessorShape_Stage(
	inst *models.TaskPredecessorShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Predecessor":
			StageSetFormDivSelectFieldToField(&inst.Predecessor, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&inst.StartRatio, formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&inst.EndRatio, formDiv)
		case "StartOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.StartOrientation, formDiv)
		case "EndOrientation":
			StageSetFormDivEnumStringFieldToField(&inst.EndOrientation, formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&inst.CornerOffsetRatio, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}

func saveStageSet_TaskShape_Stage(
	inst *models.TaskShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Task":
			StageSetFormDivSelectFieldToField(&inst.Task, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "IsShowDate":
			FormDivBasicFieldToField(&inst.IsShowDate, formDiv)
		case "VerticalOffset":
			FormDivBasicFieldToField(&inst.VerticalOffset, formDiv)
		case "DisplayVerticalBar":
			FormDivBasicFieldToField(&inst.DisplayVerticalBar, formDiv)
		case "IsLayoutDirectionDifferent":
			FormDivBasicFieldToField(&inst.IsLayoutDirectionDifferent, formDiv)
		case "X":
			FormDivBasicFieldToField(&inst.X, formDiv)
		case "Y":
			FormDivBasicFieldToField(&inst.Y, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
		}
	}
}


func StageSetNewInstance_Diagram_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Diagram",
	}).Stage(probe.formStage)
	inst := new(models.Diagram)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Diagram_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Diagram_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_Note_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Note",
	}).Stage(probe.formStage)
	inst := new(models.Note)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Note_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Note_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteProductShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteProductShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteProductShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteProductShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteProductShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteResourceShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteResourceShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteResourceShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteResourceShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteResourceShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteTaskShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New NoteTaskShape",
	}).Stage(probe.formStage)
	inst := new(models.NoteTaskShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_NoteTaskShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_NoteTaskShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Product_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Product",
	}).Stage(probe.formStage)
	inst := new(models.Product)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Product_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Product_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProductCompositionShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ProductCompositionShape",
	}).Stage(probe.formStage)
	inst := new(models.ProductCompositionShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ProductCompositionShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ProductCompositionShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProductReferenceShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ProductReferenceShape",
	}).Stage(probe.formStage)
	inst := new(models.ProductReferenceShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ProductReferenceShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ProductReferenceShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProductShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ProductShape",
	}).Stage(probe.formStage)
	inst := new(models.ProductShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ProductShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ProductShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Resource_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Resource",
	}).Stage(probe.formStage)
	inst := new(models.Resource)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Resource_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Resource_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ResourceCompositionShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ResourceCompositionShape",
	}).Stage(probe.formStage)
	inst := new(models.ResourceCompositionShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ResourceCompositionShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ResourceCompositionShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ResourceShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ResourceShape",
	}).Stage(probe.formStage)
	inst := new(models.ResourceShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ResourceShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ResourceShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ResourceTaskShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ResourceTaskShape",
	}).Stage(probe.formStage)
	inst := new(models.ResourceTaskShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ResourceTaskShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ResourceTaskShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Task_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Task",
	}).Stage(probe.formStage)
	inst := new(models.Task)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Task_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Task_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskCompositionShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskCompositionShape",
	}).Stage(probe.formStage)
	inst := new(models.TaskCompositionShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskCompositionShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskCompositionShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskGroup_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskGroup",
	}).Stage(probe.formStage)
	inst := new(models.TaskGroup)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskGroup_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskGroup_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskGroupShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskGroupShape",
	}).Stage(probe.formStage)
	inst := new(models.TaskGroupShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskGroupShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskGroupShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskInputShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskInputShape",
	}).Stage(probe.formStage)
	inst := new(models.TaskInputShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskInputShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskInputShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskOutputShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskOutputShape",
	}).Stage(probe.formStage)
	inst := new(models.TaskOutputShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskOutputShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskOutputShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskPredecessorShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskPredecessorShape",
	}).Stage(probe.formStage)
	inst := new(models.TaskPredecessorShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskPredecessorShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskPredecessorShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TaskShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New TaskShape",
	}).Stage(probe.formStage)
	inst := new(models.TaskShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_TaskShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_TaskShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}


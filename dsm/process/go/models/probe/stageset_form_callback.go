// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
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
	case *models.AllocatedProcessShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AllocatedProcessShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_AllocatedProcessShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.AllocatedResourceShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_AllocatedResourceShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_AllocatedResourceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ControlFlow:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ControlFlow_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ControlFlow_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ControlFlowShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ControlFlowShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ControlFlowShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Data:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Data_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Data_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DataFlow:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DataFlow_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DataFlow_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DataFlowShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DataFlowShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DataFlowShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DataShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DataShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DataShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.DiagramProcess:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_DiagramProcess_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_DiagramProcess_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ExternalParticipantShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ExternalParticipantShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ExternalParticipantShape_Stage(probe)
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
	case *models.Participant:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Participant_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Participant_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ParticipantShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ParticipantShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ParticipantShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Process:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Process_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Process_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ProcessShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ProcessShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ProcessShape_Stage(probe)
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

func saveStageSet_AllocatedProcessShape_Stage(
	inst *models.AllocatedProcessShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Participant":
			StageSetFormDivSelectFieldToField(&inst.Participant, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), formDiv)
		case "Process":
			StageSetFormDivSelectFieldToField(&inst.Process, probe.stageSet.Stage.GetInstancesSet[*models.Process](), formDiv)
		}
	}
}

func saveStageSet_AllocatedResourceShape_Stage(
	inst *models.AllocatedResourceShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Participant":
			StageSetFormDivSelectFieldToField(&inst.Participant, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), formDiv)
		case "Resource":
			StageSetFormDivSelectFieldToField(&inst.Resource, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), formDiv)
		}
	}
}

func saveStageSet_ControlFlow_Stage(
	inst *models.ControlFlow,
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
		case "Start":
			StageSetFormDivSelectFieldToField(&inst.Start, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "End":
			StageSetFormDivSelectFieldToField(&inst.End, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		}
	}
}

func saveStageSet_ControlFlowShape_Stage(
	inst *models.ControlFlowShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ControlFlow":
			StageSetFormDivSelectFieldToField(&inst.ControlFlow, probe.stageSet.Stage.GetInstancesSet[*models.ControlFlow](), formDiv)
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

func saveStageSet_Data_Stage(
	inst *models.Data,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Acronym":
			FormDivBasicFieldToField(&inst.Acronym, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&inst.SVG_Path, formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&inst.InverseAppliedScaling, formDiv)
		}
	}
}

func saveStageSet_DataFlow_Stage(
	inst *models.DataFlow,
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
		case "Type":
			StageSetFormDivEnumStringFieldToField(&inst.Type, formDiv)
		case "StartTask":
			StageSetFormDivSelectFieldToField(&inst.StartTask, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "EndTask":
			StageSetFormDivSelectFieldToField(&inst.EndTask, probe.stageSet.Stage.GetInstancesSet[*models.Task](), formDiv)
		case "StartExternalParticipant":
			StageSetFormDivSelectFieldToField(&inst.StartExternalParticipant, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), formDiv)
		case "EndExternalParticipant":
			StageSetFormDivSelectFieldToField(&inst.EndExternalParticipant, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), formDiv)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&inst.IsDatasNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_DataFlowShape_Stage(
	inst *models.DataFlowShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DataFlow":
			StageSetFormDivSelectFieldToField(&inst.DataFlow, probe.stageSet.Stage.GetInstancesSet[*models.DataFlow](), formDiv)
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

func saveStageSet_DataShape_Stage(
	inst *models.DataShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Data":
			StageSetFormDivSelectFieldToField(&inst.Data, probe.stageSet.Stage.GetInstancesSet[*models.Data](), formDiv)
		case "DataFlow":
			StageSetFormDivSelectFieldToField(&inst.DataFlow, probe.stageSet.Stage.GetInstancesSet[*models.DataFlow](), formDiv)
		}
	}
}

func saveStageSet_DiagramProcess_Stage(
	inst *models.DiagramProcess,
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
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&inst.IsEditable_, formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&inst.IsShowPrefix, formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&inst.DefaultBoxWidth, formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&inst.DefaultBoxHeigth, formDiv)
		case "Width":
			FormDivBasicFieldToField(&inst.Width, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "IsProcesssNodeExpanded":
			FormDivBasicFieldToField(&inst.IsProcesssNodeExpanded, formDiv)
		case "IsParticipantsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsParticipantsNodeExpanded, formDiv)
		case "IsExternalParticipantsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsExternalParticipantsNodeExpanded, formDiv)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNotesNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_ExternalParticipantShape_Stage(
	inst *models.ExternalParticipantShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Participant":
			StageSetFormDivSelectFieldToField(&inst.Participant, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
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
		case "TailHeigth":
			FormDivBasicFieldToField(&inst.TailHeigth, formDiv)
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
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&inst.IsRootLibrary, formDiv)
		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsSubLibrariesNodeExpanded, formDiv)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&inst.NbPixPerCharacter, formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&inst.LogoSVGFile, formDiv)
		case "IsProcessesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsProcessesNodeExpanded, formDiv)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsDataFlowsNodeExpanded, formDiv)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&inst.IsDatasNodeExpanded, formDiv)
		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsResourcesNodeExpanded, formDiv)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNotesNodeExpanded, formDiv)
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&inst.IsExpandedTmp, formDiv)
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
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsTasksNodeExpanded":
			FormDivBasicFieldToField(&inst.IsTasksNodeExpanded, formDiv)
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

func saveStageSet_Participant_Stage(
	inst *models.Participant,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IsProcessResource":
			FormDivBasicFieldToField(&inst.IsProcessResource, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsResourcesNodeExpanded, formDiv)
		case "IsProcessesNodeExpanded":
			FormDivBasicFieldToField(&inst.IsProcessesNodeExpanded, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsTasksNodeExpanded":
			FormDivBasicFieldToField(&inst.IsTasksNodeExpanded, formDiv)
		case "IsControlFlowsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsControlFlowsNodeExpanded, formDiv)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsDataFlowsNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_ParticipantShape_Stage(
	inst *models.ParticipantShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Participant":
			StageSetFormDivSelectFieldToField(&inst.Participant, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
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
		case "WidthWeight":
			FormDivBasicFieldToField(&inst.WidthWeight, formDiv)
		}
	}
}

func saveStageSet_Process_Stage(
	inst *models.Process,
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
		case "SVG_Path":
			FormDivBasicFieldToField(&inst.SVG_Path, formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&inst.InverseAppliedScaling, formDiv)
		case "IsSubProcessNodeExpanded":
			FormDivBasicFieldToField(&inst.IsSubProcessNodeExpanded, formDiv)
		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsDataFlowsNodeExpanded, formDiv)
		}
	}
}

func saveStageSet_ProcessShape_Stage(
	inst *models.ProcessShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Process":
			StageSetFormDivSelectFieldToField(&inst.Process, probe.stageSet.Stage.GetInstancesSet[*models.Process](), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
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
		case "Acronym":
			FormDivBasicFieldToField(&inst.Acronym, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&inst.SVG_Path, formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&inst.InverseAppliedScaling, formDiv)
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsStartTask":
			FormDivBasicFieldToField(&inst.IsStartTask, formDiv)
		case "IsEndTask":
			FormDivBasicFieldToField(&inst.IsEndTask, formDiv)
		case "Type":
			StageSetFormDivSelectFieldToField(&inst.Type, probe.stageSet.Stage.GetInstancesSet[*models.Process](), formDiv)
		case "IsTaskNameNotProcessName":
			FormDivBasicFieldToField(&inst.IsTaskNameNotProcessName, formDiv)
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
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
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


func StageSetNewInstance_AllocatedProcessShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AllocatedProcessShape",
	}).Stage(probe.formStage)
	inst := new(models.AllocatedProcessShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_AllocatedProcessShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_AllocatedProcessShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AllocatedResourceShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New AllocatedResourceShape",
	}).Stage(probe.formStage)
	inst := new(models.AllocatedResourceShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_AllocatedResourceShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_AllocatedResourceShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ControlFlow_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ControlFlow",
	}).Stage(probe.formStage)
	inst := new(models.ControlFlow)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ControlFlow_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ControlFlow_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ControlFlowShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ControlFlowShape",
	}).Stage(probe.formStage)
	inst := new(models.ControlFlowShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ControlFlowShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ControlFlowShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Data_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Data",
	}).Stage(probe.formStage)
	inst := new(models.Data)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Data_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Data_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DataFlow_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DataFlow",
	}).Stage(probe.formStage)
	inst := new(models.DataFlow)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DataFlow_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DataFlow_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DataFlowShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DataFlowShape",
	}).Stage(probe.formStage)
	inst := new(models.DataFlowShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DataFlowShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DataFlowShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DataShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DataShape",
	}).Stage(probe.formStage)
	inst := new(models.DataShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DataShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DataShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DiagramProcess_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New DiagramProcess",
	}).Stage(probe.formStage)
	inst := new(models.DiagramProcess)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_DiagramProcess_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_DiagramProcess_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ExternalParticipantShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ExternalParticipantShape",
	}).Stage(probe.formStage)
	inst := new(models.ExternalParticipantShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ExternalParticipantShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ExternalParticipantShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_Participant_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Participant",
	}).Stage(probe.formStage)
	inst := new(models.Participant)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Participant_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Participant_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ParticipantShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ParticipantShape",
	}).Stage(probe.formStage)
	inst := new(models.ParticipantShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ParticipantShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ParticipantShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Process_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Process",
	}).Stage(probe.formStage)
	inst := new(models.Process)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Process_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Process_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			StageSetFillUpFormFromGongstruct(inst, probe)
		},
	}
	StageSetFillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcessShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ProcessShape",
	}).Stage(probe.formStage)
	inst := new(models.ProcessShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ProcessShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ProcessShape_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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


// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ControlFlowFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ControlFlow", true)
			} else {
				FillUpFormFromGongstruct(onSave.controlflow, probe)
			}
		case *ControlFlowShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ControlFlowShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.controlflowshape, probe)
			}
		case *DataFlowFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DataFlow", true)
			} else {
				FillUpFormFromGongstruct(onSave.dataflow, probe)
			}
		case *DiagramProcessFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiagramProcess", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagramprocess, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *ParticipantFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Participant", true)
			} else {
				FillUpFormFromGongstruct(onSave.participant, probe)
			}
		case *ParticipantShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ParticipantShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.participantshape, probe)
			}
		case *ProcessFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Process", true)
			} else {
				FillUpFormFromGongstruct(onSave.process, probe)
			}
		case *ProcessShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ProcessShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.processshape, probe)
			}
		case *TaskFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Task", true)
			} else {
				FillUpFormFromGongstruct(onSave.task, probe)
			}
		case *TaskShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TaskShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.taskshape, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "ControlFlow":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlFlow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlFlowFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlflow := new(models.ControlFlow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlflow, formGroup, probe)
	case "ControlFlowShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlFlowShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlFlowShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlflowshape := new(models.ControlFlowShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlflowshape, formGroup, probe)
	case "DataFlow":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DataFlow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataFlowFormCallback(
			nil,
			probe,
			formGroup,
		)
		dataflow := new(models.DataFlow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dataflow, formGroup, probe)
	case "DiagramProcess":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramProcess Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramProcessFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagramprocess := new(models.DiagramProcess)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagramprocess, formGroup, probe)
	case "Library":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			probe,
			formGroup,
		)
		library := new(models.Library)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(library, formGroup, probe)
	case "Participant":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Participant Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParticipantFormCallback(
			nil,
			probe,
			formGroup,
		)
		participant := new(models.Participant)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(participant, formGroup, probe)
	case "ParticipantShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParticipantShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParticipantShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		participantshape := new(models.ParticipantShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(participantshape, formGroup, probe)
	case "Process":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Process Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessFormCallback(
			nil,
			probe,
			formGroup,
		)
		process := new(models.Process)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(process, formGroup, probe)
	case "ProcessShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProcessShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		processshape := new(models.ProcessShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(processshape, formGroup, probe)
	case "Task":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Task Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			probe,
			formGroup,
		)
		task := new(models.Task)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(task, formGroup, probe)
	case "TaskShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TaskShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		taskshape := new(models.TaskShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(taskshape, formGroup, probe)
	}
	formStage.Commit()
}

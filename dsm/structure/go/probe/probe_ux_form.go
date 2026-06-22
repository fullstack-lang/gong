// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AllocatedResourceShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AllocatedResourceShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.allocatedresourceshape, probe)
			}
		case *AllocatedSystemShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AllocatedSystemShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.allocatedsystemshape, probe)
			}
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
		case *DataFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Data", true)
			} else {
				FillUpFormFromGongstruct(onSave.data, probe)
			}
		case *DataFlowFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DataFlow", true)
			} else {
				FillUpFormFromGongstruct(onSave.dataflow, probe)
			}
		case *DataFlowShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DataFlowShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.dataflowshape, probe)
			}
		case *DataShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DataShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.datashape, probe)
			}
		case *DiagramStructureFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiagramStructure", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagramstructure, probe)
			}
		case *ExternalPartShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ExternalPartShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.externalpartshape, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *NoteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note", true)
			} else {
				FillUpFormFromGongstruct(onSave.note, probe)
			}
		case *NotePortShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NotePortShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteportshape, probe)
			}
		case *NoteShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteshape, probe)
			}
		case *PartFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part", true)
			} else {
				FillUpFormFromGongstruct(onSave.part, probe)
			}
		case *PartAnchoredPathFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PartAnchoredPath", true)
			} else {
				FillUpFormFromGongstruct(onSave.partanchoredpath, probe)
			}
		case *PartShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PartShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.partshape, probe)
			}
		case *PortFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Port", true)
			} else {
				FillUpFormFromGongstruct(onSave.port, probe)
			}
		case *PortShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PortShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.portshape, probe)
			}
		case *ResourceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Resource", true)
			} else {
				FillUpFormFromGongstruct(onSave.resource, probe)
			}
		case *SystemFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "System", true)
			} else {
				FillUpFormFromGongstruct(onSave.system, probe)
			}
		case *SystemShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SystemShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.systemshape, probe)
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
	case "AllocatedResourceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AllocatedResourceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllocatedResourceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		allocatedresourceshape := new(models.AllocatedResourceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(allocatedresourceshape, formGroup, probe)
	case "AllocatedSystemShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AllocatedSystemShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllocatedSystemShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		allocatedsystemshape := new(models.AllocatedSystemShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(allocatedsystemshape, formGroup, probe)
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
	case "Data":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Data Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataFormCallback(
			nil,
			probe,
			formGroup,
		)
		data := new(models.Data)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(data, formGroup, probe)
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
	case "DataFlowShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DataFlowShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataFlowShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		dataflowshape := new(models.DataFlowShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dataflowshape, formGroup, probe)
	case "DataShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DataShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		datashape := new(models.DataShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datashape, formGroup, probe)
	case "DiagramStructure":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramStructure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramStructureFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagramstructure := new(models.DiagramStructure)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagramstructure, formGroup, probe)
	case "ExternalPartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ExternalPartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExternalPartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		externalpartshape := new(models.ExternalPartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(externalpartshape, formGroup, probe)
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
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "NotePortShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NotePortShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotePortShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteportshape := new(models.NotePortShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteportshape, formGroup, probe)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteshape := new(models.NoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshape, formGroup, probe)
	case "Part":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartFormCallback(
			nil,
			probe,
			formGroup,
		)
		part := new(models.Part)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part, formGroup, probe)
	case "PartAnchoredPath":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartAnchoredPath Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartAnchoredPathFormCallback(
			nil,
			probe,
			formGroup,
		)
		partanchoredpath := new(models.PartAnchoredPath)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partanchoredpath, formGroup, probe)
	case "PartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		partshape := new(models.PartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(partshape, formGroup, probe)
	case "Port":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Port Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PortFormCallback(
			nil,
			probe,
			formGroup,
		)
		port := new(models.Port)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(port, formGroup, probe)
	case "PortShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PortShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PortShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		portshape := new(models.PortShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(portshape, formGroup, probe)
	case "Resource":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Resource Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceFormCallback(
			nil,
			probe,
			formGroup,
		)
		resource := new(models.Resource)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(resource, formGroup, probe)
	case "System":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "System Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemFormCallback(
			nil,
			probe,
			formGroup,
		)
		system := new(models.System)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system, formGroup, probe)
	case "SystemShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SystemShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		systemshape := new(models.SystemShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(systemshape, formGroup, probe)
	}
	formStage.Commit()
}

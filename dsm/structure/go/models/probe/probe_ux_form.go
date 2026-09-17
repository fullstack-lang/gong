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
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
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
	case "DiagramLayerState":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramLayerState Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramLayerStateFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagramlayerstate := new(models.DiagramLayerState)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagramlayerstate, formGroup, probe)
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
	case "LayerDefinition":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "LayerDefinition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerDefinitionFormCallback(
			nil,
			probe,
			formGroup,
		)
		layerdefinition := new(models.LayerDefinition)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layerdefinition, formGroup, probe)
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
	case "NotePartShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NotePartShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotePartShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		notepartshape := new(models.NotePartShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notepartshape, formGroup, probe)
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
	case "SemanticTag":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SemanticTag Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SemanticTagFormCallback(
			nil,
			probe,
			formGroup,
		)
		semantictag := new(models.SemanticTag)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(semantictag, formGroup, probe)
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

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
		case *DiagramStructureFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiagramStructure", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagramstructure, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *LinkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Link", true)
			} else {
				FillUpFormFromGongstruct(onSave.link, probe)
			}
		case *LinkAssociationShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "LinkAssociationShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.linkassociationshape, probe)
			}
		case *PartFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part", true)
			} else {
				FillUpFormFromGongstruct(onSave.part, probe)
			}
		case *PartShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PartShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.partshape, probe)
			}
		case *SystemFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "System", true)
			} else {
				FillUpFormFromGongstruct(onSave.system, probe)
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
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "LinkAssociationShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "LinkAssociationShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkAssociationShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		linkassociationshape := new(models.LinkAssociationShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(linkassociationshape, formGroup, probe)
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
	}
	formStage.Commit()
}

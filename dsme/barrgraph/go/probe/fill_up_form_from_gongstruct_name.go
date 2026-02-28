// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/barrgraph/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ArtefactTypeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ArtefactType", true)
			} else {
				FillUpFormFromGongstruct(onSave.artefacttype, probe)
			}
		case *ArtefactTypeShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ArtefactTypeShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.artefacttypeshape, probe)
			}
		case *ArtistFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Artist", true)
			} else {
				FillUpFormFromGongstruct(onSave.artist, probe)
			}
		case *ArtistShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ArtistShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.artistshape, probe)
			}
		case *ControlPointShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ControlPointShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.controlpointshape, probe)
			}
		case *DeskFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Desk", true)
			} else {
				FillUpFormFromGongstruct(onSave.desk, probe)
			}
		case *DiagramFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Diagram", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagram, probe)
			}
		case *InfluenceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Influence", true)
			} else {
				FillUpFormFromGongstruct(onSave.influence, probe)
			}
		case *InfluenceShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "InfluenceShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.influenceshape, probe)
			}
		case *MovementFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Movement", true)
			} else {
				FillUpFormFromGongstruct(onSave.movement, probe)
			}
		case *MovementShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MovementShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.movementshape, probe)
			}
		case *PlaceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Place", true)
			} else {
				FillUpFormFromGongstruct(onSave.place, probe)
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
	case "ArtefactType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ArtefactType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtefactTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		artefacttype := new(models.ArtefactType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(artefacttype, formGroup, probe)
	case "ArtefactTypeShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ArtefactTypeShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtefactTypeShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		artefacttypeshape := new(models.ArtefactTypeShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(artefacttypeshape, formGroup, probe)
	case "Artist":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Artist Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtistFormCallback(
			nil,
			probe,
			formGroup,
		)
		artist := new(models.Artist)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(artist, formGroup, probe)
	case "ArtistShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ArtistShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArtistShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		artistshape := new(models.ArtistShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(artistshape, formGroup, probe)
	case "ControlPointShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ControlPointShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(controlpointshape, formGroup, probe)
	case "Desk":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Desk Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DeskFormCallback(
			nil,
			probe,
			formGroup,
		)
		desk := new(models.Desk)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(desk, formGroup, probe)
	case "Diagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagram := new(models.Diagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagram, formGroup, probe)
	case "Influence":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Influence Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InfluenceFormCallback(
			nil,
			probe,
			formGroup,
		)
		influence := new(models.Influence)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(influence, formGroup, probe)
	case "InfluenceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "InfluenceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InfluenceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		influenceshape := new(models.InfluenceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(influenceshape, formGroup, probe)
	case "Movement":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Movement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MovementFormCallback(
			nil,
			probe,
			formGroup,
		)
		movement := new(models.Movement)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(movement, formGroup, probe)
	case "MovementShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MovementShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MovementShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		movementshape := new(models.MovementShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(movementshape, formGroup, probe)
	case "Place":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Place Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlaceFormCallback(
			nil,
			probe,
			formGroup,
		)
		place := new(models.Place)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(place, formGroup, probe)
	}
	formStage.Commit()
}

// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ArtefactTypeFormCallback(
	artefacttype *models.ArtefactType,
	probe *Probe,
	formGroup *table.FormGroup,
) (artefacttypeFormCallback *ArtefactTypeFormCallback) {
	artefacttypeFormCallback = new(ArtefactTypeFormCallback)
	artefacttypeFormCallback.probe = probe
	artefacttypeFormCallback.artefacttype = artefacttype
	artefacttypeFormCallback.formGroup = formGroup

	artefacttypeFormCallback.CreationMode = (artefacttype == nil)

	return
}

type ArtefactTypeFormCallback struct {
	artefacttype *models.ArtefactType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (artefacttypeFormCallback *ArtefactTypeFormCallback) OnSave() {

	// log.Println("ArtefactTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	artefacttypeFormCallback.probe.formStage.Checkout()

	if artefacttypeFormCallback.artefacttype == nil {
		artefacttypeFormCallback.artefacttype = new(models.ArtefactType).Stage(artefacttypeFormCallback.probe.stageOfInterest)
	}
	artefacttype_ := artefacttypeFormCallback.artefacttype
	_ = artefacttype_

	for _, formDiv := range artefacttypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(artefacttype_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if artefacttypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artefacttype_.Unstage(artefacttypeFormCallback.probe.stageOfInterest)
	}

	artefacttypeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ArtefactType](
		artefacttypeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if artefacttypeFormCallback.CreationMode || artefacttypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artefacttypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(artefacttypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArtefactTypeFormCallback(
			nil,
			artefacttypeFormCallback.probe,
			newFormGroup,
		)
		artefacttype := new(models.ArtefactType)
		FillUpForm(artefacttype, newFormGroup, artefacttypeFormCallback.probe)
		artefacttypeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(artefacttypeFormCallback.probe)
}
func __gong__New__ArtefactTypeShapeFormCallback(
	artefacttypeshape *models.ArtefactTypeShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (artefacttypeshapeFormCallback *ArtefactTypeShapeFormCallback) {
	artefacttypeshapeFormCallback = new(ArtefactTypeShapeFormCallback)
	artefacttypeshapeFormCallback.probe = probe
	artefacttypeshapeFormCallback.artefacttypeshape = artefacttypeshape
	artefacttypeshapeFormCallback.formGroup = formGroup

	artefacttypeshapeFormCallback.CreationMode = (artefacttypeshape == nil)

	return
}

type ArtefactTypeShapeFormCallback struct {
	artefacttypeshape *models.ArtefactTypeShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (artefacttypeshapeFormCallback *ArtefactTypeShapeFormCallback) OnSave() {

	// log.Println("ArtefactTypeShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	artefacttypeshapeFormCallback.probe.formStage.Checkout()

	if artefacttypeshapeFormCallback.artefacttypeshape == nil {
		artefacttypeshapeFormCallback.artefacttypeshape = new(models.ArtefactTypeShape).Stage(artefacttypeshapeFormCallback.probe.stageOfInterest)
	}
	artefacttypeshape_ := artefacttypeshapeFormCallback.artefacttypeshape
	_ = artefacttypeshape_

	for _, formDiv := range artefacttypeshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(artefacttypeshape_.Name), formDiv)
		case "ArtefactType":
			FormDivSelectFieldToField(&(artefacttypeshape_.ArtefactType), artefacttypeshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(artefacttypeshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(artefacttypeshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(artefacttypeshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(artefacttypeshape_.Height), formDiv)
		case "Diagram:ArtefactTypeShapes":
			// WARNING : this form deals with the N-N association "Diagram.ArtefactTypeShapes []*ArtefactTypeShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ArtefactTypeShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ArtefactTypeShapes"
				formerAssociationSource := artefacttypeshape_.GongGetReverseFieldOwner(
					artefacttypeshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ArtefactTypeShapes []*ArtefactTypeShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ArtefactTypeShapes, artefacttypeshape_)
					formerSource.ArtefactTypeShapes = slices.Delete(formerSource.ArtefactTypeShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](artefacttypeshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ArtefactTypeShapes []*ArtefactTypeShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ArtefactTypeShapes = append(newSource.ArtefactTypeShapes, artefacttypeshape_)
		}
	}

	// manage the suppress operation
	if artefacttypeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artefacttypeshape_.Unstage(artefacttypeshapeFormCallback.probe.stageOfInterest)
	}

	artefacttypeshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ArtefactTypeShape](
		artefacttypeshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if artefacttypeshapeFormCallback.CreationMode || artefacttypeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artefacttypeshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(artefacttypeshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArtefactTypeShapeFormCallback(
			nil,
			artefacttypeshapeFormCallback.probe,
			newFormGroup,
		)
		artefacttypeshape := new(models.ArtefactTypeShape)
		FillUpForm(artefacttypeshape, newFormGroup, artefacttypeshapeFormCallback.probe)
		artefacttypeshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(artefacttypeshapeFormCallback.probe)
}
func __gong__New__ArtistFormCallback(
	artist *models.Artist,
	probe *Probe,
	formGroup *table.FormGroup,
) (artistFormCallback *ArtistFormCallback) {
	artistFormCallback = new(ArtistFormCallback)
	artistFormCallback.probe = probe
	artistFormCallback.artist = artist
	artistFormCallback.formGroup = formGroup

	artistFormCallback.CreationMode = (artist == nil)

	return
}

type ArtistFormCallback struct {
	artist *models.Artist

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (artistFormCallback *ArtistFormCallback) OnSave() {

	// log.Println("ArtistFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	artistFormCallback.probe.formStage.Checkout()

	if artistFormCallback.artist == nil {
		artistFormCallback.artist = new(models.Artist).Stage(artistFormCallback.probe.stageOfInterest)
	}
	artist_ := artistFormCallback.artist
	_ = artist_

	for _, formDiv := range artistFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(artist_.Name), formDiv)
		case "IsDead":
			FormDivBasicFieldToField(&(artist_.IsDead), formDiv)
		case "DateOfDeath":
			FormDivBasicFieldToField(&(artist_.DateOfDeath), formDiv)
		case "Place":
			FormDivSelectFieldToField(&(artist_.Place), artistFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if artistFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artist_.Unstage(artistFormCallback.probe.stageOfInterest)
	}

	artistFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Artist](
		artistFormCallback.probe,
	)

	// display a new form by reset the form stage
	if artistFormCallback.CreationMode || artistFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artistFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(artistFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArtistFormCallback(
			nil,
			artistFormCallback.probe,
			newFormGroup,
		)
		artist := new(models.Artist)
		FillUpForm(artist, newFormGroup, artistFormCallback.probe)
		artistFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(artistFormCallback.probe)
}
func __gong__New__ArtistShapeFormCallback(
	artistshape *models.ArtistShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (artistshapeFormCallback *ArtistShapeFormCallback) {
	artistshapeFormCallback = new(ArtistShapeFormCallback)
	artistshapeFormCallback.probe = probe
	artistshapeFormCallback.artistshape = artistshape
	artistshapeFormCallback.formGroup = formGroup

	artistshapeFormCallback.CreationMode = (artistshape == nil)

	return
}

type ArtistShapeFormCallback struct {
	artistshape *models.ArtistShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (artistshapeFormCallback *ArtistShapeFormCallback) OnSave() {

	// log.Println("ArtistShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	artistshapeFormCallback.probe.formStage.Checkout()

	if artistshapeFormCallback.artistshape == nil {
		artistshapeFormCallback.artistshape = new(models.ArtistShape).Stage(artistshapeFormCallback.probe.stageOfInterest)
	}
	artistshape_ := artistshapeFormCallback.artistshape
	_ = artistshape_

	for _, formDiv := range artistshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(artistshape_.Name), formDiv)
		case "Artist":
			FormDivSelectFieldToField(&(artistshape_.Artist), artistshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(artistshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(artistshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(artistshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(artistshape_.Height), formDiv)
		case "Diagram:ArtistShapes":
			// WARNING : this form deals with the N-N association "Diagram.ArtistShapes []*ArtistShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ArtistShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "ArtistShapes"
				formerAssociationSource := artistshape_.GongGetReverseFieldOwner(
					artistshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.ArtistShapes []*ArtistShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ArtistShapes, artistshape_)
					formerSource.ArtistShapes = slices.Delete(formerSource.ArtistShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](artistshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.ArtistShapes []*ArtistShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ArtistShapes = append(newSource.ArtistShapes, artistshape_)
		}
	}

	// manage the suppress operation
	if artistshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artistshape_.Unstage(artistshapeFormCallback.probe.stageOfInterest)
	}

	artistshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ArtistShape](
		artistshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if artistshapeFormCallback.CreationMode || artistshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		artistshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(artistshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArtistShapeFormCallback(
			nil,
			artistshapeFormCallback.probe,
			newFormGroup,
		)
		artistshape := new(models.ArtistShape)
		FillUpForm(artistshape, newFormGroup, artistshapeFormCallback.probe)
		artistshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(artistshapeFormCallback.probe)
}
func __gong__New__ControlPointShapeFormCallback(
	controlpointshape *models.ControlPointShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (controlpointshapeFormCallback *ControlPointShapeFormCallback) {
	controlpointshapeFormCallback = new(ControlPointShapeFormCallback)
	controlpointshapeFormCallback.probe = probe
	controlpointshapeFormCallback.controlpointshape = controlpointshape
	controlpointshapeFormCallback.formGroup = formGroup

	controlpointshapeFormCallback.CreationMode = (controlpointshape == nil)

	return
}

type ControlPointShapeFormCallback struct {
	controlpointshape *models.ControlPointShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (controlpointshapeFormCallback *ControlPointShapeFormCallback) OnSave() {

	// log.Println("ControlPointShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlpointshapeFormCallback.probe.formStage.Checkout()

	if controlpointshapeFormCallback.controlpointshape == nil {
		controlpointshapeFormCallback.controlpointshape = new(models.ControlPointShape).Stage(controlpointshapeFormCallback.probe.stageOfInterest)
	}
	controlpointshape_ := controlpointshapeFormCallback.controlpointshape
	_ = controlpointshape_

	for _, formDiv := range controlpointshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlpointshape_.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.Y_Relative), formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&(controlpointshape_.IsStartShapeTheClosestShape), formDiv)
		case "InfluenceShape:ControlPointShapes":
			// WARNING : this form deals with the N-N association "InfluenceShape.ControlPointShapes []*ControlPointShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ControlPointShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.InfluenceShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "InfluenceShape"
				rf.Fieldname = "ControlPointShapes"
				formerAssociationSource := controlpointshape_.GongGetReverseFieldOwner(
					controlpointshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.InfluenceShape)
					if !ok {
						log.Fatalln("Source of InfluenceShape.ControlPointShapes []*ControlPointShape, is not an InfluenceShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ControlPointShapes, controlpointshape_)
					formerSource.ControlPointShapes = slices.Delete(formerSource.ControlPointShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.InfluenceShape
			for _influenceshape := range *models.GetGongstructInstancesSet[models.InfluenceShape](controlpointshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _influenceshape.GetName() == newSourceName.GetName() {
					newSource = _influenceshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of InfluenceShape.ControlPointShapes []*ControlPointShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ControlPointShapes = append(newSource.ControlPointShapes, controlpointshape_)
		}
	}

	// manage the suppress operation
	if controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshape_.Unstage(controlpointshapeFormCallback.probe.stageOfInterest)
	}

	controlpointshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlPointShape](
		controlpointshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlpointshapeFormCallback.CreationMode || controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(controlpointshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			controlpointshapeFormCallback.probe,
			newFormGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		FillUpForm(controlpointshape, newFormGroup, controlpointshapeFormCallback.probe)
		controlpointshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(controlpointshapeFormCallback.probe)
}
func __gong__New__DeskFormCallback(
	desk *models.Desk,
	probe *Probe,
	formGroup *table.FormGroup,
) (deskFormCallback *DeskFormCallback) {
	deskFormCallback = new(DeskFormCallback)
	deskFormCallback.probe = probe
	deskFormCallback.desk = desk
	deskFormCallback.formGroup = formGroup

	deskFormCallback.CreationMode = (desk == nil)

	return
}

type DeskFormCallback struct {
	desk *models.Desk

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (deskFormCallback *DeskFormCallback) OnSave() {

	// log.Println("DeskFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	deskFormCallback.probe.formStage.Checkout()

	if deskFormCallback.desk == nil {
		deskFormCallback.desk = new(models.Desk).Stage(deskFormCallback.probe.stageOfInterest)
	}
	desk_ := deskFormCallback.desk
	_ = desk_

	for _, formDiv := range deskFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(desk_.Name), formDiv)
		case "SelectedDiagram":
			FormDivSelectFieldToField(&(desk_.SelectedDiagram), deskFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if deskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		desk_.Unstage(deskFormCallback.probe.stageOfInterest)
	}

	deskFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Desk](
		deskFormCallback.probe,
	)

	// display a new form by reset the form stage
	if deskFormCallback.CreationMode || deskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deskFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(deskFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DeskFormCallback(
			nil,
			deskFormCallback.probe,
			newFormGroup,
		)
		desk := new(models.Desk)
		FillUpForm(desk, newFormGroup, deskFormCallback.probe)
		deskFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(deskFormCallback.probe)
}
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *table.FormGroup,
) (diagramFormCallback *DiagramFormCallback) {
	diagramFormCallback = new(DiagramFormCallback)
	diagramFormCallback.probe = probe
	diagramFormCallback.diagram = diagram
	diagramFormCallback.formGroup = formGroup

	diagramFormCallback.CreationMode = (diagram == nil)

	return
}

type DiagramFormCallback struct {
	diagram *models.Diagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {

	// log.Println("DiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramFormCallback.probe.formStage.Checkout()

	if diagramFormCallback.diagram == nil {
		diagramFormCallback.diagram = new(models.Diagram).Stage(diagramFormCallback.probe.stageOfInterest)
	}
	diagram_ := diagramFormCallback.diagram
	_ = diagram_

	for _, formDiv := range diagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagram_.Name), formDiv)
		case "MovementShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.MovementShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.MovementShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.MovementShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.MovementShapes = instanceSlice

		case "ArtefactTypeShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ArtefactTypeShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ArtefactTypeShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ArtefactTypeShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.ArtefactTypeShapes = instanceSlice

		case "ArtistShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ArtistShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ArtistShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ArtistShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.ArtistShapes = instanceSlice

		case "InfluenceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.InfluenceShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.InfluenceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.InfluenceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.InfluenceShapes = instanceSlice

		case "IsEditable":
			FormDivBasicFieldToField(&(diagram_.IsEditable), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsNodeExpanded), formDiv)
		case "IsMovementCategoryNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsMovementCategoryNodeExpanded), formDiv)
		case "IsArtefactTypeCategoryNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsArtefactTypeCategoryNodeExpanded), formDiv)
		case "IsArtistCategoryNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsArtistCategoryNodeExpanded), formDiv)
		case "IsInfluenceCategoryNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsInfluenceCategoryNodeExpanded), formDiv)
		case "IsMovementCategoryShown":
			FormDivBasicFieldToField(&(diagram_.IsMovementCategoryShown), formDiv)
		case "IsArtefactTypeCategoryShown":
			FormDivBasicFieldToField(&(diagram_.IsArtefactTypeCategoryShown), formDiv)
		case "IsArtistCategoryShown":
			FormDivBasicFieldToField(&(diagram_.IsArtistCategoryShown), formDiv)
		case "IsInfluenceCategoryShown":
			FormDivBasicFieldToField(&(diagram_.IsInfluenceCategoryShown), formDiv)
		case "XMargin":
			FormDivBasicFieldToField(&(diagram_.XMargin), formDiv)
		case "YMargin":
			FormDivBasicFieldToField(&(diagram_.YMargin), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagram_.Height), formDiv)
		case "NextVerticalDateXMargin":
			FormDivBasicFieldToField(&(diagram_.NextVerticalDateXMargin), formDiv)
		case "RedColorCode":
			FormDivBasicFieldToField(&(diagram_.RedColorCode), formDiv)
		case "BackgroundGreyColorCode":
			FormDivBasicFieldToField(&(diagram_.BackgroundGreyColorCode), formDiv)
		case "GrayColorCode":
			FormDivBasicFieldToField(&(diagram_.GrayColorCode), formDiv)
		case "BottomBoxYOffset":
			FormDivBasicFieldToField(&(diagram_.BottomBoxYOffset), formDiv)
		case "BottomBoxWidth":
			FormDivBasicFieldToField(&(diagram_.BottomBoxWidth), formDiv)
		case "BottomBoxHeigth":
			FormDivBasicFieldToField(&(diagram_.BottomBoxHeigth), formDiv)
		case "BottomBoxFontSize":
			FormDivBasicFieldToField(&(diagram_.BottomBoxFontSize), formDiv)
		case "BottomBoxFontWeigth":
			FormDivBasicFieldToField(&(diagram_.BottomBoxFontWeigth), formDiv)
		case "BottomBoxFontFamily":
			FormDivBasicFieldToField(&(diagram_.BottomBoxFontFamily), formDiv)
		case "BottomBoxLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.BottomBoxLetterSpacing), formDiv)
		case "BottomBoxLetterColorCode":
			FormDivBasicFieldToField(&(diagram_.BottomBoxLetterColorCode), formDiv)
		case "MovementRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.MovementRectAnchorType), formDiv)
		case "MovementTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.MovementTextAnchorType), formDiv)
		case "MovementDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.MovementDominantBaselineType), formDiv)
		case "MovementFontSize":
			FormDivBasicFieldToField(&(diagram_.MovementFontSize), formDiv)
		case "MajorMovementFontSize":
			FormDivBasicFieldToField(&(diagram_.MajorMovementFontSize), formDiv)
		case "MinorMovementFontSize":
			FormDivBasicFieldToField(&(diagram_.MinorMovementFontSize), formDiv)
		case "MovementFontWeigth":
			FormDivBasicFieldToField(&(diagram_.MovementFontWeigth), formDiv)
		case "MovementFontFamily":
			FormDivBasicFieldToField(&(diagram_.MovementFontFamily), formDiv)
		case "MovementLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.MovementLetterSpacing), formDiv)
		case "AbstractMovementFontSize":
			FormDivBasicFieldToField(&(diagram_.AbstractMovementFontSize), formDiv)
		case "AbstractMovementRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.AbstractMovementRectAnchorType), formDiv)
		case "AbstractMovementTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.AbstractMovementTextAnchorType), formDiv)
		case "AbstractDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.AbstractDominantBaselineType), formDiv)
		case "MovementDateRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.MovementDateRectAnchorType), formDiv)
		case "MovementDateTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.MovementDateTextAnchorType), formDiv)
		case "MovementDateTextDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.MovementDateTextDominantBaselineType), formDiv)
		case "MovementDateAndPlacesFontSize":
			FormDivBasicFieldToField(&(diagram_.MovementDateAndPlacesFontSize), formDiv)
		case "MovementDateAndPlacesFontWeigth":
			FormDivBasicFieldToField(&(diagram_.MovementDateAndPlacesFontWeigth), formDiv)
		case "MovementDateAndPlacesFontFamily":
			FormDivBasicFieldToField(&(diagram_.MovementDateAndPlacesFontFamily), formDiv)
		case "MovementDateAndPlacesLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.MovementDateAndPlacesLetterSpacing), formDiv)
		case "MovementBelowArcY_Offset":
			FormDivBasicFieldToField(&(diagram_.MovementBelowArcY_Offset), formDiv)
		case "MovementBelowArcY_OffsetPerPlace":
			FormDivBasicFieldToField(&(diagram_.MovementBelowArcY_OffsetPerPlace), formDiv)
		case "MovementPlacesRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.MovementPlacesRectAnchorType), formDiv)
		case "MovementPlacesTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.MovementPlacesTextAnchorType), formDiv)
		case "MovementPlacesDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.MovementPlacesDominantBaselineType), formDiv)
		case "ArtefactTypeFontSize":
			FormDivBasicFieldToField(&(diagram_.ArtefactTypeFontSize), formDiv)
		case "ArtefactTypeFontWeigth":
			FormDivBasicFieldToField(&(diagram_.ArtefactTypeFontWeigth), formDiv)
		case "ArtefactTypeFontFamily":
			FormDivBasicFieldToField(&(diagram_.ArtefactTypeFontFamily), formDiv)
		case "ArtefactTypeLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.ArtefactTypeLetterSpacing), formDiv)
		case "ArtefactTypeRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtefactTypeRectAnchorType), formDiv)
		case "ArtefactDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.ArtefactDominantBaselineType), formDiv)
		case "ArtefactTypeStrokeWidth":
			FormDivBasicFieldToField(&(diagram_.ArtefactTypeStrokeWidth), formDiv)
		case "ArtistRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistRectAnchorType), formDiv)
		case "ArtistTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistTextAnchorType), formDiv)
		case "ArtistDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistDominantBaselineType), formDiv)
		case "ArtistFontSize":
			FormDivBasicFieldToField(&(diagram_.ArtistFontSize), formDiv)
		case "MajorArtistFontSize":
			FormDivBasicFieldToField(&(diagram_.MajorArtistFontSize), formDiv)
		case "MinorArtistFontSize":
			FormDivBasicFieldToField(&(diagram_.MinorArtistFontSize), formDiv)
		case "ArtistFontWeigth":
			FormDivBasicFieldToField(&(diagram_.ArtistFontWeigth), formDiv)
		case "ArtistFontFamily":
			FormDivBasicFieldToField(&(diagram_.ArtistFontFamily), formDiv)
		case "ArtistLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.ArtistLetterSpacing), formDiv)
		case "ArtistDateRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistDateRectAnchorType), formDiv)
		case "ArtistDateTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistDateTextAnchorType), formDiv)
		case "ArtistDateDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistDateDominantBaselineType), formDiv)
		case "ArtistDateAndPlacesFontSize":
			FormDivBasicFieldToField(&(diagram_.ArtistDateAndPlacesFontSize), formDiv)
		case "ArtistDateAndPlacesFontWeigth":
			FormDivBasicFieldToField(&(diagram_.ArtistDateAndPlacesFontWeigth), formDiv)
		case "ArtistDateAndPlacesFontFamily":
			FormDivBasicFieldToField(&(diagram_.ArtistDateAndPlacesFontFamily), formDiv)
		case "ArtistDateAndPlacesLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.ArtistDateAndPlacesLetterSpacing), formDiv)
		case "ArtistPlacesRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistPlacesRectAnchorType), formDiv)
		case "ArtistPlacesTextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistPlacesTextAnchorType), formDiv)
		case "ArtistPlacesDominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.ArtistPlacesDominantBaselineType), formDiv)
		case "InfluenceArrowSize":
			FormDivBasicFieldToField(&(diagram_.InfluenceArrowSize), formDiv)
		case "InfluenceArrowStartOffset":
			FormDivBasicFieldToField(&(diagram_.InfluenceArrowStartOffset), formDiv)
		case "InfluenceArrowEndOffset":
			FormDivBasicFieldToField(&(diagram_.InfluenceArrowEndOffset), formDiv)
		case "InfluenceCornerRadius":
			FormDivBasicFieldToField(&(diagram_.InfluenceCornerRadius), formDiv)
		case "InfluenceDashedLinePattern":
			FormDivBasicFieldToField(&(diagram_.InfluenceDashedLinePattern), formDiv)
		}
	}

	// manage the suppress operation
	if diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagram_.Unstage(diagramFormCallback.probe.stageOfInterest)
	}

	diagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Diagram](
		diagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramFormCallback.CreationMode || diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(diagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			diagramFormCallback.probe,
			newFormGroup,
		)
		diagram := new(models.Diagram)
		FillUpForm(diagram, newFormGroup, diagramFormCallback.probe)
		diagramFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(diagramFormCallback.probe)
}
func __gong__New__InfluenceFormCallback(
	influence *models.Influence,
	probe *Probe,
	formGroup *table.FormGroup,
) (influenceFormCallback *InfluenceFormCallback) {
	influenceFormCallback = new(InfluenceFormCallback)
	influenceFormCallback.probe = probe
	influenceFormCallback.influence = influence
	influenceFormCallback.formGroup = formGroup

	influenceFormCallback.CreationMode = (influence == nil)

	return
}

type InfluenceFormCallback struct {
	influence *models.Influence

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (influenceFormCallback *InfluenceFormCallback) OnSave() {

	// log.Println("InfluenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	influenceFormCallback.probe.formStage.Checkout()

	if influenceFormCallback.influence == nil {
		influenceFormCallback.influence = new(models.Influence).Stage(influenceFormCallback.probe.stageOfInterest)
	}
	influence_ := influenceFormCallback.influence
	_ = influence_

	for _, formDiv := range influenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(influence_.Name), formDiv)
		case "SourceMovement":
			FormDivSelectFieldToField(&(influence_.SourceMovement), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "SourceArtefactType":
			FormDivSelectFieldToField(&(influence_.SourceArtefactType), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "SourceArtist":
			FormDivSelectFieldToField(&(influence_.SourceArtist), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "TargetMovement":
			FormDivSelectFieldToField(&(influence_.TargetMovement), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "TargetArtefactType":
			FormDivSelectFieldToField(&(influence_.TargetArtefactType), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "TargetArtist":
			FormDivSelectFieldToField(&(influence_.TargetArtist), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "IsHypothtical":
			FormDivBasicFieldToField(&(influence_.IsHypothtical), formDiv)
		}
	}

	// manage the suppress operation
	if influenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influence_.Unstage(influenceFormCallback.probe.stageOfInterest)
	}

	influenceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Influence](
		influenceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if influenceFormCallback.CreationMode || influenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(influenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InfluenceFormCallback(
			nil,
			influenceFormCallback.probe,
			newFormGroup,
		)
		influence := new(models.Influence)
		FillUpForm(influence, newFormGroup, influenceFormCallback.probe)
		influenceFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(influenceFormCallback.probe)
}
func __gong__New__InfluenceShapeFormCallback(
	influenceshape *models.InfluenceShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (influenceshapeFormCallback *InfluenceShapeFormCallback) {
	influenceshapeFormCallback = new(InfluenceShapeFormCallback)
	influenceshapeFormCallback.probe = probe
	influenceshapeFormCallback.influenceshape = influenceshape
	influenceshapeFormCallback.formGroup = formGroup

	influenceshapeFormCallback.CreationMode = (influenceshape == nil)

	return
}

type InfluenceShapeFormCallback struct {
	influenceshape *models.InfluenceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (influenceshapeFormCallback *InfluenceShapeFormCallback) OnSave() {

	// log.Println("InfluenceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	influenceshapeFormCallback.probe.formStage.Checkout()

	if influenceshapeFormCallback.influenceshape == nil {
		influenceshapeFormCallback.influenceshape = new(models.InfluenceShape).Stage(influenceshapeFormCallback.probe.stageOfInterest)
	}
	influenceshape_ := influenceshapeFormCallback.influenceshape
	_ = influenceshape_

	for _, formDiv := range influenceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(influenceshape_.Name), formDiv)
		case "Influence":
			FormDivSelectFieldToField(&(influenceshape_.Influence), influenceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](influenceshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					influenceshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			influenceshape_.ControlPointShapes = instanceSlice

		case "Diagram:InfluenceShapes":
			// WARNING : this form deals with the N-N association "Diagram.InfluenceShapes []*InfluenceShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of InfluenceShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "InfluenceShapes"
				formerAssociationSource := influenceshape_.GongGetReverseFieldOwner(
					influenceshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.InfluenceShapes []*InfluenceShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.InfluenceShapes, influenceshape_)
					formerSource.InfluenceShapes = slices.Delete(formerSource.InfluenceShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](influenceshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.InfluenceShapes []*InfluenceShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.InfluenceShapes = append(newSource.InfluenceShapes, influenceshape_)
		}
	}

	// manage the suppress operation
	if influenceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influenceshape_.Unstage(influenceshapeFormCallback.probe.stageOfInterest)
	}

	influenceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.InfluenceShape](
		influenceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if influenceshapeFormCallback.CreationMode || influenceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influenceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(influenceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InfluenceShapeFormCallback(
			nil,
			influenceshapeFormCallback.probe,
			newFormGroup,
		)
		influenceshape := new(models.InfluenceShape)
		FillUpForm(influenceshape, newFormGroup, influenceshapeFormCallback.probe)
		influenceshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(influenceshapeFormCallback.probe)
}
func __gong__New__MovementFormCallback(
	movement *models.Movement,
	probe *Probe,
	formGroup *table.FormGroup,
) (movementFormCallback *MovementFormCallback) {
	movementFormCallback = new(MovementFormCallback)
	movementFormCallback.probe = probe
	movementFormCallback.movement = movement
	movementFormCallback.formGroup = formGroup

	movementFormCallback.CreationMode = (movement == nil)

	return
}

type MovementFormCallback struct {
	movement *models.Movement

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (movementFormCallback *MovementFormCallback) OnSave() {

	// log.Println("MovementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	movementFormCallback.probe.formStage.Checkout()

	if movementFormCallback.movement == nil {
		movementFormCallback.movement = new(models.Movement).Stage(movementFormCallback.probe.stageOfInterest)
	}
	movement_ := movementFormCallback.movement
	_ = movement_

	for _, formDiv := range movementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(movement_.Name), formDiv)
		case "Date":
			FormDivBasicFieldToField(&(movement_.Date), formDiv)
		case "Places":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Place](movementFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Place, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Place)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					movementFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			movement_.Places = instanceSlice

		case "IsAbstract":
			FormDivBasicFieldToField(&(movement_.IsAbstract), formDiv)
		case "IsModern":
			FormDivBasicFieldToField(&(movement_.IsModern), formDiv)
		case "IsMajor":
			FormDivBasicFieldToField(&(movement_.IsMajor), formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&(movement_.IsMinor), formDiv)
		case "AdditionnalName":
			FormDivBasicFieldToField(&(movement_.AdditionnalName), formDiv)
		case "HideDate":
			FormDivBasicFieldToField(&(movement_.HideDate), formDiv)
		}
	}

	// manage the suppress operation
	if movementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		movement_.Unstage(movementFormCallback.probe.stageOfInterest)
	}

	movementFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Movement](
		movementFormCallback.probe,
	)

	// display a new form by reset the form stage
	if movementFormCallback.CreationMode || movementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		movementFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(movementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MovementFormCallback(
			nil,
			movementFormCallback.probe,
			newFormGroup,
		)
		movement := new(models.Movement)
		FillUpForm(movement, newFormGroup, movementFormCallback.probe)
		movementFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(movementFormCallback.probe)
}
func __gong__New__MovementShapeFormCallback(
	movementshape *models.MovementShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (movementshapeFormCallback *MovementShapeFormCallback) {
	movementshapeFormCallback = new(MovementShapeFormCallback)
	movementshapeFormCallback.probe = probe
	movementshapeFormCallback.movementshape = movementshape
	movementshapeFormCallback.formGroup = formGroup

	movementshapeFormCallback.CreationMode = (movementshape == nil)

	return
}

type MovementShapeFormCallback struct {
	movementshape *models.MovementShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (movementshapeFormCallback *MovementShapeFormCallback) OnSave() {

	// log.Println("MovementShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	movementshapeFormCallback.probe.formStage.Checkout()

	if movementshapeFormCallback.movementshape == nil {
		movementshapeFormCallback.movementshape = new(models.MovementShape).Stage(movementshapeFormCallback.probe.stageOfInterest)
	}
	movementshape_ := movementshapeFormCallback.movementshape
	_ = movementshape_

	for _, formDiv := range movementshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(movementshape_.Name), formDiv)
		case "Movement":
			FormDivSelectFieldToField(&(movementshape_.Movement), movementshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(movementshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(movementshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(movementshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(movementshape_.Height), formDiv)
		case "Diagram:MovementShapes":
			// WARNING : this form deals with the N-N association "Diagram.MovementShapes []*MovementShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of MovementShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "MovementShapes"
				formerAssociationSource := movementshape_.GongGetReverseFieldOwner(
					movementshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.MovementShapes []*MovementShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.MovementShapes, movementshape_)
					formerSource.MovementShapes = slices.Delete(formerSource.MovementShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](movementshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.MovementShapes []*MovementShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.MovementShapes = append(newSource.MovementShapes, movementshape_)
		}
	}

	// manage the suppress operation
	if movementshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		movementshape_.Unstage(movementshapeFormCallback.probe.stageOfInterest)
	}

	movementshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MovementShape](
		movementshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if movementshapeFormCallback.CreationMode || movementshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		movementshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(movementshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MovementShapeFormCallback(
			nil,
			movementshapeFormCallback.probe,
			newFormGroup,
		)
		movementshape := new(models.MovementShape)
		FillUpForm(movementshape, newFormGroup, movementshapeFormCallback.probe)
		movementshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(movementshapeFormCallback.probe)
}
func __gong__New__PlaceFormCallback(
	place *models.Place,
	probe *Probe,
	formGroup *table.FormGroup,
) (placeFormCallback *PlaceFormCallback) {
	placeFormCallback = new(PlaceFormCallback)
	placeFormCallback.probe = probe
	placeFormCallback.place = place
	placeFormCallback.formGroup = formGroup

	placeFormCallback.CreationMode = (place == nil)

	return
}

type PlaceFormCallback struct {
	place *models.Place

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (placeFormCallback *PlaceFormCallback) OnSave() {

	// log.Println("PlaceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	placeFormCallback.probe.formStage.Checkout()

	if placeFormCallback.place == nil {
		placeFormCallback.place = new(models.Place).Stage(placeFormCallback.probe.stageOfInterest)
	}
	place_ := placeFormCallback.place
	_ = place_

	for _, formDiv := range placeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(place_.Name), formDiv)
		case "Movement:Places":
			// WARNING : this form deals with the N-N association "Movement.Places []*Place" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Place). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Movement
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Movement"
				rf.Fieldname = "Places"
				formerAssociationSource := place_.GongGetReverseFieldOwner(
					placeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Movement)
					if !ok {
						log.Fatalln("Source of Movement.Places []*Place, is not an Movement instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Places, place_)
					formerSource.Places = slices.Delete(formerSource.Places, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Movement
			for _movement := range *models.GetGongstructInstancesSet[models.Movement](placeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _movement.GetName() == newSourceName.GetName() {
					newSource = _movement // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Movement.Places []*Place, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Places = append(newSource.Places, place_)
		}
	}

	// manage the suppress operation
	if placeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		place_.Unstage(placeFormCallback.probe.stageOfInterest)
	}

	placeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Place](
		placeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if placeFormCallback.CreationMode || placeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		placeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(placeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlaceFormCallback(
			nil,
			placeFormCallback.probe,
			newFormGroup,
		)
		place := new(models.Place)
		FillUpForm(place, newFormGroup, placeFormCallback.probe)
		placeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(placeFormCallback.probe)
}

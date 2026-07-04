// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ArtefactTypeFormCallback(
	artefacttype *models.ArtefactType,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (artefacttypeFormCallback *ArtefactTypeFormCallback) OnSave() {
	artefacttypeFormCallback.probe.stageOfInterest.Lock()
	defer artefacttypeFormCallback.probe.stageOfInterest.Unlock()

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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(artefacttype_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(artefacttype_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(artefacttype_.LayoutDirection), formDiv)
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
		newFormGroup := (&form.FormGroup{
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

	artefacttypeFormCallback.probe.ux_tree()
}
func __gong__New__ArtefactTypeShapeFormCallback(
	artefacttypeshape *models.ArtefactTypeShape,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (artefacttypeshapeFormCallback *ArtefactTypeShapeFormCallback) OnSave() {
	artefacttypeshapeFormCallback.probe.stageOfInterest.Lock()
	defer artefacttypeshapeFormCallback.probe.stageOfInterest.Unlock()

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
		case "IsHidden":
			FormDivBasicFieldToField(&(artefacttypeshape_.IsHidden), formDiv)
		case "Diagram:ArtefactTypeShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](artefacttypeshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ArtefactTypeShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](artefacttypeshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(artefacttypeshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure artefacttypeshape_ is in _diagram.ArtefactTypeShapes
					found := false
					for _, _b := range _diagram.ArtefactTypeShapes {
						if _b == artefacttypeshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ArtefactTypeShapes = append(_diagram.ArtefactTypeShapes, artefacttypeshape_)
						artefacttypeshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ArtefactTypeShapes", &_diagram.ArtefactTypeShapes)
					}
				} else {
					// ensure artefacttypeshape_ is NOT in _diagram.ArtefactTypeShapes
					idx := slices.Index(_diagram.ArtefactTypeShapes, artefacttypeshape_)
					if idx != -1 {
						_diagram.ArtefactTypeShapes = slices.Delete(_diagram.ArtefactTypeShapes, idx, idx+1)
						artefacttypeshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ArtefactTypeShapes", &_diagram.ArtefactTypeShapes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	artefacttypeshapeFormCallback.probe.ux_tree()
}
func __gong__New__ArtistFormCallback(
	artist *models.Artist,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (artistFormCallback *ArtistFormCallback) OnSave() {
	artistFormCallback.probe.stageOfInterest.Lock()
	defer artistFormCallback.probe.stageOfInterest.Unlock()

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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(artist_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(artist_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(artist_.LayoutDirection), formDiv)
		case "IsDead":
			FormDivBasicFieldToField(&(artist_.IsDead), formDiv)
		case "DateOfDeath":
			FormDivTimeFieldToField(&(artist_.DateOfDeath), formDiv, false)
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
		newFormGroup := (&form.FormGroup{
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

	artistFormCallback.probe.ux_tree()
}
func __gong__New__ArtistShapeFormCallback(
	artistshape *models.ArtistShape,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (artistshapeFormCallback *ArtistShapeFormCallback) OnSave() {
	artistshapeFormCallback.probe.stageOfInterest.Lock()
	defer artistshapeFormCallback.probe.stageOfInterest.Unlock()

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
		case "IsHidden":
			FormDivBasicFieldToField(&(artistshape_.IsHidden), formDiv)
		case "ImagePng_X":
			FormDivBasicFieldToField(&(artistshape_.ImagePng_X), formDiv)
		case "ImagePng_Y":
			FormDivBasicFieldToField(&(artistshape_.ImagePng_Y), formDiv)
		case "ImagePng_Width":
			FormDivBasicFieldToField(&(artistshape_.ImagePng_Width), formDiv)
		case "ImagePng_Height":
			FormDivBasicFieldToField(&(artistshape_.ImagePng_Height), formDiv)
		case "ImagePng_X_Offset":
			FormDivBasicFieldToField(&(artistshape_.ImagePng_X_Offset), formDiv)
		case "ImagePng_Y_Offset":
			FormDivBasicFieldToField(&(artistshape_.ImagePng_Y_Offset), formDiv)
		case "ImagePng_RectAnchorType":
			FormDivEnumStringFieldToField(&(artistshape_.ImagePng_RectAnchorType), formDiv)
		case "ImagePngBase64Content":
			FormDivBasicFieldToField(&(artistshape_.ImagePngBase64Content), formDiv)
		case "Diagram:ArtistShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](artistshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ArtistShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](artistshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(artistshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure artistshape_ is in _diagram.ArtistShapes
					found := false
					for _, _b := range _diagram.ArtistShapes {
						if _b == artistshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ArtistShapes = append(_diagram.ArtistShapes, artistshape_)
						artistshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ArtistShapes", &_diagram.ArtistShapes)
					}
				} else {
					// ensure artistshape_ is NOT in _diagram.ArtistShapes
					idx := slices.Index(_diagram.ArtistShapes, artistshape_)
					if idx != -1 {
						_diagram.ArtistShapes = slices.Delete(_diagram.ArtistShapes, idx, idx+1)
						artistshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ArtistShapes", &_diagram.ArtistShapes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	artistshapeFormCallback.probe.ux_tree()
}
func __gong__New__ControlPointShapeFormCallback(
	controlpointshape *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (controlpointshapeFormCallback *ControlPointShapeFormCallback) OnSave() {
	controlpointshapeFormCallback.probe.stageOfInterest.Lock()
	defer controlpointshapeFormCallback.probe.stageOfInterest.Unlock()

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the InfluenceShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target InfluenceShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.InfluenceShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetInfluenceShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetInfluenceShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all InfluenceShape instances and update their ControlPointShapes slice
			for _influenceshape := range *models.GetGongstructInstancesSetFromPointerType[*models.InfluenceShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _influenceshape)
				
				// if InfluenceShape is selected
				if targetInfluenceShapeIDs[id] {
					// ensure controlpointshape_ is in _influenceshape.ControlPointShapes
					found := false
					for _, _b := range _influenceshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_influenceshape.ControlPointShapes = append(_influenceshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_influenceshape, "ControlPointShapes", &_influenceshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _influenceshape.ControlPointShapes
					idx := slices.Index(_influenceshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_influenceshape.ControlPointShapes = slices.Delete(_influenceshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_influenceshape, "ControlPointShapes", &_influenceshape.ControlPointShapes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	controlpointshapeFormCallback.probe.ux_tree()
}
func __gong__New__DeskFormCallback(
	desk *models.Desk,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (deskFormCallback *DeskFormCallback) OnSave() {
	deskFormCallback.probe.stageOfInterest.Lock()
	defer deskFormCallback.probe.stageOfInterest.Unlock()

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
		newFormGroup := (&form.FormGroup{
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

	deskFormCallback.probe.ux_tree()
}
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {
	diagramFormCallback.probe.stageOfInterest.Lock()
	defer diagramFormCallback.probe.stageOfInterest.Unlock()

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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagram_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(diagram_.LayoutDirection), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagram_.IsChecked), formDiv)
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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.MovementShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.MovementShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "MovementShapes", &diagram_.MovementShapes)

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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ArtefactTypeShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ArtefactTypeShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ArtefactTypeShapes", &diagram_.ArtefactTypeShapes)

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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ArtistShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ArtistShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ArtistShapes", &diagram_.ArtistShapes)

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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.InfluenceShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.InfluenceShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "InfluenceShapes", &diagram_.InfluenceShapes)

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
		case "IsMovementCategoryHidden":
			FormDivBasicFieldToField(&(diagram_.IsMovementCategoryHidden), formDiv)
		case "IsArtefactTypeCategoryHidden":
			FormDivBasicFieldToField(&(diagram_.IsArtefactTypeCategoryHidden), formDiv)
		case "IsArtistCategoryHidden":
			FormDivBasicFieldToField(&(diagram_.IsArtistCategoryHidden), formDiv)
		case "IsInfluenceCategoryHidden":
			FormDivBasicFieldToField(&(diagram_.IsInfluenceCategoryHidden), formDiv)
		case "StartDate":
			FormDivTimeFieldToField(&(diagram_.StartDate), formDiv, false)
		case "EndDate":
			FormDivTimeFieldToField(&(diagram_.EndDate), formDiv, false)
		case "NbYearsForIntervals":
			FormDivBasicFieldToField(&(diagram_.NbYearsForIntervals), formDiv)
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
		newFormGroup := (&form.FormGroup{
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

	diagramFormCallback.probe.ux_tree()
}
func __gong__New__InfluenceFormCallback(
	influence *models.Influence,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (influenceFormCallback *InfluenceFormCallback) OnSave() {
	influenceFormCallback.probe.stageOfInterest.Lock()
	defer influenceFormCallback.probe.stageOfInterest.Unlock()

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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(influence_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(influence_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(influence_.LayoutDirection), formDiv)
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
		newFormGroup := (&form.FormGroup{
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

	influenceFormCallback.probe.ux_tree()
}
func __gong__New__InfluenceShapeFormCallback(
	influenceshape *models.InfluenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (influenceshapeFormCallback *InfluenceShapeFormCallback) OnSave() {
	influenceshapeFormCallback.probe.stageOfInterest.Lock()
	defer influenceshapeFormCallback.probe.stageOfInterest.Unlock()

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
		case "IsHidden":
			FormDivBasicFieldToField(&(influenceshape_.IsHidden), formDiv)
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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](influenceshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			influenceshape_.ControlPointShapes = instanceSlice
			influenceshapeFormCallback.probe.UpdateSliceOfPointersCallback(influenceshape_, "ControlPointShapes", &influenceshape_.ControlPointShapes)

		case "Diagram:InfluenceShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](influenceshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their InfluenceShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](influenceshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(influenceshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure influenceshape_ is in _diagram.InfluenceShapes
					found := false
					for _, _b := range _diagram.InfluenceShapes {
						if _b == influenceshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.InfluenceShapes = append(_diagram.InfluenceShapes, influenceshape_)
						influenceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "InfluenceShapes", &_diagram.InfluenceShapes)
					}
				} else {
					// ensure influenceshape_ is NOT in _diagram.InfluenceShapes
					idx := slices.Index(_diagram.InfluenceShapes, influenceshape_)
					if idx != -1 {
						_diagram.InfluenceShapes = slices.Delete(_diagram.InfluenceShapes, idx, idx+1)
						influenceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "InfluenceShapes", &_diagram.InfluenceShapes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	influenceshapeFormCallback.probe.ux_tree()
}
func __gong__New__LibraryFormCallback(
	library *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *LibraryFormCallback) {
	libraryFormCallback = new(LibraryFormCallback)
	libraryFormCallback.probe = probe
	libraryFormCallback.library = library
	libraryFormCallback.formGroup = formGroup

	libraryFormCallback.CreationMode = (library == nil)

	return
}

type LibraryFormCallback struct {
	library *models.Library

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (libraryFormCallback *LibraryFormCallback) OnSave() {
	libraryFormCallback.probe.stageOfInterest.Lock()
	defer libraryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LibraryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	libraryFormCallback.probe.formStage.Checkout()

	if libraryFormCallback.library == nil {
		libraryFormCallback.library = new(models.Library).Stage(libraryFormCallback.probe.stageOfInterest)
	}
	library_ := libraryFormCallback.library
	_ = library_

	for _, formDiv := range libraryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(library_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(library_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(library_.LayoutDirection), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
		case "SubLibraries":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibraries = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibraries", &library_.SubLibraries)

		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibrariesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibrariesWhoseNodeIsExpanded", &library_.SubLibrariesWhoseNodeIsExpanded)

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(library_.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibraries slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibraries
					found := false
					for _, _b := range _library.SubLibraries {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibraries = append(_library.SubLibraries, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibraries
					idx := slices.Index(_library.SubLibraries, library_)
					if idx != -1 {
						_library.SubLibraries = slices.Delete(_library.SubLibraries, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				}
			}
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibrariesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibrariesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.SubLibrariesWhoseNodeIsExpanded {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibrariesWhoseNodeIsExpanded = append(_library.SubLibrariesWhoseNodeIsExpanded, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibrariesWhoseNodeIsExpanded", &_library.SubLibrariesWhoseNodeIsExpanded)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibrariesWhoseNodeIsExpanded
					idx := slices.Index(_library.SubLibrariesWhoseNodeIsExpanded, library_)
					if idx != -1 {
						_library.SubLibrariesWhoseNodeIsExpanded = slices.Delete(_library.SubLibrariesWhoseNodeIsExpanded, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibrariesWhoseNodeIsExpanded", &_library.SubLibrariesWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		library_.Unstage(libraryFormCallback.probe.stageOfInterest)
	}

	libraryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Library](
		libraryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if libraryFormCallback.CreationMode || libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		libraryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(libraryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			libraryFormCallback.probe,
			newFormGroup,
		)
		library := new(models.Library)
		FillUpForm(library, newFormGroup, libraryFormCallback.probe)
		libraryFormCallback.probe.formStage.Commit()
	}

	libraryFormCallback.probe.ux_tree()
}
func __gong__New__MovementFormCallback(
	movement *models.Movement,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (movementFormCallback *MovementFormCallback) OnSave() {
	movementFormCallback.probe.stageOfInterest.Lock()
	defer movementFormCallback.probe.stageOfInterest.Unlock()

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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(movement_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(movement_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(movement_.LayoutDirection), formDiv)
		case "Date":
			FormDivTimeFieldToField(&(movement_.Date), formDiv, false)
		case "HideDate":
			FormDivBasicFieldToField(&(movement_.HideDate), formDiv)
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

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Place](movementFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			movement_.Places = instanceSlice
			movementFormCallback.probe.UpdateSliceOfPointersCallback(movement_, "Places", &movement_.Places)

		case "HasTaxonomicFilter":
			FormDivBasicFieldToField(&(movement_.HasTaxonomicFilter), formDiv)
		case "TaxonomicFilter":
			FormDivBasicFieldToField(&(movement_.TaxonomicFilter), formDiv)
		case "IsFeatured":
			FormDivBasicFieldToField(&(movement_.IsFeatured), formDiv)
		case "FeaturePrefix":
			FormDivBasicFieldToField(&(movement_.FeaturePrefix), formDiv)
		case "IsMajor":
			FormDivBasicFieldToField(&(movement_.IsMajor), formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&(movement_.IsMinor), formDiv)
		case "AdditionnalName":
			FormDivBasicFieldToField(&(movement_.AdditionnalName), formDiv)
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
		newFormGroup := (&form.FormGroup{
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

	movementFormCallback.probe.ux_tree()
}
func __gong__New__MovementShapeFormCallback(
	movementshape *models.MovementShape,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (movementshapeFormCallback *MovementShapeFormCallback) OnSave() {
	movementshapeFormCallback.probe.stageOfInterest.Lock()
	defer movementshapeFormCallback.probe.stageOfInterest.Unlock()

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
		case "IsHidden":
			FormDivBasicFieldToField(&(movementshape_.IsHidden), formDiv)
		case "Diagram:MovementShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](movementshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their MovementShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](movementshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(movementshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure movementshape_ is in _diagram.MovementShapes
					found := false
					for _, _b := range _diagram.MovementShapes {
						if _b == movementshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.MovementShapes = append(_diagram.MovementShapes, movementshape_)
						movementshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "MovementShapes", &_diagram.MovementShapes)
					}
				} else {
					// ensure movementshape_ is NOT in _diagram.MovementShapes
					idx := slices.Index(_diagram.MovementShapes, movementshape_)
					if idx != -1 {
						_diagram.MovementShapes = slices.Delete(_diagram.MovementShapes, idx, idx+1)
						movementshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "MovementShapes", &_diagram.MovementShapes)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	movementshapeFormCallback.probe.ux_tree()
}
func __gong__New__PlaceFormCallback(
	place *models.Place,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (placeFormCallback *PlaceFormCallback) OnSave() {
	placeFormCallback.probe.stageOfInterest.Lock()
	defer placeFormCallback.probe.stageOfInterest.Unlock()

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
			// 1. Decode the AssociationStorage which contains the rowIDs of the Movement instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Movement instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Movement](placeFormCallback.probe.stageOfInterest)
			targetMovementIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetMovementIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Movement instances and update their Places slice
			for _movement := range *models.GetGongstructInstancesSetFromPointerType[*models.Movement](placeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(placeFormCallback.probe.stageOfInterest, _movement)
				
				// if Movement is selected
				if targetMovementIDs[id] {
					// ensure place_ is in _movement.Places
					found := false
					for _, _b := range _movement.Places {
						if _b == place_ {
							found = true
							break
						}
					}
					if !found {
						_movement.Places = append(_movement.Places, place_)
						placeFormCallback.probe.UpdateSliceOfPointersCallback(_movement, "Places", &_movement.Places)
					}
				} else {
					// ensure place_ is NOT in _movement.Places
					idx := slices.Index(_movement.Places, place_)
					if idx != -1 {
						_movement.Places = slices.Delete(_movement.Places, idx, idx+1)
						placeFormCallback.probe.UpdateSliceOfPointersCallback(_movement, "Places", &_movement.Places)
					}
				}
			}
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
		newFormGroup := (&form.FormGroup{
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

	placeFormCallback.probe.ux_tree()
}

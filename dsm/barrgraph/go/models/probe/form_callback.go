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

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	var zero T
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: instance == zero,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	var zero T
	if cb.Instance == zero {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__ArtefactTypeFormCallback(
	_instance *models.ArtefactType,
	probe *Probe,
	formGroup *form.FormGroup,
) (artefacttypeFormCallback *FormCallback[*models.ArtefactType]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArtefactTypeFields,
	)
}

type ArtefactTypeFormCallback = FormCallback[*models.ArtefactType]

func saveArtefactTypeFields(
	_instance *models.ArtefactType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		}
	}
}

func __gong__New__ArtefactTypeShapeFormCallback(
	_instance *models.ArtefactTypeShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (artefacttypeshapeFormCallback *FormCallback[*models.ArtefactTypeShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArtefactTypeShapeFields,
	)
}

type ArtefactTypeShapeFormCallback = FormCallback[*models.ArtefactTypeShape]

func saveArtefactTypeShapeFields(
	_instance *models.ArtefactTypeShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ArtefactType":
			FormDivSelectFieldToField(&(_instance.ArtefactType), probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:ArtefactTypeShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ArtefactTypeShapes", func(owner *models.Diagram) *[]*models.ArtefactTypeShape { return &owner.ArtefactTypeShapes })
		}
	}
}

func __gong__New__ArtistFormCallback(
	_instance *models.Artist,
	probe *Probe,
	formGroup *form.FormGroup,
) (artistFormCallback *FormCallback[*models.Artist]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArtistFields,
	)
}

type ArtistFormCallback = FormCallback[*models.Artist]

func saveArtistFields(
	_instance *models.Artist,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsDead":
			FormDivBasicFieldToField(&(_instance.IsDead), formDiv)
		case "DateOfDeath":
			FormDivTimeFieldToField(&(_instance.DateOfDeath), formDiv, false)
		case "Place":
			FormDivSelectFieldToField(&(_instance.Place), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__ArtistShapeFormCallback(
	_instance *models.ArtistShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (artistshapeFormCallback *FormCallback[*models.ArtistShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArtistShapeFields,
	)
}

type ArtistShapeFormCallback = FormCallback[*models.ArtistShape]

func saveArtistShapeFields(
	_instance *models.ArtistShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Artist":
			FormDivSelectFieldToField(&(_instance.Artist), probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "ImagePng_X":
			FormDivBasicFieldToField(&(_instance.ImagePng_X), formDiv)
		case "ImagePng_Y":
			FormDivBasicFieldToField(&(_instance.ImagePng_Y), formDiv)
		case "ImagePng_Width":
			FormDivBasicFieldToField(&(_instance.ImagePng_Width), formDiv)
		case "ImagePng_Height":
			FormDivBasicFieldToField(&(_instance.ImagePng_Height), formDiv)
		case "ImagePng_X_Offset":
			FormDivBasicFieldToField(&(_instance.ImagePng_X_Offset), formDiv)
		case "ImagePng_Y_Offset":
			FormDivBasicFieldToField(&(_instance.ImagePng_Y_Offset), formDiv)
		case "ImagePng_RectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ImagePng_RectAnchorType), formDiv)
		case "ImagePngBase64Content":
			FormDivBasicFieldToField(&(_instance.ImagePngBase64Content), formDiv)
		case "Diagram:ArtistShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ArtistShapes", func(owner *models.Diagram) *[]*models.ArtistShape { return &owner.ArtistShapes })
		}
	}
}

func __gong__New__ControlPointShapeFormCallback(
	_instance *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlpointshapeFormCallback *FormCallback[*models.ControlPointShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveControlPointShapeFields,
	)
}

type ControlPointShapeFormCallback = FormCallback[*models.ControlPointShape]

func saveControlPointShapeFields(
	_instance *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(_instance.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(_instance.Y_Relative), formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&(_instance.IsStartShapeTheClosestShape), formDiv)
		case "InfluenceShape:ControlPointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPointShapes", func(owner *models.InfluenceShape) *[]*models.ControlPointShape { return &owner.ControlPointShapes })
		}
	}
}

func __gong__New__DeskFormCallback(
	_instance *models.Desk,
	probe *Probe,
	formGroup *form.FormGroup,
) (deskFormCallback *FormCallback[*models.Desk]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDeskFields,
	)
}

type DeskFormCallback = FormCallback[*models.Desk]

func saveDeskFields(
	_instance *models.Desk,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "SelectedDiagram":
			FormDivSelectFieldToField(&(_instance.SelectedDiagram), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__DiagramFormCallback(
	_instance *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramFormCallback *FormCallback[*models.Diagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramFields,
	)
}

type DiagramFormCallback = FormCallback[*models.Diagram]

func saveDiagramFields(
	_instance *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "MovementShapes":
			FormDivSliceOfPointersToField(_instance, "MovementShapes", &(_instance.MovementShapes), formDiv, probe)
		case "ArtefactTypeShapes":
			FormDivSliceOfPointersToField(_instance, "ArtefactTypeShapes", &(_instance.ArtefactTypeShapes), formDiv, probe)
		case "ArtistShapes":
			FormDivSliceOfPointersToField(_instance, "ArtistShapes", &(_instance.ArtistShapes), formDiv, probe)
		case "InfluenceShapes":
			FormDivSliceOfPointersToField(_instance, "InfluenceShapes", &(_instance.InfluenceShapes), formDiv, probe)
		case "IsEditable":
			FormDivBasicFieldToField(&(_instance.IsEditable), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNodeExpanded), formDiv)
		case "IsMovementCategoryNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsMovementCategoryNodeExpanded), formDiv)
		case "IsArtefactTypeCategoryNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsArtefactTypeCategoryNodeExpanded), formDiv)
		case "IsArtistCategoryNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsArtistCategoryNodeExpanded), formDiv)
		case "IsInfluenceCategoryNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsInfluenceCategoryNodeExpanded), formDiv)
		case "IsMovementCategoryHidden":
			FormDivBasicFieldToField(&(_instance.IsMovementCategoryHidden), formDiv)
		case "IsArtefactTypeCategoryHidden":
			FormDivBasicFieldToField(&(_instance.IsArtefactTypeCategoryHidden), formDiv)
		case "IsArtistCategoryHidden":
			FormDivBasicFieldToField(&(_instance.IsArtistCategoryHidden), formDiv)
		case "IsInfluenceCategoryHidden":
			FormDivBasicFieldToField(&(_instance.IsInfluenceCategoryHidden), formDiv)
		case "StartDate":
			FormDivTimeFieldToField(&(_instance.StartDate), formDiv, false)
		case "EndDate":
			FormDivTimeFieldToField(&(_instance.EndDate), formDiv, false)
		case "NbYearsForIntervals":
			FormDivBasicFieldToField(&(_instance.NbYearsForIntervals), formDiv)
		case "XMargin":
			FormDivBasicFieldToField(&(_instance.XMargin), formDiv)
		case "YMargin":
			FormDivBasicFieldToField(&(_instance.YMargin), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "NextVerticalDateXMargin":
			FormDivBasicFieldToField(&(_instance.NextVerticalDateXMargin), formDiv)
		case "RedColorCode":
			FormDivBasicFieldToField(&(_instance.RedColorCode), formDiv)
		case "BackgroundGreyColorCode":
			FormDivBasicFieldToField(&(_instance.BackgroundGreyColorCode), formDiv)
		case "GrayColorCode":
			FormDivBasicFieldToField(&(_instance.GrayColorCode), formDiv)
		case "BottomBoxYOffset":
			FormDivBasicFieldToField(&(_instance.BottomBoxYOffset), formDiv)
		case "BottomBoxWidth":
			FormDivBasicFieldToField(&(_instance.BottomBoxWidth), formDiv)
		case "BottomBoxHeigth":
			FormDivBasicFieldToField(&(_instance.BottomBoxHeigth), formDiv)
		case "BottomBoxFontSize":
			FormDivBasicFieldToField(&(_instance.BottomBoxFontSize), formDiv)
		case "BottomBoxFontWeigth":
			FormDivBasicFieldToField(&(_instance.BottomBoxFontWeigth), formDiv)
		case "BottomBoxFontFamily":
			FormDivBasicFieldToField(&(_instance.BottomBoxFontFamily), formDiv)
		case "BottomBoxLetterSpacing":
			FormDivBasicFieldToField(&(_instance.BottomBoxLetterSpacing), formDiv)
		case "BottomBoxLetterColorCode":
			FormDivBasicFieldToField(&(_instance.BottomBoxLetterColorCode), formDiv)
		case "MovementRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.MovementRectAnchorType), formDiv)
		case "MovementTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.MovementTextAnchorType), formDiv)
		case "MovementDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.MovementDominantBaselineType), formDiv)
		case "MovementFontSize":
			FormDivBasicFieldToField(&(_instance.MovementFontSize), formDiv)
		case "MajorMovementFontSize":
			FormDivBasicFieldToField(&(_instance.MajorMovementFontSize), formDiv)
		case "MinorMovementFontSize":
			FormDivBasicFieldToField(&(_instance.MinorMovementFontSize), formDiv)
		case "MovementFontWeigth":
			FormDivBasicFieldToField(&(_instance.MovementFontWeigth), formDiv)
		case "MovementFontFamily":
			FormDivBasicFieldToField(&(_instance.MovementFontFamily), formDiv)
		case "MovementLetterSpacing":
			FormDivBasicFieldToField(&(_instance.MovementLetterSpacing), formDiv)
		case "AbstractMovementFontSize":
			FormDivBasicFieldToField(&(_instance.AbstractMovementFontSize), formDiv)
		case "AbstractMovementRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.AbstractMovementRectAnchorType), formDiv)
		case "AbstractMovementTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.AbstractMovementTextAnchorType), formDiv)
		case "AbstractDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.AbstractDominantBaselineType), formDiv)
		case "MovementDateRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.MovementDateRectAnchorType), formDiv)
		case "MovementDateTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.MovementDateTextAnchorType), formDiv)
		case "MovementDateTextDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.MovementDateTextDominantBaselineType), formDiv)
		case "MovementDateAndPlacesFontSize":
			FormDivBasicFieldToField(&(_instance.MovementDateAndPlacesFontSize), formDiv)
		case "MovementDateAndPlacesFontWeigth":
			FormDivBasicFieldToField(&(_instance.MovementDateAndPlacesFontWeigth), formDiv)
		case "MovementDateAndPlacesFontFamily":
			FormDivBasicFieldToField(&(_instance.MovementDateAndPlacesFontFamily), formDiv)
		case "MovementDateAndPlacesLetterSpacing":
			FormDivBasicFieldToField(&(_instance.MovementDateAndPlacesLetterSpacing), formDiv)
		case "MovementBelowArcY_Offset":
			FormDivBasicFieldToField(&(_instance.MovementBelowArcY_Offset), formDiv)
		case "MovementBelowArcY_OffsetPerPlace":
			FormDivBasicFieldToField(&(_instance.MovementBelowArcY_OffsetPerPlace), formDiv)
		case "MovementPlacesRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.MovementPlacesRectAnchorType), formDiv)
		case "MovementPlacesTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.MovementPlacesTextAnchorType), formDiv)
		case "MovementPlacesDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.MovementPlacesDominantBaselineType), formDiv)
		case "ArtefactTypeFontSize":
			FormDivBasicFieldToField(&(_instance.ArtefactTypeFontSize), formDiv)
		case "ArtefactTypeFontWeigth":
			FormDivBasicFieldToField(&(_instance.ArtefactTypeFontWeigth), formDiv)
		case "ArtefactTypeFontFamily":
			FormDivBasicFieldToField(&(_instance.ArtefactTypeFontFamily), formDiv)
		case "ArtefactTypeLetterSpacing":
			FormDivBasicFieldToField(&(_instance.ArtefactTypeLetterSpacing), formDiv)
		case "ArtefactTypeRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtefactTypeRectAnchorType), formDiv)
		case "ArtefactDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.ArtefactDominantBaselineType), formDiv)
		case "ArtefactTypeStrokeWidth":
			FormDivBasicFieldToField(&(_instance.ArtefactTypeStrokeWidth), formDiv)
		case "ArtistRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtistRectAnchorType), formDiv)
		case "ArtistTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtistTextAnchorType), formDiv)
		case "ArtistDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.ArtistDominantBaselineType), formDiv)
		case "ArtistFontSize":
			FormDivBasicFieldToField(&(_instance.ArtistFontSize), formDiv)
		case "MajorArtistFontSize":
			FormDivBasicFieldToField(&(_instance.MajorArtistFontSize), formDiv)
		case "MinorArtistFontSize":
			FormDivBasicFieldToField(&(_instance.MinorArtistFontSize), formDiv)
		case "ArtistFontWeigth":
			FormDivBasicFieldToField(&(_instance.ArtistFontWeigth), formDiv)
		case "ArtistFontFamily":
			FormDivBasicFieldToField(&(_instance.ArtistFontFamily), formDiv)
		case "ArtistLetterSpacing":
			FormDivBasicFieldToField(&(_instance.ArtistLetterSpacing), formDiv)
		case "ArtistDateRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtistDateRectAnchorType), formDiv)
		case "ArtistDateTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtistDateTextAnchorType), formDiv)
		case "ArtistDateDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.ArtistDateDominantBaselineType), formDiv)
		case "ArtistDateAndPlacesFontSize":
			FormDivBasicFieldToField(&(_instance.ArtistDateAndPlacesFontSize), formDiv)
		case "ArtistDateAndPlacesFontWeigth":
			FormDivBasicFieldToField(&(_instance.ArtistDateAndPlacesFontWeigth), formDiv)
		case "ArtistDateAndPlacesFontFamily":
			FormDivBasicFieldToField(&(_instance.ArtistDateAndPlacesFontFamily), formDiv)
		case "ArtistDateAndPlacesLetterSpacing":
			FormDivBasicFieldToField(&(_instance.ArtistDateAndPlacesLetterSpacing), formDiv)
		case "ArtistPlacesRectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtistPlacesRectAnchorType), formDiv)
		case "ArtistPlacesTextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.ArtistPlacesTextAnchorType), formDiv)
		case "ArtistPlacesDominantBaselineType":
			FormDivEnumStringFieldToField(&(_instance.ArtistPlacesDominantBaselineType), formDiv)
		case "InfluenceArrowSize":
			FormDivBasicFieldToField(&(_instance.InfluenceArrowSize), formDiv)
		case "InfluenceArrowStartOffset":
			FormDivBasicFieldToField(&(_instance.InfluenceArrowStartOffset), formDiv)
		case "InfluenceArrowEndOffset":
			FormDivBasicFieldToField(&(_instance.InfluenceArrowEndOffset), formDiv)
		case "InfluenceCornerRadius":
			FormDivBasicFieldToField(&(_instance.InfluenceCornerRadius), formDiv)
		case "InfluenceDashedLinePattern":
			FormDivBasicFieldToField(&(_instance.InfluenceDashedLinePattern), formDiv)
		}
	}
}

func __gong__New__InfluenceFormCallback(
	_instance *models.Influence,
	probe *Probe,
	formGroup *form.FormGroup,
) (influenceFormCallback *FormCallback[*models.Influence]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInfluenceFields,
	)
}

type InfluenceFormCallback = FormCallback[*models.Influence]

func saveInfluenceFields(
	_instance *models.Influence,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SourceMovement":
			FormDivSelectFieldToField(&(_instance.SourceMovement), probe.stageOfInterest, formDiv)
		case "SourceArtefactType":
			FormDivSelectFieldToField(&(_instance.SourceArtefactType), probe.stageOfInterest, formDiv)
		case "SourceArtist":
			FormDivSelectFieldToField(&(_instance.SourceArtist), probe.stageOfInterest, formDiv)
		case "TargetMovement":
			FormDivSelectFieldToField(&(_instance.TargetMovement), probe.stageOfInterest, formDiv)
		case "TargetArtefactType":
			FormDivSelectFieldToField(&(_instance.TargetArtefactType), probe.stageOfInterest, formDiv)
		case "TargetArtist":
			FormDivSelectFieldToField(&(_instance.TargetArtist), probe.stageOfInterest, formDiv)
		case "IsHypothtical":
			FormDivBasicFieldToField(&(_instance.IsHypothtical), formDiv)
		}
	}
}

func __gong__New__InfluenceShapeFormCallback(
	_instance *models.InfluenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (influenceshapeFormCallback *FormCallback[*models.InfluenceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInfluenceShapeFields,
	)
}

type InfluenceShapeFormCallback = FormCallback[*models.InfluenceShape]

func saveInfluenceShapeFields(
	_instance *models.InfluenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Influence":
			FormDivSelectFieldToField(&(_instance.Influence), probe.stageOfInterest, formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "ControlPointShapes":
			FormDivSliceOfPointersToField(_instance, "ControlPointShapes", &(_instance.ControlPointShapes), formDiv, probe)
		case "Diagram:InfluenceShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "InfluenceShapes", func(owner *models.Diagram) *[]*models.InfluenceShape { return &owner.InfluenceShapes })
		}
	}
}

func __gong__New__LibraryFormCallback(
	_instance *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *FormCallback[*models.Library]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLibraryFields,
	)
}

type LibraryFormCallback = FormCallback[*models.Library]

func saveLibraryFields(
	_instance *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(_instance.IsRootLibrary), formDiv)
		case "SubLibraries":
			FormDivSliceOfPointersToField(_instance, "SubLibraries", &(_instance.SubLibraries), formDiv, probe)
		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SubLibrariesWhoseNodeIsExpanded", &(_instance.SubLibrariesWhoseNodeIsExpanded), formDiv, probe)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(_instance.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(_instance.LogoSVGFile), formDiv)
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(_instance.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibrariesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Library { return &owner.SubLibrariesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__MovementFormCallback(
	_instance *models.Movement,
	probe *Probe,
	formGroup *form.FormGroup,
) (movementFormCallback *FormCallback[*models.Movement]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMovementFields,
	)
}

type MovementFormCallback = FormCallback[*models.Movement]

func saveMovementFields(
	_instance *models.Movement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Date":
			FormDivTimeFieldToField(&(_instance.Date), formDiv, false)
		case "HideDate":
			FormDivBasicFieldToField(&(_instance.HideDate), formDiv)
		case "Places":
			FormDivSliceOfPointersToField(_instance, "Places", &(_instance.Places), formDiv, probe)
		case "HasTaxonomicFilter":
			FormDivBasicFieldToField(&(_instance.HasTaxonomicFilter), formDiv)
		case "TaxonomicFilter":
			FormDivBasicFieldToField(&(_instance.TaxonomicFilter), formDiv)
		case "IsFeatured":
			FormDivBasicFieldToField(&(_instance.IsFeatured), formDiv)
		case "FeaturePrefix":
			FormDivBasicFieldToField(&(_instance.FeaturePrefix), formDiv)
		case "IsMajor":
			FormDivBasicFieldToField(&(_instance.IsMajor), formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&(_instance.IsMinor), formDiv)
		case "AdditionnalName":
			FormDivBasicFieldToField(&(_instance.AdditionnalName), formDiv)
		}
	}
}

func __gong__New__MovementShapeFormCallback(
	_instance *models.MovementShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (movementshapeFormCallback *FormCallback[*models.MovementShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMovementShapeFields,
	)
}

type MovementShapeFormCallback = FormCallback[*models.MovementShape]

func saveMovementShapeFields(
	_instance *models.MovementShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Movement":
			FormDivSelectFieldToField(&(_instance.Movement), probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:MovementShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "MovementShapes", func(owner *models.Diagram) *[]*models.MovementShape { return &owner.MovementShapes })
		}
	}
}

func __gong__New__PlaceFormCallback(
	_instance *models.Place,
	probe *Probe,
	formGroup *form.FormGroup,
) (placeFormCallback *FormCallback[*models.Place]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlaceFields,
	)
}

type PlaceFormCallback = FormCallback[*models.Place]

func savePlaceFields(
	_instance *models.Place,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Movement:Places":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Places", func(owner *models.Movement) *[]*models.Place { return &owner.Places })
		}
	}
}


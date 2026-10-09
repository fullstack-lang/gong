// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
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
	case *models.ArtefactType:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArtefactType_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ArtefactType_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ArtefactTypeShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArtefactTypeShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ArtefactTypeShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Artist:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Artist_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Artist_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ArtistShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ArtistShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ArtistShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.ControlPointShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_ControlPointShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_ControlPointShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Desk:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Desk_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Desk_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
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
	case *models.Influence:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Influence_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Influence_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.InfluenceShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_InfluenceShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_InfluenceShape_Stage(probe)
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
	case *models.Movement:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Movement_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Movement_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.MovementShape:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_MovementShape_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_MovementShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
	case *models.Place:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.Stage.Lock()
				defer probe.stageSet.Stage.Unlock()
				probe.formStage.Checkout()
				saveStageSet_Place_Stage(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.Stage)
				}
				probe.stageSet.Stage.Commit()
				updateStageSetTable_Place_Stage(probe)
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

func saveStageSet_ArtefactType_Stage(
	inst *models.ArtefactType,
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

func saveStageSet_ArtefactTypeShape_Stage(
	inst *models.ArtefactTypeShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ArtefactType":
			StageSetFormDivSelectFieldToField(&inst.ArtefactType, probe.stageSet.Stage.GetInstancesSet[*models.ArtefactType](), formDiv)
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

func saveStageSet_Artist_Stage(
	inst *models.Artist,
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
		case "IsDead":
			FormDivBasicFieldToField(&inst.IsDead, formDiv)
		case "DateOfDeath":
			FormDivTimeFieldToField(&inst.DateOfDeath, formDiv, false)
		case "Place":
			StageSetFormDivSelectFieldToField(&inst.Place, probe.stageSet.Stage.GetInstancesSet[*models.Place](), formDiv)
		}
	}
}

func saveStageSet_ArtistShape_Stage(
	inst *models.ArtistShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Artist":
			StageSetFormDivSelectFieldToField(&inst.Artist, probe.stageSet.Stage.GetInstancesSet[*models.Artist](), formDiv)
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
		case "ImagePng_X":
			FormDivBasicFieldToField(&inst.ImagePng_X, formDiv)
		case "ImagePng_Y":
			FormDivBasicFieldToField(&inst.ImagePng_Y, formDiv)
		case "ImagePng_Width":
			FormDivBasicFieldToField(&inst.ImagePng_Width, formDiv)
		case "ImagePng_Height":
			FormDivBasicFieldToField(&inst.ImagePng_Height, formDiv)
		case "ImagePng_X_Offset":
			FormDivBasicFieldToField(&inst.ImagePng_X_Offset, formDiv)
		case "ImagePng_Y_Offset":
			FormDivBasicFieldToField(&inst.ImagePng_Y_Offset, formDiv)
		case "ImagePng_RectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ImagePng_RectAnchorType, formDiv)
		case "ImagePngBase64Content":
			FormDivBasicFieldToField(&inst.ImagePngBase64Content, formDiv)
		}
	}
}

func saveStageSet_ControlPointShape_Stage(
	inst *models.ControlPointShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&inst.X_Relative, formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&inst.Y_Relative, formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&inst.IsStartShapeTheClosestShape, formDiv)
		}
	}
}

func saveStageSet_Desk_Stage(
	inst *models.Desk,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SelectedDiagram":
			StageSetFormDivSelectFieldToField(&inst.SelectedDiagram, probe.stageSet.Stage.GetInstancesSet[*models.Diagram](), formDiv)
		}
	}
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&inst.ComputedPrefix, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&inst.IsExpanded, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&inst.IsChecked, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&inst.IsEditable, formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&inst.IsNodeExpanded, formDiv)
		case "IsMovementCategoryNodeExpanded":
			FormDivBasicFieldToField(&inst.IsMovementCategoryNodeExpanded, formDiv)
		case "IsArtefactTypeCategoryNodeExpanded":
			FormDivBasicFieldToField(&inst.IsArtefactTypeCategoryNodeExpanded, formDiv)
		case "IsArtistCategoryNodeExpanded":
			FormDivBasicFieldToField(&inst.IsArtistCategoryNodeExpanded, formDiv)
		case "IsInfluenceCategoryNodeExpanded":
			FormDivBasicFieldToField(&inst.IsInfluenceCategoryNodeExpanded, formDiv)
		case "IsMovementCategoryHidden":
			FormDivBasicFieldToField(&inst.IsMovementCategoryHidden, formDiv)
		case "IsArtefactTypeCategoryHidden":
			FormDivBasicFieldToField(&inst.IsArtefactTypeCategoryHidden, formDiv)
		case "IsArtistCategoryHidden":
			FormDivBasicFieldToField(&inst.IsArtistCategoryHidden, formDiv)
		case "IsInfluenceCategoryHidden":
			FormDivBasicFieldToField(&inst.IsInfluenceCategoryHidden, formDiv)
		case "StartDate":
			FormDivTimeFieldToField(&inst.StartDate, formDiv, false)
		case "EndDate":
			FormDivTimeFieldToField(&inst.EndDate, formDiv, false)
		case "NbYearsForIntervals":
			FormDivBasicFieldToField(&inst.NbYearsForIntervals, formDiv)
		case "XMargin":
			FormDivBasicFieldToField(&inst.XMargin, formDiv)
		case "YMargin":
			FormDivBasicFieldToField(&inst.YMargin, formDiv)
		case "Height":
			FormDivBasicFieldToField(&inst.Height, formDiv)
		case "NextVerticalDateXMargin":
			FormDivBasicFieldToField(&inst.NextVerticalDateXMargin, formDiv)
		case "RedColorCode":
			FormDivBasicFieldToField(&inst.RedColorCode, formDiv)
		case "BackgroundGreyColorCode":
			FormDivBasicFieldToField(&inst.BackgroundGreyColorCode, formDiv)
		case "GrayColorCode":
			FormDivBasicFieldToField(&inst.GrayColorCode, formDiv)
		case "BottomBoxYOffset":
			FormDivBasicFieldToField(&inst.BottomBoxYOffset, formDiv)
		case "BottomBoxWidth":
			FormDivBasicFieldToField(&inst.BottomBoxWidth, formDiv)
		case "BottomBoxHeigth":
			FormDivBasicFieldToField(&inst.BottomBoxHeigth, formDiv)
		case "BottomBoxFontSize":
			FormDivBasicFieldToField(&inst.BottomBoxFontSize, formDiv)
		case "BottomBoxFontWeigth":
			FormDivBasicFieldToField(&inst.BottomBoxFontWeigth, formDiv)
		case "BottomBoxFontFamily":
			FormDivBasicFieldToField(&inst.BottomBoxFontFamily, formDiv)
		case "BottomBoxLetterSpacing":
			FormDivBasicFieldToField(&inst.BottomBoxLetterSpacing, formDiv)
		case "BottomBoxLetterColorCode":
			FormDivBasicFieldToField(&inst.BottomBoxLetterColorCode, formDiv)
		case "MovementRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementRectAnchorType, formDiv)
		case "MovementTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementTextAnchorType, formDiv)
		case "MovementDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementDominantBaselineType, formDiv)
		case "MovementFontSize":
			FormDivBasicFieldToField(&inst.MovementFontSize, formDiv)
		case "MajorMovementFontSize":
			FormDivBasicFieldToField(&inst.MajorMovementFontSize, formDiv)
		case "MinorMovementFontSize":
			FormDivBasicFieldToField(&inst.MinorMovementFontSize, formDiv)
		case "MovementFontWeigth":
			FormDivBasicFieldToField(&inst.MovementFontWeigth, formDiv)
		case "MovementFontFamily":
			FormDivBasicFieldToField(&inst.MovementFontFamily, formDiv)
		case "MovementLetterSpacing":
			FormDivBasicFieldToField(&inst.MovementLetterSpacing, formDiv)
		case "AbstractMovementFontSize":
			FormDivBasicFieldToField(&inst.AbstractMovementFontSize, formDiv)
		case "AbstractMovementRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.AbstractMovementRectAnchorType, formDiv)
		case "AbstractMovementTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.AbstractMovementTextAnchorType, formDiv)
		case "AbstractDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.AbstractDominantBaselineType, formDiv)
		case "MovementDateRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementDateRectAnchorType, formDiv)
		case "MovementDateTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementDateTextAnchorType, formDiv)
		case "MovementDateTextDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementDateTextDominantBaselineType, formDiv)
		case "MovementDateAndPlacesFontSize":
			FormDivBasicFieldToField(&inst.MovementDateAndPlacesFontSize, formDiv)
		case "MovementDateAndPlacesFontWeigth":
			FormDivBasicFieldToField(&inst.MovementDateAndPlacesFontWeigth, formDiv)
		case "MovementDateAndPlacesFontFamily":
			FormDivBasicFieldToField(&inst.MovementDateAndPlacesFontFamily, formDiv)
		case "MovementDateAndPlacesLetterSpacing":
			FormDivBasicFieldToField(&inst.MovementDateAndPlacesLetterSpacing, formDiv)
		case "MovementBelowArcY_Offset":
			FormDivBasicFieldToField(&inst.MovementBelowArcY_Offset, formDiv)
		case "MovementBelowArcY_OffsetPerPlace":
			FormDivBasicFieldToField(&inst.MovementBelowArcY_OffsetPerPlace, formDiv)
		case "MovementPlacesRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementPlacesRectAnchorType, formDiv)
		case "MovementPlacesTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementPlacesTextAnchorType, formDiv)
		case "MovementPlacesDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.MovementPlacesDominantBaselineType, formDiv)
		case "ArtefactTypeFontSize":
			FormDivBasicFieldToField(&inst.ArtefactTypeFontSize, formDiv)
		case "ArtefactTypeFontWeigth":
			FormDivBasicFieldToField(&inst.ArtefactTypeFontWeigth, formDiv)
		case "ArtefactTypeFontFamily":
			FormDivBasicFieldToField(&inst.ArtefactTypeFontFamily, formDiv)
		case "ArtefactTypeLetterSpacing":
			FormDivBasicFieldToField(&inst.ArtefactTypeLetterSpacing, formDiv)
		case "ArtefactTypeRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtefactTypeRectAnchorType, formDiv)
		case "ArtefactDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtefactDominantBaselineType, formDiv)
		case "ArtefactTypeStrokeWidth":
			FormDivBasicFieldToField(&inst.ArtefactTypeStrokeWidth, formDiv)
		case "ArtistRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistRectAnchorType, formDiv)
		case "ArtistTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistTextAnchorType, formDiv)
		case "ArtistDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistDominantBaselineType, formDiv)
		case "ArtistFontSize":
			FormDivBasicFieldToField(&inst.ArtistFontSize, formDiv)
		case "MajorArtistFontSize":
			FormDivBasicFieldToField(&inst.MajorArtistFontSize, formDiv)
		case "MinorArtistFontSize":
			FormDivBasicFieldToField(&inst.MinorArtistFontSize, formDiv)
		case "ArtistFontWeigth":
			FormDivBasicFieldToField(&inst.ArtistFontWeigth, formDiv)
		case "ArtistFontFamily":
			FormDivBasicFieldToField(&inst.ArtistFontFamily, formDiv)
		case "ArtistLetterSpacing":
			FormDivBasicFieldToField(&inst.ArtistLetterSpacing, formDiv)
		case "ArtistDateRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistDateRectAnchorType, formDiv)
		case "ArtistDateTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistDateTextAnchorType, formDiv)
		case "ArtistDateDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistDateDominantBaselineType, formDiv)
		case "ArtistDateAndPlacesFontSize":
			FormDivBasicFieldToField(&inst.ArtistDateAndPlacesFontSize, formDiv)
		case "ArtistDateAndPlacesFontWeigth":
			FormDivBasicFieldToField(&inst.ArtistDateAndPlacesFontWeigth, formDiv)
		case "ArtistDateAndPlacesFontFamily":
			FormDivBasicFieldToField(&inst.ArtistDateAndPlacesFontFamily, formDiv)
		case "ArtistDateAndPlacesLetterSpacing":
			FormDivBasicFieldToField(&inst.ArtistDateAndPlacesLetterSpacing, formDiv)
		case "ArtistPlacesRectAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistPlacesRectAnchorType, formDiv)
		case "ArtistPlacesTextAnchorType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistPlacesTextAnchorType, formDiv)
		case "ArtistPlacesDominantBaselineType":
			StageSetFormDivEnumStringFieldToField(&inst.ArtistPlacesDominantBaselineType, formDiv)
		case "InfluenceArrowSize":
			FormDivBasicFieldToField(&inst.InfluenceArrowSize, formDiv)
		case "InfluenceArrowStartOffset":
			FormDivBasicFieldToField(&inst.InfluenceArrowStartOffset, formDiv)
		case "InfluenceArrowEndOffset":
			FormDivBasicFieldToField(&inst.InfluenceArrowEndOffset, formDiv)
		case "InfluenceCornerRadius":
			FormDivBasicFieldToField(&inst.InfluenceCornerRadius, formDiv)
		case "InfluenceDashedLinePattern":
			FormDivBasicFieldToField(&inst.InfluenceDashedLinePattern, formDiv)
		}
	}
}

func saveStageSet_Influence_Stage(
	inst *models.Influence,
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
		case "SourceMovement":
			StageSetFormDivSelectFieldToField(&inst.SourceMovement, probe.stageSet.Stage.GetInstancesSet[*models.Movement](), formDiv)
		case "SourceArtefactType":
			StageSetFormDivSelectFieldToField(&inst.SourceArtefactType, probe.stageSet.Stage.GetInstancesSet[*models.ArtefactType](), formDiv)
		case "SourceArtist":
			StageSetFormDivSelectFieldToField(&inst.SourceArtist, probe.stageSet.Stage.GetInstancesSet[*models.Artist](), formDiv)
		case "TargetMovement":
			StageSetFormDivSelectFieldToField(&inst.TargetMovement, probe.stageSet.Stage.GetInstancesSet[*models.Movement](), formDiv)
		case "TargetArtefactType":
			StageSetFormDivSelectFieldToField(&inst.TargetArtefactType, probe.stageSet.Stage.GetInstancesSet[*models.ArtefactType](), formDiv)
		case "TargetArtist":
			StageSetFormDivSelectFieldToField(&inst.TargetArtist, probe.stageSet.Stage.GetInstancesSet[*models.Artist](), formDiv)
		case "IsHypothtical":
			FormDivBasicFieldToField(&inst.IsHypothtical, formDiv)
		}
	}
}

func saveStageSet_InfluenceShape_Stage(
	inst *models.InfluenceShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Influence":
			StageSetFormDivSelectFieldToField(&inst.Influence, probe.stageSet.Stage.GetInstancesSet[*models.Influence](), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&inst.IsHidden, formDiv)
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
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&inst.IsExpandedTmp, formDiv)
		}
	}
}

func saveStageSet_Movement_Stage(
	inst *models.Movement,
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
		case "Date":
			FormDivTimeFieldToField(&inst.Date, formDiv, false)
		case "HideDate":
			FormDivBasicFieldToField(&inst.HideDate, formDiv)
		case "HasTaxonomicFilter":
			FormDivBasicFieldToField(&inst.HasTaxonomicFilter, formDiv)
		case "TaxonomicFilter":
			FormDivBasicFieldToField(&inst.TaxonomicFilter, formDiv)
		case "IsFeatured":
			FormDivBasicFieldToField(&inst.IsFeatured, formDiv)
		case "FeaturePrefix":
			FormDivBasicFieldToField(&inst.FeaturePrefix, formDiv)
		case "IsMajor":
			FormDivBasicFieldToField(&inst.IsMajor, formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&inst.IsMinor, formDiv)
		case "AdditionnalName":
			FormDivBasicFieldToField(&inst.AdditionnalName, formDiv)
		}
	}
}

func saveStageSet_MovementShape_Stage(
	inst *models.MovementShape,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Movement":
			StageSetFormDivSelectFieldToField(&inst.Movement, probe.stageSet.Stage.GetInstancesSet[*models.Movement](), formDiv)
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

func saveStageSet_Place_Stage(
	inst *models.Place,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}


func StageSetNewInstance_ArtefactType_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ArtefactType",
	}).Stage(probe.formStage)
	inst := new(models.ArtefactType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ArtefactType_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ArtefactType_Stage(probe)
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

func StageSetNewInstance_ArtefactTypeShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ArtefactTypeShape",
	}).Stage(probe.formStage)
	inst := new(models.ArtefactTypeShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ArtefactTypeShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ArtefactTypeShape_Stage(probe)
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

func StageSetNewInstance_Artist_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Artist",
	}).Stage(probe.formStage)
	inst := new(models.Artist)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Artist_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Artist_Stage(probe)
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

func StageSetNewInstance_ArtistShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ArtistShape",
	}).Stage(probe.formStage)
	inst := new(models.ArtistShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ArtistShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ArtistShape_Stage(probe)
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

func StageSetNewInstance_ControlPointShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New ControlPointShape",
	}).Stage(probe.formStage)
	inst := new(models.ControlPointShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_ControlPointShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_ControlPointShape_Stage(probe)
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

func StageSetNewInstance_Desk_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Desk",
	}).Stage(probe.formStage)
	inst := new(models.Desk)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Desk_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Desk_Stage(probe)
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

func StageSetNewInstance_Influence_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Influence",
	}).Stage(probe.formStage)
	inst := new(models.Influence)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Influence_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Influence_Stage(probe)
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

func StageSetNewInstance_InfluenceShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New InfluenceShape",
	}).Stage(probe.formStage)
	inst := new(models.InfluenceShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_InfluenceShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_InfluenceShape_Stage(probe)
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

func StageSetNewInstance_Movement_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Movement",
	}).Stage(probe.formStage)
	inst := new(models.Movement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Movement_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Movement_Stage(probe)
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

func StageSetNewInstance_MovementShape_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New MovementShape",
	}).Stage(probe.formStage)
	inst := new(models.MovementShape)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_MovementShape_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_MovementShape_Stage(probe)
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

func StageSetNewInstance_Place_Stage(probe *StageSetProbe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  "Form",
		Label: "New Place",
	}).Stage(probe.formStage)
	inst := new(models.Place)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageSet.Stage.Lock()
			defer probe.stageSet.Stage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageSet.Stage)
			saveStageSet_Place_Stage(inst, probe, formGroup)
			probe.stageSet.Stage.Commit()
			updateStageSetTable_Place_Stage(probe)
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


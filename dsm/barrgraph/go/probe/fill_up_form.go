// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.ArtefactType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ArtefactTypeShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ArtefactType", instanceWithInferedType.ArtefactType, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ArtefactTypeShape](
				"Diagram",
				"ArtefactTypeShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ArtefactTypeShape {
					return owner.ArtefactTypeShapes
				})
		}

	case *models.Artist:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsDead", instanceWithInferedType.IsDead, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DateOfDeath", instanceWithInferedType.DateOfDeath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Place", instanceWithInferedType.Place, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ArtistShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Artist", instanceWithInferedType.Artist, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ImagePng_X", instanceWithInferedType.ImagePng_X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ImagePng_Y", instanceWithInferedType.ImagePng_Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ImagePng_Width", instanceWithInferedType.ImagePng_Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ImagePng_Height", instanceWithInferedType.ImagePng_Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ImagePng_X_Offset", instanceWithInferedType.ImagePng_X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ImagePng_Y_Offset", instanceWithInferedType.ImagePng_Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("ImagePng_RectAnchorType", instanceWithInferedType.ImagePng_RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ImagePngBase64Content", instanceWithInferedType.ImagePngBase64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ArtistShape](
				"Diagram",
				"ArtistShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ArtistShape {
					return owner.ArtistShapes
				})
		}

	case *models.ControlPointShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X_Relative", instanceWithInferedType.X_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y_Relative", instanceWithInferedType.Y_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsStartShapeTheClosestShape", instanceWithInferedType.IsStartShapeTheClosestShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.InfluenceShape, *models.ControlPointShape](
				"InfluenceShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.InfluenceShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}

	case *models.Desk:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("SelectedDiagram", instanceWithInferedType.SelectedDiagram, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("MovementShapes", instanceWithInferedType, &instanceWithInferedType.MovementShapes, formGroup, probe)
		AssociationSliceToForm("ArtefactTypeShapes", instanceWithInferedType, &instanceWithInferedType.ArtefactTypeShapes, formGroup, probe)
		AssociationSliceToForm("ArtistShapes", instanceWithInferedType, &instanceWithInferedType.ArtistShapes, formGroup, probe)
		AssociationSliceToForm("InfluenceShapes", instanceWithInferedType, &instanceWithInferedType.InfluenceShapes, formGroup, probe)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsNodeExpanded", instanceWithInferedType.IsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsMovementCategoryNodeExpanded", instanceWithInferedType.IsMovementCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsArtefactTypeCategoryNodeExpanded", instanceWithInferedType.IsArtefactTypeCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsArtistCategoryNodeExpanded", instanceWithInferedType.IsArtistCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsInfluenceCategoryNodeExpanded", instanceWithInferedType.IsInfluenceCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsMovementCategoryHidden", instanceWithInferedType.IsMovementCategoryHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsArtefactTypeCategoryHidden", instanceWithInferedType.IsArtefactTypeCategoryHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsArtistCategoryHidden", instanceWithInferedType.IsArtistCategoryHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsInfluenceCategoryHidden", instanceWithInferedType.IsInfluenceCategoryHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartDate", instanceWithInferedType.StartDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndDate", instanceWithInferedType.EndDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NbYearsForIntervals", instanceWithInferedType.NbYearsForIntervals, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XMargin", instanceWithInferedType.XMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("YMargin", instanceWithInferedType.YMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NextVerticalDateXMargin", instanceWithInferedType.NextVerticalDateXMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RedColorCode", instanceWithInferedType.RedColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BackgroundGreyColorCode", instanceWithInferedType.BackgroundGreyColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("GrayColorCode", instanceWithInferedType.GrayColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxYOffset", instanceWithInferedType.BottomBoxYOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxWidth", instanceWithInferedType.BottomBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxHeigth", instanceWithInferedType.BottomBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxFontSize", instanceWithInferedType.BottomBoxFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxFontWeigth", instanceWithInferedType.BottomBoxFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxFontFamily", instanceWithInferedType.BottomBoxFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxLetterSpacing", instanceWithInferedType.BottomBoxLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BottomBoxLetterColorCode", instanceWithInferedType.BottomBoxLetterColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("MovementRectAnchorType", instanceWithInferedType.MovementRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementTextAnchorType", instanceWithInferedType.MovementTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDominantBaselineType", instanceWithInferedType.MovementDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("MovementFontSize", instanceWithInferedType.MovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MajorMovementFontSize", instanceWithInferedType.MajorMovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinorMovementFontSize", instanceWithInferedType.MinorMovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementFontWeigth", instanceWithInferedType.MovementFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementFontFamily", instanceWithInferedType.MovementFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementLetterSpacing", instanceWithInferedType.MovementLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("AbstractMovementFontSize", instanceWithInferedType.AbstractMovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("AbstractMovementRectAnchorType", instanceWithInferedType.AbstractMovementRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("AbstractMovementTextAnchorType", instanceWithInferedType.AbstractMovementTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("AbstractDominantBaselineType", instanceWithInferedType.AbstractDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDateRectAnchorType", instanceWithInferedType.MovementDateRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDateTextAnchorType", instanceWithInferedType.MovementDateTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDateTextDominantBaselineType", instanceWithInferedType.MovementDateTextDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("MovementDateAndPlacesFontSize", instanceWithInferedType.MovementDateAndPlacesFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementDateAndPlacesFontWeigth", instanceWithInferedType.MovementDateAndPlacesFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementDateAndPlacesFontFamily", instanceWithInferedType.MovementDateAndPlacesFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementDateAndPlacesLetterSpacing", instanceWithInferedType.MovementDateAndPlacesLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementBelowArcY_Offset", instanceWithInferedType.MovementBelowArcY_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MovementBelowArcY_OffsetPerPlace", instanceWithInferedType.MovementBelowArcY_OffsetPerPlace, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("MovementPlacesRectAnchorType", instanceWithInferedType.MovementPlacesRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementPlacesTextAnchorType", instanceWithInferedType.MovementPlacesTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementPlacesDominantBaselineType", instanceWithInferedType.MovementPlacesDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtefactTypeFontSize", instanceWithInferedType.ArtefactTypeFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtefactTypeFontWeigth", instanceWithInferedType.ArtefactTypeFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtefactTypeFontFamily", instanceWithInferedType.ArtefactTypeFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtefactTypeLetterSpacing", instanceWithInferedType.ArtefactTypeLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("ArtefactTypeRectAnchorType", instanceWithInferedType.ArtefactTypeRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtefactDominantBaselineType", instanceWithInferedType.ArtefactDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtefactTypeStrokeWidth", instanceWithInferedType.ArtefactTypeStrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("ArtistRectAnchorType", instanceWithInferedType.ArtistRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistTextAnchorType", instanceWithInferedType.ArtistTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistDominantBaselineType", instanceWithInferedType.ArtistDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtistFontSize", instanceWithInferedType.ArtistFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MajorArtistFontSize", instanceWithInferedType.MajorArtistFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinorArtistFontSize", instanceWithInferedType.MinorArtistFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtistFontWeigth", instanceWithInferedType.ArtistFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtistFontFamily", instanceWithInferedType.ArtistFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtistLetterSpacing", instanceWithInferedType.ArtistLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("ArtistDateRectAnchorType", instanceWithInferedType.ArtistDateRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistDateTextAnchorType", instanceWithInferedType.ArtistDateTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistDateDominantBaselineType", instanceWithInferedType.ArtistDateDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtistDateAndPlacesFontSize", instanceWithInferedType.ArtistDateAndPlacesFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtistDateAndPlacesFontWeigth", instanceWithInferedType.ArtistDateAndPlacesFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtistDateAndPlacesFontFamily", instanceWithInferedType.ArtistDateAndPlacesFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ArtistDateAndPlacesLetterSpacing", instanceWithInferedType.ArtistDateAndPlacesLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("ArtistPlacesRectAnchorType", instanceWithInferedType.ArtistPlacesRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistPlacesTextAnchorType", instanceWithInferedType.ArtistPlacesTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistPlacesDominantBaselineType", instanceWithInferedType.ArtistPlacesDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("InfluenceArrowSize", instanceWithInferedType.InfluenceArrowSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("InfluenceArrowStartOffset", instanceWithInferedType.InfluenceArrowStartOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("InfluenceArrowEndOffset", instanceWithInferedType.InfluenceArrowEndOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("InfluenceCornerRadius", instanceWithInferedType.InfluenceCornerRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("InfluenceDashedLinePattern", instanceWithInferedType.InfluenceDashedLinePattern, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Influence:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("SourceMovement", instanceWithInferedType.SourceMovement, formGroup, probe)
		AssociationFieldToForm("SourceArtefactType", instanceWithInferedType.SourceArtefactType, formGroup, probe)
		AssociationFieldToForm("SourceArtist", instanceWithInferedType.SourceArtist, formGroup, probe)
		AssociationFieldToForm("TargetMovement", instanceWithInferedType.TargetMovement, formGroup, probe)
		AssociationFieldToForm("TargetArtefactType", instanceWithInferedType.TargetArtefactType, formGroup, probe)
		AssociationFieldToForm("TargetArtist", instanceWithInferedType.TargetArtist, formGroup, probe)
		BasicFieldtoForm("IsHypothtical", instanceWithInferedType.IsHypothtical, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.InfluenceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Influence", instanceWithInferedType.Influence, formGroup, probe)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.InfluenceShape](
				"Diagram",
				"InfluenceShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.InfluenceShape {
					return owner.InfluenceShapes
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("IsExpandedTmp", instanceWithInferedType.IsExpandedTmp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibraries",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibraries
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibrariesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibrariesWhoseNodeIsExpanded
				})
		}

	case *models.Movement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("HideDate", instanceWithInferedType.HideDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Places", instanceWithInferedType, &instanceWithInferedType.Places, formGroup, probe)
		BasicFieldtoForm("HasTaxonomicFilter", instanceWithInferedType.HasTaxonomicFilter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("TaxonomicFilter", instanceWithInferedType.TaxonomicFilter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsFeatured", instanceWithInferedType.IsFeatured, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("FeaturePrefix", instanceWithInferedType.FeaturePrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsMajor", instanceWithInferedType.IsMajor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsMinor", instanceWithInferedType.IsMinor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("AdditionnalName", instanceWithInferedType.AdditionnalName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.MovementShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Movement", instanceWithInferedType.Movement, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.MovementShape](
				"Diagram",
				"MovementShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.MovementShape {
					return owner.MovementShapes
				})
		}

	case *models.Place:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Movement, *models.Place](
				"Movement",
				"Places",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Movement) []*models.Place {
					return owner.Places
				})
		}

	default:
		_ = instanceWithInferedType
	}
}

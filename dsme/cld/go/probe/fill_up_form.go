// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
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
			false, false, 0, false, 0)

	case *models.ArtefactTypeShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ArtefactType", instanceWithInferedType.ArtefactType, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "ArtefactTypeShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"ArtefactTypeShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"ArtefactTypeShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Artist:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.ArtistShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Artist", instanceWithInferedType.Artist, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "ArtistShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"ArtistShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"ArtistShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ControlPointShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Relative", instanceWithInferedType.X_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Relative", instanceWithInferedType.Y_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsStartShapeTheClosestShape", instanceWithInferedType.IsStartShapeTheClosestShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "InfluenceShape"
			rf.Fieldname = "ControlPointShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.InfluenceShape),
					"ControlPointShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.InfluenceShape](
					nil,
					"ControlPointShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Desk:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SelectedDiagram", instanceWithInferedType.SelectedDiagram, formGroup, probe)

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("MovementShapes", instanceWithInferedType, &instanceWithInferedType.MovementShapes, formGroup, probe)
		AssociationSliceToForm("ArtefactTypeShapes", instanceWithInferedType, &instanceWithInferedType.ArtefactTypeShapes, formGroup, probe)
		AssociationSliceToForm("ArtistShapes", instanceWithInferedType, &instanceWithInferedType.ArtistShapes, formGroup, probe)
		AssociationSliceToForm("InfluenceShapes", instanceWithInferedType, &instanceWithInferedType.InfluenceShapes, formGroup, probe)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsNodeExpanded", instanceWithInferedType.IsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsMovementCategoryNodeExpanded", instanceWithInferedType.IsMovementCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsArtefactTypeCategoryNodeExpanded", instanceWithInferedType.IsArtefactTypeCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsArtistCategoryNodeExpanded", instanceWithInferedType.IsArtistCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInfluenceCategoryNodeExpanded", instanceWithInferedType.IsInfluenceCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsMovementCategoryShown", instanceWithInferedType.IsMovementCategoryShown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsArtefactTypeCategoryShown", instanceWithInferedType.IsArtefactTypeCategoryShown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsArtistCategoryShown", instanceWithInferedType.IsArtistCategoryShown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInfluenceCategoryShown", instanceWithInferedType.IsInfluenceCategoryShown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XMargin", instanceWithInferedType.XMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YMargin", instanceWithInferedType.YMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NextVerticalDateXMargin", instanceWithInferedType.NextVerticalDateXMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RedColorCode", instanceWithInferedType.RedColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BackgroundGreyColorCode", instanceWithInferedType.BackgroundGreyColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GrayColorCode", instanceWithInferedType.GrayColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxYOffset", instanceWithInferedType.BottomBoxYOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxWidth", instanceWithInferedType.BottomBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxHeigth", instanceWithInferedType.BottomBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxFontSize", instanceWithInferedType.BottomBoxFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxFontWeigth", instanceWithInferedType.BottomBoxFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxFontFamily", instanceWithInferedType.BottomBoxFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxLetterSpacing", instanceWithInferedType.BottomBoxLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BottomBoxLetterColorCode", instanceWithInferedType.BottomBoxLetterColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("MovementRectAnchorType", instanceWithInferedType.MovementRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementTextAnchorType", instanceWithInferedType.MovementTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDominantBaselineType", instanceWithInferedType.MovementDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("MovementFontSize", instanceWithInferedType.MovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MajorMovementFontSize", instanceWithInferedType.MajorMovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinorMovementFontSize", instanceWithInferedType.MinorMovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementFontWeigth", instanceWithInferedType.MovementFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementFontFamily", instanceWithInferedType.MovementFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementLetterSpacing", instanceWithInferedType.MovementLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AbstractMovementFontSize", instanceWithInferedType.AbstractMovementFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("AbstractMovementRectAnchorType", instanceWithInferedType.AbstractMovementRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("AbstractMovementTextAnchorType", instanceWithInferedType.AbstractMovementTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("AbstractDominantBaselineType", instanceWithInferedType.AbstractDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDateRectAnchorType", instanceWithInferedType.MovementDateRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDateTextAnchorType", instanceWithInferedType.MovementDateTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementDateTextDominantBaselineType", instanceWithInferedType.MovementDateTextDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("MovementDateAndPlacesFontSize", instanceWithInferedType.MovementDateAndPlacesFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementDateAndPlacesFontWeigth", instanceWithInferedType.MovementDateAndPlacesFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementDateAndPlacesFontFamily", instanceWithInferedType.MovementDateAndPlacesFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementDateAndPlacesLetterSpacing", instanceWithInferedType.MovementDateAndPlacesLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementBelowArcY_Offset", instanceWithInferedType.MovementBelowArcY_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MovementBelowArcY_OffsetPerPlace", instanceWithInferedType.MovementBelowArcY_OffsetPerPlace, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("MovementPlacesRectAnchorType", instanceWithInferedType.MovementPlacesRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementPlacesTextAnchorType", instanceWithInferedType.MovementPlacesTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("MovementPlacesDominantBaselineType", instanceWithInferedType.MovementPlacesDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtefactTypeFontSize", instanceWithInferedType.ArtefactTypeFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtefactTypeFontWeigth", instanceWithInferedType.ArtefactTypeFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtefactTypeFontFamily", instanceWithInferedType.ArtefactTypeFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtefactTypeLetterSpacing", instanceWithInferedType.ArtefactTypeLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ArtefactTypeRectAnchorType", instanceWithInferedType.ArtefactTypeRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtefactDominantBaselineType", instanceWithInferedType.ArtefactDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtefactTypeStrokeWidth", instanceWithInferedType.ArtefactTypeStrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ArtistRectAnchorType", instanceWithInferedType.ArtistRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistTextAnchorType", instanceWithInferedType.ArtistTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistDominantBaselineType", instanceWithInferedType.ArtistDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtistFontSize", instanceWithInferedType.ArtistFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MajorArtistFontSize", instanceWithInferedType.MajorArtistFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinorArtistFontSize", instanceWithInferedType.MinorArtistFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtistFontWeigth", instanceWithInferedType.ArtistFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtistFontFamily", instanceWithInferedType.ArtistFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtistLetterSpacing", instanceWithInferedType.ArtistLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ArtistDateRectAnchorType", instanceWithInferedType.ArtistDateRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistDateTextAnchorType", instanceWithInferedType.ArtistDateTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistDateDominantBaselineType", instanceWithInferedType.ArtistDateDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ArtistDateAndPlacesFontSize", instanceWithInferedType.ArtistDateAndPlacesFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtistDateAndPlacesFontWeigth", instanceWithInferedType.ArtistDateAndPlacesFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtistDateAndPlacesFontFamily", instanceWithInferedType.ArtistDateAndPlacesFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArtistDateAndPlacesLetterSpacing", instanceWithInferedType.ArtistDateAndPlacesLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ArtistPlacesRectAnchorType", instanceWithInferedType.ArtistPlacesRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistPlacesTextAnchorType", instanceWithInferedType.ArtistPlacesTextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("ArtistPlacesDominantBaselineType", instanceWithInferedType.ArtistPlacesDominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("InfluenceArrowSize", instanceWithInferedType.InfluenceArrowSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceArrowStartOffset", instanceWithInferedType.InfluenceArrowStartOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceArrowEndOffset", instanceWithInferedType.InfluenceArrowEndOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceCornerRadius", instanceWithInferedType.InfluenceCornerRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceDashedLinePattern", instanceWithInferedType.InfluenceDashedLinePattern, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Influence:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SourceMovement", instanceWithInferedType.SourceMovement, formGroup, probe)
		AssociationFieldToForm("SourceArtefactType", instanceWithInferedType.SourceArtefactType, formGroup, probe)
		AssociationFieldToForm("SourceArtist", instanceWithInferedType.SourceArtist, formGroup, probe)
		AssociationFieldToForm("TargetMovement", instanceWithInferedType.TargetMovement, formGroup, probe)
		AssociationFieldToForm("TargetArtefactType", instanceWithInferedType.TargetArtefactType, formGroup, probe)
		AssociationFieldToForm("TargetArtist", instanceWithInferedType.TargetArtist, formGroup, probe)
		BasicFieldtoForm("IsHypothtical", instanceWithInferedType.IsHypothtical, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.InfluenceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Influence", instanceWithInferedType.Influence, formGroup, probe)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "InfluenceShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"InfluenceShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"InfluenceShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Movement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.MovementShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Movement", instanceWithInferedType.Movement, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "MovementShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"MovementShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"MovementShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}

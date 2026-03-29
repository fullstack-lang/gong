package models

import (
	"time"
)

func (stager *Stager) enforce_semantic_diagrams() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		var zeroDate time.Time
		if diagram.StartDate.Equal(zeroDate) {
			diagram.StartDate = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
			needCommit = true
		}
		if diagram.EndDate.Equal(zeroDate) {
			diagram.EndDate = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
			needCommit = true
		}

		if diagram.NbYearsForIntervals == 0 {
			diagram.NbYearsForIntervals = 5
			needCommit = true
		}

		needCommit = enforceDiagramFields(diagram) || needCommit

		if diagram.AlignDates() {
			needCommit = true
		}

		if diagram.EndDate.Before(diagram.StartDate) {
			diagram.EndDate = diagram.StartDate
			needCommit = true
		}

		if diagram.MovementRectAnchorType == "" {
			diagram.MovementRectAnchorType = RECT_RIGHT
			needCommit = true
		}
		if diagram.MovementTextAnchorType == "" {
			diagram.MovementTextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}
		if diagram.MovementDateRectAnchorType == "" {
			diagram.MovementDateRectAnchorType = RECT_BOTTOM_LEFT
			needCommit = true
		}
		if diagram.MovementDateTextAnchorType == "" {
			diagram.MovementDateTextAnchorType = TEXT_ANCHOR_START
			needCommit = true
		}
		if diagram.MovementPlacesRectAnchorType == "" {
			diagram.MovementPlacesRectAnchorType = RECT_BOTTOM_RIGHT
			needCommit = true
		}
		if diagram.MovementPlacesTextAnchorType == "" {
			diagram.MovementPlacesTextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}
	}

	return
}

func enforceDiagramFields(diagram *Diagram) (needCommit bool) {
	if diagram.XMargin == 0 {
		diagram.XMargin = 20.000000
		needCommit = true
	}
	if diagram.YMargin == 0 {
		diagram.YMargin = 40.000000
		needCommit = true
	}
	if diagram.Height == 0 {
		diagram.Height = 900.000000
		needCommit = true
	}
	if diagram.NextVerticalDateXMargin == 0 {
		diagram.NextVerticalDateXMargin = 700.000000
		needCommit = true
	}
	if diagram.RedColorCode == "" {
		diagram.RedColorCode = "#D23B22"
		needCommit = true
	}
	if diagram.BackgroundGreyColorCode == "" {
		diagram.BackgroundGreyColorCode = "#DED6CA"
		needCommit = true
	}
	if diagram.GrayColorCode == "" {
		diagram.GrayColorCode = "#343434"
		needCommit = true
	}
	if diagram.BottomBoxYOffset == 0 {
		diagram.BottomBoxYOffset = 50.000000
		needCommit = true
	}
	if diagram.BottomBoxWidth == 0 {
		diagram.BottomBoxWidth = 770.000000
		needCommit = true
	}
	if diagram.BottomBoxHeigth == 0 {
		diagram.BottomBoxHeigth = 150.000000
		needCommit = true
	}
	if diagram.BottomBoxFontSize == "" {
		diagram.BottomBoxFontSize = "74px"
		needCommit = true
	}
	if diagram.BottomBoxFontWeigth == "" {
		diagram.BottomBoxFontWeigth = "500"
		needCommit = true
	}
	if diagram.BottomBoxFontFamily == "" {
		diagram.BottomBoxFontFamily = "ChunkFive, sans-serif"
		needCommit = true
	}
	if diagram.BottomBoxLetterSpacing == "" {
		diagram.BottomBoxLetterSpacing = "1"
		needCommit = true
	}
	if diagram.BottomBoxLetterColorCode == "" {
		diagram.BottomBoxLetterColorCode = "#debdaaff"
		needCommit = true
	}
	if diagram.MovementRectAnchorType == "" {
		diagram.MovementRectAnchorType = RECT_RIGHT
		needCommit = true
	}
	if diagram.MovementTextAnchorType == "" {
		diagram.MovementTextAnchorType = TEXT_ANCHOR_END
		needCommit = true
	}
	if diagram.MovementDominantBaselineType == "" {
		diagram.MovementDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.MovementFontSize == "" {
		diagram.MovementFontSize = "14px"
		needCommit = true
	}
	if diagram.MajorMovementFontSize == "" {
		diagram.MajorMovementFontSize = "18px"
		needCommit = true
	}
	if diagram.MinorMovementFontSize == "" {
		diagram.MinorMovementFontSize = "10px"
		needCommit = true
	}
	if diagram.MovementFontWeigth == "" {
		diagram.MovementFontWeigth = "Thin"
		needCommit = true
	}
	if diagram.MovementFontFamily == "" {
		diagram.MovementFontFamily = "Futura, sans serif"
		needCommit = true
	}
	if diagram.MovementLetterSpacing == "" {
		diagram.MovementLetterSpacing = "1"
		needCommit = true
	}
	if diagram.AbstractMovementFontSize == "" {
		diagram.AbstractMovementFontSize = "12px"
		needCommit = true
	}
	if diagram.AbstractMovementRectAnchorType == "" {
		diagram.AbstractMovementRectAnchorType = RECT_TOP_RIGHT
		needCommit = true
	}
	if diagram.AbstractMovementTextAnchorType == "" {
		diagram.AbstractMovementTextAnchorType = TEXT_ANCHOR_END
		needCommit = true
	}
	if diagram.AbstractDominantBaselineType == "" {
		diagram.AbstractDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.MovementDateRectAnchorType == "" {
		diagram.MovementDateRectAnchorType = RECT_BOTTOM_LEFT
		needCommit = true
	}
	if diagram.MovementDateTextAnchorType == "" {
		diagram.MovementDateTextAnchorType = TEXT_ANCHOR_START
		needCommit = true
	}
	if diagram.MovementDateTextDominantBaselineType == "" {
		diagram.MovementDateTextDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.MovementDateAndPlacesFontSize == "" {
		diagram.MovementDateAndPlacesFontSize = "9px"
		needCommit = true
	}
	if diagram.MovementDateAndPlacesFontFamily == "" {
		diagram.MovementDateAndPlacesFontFamily = "Futura"
		needCommit = true
	}
	if diagram.MovementDateAndPlacesLetterSpacing == "" {
		diagram.MovementDateAndPlacesLetterSpacing = "0"
		needCommit = true
	}
	if diagram.MovementBelowArcY_Offset == 0 {
		diagram.MovementBelowArcY_Offset = 6.000000
		needCommit = true
	}
	if diagram.MovementBelowArcY_OffsetPerPlace == 0 {
		diagram.MovementBelowArcY_OffsetPerPlace = 6.000000
		needCommit = true
	}
	if diagram.MovementPlacesRectAnchorType == "" {
		diagram.MovementPlacesRectAnchorType = RECT_BOTTOM_RIGHT
		needCommit = true
	}
	if diagram.MovementPlacesTextAnchorType == "" {
		diagram.MovementPlacesTextAnchorType = TEXT_ANCHOR_END
		needCommit = true
	}
	if diagram.MovementPlacesDominantBaselineType == "" {
		diagram.MovementPlacesDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.ArtefactTypeFontSize == "" {
		diagram.ArtefactTypeFontSize = "12px"
		needCommit = true
	}
	if diagram.ArtefactTypeFontFamily == "" {
		diagram.ArtefactTypeFontFamily = "Futura"
		needCommit = true
	}
	if diagram.ArtefactTypeRectAnchorType == "" {
		diagram.ArtefactTypeRectAnchorType = RECT_CENTER
		needCommit = true
	}
	if diagram.ArtefactDominantBaselineType == "" {
		diagram.ArtefactDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.ArtefactTypeStrokeWidth == 0 {
		diagram.ArtefactTypeStrokeWidth = 3.000000
		needCommit = true
	}
	if diagram.ArtistRectAnchorType == "" {
		diagram.ArtistRectAnchorType = RECT_CENTER
		needCommit = true
	}
	if diagram.ArtistTextAnchorType == "" {
		diagram.ArtistTextAnchorType = TEXT_ANCHOR_CENTER
		needCommit = true
	}
	if diagram.ArtistDominantBaselineType == "" {
		diagram.ArtistDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.ArtistFontSize == "" {
		diagram.ArtistFontSize = "14px"
		needCommit = true
	}
	if diagram.ArtistFontWeigth == "" {
		diagram.ArtistFontWeigth = "100"
		needCommit = true
	}
	if diagram.ArtistFontFamily == "" {
		diagram.ArtistFontFamily = "Futura"
		needCommit = true
	}
	if diagram.ArtistLetterSpacing == "" {
		diagram.ArtistLetterSpacing = "0"
		needCommit = true
	}
	if diagram.ArtistDateRectAnchorType == "" {
		diagram.ArtistDateRectAnchorType = RECT_BOTTOM_LEFT
		needCommit = true
	}
	if diagram.ArtistDateTextAnchorType == "" {
		diagram.ArtistDateTextAnchorType = TEXT_ANCHOR_START
		needCommit = true
	}
	if diagram.ArtistDateDominantBaselineType == "" {
		diagram.ArtistDateDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.ArtistDateAndPlacesFontSize == "" {
		diagram.ArtistDateAndPlacesFontSize = "10px"
		needCommit = true
	}
	if diagram.ArtistDateAndPlacesFontWeigth == "" {
		diagram.ArtistDateAndPlacesFontWeigth = "100"
		needCommit = true
	}
	if diagram.ArtistDateAndPlacesFontFamily == "" {
		diagram.ArtistDateAndPlacesFontFamily = "Futura"
		needCommit = true
	}
	if diagram.ArtistPlacesRectAnchorType == "" {
		diagram.ArtistPlacesRectAnchorType = RECT_BOTTOM_RIGHT
		needCommit = true
	}
	if diagram.ArtistPlacesTextAnchorType == "" {
		diagram.ArtistPlacesTextAnchorType = TEXT_ANCHOR_END
		needCommit = true
	}
	if diagram.ArtistPlacesDominantBaselineType == "" {
		diagram.ArtistPlacesDominantBaselineType = DominantBaselineCentral
		needCommit = true
	}
	if diagram.InfluenceArrowSize == 0 {
		diagram.InfluenceArrowSize = 6.000000
		needCommit = true
	}
	if diagram.InfluenceArrowStartOffset == 0 {
		diagram.InfluenceArrowStartOffset = 19.000000
		needCommit = true
	}
	if diagram.InfluenceArrowEndOffset == 0 {
		diagram.InfluenceArrowEndOffset = 9.000000
		needCommit = true
	}
	if diagram.InfluenceCornerRadius == 0 {
		diagram.InfluenceCornerRadius = 20.000000
		needCommit = true
	}
	if diagram.InfluenceDashedLinePattern == "" {
		diagram.InfluenceDashedLinePattern = "7 3"
		needCommit = true
	}
	return
}

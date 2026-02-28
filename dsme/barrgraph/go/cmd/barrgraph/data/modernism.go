package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/barrgraph/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Desk__00000000_ := (&models.Desk{Name: `Desk`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default`}).Stage(stage)

	// insertion point for initialization of values

	__Desk__00000000_.Name = `Desk`

	__Diagram__00000000_.Name = `Default`
	__Diagram__00000000_.IsEditable = false
	__Diagram__00000000_.IsNodeExpanded = false
	__Diagram__00000000_.IsMovementCategoryNodeExpanded = false
	__Diagram__00000000_.IsArtefactTypeCategoryNodeExpanded = false
	__Diagram__00000000_.IsArtistCategoryNodeExpanded = false
	__Diagram__00000000_.IsInfluenceCategoryNodeExpanded = false
	__Diagram__00000000_.IsMovementCategoryShown = false
	__Diagram__00000000_.IsArtefactTypeCategoryShown = false
	__Diagram__00000000_.IsArtistCategoryShown = false
	__Diagram__00000000_.IsInfluenceCategoryShown = false
	__Diagram__00000000_.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1900-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.XMargin = 0.000000
	__Diagram__00000000_.YMargin = 0.000000
	__Diagram__00000000_.Height = 0.000000
	__Diagram__00000000_.NextVerticalDateXMargin = 0.000000
	__Diagram__00000000_.RedColorCode = ``
	__Diagram__00000000_.BackgroundGreyColorCode = ``
	__Diagram__00000000_.GrayColorCode = ``
	__Diagram__00000000_.BottomBoxYOffset = 0.000000
	__Diagram__00000000_.BottomBoxWidth = 0.000000
	__Diagram__00000000_.BottomBoxHeigth = 0.000000
	__Diagram__00000000_.BottomBoxFontSize = ``
	__Diagram__00000000_.BottomBoxFontWeigth = ``
	__Diagram__00000000_.BottomBoxFontFamily = ``
	__Diagram__00000000_.BottomBoxLetterSpacing = ``
	__Diagram__00000000_.BottomBoxLetterColorCode = ``
	__Diagram__00000000_.MovementRectAnchorType = models.RECT_RIGHT
	__Diagram__00000000_.MovementTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.MovementDominantBaselineType = ""
	__Diagram__00000000_.MovementFontSize = ``
	__Diagram__00000000_.MajorMovementFontSize = ``
	__Diagram__00000000_.MinorMovementFontSize = ``
	__Diagram__00000000_.MovementFontWeigth = ``
	__Diagram__00000000_.MovementFontFamily = ``
	__Diagram__00000000_.MovementLetterSpacing = ``
	__Diagram__00000000_.AbstractMovementFontSize = ``
	__Diagram__00000000_.AbstractMovementRectAnchorType = ""
	__Diagram__00000000_.AbstractMovementTextAnchorType = ""
	__Diagram__00000000_.AbstractDominantBaselineType = ""
	__Diagram__00000000_.MovementDateRectAnchorType = models.RECT_BOTTOM_LEFT
	__Diagram__00000000_.MovementDateTextAnchorType = models.TEXT_ANCHOR_START
	__Diagram__00000000_.MovementDateTextDominantBaselineType = ""
	__Diagram__00000000_.MovementDateAndPlacesFontSize = ``
	__Diagram__00000000_.MovementDateAndPlacesFontWeigth = ``
	__Diagram__00000000_.MovementDateAndPlacesFontFamily = ``
	__Diagram__00000000_.MovementDateAndPlacesLetterSpacing = ``
	__Diagram__00000000_.MovementBelowArcY_Offset = 0.000000
	__Diagram__00000000_.MovementBelowArcY_OffsetPerPlace = 0.000000
	__Diagram__00000000_.MovementPlacesRectAnchorType = models.RECT_BOTTOM_RIGHT
	__Diagram__00000000_.MovementPlacesTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.MovementPlacesDominantBaselineType = ""
	__Diagram__00000000_.ArtefactTypeFontSize = ``
	__Diagram__00000000_.ArtefactTypeFontWeigth = ``
	__Diagram__00000000_.ArtefactTypeFontFamily = ``
	__Diagram__00000000_.ArtefactTypeLetterSpacing = ``
	__Diagram__00000000_.ArtefactTypeRectAnchorType = ""
	__Diagram__00000000_.ArtefactDominantBaselineType = ""
	__Diagram__00000000_.ArtefactTypeStrokeWidth = 0.000000
	__Diagram__00000000_.ArtistRectAnchorType = ""
	__Diagram__00000000_.ArtistTextAnchorType = ""
	__Diagram__00000000_.ArtistDominantBaselineType = ""
	__Diagram__00000000_.ArtistFontSize = ``
	__Diagram__00000000_.MajorArtistFontSize = ``
	__Diagram__00000000_.MinorArtistFontSize = ``
	__Diagram__00000000_.ArtistFontWeigth = ``
	__Diagram__00000000_.ArtistFontFamily = ``
	__Diagram__00000000_.ArtistLetterSpacing = ``
	__Diagram__00000000_.ArtistDateRectAnchorType = ""
	__Diagram__00000000_.ArtistDateTextAnchorType = ""
	__Diagram__00000000_.ArtistDateDominantBaselineType = ""
	__Diagram__00000000_.ArtistDateAndPlacesFontSize = ``
	__Diagram__00000000_.ArtistDateAndPlacesFontWeigth = ``
	__Diagram__00000000_.ArtistDateAndPlacesFontFamily = ``
	__Diagram__00000000_.ArtistDateAndPlacesLetterSpacing = ``
	__Diagram__00000000_.ArtistPlacesRectAnchorType = ""
	__Diagram__00000000_.ArtistPlacesTextAnchorType = ""
	__Diagram__00000000_.ArtistPlacesDominantBaselineType = ""
	__Diagram__00000000_.InfluenceArrowSize = 0.000000
	__Diagram__00000000_.InfluenceArrowStartOffset = 0.000000
	__Diagram__00000000_.InfluenceArrowEndOffset = 0.000000
	__Diagram__00000000_.InfluenceCornerRadius = 0.000000
	__Diagram__00000000_.InfluenceDashedLinePattern = ``

	// insertion point for setup of pointers
	__Desk__00000000_.SelectedDiagram = __Diagram__00000000_
}

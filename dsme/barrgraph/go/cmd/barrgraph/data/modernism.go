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
	__Diagram__00000000_.IsNodeExpanded = true
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
	__Diagram__00000000_.XMargin = 20.000000
	__Diagram__00000000_.YMargin = 40.000000
	__Diagram__00000000_.Height = 900.000000
	__Diagram__00000000_.NextVerticalDateXMargin = 700.000000
	__Diagram__00000000_.RedColorCode = `#D23B22`
	__Diagram__00000000_.BackgroundGreyColorCode = `#DED6CA`
	__Diagram__00000000_.GrayColorCode = `#343434`
	__Diagram__00000000_.BottomBoxYOffset = 50.000000
	__Diagram__00000000_.BottomBoxWidth = 770.000000
	__Diagram__00000000_.BottomBoxHeigth = 150.000000
	__Diagram__00000000_.BottomBoxFontSize = `74px`
	__Diagram__00000000_.BottomBoxFontWeigth = `500`
	__Diagram__00000000_.BottomBoxFontFamily = `ChunkFive, sans-serif`
	__Diagram__00000000_.BottomBoxLetterSpacing = `1`
	__Diagram__00000000_.BottomBoxLetterColorCode = `#debdaaff`
	__Diagram__00000000_.MovementRectAnchorType = models.RECT_RIGHT
	__Diagram__00000000_.MovementTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.MovementDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.MovementFontSize = `14px`
	__Diagram__00000000_.MajorMovementFontSize = `18px`
	__Diagram__00000000_.MinorMovementFontSize = `10px`
	__Diagram__00000000_.MovementFontWeigth = `Thin`
	__Diagram__00000000_.MovementFontFamily = `Futura, sans serif`
	__Diagram__00000000_.MovementLetterSpacing = `1`
	__Diagram__00000000_.AbstractMovementFontSize = `12px`
	__Diagram__00000000_.AbstractMovementRectAnchorType = models.RECT_TOP_RIGHT
	__Diagram__00000000_.AbstractMovementTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.AbstractDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.MovementDateRectAnchorType = models.RECT_BOTTOM_LEFT
	__Diagram__00000000_.MovementDateTextAnchorType = models.TEXT_ANCHOR_START
	__Diagram__00000000_.MovementDateTextDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.MovementDateAndPlacesFontSize = `9px`
	__Diagram__00000000_.MovementDateAndPlacesFontWeigth = ``
	__Diagram__00000000_.MovementDateAndPlacesFontFamily = `Futura`
	__Diagram__00000000_.MovementDateAndPlacesLetterSpacing = `0`
	__Diagram__00000000_.MovementBelowArcY_Offset = 6.000000
	__Diagram__00000000_.MovementBelowArcY_OffsetPerPlace = 6.000000
	__Diagram__00000000_.MovementPlacesRectAnchorType = models.RECT_BOTTOM_RIGHT
	__Diagram__00000000_.MovementPlacesTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.MovementPlacesDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtefactTypeFontSize = `12px`
	__Diagram__00000000_.ArtefactTypeFontWeigth = ``
	__Diagram__00000000_.ArtefactTypeFontFamily = `Futura`
	__Diagram__00000000_.ArtefactTypeLetterSpacing = ``
	__Diagram__00000000_.ArtefactTypeRectAnchorType = models.RECT_CENTER
	__Diagram__00000000_.ArtefactDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtefactTypeStrokeWidth = 3.000000
	__Diagram__00000000_.ArtistRectAnchorType = models.RECT_CENTER
	__Diagram__00000000_.ArtistTextAnchorType = models.TEXT_ANCHOR_CENTER
	__Diagram__00000000_.ArtistDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtistFontSize = `14px`
	__Diagram__00000000_.MajorArtistFontSize = ``
	__Diagram__00000000_.MinorArtistFontSize = ``
	__Diagram__00000000_.ArtistFontWeigth = `100`
	__Diagram__00000000_.ArtistFontFamily = `Futura`
	__Diagram__00000000_.ArtistLetterSpacing = `0`
	__Diagram__00000000_.ArtistDateRectAnchorType = models.RECT_BOTTOM_LEFT
	__Diagram__00000000_.ArtistDateTextAnchorType = models.TEXT_ANCHOR_START
	__Diagram__00000000_.ArtistDateDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtistDateAndPlacesFontSize = `10px`
	__Diagram__00000000_.ArtistDateAndPlacesFontWeigth = `100`
	__Diagram__00000000_.ArtistDateAndPlacesFontFamily = `Futura`
	__Diagram__00000000_.ArtistDateAndPlacesLetterSpacing = ``
	__Diagram__00000000_.ArtistPlacesRectAnchorType = models.RECT_BOTTOM_RIGHT
	__Diagram__00000000_.ArtistPlacesTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.ArtistPlacesDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.InfluenceArrowSize = 6.000000
	__Diagram__00000000_.InfluenceArrowStartOffset = 19.000000
	__Diagram__00000000_.InfluenceArrowEndOffset = 9.000000
	__Diagram__00000000_.InfluenceCornerRadius = 20.000000
	__Diagram__00000000_.InfluenceDashedLinePattern = `7 3`

	// insertion point for setup of pointers
	__Desk__00000000_.SelectedDiagram = __Diagram__00000000_
}

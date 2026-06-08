package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__ArtefactType__00000000_ := (&models.ArtefactType{Name: `History Painting`}).Stage(stage)
	__ArtefactType__00000001_ := (&models.ArtefactType{Name: `Portraiture`}).Stage(stage)
	__ArtefactType__00000002_ := (&models.ArtefactType{Name: `Genre Painting`}).Stage(stage)
	__ArtefactType__00000003_ := (&models.ArtefactType{Name: `Landscape Painting`}).Stage(stage)
	__ArtefactType__00000004_ := (&models.ArtefactType{Name: `Protestantism`}).Stage(stage)
	__ArtefactType__00000005_ := (&models.ArtefactType{Name: `Catholicism`}).Stage(stage)
	__ArtefactType__00000006_ := (&models.ArtefactType{Name: `Eighty Years' War`}).Stage(stage)

	__ArtefactTypeShape__00000000_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000001_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000002_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000003_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000004_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000005_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000006_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)

	__Artist__00000000_ := (&models.Artist{Name: `Frans Hals`}).Stage(stage)
	__Artist__00000001_ := (&models.Artist{Name: `Rembrandt van Rijn`}).Stage(stage)
	__Artist__00000002_ := (&models.Artist{Name: `Johannes Vermeer`}).Stage(stage)
	__Artist__00000003_ := (&models.Artist{Name: `Jan Steen`}).Stage(stage)
	__Artist__00000004_ := (&models.Artist{Name: `Pieter de Hooch`}).Stage(stage)
	__Artist__00000005_ := (&models.Artist{Name: `Jacob van Ruisdael`}).Stage(stage)

	__ArtistShape__00000000_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000001_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000002_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000003_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000004_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000005_ := (&models.ArtistShape{Name: ``}).Stage(stage)

	__Desk__00000000_ := (&models.Desk{Name: `Desk`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Netherland Golden Century`}).Stage(stage)

	__Influence__00000000_ := (&models.Influence{Name: `Haarlem Mannerism to Pre-Rembrandtists`}).Stage(stage)
	__Influence__00000001_ := (&models.Influence{Name: `Pre-Rembrandtists to Rembrandt van Rijn`}).Stage(stage)
	__Influence__00000002_ := (&models.Influence{Name: `Dutch Caravaggisti to Rembrandt van Rijn`}).Stage(stage)
	__Influence__00000003_ := (&models.Influence{Name: `Haarlem Mannerism to Frans Hals`}).Stage(stage)
	__Influence__00000004_ := (&models.Influence{Name: `Frans Hals to Jan Steen`}).Stage(stage)
	__Influence__00000005_ := (&models.Influence{Name: `Rembrandt van Rijn to Delft School`}).Stage(stage)
	__Influence__00000006_ := (&models.Influence{Name: `Delft School to Johannes Vermeer`}).Stage(stage)
	__Influence__00000007_ := (&models.Influence{Name: `Delft School to Pieter de Hooch`}).Stage(stage)
	__Influence__00000008_ := (&models.Influence{Name: `Catholicism to Dutch Caravaggisti`}).Stage(stage)
	__Influence__00000009_ := (&models.Influence{Name: `Eighty Years' War to Protestantism`}).Stage(stage)
	__Influence__00000010_ := (&models.Influence{Name: `Eighty Years' War to Catholicism`}).Stage(stage)
	__Influence__00000011_ := (&models.Influence{Name: `Protestantism to Portraiture`}).Stage(stage)
	__Influence__00000012_ := (&models.Influence{Name: `Protestantism to Genre Painting`}).Stage(stage)
	__Influence__00000013_ := (&models.Influence{Name: `Protestantism to Landscape Painting`}).Stage(stage)
	__Influence__00000014_ := (&models.Influence{Name: `Portraiture to Frans Hals`}).Stage(stage)
	__Influence__00000015_ := (&models.Influence{Name: `History Painting to Pre-Rembrandtists`}).Stage(stage)
	__Influence__00000016_ := (&models.Influence{Name: `Landscape Painting to Jacob van Ruisdael`}).Stage(stage)
	__Influence__00000017_ := (&models.Influence{Name: `Genre Painting to Jan Steen`}).Stage(stage)

	__InfluenceShape__00000000_ := (&models.InfluenceShape{Name: `Haarlem Mannerism to Pre-Rembrandtists`}).Stage(stage)
	__InfluenceShape__00000001_ := (&models.InfluenceShape{Name: `Pre-Rembrandtists to Rembrandt van Rijn`}).Stage(stage)
	__InfluenceShape__00000002_ := (&models.InfluenceShape{Name: `Dutch Caravaggisti to Rembrandt van Rijn`}).Stage(stage)
	__InfluenceShape__00000003_ := (&models.InfluenceShape{Name: `Haarlem Mannerism to Frans Hals`}).Stage(stage)
	__InfluenceShape__00000004_ := (&models.InfluenceShape{Name: `Frans Hals to Jan Steen`}).Stage(stage)
	__InfluenceShape__00000005_ := (&models.InfluenceShape{Name: `Rembrandt van Rijn to Delft School`}).Stage(stage)
	__InfluenceShape__00000006_ := (&models.InfluenceShape{Name: `Delft School to Johannes Vermeer`}).Stage(stage)
	__InfluenceShape__00000007_ := (&models.InfluenceShape{Name: `Delft School to Pieter de Hooch`}).Stage(stage)
	__InfluenceShape__00000008_ := (&models.InfluenceShape{Name: `Catholicism to Dutch Caravaggisti`}).Stage(stage)
	__InfluenceShape__00000009_ := (&models.InfluenceShape{Name: `Eighty Years' War to Protestantism`}).Stage(stage)
	__InfluenceShape__00000010_ := (&models.InfluenceShape{Name: `Eighty Years' War to Catholicism`}).Stage(stage)
	__InfluenceShape__00000011_ := (&models.InfluenceShape{Name: `Protestantism to Portraiture`}).Stage(stage)
	__InfluenceShape__00000012_ := (&models.InfluenceShape{Name: `Protestantism to Genre Painting`}).Stage(stage)
	__InfluenceShape__00000013_ := (&models.InfluenceShape{Name: `Protestantism to Landscape Painting`}).Stage(stage)
	__InfluenceShape__00000014_ := (&models.InfluenceShape{Name: `Portraiture to Frans Hals`}).Stage(stage)
	__InfluenceShape__00000015_ := (&models.InfluenceShape{Name: `History Painting to Pre-Rembrandtists`}).Stage(stage)
	__InfluenceShape__00000016_ := (&models.InfluenceShape{Name: `Landscape Painting to Jacob van Ruisdael`}).Stage(stage)
	__InfluenceShape__00000017_ := (&models.InfluenceShape{Name: `Genre Painting to Jan Steen`}).Stage(stage)

	__Movement__00000000_ := (&models.Movement{Name: `Haarlem Mannerism`}).Stage(stage)
	__Movement__00000001_ := (&models.Movement{Name: `Pre-Rembrandtists`}).Stage(stage)
	__Movement__00000002_ := (&models.Movement{Name: `Dutch Caravaggisti`}).Stage(stage)
	__Movement__00000003_ := (&models.Movement{Name: `Delft School`}).Stage(stage)

	__MovementShape__00000000_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000001_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000002_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000003_ := (&models.MovementShape{Name: ``}).Stage(stage)

	__Place__00000000_ := (&models.Place{Name: `Haarlem`}).Stage(stage)
	__Place__00000001_ := (&models.Place{Name: `Amsterdam`}).Stage(stage)
	__Place__00000002_ := (&models.Place{Name: `Delft`}).Stage(stage)
	__Place__00000003_ := (&models.Place{Name: `Leiden`}).Stage(stage)
	__Place__00000004_ := (&models.Place{Name: `Utrecht`}).Stage(stage)

	// insertion point for initialization of values

	__ArtefactType__00000000_.Name = `History Painting`
	__ArtefactType__00000000_.ComputedPrefix = ``

	__ArtefactType__00000001_.Name = `Portraiture`
	__ArtefactType__00000001_.ComputedPrefix = ``

	__ArtefactType__00000002_.Name = `Genre Painting`
	__ArtefactType__00000002_.ComputedPrefix = ``

	__ArtefactType__00000003_.Name = `Landscape Painting`
	__ArtefactType__00000003_.ComputedPrefix = ``

	__ArtefactType__00000004_.Name = `Protestantism`
	__ArtefactType__00000004_.ComputedPrefix = ``

	__ArtefactType__00000005_.Name = `Catholicism`
	__ArtefactType__00000005_.ComputedPrefix = ``

	__ArtefactType__00000006_.Name = `Eighty Years' War`
	__ArtefactType__00000006_.ComputedPrefix = ``

	__ArtefactTypeShape__00000000_.Name = ``
	__ArtefactTypeShape__00000000_.X = 80.000000
	__ArtefactTypeShape__00000000_.Y = 240.000000
	__ArtefactTypeShape__00000000_.Width = 120.000000
	__ArtefactTypeShape__00000000_.Height = 25.000000
	__ArtefactTypeShape__00000000_.IsHidden = false

	__ArtefactTypeShape__00000001_.Name = ``
	__ArtefactTypeShape__00000001_.X = 220.000000
	__ArtefactTypeShape__00000001_.Y = 240.000000
	__ArtefactTypeShape__00000001_.Width = 120.000000
	__ArtefactTypeShape__00000001_.Height = 25.000000
	__ArtefactTypeShape__00000001_.IsHidden = false

	__ArtefactTypeShape__00000002_.Name = ``
	__ArtefactTypeShape__00000002_.X = 380.000000
	__ArtefactTypeShape__00000002_.Y = 240.000000
	__ArtefactTypeShape__00000002_.Width = 120.000000
	__ArtefactTypeShape__00000002_.Height = 25.000000
	__ArtefactTypeShape__00000002_.IsHidden = false

	__ArtefactTypeShape__00000003_.Name = ``
	__ArtefactTypeShape__00000003_.X = 550.000000
	__ArtefactTypeShape__00000003_.Y = 240.000000
	__ArtefactTypeShape__00000003_.Width = 120.000000
	__ArtefactTypeShape__00000003_.Height = 25.000000
	__ArtefactTypeShape__00000003_.IsHidden = false

	__ArtefactTypeShape__00000004_.Name = ``
	__ArtefactTypeShape__00000004_.X = 300.000000
	__ArtefactTypeShape__00000004_.Y = 115.000000
	__ArtefactTypeShape__00000004_.Width = 120.000000
	__ArtefactTypeShape__00000004_.Height = 25.000000
	__ArtefactTypeShape__00000004_.IsHidden = false

	__ArtefactTypeShape__00000005_.Name = ``
	__ArtefactTypeShape__00000005_.X = 650.000000
	__ArtefactTypeShape__00000005_.Y = 115.000000
	__ArtefactTypeShape__00000005_.Width = 120.000000
	__ArtefactTypeShape__00000005_.Height = 25.000000
	__ArtefactTypeShape__00000005_.IsHidden = false

	__ArtefactTypeShape__00000006_.Name = ``
	__ArtefactTypeShape__00000006_.X = 450.000000
	__ArtefactTypeShape__00000006_.Y = 80.000000
	__ArtefactTypeShape__00000006_.Width = 120.000000
	__ArtefactTypeShape__00000006_.Height = 25.000000
	__ArtefactTypeShape__00000006_.IsHidden = false

	__Artist__00000000_.Name = `Frans Hals`
	__Artist__00000000_.ComputedPrefix = ``
	__Artist__00000000_.IsDead = true
	__Artist__00000000_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1666-01-01 00:00:00 +0000 UTC")

	__Artist__00000001_.Name = `Rembrandt van Rijn`
	__Artist__00000001_.ComputedPrefix = ``
	__Artist__00000001_.IsDead = true
	__Artist__00000001_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1669-01-01 00:00:00 +0000 UTC")

	__Artist__00000002_.Name = `Johannes Vermeer`
	__Artist__00000002_.ComputedPrefix = ``
	__Artist__00000002_.IsDead = true
	__Artist__00000002_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1675-01-01 00:00:00 +0000 UTC")

	__Artist__00000003_.Name = `Jan Steen`
	__Artist__00000003_.ComputedPrefix = ``
	__Artist__00000003_.IsDead = true
	__Artist__00000003_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1679-01-01 00:00:00 +0000 UTC")

	__Artist__00000004_.Name = `Pieter de Hooch`
	__Artist__00000004_.ComputedPrefix = ``
	__Artist__00000004_.IsDead = true
	__Artist__00000004_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1684-01-01 00:00:00 +0000 UTC")

	__Artist__00000005_.Name = `Jacob van Ruisdael`
	__Artist__00000005_.ComputedPrefix = ``
	__Artist__00000005_.IsDead = true
	__Artist__00000005_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1682-01-01 00:00:00 +0000 UTC")

	__ArtistShape__00000000_.Name = ``
	__ArtistShape__00000000_.X = 50.000000
	__ArtistShape__00000000_.Y = 365.000000
	__ArtistShape__00000000_.Width = 120.000000
	__ArtistShape__00000000_.Height = 30.000000
	__ArtistShape__00000000_.IsHidden = false
	__ArtistShape__00000000_.ImagePng_X = 0.000000
	__ArtistShape__00000000_.ImagePng_Y = 0.000000
	__ArtistShape__00000000_.ImagePng_Width = 0.000000
	__ArtistShape__00000000_.ImagePng_Height = 0.000000
	__ArtistShape__00000000_.ImagePng_X_Offset = 0.000000
	__ArtistShape__00000000_.ImagePng_Y_Offset = 0.000000
	__ArtistShape__00000000_.ImagePng_RectAnchorType = ""
	__ArtistShape__00000000_.ImagePngBase64Content = ``

	__ArtistShape__00000001_.Name = ``
	__ArtistShape__00000001_.X = 300.000000
	__ArtistShape__00000001_.Y = 440.000000
	__ArtistShape__00000001_.Width = 120.000000
	__ArtistShape__00000001_.Height = 30.000000
	__ArtistShape__00000001_.IsHidden = false
	__ArtistShape__00000001_.ImagePng_X = 0.000000
	__ArtistShape__00000001_.ImagePng_Y = 0.000000
	__ArtistShape__00000001_.ImagePng_Width = 0.000000
	__ArtistShape__00000001_.ImagePng_Height = 0.000000
	__ArtistShape__00000001_.ImagePng_X_Offset = 0.000000
	__ArtistShape__00000001_.ImagePng_Y_Offset = 0.000000
	__ArtistShape__00000001_.ImagePng_RectAnchorType = ""
	__ArtistShape__00000001_.ImagePngBase64Content = ``

	__ArtistShape__00000002_.Name = ``
	__ArtistShape__00000002_.X = 350.000000
	__ArtistShape__00000002_.Y = 565.000000
	__ArtistShape__00000002_.Width = 120.000000
	__ArtistShape__00000002_.Height = 30.000000
	__ArtistShape__00000002_.IsHidden = false
	__ArtistShape__00000002_.ImagePng_X = 0.000000
	__ArtistShape__00000002_.ImagePng_Y = 0.000000
	__ArtistShape__00000002_.ImagePng_Width = 0.000000
	__ArtistShape__00000002_.ImagePng_Height = 0.000000
	__ArtistShape__00000002_.ImagePng_X_Offset = 0.000000
	__ArtistShape__00000002_.ImagePng_Y_Offset = 0.000000
	__ArtistShape__00000002_.ImagePng_RectAnchorType = ""
	__ArtistShape__00000002_.ImagePngBase64Content = ``

	__ArtistShape__00000003_.Name = ``
	__ArtistShape__00000003_.X = 150.000000
	__ArtistShape__00000003_.Y = 540.000000
	__ArtistShape__00000003_.Width = 120.000000
	__ArtistShape__00000003_.Height = 30.000000
	__ArtistShape__00000003_.IsHidden = false
	__ArtistShape__00000003_.ImagePng_X = 0.000000
	__ArtistShape__00000003_.ImagePng_Y = 0.000000
	__ArtistShape__00000003_.ImagePng_Width = 0.000000
	__ArtistShape__00000003_.ImagePng_Height = 0.000000
	__ArtistShape__00000003_.ImagePng_X_Offset = 0.000000
	__ArtistShape__00000003_.ImagePng_Y_Offset = 0.000000
	__ArtistShape__00000003_.ImagePng_RectAnchorType = ""
	__ArtistShape__00000003_.ImagePngBase64Content = ``

	__ArtistShape__00000004_.Name = ``
	__ArtistShape__00000004_.X = 450.000000
	__ArtistShape__00000004_.Y = 565.000000
	__ArtistShape__00000004_.Width = 120.000000
	__ArtistShape__00000004_.Height = 30.000000
	__ArtistShape__00000004_.IsHidden = false
	__ArtistShape__00000004_.ImagePng_X = 0.000000
	__ArtistShape__00000004_.ImagePng_Y = 0.000000
	__ArtistShape__00000004_.ImagePng_Width = 0.000000
	__ArtistShape__00000004_.ImagePng_Height = 0.000000
	__ArtistShape__00000004_.ImagePng_X_Offset = 0.000000
	__ArtistShape__00000004_.ImagePng_Y_Offset = 0.000000
	__ArtistShape__00000004_.ImagePng_RectAnchorType = ""
	__ArtistShape__00000004_.ImagePngBase64Content = ``

	__ArtistShape__00000005_.Name = ``
	__ArtistShape__00000005_.X = 550.000000
	__ArtistShape__00000005_.Y = 515.000000
	__ArtistShape__00000005_.Width = 120.000000
	__ArtistShape__00000005_.Height = 30.000000
	__ArtistShape__00000005_.IsHidden = false
	__ArtistShape__00000005_.ImagePng_X = 0.000000
	__ArtistShape__00000005_.ImagePng_Y = 0.000000
	__ArtistShape__00000005_.ImagePng_Width = 0.000000
	__ArtistShape__00000005_.ImagePng_Height = 0.000000
	__ArtistShape__00000005_.ImagePng_X_Offset = 0.000000
	__ArtistShape__00000005_.ImagePng_Y_Offset = 0.000000
	__ArtistShape__00000005_.ImagePng_RectAnchorType = ""
	__ArtistShape__00000005_.ImagePngBase64Content = ``

	__Desk__00000000_.Name = `Desk`

	__Diagram__00000000_.Name = `Netherland Golden Century`
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable = false
	__Diagram__00000000_.IsNodeExpanded = false
	__Diagram__00000000_.IsMovementCategoryNodeExpanded = false
	__Diagram__00000000_.IsArtefactTypeCategoryNodeExpanded = false
	__Diagram__00000000_.IsArtistCategoryNodeExpanded = false
	__Diagram__00000000_.IsInfluenceCategoryNodeExpanded = false
	__Diagram__00000000_.IsMovementCategoryHidden = false
	__Diagram__00000000_.IsArtefactTypeCategoryHidden = false
	__Diagram__00000000_.IsArtistCategoryHidden = false
	__Diagram__00000000_.IsInfluenceCategoryHidden = false
	__Diagram__00000000_.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1560-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1700-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.NbYearsForIntervals = 10
	__Diagram__00000000_.XMargin = 20.000000
	__Diagram__00000000_.YMargin = 40.000000
	__Diagram__00000000_.Height = 900.000000
	__Diagram__00000000_.NextVerticalDateXMargin = 800.000000
	__Diagram__00000000_.RedColorCode = `#D23B22`
	__Diagram__00000000_.BackgroundGreyColorCode = `#DED6CA`
	__Diagram__00000000_.GrayColorCode = `#343434`
	__Diagram__00000000_.BottomBoxYOffset = 50.000000
	__Diagram__00000000_.BottomBoxWidth = 870.000000
	__Diagram__00000000_.BottomBoxHeigth = 150.000000
	__Diagram__00000000_.BottomBoxFontSize = `50px`
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

	__Influence__00000000_.Name = `Haarlem Mannerism to Pre-Rembrandtists`
	__Influence__00000000_.ComputedPrefix = ``
	__Influence__00000000_.IsHypothtical = false

	__Influence__00000001_.Name = `Pre-Rembrandtists to Rembrandt van Rijn`
	__Influence__00000001_.ComputedPrefix = ``
	__Influence__00000001_.IsHypothtical = false

	__Influence__00000002_.Name = `Dutch Caravaggisti to Rembrandt van Rijn`
	__Influence__00000002_.ComputedPrefix = ``
	__Influence__00000002_.IsHypothtical = false

	__Influence__00000003_.Name = `Haarlem Mannerism to Frans Hals`
	__Influence__00000003_.ComputedPrefix = ``
	__Influence__00000003_.IsHypothtical = false

	__Influence__00000004_.Name = `Frans Hals to Jan Steen`
	__Influence__00000004_.ComputedPrefix = ``
	__Influence__00000004_.IsHypothtical = false

	__Influence__00000005_.Name = `Rembrandt van Rijn to Delft School`
	__Influence__00000005_.ComputedPrefix = ``
	__Influence__00000005_.IsHypothtical = false

	__Influence__00000006_.Name = `Delft School to Johannes Vermeer`
	__Influence__00000006_.ComputedPrefix = ``
	__Influence__00000006_.IsHypothtical = false

	__Influence__00000007_.Name = `Delft School to Pieter de Hooch`
	__Influence__00000007_.ComputedPrefix = ``
	__Influence__00000007_.IsHypothtical = false

	__Influence__00000008_.Name = `Catholicism to Dutch Caravaggisti`
	__Influence__00000008_.ComputedPrefix = ``
	__Influence__00000008_.IsHypothtical = false

	__Influence__00000009_.Name = `Eighty Years' War to Protestantism`
	__Influence__00000009_.ComputedPrefix = ``
	__Influence__00000009_.IsHypothtical = false

	__Influence__00000010_.Name = `Eighty Years' War to Catholicism`
	__Influence__00000010_.ComputedPrefix = ``
	__Influence__00000010_.IsHypothtical = false

	__Influence__00000011_.Name = `Protestantism to Portraiture`
	__Influence__00000011_.ComputedPrefix = ``
	__Influence__00000011_.IsHypothtical = false

	__Influence__00000012_.Name = `Protestantism to Genre Painting`
	__Influence__00000012_.ComputedPrefix = ``
	__Influence__00000012_.IsHypothtical = false

	__Influence__00000013_.Name = `Protestantism to Landscape Painting`
	__Influence__00000013_.ComputedPrefix = ``
	__Influence__00000013_.IsHypothtical = false

	__Influence__00000014_.Name = `Portraiture to Frans Hals`
	__Influence__00000014_.ComputedPrefix = ``
	__Influence__00000014_.IsHypothtical = false

	__Influence__00000015_.Name = `History Painting to Pre-Rembrandtists`
	__Influence__00000015_.ComputedPrefix = ``
	__Influence__00000015_.IsHypothtical = false

	__Influence__00000016_.Name = `Landscape Painting to Jacob van Ruisdael`
	__Influence__00000016_.ComputedPrefix = ``
	__Influence__00000016_.IsHypothtical = false

	__Influence__00000017_.Name = `Genre Painting to Jan Steen`
	__Influence__00000017_.ComputedPrefix = ``
	__Influence__00000017_.IsHypothtical = false

	__InfluenceShape__00000000_.Name = `Haarlem Mannerism to Pre-Rembrandtists`
	__InfluenceShape__00000000_.IsHidden = false

	__InfluenceShape__00000001_.Name = `Pre-Rembrandtists to Rembrandt van Rijn`
	__InfluenceShape__00000001_.IsHidden = false

	__InfluenceShape__00000002_.Name = `Dutch Caravaggisti to Rembrandt van Rijn`
	__InfluenceShape__00000002_.IsHidden = false

	__InfluenceShape__00000003_.Name = `Haarlem Mannerism to Frans Hals`
	__InfluenceShape__00000003_.IsHidden = false

	__InfluenceShape__00000004_.Name = `Frans Hals to Jan Steen`
	__InfluenceShape__00000004_.IsHidden = false

	__InfluenceShape__00000005_.Name = `Rembrandt van Rijn to Delft School`
	__InfluenceShape__00000005_.IsHidden = false

	__InfluenceShape__00000006_.Name = `Delft School to Johannes Vermeer`
	__InfluenceShape__00000006_.IsHidden = false

	__InfluenceShape__00000007_.Name = `Delft School to Pieter de Hooch`
	__InfluenceShape__00000007_.IsHidden = false

	__InfluenceShape__00000008_.Name = `Catholicism to Dutch Caravaggisti`
	__InfluenceShape__00000008_.IsHidden = false

	__InfluenceShape__00000009_.Name = `Eighty Years' War to Protestantism`
	__InfluenceShape__00000009_.IsHidden = false

	__InfluenceShape__00000010_.Name = `Eighty Years' War to Catholicism`
	__InfluenceShape__00000010_.IsHidden = false

	__InfluenceShape__00000011_.Name = `Protestantism to Portraiture`
	__InfluenceShape__00000011_.IsHidden = false

	__InfluenceShape__00000012_.Name = `Protestantism to Genre Painting`
	__InfluenceShape__00000012_.IsHidden = false

	__InfluenceShape__00000013_.Name = `Protestantism to Landscape Painting`
	__InfluenceShape__00000013_.IsHidden = false

	__InfluenceShape__00000014_.Name = `Portraiture to Frans Hals`
	__InfluenceShape__00000014_.IsHidden = false

	__InfluenceShape__00000015_.Name = `History Painting to Pre-Rembrandtists`
	__InfluenceShape__00000015_.IsHidden = false

	__InfluenceShape__00000016_.Name = `Landscape Painting to Jacob van Ruisdael`
	__InfluenceShape__00000016_.IsHidden = false

	__InfluenceShape__00000017_.Name = `Genre Painting to Jan Steen`
	__InfluenceShape__00000017_.IsHidden = false

	__Movement__00000000_.Name = `Haarlem Mannerism`
	__Movement__00000000_.ComputedPrefix = ``
	__Movement__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1585-01-01 00:00:00 +0000 UTC")
	__Movement__00000000_.HideDate = false
	__Movement__00000000_.HasTaxonomicFilter = false
	__Movement__00000000_.TaxonomicFilter = ``
	__Movement__00000000_.IsFeatured = false
	__Movement__00000000_.FeaturePrefix = ``
	__Movement__00000000_.IsMajor = false
	__Movement__00000000_.IsMinor = false
	__Movement__00000000_.AdditionnalName = ``

	__Movement__00000001_.Name = `Pre-Rembrandtists`
	__Movement__00000001_.ComputedPrefix = ``
	__Movement__00000001_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1610-01-01 00:00:00 +0000 UTC")
	__Movement__00000001_.HideDate = false
	__Movement__00000001_.HasTaxonomicFilter = false
	__Movement__00000001_.TaxonomicFilter = ``
	__Movement__00000001_.IsFeatured = false
	__Movement__00000001_.FeaturePrefix = ``
	__Movement__00000001_.IsMajor = false
	__Movement__00000001_.IsMinor = false
	__Movement__00000001_.AdditionnalName = ``

	__Movement__00000002_.Name = `Dutch Caravaggisti`
	__Movement__00000002_.ComputedPrefix = ``
	__Movement__00000002_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1620-01-01 00:00:00 +0000 UTC")
	__Movement__00000002_.HideDate = false
	__Movement__00000002_.HasTaxonomicFilter = false
	__Movement__00000002_.TaxonomicFilter = ``
	__Movement__00000002_.IsFeatured = false
	__Movement__00000002_.FeaturePrefix = ``
	__Movement__00000002_.IsMajor = false
	__Movement__00000002_.IsMinor = false
	__Movement__00000002_.AdditionnalName = ``

	__Movement__00000003_.Name = `Delft School`
	__Movement__00000003_.ComputedPrefix = ``
	__Movement__00000003_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1650-01-01 00:00:00 +0000 UTC")
	__Movement__00000003_.HideDate = false
	__Movement__00000003_.HasTaxonomicFilter = false
	__Movement__00000003_.TaxonomicFilter = ``
	__Movement__00000003_.IsFeatured = false
	__Movement__00000003_.FeaturePrefix = ``
	__Movement__00000003_.IsMajor = false
	__Movement__00000003_.IsMinor = false
	__Movement__00000003_.AdditionnalName = ``

	__MovementShape__00000000_.Name = ``
	__MovementShape__00000000_.X = 100.000000
	__MovementShape__00000000_.Y = 165.000000
	__MovementShape__00000000_.Width = 120.000000
	__MovementShape__00000000_.Height = 30.000000
	__MovementShape__00000000_.IsHidden = false

	__MovementShape__00000001_.Name = ``
	__MovementShape__00000001_.X = 200.000000
	__MovementShape__00000001_.Y = 290.000000
	__MovementShape__00000001_.Width = 120.000000
	__MovementShape__00000001_.Height = 30.000000
	__MovementShape__00000001_.IsHidden = false

	__MovementShape__00000002_.Name = ``
	__MovementShape__00000002_.X = 650.000000
	__MovementShape__00000002_.Y = 340.000000
	__MovementShape__00000002_.Width = 120.000000
	__MovementShape__00000002_.Height = 30.000000
	__MovementShape__00000002_.IsHidden = false

	__MovementShape__00000003_.Name = ``
	__MovementShape__00000003_.X = 400.000000
	__MovementShape__00000003_.Y = 490.000000
	__MovementShape__00000003_.Width = 120.000000
	__MovementShape__00000003_.Height = 30.000000
	__MovementShape__00000003_.IsHidden = false

	__Place__00000000_.Name = `Haarlem`

	__Place__00000001_.Name = `Amsterdam`

	__Place__00000002_.Name = `Delft`

	__Place__00000003_.Name = `Leiden`

	__Place__00000004_.Name = `Utrecht`

	// insertion point for setup of pointers
	__ArtefactTypeShape__00000000_.ArtefactType = __ArtefactType__00000000_
	__ArtefactTypeShape__00000001_.ArtefactType = __ArtefactType__00000001_
	__ArtefactTypeShape__00000002_.ArtefactType = __ArtefactType__00000002_
	__ArtefactTypeShape__00000003_.ArtefactType = __ArtefactType__00000003_
	__ArtefactTypeShape__00000004_.ArtefactType = __ArtefactType__00000004_
	__ArtefactTypeShape__00000005_.ArtefactType = __ArtefactType__00000005_
	__ArtefactTypeShape__00000006_.ArtefactType = __ArtefactType__00000006_
	__Artist__00000000_.Place = __Place__00000000_
	__Artist__00000001_.Place = __Place__00000001_
	__Artist__00000002_.Place = __Place__00000002_
	__Artist__00000003_.Place = __Place__00000003_
	__Artist__00000004_.Place = __Place__00000002_
	__Artist__00000005_.Place = __Place__00000000_
	__ArtistShape__00000000_.Artist = __Artist__00000000_
	__ArtistShape__00000001_.Artist = __Artist__00000001_
	__ArtistShape__00000002_.Artist = __Artist__00000002_
	__ArtistShape__00000003_.Artist = __Artist__00000003_
	__ArtistShape__00000004_.Artist = __Artist__00000004_
	__ArtistShape__00000005_.Artist = __Artist__00000005_
	__Desk__00000000_.SelectedDiagram = __Diagram__00000000_
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000000_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000001_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000002_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000003_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000000_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000001_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000002_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000003_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000004_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000005_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000006_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000000_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000001_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000002_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000003_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000004_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000005_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000000_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000001_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000002_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000003_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000004_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000005_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000006_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000007_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000008_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000009_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000010_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000011_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000012_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000013_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000014_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000015_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000016_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000017_)
	__Influence__00000000_.SourceMovement = __Movement__00000000_
	__Influence__00000000_.SourceArtefactType = nil
	__Influence__00000000_.SourceArtist = nil
	__Influence__00000000_.TargetMovement = __Movement__00000001_
	__Influence__00000000_.TargetArtefactType = nil
	__Influence__00000000_.TargetArtist = nil
	__Influence__00000001_.SourceMovement = __Movement__00000001_
	__Influence__00000001_.SourceArtefactType = nil
	__Influence__00000001_.SourceArtist = nil
	__Influence__00000001_.TargetMovement = nil
	__Influence__00000001_.TargetArtefactType = nil
	__Influence__00000001_.TargetArtist = __Artist__00000001_
	__Influence__00000002_.SourceMovement = __Movement__00000002_
	__Influence__00000002_.SourceArtefactType = nil
	__Influence__00000002_.SourceArtist = nil
	__Influence__00000002_.TargetMovement = nil
	__Influence__00000002_.TargetArtefactType = nil
	__Influence__00000002_.TargetArtist = __Artist__00000001_
	__Influence__00000003_.SourceMovement = __Movement__00000000_
	__Influence__00000003_.SourceArtefactType = nil
	__Influence__00000003_.SourceArtist = nil
	__Influence__00000003_.TargetMovement = nil
	__Influence__00000003_.TargetArtefactType = nil
	__Influence__00000003_.TargetArtist = __Artist__00000000_
	__Influence__00000004_.SourceMovement = nil
	__Influence__00000004_.SourceArtefactType = nil
	__Influence__00000004_.SourceArtist = __Artist__00000000_
	__Influence__00000004_.TargetMovement = nil
	__Influence__00000004_.TargetArtefactType = nil
	__Influence__00000004_.TargetArtist = __Artist__00000003_
	__Influence__00000005_.SourceMovement = nil
	__Influence__00000005_.SourceArtefactType = nil
	__Influence__00000005_.SourceArtist = __Artist__00000001_
	__Influence__00000005_.TargetMovement = __Movement__00000003_
	__Influence__00000005_.TargetArtefactType = nil
	__Influence__00000005_.TargetArtist = nil
	__Influence__00000006_.SourceMovement = __Movement__00000003_
	__Influence__00000006_.SourceArtefactType = nil
	__Influence__00000006_.SourceArtist = nil
	__Influence__00000006_.TargetMovement = nil
	__Influence__00000006_.TargetArtefactType = nil
	__Influence__00000006_.TargetArtist = __Artist__00000002_
	__Influence__00000007_.SourceMovement = __Movement__00000003_
	__Influence__00000007_.SourceArtefactType = nil
	__Influence__00000007_.SourceArtist = nil
	__Influence__00000007_.TargetMovement = nil
	__Influence__00000007_.TargetArtefactType = nil
	__Influence__00000007_.TargetArtist = __Artist__00000004_
	__Influence__00000008_.SourceMovement = nil
	__Influence__00000008_.SourceArtefactType = __ArtefactType__00000005_
	__Influence__00000008_.SourceArtist = nil
	__Influence__00000008_.TargetMovement = __Movement__00000002_
	__Influence__00000008_.TargetArtefactType = nil
	__Influence__00000008_.TargetArtist = nil
	__Influence__00000009_.SourceMovement = nil
	__Influence__00000009_.SourceArtefactType = __ArtefactType__00000006_
	__Influence__00000009_.SourceArtist = nil
	__Influence__00000009_.TargetMovement = nil
	__Influence__00000009_.TargetArtefactType = __ArtefactType__00000004_
	__Influence__00000009_.TargetArtist = nil
	__Influence__00000010_.SourceMovement = nil
	__Influence__00000010_.SourceArtefactType = __ArtefactType__00000006_
	__Influence__00000010_.SourceArtist = nil
	__Influence__00000010_.TargetMovement = nil
	__Influence__00000010_.TargetArtefactType = __ArtefactType__00000005_
	__Influence__00000010_.TargetArtist = nil
	__Influence__00000011_.SourceMovement = nil
	__Influence__00000011_.SourceArtefactType = __ArtefactType__00000004_
	__Influence__00000011_.SourceArtist = nil
	__Influence__00000011_.TargetMovement = nil
	__Influence__00000011_.TargetArtefactType = __ArtefactType__00000001_
	__Influence__00000011_.TargetArtist = nil
	__Influence__00000012_.SourceMovement = nil
	__Influence__00000012_.SourceArtefactType = __ArtefactType__00000004_
	__Influence__00000012_.SourceArtist = nil
	__Influence__00000012_.TargetMovement = nil
	__Influence__00000012_.TargetArtefactType = __ArtefactType__00000002_
	__Influence__00000012_.TargetArtist = nil
	__Influence__00000013_.SourceMovement = nil
	__Influence__00000013_.SourceArtefactType = __ArtefactType__00000004_
	__Influence__00000013_.SourceArtist = nil
	__Influence__00000013_.TargetMovement = nil
	__Influence__00000013_.TargetArtefactType = __ArtefactType__00000003_
	__Influence__00000013_.TargetArtist = nil
	__Influence__00000014_.SourceMovement = nil
	__Influence__00000014_.SourceArtefactType = __ArtefactType__00000001_
	__Influence__00000014_.SourceArtist = nil
	__Influence__00000014_.TargetMovement = nil
	__Influence__00000014_.TargetArtefactType = nil
	__Influence__00000014_.TargetArtist = __Artist__00000000_
	__Influence__00000015_.SourceMovement = nil
	__Influence__00000015_.SourceArtefactType = __ArtefactType__00000000_
	__Influence__00000015_.SourceArtist = nil
	__Influence__00000015_.TargetMovement = __Movement__00000001_
	__Influence__00000015_.TargetArtefactType = nil
	__Influence__00000015_.TargetArtist = nil
	__Influence__00000016_.SourceMovement = nil
	__Influence__00000016_.SourceArtefactType = __ArtefactType__00000003_
	__Influence__00000016_.SourceArtist = nil
	__Influence__00000016_.TargetMovement = nil
	__Influence__00000016_.TargetArtefactType = nil
	__Influence__00000016_.TargetArtist = __Artist__00000005_
	__Influence__00000017_.SourceMovement = nil
	__Influence__00000017_.SourceArtefactType = __ArtefactType__00000002_
	__Influence__00000017_.SourceArtist = nil
	__Influence__00000017_.TargetMovement = nil
	__Influence__00000017_.TargetArtefactType = nil
	__Influence__00000017_.TargetArtist = __Artist__00000003_
	__InfluenceShape__00000000_.Influence = __Influence__00000000_
	__InfluenceShape__00000001_.Influence = __Influence__00000001_
	__InfluenceShape__00000002_.Influence = __Influence__00000002_
	__InfluenceShape__00000003_.Influence = __Influence__00000003_
	__InfluenceShape__00000004_.Influence = __Influence__00000004_
	__InfluenceShape__00000005_.Influence = __Influence__00000005_
	__InfluenceShape__00000006_.Influence = __Influence__00000006_
	__InfluenceShape__00000007_.Influence = __Influence__00000007_
	__InfluenceShape__00000008_.Influence = __Influence__00000008_
	__InfluenceShape__00000009_.Influence = __Influence__00000009_
	__InfluenceShape__00000010_.Influence = __Influence__00000010_
	__InfluenceShape__00000011_.Influence = __Influence__00000011_
	__InfluenceShape__00000012_.Influence = __Influence__00000012_
	__InfluenceShape__00000013_.Influence = __Influence__00000013_
	__InfluenceShape__00000014_.Influence = __Influence__00000014_
	__InfluenceShape__00000015_.Influence = __Influence__00000015_
	__InfluenceShape__00000016_.Influence = __Influence__00000016_
	__InfluenceShape__00000017_.Influence = __Influence__00000017_
	__Movement__00000000_.Places = append(__Movement__00000000_.Places, __Place__00000000_)
	__Movement__00000001_.Places = append(__Movement__00000001_.Places, __Place__00000001_)
	__Movement__00000002_.Places = append(__Movement__00000002_.Places, __Place__00000004_)
	__Movement__00000003_.Places = append(__Movement__00000003_.Places, __Place__00000002_)
	__MovementShape__00000000_.Movement = __Movement__00000000_
	__MovementShape__00000001_.Movement = __Movement__00000001_
	__MovementShape__00000002_.Movement = __Movement__00000002_
	__MovementShape__00000003_.Movement = __Movement__00000003_
}

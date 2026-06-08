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

	__ArtefactTypeShape__00000000_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000001_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000002_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000003_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)

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

	__InfluenceShape__00000000_ := (&models.InfluenceShape{Name: `Haarlem Mannerism to Pre-Rembrandtists`}).Stage(stage)
	__InfluenceShape__00000001_ := (&models.InfluenceShape{Name: `Pre-Rembrandtists to Rembrandt van Rijn`}).Stage(stage)
	__InfluenceShape__00000002_ := (&models.InfluenceShape{Name: `Dutch Caravaggisti to Rembrandt van Rijn`}).Stage(stage)
	__InfluenceShape__00000003_ := (&models.InfluenceShape{Name: `Haarlem Mannerism to Frans Hals`}).Stage(stage)
	__InfluenceShape__00000004_ := (&models.InfluenceShape{Name: `Frans Hals to Jan Steen`}).Stage(stage)
	__InfluenceShape__00000005_ := (&models.InfluenceShape{Name: `Rembrandt van Rijn to Delft School`}).Stage(stage)
	__InfluenceShape__00000006_ := (&models.InfluenceShape{Name: `Delft School to Johannes Vermeer`}).Stage(stage)
	__InfluenceShape__00000007_ := (&models.InfluenceShape{Name: `Delft School to Pieter de Hooch`}).Stage(stage)

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

	// insertion point for initialization of values

	__ArtefactType__00000000_.Name = `History Painting`
	__ArtefactType__00000001_.Name = `Portraiture`
	__ArtefactType__00000002_.Name = `Genre Painting`
	__ArtefactType__00000003_.Name = `Landscape Painting`

	__ArtefactTypeShape__00000000_.X = 100.0
	__ArtefactTypeShape__00000000_.Y = 50.0
	__ArtefactTypeShape__00000000_.Width = 120.0
	__ArtefactTypeShape__00000000_.Height = 25.0

	__ArtefactTypeShape__00000001_.X = 250.0
	__ArtefactTypeShape__00000001_.Y = 50.0
	__ArtefactTypeShape__00000001_.Width = 120.0
	__ArtefactTypeShape__00000001_.Height = 25.0

	__ArtefactTypeShape__00000002_.X = 400.0
	__ArtefactTypeShape__00000002_.Y = 50.0
	__ArtefactTypeShape__00000002_.Width = 120.0
	__ArtefactTypeShape__00000002_.Height = 25.0

	__ArtefactTypeShape__00000003_.X = 550.0
	__ArtefactTypeShape__00000003_.Y = 50.0
	__ArtefactTypeShape__00000003_.Width = 120.0
	__ArtefactTypeShape__00000003_.Height = 25.0

	__Artist__00000000_.Name = `Frans Hals`
	__Artist__00000000_.IsDead = true
	__Artist__00000000_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1666-01-01 00:00:00 +0000 UTC")

	__Artist__00000001_.Name = `Rembrandt van Rijn`
	__Artist__00000001_.IsDead = true
	__Artist__00000001_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1669-01-01 00:00:00 +0000 UTC")

	__Artist__00000002_.Name = `Johannes Vermeer`
	__Artist__00000002_.IsDead = true
	__Artist__00000002_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1675-01-01 00:00:00 +0000 UTC")

	__Artist__00000003_.Name = `Jan Steen`
	__Artist__00000003_.IsDead = true
	__Artist__00000003_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1679-01-01 00:00:00 +0000 UTC")

	__Artist__00000004_.Name = `Pieter de Hooch`
	__Artist__00000004_.IsDead = true
	__Artist__00000004_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1684-01-01 00:00:00 +0000 UTC")

	__Artist__00000005_.Name = `Jacob van Ruisdael`
	__Artist__00000005_.IsDead = true
	__Artist__00000005_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1682-01-01 00:00:00 +0000 UTC")

	__ArtistShape__00000000_.X = 150.0
	__ArtistShape__00000000_.Y = 300.0
	__ArtistShape__00000000_.Width = 80.0
	__ArtistShape__00000000_.Height = 30.0

	__ArtistShape__00000001_.X = 300.0
	__ArtistShape__00000001_.Y = 400.0
	__ArtistShape__00000001_.Width = 120.0
	__ArtistShape__00000001_.Height = 30.0

	__ArtistShape__00000002_.X = 450.0
	__ArtistShape__00000002_.Y = 600.0
	__ArtistShape__00000002_.Width = 120.0
	__ArtistShape__00000002_.Height = 30.0

	__ArtistShape__00000003_.X = 150.0
	__ArtistShape__00000003_.Y = 500.0
	__ArtistShape__00000003_.Width = 80.0
	__ArtistShape__00000003_.Height = 30.0

	__ArtistShape__00000004_.X = 600.0
	__ArtistShape__00000004_.Y = 600.0
	__ArtistShape__00000004_.Width = 120.0
	__ArtistShape__00000004_.Height = 30.0

	__ArtistShape__00000005_.X = 600.0
	__ArtistShape__00000005_.Y = 400.0
	__ArtistShape__00000005_.Width = 120.0
	__ArtistShape__00000005_.Height = 30.0

	__Desk__00000000_.Name = `Desk`

	__Diagram__00000000_.Name = `Netherland Golden Century`
	__Diagram__00000000_.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1580-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1700-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.NbYearsForIntervals = 10
	__Diagram__00000000_.XMargin = 20.0
	__Diagram__00000000_.YMargin = 40.0
	__Diagram__00000000_.Height = 900.0
	__Diagram__00000000_.NextVerticalDateXMargin = 800.0
	__Diagram__00000000_.RedColorCode = `#D23B22`
	__Diagram__00000000_.BackgroundGreyColorCode = `#DED6CA`
	__Diagram__00000000_.GrayColorCode = `#343434`
	__Diagram__00000000_.BottomBoxYOffset = 50.0
	__Diagram__00000000_.BottomBoxWidth = 770.0
	__Diagram__00000000_.BottomBoxHeigth = 150.0
	__Diagram__00000000_.BottomBoxFontSize = `50px`
	__Diagram__00000000_.BottomBoxFontWeigth = `500`
	__Diagram__00000000_.BottomBoxFontFamily = `ChunkFive, sans-serif`
	__Diagram__00000000_.BottomBoxLetterSpacing = `1`
	__Diagram__00000000_.BottomBoxLetterColorCode = `#debdaaff`
	__Diagram__00000000_.MovementRectAnchorType = models.RECT_RIGHT
	__Diagram__00000000_.MovementTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.MovementDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.MovementFontSize = `14px`
	__Diagram__00000000_.MovementFontWeigth = `Thin`
	__Diagram__00000000_.MovementFontFamily = `Futura, sans serif`
	__Diagram__00000000_.MovementDateRectAnchorType = models.RECT_BOTTOM_LEFT
	__Diagram__00000000_.MovementDateTextAnchorType = models.TEXT_ANCHOR_START
	__Diagram__00000000_.MovementDateTextDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.MovementPlacesRectAnchorType = models.RECT_BOTTOM_RIGHT
	__Diagram__00000000_.MovementPlacesTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.MovementPlacesDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtefactTypeFontSize = `12px`
	__Diagram__00000000_.ArtefactTypeFontFamily = `Futura`
	__Diagram__00000000_.ArtefactTypeRectAnchorType = models.RECT_CENTER
	__Diagram__00000000_.ArtefactDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtefactTypeStrokeWidth = 3.0
	__Diagram__00000000_.ArtistRectAnchorType = models.RECT_CENTER
	__Diagram__00000000_.ArtistTextAnchorType = models.TEXT_ANCHOR_CENTER
	__Diagram__00000000_.ArtistDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtistFontSize = `14px`
	__Diagram__00000000_.ArtistFontWeigth = `100`
	__Diagram__00000000_.ArtistFontFamily = `Futura`
	__Diagram__00000000_.ArtistDateRectAnchorType = models.RECT_BOTTOM_LEFT
	__Diagram__00000000_.ArtistDateTextAnchorType = models.TEXT_ANCHOR_START
	__Diagram__00000000_.ArtistDateDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.ArtistPlacesRectAnchorType = models.RECT_BOTTOM_RIGHT
	__Diagram__00000000_.ArtistPlacesTextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__00000000_.ArtistPlacesDominantBaselineType = models.DominantBaselineCentral
	__Diagram__00000000_.InfluenceArrowSize = 6.0
	__Diagram__00000000_.InfluenceArrowStartOffset = 19.0
	__Diagram__00000000_.InfluenceArrowEndOffset = 9.0
	__Diagram__00000000_.InfluenceCornerRadius = 20.0
	__Diagram__00000000_.InfluenceDashedLinePattern = `7 3`

	__Influence__00000000_.Name = `Haarlem Mannerism to Pre-Rembrandtists`
	__Influence__00000001_.Name = `Pre-Rembrandtists to Rembrandt van Rijn`
	__Influence__00000002_.Name = `Dutch Caravaggisti to Rembrandt van Rijn`
	__Influence__00000003_.Name = `Haarlem Mannerism to Frans Hals`
	__Influence__00000004_.Name = `Frans Hals to Jan Steen`
	__Influence__00000005_.Name = `Rembrandt van Rijn to Delft School`
	__Influence__00000006_.Name = `Delft School to Johannes Vermeer`
	__Influence__00000007_.Name = `Delft School to Pieter de Hooch`

	__InfluenceShape__00000000_.Name = `Haarlem Mannerism to Pre-Rembrandtists`
	__InfluenceShape__00000001_.Name = `Pre-Rembrandtists to Rembrandt van Rijn`
	__InfluenceShape__00000002_.Name = `Dutch Caravaggisti to Rembrandt van Rijn`
	__InfluenceShape__00000003_.Name = `Haarlem Mannerism to Frans Hals`
	__InfluenceShape__00000004_.Name = `Frans Hals to Jan Steen`
	__InfluenceShape__00000005_.Name = `Rembrandt van Rijn to Delft School`
	__InfluenceShape__00000006_.Name = `Delft School to Johannes Vermeer`
	__InfluenceShape__00000007_.Name = `Delft School to Pieter de Hooch`

	__Movement__00000000_.Name = `Haarlem Mannerism`
	__Movement__00000001_.Name = `Pre-Rembrandtists`
	__Movement__00000002_.Name = `Dutch Caravaggisti`
	__Movement__00000003_.Name = `Delft School`

	__MovementShape__00000000_.X = 100.0
	__MovementShape__00000000_.Y = 150.0
	__MovementShape__00000000_.Width = 120.0
	__MovementShape__00000000_.Height = 30.0

	__MovementShape__00000001_.X = 250.0
	__MovementShape__00000001_.Y = 200.0
	__MovementShape__00000001_.Width = 120.0
	__MovementShape__00000001_.Height = 30.0

	__MovementShape__00000002_.X = 400.0
	__MovementShape__00000002_.Y = 250.0
	__MovementShape__00000002_.Width = 120.0
	__MovementShape__00000002_.Height = 30.0

	__MovementShape__00000003_.X = 550.0
	__MovementShape__00000003_.Y = 450.0
	__MovementShape__00000003_.Width = 120.0
	__MovementShape__00000003_.Height = 30.0

	__Place__00000000_.Name = `Haarlem`
	__Place__00000001_.Name = `Amsterdam`
	__Place__00000002_.Name = `Delft`
	__Place__00000003_.Name = `Leiden`

	// insertion point for setup of pointers
	__ArtefactTypeShape__00000000_.ArtefactType = __ArtefactType__00000000_
	__ArtefactTypeShape__00000001_.ArtefactType = __ArtefactType__00000001_
	__ArtefactTypeShape__00000002_.ArtefactType = __ArtefactType__00000002_
	__ArtefactTypeShape__00000003_.ArtefactType = __ArtefactType__00000003_

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

	__InfluenceShape__00000000_.Influence = __Influence__00000000_
	__InfluenceShape__00000001_.Influence = __Influence__00000001_
	__InfluenceShape__00000002_.Influence = __Influence__00000002_
	__InfluenceShape__00000003_.Influence = __Influence__00000003_
	__InfluenceShape__00000004_.Influence = __Influence__00000004_
	__InfluenceShape__00000005_.Influence = __Influence__00000005_
	__InfluenceShape__00000006_.Influence = __Influence__00000006_
	__InfluenceShape__00000007_.Influence = __Influence__00000007_

	__MovementShape__00000000_.Movement = __Movement__00000000_
	__MovementShape__00000001_.Movement = __Movement__00000001_
	__MovementShape__00000002_.Movement = __Movement__00000002_
	__MovementShape__00000003_.Movement = __Movement__00000003_
}

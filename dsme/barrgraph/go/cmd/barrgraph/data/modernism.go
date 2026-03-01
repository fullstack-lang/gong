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
	// ---------------------------------------------------------
	// 1. DECLARATIONS
	// ---------------------------------------------------------

	__ArtefactType__00000000_ := (&models.ArtefactType{Name: `CRAFTSMANSHIP`}).Stage(stage)
	__ArtefactType__00000001_ := (&models.ArtefactType{Name: `INDUSTRIAL PRODUCTION`}).Stage(stage)
	__ArtefactType__00000002_ := (&models.ArtefactType{Name: `NEW MATERIALS`}).Stage(stage)

	__ArtefactTypeShape__00000000_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000001_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__ArtefactTypeShape__00000002_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)

	__Artist__00000000_ := (&models.Artist{Name: `Mackintosh`}).Stage(stage)
	__Artist__00000001_ := (&models.Artist{Name: `Behrens`}).Stage(stage)
	__Artist__00000002_ := (&models.Artist{Name: `Gropius`}).Stage(stage)
	__Artist__00000003_ := (&models.Artist{Name: `Le Corbusier`}).Stage(stage)
	__Artist__00000004_ := (&models.Artist{Name: `Aalto`}).Stage(stage)

	__ArtistShape__00000000_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000001_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000002_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000003_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtistShape__00000004_ := (&models.ArtistShape{Name: ``}).Stage(stage)

	__Desk__00000000_ := (&models.Desk{Name: `Desk`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Modernism in Design`}).Stage(stage)

	__Influence__00000000_ := (&models.Influence{Name: `Craft to Arts & Crafts`}).Stage(stage)
	__Influence__00000001_ := (&models.Influence{Name: `Arts & Crafts to Wiener`}).Stage(stage)
	__Influence__00000002_ := (&models.Influence{Name: `Ind to Werkbund`}).Stage(stage)
	__Influence__00000003_ := (&models.Influence{Name: `Wiener to Werkbund`}).Stage(stage)
	__Influence__00000004_ := (&models.Influence{Name: `Werkbund to Bauhaus`}).Stage(stage)
	__Influence__00000005_ := (&models.Influence{Name: `Constructivism to Bauhaus`}).Stage(stage)
	__Influence__00000006_ := (&models.Influence{Name: `DeStijl to Bauhaus`}).Stage(stage)
	__Influence__00000007_ := (&models.Influence{Name: `Bauhaus to IntStyle`}).Stage(stage)
	__Influence__00000008_ := (&models.Influence{Name: `Mat to Organic`}).Stage(stage)
	__Influence__00000009_ := (&models.Influence{Name: `Organic to Scandi`}).Stage(stage)
	__Influence__00000010_ := (&models.Influence{Name: `IntStyle to Scandi`}).Stage(stage)

	__InfluenceShape__00000000_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000001_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000002_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000003_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000004_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000005_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000006_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000007_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000008_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000009_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__InfluenceShape__00000010_ := (&models.InfluenceShape{Name: ``}).Stage(stage)

	__Movement__00000000_ := (&models.Movement{Name: `ARTS & CRAFTS`}).Stage(stage)
	__Movement__00000001_ := (&models.Movement{Name: `WIENER WERKSTÄTTE`}).Stage(stage)
	__Movement__00000002_ := (&models.Movement{Name: `DEUTSCHER WERKBUND`}).Stage(stage)
	__Movement__00000003_ := (&models.Movement{Name: `CONSTRUCTIVISM`}).Stage(stage)
	__Movement__00000004_ := (&models.Movement{Name: `DE STIJL`}).Stage(stage)
	__Movement__00000005_ := (&models.Movement{Name: `BAUHAUS`}).Stage(stage)
	__Movement__00000006_ := (&models.Movement{Name: `INTERNATIONAL STYLE`}).Stage(stage)
	__Movement__00000007_ := (&models.Movement{Name: `ORGANIC DESIGN`}).Stage(stage)
	__Movement__00000008_ := (&models.Movement{Name: `SCANDINAVIAN MODERN`}).Stage(stage)

	__MovementShape__00000000_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000001_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000002_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000003_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000004_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000005_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000006_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000007_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MovementShape__00000008_ := (&models.MovementShape{Name: ``}).Stage(stage)

	__Place__00000000_ := (&models.Place{Name: `London`}).Stage(stage)
	__Place__00000001_ := (&models.Place{Name: `Vienna`}).Stage(stage)
	__Place__00000002_ := (&models.Place{Name: `Munich`}).Stage(stage)
	__Place__00000003_ := (&models.Place{Name: `Weimar`}).Stage(stage)
	__Place__00000004_ := (&models.Place{Name: `Dessau`}).Stage(stage)
	__Place__00000005_ := (&models.Place{Name: `Moscow`}).Stage(stage)
	__Place__00000006_ := (&models.Place{Name: `Paris`}).Stage(stage)
	__Place__00000007_ := (&models.Place{Name: `Helsinki`}).Stage(stage)
	__Place__00000008_ := (&models.Place{Name: `Copenhagen`}).Stage(stage)

	// ---------------------------------------------------------
	// 2. INITIALIZATIONS
	// ---------------------------------------------------------

	__ArtefactType__00000000_.Name = `CRAFTSMANSHIP`
	__ArtefactType__00000001_.Name = `INDUSTRIAL PRODUCTION`
	__ArtefactType__00000002_.Name = `NEW MATERIALS`

	__ArtefactTypeShape__00000000_.Name = ``
	__ArtefactTypeShape__00000000_.X = 100.000000
	__ArtefactTypeShape__00000000_.Y = 50.000000
	__ArtefactTypeShape__00000000_.Width = 140.000000
	__ArtefactTypeShape__00000000_.Height = 25.000000

	__ArtefactTypeShape__00000001_.Name = ``
	__ArtefactTypeShape__00000001_.X = 550.000000
	__ArtefactTypeShape__00000001_.Y = 50.000000
	__ArtefactTypeShape__00000001_.Width = 180.000000
	__ArtefactTypeShape__00000001_.Height = 25.000000

	__ArtefactTypeShape__00000002_.Name = ``
	__ArtefactTypeShape__00000002_.X = 100.000000
	__ArtefactTypeShape__00000002_.Y = 300.000000
	__ArtefactTypeShape__00000002_.Width = 140.000000
	__ArtefactTypeShape__00000002_.Height = 25.000000

	__Artist__00000000_.Name = `Mackintosh`
	__Artist__00000000_.IsDead = true
	__Artist__00000000_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1928-01-01 00:00:00 +0000 UTC")

	__Artist__00000001_.Name = `Behrens`
	__Artist__00000001_.IsDead = true
	__Artist__00000001_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1940-01-01 00:00:00 +0000 UTC")

	__Artist__00000002_.Name = `Gropius`
	__Artist__00000002_.IsDead = true
	__Artist__00000002_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1969-01-01 00:00:00 +0000 UTC")

	__Artist__00000003_.Name = `Le Corbusier`
	__Artist__00000003_.IsDead = true
	__Artist__00000003_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1965-01-01 00:00:00 +0000 UTC")

	__Artist__00000004_.Name = `Aalto`
	__Artist__00000004_.IsDead = true
	__Artist__00000004_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1976-01-01 00:00:00 +0000 UTC")

	__ArtistShape__00000000_.Name = ``
	__ArtistShape__00000000_.X = 250.000000
	__ArtistShape__00000000_.Y = 100.000000
	__ArtistShape__00000000_.Width = 80.000000
	__ArtistShape__00000000_.Height = 30.000000

	__ArtistShape__00000001_.Name = ``
	__ArtistShape__00000001_.X = 550.000000
	__ArtistShape__00000001_.Y = 150.000000
	__ArtistShape__00000001_.Width = 80.000000
	__ArtistShape__00000001_.Height = 30.000000

	__ArtistShape__00000002_.Name = ``
	__ArtistShape__00000002_.X = 350.000000
	__ArtistShape__00000002_.Y = 270.000000
	__ArtistShape__00000002_.Width = 80.000000
	__ArtistShape__00000002_.Height = 30.000000

	__ArtistShape__00000003_.Name = ``
	__ArtistShape__00000003_.X = 600.000000
	__ArtistShape__00000003_.Y = 320.000000
	__ArtistShape__00000003_.Width = 80.000000
	__ArtistShape__00000003_.Height = 30.000000

	__ArtistShape__00000004_.Name = ``
	__ArtistShape__00000004_.X = 150.000000
	__ArtistShape__00000004_.Y = 380.000000
	__ArtistShape__00000004_.Width = 80.000000
	__ArtistShape__00000004_.Height = 30.000000

	__Desk__00000000_.Name = `Desk`

	__Diagram__00000000_.Name = `Modernism in Design`
	__Diagram__00000000_.IsEditable = true
	__Diagram__00000000_.IsMovementCategoryShown = true
	__Diagram__00000000_.IsArtefactTypeCategoryShown = true
	__Diagram__00000000_.IsArtistCategoryShown = true
	__Diagram__00000000_.IsInfluenceCategoryShown = true
	__Diagram__00000000_.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1900-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1965-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.RedColorCode = `#D23B22`
	__Diagram__00000000_.BackgroundGreyColorCode = `#DED6CA`

	__Influence__00000000_.Name = `Craft to Arts & Crafts`
	__Influence__00000001_.Name = `Arts & Crafts to Wiener`
	__Influence__00000002_.Name = `Ind to Werkbund`
	__Influence__00000003_.Name = `Wiener to Werkbund`
	__Influence__00000004_.Name = `Werkbund to Bauhaus`
	__Influence__00000005_.Name = `Constructivism to Bauhaus`
	__Influence__00000006_.Name = `DeStijl to Bauhaus`
	__Influence__00000007_.Name = `Bauhaus to IntStyle`
	__Influence__00000008_.Name = `Mat to Organic`
	__Influence__00000009_.Name = `Organic to Scandi`
	__Influence__00000010_.Name = `IntStyle to Scandi`

	__Movement__00000000_.Name = `ARTS & CRAFTS`
	__Movement__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1900-01-01 00:00:00 +0000 UTC")
	__Movement__00000000_.IsMajor = true

	__Movement__00000001_.Name = `WIENER WERKSTÄTTE`
	__Movement__00000001_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1903-01-01 00:00:00 +0000 UTC")

	__Movement__00000002_.Name = `DEUTSCHER WERKBUND`
	__Movement__00000002_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1907-01-01 00:00:00 +0000 UTC")

	__Movement__00000003_.Name = `CONSTRUCTIVISM`
	__Movement__00000003_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1915-01-01 00:00:00 +0000 UTC")

	__Movement__00000004_.Name = `DE STIJL`
	__Movement__00000004_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1917-01-01 00:00:00 +0000 UTC")

	__Movement__00000005_.Name = `BAUHAUS`
	__Movement__00000005_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1919-01-01 00:00:00 +0000 UTC")
	__Movement__00000005_.IsMajor = true

	__Movement__00000006_.Name = `INTERNATIONAL STYLE`
	__Movement__00000006_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1925-01-01 00:00:00 +0000 UTC")
	__Movement__00000006_.IsMajor = true

	__Movement__00000007_.Name = `ORGANIC DESIGN`
	__Movement__00000007_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1930-01-01 00:00:00 +0000 UTC")

	__Movement__00000008_.Name = `SCANDINAVIAN MODERN`
	__Movement__00000008_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1950-01-01 00:00:00 +0000 UTC")
	__Movement__00000008_.IsMajor = true

	__MovementShape__00000000_.Name = ``
	__MovementShape__00000000_.X = 100.000000
	__MovementShape__00000000_.Y = 100.000000
	__MovementShape__00000000_.Width = 100.000000
	__MovementShape__00000000_.Height = 30.000000

	__MovementShape__00000001_.Name = ``
	__MovementShape__00000001_.X = 300.000000
	__MovementShape__00000001_.Y = 130.000000
	__MovementShape__00000001_.Width = 120.000000
	__MovementShape__00000001_.Height = 30.000000

	__MovementShape__00000002_.Name = ``
	__MovementShape__00000002_.X = 450.000000
	__MovementShape__00000002_.Y = 170.000000
	__MovementShape__00000002_.Width = 140.000000
	__MovementShape__00000002_.Height = 30.000000

	__MovementShape__00000003_.Name = ``
	__MovementShape__00000003_.X = 600.000000
	__MovementShape__00000003_.Y = 250.000000
	__MovementShape__00000003_.Width = 120.000000
	__MovementShape__00000003_.Height = 30.000000

	__MovementShape__00000004_.Name = ``
	__MovementShape__00000004_.X = 500.000000
	__MovementShape__00000004_.Y = 270.000000
	__MovementShape__00000004_.Width = 80.000000
	__MovementShape__00000004_.Height = 30.000000

	__MovementShape__00000005_.Name = ``
	__MovementShape__00000005_.X = 400.000000
	__MovementShape__00000005_.Y = 290.000000
	__MovementShape__00000005_.Width = 90.000000
	__MovementShape__00000005_.Height = 30.000000

	__MovementShape__00000006_.Name = ``
	__MovementShape__00000006_.X = 450.000000
	__MovementShape__00000006_.Y = 350.000000
	__MovementShape__00000006_.Width = 150.000000
	__MovementShape__00000006_.Height = 30.000000

	__MovementShape__00000007_.Name = ``
	__MovementShape__00000007_.X = 200.000000
	__MovementShape__00000007_.Y = 400.000000
	__MovementShape__00000007_.Width = 120.000000
	__MovementShape__00000007_.Height = 30.000000

	__MovementShape__00000008_.Name = ``
	__MovementShape__00000008_.X = 350.000000
	__MovementShape__00000008_.Y = 600.000000
	__MovementShape__00000008_.Width = 160.000000
	__MovementShape__00000008_.Height = 30.000000

	// ---------------------------------------------------------
	// 3. POINTERS
	// ---------------------------------------------------------

	__ArtefactTypeShape__00000000_.ArtefactType = __ArtefactType__00000000_
	__ArtefactTypeShape__00000001_.ArtefactType = __ArtefactType__00000001_
	__ArtefactTypeShape__00000002_.ArtefactType = __ArtefactType__00000002_

	__Artist__00000000_.Place = __Place__00000000_
	__Artist__00000001_.Place = __Place__00000002_
	__Artist__00000002_.Place = __Place__00000003_
	__Artist__00000003_.Place = __Place__00000006_
	__Artist__00000004_.Place = __Place__00000007_

	__ArtistShape__00000000_.Artist = __Artist__00000000_
	__ArtistShape__00000001_.Artist = __Artist__00000001_
	__ArtistShape__00000002_.Artist = __Artist__00000002_
	__ArtistShape__00000003_.Artist = __Artist__00000003_
	__ArtistShape__00000004_.Artist = __Artist__00000004_

	__Desk__00000000_.SelectedDiagram = __Diagram__00000000_

	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000000_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000001_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __ArtefactTypeShape__00000002_)

	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000000_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000001_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000002_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000003_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtistShape__00000004_)

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

	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000000_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000001_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000002_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000003_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000004_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000005_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000006_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000007_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MovementShape__00000008_)

	__Influence__00000000_.SourceArtefactType = __ArtefactType__00000000_
	__Influence__00000000_.TargetMovement = __Movement__00000000_

	__Influence__00000001_.SourceMovement = __Movement__00000000_
	__Influence__00000001_.TargetMovement = __Movement__00000001_

	__Influence__00000002_.SourceArtefactType = __ArtefactType__00000001_
	__Influence__00000002_.TargetMovement = __Movement__00000002_

	__Influence__00000003_.SourceMovement = __Movement__00000001_
	__Influence__00000003_.TargetMovement = __Movement__00000002_

	__Influence__00000004_.SourceMovement = __Movement__00000002_
	__Influence__00000004_.TargetMovement = __Movement__00000005_

	__Influence__00000005_.SourceMovement = __Movement__00000003_
	__Influence__00000005_.TargetMovement = __Movement__00000005_

	__Influence__00000006_.SourceMovement = __Movement__00000004_
	__Influence__00000006_.TargetMovement = __Movement__00000005_

	__Influence__00000007_.SourceMovement = __Movement__00000005_
	__Influence__00000007_.TargetMovement = __Movement__00000006_

	__Influence__00000008_.SourceArtefactType = __ArtefactType__00000002_
	__Influence__00000008_.TargetMovement = __Movement__00000007_

	__Influence__00000009_.SourceMovement = __Movement__00000007_
	__Influence__00000009_.TargetMovement = __Movement__00000008_

	__Influence__00000010_.SourceMovement = __Movement__00000006_
	__Influence__00000010_.TargetMovement = __Movement__00000008_

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

	__Movement__00000000_.Places = append(__Movement__00000000_.Places, __Place__00000000_)

	__Movement__00000001_.Places = append(__Movement__00000001_.Places, __Place__00000001_)

	__Movement__00000002_.Places = append(__Movement__00000002_.Places, __Place__00000002_)

	__Movement__00000003_.Places = append(__Movement__00000003_.Places, __Place__00000005_)

	__Movement__00000004_.Places = append(__Movement__00000004_.Places, __Place__00000000_)

	__Movement__00000005_.Places = append(__Movement__00000005_.Places, __Place__00000003_)
	__Movement__00000005_.Places = append(__Movement__00000005_.Places, __Place__00000004_)

	__Movement__00000006_.Places = append(__Movement__00000006_.Places, __Place__00000006_)

	__Movement__00000007_.Places = append(__Movement__00000007_.Places, __Place__00000007_)

	__Movement__00000008_.Places = append(__Movement__00000008_.Places, __Place__00000008_)
	__Movement__00000008_.Places = append(__Movement__00000008_.Places, __Place__00000007_)

	__MovementShape__00000000_.Movement = __Movement__00000000_
	__MovementShape__00000001_.Movement = __Movement__00000001_
	__MovementShape__00000002_.Movement = __Movement__00000002_
	__MovementShape__00000003_.Movement = __Movement__00000003_
	__MovementShape__00000004_.Movement = __Movement__00000004_
	__MovementShape__00000005_.Movement = __Movement__00000005_
	__MovementShape__00000006_.Movement = __Movement__00000006_
	__MovementShape__00000007_.Movement = __Movement__00000007_
	__MovementShape__00000008_.Movement = __Movement__00000008_
}

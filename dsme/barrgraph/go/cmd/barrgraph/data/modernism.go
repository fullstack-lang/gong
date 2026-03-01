package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/barrgraph/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

var (
	_ time.Time
	_ map[string]any = map[string]any{}
)

func _(stage *models.Stage) {
	// ---------------------------------------------------------
	// 1. DECLARATIONS
	// ---------------------------------------------------------
	__Desk__00000000_ := (&models.Desk{Name: `Desk`}).Stage(stage)
	__Diagram__00000000_ := (&models.Diagram{Name: `Modernism in Design`}).Stage(stage)

	// Artefact Types
	__Artefact__Craft_ := (&models.ArtefactType{Name: `CRAFTSMANSHIP`}).Stage(stage)
	__Artefact__Industrial_ := (&models.ArtefactType{Name: `INDUSTRIAL PRODUCTION`}).Stage(stage)
	__Artefact__Materials_ := (&models.ArtefactType{Name: `NEW MATERIALS`}).Stage(stage)
	__AShape__Craft_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__AShape__Industrial_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)
	__AShape__Materials_ := (&models.ArtefactTypeShape{Name: ``}).Stage(stage)

	// Places
	__Place__London_ := (&models.Place{Name: `London`}).Stage(stage)
	__Place__Vienna_ := (&models.Place{Name: `Vienna`}).Stage(stage)
	__Place__Munich_ := (&models.Place{Name: `Munich`}).Stage(stage)
	__Place__Weimar_ := (&models.Place{Name: `Weimar`}).Stage(stage)
	__Place__Dessau_ := (&models.Place{Name: `Dessau`}).Stage(stage)
	__Place__Moscow_ := (&models.Place{Name: `Moscow`}).Stage(stage)
	__Place__Paris_ := (&models.Place{Name: `Paris`}).Stage(stage)
	__Place__Helsinki_ := (&models.Place{Name: `Helsinki`}).Stage(stage)
	__Place__Copenhagen_ := (&models.Place{Name: `Copenhagen`}).Stage(stage)

	// Artists
	__Artist__Mackintosh_ := (&models.Artist{Name: `Mackintosh`}).Stage(stage)
	__Artist__Behrens_ := (&models.Artist{Name: `Behrens`}).Stage(stage)
	__Artist__Gropius_ := (&models.Artist{Name: `Gropius`}).Stage(stage)
	__Artist__Corbusier_ := (&models.Artist{Name: `Le Corbusier`}).Stage(stage)
	__Artist__Aalto_ := (&models.Artist{Name: `Aalto`}).Stage(stage)
	__ArtShape__Mackintosh_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtShape__Behrens_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtShape__Gropius_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtShape__Corbusier_ := (&models.ArtistShape{Name: ``}).Stage(stage)
	__ArtShape__Aalto_ := (&models.ArtistShape{Name: ``}).Stage(stage)

	// Movements
	__Mov__ArtsCrafts_ := (&models.Movement{Name: `ARTS & CRAFTS`}).Stage(stage)
	__Mov__Wiener_ := (&models.Movement{Name: `WIENER WERKSTÄTTE`}).Stage(stage)
	__Mov__Werkbund_ := (&models.Movement{Name: `DEUTSCHER WERKBUND`}).Stage(stage)
	__Mov__Constructivism_ := (&models.Movement{Name: `CONSTRUCTIVISM`}).Stage(stage)
	__Mov__DeStijl_ := (&models.Movement{Name: `DE STIJL`}).Stage(stage)
	__Mov__Bauhaus_ := (&models.Movement{Name: `BAUHAUS`}).Stage(stage)
	__Mov__IntStyle_ := (&models.Movement{Name: `INTERNATIONAL STYLE`}).Stage(stage)
	__Mov__Organic_ := (&models.Movement{Name: `ORGANIC DESIGN`}).Stage(stage)
	__Mov__Scandi_ := (&models.Movement{Name: `SCANDINAVIAN MODERN`}).Stage(stage)

	__MShape__ArtsCrafts_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__Wiener_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__Werkbund_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__Constructivism_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__DeStijl_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__Bauhaus_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__IntStyle_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__Organic_ := (&models.MovementShape{Name: ``}).Stage(stage)
	__MShape__Scandi_ := (&models.MovementShape{Name: ``}).Stage(stage)

	// Influences
	__Inf__01_ := (&models.Influence{Name: `Craft to Arts & Crafts`}).Stage(stage)
	__Inf__02_ := (&models.Influence{Name: `Arts & Crafts to Wiener`}).Stage(stage)
	__Inf__03_ := (&models.Influence{Name: `Ind to Werkbund`}).Stage(stage)
	__Inf__04_ := (&models.Influence{Name: `Wiener to Werkbund`}).Stage(stage)
	__Inf__05_ := (&models.Influence{Name: `Werkbund to Bauhaus`}).Stage(stage)
	__Inf__06_ := (&models.Influence{Name: `Constructivism to Bauhaus`}).Stage(stage)
	__Inf__07_ := (&models.Influence{Name: `DeStijl to Bauhaus`}).Stage(stage)
	__Inf__08_ := (&models.Influence{Name: `Bauhaus to IntStyle`}).Stage(stage)
	__Inf__09_ := (&models.Influence{Name: `Mat to Organic`}).Stage(stage)
	__Inf__10_ := (&models.Influence{Name: `Organic to Scandi`}).Stage(stage)
	__Inf__11_ := (&models.Influence{Name: `IntStyle to Scandi`}).Stage(stage)

	__IShape__01_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__02_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__03_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__04_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__05_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__06_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__07_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__08_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__09_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__10_ := (&models.InfluenceShape{Name: ``}).Stage(stage)
	__IShape__11_ := (&models.InfluenceShape{Name: ``}).Stage(stage)

	// ---------------------------------------------------------
	// 2. INITIALIZATIONS
	// ---------------------------------------------------------

	__Diagram__00000000_.IsEditable = true
	__Diagram__00000000_.IsMovementCategoryShown = true
	__Diagram__00000000_.IsArtefactTypeCategoryShown = true
	__Diagram__00000000_.IsArtistCategoryShown = true
	__Diagram__00000000_.IsInfluenceCategoryShown = true
	__Diagram__00000000_.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1900-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1965-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.RedColorCode = `#D23B22`
	__Diagram__00000000_.BackgroundGreyColorCode = `#DED6CA`

	__AShape__Craft_.X = 100
	__AShape__Craft_.Y = 50
	__AShape__Craft_.Width = 140
	__AShape__Craft_.Height = 25
	__AShape__Industrial_.X = 550
	__AShape__Industrial_.Y = 50
	__AShape__Industrial_.Width = 180
	__AShape__Industrial_.Height = 25
	__AShape__Materials_.X = 100
	__AShape__Materials_.Y = 300
	__AShape__Materials_.Width = 140
	__AShape__Materials_.Height = 25

	__Artist__Mackintosh_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1928-01-01 00:00:00 +0000 UTC")
	__Artist__Mackintosh_.IsDead = true
	__Artist__Behrens_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1940-01-01 00:00:00 +0000 UTC")
	__Artist__Behrens_.IsDead = true
	__Artist__Gropius_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1969-01-01 00:00:00 +0000 UTC")
	__Artist__Gropius_.IsDead = true
	__Artist__Corbusier_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1965-01-01 00:00:00 +0000 UTC")
	__Artist__Corbusier_.IsDead = true
	__Artist__Aalto_.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1976-01-01 00:00:00 +0000 UTC")
	__Artist__Aalto_.IsDead = true

	__ArtShape__Mackintosh_.X = 250
	__ArtShape__Mackintosh_.Y = 100
	__ArtShape__Mackintosh_.Width = 80
	__ArtShape__Mackintosh_.Height = 30
	__ArtShape__Behrens_.X = 550
	__ArtShape__Behrens_.Y = 150
	__ArtShape__Behrens_.Width = 80
	__ArtShape__Behrens_.Height = 30
	__ArtShape__Gropius_.X = 350
	__ArtShape__Gropius_.Y = 270
	__ArtShape__Gropius_.Width = 80
	__ArtShape__Gropius_.Height = 30
	__ArtShape__Corbusier_.X = 600
	__ArtShape__Corbusier_.Y = 320
	__ArtShape__Corbusier_.Width = 80
	__ArtShape__Corbusier_.Height = 30
	__ArtShape__Aalto_.X = 150
	__ArtShape__Aalto_.Y = 380
	__ArtShape__Aalto_.Width = 80
	__ArtShape__Aalto_.Height = 30

	__Mov__ArtsCrafts_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1900-01-01 00:00:00 +0000 UTC")
	__Mov__ArtsCrafts_.IsMajor = true
	__Mov__Wiener_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1903-01-01 00:00:00 +0000 UTC")
	__Mov__Werkbund_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1907-01-01 00:00:00 +0000 UTC")
	__Mov__Constructivism_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1915-01-01 00:00:00 +0000 UTC")
	__Mov__DeStijl_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1917-01-01 00:00:00 +0000 UTC")
	__Mov__Bauhaus_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1919-01-01 00:00:00 +0000 UTC")
	__Mov__Bauhaus_.IsMajor = true
	__Mov__IntStyle_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1925-01-01 00:00:00 +0000 UTC")
	__Mov__IntStyle_.IsMajor = true
	__Mov__Organic_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1930-01-01 00:00:00 +0000 UTC")
	__Mov__Scandi_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1950-01-01 00:00:00 +0000 UTC")
	__Mov__Scandi_.IsMajor = true

	__MShape__ArtsCrafts_.X = 100
	__MShape__ArtsCrafts_.Y = 100
	__MShape__ArtsCrafts_.Width = 100
	__MShape__ArtsCrafts_.Height = 30
	__MShape__Wiener_.X = 300
	__MShape__Wiener_.Y = 130
	__MShape__Wiener_.Width = 120
	__MShape__Wiener_.Height = 30
	__MShape__Werkbund_.X = 450
	__MShape__Werkbund_.Y = 170
	__MShape__Werkbund_.Width = 140
	__MShape__Werkbund_.Height = 30
	__MShape__Constructivism_.X = 600
	__MShape__Constructivism_.Y = 250
	__MShape__Constructivism_.Width = 120
	__MShape__Constructivism_.Height = 30
	__MShape__DeStijl_.X = 500
	__MShape__DeStijl_.Y = 270
	__MShape__DeStijl_.Width = 80
	__MShape__DeStijl_.Height = 30
	__MShape__Bauhaus_.X = 400
	__MShape__Bauhaus_.Y = 290
	__MShape__Bauhaus_.Width = 90
	__MShape__Bauhaus_.Height = 30
	__MShape__IntStyle_.X = 450
	__MShape__IntStyle_.Y = 350
	__MShape__IntStyle_.Width = 150
	__MShape__IntStyle_.Height = 30
	__MShape__Organic_.X = 200
	__MShape__Organic_.Y = 400
	__MShape__Organic_.Width = 120
	__MShape__Organic_.Height = 30
	__MShape__Scandi_.X = 350
	__MShape__Scandi_.Y = 600
	__MShape__Scandi_.Width = 160
	__MShape__Scandi_.Height = 30

	// ---------------------------------------------------------
	// 3. POINTERS
	// ---------------------------------------------------------

	__Desk__00000000_.SelectedDiagram = __Diagram__00000000_

	__AShape__Craft_.ArtefactType = __Artefact__Craft_
	__AShape__Industrial_.ArtefactType = __Artefact__Industrial_
	__AShape__Materials_.ArtefactType = __Artefact__Materials_

	__Artist__Mackintosh_.Place = __Place__London_
	__Artist__Behrens_.Place = __Place__Munich_
	__Artist__Gropius_.Place = __Place__Weimar_
	__Artist__Corbusier_.Place = __Place__Paris_
	__Artist__Aalto_.Place = __Place__Helsinki_

	__ArtShape__Mackintosh_.Artist = __Artist__Mackintosh_
	__ArtShape__Behrens_.Artist = __Artist__Behrens_
	__ArtShape__Gropius_.Artist = __Artist__Gropius_
	__ArtShape__Corbusier_.Artist = __Artist__Corbusier_
	__ArtShape__Aalto_.Artist = __Artist__Aalto_

	__Mov__ArtsCrafts_.Places = append(__Mov__ArtsCrafts_.Places, __Place__London_)
	__Mov__Wiener_.Places = append(__Mov__Wiener_.Places, __Place__Vienna_)
	__Mov__Werkbund_.Places = append(__Mov__Werkbund_.Places, __Place__Munich_)
	__Mov__Constructivism_.Places = append(__Mov__Constructivism_.Places, __Place__Moscow_)
	__Mov__DeStijl_.Places = append(__Mov__DeStijl_.Places, __Place__London_)
	__Mov__Bauhaus_.Places = append(__Mov__Bauhaus_.Places, __Place__Weimar_)
	__Mov__Bauhaus_.Places = append(__Mov__Bauhaus_.Places, __Place__Dessau_)
	__Mov__IntStyle_.Places = append(__Mov__IntStyle_.Places, __Place__Paris_)
	__Mov__Organic_.Places = append(__Mov__Organic_.Places, __Place__Helsinki_)
	__Mov__Scandi_.Places = append(__Mov__Scandi_.Places, __Place__Copenhagen_)
	__Mov__Scandi_.Places = append(__Mov__Scandi_.Places, __Place__Helsinki_)

	__MShape__ArtsCrafts_.Movement = __Mov__ArtsCrafts_
	__MShape__Wiener_.Movement = __Mov__Wiener_
	__MShape__Werkbund_.Movement = __Mov__Werkbund_
	__MShape__Constructivism_.Movement = __Mov__Constructivism_
	__MShape__DeStijl_.Movement = __Mov__DeStijl_
	__MShape__Bauhaus_.Movement = __Mov__Bauhaus_
	__MShape__IntStyle_.Movement = __Mov__IntStyle_
	__MShape__Organic_.Movement = __Mov__Organic_
	__MShape__Scandi_.Movement = __Mov__Scandi_

	__Inf__01_.SourceArtefactType = __Artefact__Craft_
	__Inf__01_.TargetMovement = __Mov__ArtsCrafts_
	__Inf__02_.SourceMovement = __Mov__ArtsCrafts_
	__Inf__02_.TargetMovement = __Mov__Wiener_
	__Inf__03_.SourceArtefactType = __Artefact__Industrial_
	__Inf__03_.TargetMovement = __Mov__Werkbund_
	__Inf__04_.SourceMovement = __Mov__Wiener_
	__Inf__04_.TargetMovement = __Mov__Werkbund_
	__Inf__05_.SourceMovement = __Mov__Werkbund_
	__Inf__05_.TargetMovement = __Mov__Bauhaus_
	__Inf__06_.SourceMovement = __Mov__Constructivism_
	__Inf__06_.TargetMovement = __Mov__Bauhaus_
	__Inf__07_.SourceMovement = __Mov__DeStijl_
	__Inf__07_.TargetMovement = __Mov__Bauhaus_
	__Inf__08_.SourceMovement = __Mov__Bauhaus_
	__Inf__08_.TargetMovement = __Mov__IntStyle_
	__Inf__09_.SourceArtefactType = __Artefact__Materials_
	__Inf__09_.TargetMovement = __Mov__Organic_
	__Inf__10_.SourceMovement = __Mov__Organic_
	__Inf__10_.TargetMovement = __Mov__Scandi_
	__Inf__11_.SourceMovement = __Mov__IntStyle_
	__Inf__11_.TargetMovement = __Mov__Scandi_

	__IShape__01_.Influence = __Inf__01_
	__IShape__02_.Influence = __Inf__02_
	__IShape__03_.Influence = __Inf__03_
	__IShape__04_.Influence = __Inf__04_
	__IShape__05_.Influence = __Inf__05_
	__IShape__06_.Influence = __Inf__06_
	__IShape__07_.Influence = __Inf__07_
	__IShape__08_.Influence = __Inf__08_
	__IShape__09_.Influence = __Inf__09_
	__IShape__10_.Influence = __Inf__10_
	__IShape__11_.Influence = __Inf__11_

	// Unrolled Diagram Shape Appends
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __AShape__Craft_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __AShape__Industrial_)
	__Diagram__00000000_.ArtefactTypeShapes = append(__Diagram__00000000_.ArtefactTypeShapes, __AShape__Materials_)

	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtShape__Mackintosh_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtShape__Behrens_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtShape__Gropius_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtShape__Corbusier_)
	__Diagram__00000000_.ArtistShapes = append(__Diagram__00000000_.ArtistShapes, __ArtShape__Aalto_)

	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__ArtsCrafts_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__Wiener_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__Werkbund_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__Constructivism_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__DeStijl_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__Bauhaus_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__IntStyle_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__Organic_)
	__Diagram__00000000_.MovementShapes = append(__Diagram__00000000_.MovementShapes, __MShape__Scandi_)

	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__01_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__02_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__03_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__04_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__05_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__06_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__07_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__08_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__09_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__10_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __IShape__11_)
}

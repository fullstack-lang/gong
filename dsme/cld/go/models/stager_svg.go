package models

import (
	"fmt"
	"log"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) SvgStageUpdate() {
	stager.svgStage.Reset()

	log.Println("SVG update")

	// creates a map of art history element
	map_ArtElement_Rect := make(map[ArtElement]*svg.Rect)

	map_Artist_Influences := GetPointerReverseMap[Influence, Artist](GetAssociationName[Influence]().SourceArtist.Name, stager.stage)

	diagram := stager.desk.SelectedDiagram
	svg_ := &svg.SVG{
		Name:       diagram.Name,
		IsEditable: diagram.IsEditable,
	}

	layer := &svg.Layer{Name: "Default"}
	svg_.Layers = append(svg_.Layers, layer)

	arcLayer := &svg.Layer{Name: "Arc Layer"}
	svg_.Layers = append(svg_.Layers, arcLayer)

	y := diagram.Height + diagram.BottomBoxYOffset

	backgroundRect := &svg.Rect{
		Name:   diagram.Name,
		X:      0.0,
		Y:      0,
		Width:  diagram.BottomBoxWidth,
		Height: y,
		Presentation: svg.Presentation{
			Color:       diagram.BackgroundGreyColorCode,
			FillOpacity: 1.0,
		},
		Impl: &BackgroundRectProxy{
			stager:  stager,
			diagram: diagram,
		},
	}
	layer.Rects = append(layer.Rects, backgroundRect)

	for _, movementShape := range diagram.MovementShapes {
		movement := movementShape.Movement
		titleRectAnchoredText := &svg.RectAnchoredText{
			Name:             movement.Name,
			Content:          strings.ToUpper(movement.Name),
			RectAnchorType:   svg.RectAnchorType(diagram.MovementRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.MovementTextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.MovementDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.MovementFontWeigth,
				FontSize:      diagram.MovementFontSize,
				FontFamily:    diagram.MovementFontFamily,
				LetterSpacing: diagram.MovementLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}

		if movement.IsModern {
			titleRectAnchoredText.Content = strings.ToUpper("modern") + "\n" + titleRectAnchoredText.Content
			titleRectAnchoredText.Y_Offset = -10
			titleRectAnchoredText.X_Offset = -6
		}
		if movement.AdditionnalName != "" {
			titleRectAnchoredText.Content += " and" + "\n" + strings.ToUpper(movement.AdditionnalName)
		}
		dateRectAnchoredText := &svg.RectAnchoredText{
			Name:             movement.Name + " date",
			Content:          movement.Date.Format("2006"),
			RectAnchorType:   svg.RectAnchorType(diagram.MovementDateRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.MovementDateTextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.MovementDateTextDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.MovementDateAndPlacesFontWeigth,
				FontSize:      diagram.MovementDateAndPlacesFontSize,
				FontFamily:    diagram.MovementDateAndPlacesFontFamily,
				LetterSpacing: diagram.MovementDateAndPlacesLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}
		placesRectAnchoredText := &svg.RectAnchoredText{
			Name:             movement.Name + " places",
			RectAnchorType:   svg.RectAnchorType(diagram.MovementPlacesRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.MovementPlacesTextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.MovementPlacesDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.MovementDateAndPlacesFontWeigth,
				FontSize:      diagram.MovementDateAndPlacesFontSize,
				FontFamily:    diagram.MovementDateAndPlacesFontFamily,
				LetterSpacing: diagram.MovementDateAndPlacesLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}
		for _, place := range movement.Places {
			placesRectAnchoredText.Content += place.Name + "\n"
		}
		rect := &svg.Rect{
			Name:                movement.Name,
			X:                   movementShape.X,
			Y:                   movementShape.Y,
			Width:               movementShape.Width,
			Height:              movementShape.Height,
			CanMoveHorizontaly:  true,
			CanMoveVerticaly:    true,
			IsSelectable:        true,
			CanHaveLeftHandle:   true,
			CanHaveRightHandle:  true,
			CanHaveTopHandle:    true,
			CanHaveBottomHandle: true,
			Presentation: svg.Presentation{
				Color: svg.White.ToString(),
			},
			RectAnchoredTexts: []*svg.RectAnchoredText{
				titleRectAnchoredText,
				placesRectAnchoredText},
		}
		map_ArtElement_Rect[movement] = rect

		if movement.IsMajor {
			titleRectAnchoredText.FontSize = diagram.MajorMovementFontSize
		}
		if movement.IsMinor {
			titleRectAnchoredText.FontSize = diagram.MinorMovementFontSize
		}
		if !movement.HideDate {
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, dateRectAnchoredText)
		}
		if movement.IsModern {
			rect.Stroke = diagram.GrayColorCode
			rect.StrokeOpacity = 1.0
			rect.StrokeWidth = diagram.ArtefactTypeStrokeWidth
		}
		if movement.IsAbstract {
			abstractRectAnchoredText := &svg.RectAnchoredText{
				Name:             movement.Name,
				Content:          strings.ToUpper("(abstract)"),
				RectAnchorType:   svg.RectAnchorType(diagram.AbstractMovementRectAnchorType),
				TextAnchorType:   svg.TextAnchorType(diagram.AbstractMovementTextAnchorType),
				DominantBaseline: svg.DominantBaselineType(diagram.AbstractDominantBaselineType),
				TextAttributes: svg.TextAttributes{
					FontWeight:    diagram.MovementFontWeigth,
					FontSize:      diagram.AbstractMovementFontSize,
					FontFamily:    diagram.MovementFontFamily,
					LetterSpacing: diagram.MovementLetterSpacing,
				},
				Presentation: svg.Presentation{
					Color:       diagram.GrayColorCode,
					FillOpacity: 1.0,
				},
			}
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, abstractRectAnchoredText)
		}

		rect.Impl = &MovementShapeProxy{
			stager: stager,
			shape:  movementShape,
		}
		if diagram.IsMovementCategoryShown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, artefactTypeShape := range diagram.ArtefactTypeShapes {
		artefactType := artefactTypeShape.ArtefactType
		titleRectAnchoredText := &svg.RectAnchoredText{
			Name:             artefactType.Name,
			Content:          strings.ToUpper(artefactType.Name),
			RectAnchorType:   svg.RectAnchorType(diagram.ArtefactTypeRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(TEXT_ANCHOR_CENTER),
			DominantBaseline: svg.DominantBaselineType(diagram.ArtefactDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.ArtefactTypeFontWeigth,
				FontSize:      diagram.ArtefactTypeFontSize,
				FontFamily:    diagram.ArtefactTypeFontFamily,
				LetterSpacing: diagram.ArtefactTypeLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.RedColorCode,
				FillOpacity: 1.0,
			},
		}

		rect := &svg.Rect{
			Name:                artefactType.Name,
			X:                   artefactTypeShape.X,
			Y:                   artefactTypeShape.Y,
			Width:               artefactTypeShape.Width,
			Height:              artefactTypeShape.Height,
			CanMoveHorizontaly:  true,
			CanMoveVerticaly:    true,
			IsSelectable:        true,
			CanHaveLeftHandle:   true,
			CanHaveRightHandle:  true,
			CanHaveTopHandle:    true,
			CanHaveBottomHandle: true,
			Presentation: svg.Presentation{
				Color:         svg.White.ToString(),
				Stroke:        diagram.RedColorCode,
				StrokeOpacity: 1.0,
				StrokeWidth:   diagram.ArtefactTypeStrokeWidth,
			},
			RectAnchoredTexts: []*svg.RectAnchoredText{
				titleRectAnchoredText},
		}
		map_ArtElement_Rect[artefactType] = rect

		rect.Impl = &ArtefactTypeShapeProxy{
			stager: stager,
			shape:  artefactTypeShape,
		}
		if diagram.IsArtefactTypeCategoryShown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, artistShape := range diagram.ArtistShapes {
		artist := artistShape.Artist
		titleRectAnchoredText := &svg.RectAnchoredText{
			Name:             artist.Name,
			Content:          artist.Name,
			RectAnchorType:   svg.RectAnchorType(diagram.ArtistRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.ArtistTextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.ArtistDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.ArtistFontWeigth,
				FontSize:      diagram.ArtistFontSize,
				FontFamily:    diagram.ArtistFontFamily,
				LetterSpacing: diagram.ArtistLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}

		dateRectAnchoredText := &svg.RectAnchoredText{
			Name:             artist.Name + " date",
			Content:          "d. " + artist.DateOfDeath.Format("2006"),
			RectAnchorType:   svg.RectAnchorType(diagram.ArtistDateRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.ArtistDateTextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.ArtefactDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.ArtistDateAndPlacesFontWeigth,
				FontSize:      diagram.ArtistDateAndPlacesFontSize,
				FontFamily:    diagram.ArtistDateAndPlacesFontFamily,
				LetterSpacing: diagram.ArtistDateAndPlacesLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}
		placesRectAnchoredText := &svg.RectAnchoredText{
			Name:             artist.Name + " places",
			RectAnchorType:   svg.RectAnchorType(diagram.ArtistPlacesRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.ArtistPlacesTextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.ArtistPlacesDominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.ArtistDateAndPlacesFontWeigth,
				FontSize:      diagram.ArtistDateAndPlacesFontSize,
				FontFamily:    diagram.ArtistDateAndPlacesFontFamily,
				LetterSpacing: diagram.ArtistDateAndPlacesLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}
		if artist.Place != nil {
			placesRectAnchoredText.Content += artist.Place.Name
		}

		x := artistShape.X + artistShape.Width/2.0
		y := artistShape.Y + artistShape.Height +
			diagram.MovementBelowArcY_Offset +
			diagram.MovementBelowArcY_OffsetPerPlace
		path := &svg.Path{
			Name: artist.Name,
			Definition: fmt.Sprintf("M %f %f A %f %f 0 0 0 %f %f",
				// move to
				x-artistShape.Width/2.0,
				y,

				// rx, ry
				artistShape.Width/2.0,
				artistShape.Height/2.0,

				// end of the stroke
				x+artistShape.Width/2.0,
				y,
			),
			Presentation: svg.Presentation{
				Stroke:        diagram.GrayColorCode,
				StrokeOpacity: 1.0,
				StrokeWidth:   diagram.ArtefactTypeStrokeWidth,
				Color:         diagram.BackgroundGreyColorCode,
				FillOpacity:   1.0,
			},
		}

		rect := &svg.Rect{
			Name:                artist.Name,
			X:                   artistShape.X,
			Y:                   artistShape.Y,
			Width:               artistShape.Width,
			Height:              artistShape.Height,
			CanMoveHorizontaly:  true,
			CanMoveVerticaly:    true,
			IsSelectable:        true,
			CanHaveLeftHandle:   true,
			CanHaveRightHandle:  true,
			CanHaveTopHandle:    true,
			CanHaveBottomHandle: true,
			Presentation: svg.Presentation{
				Color: svg.White.ToString(),
			},
			RectAnchoredTexts: []*svg.RectAnchoredText{
				titleRectAnchoredText,
				placesRectAnchoredText},
		}
		map_ArtElement_Rect[artist] = rect

		if artist.IsDead {
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, dateRectAnchoredText)
		}

		rect.Impl = &ArtistShapeProxy{
			stager: stager,
			shape:  artistShape,
		}

		if diagram.IsInfluenceCategoryShown {
			// only artists with influence have a bottom arc
			if len(map_Artist_Influences[artist]) > 0 {
				arcLayer.Paths = append(arcLayer.Paths, path)
			}
		}
		if diagram.IsArtistCategoryShown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, influenceShape := range diagram.InfluenceShapes {

		influence := influenceShape.Influence

		link := new(svg.Link)
		link.Name = influenceShape.Name

		link.Start = map_ArtElement_Rect[influence.source]

		link.StartArrowOffset = diagram.InfluenceArrowStartOffset

		link.End = map_ArtElement_Rect[influence.target]
		link.HasEndArrow = true
		link.EndArrowSize = diagram.InfluenceArrowSize

		link.EndArrowOffset = diagram.InfluenceArrowEndOffset

		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER

		link.CornerRadius = 25

		link.Stroke = diagram.GrayColorCode

		// exception for artefact types
		if influence.SourceArtefactType != nil {
			link.Stroke = diagram.RedColorCode
			link.StartArrowOffset = 0.0
		}

		if influence.TargetMovement != nil && influence.TargetMovement.IsModern {
			if influence.SourceArtefactType == nil {
				link.EndArrowOffset = 0.0
			} else {
				link.EndArrowOffset *= 2.0
			}
		}

		if influence.SourceMovement != nil {
			link.StartArrowOffset += float64(len(influence.SourceMovement.Places) *
				int(diagram.MovementBelowArcY_OffsetPerPlace))
		}

		if influence.TargetMovement != nil && influence.TargetMovement.IsAbstract {
			link.EndArrowOffset *= 2.0
		}

		link.StrokeWidth = diagram.ArtefactTypeStrokeWidth
		link.StrokeOpacity = 1

		if influence.IsHypothtical {
			link.StrokeDashArray = diagram.InfluenceDashedLinePattern
		}

		for _, controlPointShape := range influenceShape.ControlPointShapes {

			closestRect := link.Start
			if !controlPointShape.IsStartShapeTheClosestShape {
				closestRect = link.End
			}

			link.ControlPoints = append(link.ControlPoints, &svg.ControlPoint{
				Name:        controlPointShape.GetName(),
				X_Relative:  controlPointShape.X_Relative,
				Y_Relative:  controlPointShape.Y_Relative,
				ClosestRect: closestRect,
				Impl: &ControlPointShapeProxy{
					stager:            stager,
					influenceShape:    influenceShape,
					controlPointShape: controlPointShape,
				},
			})
		}
		link.CornerRadius = diagram.InfluenceCornerRadius

		// callback
		link.Impl = &InfluenceShapeProxy{
			influenceShape: influenceShape,
			stager:         stager,
		}

		if diagram.IsInfluenceCategoryShown {
			layer.Links = append(layer.Links, link)
		}
	}

	svg.StageBranch(stager.svgStage, svg_)

	stager.svgStage.Commit()
}

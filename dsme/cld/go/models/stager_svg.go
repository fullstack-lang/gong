package models

import (
	"log"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) SvgStageUpdate() {
	stager.svgStage.Reset()

	log.Println("SVG update")

	// creates a map of art history element
	map_ArtElement_Rect := make(map[Category]*svg.Rect)

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

	for _, movementShape := range diagram.Category1Shapes {
		movement := movementShape.Category1
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
				titleRectAnchoredText},
		}
		map_ArtElement_Rect[movement] = rect

		rect.Impl = &Category1ShapeProxy{
			stager: stager,
			shape:  movementShape,
		}
		if diagram.IsCategory1Shown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, artefactTypeShape := range diagram.Category3Shapes {
		artefactType := artefactTypeShape.Category3
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

		rect.Impl = &Category3ShapeProxy{
			stager: stager,
			shape:  artefactTypeShape,
		}
		if diagram.IsCategory2Shown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, artistShape := range diagram.Category2Shapes {
		artist := artistShape.Category2
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
				titleRectAnchoredText},
		}
		map_ArtElement_Rect[artist] = rect

		rect.Impl = &Category2Proxy{
			stager: stager,
			shape:  artistShape,
		}

		if diagram.IsCategory3Shown {
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
		if influence.SourceCategory2 != nil {
			link.Stroke = diagram.RedColorCode
			link.StartArrowOffset = 0.0
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

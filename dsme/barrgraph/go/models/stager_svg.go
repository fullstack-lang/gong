package models

import (
	"log"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/pkg/strutils"
)

func (stager *Stager) SvgStageUpdate() {
	stager.svgStage.Reset()

	log.Println("SVG update")

	// creates a map of art history element
	map_Category_Rect := make(map[Category]*svg.Rect)

	diagram := stager.desk.SelectedDiagram
	svg_ := &svg.SVG{
		Name:       diagram.Name,
		IsEditable: diagram.IsEditable,
	}

	layer := &svg.Layer{Name: "Default"}
	svg_.Layers = append(svg_.Layers, layer)

	y := diagram.Height

	backgroundRect := &svg.Rect{
		Name:   diagram.Name,
		X:      0.0,
		Y:      0,
		Width:  diagram.Width,
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

	for _, s := range diagram.Category1Shapes {
		category := s.Category1
		titleRectAnchoredText := &svg.RectAnchoredText{
			Name:             category.Name,
			Content:          strutils.WrapString(strings.ToUpper(category.Name), int(s.Width/diagram.NbPixPerCharacter)),
			RectAnchorType:   svg.RectAnchorType(diagram.Category1RectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.Category1TextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.Category1DominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.Category1FontWeigth,
				FontSize:      diagram.Category1FontSize,
				FontFamily:    diagram.Category1FontFamily,
				LetterSpacing: diagram.Category1LetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}

		rect := &svg.Rect{
			Name:                category.Name,
			X:                   s.X,
			Y:                   s.Y,
			Width:               s.Width,
			Height:              s.Height,
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
		map_Category_Rect[category] = rect

		rect.Impl = &Category1ShapeProxy{
			stager: stager,
			shape:  s,
		}
		if diagram.IsCategory1Shown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, s := range diagram.Category3Shapes {
		category := s.Category3
		titleRectAnchoredText := &svg.RectAnchoredText{
			Name:             category.Name,
			Content:          strutils.WrapString(strings.ToUpper(category.Name), int(s.Width/diagram.NbPixPerCharacter)),
			RectAnchorType:   svg.RectAnchorType(diagram.Category2TypeRectAnchorType),
			TextAnchorType:   svg.TextAnchorType(TEXT_ANCHOR_CENTER),
			DominantBaseline: svg.DominantBaselineType(diagram.Category2DominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.Category2TypeFontWeigth,
				FontSize:      diagram.Category2TypeFontSize,
				FontFamily:    diagram.Category2TypeFontFamily,
				LetterSpacing: diagram.Category2TypeLetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.RedColorCode,
				FillOpacity: 1.0,
			},
		}

		rect := &svg.Rect{
			Name:                category.Name,
			X:                   s.X,
			Y:                   s.Y,
			Width:               s.Width,
			Height:              s.Height,
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
				StrokeWidth:   diagram.Category2StrokeWidth,
			},
			RectAnchoredTexts: []*svg.RectAnchoredText{
				titleRectAnchoredText},
		}
		map_Category_Rect[category] = rect

		rect.Impl = &Category3ShapeProxy{
			stager: stager,
			shape:  s,
		}
		if diagram.IsCategory2Shown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, s := range diagram.Category2Shapes {
		category := s.Category2
		titleRectAnchoredText := &svg.RectAnchoredText{
			Name:             category.Name,
			Content:          strutils.WrapString(strings.ToUpper(category.Name), int(s.Width/diagram.NbPixPerCharacter)),
			RectAnchorType:   svg.RectAnchorType(diagram.Category3RectAnchorType),
			TextAnchorType:   svg.TextAnchorType(diagram.Category3TextAnchorType),
			DominantBaseline: svg.DominantBaselineType(diagram.Category3DominantBaselineType),
			TextAttributes: svg.TextAttributes{
				FontWeight:    diagram.Category3FontWeigth,
				FontSize:      diagram.Category3FontSize,
				FontFamily:    diagram.Category3FontFamily,
				LetterSpacing: diagram.Category3LetterSpacing,
			},
			Presentation: svg.Presentation{
				Color:       diagram.GrayColorCode,
				FillOpacity: 1.0,
			},
		}
		rect := &svg.Rect{
			Name:                category.Name,
			X:                   s.X,
			Y:                   s.Y,
			Width:               s.Width,
			Height:              s.Height,
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
		map_Category_Rect[category] = rect

		rect.Impl = &Category2Proxy{
			stager: stager,
			shape:  s,
		}

		if diagram.IsCategory3Shown {
			layer.Rects = append(layer.Rects, rect)
		}
	}

	for _, influenceShape := range diagram.InfluenceShapes {

		influence := influenceShape.Influence

		link := new(svg.Link)
		link.Name = influenceShape.Name

		link.Start = map_Category_Rect[influence.source]

		link.StartArrowOffset = diagram.InfluenceArrowStartOffset

		link.End = map_Category_Rect[influence.target]
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

		link.StrokeWidth = diagram.InfluenceStrokeWidth
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

		if text := influence.TextAtEndOfArrow; text != "" {
			textDisplayed := &svg.LinkAnchoredText{
				Name:            text,
				Content:         text,
				LinkAnchorType:  svg.LINK_LEFT_OR_TOP,
				AutomaticLayout: true,
				Presentation: svg.Presentation{
					Stroke:        diagram.GrayColorCode,
					StrokeOpacity: 1.0,
					StrokeWidth:   diagram.InfluenceStrokeWidth,
					Color:         diagram.GrayColorCode,
					FillOpacity:   1.0,
				},
				TextAttributes: svg.TextAttributes{
					FontWeight:    diagram.InfluenceFontWeigth,
					FontSize:      diagram.InfluenceFontSize,
					FontFamily:    diagram.InfluenceFontFamily,
					LetterSpacing: diagram.InfluenceLetterSpacing,
				},
			}
			link.TextAtArrowEnd = append(link.TextAtArrowEnd, textDisplayed)
		}
	}

	svg.StageBranch(stager.svgStage, svg_)

	stager.svgStage.Commit()
}

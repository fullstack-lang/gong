package models

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

const SVGName string = `SVG`

const ClassBoxStrokeWidth = 3

func (stager *Stager) UpdateAndCommitSVGStage() {

	stager.svgStage.Reset()
	// log.Println("stager.GenerateSvg")

	stager.map_GongstructShape_Rect = make(map[*GongStructShape]*svg_models.Rect)
	stager.map_GongenumShape_Rect = make(map[*GongEnumShape]*svg_models.Rect)
	stager.map_NoteShape_Rect = make(map[*GongNoteShape]*svg_models.Rect)

	stager.map_Structname_Rect = make(map[string]*svg_models.Rect)
	stager.map_Fieldname_Link = make(map[string]*svg_models.Link)

	stager.svgStage.Reset()

	var classdiagram *Classdiagram

	var diagramPackage *DiagramPackage
	for diagramPackage = range *GetGongstructInstancesSet[DiagramPackage](stager.stage) {

		classdiagram = diagramPackage.SelectedClassdiagram

		// if no class diagram is selected generate a blank diagram
		if classdiagram == nil {
			classdiagram = new(Classdiagram)
		}
	}
	if diagramPackage == nil {
		return
	}

	svg := new(svg_models.SVG)
	svg.Name = classdiagram.Name
	svg.IsEditable = !stager.embeddedDiagrams

	for _, gongstructShape := range classdiagram.GongStructShapes {

		rectLayer := new(svg_models.Layer)

		gongStructIdentifier := IdentifierMetaToGongStructName(gongstructShape.IdentifierMeta)

		rectLayer.Name = "Layer" + gongStructIdentifier
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(svg_models.Rect)
		rect.Name = gongStructIdentifier

		// hook a callback on rect modifications
		rect.Impl = NewRectImplGongstructShape(gongstructShape, stager.stage)

		stager.map_GongstructShape_Rect[gongstructShape] = rect
		stager.map_Structname_Rect[gongStructIdentifier] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = gongstructShape.X
		rect.Y = gongstructShape.Y

		rect.Width = gongstructShape.Width
		rect.Height = gongstructShape.Height

		rect.Stroke = svg_models.Lightsalmon.ToString()
		rect.Stroke = svg_models.Lightgrey.ToString()
		rect.StrokeOpacity = 1
		rect.StrokeWidth = ClassBoxStrokeWidth
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = svg_models.Lightsalmon.ToString()
		rect.Color = "white"
		rect.RX = 8

		// moveability
		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true

		// resizing
		rect.IsSelectable = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveTopHandle = true
		rect.CanHaveRightHandle = true

		//
		// Title
		//
		title := new(svg_models.RectAnchoredText)
		title.Name = gongStructIdentifier
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = svg_models.RECT_TOP
		title.TextAnchorType = svg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.Color = svg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		// additional box to hightlight the title
		titleBox := new(svg_models.RectAnchoredRect)
		titleBox.Name = gongStructIdentifier
		titleBox.X_Offset = rect.StrokeWidth
		titleBox.Y_Offset = rect.StrokeWidth
		titleBox.Width = rect.Width - 2*rect.StrokeWidth - 30
		titleBox.Height = 30
		titleBox.StrokeOpacity = 0
		titleBox.RectAnchorType = svg_models.RECT_TOP_LEFT
		titleBox.Color = "#ff8450"
		titleBox.Color = svg_models.White.ToString()
		titleBox.WidthFollowRect = true
		titleBox.FillOpacity = 0

		titleBox.Stroke = svg_models.White.ToString()
		titleBox.StrokeOpacity = 1
		titleBox.StrokeWidth = 0

		rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)

		//
		// fields
		//
		for idx, field := range gongstructShape.AttributeShapes {
			fieldText := new(svg_models.RectAnchoredText)
			fieldText.Name = IdentifierMetaToFieldName(field.IdentifierMeta) + " : " + field.Fieldtypename
			fieldText.Content = IdentifierMetaToFieldName(field.IdentifierMeta) + " : " + field.Fieldtypename

			// field position
			fieldText.X_Offset = 10
			fieldText.Y_Offset = 20 + 30 + float64(idx)*HeightBetween2AttributeShapes
			fieldText.RectAnchorType = svg_models.RECT_TOP_LEFT
			fieldText.TextAnchorType = svg_models.TEXT_ANCHOR_START

			fieldText.Color = "black"
			fieldText.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, fieldText)
		}

		//
		// number of instance (x%d)
		//
		if classdiagram.ShowNbInstances {

			if nbInstance, ok := stager.map_GongStructName_InstancesNb[gongStructIdentifier]; ok {

				nbInstancesText := new(svg_models.RectAnchoredText)
				nbInstancesText.Name = fmt.Sprintf("(%d)", nbInstance)
				nbInstancesText.Content = fmt.Sprintf("(%d)", nbInstance)

				// text position
				nbInstancesText.X_Offset = -5 - 2*rect.StrokeWidth
				nbInstancesText.Y_Offset = 20
				nbInstancesText.RectAnchorType = svg_models.RECT_TOP_RIGHT
				nbInstancesText.TextAnchorType = svg_models.TEXT_ANCHOR_END

				nbInstancesText.Color = "black"
				nbInstancesText.FillOpacity = 1.0
				rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, nbInstancesText)
			}
		}
	}

	// display links between gongstruct shapes
	for _, gongstructShape := range classdiagram.GongStructShapes {

		startRect := stager.map_GongstructShape_Rect[gongstructShape]
		for _, linkShape := range gongstructShape.LinkShapes {
			var ok bool
			var fieldMetaIdentifierString string
			if fieldMetaIdentifierString, ok = linkShape.IdentifierMeta.(string); !ok {
				// some legacy diagram have not the meta identifier
				// therefore, we continue
				log.Println("not a identifier meta as string")
				continue
			}

			endRectGongStructName := IdentifierMetaToGongStructName(linkShape.FieldTypeIdentifierMeta)
			endRect, ok := stager.map_Structname_Rect[endRectGongStructName]

			// if some renaming of field type name has occured, end rect might be nil
			if !ok {
				continue
			}

			link := new(svg_models.Link)
			link.Name = startRect.Name + " - to - " + endRect.Name

			link.Impl = NewLinkImplLink(linkShape, stager.stage)

			linkLayer := new(svg_models.Layer)

			stager.map_Fieldname_Link[fieldMetaIdentifierString] = link

			linkLayer.Links = append(linkLayer.Links, link)
			svg.Layers = append(svg.Layers, linkLayer)

			// configuration
			link.Stroke = svg_models.Slategray.ToString()
			link.StrokeOpacity = 1
			link.StrokeWidth = 3
			link.HasEndArrow = true
			link.EndArrowSize = 8
			link.Type = svg_models.LINK_TYPE_FLOATING_ORTHOGONAL

			link.StartOrientation = svg_models.OrientationType(linkShape.StartOrientation)

			// take into accound legacy links
			if link.StartOrientation == "" {
				link.StartOrientation = svg_models.ORIENTATION_HORIZONTAL
			}

			link.StartRatio = linkShape.StartRatio
			link.EndOrientation = svg_models.OrientationType(linkShape.EndOrientation)
			// take into accound legacy links
			if link.EndOrientation == "" {
				link.EndOrientation = svg_models.ORIENTATION_HORIZONTAL
			}

			link.EndRatio = linkShape.EndRatio

			link.CornerOffsetRatio = linkShape.CornerOffsetRatio

			link.CornerRadius = 8

			link.Start = startRect
			link.End = endRect

			// add text to the arrow
			if classdiagram.ShowMultiplicity {
				targetMulitplicity := new(svg_models.LinkAnchoredText)
				targetMulitplicity.AutomaticLayout = true
				targetMulitplicity.LinkAnchorType = svg_models.LINK_RIGHT_OR_BOTTOM

				targetMulitplicity.Impl = NewAnchoredTextImplLinkTargetMultiplicity(linkShape, stager.stage)
				link.TextAtArrowEnd = append(link.TextAtArrowEnd, targetMulitplicity)
				targetMulitplicity.Name = linkShape.TargetMultiplicity.ToString()
				targetMulitplicity.Content = targetMulitplicity.Name
				targetMulitplicity.X_Offset = linkShape.TargetMultiplicityOffsetX
				targetMulitplicity.Y_Offset = linkShape.TargetMultiplicityOffsetY
				targetMulitplicity.Stroke = svg_models.Black.ToString()
				targetMulitplicity.StrokeOpacity = 1
				targetMulitplicity.StrokeWidth = 1
				targetMulitplicity.Color = svg_models.Black.ToString()
				targetMulitplicity.FillOpacity = 100
				targetMulitplicity.FontWeight = "300"
				targetMulitplicity.FontSize = "15"
				targetMulitplicity.LetterSpacing = "0.1em"
			}

			if classdiagram.ShowLinkNames {
				fieldName := new(svg_models.LinkAnchoredText)
				fieldName.AutomaticLayout = true
				fieldName.LinkAnchorType = svg_models.LINK_LEFT_OR_TOP

				fieldName.Impl = NewAnchoredTextImplLinkFieldName(linkShape, stager.stage)

				link.TextAtArrowEnd = append(link.TextAtArrowEnd, fieldName)
				fieldName.Name = linkShape.GetName()
				fieldName.Content = IdentifierMetaToFieldName(linkShape.IdentifierMeta)
				fieldName.Y_Offset = linkShape.FieldOffsetY
				fieldName.X_Offset = linkShape.FieldOffsetX
				fieldName.Stroke = svg_models.Black.ToString()
				fieldName.StrokeOpacity = 1
				fieldName.StrokeWidth = 1
				fieldName.Color = svg_models.Black.ToString()
				fieldName.FillOpacity = 100
				fieldName.FontWeight = "300"
				fieldName.FontSize = "15"
				fieldName.LetterSpacing = "0.1em"
			}

			// add the callback
			if classdiagram.ShowMultiplicity {
				sourceMultiplicity := new(svg_models.LinkAnchoredText)
				sourceMultiplicity.AutomaticLayout = true
				sourceMultiplicity.LinkAnchorType = svg_models.LINK_RIGHT_OR_BOTTOM

				sourceMultiplicity.Impl = NewAnchoredTextImplLinkSourceMultiplicity(linkShape, stager.stage)

				link.TextAtArrowStart = append(link.TextAtArrowStart, sourceMultiplicity)
				sourceMultiplicity.Name = linkShape.SourceMultiplicity.ToString()
				sourceMultiplicity.Content = sourceMultiplicity.Name
				sourceMultiplicity.X_Offset = linkShape.SourceMultiplicityOffsetX
				sourceMultiplicity.Y_Offset = linkShape.SourceMultiplicityOffsetY
				sourceMultiplicity.Stroke = svg_models.Black.ToString()
				sourceMultiplicity.StrokeOpacity = 1
				sourceMultiplicity.StrokeWidth = 1
				sourceMultiplicity.FontWeight = "300"
				sourceMultiplicity.FontSize = "15"
				sourceMultiplicity.LetterSpacing = "0.1em"
			}
		}
	}

	//
	// GongEnumShapes
	//
	for _, gongenumShape := range classdiagram.GongEnumShapes {

		rectLayer := new(svg_models.Layer)
		rectLayer.Name = "Layer" + GongEnumIdentifierMetaToGongEnumName(gongenumShape.IdentifierMeta)
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(svg_models.Rect)
		rect.Name = GongEnumIdentifierMetaToGongEnumName(gongenumShape.IdentifierMeta)

		rect.Impl = NewRectImplGongenumShape(gongenumShape, stager.stage)

		stager.map_GongenumShape_Rect[gongenumShape] = rect
		stager.map_Structname_Rect[GongEnumIdentifierMetaToGongEnumName(gongenumShape.IdentifierMeta)] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = gongenumShape.X
		rect.Y = gongenumShape.Y

		rect.Width = gongenumShape.Width
		rect.Height = gongenumShape.Height

		rect.Stroke = svg_models.Lightsteelblue.ToString()
		rect.StrokeOpacity = 1
		rect.StrokeWidth = ClassBoxStrokeWidth
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = svg_models.White.ToString()

		// moveability
		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true

		// resizing
		rect.IsSelectable = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveTopHandle = true
		rect.CanHaveRightHandle = true

		//
		// Title
		//
		title := new(svg_models.RectAnchoredText)
		title.Name = IdentifierToGongStructName(GongEnumIdentifierMetaToGongEnumName(gongenumShape.IdentifierMeta))
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = svg_models.RECT_TOP
		title.TextAnchorType = svg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.Color = svg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		// additional box to hightlight the title
		titleBox := new(svg_models.RectAnchoredRect)
		titleBox.Name = IdentifierToGongStructName(GongEnumIdentifierMetaToGongEnumName(gongenumShape.IdentifierMeta))
		titleBox.X_Offset = 0
		titleBox.Y_Offset = 0
		titleBox.Width = rect.Width
		titleBox.Height = 30
		titleBox.RectAnchorType = svg_models.RECT_TOP_LEFT
		titleBox.Color = svg_models.White.ToString()
		titleBox.WidthFollowRect = true
		titleBox.FillOpacity = 0.0
		titleBox.StrokeWidth = 0

		rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)

		//
		// fields
		//
		for idx, gongEnumValueShape := range gongenumShape.GongEnumValueShapes {
			fieldText := new(svg_models.RectAnchoredText)
			fieldText.Name = GongEnumValueShapeIdentifierMetaToValueName(gongEnumValueShape.IdentifierMeta)
			fieldText.Content = GongEnumValueShapeIdentifierMetaToValueName(gongEnumValueShape.IdentifierMeta)

			// field position
			fieldText.X_Offset = 10
			fieldText.Y_Offset = 20 + 30 + float64(idx)*HeightBetween2AttributeShapes
			fieldText.RectAnchorType = svg_models.RECT_TOP_LEFT
			fieldText.TextAnchorType = svg_models.TEXT_ANCHOR_START

			fieldText.Color = "black"
			fieldText.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, fieldText)
		}
	}

	//
	// Notes
	//
	for _, noteShape := range classdiagram.GongNoteShapes {

		rectLayer := new(svg_models.Layer)
		rectLayer.Name = "Layer" + noteShape.Identifier
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(svg_models.Rect)
		rect.Name = noteShape.Identifier

		rect.Impl = NewRectImplNoteShape(noteShape, stager.stage)

		stager.map_NoteShape_Rect[noteShape] = rect
		stager.map_Structname_Rect[noteShape.Identifier] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = noteShape.X
		rect.Y = noteShape.Y

		rect.Width = noteShape.Width
		rect.Height = noteShape.Height

		rect.Stroke = svg_models.Lightskyblue.ToString()
		rect.StrokeOpacity = 1
		rect.StrokeWidth = ClassBoxStrokeWidth
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = svg_models.White.ToString()

		// moveability
		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true

		// resizing
		rect.IsSelectable = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveTopHandle = true
		rect.CanHaveRightHandle = true

		// Split the body into title and the rest of the content
		bodyLines := strings.SplitN(noteShape.Body, "\n", 2)
		noteTitle := ""
		noteContent := ""
		if len(bodyLines) > 0 {
			noteTitle = bodyLines[0]
		}
		if len(bodyLines) > 1 {
			// Trim leading/trailing whitespace and newlines from the rest of the content
			noteContent = strings.TrimSpace(bodyLines[1])

			// remove [models. prefix and ] suffix
			re := regexp.MustCompile(`\[models\.([^\]]+)\]`)
			noteContent = re.ReplaceAllString(noteContent, "$1")
		}

		//
		// Title
		//
		title := new(svg_models.RectAnchoredText)
		title.Name = noteTitle
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = svg_models.RECT_TOP
		title.TextAnchorType = svg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.FontStyle = "oblique"
		title.Color = svg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		//
		// Content
		//
		content := new(svg_models.RectAnchoredText)
		content.Name = noteContent
		content.Content = content.Name
		content.X_Offset = 0
		content.Y_Offset = 40
		content.RectAnchorType = svg_models.RECT_TOP
		content.TextAnchorType = svg_models.TEXT_ANCHOR_CENTER
		content.FontWeight = "normal"
		content.Color = svg_models.Black.ToString()
		content.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, content)

		// // additional box to hightlight the title
		// titleBox := new(gongsvg_models.RectAnchoredRect).Stage(stager.gongsvgStage)
		// titleBox.Name = IdentifierToGongObjectName(noteShape.Identifier)
		// titleBox.X_Offset = 0
		// titleBox.Y_Offset = 0
		// titleBox.Width = rect.Width
		// titleBox.Height = 30
		// titleBox.RectAnchorType = gongsvg_models.RECT_TOP_LEFT
		// titleBox.Color = gongsvg_models.Skyblue.ToString()
		// titleBox.WidthFollowRect = true
		// titleBox.FillOpacity = 100

		// rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)
	}

	//
	// Links between notes and othe shapes
	//
	for _, noteShape := range classdiagram.GongNoteShapes {

		startRect := stager.map_NoteShape_Rect[noteShape]
		_ = startRect

		for _, noteLink := range noteShape.GongNoteLinkShapes {

			if noteLink.Type == NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE {

				var endRect *svg_models.Rect
				var ok bool
				// find the endRect
				endRect, ok = stager.map_Structname_Rect[noteLink.Identifier]
				if ok {
					// create the link
					link := new(svg_models.Link)
					link.Name = startRect.Name + " - to - " + endRect.Name
					linkLayer := new(svg_models.Layer)

					linkLayer.Links = append(linkLayer.Links, link)
					svg.Layers = append(svg.Layers, linkLayer)

					// configuration
					link.Stroke = svg_models.Slategray.ToString()
					link.StrokeOpacity = 1
					link.StrokeWidth = 2
					link.StrokeDashArray = "2 2"

					// first version
					link.Type = svg_models.LINK_TYPE_FLOATING_ORTHOGONAL
					link.StartOrientation = svg_models.ORIENTATION_HORIZONTAL
					link.StartRatio = 0.5
					link.EndOrientation = svg_models.ORIENTATION_HORIZONTAL
					link.EndRatio = 0.5

					// more elegent
					link.Type = svg_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
					link.StartAnchorType = svg_models.ANCHOR_CENTER
					link.EndAnchorType = svg_models.ANCHOR_CENTER

					link.CornerOffsetRatio = (endRect.X - startRect.X - 80) / startRect.Width

					link.CornerRadius = 3

					link.Start = startRect
					link.End = endRect
				}

			}
			if noteLink.Type == NOTE_SHAPE_LINK_TO_GONG_FIELD {

				var endLink *svg_models.Link
				_ = endLink
				var ok bool
				// find the endLink
				endLink, ok = stager.map_Fieldname_Link[noteLink.Identifier]
				if ok {
					rectLinkLink := new(svg_models.RectLinkLink)
					rectLinkLink.Name = startRect.Name + " - to - " + endLink.Name
					rectLinkLinkLayer := new(svg_models.Layer)

					rectLinkLinkLayer.RectLinkLinks = append(rectLinkLinkLayer.RectLinkLinks, rectLinkLink)
					svg.Layers = append(svg.Layers, rectLinkLinkLayer)

					// configuration
					rectLinkLink.Stroke = svg_models.Slategray.ToString()
					rectLinkLink.StrokeOpacity = 1
					rectLinkLink.StrokeWidth = 2
					rectLinkLink.StrokeDashArray = "2 2"
					rectLinkLink.TargetAnchorPosition = 0.5

					rectLinkLink.Start = startRect
					rectLinkLink.End = endLink
				}
			}

		}

	}

	svg_models.StageBranch(stager.svgStage, svg)

	stager.svgStage.Commit()

}

package doc2svg

import (
	"fmt"

	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type DocSVGMapper struct {
	map_GongstructShape_Rect map[*gongdoc_models.GongStructShape]*gongsvg_models.Rect
	map_GongenumShape_Rect   map[*gongdoc_models.GongEnumShape]*gongsvg_models.Rect
	map_NoteShape_Rect       map[*gongdoc_models.NoteShape]*gongsvg_models.Rect
	map_Structname_Rect      map[string]*gongsvg_models.Rect
	map_Fieldname_Link       map[string]*gongsvg_models.Link

	gongsvgStage *gongsvg_models.Stage
}

func NewDocSVGMapper(
	gongsvgStage *gongsvg_models.Stage) (docSVGMapper *DocSVGMapper) {

	docSVGMapper = new(DocSVGMapper)
	docSVGMapper.gongsvgStage = gongsvgStage

	return
}

const ClassBoxStrokeWidth = 3

func (docSVGMapper *DocSVGMapper) GenerateSvg(
	gongdocStage *gongdoc_models.Stage,
) {

	// log.Println("DocSVGMapper.GenerateSvg")

	docSVGMapper.map_GongstructShape_Rect = make(map[*gongdoc_models.GongStructShape]*gongsvg_models.Rect)
	docSVGMapper.map_GongenumShape_Rect = make(map[*gongdoc_models.GongEnumShape]*gongsvg_models.Rect)
	docSVGMapper.map_NoteShape_Rect = make(map[*gongdoc_models.NoteShape]*gongsvg_models.Rect)

	docSVGMapper.map_Structname_Rect = make(map[string]*gongsvg_models.Rect)
	docSVGMapper.map_Fieldname_Link = make(map[string]*gongsvg_models.Link)

	docSVGMapper.gongsvgStage.Reset()

	var selectedDiagram *gongdoc_models.Classdiagram

	var diagramPackage *gongdoc_models.DiagramPackage
	for diagramPackage = range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](gongdocStage) {

		selectedDiagram = diagramPackage.SelectedClassdiagram

		// if no class diagram is selected generate a blank diagram
		if selectedDiagram == nil {
			selectedDiagram = new(gongdoc_models.Classdiagram)
		}
	}
	if diagramPackage == nil {
		return
	}

	svg := new(gongsvg_models.SVG).Stage(docSVGMapper.gongsvgStage)
	svg.Name = selectedDiagram.Name
	svg.IsEditable = selectedDiagram.IsInDrawMode

	for _, gongstructShape := range selectedDiagram.GongStructShapes {

		rectLayer := new(gongsvg_models.Layer).Stage(docSVGMapper.gongsvgStage)
		rectLayer.Name = "Layer" + gongstructShape.Identifier
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(gongsvg_models.Rect).Stage(docSVGMapper.gongsvgStage)
		rect.Name = gongstructShape.Identifier

		// hook a callback on rect modifications
		rect.Impl = NewRectImplGongstructShape(gongstructShape, gongdocStage)

		docSVGMapper.map_GongstructShape_Rect[gongstructShape] = rect
		docSVGMapper.map_Structname_Rect[gongstructShape.Identifier] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = gongstructShape.Position.X
		rect.Y = gongstructShape.Position.Y

		rect.Width = gongstructShape.Width
		rect.Height = gongstructShape.Height

		rect.Stroke = gongsvg_models.Lightsalmon.ToString()
		rect.Stroke = gongsvg_models.Lightgrey.ToString()
		rect.StrokeOpacity = 1
		rect.StrokeWidth = ClassBoxStrokeWidth
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = gongsvg_models.Lightsalmon.ToString()
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
		title := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
		title.Name = gongdoc_models.IdentifierToGongObjectName(gongstructShape.Identifier)
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = gongsvg_models.RECT_TOP
		title.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.Color = gongsvg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		// additional box to hightlight the title
		titleBox := new(gongsvg_models.RectAnchoredRect).Stage(docSVGMapper.gongsvgStage)
		titleBox.Name = gongdoc_models.IdentifierToGongObjectName(gongstructShape.Identifier)
		titleBox.X_Offset = rect.StrokeWidth
		titleBox.Y_Offset = rect.StrokeWidth
		titleBox.Width = rect.Width - 2*rect.StrokeWidth - 30
		titleBox.Height = 30
		titleBox.StrokeOpacity = 0
		titleBox.RectAnchorType = gongsvg_models.RECT_TOP_LEFT
		titleBox.Color = "#ff8450"
		titleBox.Color = gongsvg_models.White.ToString()
		titleBox.WidthFollowRect = true
		titleBox.FillOpacity = 0

		titleBox.Stroke = gongsvg_models.White.ToString()
		titleBox.StrokeOpacity = 1
		titleBox.StrokeWidth = 0

		rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)

		//
		// fields
		//
		for idx, field := range gongstructShape.Fields {
			fieldText := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
			fieldText.Name = field.Name + " : " + field.Fieldtypename
			fieldText.Content = field.Name + " : " + field.Fieldtypename

			// field position
			fieldText.X_Offset = 10
			fieldText.Y_Offset = 20 + 30 + float64(idx)*15
			fieldText.RectAnchorType = gongsvg_models.RECT_TOP_LEFT
			fieldText.TextAnchorType = gongsvg_models.TEXT_ANCHOR_START

			fieldText.Color = "black"
			fieldText.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, fieldText)
		}

		//
		// number of instance (x%d)
		//
		if gongstructShape.ShowNbInstances {
			nbInstancesText := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
			nbInstancesText.Name = fmt.Sprintf("(x%d)", gongstructShape.NbInstances)
			nbInstancesText.Content = fmt.Sprintf("(x%d)", gongstructShape.NbInstances)

			// text position
			nbInstancesText.X_Offset = -5 - 2*rect.StrokeWidth
			nbInstancesText.Y_Offset = 20
			nbInstancesText.RectAnchorType = gongsvg_models.RECT_TOP_RIGHT
			nbInstancesText.TextAnchorType = gongsvg_models.TEXT_ANCHOR_END

			nbInstancesText.Color = "black"
			nbInstancesText.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, nbInstancesText)
		}
	}

	// display links between gongstruct shapes
	for _, gongstructShape := range selectedDiagram.GongStructShapes {

		startRect := docSVGMapper.map_GongstructShape_Rect[gongstructShape]
		for _, linkOfDocLib := range gongstructShape.Links {

			endRect, ok := docSVGMapper.map_Structname_Rect[linkOfDocLib.Fieldtypename]

			// if some renaming of field type name has occured, end rect might be nil
			if !ok {
				continue
			}

			link := new(gongsvg_models.Link).Stage(docSVGMapper.gongsvgStage)
			link.Name = startRect.Name + " - to - " + endRect.Name

			link.Impl = NewLinkImplLink(linkOfDocLib, gongdocStage)

			linkLayer := new(gongsvg_models.Layer).Stage(docSVGMapper.gongsvgStage)
			docSVGMapper.map_Fieldname_Link[linkOfDocLib.Identifier] = link

			linkLayer.Links = append(linkLayer.Links, link)
			svg.Layers = append(svg.Layers, linkLayer)

			// configuration
			link.Stroke = gongsvg_models.Slategray.ToString()
			link.StrokeOpacity = 1
			link.StrokeWidth = 3
			link.HasEndArrow = true
			link.EndArrowSize = 8
			link.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL

			link.StartOrientation = gongsvg_models.OrientationType(linkOfDocLib.StartOrientation)

			// take into accound legacy links
			if link.StartOrientation == "" {
				link.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
			}

			link.StartRatio = linkOfDocLib.StartRatio
			link.EndOrientation = gongsvg_models.OrientationType(linkOfDocLib.EndOrientation)
			// take into accound legacy links
			if link.EndOrientation == "" {
				link.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
			}

			link.EndRatio = linkOfDocLib.EndRatio

			link.CornerOffsetRatio = linkOfDocLib.CornerOffsetRatio

			link.CornerRadius = 8

			link.Start = startRect
			link.End = endRect

			// add text to the arrow
			targetMulitplicity := new(gongsvg_models.LinkAnchoredText).Stage(docSVGMapper.gongsvgStage)
			targetMulitplicity.AutomaticLayout = true
			targetMulitplicity.LinkAnchorType = gongsvg_models.LINK_RIGHT_OR_BOTTOM

			targetMulitplicity.Impl = NewAnchoredTextImplLinkTargetMultiplicity(linkOfDocLib, gongdocStage)
			link.TextAtArrowEnd = append(link.TextAtArrowEnd, targetMulitplicity)
			targetMulitplicity.Name = linkOfDocLib.TargetMultiplicity.ToString()
			targetMulitplicity.Content = targetMulitplicity.Name
			targetMulitplicity.X_Offset = linkOfDocLib.TargetMultiplicityOffsetX
			targetMulitplicity.Y_Offset = linkOfDocLib.TargetMultiplicityOffsetY
			targetMulitplicity.Stroke = gongsvg_models.Black.ToString()
			targetMulitplicity.StrokeOpacity = 1
			targetMulitplicity.StrokeWidth = 1
			targetMulitplicity.Color = gongsvg_models.Black.ToString()
			targetMulitplicity.FillOpacity = 100
			targetMulitplicity.FontWeight = "300"
			targetMulitplicity.FontSize = "15"
			targetMulitplicity.LetterSpacing = "0.1em"

			fieldName := new(gongsvg_models.LinkAnchoredText).Stage(docSVGMapper.gongsvgStage)
			fieldName.AutomaticLayout = true
			fieldName.LinkAnchorType = gongsvg_models.LINK_LEFT_OR_TOP

			fieldName.Impl = NewAnchoredTextImplLinkFieldName(linkOfDocLib, gongdocStage)

			link.TextAtArrowEnd = append(link.TextAtArrowEnd, fieldName)
			fieldName.Name = linkOfDocLib.GetName()
			fieldName.Content = fieldName.Name
			fieldName.Y_Offset = linkOfDocLib.FieldOffsetY
			fieldName.X_Offset = linkOfDocLib.FieldOffsetX
			fieldName.Stroke = gongsvg_models.Black.ToString()
			fieldName.StrokeOpacity = 1
			fieldName.StrokeWidth = 1
			fieldName.Color = gongsvg_models.Black.ToString()
			fieldName.FillOpacity = 100
			fieldName.FontWeight = "300"
			fieldName.FontSize = "15"
			fieldName.LetterSpacing = "0.1em"

			// add the callback

			sourceMultiplicity := new(gongsvg_models.LinkAnchoredText).Stage(docSVGMapper.gongsvgStage)
			sourceMultiplicity.AutomaticLayout = true
			sourceMultiplicity.LinkAnchorType = gongsvg_models.LINK_RIGHT_OR_BOTTOM

			sourceMultiplicity.Impl = NewAnchoredTextImplLinkSourceMultiplicity(linkOfDocLib, gongdocStage)

			link.TextAtArrowStart = append(link.TextAtArrowStart, sourceMultiplicity)
			sourceMultiplicity.Name = linkOfDocLib.SourceMultiplicity.ToString()
			sourceMultiplicity.Content = sourceMultiplicity.Name
			sourceMultiplicity.X_Offset = linkOfDocLib.SourceMultiplicityOffsetX
			sourceMultiplicity.Y_Offset = linkOfDocLib.SourceMultiplicityOffsetY
			sourceMultiplicity.Stroke = gongsvg_models.Black.ToString()
			sourceMultiplicity.StrokeOpacity = 1
			sourceMultiplicity.StrokeWidth = 1
			sourceMultiplicity.FontWeight = "300"
			sourceMultiplicity.FontSize = "15"
			sourceMultiplicity.LetterSpacing = "0.1em"
		}
	}

	//
	// GongEnumShapes
	//
	for _, gongenumShape := range selectedDiagram.GongEnumShapes {

		rectLayer := new(gongsvg_models.Layer).Stage(docSVGMapper.gongsvgStage)
		rectLayer.Name = "Layer" + gongenumShape.Identifier
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(gongsvg_models.Rect).Stage(docSVGMapper.gongsvgStage)
		rect.Name = gongenumShape.Identifier

		rect.Impl = NewRectImplGongenumShape(gongenumShape, gongdocStage)

		docSVGMapper.map_GongenumShape_Rect[gongenumShape] = rect
		docSVGMapper.map_Structname_Rect[gongenumShape.Identifier] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = gongenumShape.Position.X
		rect.Y = gongenumShape.Position.Y

		rect.Width = gongenumShape.Width
		rect.Height = gongenumShape.Height

		rect.Stroke = gongsvg_models.Lightsteelblue.ToString()
		rect.StrokeOpacity = 1
		rect.StrokeWidth = ClassBoxStrokeWidth
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = gongsvg_models.White.ToString()

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
		title := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
		title.Name = gongdoc_models.IdentifierToGongObjectName(gongenumShape.Identifier)
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = gongsvg_models.RECT_TOP
		title.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.Color = gongsvg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		// additional box to hightlight the title
		titleBox := new(gongsvg_models.RectAnchoredRect).Stage(docSVGMapper.gongsvgStage)
		titleBox.Name = gongdoc_models.IdentifierToGongObjectName(gongenumShape.Identifier)
		titleBox.X_Offset = 0
		titleBox.Y_Offset = 0
		titleBox.Width = rect.Width
		titleBox.Height = 30
		titleBox.RectAnchorType = gongsvg_models.RECT_TOP_LEFT
		titleBox.Color = gongsvg_models.White.ToString()
		titleBox.WidthFollowRect = true
		titleBox.FillOpacity = 0.0
		titleBox.StrokeWidth = 0

		rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)

		//
		// fields
		//
		for idx, field := range gongenumShape.GongEnumValueEntrys {
			fieldText := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
			fieldText.Name = field.Name
			fieldText.Content = field.Name

			// field position
			fieldText.X_Offset = 10
			fieldText.Y_Offset = 20 + 30 + float64(idx)*15
			fieldText.RectAnchorType = gongsvg_models.RECT_TOP_LEFT
			fieldText.TextAnchorType = gongsvg_models.TEXT_ANCHOR_START

			fieldText.Color = "black"
			fieldText.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, fieldText)
		}
	}

	//
	// Notes
	//
	for _, noteShape := range selectedDiagram.NoteShapes {

		rectLayer := new(gongsvg_models.Layer).Stage(docSVGMapper.gongsvgStage)
		rectLayer.Name = "Layer" + noteShape.Identifier
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(gongsvg_models.Rect).Stage(docSVGMapper.gongsvgStage)
		rect.Name = noteShape.Identifier

		rect.Impl = NewRectImplNoteShape(noteShape, gongdocStage)

		docSVGMapper.map_NoteShape_Rect[noteShape] = rect
		docSVGMapper.map_Structname_Rect[noteShape.Identifier] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = noteShape.X
		rect.Y = noteShape.Y

		rect.Width = noteShape.Width
		rect.Height = noteShape.Height

		rect.Stroke = gongsvg_models.Lightskyblue.ToString()
		rect.StrokeOpacity = 1
		rect.StrokeWidth = ClassBoxStrokeWidth
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = gongsvg_models.White.ToString()

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
		title := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
		title.Name = gongdoc_models.IdentifierToGongObjectName(noteShape.Identifier)
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = gongsvg_models.RECT_TOP
		title.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.FontStyle = "oblique"
		title.Color = gongsvg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		//
		// Title
		//
		content := new(gongsvg_models.RectAnchoredText).Stage(docSVGMapper.gongsvgStage)
		content.Name = noteShape.Body
		content.Content = content.Name
		content.X_Offset = 0
		content.Y_Offset = 40
		content.RectAnchorType = gongsvg_models.RECT_TOP
		content.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
		content.FontWeight = "normal"
		content.Color = gongsvg_models.Black.ToString()
		content.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, content)

		// // additional box to hightlight the title
		// titleBox := new(gongsvg_models.RectAnchoredRect).Stage(docSVGMapper.gongsvgStage)
		// titleBox.Name = gongdoc_models.IdentifierToGongObjectName(noteShape.Identifier)
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
	for _, noteShape := range selectedDiagram.NoteShapes {

		startRect := docSVGMapper.map_NoteShape_Rect[noteShape]
		_ = startRect

		for _, noteLink := range noteShape.NoteShapeLinks {

			if noteLink.Type == gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE {

				var endRect *gongsvg_models.Rect
				var ok bool
				// find the endRect
				endRect, ok = docSVGMapper.map_Structname_Rect[noteLink.Identifier]
				if ok {
					// create the link
					link := new(gongsvg_models.Link).Stage(docSVGMapper.gongsvgStage)
					link.Name = startRect.Name + " - to - " + endRect.Name
					linkLayer := new(gongsvg_models.Layer).Stage(docSVGMapper.gongsvgStage)

					linkLayer.Links = append(linkLayer.Links, link)
					svg.Layers = append(svg.Layers, linkLayer)

					// configuration
					link.Stroke = gongsvg_models.Slategray.ToString()
					link.StrokeOpacity = 1
					link.StrokeWidth = 2
					link.StrokeDashArray = "2 2"

					// first version
					link.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL
					link.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
					link.StartRatio = 0.5
					link.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
					link.EndRatio = 0.5

					// more elegent
					link.Type = gongsvg_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
					link.StartAnchorType = gongsvg_models.ANCHOR_CENTER
					link.EndAnchorType = gongsvg_models.ANCHOR_CENTER

					link.CornerOffsetRatio = (endRect.X - startRect.X - 80) / startRect.Width

					link.CornerRadius = 3

					link.Start = startRect
					link.End = endRect
				}

			}
			if noteLink.Type == gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD {

				var endLink *gongsvg_models.Link
				_ = endLink
				var ok bool
				// find the endLink
				endLink, ok = docSVGMapper.map_Fieldname_Link[noteLink.Identifier]
				if ok {
					rectLinkLink := new(gongsvg_models.RectLinkLink).Stage(docSVGMapper.gongsvgStage)
					rectLinkLink.Name = startRect.Name + " - to - " + endLink.Name
					rectLinkLinkLayer := new(gongsvg_models.Layer).Stage(docSVGMapper.gongsvgStage)

					rectLinkLinkLayer.RectLinkLinks = append(rectLinkLinkLayer.RectLinkLinks, rectLinkLink)
					svg.Layers = append(svg.Layers, rectLinkLinkLayer)

					// configuration
					rectLinkLink.Stroke = gongsvg_models.Slategray.ToString()
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
	docSVGMapper.gongsvgStage.Commit()
}

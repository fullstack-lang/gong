package models

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svgFlossEquation() {
	stager.flossDiagramSvgStage.Reset()

	var diagramFlossEquation *DiagramFlossEquation
	for diagram_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if diagram_.IsChecked {
			diagramFlossEquation = diagram_
			break
		}
	}

	if diagramFlossEquation == nil {
		stager.flossDiagramSvgStage.Commit()
		return
	}

	svgObject := stager.generateSvgObjectFlossEquation(diagramFlossEquation)
	svg.StageBranch(stager.flossDiagramSvgStage, svgObject)
	stager.flossDiagramSvgStage.Commit()
}

func (stager *Stager) generateSvgObjectFlossEquation(diagram *DiagramFlossEquation) *svg.SVG {
	svgObject := &svg.SVG{Name: "FLOSS Equation SVG"}

	diagWidth := math.Max(diagram.Width, 1050.0)
	diagHeight := math.Max(diagram.Height, 750.0)

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagWidth
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagHeight

	layer := &svg.Layer{Name: "Equation Layer"}
	svgObject.Layers = append(svgObject.Layers, layer)

	compareAnalysis := diagram.GetOwningCompareAnalysis()
	if compareAnalysis == nil {
		for ca := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
			for _, d := range ca.DiagramFlossEquations {
				if d == diagram {
					compareAnalysis = ca
					break
				}
			}
		}
	}

	if compareAnalysis == nil || compareAnalysis.FromSystem == nil || compareAnalysis.ToSystem == nil {
		infoRect := &svg.Rect{
			Name:   "Info Background",
			X:      100,
			Y:      250,
			Width:  750,
			Height: 120,
			RX:     10,
			Presentation: svg.Presentation{
				Color:         "#F5F5F5",
				FillOpacity:   0.9,
				Stroke:        "#9E9E9E",
				StrokeWidth:   1.5,
				StrokeOpacity: 1.0,
			},
		}
		layer.Rects = append(layer.Rects, infoRect)

		infoText := &svg.Text{
			Name:    "Info Text",
			X:       150,
			Y:       315,
			Content: "Please select FromSystem and ToSystem in CompareAnalysis to render the FLOSS Equation diagram.",
			Presentation: svg.Presentation{
				Color:       "#424242",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, infoText)
		return svgObject
	}

	fromSys := compareAnalysis.FromSystem
	toSys := compareAnalysis.ToSystem
	alpha := compareAnalysis.Alpha
	if alpha == 0 {
		alpha = 1.0
	}

	var cFrom, cTo, pFrom, pTo, eFrom, eTo float64
	for _, c := range fromSys.Complexities {
		cFrom += c.Strength
	}
	for _, c := range toSys.Complexities {
		cTo += c.Strength
	}
	for _, p := range fromSys.Performances {
		pFrom += p.Strength
	}
	for _, p := range toSys.Performances {
		pTo += p.Strength
	}
	for _, e := range fromSys.Efforts {
		eFrom += e.Strength
	}
	for _, e := range toSys.Efforts {
		eTo += e.Strength
	}

	deltaC := cTo - cFrom
	deltaP := pTo - pFrom
	deltaE := eTo - eFrom

	beta := 0.0
	if deltaE != 0 {
		beta = (alpha*deltaP - deltaC) / deltaE
	}

	scale := diagram.Scale
	if scale <= 0 {
		scale = 5.0
	}

	// Layout coordinates adapted to diagram dimensions
	yGround := diagHeight - 80.0
	colWidth := 250.0
	xCol1 := 70.0  // Delta C
	xCol2 := 390.0 // Alpha * Delta P
	xCol3 := 710.0 // Beta * Delta E

	root := stager.getRootLibrary()
	nbPixPerChar := root.NbPixPerCharacter
	if nbPixPerChar <= 0 {
		nbPixPerChar = 8.0
	}

	// Header banner
	headerRect := &svg.Rect{
		Name:   "Header Card",
		X:      40,
		Y:      20,
		Width:  980,
		Height: 80,
		RX:     8,
		Presentation: svg.Presentation{
			Color:         "#FAFAFA",
			FillOpacity:   1.0,
			Stroke:        "#B0BEC5",
			StrokeWidth:   1.5,
			StrokeOpacity: 1.0,
		},
	}
	layer.Rects = append(layer.Rects, headerRect)

	headerTitle := &svg.Text{
		Name:    "Header Title",
		X:       60,
		Y:       50,
		Content: fmt.Sprintf("FLOSS Equation: ΔC = α · ΔP - β · ΔE  ⟺  ΔC + β · ΔE = α · ΔP (%s → %s)", fromSys.Name, toSys.Name),
		Presentation: svg.Presentation{
			Color:       "#1A237E",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, headerTitle)

	headerCalc := &svg.Text{
		Name:    "Header Calculation",
		X:       60,
		Y:       80,
		Content: fmt.Sprintf("Values: ΔC = %.2f | α·ΔP = %.2f (α = %.2f, ΔP = %.2f) | β·ΔE = %.2f (β = %.2f, ΔE = %.2f)", deltaC, alpha*deltaP, alpha, deltaP, beta*deltaE, beta, deltaE),
		Presentation: svg.Presentation{
			Color:       "#37474F",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, headerCalc)

	// Ground line (Abscissa)
	groundLine := &svg.Line{
		Name: "Ground Baseline",
		X1:   40,
		Y1:   yGround,
		X2:   1010,
		Y2:   yGround,
		Presentation: svg.Presentation{
			Stroke:        "#37474F",
			StrokeWidth:   2.5,
			StrokeOpacity: 1.0,
		},
	}
	layer.Lines = append(layer.Lines, groundLine)

	groundLabel := &svg.Text{
		Name:    "Ground Label",
		X:       920,
		Y:       yGround + 20,
		Content: "Abscissa line",
		Presentation: svg.Presentation{
			Color:       "#78909C",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, groundLabel)

	// Map to hold element to rect pointers for Note link connections
	map_Element_Rect := make(map[any]*svg.Rect)
	diagram.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)
	diagram.map_SvgRect_Complexity = make(map[*svg.Rect]*Complexity)
	diagram.map_SvgRect_Performance = make(map[*svg.Rect]*Performance)
	diagram.map_SvgRect_Effort = make(map[*svg.Rect]*Effort)

	//
	// Column 1: Delta C (Amber / Orange) - Bottom on ground
	//
	heightC := math.Max(math.Abs(deltaC)*scale, 24.0)
	yTip1 := yGround - heightC

	col1BaseY := yGround
	if len(toSys.Complexities) > 0 {
		currY := col1BaseY
		for idx, c := range toSys.Complexities {
			itemH := (c.Strength / math.Max(cTo, 1.0)) * heightC
			itemRect := &svg.Rect{
				Name:   fmt.Sprintf("C_%d", idx),
				X:      xCol1,
				Y:      currY - itemH,
				Width:  colWidth,
				Height: itemH,
				RX:     4,
				Presentation: svg.Presentation{
					Color:         "#FFF8E1",
					FillOpacity:   0.95,
					Stroke:        "#FFA000",
					StrokeWidth:   1.5,
					StrokeOpacity: 1.0,
				},
			}
			layer.Rects = append(layer.Rects, itemRect)
			map_Element_Rect[c] = itemRect
			diagram.map_SvgRect_Complexity[itemRect] = c

			itemRect.CanHaveBottomHandle = true
			itemRect.CanHaveLeftHandle = true
			itemRect.CanHaveRightHandle = true
			itemRect.CanHaveTopHandle = true


			content := fmt.Sprintf("%s (%.1f)", c.Name, c.Strength)
			content = strutils.WrapStringPreservingNewlinesScaled(content, colWidth-20, nbPixPerChar, 13.0, 15.0)

			fontSize := "13px"
			if itemH < 30 {
				fontSize = "11px"
			}

			title := new(svg.RectAnchoredText)
			title.Name = fmt.Sprintf("C_text_%d", idx)
			title.Content = content
			title.FontSize = fontSize
			title.FontWeight = "600"
			title.Color = "#B78103"
			title.FillOpacity = 1.0
			title.Stroke = "#B78103"
			title.StrokeWidth = 0
			title.StrokeOpacity = 1.0
			title.RectAnchorType = svg.RECT_CENTER_MIDDLE
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
			itemRect.RectAnchoredTexts = append(itemRect.RectAnchoredTexts, title)

			currY -= itemH
		}
	} else {
		col1Rect := &svg.Rect{
			Name:   "Col 1 Rect",
			X:      xCol1,
			Y:      yTip1,
			Width:  colWidth,
			Height: heightC,
			RX:     4,
			Presentation: svg.Presentation{
				Color:         "#FFF8E1",
				FillOpacity:   0.95,
				Stroke:        "#FFA000",
				StrokeWidth:   1.5,
				StrokeOpacity: 1.0,
			},
		}
		layer.Rects = append(layer.Rects, col1Rect)
	}

	col1Label := &svg.Text{
		Name:    "Col 1 Label",
		X:       xCol1 + 45,
		Y:       yGround + 30,
		Content: fmt.Sprintf("ΔC = %.2f", deltaC),
		Presentation: svg.Presentation{
			Color:       "#E65100",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, col1Label)

	//
	// Column 2: alpha * Delta P (Emerald / Green) - Bottom on ground
	//
	heightP := math.Max(math.Abs(alpha*deltaP)*scale, 24.0)
	yTip2 := yGround - heightP

	col2BaseY := yGround
	if len(toSys.Performances) > 0 {
		currY := col2BaseY
		for idx, p := range toSys.Performances {
			itemH := (p.Strength / math.Max(pTo, 1.0)) * heightP
			itemRect := &svg.Rect{
				Name:   fmt.Sprintf("P_%d", idx),
				X:      xCol2,
				Y:      currY - itemH,
				Width:  colWidth,
				Height: itemH,
				RX:     4,
				Presentation: svg.Presentation{
					Color:         "#E8F5E9",
					FillOpacity:   0.95,
					Stroke:        "#2E7D32",
					StrokeWidth:   1.5,
					StrokeOpacity: 1.0,
				},
			}
			layer.Rects = append(layer.Rects, itemRect)
			map_Element_Rect[p] = itemRect
			diagram.map_SvgRect_Performance[itemRect] = p

			itemRect.CanHaveBottomHandle = true
			itemRect.CanHaveLeftHandle = true
			itemRect.CanHaveRightHandle = true
			itemRect.CanHaveTopHandle = true


			content := fmt.Sprintf("%s (%.1f · α=%.1f)", p.Name, p.Strength, p.Strength*alpha)
			content = strutils.WrapStringPreservingNewlinesScaled(content, colWidth-20, nbPixPerChar, 13.0, 15.0)

			fontSize := "13px"
			if itemH < 30 {
				fontSize = "11px"
			}

			title := new(svg.RectAnchoredText)
			title.Name = fmt.Sprintf("P_text_%d", idx)
			title.Content = content
			title.FontSize = fontSize
			title.FontWeight = "600"
			title.Color = "#1B5E20"
			title.FillOpacity = 1.0
			title.Stroke = "#1B5E20"
			title.StrokeWidth = 0
			title.StrokeOpacity = 1.0
			title.RectAnchorType = svg.RECT_CENTER_MIDDLE
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
			itemRect.RectAnchoredTexts = append(itemRect.RectAnchoredTexts, title)

			currY -= itemH
		}
	} else {
		col2Rect := &svg.Rect{
			Name:   "Col 2 Rect",
			X:      xCol2,
			Y:      yTip2,
			Width:  colWidth,
			Height: heightP,
			RX:     4,
			Presentation: svg.Presentation{
				Color:         "#E8F5E9",
				FillOpacity:   0.95,
				Stroke:        "#2E7D32",
				StrokeWidth:   1.5,
				StrokeOpacity: 1.0,
			},
		}
		layer.Rects = append(layer.Rects, col2Rect)
	}

	col2Label := &svg.Text{
		Name:    "Col 2 Label",
		X:       xCol2 + 25,
		Y:       yGround + 30,
		Content: fmt.Sprintf("α · ΔP = %.2f (α=%.2f)", alpha*deltaP, alpha),
		Presentation: svg.Presentation{
			Color:       "#2E7D32",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, col2Label)

	//
	// Column 3: beta * Delta E (Slate / Blue) - Bottom is on the tip of the first column
	//
	heightE := math.Max(math.Abs(beta*deltaE)*scale, 24.0)
	col3BaseY := yTip1
	yTip3 := col3BaseY - heightE

	if len(toSys.Efforts) > 0 {
		currY := col3BaseY
		for idx, e := range toSys.Efforts {
			itemH := (e.Strength / math.Max(eTo, 1.0)) * heightE
			itemRect := &svg.Rect{
				Name:   fmt.Sprintf("E_%d", idx),
				X:      xCol3,
				Y:      currY - itemH,
				Width:  colWidth,
				Height: itemH,
				RX:     4,
				Presentation: svg.Presentation{
					Color:         "#E3F2FD",
					FillOpacity:   0.95,
					Stroke:        "#1976D2",
					StrokeWidth:   1.5,
					StrokeOpacity: 1.0,
				},
			}
			layer.Rects = append(layer.Rects, itemRect)
			map_Element_Rect[e] = itemRect
			diagram.map_SvgRect_Effort[itemRect] = e

			itemRect.CanHaveBottomHandle = true
			itemRect.CanHaveLeftHandle = true
			itemRect.CanHaveRightHandle = true
			itemRect.CanHaveTopHandle = true


			content := fmt.Sprintf("%s (%.1f · β=%.1f)", e.Name, e.Strength, e.Strength*beta)
			content = strutils.WrapStringPreservingNewlinesScaled(content, colWidth-20, nbPixPerChar, 13.0, 15.0)

			fontSize := "13px"
			if itemH < 30 {
				fontSize = "11px"
			}

			title := new(svg.RectAnchoredText)
			title.Name = fmt.Sprintf("E_text_%d", idx)
			title.Content = content
			title.FontSize = fontSize
			title.FontWeight = "600"
			title.Color = "#0D47A1"
			title.FillOpacity = 1.0
			title.Stroke = "#0D47A1"
			title.StrokeWidth = 0
			title.StrokeOpacity = 1.0
			title.RectAnchorType = svg.RECT_CENTER_MIDDLE
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
			itemRect.RectAnchoredTexts = append(itemRect.RectAnchoredTexts, title)

			currY -= itemH
		}
	} else {
		col3Rect := &svg.Rect{
			Name:   "Col 3 Rect",
			X:      xCol3,
			Y:      yTip3,
			Width:  colWidth,
			Height: heightE,
			RX:     4,
			Presentation: svg.Presentation{
				Color:         "#E3F2FD",
				FillOpacity:   0.95,
				Stroke:        "#1976D2",
				StrokeWidth:   1.5,
				StrokeOpacity: 1.0,
			},
		}
		layer.Rects = append(layer.Rects, col3Rect)
	}

	col3Label := &svg.Text{
		Name:    "Col 3 Label",
		X:       xCol3 + 25,
		Y:       col3BaseY + 20,
		Content: fmt.Sprintf("β · ΔE = %.2f (β=%.2f)", beta*deltaE, beta),
		Presentation: svg.Presentation{
			Color:       "#1565C0",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, col3Label)

	// Step line from tip of Col 1 to bottom of Col 3
	stepLine := &svg.Line{
		Name: "Col 1 Tip to Col 3 Base Guide",
		X1:   xCol1 + colWidth,
		Y1:   yTip1,
		X2:   xCol3,
		Y2:   col3BaseY,
		Presentation: svg.Presentation{
			Stroke:          "#90A4AE",
			StrokeWidth:     1.5,
			StrokeDashArray: "4,4",
			StrokeOpacity:   0.8,
		},
	}
	layer.Lines = append(layer.Lines, stepLine)

	// Equality alignment line from top of Col 2 to top of Col 3
	equalityLine := &svg.Line{
		Name: "Equality Alignment Guide",
		X1:   xCol2 + colWidth,
		Y1:   yTip2,
		X2:   xCol3,
		Y2:   yTip3,
		Presentation: svg.Presentation{
			Stroke:          "#00897B",
			StrokeWidth:     2.0,
			StrokeDashArray: "5,5",
			StrokeOpacity:   0.9,
		},
	}
	layer.Lines = append(layer.Lines, equalityLine)

	equalityText := &svg.Text{
		Name:    "Equality Label",
		X:       (xCol2 + colWidth + xCol3) / 2 - 40,
		Y:       (yTip2+yTip3)/2 - 8,
		Content: "ΔC + β·ΔE ≡ α·ΔP",
		Presentation: svg.Presentation{
			Color:       "#00695C",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, equalityText)

	//
	// Notes & Link Shapes on Equation Diagram
	//
	diagram.map_Note_Rect = make(map[*Note]*svg.Rect)
	for _, noteShape := range diagram.Note_Shapes {
		if noteShape.IsHidden || noteShape.Note == nil {
			continue
		}

		rect := new(svg.Rect)
		layer.Rects = append(layer.Rects, rect)
		diagram.map_Note_Rect[noteShape.Note] = rect
		diagram.map_SvgRect_NoteShape[rect] = noteShape


		rect.Name = noteShape.GetName()
		rect.X = noteShape.GetX()
		rect.Y = noteShape.GetY()
		rect.Width = noteShape.GetWidth()
		rect.Height = noteShape.GetHeight()

		rect.Color = "#FFF9C4"
		rect.FillOpacity = 1.0
		rect.Stroke = "#FBC02D"
		rect.StrokeWidth = 1.5
		rect.StrokeOpacity = 1.0
		rect.RX = 2

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(noteShape.Note, GetPointerToGongstructName[*Note]())
		}
		rect.OnMove = onMoveRectElement(stager, noteShape, true)
		rect.OnResize = onResizeRectElement(stager, noteShape)

		title := new(svg.RectAnchoredText)
		title.Name = noteShape.Note.Name
		content := "📝 " + noteShape.Note.Name
		if noteShape.Note.Description != "" {
			content += "\n" + noteShape.Note.Description
		}

		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlinesScaled(content, rect.Width-20, nbPixPerChar, 13.0, 15.0)
		}

		title.Content = content
		title.Color = "#5D4037"
		title.FillOpacity = 1.0
		title.Stroke = "#5D4037"
		title.StrokeWidth = 0
		title.StrokeOpacity = 1.0
		title.FontSize = "13px"
		title.FontStyle = "italic"
		title.FontWeight = "normal"
		title.RectAnchorType = svg.RECT_TOP_LEFT
		title.TextAnchorType = svg.TEXT_ANCHOR_START
		title.X_Offset = 10
		title.Y_Offset = 18
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}

	for _, shape := range diagram.NoteComplexityShapes {
		if shape.IsHidden || shape.Note == nil || shape.Complexity == nil {
			continue
		}
		startRect := diagram.map_Note_Rect[shape.Note]
		endRect := map_Element_Rect[shape.Complexity]
		if startRect == nil || endRect == nil {
			continue
		}

		startOrientation := shape.GetStartOrientation()
		if startOrientation == "" {
			startOrientation = ORIENTATION_HORIZONTAL
		}
		endOrientation := shape.GetEndOrientation()
		if endOrientation == "" {
			endOrientation = ORIENTATION_HORIZONTAL
		}
		startRatio := shape.GetStartRatio()
		if startRatio == 0 {
			startRatio = 0.5
		}
		endRatio := shape.GetEndRatio()
		if endRatio == 0 {
			endRatio = 0.5
		}
		cornerOffsetRatio := shape.GetCornerOffsetRatio()
		if cornerOffsetRatio == 0 {
			cornerOffsetRatio = 1.5
		}

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.StartOrientation = svg.OrientationType(startOrientation)
		link.StartRatio = startRatio
		link.End = endRect
		link.EndOrientation = svg.OrientationType(endOrientation)
		link.EndRatio = endRatio
		link.CornerOffsetRatio = cornerOffsetRatio
		link.CornerRadius = 5
		link.Type = svg.LINK_TYPE_FLOATING_ORTHOGONAL
		link.Stroke = "#FFA000"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0

		link.OnChange = func(updatedLink *svg.Link) {
			shape.SetStartRatio(updatedLink.StartRatio)
			shape.SetEndRatio(updatedLink.EndRatio)
			shape.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
			shape.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
			shape.SetEndOrientation(OrientationType(updatedLink.EndOrientation))
			stager.stage.Commit()
		}
	}

	for _, shape := range diagram.NotePerformanceShapes {
		if shape.IsHidden || shape.Note == nil || shape.Performance == nil {
			continue
		}
		startRect := diagram.map_Note_Rect[shape.Note]
		endRect := map_Element_Rect[shape.Performance]
		if startRect == nil || endRect == nil {
			continue
		}

		startOrientation := shape.GetStartOrientation()
		if startOrientation == "" {
			startOrientation = ORIENTATION_HORIZONTAL
		}
		endOrientation := shape.GetEndOrientation()
		if endOrientation == "" {
			endOrientation = ORIENTATION_HORIZONTAL
		}
		startRatio := shape.GetStartRatio()
		if startRatio == 0 {
			startRatio = 0.5
		}
		endRatio := shape.GetEndRatio()
		if endRatio == 0 {
			endRatio = 0.5
		}
		cornerOffsetRatio := shape.GetCornerOffsetRatio()
		if cornerOffsetRatio == 0 {
			cornerOffsetRatio = 1.5
		}

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.StartOrientation = svg.OrientationType(startOrientation)
		link.StartRatio = startRatio
		link.End = endRect
		link.EndOrientation = svg.OrientationType(endOrientation)
		link.EndRatio = endRatio
		link.CornerOffsetRatio = cornerOffsetRatio
		link.CornerRadius = 5
		link.Type = svg.LINK_TYPE_FLOATING_ORTHOGONAL
		link.Stroke = "#2E7D32"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0

		link.OnChange = func(updatedLink *svg.Link) {
			shape.SetStartRatio(updatedLink.StartRatio)
			shape.SetEndRatio(updatedLink.EndRatio)
			shape.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
			shape.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
			shape.SetEndOrientation(OrientationType(updatedLink.EndOrientation))
			stager.stage.Commit()
		}
	}

	for _, shape := range diagram.NoteEffortShapes {
		if shape.IsHidden || shape.Note == nil || shape.Effort == nil {
			continue
		}
		startRect := diagram.map_Note_Rect[shape.Note]
		endRect := map_Element_Rect[shape.Effort]
		if startRect == nil || endRect == nil {
			continue
		}

		startOrientation := shape.GetStartOrientation()
		if startOrientation == "" {
			startOrientation = ORIENTATION_HORIZONTAL
		}
		endOrientation := shape.GetEndOrientation()
		if endOrientation == "" {
			endOrientation = ORIENTATION_HORIZONTAL
		}
		startRatio := shape.GetStartRatio()
		if startRatio == 0 {
			startRatio = 0.5
		}
		endRatio := shape.GetEndRatio()
		if endRatio == 0 {
			endRatio = 0.5
		}
		cornerOffsetRatio := shape.GetCornerOffsetRatio()
		if cornerOffsetRatio == 0 {
			cornerOffsetRatio = 1.5
		}

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.StartOrientation = svg.OrientationType(startOrientation)
		link.StartRatio = startRatio
		link.End = endRect
		link.EndOrientation = svg.OrientationType(endOrientation)
		link.EndRatio = endRatio
		link.CornerOffsetRatio = cornerOffsetRatio
		link.CornerRadius = 5
		link.Type = svg.LINK_TYPE_FLOATING_ORTHOGONAL
		link.Stroke = "#1976D2"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0

		link.OnChange = func(updatedLink *svg.Link) {
			shape.SetStartRatio(updatedLink.StartRatio)
			shape.SetEndRatio(updatedLink.EndRatio)
			shape.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
			shape.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
			shape.SetEndOrientation(OrientationType(updatedLink.EndOrientation))
			stager.stage.Commit()
		}
	}


	return svgObject
}

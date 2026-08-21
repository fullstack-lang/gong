package models

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	stager.systemDiagramSvgStage.Reset()

	var diagramFlossEquation *DiagramFlossEquation
	for diagram_ := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if diagram_.IsChecked {
			diagramFlossEquation = diagram_
			break
		}
	}

	if diagramFlossEquation != nil {
		svgObject := stager.generateSvgObjectFlossEquation(diagramFlossEquation)
		svg.StageBranch(stager.systemDiagramSvgStage, svgObject)
		stager.svgObjectDiagramFloss = svgObject
		stager.svgObjectDiagramFloss.OnUpdate = stager.onUpdateSVG
	}

	stager.systemDiagramSvgStage.Commit()
}

func (stager *Stager) svgFlossEquation() {
	stager.svg()
}

func (stager *Stager) generateSvgObjectFlossEquation(diagram *DiagramFlossEquation) *svg.SVG {
	svgObject := &svg.SVG{Name: "FLOSS Equation SVG"}

	diagWidth := math.Max(diagram.Width, 1050.0)
	diagHeight := math.Max(diagram.Height, 750.0)

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagWidth
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagHeight
	svgObject.IsEditable = diagram.IsEditable()

	layer := &svg.Layer{Name: "Equation Layer"}
	svgObject.Layers = append(svgObject.Layers, layer)

	compareAnalysis := diagram.GetOwningCompareAnalysis()
	owningSystem := diagram.GetOwningSystem()

	if compareAnalysis == nil && owningSystem == nil {
		for ca := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
			for _, d := range ca.DiagramFlossEquations {
				if d == diagram {
					compareAnalysis = ca
					break
				}
			}
		}
		if compareAnalysis == nil {
			for sys := range *GetGongstructInstancesSet[System](stager.stage) {
				for _, d := range sys.DiagramFlossEquations {
					if d == diagram {
						owningSystem = sys
						break
					}
				}
			}
		}
	}

	var fromSys *System
	var toSys *System
	alpha := 1.0
	beta := 1.0

	if compareAnalysis != nil {
		fromSys = compareAnalysis.FromSystem
		toSys = compareAnalysis.ToSystem
		if compareAnalysis.Alpha != 0 {
			alpha = compareAnalysis.Alpha
		}
		if compareAnalysis.Beta != 0 {
			beta = compareAnalysis.Beta
		}
	} else if owningSystem != nil {
		fromSys = nil
		toSys = owningSystem
		alpha = 1.0
		beta = 1.0
	}

	if toSys == nil || (compareAnalysis != nil && fromSys == nil) {
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

		infoTitle := new(svg.RectAnchoredText)
		infoTitle.Name = "Info Text"
		if compareAnalysis != nil {
			infoTitle.Content = "Please select FromSystem and ToSystem in CompareAnalysis to render the FLOSS Equation diagram."
		} else {
			infoTitle.Content = "Please select a System to render the FLOSS Equation diagram."
		}
		infoTitle.FontSize = "14px"
		infoTitle.FontWeight = "500"
		infoTitle.Color = "#424242"
		infoTitle.FillOpacity = 1.0
		infoTitle.Stroke = "#424242"
		infoTitle.StrokeWidth = 0
		infoTitle.StrokeOpacity = 1.0
		infoTitle.RectAnchorType = svg.RECT_CENTER_MIDDLE
		infoTitle.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		infoRect.RectAnchoredTexts = append(infoRect.RectAnchoredTexts, infoTitle)
		return svgObject
	}

	var cFrom, cTo, pFrom, pTo, eFrom, eTo float64
	if fromSys != nil {
		for _, c := range fromSys.Complexities {
			cFrom += c.Strength
		}
		for _, p := range fromSys.Performances {
			pFrom += p.Strength
		}
		for _, e := range fromSys.Efforts {
			eFrom += e.Strength
		}
	}
	for _, c := range toSys.Complexities {
		cTo += c.Strength
	}
	for _, p := range toSys.Performances {
		pTo += p.Strength
	}
	for _, e := range toSys.Efforts {
		eTo += e.Strength
	}

	deltaC := cTo - cFrom
	deltaP := pTo - pFrom
	deltaE := eTo
	_ = eFrom

	scale := diagram.Scale
	if scale <= 0 {
		scale = 5.0
	}

	rhs := alpha*deltaP - beta*deltaE
	diff := deltaC - rhs

	// Layout coordinates adapted to diagram dimensions and DefaultBoxWidth
	yGround := diagHeight - 80.0
	colWidth := diagram.GetDefaultBoxWidth()
	if colWidth <= 0 {
		colWidth = 250.0
	}
	colSpacing := 40.0
	xMargin := 80.0
	xCol1 := xMargin                              // Column 1: Delta C
	xCol2 := xCol1 + colWidth + colSpacing        // Column 2: Alpha * Delta P
	xCol3 := xCol2 + colWidth + colSpacing        // Column 3: Beta * Delta E
	columnsRight := xCol3 + colWidth

	root := stager.getRootLibrary()
	nbPixPerChar := root.NbPixPerCharacter
	if nbPixPerChar <= 0 {
		nbPixPerChar = 8.0
	}

	// Header banner
	headerWidth := math.Max(columnsRight+220.0, diagWidth-80.0)
	headerRect := &svg.Rect{
		Name:   "Header Card",
		X:      40,
		Y:      20,
		Width:  headerWidth,
		Height: 90,
		RX:     8,
		Presentation: svg.Presentation{
			Color:         "#FAFAFA",
			FillOpacity:   1.0,
			Stroke:        "#B0BEC5",
			StrokeWidth:   1.5,
			StrokeOpacity: 1.0,
		},
	}
	if compareAnalysis != nil {
		headerRect.OnSelect = onSelectRectElement(stager, compareAnalysis)
	} else if owningSystem != nil {
		headerRect.OnSelect = onSelectRectElement(stager, owningSystem)
	}
	layer.Rects = append(layer.Rects, headerRect)

	// Compare Analysis or System Title
	headerTitle := new(svg.RectAnchoredText)
	headerTitle.Name = "Header Title"
	if compareAnalysis != nil {
		headerTitle.Content = compareAnalysis.Name
	} else {
		headerTitle.Content = toSys.Name
	}
	headerTitle.FontSize = "17px"
	headerTitle.FontWeight = "700"
	headerTitle.Color = "#0D47A1"
	headerTitle.FillOpacity = 1.0
	headerTitle.Stroke = "#0D47A1"
	headerTitle.StrokeWidth = 0
	headerTitle.StrokeOpacity = 1.0
	headerTitle.RectAnchorType = svg.RECT_TOP_LEFT
	headerTitle.TextAnchorType = svg.TEXT_ANCHOR_START
	headerTitle.X_Offset = 20
	headerTitle.Y_Offset = 26
	headerRect.RectAnchoredTexts = append(headerRect.RectAnchoredTexts, headerTitle)

	// Formula line
	headerFormula := new(svg.RectAnchoredText)
	headerFormula.Name = "Header Formula"
	if fromSys == nil {
		headerFormula.Content = fmt.Sprintf("FLOSS Equation: C = α·P - β·E   |   System: %s", toSys.Name)
	} else {
		headerFormula.Content = fmt.Sprintf("FLOSS Equation: ΔC = α·ΔP - β·ΔE   (From: %s  →  To: %s)", fromSys.Name, toSys.Name)
	}
	headerFormula.FontSize = "14px"
	headerFormula.FontWeight = "600"
	headerFormula.Color = "#37474F"
	headerFormula.FillOpacity = 1.0
	headerFormula.Stroke = "#37474F"
	headerFormula.StrokeWidth = 0
	headerFormula.StrokeOpacity = 1.0
	headerFormula.RectAnchorType = svg.RECT_TOP_LEFT
	headerFormula.TextAnchorType = svg.TEXT_ANCHOR_START
	headerFormula.X_Offset = 20
	headerFormula.Y_Offset = 52
	headerRect.RectAnchoredTexts = append(headerRect.RectAnchoredTexts, headerFormula)

	// Values summary
	headerValues := new(svg.RectAnchoredText)
	headerValues.Name = "Header Values"
	if fromSys == nil {
		headerValues.Content = fmt.Sprintf("C=%.2f   P=%.2f (α=%.2f → α·P=%.2f)   E=%.2f (β=%.2f → β·E=%.2f)   [α·P - β·E = %.2f,  Diff = %.2f]",
			deltaC, deltaP, alpha, alpha*deltaP, deltaE, beta, beta*deltaE, rhs, diff)
	} else {
		headerValues.Content = fmt.Sprintf("ΔC=%.2f   ΔP=%.2f (α=%.2f → α·ΔP=%.2f)   ΔE=%.2f (β=%.2f → β·ΔE=%.2f)   [RHS = %.2f,  Diff = %.2f]",
			deltaC, deltaP, alpha, alpha*deltaP, deltaE, beta, beta*deltaE, rhs, diff)
	}
	headerValues.FontSize = "13px"
	headerValues.FontWeight = "500"
	headerValues.Color = "#616161"
	headerValues.FillOpacity = 1.0
	headerValues.Stroke = "#616161"
	headerValues.StrokeWidth = 0
	headerValues.StrokeOpacity = 1.0
	headerValues.RectAnchorType = svg.RECT_TOP_LEFT
	headerValues.TextAnchorType = svg.TEXT_ANCHOR_START
	headerValues.X_Offset = 20
	headerValues.Y_Offset = 76
	headerRect.RectAnchoredTexts = append(headerRect.RectAnchoredTexts, headerValues)

	// -------------------------------------------------------------
	// Ground baseline
	// -------------------------------------------------------------

	groundLine := &svg.Line{
		Name: "Ground Baseline",
		X1:   40,
		Y1:   yGround,
		X2:   columnsRight + 80.0,
		Y2:   yGround,
		Presentation: svg.Presentation{
			Stroke:        "#78909C",
			StrokeWidth:   2.0,
			StrokeOpacity: 1.0,
		},
	}
	layer.Lines = append(layer.Lines, groundLine)

	groundText := &svg.Text{
		Name:    "Ground Text",
		X:       40,
		Y:       yGround - 8,
		Content: "Baseline (0.0)",
		Presentation: svg.Presentation{
			Color:       "#90A4AE",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, groundText)

	// Map to hold element to rect pointers
	map_Element_Rect := make(map[any]*svg.Rect)
	diagram.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)
	diagram.map_SvgRect_Complexity = make(map[*svg.Rect]*Complexity)
	diagram.map_SvgRect_Performance = make(map[*svg.Rect]*Performance)
	diagram.map_SvgRect_Effort = make(map[*svg.Rect]*Effort)

	//
	// Column 1: Delta C (Amber / Orange) - Bottom on ground
	//
	heightC := math.Max(math.Abs(deltaC)*scale, 24.0)

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
			itemRect.OnSelect = onSelectRectElement(stager, c)

			itemRect.CanHaveBottomHandle = true
			itemRect.CanHaveLeftHandle = true
			itemRect.CanHaveRightHandle = true
			itemRect.CanHaveTopHandle = true

			content := c.Name
			if diagram.AreQuantitativeElementsVisible {
				content = fmt.Sprintf("%s (%.1f)", c.Name, c.Strength)
			}
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
			Y:      yGround - heightC,
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

	col1LabelText := "ΔC"
	if fromSys == nil {
		col1LabelText = "C"
	}
	if diagram.AreQuantitativeElementsVisible {
		if fromSys == nil {
			col1LabelText = fmt.Sprintf("C = %.2f", deltaC)
		} else {
			col1LabelText = fmt.Sprintf("ΔC = %.2f", deltaC)
		}
	}
	col1Label := &svg.Text{
		Name:    "Col 1 Label",
		X:       xCol1 + 45,
		Y:       yGround + 30,
		Content: col1LabelText,
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
			itemRect.OnSelect = onSelectRectElement(stager, p)

			itemRect.CanHaveBottomHandle = true
			itemRect.CanHaveLeftHandle = true
			itemRect.CanHaveRightHandle = true
			itemRect.CanHaveTopHandle = true

			content := p.Name
			if diagram.AreQuantitativeElementsVisible {
				content = fmt.Sprintf("%s (%.1f · α=%.1f)", p.Name, p.Strength, p.Strength*alpha)
			}
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

	col2LabelText := "α · ΔP"
	if fromSys == nil {
		col2LabelText = "α · P"
	}
	if diagram.AreQuantitativeElementsVisible {
		if fromSys == nil {
			col2LabelText = fmt.Sprintf("α · P = %.2f", alpha*deltaP)
		} else {
			col2LabelText = fmt.Sprintf("α · ΔP = %.2f", alpha*deltaP)
		}
	}
	col2Label := &svg.Text{
		Name:    "Col 2 Label",
		X:       xCol2 + 45,
		Y:       yGround + 30,
		Content: col2LabelText,
		Presentation: svg.Presentation{
			Color:       "#2E7D32",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, col2Label)

	//
	// Column 3: beta * Delta E (Slate / Blue)
	//
	heightE := math.Max(math.Abs(beta*deltaE)*scale, 24.0)
	yTop3 := yTip2
	yBottom3 := yTop3 + heightE

	if len(toSys.Efforts) > 0 {
		currY := yTop3
		for idx, e := range toSys.Efforts {
			itemH := (e.Strength / math.Max(eTo, 1.0)) * heightE
			itemRect := &svg.Rect{
				Name:   fmt.Sprintf("E_%d", idx),
				X:      xCol3,
				Y:      currY,
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
			itemRect.OnSelect = onSelectRectElement(stager, e)

			itemRect.CanHaveBottomHandle = true
			itemRect.CanHaveLeftHandle = true
			itemRect.CanHaveRightHandle = true
			itemRect.CanHaveTopHandle = true

			content := e.Name
			if diagram.AreQuantitativeElementsVisible {
				content = fmt.Sprintf("%s (%.1f · β=%.1f)", e.Name, e.Strength, e.Strength*beta)
			}
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

			currY += itemH
		}
	} else {
		col3Rect := &svg.Rect{
			Name:   "Col 3 Rect",
			X:      xCol3,
			Y:      yTop3,
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

	col3LabelY := math.Max(yBottom3, yGround) + 20
	col3LabelText := "β · ΔE"
	if fromSys == nil {
		col3LabelText = "β · E"
	}
	if diagram.AreQuantitativeElementsVisible {
		if fromSys == nil {
			col3LabelText = fmt.Sprintf("β · E = %.2f", beta*deltaE)
		} else {
			col3LabelText = fmt.Sprintf("β · ΔE = %.2f", beta*deltaE)
		}
	}
	col3Label := &svg.Text{
		Name:    "Col 3 Label",
		X:       xCol3 + 45,
		Y:       col3LabelY,
		Content: col3LabelText,
		Presentation: svg.Presentation{
			Color:       "#1976D2",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, col3Label)

	//
	// Guides and Indicators
	//
	peakLine := &svg.Line{
		Name: "Peak Indicator Line",
		X1:   xCol2,
		Y1:   yTip2,
		X2:   xCol3 + colWidth,
		Y2:   yTip2,
		Presentation: svg.Presentation{
			Stroke:           "#388E3C",
			StrokeWidth:      1.5,
			StrokeOpacity:    1.0,
			StrokeDashArray: "4 3",
		},
	}
	layer.Lines = append(layer.Lines, peakLine)

	peakLabelText := "α · ΔP peak"
	if fromSys == nil {
		peakLabelText = "α · P peak"
	}
	peakText := &svg.Text{
		Name:    "Peak Label",
		X:       xCol3 + colWidth + 10,
		Y:       yTip2 + 4,
		Content: peakLabelText,
		Presentation: svg.Presentation{
			Color:       "#388E3C",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, peakText)

	rhsLine := &svg.Line{
		Name: "RHS Level Line",
		X1:   xCol1,
		Y1:   yBottom3,
		X2:   xCol3 + colWidth,
		Y2:   yBottom3,
		Presentation: svg.Presentation{
			Stroke:           "#1565C0",
			StrokeWidth:      1.5,
			StrokeOpacity:    1.0,
			StrokeDashArray: "4 3",
		},
	}
	layer.Lines = append(layer.Lines, rhsLine)

	rhsLabelText := "α·ΔP - β·ΔE"
	if fromSys == nil {
		rhsLabelText = "α·P - β·E"
	}
	rhsText := &svg.Text{
		Name:    "RHS Label",
		X:       xCol3 + colWidth + 10,
		Y:       yBottom3 + 4,
		Content: rhsLabelText,
		Presentation: svg.Presentation{
			Color:       "#1565C0",
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, rhsText)

	diffColor := "#43A047"
	diffTextMsg := "Equilibrium (ΔC ≈ α·ΔP - β·ΔE)"
	if fromSys == nil {
		diffTextMsg = "Equilibrium (C ≈ α·P - β·E)"
	}
	if math.Abs(diff) > 2.0 {
		if diff > 0 {
			diffColor = "#E53935"
			diffTextMsg = fmt.Sprintf("Under-performing: Complexity exceeds Net Performance by %.2f", diff)
		} else {
			diffColor = "#1E88E5"
			diffTextMsg = fmt.Sprintf("Favorable Margin: Net Performance exceeds Complexity by %.2f", -diff)
		}
	}

	diffLine := &svg.Line{
		Name: "Delta Indicator Line",
		X1:   xCol1 + colWidth/2,
		Y1:   yGround - heightC,
		X2:   xCol1 + colWidth/2,
		Y2:   yBottom3,
		Presentation: svg.Presentation{
			Stroke:        diffColor,
			StrokeWidth:   2.5,
			StrokeOpacity: 1.0,
		},
	}
	layer.Lines = append(layer.Lines, diffLine)

	indicatorLabel := &svg.Text{
		Name:    "Indicator Message",
		X:       40,
		Y:       yGround + 65,
		Content: diffTextMsg,
		Presentation: svg.Presentation{
			Color:       diffColor,
			FillOpacity: 1.0,
			Stroke:      "transparent",
		},
	}
	layer.Texts = append(layer.Texts, indicatorLabel)

	//
	// Notes & Link Shapes
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
		rect.OnSelect = onSelectRectElement(stager, noteShape.Note)

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

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

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.End = endRect
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.HasEndArrow = false
		link.Stroke = "#FFA000"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0
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

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.End = endRect
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.HasEndArrow = false
		link.Stroke = "#2E7D32"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0
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

		link := new(svg.Link)
		layer.Links = append(layer.Links, link)
		link.Name = startRect.Name + " to " + endRect.Name
		link.Start = startRect
		link.End = endRect
		link.StartAnchorType = svg.ANCHOR_CENTER
		link.EndAnchorType = svg.ANCHOR_CENTER
		link.Type = svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS
		link.HasEndArrow = false
		link.Stroke = "#1976D2"
		link.StrokeWidth = 1.5
		link.StrokeDashArray = "5 5"
		link.StrokeOpacity = 1.0
	}

	return svgObject
}


package models

import (
	"fmt"
	"math"
	"strings"

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

	var toComplexities []*Complexity
	var toComplexitiesSysMap map[*Complexity]*System
	var toPerformances []*Performance
	var toPerformancesSysMap map[*Performance]*System
	var toEfforts []*Effort
	var toEffortsSysMap map[*Effort]*System

	if diagram.AreSubsystemsVisible && toSys.AreCPEsCompoundedFromSubSystems {
		toComplexities, toComplexitiesSysMap = toSys.GetEffectiveComplexities()
		toPerformances, toPerformancesSysMap = toSys.GetEffectivePerformances()
		toEfforts, toEffortsSysMap = toSys.GetEffectiveEfforts()
	} else {
		toComplexities = toSys.Complexities
		toComplexitiesSysMap = make(map[*Complexity]*System)
		for _, c := range toSys.Complexities {
			toComplexitiesSysMap[c] = toSys
		}
		toPerformances = toSys.Performances
		toPerformancesSysMap = make(map[*Performance]*System)
		for _, p := range toSys.Performances {
			toPerformancesSysMap[p] = toSys
		}
		toEfforts = toSys.Efforts
		toEffortsSysMap = make(map[*Effort]*System)
		for _, e := range toSys.Efforts {
			toEffortsSysMap[e] = toSys
		}
	}

	var fromComplexities []*Complexity
	var fromPerformances []*Performance
	var fromEfforts []*Effort
	var fromComplexitiesSysMap map[*Complexity]*System
	var fromPerformancesSysMap map[*Performance]*System
	var fromEffortsSysMap map[*Effort]*System
	if fromSys != nil {
		if diagram.AreSubsystemsVisible && fromSys.AreCPEsCompoundedFromSubSystems {
			fromComplexities, fromComplexitiesSysMap = fromSys.GetEffectiveComplexities()
			fromPerformances, fromPerformancesSysMap = fromSys.GetEffectivePerformances()
			fromEfforts, fromEffortsSysMap = fromSys.GetEffectiveEfforts()
		} else {
			fromComplexities = fromSys.Complexities
			fromComplexitiesSysMap = make(map[*Complexity]*System)
			for _, c := range fromSys.Complexities {
				fromComplexitiesSysMap[c] = fromSys
			}
			fromPerformances = fromSys.Performances
			fromPerformancesSysMap = make(map[*Performance]*System)
			for _, p := range fromSys.Performances {
				fromPerformancesSysMap[p] = fromSys
			}
			fromEfforts = fromSys.Efforts
			fromEffortsSysMap = make(map[*Effort]*System)
			for _, e := range fromSys.Efforts {
				fromEffortsSysMap[e] = fromSys
			}
		}
	}

	var cFrom, cTo, pFrom, pTo, eFrom, eTo float64
	for _, c := range fromComplexities {
		cFrom += c.Strength
	}
	for _, p := range fromPerformances {
		pFrom += p.Strength
	}
	for _, e := range fromEfforts {
		eFrom += e.Strength
	}
	for _, c := range toComplexities {
		cTo += c.Strength
	}
	for _, p := range toPerformances {
		pTo += p.Strength
	}
	for _, e := range toEfforts {
		eTo += e.Strength
	}

	deltaC := cTo - cFrom
	deltaP := pTo - pFrom
	deltaE := eTo - eFrom

	scale := diagram.Scale
	if scale <= 0 {
		scale = 5.0
	}

	rhs := alpha*deltaP - beta*deltaE
	diff := deltaC - rhs

	isDelta := fromSys != nil

	// Layout coordinates adapted to diagram dimensions and DefaultBoxWidth
	yGround := diagHeight - 80.0
	colWidth := diagram.GetDefaultBoxWidth()
	if colWidth <= 0 {
		colWidth = 250.0
	}

	pairGap := 15.0
	colSpacing := 40.0
	if isDelta {
		colSpacing = 50.0
	}
	xMargin := 80.0

	var xCol1_V2, xCol1_V1, xCol2_V2, xCol2_V1, xCol3_V2, xCol3_V1 float64
	var columnsRight float64

	if isDelta && !diagram.IsInDelta3ColumnsMode {
		xCol1_V2 = xMargin
		xCol1_V1 = xCol1_V2 + colWidth + pairGap
		xCol2_V2 = xCol1_V1 + colWidth + colSpacing
		xCol2_V1 = xCol2_V2 + colWidth + pairGap
		xCol3_V2 = xCol2_V1 + colWidth + colSpacing
		xCol3_V1 = xCol3_V2 + colWidth + pairGap
		columnsRight = xCol3_V1 + colWidth
	} else {
		xCol1_V2 = xMargin
		xCol2_V2 = xCol1_V2 + colWidth + colSpacing
		xCol3_V2 = xCol2_V2 + colWidth + colSpacing
		columnsRight = xCol3_V2 + colWidth
	}

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
	isCompoundedBreakdown := diagram.AreSubsystemsVisible && toSys.AreCPEsCompoundedFromSubSystems
	headerTitle := new(svg.RectAnchoredText)
	headerTitle.Name = "Header Title"
	if compareAnalysis != nil {
		headerTitle.Content = compareAnalysis.Name
	} else if isCompoundedBreakdown {
		headerTitle.Content = fmt.Sprintf("%s (Subsystems Breakdown)", toSys.Name)
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
	if !isDelta {
		if isCompoundedBreakdown {
			headerFormula.Content = fmt.Sprintf("FLOSS Equation: C = α·P - β·E   |   System: %s [Subsystems Breakdown]", toSys.Name)
		} else {
			headerFormula.Content = fmt.Sprintf("FLOSS Equation: C = α·P - β·E   |   System: %s", toSys.Name)
		}
	} else {
		headerFormula.Content = fmt.Sprintf("FLOSS Equation: ΔC = α·ΔP - β·ΔE   (V1: %s  →  V2: %s)", fromSys.Name, toSys.Name)
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
	if !isDelta {
		headerValues.Content = fmt.Sprintf("C=%.2f   P=%.2f (α=%.2f → α·P=%.2f)   E=%.2f (β=%.2f → β·E=%.2f)   [α·P - β·E = %.2f,  Diff = %.2f]",
			cTo, pTo, alpha, alpha*pTo, eTo, beta, beta*eTo, rhs, diff)
	} else {
		headerValues.Content = fmt.Sprintf("ΔC=%.2f (V2:%.2f - V1:%.2f)   α·ΔP=%.2f (V2:%.2f - V1:%.2f)   β·ΔE=%.2f (V2:%.2f - V1:%.2f)   [RHS = %.2f,  Diff = %.2f]",
			deltaC, cTo, cFrom, alpha*deltaP, alpha*pTo, alpha*pFrom, beta*deltaE, beta*eTo, beta*eFrom, rhs, diff)
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

	// Helper for rendering stack of rects inside a column
	renderColumnStack := func(
		xPos, baseY, totalH, totalVal float64,
		isV1 bool,
		category string, // "C", "P", "E"
		items []any,
		sysMap map[any]*System,
		ownerSys *System,
		multiplier float64, // 1.0 for C, alpha for P, beta for E
	) {
		fillColor := "#FFF8E1"
		strokeColor := "#FFA000"
		textColor := "#B78103"
		if isV1 {
			fillColor = "#FFFDE7"
		}
		if category == "P" {
			fillColor = "#E8F5E9"
			strokeColor = "#2E7D32"
			textColor = "#1B5E20"
			if isV1 {
				fillColor = "#F1F8E9"
			}
		} else if category == "E" {
			fillColor = "#E3F2FD"
			strokeColor = "#1976D2"
			textColor = "#0D47A1"
			if isV1 {
				fillColor = "#EBF5FB"
			}
		}

		strokeDash := ""
		if isV1 {
			strokeDash = "4 2"
		}

		if len(items) > 0 {
			currY := baseY
			for idx, item := range items {
				var strength float64
				var name string
				switch v := item.(type) {
				case *Complexity:
					strength = v.Strength
					name = v.Name
				case *Performance:
					strength = v.Strength
					name = v.Name
				case *Effort:
					strength = v.Strength
					name = v.Name
				}

				var itemH float64
				if totalVal > 0.000001 {
					itemH = (strength / totalVal) * totalH
				} else {
					itemH = totalH / float64(len(items))
				}

				itemRect := &svg.Rect{
					Name:   fmt.Sprintf("%s_%s_%d", category, ownerSys.Name, idx),
					X:      xPos,
					Y:      currY,
					Width:  colWidth,
					Height: itemH,
					RX:     4,
					Presentation: svg.Presentation{
						Color:           fillColor,
						FillOpacity:     0.95,
						Stroke:          strokeColor,
						StrokeWidth:     1.5,
						StrokeOpacity:   1.0,
						StrokeDashArray: strokeDash,
					},
				}
				layer.Rects = append(layer.Rects, itemRect)
				map_Element_Rect[item] = itemRect
				switch v := item.(type) {
				case *Complexity:
					diagram.map_SvgRect_Complexity[itemRect] = v
					itemRect.OnSelect = onSelectRectElement(stager, v)
				case *Performance:
					diagram.map_SvgRect_Performance[itemRect] = v
					itemRect.OnSelect = onSelectRectElement(stager, v)
				case *Effort:
					diagram.map_SvgRect_Effort[itemRect] = v
					itemRect.OnSelect = onSelectRectElement(stager, v)
				}

				itemRect.CanHaveBottomHandle = true
				itemRect.CanHaveLeftHandle = true
				itemRect.CanHaveRightHandle = true
				itemRect.CanHaveTopHandle = true

				content := name
				if sysOwner := sysMap[item]; sysOwner != nil && sysOwner != ownerSys {
					content = fmt.Sprintf("[%s] %s", sysOwner.Name, name)
				}
				if diagram.AreQuantitativeElementsVisible {
					if multiplier != 1.0 {
						if sysOwner := sysMap[item]; sysOwner != nil && sysOwner != ownerSys {
							content = fmt.Sprintf("[%s] %s (%.2f · factor=%.2f)", sysOwner.Name, name, strength, strength*multiplier)
						} else {
							content = fmt.Sprintf("%s (%.2f · factor=%.2f)", name, strength, strength*multiplier)
						}
					} else {
						if sysOwner := sysMap[item]; sysOwner != nil && sysOwner != ownerSys {
							content = fmt.Sprintf("[%s] %s (%.2f)", sysOwner.Name, name, strength)
						} else {
							content = fmt.Sprintf("%s (%.2f)", name, strength)
						}
					}
				}
				content = strutils.WrapStringPreservingNewlinesScaled(content, colWidth-20, nbPixPerChar, 13.0, 15.0)

				fontSize := "13px"
				if itemH < 30 {
					fontSize = "11px"
				}

				title := new(svg.RectAnchoredText)
				title.Name = fmt.Sprintf("%s_text_%d", category, idx)
				title.Content = content
				title.FontSize = fontSize
				title.FontWeight = "600"
				title.Color = textColor
				title.FillOpacity = 1.0
				title.Stroke = textColor
				title.StrokeWidth = 0
				title.StrokeOpacity = 1.0
				title.RectAnchorType = svg.RECT_CENTER_MIDDLE
				title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
				itemRect.RectAnchoredTexts = append(itemRect.RectAnchoredTexts, title)

				currY += itemH
			}
		} else {
			colRect := &svg.Rect{
				Name:   fmt.Sprintf("%s Col Rect", category),
				X:      xPos,
				Y:      baseY,
				Width:  colWidth,
				Height: totalH,
				RX:     4,
				Presentation: svg.Presentation{
					Color:           fillColor,
					FillOpacity:     0.95,
					Stroke:          strokeColor,
					StrokeWidth:     1.5,
					StrokeOpacity:   1.0,
					StrokeDashArray: strokeDash,
				},
			}
			layer.Rects = append(layer.Rects, colRect)
		}
	}

	var toCompItems []any
	toMapC := make(map[any]*System)
	for _, c := range toComplexities {
		toCompItems = append(toCompItems, c)
		toMapC[c] = toComplexitiesSysMap[c]
	}

	var fromCompItems []any
	fromMapC := make(map[any]*System)
	for _, c := range fromComplexities {
		fromCompItems = append(fromCompItems, c)
		fromMapC[c] = fromComplexitiesSysMap[c]
	}

	var toPerfItems []any
	toMapP := make(map[any]*System)
	for _, p := range toPerformances {
		toPerfItems = append(toPerfItems, p)
		toMapP[p] = toPerformancesSysMap[p]
	}

	var fromPerfItems []any
	fromMapP := make(map[any]*System)
	for _, p := range fromPerformances {
		fromPerfItems = append(fromPerfItems, p)
		fromMapP[p] = fromPerformancesSysMap[p]
	}

	var toEffItems []any
	toMapE := make(map[any]*System)
	for _, e := range toEfforts {
		toEffItems = append(toEffItems, e)
		toMapE[e] = toEffortsSysMap[e]
	}

	var fromEffItems []any
	fromMapE := make(map[any]*System)
	for _, e := range fromEfforts {
		fromEffItems = append(fromEffItems, e)
		fromMapE[e] = fromEffortsSysMap[e]
	}

	var yTipC_Indicator, yRHS_Indicator float64
	var indicatorX float64 = xCol1_V2 + colWidth/2

	if !isDelta {
		// -------------------------------------------------------------
		// SINGLE SYSTEM MODE (3 Columns)
		// -------------------------------------------------------------
		heightC_V2 := math.Max(cTo*scale, 24.0)
		yTipC_V2 := yGround - heightC_V2
		renderColumnStack(xCol1_V2, yTipC_V2, heightC_V2, cTo, false, "C", toCompItems, toMapC, toSys, 1.0)

		col1LabelText := "C"
		if diagram.AreQuantitativeElementsVisible {
			col1LabelText = fmt.Sprintf("C = %.2f", cTo)
		}
		col1Label := &svg.Text{
			Name:    "Col 1 Label",
			X:       xCol1_V2 + colWidth/2,
			Y:       yGround + 40,
			Content: col1LabelText,
			Presentation: svg.Presentation{
				Color:       "#E65100",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col1Label)

		heightP_V2 := math.Max(alpha*pTo*scale, 24.0)
		yTipP_V2 := yGround - heightP_V2
		renderColumnStack(xCol2_V2, yTipP_V2, heightP_V2, pTo, false, "P", toPerfItems, toMapP, toSys, alpha)

		col2LabelText := "α · P"
		if diagram.AreQuantitativeElementsVisible {
			col2LabelText = fmt.Sprintf("α · P = %.2f", alpha*pTo)
		}
		col2Label := &svg.Text{
			Name:    "Col 2 Label",
			X:       xCol2_V2 + colWidth/2,
			Y:       yGround + 40,
			Content: col2LabelText,
			Presentation: svg.Presentation{
				Color:       "#2E7D32",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col2Label)

		heightE_V2 := math.Max(beta*eTo*scale, 24.0)
		yTopE_V2 := yTipP_V2
		yBottomE_V2 := yTopE_V2 + heightE_V2
		renderColumnStack(xCol3_V2, yTopE_V2, heightE_V2, eTo, false, "E", toEffItems, toMapE, toSys, beta)

		col3LabelText := "β · E"
		if diagram.AreQuantitativeElementsVisible {
			col3LabelText = fmt.Sprintf("β · E = %.2f", beta*eTo)
		}
		col3Label := &svg.Text{
			Name:    "Col 3 Label",
			X:       xCol3_V2 + colWidth/2,
			Y:       math.Max(yBottomE_V2, yGround) + 40,
			Content: col3LabelText,
			Presentation: svg.Presentation{
				Color:       "#1976D2",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col3Label)

		// Guides & Indicators
		peakLine := &svg.Line{
			Name: "Peak Indicator Line",
			X1:   xCol2_V2,
			Y1:   yTipP_V2,
			X2:   columnsRight,
			Y2:   yTipP_V2,
			Presentation: svg.Presentation{
				Stroke:          "#388E3C",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, peakLine)

		peakText := &svg.Text{
			Name:    "Peak Label",
			X:       columnsRight + 10,
			Y:       yTipP_V2 + 4,
			Content: "α · P peak",
			Presentation: svg.Presentation{
				Color:       "#388E3C",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, peakText)

		rhsLine := &svg.Line{
			Name: "RHS Level Line",
			X1:   xCol1_V2,
			Y1:   yBottomE_V2,
			X2:   columnsRight,
			Y2:   yBottomE_V2,
			Presentation: svg.Presentation{
				Stroke:          "#1565C0",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, rhsLine)

		rhsText := &svg.Text{
			Name:    "RHS Label",
			X:       columnsRight + 10,
			Y:       yBottomE_V2 + 4,
			Content: "α·P - β·E level",
			Presentation: svg.Presentation{
				Color:       "#1565C0",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, rhsText)

		yTipC_Indicator = yTipC_V2
		yRHS_Indicator = yBottomE_V2

	} else if !diagram.IsInDelta3ColumnsMode {
		// -------------------------------------------------------------
		// 6 COLUMNS DELTA PAIR MODE (Visual closed loop (C2-C1) = (P2-P1) - (E2-E1))
		// -------------------------------------------------------------

		// Column 1: Complexity Pair
		heightC_V1 := math.Max(cFrom*scale, 24.0)
		yTipC_V1 := yGround - heightC_V1 // V1 bottom sits on abscissa baseline (yGround)

		heightC_V2 := math.Max(cTo*scale, 24.0)
		yTipC_V2 := yTipC_V1 // Top tips aligned!
		yBottomC_V2 := yTipC_V2 + heightC_V2

		renderColumnStack(xCol1_V1, yTipC_V1, heightC_V1, cFrom, true, "C", fromCompItems, fromMapC, fromSys, 1.0)
		renderColumnStack(xCol1_V2, yTipC_V2, heightC_V2, cTo, false, "C", toCompItems, toMapC, toSys, 1.0)

		// Dashed guide across tips of V2 and V1
		tipCLine := &svg.Line{
			Name: "C Tip Guide Line",
			X1:   xCol1_V2,
			Y1:   yTipC_V1,
			X2:   xCol1_V1 + colWidth,
			Y2:   yTipC_V1,
			Presentation: svg.Presentation{
				Stroke:          "#FFA000",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, tipCLine)

		// Dashed guide starting from C2 bottom extending to P2 bottom
		c2BottomLine := &svg.Line{
			Name: "C2 Bottom Guide Line",
			X1:   xCol1_V2,
			Y1:   yBottomC_V2,
			X2:   xCol2_V2 + colWidth,
			Y2:   yBottomC_V2,
			Presentation: svg.Presentation{
				Stroke:          "#2196F3",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, c2BottomLine)

		v2Label := &svg.Text{
			Name:    "C V2 Sub-Label",
			X:       xCol1_V2 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V2 (C=%.2f)", cTo),
			Presentation: svg.Presentation{
				Color:       "#B78103",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v2Label)

		v1Label := &svg.Text{
			Name:    "C V1 Sub-Label",
			X:       xCol1_V1 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V1 (C=%.2f)", cFrom),
			Presentation: svg.Presentation{
				Color:       "#B78103",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v1Label)

		col1CenterX := (xCol1_V2 + xCol1_V1 + colWidth) / 2
		col1Label := &svg.Text{
			Name:    "Col 1 Label",
			X:       col1CenterX,
			Y:       yGround + 40,
			Content: fmt.Sprintf("ΔC = %.2f", deltaC),
			Presentation: svg.Presentation{
				Color:       "#E65100",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col1Label)

		// Column 2: Performance Pair
		// "P2 bottom shall be at the dashed line."
		// "P1 top shall be aligned with P2 top"
		heightP_V2 := math.Max(alpha*pTo*scale, 24.0)
		yBottomP_V2 := yBottomC_V2
		yTopP_V2 := yBottomP_V2 - heightP_V2

		heightP_V1 := math.Max(alpha*pFrom*scale, 24.0)
		yTopP_V1 := yTopP_V2 // P1 top aligned with P2 top
		yBottomP_V1 := yTopP_V1 + heightP_V1

		renderColumnStack(xCol2_V2, yTopP_V2, heightP_V2, pTo, false, "P", toPerfItems, toMapP, toSys, alpha)
		renderColumnStack(xCol2_V1, yTopP_V1, heightP_V1, pFrom, true, "P", fromPerfItems, fromMapP, fromSys, alpha)

		// Top peak line across P2 and P1
		peakLine := &svg.Line{
			Name: "Peak Indicator Line",
			X1:   xCol2_V2,
			Y1:   yTopP_V2,
			X2:   xCol2_V1 + colWidth,
			Y2:   yTopP_V2,
			Presentation: svg.Presentation{
				Stroke:          "#388E3C",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, peakLine)

		// "From P1 bottom should start a dashed line that will serve as E1 bottom"
		p1BottomLine := &svg.Line{
			Name: "P1 Bottom Guide Line",
			X1:   xCol2_V1,
			Y1:   yBottomP_V1,
			X2:   xCol3_V1 + colWidth,
			Y2:   yBottomP_V1,
			Presentation: svg.Presentation{
				Stroke:          "#4CAF50",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, p1BottomLine)

		v2PLabel := &svg.Text{
			Name:    "P V2 Sub-Label",
			X:       xCol2_V2 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V2 (α·P=%.2f)", alpha*pTo),
			Presentation: svg.Presentation{
				Color:       "#2E7D32",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v2PLabel)

		v1PLabel := &svg.Text{
			Name:    "P V1 Sub-Label",
			X:       xCol2_V1 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V1 (α·P=%.2f)", alpha*pFrom),
			Presentation: svg.Presentation{
				Color:       "#2E7D32",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v1PLabel)

		col2CenterX := (xCol2_V2 + xCol2_V1 + colWidth) / 2
		col2Label := &svg.Text{
			Name:    "Col 2 Label",
			X:       col2CenterX,
			Y:       yGround + 40,
			Content: fmt.Sprintf("α · ΔP = %.2f", alpha*deltaP),
			Presentation: svg.Presentation{
				Color:       "#2E7D32",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col2Label)

		// Column 3: Effort Pair
		// "From P1 bottom should start a dashed line that will serve as E1 bottom"
		// "E2 top is aligned with E1 top"
		heightE_V1 := math.Max(beta*eFrom*scale, 24.0)
		yBottomE_V1 := yBottomP_V1
		yTopE_V1 := yBottomE_V1 - heightE_V1

		heightE_V2 := math.Max(beta*eTo*scale, 24.0)
		yTopE_V2 := yTopE_V1 // E2 top is aligned with E1 top
		yBottomE_V2 := yTopE_V2 + heightE_V2

		renderColumnStack(xCol3_V1, yTopE_V1, heightE_V1, eFrom, true, "E", fromEffItems, fromMapE, fromSys, beta)
		renderColumnStack(xCol3_V2, yTopE_V2, heightE_V2, eTo, false, "E", toEffItems, toMapE, toSys, beta)

		// Dashed guide across tops of E2 and E1
		tipELine := &svg.Line{
			Name: "E Tip Guide Line",
			X1:   xCol3_V2,
			Y1:   yTopE_V2,
			X2:   xCol3_V1 + colWidth,
			Y2:   yTopE_V2,
			Presentation: svg.Presentation{
				Stroke:          "#1976D2",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, tipELine)

		v2ELabel := &svg.Text{
			Name:    "E V2 Sub-Label",
			X:       xCol3_V2 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V2 (β·E=%.2f)", beta*eTo),
			Presentation: svg.Presentation{
				Color:       "#1976D2",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v2ELabel)

		v1ELabel := &svg.Text{
			Name:    "E V1 Sub-Label",
			X:       xCol3_V1 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V1 (β·E=%.2f)", beta*eFrom),
			Presentation: svg.Presentation{
				Color:       "#1976D2",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v1ELabel)

		col3CenterX := (xCol3_V2 + xCol3_V1 + colWidth) / 2
		col3Label := &svg.Text{
			Name:    "Col 3 Label",
			X:       col3CenterX,
			Y:       yGround + 40,
			Content: fmt.Sprintf("β · ΔE = %.2f", beta*deltaE),
			Presentation: svg.Presentation{
				Color:       "#1976D2",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col3Label)

		// Level line at E2 bottom
		rhsLine := &svg.Line{
			Name: "E2 Bottom Level Line",
			X1:   xCol3_V2,
			Y1:   yBottomE_V2,
			X2:   columnsRight,
			Y2:   yBottomE_V2,
			Presentation: svg.Presentation{
				Stroke:          "#1565C0",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, rhsLine)

		rhsText := &svg.Text{
			Name:    "E2 Bottom Label",
			X:       columnsRight + 10,
			Y:       yBottomE_V2 + 4,
			Content: fmt.Sprintf("E2 bottom (diff = %.2f)", diff),
			Presentation: svg.Presentation{
				Color:       "#1565C0",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, rhsText)

		yTipC_Indicator = yBottomE_V2
		yRHS_Indicator = yGround
		indicatorX = xCol3_V2 + colWidth/2

	} else {
		// -------------------------------------------------------------
		// 3 COLUMNS DELTA MODE (V2 - V1 computation view)
		// -------------------------------------------------------------

		// Helper for single delta rect
		renderDeltaRect := func(
			xPos, yPos, height float64,
			category string,
			fillColor, strokeColor, textColor string,
			lines []string,
			items []any,
		) *svg.Rect {
			rect := &svg.Rect{
				Name:   fmt.Sprintf("Delta %s Rect", category),
				X:      xPos,
				Y:      yPos,
				Width:  colWidth,
				Height: height,
				RX:     4,
				Presentation: svg.Presentation{
					Color:         fillColor,
					FillOpacity:   0.95,
					Stroke:        strokeColor,
					StrokeWidth:   1.5,
					StrokeOpacity: 1.0,
				},
			}
			layer.Rects = append(layer.Rects, rect)
			for _, item := range items {
				map_Element_Rect[item] = rect
				switch v := item.(type) {
				case *Complexity:
					diagram.map_SvgRect_Complexity[rect] = v
				case *Performance:
					diagram.map_SvgRect_Performance[rect] = v
				case *Effort:
					diagram.map_SvgRect_Effort[rect] = v
				}
			}
			rect.CanHaveBottomHandle = true
			rect.CanHaveLeftHandle = true
			rect.CanHaveRightHandle = true
			rect.CanHaveTopHandle = true

			rawContent := strings.Join(lines, "\n")
			content := strutils.WrapStringPreservingNewlinesScaled(rawContent, colWidth-20, nbPixPerChar, 13.0, 15.0)

			fontSize := "13px"
			if height < 40 {
				fontSize = "11px"
			}

			title := new(svg.RectAnchoredText)
			title.Name = fmt.Sprintf("Delta_%s_text", category)
			title.Content = content
			title.FontSize = fontSize
			title.FontWeight = "600"
			title.Color = textColor
			title.FillOpacity = 1.0
			title.Stroke = textColor
			title.StrokeWidth = 0
			title.StrokeOpacity = 1.0
			title.RectAnchorType = svg.RECT_CENTER_MIDDLE
			title.TextAnchorType = svg.TEXT_ANCHOR_CENTER
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
			return rect
		}

		// Column 1: Net Complexity Change (ΔC)
		heightDeltaC := math.Max(math.Abs(deltaC)*scale, 50.0)
		var yTopDeltaC, yBottomDeltaC float64
		if deltaC <= 0 {
			yTopDeltaC = yGround
			yBottomDeltaC = yGround + heightDeltaC
		} else {
			yTopDeltaC = yGround - heightDeltaC
			yBottomDeltaC = yGround
		}

		var linesC []string
		for _, c := range toComplexities {
			if diagram.AreQuantitativeElementsVisible {
				linesC = append(linesC, fmt.Sprintf("V2: %s (%.2f)", c.Name, c.Strength))
			} else {
				linesC = append(linesC, fmt.Sprintf("V2: %s", c.Name))
			}
		}
		for _, c := range fromComplexities {
			if diagram.AreQuantitativeElementsVisible {
				linesC = append(linesC, fmt.Sprintf("V1: %s (%.2f)", c.Name, c.Strength))
			} else {
				linesC = append(linesC, fmt.Sprintf("V1: %s", c.Name))
			}
		}
		var allCompItems []any
		for _, c := range toComplexities {
			allCompItems = append(allCompItems, c)
		}
		for _, c := range fromComplexities {
			allCompItems = append(allCompItems, c)
		}
		renderDeltaRect(xCol1_V2, yTopDeltaC, heightDeltaC, "C", "#FFF8E1", "#FFA000", "#B78103", linesC, allCompItems)

		v1v2SubC := &svg.Text{
			Name:    "ΔC Sub-Label",
			X:       xCol1_V2 + colWidth/2,
			Y:       math.Max(yBottomDeltaC, yGround) + 20,
			Content: fmt.Sprintf("V2:%.2f - V1:%.2f", cTo, cFrom),
			Presentation: svg.Presentation{
				Color:       "#B78103",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v1v2SubC)

		col1Label := &svg.Text{
			Name:    "Col 1 Label",
			X:       xCol1_V2 + colWidth/2,
			Y:       math.Max(yBottomDeltaC, yGround) + 40,
			Content: fmt.Sprintf("ΔC = %.2f", deltaC),
			Presentation: svg.Presentation{
				Color:       "#E65100",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col1Label)

		// Column 2: Net Performance Gain (α·ΔP)
		heightDeltaP := math.Max(math.Abs(alpha*deltaP)*scale, 50.0)
		yTopDeltaP := yGround - heightDeltaP // Top peak

		var linesP []string
		for _, p := range toPerformances {
			if diagram.AreQuantitativeElementsVisible {
				if alpha != 1.0 {
					linesP = append(linesP, fmt.Sprintf("V2: %s (%.2f · factor=%.2f)", p.Name, p.Strength, p.Strength*alpha))
				} else {
					linesP = append(linesP, fmt.Sprintf("V2: %s (%.2f)", p.Name, p.Strength))
				}
			} else {
				linesP = append(linesP, fmt.Sprintf("V2: %s", p.Name))
			}
		}
		for _, p := range fromPerformances {
			if diagram.AreQuantitativeElementsVisible {
				if alpha != 1.0 {
					linesP = append(linesP, fmt.Sprintf("V1: %s (%.2f · factor=%.2f)", p.Name, p.Strength, p.Strength*alpha))
				} else {
					linesP = append(linesP, fmt.Sprintf("V1: %s (%.2f)", p.Name, p.Strength))
				}
			} else {
				linesP = append(linesP, fmt.Sprintf("V1: %s", p.Name))
			}
		}
		var allPerfItems []any
		for _, p := range toPerformances {
			allPerfItems = append(allPerfItems, p)
		}
		for _, p := range fromPerformances {
			allPerfItems = append(allPerfItems, p)
		}
		renderDeltaRect(xCol2_V2, yTopDeltaP, heightDeltaP, "P", "#E8F5E9", "#2E7D32", "#1B5E20", linesP, allPerfItems)

		v1v2SubP := &svg.Text{
			Name:    "α·ΔP Sub-Label",
			X:       xCol2_V2 + colWidth/2,
			Y:       yGround + 20,
			Content: fmt.Sprintf("V2:%.2f - V1:%.2f", alpha*pTo, alpha*pFrom),
			Presentation: svg.Presentation{
				Color:       "#2E7D32",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v1v2SubP)

		col2Label := &svg.Text{
			Name:    "Col 2 Label",
			X:       xCol2_V2 + colWidth/2,
			Y:       yGround + 40,
			Content: fmt.Sprintf("α · ΔP = %.2f", alpha*deltaP),
			Presentation: svg.Presentation{
				Color:       "#2E7D32",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col2Label)

		// Column 3: Net Effort Investment (β·ΔE) dropping from peak
		heightDeltaE := math.Max(math.Abs(beta*deltaE)*scale, 50.0)
		yTopDeltaE := yTopDeltaP // drops from peak
		yBottomDeltaE := yTopDeltaE + heightDeltaE

		var linesE []string
		for _, e := range toEfforts {
			if diagram.AreQuantitativeElementsVisible {
				if beta != 1.0 {
					linesE = append(linesE, fmt.Sprintf("V2: %s (%.2f · factor=%.2f)", e.Name, e.Strength, e.Strength*beta))
				} else {
					linesE = append(linesE, fmt.Sprintf("V2: %s (%.2f)", e.Name, e.Strength))
				}
			} else {
				linesE = append(linesE, fmt.Sprintf("V2: %s", e.Name))
			}
		}
		for _, e := range fromEfforts {
			if diagram.AreQuantitativeElementsVisible {
				if beta != 1.0 {
					linesE = append(linesE, fmt.Sprintf("V1: %s (%.2f · factor=%.2f)", e.Name, e.Strength, e.Strength*beta))
				} else {
					linesE = append(linesE, fmt.Sprintf("V1: %s (%.2f)", e.Name, e.Strength))
				}
			} else {
				linesE = append(linesE, fmt.Sprintf("V1: %s", e.Name))
			}
		}
		var allEffItems []any
		for _, e := range toEfforts {
			allEffItems = append(allEffItems, e)
		}
		for _, e := range fromEfforts {
			allEffItems = append(allEffItems, e)
		}
		renderDeltaRect(xCol3_V2, yTopDeltaE, heightDeltaE, "E", "#E3F2FD", "#1976D2", "#0D47A1", linesE, allEffItems)

		v1v2SubE := &svg.Text{
			Name:    "β·ΔE Sub-Label",
			X:       xCol3_V2 + colWidth/2,
			Y:       math.Max(yBottomDeltaE, yGround) + 20,
			Content: fmt.Sprintf("V2:%.2f - V1:%.2f", beta*eTo, beta*eFrom),
			Presentation: svg.Presentation{
				Color:       "#1976D2",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, v1v2SubE)

		col3Label := &svg.Text{
			Name:    "Col 3 Label",
			X:       xCol3_V2 + colWidth/2,
			Y:       math.Max(yBottomDeltaE, yGround) + 40,
			Content: fmt.Sprintf("β · ΔE = %.2f", beta*deltaE),
			Presentation: svg.Presentation{
				Color:       "#1976D2",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, col3Label)

		// Guides and Indicators
		peakLine := &svg.Line{
			Name: "Peak Indicator Line",
			X1:   xCol2_V2,
			Y1:   yTopDeltaP,
			X2:   columnsRight,
			Y2:   yTopDeltaP,
			Presentation: svg.Presentation{
				Stroke:          "#388E3C",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, peakLine)

		peakText := &svg.Text{
			Name:    "Peak Label",
			X:       columnsRight + 10,
			Y:       yTopDeltaP + 4,
			Content: "α · ΔP peak",
			Presentation: svg.Presentation{
				Color:       "#388E3C",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, peakText)

		rhsLine := &svg.Line{
			Name: "RHS Level Line",
			X1:   xCol1_V2,
			Y1:   yBottomDeltaE,
			X2:   columnsRight,
			Y2:   yBottomDeltaE,
			Presentation: svg.Presentation{
				Stroke:          "#1565C0",
				StrokeWidth:     1.5,
				StrokeOpacity:   1.0,
				StrokeDashArray: "4 3",
			},
		}
		layer.Lines = append(layer.Lines, rhsLine)

		rhsText := &svg.Text{
			Name:    "RHS Label",
			X:       columnsRight + 10,
			Y:       yBottomDeltaE + 4,
			Content: fmt.Sprintf("α·ΔP - β·ΔE level (%.2f)", rhs),
			Presentation: svg.Presentation{
				Color:       "#1565C0",
				FillOpacity: 1.0,
				Stroke:      "transparent",
			},
		}
		layer.Texts = append(layer.Texts, rhsText)

		if deltaC <= 0 {
			yTipC_Indicator = yBottomDeltaC
		} else {
			yTipC_Indicator = yTopDeltaC
		}
		yRHS_Indicator = yBottomDeltaE
		indicatorX = xCol1_V2 + colWidth/2
	}

	diffColor := "#43A047"
	diffTextMsg := "Equilibrium (ΔC ≈ α·ΔP - β·ΔE)"
	if !isDelta {
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
		X1:   indicatorX,
		Y1:   yTipC_Indicator,
		X2:   indicatorX,
		Y2:   yRHS_Indicator,
		Presentation: svg.Presentation{
			Stroke:        diffColor,
			StrokeWidth:   2.5,
			StrokeOpacity: 1.0,
		},
	}
	layer.Lines = append(layer.Lines, diffLine)

	maxBottom := math.Max(math.Max(yTipC_Indicator, yRHS_Indicator), yGround)
	indicatorLabel := &svg.Text{
		Name:    "Indicator Message",
		X:       40,
		Y:       maxBottom + 65,
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

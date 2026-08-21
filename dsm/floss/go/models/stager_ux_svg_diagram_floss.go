package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.systemDiagramSvgStage.Reset()

	var diagramFloss *DiagramFloss
	{
		for diagramsystem_ := range *GetGongstructInstancesSet[DiagramFloss](stager.stage) {
			if diagramsystem_.IsChecked {
				diagramFloss = diagramsystem_
			}
		}
	}

	if diagramFloss == nil {
		stager.systemDiagramSvgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(diagramFloss)

	svg.StageBranch(stager.systemDiagramSvgStage, svgObject)
	stager.svgObjectDiagramFloss = svgObject
	stager.svgObjectDiagramFloss.OnUpdate = stager.onUpdateSVG

	stager.systemDiagramSvgStage.Commit()
}

func (stager *Stager) generateSvgObject(diagramFloss *DiagramFloss) *svg.SVG {

	svgObject := (&svg.SVG{Name: `SVG`})
	stager.diagramFloss = diagramFloss

	svgObject.OverrideWidth = true
	svgObject.OverriddenWidth = diagramFloss.Width
	svgObject.OverrideHeight = true
	svgObject.OverriddenHeight = diagramFloss.Height

	svgObject.Name = diagramFloss.Name
	svgObject.IsEditable = diagramFloss.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	svgObject.Layers = append(svgObject.Layers, layer)

	stager.drawSystemShapes(diagramFloss, layer)

	return svgObject
}

func (stager *Stager) drawSystemShapes(diagramFloss *DiagramFloss, layer *svg.Layer) {
	diagramFloss.map_System_Rect = make(map[*System]*svg.Rect)
	for _, systemShape := range diagramFloss.System_Shapes {
		if systemShape.IsHidden {
			continue
		}

		rect := svgRect(
			stager,
			diagramFloss,
			systemShape,
			layer)
		rect.RX = 3

		if len(rect.RectAnchoredTexts) > 0 {
			title := rect.RectAnchoredTexts[0]
			title.RectAnchorType = svg.RECT_TOP
			title.Y_Offset = +30
			title.FontWeight = "500"
			title.FontSize = "18px"
		}

		rect.Color = "#F8F9FA"
		rect.FillOpacity = 1.0
		rect.Stroke = "#E0E0E0"
		rect.StrokeWidth = 1.5

		rect.OnSelect = func() {
			stager.probeForm.FillUpFormFromGongstruct(systemShape.System, GetPointerToGongstructName[*System]())
		}
		rect.OnMove = onMoveRectElement(stager, systemShape, false)
		rect.OnResize = onResizeRectElement(stager, systemShape)
		diagramFloss.map_System_Rect[systemShape.System] = rect
	}
}

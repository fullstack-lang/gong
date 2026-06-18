package models

import (
	"log"
	"strings"

	"github.com/fullstack-lang/gong/dsm/scenario/go/icons"
	"github.com/fullstack-lang/gong/dsm/scenario/go/icons/path"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.svgStage.Reset()

	var diagram *Diagram
	for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
		if diagram_.IsChecked {
			diagram = diagram_
		}
	}

	if diagram == nil {
		stager.svgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(diagram)

	svg.StageBranch(stager.svgStage, svgObject)
	stager.svgObject = svgObject

	stager.svgStage.Commit()
}

// splitIntoLines takes a string of plain English words and punctuation, and a maximum length.
// It breaks down the string into substrings of words, where each substring has the longest length
// but a length less than or equal to lengthMax.
func splitIntoLines(inputStr string, lengthMax int) []string {
	// Split the input string into words.
	words := strings.Fields(inputStr)

	var lines []string
	var currentSubstring string

	// Iterate over each word.
	for _, word := range words {
		// Check if adding the current word exceeds the maximum length.
		if len(currentSubstring)+len(word)+1 > lengthMax {
			// If currentSubstring is not empty, append it to substrings.
			if currentSubstring != "" {
				lines = append(lines, strings.TrimSpace(currentSubstring))
				currentSubstring = ""
			}
		}

		// Add the word to the current substring.
		currentSubstring += word + " "

		// Handling the case where a single word is longer than lengthMax.
		if len(word) > lengthMax {
			lines = append(lines, strings.TrimSpace(currentSubstring))
			currentSubstring = ""
		}
	}

	// Append the last substring if it's not empty.
	if currentSubstring != "" {
		lines = append(lines, strings.TrimSpace(currentSubstring))
	}

	return lines
}

func (stager *Stager) formatText(rect *svg.Rect, yOffset float64) {
	if len(rect.RectAnchoredTexts) > 0 {
		numberOfPixelsPerCharacters := 7.5
		padding := 20.0 // 10px padding on each side
		maxChars := int((rect.Width - padding) / numberOfPixelsPerCharacters)
		if maxChars < 10 {
			maxChars = 10
		}
		lines := splitIntoLines(rect.Name, maxChars)
		content := ""
		for _, line := range lines {
			content = content + line + "\n"
		}
		text := rect.RectAnchoredTexts[0]
		text.Name = content
		text.Content = content
		text.X_Offset = 0
		text.Y_Offset = yOffset
		text.RectAnchorType = svg.RECT_TOP
		text.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		text.FontWeight = "normal"
		text.FontSize = "14"
		text.FontFamily = "Roboto, sans-serif"
		text.Color = "#333333"
	}
}

func (stager *Stager) appendArrows(rect *svg.Rect, direction DirectionType) {
	if direction != DIRECTION_UP && direction != DIRECTION_DOWN {
		return
	}

	rectAnchoredPath := new(svg.RectAnchoredPath)
	rect.IsScalingProportionally = false
	rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, rectAnchoredPath)

	if direction == DIRECTION_UP {
		paths, err := path.ExtractSVGPaths(icons.Parameter_right_Arrow_Up_Icon.SVG)
		if err == nil && len(paths) > 0 {
			rectAnchoredPath.Definition = paths[0]
		}
	}
	if direction == DIRECTION_DOWN {
		paths, err := path.ExtractSVGPaths(icons.Parameter_right_Arrow_Down_Icon.SVG)
		if err == nil && len(paths) > 0 {
			rectAnchoredPath.Definition = paths[0]
		}
	}
	rectAnchoredPath.Name = "Arrow"

	rectAnchoredPath.ScalePropotionnally = true
	rectAnchoredPath.AppliedScaling = 0.06 // Make arrow slightly larger

	rectAnchoredPath.Stroke = "#455A64" // Much darker stroke for better contrast
	rectAnchoredPath.StrokeWidth = 2.0
	rectAnchoredPath.StrokeOpacity = 1

	rectAnchoredPath.Color = "#90A4AE" // Darker fill color
	rectAnchoredPath.FillOpacity = 1.0

	distanceFromBorder := 10.0
	iconWidth := 25.0

	rectAnchoredPath.X_Offset = distanceFromBorder
	rectAnchoredPath.Y_Offset = -distanceFromBorder

	rectAnchoredPath.RectAnchorType = svg.RECT_BOTTOM_LEFT

	// shift the text on the right
	if len(rect.RectAnchoredTexts) > 0 {
		title := rect.RectAnchoredTexts[0]
		if rect.Width > (distanceFromBorder + iconWidth) {
			numberOfPixelsPerCharacters := 7.5
			padding := 20.0
			maxChars := int((rect.Width - (distanceFromBorder + iconWidth) - padding) / numberOfPixelsPerCharacters)
			if maxChars < 10 {
				maxChars = 10
			}
			lines := splitIntoLines(rect.Name, maxChars)
			content := ""
			for _, line := range lines {
				content = content + line + "\n"
			}
			title.Content = content
			title.Name = content
		}
		title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
	}
}

func (stager *Stager) generateSvgObject(diagram *Diagram) *svg.SVG {
	svgObject := (&svg.SVG{Name: `SVG`})
	stager.diagram = diagram

	svgObject.Name = diagram.Name
	svgObject.IsEditable = diagram.IsInDrawMode

	//
	// Layers
	//
	backgroundLayer := new(svg.Layer)
	backgroundLayer.Name = "Background Layer"
	svgObject.Layers = append(svgObject.Layers, backgroundLayer)

	backgroundRect := new(svg.Rect)
	backgroundRect.Name = "Background"
	backgroundRect.X = -100000
	backgroundRect.Y = -100000
	backgroundRect.Width = 200000
	backgroundRect.Height = 200000
	backgroundRect.Color = "white"
	backgroundRect.FillOpacity = 0.0
	backgroundRect.OnUpdate = func(frontRect *svg.Rect) {
		diagram.IsInDrawMode = !diagram.IsInDrawMode
		stager.stage.Commit()
	}
	backgroundLayer.Rects = append(backgroundLayer.Rects, backgroundRect)

	axisLayer := new(svg.Layer)
	axisLayer.Name = "Axis Layer"
	svgObject.Layers = append(svgObject.Layers, axisLayer)

	evolutionDirectionLayer := new(svg.Layer)
	evolutionDirectionLayer.Name = "Evolution Direction Layer"
	svgObject.Layers = append(svgObject.Layers, evolutionDirectionLayer)

	parameterLayer := new(svg.Layer)
	parameterLayer.Name = "Parameter Layer"
	svgObject.Layers = append(svgObject.Layers, parameterLayer)

	scenarioParameterLayer := new(svg.Layer)
	scenarioParameterLayer.Name = "Scenario Parameter Layer"
	svgObject.Layers = append(svgObject.Layers, scenarioParameterLayer)

	actorStateShapesLayer := new(svg.Layer)
	actorStateShapesLayer.Name = "Actor State Layer"
	svgObject.Layers = append(svgObject.Layers, actorStateShapesLayer)

	simAgentTrajectoryShapesLayer := new(svg.Layer)
	simAgentTrajectoryShapesLayer.Name = "Sim Agent Trajectory Layer"
	svgObject.Layers = append(svgObject.Layers, simAgentTrajectoryShapesLayer)

	//
	// Axes
	//
	stager.drawVerticalAxis(diagram, axisLayer)
	stager.drawHorizontalAxis(diagram, axisLayer)

	map_ActorStateShape_Rect := make(map[*ActorStateShape]*svg.Rect)
	map_ParameterShape_Rect := make(map[*ParameterShape]*svg.Rect)
	map_ScenarioParameterShape_Rect := make(map[*ParametersAggregateShape]*svg.Rect)

	//
	// Rectangles
	//
	for _, shape := range diagram.ActorStateShapes {
		if shape.GetIsHidden() {
			continue
		}
		rect := svgRect(stager, diagram, shape, actorStateShapesLayer)
		rect.RX = 8
		rect.Color = "#FFF0E0" // Soft peach/orange
		rect.Stroke = "#FFDDBB"
		rect.StrokeWidth = 1.5
		rect.FillOpacity = 1.0
		stager.formatText(rect, 20)

		if shape.ActorState.IsWithProbaility {
			rectAnchoredPath := new(svg.RectAnchoredPath)
			rect.IsScalingProportionally = false
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, rectAnchoredPath)

			switch shape.ActorState.Probability {
			case PERCENT_100:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.38,4.16 v 6.85 h 11.11 v -6.85 z m 0,11.35 v 6.68 h 11.11 v -6.68 z m 0,11.08 v 6.60 h 11.11 v -6.60 z m 0,11.05 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_75:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.38,15.51 v 6.68 h 11.11 v -6.68 z m 0,11.09 v 6.60 h 11.11 v -6.60 z m 0,11.04 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_50:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z m 1.39,26.6 v 6.6 h 11.11 v -6.6 z m 0,11.05 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_25:
				rectAnchoredPath.Definition = "M 5.56,0 c -3.08,0 -5.56,2.48 -5.56,5.56 v 42.97 c 0,3.08 2.48,5.56 5.56,5.56 h 13.89 c 3.08,0 5.56,-2.48 5.56,-5.56 v -42.97 c 0,-3.08 -2.48,-5.56 -5.56,-5.56 z M 5.56,2.70 h 13.89 c 1.54,0 2.78,1.24 2.78,2.78 v 43.06 c 0,1.54 -1.24,2.78 -2.78,2.78 H 5.56 c -1.54,0 -2.78,-1.24 -2.78,-2.78 v -43.06 c 0,-1.54 1.24,-2.78 2.78,-2.78 z M 6.94,40.34 v 6.67 h 11.11 v -6.67 z"
			case PERCENT_00:
				rectAnchoredPath.Definition = "M 4.99,0 c -2.81,0.28 -4.99,2.65 -4.99,5.53 v 42.95 c 0,3.08 2.51,5.53 5.59,5.53 h 13.88 c 3.08,0 5.53,-2.45 5.53,-5.53 v -42.95 c 0,-3.08 -2.46,-5.53 -5.53,-5.53 h -13.88 c -0.19,0 -0.41,-0.02 -0.60,0 z m 0.60,2.66 h 13.88 c 1.54,0 2.77,1.23 2.77,2.77 v 43.06 c 0,1.54 -1.23,2.77 -2.77,2.77 h -13.88 c -1.54,0 -2.82,-1.23 -2.82,-2.77 v -43.06 c 0,-1.54 1.28,-2.77 2.82,-2.77 z"
			}

			rectAnchoredPath.Name = "Percentage"

			rectAnchoredPath.Stroke = "#B0BEC5" // Soft blue-gray
			rectAnchoredPath.StrokeWidth = 1.5
			rectAnchoredPath.StrokeOpacity = 1

			rectAnchoredPath.Color = "#CFD8DC" // Lighter blue-gray
			rectAnchoredPath.FillOpacity = 0.8

			distanceFromBorder := 10.0
			iconWidth := 25.0

			rectAnchoredPath.X_Offset = distanceFromBorder
			rectAnchoredPath.Y_Offset = distanceFromBorder

			rectAnchoredPath.RectAnchorType = svg.RECT_TOP_LEFT

			// shift the text on the right
			if len(rect.RectAnchoredTexts) > 0 {
				title := rect.RectAnchoredTexts[0]
				if rect.Width > (distanceFromBorder + iconWidth) {
					numberOfPixelsPerCharacters := 7.5
					padding := 20.0
					maxChars := int((rect.Width - (distanceFromBorder + iconWidth) - padding) / numberOfPixelsPerCharacters)
					if maxChars < 10 {
						maxChars = 10
					}
					lines := splitIntoLines(rect.Name, maxChars)
					content := ""
					for _, line := range lines {
						content = content + line + "\n"
					}
					title.Content = content
					title.Name = content
				}
				title.X_Offset = (distanceFromBorder + iconWidth) / 2.0
			}
		}

		map_ActorStateShape_Rect[shape] = rect
	}

	for _, shape := range diagram.ParameterShapes {
		if shape.GetIsHidden() {
			continue
		}
		rect := svgRect(stager, diagram, shape, parameterLayer)
		rect.RX = 8
		rect.Color = "#E3F2FD" // Soft light blue
		rect.Stroke = "#BBDEFB"
		rect.StrokeWidth = 1.5
		rect.FillOpacity = 1.0
		stager.formatText(rect, 20)

		stager.appendArrows(rect, shape.Direction)

		map_ParameterShape_Rect[shape] = rect
	}

	for _, shape := range diagram.ScenarioParameterShapes {
		if shape.GetIsHidden() {
			continue
		}
		rect := svgRect(stager, diagram, shape, scenarioParameterLayer)
		rect.RX = 8
		rect.Color = "#E8F5E9" // Soft light green
		rect.Stroke = "#C8E6C9"
		rect.StrokeWidth = 1.5
		rect.FillOpacity = 1.0
		stager.formatText(rect, 20)

		stager.appendArrows(rect, shape.Direction)

		map_ScenarioParameterShape_Rect[shape] = rect
	}

	for _, shape := range diagram.EvolutionDirectionShapes {
		if shape.GetIsHidden() {
			continue
		}
		rect := svgRect(stager, diagram, shape, evolutionDirectionLayer)
		rect.RX = 8
		rect.Color = "#F5F5F5" // Soft light gray
		rect.Stroke = "#E0E0E0"
		rect.StrokeWidth = 1.5
		rect.FillOpacity = 1.0
		stager.formatText(rect, 20)
	}

	//
	// Links
	//
	stager.drawActorStateTransitionShapes(diagram, actorStateShapesLayer, map_ActorStateShape_Rect)
	stager.drawScenarioParameterToParameterLinks(diagram, scenarioParameterLayer, map_ScenarioParameterShape_Rect, map_ParameterShape_Rect)

	return svgObject
}

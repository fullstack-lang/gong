package models

import (
	"log"
	"math"
)

type SVG struct {
	Name   string
	Layers []*Layer

	// When the user draws a link between two rects, he ALT+RMB clicks on the source
	// rect than drag the cursor to the target rect.
	// A dotted line appear between the cursor start point and the cursor current point

	// Drawing indicates wether it draws the line
	DrawingState DrawingState

	// When the mouse is up on the end rect, the SVG backend will create the link
	// between both rects
	StartRect *Rect
	EndRect   *Rect

	IsEditable bool
}

// OnAfterUpdate, notice that rect == stagedRect
func (svg *SVG) OnAfterUpdate(stage *StageStruct, _, frontSVG *SVG) {

	log.Println("SVG, OnAfterUpdate", svg.Name)

	// if it is the end of a drawing state
	if svg.DrawingState == NOT_DRAWING_LINE && frontSVG.DrawingState == DRAWING_LINE {
		svg.DrawingState = frontSVG.DrawingState

		// let's create a new layer with a line in it that connects both rectangles
		layer := new(Layer).Stage(stage)
		layer.Name = "Line layer"
		svg.Layers = append(svg.Layers, layer)

		line := closestMidpoints(svg.StartRect, svg.EndRect).Stage(stage)
		line.Name = "Line connecting rect " + svg.StartRect.Name + " to " + svg.EndRect.Name
		line.Color = "olivedrab"
		line.Stroke = "olivedrab"
		line.StrokeWidth = 4
		layer.Lines = append(layer.Lines, line)

		stage.Commit()
	}
}

func closestMidpoints(r1, r2 *Rect) *Line {
	midpoints1 := [][2]float64{
		{r1.X + r1.Width/2, r1.Y},
		{r1.X + r1.Width, r1.Y + r1.Height/2},
		{r1.X + r1.Width/2, r1.Y + r1.Height},
		{r1.X, r1.Y + r1.Height/2},
	}

	midpoints2 := [][2]float64{
		{r2.X + r2.Width/2, r2.Y},
		{r2.X + r2.Width, r2.Y + r2.Height/2},
		{r2.X + r2.Width/2, r2.Y + r2.Height},
		{r2.X, r2.Y + r2.Height/2},
	}

	minDistance := math.MaxFloat64
	var line Line

	for _, p1 := range midpoints1 {
		for _, p2 := range midpoints2 {
			distance := math.Sqrt(math.Pow(p2[0]-p1[0], 2) + math.Pow(p2[1]-p1[1], 2))
			if distance < minDistance {
				minDistance = distance
				line = Line{
					X1: p1[0],
					Y1: p1[1],
					X2: p2[0],
					Y2: p2[1]}
			}
		}
	}

	return &line
}

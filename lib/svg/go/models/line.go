package models

import "log"

type Line struct {
	Name           string
	X1, Y1, X2, Y2 float64
	Presentation

	Animates []*Animate

	// when an end user interacts with the line,
	// it updates those fields
	MouseClickX, MouseClickY float64
}

func (line *Line) OnAfterUpdate(stage *StageStruct, _, frontLine *Line) {

	if line.MouseClickX != frontLine.MouseClickX ||
		line.MouseClickY != frontLine.MouseClickY {

		line.MouseClickX = frontLine.MouseClickX
		line.MouseClickY = frontLine.MouseClickY

		log.Println("Line, OnAfterUpdate, with new Mouse Click", line.Name)
	}
}

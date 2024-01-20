// generated code - do not edit

import { LineDB } from './line-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Line {

	static GONGSTRUCT_NAME = "Line"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X1: number = 0
	Y1: number = 0
	X2: number = 0
	Y2: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""
	MouseClickX: number = 0
	MouseClickY: number = 0

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyLineToLineDB(line: Line, lineDB: LineDB) {

	lineDB.CreatedAt = line.CreatedAt
	lineDB.DeletedAt = line.DeletedAt
	lineDB.ID = line.ID
	
	// insertion point for basic fields copy operations
	lineDB.Name = line.Name
	lineDB.X1 = line.X1
	lineDB.Y1 = line.Y1
	lineDB.X2 = line.X2
	lineDB.Y2 = line.Y2
	lineDB.Color = line.Color
	lineDB.FillOpacity = line.FillOpacity
	lineDB.Stroke = line.Stroke
	lineDB.StrokeWidth = line.StrokeWidth
	lineDB.StrokeDashArray = line.StrokeDashArray
	lineDB.StrokeDashArrayWhenSelected = line.StrokeDashArrayWhenSelected
	lineDB.Transform = line.Transform
	lineDB.MouseClickX = line.MouseClickX
	lineDB.MouseClickY = line.MouseClickY

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	lineDB.LinePointersEncoding.Animates = []
    for (let _animate of line.Animates) {
		lineDB.LinePointersEncoding.Animates.push(_animate.ID)
    }
	
}

export function CopyLineDBToLine(lineDB: LineDB, line: Line, frontRepo: FrontRepo) {

	line.CreatedAt = lineDB.CreatedAt
	line.DeletedAt = lineDB.DeletedAt
	line.ID = lineDB.ID
	
	// insertion point for basic fields copy operations
	line.Name = lineDB.Name
	line.X1 = lineDB.X1
	line.Y1 = lineDB.Y1
	line.X2 = lineDB.X2
	line.Y2 = lineDB.Y2
	line.Color = lineDB.Color
	line.FillOpacity = lineDB.FillOpacity
	line.Stroke = lineDB.Stroke
	line.StrokeWidth = lineDB.StrokeWidth
	line.StrokeDashArray = lineDB.StrokeDashArray
	line.StrokeDashArrayWhenSelected = lineDB.StrokeDashArrayWhenSelected
	line.Transform = lineDB.Transform
	line.MouseClickX = lineDB.MouseClickX
	line.MouseClickY = lineDB.MouseClickY

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	line.Animates = new Array<Animate>()
	for (let _id of lineDB.LinePointersEncoding.Animates) {
	  let _animate = frontRepo.Animates.get(_id)
	  if (_animate != undefined) {
		line.Animates.push(_animate!)
	  }
	}
}

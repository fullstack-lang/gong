// generated code - do not edit

import { EllipseDB } from './ellipse-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Ellipse {

	static GONGSTRUCT_NAME = "Ellipse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CX: number = 0
	CY: number = 0
	RX: number = 0
	RY: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyEllipseToEllipseDB(ellipse: Ellipse, ellipseDB: EllipseDB) {

	ellipseDB.CreatedAt = ellipse.CreatedAt
	ellipseDB.DeletedAt = ellipse.DeletedAt
	ellipseDB.ID = ellipse.ID

	// insertion point for basic fields copy operations
	ellipseDB.Name = ellipse.Name
	ellipseDB.CX = ellipse.CX
	ellipseDB.CY = ellipse.CY
	ellipseDB.RX = ellipse.RX
	ellipseDB.RY = ellipse.RY
	ellipseDB.Color = ellipse.Color
	ellipseDB.FillOpacity = ellipse.FillOpacity
	ellipseDB.Stroke = ellipse.Stroke
	ellipseDB.StrokeWidth = ellipse.StrokeWidth
	ellipseDB.StrokeDashArray = ellipse.StrokeDashArray
	ellipseDB.StrokeDashArrayWhenSelected = ellipse.StrokeDashArrayWhenSelected
	ellipseDB.Transform = ellipse.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	ellipseDB.EllipsePointersEncoding.Animates = []
	for (let _animate of ellipse.Animates) {
		ellipseDB.EllipsePointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyEllipseDBToEllipse update basic, pointers and slice of pointers fields of ellipse
// from respectively the basic fields and encoded fields of pointers and slices of pointers of ellipseDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyEllipseDBToEllipse(ellipseDB: EllipseDB, ellipse: Ellipse, frontRepo: FrontRepo) {

	ellipse.CreatedAt = ellipseDB.CreatedAt
	ellipse.DeletedAt = ellipseDB.DeletedAt
	ellipse.ID = ellipseDB.ID

	// insertion point for basic fields copy operations
	ellipse.Name = ellipseDB.Name
	ellipse.CX = ellipseDB.CX
	ellipse.CY = ellipseDB.CY
	ellipse.RX = ellipseDB.RX
	ellipse.RY = ellipseDB.RY
	ellipse.Color = ellipseDB.Color
	ellipse.FillOpacity = ellipseDB.FillOpacity
	ellipse.Stroke = ellipseDB.Stroke
	ellipse.StrokeWidth = ellipseDB.StrokeWidth
	ellipse.StrokeDashArray = ellipseDB.StrokeDashArray
	ellipse.StrokeDashArrayWhenSelected = ellipseDB.StrokeDashArrayWhenSelected
	ellipse.Transform = ellipseDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	ellipse.Animates = new Array<Animate>()
	for (let _id of ellipseDB.EllipsePointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			ellipse.Animates.push(_animate!)
		}
	}
}

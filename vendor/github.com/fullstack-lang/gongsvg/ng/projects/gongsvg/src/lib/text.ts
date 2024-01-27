// generated code - do not edit

import { TextDB } from './text-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Text {

	static GONGSTRUCT_NAME = "Text"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0
	Content: string = ""
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

export function CopyTextToTextDB(text: Text, textDB: TextDB) {

	textDB.CreatedAt = text.CreatedAt
	textDB.DeletedAt = text.DeletedAt
	textDB.ID = text.ID

	// insertion point for basic fields copy operations
	textDB.Name = text.Name
	textDB.X = text.X
	textDB.Y = text.Y
	textDB.Content = text.Content
	textDB.Color = text.Color
	textDB.FillOpacity = text.FillOpacity
	textDB.Stroke = text.Stroke
	textDB.StrokeWidth = text.StrokeWidth
	textDB.StrokeDashArray = text.StrokeDashArray
	textDB.StrokeDashArrayWhenSelected = text.StrokeDashArrayWhenSelected
	textDB.Transform = text.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	textDB.TextPointersEncoding.Animates = []
	for (let _animate of text.Animates) {
		textDB.TextPointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyTextDBToText update basic, pointers and slice of pointers fields of text
// from respectively the basic fields and encoded fields of pointers and slices of pointers of textDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTextDBToText(textDB: TextDB, text: Text, frontRepo: FrontRepo) {

	text.CreatedAt = textDB.CreatedAt
	text.DeletedAt = textDB.DeletedAt
	text.ID = textDB.ID

	// insertion point for basic fields copy operations
	text.Name = textDB.Name
	text.X = textDB.X
	text.Y = textDB.Y
	text.Content = textDB.Content
	text.Color = textDB.Color
	text.FillOpacity = textDB.FillOpacity
	text.Stroke = textDB.Stroke
	text.StrokeWidth = textDB.StrokeWidth
	text.StrokeDashArray = textDB.StrokeDashArray
	text.StrokeDashArrayWhenSelected = textDB.StrokeDashArrayWhenSelected
	text.Transform = textDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	text.Animates = new Array<Animate>()
	for (let _id of textDB.TextPointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			text.Animates.push(_animate!)
		}
	}
}

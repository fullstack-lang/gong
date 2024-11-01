// generated code - do not edit

import { TextAPI } from './text-api'
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
	StrokeOpacity: number = 0
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	StrokeDashArrayWhenSelected: string = ""
	Transform: string = ""

	// insertion point for pointers and slices of pointers declarations
	Animates: Array<Animate> = []
}

export function CopyTextToTextAPI(text: Text, textAPI: TextAPI) {

	textAPI.CreatedAt = text.CreatedAt
	textAPI.DeletedAt = text.DeletedAt
	textAPI.ID = text.ID

	// insertion point for basic fields copy operations
	textAPI.Name = text.Name
	textAPI.X = text.X
	textAPI.Y = text.Y
	textAPI.Content = text.Content
	textAPI.Color = text.Color
	textAPI.FillOpacity = text.FillOpacity
	textAPI.Stroke = text.Stroke
	textAPI.StrokeOpacity = text.StrokeOpacity
	textAPI.StrokeWidth = text.StrokeWidth
	textAPI.StrokeDashArray = text.StrokeDashArray
	textAPI.StrokeDashArrayWhenSelected = text.StrokeDashArrayWhenSelected
	textAPI.Transform = text.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	textAPI.TextPointersEncoding.Animates = []
	for (let _animate of text.Animates) {
		textAPI.TextPointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyTextAPIToText update basic, pointers and slice of pointers fields of text
// from respectively the basic fields and encoded fields of pointers and slices of pointers of textAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTextAPIToText(textAPI: TextAPI, text: Text, frontRepo: FrontRepo) {

	text.CreatedAt = textAPI.CreatedAt
	text.DeletedAt = textAPI.DeletedAt
	text.ID = textAPI.ID

	// insertion point for basic fields copy operations
	text.Name = textAPI.Name
	text.X = textAPI.X
	text.Y = textAPI.Y
	text.Content = textAPI.Content
	text.Color = textAPI.Color
	text.FillOpacity = textAPI.FillOpacity
	text.Stroke = textAPI.Stroke
	text.StrokeOpacity = textAPI.StrokeOpacity
	text.StrokeWidth = textAPI.StrokeWidth
	text.StrokeDashArray = textAPI.StrokeDashArray
	text.StrokeDashArrayWhenSelected = textAPI.StrokeDashArrayWhenSelected
	text.Transform = textAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(textAPI.TextPointersEncoding.Animates)) {
		console.error('Rects is not an array:', textAPI.TextPointersEncoding.Animates);
		return;
	}

	text.Animates = new Array<Animate>()
	for (let _id of textAPI.TextPointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			text.Animates.push(_animate!)
		}
	}
}

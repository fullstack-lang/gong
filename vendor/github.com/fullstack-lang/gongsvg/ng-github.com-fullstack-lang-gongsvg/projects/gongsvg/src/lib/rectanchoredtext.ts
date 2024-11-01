// generated code - do not edit

import { RectAnchoredTextAPI } from './rectanchoredtext-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RectAnchoredText {

	static GONGSTRUCT_NAME = "RectAnchoredText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	FontWeight: string = ""
	FontSize: number = 0
	FontStyle: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	TextAnchorType: string = ""
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

export function CopyRectAnchoredTextToRectAnchoredTextAPI(rectanchoredtext: RectAnchoredText, rectanchoredtextAPI: RectAnchoredTextAPI) {

	rectanchoredtextAPI.CreatedAt = rectanchoredtext.CreatedAt
	rectanchoredtextAPI.DeletedAt = rectanchoredtext.DeletedAt
	rectanchoredtextAPI.ID = rectanchoredtext.ID

	// insertion point for basic fields copy operations
	rectanchoredtextAPI.Name = rectanchoredtext.Name
	rectanchoredtextAPI.Content = rectanchoredtext.Content
	rectanchoredtextAPI.FontWeight = rectanchoredtext.FontWeight
	rectanchoredtextAPI.FontSize = rectanchoredtext.FontSize
	rectanchoredtextAPI.FontStyle = rectanchoredtext.FontStyle
	rectanchoredtextAPI.X_Offset = rectanchoredtext.X_Offset
	rectanchoredtextAPI.Y_Offset = rectanchoredtext.Y_Offset
	rectanchoredtextAPI.RectAnchorType = rectanchoredtext.RectAnchorType
	rectanchoredtextAPI.TextAnchorType = rectanchoredtext.TextAnchorType
	rectanchoredtextAPI.Color = rectanchoredtext.Color
	rectanchoredtextAPI.FillOpacity = rectanchoredtext.FillOpacity
	rectanchoredtextAPI.Stroke = rectanchoredtext.Stroke
	rectanchoredtextAPI.StrokeOpacity = rectanchoredtext.StrokeOpacity
	rectanchoredtextAPI.StrokeWidth = rectanchoredtext.StrokeWidth
	rectanchoredtextAPI.StrokeDashArray = rectanchoredtext.StrokeDashArray
	rectanchoredtextAPI.StrokeDashArrayWhenSelected = rectanchoredtext.StrokeDashArrayWhenSelected
	rectanchoredtextAPI.Transform = rectanchoredtext.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rectanchoredtextAPI.RectAnchoredTextPointersEncoding.Animates = []
	for (let _animate of rectanchoredtext.Animates) {
		rectanchoredtextAPI.RectAnchoredTextPointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyRectAnchoredTextAPIToRectAnchoredText update basic, pointers and slice of pointers fields of rectanchoredtext
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredtextAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredTextAPIToRectAnchoredText(rectanchoredtextAPI: RectAnchoredTextAPI, rectanchoredtext: RectAnchoredText, frontRepo: FrontRepo) {

	rectanchoredtext.CreatedAt = rectanchoredtextAPI.CreatedAt
	rectanchoredtext.DeletedAt = rectanchoredtextAPI.DeletedAt
	rectanchoredtext.ID = rectanchoredtextAPI.ID

	// insertion point for basic fields copy operations
	rectanchoredtext.Name = rectanchoredtextAPI.Name
	rectanchoredtext.Content = rectanchoredtextAPI.Content
	rectanchoredtext.FontWeight = rectanchoredtextAPI.FontWeight
	rectanchoredtext.FontSize = rectanchoredtextAPI.FontSize
	rectanchoredtext.FontStyle = rectanchoredtextAPI.FontStyle
	rectanchoredtext.X_Offset = rectanchoredtextAPI.X_Offset
	rectanchoredtext.Y_Offset = rectanchoredtextAPI.Y_Offset
	rectanchoredtext.RectAnchorType = rectanchoredtextAPI.RectAnchorType
	rectanchoredtext.TextAnchorType = rectanchoredtextAPI.TextAnchorType
	rectanchoredtext.Color = rectanchoredtextAPI.Color
	rectanchoredtext.FillOpacity = rectanchoredtextAPI.FillOpacity
	rectanchoredtext.Stroke = rectanchoredtextAPI.Stroke
	rectanchoredtext.StrokeOpacity = rectanchoredtextAPI.StrokeOpacity
	rectanchoredtext.StrokeWidth = rectanchoredtextAPI.StrokeWidth
	rectanchoredtext.StrokeDashArray = rectanchoredtextAPI.StrokeDashArray
	rectanchoredtext.StrokeDashArrayWhenSelected = rectanchoredtextAPI.StrokeDashArrayWhenSelected
	rectanchoredtext.Transform = rectanchoredtextAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(rectanchoredtextAPI.RectAnchoredTextPointersEncoding.Animates)) {
		console.error('Rects is not an array:', rectanchoredtextAPI.RectAnchoredTextPointersEncoding.Animates);
		return;
	}

	rectanchoredtext.Animates = new Array<Animate>()
	for (let _id of rectanchoredtextAPI.RectAnchoredTextPointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			rectanchoredtext.Animates.push(_animate!)
		}
	}
}

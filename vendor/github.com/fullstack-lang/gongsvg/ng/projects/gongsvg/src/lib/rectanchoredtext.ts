// generated code - do not edit

import { RectAnchoredTextDB } from './rectanchoredtext-db'
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
	X_Offset: number = 0
	Y_Offset: number = 0
	RectAnchorType: string = ""
	TextAnchorType: string = ""
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

export function CopyRectAnchoredTextToRectAnchoredTextDB(rectanchoredtext: RectAnchoredText, rectanchoredtextDB: RectAnchoredTextDB) {

	rectanchoredtextDB.CreatedAt = rectanchoredtext.CreatedAt
	rectanchoredtextDB.DeletedAt = rectanchoredtext.DeletedAt
	rectanchoredtextDB.ID = rectanchoredtext.ID

	// insertion point for basic fields copy operations
	rectanchoredtextDB.Name = rectanchoredtext.Name
	rectanchoredtextDB.Content = rectanchoredtext.Content
	rectanchoredtextDB.FontWeight = rectanchoredtext.FontWeight
	rectanchoredtextDB.FontSize = rectanchoredtext.FontSize
	rectanchoredtextDB.X_Offset = rectanchoredtext.X_Offset
	rectanchoredtextDB.Y_Offset = rectanchoredtext.Y_Offset
	rectanchoredtextDB.RectAnchorType = rectanchoredtext.RectAnchorType
	rectanchoredtextDB.TextAnchorType = rectanchoredtext.TextAnchorType
	rectanchoredtextDB.Color = rectanchoredtext.Color
	rectanchoredtextDB.FillOpacity = rectanchoredtext.FillOpacity
	rectanchoredtextDB.Stroke = rectanchoredtext.Stroke
	rectanchoredtextDB.StrokeWidth = rectanchoredtext.StrokeWidth
	rectanchoredtextDB.StrokeDashArray = rectanchoredtext.StrokeDashArray
	rectanchoredtextDB.StrokeDashArrayWhenSelected = rectanchoredtext.StrokeDashArrayWhenSelected
	rectanchoredtextDB.Transform = rectanchoredtext.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rectanchoredtextDB.RectAnchoredTextPointersEncoding.Animates = []
	for (let _animate of rectanchoredtext.Animates) {
		rectanchoredtextDB.RectAnchoredTextPointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyRectAnchoredTextDBToRectAnchoredText update basic, pointers and slice of pointers fields of rectanchoredtext
// from respectively the basic fields and encoded fields of pointers and slices of pointers of rectanchoredtextDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyRectAnchoredTextDBToRectAnchoredText(rectanchoredtextDB: RectAnchoredTextDB, rectanchoredtext: RectAnchoredText, frontRepo: FrontRepo) {

	rectanchoredtext.CreatedAt = rectanchoredtextDB.CreatedAt
	rectanchoredtext.DeletedAt = rectanchoredtextDB.DeletedAt
	rectanchoredtext.ID = rectanchoredtextDB.ID

	// insertion point for basic fields copy operations
	rectanchoredtext.Name = rectanchoredtextDB.Name
	rectanchoredtext.Content = rectanchoredtextDB.Content
	rectanchoredtext.FontWeight = rectanchoredtextDB.FontWeight
	rectanchoredtext.FontSize = rectanchoredtextDB.FontSize
	rectanchoredtext.X_Offset = rectanchoredtextDB.X_Offset
	rectanchoredtext.Y_Offset = rectanchoredtextDB.Y_Offset
	rectanchoredtext.RectAnchorType = rectanchoredtextDB.RectAnchorType
	rectanchoredtext.TextAnchorType = rectanchoredtextDB.TextAnchorType
	rectanchoredtext.Color = rectanchoredtextDB.Color
	rectanchoredtext.FillOpacity = rectanchoredtextDB.FillOpacity
	rectanchoredtext.Stroke = rectanchoredtextDB.Stroke
	rectanchoredtext.StrokeWidth = rectanchoredtextDB.StrokeWidth
	rectanchoredtext.StrokeDashArray = rectanchoredtextDB.StrokeDashArray
	rectanchoredtext.StrokeDashArrayWhenSelected = rectanchoredtextDB.StrokeDashArrayWhenSelected
	rectanchoredtext.Transform = rectanchoredtextDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	rectanchoredtext.Animates = new Array<Animate>()
	for (let _id of rectanchoredtextDB.RectAnchoredTextPointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			rectanchoredtext.Animates.push(_animate!)
		}
	}
}

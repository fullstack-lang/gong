// generated code - do not edit

import { LinkAnchoredTextAPI } from './linkanchoredtext-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Animate } from './animate'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAnchoredText {

	static GONGSTRUCT_NAME = "LinkAnchoredText"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""
	AutomaticLayout: boolean = false
	LinkAnchorType: string = ""
	X_Offset: number = 0
	Y_Offset: number = 0
	FontWeight: string = ""
	FontSize: string = ""
	LetterSpacing: string = ""
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

export function CopyLinkAnchoredTextToLinkAnchoredTextAPI(linkanchoredtext: LinkAnchoredText, linkanchoredtextAPI: LinkAnchoredTextAPI) {

	linkanchoredtextAPI.CreatedAt = linkanchoredtext.CreatedAt
	linkanchoredtextAPI.DeletedAt = linkanchoredtext.DeletedAt
	linkanchoredtextAPI.ID = linkanchoredtext.ID

	// insertion point for basic fields copy operations
	linkanchoredtextAPI.Name = linkanchoredtext.Name
	linkanchoredtextAPI.Content = linkanchoredtext.Content
	linkanchoredtextAPI.AutomaticLayout = linkanchoredtext.AutomaticLayout
	linkanchoredtextAPI.LinkAnchorType = linkanchoredtext.LinkAnchorType
	linkanchoredtextAPI.X_Offset = linkanchoredtext.X_Offset
	linkanchoredtextAPI.Y_Offset = linkanchoredtext.Y_Offset
	linkanchoredtextAPI.FontWeight = linkanchoredtext.FontWeight
	linkanchoredtextAPI.FontSize = linkanchoredtext.FontSize
	linkanchoredtextAPI.LetterSpacing = linkanchoredtext.LetterSpacing
	linkanchoredtextAPI.Color = linkanchoredtext.Color
	linkanchoredtextAPI.FillOpacity = linkanchoredtext.FillOpacity
	linkanchoredtextAPI.Stroke = linkanchoredtext.Stroke
	linkanchoredtextAPI.StrokeOpacity = linkanchoredtext.StrokeOpacity
	linkanchoredtextAPI.StrokeWidth = linkanchoredtext.StrokeWidth
	linkanchoredtextAPI.StrokeDashArray = linkanchoredtext.StrokeDashArray
	linkanchoredtextAPI.StrokeDashArrayWhenSelected = linkanchoredtext.StrokeDashArrayWhenSelected
	linkanchoredtextAPI.Transform = linkanchoredtext.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	linkanchoredtextAPI.LinkAnchoredTextPointersEncoding.Animates = []
	for (let _animate of linkanchoredtext.Animates) {
		linkanchoredtextAPI.LinkAnchoredTextPointersEncoding.Animates.push(_animate.ID)
	}

}

// CopyLinkAnchoredTextAPIToLinkAnchoredText update basic, pointers and slice of pointers fields of linkanchoredtext
// from respectively the basic fields and encoded fields of pointers and slices of pointers of linkanchoredtextAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLinkAnchoredTextAPIToLinkAnchoredText(linkanchoredtextAPI: LinkAnchoredTextAPI, linkanchoredtext: LinkAnchoredText, frontRepo: FrontRepo) {

	linkanchoredtext.CreatedAt = linkanchoredtextAPI.CreatedAt
	linkanchoredtext.DeletedAt = linkanchoredtextAPI.DeletedAt
	linkanchoredtext.ID = linkanchoredtextAPI.ID

	// insertion point for basic fields copy operations
	linkanchoredtext.Name = linkanchoredtextAPI.Name
	linkanchoredtext.Content = linkanchoredtextAPI.Content
	linkanchoredtext.AutomaticLayout = linkanchoredtextAPI.AutomaticLayout
	linkanchoredtext.LinkAnchorType = linkanchoredtextAPI.LinkAnchorType
	linkanchoredtext.X_Offset = linkanchoredtextAPI.X_Offset
	linkanchoredtext.Y_Offset = linkanchoredtextAPI.Y_Offset
	linkanchoredtext.FontWeight = linkanchoredtextAPI.FontWeight
	linkanchoredtext.FontSize = linkanchoredtextAPI.FontSize
	linkanchoredtext.LetterSpacing = linkanchoredtextAPI.LetterSpacing
	linkanchoredtext.Color = linkanchoredtextAPI.Color
	linkanchoredtext.FillOpacity = linkanchoredtextAPI.FillOpacity
	linkanchoredtext.Stroke = linkanchoredtextAPI.Stroke
	linkanchoredtext.StrokeOpacity = linkanchoredtextAPI.StrokeOpacity
	linkanchoredtext.StrokeWidth = linkanchoredtextAPI.StrokeWidth
	linkanchoredtext.StrokeDashArray = linkanchoredtextAPI.StrokeDashArray
	linkanchoredtext.StrokeDashArrayWhenSelected = linkanchoredtextAPI.StrokeDashArrayWhenSelected
	linkanchoredtext.Transform = linkanchoredtextAPI.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(linkanchoredtextAPI.LinkAnchoredTextPointersEncoding.Animates)) {
		console.error('Rects is not an array:', linkanchoredtextAPI.LinkAnchoredTextPointersEncoding.Animates);
		return;
	}

	linkanchoredtext.Animates = new Array<Animate>()
	for (let _id of linkanchoredtextAPI.LinkAnchoredTextPointersEncoding.Animates) {
		let _animate = frontRepo.map_ID_Animate.get(_id)
		if (_animate != undefined) {
			linkanchoredtext.Animates.push(_animate!)
		}
	}
}

// generated code - do not edit

import { LinkAnchoredTextDB } from './linkanchoredtext-db'
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

export function CopyLinkAnchoredTextToLinkAnchoredTextDB(linkanchoredtext: LinkAnchoredText, linkanchoredtextDB: LinkAnchoredTextDB) {

	linkanchoredtextDB.CreatedAt = linkanchoredtext.CreatedAt
	linkanchoredtextDB.DeletedAt = linkanchoredtext.DeletedAt
	linkanchoredtextDB.ID = linkanchoredtext.ID
	
	// insertion point for basic fields copy operations
	linkanchoredtextDB.Name = linkanchoredtext.Name
	linkanchoredtextDB.Content = linkanchoredtext.Content
	linkanchoredtextDB.AutomaticLayout = linkanchoredtext.AutomaticLayout
	linkanchoredtextDB.LinkAnchorType = linkanchoredtext.LinkAnchorType
	linkanchoredtextDB.X_Offset = linkanchoredtext.X_Offset
	linkanchoredtextDB.Y_Offset = linkanchoredtext.Y_Offset
	linkanchoredtextDB.FontWeight = linkanchoredtext.FontWeight
	linkanchoredtextDB.Color = linkanchoredtext.Color
	linkanchoredtextDB.FillOpacity = linkanchoredtext.FillOpacity
	linkanchoredtextDB.Stroke = linkanchoredtext.Stroke
	linkanchoredtextDB.StrokeWidth = linkanchoredtext.StrokeWidth
	linkanchoredtextDB.StrokeDashArray = linkanchoredtext.StrokeDashArray
	linkanchoredtextDB.StrokeDashArrayWhenSelected = linkanchoredtext.StrokeDashArrayWhenSelected
	linkanchoredtextDB.Transform = linkanchoredtext.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	linkanchoredtextDB.LinkAnchoredTextPointersEncoding.Animates = []
    for (let _animate of linkanchoredtext.Animates) {
		linkanchoredtextDB.LinkAnchoredTextPointersEncoding.Animates.push(_animate.ID)
    }
	
}

export function CopyLinkAnchoredTextDBToLinkAnchoredText(linkanchoredtextDB: LinkAnchoredTextDB, linkanchoredtext: LinkAnchoredText, frontRepo: FrontRepo) {

	linkanchoredtext.CreatedAt = linkanchoredtextDB.CreatedAt
	linkanchoredtext.DeletedAt = linkanchoredtextDB.DeletedAt
	linkanchoredtext.ID = linkanchoredtextDB.ID
	
	// insertion point for basic fields copy operations
	linkanchoredtext.Name = linkanchoredtextDB.Name
	linkanchoredtext.Content = linkanchoredtextDB.Content
	linkanchoredtext.AutomaticLayout = linkanchoredtextDB.AutomaticLayout
	linkanchoredtext.LinkAnchorType = linkanchoredtextDB.LinkAnchorType
	linkanchoredtext.X_Offset = linkanchoredtextDB.X_Offset
	linkanchoredtext.Y_Offset = linkanchoredtextDB.Y_Offset
	linkanchoredtext.FontWeight = linkanchoredtextDB.FontWeight
	linkanchoredtext.Color = linkanchoredtextDB.Color
	linkanchoredtext.FillOpacity = linkanchoredtextDB.FillOpacity
	linkanchoredtext.Stroke = linkanchoredtextDB.Stroke
	linkanchoredtext.StrokeWidth = linkanchoredtextDB.StrokeWidth
	linkanchoredtext.StrokeDashArray = linkanchoredtextDB.StrokeDashArray
	linkanchoredtext.StrokeDashArrayWhenSelected = linkanchoredtextDB.StrokeDashArrayWhenSelected
	linkanchoredtext.Transform = linkanchoredtextDB.Transform

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	linkanchoredtext.Animates = new Array<Animate>()
	for (let _id of linkanchoredtextDB.LinkAnchoredTextPointersEncoding.Animates) {
	  let _animate = frontRepo.Animates.get(_id)
	  if (_animate != undefined) {
		linkanchoredtext.Animates.push(_animate!)
	  }
	}
}

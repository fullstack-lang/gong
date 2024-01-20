// generated code - do not edit

import { SVGIconDB } from './svgicon-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGIcon {

	static GONGSTRUCT_NAME = "SVGIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySVGIconToSVGIconDB(svgicon: SVGIcon, svgiconDB: SVGIconDB) {

	svgiconDB.CreatedAt = svgicon.CreatedAt
	svgiconDB.DeletedAt = svgicon.DeletedAt
	svgiconDB.ID = svgicon.ID
	
	// insertion point for basic fields copy operations
	svgiconDB.Name = svgicon.Name
	svgiconDB.SVG = svgicon.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopySVGIconDBToSVGIcon(svgiconDB: SVGIconDB, svgicon: SVGIcon, frontRepo: FrontRepo) {

	svgicon.CreatedAt = svgiconDB.CreatedAt
	svgicon.DeletedAt = svgiconDB.DeletedAt
	svgicon.ID = svgiconDB.ID
	
	// insertion point for basic fields copy operations
	svgicon.Name = svgiconDB.Name
	svgicon.SVG = svgiconDB.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

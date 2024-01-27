// generated code - do not edit

import { AnimateDB } from './animate-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Animate {

	static GONGSTRUCT_NAME = "Animate"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	AttributeName: string = ""
	Values: string = ""
	Dur: string = ""
	RepeatCount: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyAnimateToAnimateDB(animate: Animate, animateDB: AnimateDB) {

	animateDB.CreatedAt = animate.CreatedAt
	animateDB.DeletedAt = animate.DeletedAt
	animateDB.ID = animate.ID

	// insertion point for basic fields copy operations
	animateDB.Name = animate.Name
	animateDB.AttributeName = animate.AttributeName
	animateDB.Values = animate.Values
	animateDB.Dur = animate.Dur
	animateDB.RepeatCount = animate.RepeatCount

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAnimateDBToAnimate update basic, pointers and slice of pointers fields of animate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of animateDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAnimateDBToAnimate(animateDB: AnimateDB, animate: Animate, frontRepo: FrontRepo) {

	animate.CreatedAt = animateDB.CreatedAt
	animate.DeletedAt = animateDB.DeletedAt
	animate.ID = animateDB.ID

	// insertion point for basic fields copy operations
	animate.Name = animateDB.Name
	animate.AttributeName = animateDB.AttributeName
	animate.Values = animateDB.Values
	animate.Dur = animateDB.Dur
	animate.RepeatCount = animateDB.RepeatCount

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

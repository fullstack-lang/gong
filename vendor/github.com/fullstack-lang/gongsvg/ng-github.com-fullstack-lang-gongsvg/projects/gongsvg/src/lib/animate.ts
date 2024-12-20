// generated code - do not edit

import { AnimateAPI } from './animate-api'
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
	From: string = ""
	To: string = ""
	Dur: string = ""
	RepeatCount: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyAnimateToAnimateAPI(animate: Animate, animateAPI: AnimateAPI) {

	animateAPI.CreatedAt = animate.CreatedAt
	animateAPI.DeletedAt = animate.DeletedAt
	animateAPI.ID = animate.ID

	// insertion point for basic fields copy operations
	animateAPI.Name = animate.Name
	animateAPI.AttributeName = animate.AttributeName
	animateAPI.Values = animate.Values
	animateAPI.From = animate.From
	animateAPI.To = animate.To
	animateAPI.Dur = animate.Dur
	animateAPI.RepeatCount = animate.RepeatCount

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyAnimateAPIToAnimate update basic, pointers and slice of pointers fields of animate
// from respectively the basic fields and encoded fields of pointers and slices of pointers of animateAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAnimateAPIToAnimate(animateAPI: AnimateAPI, animate: Animate, frontRepo: FrontRepo) {

	animate.CreatedAt = animateAPI.CreatedAt
	animate.DeletedAt = animateAPI.DeletedAt
	animate.ID = animateAPI.ID

	// insertion point for basic fields copy operations
	animate.Name = animateAPI.Name
	animate.AttributeName = animateAPI.AttributeName
	animate.Values = animateAPI.Values
	animate.From = animateAPI.From
	animate.To = animateAPI.To
	animate.Dur = animateAPI.Dur
	animate.RepeatCount = animateAPI.RepeatCount

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

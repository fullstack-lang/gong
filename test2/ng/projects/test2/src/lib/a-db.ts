// insertion point for imports
import { BDB } from './b-db'
import { FrontRepo } from './front-repo.service'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ADB {

	static GONGSTRUCT_NAME = "A"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NumberField: number = 0

	// insertion point for pointers and slices of pointers declarations
	B?: BDB

	Bs: Array<BDB> = []

	APointersEncoding: APointersEncoding = new APointersEncoding

	redeemPointers(frontRepo: FrontRepo) {
		// insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
		// insertion point for pointer field B redeeming
		{
			let _b = frontRepo.Bs.get(this.APointersEncoding.BID.Int64)
			if (_b) {
				this.B = _b
			}
		}

		// insertion point for redeeming ONE-MANY associations
		this.Bs = new Array<BDB>()
		for (let _bId of this.APointersEncoding.Bs) {
			let _b = frontRepo.Bs.get(_bId)
			this.Bs.push(_b!)
		}
	}

	toto() {
		console.log("toto")
	}
}

export class APointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	BID: NullInt64 = new NullInt64 // if pointer is null, B.ID = 0

	Bs: number[] = []
}

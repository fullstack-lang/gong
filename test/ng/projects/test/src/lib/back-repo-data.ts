// generated code - do not edit

//insertion point for imports
import { AstructDB } from './astruct-db'

import { AstructBstruct2UseDB } from './astructbstruct2use-db'

import { AstructBstructUseDB } from './astructbstructuse-db'

import { BstructDB } from './bstruct-db'

import { DstructDB } from './dstruct-db'


export class BackRepoData {
	// insertion point for declarations
	AstructDBs = new Array<AstructDB>()

	AstructBstruct2UseDBs = new Array<AstructBstruct2UseDB>()

	AstructBstructUseDBs = new Array<AstructBstructUseDB>()

	BstructDBs = new Array<BstructDB>()

	DstructDBs = new Array<DstructDB>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AstructDBs = data?.AstructDBs || [];

		this.AstructBstruct2UseDBs = data?.AstructBstruct2UseDBs || [];

		this.AstructBstructUseDBs = data?.AstructBstructUseDBs || [];

		this.BstructDBs = data?.BstructDBs || [];

		this.DstructDBs = data?.DstructDBs || [];

	}

}
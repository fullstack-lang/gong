// generated code - do not edit

//insertion point for imports
import { AstructAPI } from './astruct-api'

import { AstructBstruct2UseAPI } from './astructbstruct2use-api'

import { AstructBstructUseAPI } from './astructbstructuse-api'

import { BstructAPI } from './bstruct-api'

import { DstructAPI } from './dstruct-api'

import { GstructAPI } from './gstruct-api'


export class BackRepoData {
	// insertion point for declarations
	AstructAPIs = new Array<AstructAPI>()

	AstructBstruct2UseAPIs = new Array<AstructBstruct2UseAPI>()

	AstructBstructUseAPIs = new Array<AstructBstructUseAPI>()

	BstructAPIs = new Array<BstructAPI>()

	DstructAPIs = new Array<DstructAPI>()

	GstructAPIs = new Array<GstructAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AstructAPIs = data?.AstructAPIs || [];

		this.AstructBstruct2UseAPIs = data?.AstructBstruct2UseAPIs || [];

		this.AstructBstructUseAPIs = data?.AstructBstructUseAPIs || [];

		this.BstructAPIs = data?.BstructAPIs || [];

		this.DstructAPIs = data?.DstructAPIs || [];

		this.GstructAPIs = data?.GstructAPIs || [];

	}

}
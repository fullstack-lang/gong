// generated code - do not edit

//insertion point for imports
import { AAPI } from './a-api'

import { BAPI } from './b-api'


export class BackRepoData {
	// insertion point for declarations
	AAPIs = new Array<AAPI>()

	BAPIs = new Array<BAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AAPIs = data?.AAPIs || [];

		this.BAPIs = data?.BAPIs || [];

	}

}
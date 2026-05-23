// generated code - do not edit

//insertion point for imports
import { AstructAPI } from './astruct-api'


export class BackRepoData {
	// insertion point for declarations
	AstructAPIs = new Array<AstructAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AstructAPIs = data?.AstructAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
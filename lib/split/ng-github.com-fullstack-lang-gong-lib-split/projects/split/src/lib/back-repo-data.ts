// generated code - do not edit

//insertion point for imports
import { SplitAreaAPI } from './splitarea-api'


export class BackRepoData {
	// insertion point for declarations
	SplitAreaAPIs = new Array<SplitAreaAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.SplitAreaAPIs = data?.SplitAreaAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
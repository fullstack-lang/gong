// generated code - do not edit

//insertion point for imports
import { CursorAPI } from './cursor-api'


export class BackRepoData {
	// insertion point for declarations
	CursorAPIs = new Array<CursorAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CursorAPIs = data?.CursorAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
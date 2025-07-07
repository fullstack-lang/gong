// generated code - do not edit

//insertion point for imports
import { ContentAPI } from './content-api'


export class BackRepoData {
	// insertion point for declarations
	ContentAPIs = new Array<ContentAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.ContentAPIs = data?.ContentAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
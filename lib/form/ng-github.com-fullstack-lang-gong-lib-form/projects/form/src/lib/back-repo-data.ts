// generated code - do not edit

//insertion point for imports
import { Form2API } from './form2-api'


export class BackRepoData {
	// insertion point for declarations
	Form2APIs = new Array<Form2API>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.Form2APIs = data?.Form2APIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
// generated code - do not edit

//insertion point for imports
import { ButtonAPI } from './button-api'

import { GroupAPI } from './group-api'

import { LayoutAPI } from './layout-api'


export class BackRepoData {
	// insertion point for declarations
	ButtonAPIs = new Array<ButtonAPI>()

	GroupAPIs = new Array<GroupAPI>()

	LayoutAPIs = new Array<LayoutAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.ButtonAPIs = data?.ButtonAPIs || [];

		this.GroupAPIs = data?.GroupAPIs || [];

		this.LayoutAPIs = data?.LayoutAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
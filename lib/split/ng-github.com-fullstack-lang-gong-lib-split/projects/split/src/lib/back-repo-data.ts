// generated code - do not edit

//insertion point for imports
import { AsSplitAPI } from './assplit-api'

import { AsSplitAreaAPI } from './assplitarea-api'

import { FormAPI } from './form-api'

import { TableAPI } from './table-api'

import { TreeAPI } from './tree-api'

import { ViewAPI } from './view-api'


export class BackRepoData {
	// insertion point for declarations
	AsSplitAPIs = new Array<AsSplitAPI>()

	AsSplitAreaAPIs = new Array<AsSplitAreaAPI>()

	FormAPIs = new Array<FormAPI>()

	TableAPIs = new Array<TableAPI>()

	TreeAPIs = new Array<TreeAPI>()

	ViewAPIs = new Array<ViewAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AsSplitAPIs = data?.AsSplitAPIs || [];

		this.AsSplitAreaAPIs = data?.AsSplitAreaAPIs || [];

		this.FormAPIs = data?.FormAPIs || [];

		this.TableAPIs = data?.TableAPIs || [];

		this.TreeAPIs = data?.TreeAPIs || [];

		this.ViewAPIs = data?.ViewAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
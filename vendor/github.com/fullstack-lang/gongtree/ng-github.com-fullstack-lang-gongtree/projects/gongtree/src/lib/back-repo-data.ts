// generated code - do not edit

//insertion point for imports
import { ButtonAPI } from './button-api'

import { NodeAPI } from './node-api'

import { SVGIconAPI } from './svgicon-api'

import { TreeAPI } from './tree-api'


export class BackRepoData {
	// insertion point for declarations
	ButtonAPIs = new Array<ButtonAPI>()

	NodeAPIs = new Array<NodeAPI>()

	SVGIconAPIs = new Array<SVGIconAPI>()

	TreeAPIs = new Array<TreeAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.ButtonAPIs = data?.ButtonAPIs || [];

		this.NodeAPIs = data?.NodeAPIs || [];

		this.SVGIconAPIs = data?.SVGIconAPIs || [];

		this.TreeAPIs = data?.TreeAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
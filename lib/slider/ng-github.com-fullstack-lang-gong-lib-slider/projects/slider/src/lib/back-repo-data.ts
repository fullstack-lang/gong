// generated code - do not edit

//insertion point for imports
import { CheckboxAPI } from './checkbox-api'

import { GroupAPI } from './group-api'

import { LayoutAPI } from './layout-api'

import { SliderAPI } from './slider-api'


export class BackRepoData {
	// insertion point for declarations
	CheckboxAPIs = new Array<CheckboxAPI>()

	GroupAPIs = new Array<GroupAPI>()

	LayoutAPIs = new Array<LayoutAPI>()

	SliderAPIs = new Array<SliderAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CheckboxAPIs = data?.CheckboxAPIs || [];

		this.GroupAPIs = data?.GroupAPIs || [];

		this.LayoutAPIs = data?.LayoutAPIs || [];

		this.SliderAPIs = data?.SliderAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
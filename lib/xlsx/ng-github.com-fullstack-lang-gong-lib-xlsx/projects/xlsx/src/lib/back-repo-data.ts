// generated code - do not edit

//insertion point for imports
import { DisplaySelectionAPI } from './displayselection-api'

import { XLCellAPI } from './xlcell-api'

import { XLFileAPI } from './xlfile-api'

import { XLRowAPI } from './xlrow-api'

import { XLSheetAPI } from './xlsheet-api'


export class BackRepoData {
	// insertion point for declarations
	DisplaySelectionAPIs = new Array<DisplaySelectionAPI>()

	XLCellAPIs = new Array<XLCellAPI>()

	XLFileAPIs = new Array<XLFileAPI>()

	XLRowAPIs = new Array<XLRowAPI>()

	XLSheetAPIs = new Array<XLSheetAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.DisplaySelectionAPIs = data?.DisplaySelectionAPIs || [];

		this.XLCellAPIs = data?.XLCellAPIs || [];

		this.XLFileAPIs = data?.XLFileAPIs || [];

		this.XLRowAPIs = data?.XLRowAPIs || [];

		this.XLSheetAPIs = data?.XLSheetAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
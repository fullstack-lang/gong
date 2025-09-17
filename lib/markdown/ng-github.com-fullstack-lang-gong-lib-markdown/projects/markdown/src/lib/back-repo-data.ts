// generated code - do not edit

//insertion point for imports
import { ContentAPI } from './content-api'

import { JpgImageAPI } from './jpgimage-api'

import { PngImageAPI } from './pngimage-api'

import { SvgImageAPI } from './svgimage-api'


export class BackRepoData {
	// insertion point for declarations
	ContentAPIs = new Array<ContentAPI>()

	JpgImageAPIs = new Array<JpgImageAPI>()

	PngImageAPIs = new Array<PngImageAPI>()

	SvgImageAPIs = new Array<SvgImageAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.ContentAPIs = data?.ContentAPIs || [];

		this.JpgImageAPIs = data?.JpgImageAPIs || [];

		this.PngImageAPIs = data?.PngImageAPIs || [];

		this.SvgImageAPIs = data?.SvgImageAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
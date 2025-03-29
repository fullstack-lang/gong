// generated code - do not edit

//insertion point for imports
import { AsSplitAPI } from './assplit-api'

import { AsSplitAreaAPI } from './assplitarea-api'

import { ButtonAPI } from './button-api'

import { CursorAPI } from './cursor-api'

import { DocAPI } from './doc-api'

import { FormAPI } from './form-api'

import { LoadAPI } from './load-api'

import { SliderAPI } from './slider-api'

import { SplitAPI } from './split-api'

import { SvgAPI } from './svg-api'

import { TableAPI } from './table-api'

import { ToneAPI } from './tone-api'

import { TreeAPI } from './tree-api'

import { ViewAPI } from './view-api'

import { XlsxAPI } from './xlsx-api'


export class BackRepoData {
	// insertion point for declarations
	AsSplitAPIs = new Array<AsSplitAPI>()

	AsSplitAreaAPIs = new Array<AsSplitAreaAPI>()

	ButtonAPIs = new Array<ButtonAPI>()

	CursorAPIs = new Array<CursorAPI>()

	DocAPIs = new Array<DocAPI>()

	FormAPIs = new Array<FormAPI>()

	LoadAPIs = new Array<LoadAPI>()

	SliderAPIs = new Array<SliderAPI>()

	SplitAPIs = new Array<SplitAPI>()

	SvgAPIs = new Array<SvgAPI>()

	TableAPIs = new Array<TableAPI>()

	ToneAPIs = new Array<ToneAPI>()

	TreeAPIs = new Array<TreeAPI>()

	ViewAPIs = new Array<ViewAPI>()

	XlsxAPIs = new Array<XlsxAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AsSplitAPIs = data?.AsSplitAPIs || [];

		this.AsSplitAreaAPIs = data?.AsSplitAreaAPIs || [];

		this.ButtonAPIs = data?.ButtonAPIs || [];

		this.CursorAPIs = data?.CursorAPIs || [];

		this.DocAPIs = data?.DocAPIs || [];

		this.FormAPIs = data?.FormAPIs || [];

		this.LoadAPIs = data?.LoadAPIs || [];

		this.SliderAPIs = data?.SliderAPIs || [];

		this.SplitAPIs = data?.SplitAPIs || [];

		this.SvgAPIs = data?.SvgAPIs || [];

		this.TableAPIs = data?.TableAPIs || [];

		this.ToneAPIs = data?.ToneAPIs || [];

		this.TreeAPIs = data?.TreeAPIs || [];

		this.ViewAPIs = data?.ViewAPIs || [];

		this.XlsxAPIs = data?.XlsxAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
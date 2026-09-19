// generated code - do not edit

//insertion point for imports
import { AsSplitAPI } from './assplit-api'

import { AsSplitAreaAPI } from './assplitarea-api'

import { ButtonAPI } from './button-api'

import { FavIconAPI } from './favicon-api'

import { FormAPI } from './form-api'

import { LoadAPI } from './load-api'

import { LogoOnTheLeftAPI } from './logoontheleft-api'

import { LogoOnTheRightAPI } from './logoontheright-api'

import { SplitAPI } from './split-api'

import { SvgAPI } from './svg-api'

import { TableAPI } from './table-api'

import { TitleAPI } from './title-api'

import { TreeAPI } from './tree-api'

import { ViewAPI } from './view-api'


export class BackRepoData {
	// insertion point for declarations
	AsSplitAPIs = new Array<AsSplitAPI>()

	AsSplitAreaAPIs = new Array<AsSplitAreaAPI>()

	ButtonAPIs = new Array<ButtonAPI>()

	FavIconAPIs = new Array<FavIconAPI>()

	FormAPIs = new Array<FormAPI>()

	LoadAPIs = new Array<LoadAPI>()

	LogoOnTheLeftAPIs = new Array<LogoOnTheLeftAPI>()

	LogoOnTheRightAPIs = new Array<LogoOnTheRightAPI>()

	SplitAPIs = new Array<SplitAPI>()

	SvgAPIs = new Array<SvgAPI>()

	TableAPIs = new Array<TableAPI>()

	TitleAPIs = new Array<TitleAPI>()

	TreeAPIs = new Array<TreeAPI>()

	ViewAPIs = new Array<ViewAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AsSplitAPIs = data?.AsSplitAPIs || [];

		this.AsSplitAreaAPIs = data?.AsSplitAreaAPIs || [];

		this.ButtonAPIs = data?.ButtonAPIs || [];

		this.FavIconAPIs = data?.FavIconAPIs || [];

		this.FormAPIs = data?.FormAPIs || [];

		this.LoadAPIs = data?.LoadAPIs || [];

		this.LogoOnTheLeftAPIs = data?.LogoOnTheLeftAPIs || [];

		this.LogoOnTheRightAPIs = data?.LogoOnTheRightAPIs || [];

		this.SplitAPIs = data?.SplitAPIs || [];

		this.SvgAPIs = data?.SvgAPIs || [];

		this.TableAPIs = data?.TableAPIs || [];

		this.TitleAPIs = data?.TitleAPIs || [];

		this.TreeAPIs = data?.TreeAPIs || [];

		this.ViewAPIs = data?.ViewAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
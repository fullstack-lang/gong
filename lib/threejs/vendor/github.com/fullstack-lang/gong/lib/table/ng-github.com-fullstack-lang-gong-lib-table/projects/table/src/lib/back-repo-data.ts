// generated code - do not edit

//insertion point for imports
import { ButtonAPI } from './button-api'

import { CellAPI } from './cell-api'

import { CellBooleanAPI } from './cellboolean-api'

import { CellFloat64API } from './cellfloat64-api'

import { CellIconAPI } from './cellicon-api'

import { CellIntAPI } from './cellint-api'

import { CellStringAPI } from './cellstring-api'

import { DisplayedColumnAPI } from './displayedcolumn-api'

import { RowAPI } from './row-api'

import { SVGIconAPI } from './svgicon-api'

import { TableAPI } from './table-api'


export class BackRepoData {
	// insertion point for declarations
	ButtonAPIs = new Array<ButtonAPI>()

	CellAPIs = new Array<CellAPI>()

	CellBooleanAPIs = new Array<CellBooleanAPI>()

	CellFloat64APIs = new Array<CellFloat64API>()

	CellIconAPIs = new Array<CellIconAPI>()

	CellIntAPIs = new Array<CellIntAPI>()

	CellStringAPIs = new Array<CellStringAPI>()

	DisplayedColumnAPIs = new Array<DisplayedColumnAPI>()

	RowAPIs = new Array<RowAPI>()

	SVGIconAPIs = new Array<SVGIconAPI>()

	TableAPIs = new Array<TableAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.ButtonAPIs = data?.ButtonAPIs || [];

		this.CellAPIs = data?.CellAPIs || [];

		this.CellBooleanAPIs = data?.CellBooleanAPIs || [];

		this.CellFloat64APIs = data?.CellFloat64APIs || [];

		this.CellIconAPIs = data?.CellIconAPIs || [];

		this.CellIntAPIs = data?.CellIntAPIs || [];

		this.CellStringAPIs = data?.CellStringAPIs || [];

		this.DisplayedColumnAPIs = data?.DisplayedColumnAPIs || [];

		this.RowAPIs = data?.RowAPIs || [];

		this.SVGIconAPIs = data?.SVGIconAPIs || [];

		this.TableAPIs = data?.TableAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
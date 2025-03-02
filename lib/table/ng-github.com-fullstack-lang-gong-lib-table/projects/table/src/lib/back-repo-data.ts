// generated code - do not edit

//insertion point for imports
import { CellAPI } from './cell-api'

import { CellBooleanAPI } from './cellboolean-api'

import { CellFloat64API } from './cellfloat64-api'

import { CellIconAPI } from './cellicon-api'

import { CellIntAPI } from './cellint-api'

import { CellStringAPI } from './cellstring-api'

import { CheckBoxAPI } from './checkbox-api'

import { DisplayedColumnAPI } from './displayedcolumn-api'

import { FormDivAPI } from './formdiv-api'

import { FormEditAssocButtonAPI } from './formeditassocbutton-api'

import { FormFieldAPI } from './formfield-api'

import { FormFieldDateAPI } from './formfielddate-api'

import { FormFieldDateTimeAPI } from './formfielddatetime-api'

import { FormFieldFloat64API } from './formfieldfloat64-api'

import { FormFieldIntAPI } from './formfieldint-api'

import { FormFieldSelectAPI } from './formfieldselect-api'

import { FormFieldStringAPI } from './formfieldstring-api'

import { FormFieldTimeAPI } from './formfieldtime-api'

import { FormGroupAPI } from './formgroup-api'

import { FormSortAssocButtonAPI } from './formsortassocbutton-api'

import { OptionAPI } from './option-api'

import { RowAPI } from './row-api'

import { TableAPI } from './table-api'


export class BackRepoData {
	// insertion point for declarations
	CellAPIs = new Array<CellAPI>()

	CellBooleanAPIs = new Array<CellBooleanAPI>()

	CellFloat64APIs = new Array<CellFloat64API>()

	CellIconAPIs = new Array<CellIconAPI>()

	CellIntAPIs = new Array<CellIntAPI>()

	CellStringAPIs = new Array<CellStringAPI>()

	CheckBoxAPIs = new Array<CheckBoxAPI>()

	DisplayedColumnAPIs = new Array<DisplayedColumnAPI>()

	FormDivAPIs = new Array<FormDivAPI>()

	FormEditAssocButtonAPIs = new Array<FormEditAssocButtonAPI>()

	FormFieldAPIs = new Array<FormFieldAPI>()

	FormFieldDateAPIs = new Array<FormFieldDateAPI>()

	FormFieldDateTimeAPIs = new Array<FormFieldDateTimeAPI>()

	FormFieldFloat64APIs = new Array<FormFieldFloat64API>()

	FormFieldIntAPIs = new Array<FormFieldIntAPI>()

	FormFieldSelectAPIs = new Array<FormFieldSelectAPI>()

	FormFieldStringAPIs = new Array<FormFieldStringAPI>()

	FormFieldTimeAPIs = new Array<FormFieldTimeAPI>()

	FormGroupAPIs = new Array<FormGroupAPI>()

	FormSortAssocButtonAPIs = new Array<FormSortAssocButtonAPI>()

	OptionAPIs = new Array<OptionAPI>()

	RowAPIs = new Array<RowAPI>()

	TableAPIs = new Array<TableAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CellAPIs = data?.CellAPIs || [];

		this.CellBooleanAPIs = data?.CellBooleanAPIs || [];

		this.CellFloat64APIs = data?.CellFloat64APIs || [];

		this.CellIconAPIs = data?.CellIconAPIs || [];

		this.CellIntAPIs = data?.CellIntAPIs || [];

		this.CellStringAPIs = data?.CellStringAPIs || [];

		this.CheckBoxAPIs = data?.CheckBoxAPIs || [];

		this.DisplayedColumnAPIs = data?.DisplayedColumnAPIs || [];

		this.FormDivAPIs = data?.FormDivAPIs || [];

		this.FormEditAssocButtonAPIs = data?.FormEditAssocButtonAPIs || [];

		this.FormFieldAPIs = data?.FormFieldAPIs || [];

		this.FormFieldDateAPIs = data?.FormFieldDateAPIs || [];

		this.FormFieldDateTimeAPIs = data?.FormFieldDateTimeAPIs || [];

		this.FormFieldFloat64APIs = data?.FormFieldFloat64APIs || [];

		this.FormFieldIntAPIs = data?.FormFieldIntAPIs || [];

		this.FormFieldSelectAPIs = data?.FormFieldSelectAPIs || [];

		this.FormFieldStringAPIs = data?.FormFieldStringAPIs || [];

		this.FormFieldTimeAPIs = data?.FormFieldTimeAPIs || [];

		this.FormGroupAPIs = data?.FormGroupAPIs || [];

		this.FormSortAssocButtonAPIs = data?.FormSortAssocButtonAPIs || [];

		this.OptionAPIs = data?.OptionAPIs || [];

		this.RowAPIs = data?.RowAPIs || [];

		this.TableAPIs = data?.TableAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
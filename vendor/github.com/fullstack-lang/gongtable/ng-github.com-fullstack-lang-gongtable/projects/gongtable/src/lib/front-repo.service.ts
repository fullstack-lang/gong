import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CellAPI } from './cell-api'
import { Cell, CopyCellAPIToCell } from './cell'
import { CellService } from './cell.service'

import { CellBooleanAPI } from './cellboolean-api'
import { CellBoolean, CopyCellBooleanAPIToCellBoolean } from './cellboolean'
import { CellBooleanService } from './cellboolean.service'

import { CellFloat64API } from './cellfloat64-api'
import { CellFloat64, CopyCellFloat64APIToCellFloat64 } from './cellfloat64'
import { CellFloat64Service } from './cellfloat64.service'

import { CellIconAPI } from './cellicon-api'
import { CellIcon, CopyCellIconAPIToCellIcon } from './cellicon'
import { CellIconService } from './cellicon.service'

import { CellIntAPI } from './cellint-api'
import { CellInt, CopyCellIntAPIToCellInt } from './cellint'
import { CellIntService } from './cellint.service'

import { CellStringAPI } from './cellstring-api'
import { CellString, CopyCellStringAPIToCellString } from './cellstring'
import { CellStringService } from './cellstring.service'

import { CheckBoxAPI } from './checkbox-api'
import { CheckBox, CopyCheckBoxAPIToCheckBox } from './checkbox'
import { CheckBoxService } from './checkbox.service'

import { DisplayedColumnAPI } from './displayedcolumn-api'
import { DisplayedColumn, CopyDisplayedColumnAPIToDisplayedColumn } from './displayedcolumn'
import { DisplayedColumnService } from './displayedcolumn.service'

import { FormDivAPI } from './formdiv-api'
import { FormDiv, CopyFormDivAPIToFormDiv } from './formdiv'
import { FormDivService } from './formdiv.service'

import { FormEditAssocButtonAPI } from './formeditassocbutton-api'
import { FormEditAssocButton, CopyFormEditAssocButtonAPIToFormEditAssocButton } from './formeditassocbutton'
import { FormEditAssocButtonService } from './formeditassocbutton.service'

import { FormFieldAPI } from './formfield-api'
import { FormField, CopyFormFieldAPIToFormField } from './formfield'
import { FormFieldService } from './formfield.service'

import { FormFieldDateAPI } from './formfielddate-api'
import { FormFieldDate, CopyFormFieldDateAPIToFormFieldDate } from './formfielddate'
import { FormFieldDateService } from './formfielddate.service'

import { FormFieldDateTimeAPI } from './formfielddatetime-api'
import { FormFieldDateTime, CopyFormFieldDateTimeAPIToFormFieldDateTime } from './formfielddatetime'
import { FormFieldDateTimeService } from './formfielddatetime.service'

import { FormFieldFloat64API } from './formfieldfloat64-api'
import { FormFieldFloat64, CopyFormFieldFloat64APIToFormFieldFloat64 } from './formfieldfloat64'
import { FormFieldFloat64Service } from './formfieldfloat64.service'

import { FormFieldIntAPI } from './formfieldint-api'
import { FormFieldInt, CopyFormFieldIntAPIToFormFieldInt } from './formfieldint'
import { FormFieldIntService } from './formfieldint.service'

import { FormFieldSelectAPI } from './formfieldselect-api'
import { FormFieldSelect, CopyFormFieldSelectAPIToFormFieldSelect } from './formfieldselect'
import { FormFieldSelectService } from './formfieldselect.service'

import { FormFieldStringAPI } from './formfieldstring-api'
import { FormFieldString, CopyFormFieldStringAPIToFormFieldString } from './formfieldstring'
import { FormFieldStringService } from './formfieldstring.service'

import { FormFieldTimeAPI } from './formfieldtime-api'
import { FormFieldTime, CopyFormFieldTimeAPIToFormFieldTime } from './formfieldtime'
import { FormFieldTimeService } from './formfieldtime.service'

import { FormGroupAPI } from './formgroup-api'
import { FormGroup, CopyFormGroupAPIToFormGroup } from './formgroup'
import { FormGroupService } from './formgroup.service'

import { FormSortAssocButtonAPI } from './formsortassocbutton-api'
import { FormSortAssocButton, CopyFormSortAssocButtonAPIToFormSortAssocButton } from './formsortassocbutton'
import { FormSortAssocButtonService } from './formsortassocbutton.service'

import { OptionAPI } from './option-api'
import { Option, CopyOptionAPIToOption } from './option'
import { OptionService } from './option.service'

import { RowAPI } from './row-api'
import { Row, CopyRowAPIToRow } from './row'
import { RowService } from './row.service'

import { TableAPI } from './table-api'
import { Table, CopyTableAPIToTable } from './table'
import { TableService } from './table.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongtable/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Cells = new Array<Cell>() // array of front instances
	map_ID_Cell = new Map<number, Cell>() // map of front instances

	array_CellBooleans = new Array<CellBoolean>() // array of front instances
	map_ID_CellBoolean = new Map<number, CellBoolean>() // map of front instances

	array_CellFloat64s = new Array<CellFloat64>() // array of front instances
	map_ID_CellFloat64 = new Map<number, CellFloat64>() // map of front instances

	array_CellIcons = new Array<CellIcon>() // array of front instances
	map_ID_CellIcon = new Map<number, CellIcon>() // map of front instances

	array_CellInts = new Array<CellInt>() // array of front instances
	map_ID_CellInt = new Map<number, CellInt>() // map of front instances

	array_CellStrings = new Array<CellString>() // array of front instances
	map_ID_CellString = new Map<number, CellString>() // map of front instances

	array_CheckBoxs = new Array<CheckBox>() // array of front instances
	map_ID_CheckBox = new Map<number, CheckBox>() // map of front instances

	array_DisplayedColumns = new Array<DisplayedColumn>() // array of front instances
	map_ID_DisplayedColumn = new Map<number, DisplayedColumn>() // map of front instances

	array_FormDivs = new Array<FormDiv>() // array of front instances
	map_ID_FormDiv = new Map<number, FormDiv>() // map of front instances

	array_FormEditAssocButtons = new Array<FormEditAssocButton>() // array of front instances
	map_ID_FormEditAssocButton = new Map<number, FormEditAssocButton>() // map of front instances

	array_FormFields = new Array<FormField>() // array of front instances
	map_ID_FormField = new Map<number, FormField>() // map of front instances

	array_FormFieldDates = new Array<FormFieldDate>() // array of front instances
	map_ID_FormFieldDate = new Map<number, FormFieldDate>() // map of front instances

	array_FormFieldDateTimes = new Array<FormFieldDateTime>() // array of front instances
	map_ID_FormFieldDateTime = new Map<number, FormFieldDateTime>() // map of front instances

	array_FormFieldFloat64s = new Array<FormFieldFloat64>() // array of front instances
	map_ID_FormFieldFloat64 = new Map<number, FormFieldFloat64>() // map of front instances

	array_FormFieldInts = new Array<FormFieldInt>() // array of front instances
	map_ID_FormFieldInt = new Map<number, FormFieldInt>() // map of front instances

	array_FormFieldSelects = new Array<FormFieldSelect>() // array of front instances
	map_ID_FormFieldSelect = new Map<number, FormFieldSelect>() // map of front instances

	array_FormFieldStrings = new Array<FormFieldString>() // array of front instances
	map_ID_FormFieldString = new Map<number, FormFieldString>() // map of front instances

	array_FormFieldTimes = new Array<FormFieldTime>() // array of front instances
	map_ID_FormFieldTime = new Map<number, FormFieldTime>() // map of front instances

	array_FormGroups = new Array<FormGroup>() // array of front instances
	map_ID_FormGroup = new Map<number, FormGroup>() // map of front instances

	array_FormSortAssocButtons = new Array<FormSortAssocButton>() // array of front instances
	map_ID_FormSortAssocButton = new Map<number, FormSortAssocButton>() // map of front instances

	array_Options = new Array<Option>() // array of front instances
	map_ID_Option = new Map<number, Option>() // map of front instances

	array_Rows = new Array<Row>() // array of front instances
	map_ID_Row = new Map<number, Row>() // map of front instances

	array_Tables = new Array<Table>() // array of front instances
	map_ID_Table = new Map<number, Table>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Cell':
				return this.array_Cells as unknown as Array<Type>
			case 'CellBoolean':
				return this.array_CellBooleans as unknown as Array<Type>
			case 'CellFloat64':
				return this.array_CellFloat64s as unknown as Array<Type>
			case 'CellIcon':
				return this.array_CellIcons as unknown as Array<Type>
			case 'CellInt':
				return this.array_CellInts as unknown as Array<Type>
			case 'CellString':
				return this.array_CellStrings as unknown as Array<Type>
			case 'CheckBox':
				return this.array_CheckBoxs as unknown as Array<Type>
			case 'DisplayedColumn':
				return this.array_DisplayedColumns as unknown as Array<Type>
			case 'FormDiv':
				return this.array_FormDivs as unknown as Array<Type>
			case 'FormEditAssocButton':
				return this.array_FormEditAssocButtons as unknown as Array<Type>
			case 'FormField':
				return this.array_FormFields as unknown as Array<Type>
			case 'FormFieldDate':
				return this.array_FormFieldDates as unknown as Array<Type>
			case 'FormFieldDateTime':
				return this.array_FormFieldDateTimes as unknown as Array<Type>
			case 'FormFieldFloat64':
				return this.array_FormFieldFloat64s as unknown as Array<Type>
			case 'FormFieldInt':
				return this.array_FormFieldInts as unknown as Array<Type>
			case 'FormFieldSelect':
				return this.array_FormFieldSelects as unknown as Array<Type>
			case 'FormFieldString':
				return this.array_FormFieldStrings as unknown as Array<Type>
			case 'FormFieldTime':
				return this.array_FormFieldTimes as unknown as Array<Type>
			case 'FormGroup':
				return this.array_FormGroups as unknown as Array<Type>
			case 'FormSortAssocButton':
				return this.array_FormSortAssocButtons as unknown as Array<Type>
			case 'Option':
				return this.array_Options as unknown as Array<Type>
			case 'Row':
				return this.array_Rows as unknown as Array<Type>
			case 'Table':
				return this.array_Tables as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Cell':
				return this.map_ID_Cell as unknown as Map<number, Type>
			case 'CellBoolean':
				return this.map_ID_CellBoolean as unknown as Map<number, Type>
			case 'CellFloat64':
				return this.map_ID_CellFloat64 as unknown as Map<number, Type>
			case 'CellIcon':
				return this.map_ID_CellIcon as unknown as Map<number, Type>
			case 'CellInt':
				return this.map_ID_CellInt as unknown as Map<number, Type>
			case 'CellString':
				return this.map_ID_CellString as unknown as Map<number, Type>
			case 'CheckBox':
				return this.map_ID_CheckBox as unknown as Map<number, Type>
			case 'DisplayedColumn':
				return this.map_ID_DisplayedColumn as unknown as Map<number, Type>
			case 'FormDiv':
				return this.map_ID_FormDiv as unknown as Map<number, Type>
			case 'FormEditAssocButton':
				return this.map_ID_FormEditAssocButton as unknown as Map<number, Type>
			case 'FormField':
				return this.map_ID_FormField as unknown as Map<number, Type>
			case 'FormFieldDate':
				return this.map_ID_FormFieldDate as unknown as Map<number, Type>
			case 'FormFieldDateTime':
				return this.map_ID_FormFieldDateTime as unknown as Map<number, Type>
			case 'FormFieldFloat64':
				return this.map_ID_FormFieldFloat64 as unknown as Map<number, Type>
			case 'FormFieldInt':
				return this.map_ID_FormFieldInt as unknown as Map<number, Type>
			case 'FormFieldSelect':
				return this.map_ID_FormFieldSelect as unknown as Map<number, Type>
			case 'FormFieldString':
				return this.map_ID_FormFieldString as unknown as Map<number, Type>
			case 'FormFieldTime':
				return this.map_ID_FormFieldTime as unknown as Map<number, Type>
			case 'FormGroup':
				return this.map_ID_FormGroup as unknown as Map<number, Type>
			case 'FormSortAssocButton':
				return this.map_ID_FormSortAssocButton as unknown as Map<number, Type>
			case 'Option':
				return this.map_ID_Option as unknown as Map<number, Type>
			case 'Row':
				return this.map_ID_Row as unknown as Map<number, Type>
			case 'Table':
				return this.map_ID_Table as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private cellService: CellService,
		private cellbooleanService: CellBooleanService,
		private cellfloat64Service: CellFloat64Service,
		private celliconService: CellIconService,
		private cellintService: CellIntService,
		private cellstringService: CellStringService,
		private checkboxService: CheckBoxService,
		private displayedcolumnService: DisplayedColumnService,
		private formdivService: FormDivService,
		private formeditassocbuttonService: FormEditAssocButtonService,
		private formfieldService: FormFieldService,
		private formfielddateService: FormFieldDateService,
		private formfielddatetimeService: FormFieldDateTimeService,
		private formfieldfloat64Service: FormFieldFloat64Service,
		private formfieldintService: FormFieldIntService,
		private formfieldselectService: FormFieldSelectService,
		private formfieldstringService: FormFieldStringService,
		private formfieldtimeService: FormFieldTimeService,
		private formgroupService: FormGroupService,
		private formsortassocbuttonService: FormSortAssocButtonService,
		private optionService: OptionService,
		private rowService: RowService,
		private tableService: TableService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<CellAPI[]>,
		Observable<CellBooleanAPI[]>,
		Observable<CellFloat64API[]>,
		Observable<CellIconAPI[]>,
		Observable<CellIntAPI[]>,
		Observable<CellStringAPI[]>,
		Observable<CheckBoxAPI[]>,
		Observable<DisplayedColumnAPI[]>,
		Observable<FormDivAPI[]>,
		Observable<FormEditAssocButtonAPI[]>,
		Observable<FormFieldAPI[]>,
		Observable<FormFieldDateAPI[]>,
		Observable<FormFieldDateTimeAPI[]>,
		Observable<FormFieldFloat64API[]>,
		Observable<FormFieldIntAPI[]>,
		Observable<FormFieldSelectAPI[]>,
		Observable<FormFieldStringAPI[]>,
		Observable<FormFieldTimeAPI[]>,
		Observable<FormGroupAPI[]>,
		Observable<FormSortAssocButtonAPI[]>,
		Observable<OptionAPI[]>,
		Observable<RowAPI[]>,
		Observable<TableAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.cellService.getCells(this.GONG__StackPath, this.frontRepo),
			this.cellbooleanService.getCellBooleans(this.GONG__StackPath, this.frontRepo),
			this.cellfloat64Service.getCellFloat64s(this.GONG__StackPath, this.frontRepo),
			this.celliconService.getCellIcons(this.GONG__StackPath, this.frontRepo),
			this.cellintService.getCellInts(this.GONG__StackPath, this.frontRepo),
			this.cellstringService.getCellStrings(this.GONG__StackPath, this.frontRepo),
			this.checkboxService.getCheckBoxs(this.GONG__StackPath, this.frontRepo),
			this.displayedcolumnService.getDisplayedColumns(this.GONG__StackPath, this.frontRepo),
			this.formdivService.getFormDivs(this.GONG__StackPath, this.frontRepo),
			this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath, this.frontRepo),
			this.formfieldService.getFormFields(this.GONG__StackPath, this.frontRepo),
			this.formfielddateService.getFormFieldDates(this.GONG__StackPath, this.frontRepo),
			this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath, this.frontRepo),
			this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath, this.frontRepo),
			this.formfieldintService.getFormFieldInts(this.GONG__StackPath, this.frontRepo),
			this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath, this.frontRepo),
			this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath, this.frontRepo),
			this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath, this.frontRepo),
			this.formgroupService.getFormGroups(this.GONG__StackPath, this.frontRepo),
			this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath, this.frontRepo),
			this.optionService.getOptions(this.GONG__StackPath, this.frontRepo),
			this.rowService.getRows(this.GONG__StackPath, this.frontRepo),
			this.tableService.getTables(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.cellService.getCells(this.GONG__StackPath, this.frontRepo),
			this.cellbooleanService.getCellBooleans(this.GONG__StackPath, this.frontRepo),
			this.cellfloat64Service.getCellFloat64s(this.GONG__StackPath, this.frontRepo),
			this.celliconService.getCellIcons(this.GONG__StackPath, this.frontRepo),
			this.cellintService.getCellInts(this.GONG__StackPath, this.frontRepo),
			this.cellstringService.getCellStrings(this.GONG__StackPath, this.frontRepo),
			this.checkboxService.getCheckBoxs(this.GONG__StackPath, this.frontRepo),
			this.displayedcolumnService.getDisplayedColumns(this.GONG__StackPath, this.frontRepo),
			this.formdivService.getFormDivs(this.GONG__StackPath, this.frontRepo),
			this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath, this.frontRepo),
			this.formfieldService.getFormFields(this.GONG__StackPath, this.frontRepo),
			this.formfielddateService.getFormFieldDates(this.GONG__StackPath, this.frontRepo),
			this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath, this.frontRepo),
			this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath, this.frontRepo),
			this.formfieldintService.getFormFieldInts(this.GONG__StackPath, this.frontRepo),
			this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath, this.frontRepo),
			this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath, this.frontRepo),
			this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath, this.frontRepo),
			this.formgroupService.getFormGroups(this.GONG__StackPath, this.frontRepo),
			this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath, this.frontRepo),
			this.optionService.getOptions(this.GONG__StackPath, this.frontRepo),
			this.rowService.getRows(this.GONG__StackPath, this.frontRepo),
			this.tableService.getTables(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						cells_,
						cellbooleans_,
						cellfloat64s_,
						cellicons_,
						cellints_,
						cellstrings_,
						checkboxs_,
						displayedcolumns_,
						formdivs_,
						formeditassocbuttons_,
						formfields_,
						formfielddates_,
						formfielddatetimes_,
						formfieldfloat64s_,
						formfieldints_,
						formfieldselects_,
						formfieldstrings_,
						formfieldtimes_,
						formgroups_,
						formsortassocbuttons_,
						options_,
						rows_,
						tables_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var cells: CellAPI[]
						cells = cells_ as CellAPI[]
						var cellbooleans: CellBooleanAPI[]
						cellbooleans = cellbooleans_ as CellBooleanAPI[]
						var cellfloat64s: CellFloat64API[]
						cellfloat64s = cellfloat64s_ as CellFloat64API[]
						var cellicons: CellIconAPI[]
						cellicons = cellicons_ as CellIconAPI[]
						var cellints: CellIntAPI[]
						cellints = cellints_ as CellIntAPI[]
						var cellstrings: CellStringAPI[]
						cellstrings = cellstrings_ as CellStringAPI[]
						var checkboxs: CheckBoxAPI[]
						checkboxs = checkboxs_ as CheckBoxAPI[]
						var displayedcolumns: DisplayedColumnAPI[]
						displayedcolumns = displayedcolumns_ as DisplayedColumnAPI[]
						var formdivs: FormDivAPI[]
						formdivs = formdivs_ as FormDivAPI[]
						var formeditassocbuttons: FormEditAssocButtonAPI[]
						formeditassocbuttons = formeditassocbuttons_ as FormEditAssocButtonAPI[]
						var formfields: FormFieldAPI[]
						formfields = formfields_ as FormFieldAPI[]
						var formfielddates: FormFieldDateAPI[]
						formfielddates = formfielddates_ as FormFieldDateAPI[]
						var formfielddatetimes: FormFieldDateTimeAPI[]
						formfielddatetimes = formfielddatetimes_ as FormFieldDateTimeAPI[]
						var formfieldfloat64s: FormFieldFloat64API[]
						formfieldfloat64s = formfieldfloat64s_ as FormFieldFloat64API[]
						var formfieldints: FormFieldIntAPI[]
						formfieldints = formfieldints_ as FormFieldIntAPI[]
						var formfieldselects: FormFieldSelectAPI[]
						formfieldselects = formfieldselects_ as FormFieldSelectAPI[]
						var formfieldstrings: FormFieldStringAPI[]
						formfieldstrings = formfieldstrings_ as FormFieldStringAPI[]
						var formfieldtimes: FormFieldTimeAPI[]
						formfieldtimes = formfieldtimes_ as FormFieldTimeAPI[]
						var formgroups: FormGroupAPI[]
						formgroups = formgroups_ as FormGroupAPI[]
						var formsortassocbuttons: FormSortAssocButtonAPI[]
						formsortassocbuttons = formsortassocbuttons_ as FormSortAssocButtonAPI[]
						var options: OptionAPI[]
						options = options_ as OptionAPI[]
						var rows: RowAPI[]
						rows = rows_ as RowAPI[]
						var tables: TableAPI[]
						tables = tables_ as TableAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Cells = []
						this.frontRepo.map_ID_Cell.clear()

						cells.forEach(
							cellAPI => {
								let cell = new Cell
								this.frontRepo.array_Cells.push(cell)
								this.frontRepo.map_ID_Cell.set(cellAPI.ID, cell)
							}
						)

						// init the arrays
						this.frontRepo.array_CellBooleans = []
						this.frontRepo.map_ID_CellBoolean.clear()

						cellbooleans.forEach(
							cellbooleanAPI => {
								let cellboolean = new CellBoolean
								this.frontRepo.array_CellBooleans.push(cellboolean)
								this.frontRepo.map_ID_CellBoolean.set(cellbooleanAPI.ID, cellboolean)
							}
						)

						// init the arrays
						this.frontRepo.array_CellFloat64s = []
						this.frontRepo.map_ID_CellFloat64.clear()

						cellfloat64s.forEach(
							cellfloat64API => {
								let cellfloat64 = new CellFloat64
								this.frontRepo.array_CellFloat64s.push(cellfloat64)
								this.frontRepo.map_ID_CellFloat64.set(cellfloat64API.ID, cellfloat64)
							}
						)

						// init the arrays
						this.frontRepo.array_CellIcons = []
						this.frontRepo.map_ID_CellIcon.clear()

						cellicons.forEach(
							celliconAPI => {
								let cellicon = new CellIcon
								this.frontRepo.array_CellIcons.push(cellicon)
								this.frontRepo.map_ID_CellIcon.set(celliconAPI.ID, cellicon)
							}
						)

						// init the arrays
						this.frontRepo.array_CellInts = []
						this.frontRepo.map_ID_CellInt.clear()

						cellints.forEach(
							cellintAPI => {
								let cellint = new CellInt
								this.frontRepo.array_CellInts.push(cellint)
								this.frontRepo.map_ID_CellInt.set(cellintAPI.ID, cellint)
							}
						)

						// init the arrays
						this.frontRepo.array_CellStrings = []
						this.frontRepo.map_ID_CellString.clear()

						cellstrings.forEach(
							cellstringAPI => {
								let cellstring = new CellString
								this.frontRepo.array_CellStrings.push(cellstring)
								this.frontRepo.map_ID_CellString.set(cellstringAPI.ID, cellstring)
							}
						)

						// init the arrays
						this.frontRepo.array_CheckBoxs = []
						this.frontRepo.map_ID_CheckBox.clear()

						checkboxs.forEach(
							checkboxAPI => {
								let checkbox = new CheckBox
								this.frontRepo.array_CheckBoxs.push(checkbox)
								this.frontRepo.map_ID_CheckBox.set(checkboxAPI.ID, checkbox)
							}
						)

						// init the arrays
						this.frontRepo.array_DisplayedColumns = []
						this.frontRepo.map_ID_DisplayedColumn.clear()

						displayedcolumns.forEach(
							displayedcolumnAPI => {
								let displayedcolumn = new DisplayedColumn
								this.frontRepo.array_DisplayedColumns.push(displayedcolumn)
								this.frontRepo.map_ID_DisplayedColumn.set(displayedcolumnAPI.ID, displayedcolumn)
							}
						)

						// init the arrays
						this.frontRepo.array_FormDivs = []
						this.frontRepo.map_ID_FormDiv.clear()

						formdivs.forEach(
							formdivAPI => {
								let formdiv = new FormDiv
								this.frontRepo.array_FormDivs.push(formdiv)
								this.frontRepo.map_ID_FormDiv.set(formdivAPI.ID, formdiv)
							}
						)

						// init the arrays
						this.frontRepo.array_FormEditAssocButtons = []
						this.frontRepo.map_ID_FormEditAssocButton.clear()

						formeditassocbuttons.forEach(
							formeditassocbuttonAPI => {
								let formeditassocbutton = new FormEditAssocButton
								this.frontRepo.array_FormEditAssocButtons.push(formeditassocbutton)
								this.frontRepo.map_ID_FormEditAssocButton.set(formeditassocbuttonAPI.ID, formeditassocbutton)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFields = []
						this.frontRepo.map_ID_FormField.clear()

						formfields.forEach(
							formfieldAPI => {
								let formfield = new FormField
								this.frontRepo.array_FormFields.push(formfield)
								this.frontRepo.map_ID_FormField.set(formfieldAPI.ID, formfield)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldDates = []
						this.frontRepo.map_ID_FormFieldDate.clear()

						formfielddates.forEach(
							formfielddateAPI => {
								let formfielddate = new FormFieldDate
								this.frontRepo.array_FormFieldDates.push(formfielddate)
								this.frontRepo.map_ID_FormFieldDate.set(formfielddateAPI.ID, formfielddate)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldDateTimes = []
						this.frontRepo.map_ID_FormFieldDateTime.clear()

						formfielddatetimes.forEach(
							formfielddatetimeAPI => {
								let formfielddatetime = new FormFieldDateTime
								this.frontRepo.array_FormFieldDateTimes.push(formfielddatetime)
								this.frontRepo.map_ID_FormFieldDateTime.set(formfielddatetimeAPI.ID, formfielddatetime)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldFloat64s = []
						this.frontRepo.map_ID_FormFieldFloat64.clear()

						formfieldfloat64s.forEach(
							formfieldfloat64API => {
								let formfieldfloat64 = new FormFieldFloat64
								this.frontRepo.array_FormFieldFloat64s.push(formfieldfloat64)
								this.frontRepo.map_ID_FormFieldFloat64.set(formfieldfloat64API.ID, formfieldfloat64)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldInts = []
						this.frontRepo.map_ID_FormFieldInt.clear()

						formfieldints.forEach(
							formfieldintAPI => {
								let formfieldint = new FormFieldInt
								this.frontRepo.array_FormFieldInts.push(formfieldint)
								this.frontRepo.map_ID_FormFieldInt.set(formfieldintAPI.ID, formfieldint)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldSelects = []
						this.frontRepo.map_ID_FormFieldSelect.clear()

						formfieldselects.forEach(
							formfieldselectAPI => {
								let formfieldselect = new FormFieldSelect
								this.frontRepo.array_FormFieldSelects.push(formfieldselect)
								this.frontRepo.map_ID_FormFieldSelect.set(formfieldselectAPI.ID, formfieldselect)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldStrings = []
						this.frontRepo.map_ID_FormFieldString.clear()

						formfieldstrings.forEach(
							formfieldstringAPI => {
								let formfieldstring = new FormFieldString
								this.frontRepo.array_FormFieldStrings.push(formfieldstring)
								this.frontRepo.map_ID_FormFieldString.set(formfieldstringAPI.ID, formfieldstring)
							}
						)

						// init the arrays
						this.frontRepo.array_FormFieldTimes = []
						this.frontRepo.map_ID_FormFieldTime.clear()

						formfieldtimes.forEach(
							formfieldtimeAPI => {
								let formfieldtime = new FormFieldTime
								this.frontRepo.array_FormFieldTimes.push(formfieldtime)
								this.frontRepo.map_ID_FormFieldTime.set(formfieldtimeAPI.ID, formfieldtime)
							}
						)

						// init the arrays
						this.frontRepo.array_FormGroups = []
						this.frontRepo.map_ID_FormGroup.clear()

						formgroups.forEach(
							formgroupAPI => {
								let formgroup = new FormGroup
								this.frontRepo.array_FormGroups.push(formgroup)
								this.frontRepo.map_ID_FormGroup.set(formgroupAPI.ID, formgroup)
							}
						)

						// init the arrays
						this.frontRepo.array_FormSortAssocButtons = []
						this.frontRepo.map_ID_FormSortAssocButton.clear()

						formsortassocbuttons.forEach(
							formsortassocbuttonAPI => {
								let formsortassocbutton = new FormSortAssocButton
								this.frontRepo.array_FormSortAssocButtons.push(formsortassocbutton)
								this.frontRepo.map_ID_FormSortAssocButton.set(formsortassocbuttonAPI.ID, formsortassocbutton)
							}
						)

						// init the arrays
						this.frontRepo.array_Options = []
						this.frontRepo.map_ID_Option.clear()

						options.forEach(
							optionAPI => {
								let option = new Option
								this.frontRepo.array_Options.push(option)
								this.frontRepo.map_ID_Option.set(optionAPI.ID, option)
							}
						)

						// init the arrays
						this.frontRepo.array_Rows = []
						this.frontRepo.map_ID_Row.clear()

						rows.forEach(
							rowAPI => {
								let row = new Row
								this.frontRepo.array_Rows.push(row)
								this.frontRepo.map_ID_Row.set(rowAPI.ID, row)
							}
						)

						// init the arrays
						this.frontRepo.array_Tables = []
						this.frontRepo.map_ID_Table.clear()

						tables.forEach(
							tableAPI => {
								let table = new Table
								this.frontRepo.array_Tables.push(table)
								this.frontRepo.map_ID_Table.set(tableAPI.ID, table)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						cells.forEach(
							cellAPI => {
								let cell = this.frontRepo.map_ID_Cell.get(cellAPI.ID)
								CopyCellAPIToCell(cellAPI, cell!, this.frontRepo)
							}
						)

						// fill up front objects
						cellbooleans.forEach(
							cellbooleanAPI => {
								let cellboolean = this.frontRepo.map_ID_CellBoolean.get(cellbooleanAPI.ID)
								CopyCellBooleanAPIToCellBoolean(cellbooleanAPI, cellboolean!, this.frontRepo)
							}
						)

						// fill up front objects
						cellfloat64s.forEach(
							cellfloat64API => {
								let cellfloat64 = this.frontRepo.map_ID_CellFloat64.get(cellfloat64API.ID)
								CopyCellFloat64APIToCellFloat64(cellfloat64API, cellfloat64!, this.frontRepo)
							}
						)

						// fill up front objects
						cellicons.forEach(
							celliconAPI => {
								let cellicon = this.frontRepo.map_ID_CellIcon.get(celliconAPI.ID)
								CopyCellIconAPIToCellIcon(celliconAPI, cellicon!, this.frontRepo)
							}
						)

						// fill up front objects
						cellints.forEach(
							cellintAPI => {
								let cellint = this.frontRepo.map_ID_CellInt.get(cellintAPI.ID)
								CopyCellIntAPIToCellInt(cellintAPI, cellint!, this.frontRepo)
							}
						)

						// fill up front objects
						cellstrings.forEach(
							cellstringAPI => {
								let cellstring = this.frontRepo.map_ID_CellString.get(cellstringAPI.ID)
								CopyCellStringAPIToCellString(cellstringAPI, cellstring!, this.frontRepo)
							}
						)

						// fill up front objects
						checkboxs.forEach(
							checkboxAPI => {
								let checkbox = this.frontRepo.map_ID_CheckBox.get(checkboxAPI.ID)
								CopyCheckBoxAPIToCheckBox(checkboxAPI, checkbox!, this.frontRepo)
							}
						)

						// fill up front objects
						displayedcolumns.forEach(
							displayedcolumnAPI => {
								let displayedcolumn = this.frontRepo.map_ID_DisplayedColumn.get(displayedcolumnAPI.ID)
								CopyDisplayedColumnAPIToDisplayedColumn(displayedcolumnAPI, displayedcolumn!, this.frontRepo)
							}
						)

						// fill up front objects
						formdivs.forEach(
							formdivAPI => {
								let formdiv = this.frontRepo.map_ID_FormDiv.get(formdivAPI.ID)
								CopyFormDivAPIToFormDiv(formdivAPI, formdiv!, this.frontRepo)
							}
						)

						// fill up front objects
						formeditassocbuttons.forEach(
							formeditassocbuttonAPI => {
								let formeditassocbutton = this.frontRepo.map_ID_FormEditAssocButton.get(formeditassocbuttonAPI.ID)
								CopyFormEditAssocButtonAPIToFormEditAssocButton(formeditassocbuttonAPI, formeditassocbutton!, this.frontRepo)
							}
						)

						// fill up front objects
						formfields.forEach(
							formfieldAPI => {
								let formfield = this.frontRepo.map_ID_FormField.get(formfieldAPI.ID)
								CopyFormFieldAPIToFormField(formfieldAPI, formfield!, this.frontRepo)
							}
						)

						// fill up front objects
						formfielddates.forEach(
							formfielddateAPI => {
								let formfielddate = this.frontRepo.map_ID_FormFieldDate.get(formfielddateAPI.ID)
								CopyFormFieldDateAPIToFormFieldDate(formfielddateAPI, formfielddate!, this.frontRepo)
							}
						)

						// fill up front objects
						formfielddatetimes.forEach(
							formfielddatetimeAPI => {
								let formfielddatetime = this.frontRepo.map_ID_FormFieldDateTime.get(formfielddatetimeAPI.ID)
								CopyFormFieldDateTimeAPIToFormFieldDateTime(formfielddatetimeAPI, formfielddatetime!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldfloat64s.forEach(
							formfieldfloat64API => {
								let formfieldfloat64 = this.frontRepo.map_ID_FormFieldFloat64.get(formfieldfloat64API.ID)
								CopyFormFieldFloat64APIToFormFieldFloat64(formfieldfloat64API, formfieldfloat64!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldints.forEach(
							formfieldintAPI => {
								let formfieldint = this.frontRepo.map_ID_FormFieldInt.get(formfieldintAPI.ID)
								CopyFormFieldIntAPIToFormFieldInt(formfieldintAPI, formfieldint!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldselects.forEach(
							formfieldselectAPI => {
								let formfieldselect = this.frontRepo.map_ID_FormFieldSelect.get(formfieldselectAPI.ID)
								CopyFormFieldSelectAPIToFormFieldSelect(formfieldselectAPI, formfieldselect!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldstrings.forEach(
							formfieldstringAPI => {
								let formfieldstring = this.frontRepo.map_ID_FormFieldString.get(formfieldstringAPI.ID)
								CopyFormFieldStringAPIToFormFieldString(formfieldstringAPI, formfieldstring!, this.frontRepo)
							}
						)

						// fill up front objects
						formfieldtimes.forEach(
							formfieldtimeAPI => {
								let formfieldtime = this.frontRepo.map_ID_FormFieldTime.get(formfieldtimeAPI.ID)
								CopyFormFieldTimeAPIToFormFieldTime(formfieldtimeAPI, formfieldtime!, this.frontRepo)
							}
						)

						// fill up front objects
						formgroups.forEach(
							formgroupAPI => {
								let formgroup = this.frontRepo.map_ID_FormGroup.get(formgroupAPI.ID)
								CopyFormGroupAPIToFormGroup(formgroupAPI, formgroup!, this.frontRepo)
							}
						)

						// fill up front objects
						formsortassocbuttons.forEach(
							formsortassocbuttonAPI => {
								let formsortassocbutton = this.frontRepo.map_ID_FormSortAssocButton.get(formsortassocbuttonAPI.ID)
								CopyFormSortAssocButtonAPIToFormSortAssocButton(formsortassocbuttonAPI, formsortassocbutton!, this.frontRepo)
							}
						)

						// fill up front objects
						options.forEach(
							optionAPI => {
								let option = this.frontRepo.map_ID_Option.get(optionAPI.ID)
								CopyOptionAPIToOption(optionAPI, option!, this.frontRepo)
							}
						)

						// fill up front objects
						rows.forEach(
							rowAPI => {
								let row = this.frontRepo.map_ID_Row.get(rowAPI.ID)
								CopyRowAPIToRow(rowAPI, row!, this.frontRepo)
							}
						)

						// fill up front objects
						tables.forEach(
							tableAPI => {
								let table = this.frontRepo.map_ID_Table.get(tableAPI.ID)
								CopyTableAPIToTable(tableAPI, table!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongtable/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {
				let _this = this

				const backRepoData = new BackRepoData(JSON.parse(event.data))

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				this.frontRepo.array_Cells = []
				this.frontRepo.map_ID_Cell.clear()

				backRepoData.CellAPIs.forEach(
					cellAPI => {
						let cell = new Cell
						this.frontRepo.array_Cells.push(cell)
						this.frontRepo.map_ID_Cell.set(cellAPI.ID, cell)
					}
				)

				// init the arrays
				this.frontRepo.array_CellBooleans = []
				this.frontRepo.map_ID_CellBoolean.clear()

				backRepoData.CellBooleanAPIs.forEach(
					cellbooleanAPI => {
						let cellboolean = new CellBoolean
						this.frontRepo.array_CellBooleans.push(cellboolean)
						this.frontRepo.map_ID_CellBoolean.set(cellbooleanAPI.ID, cellboolean)
					}
				)

				// init the arrays
				this.frontRepo.array_CellFloat64s = []
				this.frontRepo.map_ID_CellFloat64.clear()

				backRepoData.CellFloat64APIs.forEach(
					cellfloat64API => {
						let cellfloat64 = new CellFloat64
						this.frontRepo.array_CellFloat64s.push(cellfloat64)
						this.frontRepo.map_ID_CellFloat64.set(cellfloat64API.ID, cellfloat64)
					}
				)

				// init the arrays
				this.frontRepo.array_CellIcons = []
				this.frontRepo.map_ID_CellIcon.clear()

				backRepoData.CellIconAPIs.forEach(
					celliconAPI => {
						let cellicon = new CellIcon
						this.frontRepo.array_CellIcons.push(cellicon)
						this.frontRepo.map_ID_CellIcon.set(celliconAPI.ID, cellicon)
					}
				)

				// init the arrays
				this.frontRepo.array_CellInts = []
				this.frontRepo.map_ID_CellInt.clear()

				backRepoData.CellIntAPIs.forEach(
					cellintAPI => {
						let cellint = new CellInt
						this.frontRepo.array_CellInts.push(cellint)
						this.frontRepo.map_ID_CellInt.set(cellintAPI.ID, cellint)
					}
				)

				// init the arrays
				this.frontRepo.array_CellStrings = []
				this.frontRepo.map_ID_CellString.clear()

				backRepoData.CellStringAPIs.forEach(
					cellstringAPI => {
						let cellstring = new CellString
						this.frontRepo.array_CellStrings.push(cellstring)
						this.frontRepo.map_ID_CellString.set(cellstringAPI.ID, cellstring)
					}
				)

				// init the arrays
				this.frontRepo.array_CheckBoxs = []
				this.frontRepo.map_ID_CheckBox.clear()

				backRepoData.CheckBoxAPIs.forEach(
					checkboxAPI => {
						let checkbox = new CheckBox
						this.frontRepo.array_CheckBoxs.push(checkbox)
						this.frontRepo.map_ID_CheckBox.set(checkboxAPI.ID, checkbox)
					}
				)

				// init the arrays
				this.frontRepo.array_DisplayedColumns = []
				this.frontRepo.map_ID_DisplayedColumn.clear()

				backRepoData.DisplayedColumnAPIs.forEach(
					displayedcolumnAPI => {
						let displayedcolumn = new DisplayedColumn
						this.frontRepo.array_DisplayedColumns.push(displayedcolumn)
						this.frontRepo.map_ID_DisplayedColumn.set(displayedcolumnAPI.ID, displayedcolumn)
					}
				)

				// init the arrays
				this.frontRepo.array_FormDivs = []
				this.frontRepo.map_ID_FormDiv.clear()

				backRepoData.FormDivAPIs.forEach(
					formdivAPI => {
						let formdiv = new FormDiv
						this.frontRepo.array_FormDivs.push(formdiv)
						this.frontRepo.map_ID_FormDiv.set(formdivAPI.ID, formdiv)
					}
				)

				// init the arrays
				this.frontRepo.array_FormEditAssocButtons = []
				this.frontRepo.map_ID_FormEditAssocButton.clear()

				backRepoData.FormEditAssocButtonAPIs.forEach(
					formeditassocbuttonAPI => {
						let formeditassocbutton = new FormEditAssocButton
						this.frontRepo.array_FormEditAssocButtons.push(formeditassocbutton)
						this.frontRepo.map_ID_FormEditAssocButton.set(formeditassocbuttonAPI.ID, formeditassocbutton)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFields = []
				this.frontRepo.map_ID_FormField.clear()

				backRepoData.FormFieldAPIs.forEach(
					formfieldAPI => {
						let formfield = new FormField
						this.frontRepo.array_FormFields.push(formfield)
						this.frontRepo.map_ID_FormField.set(formfieldAPI.ID, formfield)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldDates = []
				this.frontRepo.map_ID_FormFieldDate.clear()

				backRepoData.FormFieldDateAPIs.forEach(
					formfielddateAPI => {
						let formfielddate = new FormFieldDate
						this.frontRepo.array_FormFieldDates.push(formfielddate)
						this.frontRepo.map_ID_FormFieldDate.set(formfielddateAPI.ID, formfielddate)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldDateTimes = []
				this.frontRepo.map_ID_FormFieldDateTime.clear()

				backRepoData.FormFieldDateTimeAPIs.forEach(
					formfielddatetimeAPI => {
						let formfielddatetime = new FormFieldDateTime
						this.frontRepo.array_FormFieldDateTimes.push(formfielddatetime)
						this.frontRepo.map_ID_FormFieldDateTime.set(formfielddatetimeAPI.ID, formfielddatetime)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldFloat64s = []
				this.frontRepo.map_ID_FormFieldFloat64.clear()

				backRepoData.FormFieldFloat64APIs.forEach(
					formfieldfloat64API => {
						let formfieldfloat64 = new FormFieldFloat64
						this.frontRepo.array_FormFieldFloat64s.push(formfieldfloat64)
						this.frontRepo.map_ID_FormFieldFloat64.set(formfieldfloat64API.ID, formfieldfloat64)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldInts = []
				this.frontRepo.map_ID_FormFieldInt.clear()

				backRepoData.FormFieldIntAPIs.forEach(
					formfieldintAPI => {
						let formfieldint = new FormFieldInt
						this.frontRepo.array_FormFieldInts.push(formfieldint)
						this.frontRepo.map_ID_FormFieldInt.set(formfieldintAPI.ID, formfieldint)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldSelects = []
				this.frontRepo.map_ID_FormFieldSelect.clear()

				backRepoData.FormFieldSelectAPIs.forEach(
					formfieldselectAPI => {
						let formfieldselect = new FormFieldSelect
						this.frontRepo.array_FormFieldSelects.push(formfieldselect)
						this.frontRepo.map_ID_FormFieldSelect.set(formfieldselectAPI.ID, formfieldselect)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldStrings = []
				this.frontRepo.map_ID_FormFieldString.clear()

				backRepoData.FormFieldStringAPIs.forEach(
					formfieldstringAPI => {
						let formfieldstring = new FormFieldString
						this.frontRepo.array_FormFieldStrings.push(formfieldstring)
						this.frontRepo.map_ID_FormFieldString.set(formfieldstringAPI.ID, formfieldstring)
					}
				)

				// init the arrays
				this.frontRepo.array_FormFieldTimes = []
				this.frontRepo.map_ID_FormFieldTime.clear()

				backRepoData.FormFieldTimeAPIs.forEach(
					formfieldtimeAPI => {
						let formfieldtime = new FormFieldTime
						this.frontRepo.array_FormFieldTimes.push(formfieldtime)
						this.frontRepo.map_ID_FormFieldTime.set(formfieldtimeAPI.ID, formfieldtime)
					}
				)

				// init the arrays
				this.frontRepo.array_FormGroups = []
				this.frontRepo.map_ID_FormGroup.clear()

				backRepoData.FormGroupAPIs.forEach(
					formgroupAPI => {
						let formgroup = new FormGroup
						this.frontRepo.array_FormGroups.push(formgroup)
						this.frontRepo.map_ID_FormGroup.set(formgroupAPI.ID, formgroup)
					}
				)

				// init the arrays
				this.frontRepo.array_FormSortAssocButtons = []
				this.frontRepo.map_ID_FormSortAssocButton.clear()

				backRepoData.FormSortAssocButtonAPIs.forEach(
					formsortassocbuttonAPI => {
						let formsortassocbutton = new FormSortAssocButton
						this.frontRepo.array_FormSortAssocButtons.push(formsortassocbutton)
						this.frontRepo.map_ID_FormSortAssocButton.set(formsortassocbuttonAPI.ID, formsortassocbutton)
					}
				)

				// init the arrays
				this.frontRepo.array_Options = []
				this.frontRepo.map_ID_Option.clear()

				backRepoData.OptionAPIs.forEach(
					optionAPI => {
						let option = new Option
						this.frontRepo.array_Options.push(option)
						this.frontRepo.map_ID_Option.set(optionAPI.ID, option)
					}
				)

				// init the arrays
				this.frontRepo.array_Rows = []
				this.frontRepo.map_ID_Row.clear()

				backRepoData.RowAPIs.forEach(
					rowAPI => {
						let row = new Row
						this.frontRepo.array_Rows.push(row)
						this.frontRepo.map_ID_Row.set(rowAPI.ID, row)
					}
				)

				// init the arrays
				this.frontRepo.array_Tables = []
				this.frontRepo.map_ID_Table.clear()

				backRepoData.TableAPIs.forEach(
					tableAPI => {
						let table = new Table
						this.frontRepo.array_Tables.push(table)
						this.frontRepo.map_ID_Table.set(tableAPI.ID, table)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.CellAPIs.forEach(
					cellAPI => {
						let cell = this.frontRepo.map_ID_Cell.get(cellAPI.ID)
						CopyCellAPIToCell(cellAPI, cell!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellBooleanAPIs.forEach(
					cellbooleanAPI => {
						let cellboolean = this.frontRepo.map_ID_CellBoolean.get(cellbooleanAPI.ID)
						CopyCellBooleanAPIToCellBoolean(cellbooleanAPI, cellboolean!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellFloat64APIs.forEach(
					cellfloat64API => {
						let cellfloat64 = this.frontRepo.map_ID_CellFloat64.get(cellfloat64API.ID)
						CopyCellFloat64APIToCellFloat64(cellfloat64API, cellfloat64!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellIconAPIs.forEach(
					celliconAPI => {
						let cellicon = this.frontRepo.map_ID_CellIcon.get(celliconAPI.ID)
						CopyCellIconAPIToCellIcon(celliconAPI, cellicon!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellIntAPIs.forEach(
					cellintAPI => {
						let cellint = this.frontRepo.map_ID_CellInt.get(cellintAPI.ID)
						CopyCellIntAPIToCellInt(cellintAPI, cellint!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellStringAPIs.forEach(
					cellstringAPI => {
						let cellstring = this.frontRepo.map_ID_CellString.get(cellstringAPI.ID)
						CopyCellStringAPIToCellString(cellstringAPI, cellstring!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CheckBoxAPIs.forEach(
					checkboxAPI => {
						let checkbox = this.frontRepo.map_ID_CheckBox.get(checkboxAPI.ID)
						CopyCheckBoxAPIToCheckBox(checkboxAPI, checkbox!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DisplayedColumnAPIs.forEach(
					displayedcolumnAPI => {
						let displayedcolumn = this.frontRepo.map_ID_DisplayedColumn.get(displayedcolumnAPI.ID)
						CopyDisplayedColumnAPIToDisplayedColumn(displayedcolumnAPI, displayedcolumn!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormDivAPIs.forEach(
					formdivAPI => {
						let formdiv = this.frontRepo.map_ID_FormDiv.get(formdivAPI.ID)
						CopyFormDivAPIToFormDiv(formdivAPI, formdiv!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormEditAssocButtonAPIs.forEach(
					formeditassocbuttonAPI => {
						let formeditassocbutton = this.frontRepo.map_ID_FormEditAssocButton.get(formeditassocbuttonAPI.ID)
						CopyFormEditAssocButtonAPIToFormEditAssocButton(formeditassocbuttonAPI, formeditassocbutton!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldAPIs.forEach(
					formfieldAPI => {
						let formfield = this.frontRepo.map_ID_FormField.get(formfieldAPI.ID)
						CopyFormFieldAPIToFormField(formfieldAPI, formfield!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldDateAPIs.forEach(
					formfielddateAPI => {
						let formfielddate = this.frontRepo.map_ID_FormFieldDate.get(formfielddateAPI.ID)
						CopyFormFieldDateAPIToFormFieldDate(formfielddateAPI, formfielddate!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldDateTimeAPIs.forEach(
					formfielddatetimeAPI => {
						let formfielddatetime = this.frontRepo.map_ID_FormFieldDateTime.get(formfielddatetimeAPI.ID)
						CopyFormFieldDateTimeAPIToFormFieldDateTime(formfielddatetimeAPI, formfielddatetime!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldFloat64APIs.forEach(
					formfieldfloat64API => {
						let formfieldfloat64 = this.frontRepo.map_ID_FormFieldFloat64.get(formfieldfloat64API.ID)
						CopyFormFieldFloat64APIToFormFieldFloat64(formfieldfloat64API, formfieldfloat64!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldIntAPIs.forEach(
					formfieldintAPI => {
						let formfieldint = this.frontRepo.map_ID_FormFieldInt.get(formfieldintAPI.ID)
						CopyFormFieldIntAPIToFormFieldInt(formfieldintAPI, formfieldint!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldSelectAPIs.forEach(
					formfieldselectAPI => {
						let formfieldselect = this.frontRepo.map_ID_FormFieldSelect.get(formfieldselectAPI.ID)
						CopyFormFieldSelectAPIToFormFieldSelect(formfieldselectAPI, formfieldselect!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldStringAPIs.forEach(
					formfieldstringAPI => {
						let formfieldstring = this.frontRepo.map_ID_FormFieldString.get(formfieldstringAPI.ID)
						CopyFormFieldStringAPIToFormFieldString(formfieldstringAPI, formfieldstring!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormFieldTimeAPIs.forEach(
					formfieldtimeAPI => {
						let formfieldtime = this.frontRepo.map_ID_FormFieldTime.get(formfieldtimeAPI.ID)
						CopyFormFieldTimeAPIToFormFieldTime(formfieldtimeAPI, formfieldtime!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormGroupAPIs.forEach(
					formgroupAPI => {
						let formgroup = this.frontRepo.map_ID_FormGroup.get(formgroupAPI.ID)
						CopyFormGroupAPIToFormGroup(formgroupAPI, formgroup!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.FormSortAssocButtonAPIs.forEach(
					formsortassocbuttonAPI => {
						let formsortassocbutton = this.frontRepo.map_ID_FormSortAssocButton.get(formsortassocbuttonAPI.ID)
						CopyFormSortAssocButtonAPIToFormSortAssocButton(formsortassocbuttonAPI, formsortassocbutton!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.OptionAPIs.forEach(
					optionAPI => {
						let option = this.frontRepo.map_ID_Option.get(optionAPI.ID)
						CopyOptionAPIToOption(optionAPI, option!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RowAPIs.forEach(
					rowAPI => {
						let row = this.frontRepo.map_ID_Row.get(rowAPI.ID)
						CopyRowAPIToRow(rowAPI, row!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TableAPIs.forEach(
					tableAPI => {
						let table = this.frontRepo.map_ID_Table.get(tableAPI.ID)
						CopyTableAPIToTable(tableAPI, table!, this.frontRepo)
					}
				)



				observer.next(this.frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getCellUniqueID(id: number): number {
	return 31 * id
}
export function getCellBooleanUniqueID(id: number): number {
	return 37 * id
}
export function getCellFloat64UniqueID(id: number): number {
	return 41 * id
}
export function getCellIconUniqueID(id: number): number {
	return 43 * id
}
export function getCellIntUniqueID(id: number): number {
	return 47 * id
}
export function getCellStringUniqueID(id: number): number {
	return 53 * id
}
export function getCheckBoxUniqueID(id: number): number {
	return 59 * id
}
export function getDisplayedColumnUniqueID(id: number): number {
	return 61 * id
}
export function getFormDivUniqueID(id: number): number {
	return 67 * id
}
export function getFormEditAssocButtonUniqueID(id: number): number {
	return 71 * id
}
export function getFormFieldUniqueID(id: number): number {
	return 73 * id
}
export function getFormFieldDateUniqueID(id: number): number {
	return 79 * id
}
export function getFormFieldDateTimeUniqueID(id: number): number {
	return 83 * id
}
export function getFormFieldFloat64UniqueID(id: number): number {
	return 89 * id
}
export function getFormFieldIntUniqueID(id: number): number {
	return 97 * id
}
export function getFormFieldSelectUniqueID(id: number): number {
	return 101 * id
}
export function getFormFieldStringUniqueID(id: number): number {
	return 103 * id
}
export function getFormFieldTimeUniqueID(id: number): number {
	return 107 * id
}
export function getFormGroupUniqueID(id: number): number {
	return 109 * id
}
export function getFormSortAssocButtonUniqueID(id: number): number {
	return 113 * id
}
export function getOptionUniqueID(id: number): number {
	return 127 * id
}
export function getRowUniqueID(id: number): number {
	return 131 * id
}
export function getTableUniqueID(id: number): number {
	return 137 * id
}

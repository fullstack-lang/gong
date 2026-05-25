import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'
import { ButtonService } from './button.service'

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

import { DisplayedColumnAPI } from './displayedcolumn-api'
import { DisplayedColumn, CopyDisplayedColumnAPIToDisplayedColumn } from './displayedcolumn'
import { DisplayedColumnService } from './displayedcolumn.service'

import { RowAPI } from './row-api'
import { Row, CopyRowAPIToRow } from './row'
import { RowService } from './row.service'

import { SVGIconAPI } from './svgicon-api'
import { SVGIcon, CopySVGIconAPIToSVGIcon } from './svgicon'
import { SVGIconService } from './svgicon.service'

import { TableAPI } from './table-api'
import { Table, CopyTableAPIToTable } from './table'
import { TableService } from './table.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/table/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Buttons = new Array<Button>() // array of front instances
	map_ID_Button = new Map<number, Button>() // map of front instances

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

	array_DisplayedColumns = new Array<DisplayedColumn>() // array of front instances
	map_ID_DisplayedColumn = new Map<number, DisplayedColumn>() // map of front instances

	array_Rows = new Array<Row>() // array of front instances
	map_ID_Row = new Map<number, Row>() // map of front instances

	array_SVGIcons = new Array<SVGIcon>() // array of front instances
	map_ID_SVGIcon = new Map<number, SVGIcon>() // map of front instances

	array_Tables = new Array<Table>() // array of front instances
	map_ID_Table = new Map<number, Table>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.array_Buttons as unknown as Array<Type>
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
			case 'DisplayedColumn':
				return this.array_DisplayedColumns as unknown as Array<Type>
			case 'Row':
				return this.array_Rows as unknown as Array<Type>
			case 'SVGIcon':
				return this.array_SVGIcons as unknown as Array<Type>
			case 'Table':
				return this.array_Tables as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.map_ID_Button as unknown as Map<number, Type>
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
			case 'DisplayedColumn':
				return this.map_ID_DisplayedColumn as unknown as Map<number, Type>
			case 'Row':
				return this.map_ID_Row as unknown as Map<number, Type>
			case 'SVGIcon':
				return this.map_ID_SVGIcon as unknown as Map<number, Type>
			case 'Table':
				return this.map_ID_Table as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized")
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

	Name: string = ""
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

	Name: string = ""

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	}

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	// Manage open WebSocket connections
	private webSocketConnections = new Map<string, Observable<FrontRepo>>()


	constructor(
		private http: HttpClient, // insertion point sub template 
		private buttonService: ButtonService,
		private cellService: CellService,
		private cellbooleanService: CellBooleanService,
		private cellfloat64Service: CellFloat64Service,
		private celliconService: CellIconService,
		private cellintService: CellIntService,
		private cellstringService: CellStringService,
		private displayedcolumnService: DisplayedColumnService,
		private rowService: RowService,
		private svgiconService: SVGIconService,
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
		)
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
		)
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<ButtonAPI[]>,
		Observable<CellAPI[]>,
		Observable<CellBooleanAPI[]>,
		Observable<CellFloat64API[]>,
		Observable<CellIconAPI[]>,
		Observable<CellIntAPI[]>,
		Observable<CellStringAPI[]>,
		Observable<DisplayedColumnAPI[]>,
		Observable<RowAPI[]>,
		Observable<SVGIconAPI[]>,
		Observable<TableAPI[]>,
	]

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(Name: string = ""): Observable<FrontRepo> {

		this.Name = Name

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.buttonService.getButtons(this.Name, this.frontRepo),
			this.cellService.getCells(this.Name, this.frontRepo),
			this.cellbooleanService.getCellBooleans(this.Name, this.frontRepo),
			this.cellfloat64Service.getCellFloat64s(this.Name, this.frontRepo),
			this.celliconService.getCellIcons(this.Name, this.frontRepo),
			this.cellintService.getCellInts(this.Name, this.frontRepo),
			this.cellstringService.getCellStrings(this.Name, this.frontRepo),
			this.displayedcolumnService.getDisplayedColumns(this.Name, this.frontRepo),
			this.rowService.getRows(this.Name, this.frontRepo),
			this.svgiconService.getSVGIcons(this.Name, this.frontRepo),
			this.tableService.getTables(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						buttons_,
						cells_,
						cellbooleans_,
						cellfloat64s_,
						cellicons_,
						cellints_,
						cellstrings_,
						displayedcolumns_,
						rows_,
						svgicons_,
						tables_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var buttons: ButtonAPI[]
						buttons = buttons_ as ButtonAPI[]
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
						var displayedcolumns: DisplayedColumnAPI[]
						displayedcolumns = displayedcolumns_ as DisplayedColumnAPI[]
						var rows: RowAPI[]
						rows = rows_ as RowAPI[]
						var svgicons: SVGIconAPI[]
						svgicons = svgicons_ as SVGIconAPI[]
						var tables: TableAPI[]
						tables = tables_ as TableAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Buttons = []
						this.frontRepo.map_ID_Button.clear()

						buttons.forEach(
							buttonAPI => {
								let button = new Button
								this.frontRepo.array_Buttons.push(button)
								this.frontRepo.map_ID_Button.set(buttonAPI.ID, button)
							}
						)

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
						this.frontRepo.array_SVGIcons = []
						this.frontRepo.map_ID_SVGIcon.clear()

						svgicons.forEach(
							svgiconAPI => {
								let svgicon = new SVGIcon
								this.frontRepo.array_SVGIcons.push(svgicon)
								this.frontRepo.map_ID_SVGIcon.set(svgiconAPI.ID, svgicon)
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
						buttons.forEach(
							buttonAPI => {
								let button = this.frontRepo.map_ID_Button.get(buttonAPI.ID)
								CopyButtonAPIToButton(buttonAPI, button!, this.frontRepo)
							}
						)

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
						displayedcolumns.forEach(
							displayedcolumnAPI => {
								let displayedcolumn = this.frontRepo.map_ID_DisplayedColumn.get(displayedcolumnAPI.ID)
								CopyDisplayedColumnAPIToDisplayedColumn(displayedcolumnAPI, displayedcolumn!, this.frontRepo)
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
						svgicons.forEach(
							svgiconAPI => {
								let svgicon = this.frontRepo.map_ID_SVGIcon.get(svgiconAPI.ID)
								CopySVGIconAPIToSVGIcon(svgiconAPI, svgicon!, this.frontRepo)
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

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		console.log("github.com/fullstack-lang/gong/lib/table/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			console.log("github.com/fullstack-lang/gong/lib/table/go; connectToWebSocket: returning existing connection")
			return this.webSocketConnections.get(Name)!
		}

		//
		// Create a new connection
		//
		let host = window.location.host
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'

		if (host === 'localhost:4200') {
			host = 'localhost:8080'
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/table/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			console.log("github.com/fullstack-lang/gong/lib/table/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:'

			const processData = (dataString: string) => {
				console.log("github.com/fullstack-lang/gong/lib/table/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Buttons = []
				frontRepo.map_ID_Button.clear()

				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = new Button
						frontRepo.array_Buttons.push(button)
						frontRepo.map_ID_Button.set(buttonAPI.ID, button)
					}
				)

				// init the arrays
				frontRepo.array_Cells = []
				frontRepo.map_ID_Cell.clear()

				backRepoData.CellAPIs.forEach(
					cellAPI => {
						let cell = new Cell
						frontRepo.array_Cells.push(cell)
						frontRepo.map_ID_Cell.set(cellAPI.ID, cell)
					}
				)

				// init the arrays
				frontRepo.array_CellBooleans = []
				frontRepo.map_ID_CellBoolean.clear()

				backRepoData.CellBooleanAPIs.forEach(
					cellbooleanAPI => {
						let cellboolean = new CellBoolean
						frontRepo.array_CellBooleans.push(cellboolean)
						frontRepo.map_ID_CellBoolean.set(cellbooleanAPI.ID, cellboolean)
					}
				)

				// init the arrays
				frontRepo.array_CellFloat64s = []
				frontRepo.map_ID_CellFloat64.clear()

				backRepoData.CellFloat64APIs.forEach(
					cellfloat64API => {
						let cellfloat64 = new CellFloat64
						frontRepo.array_CellFloat64s.push(cellfloat64)
						frontRepo.map_ID_CellFloat64.set(cellfloat64API.ID, cellfloat64)
					}
				)

				// init the arrays
				frontRepo.array_CellIcons = []
				frontRepo.map_ID_CellIcon.clear()

				backRepoData.CellIconAPIs.forEach(
					celliconAPI => {
						let cellicon = new CellIcon
						frontRepo.array_CellIcons.push(cellicon)
						frontRepo.map_ID_CellIcon.set(celliconAPI.ID, cellicon)
					}
				)

				// init the arrays
				frontRepo.array_CellInts = []
				frontRepo.map_ID_CellInt.clear()

				backRepoData.CellIntAPIs.forEach(
					cellintAPI => {
						let cellint = new CellInt
						frontRepo.array_CellInts.push(cellint)
						frontRepo.map_ID_CellInt.set(cellintAPI.ID, cellint)
					}
				)

				// init the arrays
				frontRepo.array_CellStrings = []
				frontRepo.map_ID_CellString.clear()

				backRepoData.CellStringAPIs.forEach(
					cellstringAPI => {
						let cellstring = new CellString
						frontRepo.array_CellStrings.push(cellstring)
						frontRepo.map_ID_CellString.set(cellstringAPI.ID, cellstring)
					}
				)

				// init the arrays
				frontRepo.array_DisplayedColumns = []
				frontRepo.map_ID_DisplayedColumn.clear()

				backRepoData.DisplayedColumnAPIs.forEach(
					displayedcolumnAPI => {
						let displayedcolumn = new DisplayedColumn
						frontRepo.array_DisplayedColumns.push(displayedcolumn)
						frontRepo.map_ID_DisplayedColumn.set(displayedcolumnAPI.ID, displayedcolumn)
					}
				)

				// init the arrays
				frontRepo.array_Rows = []
				frontRepo.map_ID_Row.clear()

				backRepoData.RowAPIs.forEach(
					rowAPI => {
						let row = new Row
						frontRepo.array_Rows.push(row)
						frontRepo.map_ID_Row.set(rowAPI.ID, row)
					}
				)

				// init the arrays
				frontRepo.array_SVGIcons = []
				frontRepo.map_ID_SVGIcon.clear()

				backRepoData.SVGIconAPIs.forEach(
					svgiconAPI => {
						let svgicon = new SVGIcon
						frontRepo.array_SVGIcons.push(svgicon)
						frontRepo.map_ID_SVGIcon.set(svgiconAPI.ID, svgicon)
					}
				)

				// init the arrays
				frontRepo.array_Tables = []
				frontRepo.map_ID_Table.clear()

				backRepoData.TableAPIs.forEach(
					tableAPI => {
						let table = new Table
						frontRepo.array_Tables.push(table)
						frontRepo.map_ID_Table.set(tableAPI.ID, table)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = frontRepo.map_ID_Button.get(buttonAPI.ID)
						CopyButtonAPIToButton(buttonAPI, button!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellAPIs.forEach(
					cellAPI => {
						let cell = frontRepo.map_ID_Cell.get(cellAPI.ID)
						CopyCellAPIToCell(cellAPI, cell!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellBooleanAPIs.forEach(
					cellbooleanAPI => {
						let cellboolean = frontRepo.map_ID_CellBoolean.get(cellbooleanAPI.ID)
						CopyCellBooleanAPIToCellBoolean(cellbooleanAPI, cellboolean!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellFloat64APIs.forEach(
					cellfloat64API => {
						let cellfloat64 = frontRepo.map_ID_CellFloat64.get(cellfloat64API.ID)
						CopyCellFloat64APIToCellFloat64(cellfloat64API, cellfloat64!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellIconAPIs.forEach(
					celliconAPI => {
						let cellicon = frontRepo.map_ID_CellIcon.get(celliconAPI.ID)
						CopyCellIconAPIToCellIcon(celliconAPI, cellicon!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellIntAPIs.forEach(
					cellintAPI => {
						let cellint = frontRepo.map_ID_CellInt.get(cellintAPI.ID)
						CopyCellIntAPIToCellInt(cellintAPI, cellint!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CellStringAPIs.forEach(
					cellstringAPI => {
						let cellstring = frontRepo.map_ID_CellString.get(cellstringAPI.ID)
						CopyCellStringAPIToCellString(cellstringAPI, cellstring!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DisplayedColumnAPIs.forEach(
					displayedcolumnAPI => {
						let displayedcolumn = frontRepo.map_ID_DisplayedColumn.get(displayedcolumnAPI.ID)
						CopyDisplayedColumnAPIToDisplayedColumn(displayedcolumnAPI, displayedcolumn!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.RowAPIs.forEach(
					rowAPI => {
						let row = frontRepo.map_ID_Row.get(rowAPI.ID)
						CopyRowAPIToRow(rowAPI, row!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SVGIconAPIs.forEach(
					svgiconAPI => {
						let svgicon = frontRepo.map_ID_SVGIcon.get(svgiconAPI.ID)
						CopySVGIconAPIToSVGIcon(svgiconAPI, svgicon!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TableAPIs.forEach(
					tableAPI => {
						let table = frontRepo.map_ID_Table.get(tableAPI.ID)
						CopyTableAPIToTable(tableAPI, table!, frontRepo)
					}
				)


				observer.next(frontRepo)
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				console.log("github.com/fullstack-lang/gong/lib/table/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					console.log("github.com/fullstack-lang/gong/lib/table/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket(Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					console.log("github.com/fullstack-lang/gong/lib/table/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					console.log("github.com/fullstack-lang/gong/lib/table/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						console.log("github.com/fullstack-lang/gong/lib/table/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						console.log("github.com/fullstack-lang/gong/lib/table/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/table/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						console.log("github.com/fullstack-lang/gong/lib/table/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/table/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.");
				}
			};

			attemptConnection(50);

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				if (socket) {
					socket.close()
				}
			}
		}).pipe(
			// This is the key:
			// - shareReplay makes this a "multicast" observable, sharing the single WebSocket among subscribers.
			// - { bufferSize: 1, refCount: true } means:
			//   - bufferSize: 1 => new subscribers get the last emitted value immediately.
			//   - refCount: true => the connection starts with the first subscriber and stops with the last.
			shareReplay({ bufferSize: 1, refCount: true })
		)

		// Store the new connection observable in the map
		this.webSocketConnections.set(Name, newConnection$)
		return newConnection$
	}
}

// insertion point for get unique ID per struct 
export function getButtonUniqueID(id: number): number {
	return 31 * id
}
export function getCellUniqueID(id: number): number {
	return 37 * id
}
export function getCellBooleanUniqueID(id: number): number {
	return 41 * id
}
export function getCellFloat64UniqueID(id: number): number {
	return 43 * id
}
export function getCellIconUniqueID(id: number): number {
	return 47 * id
}
export function getCellIntUniqueID(id: number): number {
	return 53 * id
}
export function getCellStringUniqueID(id: number): number {
	return 59 * id
}
export function getDisplayedColumnUniqueID(id: number): number {
	return 61 * id
}
export function getRowUniqueID(id: number): number {
	return 67 * id
}
export function getSVGIconUniqueID(id: number): number {
	return 71 * id
}
export function getTableUniqueID(id: number): number {
	return 73 * id
}

import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { DisplaySelectionAPI } from './displayselection-api'
import { DisplaySelection, CopyDisplaySelectionAPIToDisplaySelection } from './displayselection'
import { DisplaySelectionService } from './displayselection.service'

import { XLCellAPI } from './xlcell-api'
import { XLCell, CopyXLCellAPIToXLCell } from './xlcell'
import { XLCellService } from './xlcell.service'

import { XLFileAPI } from './xlfile-api'
import { XLFile, CopyXLFileAPIToXLFile } from './xlfile'
import { XLFileService } from './xlfile.service'

import { XLRowAPI } from './xlrow-api'
import { XLRow, CopyXLRowAPIToXLRow } from './xlrow'
import { XLRowService } from './xlrow.service'

import { XLSheetAPI } from './xlsheet-api'
import { XLSheet, CopyXLSheetAPIToXLSheet } from './xlsheet'
import { XLSheetService } from './xlsheet.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/xlsx/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_DisplaySelections = new Array<DisplaySelection>() // array of front instances
	map_ID_DisplaySelection = new Map<number, DisplaySelection>() // map of front instances

	array_XLCells = new Array<XLCell>() // array of front instances
	map_ID_XLCell = new Map<number, XLCell>() // map of front instances

	array_XLFiles = new Array<XLFile>() // array of front instances
	map_ID_XLFile = new Map<number, XLFile>() // map of front instances

	array_XLRows = new Array<XLRow>() // array of front instances
	map_ID_XLRow = new Map<number, XLRow>() // map of front instances

	array_XLSheets = new Array<XLSheet>() // array of front instances
	map_ID_XLSheet = new Map<number, XLSheet>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'DisplaySelection':
				return this.array_DisplaySelections as unknown as Array<Type>
			case 'XLCell':
				return this.array_XLCells as unknown as Array<Type>
			case 'XLFile':
				return this.array_XLFiles as unknown as Array<Type>
			case 'XLRow':
				return this.array_XLRows as unknown as Array<Type>
			case 'XLSheet':
				return this.array_XLSheets as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'DisplaySelection':
				return this.map_ID_DisplaySelection as unknown as Map<number, Type>
			case 'XLCell':
				return this.map_ID_XLCell as unknown as Map<number, Type>
			case 'XLFile':
				return this.map_ID_XLFile as unknown as Map<number, Type>
			case 'XLRow':
				return this.map_ID_XLRow as unknown as Map<number, Type>
			case 'XLSheet':
				return this.map_ID_XLSheet as unknown as Map<number, Type>
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
		private displayselectionService: DisplaySelectionService,
		private xlcellService: XLCellService,
		private xlfileService: XLFileService,
		private xlrowService: XLRowService,
		private xlsheetService: XLSheetService,
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
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<DisplaySelectionAPI[]>,
		Observable<XLCellAPI[]>,
		Observable<XLFileAPI[]>,
		Observable<XLRowAPI[]>,
		Observable<XLSheetAPI[]>,
	];

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
			this.displayselectionService.getDisplaySelections(this.Name, this.frontRepo),
			this.xlcellService.getXLCells(this.Name, this.frontRepo),
			this.xlfileService.getXLFiles(this.Name, this.frontRepo),
			this.xlrowService.getXLRows(this.Name, this.frontRepo),
			this.xlsheetService.getXLSheets(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						displayselections_,
						xlcells_,
						xlfiles_,
						xlrows_,
						xlsheets_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var displayselections: DisplaySelectionAPI[]
						displayselections = displayselections_ as DisplaySelectionAPI[]
						var xlcells: XLCellAPI[]
						xlcells = xlcells_ as XLCellAPI[]
						var xlfiles: XLFileAPI[]
						xlfiles = xlfiles_ as XLFileAPI[]
						var xlrows: XLRowAPI[]
						xlrows = xlrows_ as XLRowAPI[]
						var xlsheets: XLSheetAPI[]
						xlsheets = xlsheets_ as XLSheetAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_DisplaySelections = []
						this.frontRepo.map_ID_DisplaySelection.clear()

						displayselections.forEach(
							displayselectionAPI => {
								let displayselection = new DisplaySelection
								this.frontRepo.array_DisplaySelections.push(displayselection)
								this.frontRepo.map_ID_DisplaySelection.set(displayselectionAPI.ID, displayselection)
							}
						)

						// init the arrays
						this.frontRepo.array_XLCells = []
						this.frontRepo.map_ID_XLCell.clear()

						xlcells.forEach(
							xlcellAPI => {
								let xlcell = new XLCell
								this.frontRepo.array_XLCells.push(xlcell)
								this.frontRepo.map_ID_XLCell.set(xlcellAPI.ID, xlcell)
							}
						)

						// init the arrays
						this.frontRepo.array_XLFiles = []
						this.frontRepo.map_ID_XLFile.clear()

						xlfiles.forEach(
							xlfileAPI => {
								let xlfile = new XLFile
								this.frontRepo.array_XLFiles.push(xlfile)
								this.frontRepo.map_ID_XLFile.set(xlfileAPI.ID, xlfile)
							}
						)

						// init the arrays
						this.frontRepo.array_XLRows = []
						this.frontRepo.map_ID_XLRow.clear()

						xlrows.forEach(
							xlrowAPI => {
								let xlrow = new XLRow
								this.frontRepo.array_XLRows.push(xlrow)
								this.frontRepo.map_ID_XLRow.set(xlrowAPI.ID, xlrow)
							}
						)

						// init the arrays
						this.frontRepo.array_XLSheets = []
						this.frontRepo.map_ID_XLSheet.clear()

						xlsheets.forEach(
							xlsheetAPI => {
								let xlsheet = new XLSheet
								this.frontRepo.array_XLSheets.push(xlsheet)
								this.frontRepo.map_ID_XLSheet.set(xlsheetAPI.ID, xlsheet)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						displayselections.forEach(
							displayselectionAPI => {
								let displayselection = this.frontRepo.map_ID_DisplaySelection.get(displayselectionAPI.ID)
								CopyDisplaySelectionAPIToDisplaySelection(displayselectionAPI, displayselection!, this.frontRepo)
							}
						)

						// fill up front objects
						xlcells.forEach(
							xlcellAPI => {
								let xlcell = this.frontRepo.map_ID_XLCell.get(xlcellAPI.ID)
								CopyXLCellAPIToXLCell(xlcellAPI, xlcell!, this.frontRepo)
							}
						)

						// fill up front objects
						xlfiles.forEach(
							xlfileAPI => {
								let xlfile = this.frontRepo.map_ID_XLFile.get(xlfileAPI.ID)
								CopyXLFileAPIToXLFile(xlfileAPI, xlfile!, this.frontRepo)
							}
						)

						// fill up front objects
						xlrows.forEach(
							xlrowAPI => {
								let xlrow = this.frontRepo.map_ID_XLRow.get(xlrowAPI.ID)
								CopyXLRowAPIToXLRow(xlrowAPI, xlrow!, this.frontRepo)
							}
						)

						// fill up front objects
						xlsheets.forEach(
							xlsheetAPI => {
								let xlsheet = this.frontRepo.map_ID_XLSheet.get(xlsheetAPI.ID)
								CopyXLSheetAPIToXLSheet(xlsheetAPI, xlsheet!, this.frontRepo)
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

		this.Name = Name


		let params = new HttpParams().set("Name", this.Name)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/lib/xlsx/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_DisplaySelections = []
				frontRepo.map_ID_DisplaySelection.clear()

				backRepoData.DisplaySelectionAPIs.forEach(
					displayselectionAPI => {
						let displayselection = new DisplaySelection
						frontRepo.array_DisplaySelections.push(displayselection)
						frontRepo.map_ID_DisplaySelection.set(displayselectionAPI.ID, displayselection)
					}
				)

				// init the arrays
				frontRepo.array_XLCells = []
				frontRepo.map_ID_XLCell.clear()

				backRepoData.XLCellAPIs.forEach(
					xlcellAPI => {
						let xlcell = new XLCell
						frontRepo.array_XLCells.push(xlcell)
						frontRepo.map_ID_XLCell.set(xlcellAPI.ID, xlcell)
					}
				)

				// init the arrays
				frontRepo.array_XLFiles = []
				frontRepo.map_ID_XLFile.clear()

				backRepoData.XLFileAPIs.forEach(
					xlfileAPI => {
						let xlfile = new XLFile
						frontRepo.array_XLFiles.push(xlfile)
						frontRepo.map_ID_XLFile.set(xlfileAPI.ID, xlfile)
					}
				)

				// init the arrays
				frontRepo.array_XLRows = []
				frontRepo.map_ID_XLRow.clear()

				backRepoData.XLRowAPIs.forEach(
					xlrowAPI => {
						let xlrow = new XLRow
						frontRepo.array_XLRows.push(xlrow)
						frontRepo.map_ID_XLRow.set(xlrowAPI.ID, xlrow)
					}
				)

				// init the arrays
				frontRepo.array_XLSheets = []
				frontRepo.map_ID_XLSheet.clear()

				backRepoData.XLSheetAPIs.forEach(
					xlsheetAPI => {
						let xlsheet = new XLSheet
						frontRepo.array_XLSheets.push(xlsheet)
						frontRepo.map_ID_XLSheet.set(xlsheetAPI.ID, xlsheet)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.DisplaySelectionAPIs.forEach(
					displayselectionAPI => {
						let displayselection = frontRepo.map_ID_DisplaySelection.get(displayselectionAPI.ID)
						CopyDisplaySelectionAPIToDisplaySelection(displayselectionAPI, displayselection!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.XLCellAPIs.forEach(
					xlcellAPI => {
						let xlcell = frontRepo.map_ID_XLCell.get(xlcellAPI.ID)
						CopyXLCellAPIToXLCell(xlcellAPI, xlcell!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.XLFileAPIs.forEach(
					xlfileAPI => {
						let xlfile = frontRepo.map_ID_XLFile.get(xlfileAPI.ID)
						CopyXLFileAPIToXLFile(xlfileAPI, xlfile!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.XLRowAPIs.forEach(
					xlrowAPI => {
						let xlrow = frontRepo.map_ID_XLRow.get(xlrowAPI.ID)
						CopyXLRowAPIToXLRow(xlrowAPI, xlrow!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.XLSheetAPIs.forEach(
					xlsheetAPI => {
						let xlsheet = frontRepo.map_ID_XLSheet.get(xlsheetAPI.ID)
						CopyXLSheetAPIToXLSheet(xlsheetAPI, xlsheet!, frontRepo)
					}
				)



				observer.next(frontRepo)
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
export function getDisplaySelectionUniqueID(id: number): number {
	return 31 * id
}
export function getXLCellUniqueID(id: number): number {
	return 37 * id
}
export function getXLFileUniqueID(id: number): number {
	return 41 * id
}
export function getXLRowUniqueID(id: number): number {
	return 43 * id
}
export function getXLSheetUniqueID(id: number): number {
	return 47 * id
}

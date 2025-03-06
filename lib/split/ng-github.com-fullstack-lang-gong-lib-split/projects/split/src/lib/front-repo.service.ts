import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AsSplitAPI } from './assplit-api'
import { AsSplit, CopyAsSplitAPIToAsSplit } from './assplit'
import { AsSplitService } from './assplit.service'

import { AsSplitAreaAPI } from './assplitarea-api'
import { AsSplitArea, CopyAsSplitAreaAPIToAsSplitArea } from './assplitarea'
import { AsSplitAreaService } from './assplitarea.service'

import { ViewAPI } from './view-api'
import { View, CopyViewAPIToView } from './view'
import { ViewService } from './view.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/split/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_AsSplits = new Array<AsSplit>() // array of front instances
	map_ID_AsSplit = new Map<number, AsSplit>() // map of front instances

	array_AsSplitAreas = new Array<AsSplitArea>() // array of front instances
	map_ID_AsSplitArea = new Map<number, AsSplitArea>() // map of front instances

	array_Views = new Array<View>() // array of front instances
	map_ID_View = new Map<number, View>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'AsSplit':
				return this.array_AsSplits as unknown as Array<Type>
			case 'AsSplitArea':
				return this.array_AsSplitAreas as unknown as Array<Type>
			case 'View':
				return this.array_Views as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'AsSplit':
				return this.map_ID_AsSplit as unknown as Map<number, Type>
			case 'AsSplitArea':
				return this.map_ID_AsSplitArea as unknown as Map<number, Type>
			case 'View':
				return this.map_ID_View as unknown as Map<number, Type>
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
		private assplitService: AsSplitService,
		private assplitareaService: AsSplitAreaService,
		private viewService: ViewService,
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
		Observable<AsSplitAPI[]>,
		Observable<AsSplitAreaAPI[]>,
		Observable<ViewAPI[]>,
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
			this.assplitService.getAsSplits(this.GONG__StackPath, this.frontRepo),
			this.assplitareaService.getAsSplitAreas(this.GONG__StackPath, this.frontRepo),
			this.viewService.getViews(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						assplits_,
						assplitareas_,
						views_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var assplits: AsSplitAPI[]
						assplits = assplits_ as AsSplitAPI[]
						var assplitareas: AsSplitAreaAPI[]
						assplitareas = assplitareas_ as AsSplitAreaAPI[]
						var views: ViewAPI[]
						views = views_ as ViewAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_AsSplits = []
						this.frontRepo.map_ID_AsSplit.clear()

						assplits.forEach(
							assplitAPI => {
								let assplit = new AsSplit
								this.frontRepo.array_AsSplits.push(assplit)
								this.frontRepo.map_ID_AsSplit.set(assplitAPI.ID, assplit)
							}
						)

						// init the arrays
						this.frontRepo.array_AsSplitAreas = []
						this.frontRepo.map_ID_AsSplitArea.clear()

						assplitareas.forEach(
							assplitareaAPI => {
								let assplitarea = new AsSplitArea
								this.frontRepo.array_AsSplitAreas.push(assplitarea)
								this.frontRepo.map_ID_AsSplitArea.set(assplitareaAPI.ID, assplitarea)
							}
						)

						// init the arrays
						this.frontRepo.array_Views = []
						this.frontRepo.map_ID_View.clear()

						views.forEach(
							viewAPI => {
								let view = new View
								this.frontRepo.array_Views.push(view)
								this.frontRepo.map_ID_View.set(viewAPI.ID, view)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						assplits.forEach(
							assplitAPI => {
								let assplit = this.frontRepo.map_ID_AsSplit.get(assplitAPI.ID)
								CopyAsSplitAPIToAsSplit(assplitAPI, assplit!, this.frontRepo)
							}
						)

						// fill up front objects
						assplitareas.forEach(
							assplitareaAPI => {
								let assplitarea = this.frontRepo.map_ID_AsSplitArea.get(assplitareaAPI.ID)
								CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea!, this.frontRepo)
							}
						)

						// fill up front objects
						views.forEach(
							viewAPI => {
								let view = this.frontRepo.map_ID_View.get(viewAPI.ID)
								CopyViewAPIToView(viewAPI, view!, this.frontRepo)
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
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/lib/split/go/v1/ws/stage'
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
				frontRepo.array_AsSplits = []
				frontRepo.map_ID_AsSplit.clear()

				backRepoData.AsSplitAPIs.forEach(
					assplitAPI => {
						let assplit = new AsSplit
						frontRepo.array_AsSplits.push(assplit)
						frontRepo.map_ID_AsSplit.set(assplitAPI.ID, assplit)
					}
				)

				// init the arrays
				frontRepo.array_AsSplitAreas = []
				frontRepo.map_ID_AsSplitArea.clear()

				backRepoData.AsSplitAreaAPIs.forEach(
					assplitareaAPI => {
						let assplitarea = new AsSplitArea
						frontRepo.array_AsSplitAreas.push(assplitarea)
						frontRepo.map_ID_AsSplitArea.set(assplitareaAPI.ID, assplitarea)
					}
				)

				// init the arrays
				frontRepo.array_Views = []
				frontRepo.map_ID_View.clear()

				backRepoData.ViewAPIs.forEach(
					viewAPI => {
						let view = new View
						frontRepo.array_Views.push(view)
						frontRepo.map_ID_View.set(viewAPI.ID, view)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AsSplitAPIs.forEach(
					assplitAPI => {
						let assplit = frontRepo.map_ID_AsSplit.get(assplitAPI.ID)
						CopyAsSplitAPIToAsSplit(assplitAPI, assplit!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AsSplitAreaAPIs.forEach(
					assplitareaAPI => {
						let assplitarea = frontRepo.map_ID_AsSplitArea.get(assplitareaAPI.ID)
						CopyAsSplitAreaAPIToAsSplitArea(assplitareaAPI, assplitarea!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ViewAPIs.forEach(
					viewAPI => {
						let view = frontRepo.map_ID_View.get(viewAPI.ID)
						CopyViewAPIToView(viewAPI, view!, frontRepo)
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
export function getAsSplitUniqueID(id: number): number {
	return 31 * id
}
export function getAsSplitAreaUniqueID(id: number): number {
	return 37 * id
}
export function getViewUniqueID(id: number): number {
	return 41 * id
}

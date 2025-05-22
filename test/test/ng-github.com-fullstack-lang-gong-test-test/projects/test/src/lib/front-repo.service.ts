import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AstructAPI } from './astruct-api'
import { Astruct, CopyAstructAPIToAstruct } from './astruct'
import { AstructService } from './astruct.service'

import { AstructBstruct2UseAPI } from './astructbstruct2use-api'
import { AstructBstruct2Use, CopyAstructBstruct2UseAPIToAstructBstruct2Use } from './astructbstruct2use'
import { AstructBstruct2UseService } from './astructbstruct2use.service'

import { AstructBstructUseAPI } from './astructbstructuse-api'
import { AstructBstructUse, CopyAstructBstructUseAPIToAstructBstructUse } from './astructbstructuse'
import { AstructBstructUseService } from './astructbstructuse.service'

import { BstructAPI } from './bstruct-api'
import { Bstruct, CopyBstructAPIToBstruct } from './bstruct'
import { BstructService } from './bstruct.service'

import { DstructAPI } from './dstruct-api'
import { Dstruct, CopyDstructAPIToDstruct } from './dstruct'
import { DstructService } from './dstruct.service'

import { GstructAPI } from './gstruct-api'
import { Gstruct, CopyGstructAPIToGstruct } from './gstruct'
import { GstructService } from './gstruct.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/test/test/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Astructs = new Array<Astruct>() // array of front instances
	map_ID_Astruct = new Map<number, Astruct>() // map of front instances

	array_AstructBstruct2Uses = new Array<AstructBstruct2Use>() // array of front instances
	map_ID_AstructBstruct2Use = new Map<number, AstructBstruct2Use>() // map of front instances

	array_AstructBstructUses = new Array<AstructBstructUse>() // array of front instances
	map_ID_AstructBstructUse = new Map<number, AstructBstructUse>() // map of front instances

	array_Bstructs = new Array<Bstruct>() // array of front instances
	map_ID_Bstruct = new Map<number, Bstruct>() // map of front instances

	array_Dstructs = new Array<Dstruct>() // array of front instances
	map_ID_Dstruct = new Map<number, Dstruct>() // map of front instances

	array_Gstructs = new Array<Gstruct>() // array of front instances
	map_ID_Gstruct = new Map<number, Gstruct>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Astruct':
				return this.array_Astructs as unknown as Array<Type>
			case 'AstructBstruct2Use':
				return this.array_AstructBstruct2Uses as unknown as Array<Type>
			case 'AstructBstructUse':
				return this.array_AstructBstructUses as unknown as Array<Type>
			case 'Bstruct':
				return this.array_Bstructs as unknown as Array<Type>
			case 'Dstruct':
				return this.array_Dstructs as unknown as Array<Type>
			case 'Gstruct':
				return this.array_Gstructs as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Astruct':
				return this.map_ID_Astruct as unknown as Map<number, Type>
			case 'AstructBstruct2Use':
				return this.map_ID_AstructBstruct2Use as unknown as Map<number, Type>
			case 'AstructBstructUse':
				return this.map_ID_AstructBstructUse as unknown as Map<number, Type>
			case 'Bstruct':
				return this.map_ID_Bstruct as unknown as Map<number, Type>
			case 'Dstruct':
				return this.map_ID_Dstruct as unknown as Map<number, Type>
			case 'Gstruct':
				return this.map_ID_Gstruct as unknown as Map<number, Type>
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
		private astructService: AstructService,
		private astructbstruct2useService: AstructBstruct2UseService,
		private astructbstructuseService: AstructBstructUseService,
		private bstructService: BstructService,
		private dstructService: DstructService,
		private gstructService: GstructService,
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
		Observable<AstructAPI[]>,
		Observable<AstructBstruct2UseAPI[]>,
		Observable<AstructBstructUseAPI[]>,
		Observable<BstructAPI[]>,
		Observable<DstructAPI[]>,
		Observable<GstructAPI[]>,
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
			this.astructService.getAstructs(this.Name, this.frontRepo),
			this.astructbstruct2useService.getAstructBstruct2Uses(this.Name, this.frontRepo),
			this.astructbstructuseService.getAstructBstructUses(this.Name, this.frontRepo),
			this.bstructService.getBstructs(this.Name, this.frontRepo),
			this.dstructService.getDstructs(this.Name, this.frontRepo),
			this.gstructService.getGstructs(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						astructs_,
						astructbstruct2uses_,
						astructbstructuses_,
						bstructs_,
						dstructs_,
						gstructs_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var astructs: AstructAPI[]
						astructs = astructs_ as AstructAPI[]
						var astructbstruct2uses: AstructBstruct2UseAPI[]
						astructbstruct2uses = astructbstruct2uses_ as AstructBstruct2UseAPI[]
						var astructbstructuses: AstructBstructUseAPI[]
						astructbstructuses = astructbstructuses_ as AstructBstructUseAPI[]
						var bstructs: BstructAPI[]
						bstructs = bstructs_ as BstructAPI[]
						var dstructs: DstructAPI[]
						dstructs = dstructs_ as DstructAPI[]
						var gstructs: GstructAPI[]
						gstructs = gstructs_ as GstructAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Astructs = []
						this.frontRepo.map_ID_Astruct.clear()

						astructs.forEach(
							astructAPI => {
								let astruct = new Astruct
								this.frontRepo.array_Astructs.push(astruct)
								this.frontRepo.map_ID_Astruct.set(astructAPI.ID, astruct)
							}
						)

						// init the arrays
						this.frontRepo.array_AstructBstruct2Uses = []
						this.frontRepo.map_ID_AstructBstruct2Use.clear()

						astructbstruct2uses.forEach(
							astructbstruct2useAPI => {
								let astructbstruct2use = new AstructBstruct2Use
								this.frontRepo.array_AstructBstruct2Uses.push(astructbstruct2use)
								this.frontRepo.map_ID_AstructBstruct2Use.set(astructbstruct2useAPI.ID, astructbstruct2use)
							}
						)

						// init the arrays
						this.frontRepo.array_AstructBstructUses = []
						this.frontRepo.map_ID_AstructBstructUse.clear()

						astructbstructuses.forEach(
							astructbstructuseAPI => {
								let astructbstructuse = new AstructBstructUse
								this.frontRepo.array_AstructBstructUses.push(astructbstructuse)
								this.frontRepo.map_ID_AstructBstructUse.set(astructbstructuseAPI.ID, astructbstructuse)
							}
						)

						// init the arrays
						this.frontRepo.array_Bstructs = []
						this.frontRepo.map_ID_Bstruct.clear()

						bstructs.forEach(
							bstructAPI => {
								let bstruct = new Bstruct
								this.frontRepo.array_Bstructs.push(bstruct)
								this.frontRepo.map_ID_Bstruct.set(bstructAPI.ID, bstruct)
							}
						)

						// init the arrays
						this.frontRepo.array_Dstructs = []
						this.frontRepo.map_ID_Dstruct.clear()

						dstructs.forEach(
							dstructAPI => {
								let dstruct = new Dstruct
								this.frontRepo.array_Dstructs.push(dstruct)
								this.frontRepo.map_ID_Dstruct.set(dstructAPI.ID, dstruct)
							}
						)

						// init the arrays
						this.frontRepo.array_Gstructs = []
						this.frontRepo.map_ID_Gstruct.clear()

						gstructs.forEach(
							gstructAPI => {
								let gstruct = new Gstruct
								this.frontRepo.array_Gstructs.push(gstruct)
								this.frontRepo.map_ID_Gstruct.set(gstructAPI.ID, gstruct)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						astructs.forEach(
							astructAPI => {
								let astruct = this.frontRepo.map_ID_Astruct.get(astructAPI.ID)
								CopyAstructAPIToAstruct(astructAPI, astruct!, this.frontRepo)
							}
						)

						// fill up front objects
						astructbstruct2uses.forEach(
							astructbstruct2useAPI => {
								let astructbstruct2use = this.frontRepo.map_ID_AstructBstruct2Use.get(astructbstruct2useAPI.ID)
								CopyAstructBstruct2UseAPIToAstructBstruct2Use(astructbstruct2useAPI, astructbstruct2use!, this.frontRepo)
							}
						)

						// fill up front objects
						astructbstructuses.forEach(
							astructbstructuseAPI => {
								let astructbstructuse = this.frontRepo.map_ID_AstructBstructUse.get(astructbstructuseAPI.ID)
								CopyAstructBstructUseAPIToAstructBstructUse(astructbstructuseAPI, astructbstructuse!, this.frontRepo)
							}
						)

						// fill up front objects
						bstructs.forEach(
							bstructAPI => {
								let bstruct = this.frontRepo.map_ID_Bstruct.get(bstructAPI.ID)
								CopyBstructAPIToBstruct(bstructAPI, bstruct!, this.frontRepo)
							}
						)

						// fill up front objects
						dstructs.forEach(
							dstructAPI => {
								let dstruct = this.frontRepo.map_ID_Dstruct.get(dstructAPI.ID)
								CopyDstructAPIToDstruct(dstructAPI, dstruct!, this.frontRepo)
							}
						)

						// fill up front objects
						gstructs.forEach(
							gstructAPI => {
								let gstruct = this.frontRepo.map_ID_Gstruct.get(gstructAPI.ID)
								CopyGstructAPIToGstruct(gstructAPI, gstruct!, this.frontRepo)
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


		// Determine the base URL for the WebSocket connection dynamically
		// window.location.host includes hostname and port (e.g., "localhost:8080" or "yourdomain.com:8090")
		// If running on standard ports (80 for http, 443 for https), the port might not be explicitly in window.location.host
		// but WebSocket constructor handles 'ws://' and 'wss://' correctly with host.
		let host = window.location.host; // e.g., localhost:4200 or myapp.com
		const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'; // Use wss for https, ws for http

		// Check if the host is localhost:4200 and change it to localhost:8080 (when using ng serve)
		if (host === 'localhost:4200') {
			host = 'localhost:8080';
		}

		// Construct the base path using the dynamic host and protocol
		// The API path remains the same.
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/test/test/go/v1/ws/stage`

		let params = new HttpParams().set("Name", this.Name)
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
				frontRepo.array_Astructs = []
				frontRepo.map_ID_Astruct.clear()

				backRepoData.AstructAPIs.forEach(
					astructAPI => {
						let astruct = new Astruct
						frontRepo.array_Astructs.push(astruct)
						frontRepo.map_ID_Astruct.set(astructAPI.ID, astruct)
					}
				)

				// init the arrays
				frontRepo.array_AstructBstruct2Uses = []
				frontRepo.map_ID_AstructBstruct2Use.clear()

				backRepoData.AstructBstruct2UseAPIs.forEach(
					astructbstruct2useAPI => {
						let astructbstruct2use = new AstructBstruct2Use
						frontRepo.array_AstructBstruct2Uses.push(astructbstruct2use)
						frontRepo.map_ID_AstructBstruct2Use.set(astructbstruct2useAPI.ID, astructbstruct2use)
					}
				)

				// init the arrays
				frontRepo.array_AstructBstructUses = []
				frontRepo.map_ID_AstructBstructUse.clear()

				backRepoData.AstructBstructUseAPIs.forEach(
					astructbstructuseAPI => {
						let astructbstructuse = new AstructBstructUse
						frontRepo.array_AstructBstructUses.push(astructbstructuse)
						frontRepo.map_ID_AstructBstructUse.set(astructbstructuseAPI.ID, astructbstructuse)
					}
				)

				// init the arrays
				frontRepo.array_Bstructs = []
				frontRepo.map_ID_Bstruct.clear()

				backRepoData.BstructAPIs.forEach(
					bstructAPI => {
						let bstruct = new Bstruct
						frontRepo.array_Bstructs.push(bstruct)
						frontRepo.map_ID_Bstruct.set(bstructAPI.ID, bstruct)
					}
				)

				// init the arrays
				frontRepo.array_Dstructs = []
				frontRepo.map_ID_Dstruct.clear()

				backRepoData.DstructAPIs.forEach(
					dstructAPI => {
						let dstruct = new Dstruct
						frontRepo.array_Dstructs.push(dstruct)
						frontRepo.map_ID_Dstruct.set(dstructAPI.ID, dstruct)
					}
				)

				// init the arrays
				frontRepo.array_Gstructs = []
				frontRepo.map_ID_Gstruct.clear()

				backRepoData.GstructAPIs.forEach(
					gstructAPI => {
						let gstruct = new Gstruct
						frontRepo.array_Gstructs.push(gstruct)
						frontRepo.map_ID_Gstruct.set(gstructAPI.ID, gstruct)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AstructAPIs.forEach(
					astructAPI => {
						let astruct = frontRepo.map_ID_Astruct.get(astructAPI.ID)
						CopyAstructAPIToAstruct(astructAPI, astruct!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AstructBstruct2UseAPIs.forEach(
					astructbstruct2useAPI => {
						let astructbstruct2use = frontRepo.map_ID_AstructBstruct2Use.get(astructbstruct2useAPI.ID)
						CopyAstructBstruct2UseAPIToAstructBstruct2Use(astructbstruct2useAPI, astructbstruct2use!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AstructBstructUseAPIs.forEach(
					astructbstructuseAPI => {
						let astructbstructuse = frontRepo.map_ID_AstructBstructUse.get(astructbstructuseAPI.ID)
						CopyAstructBstructUseAPIToAstructBstructUse(astructbstructuseAPI, astructbstructuse!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BstructAPIs.forEach(
					bstructAPI => {
						let bstruct = frontRepo.map_ID_Bstruct.get(bstructAPI.ID)
						CopyBstructAPIToBstruct(bstructAPI, bstruct!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DstructAPIs.forEach(
					dstructAPI => {
						let dstruct = frontRepo.map_ID_Dstruct.get(dstructAPI.ID)
						CopyDstructAPIToDstruct(dstructAPI, dstruct!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GstructAPIs.forEach(
					gstructAPI => {
						let gstruct = frontRepo.map_ID_Gstruct.get(gstructAPI.ID)
						CopyGstructAPIToGstruct(gstructAPI, gstruct!, frontRepo)
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
export function getAstructUniqueID(id: number): number {
	return 31 * id
}
export function getAstructBstruct2UseUniqueID(id: number): number {
	return 37 * id
}
export function getAstructBstructUseUniqueID(id: number): number {
	return 41 * id
}
export function getBstructUniqueID(id: number): number {
	return 43 * id
}
export function getDstructUniqueID(id: number): number {
	return 47 * id
}
export function getGstructUniqueID(id: number): number {
	return 53 * id
}

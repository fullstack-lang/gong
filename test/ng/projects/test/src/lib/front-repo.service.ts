import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AstructDB } from './astruct-db'
import { Astruct, CopyAstructDBToAstruct } from './astruct'
import { AstructService } from './astruct.service'

import { AstructBstruct2UseDB } from './astructbstruct2use-db'
import { AstructBstruct2Use, CopyAstructBstruct2UseDBToAstructBstruct2Use } from './astructbstruct2use'
import { AstructBstruct2UseService } from './astructbstruct2use.service'

import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstructUse, CopyAstructBstructUseDBToAstructBstructUse } from './astructbstructuse'
import { AstructBstructUseService } from './astructbstructuse.service'

import { BstructDB } from './bstruct-db'
import { Bstruct, CopyBstructDBToBstruct } from './bstruct'
import { BstructService } from './bstruct.service'

import { DstructDB } from './dstruct-db'
import { Dstruct, CopyDstructDBToDstruct } from './dstruct'
import { DstructService } from './dstruct.service'
import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/test/go/models"

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
		private astructService: AstructService,
		private astructbstruct2useService: AstructBstruct2UseService,
		private astructbstructuseService: AstructBstructUseService,
		private bstructService: BstructService,
		private dstructService: DstructService,
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
		Observable<AstructDB[]>,
		Observable<AstructBstruct2UseDB[]>,
		Observable<AstructBstructUseDB[]>,
		Observable<BstructDB[]>,
		Observable<DstructDB[]>,
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
			this.astructService.getAstructs(this.GONG__StackPath, this.frontRepo),
			this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath, this.frontRepo),
			this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath, this.frontRepo),
			this.bstructService.getBstructs(this.GONG__StackPath, this.frontRepo),
			this.dstructService.getDstructs(this.GONG__StackPath, this.frontRepo),
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
			this.astructService.getAstructs(this.GONG__StackPath, this.frontRepo),
			this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath, this.frontRepo),
			this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath, this.frontRepo),
			this.bstructService.getBstructs(this.GONG__StackPath, this.frontRepo),
			this.dstructService.getDstructs(this.GONG__StackPath, this.frontRepo),
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
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var astructs: AstructDB[]
						astructs = astructs_ as AstructDB[]
						var astructbstruct2uses: AstructBstruct2UseDB[]
						astructbstruct2uses = astructbstruct2uses_ as AstructBstruct2UseDB[]
						var astructbstructuses: AstructBstructUseDB[]
						astructbstructuses = astructbstructuses_ as AstructBstructUseDB[]
						var bstructs: BstructDB[]
						bstructs = bstructs_ as BstructDB[]
						var dstructs: DstructDB[]
						dstructs = dstructs_ as DstructDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Astructs = []
						this.frontRepo.map_ID_Astruct.clear()

						astructs.forEach(
							astructDB => {
								let astruct = new Astruct
								this.frontRepo.array_Astructs.push(astruct)
								this.frontRepo.map_ID_Astruct.set(astructDB.ID, astruct)
							}
						)

						// init the arrays
						this.frontRepo.array_AstructBstruct2Uses = []
						this.frontRepo.map_ID_AstructBstruct2Use.clear()

						astructbstruct2uses.forEach(
							astructbstruct2useDB => {
								let astructbstruct2use = new AstructBstruct2Use
								this.frontRepo.array_AstructBstruct2Uses.push(astructbstruct2use)
								this.frontRepo.map_ID_AstructBstruct2Use.set(astructbstruct2useDB.ID, astructbstruct2use)
							}
						)

						// init the arrays
						this.frontRepo.array_AstructBstructUses = []
						this.frontRepo.map_ID_AstructBstructUse.clear()

						astructbstructuses.forEach(
							astructbstructuseDB => {
								let astructbstructuse = new AstructBstructUse
								this.frontRepo.array_AstructBstructUses.push(astructbstructuse)
								this.frontRepo.map_ID_AstructBstructUse.set(astructbstructuseDB.ID, astructbstructuse)
							}
						)

						// init the arrays
						this.frontRepo.array_Bstructs = []
						this.frontRepo.map_ID_Bstruct.clear()

						bstructs.forEach(
							bstructDB => {
								let bstruct = new Bstruct
								this.frontRepo.array_Bstructs.push(bstruct)
								this.frontRepo.map_ID_Bstruct.set(bstructDB.ID, bstruct)
							}
						)

						// init the arrays
						this.frontRepo.array_Dstructs = []
						this.frontRepo.map_ID_Dstruct.clear()

						dstructs.forEach(
							dstructDB => {
								let dstruct = new Dstruct
								this.frontRepo.array_Dstructs.push(dstruct)
								this.frontRepo.map_ID_Dstruct.set(dstructDB.ID, dstruct)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						astructs.forEach(
							astructDB => {
								let astruct = this.frontRepo.map_ID_Astruct.get(astructDB.ID)
								CopyAstructDBToAstruct(astructDB, astruct!, this.frontRepo)
							}
						)

						// fill up front objects
						astructbstruct2uses.forEach(
							astructbstruct2useDB => {
								let astructbstruct2use = this.frontRepo.map_ID_AstructBstruct2Use.get(astructbstruct2useDB.ID)
								CopyAstructBstruct2UseDBToAstructBstruct2Use(astructbstruct2useDB, astructbstruct2use!, this.frontRepo)
							}
						)

						// fill up front objects
						astructbstructuses.forEach(
							astructbstructuseDB => {
								let astructbstructuse = this.frontRepo.map_ID_AstructBstructUse.get(astructbstructuseDB.ID)
								CopyAstructBstructUseDBToAstructBstructUse(astructbstructuseDB, astructbstructuse!, this.frontRepo)
							}
						)

						// fill up front objects
						bstructs.forEach(
							bstructDB => {
								let bstruct = this.frontRepo.map_ID_Bstruct.get(bstructDB.ID)
								CopyBstructDBToBstruct(bstructDB, bstruct!, this.frontRepo)
							}
						)

						// fill up front objects
						dstructs.forEach(
							dstructDB => {
								let dstruct = this.frontRepo.map_ID_Dstruct.get(dstructDB.ID)
								CopyDstructDBToDstruct(dstructDB, dstruct!, this.frontRepo)
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
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/test/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)


		return new Observable(observer => {
			this.socket!.onmessage = event => {
				let _this = this

				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let json = event.data
				console.log(event.data)

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				this.frontRepo.array_Astructs = []
				this.frontRepo.map_ID_Astruct.clear()

				backRepoData.AstructDBs?.forEach(
					astructDB => {
						let astruct = new Astruct
						this.frontRepo.array_Astructs.push(astruct)
						this.frontRepo.map_ID_Astruct.set(astructDB.ID, astruct)
					}
				)

				// init the arrays
				this.frontRepo.array_AstructBstruct2Uses = []
				this.frontRepo.map_ID_AstructBstruct2Use.clear()

				backRepoData.AstructBstructUseDBs?.forEach(
					astructbstruct2useDB => {
						let astructbstruct2use = new AstructBstruct2Use
						this.frontRepo.array_AstructBstruct2Uses.push(astructbstruct2use)
						this.frontRepo.map_ID_AstructBstruct2Use.set(astructbstruct2useDB.ID, astructbstruct2use)
					}
				)

				// init the arrays
				this.frontRepo.array_AstructBstructUses = []
				this.frontRepo.map_ID_AstructBstructUse.clear()

				backRepoData.AstructBstructUseDBs?.forEach(
					astructbstructuseDB => {
						let astructbstructuse = new AstructBstructUse
						this.frontRepo.array_AstructBstructUses.push(astructbstructuse)
						this.frontRepo.map_ID_AstructBstructUse.set(astructbstructuseDB.ID, astructbstructuse)
					}
				)

				// init the arrays
				this.frontRepo.array_Bstructs = []
				this.frontRepo.map_ID_Bstruct.clear()

				backRepoData.BstructDBs?.forEach(
					bstructDB => {
						let bstruct = new Bstruct
						this.frontRepo.array_Bstructs.push(bstruct)
						this.frontRepo.map_ID_Bstruct.set(bstructDB.ID, bstruct)
					}
				)

				// init the arrays
				this.frontRepo.array_Dstructs = []
				this.frontRepo.map_ID_Dstruct.clear()

				backRepoData.DstructDBs?.forEach(
					dstructDB => {
						let dstruct = new Dstruct
						this.frontRepo.array_Dstructs.push(dstruct)
						this.frontRepo.map_ID_Dstruct.set(dstructDB.ID, dstruct)
					}
				)

				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AstructDBs?.forEach(
					astructDB => {
						let astruct = this.frontRepo.map_ID_Astruct.get(astructDB.ID)
						CopyAstructDBToAstruct(astructDB, astruct!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AstructBstruct2UseDBs?.forEach(
					astructbstruct2useDB => {
						let astructbstruct2use = this.frontRepo.map_ID_AstructBstruct2Use.get(astructbstruct2useDB.ID)
						CopyAstructBstruct2UseDBToAstructBstruct2Use(astructbstruct2useDB, astructbstruct2use!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.AstructBstructUseDBs?.forEach(
					astructbstructuseDB => {
						let astructbstructuse = this.frontRepo.map_ID_AstructBstructUse.get(astructbstructuseDB.ID)
						CopyAstructBstructUseDBToAstructBstructUse(astructbstructuseDB, astructbstructuse!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BstructDBs?.forEach(
					bstructDB => {
						let bstruct = this.frontRepo.map_ID_Bstruct.get(bstructDB.ID)
						CopyBstructDBToBstruct(bstructDB, bstruct!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DstructDBs?.forEach(
					dstructDB => {
						let dstruct = this.frontRepo.map_ID_Dstruct.get(dstructDB.ID)
						CopyDstructDBToDstruct(dstructDB, dstruct!, this.frontRepo)
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

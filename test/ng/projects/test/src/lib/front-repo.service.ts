import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

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

export const StackType = "github.com/fullstack-lang/gong/test/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Astructs = new Array<Astruct>() // array of front instances
	map_ID_Astruct = new Map<number, Astruct>() // map of front instances

	AstructBstruct2Uses_array = new Array<AstructBstruct2UseDB>() // array of repo instances
	AstructBstruct2UseDBs = new Map<number, AstructBstruct2UseDB>() // map of repo instances
	AstructBstruct2Uses_batch = new Map<number, AstructBstruct2UseDB>() // same but only in last GET (for finding repo instances to delete)

	array_AstructBstruct2Uses = new Array<AstructBstruct2Use>() // array of front instances
	map_ID_AstructBstruct2Use = new Map<number, AstructBstruct2Use>() // map of front instances

	AstructBstructUses_array = new Array<AstructBstructUseDB>() // array of repo instances
	AstructBstructUseDBs = new Map<number, AstructBstructUseDB>() // map of repo instances
	AstructBstructUses_batch = new Map<number, AstructBstructUseDB>() // same but only in last GET (for finding repo instances to delete)

	array_AstructBstructUses = new Array<AstructBstructUse>() // array of front instances
	map_ID_AstructBstructUse = new Map<number, AstructBstructUse>() // map of front instances

	Bstructs_array = new Array<BstructDB>() // array of repo instances
	BstructDBs = new Map<number, BstructDB>() // map of repo instances
	Bstructs_batch = new Map<number, BstructDB>() // same but only in last GET (for finding repo instances to delete)

	array_Bstructs = new Array<Bstruct>() // array of front instances
	map_ID_Bstruct = new Map<number, Bstruct>() // map of front instances

	Dstructs_array = new Array<DstructDB>() // array of repo instances
	DstructDBs = new Map<number, DstructDB>() // map of repo instances
	Dstructs_batch = new Map<number, DstructDB>() // same but only in last GET (for finding repo instances to delete)

	array_Dstructs = new Array<Dstruct>() // array of front instances
	map_ID_Dstruct = new Map<number, Dstruct>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getFrontArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
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
						this.frontRepo.AstructBstruct2Uses_array = astructbstruct2uses
						this.frontRepo.array_AstructBstruct2Uses = []
						this.frontRepo.map_ID_AstructBstruct2Use.clear()

						// clear the map that counts AstructBstruct2Use in the GET
						this.frontRepo.AstructBstruct2Uses_batch.clear()

						astructbstruct2uses.forEach(
							astructbstruct2useDB => {
								this.frontRepo.AstructBstruct2UseDBs.set(astructbstruct2useDB.ID, astructbstruct2useDB)
								this.frontRepo.AstructBstruct2Uses_batch.set(astructbstruct2useDB.ID, astructbstruct2useDB)
								let astructbstruct2use = new AstructBstruct2Use
								this.frontRepo.array_AstructBstruct2Uses.push(astructbstruct2use)
								this.frontRepo.map_ID_AstructBstruct2Use.set(astructbstruct2useDB.ID, astructbstruct2use)
							}
						)

						// clear astructbstruct2uses that are absent from the batch
						this.frontRepo.AstructBstruct2UseDBs.forEach(
							astructbstruct2useDB => {
								if (this.frontRepo.AstructBstruct2Uses_batch.get(astructbstruct2useDB.ID) == undefined) {
									this.frontRepo.AstructBstruct2UseDBs.delete(astructbstruct2useDB.ID)
								}
							}
						)

						// sort AstructBstruct2Uses_array array
						this.frontRepo.AstructBstruct2Uses_array.sort((t1, t2) => {
							if (t1.Name > t2.Name) {
								return 1;
							}
							if (t1.Name < t2.Name) {
								return -1;
							}
							return 0;
						});

						// init the arrays
						this.frontRepo.AstructBstructUses_array = astructbstructuses
						this.frontRepo.array_AstructBstructUses = []
						this.frontRepo.map_ID_AstructBstructUse.clear()

						// clear the map that counts AstructBstructUse in the GET
						this.frontRepo.AstructBstructUses_batch.clear()

						astructbstructuses.forEach(
							astructbstructuseDB => {
								this.frontRepo.AstructBstructUseDBs.set(astructbstructuseDB.ID, astructbstructuseDB)
								this.frontRepo.AstructBstructUses_batch.set(astructbstructuseDB.ID, astructbstructuseDB)
								let astructbstructuse = new AstructBstructUse
								this.frontRepo.array_AstructBstructUses.push(astructbstructuse)
								this.frontRepo.map_ID_AstructBstructUse.set(astructbstructuseDB.ID, astructbstructuse)
							}
						)

						// clear astructbstructuses that are absent from the batch
						this.frontRepo.AstructBstructUseDBs.forEach(
							astructbstructuseDB => {
								if (this.frontRepo.AstructBstructUses_batch.get(astructbstructuseDB.ID) == undefined) {
									this.frontRepo.AstructBstructUseDBs.delete(astructbstructuseDB.ID)
								}
							}
						)

						// sort AstructBstructUses_array array
						this.frontRepo.AstructBstructUses_array.sort((t1, t2) => {
							if (t1.Name > t2.Name) {
								return 1;
							}
							if (t1.Name < t2.Name) {
								return -1;
							}
							return 0;
						});

						// init the arrays
						this.frontRepo.Bstructs_array = bstructs
						this.frontRepo.array_Bstructs = []
						this.frontRepo.map_ID_Bstruct.clear()

						// clear the map that counts Bstruct in the GET
						this.frontRepo.Bstructs_batch.clear()

						bstructs.forEach(
							bstructDB => {
								this.frontRepo.BstructDBs.set(bstructDB.ID, bstructDB)
								this.frontRepo.Bstructs_batch.set(bstructDB.ID, bstructDB)
								let bstruct = new Bstruct
								this.frontRepo.array_Bstructs.push(bstruct)
								this.frontRepo.map_ID_Bstruct.set(bstructDB.ID, bstruct)
							}
						)

						// clear bstructs that are absent from the batch
						this.frontRepo.BstructDBs.forEach(
							bstructDB => {
								if (this.frontRepo.Bstructs_batch.get(bstructDB.ID) == undefined) {
									this.frontRepo.BstructDBs.delete(bstructDB.ID)
								}
							}
						)

						// sort Bstructs_array array
						this.frontRepo.Bstructs_array.sort((t1, t2) => {
							if (t1.Name > t2.Name) {
								return 1;
							}
							if (t1.Name < t2.Name) {
								return -1;
							}
							return 0;
						});

						// init the arrays
						this.frontRepo.Dstructs_array = dstructs
						this.frontRepo.array_Dstructs = []
						this.frontRepo.map_ID_Dstruct.clear()

						// clear the map that counts Dstruct in the GET
						this.frontRepo.Dstructs_batch.clear()

						dstructs.forEach(
							dstructDB => {
								this.frontRepo.DstructDBs.set(dstructDB.ID, dstructDB)
								this.frontRepo.Dstructs_batch.set(dstructDB.ID, dstructDB)
								let dstruct = new Dstruct
								this.frontRepo.array_Dstructs.push(dstruct)
								this.frontRepo.map_ID_Dstruct.set(dstructDB.ID, dstruct)
							}
						)

						// clear dstructs that are absent from the batch
						this.frontRepo.DstructDBs.forEach(
							dstructDB => {
								if (this.frontRepo.Dstructs_batch.get(dstructDB.ID) == undefined) {
									this.frontRepo.DstructDBs.delete(dstructDB.ID)
								}
							}
						)

						// sort Dstructs_array array
						this.frontRepo.Dstructs_array.sort((t1, t2) => {
							if (t1.Name > t2.Name) {
								return 1;
							}
							if (t1.Name < t2.Name) {
								return -1;
							}
							return 0;
						});


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// init front objects
						astructs.forEach(
							astructDB => {
								let astruct = this.frontRepo.map_ID_Astruct.get(astructDB.ID)
								CopyAstructDBToAstruct(astructDB, astruct!, this.frontRepo)
							}
						)

						// init front objects
						this.frontRepo.AstructBstruct2Uses_array.forEach(
							astructbstruct2useDB => {
								let astructbstruct2use = this.frontRepo.map_ID_AstructBstruct2Use.get(astructbstruct2useDB.ID)
								CopyAstructBstruct2UseDBToAstructBstruct2Use(astructbstruct2useDB, astructbstruct2use!, this.frontRepo)
							}
						)

						// init front objects
						this.frontRepo.AstructBstructUses_array.forEach(
							astructbstructuseDB => {
								let astructbstructuse = this.frontRepo.map_ID_AstructBstructUse.get(astructbstructuseDB.ID)
								CopyAstructBstructUseDBToAstructBstructUse(astructbstructuseDB, astructbstructuse!, this.frontRepo)
							}
						)

						// init front objects
						this.frontRepo.Bstructs_array.forEach(
							bstructDB => {
								let bstruct = this.frontRepo.map_ID_Bstruct.get(bstructDB.ID)
								CopyBstructDBToBstruct(bstructDB, bstruct!, this.frontRepo)
							}
						)

						// init front objects
						this.frontRepo.Dstructs_array.forEach(
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

	// insertion point for pull per struct 

	// AstructBstruct2UsePull performs a GET on AstructBstruct2Use of the stack and redeem association pointers 
	AstructBstruct2UsePull(): Observable<FrontRepo> {
		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest([
					this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath, this.frontRepo)
				]).subscribe(
					([ // insertion point sub template 
						astructbstruct2uses,
					]) => {
						// init the array
						this.frontRepo.AstructBstruct2Uses_array = astructbstruct2uses

						// clear the map that counts AstructBstruct2Use in the GET
						this.frontRepo.AstructBstruct2Uses_batch.clear()

						// 
						// First Step: init map of instances
						// insertion point sub template 
						astructbstruct2uses.forEach(
							astructbstruct2use => {
								this.frontRepo.AstructBstruct2UseDBs.set(astructbstruct2use.ID, astructbstruct2use)
								this.frontRepo.AstructBstruct2Uses_batch.set(astructbstruct2use.ID, astructbstruct2use)
							}
						)

						// clear astructbstruct2uses that are absent from the GET
						this.frontRepo.AstructBstruct2UseDBs.forEach(
							astructbstruct2use => {
								if (this.frontRepo.AstructBstruct2Uses_batch.get(astructbstruct2use.ID) == undefined) {
									this.frontRepo.AstructBstruct2UseDBs.delete(astructbstruct2use.ID)
								}
							}
						)

						// 
						// Second Step: redeem pointers between instances (thanks to maps in the First Step)
						// insertion point sub template 

						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	// AstructBstructUsePull performs a GET on AstructBstructUse of the stack and redeem association pointers 
	AstructBstructUsePull(): Observable<FrontRepo> {
		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest([
					this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath, this.frontRepo)
				]).subscribe(
					([ // insertion point sub template 
						astructbstructuses,
					]) => {
						// init the array
						this.frontRepo.AstructBstructUses_array = astructbstructuses

						// clear the map that counts AstructBstructUse in the GET
						this.frontRepo.AstructBstructUses_batch.clear()

						// 
						// First Step: init map of instances
						// insertion point sub template 
						astructbstructuses.forEach(
							astructbstructuse => {
								this.frontRepo.AstructBstructUseDBs.set(astructbstructuse.ID, astructbstructuse)
								this.frontRepo.AstructBstructUses_batch.set(astructbstructuse.ID, astructbstructuse)
							}
						)

						// clear astructbstructuses that are absent from the GET
						this.frontRepo.AstructBstructUseDBs.forEach(
							astructbstructuse => {
								if (this.frontRepo.AstructBstructUses_batch.get(astructbstructuse.ID) == undefined) {
									this.frontRepo.AstructBstructUseDBs.delete(astructbstructuse.ID)
								}
							}
						)

						// 
						// Second Step: redeem pointers between instances (thanks to maps in the First Step)
						// insertion point sub template 

						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	// BstructPull performs a GET on Bstruct of the stack and redeem association pointers 
	BstructPull(): Observable<FrontRepo> {
		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest([
					this.bstructService.getBstructs(this.GONG__StackPath, this.frontRepo)
				]).subscribe(
					([ // insertion point sub template 
						bstructs,
					]) => {
						// init the array
						this.frontRepo.Bstructs_array = bstructs

						// clear the map that counts Bstruct in the GET
						this.frontRepo.Bstructs_batch.clear()

						// 
						// First Step: init map of instances
						// insertion point sub template 
						bstructs.forEach(
							bstruct => {
								this.frontRepo.BstructDBs.set(bstruct.ID, bstruct)
								this.frontRepo.Bstructs_batch.set(bstruct.ID, bstruct)
							}
						)

						// clear bstructs that are absent from the GET
						this.frontRepo.BstructDBs.forEach(
							bstruct => {
								if (this.frontRepo.Bstructs_batch.get(bstruct.ID) == undefined) {
									this.frontRepo.BstructDBs.delete(bstruct.ID)
								}
							}
						)

						// 
						// Second Step: redeem pointers between instances (thanks to maps in the First Step)
						// insertion point sub template 

						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	// DstructPull performs a GET on Dstruct of the stack and redeem association pointers 
	DstructPull(): Observable<FrontRepo> {
		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest([
					this.dstructService.getDstructs(this.GONG__StackPath, this.frontRepo)
				]).subscribe(
					([ // insertion point sub template 
						dstructs,
					]) => {
						// init the array
						this.frontRepo.Dstructs_array = dstructs

						// clear the map that counts Dstruct in the GET
						this.frontRepo.Dstructs_batch.clear()

						// 
						// First Step: init map of instances
						// insertion point sub template 
						dstructs.forEach(
							dstruct => {
								this.frontRepo.DstructDBs.set(dstruct.ID, dstruct)
								this.frontRepo.Dstructs_batch.set(dstruct.ID, dstruct)
							}
						)

						// clear dstructs that are absent from the GET
						this.frontRepo.DstructDBs.forEach(
							dstruct => {
								if (this.frontRepo.Dstructs_batch.get(dstruct.ID) == undefined) {
									this.frontRepo.DstructDBs.delete(dstruct.ID)
								}
							}
						)

						// 
						// Second Step: redeem pointers between instances (thanks to maps in the First Step)
						// insertion point sub template 

						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
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

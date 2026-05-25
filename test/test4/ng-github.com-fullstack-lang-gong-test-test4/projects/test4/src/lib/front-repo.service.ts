import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AstructAPI } from './astruct-api'
import { Astruct, CopyAstructAPIToAstruct } from './astruct'
import { AstructService } from './astruct.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/test/test4/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Astructs = new Array<Astruct>() // array of front instances
	map_ID_Astruct = new Map<number, Astruct>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Astruct':
				return this.array_Astructs as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Astruct':
				return this.map_ID_Astruct as unknown as Map<number, Type>
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
		private astructService: AstructService,
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
		Observable<AstructAPI[]>,
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
			this.astructService.getAstructs(this.Name, this.frontRepo),
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
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var astructs: AstructAPI[]
						astructs = astructs_ as AstructAPI[]

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


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		console.log("github.com/fullstack-lang/gong/test/test4/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			console.log("github.com/fullstack-lang/gong/test/test4/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/test/test4/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			console.log("github.com/fullstack-lang/gong/test/test4/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:'

			const processData = (dataString: string) => {
				console.log("github.com/fullstack-lang/gong/test/test4/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
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


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AstructAPIs.forEach(
					astructAPI => {
						let astruct = frontRepo.map_ID_Astruct.get(astructAPI.ID)
						CopyAstructAPIToAstruct(astructAPI, astruct!, frontRepo)
					}
				)


				observer.next(frontRepo)
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				console.log("github.com/fullstack-lang/gong/test/test4/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					console.log("github.com/fullstack-lang/gong/test/test4/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket("github.com/fullstack-lang/gong/test/test4/go", Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					console.log("github.com/fullstack-lang/gong/test/test4/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					console.log("github.com/fullstack-lang/gong/test/test4/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						console.log("github.com/fullstack-lang/gong/test/test4/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						console.log("github.com/fullstack-lang/gong/test/test4/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/test/test4/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						console.log("github.com/fullstack-lang/gong/test/test4/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/test/test4/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
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
export function getAstructUniqueID(id: number): number {
	return 31 * id
}

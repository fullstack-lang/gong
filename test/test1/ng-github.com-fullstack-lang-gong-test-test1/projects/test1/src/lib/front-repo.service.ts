// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { AstructAPI } from './astruct-api'
import { Astruct, CopyAstructAPIToAstruct } from './astruct'

import { AstructBstruct2UseAPI } from './astructbstruct2use-api'
import { AstructBstruct2Use, CopyAstructBstruct2UseAPIToAstructBstruct2Use } from './astructbstruct2use'

import { AstructBstructUseAPI } from './astructbstructuse-api'
import { AstructBstructUse, CopyAstructBstructUseAPIToAstructBstructUse } from './astructbstructuse'

import { BstructAPI } from './bstruct-api'
import { Bstruct, CopyBstructAPIToBstruct } from './bstruct'

import { DstructAPI } from './dstruct-api'
import { Dstruct, CopyDstructAPIToDstruct } from './dstruct'

import { GstructAPI } from './gstruct-api'
import { Gstruct, CopyGstructAPIToGstruct } from './gstruct'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/test/test1/go/models"

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
				throw new Error("Type not recognized")
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
		private http: HttpClient,
		private ngZone: NgZone,
	) { }

	//
	// pull returns the FrontRepo Observable
	//
	pull(Name: string = ""): Observable<FrontRepo> {
		this.Name = Name
		return of(this.frontRepo)
	}

	public connectToWebSocket(Name: string): Observable<FrontRepo> {

		// console.log("github.com/fullstack-lang/gong/test/test1/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/test/test1/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/test/test1/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/test/test1/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/test/test1/go; connectToWebSocket: processData called")
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


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// Offline mode handling: Listen to the global event
			if (isOfflineMode) {
				console.log("github.com/fullstack-lang/gong/test/test1/go; Offline mode detected. Skipping WebSocket connection.")

				window.addEventListener('message', (event) => {
					if (event.data && event.data.type === 'STAGE_UPDATE') {
						console.log("github.com/fullstack-lang/gong/test/test1/go; Received STAGE_UPDATE message.")
						processData(JSON.stringify(event.data.data))
					}
				})

				return () => {
					console.log("github.com/fullstack-lang/gong/test/test1/go; Cleaning up offline message listener.")
				}
			}

			// Fallback: If not offline, create normal WebSocket
			const attemptConnection = () => {
				// Offline check inside attemptConnection: if window.openWasmSocket is available, use it!
				if (typeof window !== 'undefined' && (window as any).openWasmSocket) {
					(window as any).openWasmSocket('github.com/fullstack-lang/gong/test/test1/go', Name, (data: any) => {
						processData(data)
					})
					return
				}

				if (isOfflineMode && retryCount > 0) {
					console.log("github.com/fullstack-lang/gong/test/test1/go; Waiting for wasm socket provider...")
					setTimeout(() => attemptConnection(), 100)
					return
				}

				if (isOfflineMode) {
					console.error("github.com/fullstack-lang/gong/test/test1/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.")
					return
				}

				socket = new WebSocket(url)

				socket.onopen = () => {
					// console.log("github.com/fullstack-lang/gong/test/test1/go; WebSocket connection opened successfully:", url)
				}

				socket.onmessage = (event) => {
					// console.log("github.com/fullstack-lang/gong/test/test1/go; WebSocket message received:", event.data)
					processData(event.data)
				}

				socket.onerror = (error) => {
					console.error("github.com/fullstack-lang/gong/test/test1/go WebSocket: onerror", error)
					observer.error(error)
				}

				socket.onclose = (event) => {
					// console.log("github.com/fullstack-lang/gong/test/test1/go; WebSocket connection closed:", event)
					observer.complete()
				}
			}

			let retryCount = 10
			attemptConnection()

			return () => {
				// console.log("github.com/fullstack-lang/gong/test/test1/go; Cleaning up WebSocket connection")
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

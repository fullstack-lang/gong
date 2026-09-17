// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { FreqencyAPI } from './freqency-api'
import { Freqency, CopyFreqencyAPIToFreqency } from './freqency'

import { NoteAPI } from './note-api'
import { Note, CopyNoteAPIToNote } from './note'

import { PlayerAPI } from './player-api'
import { Player, CopyPlayerAPIToPlayer } from './player'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/tone/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Freqencys = new Array<Freqency>() // array of front instances
	map_ID_Freqency = new Map<number, Freqency>() // map of front instances

	array_Notes = new Array<Note>() // array of front instances
	map_ID_Note = new Map<number, Note>() // map of front instances

	array_Players = new Array<Player>() // array of front instances
	map_ID_Player = new Map<number, Player>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Freqency':
				return this.array_Freqencys as unknown as Array<Type>
			case 'Note':
				return this.array_Notes as unknown as Array<Type>
			case 'Player':
				return this.array_Players as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Freqency':
				return this.map_ID_Freqency as unknown as Map<number, Type>
			case 'Note':
				return this.map_ID_Note as unknown as Map<number, Type>
			case 'Player':
				return this.map_ID_Player as unknown as Map<number, Type>
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

		// console.log("github.com/fullstack-lang/gong/lib/tone/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/tone/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/tone/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/tone/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/tone/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Freqencys = []
				frontRepo.map_ID_Freqency.clear()

				backRepoData.FreqencyAPIs.forEach(
					freqencyAPI => {
						let freqency = new Freqency
						frontRepo.array_Freqencys.push(freqency)
						frontRepo.map_ID_Freqency.set(freqencyAPI.ID, freqency)
					}
				)

				// init the arrays
				frontRepo.array_Notes = []
				frontRepo.map_ID_Note.clear()

				backRepoData.NoteAPIs.forEach(
					noteAPI => {
						let note = new Note
						frontRepo.array_Notes.push(note)
						frontRepo.map_ID_Note.set(noteAPI.ID, note)
					}
				)

				// init the arrays
				frontRepo.array_Players = []
				frontRepo.map_ID_Player.clear()

				backRepoData.PlayerAPIs.forEach(
					playerAPI => {
						let player = new Player
						frontRepo.array_Players.push(player)
						frontRepo.map_ID_Player.set(playerAPI.ID, player)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.FreqencyAPIs.forEach(
					freqencyAPI => {
						let freqency = frontRepo.map_ID_Freqency.get(freqencyAPI.ID)
						CopyFreqencyAPIToFreqency(freqencyAPI, freqency!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.NoteAPIs.forEach(
					noteAPI => {
						let note = frontRepo.map_ID_Note.get(noteAPI.ID)
						CopyNoteAPIToNote(noteAPI, note!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PlayerAPIs.forEach(
					playerAPI => {
						let player = frontRepo.map_ID_Player.get(playerAPI.ID)
						CopyPlayerAPIToPlayer(playerAPI, player!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// 3. Connection Loop
			const attemptConnection = (retries: number): void => {
				// console.log("github.com/fullstack-lang/gong/lib/tone/go; attemptConnection: retries =", retries, "isOfflineMode =", isOfflineMode)

				// A. WASM OFFLINE MODE (Check if Go is ready)
				if ((window as any).openWasmSocket) {
					// console.log("github.com/fullstack-lang/gong/lib/tone/go; attemptConnection: openWasmSocket exists, calling it");
					(window as any).openWasmSocket("github.com/fullstack-lang/gong/lib/tone/go", Name, processData);
					return;
				}

				// B. WAITING FOR WASM
				if (isOfflineMode && retries > 0) {
					// console.log("github.com/fullstack-lang/gong/lib/tone/go; attemptConnection: WAITING FOR WASM. Retries left:", retries)
					setTimeout(() => attemptConnection(retries - 1), 100);
					return;
				}

				// C. STANDARD SERVER MODE
				if (!isOfflineMode) {
					// console.log("github.com/fullstack-lang/gong/lib/tone/go; attemptConnection: STANDARD SERVER MODE. url =", url)
					socket = new WebSocket(url)
					socket.onopen = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/tone/go; WebSocket: onopen", event)
					}
					socket.onmessage = event => {
						// console.log("github.com/fullstack-lang/gong/lib/tone/go; WebSocket: onmessage")
						processData(event.data)
					}
					socket.onerror = event => {
						console.error("github.com/fullstack-lang/gong/lib/tone/go WebSocket: onerror", event)
						observer.error(event)
					}
					socket.onclose = (event) => {
						// console.log("github.com/fullstack-lang/gong/lib/tone/go; WebSocket: onclose", event)
						observer.complete()
					}
				} else {
					console.error("github.com/fullstack-lang/gong/lib/tone/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
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
export function getFreqencyUniqueID(id: number): number {
	return 31 * id
}
export function getNoteUniqueID(id: number): number {
	return 37 * id
}
export function getPlayerUniqueID(id: number): number {
	return 41 * id
}

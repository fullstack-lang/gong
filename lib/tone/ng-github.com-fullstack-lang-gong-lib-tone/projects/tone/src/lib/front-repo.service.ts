import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { FreqencyAPI } from './freqency-api'
import { Freqency, CopyFreqencyAPIToFreqency } from './freqency'
import { FreqencyService } from './freqency.service'

import { NoteAPI } from './note-api'
import { Note, CopyNoteAPIToNote } from './note'
import { NoteService } from './note.service'

import { PlayerAPI } from './player-api'
import { Player, CopyPlayerAPIToPlayer } from './player'
import { PlayerService } from './player.service'


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
				throw new Error("Type not recognized");
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
		private freqencyService: FreqencyService,
		private noteService: NoteService,
		private playerService: PlayerService,
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
		Observable<FreqencyAPI[]>,
		Observable<NoteAPI[]>,
		Observable<PlayerAPI[]>,
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
			this.freqencyService.getFreqencys(this.GONG__StackPath, this.frontRepo),
			this.noteService.getNotes(this.GONG__StackPath, this.frontRepo),
			this.playerService.getPlayers(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						freqencys_,
						notes_,
						players_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var freqencys: FreqencyAPI[]
						freqencys = freqencys_ as FreqencyAPI[]
						var notes: NoteAPI[]
						notes = notes_ as NoteAPI[]
						var players: PlayerAPI[]
						players = players_ as PlayerAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Freqencys = []
						this.frontRepo.map_ID_Freqency.clear()

						freqencys.forEach(
							freqencyAPI => {
								let freqency = new Freqency
								this.frontRepo.array_Freqencys.push(freqency)
								this.frontRepo.map_ID_Freqency.set(freqencyAPI.ID, freqency)
							}
						)

						// init the arrays
						this.frontRepo.array_Notes = []
						this.frontRepo.map_ID_Note.clear()

						notes.forEach(
							noteAPI => {
								let note = new Note
								this.frontRepo.array_Notes.push(note)
								this.frontRepo.map_ID_Note.set(noteAPI.ID, note)
							}
						)

						// init the arrays
						this.frontRepo.array_Players = []
						this.frontRepo.map_ID_Player.clear()

						players.forEach(
							playerAPI => {
								let player = new Player
								this.frontRepo.array_Players.push(player)
								this.frontRepo.map_ID_Player.set(playerAPI.ID, player)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						freqencys.forEach(
							freqencyAPI => {
								let freqency = this.frontRepo.map_ID_Freqency.get(freqencyAPI.ID)
								CopyFreqencyAPIToFreqency(freqencyAPI, freqency!, this.frontRepo)
							}
						)

						// fill up front objects
						notes.forEach(
							noteAPI => {
								let note = this.frontRepo.map_ID_Note.get(noteAPI.ID)
								CopyNoteAPIToNote(noteAPI, note!, this.frontRepo)
							}
						)

						// fill up front objects
						players.forEach(
							playerAPI => {
								let player = this.frontRepo.map_ID_Player.get(playerAPI.ID)
								CopyPlayerAPIToPlayer(playerAPI, player!, this.frontRepo)
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
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/lib/tone/go/v1/ws/stage'
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
export function getFreqencyUniqueID(id: number): number {
	return 31 * id
}
export function getNoteUniqueID(id: number): number {
	return 37 * id
}
export function getPlayerUniqueID(id: number): number {
	return 41 * id
}

// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { CommandAPI } from './command-api'
import { Command, CopyCommandAPIToCommand } from './command'

import { DummyAgentAPI } from './dummyagent-api'
import { DummyAgent, CopyDummyAgentAPIToDummyAgent } from './dummyagent'

import { EngineAPI } from './engine-api'
import { Engine, CopyEngineAPIToEngine } from './engine'

import { EventAPI } from './event-api'
import { Event, CopyEventAPIToEvent } from './event'

import { StatusAPI } from './status-api'
import { Status, CopyStatusAPIToStatus } from './status'

import { UpdateStateAPI } from './updatestate-api'
import { UpdateState, CopyUpdateStateAPIToUpdateState } from './updatestate'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/sim/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Commands = new Array<Command>() // array of front instances
	map_ID_Command = new Map<number, Command>() // map of front instances

	array_DummyAgents = new Array<DummyAgent>() // array of front instances
	map_ID_DummyAgent = new Map<number, DummyAgent>() // map of front instances

	array_Engines = new Array<Engine>() // array of front instances
	map_ID_Engine = new Map<number, Engine>() // map of front instances

	array_Events = new Array<Event>() // array of front instances
	map_ID_Event = new Map<number, Event>() // map of front instances

	array_Statuss = new Array<Status>() // array of front instances
	map_ID_Status = new Map<number, Status>() // map of front instances

	array_UpdateStates = new Array<UpdateState>() // array of front instances
	map_ID_UpdateState = new Map<number, UpdateState>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Command':
				return this.array_Commands as unknown as Array<Type>
			case 'DummyAgent':
				return this.array_DummyAgents as unknown as Array<Type>
			case 'Engine':
				return this.array_Engines as unknown as Array<Type>
			case 'Event':
				return this.array_Events as unknown as Array<Type>
			case 'Status':
				return this.array_Statuss as unknown as Array<Type>
			case 'UpdateState':
				return this.array_UpdateStates as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Command':
				return this.map_ID_Command as unknown as Map<number, Type>
			case 'DummyAgent':
				return this.map_ID_DummyAgent as unknown as Map<number, Type>
			case 'Engine':
				return this.map_ID_Engine as unknown as Map<number, Type>
			case 'Event':
				return this.map_ID_Event as unknown as Map<number, Type>
			case 'Status':
				return this.map_ID_Status as unknown as Map<number, Type>
			case 'UpdateState':
				return this.map_ID_UpdateState as unknown as Map<number, Type>
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

		// console.log("github.com/fullstack-lang/gong/lib/sim/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/sim/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/sim/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/sim/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/sim/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Commands = []
				frontRepo.map_ID_Command.clear()

				backRepoData.CommandAPIs.forEach(
					commandAPI => {
						let command = new Command
						frontRepo.array_Commands.push(command)
						frontRepo.map_ID_Command.set(commandAPI.ID, command)
					}
				)

				// init the arrays
				frontRepo.array_DummyAgents = []
				frontRepo.map_ID_DummyAgent.clear()

				backRepoData.DummyAgentAPIs.forEach(
					dummyagentAPI => {
						let dummyagent = new DummyAgent
						frontRepo.array_DummyAgents.push(dummyagent)
						frontRepo.map_ID_DummyAgent.set(dummyagentAPI.ID, dummyagent)
					}
				)

				// init the arrays
				frontRepo.array_Engines = []
				frontRepo.map_ID_Engine.clear()

				backRepoData.EngineAPIs.forEach(
					engineAPI => {
						let engine = new Engine
						frontRepo.array_Engines.push(engine)
						frontRepo.map_ID_Engine.set(engineAPI.ID, engine)
					}
				)

				// init the arrays
				frontRepo.array_Events = []
				frontRepo.map_ID_Event.clear()

				backRepoData.EventAPIs.forEach(
					eventAPI => {
						let event = new Event
						frontRepo.array_Events.push(event)
						frontRepo.map_ID_Event.set(eventAPI.ID, event)
					}
				)

				// init the arrays
				frontRepo.array_Statuss = []
				frontRepo.map_ID_Status.clear()

				backRepoData.StatusAPIs.forEach(
					statusAPI => {
						let status = new Status
						frontRepo.array_Statuss.push(status)
						frontRepo.map_ID_Status.set(statusAPI.ID, status)
					}
				)

				// init the arrays
				frontRepo.array_UpdateStates = []
				frontRepo.map_ID_UpdateState.clear()

				backRepoData.UpdateStateAPIs.forEach(
					updatestateAPI => {
						let updatestate = new UpdateState
						frontRepo.array_UpdateStates.push(updatestate)
						frontRepo.map_ID_UpdateState.set(updatestateAPI.ID, updatestate)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.CommandAPIs.forEach(
					commandAPI => {
						let command = frontRepo.map_ID_Command.get(commandAPI.ID)
						CopyCommandAPIToCommand(commandAPI, command!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DummyAgentAPIs.forEach(
					dummyagentAPI => {
						let dummyagent = frontRepo.map_ID_DummyAgent.get(dummyagentAPI.ID)
						CopyDummyAgentAPIToDummyAgent(dummyagentAPI, dummyagent!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.EngineAPIs.forEach(
					engineAPI => {
						let engine = frontRepo.map_ID_Engine.get(engineAPI.ID)
						CopyEngineAPIToEngine(engineAPI, engine!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.EventAPIs.forEach(
					eventAPI => {
						let event = frontRepo.map_ID_Event.get(eventAPI.ID)
						CopyEventAPIToEvent(eventAPI, event!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.StatusAPIs.forEach(
					statusAPI => {
						let status = frontRepo.map_ID_Status.get(statusAPI.ID)
						CopyStatusAPIToStatus(statusAPI, status!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.UpdateStateAPIs.forEach(
					updatestateAPI => {
						let updatestate = frontRepo.map_ID_UpdateState.get(updatestateAPI.ID)
						CopyUpdateStateAPIToUpdateState(updatestateAPI, updatestate!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// Offline mode handling: Listen to the global event
			if (isOfflineMode) {
				console.log("github.com/fullstack-lang/gong/lib/sim/go; Offline mode detected. Skipping WebSocket connection.")

				window.addEventListener('message', (event) => {
					if (event.data && event.data.type === 'STAGE_UPDATE') {
						console.log("github.com/fullstack-lang/gong/lib/sim/go; Received STAGE_UPDATE message.")
						processData(JSON.stringify(event.data.data))
					}
				})

				return () => {
					console.log("github.com/fullstack-lang/gong/lib/sim/go; Cleaning up offline message listener.")
				}
			}

			// Fallback: If not offline, create normal WebSocket
			const attemptConnection = () => {
				// Offline check inside attemptConnection: if window.openWasmSocket is available, use it!
				if (typeof window !== 'undefined' && (window as any).openWasmSocket) {
					(window as any).openWasmSocket('github.com/fullstack-lang/gong/lib/sim/go', Name, (data: any) => {
						processData(data)
					})
					return
				}

				if (isOfflineMode && retryCount > 0) {
					console.log("github.com/fullstack-lang/gong/lib/sim/go; Waiting for wasm socket provider...")
					setTimeout(() => attemptConnection(), 100)
					return
				}

				if (isOfflineMode) {
					console.error("github.com/fullstack-lang/gong/lib/sim/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.")
					return
				}

				socket = new WebSocket(url)

				socket.onopen = () => {
					// console.log("github.com/fullstack-lang/gong/lib/sim/go; WebSocket connection opened successfully:", url)
				}

				socket.onmessage = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/sim/go; WebSocket message received:", event.data)
					processData(event.data)
				}

				socket.onerror = (error) => {
					console.error("github.com/fullstack-lang/gong/lib/sim/go WebSocket: onerror", error)
					observer.error(error)
				}

				socket.onclose = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/sim/go; WebSocket connection closed:", event)
					observer.complete()
				}
			}

			let retryCount = 10
			attemptConnection()

			return () => {
				// console.log("github.com/fullstack-lang/gong/lib/sim/go; Cleaning up WebSocket connection")
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
export function getCommandUniqueID(id: number): number {
	return 31 * id
}
export function getDummyAgentUniqueID(id: number): number {
	return 37 * id
}
export function getEngineUniqueID(id: number): number {
	return 41 * id
}
export function getEventUniqueID(id: number): number {
	return 43 * id
}
export function getStatusUniqueID(id: number): number {
	return 47 * id
}
export function getUpdateStateUniqueID(id: number): number {
	return 53 * id
}

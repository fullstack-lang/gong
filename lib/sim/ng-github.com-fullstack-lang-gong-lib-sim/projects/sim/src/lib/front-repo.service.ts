import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CommandAPI } from './command-api'
import { Command, CopyCommandAPIToCommand } from './command'
import { CommandService } from './command.service'

import { DummyAgentAPI } from './dummyagent-api'
import { DummyAgent, CopyDummyAgentAPIToDummyAgent } from './dummyagent'
import { DummyAgentService } from './dummyagent.service'

import { EngineAPI } from './engine-api'
import { Engine, CopyEngineAPIToEngine } from './engine'
import { EngineService } from './engine.service'

import { EventAPI } from './event-api'
import { Event, CopyEventAPIToEvent } from './event'
import { EventService } from './event.service'

import { StatusAPI } from './status-api'
import { Status, CopyStatusAPIToStatus } from './status'
import { StatusService } from './status.service'

import { UpdateStateAPI } from './updatestate-api'
import { UpdateState, CopyUpdateStateAPIToUpdateState } from './updatestate'
import { UpdateStateService } from './updatestate.service'


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
				throw new Error("Type not recognized");
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
		private commandService: CommandService,
		private dummyagentService: DummyAgentService,
		private engineService: EngineService,
		private eventService: EventService,
		private statusService: StatusService,
		private updatestateService: UpdateStateService,
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
		Observable<CommandAPI[]>,
		Observable<DummyAgentAPI[]>,
		Observable<EngineAPI[]>,
		Observable<EventAPI[]>,
		Observable<StatusAPI[]>,
		Observable<UpdateStateAPI[]>,
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
			this.commandService.getCommands(this.Name, this.frontRepo),
			this.dummyagentService.getDummyAgents(this.Name, this.frontRepo),
			this.engineService.getEngines(this.Name, this.frontRepo),
			this.eventService.getEvents(this.Name, this.frontRepo),
			this.statusService.getStatuss(this.Name, this.frontRepo),
			this.updatestateService.getUpdateStates(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						commands_,
						dummyagents_,
						engines_,
						events_,
						statuss_,
						updatestates_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var commands: CommandAPI[]
						commands = commands_ as CommandAPI[]
						var dummyagents: DummyAgentAPI[]
						dummyagents = dummyagents_ as DummyAgentAPI[]
						var engines: EngineAPI[]
						engines = engines_ as EngineAPI[]
						var events: EventAPI[]
						events = events_ as EventAPI[]
						var statuss: StatusAPI[]
						statuss = statuss_ as StatusAPI[]
						var updatestates: UpdateStateAPI[]
						updatestates = updatestates_ as UpdateStateAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Commands = []
						this.frontRepo.map_ID_Command.clear()

						commands.forEach(
							commandAPI => {
								let command = new Command
								this.frontRepo.array_Commands.push(command)
								this.frontRepo.map_ID_Command.set(commandAPI.ID, command)
							}
						)

						// init the arrays
						this.frontRepo.array_DummyAgents = []
						this.frontRepo.map_ID_DummyAgent.clear()

						dummyagents.forEach(
							dummyagentAPI => {
								let dummyagent = new DummyAgent
								this.frontRepo.array_DummyAgents.push(dummyagent)
								this.frontRepo.map_ID_DummyAgent.set(dummyagentAPI.ID, dummyagent)
							}
						)

						// init the arrays
						this.frontRepo.array_Engines = []
						this.frontRepo.map_ID_Engine.clear()

						engines.forEach(
							engineAPI => {
								let engine = new Engine
								this.frontRepo.array_Engines.push(engine)
								this.frontRepo.map_ID_Engine.set(engineAPI.ID, engine)
							}
						)

						// init the arrays
						this.frontRepo.array_Events = []
						this.frontRepo.map_ID_Event.clear()

						events.forEach(
							eventAPI => {
								let event = new Event
								this.frontRepo.array_Events.push(event)
								this.frontRepo.map_ID_Event.set(eventAPI.ID, event)
							}
						)

						// init the arrays
						this.frontRepo.array_Statuss = []
						this.frontRepo.map_ID_Status.clear()

						statuss.forEach(
							statusAPI => {
								let status = new Status
								this.frontRepo.array_Statuss.push(status)
								this.frontRepo.map_ID_Status.set(statusAPI.ID, status)
							}
						)

						// init the arrays
						this.frontRepo.array_UpdateStates = []
						this.frontRepo.map_ID_UpdateState.clear()

						updatestates.forEach(
							updatestateAPI => {
								let updatestate = new UpdateState
								this.frontRepo.array_UpdateStates.push(updatestate)
								this.frontRepo.map_ID_UpdateState.set(updatestateAPI.ID, updatestate)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						commands.forEach(
							commandAPI => {
								let command = this.frontRepo.map_ID_Command.get(commandAPI.ID)
								CopyCommandAPIToCommand(commandAPI, command!, this.frontRepo)
							}
						)

						// fill up front objects
						dummyagents.forEach(
							dummyagentAPI => {
								let dummyagent = this.frontRepo.map_ID_DummyAgent.get(dummyagentAPI.ID)
								CopyDummyAgentAPIToDummyAgent(dummyagentAPI, dummyagent!, this.frontRepo)
							}
						)

						// fill up front objects
						engines.forEach(
							engineAPI => {
								let engine = this.frontRepo.map_ID_Engine.get(engineAPI.ID)
								CopyEngineAPIToEngine(engineAPI, engine!, this.frontRepo)
							}
						)

						// fill up front objects
						events.forEach(
							eventAPI => {
								let event = this.frontRepo.map_ID_Event.get(eventAPI.ID)
								CopyEventAPIToEvent(eventAPI, event!, this.frontRepo)
							}
						)

						// fill up front objects
						statuss.forEach(
							statusAPI => {
								let status = this.frontRepo.map_ID_Status.get(statusAPI.ID)
								CopyStatusAPIToStatus(statusAPI, status!, this.frontRepo)
							}
						)

						// fill up front objects
						updatestates.forEach(
							updatestateAPI => {
								let updatestate = this.frontRepo.map_ID_UpdateState.get(updatestateAPI.ID)
								CopyUpdateStateAPIToUpdateState(updatestateAPI, updatestate!, this.frontRepo)
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/sim/go/v1/ws/stage`

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

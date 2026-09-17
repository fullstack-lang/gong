// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'

import { ButtonToggleAPI } from './buttontoggle-api'
import { ButtonToggle, CopyButtonToggleAPIToButtonToggle } from './buttontoggle'

import { GroupAPI } from './group-api'
import { Group, CopyGroupAPIToGroup } from './group'

import { GroupToogleAPI } from './grouptoogle-api'
import { GroupToogle, CopyGroupToogleAPIToGroupToogle } from './grouptoogle'

import { LayoutAPI } from './layout-api'
import { Layout, CopyLayoutAPIToLayout } from './layout'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/button/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Buttons = new Array<Button>() // array of front instances
	map_ID_Button = new Map<number, Button>() // map of front instances

	array_ButtonToggles = new Array<ButtonToggle>() // array of front instances
	map_ID_ButtonToggle = new Map<number, ButtonToggle>() // map of front instances

	array_Groups = new Array<Group>() // array of front instances
	map_ID_Group = new Map<number, Group>() // map of front instances

	array_GroupToogles = new Array<GroupToogle>() // array of front instances
	map_ID_GroupToogle = new Map<number, GroupToogle>() // map of front instances

	array_Layouts = new Array<Layout>() // array of front instances
	map_ID_Layout = new Map<number, Layout>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.array_Buttons as unknown as Array<Type>
			case 'ButtonToggle':
				return this.array_ButtonToggles as unknown as Array<Type>
			case 'Group':
				return this.array_Groups as unknown as Array<Type>
			case 'GroupToogle':
				return this.array_GroupToogles as unknown as Array<Type>
			case 'Layout':
				return this.array_Layouts as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.map_ID_Button as unknown as Map<number, Type>
			case 'ButtonToggle':
				return this.map_ID_ButtonToggle as unknown as Map<number, Type>
			case 'Group':
				return this.map_ID_Group as unknown as Map<number, Type>
			case 'GroupToogle':
				return this.map_ID_GroupToogle as unknown as Map<number, Type>
			case 'Layout':
				return this.map_ID_Layout as unknown as Map<number, Type>
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

		// console.log("github.com/fullstack-lang/gong/lib/button/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/button/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/button/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/button/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/button/go; connectToWebSocket: processData called")
				const backRepoData = new BackRepoData(JSON.parse(dataString))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_Buttons = []
				frontRepo.map_ID_Button.clear()

				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = new Button
						frontRepo.array_Buttons.push(button)
						frontRepo.map_ID_Button.set(buttonAPI.ID, button)
					}
				)

				// init the arrays
				frontRepo.array_ButtonToggles = []
				frontRepo.map_ID_ButtonToggle.clear()

				backRepoData.ButtonToggleAPIs.forEach(
					buttontoggleAPI => {
						let buttontoggle = new ButtonToggle
						frontRepo.array_ButtonToggles.push(buttontoggle)
						frontRepo.map_ID_ButtonToggle.set(buttontoggleAPI.ID, buttontoggle)
					}
				)

				// init the arrays
				frontRepo.array_Groups = []
				frontRepo.map_ID_Group.clear()

				backRepoData.GroupAPIs.forEach(
					groupAPI => {
						let group = new Group
						frontRepo.array_Groups.push(group)
						frontRepo.map_ID_Group.set(groupAPI.ID, group)
					}
				)

				// init the arrays
				frontRepo.array_GroupToogles = []
				frontRepo.map_ID_GroupToogle.clear()

				backRepoData.GroupToogleAPIs.forEach(
					grouptoogleAPI => {
						let grouptoogle = new GroupToogle
						frontRepo.array_GroupToogles.push(grouptoogle)
						frontRepo.map_ID_GroupToogle.set(grouptoogleAPI.ID, grouptoogle)
					}
				)

				// init the arrays
				frontRepo.array_Layouts = []
				frontRepo.map_ID_Layout.clear()

				backRepoData.LayoutAPIs.forEach(
					layoutAPI => {
						let layout = new Layout
						frontRepo.array_Layouts.push(layout)
						frontRepo.map_ID_Layout.set(layoutAPI.ID, layout)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = frontRepo.map_ID_Button.get(buttonAPI.ID)
						CopyButtonAPIToButton(buttonAPI, button!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ButtonToggleAPIs.forEach(
					buttontoggleAPI => {
						let buttontoggle = frontRepo.map_ID_ButtonToggle.get(buttontoggleAPI.ID)
						CopyButtonToggleAPIToButtonToggle(buttontoggleAPI, buttontoggle!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GroupAPIs.forEach(
					groupAPI => {
						let group = frontRepo.map_ID_Group.get(groupAPI.ID)
						CopyGroupAPIToGroup(groupAPI, group!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GroupToogleAPIs.forEach(
					grouptoogleAPI => {
						let grouptoogle = frontRepo.map_ID_GroupToogle.get(grouptoogleAPI.ID)
						CopyGroupToogleAPIToGroupToogle(grouptoogleAPI, grouptoogle!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LayoutAPIs.forEach(
					layoutAPI => {
						let layout = frontRepo.map_ID_Layout.get(layoutAPI.ID)
						CopyLayoutAPIToLayout(layoutAPI, layout!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// Offline mode handling: Listen to the global event
			if (isOfflineMode) {
				console.log("github.com/fullstack-lang/gong/lib/button/go; Offline mode detected. Skipping WebSocket connection.")

				window.addEventListener('message', (event) => {
					if (event.data && event.data.type === 'STAGE_UPDATE') {
						console.log("github.com/fullstack-lang/gong/lib/button/go; Received STAGE_UPDATE message.")
						processData(JSON.stringify(event.data.data))
					}
				})

				return () => {
					console.log("github.com/fullstack-lang/gong/lib/button/go; Cleaning up offline message listener.")
				}
			}

			// Fallback: If not offline, create normal WebSocket
			const attemptConnection = () => {
				// Offline check inside attemptConnection: if window.openWasmSocket is available, use it!
				if (typeof window !== 'undefined' && (window as any).openWasmSocket) {
					(window as any).openWasmSocket('github.com/fullstack-lang/gong/lib/button/go', Name, (data: any) => {
						processData(data)
					})
					return
				}

				if (isOfflineMode && retryCount > 0) {
					console.log("github.com/fullstack-lang/gong/lib/button/go; Waiting for wasm socket provider...")
					setTimeout(() => attemptConnection(), 100)
					return
				}

				if (isOfflineMode) {
					console.error("github.com/fullstack-lang/gong/lib/button/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.")
					return
				}

				socket = new WebSocket(url)

				socket.onopen = () => {
					// console.log("github.com/fullstack-lang/gong/lib/button/go; WebSocket connection opened successfully:", url)
				}

				socket.onmessage = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/button/go; WebSocket message received:", event.data)
					processData(event.data)
				}

				socket.onerror = (error) => {
					console.error("github.com/fullstack-lang/gong/lib/button/go WebSocket: onerror", error)
					observer.error(error)
				}

				socket.onclose = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/button/go; WebSocket connection closed:", event)
					observer.complete()
				}
			}

			let retryCount = 10
			attemptConnection()

			return () => {
				// console.log("github.com/fullstack-lang/gong/lib/button/go; Cleaning up WebSocket connection")
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
export function getButtonUniqueID(id: number): number {
	return 31 * id
}
export function getButtonToggleUniqueID(id: number): number {
	return 37 * id
}
export function getGroupUniqueID(id: number): number {
	return 41 * id
}
export function getGroupToogleUniqueID(id: number): number {
	return 43 * id
}
export function getLayoutUniqueID(id: number): number {
	return 47 * id
}

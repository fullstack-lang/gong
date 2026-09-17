// generated code - do not edit
import { Injectable, NgZone } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'

import { MenuAPI } from './menu-api'
import { Menu, CopyMenuAPIToMenu } from './menu'

import { NodeAPI } from './node-api'
import { Node, CopyNodeAPIToNode } from './node'

import { SVGIconAPI } from './svgicon-api'
import { SVGIcon, CopySVGIconAPIToSVGIcon } from './svgicon'

import { TreeAPI } from './tree-api'
import { Tree, CopyTreeAPIToTree } from './tree'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/tree/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Buttons = new Array<Button>() // array of front instances
	map_ID_Button = new Map<number, Button>() // map of front instances

	array_Menus = new Array<Menu>() // array of front instances
	map_ID_Menu = new Map<number, Menu>() // map of front instances

	array_Nodes = new Array<Node>() // array of front instances
	map_ID_Node = new Map<number, Node>() // map of front instances

	array_SVGIcons = new Array<SVGIcon>() // array of front instances
	map_ID_SVGIcon = new Map<number, SVGIcon>() // map of front instances

	array_Trees = new Array<Tree>() // array of front instances
	map_ID_Tree = new Map<number, Tree>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.array_Buttons as unknown as Array<Type>
			case 'Menu':
				return this.array_Menus as unknown as Array<Type>
			case 'Node':
				return this.array_Nodes as unknown as Array<Type>
			case 'SVGIcon':
				return this.array_SVGIcons as unknown as Array<Type>
			case 'Tree':
				return this.array_Trees as unknown as Array<Type>
			default:
				throw new Error("Type not recognized")
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.map_ID_Button as unknown as Map<number, Type>
			case 'Menu':
				return this.map_ID_Menu as unknown as Map<number, Type>
			case 'Node':
				return this.map_ID_Node as unknown as Map<number, Type>
			case 'SVGIcon':
				return this.map_ID_SVGIcon as unknown as Map<number, Type>
			case 'Tree':
				return this.map_ID_Tree as unknown as Map<number, Type>
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

		// console.log("github.com/fullstack-lang/gong/lib/tree/go; connectToWebSocket: started", Name)

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
			// console.log("github.com/fullstack-lang/gong/lib/tree/go; connectToWebSocket: returning existing connection")
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
		let basePath = `${protocol}//${host}/api/github.com/fullstack-lang/gong/lib/tree/go/v1/ws/stage`

		let params = new HttpParams().set("Name", Name)
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`

		const newConnection$ = new Observable<FrontRepo>(observer => {
			// console.log("github.com/fullstack-lang/gong/lib/tree/go; connectToWebSocket: new Observable created")

			let socket: WebSocket | undefined

			const isOfflineMode = window.location.protocol === 'file:' || window.document.getElementById('wasm-progress-container') !== null

			const processData = (dataString: string) => {
				// console.log("github.com/fullstack-lang/gong/lib/tree/go; connectToWebSocket: processData called")
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
				frontRepo.array_Menus = []
				frontRepo.map_ID_Menu.clear()

				backRepoData.MenuAPIs.forEach(
					menuAPI => {
						let menu = new Menu
						frontRepo.array_Menus.push(menu)
						frontRepo.map_ID_Menu.set(menuAPI.ID, menu)
					}
				)

				// init the arrays
				frontRepo.array_Nodes = []
				frontRepo.map_ID_Node.clear()

				backRepoData.NodeAPIs.forEach(
					nodeAPI => {
						let node = new Node
						frontRepo.array_Nodes.push(node)
						frontRepo.map_ID_Node.set(nodeAPI.ID, node)
					}
				)

				// init the arrays
				frontRepo.array_SVGIcons = []
				frontRepo.map_ID_SVGIcon.clear()

				backRepoData.SVGIconAPIs.forEach(
					svgiconAPI => {
						let svgicon = new SVGIcon
						frontRepo.array_SVGIcons.push(svgicon)
						frontRepo.map_ID_SVGIcon.set(svgiconAPI.ID, svgicon)
					}
				)

				// init the arrays
				frontRepo.array_Trees = []
				frontRepo.map_ID_Tree.clear()

				backRepoData.TreeAPIs.forEach(
					treeAPI => {
						let tree = new Tree
						frontRepo.array_Trees.push(tree)
						frontRepo.map_ID_Tree.set(treeAPI.ID, tree)
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
				backRepoData.MenuAPIs.forEach(
					menuAPI => {
						let menu = frontRepo.map_ID_Menu.get(menuAPI.ID)
						CopyMenuAPIToMenu(menuAPI, menu!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.NodeAPIs.forEach(
					nodeAPI => {
						let node = frontRepo.map_ID_Node.get(nodeAPI.ID)
						CopyNodeAPIToNode(nodeAPI, node!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SVGIconAPIs.forEach(
					svgiconAPI => {
						let svgicon = frontRepo.map_ID_SVGIcon.get(svgiconAPI.ID)
						CopySVGIconAPIToSVGIcon(svgiconAPI, svgicon!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TreeAPIs.forEach(
					treeAPI => {
						let tree = frontRepo.map_ID_Tree.get(treeAPI.ID)
						CopyTreeAPIToTree(treeAPI, tree!, frontRepo)
					}
				)


				this.ngZone.run(() => {
					observer.next(frontRepo)
				})
			}

			// Offline mode handling: Listen to the global event
			if (isOfflineMode) {
				console.log("github.com/fullstack-lang/gong/lib/tree/go; Offline mode detected. Skipping WebSocket connection.")

				window.addEventListener('message', (event) => {
					if (event.data && event.data.type === 'STAGE_UPDATE') {
						console.log("github.com/fullstack-lang/gong/lib/tree/go; Received STAGE_UPDATE message.")
						processData(JSON.stringify(event.data.data))
					}
				})

				return () => {
					console.log("github.com/fullstack-lang/gong/lib/tree/go; Cleaning up offline message listener.")
				}
			}

			// Fallback: If not offline, create normal WebSocket
			const attemptConnection = () => {
				// Offline check inside attemptConnection: if window.openWasmSocket is available, use it!
				if (typeof window !== 'undefined' && (window as any).openWasmSocket) {
					(window as any).openWasmSocket('github.com/fullstack-lang/gong/lib/tree/go', Name, (data: any) => {
						processData(data)
					})
					return
				}

				if (isOfflineMode && retryCount > 0) {
					console.log("github.com/fullstack-lang/gong/lib/tree/go; Waiting for wasm socket provider...")
					setTimeout(() => attemptConnection(), 100)
					return
				}

				if (isOfflineMode) {
					console.error("github.com/fullstack-lang/gong/lib/tree/go, attemptConnection: Offline mode detected, but WASM backend failed to load.")
					observer.error("Offline mode detected, but WASM backend failed to load.")
					return
				}

				socket = new WebSocket(url)

				socket.onopen = () => {
					// console.log("github.com/fullstack-lang/gong/lib/tree/go; WebSocket connection opened successfully:", url)
				}

				socket.onmessage = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/tree/go; WebSocket message received:", event.data)
					processData(event.data)
				}

				socket.onerror = (error) => {
					console.error("github.com/fullstack-lang/gong/lib/tree/go WebSocket: onerror", error)
					observer.error(error)
				}

				socket.onclose = (event) => {
					// console.log("github.com/fullstack-lang/gong/lib/tree/go; WebSocket connection closed:", event)
					observer.complete()
				}
			}

			let retryCount = 10
			attemptConnection()

			return () => {
				// console.log("github.com/fullstack-lang/gong/lib/tree/go; Cleaning up WebSocket connection")
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
export function getMenuUniqueID(id: number): number {
	return 37 * id
}
export function getNodeUniqueID(id: number): number {
	return 41 * id
}
export function getSVGIconUniqueID(id: number): number {
	return 43 * id
}
export function getTreeUniqueID(id: number): number {
	return 47 * id
}

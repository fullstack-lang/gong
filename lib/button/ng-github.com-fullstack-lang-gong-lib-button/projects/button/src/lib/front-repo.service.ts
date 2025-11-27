import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'
import { shareReplay } from 'rxjs/operators'

// insertion point sub template for services imports
import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'
import { ButtonService } from './button.service'

import { ButtonToggleAPI } from './buttontoggle-api'
import { ButtonToggle, CopyButtonToggleAPIToButtonToggle } from './buttontoggle'
import { ButtonToggleService } from './buttontoggle.service'

import { GroupAPI } from './group-api'
import { Group, CopyGroupAPIToGroup } from './group'
import { GroupService } from './group.service'

import { GroupToogleAPI } from './grouptoogle-api'
import { GroupToogle, CopyGroupToogleAPIToGroupToogle } from './grouptoogle'
import { GroupToogleService } from './grouptoogle.service'

import { LayoutAPI } from './layout-api'
import { Layout, CopyLayoutAPIToLayout } from './layout'
import { LayoutService } from './layout.service'


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
		private http: HttpClient, // insertion point sub template 
		private buttonService: ButtonService,
		private buttontoggleService: ButtonToggleService,
		private groupService: GroupService,
		private grouptoogleService: GroupToogleService,
		private layoutService: LayoutService,
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
		Observable<ButtonAPI[]>,
		Observable<ButtonToggleAPI[]>,
		Observable<GroupAPI[]>,
		Observable<GroupToogleAPI[]>,
		Observable<LayoutAPI[]>,
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
			this.buttonService.getButtons(this.Name, this.frontRepo),
			this.buttontoggleService.getButtonToggles(this.Name, this.frontRepo),
			this.groupService.getGroups(this.Name, this.frontRepo),
			this.grouptoogleService.getGroupToogles(this.Name, this.frontRepo),
			this.layoutService.getLayouts(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						buttons_,
						buttontoggles_,
						groups_,
						grouptoogles_,
						layouts_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var buttons: ButtonAPI[]
						buttons = buttons_ as ButtonAPI[]
						var buttontoggles: ButtonToggleAPI[]
						buttontoggles = buttontoggles_ as ButtonToggleAPI[]
						var groups: GroupAPI[]
						groups = groups_ as GroupAPI[]
						var grouptoogles: GroupToogleAPI[]
						grouptoogles = grouptoogles_ as GroupToogleAPI[]
						var layouts: LayoutAPI[]
						layouts = layouts_ as LayoutAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Buttons = []
						this.frontRepo.map_ID_Button.clear()

						buttons.forEach(
							buttonAPI => {
								let button = new Button
								this.frontRepo.array_Buttons.push(button)
								this.frontRepo.map_ID_Button.set(buttonAPI.ID, button)
							}
						)

						// init the arrays
						this.frontRepo.array_ButtonToggles = []
						this.frontRepo.map_ID_ButtonToggle.clear()

						buttontoggles.forEach(
							buttontoggleAPI => {
								let buttontoggle = new ButtonToggle
								this.frontRepo.array_ButtonToggles.push(buttontoggle)
								this.frontRepo.map_ID_ButtonToggle.set(buttontoggleAPI.ID, buttontoggle)
							}
						)

						// init the arrays
						this.frontRepo.array_Groups = []
						this.frontRepo.map_ID_Group.clear()

						groups.forEach(
							groupAPI => {
								let group = new Group
								this.frontRepo.array_Groups.push(group)
								this.frontRepo.map_ID_Group.set(groupAPI.ID, group)
							}
						)

						// init the arrays
						this.frontRepo.array_GroupToogles = []
						this.frontRepo.map_ID_GroupToogle.clear()

						grouptoogles.forEach(
							grouptoogleAPI => {
								let grouptoogle = new GroupToogle
								this.frontRepo.array_GroupToogles.push(grouptoogle)
								this.frontRepo.map_ID_GroupToogle.set(grouptoogleAPI.ID, grouptoogle)
							}
						)

						// init the arrays
						this.frontRepo.array_Layouts = []
						this.frontRepo.map_ID_Layout.clear()

						layouts.forEach(
							layoutAPI => {
								let layout = new Layout
								this.frontRepo.array_Layouts.push(layout)
								this.frontRepo.map_ID_Layout.set(layoutAPI.ID, layout)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						buttons.forEach(
							buttonAPI => {
								let button = this.frontRepo.map_ID_Button.get(buttonAPI.ID)
								CopyButtonAPIToButton(buttonAPI, button!, this.frontRepo)
							}
						)

						// fill up front objects
						buttontoggles.forEach(
							buttontoggleAPI => {
								let buttontoggle = this.frontRepo.map_ID_ButtonToggle.get(buttontoggleAPI.ID)
								CopyButtonToggleAPIToButtonToggle(buttontoggleAPI, buttontoggle!, this.frontRepo)
							}
						)

						// fill up front objects
						groups.forEach(
							groupAPI => {
								let group = this.frontRepo.map_ID_Group.get(groupAPI.ID)
								CopyGroupAPIToGroup(groupAPI, group!, this.frontRepo)
							}
						)

						// fill up front objects
						grouptoogles.forEach(
							grouptoogleAPI => {
								let grouptoogle = this.frontRepo.map_ID_GroupToogle.get(grouptoogleAPI.ID)
								CopyGroupToogleAPIToGroupToogle(grouptoogleAPI, grouptoogle!, this.frontRepo)
							}
						)

						// fill up front objects
						layouts.forEach(
							layoutAPI => {
								let layout = this.frontRepo.map_ID_Layout.get(layoutAPI.ID)
								CopyLayoutAPIToLayout(layoutAPI, layout!, this.frontRepo)
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

		// Check if a connection for this name already exists
		if (this.webSocketConnections.has(Name)) {
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
			const socket = new WebSocket(url)

			socket.onmessage = event => {
				const backRepoData = new BackRepoData(JSON.parse(event.data))
				let frontRepo = new (FrontRepo)()
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
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


				observer.next(frontRepo)
			}

			socket.onerror = event => observer.error(event)
			socket.onclose = () => observer.complete()

			// Teardown logic: Called when the last subscriber unsubscribes.
			return () => {
				this.webSocketConnections.delete(Name) // Remove from cache
				socket.close()
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

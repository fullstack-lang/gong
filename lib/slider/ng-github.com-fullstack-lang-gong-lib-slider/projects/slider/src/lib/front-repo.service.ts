import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CheckboxAPI } from './checkbox-api'
import { Checkbox, CopyCheckboxAPIToCheckbox } from './checkbox'
import { CheckboxService } from './checkbox.service'

import { GroupAPI } from './group-api'
import { Group, CopyGroupAPIToGroup } from './group'
import { GroupService } from './group.service'

import { LayoutAPI } from './layout-api'
import { Layout, CopyLayoutAPIToLayout } from './layout'
import { LayoutService } from './layout.service'

import { SliderAPI } from './slider-api'
import { Slider, CopySliderAPIToSlider } from './slider'
import { SliderService } from './slider.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/lib/slider/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Checkboxs = new Array<Checkbox>() // array of front instances
	map_ID_Checkbox = new Map<number, Checkbox>() // map of front instances

	array_Groups = new Array<Group>() // array of front instances
	map_ID_Group = new Map<number, Group>() // map of front instances

	array_Layouts = new Array<Layout>() // array of front instances
	map_ID_Layout = new Map<number, Layout>() // map of front instances

	array_Sliders = new Array<Slider>() // array of front instances
	map_ID_Slider = new Map<number, Slider>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Checkbox':
				return this.array_Checkboxs as unknown as Array<Type>
			case 'Group':
				return this.array_Groups as unknown as Array<Type>
			case 'Layout':
				return this.array_Layouts as unknown as Array<Type>
			case 'Slider':
				return this.array_Sliders as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Checkbox':
				return this.map_ID_Checkbox as unknown as Map<number, Type>
			case 'Group':
				return this.map_ID_Group as unknown as Map<number, Type>
			case 'Layout':
				return this.map_ID_Layout as unknown as Map<number, Type>
			case 'Slider':
				return this.map_ID_Slider as unknown as Map<number, Type>
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
		private checkboxService: CheckboxService,
		private groupService: GroupService,
		private layoutService: LayoutService,
		private sliderService: SliderService,
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
		Observable<CheckboxAPI[]>,
		Observable<GroupAPI[]>,
		Observable<LayoutAPI[]>,
		Observable<SliderAPI[]>,
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
			this.checkboxService.getCheckboxs(this.Name, this.frontRepo),
			this.groupService.getGroups(this.Name, this.frontRepo),
			this.layoutService.getLayouts(this.Name, this.frontRepo),
			this.sliderService.getSliders(this.Name, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						checkboxs_,
						groups_,
						layouts_,
						sliders_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var checkboxs: CheckboxAPI[]
						checkboxs = checkboxs_ as CheckboxAPI[]
						var groups: GroupAPI[]
						groups = groups_ as GroupAPI[]
						var layouts: LayoutAPI[]
						layouts = layouts_ as LayoutAPI[]
						var sliders: SliderAPI[]
						sliders = sliders_ as SliderAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Checkboxs = []
						this.frontRepo.map_ID_Checkbox.clear()

						checkboxs.forEach(
							checkboxAPI => {
								let checkbox = new Checkbox
								this.frontRepo.array_Checkboxs.push(checkbox)
								this.frontRepo.map_ID_Checkbox.set(checkboxAPI.ID, checkbox)
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
						this.frontRepo.array_Layouts = []
						this.frontRepo.map_ID_Layout.clear()

						layouts.forEach(
							layoutAPI => {
								let layout = new Layout
								this.frontRepo.array_Layouts.push(layout)
								this.frontRepo.map_ID_Layout.set(layoutAPI.ID, layout)
							}
						)

						// init the arrays
						this.frontRepo.array_Sliders = []
						this.frontRepo.map_ID_Slider.clear()

						sliders.forEach(
							sliderAPI => {
								let slider = new Slider
								this.frontRepo.array_Sliders.push(slider)
								this.frontRepo.map_ID_Slider.set(sliderAPI.ID, slider)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						checkboxs.forEach(
							checkboxAPI => {
								let checkbox = this.frontRepo.map_ID_Checkbox.get(checkboxAPI.ID)
								CopyCheckboxAPIToCheckbox(checkboxAPI, checkbox!, this.frontRepo)
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
						layouts.forEach(
							layoutAPI => {
								let layout = this.frontRepo.map_ID_Layout.get(layoutAPI.ID)
								CopyLayoutAPIToLayout(layoutAPI, layout!, this.frontRepo)
							}
						)

						// fill up front objects
						sliders.forEach(
							sliderAPI => {
								let slider = this.frontRepo.map_ID_Slider.get(sliderAPI.ID)
								CopySliderAPIToSlider(sliderAPI, slider!, this.frontRepo)
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


		let params = new HttpParams().set("Name", this.Name)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/lib/slider/go/v1/ws/stage'
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
				frontRepo.array_Checkboxs = []
				frontRepo.map_ID_Checkbox.clear()

				backRepoData.CheckboxAPIs.forEach(
					checkboxAPI => {
						let checkbox = new Checkbox
						frontRepo.array_Checkboxs.push(checkbox)
						frontRepo.map_ID_Checkbox.set(checkboxAPI.ID, checkbox)
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
				frontRepo.array_Layouts = []
				frontRepo.map_ID_Layout.clear()

				backRepoData.LayoutAPIs.forEach(
					layoutAPI => {
						let layout = new Layout
						frontRepo.array_Layouts.push(layout)
						frontRepo.map_ID_Layout.set(layoutAPI.ID, layout)
					}
				)

				// init the arrays
				frontRepo.array_Sliders = []
				frontRepo.map_ID_Slider.clear()

				backRepoData.SliderAPIs.forEach(
					sliderAPI => {
						let slider = new Slider
						frontRepo.array_Sliders.push(slider)
						frontRepo.map_ID_Slider.set(sliderAPI.ID, slider)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.CheckboxAPIs.forEach(
					checkboxAPI => {
						let checkbox = frontRepo.map_ID_Checkbox.get(checkboxAPI.ID)
						CopyCheckboxAPIToCheckbox(checkboxAPI, checkbox!, frontRepo)
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
				backRepoData.LayoutAPIs.forEach(
					layoutAPI => {
						let layout = frontRepo.map_ID_Layout.get(layoutAPI.ID)
						CopyLayoutAPIToLayout(layoutAPI, layout!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SliderAPIs.forEach(
					sliderAPI => {
						let slider = frontRepo.map_ID_Slider.get(sliderAPI.ID)
						CopySliderAPIToSlider(sliderAPI, slider!, frontRepo)
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
export function getCheckboxUniqueID(id: number): number {
	return 31 * id
}
export function getGroupUniqueID(id: number): number {
	return 37 * id
}
export function getLayoutUniqueID(id: number): number {
	return 41 * id
}
export function getSliderUniqueID(id: number): number {
	return 43 * id
}

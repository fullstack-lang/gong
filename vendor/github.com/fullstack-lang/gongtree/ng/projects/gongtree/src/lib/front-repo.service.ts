import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { ButtonAPI } from './button-api'
import { Button, CopyButtonAPIToButton } from './button'
import { ButtonService } from './button.service'

import { NodeAPI } from './node-api'
import { Node, CopyNodeAPIToNode } from './node'
import { NodeService } from './node.service'

import { SVGIconAPI } from './svgicon-api'
import { SVGIcon, CopySVGIconAPIToSVGIcon } from './svgicon'
import { SVGIconService } from './svgicon.service'

import { TreeAPI } from './tree-api'
import { Tree, CopyTreeAPIToTree } from './tree'
import { TreeService } from './tree.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongtree/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Buttons = new Array<Button>() // array of front instances
	map_ID_Button = new Map<number, Button>() // map of front instances

	array_Nodes = new Array<Node>() // array of front instances
	map_ID_Node = new Map<number, Node>() // map of front instances

	array_SVGIcons = new Array<SVGIcon>() // array of front instances
	map_ID_SVGIcon = new Map<number, SVGIcon>() // map of front instances

	array_Trees = new Array<Tree>() // array of front instances
	map_ID_Tree = new Map<number, Tree>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.array_Buttons as unknown as Array<Type>
			case 'Node':
				return this.array_Nodes as unknown as Array<Type>
			case 'SVGIcon':
				return this.array_SVGIcons as unknown as Array<Type>
			case 'Tree':
				return this.array_Trees as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Button':
				return this.map_ID_Button as unknown as Map<number, Type>
			case 'Node':
				return this.map_ID_Node as unknown as Map<number, Type>
			case 'SVGIcon':
				return this.map_ID_SVGIcon as unknown as Map<number, Type>
			case 'Tree':
				return this.map_ID_Tree as unknown as Map<number, Type>
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
		private buttonService: ButtonService,
		private nodeService: NodeService,
		private svgiconService: SVGIconService,
		private treeService: TreeService,
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
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<ButtonAPI[]>,
		Observable<NodeAPI[]>,
		Observable<SVGIconAPI[]>,
		Observable<TreeAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.buttonService.getButtons(this.GONG__StackPath, this.frontRepo),
			this.nodeService.getNodes(this.GONG__StackPath, this.frontRepo),
			this.svgiconService.getSVGIcons(this.GONG__StackPath, this.frontRepo),
			this.treeService.getTrees(this.GONG__StackPath, this.frontRepo),
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
			this.buttonService.getButtons(this.GONG__StackPath, this.frontRepo),
			this.nodeService.getNodes(this.GONG__StackPath, this.frontRepo),
			this.svgiconService.getSVGIcons(this.GONG__StackPath, this.frontRepo),
			this.treeService.getTrees(this.GONG__StackPath, this.frontRepo),
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
						nodes_,
						svgicons_,
						trees_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var buttons: ButtonAPI[]
						buttons = buttons_ as ButtonAPI[]
						var nodes: NodeAPI[]
						nodes = nodes_ as NodeAPI[]
						var svgicons: SVGIconAPI[]
						svgicons = svgicons_ as SVGIconAPI[]
						var trees: TreeAPI[]
						trees = trees_ as TreeAPI[]

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
						this.frontRepo.array_Nodes = []
						this.frontRepo.map_ID_Node.clear()

						nodes.forEach(
							nodeAPI => {
								let node = new Node
								this.frontRepo.array_Nodes.push(node)
								this.frontRepo.map_ID_Node.set(nodeAPI.ID, node)
							}
						)

						// init the arrays
						this.frontRepo.array_SVGIcons = []
						this.frontRepo.map_ID_SVGIcon.clear()

						svgicons.forEach(
							svgiconAPI => {
								let svgicon = new SVGIcon
								this.frontRepo.array_SVGIcons.push(svgicon)
								this.frontRepo.map_ID_SVGIcon.set(svgiconAPI.ID, svgicon)
							}
						)

						// init the arrays
						this.frontRepo.array_Trees = []
						this.frontRepo.map_ID_Tree.clear()

						trees.forEach(
							treeAPI => {
								let tree = new Tree
								this.frontRepo.array_Trees.push(tree)
								this.frontRepo.map_ID_Tree.set(treeAPI.ID, tree)
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
						nodes.forEach(
							nodeAPI => {
								let node = this.frontRepo.map_ID_Node.get(nodeAPI.ID)
								CopyNodeAPIToNode(nodeAPI, node!, this.frontRepo)
							}
						)

						// fill up front objects
						svgicons.forEach(
							svgiconAPI => {
								let svgicon = this.frontRepo.map_ID_SVGIcon.get(svgiconAPI.ID)
								CopySVGIconAPIToSVGIcon(svgiconAPI, svgicon!, this.frontRepo)
							}
						)

						// fill up front objects
						trees.forEach(
							treeAPI => {
								let tree = this.frontRepo.map_ID_Tree.get(treeAPI.ID)
								CopyTreeAPIToTree(treeAPI, tree!, this.frontRepo)
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
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongtree/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {
				let _this = this

				const backRepoData = new BackRepoData(JSON.parse(event.data))

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				this.frontRepo.array_Buttons = []
				this.frontRepo.map_ID_Button.clear()

				backRepoData.ButtonAPIs.forEach(
					buttonAPI => {
						let button = new Button
						this.frontRepo.array_Buttons.push(button)
						this.frontRepo.map_ID_Button.set(buttonAPI.ID, button)
					}
				)

				// init the arrays
				this.frontRepo.array_Nodes = []
				this.frontRepo.map_ID_Node.clear()

				backRepoData.NodeAPIs.forEach(
					nodeAPI => {
						let node = new Node
						this.frontRepo.array_Nodes.push(node)
						this.frontRepo.map_ID_Node.set(nodeAPI.ID, node)
					}
				)

				// init the arrays
				this.frontRepo.array_SVGIcons = []
				this.frontRepo.map_ID_SVGIcon.clear()

				backRepoData.SVGIconAPIs.forEach(
					svgiconAPI => {
						let svgicon = new SVGIcon
						this.frontRepo.array_SVGIcons.push(svgicon)
						this.frontRepo.map_ID_SVGIcon.set(svgiconAPI.ID, svgicon)
					}
				)

				// init the arrays
				this.frontRepo.array_Trees = []
				this.frontRepo.map_ID_Tree.clear()

				backRepoData.TreeAPIs.forEach(
					treeAPI => {
						let tree = new Tree
						this.frontRepo.array_Trees.push(tree)
						this.frontRepo.map_ID_Tree.set(treeAPI.ID, tree)
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
						let button = this.frontRepo.map_ID_Button.get(buttonAPI.ID)
						CopyButtonAPIToButton(buttonAPI, button!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.NodeAPIs.forEach(
					nodeAPI => {
						let node = this.frontRepo.map_ID_Node.get(nodeAPI.ID)
						CopyNodeAPIToNode(nodeAPI, node!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SVGIconAPIs.forEach(
					svgiconAPI => {
						let svgicon = this.frontRepo.map_ID_SVGIcon.get(svgiconAPI.ID)
						CopySVGIconAPIToSVGIcon(svgiconAPI, svgicon!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.TreeAPIs.forEach(
					treeAPI => {
						let tree = this.frontRepo.map_ID_Tree.get(treeAPI.ID)
						CopyTreeAPIToTree(treeAPI, tree!, this.frontRepo)
					}
				)



				observer.next(this.frontRepo)
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
export function getButtonUniqueID(id: number): number {
	return 31 * id
}
export function getNodeUniqueID(id: number): number {
	return 37 * id
}
export function getSVGIconUniqueID(id: number): number {
	return 41 * id
}
export function getTreeUniqueID(id: number): number {
	return 43 * id
}

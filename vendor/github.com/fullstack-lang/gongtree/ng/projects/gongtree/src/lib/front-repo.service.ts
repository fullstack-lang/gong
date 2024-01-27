import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { ButtonDB } from './button-db'
import { Button, CopyButtonDBToButton } from './button'
import { ButtonService } from './button.service'

import { NodeDB } from './node-db'
import { Node, CopyNodeDBToNode } from './node'
import { NodeService } from './node.service'

import { SVGIconDB } from './svgicon-db'
import { SVGIcon, CopySVGIconDBToSVGIcon } from './svgicon'
import { SVGIconService } from './svgicon.service'

import { TreeDB } from './tree-db'
import { Tree, CopyTreeDBToTree } from './tree'
import { TreeService } from './tree.service'

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
		Observable<ButtonDB[]>,
		Observable<NodeDB[]>,
		Observable<SVGIconDB[]>,
		Observable<TreeDB[]>,
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
						var buttons: ButtonDB[]
						buttons = buttons_ as ButtonDB[]
						var nodes: NodeDB[]
						nodes = nodes_ as NodeDB[]
						var svgicons: SVGIconDB[]
						svgicons = svgicons_ as SVGIconDB[]
						var trees: TreeDB[]
						trees = trees_ as TreeDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Buttons = []
						this.frontRepo.map_ID_Button.clear()

						buttons.forEach(
							buttonDB => {
								let button = new Button
								this.frontRepo.array_Buttons.push(button)
								this.frontRepo.map_ID_Button.set(buttonDB.ID, button)
							}
						)

						// init the arrays
						this.frontRepo.array_Nodes = []
						this.frontRepo.map_ID_Node.clear()

						nodes.forEach(
							nodeDB => {
								let node = new Node
								this.frontRepo.array_Nodes.push(node)
								this.frontRepo.map_ID_Node.set(nodeDB.ID, node)
							}
						)

						// init the arrays
						this.frontRepo.array_SVGIcons = []
						this.frontRepo.map_ID_SVGIcon.clear()

						svgicons.forEach(
							svgiconDB => {
								let svgicon = new SVGIcon
								this.frontRepo.array_SVGIcons.push(svgicon)
								this.frontRepo.map_ID_SVGIcon.set(svgiconDB.ID, svgicon)
							}
						)

						// init the arrays
						this.frontRepo.array_Trees = []
						this.frontRepo.map_ID_Tree.clear()

						trees.forEach(
							treeDB => {
								let tree = new Tree
								this.frontRepo.array_Trees.push(tree)
								this.frontRepo.map_ID_Tree.set(treeDB.ID, tree)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						buttons.forEach(
							buttonDB => {
								let button = this.frontRepo.map_ID_Button.get(buttonDB.ID)
								CopyButtonDBToButton(buttonDB, button!, this.frontRepo)
							}
						)

						// fill up front objects
						nodes.forEach(
							nodeDB => {
								let node = this.frontRepo.map_ID_Node.get(nodeDB.ID)
								CopyNodeDBToNode(nodeDB, node!, this.frontRepo)
							}
						)

						// fill up front objects
						svgicons.forEach(
							svgiconDB => {
								let svgicon = this.frontRepo.map_ID_SVGIcon.get(svgiconDB.ID)
								CopySVGIconDBToSVGIcon(svgiconDB, svgicon!, this.frontRepo)
							}
						)

						// fill up front objects
						trees.forEach(
							treeDB => {
								let tree = this.frontRepo.map_ID_Tree.get(treeDB.ID)
								CopyTreeDBToTree(treeDB, tree!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
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

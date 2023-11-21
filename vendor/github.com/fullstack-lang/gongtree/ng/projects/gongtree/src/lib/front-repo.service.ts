import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { ButtonDB } from './button-db'
import { ButtonService } from './button.service'

import { NodeDB } from './node-db'
import { NodeService } from './node.service'

import { SVGIconDB } from './svgicon-db'
import { SVGIconService } from './svgicon.service'

import { TreeDB } from './tree-db'
import { TreeService } from './tree.service'

export const StackType = "github.com/fullstack-lang/gongtree/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  Buttons_array = new Array<ButtonDB>() // array of repo instances
  Buttons = new Map<number, ButtonDB>() // map of repo instances
  Buttons_batch = new Map<number, ButtonDB>() // same but only in last GET (for finding repo instances to delete)

  Nodes_array = new Array<NodeDB>() // array of repo instances
  Nodes = new Map<number, NodeDB>() // map of repo instances
  Nodes_batch = new Map<number, NodeDB>() // same but only in last GET (for finding repo instances to delete)

  SVGIcons_array = new Array<SVGIconDB>() // array of repo instances
  SVGIcons = new Map<number, SVGIconDB>() // map of repo instances
  SVGIcons_batch = new Map<number, SVGIconDB>() // same but only in last GET (for finding repo instances to delete)

  Trees_array = new Array<TreeDB>() // array of repo instances
  Trees = new Map<number, TreeDB>() // map of repo instances
  Trees_batch = new Map<number, TreeDB>() // same but only in last GET (for finding repo instances to delete)


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'Button':
        return this.Buttons_array as unknown as Array<Type>
      case 'Node':
        return this.Nodes_array as unknown as Array<Type>
      case 'SVGIcon':
        return this.SVGIcons_array as unknown as Array<Type>
      case 'Tree':
        return this.Trees_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'Button':
        return this.Buttons as unknown as Map<number, Type>
      case 'Node':
        return this.Nodes as unknown as Map<number, Type>
      case 'SVGIcon':
        return this.SVGIcons as unknown as Map<number, Type>
      case 'Tree':
        return this.Trees as unknown as Map<number, Type>
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
  SourceStruct: string = ""  // The "Aclass"
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
            // init the array
            this.frontRepo.Buttons_array = buttons

            // clear the map that counts Button in the GET
            this.frontRepo.Buttons_batch.clear()

            buttons.forEach(
              button => {
                this.frontRepo.Buttons.set(button.ID, button)
                this.frontRepo.Buttons_batch.set(button.ID, button)
              }
            )

            // clear buttons that are absent from the batch
            this.frontRepo.Buttons.forEach(
              button => {
                if (this.frontRepo.Buttons_batch.get(button.ID) == undefined) {
                  this.frontRepo.Buttons.delete(button.ID)
                }
              }
            )

            // sort Buttons_array array
            this.frontRepo.Buttons_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Nodes_array = nodes

            // clear the map that counts Node in the GET
            this.frontRepo.Nodes_batch.clear()

            nodes.forEach(
              node => {
                this.frontRepo.Nodes.set(node.ID, node)
                this.frontRepo.Nodes_batch.set(node.ID, node)
              }
            )

            // clear nodes that are absent from the batch
            this.frontRepo.Nodes.forEach(
              node => {
                if (this.frontRepo.Nodes_batch.get(node.ID) == undefined) {
                  this.frontRepo.Nodes.delete(node.ID)
                }
              }
            )

            // sort Nodes_array array
            this.frontRepo.Nodes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.SVGIcons_array = svgicons

            // clear the map that counts SVGIcon in the GET
            this.frontRepo.SVGIcons_batch.clear()

            svgicons.forEach(
              svgicon => {
                this.frontRepo.SVGIcons.set(svgicon.ID, svgicon)
                this.frontRepo.SVGIcons_batch.set(svgicon.ID, svgicon)
              }
            )

            // clear svgicons that are absent from the batch
            this.frontRepo.SVGIcons.forEach(
              svgicon => {
                if (this.frontRepo.SVGIcons_batch.get(svgicon.ID) == undefined) {
                  this.frontRepo.SVGIcons.delete(svgicon.ID)
                }
              }
            )

            // sort SVGIcons_array array
            this.frontRepo.SVGIcons_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Trees_array = trees

            // clear the map that counts Tree in the GET
            this.frontRepo.Trees_batch.clear()

            trees.forEach(
              tree => {
                this.frontRepo.Trees.set(tree.ID, tree)
                this.frontRepo.Trees_batch.set(tree.ID, tree)
              }
            )

            // clear trees that are absent from the batch
            this.frontRepo.Trees.forEach(
              tree => {
                if (this.frontRepo.Trees_batch.get(tree.ID) == undefined) {
                  this.frontRepo.Trees.delete(tree.ID)
                }
              }
            )

            // sort Trees_array array
            this.frontRepo.Trees_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            buttons.forEach(
              button => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field SVGIcon redeeming
                {
                  let _svgicon = this.frontRepo.SVGIcons.get(button.ButtonPointersEncoding.SVGIconID.Int64)
                  if (_svgicon) {
                    button.SVGIcon = _svgicon
                  }
                }
                // insertion point for pointers decoding
              }
            )
            nodes.forEach(
              node => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field PreceedingSVGIcon redeeming
                {
                  let _svgicon = this.frontRepo.SVGIcons.get(node.NodePointersEncoding.PreceedingSVGIconID.Int64)
                  if (_svgicon) {
                    node.PreceedingSVGIcon = _svgicon
                  }
                }
                // insertion point for pointers decoding
                node.Children = new Array<NodeDB>()
                for (let _id of node.NodePointersEncoding.Children) {
                  let _node = this.frontRepo.Nodes.get(_id)
                  if (_node != undefined) {
                    node.Children.push(_node!)
                  }
                }
                node.Buttons = new Array<ButtonDB>()
                for (let _id of node.NodePointersEncoding.Buttons) {
                  let _button = this.frontRepo.Buttons.get(_id)
                  if (_button != undefined) {
                    node.Buttons.push(_button!)
                  }
                }
              }
            )
            svgicons.forEach(
              svgicon => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            trees.forEach(
              tree => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                tree.RootNodes = new Array<NodeDB>()
                for (let _id of tree.TreePointersEncoding.RootNodes) {
                  let _node = this.frontRepo.Nodes.get(_id)
                  if (_node != undefined) {
                    tree.RootNodes.push(_node!)
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // ButtonPull performs a GET on Button of the stack and redeem association pointers 
  ButtonPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.buttonService.getButtons(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            buttons,
          ]) => {
            // init the array
            this.frontRepo.Buttons_array = buttons

            // clear the map that counts Button in the GET
            this.frontRepo.Buttons_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            buttons.forEach(
              button => {
                this.frontRepo.Buttons.set(button.ID, button)
                this.frontRepo.Buttons_batch.set(button.ID, button)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field SVGIcon redeeming
                {
                  let _svgicon = this.frontRepo.SVGIcons.get(button.ButtonPointersEncoding.SVGIconID.Int64)
                  if (_svgicon) {
                    button.SVGIcon = _svgicon
                  }
                }
              }
            )

            // clear buttons that are absent from the GET
            this.frontRepo.Buttons.forEach(
              button => {
                if (this.frontRepo.Buttons_batch.get(button.ID) == undefined) {
                  this.frontRepo.Buttons.delete(button.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // NodePull performs a GET on Node of the stack and redeem association pointers 
  NodePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.nodeService.getNodes(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            nodes,
          ]) => {
            // init the array
            this.frontRepo.Nodes_array = nodes

            // clear the map that counts Node in the GET
            this.frontRepo.Nodes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            nodes.forEach(
              node => {
                this.frontRepo.Nodes.set(node.ID, node)
                this.frontRepo.Nodes_batch.set(node.ID, node)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field PreceedingSVGIcon redeeming
                {
                  let _svgicon = this.frontRepo.SVGIcons.get(node.NodePointersEncoding.PreceedingSVGIconID.Int64)
                  if (_svgicon) {
                    node.PreceedingSVGIcon = _svgicon
                  }
                }
              }
            )

            // clear nodes that are absent from the GET
            this.frontRepo.Nodes.forEach(
              node => {
                if (this.frontRepo.Nodes_batch.get(node.ID) == undefined) {
                  this.frontRepo.Nodes.delete(node.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // SVGIconPull performs a GET on SVGIcon of the stack and redeem association pointers 
  SVGIconPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.svgiconService.getSVGIcons(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            svgicons,
          ]) => {
            // init the array
            this.frontRepo.SVGIcons_array = svgicons

            // clear the map that counts SVGIcon in the GET
            this.frontRepo.SVGIcons_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            svgicons.forEach(
              svgicon => {
                this.frontRepo.SVGIcons.set(svgicon.ID, svgicon)
                this.frontRepo.SVGIcons_batch.set(svgicon.ID, svgicon)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear svgicons that are absent from the GET
            this.frontRepo.SVGIcons.forEach(
              svgicon => {
                if (this.frontRepo.SVGIcons_batch.get(svgicon.ID) == undefined) {
                  this.frontRepo.SVGIcons.delete(svgicon.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TreePull performs a GET on Tree of the stack and redeem association pointers 
  TreePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.treeService.getTrees(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            trees,
          ]) => {
            // init the array
            this.frontRepo.Trees_array = trees

            // clear the map that counts Tree in the GET
            this.frontRepo.Trees_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            trees.forEach(
              tree => {
                this.frontRepo.Trees.set(tree.ID, tree)
                this.frontRepo.Trees_batch.set(tree.ID, tree)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear trees that are absent from the GET
            this.frontRepo.Trees.forEach(
              tree => {
                if (this.frontRepo.Trees_batch.get(tree.ID) == undefined) {
                  this.frontRepo.Trees.delete(tree.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

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

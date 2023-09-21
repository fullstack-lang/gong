import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { AstructService } from '../astruct.service'
import { getAstructUniqueID } from '../front-repo.service'
import { AstructBstruct2UseService } from '../astructbstruct2use.service'
import { getAstructBstruct2UseUniqueID } from '../front-repo.service'
import { AstructBstructUseService } from '../astructbstructuse.service'
import { getAstructBstructUseUniqueID } from '../front-repo.service'
import { BstructService } from '../bstruct.service'
import { getBstructUniqueID } from '../front-repo.service'
import { DstructService } from '../dstruct.service'
import { getDstructUniqueID } from '../front-repo.service'

import { RouteService } from '../route-service';

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-test-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private astructService: AstructService,
    private astructbstruct2useService: AstructBstruct2UseService,
    private astructbstructuseService: AstructBstructUseService,
    private bstructService: BstructService,
    private dstructService: DstructService,

    private routeService: RouteService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    console.log("Sidebar init: " + this.GONG__StackPath)

    // add the routes that will used by this side panel component and
    // by the component that are called from this component
    this.routeService.addDataPanelRoutes(this.GONG__StackPath)

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.astructService.AstructServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.astructbstruct2useService.AstructBstruct2UseServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.astructbstructuseService.AstructBstructUseServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.bstructService.BstructServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.dstructService.DstructServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Astruct part of the mat tree
      */
      let astructGongNodeStruct: GongNode = {
        name: "Astruct",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Astruct",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(astructGongNodeStruct)

      this.frontRepo.Astructs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Astructs_array.forEach(
        astructDB => {
          let astructGongNodeInstance: GongNode = {
            name: astructDB.Name,
            type: GongNodeType.INSTANCE,
            id: astructDB.ID,
            uniqueIdPerStack: getAstructUniqueID(astructDB.ID),
            structName: "Astruct",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          astructGongNodeStruct.children!.push(astructGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Associationtob
          */
          let AssociationtobGongNodeAssociation: GongNode = {
            name: "(Bstruct) Associationtob",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Associationtob",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(AssociationtobGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Associationtob
            */
          if (astructDB.Associationtob != undefined) {
            let astructGongNodeInstance_Associationtob: GongNode = {
              name: astructDB.Associationtob.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Associationtob.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getBstructUniqueID(astructDB.Associationtob.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AssociationtobGongNodeAssociation.children.push(astructGongNodeInstance_Associationtob)
          }

          /**
          * let append a node for the slide of pointer Anarrayofb
          */
          let AnarrayofbGongNodeAssociation: GongNode = {
            name: "(Bstruct) Anarrayofb",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Anarrayofb",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children.push(AnarrayofbGongNodeAssociation)

          astructDB.Anarrayofb?.forEach(bstructDB => {
            let bstructNode: GongNode = {
              name: bstructDB.Name,
              type: GongNodeType.INSTANCE,
              id: bstructDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAstructUniqueID(astructDB.ID)
                + 11 * getBstructUniqueID(bstructDB.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnarrayofbGongNodeAssociation.children.push(bstructNode)
          })

          /**
          * let append a node for the association Anotherassociationtob_2
          */
          let Anotherassociationtob_2GongNodeAssociation: GongNode = {
            name: "(Bstruct) Anotherassociationtob_2",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Anotherassociationtob_2",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(Anotherassociationtob_2GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Anotherassociationtob_2
            */
          if (astructDB.Anotherassociationtob_2 != undefined) {
            let astructGongNodeInstance_Anotherassociationtob_2: GongNode = {
              name: astructDB.Anotherassociationtob_2.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Anotherassociationtob_2.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getBstructUniqueID(astructDB.Anotherassociationtob_2.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Anotherassociationtob_2GongNodeAssociation.children.push(astructGongNodeInstance_Anotherassociationtob_2)
          }

          /**
          * let append a node for the association Bstruct
          */
          let BstructGongNodeAssociation: GongNode = {
            name: "(Bstruct) Bstruct",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Bstruct",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(BstructGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Bstruct
            */
          if (astructDB.Bstruct != undefined) {
            let astructGongNodeInstance_Bstruct: GongNode = {
              name: astructDB.Bstruct.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Bstruct.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getBstructUniqueID(astructDB.Bstruct.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            BstructGongNodeAssociation.children.push(astructGongNodeInstance_Bstruct)
          }

          /**
          * let append a node for the association Bstruct2
          */
          let Bstruct2GongNodeAssociation: GongNode = {
            name: "(Bstruct) Bstruct2",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Bstruct2",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(Bstruct2GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Bstruct2
            */
          if (astructDB.Bstruct2 != undefined) {
            let astructGongNodeInstance_Bstruct2: GongNode = {
              name: astructDB.Bstruct2.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Bstruct2.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getBstructUniqueID(astructDB.Bstruct2.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Bstruct2GongNodeAssociation.children.push(astructGongNodeInstance_Bstruct2)
          }

          /**
          * let append a node for the association Dstruct
          */
          let DstructGongNodeAssociation: GongNode = {
            name: "(Dstruct) Dstruct",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Dstruct",
            associatedStructName: "Dstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(DstructGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Dstruct
            */
          if (astructDB.Dstruct != undefined) {
            let astructGongNodeInstance_Dstruct: GongNode = {
              name: astructDB.Dstruct.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Dstruct.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getDstructUniqueID(astructDB.Dstruct.ID),
              structName: "Dstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            DstructGongNodeAssociation.children.push(astructGongNodeInstance_Dstruct)
          }

          /**
          * let append a node for the association Dstruct2
          */
          let Dstruct2GongNodeAssociation: GongNode = {
            name: "(Dstruct) Dstruct2",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Dstruct2",
            associatedStructName: "Dstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(Dstruct2GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Dstruct2
            */
          if (astructDB.Dstruct2 != undefined) {
            let astructGongNodeInstance_Dstruct2: GongNode = {
              name: astructDB.Dstruct2.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Dstruct2.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getDstructUniqueID(astructDB.Dstruct2.ID),
              structName: "Dstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Dstruct2GongNodeAssociation.children.push(astructGongNodeInstance_Dstruct2)
          }

          /**
          * let append a node for the association Dstruct3
          */
          let Dstruct3GongNodeAssociation: GongNode = {
            name: "(Dstruct) Dstruct3",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Dstruct3",
            associatedStructName: "Dstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(Dstruct3GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Dstruct3
            */
          if (astructDB.Dstruct3 != undefined) {
            let astructGongNodeInstance_Dstruct3: GongNode = {
              name: astructDB.Dstruct3.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Dstruct3.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getDstructUniqueID(astructDB.Dstruct3.ID),
              structName: "Dstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Dstruct3GongNodeAssociation.children.push(astructGongNodeInstance_Dstruct3)
          }

          /**
          * let append a node for the association Dstruct4
          */
          let Dstruct4GongNodeAssociation: GongNode = {
            name: "(Dstruct) Dstruct4",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Dstruct4",
            associatedStructName: "Dstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(Dstruct4GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Dstruct4
            */
          if (astructDB.Dstruct4 != undefined) {
            let astructGongNodeInstance_Dstruct4: GongNode = {
              name: astructDB.Dstruct4.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.Dstruct4.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getDstructUniqueID(astructDB.Dstruct4.ID),
              structName: "Dstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Dstruct4GongNodeAssociation.children.push(astructGongNodeInstance_Dstruct4)
          }

          /**
          * let append a node for the slide of pointer Anarrayofa
          */
          let AnarrayofaGongNodeAssociation: GongNode = {
            name: "(Astruct) Anarrayofa",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Anarrayofa",
            associatedStructName: "Astruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children.push(AnarrayofaGongNodeAssociation)

          astructDB.Anarrayofa?.forEach(astructDB => {
            let astructNode: GongNode = {
              name: astructDB.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAstructUniqueID(astructDB.ID)
                + 11 * getAstructUniqueID(astructDB.ID),
              structName: "Astruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnarrayofaGongNodeAssociation.children.push(astructNode)
          })

          /**
          * let append a node for the slide of pointer Anotherarrayofb
          */
          let AnotherarrayofbGongNodeAssociation: GongNode = {
            name: "(Bstruct) Anotherarrayofb",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Anotherarrayofb",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children.push(AnotherarrayofbGongNodeAssociation)

          astructDB.Anotherarrayofb?.forEach(bstructDB => {
            let bstructNode: GongNode = {
              name: bstructDB.Name,
              type: GongNodeType.INSTANCE,
              id: bstructDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAstructUniqueID(astructDB.ID)
                + 11 * getBstructUniqueID(bstructDB.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnotherarrayofbGongNodeAssociation.children.push(bstructNode)
          })

          /**
          * let append a node for the slide of pointer AnarrayofbUse
          */
          let AnarrayofbUseGongNodeAssociation: GongNode = {
            name: "(AstructBstructUse) AnarrayofbUse",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "AnarrayofbUse",
            associatedStructName: "AstructBstructUse",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children.push(AnarrayofbUseGongNodeAssociation)

          astructDB.AnarrayofbUse?.forEach(astructbstructuseDB => {
            let astructbstructuseNode: GongNode = {
              name: astructbstructuseDB.Name,
              type: GongNodeType.INSTANCE,
              id: astructbstructuseDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAstructUniqueID(astructDB.ID)
                + 11 * getAstructBstructUseUniqueID(astructbstructuseDB.ID),
              structName: "AstructBstructUse",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnarrayofbUseGongNodeAssociation.children.push(astructbstructuseNode)
          })

          /**
          * let append a node for the slide of pointer Anarrayofb2Use
          */
          let Anarrayofb2UseGongNodeAssociation: GongNode = {
            name: "(AstructBstruct2Use) Anarrayofb2Use",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "Anarrayofb2Use",
            associatedStructName: "AstructBstruct2Use",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children.push(Anarrayofb2UseGongNodeAssociation)

          astructDB.Anarrayofb2Use?.forEach(astructbstruct2useDB => {
            let astructbstruct2useNode: GongNode = {
              name: astructbstruct2useDB.Name,
              type: GongNodeType.INSTANCE,
              id: astructbstruct2useDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAstructUniqueID(astructDB.ID)
                + 11 * getAstructBstruct2UseUniqueID(astructbstruct2useDB.ID),
              structName: "AstructBstruct2Use",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Anarrayofb2UseGongNodeAssociation.children.push(astructbstruct2useNode)
          })

          /**
          * let append a node for the association AnAstruct
          */
          let AnAstructGongNodeAssociation: GongNode = {
            name: "(Astruct) AnAstruct",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Astruct",
            associationField: "AnAstruct",
            associatedStructName: "Astruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructGongNodeInstance.children!.push(AnAstructGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation AnAstruct
            */
          if (astructDB.AnAstruct != undefined) {
            let astructGongNodeInstance_AnAstruct: GongNode = {
              name: astructDB.AnAstruct.Name,
              type: GongNodeType.INSTANCE,
              id: astructDB.AnAstruct.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructUniqueID(astructDB.ID)
                + 5 * getAstructUniqueID(astructDB.AnAstruct.ID),
              structName: "Astruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnAstructGongNodeAssociation.children.push(astructGongNodeInstance_AnAstruct)
          }

        }
      )

      /**
      * fill up the AstructBstruct2Use part of the mat tree
      */
      let astructbstruct2useGongNodeStruct: GongNode = {
        name: "AstructBstruct2Use",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "AstructBstruct2Use",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(astructbstruct2useGongNodeStruct)

      this.frontRepo.AstructBstruct2Uses_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.AstructBstruct2Uses_array.forEach(
        astructbstruct2useDB => {
          let astructbstruct2useGongNodeInstance: GongNode = {
            name: astructbstruct2useDB.Name,
            type: GongNodeType.INSTANCE,
            id: astructbstruct2useDB.ID,
            uniqueIdPerStack: getAstructBstruct2UseUniqueID(astructbstruct2useDB.ID),
            structName: "AstructBstruct2Use",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          astructbstruct2useGongNodeStruct.children!.push(astructbstruct2useGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Bstrcut2
          */
          let Bstrcut2GongNodeAssociation: GongNode = {
            name: "(Bstruct) Bstrcut2",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructbstruct2useDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "AstructBstruct2Use",
            associationField: "Bstrcut2",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructbstruct2useGongNodeInstance.children!.push(Bstrcut2GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Bstrcut2
            */
          if (astructbstruct2useDB.Bstrcut2 != undefined) {
            let astructbstruct2useGongNodeInstance_Bstrcut2: GongNode = {
              name: astructbstruct2useDB.Bstrcut2.Name,
              type: GongNodeType.INSTANCE,
              id: astructbstruct2useDB.Bstrcut2.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructBstruct2UseUniqueID(astructbstruct2useDB.ID)
                + 5 * getBstructUniqueID(astructbstruct2useDB.Bstrcut2.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Bstrcut2GongNodeAssociation.children.push(astructbstruct2useGongNodeInstance_Bstrcut2)
          }

        }
      )

      /**
      * fill up the AstructBstructUse part of the mat tree
      */
      let astructbstructuseGongNodeStruct: GongNode = {
        name: "AstructBstructUse",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "AstructBstructUse",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(astructbstructuseGongNodeStruct)

      this.frontRepo.AstructBstructUses_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.AstructBstructUses_array.forEach(
        astructbstructuseDB => {
          let astructbstructuseGongNodeInstance: GongNode = {
            name: astructbstructuseDB.Name,
            type: GongNodeType.INSTANCE,
            id: astructbstructuseDB.ID,
            uniqueIdPerStack: getAstructBstructUseUniqueID(astructbstructuseDB.ID),
            structName: "AstructBstructUse",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          astructbstructuseGongNodeStruct.children!.push(astructbstructuseGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Bstruct2
          */
          let Bstruct2GongNodeAssociation: GongNode = {
            name: "(Bstruct) Bstruct2",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: astructbstructuseDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "AstructBstructUse",
            associationField: "Bstruct2",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          astructbstructuseGongNodeInstance.children!.push(Bstruct2GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Bstruct2
            */
          if (astructbstructuseDB.Bstruct2 != undefined) {
            let astructbstructuseGongNodeInstance_Bstruct2: GongNode = {
              name: astructbstructuseDB.Bstruct2.Name,
              type: GongNodeType.INSTANCE,
              id: astructbstructuseDB.Bstruct2.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAstructBstructUseUniqueID(astructbstructuseDB.ID)
                + 5 * getBstructUniqueID(astructbstructuseDB.Bstruct2.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Bstruct2GongNodeAssociation.children.push(astructbstructuseGongNodeInstance_Bstruct2)
          }

        }
      )

      /**
      * fill up the Bstruct part of the mat tree
      */
      let bstructGongNodeStruct: GongNode = {
        name: "Bstruct",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Bstruct",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(bstructGongNodeStruct)

      this.frontRepo.Bstructs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Bstructs_array.forEach(
        bstructDB => {
          let bstructGongNodeInstance: GongNode = {
            name: bstructDB.Name,
            type: GongNodeType.INSTANCE,
            id: bstructDB.ID,
            uniqueIdPerStack: getBstructUniqueID(bstructDB.ID),
            structName: "Bstruct",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          bstructGongNodeStruct.children!.push(bstructGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Dstruct part of the mat tree
      */
      let dstructGongNodeStruct: GongNode = {
        name: "Dstruct",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Dstruct",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(dstructGongNodeStruct)

      this.frontRepo.Dstructs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Dstructs_array.forEach(
        dstructDB => {
          let dstructGongNodeInstance: GongNode = {
            name: dstructDB.Name,
            type: GongNodeType.INSTANCE,
            id: dstructDB.ID,
            uniqueIdPerStack: getDstructUniqueID(dstructDB.ID),
            structName: "Dstruct",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          dstructGongNodeStruct.children!.push(dstructGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Anarrayofb
          */
          let AnarrayofbGongNodeAssociation: GongNode = {
            name: "(Bstruct) Anarrayofb",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: dstructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Dstruct",
            associationField: "Anarrayofb",
            associatedStructName: "Bstruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          dstructGongNodeInstance.children.push(AnarrayofbGongNodeAssociation)

          dstructDB.Anarrayofb?.forEach(bstructDB => {
            let bstructNode: GongNode = {
              name: bstructDB.Name,
              type: GongNodeType.INSTANCE,
              id: bstructDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getDstructUniqueID(dstructDB.ID)
                + 11 * getBstructUniqueID(bstructDB.ID),
              structName: "Bstruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnarrayofbGongNodeAssociation.children.push(bstructNode)
          })

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    })
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path
    this.router.navigate([{
      outlets: {
        outletName: [fullPath]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
      let outletConf: any = {}
      outletConf[outletName] = [fullPath, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }

    if (type == GongNodeType.INSTANCE) {
      let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + structName.toLowerCase() + "-detail"

      let outletConf: any = {}
      outletConf[outletName] = [fullPath, id, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf : any = {}
    outletConf[outletName] = [fullPath, this.GONG__StackPath]
    this.router.navigate([{ outlets: outletConf }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + node.associatedStructName.toLowerCase() + "-adder"
    this.router.navigate([{
      outlets: {
        outletName: [fullPath, node.id, node.structName, node.associationField, this.GONG__StackPath]
      }
    }]);
  }
}

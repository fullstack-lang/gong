import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { ClassdiagramService } from '../classdiagram.service'
import { getClassdiagramUniqueID } from '../front-repo.service'
import { ClassshapeService } from '../classshape.service'
import { getClassshapeUniqueID } from '../front-repo.service'
import { FieldService } from '../field.service'
import { getFieldUniqueID } from '../front-repo.service'
import { GongdocCommandService } from '../gongdoccommand.service'
import { getGongdocCommandUniqueID } from '../front-repo.service'
import { GongdocStatusService } from '../gongdocstatus.service'
import { getGongdocStatusUniqueID } from '../front-repo.service'
import { LinkService } from '../link.service'
import { getLinkUniqueID } from '../front-repo.service'
import { PkgeltService } from '../pkgelt.service'
import { getPkgeltUniqueID } from '../front-repo.service'
import { PositionService } from '../position.service'
import { getPositionUniqueID } from '../front-repo.service'
import { StateService } from '../state.service'
import { getStateUniqueID } from '../front-repo.service'
import { UmlscService } from '../umlsc.service'
import { getUmlscUniqueID } from '../front-repo.service'
import { VerticeService } from '../vertice.service'
import { getVerticeUniqueID } from '../front-repo.service'

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
  children?: GongNode[];
  type: GongNodeType;
  structName: string;
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
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongdoc-sidebar',
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
  frontRepo: FrontRepo
  commitNb: number

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,

    // insertion point for per struct service declaration
    private classdiagramService: ClassdiagramService,
    private classshapeService: ClassshapeService,
    private fieldService: FieldService,
    private gongdoccommandService: GongdocCommandService,
    private gongdocstatusService: GongdocStatusService,
    private linkService: LinkService,
    private pkgeltService: PkgeltService,
    private positionService: PositionService,
    private stateService: StateService,
    private umlscService: UmlscService,
    private verticeService: VerticeService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.classdiagramService.ClassdiagramServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.classshapeService.ClassshapeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.fieldService.FieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongdoccommandService.GongdocCommandServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongdocstatusService.GongdocStatusServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.linkService.LinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.pkgeltService.PkgeltServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.positionService.PositionServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.stateService.StateServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.umlscService.UmlscServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.verticeService.VerticeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (this.treeControl.isExpanded(node)) {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = true
            } else {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = false
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Classdiagram part of the mat tree
      */
      let classdiagramGongNodeStruct: GongNode = {
        name: "Classdiagram",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Classdiagram",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(classdiagramGongNodeStruct)

      this.frontRepo.Classdiagrams_array.forEach(
        classdiagramDB => {
          let classdiagramGongNodeInstance: GongNode = {
            name: classdiagramDB.Name,
            type: GongNodeType.INSTANCE,
            id: classdiagramDB.ID,
            uniqueIdPerStack: getClassdiagramUniqueID(classdiagramDB.ID),
            structName: "Classdiagram",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          classdiagramGongNodeStruct.children.push(classdiagramGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Classshapes
          */
          let ClassshapesGongNodeAssociation: GongNode = {
            name: "(Classshape) Classshapes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classdiagramDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classdiagram",
            associatedStructName: "Classshape",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classdiagramGongNodeInstance.children.push(ClassshapesGongNodeAssociation)

          classdiagramDB.Classshapes?.forEach(classshapeDB => {
            let classshapeNode: GongNode = {
              name: classshapeDB.Name,
              type: GongNodeType.INSTANCE,
              id: classshapeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getClassdiagramUniqueID(classdiagramDB.ID)
                + 11 * getClassshapeUniqueID(classshapeDB.ID),
              structName: "Classshape",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ClassshapesGongNodeAssociation.children.push(classshapeNode)
          })

        }
      )

      /**
      * fill up the Classshape part of the mat tree
      */
      let classshapeGongNodeStruct: GongNode = {
        name: "Classshape",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Classshape",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(classshapeGongNodeStruct)

      this.frontRepo.Classshapes_array.forEach(
        classshapeDB => {
          let classshapeGongNodeInstance: GongNode = {
            name: classshapeDB.Name,
            type: GongNodeType.INSTANCE,
            id: classshapeDB.ID,
            uniqueIdPerStack: getClassshapeUniqueID(classshapeDB.ID),
            structName: "Classshape",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          classshapeGongNodeStruct.children.push(classshapeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Position
          */
          let PositionGongNodeAssociation: GongNode = {
            name: "(Position) Position",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: classshapeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Classshape",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classshapeGongNodeInstance.children.push(PositionGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Position
            */
          if (classshapeDB.Position != undefined) {
            let classshapeGongNodeInstance_Position: GongNode = {
              name: classshapeDB.Position.Name,
              type: GongNodeType.INSTANCE,
              id: classshapeDB.Position.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getClassshapeUniqueID(classshapeDB.ID)
                + 5 * getPositionUniqueID(classshapeDB.Position.ID),
              structName: "Position",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PositionGongNodeAssociation.children.push(classshapeGongNodeInstance_Position)
          }

          /**
          * let append a node for the slide of pointer Fields
          */
          let FieldsGongNodeAssociation: GongNode = {
            name: "(Field) Fields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classshape",
            associatedStructName: "Field",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classshapeGongNodeInstance.children.push(FieldsGongNodeAssociation)

          classshapeDB.Fields?.forEach(fieldDB => {
            let fieldNode: GongNode = {
              name: fieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: fieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getClassshapeUniqueID(classshapeDB.ID)
                + 11 * getFieldUniqueID(fieldDB.ID),
              structName: "Field",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FieldsGongNodeAssociation.children.push(fieldNode)
          })

          /**
          * let append a node for the slide of pointer Links
          */
          let LinksGongNodeAssociation: GongNode = {
            name: "(Link) Links",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classshape",
            associatedStructName: "Link",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classshapeGongNodeInstance.children.push(LinksGongNodeAssociation)

          classshapeDB.Links?.forEach(linkDB => {
            let linkNode: GongNode = {
              name: linkDB.Name,
              type: GongNodeType.INSTANCE,
              id: linkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getClassshapeUniqueID(classshapeDB.ID)
                + 11 * getLinkUniqueID(linkDB.ID),
              structName: "Link",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinksGongNodeAssociation.children.push(linkNode)
          })

        }
      )

      /**
      * fill up the Field part of the mat tree
      */
      let fieldGongNodeStruct: GongNode = {
        name: "Field",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Field",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(fieldGongNodeStruct)

      this.frontRepo.Fields_array.forEach(
        fieldDB => {
          let fieldGongNodeInstance: GongNode = {
            name: fieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: fieldDB.ID,
            uniqueIdPerStack: getFieldUniqueID(fieldDB.ID),
            structName: "Field",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          fieldGongNodeStruct.children.push(fieldGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongdocCommand part of the mat tree
      */
      let gongdoccommandGongNodeStruct: GongNode = {
        name: "GongdocCommand",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongdocCommand",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongdoccommandGongNodeStruct)

      this.frontRepo.GongdocCommands_array.forEach(
        gongdoccommandDB => {
          let gongdoccommandGongNodeInstance: GongNode = {
            name: gongdoccommandDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongdoccommandDB.ID,
            uniqueIdPerStack: getGongdocCommandUniqueID(gongdoccommandDB.ID),
            structName: "GongdocCommand",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongdoccommandGongNodeStruct.children.push(gongdoccommandGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongdocStatus part of the mat tree
      */
      let gongdocstatusGongNodeStruct: GongNode = {
        name: "GongdocStatus",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongdocStatus",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongdocstatusGongNodeStruct)

      this.frontRepo.GongdocStatuss_array.forEach(
        gongdocstatusDB => {
          let gongdocstatusGongNodeInstance: GongNode = {
            name: gongdocstatusDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongdocstatusDB.ID,
            uniqueIdPerStack: getGongdocStatusUniqueID(gongdocstatusDB.ID),
            structName: "GongdocStatus",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongdocstatusGongNodeStruct.children.push(gongdocstatusGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Link part of the mat tree
      */
      let linkGongNodeStruct: GongNode = {
        name: "Link",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Link",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(linkGongNodeStruct)

      this.frontRepo.Links_array.forEach(
        linkDB => {
          let linkGongNodeInstance: GongNode = {
            name: linkDB.Name,
            type: GongNodeType.INSTANCE,
            id: linkDB.ID,
            uniqueIdPerStack: getLinkUniqueID(linkDB.ID),
            structName: "Link",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          linkGongNodeStruct.children.push(linkGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Middlevertice
          */
          let MiddleverticeGongNodeAssociation: GongNode = {
            name: "(Vertice) Middlevertice",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: linkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Link",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children.push(MiddleverticeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Middlevertice
            */
          if (linkDB.Middlevertice != undefined) {
            let linkGongNodeInstance_Middlevertice: GongNode = {
              name: linkDB.Middlevertice.Name,
              type: GongNodeType.INSTANCE,
              id: linkDB.Middlevertice.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getLinkUniqueID(linkDB.ID)
                + 5 * getVerticeUniqueID(linkDB.Middlevertice.ID),
              structName: "Vertice",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MiddleverticeGongNodeAssociation.children.push(linkGongNodeInstance_Middlevertice)
          }

        }
      )

      /**
      * fill up the Pkgelt part of the mat tree
      */
      let pkgeltGongNodeStruct: GongNode = {
        name: "Pkgelt",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Pkgelt",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(pkgeltGongNodeStruct)

      this.frontRepo.Pkgelts_array.forEach(
        pkgeltDB => {
          let pkgeltGongNodeInstance: GongNode = {
            name: pkgeltDB.Name,
            type: GongNodeType.INSTANCE,
            id: pkgeltDB.ID,
            uniqueIdPerStack: getPkgeltUniqueID(pkgeltDB.ID),
            structName: "Pkgelt",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          pkgeltGongNodeStruct.children.push(pkgeltGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Classdiagrams
          */
          let ClassdiagramsGongNodeAssociation: GongNode = {
            name: "(Classdiagram) Classdiagrams",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: pkgeltDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Pkgelt",
            associatedStructName: "Classdiagram",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pkgeltGongNodeInstance.children.push(ClassdiagramsGongNodeAssociation)

          pkgeltDB.Classdiagrams?.forEach(classdiagramDB => {
            let classdiagramNode: GongNode = {
              name: classdiagramDB.Name,
              type: GongNodeType.INSTANCE,
              id: classdiagramDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getPkgeltUniqueID(pkgeltDB.ID)
                + 11 * getClassdiagramUniqueID(classdiagramDB.ID),
              structName: "Classdiagram",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ClassdiagramsGongNodeAssociation.children.push(classdiagramNode)
          })

          /**
          * let append a node for the slide of pointer Umlscs
          */
          let UmlscsGongNodeAssociation: GongNode = {
            name: "(Umlsc) Umlscs",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: pkgeltDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Pkgelt",
            associatedStructName: "Umlsc",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pkgeltGongNodeInstance.children.push(UmlscsGongNodeAssociation)

          pkgeltDB.Umlscs?.forEach(umlscDB => {
            let umlscNode: GongNode = {
              name: umlscDB.Name,
              type: GongNodeType.INSTANCE,
              id: umlscDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getPkgeltUniqueID(pkgeltDB.ID)
                + 11 * getUmlscUniqueID(umlscDB.ID),
              structName: "Umlsc",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            UmlscsGongNodeAssociation.children.push(umlscNode)
          })

        }
      )

      /**
      * fill up the Position part of the mat tree
      */
      let positionGongNodeStruct: GongNode = {
        name: "Position",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Position",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(positionGongNodeStruct)

      this.frontRepo.Positions_array.forEach(
        positionDB => {
          let positionGongNodeInstance: GongNode = {
            name: positionDB.Name,
            type: GongNodeType.INSTANCE,
            id: positionDB.ID,
            uniqueIdPerStack: getPositionUniqueID(positionDB.ID),
            structName: "Position",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          positionGongNodeStruct.children.push(positionGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the State part of the mat tree
      */
      let stateGongNodeStruct: GongNode = {
        name: "State",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "State",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(stateGongNodeStruct)

      this.frontRepo.States_array.forEach(
        stateDB => {
          let stateGongNodeInstance: GongNode = {
            name: stateDB.Name,
            type: GongNodeType.INSTANCE,
            id: stateDB.ID,
            uniqueIdPerStack: getStateUniqueID(stateDB.ID),
            structName: "State",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          stateGongNodeStruct.children.push(stateGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Umlsc part of the mat tree
      */
      let umlscGongNodeStruct: GongNode = {
        name: "Umlsc",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Umlsc",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(umlscGongNodeStruct)

      this.frontRepo.Umlscs_array.forEach(
        umlscDB => {
          let umlscGongNodeInstance: GongNode = {
            name: umlscDB.Name,
            type: GongNodeType.INSTANCE,
            id: umlscDB.ID,
            uniqueIdPerStack: getUmlscUniqueID(umlscDB.ID),
            structName: "Umlsc",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          umlscGongNodeStruct.children.push(umlscGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer States
          */
          let StatesGongNodeAssociation: GongNode = {
            name: "(State) States",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: umlscDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Umlsc",
            associatedStructName: "State",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          umlscGongNodeInstance.children.push(StatesGongNodeAssociation)

          umlscDB.States?.forEach(stateDB => {
            let stateNode: GongNode = {
              name: stateDB.Name,
              type: GongNodeType.INSTANCE,
              id: stateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getUmlscUniqueID(umlscDB.ID)
                + 11 * getStateUniqueID(stateDB.ID),
              structName: "State",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            StatesGongNodeAssociation.children.push(stateNode)
          })

        }
      )

      /**
      * fill up the Vertice part of the mat tree
      */
      let verticeGongNodeStruct: GongNode = {
        name: "Vertice",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Vertice",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(verticeGongNodeStruct)

      this.frontRepo.Vertices_array.forEach(
        verticeDB => {
          let verticeGongNodeInstance: GongNode = {
            name: verticeDB.Name,
            type: GongNodeType.INSTANCE,
            id: verticeDB.ID,
            uniqueIdPerStack: getVerticeUniqueID(verticeDB.ID),
            structName: "Vertice",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          verticeGongNodeStruct.children.push(verticeGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes[node.uniqueIdPerStack] != undefined) {
              if (memoryOfExpandedNodes[node.uniqueIdPerStack]) {
                this.treeControl.expand(node)
              }
            }
          }
        )
      }
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        table: [path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          table: [path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          presentation: [structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        editor: [path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet( node: GongFlatNode) {
    console.log("setEditorSpecialRouterOutlet " + node)
    this.router.navigate([{
      outlets: {
        editor: [node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName + "_" + node.name]
      }
    }]);
  }
}

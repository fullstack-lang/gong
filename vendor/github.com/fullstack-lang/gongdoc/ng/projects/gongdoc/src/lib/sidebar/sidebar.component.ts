import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { ClassdiagramService } from '../classdiagram.service'
import { getClassdiagramUniqueID } from '../front-repo.service'
import { ClassshapeService } from '../classshape.service'
import { getClassshapeUniqueID } from '../front-repo.service'
import { DiagramPackageService } from '../diagrampackage.service'
import { getDiagramPackageUniqueID } from '../front-repo.service'
import { FieldService } from '../field.service'
import { getFieldUniqueID } from '../front-repo.service'
import { LinkService } from '../link.service'
import { getLinkUniqueID } from '../front-repo.service'
import { NodeService } from '../node.service'
import { getNodeUniqueID } from '../front-repo.service'
import { NoteLinkService } from '../notelink.service'
import { getNoteLinkUniqueID } from '../front-repo.service'
import { NoteShapeService } from '../noteshape.service'
import { getNoteShapeUniqueID } from '../front-repo.service'
import { PositionService } from '../position.service'
import { getPositionUniqueID } from '../front-repo.service'
import { ReferenceService } from '../reference.service'
import { getReferenceUniqueID } from '../front-repo.service'
import { TreeService } from '../tree.service'
import { getTreeUniqueID } from '../front-repo.service'
import { UmlStateService } from '../umlstate.service'
import { getUmlStateUniqueID } from '../front-repo.service'
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

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbFromBackService: CommitNbFromBackService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private classdiagramService: ClassdiagramService,
    private classshapeService: ClassshapeService,
    private diagrampackageService: DiagramPackageService,
    private fieldService: FieldService,
    private linkService: LinkService,
    private nodeService: NodeService,
    private notelinkService: NoteLinkService,
    private noteshapeService: NoteShapeService,
    private positionService: PositionService,
    private referenceService: ReferenceService,
    private treeService: TreeService,
    private umlstateService: UmlStateService,
    private umlscService: UmlscService,
    private verticeService: VerticeService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

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
    this.diagrampackageService.DiagramPackageServiceChanged.subscribe(
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
    this.linkService.LinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.nodeService.NodeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.notelinkService.NoteLinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.noteshapeService.NoteShapeServiceChanged.subscribe(
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
    this.referenceService.ReferenceServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.treeService.TreeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.umlstateService.UmlStateServiceChanged.subscribe(
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
      * fill up the Classdiagram part of the mat tree
      */
      let classdiagramGongNodeStruct: GongNode = {
        name: "Classdiagram",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Classdiagram",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(classdiagramGongNodeStruct)

      this.frontRepo.Classdiagrams_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Classdiagrams_array.forEach(
        classdiagramDB => {
          let classdiagramGongNodeInstance: GongNode = {
            name: classdiagramDB.Name,
            type: GongNodeType.INSTANCE,
            id: classdiagramDB.ID,
            uniqueIdPerStack: getClassdiagramUniqueID(classdiagramDB.ID),
            structName: "Classdiagram",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          classdiagramGongNodeStruct.children!.push(classdiagramGongNodeInstance)

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
            associationField: "Classshapes",
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
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ClassshapesGongNodeAssociation.children.push(classshapeNode)
          })

          /**
          * let append a node for the slide of pointer Notes
          */
          let NotesGongNodeAssociation: GongNode = {
            name: "(NoteShape) Notes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classdiagramDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classdiagram",
            associationField: "Notes",
            associatedStructName: "NoteShape",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classdiagramGongNodeInstance.children.push(NotesGongNodeAssociation)

          classdiagramDB.Notes?.forEach(noteshapeDB => {
            let noteshapeNode: GongNode = {
              name: noteshapeDB.Name,
              type: GongNodeType.INSTANCE,
              id: noteshapeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getClassdiagramUniqueID(classdiagramDB.ID)
                + 11 * getNoteShapeUniqueID(noteshapeDB.ID),
              structName: "NoteShape",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NotesGongNodeAssociation.children.push(noteshapeNode)
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
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(classshapeGongNodeStruct)

      this.frontRepo.Classshapes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Classshapes_array.forEach(
        classshapeDB => {
          let classshapeGongNodeInstance: GongNode = {
            name: classshapeDB.Name,
            type: GongNodeType.INSTANCE,
            id: classshapeDB.ID,
            uniqueIdPerStack: getClassshapeUniqueID(classshapeDB.ID),
            structName: "Classshape",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          classshapeGongNodeStruct.children!.push(classshapeGongNodeInstance)

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
            associationField: "Position",
            associatedStructName: "Position",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classshapeGongNodeInstance.children!.push(PositionGongNodeAssociation)

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
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PositionGongNodeAssociation.children.push(classshapeGongNodeInstance_Position)
          }

          /**
          * let append a node for the association Reference
          */
          let ReferenceGongNodeAssociation: GongNode = {
            name: "(Reference) Reference",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: classshapeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Classshape",
            associationField: "Reference",
            associatedStructName: "Reference",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classshapeGongNodeInstance.children!.push(ReferenceGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Reference
            */
          if (classshapeDB.Reference != undefined) {
            let classshapeGongNodeInstance_Reference: GongNode = {
              name: classshapeDB.Reference.Name,
              type: GongNodeType.INSTANCE,
              id: classshapeDB.Reference.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getClassshapeUniqueID(classshapeDB.ID)
                + 5 * getReferenceUniqueID(classshapeDB.Reference.ID),
              structName: "Reference",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ReferenceGongNodeAssociation.children.push(classshapeGongNodeInstance_Reference)
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
            associationField: "Fields",
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
              associationField: "",
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
            associationField: "Links",
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
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinksGongNodeAssociation.children.push(linkNode)
          })

        }
      )

      /**
      * fill up the DiagramPackage part of the mat tree
      */
      let diagrampackageGongNodeStruct: GongNode = {
        name: "DiagramPackage",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "DiagramPackage",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(diagrampackageGongNodeStruct)

      this.frontRepo.DiagramPackages_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.DiagramPackages_array.forEach(
        diagrampackageDB => {
          let diagrampackageGongNodeInstance: GongNode = {
            name: diagrampackageDB.Name,
            type: GongNodeType.INSTANCE,
            id: diagrampackageDB.ID,
            uniqueIdPerStack: getDiagramPackageUniqueID(diagrampackageDB.ID),
            structName: "DiagramPackage",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          diagrampackageGongNodeStruct.children!.push(diagrampackageGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Classdiagrams
          */
          let ClassdiagramsGongNodeAssociation: GongNode = {
            name: "(Classdiagram) Classdiagrams",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: diagrampackageDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "DiagramPackage",
            associationField: "Classdiagrams",
            associatedStructName: "Classdiagram",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          diagrampackageGongNodeInstance.children.push(ClassdiagramsGongNodeAssociation)

          diagrampackageDB.Classdiagrams?.forEach(classdiagramDB => {
            let classdiagramNode: GongNode = {
              name: classdiagramDB.Name,
              type: GongNodeType.INSTANCE,
              id: classdiagramDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getDiagramPackageUniqueID(diagrampackageDB.ID)
                + 11 * getClassdiagramUniqueID(classdiagramDB.ID),
              structName: "Classdiagram",
              associationField: "",
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
            id: diagrampackageDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "DiagramPackage",
            associationField: "Umlscs",
            associatedStructName: "Umlsc",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          diagrampackageGongNodeInstance.children.push(UmlscsGongNodeAssociation)

          diagrampackageDB.Umlscs?.forEach(umlscDB => {
            let umlscNode: GongNode = {
              name: umlscDB.Name,
              type: GongNodeType.INSTANCE,
              id: umlscDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getDiagramPackageUniqueID(diagrampackageDB.ID)
                + 11 * getUmlscUniqueID(umlscDB.ID),
              structName: "Umlsc",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            UmlscsGongNodeAssociation.children.push(umlscNode)
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
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(fieldGongNodeStruct)

      this.frontRepo.Fields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Fields_array.forEach(
        fieldDB => {
          let fieldGongNodeInstance: GongNode = {
            name: fieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: fieldDB.ID,
            uniqueIdPerStack: getFieldUniqueID(fieldDB.ID),
            structName: "Field",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          fieldGongNodeStruct.children!.push(fieldGongNodeInstance)

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
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(linkGongNodeStruct)

      this.frontRepo.Links_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Links_array.forEach(
        linkDB => {
          let linkGongNodeInstance: GongNode = {
            name: linkDB.Name,
            type: GongNodeType.INSTANCE,
            id: linkDB.ID,
            uniqueIdPerStack: getLinkUniqueID(linkDB.ID),
            structName: "Link",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          linkGongNodeStruct.children!.push(linkGongNodeInstance)

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
            associationField: "Middlevertice",
            associatedStructName: "Vertice",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children!.push(MiddleverticeGongNodeAssociation)

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
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MiddleverticeGongNodeAssociation.children.push(linkGongNodeInstance_Middlevertice)
          }

        }
      )

      /**
      * fill up the Node part of the mat tree
      */
      let nodeGongNodeStruct: GongNode = {
        name: "Node",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Node",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(nodeGongNodeStruct)

      this.frontRepo.Nodes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Nodes_array.forEach(
        nodeDB => {
          let nodeGongNodeInstance: GongNode = {
            name: nodeDB.Name,
            type: GongNodeType.INSTANCE,
            id: nodeDB.ID,
            uniqueIdPerStack: getNodeUniqueID(nodeDB.ID),
            structName: "Node",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nodeGongNodeStruct.children!.push(nodeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Classdiagram
          */
          let ClassdiagramGongNodeAssociation: GongNode = {
            name: "(Classdiagram) Classdiagram",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: nodeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Node",
            associationField: "Classdiagram",
            associatedStructName: "Classdiagram",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          nodeGongNodeInstance.children!.push(ClassdiagramGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Classdiagram
            */
          if (nodeDB.Classdiagram != undefined) {
            let nodeGongNodeInstance_Classdiagram: GongNode = {
              name: nodeDB.Classdiagram.Name,
              type: GongNodeType.INSTANCE,
              id: nodeDB.Classdiagram.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getNodeUniqueID(nodeDB.ID)
                + 5 * getClassdiagramUniqueID(nodeDB.Classdiagram.ID),
              structName: "Classdiagram",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ClassdiagramGongNodeAssociation.children.push(nodeGongNodeInstance_Classdiagram)
          }

          /**
          * let append a node for the slide of pointer Children
          */
          let ChildrenGongNodeAssociation: GongNode = {
            name: "(Node) Children",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: nodeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Node",
            associationField: "Children",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          nodeGongNodeInstance.children.push(ChildrenGongNodeAssociation)

          nodeDB.Children?.forEach(nodeDB => {
            let nodeNode: GongNode = {
              name: nodeDB.Name,
              type: GongNodeType.INSTANCE,
              id: nodeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getNodeUniqueID(nodeDB.ID)
                + 11 * getNodeUniqueID(nodeDB.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ChildrenGongNodeAssociation.children.push(nodeNode)
          })

        }
      )

      /**
      * fill up the NoteLink part of the mat tree
      */
      let notelinkGongNodeStruct: GongNode = {
        name: "NoteLink",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "NoteLink",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(notelinkGongNodeStruct)

      this.frontRepo.NoteLinks_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.NoteLinks_array.forEach(
        notelinkDB => {
          let notelinkGongNodeInstance: GongNode = {
            name: notelinkDB.Name,
            type: GongNodeType.INSTANCE,
            id: notelinkDB.ID,
            uniqueIdPerStack: getNoteLinkUniqueID(notelinkDB.ID),
            structName: "NoteLink",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          notelinkGongNodeStruct.children!.push(notelinkGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Classshape
          */
          let ClassshapeGongNodeAssociation: GongNode = {
            name: "(Classshape) Classshape",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: notelinkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "NoteLink",
            associationField: "Classshape",
            associatedStructName: "Classshape",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          notelinkGongNodeInstance.children!.push(ClassshapeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Classshape
            */
          if (notelinkDB.Classshape != undefined) {
            let notelinkGongNodeInstance_Classshape: GongNode = {
              name: notelinkDB.Classshape.Name,
              type: GongNodeType.INSTANCE,
              id: notelinkDB.Classshape.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getNoteLinkUniqueID(notelinkDB.ID)
                + 5 * getClassshapeUniqueID(notelinkDB.Classshape.ID),
              structName: "Classshape",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ClassshapeGongNodeAssociation.children.push(notelinkGongNodeInstance_Classshape)
          }

          /**
          * let append a node for the association Link
          */
          let LinkGongNodeAssociation: GongNode = {
            name: "(Link) Link",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: notelinkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "NoteLink",
            associationField: "Link",
            associatedStructName: "Link",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          notelinkGongNodeInstance.children!.push(LinkGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Link
            */
          if (notelinkDB.Link != undefined) {
            let notelinkGongNodeInstance_Link: GongNode = {
              name: notelinkDB.Link.Name,
              type: GongNodeType.INSTANCE,
              id: notelinkDB.Link.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getNoteLinkUniqueID(notelinkDB.ID)
                + 5 * getLinkUniqueID(notelinkDB.Link.ID),
              structName: "Link",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinkGongNodeAssociation.children.push(notelinkGongNodeInstance_Link)
          }

          /**
          * let append a node for the association Middlevertice
          */
          let MiddleverticeGongNodeAssociation: GongNode = {
            name: "(Vertice) Middlevertice",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: notelinkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "NoteLink",
            associationField: "Middlevertice",
            associatedStructName: "Vertice",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          notelinkGongNodeInstance.children!.push(MiddleverticeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Middlevertice
            */
          if (notelinkDB.Middlevertice != undefined) {
            let notelinkGongNodeInstance_Middlevertice: GongNode = {
              name: notelinkDB.Middlevertice.Name,
              type: GongNodeType.INSTANCE,
              id: notelinkDB.Middlevertice.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getNoteLinkUniqueID(notelinkDB.ID)
                + 5 * getVerticeUniqueID(notelinkDB.Middlevertice.ID),
              structName: "Vertice",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MiddleverticeGongNodeAssociation.children.push(notelinkGongNodeInstance_Middlevertice)
          }

        }
      )

      /**
      * fill up the NoteShape part of the mat tree
      */
      let noteshapeGongNodeStruct: GongNode = {
        name: "NoteShape",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "NoteShape",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(noteshapeGongNodeStruct)

      this.frontRepo.NoteShapes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.NoteShapes_array.forEach(
        noteshapeDB => {
          let noteshapeGongNodeInstance: GongNode = {
            name: noteshapeDB.Name,
            type: GongNodeType.INSTANCE,
            id: noteshapeDB.ID,
            uniqueIdPerStack: getNoteShapeUniqueID(noteshapeDB.ID),
            structName: "NoteShape",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          noteshapeGongNodeStruct.children!.push(noteshapeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer NoteLinks
          */
          let NoteLinksGongNodeAssociation: GongNode = {
            name: "(NoteLink) NoteLinks",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: noteshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "NoteShape",
            associationField: "NoteLinks",
            associatedStructName: "NoteLink",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          noteshapeGongNodeInstance.children.push(NoteLinksGongNodeAssociation)

          noteshapeDB.NoteLinks?.forEach(notelinkDB => {
            let notelinkNode: GongNode = {
              name: notelinkDB.Name,
              type: GongNodeType.INSTANCE,
              id: notelinkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getNoteShapeUniqueID(noteshapeDB.ID)
                + 11 * getNoteLinkUniqueID(notelinkDB.ID),
              structName: "NoteLink",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NoteLinksGongNodeAssociation.children.push(notelinkNode)
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
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(positionGongNodeStruct)

      this.frontRepo.Positions_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Positions_array.forEach(
        positionDB => {
          let positionGongNodeInstance: GongNode = {
            name: positionDB.Name,
            type: GongNodeType.INSTANCE,
            id: positionDB.ID,
            uniqueIdPerStack: getPositionUniqueID(positionDB.ID),
            structName: "Position",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          positionGongNodeStruct.children!.push(positionGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Reference part of the mat tree
      */
      let referenceGongNodeStruct: GongNode = {
        name: "Reference",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Reference",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(referenceGongNodeStruct)

      this.frontRepo.References_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.References_array.forEach(
        referenceDB => {
          let referenceGongNodeInstance: GongNode = {
            name: referenceDB.Name,
            type: GongNodeType.INSTANCE,
            id: referenceDB.ID,
            uniqueIdPerStack: getReferenceUniqueID(referenceDB.ID),
            structName: "Reference",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          referenceGongNodeStruct.children!.push(referenceGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Tree part of the mat tree
      */
      let treeGongNodeStruct: GongNode = {
        name: "Tree",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Tree",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(treeGongNodeStruct)

      this.frontRepo.Trees_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Trees_array.forEach(
        treeDB => {
          let treeGongNodeInstance: GongNode = {
            name: treeDB.Name,
            type: GongNodeType.INSTANCE,
            id: treeDB.ID,
            uniqueIdPerStack: getTreeUniqueID(treeDB.ID),
            structName: "Tree",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          treeGongNodeStruct.children!.push(treeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer RootNodes
          */
          let RootNodesGongNodeAssociation: GongNode = {
            name: "(Node) RootNodes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: treeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Tree",
            associationField: "RootNodes",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          treeGongNodeInstance.children.push(RootNodesGongNodeAssociation)

          treeDB.RootNodes?.forEach(nodeDB => {
            let nodeNode: GongNode = {
              name: nodeDB.Name,
              type: GongNodeType.INSTANCE,
              id: nodeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getTreeUniqueID(treeDB.ID)
                + 11 * getNodeUniqueID(nodeDB.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RootNodesGongNodeAssociation.children.push(nodeNode)
          })

        }
      )

      /**
      * fill up the UmlState part of the mat tree
      */
      let umlstateGongNodeStruct: GongNode = {
        name: "UmlState",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "UmlState",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(umlstateGongNodeStruct)

      this.frontRepo.UmlStates_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.UmlStates_array.forEach(
        umlstateDB => {
          let umlstateGongNodeInstance: GongNode = {
            name: umlstateDB.Name,
            type: GongNodeType.INSTANCE,
            id: umlstateDB.ID,
            uniqueIdPerStack: getUmlStateUniqueID(umlstateDB.ID),
            structName: "UmlState",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          umlstateGongNodeStruct.children!.push(umlstateGongNodeInstance)

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
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(umlscGongNodeStruct)

      this.frontRepo.Umlscs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Umlscs_array.forEach(
        umlscDB => {
          let umlscGongNodeInstance: GongNode = {
            name: umlscDB.Name,
            type: GongNodeType.INSTANCE,
            id: umlscDB.ID,
            uniqueIdPerStack: getUmlscUniqueID(umlscDB.ID),
            structName: "Umlsc",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          umlscGongNodeStruct.children!.push(umlscGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer States
          */
          let StatesGongNodeAssociation: GongNode = {
            name: "(UmlState) States",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: umlscDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Umlsc",
            associationField: "States",
            associatedStructName: "UmlState",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          umlscGongNodeInstance.children.push(StatesGongNodeAssociation)

          umlscDB.States?.forEach(umlstateDB => {
            let umlstateNode: GongNode = {
              name: umlstateDB.Name,
              type: GongNodeType.INSTANCE,
              id: umlstateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getUmlscUniqueID(umlscDB.ID)
                + 11 * getUmlStateUniqueID(umlstateDB.ID),
              structName: "UmlState",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            StatesGongNodeAssociation.children.push(umlstateNode)
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
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(verticeGongNodeStruct)

      this.frontRepo.Vertices_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Vertices_array.forEach(
        verticeDB => {
          let verticeGongNodeInstance: GongNode = {
            name: verticeDB.Name,
            type: GongNodeType.INSTANCE,
            id: verticeDB.ID,
            uniqueIdPerStack: getVerticeUniqueID(verticeDB.ID),
            structName: "Vertice",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          verticeGongNodeStruct.children!.push(verticeGongNodeInstance)

          // insertion point for per field code
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
    });

    // fetch the number of commits
    this.commitNbFromBackService.getCommitNbFromBack().subscribe(
      commitNbFromBack => {
        this.commitNbFromBack = commitNbFromBack
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
        github_com_fullstack_lang_gongdoc_go_table: ["github_com_fullstack_lang_gongdoc_go-" + path]
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
          github_com_fullstack_lang_gongdoc_go_table: ["github_com_fullstack_lang_gongdoc_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongdoc_go_presentation: ["github_com_fullstack_lang_gongdoc_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongdoc_go_editor: ["github_com_fullstack_lang_gongdoc_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}

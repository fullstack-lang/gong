import { Component, Input, OnInit } from '@angular/core';
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
import { DiagramPackageService } from '../diagrampackage.service'
import { getDiagramPackageUniqueID } from '../front-repo.service'
import { FieldService } from '../field.service'
import { getFieldUniqueID } from '../front-repo.service'
import { GongEnumShapeService } from '../gongenumshape.service'
import { getGongEnumShapeUniqueID } from '../front-repo.service'
import { GongEnumValueEntryService } from '../gongenumvalueentry.service'
import { getGongEnumValueEntryUniqueID } from '../front-repo.service'
import { GongStructShapeService } from '../gongstructshape.service'
import { getGongStructShapeUniqueID } from '../front-repo.service'
import { LinkService } from '../link.service'
import { getLinkUniqueID } from '../front-repo.service'
import { NoteShapeService } from '../noteshape.service'
import { getNoteShapeUniqueID } from '../front-repo.service'
import { NoteShapeLinkService } from '../noteshapelink.service'
import { getNoteShapeLinkUniqueID } from '../front-repo.service'
import { PositionService } from '../position.service'
import { getPositionUniqueID } from '../front-repo.service'
import { UmlStateService } from '../umlstate.service'
import { getUmlStateUniqueID } from '../front-repo.service'
import { UmlscService } from '../umlsc.service'
import { getUmlscUniqueID } from '../front-repo.service'
import { VerticeService } from '../vertice.service'
import { getVerticeUniqueID } from '../front-repo.service'

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

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private classdiagramService: ClassdiagramService,
    private diagrampackageService: DiagramPackageService,
    private fieldService: FieldService,
    private gongenumshapeService: GongEnumShapeService,
    private gongenumvalueentryService: GongEnumValueEntryService,
    private gongstructshapeService: GongStructShapeService,
    private linkService: LinkService,
    private noteshapeService: NoteShapeService,
    private noteshapelinkService: NoteShapeLinkService,
    private positionService: PositionService,
    private umlstateService: UmlStateService,
    private umlscService: UmlscService,
    private verticeService: VerticeService,

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
    this.classdiagramService.ClassdiagramServiceChanged.subscribe(
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
    this.gongenumshapeService.GongEnumShapeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongenumvalueentryService.GongEnumValueEntryServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongstructshapeService.GongStructShapeServiceChanged.subscribe(
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
    this.noteshapeService.NoteShapeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.noteshapelinkService.NoteShapeLinkServiceChanged.subscribe(
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
          * let append a node for the slide of pointer GongStructShapes
          */
          let GongStructShapesGongNodeAssociation: GongNode = {
            name: "(GongStructShape) GongStructShapes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classdiagramDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classdiagram",
            associationField: "GongStructShapes",
            associatedStructName: "GongStructShape",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classdiagramGongNodeInstance.children.push(GongStructShapesGongNodeAssociation)

          classdiagramDB.GongStructShapes?.forEach(gongstructshapeDB => {
            let gongstructshapeNode: GongNode = {
              name: gongstructshapeDB.Name,
              type: GongNodeType.INSTANCE,
              id: gongstructshapeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getClassdiagramUniqueID(classdiagramDB.ID)
                + 11 * getGongStructShapeUniqueID(gongstructshapeDB.ID),
              structName: "GongStructShape",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongStructShapesGongNodeAssociation.children.push(gongstructshapeNode)
          })

          /**
          * let append a node for the slide of pointer GongEnumShapes
          */
          let GongEnumShapesGongNodeAssociation: GongNode = {
            name: "(GongEnumShape) GongEnumShapes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classdiagramDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classdiagram",
            associationField: "GongEnumShapes",
            associatedStructName: "GongEnumShape",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classdiagramGongNodeInstance.children.push(GongEnumShapesGongNodeAssociation)

          classdiagramDB.GongEnumShapes?.forEach(gongenumshapeDB => {
            let gongenumshapeNode: GongNode = {
              name: gongenumshapeDB.Name,
              type: GongNodeType.INSTANCE,
              id: gongenumshapeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getClassdiagramUniqueID(classdiagramDB.ID)
                + 11 * getGongEnumShapeUniqueID(gongenumshapeDB.ID),
              structName: "GongEnumShape",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongEnumShapesGongNodeAssociation.children.push(gongenumshapeNode)
          })

          /**
          * let append a node for the slide of pointer NoteShapes
          */
          let NoteShapesGongNodeAssociation: GongNode = {
            name: "(NoteShape) NoteShapes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: classdiagramDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Classdiagram",
            associationField: "NoteShapes",
            associatedStructName: "NoteShape",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          classdiagramGongNodeInstance.children.push(NoteShapesGongNodeAssociation)

          classdiagramDB.NoteShapes?.forEach(noteshapeDB => {
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
            NoteShapesGongNodeAssociation.children.push(noteshapeNode)
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
          * let append a node for the association SelectedClassdiagram
          */
          let SelectedClassdiagramGongNodeAssociation: GongNode = {
            name: "(Classdiagram) SelectedClassdiagram",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: diagrampackageDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "DiagramPackage",
            associationField: "SelectedClassdiagram",
            associatedStructName: "Classdiagram",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          diagrampackageGongNodeInstance.children!.push(SelectedClassdiagramGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation SelectedClassdiagram
            */
          if (diagrampackageDB.SelectedClassdiagram != undefined) {
            let diagrampackageGongNodeInstance_SelectedClassdiagram: GongNode = {
              name: diagrampackageDB.SelectedClassdiagram.Name,
              type: GongNodeType.INSTANCE,
              id: diagrampackageDB.SelectedClassdiagram.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getDiagramPackageUniqueID(diagrampackageDB.ID)
                + 5 * getClassdiagramUniqueID(diagrampackageDB.SelectedClassdiagram.ID),
              structName: "Classdiagram",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            SelectedClassdiagramGongNodeAssociation.children.push(diagrampackageGongNodeInstance_SelectedClassdiagram)
          }

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
      * fill up the GongEnumShape part of the mat tree
      */
      let gongenumshapeGongNodeStruct: GongNode = {
        name: "GongEnumShape",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongEnumShape",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongenumshapeGongNodeStruct)

      this.frontRepo.GongEnumShapes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongEnumShapes_array.forEach(
        gongenumshapeDB => {
          let gongenumshapeGongNodeInstance: GongNode = {
            name: gongenumshapeDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongenumshapeDB.ID,
            uniqueIdPerStack: getGongEnumShapeUniqueID(gongenumshapeDB.ID),
            structName: "GongEnumShape",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongenumshapeGongNodeStruct.children!.push(gongenumshapeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Position
          */
          let PositionGongNodeAssociation: GongNode = {
            name: "(Position) Position",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: gongenumshapeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "GongEnumShape",
            associationField: "Position",
            associatedStructName: "Position",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongenumshapeGongNodeInstance.children!.push(PositionGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Position
            */
          if (gongenumshapeDB.Position != undefined) {
            let gongenumshapeGongNodeInstance_Position: GongNode = {
              name: gongenumshapeDB.Position.Name,
              type: GongNodeType.INSTANCE,
              id: gongenumshapeDB.Position.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getGongEnumShapeUniqueID(gongenumshapeDB.ID)
                + 5 * getPositionUniqueID(gongenumshapeDB.Position.ID),
              structName: "Position",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PositionGongNodeAssociation.children.push(gongenumshapeGongNodeInstance_Position)
          }

          /**
          * let append a node for the slide of pointer GongEnumValueEntrys
          */
          let GongEnumValueEntrysGongNodeAssociation: GongNode = {
            name: "(GongEnumValueEntry) GongEnumValueEntrys",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongenumshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongEnumShape",
            associationField: "GongEnumValueEntrys",
            associatedStructName: "GongEnumValueEntry",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongenumshapeGongNodeInstance.children.push(GongEnumValueEntrysGongNodeAssociation)

          gongenumshapeDB.GongEnumValueEntrys?.forEach(gongenumvalueentryDB => {
            let gongenumvalueentryNode: GongNode = {
              name: gongenumvalueentryDB.Name,
              type: GongNodeType.INSTANCE,
              id: gongenumvalueentryDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongEnumShapeUniqueID(gongenumshapeDB.ID)
                + 11 * getGongEnumValueEntryUniqueID(gongenumvalueentryDB.ID),
              structName: "GongEnumValueEntry",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongEnumValueEntrysGongNodeAssociation.children.push(gongenumvalueentryNode)
          })

        }
      )

      /**
      * fill up the GongEnumValueEntry part of the mat tree
      */
      let gongenumvalueentryGongNodeStruct: GongNode = {
        name: "GongEnumValueEntry",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongEnumValueEntry",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongenumvalueentryGongNodeStruct)

      this.frontRepo.GongEnumValueEntrys_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongEnumValueEntrys_array.forEach(
        gongenumvalueentryDB => {
          let gongenumvalueentryGongNodeInstance: GongNode = {
            name: gongenumvalueentryDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongenumvalueentryDB.ID,
            uniqueIdPerStack: getGongEnumValueEntryUniqueID(gongenumvalueentryDB.ID),
            structName: "GongEnumValueEntry",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongenumvalueentryGongNodeStruct.children!.push(gongenumvalueentryGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongStructShape part of the mat tree
      */
      let gongstructshapeGongNodeStruct: GongNode = {
        name: "GongStructShape",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongStructShape",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongstructshapeGongNodeStruct)

      this.frontRepo.GongStructShapes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongStructShapes_array.forEach(
        gongstructshapeDB => {
          let gongstructshapeGongNodeInstance: GongNode = {
            name: gongstructshapeDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongstructshapeDB.ID,
            uniqueIdPerStack: getGongStructShapeUniqueID(gongstructshapeDB.ID),
            structName: "GongStructShape",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongstructshapeGongNodeStruct.children!.push(gongstructshapeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Position
          */
          let PositionGongNodeAssociation: GongNode = {
            name: "(Position) Position",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: gongstructshapeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "GongStructShape",
            associationField: "Position",
            associatedStructName: "Position",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructshapeGongNodeInstance.children!.push(PositionGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Position
            */
          if (gongstructshapeDB.Position != undefined) {
            let gongstructshapeGongNodeInstance_Position: GongNode = {
              name: gongstructshapeDB.Position.Name,
              type: GongNodeType.INSTANCE,
              id: gongstructshapeDB.Position.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getGongStructShapeUniqueID(gongstructshapeDB.ID)
                + 5 * getPositionUniqueID(gongstructshapeDB.Position.ID),
              structName: "Position",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PositionGongNodeAssociation.children.push(gongstructshapeGongNodeInstance_Position)
          }

          /**
          * let append a node for the slide of pointer Fields
          */
          let FieldsGongNodeAssociation: GongNode = {
            name: "(Field) Fields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongstructshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStructShape",
            associationField: "Fields",
            associatedStructName: "Field",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructshapeGongNodeInstance.children.push(FieldsGongNodeAssociation)

          gongstructshapeDB.Fields?.forEach(fieldDB => {
            let fieldNode: GongNode = {
              name: fieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: fieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongStructShapeUniqueID(gongstructshapeDB.ID)
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
            id: gongstructshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStructShape",
            associationField: "Links",
            associatedStructName: "Link",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructshapeGongNodeInstance.children.push(LinksGongNodeAssociation)

          gongstructshapeDB.Links?.forEach(linkDB => {
            let linkNode: GongNode = {
              name: linkDB.Name,
              type: GongNodeType.INSTANCE,
              id: linkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongStructShapeUniqueID(gongstructshapeDB.ID)
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
          * let append a node for the slide of pointer NoteShapeLinks
          */
          let NoteShapeLinksGongNodeAssociation: GongNode = {
            name: "(NoteShapeLink) NoteShapeLinks",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: noteshapeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "NoteShape",
            associationField: "NoteShapeLinks",
            associatedStructName: "NoteShapeLink",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          noteshapeGongNodeInstance.children.push(NoteShapeLinksGongNodeAssociation)

          noteshapeDB.NoteShapeLinks?.forEach(noteshapelinkDB => {
            let noteshapelinkNode: GongNode = {
              name: noteshapelinkDB.Name,
              type: GongNodeType.INSTANCE,
              id: noteshapelinkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getNoteShapeUniqueID(noteshapeDB.ID)
                + 11 * getNoteShapeLinkUniqueID(noteshapelinkDB.ID),
              structName: "NoteShapeLink",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NoteShapeLinksGongNodeAssociation.children.push(noteshapelinkNode)
          })

        }
      )

      /**
      * fill up the NoteShapeLink part of the mat tree
      */
      let noteshapelinkGongNodeStruct: GongNode = {
        name: "NoteShapeLink",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "NoteShapeLink",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(noteshapelinkGongNodeStruct)

      this.frontRepo.NoteShapeLinks_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.NoteShapeLinks_array.forEach(
        noteshapelinkDB => {
          let noteshapelinkGongNodeInstance: GongNode = {
            name: noteshapelinkDB.Name,
            type: GongNodeType.INSTANCE,
            id: noteshapelinkDB.ID,
            uniqueIdPerStack: getNoteShapeLinkUniqueID(noteshapelinkDB.ID),
            structName: "NoteShapeLink",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          noteshapelinkGongNodeStruct.children!.push(noteshapelinkGongNodeInstance)

          // insertion point for per field code
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

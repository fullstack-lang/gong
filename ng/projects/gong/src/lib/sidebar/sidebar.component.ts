import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { GongBasicFieldService } from '../gongbasicfield.service'
import { getGongBasicFieldUniqueID } from '../front-repo.service'
import { GongEnumService } from '../gongenum.service'
import { getGongEnumUniqueID } from '../front-repo.service'
import { GongEnumValueService } from '../gongenumvalue.service'
import { getGongEnumValueUniqueID } from '../front-repo.service'
import { GongLinkService } from '../gonglink.service'
import { getGongLinkUniqueID } from '../front-repo.service'
import { GongNoteService } from '../gongnote.service'
import { getGongNoteUniqueID } from '../front-repo.service'
import { GongStructService } from '../gongstruct.service'
import { getGongStructUniqueID } from '../front-repo.service'
import { GongTimeFieldService } from '../gongtimefield.service'
import { getGongTimeFieldUniqueID } from '../front-repo.service'
import { MetaService } from '../meta.service'
import { getMetaUniqueID } from '../front-repo.service'
import { MetaReferenceService } from '../metareference.service'
import { getMetaReferenceUniqueID } from '../front-repo.service'
import { ModelPkgService } from '../modelpkg.service'
import { getModelPkgUniqueID } from '../front-repo.service'
import { PointerToGongStructFieldService } from '../pointertogongstructfield.service'
import { getPointerToGongStructFieldUniqueID } from '../front-repo.service'
import { SliceOfPointerToGongStructFieldService } from '../sliceofpointertogongstructfield.service'
import { getSliceOfPointerToGongStructFieldUniqueID } from '../front-repo.service'

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
  selector: 'app-gong-sidebar',
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
    private gongbasicfieldService: GongBasicFieldService,
    private gongenumService: GongEnumService,
    private gongenumvalueService: GongEnumValueService,
    private gonglinkService: GongLinkService,
    private gongnoteService: GongNoteService,
    private gongstructService: GongStructService,
    private gongtimefieldService: GongTimeFieldService,
    private metaService: MetaService,
    private metareferenceService: MetaReferenceService,
    private modelpkgService: ModelPkgService,
    private pointertogongstructfieldService: PointerToGongStructFieldService,
    private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,

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
    this.gongbasicfieldService.GongBasicFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongenumService.GongEnumServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongenumvalueService.GongEnumValueServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gonglinkService.GongLinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongnoteService.GongNoteServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongstructService.GongStructServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.gongtimefieldService.GongTimeFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.metaService.MetaServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.metareferenceService.MetaReferenceServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.modelpkgService.ModelPkgServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.pointertogongstructfieldService.PointerToGongStructFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.sliceofpointertogongstructfieldService.SliceOfPointerToGongStructFieldServiceChanged.subscribe(
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
      * fill up the GongBasicField part of the mat tree
      */
      let gongbasicfieldGongNodeStruct: GongNode = {
        name: "GongBasicField",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongBasicField",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongbasicfieldGongNodeStruct)

      this.frontRepo.GongBasicFields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongBasicFields_array.forEach(
        gongbasicfieldDB => {
          let gongbasicfieldGongNodeInstance: GongNode = {
            name: gongbasicfieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongbasicfieldDB.ID,
            uniqueIdPerStack: getGongBasicFieldUniqueID(gongbasicfieldDB.ID),
            structName: "GongBasicField",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongbasicfieldGongNodeStruct.children!.push(gongbasicfieldGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association GongEnum
          */
          let GongEnumGongNodeAssociation: GongNode = {
            name: "(GongEnum) GongEnum",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: gongbasicfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "GongBasicField",
            associationField: "GongEnum",
            associatedStructName: "GongEnum",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongbasicfieldGongNodeInstance.children!.push(GongEnumGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation GongEnum
            */
          if (gongbasicfieldDB.GongEnum != undefined) {
            let gongbasicfieldGongNodeInstance_GongEnum: GongNode = {
              name: gongbasicfieldDB.GongEnum.Name,
              type: GongNodeType.INSTANCE,
              id: gongbasicfieldDB.GongEnum.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getGongBasicFieldUniqueID(gongbasicfieldDB.ID)
                + 5 * getGongEnumUniqueID(gongbasicfieldDB.GongEnum.ID),
              structName: "GongEnum",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongEnumGongNodeAssociation.children.push(gongbasicfieldGongNodeInstance_GongEnum)
          }

        }
      )

      /**
      * fill up the GongEnum part of the mat tree
      */
      let gongenumGongNodeStruct: GongNode = {
        name: "GongEnum",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongEnum",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongenumGongNodeStruct)

      this.frontRepo.GongEnums_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongEnums_array.forEach(
        gongenumDB => {
          let gongenumGongNodeInstance: GongNode = {
            name: gongenumDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongenumDB.ID,
            uniqueIdPerStack: getGongEnumUniqueID(gongenumDB.ID),
            structName: "GongEnum",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongenumGongNodeStruct.children!.push(gongenumGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer GongEnumValues
          */
          let GongEnumValuesGongNodeAssociation: GongNode = {
            name: "(GongEnumValue) GongEnumValues",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongenumDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongEnum",
            associationField: "GongEnumValues",
            associatedStructName: "GongEnumValue",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongenumGongNodeInstance.children.push(GongEnumValuesGongNodeAssociation)

          gongenumDB.GongEnumValues?.forEach(gongenumvalueDB => {
            let gongenumvalueNode: GongNode = {
              name: gongenumvalueDB.Name,
              type: GongNodeType.INSTANCE,
              id: gongenumvalueDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongEnumUniqueID(gongenumDB.ID)
                + 11 * getGongEnumValueUniqueID(gongenumvalueDB.ID),
              structName: "GongEnumValue",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongEnumValuesGongNodeAssociation.children.push(gongenumvalueNode)
          })

        }
      )

      /**
      * fill up the GongEnumValue part of the mat tree
      */
      let gongenumvalueGongNodeStruct: GongNode = {
        name: "GongEnumValue",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongEnumValue",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongenumvalueGongNodeStruct)

      this.frontRepo.GongEnumValues_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongEnumValues_array.forEach(
        gongenumvalueDB => {
          let gongenumvalueGongNodeInstance: GongNode = {
            name: gongenumvalueDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongenumvalueDB.ID,
            uniqueIdPerStack: getGongEnumValueUniqueID(gongenumvalueDB.ID),
            structName: "GongEnumValue",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongenumvalueGongNodeStruct.children!.push(gongenumvalueGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongLink part of the mat tree
      */
      let gonglinkGongNodeStruct: GongNode = {
        name: "GongLink",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongLink",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gonglinkGongNodeStruct)

      this.frontRepo.GongLinks_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongLinks_array.forEach(
        gonglinkDB => {
          let gonglinkGongNodeInstance: GongNode = {
            name: gonglinkDB.Name,
            type: GongNodeType.INSTANCE,
            id: gonglinkDB.ID,
            uniqueIdPerStack: getGongLinkUniqueID(gonglinkDB.ID),
            structName: "GongLink",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gonglinkGongNodeStruct.children!.push(gonglinkGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the GongNote part of the mat tree
      */
      let gongnoteGongNodeStruct: GongNode = {
        name: "GongNote",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongNote",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongnoteGongNodeStruct)

      this.frontRepo.GongNotes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongNotes_array.forEach(
        gongnoteDB => {
          let gongnoteGongNodeInstance: GongNode = {
            name: gongnoteDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongnoteDB.ID,
            uniqueIdPerStack: getGongNoteUniqueID(gongnoteDB.ID),
            structName: "GongNote",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongnoteGongNodeStruct.children!.push(gongnoteGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Links
          */
          let LinksGongNodeAssociation: GongNode = {
            name: "(GongLink) Links",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongnoteDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongNote",
            associationField: "Links",
            associatedStructName: "GongLink",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongnoteGongNodeInstance.children.push(LinksGongNodeAssociation)

          gongnoteDB.Links?.forEach(gonglinkDB => {
            let gonglinkNode: GongNode = {
              name: gonglinkDB.Name,
              type: GongNodeType.INSTANCE,
              id: gonglinkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongNoteUniqueID(gongnoteDB.ID)
                + 11 * getGongLinkUniqueID(gonglinkDB.ID),
              structName: "GongLink",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinksGongNodeAssociation.children.push(gonglinkNode)
          })

        }
      )

      /**
      * fill up the GongStruct part of the mat tree
      */
      let gongstructGongNodeStruct: GongNode = {
        name: "GongStruct",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongStruct",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongstructGongNodeStruct)

      this.frontRepo.GongStructs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongStructs_array.forEach(
        gongstructDB => {
          let gongstructGongNodeInstance: GongNode = {
            name: gongstructDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongstructDB.ID,
            uniqueIdPerStack: getGongStructUniqueID(gongstructDB.ID),
            structName: "GongStruct",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongstructGongNodeStruct.children!.push(gongstructGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer GongBasicFields
          */
          let GongBasicFieldsGongNodeAssociation: GongNode = {
            name: "(GongBasicField) GongBasicFields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongstructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStruct",
            associationField: "GongBasicFields",
            associatedStructName: "GongBasicField",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(GongBasicFieldsGongNodeAssociation)

          gongstructDB.GongBasicFields?.forEach(gongbasicfieldDB => {
            let gongbasicfieldNode: GongNode = {
              name: gongbasicfieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: gongbasicfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongStructUniqueID(gongstructDB.ID)
                + 11 * getGongBasicFieldUniqueID(gongbasicfieldDB.ID),
              structName: "GongBasicField",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongBasicFieldsGongNodeAssociation.children.push(gongbasicfieldNode)
          })

          /**
          * let append a node for the slide of pointer GongTimeFields
          */
          let GongTimeFieldsGongNodeAssociation: GongNode = {
            name: "(GongTimeField) GongTimeFields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongstructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStruct",
            associationField: "GongTimeFields",
            associatedStructName: "GongTimeField",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(GongTimeFieldsGongNodeAssociation)

          gongstructDB.GongTimeFields?.forEach(gongtimefieldDB => {
            let gongtimefieldNode: GongNode = {
              name: gongtimefieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: gongtimefieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongStructUniqueID(gongstructDB.ID)
                + 11 * getGongTimeFieldUniqueID(gongtimefieldDB.ID),
              structName: "GongTimeField",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongTimeFieldsGongNodeAssociation.children.push(gongtimefieldNode)
          })

          /**
          * let append a node for the slide of pointer PointerToGongStructFields
          */
          let PointerToGongStructFieldsGongNodeAssociation: GongNode = {
            name: "(PointerToGongStructField) PointerToGongStructFields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongstructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStruct",
            associationField: "PointerToGongStructFields",
            associatedStructName: "PointerToGongStructField",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(PointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.PointerToGongStructFields?.forEach(pointertogongstructfieldDB => {
            let pointertogongstructfieldNode: GongNode = {
              name: pointertogongstructfieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: pointertogongstructfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongStructUniqueID(gongstructDB.ID)
                + 11 * getPointerToGongStructFieldUniqueID(pointertogongstructfieldDB.ID),
              structName: "PointerToGongStructField",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PointerToGongStructFieldsGongNodeAssociation.children.push(pointertogongstructfieldNode)
          })

          /**
          * let append a node for the slide of pointer SliceOfPointerToGongStructFields
          */
          let SliceOfPointerToGongStructFieldsGongNodeAssociation: GongNode = {
            name: "(SliceOfPointerToGongStructField) SliceOfPointerToGongStructFields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: gongstructDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStruct",
            associationField: "SliceOfPointerToGongStructFields",
            associatedStructName: "SliceOfPointerToGongStructField",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(SliceOfPointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.SliceOfPointerToGongStructFields?.forEach(sliceofpointertogongstructfieldDB => {
            let sliceofpointertogongstructfieldNode: GongNode = {
              name: sliceofpointertogongstructfieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: sliceofpointertogongstructfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGongStructUniqueID(gongstructDB.ID)
                + 11 * getSliceOfPointerToGongStructFieldUniqueID(sliceofpointertogongstructfieldDB.ID),
              structName: "SliceOfPointerToGongStructField",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            SliceOfPointerToGongStructFieldsGongNodeAssociation.children.push(sliceofpointertogongstructfieldNode)
          })

        }
      )

      /**
      * fill up the GongTimeField part of the mat tree
      */
      let gongtimefieldGongNodeStruct: GongNode = {
        name: "GongTimeField",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongTimeField",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongtimefieldGongNodeStruct)

      this.frontRepo.GongTimeFields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.GongTimeFields_array.forEach(
        gongtimefieldDB => {
          let gongtimefieldGongNodeInstance: GongNode = {
            name: gongtimefieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: gongtimefieldDB.ID,
            uniqueIdPerStack: getGongTimeFieldUniqueID(gongtimefieldDB.ID),
            structName: "GongTimeField",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          gongtimefieldGongNodeStruct.children!.push(gongtimefieldGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Meta part of the mat tree
      */
      let metaGongNodeStruct: GongNode = {
        name: "Meta",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Meta",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(metaGongNodeStruct)

      this.frontRepo.Metas_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Metas_array.forEach(
        metaDB => {
          let metaGongNodeInstance: GongNode = {
            name: metaDB.Name,
            type: GongNodeType.INSTANCE,
            id: metaDB.ID,
            uniqueIdPerStack: getMetaUniqueID(metaDB.ID),
            structName: "Meta",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          metaGongNodeStruct.children!.push(metaGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer MetaReferences
          */
          let MetaReferencesGongNodeAssociation: GongNode = {
            name: "(MetaReference) MetaReferences",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: metaDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Meta",
            associationField: "MetaReferences",
            associatedStructName: "MetaReference",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          metaGongNodeInstance.children.push(MetaReferencesGongNodeAssociation)

          metaDB.MetaReferences?.forEach(metareferenceDB => {
            let metareferenceNode: GongNode = {
              name: metareferenceDB.Name,
              type: GongNodeType.INSTANCE,
              id: metareferenceDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getMetaUniqueID(metaDB.ID)
                + 11 * getMetaReferenceUniqueID(metareferenceDB.ID),
              structName: "MetaReference",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MetaReferencesGongNodeAssociation.children.push(metareferenceNode)
          })

        }
      )

      /**
      * fill up the MetaReference part of the mat tree
      */
      let metareferenceGongNodeStruct: GongNode = {
        name: "MetaReference",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "MetaReference",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(metareferenceGongNodeStruct)

      this.frontRepo.MetaReferences_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.MetaReferences_array.forEach(
        metareferenceDB => {
          let metareferenceGongNodeInstance: GongNode = {
            name: metareferenceDB.Name,
            type: GongNodeType.INSTANCE,
            id: metareferenceDB.ID,
            uniqueIdPerStack: getMetaReferenceUniqueID(metareferenceDB.ID),
            structName: "MetaReference",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          metareferenceGongNodeStruct.children!.push(metareferenceGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the ModelPkg part of the mat tree
      */
      let modelpkgGongNodeStruct: GongNode = {
        name: "ModelPkg",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "ModelPkg",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(modelpkgGongNodeStruct)

      this.frontRepo.ModelPkgs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.ModelPkgs_array.forEach(
        modelpkgDB => {
          let modelpkgGongNodeInstance: GongNode = {
            name: modelpkgDB.Name,
            type: GongNodeType.INSTANCE,
            id: modelpkgDB.ID,
            uniqueIdPerStack: getModelPkgUniqueID(modelpkgDB.ID),
            structName: "ModelPkg",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          modelpkgGongNodeStruct.children!.push(modelpkgGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the PointerToGongStructField part of the mat tree
      */
      let pointertogongstructfieldGongNodeStruct: GongNode = {
        name: "PointerToGongStructField",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "PointerToGongStructField",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(pointertogongstructfieldGongNodeStruct)

      this.frontRepo.PointerToGongStructFields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.PointerToGongStructFields_array.forEach(
        pointertogongstructfieldDB => {
          let pointertogongstructfieldGongNodeInstance: GongNode = {
            name: pointertogongstructfieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: pointertogongstructfieldDB.ID,
            uniqueIdPerStack: getPointerToGongStructFieldUniqueID(pointertogongstructfieldDB.ID),
            structName: "PointerToGongStructField",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          pointertogongstructfieldGongNodeStruct.children!.push(pointertogongstructfieldGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association GongStruct
          */
          let GongStructGongNodeAssociation: GongNode = {
            name: "(GongStruct) GongStruct",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: pointertogongstructfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "PointerToGongStructField",
            associationField: "GongStruct",
            associatedStructName: "GongStruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pointertogongstructfieldGongNodeInstance.children!.push(GongStructGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation GongStruct
            */
          if (pointertogongstructfieldDB.GongStruct != undefined) {
            let pointertogongstructfieldGongNodeInstance_GongStruct: GongNode = {
              name: pointertogongstructfieldDB.GongStruct.Name,
              type: GongNodeType.INSTANCE,
              id: pointertogongstructfieldDB.GongStruct.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getPointerToGongStructFieldUniqueID(pointertogongstructfieldDB.ID)
                + 5 * getGongStructUniqueID(pointertogongstructfieldDB.GongStruct.ID),
              structName: "GongStruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongStructGongNodeAssociation.children.push(pointertogongstructfieldGongNodeInstance_GongStruct)
          }

        }
      )

      /**
      * fill up the SliceOfPointerToGongStructField part of the mat tree
      */
      let sliceofpointertogongstructfieldGongNodeStruct: GongNode = {
        name: "SliceOfPointerToGongStructField",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "SliceOfPointerToGongStructField",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(sliceofpointertogongstructfieldGongNodeStruct)

      this.frontRepo.SliceOfPointerToGongStructFields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.SliceOfPointerToGongStructFields_array.forEach(
        sliceofpointertogongstructfieldDB => {
          let sliceofpointertogongstructfieldGongNodeInstance: GongNode = {
            name: sliceofpointertogongstructfieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: sliceofpointertogongstructfieldDB.ID,
            uniqueIdPerStack: getSliceOfPointerToGongStructFieldUniqueID(sliceofpointertogongstructfieldDB.ID),
            structName: "SliceOfPointerToGongStructField",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          sliceofpointertogongstructfieldGongNodeStruct.children!.push(sliceofpointertogongstructfieldGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association GongStruct
          */
          let GongStructGongNodeAssociation: GongNode = {
            name: "(GongStruct) GongStruct",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: sliceofpointertogongstructfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "SliceOfPointerToGongStructField",
            associationField: "GongStruct",
            associatedStructName: "GongStruct",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          sliceofpointertogongstructfieldGongNodeInstance.children!.push(GongStructGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation GongStruct
            */
          if (sliceofpointertogongstructfieldDB.GongStruct != undefined) {
            let sliceofpointertogongstructfieldGongNodeInstance_GongStruct: GongNode = {
              name: sliceofpointertogongstructfieldDB.GongStruct.Name,
              type: GongNodeType.INSTANCE,
              id: sliceofpointertogongstructfieldDB.GongStruct.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getSliceOfPointerToGongStructFieldUniqueID(sliceofpointertogongstructfieldDB.ID)
                + 5 * getGongStructUniqueID(sliceofpointertogongstructfieldDB.GongStruct.ID),
              structName: "GongStruct",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GongStructGongNodeAssociation.children.push(sliceofpointertogongstructfieldGongNodeInstance_GongStruct)
          }

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

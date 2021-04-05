import { Component, OnInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, RouterState, ParamMap } from '@angular/router';
import { CdkDragDrop } from '@angular/cdk/drag-drop';
import { ElementRef, Injectable } from '@angular/core';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

// insertion point for per struct import code
import * as gong from 'gong'
import * as gongdoc from 'gongdoc'

import { ClassdiagramContextSubject, ClassdiagramContext } from '../diagram-displayed-gongstruct'
import { combineLatest, Observable, timer } from 'rxjs';
import { GongBasicFieldDB } from 'gong';
import { ClassshapeDB, FieldDB } from 'gongdoc';
import { NoDataRowOutlet } from '@angular/cdk/table';

/**
 * GongNode is the "data" node that is construed from the gongFrontRepo
 */
interface GongNode {
  name: string // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[]
  type: gongdoc.GongdocNodeType
  structName: string
  id: number
  uniqueIdPerStack: number
  instancePresentInDiagram?: boolean // for "instance" type node, in order to guide the display
  gongBasicField?: GongBasicFieldDB
  draggable: boolean
}

/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
export interface GongFlatNode {
  expandable: boolean
  name: string
  level: number
  type: gongdoc.GongdocNodeType
  structName: string
  id: number
  uniqueIdPerStack: number
  present?: boolean
  gongBasicField?: GongBasicFieldDB
  draggable: boolean
}

export interface DragDropPosition {
  x: number;
  y: number;
}

@Component({
  styleUrls: ['./sidebar-gong-diagrams.css'],
  selector: 'lib-sidebar-gong-diagams',
  templateUrl: './sidebar-gong-diagrams.html'
})
export class SidebarGongDiagramsComponent implements OnInit {

  @ViewChild('innerHTMLelement') innerHTMLelement: ElementRef;

  // the current classdiagram
  currentClassdiagram: gongdoc.ClassdiagramDB

  /**
   * initial node expansion
   * 
   * by default, all nodes are collapsed. This timer expands root nodes
   */
  expandRootNodesTimer: Observable<number> = timer(1000)

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
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
      present: node.instancePresentInDiagram,
      gongBasicField: node.gongBasicField,
      draggable: node.draggable
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
    node => {
      return node.expandable
    }
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
  gongFrontRepo: gong.FrontRepo
  gongdocFrontRepo: gongdoc.FrontRepo

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // provide the current display context
  classdiagramContext: ClassdiagramContext

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private gongFrontRepoService: gong.FrontRepoService,
    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private gongdocCommandService: gongdoc.GongdocCommandService,
  ) { }

  ngOnInit(): void {

    this.expandRootNodesTimer.subscribe(
      currTime => {
        // expand nodes that were exapanded before
        if (this.treeControl.dataNodes != undefined) {
          this.treeControl.dataNodes.forEach(
            node => {
              if (node.uniqueIdPerStack == 13) {
                this.treeControl.expand(node)
              }
            }
          )
        }
      }
    )

    this.gongdocFrontRepoService.pull().subscribe(
      gongdocFrontRepo => {
        this.gongdocFrontRepo = gongdocFrontRepo

        // // listen to any new diagram draw in order to update the 
        // // gong tree appaeance
        ClassdiagramContextSubject.subscribe(
          classdiagramContext => {
            this.classdiagramContext = classdiagramContext
            this.currentClassdiagram = this.gongdocFrontRepo.Classdiagrams.get(classdiagramContext.ClassdiagramID)
            this.refresh()
          }
        )
      }
    )
  }

  refresh(): void {
    this.gongFrontRepoService.pull().subscribe(frontRepo => {
      this.gongFrontRepo = frontRepo

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
      * fill up the GongStruct part of the mat tree
      */
      let gongstructGongNodeStruct: GongNode = {
        name: "GongStruct",
        type: gongdoc.GongdocNodeType.ROOT_OF_GONG_STRUCTS,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongStruct",
        children: new Array<GongNode>(),
        instancePresentInDiagram: false,
        draggable: false,
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongstructGongNodeStruct)

      // create the set of classshapes presents in the class diagram
      let arrayOfDispaledGongStruct = new Map<string, ClassshapeDB>()
      this.currentClassdiagram.Classshapes?.forEach(
        classshape => {
          arrayOfDispaledGongStruct.set(classshape.Structname, classshape)
        }
      )

      this.gongFrontRepo.GongStructs_array.forEach(
        gongstructDB => {

          let classshape = arrayOfDispaledGongStruct.get(gongstructDB.Name)

          let gongstructGongNodeInstance: GongNode = {
            name: gongstructDB.Name,
            type: gongdoc.GongdocNodeType.GONG_STRUCT,
            id: gongstructDB.ID,
            uniqueIdPerStack: gong.getGongStructUniqueID(gongstructDB.ID),
            structName: gongstructDB.Name,
            children: new Array<GongNode>(),
            instancePresentInDiagram: arrayOfDispaledGongStruct.has(gongstructDB.Name),
            draggable: false,
          }
          gongstructGongNodeStruct.children.push(gongstructGongNodeInstance)

          // insertion point for per field code 
          /**
          * let append a node for the slide of pointer GongBasicFields
          */
          let GongBasicFieldsGongNodeAssociation: GongNode = {
            name: "GongBasicFields",
            type: gongdoc.GongdocNodeType.ROOT_OF_BASIC_FIELDS,
            id: 0,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStruct",
            children: new Array<GongNode>(),
            draggable: false,
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(GongBasicFieldsGongNodeAssociation)

          let arrayOfDispaledGongField = new Map<string, FieldDB>()
          classshape?.Fields?.forEach(
            field => {
              arrayOfDispaledGongField.set(field.Fieldname, field)
            }
          )

          gongstructDB.GongBasicFields?.forEach(gongbasicfieldDB => {
            let draggable = false
            if (classshape && !arrayOfDispaledGongField.has(gongbasicfieldDB.Name)) {
              draggable = true
            }

            let gongbasicfieldNode: GongNode = {
              name: gongbasicfieldDB.Name,
              type: gongdoc.GongdocNodeType.BASIC_FIELD,
              id: gongbasicfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getGongBasicFieldUniqueID(gongbasicfieldDB.ID),
              structName: gongstructDB.Name,
              gongBasicField: gongbasicfieldDB,
              children: new Array<GongNode>(),
              instancePresentInDiagram: arrayOfDispaledGongField.has(gongbasicfieldDB.Name),
              draggable: draggable,
            }
            GongBasicFieldsGongNodeAssociation.children.push(gongbasicfieldNode)
          })

          /**
          * let append a node for the slide of pointer PointerToGongStructFields
          */
          let PointerToGongStructFieldsGongNodeAssociation: GongNode = {
            name: "PointerToGongStructFields",
            type: gongdoc.GongdocNodeType.ROOT_OF_POINTER_TO_STRUCT_FIELDS,
            id: 0,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: gongstructDB.Name,
            children: new Array<GongNode>(),
            draggable: false,
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(PointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.PointerToGongStructFields?.forEach(pointertogongstructfieldDB => {
            let pointertogongstructfieldNode: GongNode = {
              name: pointertogongstructfieldDB.Name,
              type: gongdoc.GongdocNodeType.POINTER_TO_STRUCT,
              id: pointertogongstructfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getPointerToGongStructFieldUniqueID(pointertogongstructfieldDB.ID),
              structName: gongstructDB.Name,
              children: new Array<GongNode>(),
              draggable: false,
            }
            PointerToGongStructFieldsGongNodeAssociation.children.push(pointertogongstructfieldNode)
          })

          /**
          * let append a node for the slide of pointer SliceOfPointerToGongStructFields
          */
          let SliceOfPointerToGongStructFieldsGongNodeAssociation: GongNode = {
            name: "SliceOfPointerToGongStructFields",
            type: gongdoc.GongdocNodeType.ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS,
            id: 0,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: gongstructDB.Name,
            children: new Array<GongNode>(),
            draggable: false,
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(SliceOfPointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.SliceOfPointerToGongStructFields?.forEach(sliceofpointertogongstructfieldDB => {
            let sliceofpointertogongstructfieldNode: GongNode = {
              name: sliceofpointertogongstructfieldDB.Name,
              type: gongdoc.GongdocNodeType.SLICE_OF_POINTER_TO_STRUCT,
              id: sliceofpointertogongstructfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getSliceOfPointerToGongStructFieldUniqueID(sliceofpointertogongstructfieldDB.ID),
              structName: gongstructDB.Name,
              children: new Array<GongNode>(),
              draggable: false,
            }
            SliceOfPointerToGongStructFieldsGongNodeAssociation.children.push(sliceofpointertogongstructfieldNode)
          })
        })

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
  }

  /**
   * removeNodeFromDiagram is called from the html template
   * 
   * @param gongFlatNode 
   */
  removeNodeFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type

      switch (gongFlatNode.type) {
        case 'BASIC_FIELD':
          gongdocCommandSingloton.FieldName = gongFlatNode.name
      }

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  /**
   * dropped is called from the html template
   * 
   * @param event 
   * @param gongFlatNode 
   * @returns 
   */
  dropped(event: CdkDragDrop<ElementRef<HTMLElement>, any>, gongFlatNode: GongFlatNode): void {

    /**
     * dragDistance is the distance betweeen the point when the element is taken
     * and the point where the element is dropped
     */
    const dragVector = event.distance;

    /**
     * originDraggedElement is the position with the absolute coordinates of the starting point
     */
    const originDraggedElement = this.getPositionFromNativeElement(
      event.item.element.nativeElement
    );

    /**
     * droppedPosition = starting point + displacement
     */
    const droppedPosition = this.getDroppingPosition(
      dragVector,
      originDraggedElement
    );
    const paper = document.getElementById('jointjs-holder');
    if (paper == null) {
      console.warn(
        "DragDropService - Diagram paper must be definied in the HTML under the id 'jointjs-holder' (#jointjs-holder)"
      );
      return;
    }

    /**
     * the innerHTMLelement provide an offset in x
     */
    const sourceElement = this.getOffsetWidthHeightFromNativeElement(this.innerHTMLelement.nativeElement)

    /**
     * 64
     */
    const originPaper = this.getPositionFromNativeElement(paper);

    /**
     * dropOnPaperOffset is the computed position the jointjs paper
     */
    const dropOnPaperOffset = {
      x: droppedPosition.x - originPaper.x - sourceElement.x,
      y: droppedPosition.y - originPaper.y,
    };

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {

      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.PositionX = dropOnPaperOffset.x
      gongdocCommandSingloton.PositionY = dropOnPaperOffset.y
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type

      switch (gongFlatNode.type) {
        case 'BASIC_FIELD':
          gongdocCommandSingloton.FieldName = gongFlatNode.name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongBasicField?.BasicKindName
      }

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  getPositionFromNativeElement(el: HTMLElement): DragDropPosition {
    return {
      x: el.offsetLeft,
      y: el.offsetTop,
    };
  }

  getOffsetWidthHeightFromNativeElement(el: HTMLElement): DragDropPosition {
    return {
      x: el.offsetWidth,
      y: el.offsetHeight,
    };
  }

  getDroppingPosition(
    dragDropDistance: DragDropPosition,
    dragOrigin: DragDropPosition
  ): DragDropPosition {
    return {
      x: dragOrigin.x + dragDropDistance.x,
      y: dragOrigin.y + dragDropDistance.y,
    };

  }
}

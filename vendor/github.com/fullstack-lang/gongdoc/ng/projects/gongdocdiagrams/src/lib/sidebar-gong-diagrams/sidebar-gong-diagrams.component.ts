import { Component, OnInit, ViewChild } from '@angular/core';
import { CdkDragDrop } from '@angular/cdk/drag-drop';
import { ElementRef } from '@angular/core';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

// insertion point for per struct import code
import * as gong from 'gong'
import * as gongdoc from 'gongdoc'

import { ClassdiagramContextSubject, ClassdiagramContext } from '../diagram-displayed-gongstruct'
import { combineLatest, Observable, timer } from 'rxjs';
import { ClassshapeDB, FieldDB, GongdocCommandDB, GongNodeType, LinkDB } from 'gongdoc';
import { NoDataRowOutlet } from '@angular/cdk/table';

/**
 * GongNode is the "data" node that is construed from the gongFrontRepo
 * 
 */
class GongNode {
  name: string = "" // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[]
  type: gongdoc.GongdocNodeType = gongdoc.GongNodeType.STRUCT as unknown as gongdoc.GongdocNodeType
  structName: string = ""
  id: number = 0
  uniqueIdPerStack: number = 0

  // specific field for gongdoc
  presentInDiagram: boolean = false // if the corresponding graphic element is present in the diagram, this value is true
  gongBasicField: gong.GongBasicFieldDB = new gong.GongBasicFieldDB
  gongPointerToGongStructField: gong.PointerToGongStructFieldDB = new gong.PointerToGongStructFieldDB
  gongSliceOfPointerToGongStructField: gong.SliceOfPointerToGongStructFieldDB = new gong.SliceOfPointerToGongStructFieldDB

  // if the node is for a link between two gong struct, both have to be present in the diagram
  // When both are present, canBeIncluded is computed to true and is checked wether the node can be displayed
  canBeIncluded: boolean = false

}

/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
export class GongFlatNode {
  expandable: boolean = false
  name: string = ""
  level: number = 0
  type: gongdoc.GongdocNodeType = gongdoc.GongNodeType.STRUCT as unknown as gongdoc.GongdocNodeType
  structName: string = ""
  id: number = 0
  uniqueIdPerStack: number = 0

  // specific field for gongdoc
  presentInDiagram: boolean = false
  gongBasicField: gong.GongBasicFieldDB = new gong.GongBasicFieldDB
  gongPointerToGongStructField: gong.PointerToGongStructFieldDB = new gong.PointerToGongStructFieldDB
  gongSliceOfPointerToGongStructField: gong.SliceOfPointerToGongStructFieldDB = new gong.SliceOfPointerToGongStructFieldDB
  canBeIncluded: boolean = false
}

export interface DragDropPosition {
  x: number;
  y: number;
}

/**
 * SidebarGongDiagramsComponent is the component that displays all gongstructs of the gong model
 * that is modeled
 * 
 * each gong struct have:
 * - basic fields
 * - pointer to gong struct
 * - slice of pointer to gong struct
 * 
 * the component deals with actions on those elements
 * 
 * the sidebar component is bespoke rework of the default gong generated sidebar
 */
@Component({
  styleUrls: ['./sidebar-gong-diagrams.css'],
  selector: 'lib-sidebar-gong-diagams',
  templateUrl: './sidebar-gong-diagrams.html'
})
export class SidebarGongDiagramsComponent implements OnInit {

  /**
   * innerHTMLelement is the html elemnt of the diagram
   * it allows the Sidebar component to devise the drop position of the gongstruct
   */
  /**
   * innerHTMLelement is the html elemnt of the diagram
   * it allows the Sidebar component to devise the drop position of the gongstruct
   */

  @ViewChild('innerHTMLelement')
  innerHTMLelement!: ElementRef;

  // the classdiagram that is currently displayed
  currentClassdiagram: gongdoc.ClassdiagramDB = new gongdoc.ClassdiagramDB

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

    let gongFlatNode: GongFlatNode = new GongFlatNode

    gongFlatNode.expandable = !!node.children && node.children.length > 0
    gongFlatNode.name = node.name
    gongFlatNode.level = level
    gongFlatNode.type = node.type
    gongFlatNode.structName = node.structName
    gongFlatNode.id = node.id
    gongFlatNode.uniqueIdPerStack = node.uniqueIdPerStack

    // specific to gongdoc
    gongFlatNode.presentInDiagram = node.presentInDiagram
    gongFlatNode.gongBasicField = node.gongBasicField
    gongFlatNode.gongPointerToGongStructField = node.gongPointerToGongStructField
    gongFlatNode.gongSliceOfPointerToGongStructField = node.gongSliceOfPointerToGongStructField
    gongFlatNode.canBeIncluded = node.canBeIncluded

    return gongFlatNode
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
  gongFrontRepo: gong.FrontRepo = new gong.FrontRepo
  gongdocFrontRepo: gongdoc.FrontRepo = new gongdoc.FrontRepo

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // provide the current display context
  classdiagramContext: ClassdiagramContext = new ClassdiagramContext

  constructor(
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

        // listen to any new diagram draw in order to update the 
        // gong tree appaeance
        ClassdiagramContextSubject.subscribe(
          classdiagramContext => {
            this.classdiagramContext = classdiagramContext
            this.currentClassdiagram = this.gongdocFrontRepo.Classdiagrams.get(classdiagramContext.ClassdiagramID)!
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
              memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
            } else {
              memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction

      /**
      * fill up the GongStruct part of the mat tree
      */
      let gongstructGongNodeStruct: GongNode = new GongNode

      gongstructGongNodeStruct.name = "GongStruct"
      gongstructGongNodeStruct.type = gongdoc.GongdocNodeType.ROOT_OF_GONG_STRUCTS
      gongstructGongNodeStruct.id = 0
      gongstructGongNodeStruct.uniqueIdPerStack = 13 * nonInstanceNodeId
      nonInstanceNodeId = nonInstanceNodeId + 1

      gongstructGongNodeStruct.structName = "GongStruct"
      gongstructGongNodeStruct.children = new Array<GongNode>()

      // the root node is neither present not draggable
      gongstructGongNodeStruct.presentInDiagram = false,


        this.gongNodeTree.push(gongstructGongNodeStruct)

      // create the set of classshapes presents in the class diagram
      // important for knowing which shapes are already displayed are 
      let arrayOfDisplayedClassshape = new Map<string, ClassshapeDB>()
      let arrayOfDisplayedBasicField = new Map<string, FieldDB>()
      let arrayOfDisplayedLink = new Map<string, LinkDB>()

      this.currentClassdiagram.Classshapes?.forEach(
        classshape => {
          arrayOfDisplayedClassshape.set(classshape.Structname, classshape)
          classshape?.Fields?.forEach(
            field => {
              arrayOfDisplayedBasicField.set(classshape.Structname + "." + field.Fieldname, field)
              // console.log("Adding " + classshape.Structname + "." + field.Fieldname)
            }
          )
          classshape?.Links?.forEach(
            link => {
              arrayOfDisplayedLink.set(classshape.Structname + "." + link.Fieldname, link)
            }
          )
        }
      )

      this.gongFrontRepo.GongStructs_array.forEach(
        gongstructDB => {

          let classshape = arrayOfDisplayedClassshape.get(gongstructDB.Name)

          let gongstructGongNodeInstance: GongNode = new GongNode
          gongstructGongNodeInstance.name = gongstructDB.Name
          gongstructGongNodeInstance.type = gongdoc.GongdocNodeType.GONG_STRUCT
          gongstructGongNodeInstance.id = gongstructDB.ID
          gongstructGongNodeInstance.uniqueIdPerStack = gong.getGongStructUniqueID(gongstructDB.ID)
          gongstructGongNodeInstance.structName = gongstructDB.Name
          gongstructGongNodeInstance.children = new Array<GongNode>()

          // specific to gongdoc
          gongstructGongNodeInstance.presentInDiagram = arrayOfDisplayedClassshape.has(gongstructDB.Name)


          gongstructGongNodeStruct.children!.push(gongstructGongNodeInstance)

          // insertion point for per field code 
          /**
          * let append a node for the slide of pointer GongBasicFields
          */
          let gongBasicFieldsGongNodeAssociation: GongNode = new GongNode


          gongBasicFieldsGongNodeAssociation.name = "GongBasicFields"
          gongBasicFieldsGongNodeAssociation.type = gongdoc.GongdocNodeType.ROOT_OF_BASIC_FIELDS
          gongBasicFieldsGongNodeAssociation.id = 0
          gongBasicFieldsGongNodeAssociation.uniqueIdPerStack = 19 * nonInstanceNodeId
          gongBasicFieldsGongNodeAssociation.structName = "GongStruct"
          gongBasicFieldsGongNodeAssociation.children = new Array<GongNode>()


          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(gongBasicFieldsGongNodeAssociation)

          gongstructDB.GongBasicFields?.forEach(gongbasicfieldDB => {

            let structIsPresent = arrayOfDisplayedClassshape.has(gongbasicfieldDB.GongStruct_GongBasicFields_reverse!.Name)

            let presentInDiagram = arrayOfDisplayedBasicField.has(gongstructDB.Name + "." + gongbasicfieldDB.Name)

            let gongbasicfieldNode: GongNode = new GongNode

            gongbasicfieldNode.name = gongbasicfieldDB.Name
            gongbasicfieldNode.type = gongdoc.GongdocNodeType.BASIC_FIELD
            gongbasicfieldNode.id = gongbasicfieldDB.ID
            gongbasicfieldNode.uniqueIdPerStack =  // godel numbering (thank you kurt)
              7 * gong.getGongStructUniqueID(gongstructDB.ID)
              + 11 * gong.getGongBasicFieldUniqueID(gongbasicfieldDB.ID)
            gongbasicfieldNode.structName = gongstructDB.Name
            gongbasicfieldNode.gongBasicField = gongbasicfieldDB
            gongbasicfieldNode.children = new Array<GongNode>()
            gongbasicfieldNode.presentInDiagram = presentInDiagram
            gongbasicfieldNode.canBeIncluded = structIsPresent

            gongBasicFieldsGongNodeAssociation.children!.push(gongbasicfieldNode)
          })

          /**
          * let append a node for the slide of pointer GongTimeFields
          */
          let gongTimeFieldsGongNodeAssociation: GongNode = new GongNode

          gongTimeFieldsGongNodeAssociation.name = "GongTimeFields"
          gongTimeFieldsGongNodeAssociation.type = gongdoc.GongdocNodeType.ROOT_OF_TIME_FIELDS
          gongTimeFieldsGongNodeAssociation.id = 0
          gongTimeFieldsGongNodeAssociation.uniqueIdPerStack = 19 * nonInstanceNodeId
          gongTimeFieldsGongNodeAssociation.structName = "GongStruct"
          gongTimeFieldsGongNodeAssociation.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(gongTimeFieldsGongNodeAssociation)

          gongstructDB.GongTimeFields?.forEach(
            gongtimefieldDB => {

              let structIsPresent = arrayOfDisplayedClassshape.has(gongtimefieldDB.GongStruct_GongTimeFields_reverse!.Name)

              let presentInDiagram = arrayOfDisplayedBasicField.has(gongstructDB.Name + "." + gongtimefieldDB.Name)

              let gongbasicfieldNode: GongNode = new GongNode

              gongbasicfieldNode.name = gongtimefieldDB.Name
              gongbasicfieldNode.type = gongdoc.GongdocNodeType.TIME_FIELD
              gongbasicfieldNode.id = gongtimefieldDB.ID
              gongbasicfieldNode.uniqueIdPerStack = // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getGongBasicFieldUniqueID(gongtimefieldDB.ID)
              gongbasicfieldNode.structName = gongstructDB.Name
              gongbasicfieldNode.children = new Array<GongNode>()
              gongbasicfieldNode.presentInDiagram = presentInDiagram
              gongbasicfieldNode.canBeIncluded = structIsPresent

              gongTimeFieldsGongNodeAssociation.children!.push(gongbasicfieldNode)
            })

          /**
          * let append a node for the slide of pointer PointerToGongStructFields
          */
          let pointerToGongStructFieldsGongNodeAssociation: GongNode = new GongNode
          pointerToGongStructFieldsGongNodeAssociation.name = "PointerToGongStructFields"
          pointerToGongStructFieldsGongNodeAssociation.type = gongdoc.GongdocNodeType.ROOT_OF_POINTER_TO_STRUCT_FIELDS
          pointerToGongStructFieldsGongNodeAssociation.id = 0
          pointerToGongStructFieldsGongNodeAssociation.uniqueIdPerStack = 19 * nonInstanceNodeId
          pointerToGongStructFieldsGongNodeAssociation.structName = gongstructDB.Name
          pointerToGongStructFieldsGongNodeAssociation.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(pointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.PointerToGongStructFields?.forEach(pointerToGongstructFieldDB => {

            // console.log("field " + pointertogongstructfieldDB.Name)
            let sourceIsPresent = arrayOfDisplayedClassshape.has(pointerToGongstructFieldDB.GongStruct_PointerToGongStructFields_reverse!.Name)
            // console.log("source " + pointertogongstructfieldDB.GongStruct_PointerToGongStructFields_reverse.Name + " is present " + sourceIsPresent)

            // compute wether the link can be included in the diagram
            let destinationIsPresent = arrayOfDisplayedClassshape.has(pointerToGongstructFieldDB.GongStruct!.Name)
            // console.log("destination " + pointertogongstructfieldDB.GongStruct.Name + " is present " + destinationIsPresent)

            let canBeIncluded = sourceIsPresent && destinationIsPresent

            let pointertogongstructfieldNode: GongNode = new GongNode

            pointertogongstructfieldNode.name = pointerToGongstructFieldDB.Name + " (" + pointerToGongstructFieldDB.GongStruct!.Name + ")"
            pointertogongstructfieldNode.type = gongdoc.GongdocNodeType.POINTER_TO_STRUCT
            pointertogongstructfieldNode.id = pointerToGongstructFieldDB.ID
            pointertogongstructfieldNode.uniqueIdPerStack = // godel numbering (thank you kurt)
              7 * gong.getGongStructUniqueID(gongstructDB.ID)
              + 11 * gong.getPointerToGongStructFieldUniqueID(pointerToGongstructFieldDB.ID)
            pointertogongstructfieldNode.structName = gongstructDB.Name
            pointertogongstructfieldNode.children = new Array<GongNode>()
            pointertogongstructfieldNode.gongPointerToGongStructField = pointerToGongstructFieldDB
            pointertogongstructfieldNode.presentInDiagram = arrayOfDisplayedLink.has(gongstructDB.Name + "." + pointerToGongstructFieldDB.Name)
            pointertogongstructfieldNode.canBeIncluded = canBeIncluded

            // console.log("can be included ? " + pointertogongstructfieldNode.name + " " +
            //   pointertogongstructfieldNode.canBeIncluded + " " +
            //   canBeIncluded)
            pointerToGongStructFieldsGongNodeAssociation.children!.push(pointertogongstructfieldNode)
          })

          /**
          * let append a node for the slide of pointer SliceOfPointerToGongStructFields
          */
          let sliceOfPointerToGongStructFieldsGongNodeAssociation: GongNode = new GongNode
          sliceOfPointerToGongStructFieldsGongNodeAssociation.name = "SliceOfPointerToGongStructFields"
          sliceOfPointerToGongStructFieldsGongNodeAssociation.type = gongdoc.GongdocNodeType.ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS
          sliceOfPointerToGongStructFieldsGongNodeAssociation.id = 0
          sliceOfPointerToGongStructFieldsGongNodeAssociation.uniqueIdPerStack = 19 * nonInstanceNodeId
          sliceOfPointerToGongStructFieldsGongNodeAssociation.structName = gongstructDB.Name
          sliceOfPointerToGongStructFieldsGongNodeAssociation.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(sliceOfPointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.SliceOfPointerToGongStructFields?.forEach(sliceofpointertogongstructfieldDB => {

            let sourceIsPresent = arrayOfDisplayedClassshape.has(sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFields_reverse!.Name)
            let destinationIsPresent = arrayOfDisplayedClassshape.has(sliceofpointertogongstructfieldDB.GongStruct!.Name)
            let canBeIncluded = sourceIsPresent && destinationIsPresent
            let presentInDiagram = arrayOfDisplayedLink.has(gongstructDB.Name + "." + sliceofpointertogongstructfieldDB.Name)

            let sliceofpointertogongstructfieldNode: GongNode = new GongNode

            sliceofpointertogongstructfieldNode.name = sliceofpointertogongstructfieldDB.Name + " (" + sliceofpointertogongstructfieldDB.GongStruct!.Name + ")"
            sliceofpointertogongstructfieldNode.type = gongdoc.GongdocNodeType.SLICE_OF_POINTER_TO_STRUCT
            sliceofpointertogongstructfieldNode.id = sliceofpointertogongstructfieldDB.ID
            sliceofpointertogongstructfieldNode.uniqueIdPerStack =// godel numbering (thank you kurt)
              7 * gong.getGongStructUniqueID(gongstructDB.ID)
              + 11 * gong.getSliceOfPointerToGongStructFieldUniqueID(sliceofpointertogongstructfieldDB.ID)
            sliceofpointertogongstructfieldNode.structName = gongstructDB.Name
            sliceofpointertogongstructfieldNode.children = new Array<GongNode>()
            sliceofpointertogongstructfieldNode.gongSliceOfPointerToGongStructField = sliceofpointertogongstructfieldDB
            sliceofpointertogongstructfieldNode.presentInDiagram = presentInDiagram
            sliceofpointertogongstructfieldNode.canBeIncluded = canBeIncluded

            sliceOfPointerToGongStructFieldsGongNodeAssociation.children!.push(sliceofpointertogongstructfieldNode)
          })
        })

      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
              this.treeControl.expand(node)
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
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

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
    )
  }

  /**
   * removeBasicFieldFromDiagram is called from the html template
   * 
   * @param gongFlatNode 
   */
  removeBasicFieldFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.FieldName = gongFlatNode.name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }

  addBasicFieldToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.FieldName = gongFlatNode.name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongBasicField?.BasicKindName

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }

  removePointerToStructFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.FieldName = gongFlatNode.gongPointerToGongStructField.Name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongPointerToGongStructField.GongStruct!.Name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }

  addPointerToStructToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.FieldName = gongFlatNode.gongPointerToGongStructField?.Name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongPointerToGongStructField?.GongStruct!.Name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }

  removeSliceOfPointerToStructFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.FieldName = gongFlatNode.gongSliceOfPointerToGongStructField.Name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongSliceOfPointerToGongStructField.GongStruct!.Name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }

  addSliceOfPointerToStructToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.FieldName = gongFlatNode.gongSliceOfPointerToGongStructField?.Name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongSliceOfPointerToGongStructField?.GongStruct!.Name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
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
      x: droppedPosition.x - originPaper.x,
      y: droppedPosition.y - originPaper.y,
    };

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand


        if (gongdocCommandSingloton != undefined) {

          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
          gongdocCommandSingloton.StructName = gongFlatNode.structName
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.Date = Date.now().toString()

          // does not work, put default position
          gongdocCommandSingloton.PositionX = dropOnPaperOffset.x
          gongdocCommandSingloton.PositionY = dropOnPaperOffset.y
          // temporary
          gongdocCommandSingloton.PositionX = 20
          gongdocCommandSingloton.PositionY = 20

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
    )
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

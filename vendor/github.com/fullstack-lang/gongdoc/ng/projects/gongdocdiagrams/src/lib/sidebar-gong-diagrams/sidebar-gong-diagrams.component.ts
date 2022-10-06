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
import { ClassshapeDB, FieldDB, GongdocCommandDB, GongNodeType, LinkDB, NoteDB } from 'gongdoc';
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
              if (node.uniqueIdPerStack == 21) {
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
      gongstructGongNodeStruct.presentInDiagram = false

      this.gongNodeTree.push(gongstructGongNodeStruct)

      // create the set of classshapes presents in the class diagram
      // important for knowing which shapes are already displayed are 
      let arrayOfDisplayedClassshape = new Map<string, ClassshapeDB>()
      let arrayOfDisplayedBasicField = new Map<string, FieldDB>()
      let arrayOfDisplayedLink = new Map<string, LinkDB>()
      let arrayOfDisplayedNote = new Map<string, NoteDB>()

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
              let key = classshape.Structname + "." + link.Fieldname + "-"+ link.Fieldtypename
              // console.log("key for set: " + key)
              arrayOfDisplayedLink.set(key, link)
            }
          )
        }
      )

      this.currentClassdiagram.Notes?.forEach(
        note => {
          arrayOfDisplayedNote.set(note.Name, note)
        }
      )

      this.gongFrontRepo.GongStructs_array.forEach(
        gongstructDB => {

          let rootNode: GongNode = new GongNode
          rootNode.name = gongstructDB.Name
          rootNode.type = gongdoc.GongdocNodeType.GONG_STRUCT
          rootNode.id = gongstructDB.ID
          rootNode.uniqueIdPerStack = gong.getGongStructUniqueID(gongstructDB.ID)
          rootNode.structName = gongstructDB.Name
          rootNode.children = new Array<GongNode>()

          // specific to gongdoc
          rootNode.presentInDiagram = arrayOfDisplayedClassshape.has(gongstructDB.Name)


          gongstructGongNodeStruct.children!.push(rootNode)

          // insertion point for per field code 
          /**
          * let append a node for the slide of pointer GongBasicFields
          */
          let rootOfBasicFieldsNode: GongNode = new GongNode


          rootOfBasicFieldsNode.name = "Basic Fields"
          rootOfBasicFieldsNode.type = gongdoc.GongdocNodeType.ROOT_OF_BASIC_FIELDS
          rootOfBasicFieldsNode.id = 0
          rootOfBasicFieldsNode.uniqueIdPerStack = 19 * nonInstanceNodeId
          rootOfBasicFieldsNode.structName = "GongStruct"
          rootOfBasicFieldsNode.children = new Array<GongNode>()


          nonInstanceNodeId = nonInstanceNodeId + 1
          rootNode.children.push(rootOfBasicFieldsNode)

          gongstructDB.GongBasicFields?.forEach(gongbasicfieldDB => {

            let structIsPresentInDiagram = arrayOfDisplayedClassshape.has(gongbasicfieldDB.GongStruct_GongBasicFields_reverse!.Name)
            let fieldPresentInDiagram = arrayOfDisplayedBasicField.has(gongstructDB.Name + "." + gongbasicfieldDB.Name)
            let basicfieldNode: GongNode = new GongNode

            basicfieldNode.name = gongbasicfieldDB.Name
            basicfieldNode.type = gongdoc.GongdocNodeType.BASIC_FIELD
            basicfieldNode.id = gongbasicfieldDB.ID
            basicfieldNode.uniqueIdPerStack =  // godel numbering (thank you kurt)
              7 * gong.getGongStructUniqueID(gongstructDB.ID)
              + 11 * gong.getGongBasicFieldUniqueID(gongbasicfieldDB.ID)
            basicfieldNode.structName = gongstructDB.Name
            basicfieldNode.gongBasicField = gongbasicfieldDB
            basicfieldNode.children = new Array<GongNode>()
            basicfieldNode.presentInDiagram = fieldPresentInDiagram
            basicfieldNode.canBeIncluded = structIsPresentInDiagram

            rootOfBasicFieldsNode.children!.push(basicfieldNode)
          })

          /**
          * let append a node for GongTimeFields
          */
          let rootOfTimeFieldsNode: GongNode = new GongNode

          rootOfTimeFieldsNode.name = "Time Fields"
          rootOfTimeFieldsNode.type = gongdoc.GongdocNodeType.ROOT_OF_TIME_FIELDS
          rootOfTimeFieldsNode.id = 0
          rootOfTimeFieldsNode.uniqueIdPerStack = 23 * nonInstanceNodeId
          rootOfTimeFieldsNode.structName = "GongStruct"
          rootOfTimeFieldsNode.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          rootNode.children.push(rootOfTimeFieldsNode)

          gongstructDB.GongTimeFields?.forEach(
            gongtimefieldDB => {

              let structIsPresentInDiagram = arrayOfDisplayedClassshape.has(gongtimefieldDB.GongStruct_GongTimeFields_reverse!.Name)
              let fieldPresentInDiagram = arrayOfDisplayedBasicField.has(gongstructDB.Name + "." + gongtimefieldDB.Name)
              let timefieldNode: GongNode = new GongNode

              timefieldNode.name = gongtimefieldDB.Name
              timefieldNode.type = gongdoc.GongdocNodeType.TIME_FIELD
              timefieldNode.id = gongtimefieldDB.ID
              timefieldNode.uniqueIdPerStack = // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getGongBasicFieldUniqueID(gongtimefieldDB.ID)
              timefieldNode.structName = gongstructDB.Name
              timefieldNode.children = new Array<GongNode>()
              timefieldNode.presentInDiagram = fieldPresentInDiagram
              timefieldNode.canBeIncluded = structIsPresentInDiagram

              rootOfTimeFieldsNode.children!.push(timefieldNode)
            })

          /**
          * let append a node for the slide of pointer PointerToGongStructFields
          */
          let rootOfPointerFieldsNode: GongNode = new GongNode
          rootOfPointerFieldsNode.name = "* - 0..1 associations (pointers)"
          rootOfPointerFieldsNode.type = gongdoc.GongdocNodeType.ROOT_OF_POINTER_TO_STRUCT_FIELDS
          rootOfPointerFieldsNode.id = 0
          rootOfPointerFieldsNode.uniqueIdPerStack = 29 * nonInstanceNodeId
          rootOfPointerFieldsNode.structName = gongstructDB.Name
          rootOfPointerFieldsNode.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          rootNode.children.push(rootOfPointerFieldsNode)

          gongstructDB.PointerToGongStructFields?.forEach(pointerToGongstructFieldDB => {

            // console.log("field " + pointertogongstructfieldDB.Name)
            let sourceIsPresent = arrayOfDisplayedClassshape.has(pointerToGongstructFieldDB.GongStruct_PointerToGongStructFields_reverse!.Name)
            // console.log("source " + pointertogongstructfieldDB.GongStruct_PointerToGongStructFields_reverse.Name + " is present " + sourceIsPresent)

            // compute wether the link can be included in the diagram
            let destinationIsPresent = arrayOfDisplayedClassshape.has(pointerToGongstructFieldDB.GongStruct!.Name)
            // console.log("destination " + pointertogongstructfieldDB.GongStruct.Name + " is present " + destinationIsPresent)

            let canBeIncluded = sourceIsPresent && destinationIsPresent

            let pointerFieldNode: GongNode = new GongNode

            pointerFieldNode.name = pointerToGongstructFieldDB.Name + " (" + pointerToGongstructFieldDB.GongStruct!.Name + ")"
            pointerFieldNode.type = gongdoc.GongdocNodeType.POINTER_TO_STRUCT
            pointerFieldNode.id = pointerToGongstructFieldDB.ID
            pointerFieldNode.uniqueIdPerStack = // godel numbering (thank you kurt)
              7 * gong.getGongStructUniqueID(gongstructDB.ID)
              + 11 * gong.getPointerToGongStructFieldUniqueID(pointerToGongstructFieldDB.ID)
            pointerFieldNode.structName = gongstructDB.Name
            pointerFieldNode.children = new Array<GongNode>()
            pointerFieldNode.gongPointerToGongStructField = pointerToGongstructFieldDB
            let key = gongstructDB.Name + 
            "." + pointerToGongstructFieldDB.Name +
            "-" + pointerToGongstructFieldDB.GongStruct!.Name
            // console.log("key for has, pointers: " + key)
            pointerFieldNode.presentInDiagram = arrayOfDisplayedLink.has(key)
            pointerFieldNode.canBeIncluded = canBeIncluded

            // console.log("can be included ? " + pointertogongstructfieldNode.name + " " +
            //   pointertogongstructfieldNode.canBeIncluded + " " +
            //   canBeIncluded)
            rootOfPointerFieldsNode.children!.push(pointerFieldNode)
          })

          /**
          * let append a node for the slide of pointer SliceOfPointerToGongStructFields
          */
          let rootOfSliceOfPointersNode: GongNode = new GongNode
          rootOfSliceOfPointersNode.name = "0..1 - * associations (slice of pointers)"
          rootOfSliceOfPointersNode.type = gongdoc.GongdocNodeType.ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS
          rootOfSliceOfPointersNode.id = 0
          rootOfSliceOfPointersNode.uniqueIdPerStack = 31 * nonInstanceNodeId
          rootOfSliceOfPointersNode.structName = gongstructDB.Name
          rootOfSliceOfPointersNode.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          rootNode.children.push(rootOfSliceOfPointersNode)

          /**
          * let append a node for the root of N-M associations
          */
          let rootOfN_M_AssocNode: GongNode = new GongNode
          rootOfN_M_AssocNode.name = "N-M associations (* - *)"
          rootOfN_M_AssocNode.type = gongdoc.GongdocNodeType.ROOT_OF_M_N_ASSOCIATION_FIELDS
          rootOfN_M_AssocNode.id = 0
          rootOfN_M_AssocNode.uniqueIdPerStack = 37 * nonInstanceNodeId
          rootOfN_M_AssocNode.structName = gongstructDB.Name
          rootOfN_M_AssocNode.children = new Array<GongNode>()

          nonInstanceNodeId = nonInstanceNodeId + 1
          rootNode.children.push(rootOfN_M_AssocNode)

          gongstructDB.SliceOfPointerToGongStructFields?.forEach(sliceOfPointerField => {

            let sourceIsPresent = arrayOfDisplayedClassshape.has(sliceOfPointerField.GongStruct_SliceOfPointerToGongStructFields_reverse!.Name)
            {
              let destinationIsPresent = arrayOfDisplayedClassshape.has(sliceOfPointerField.GongStruct!.Name)
              let canBeIncluded = sourceIsPresent && destinationIsPresent
              let key = gongstructDB.Name + "." + sliceOfPointerField.Name +
              "-"+ sliceOfPointerField.GongStruct!.Name
              // console.log("key for has, slice of pointers : " + key)
              let presentInDiagram = arrayOfDisplayedLink.has(key)
  
              let sliceOfPointerFieldNode: GongNode = new GongNode
  
              sliceOfPointerFieldNode.name = sliceOfPointerField.Name + " (" + sliceOfPointerField.GongStruct!.Name + ")"
              sliceOfPointerFieldNode.type = gongdoc.GongdocNodeType.SLICE_OF_POINTER_TO_STRUCT
              sliceOfPointerFieldNode.id = sliceOfPointerField.ID
              sliceOfPointerFieldNode.uniqueIdPerStack =// godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getSliceOfPointerToGongStructFieldUniqueID(sliceOfPointerField.ID)
              sliceOfPointerFieldNode.structName = gongstructDB.Name
              sliceOfPointerFieldNode.children = new Array<GongNode>()
              sliceOfPointerFieldNode.gongSliceOfPointerToGongStructField = sliceOfPointerField
              sliceOfPointerFieldNode.presentInDiagram = presentInDiagram
              sliceOfPointerFieldNode.canBeIncluded = canBeIncluded
  
              rootOfSliceOfPointersNode.children!.push(sliceOfPointerFieldNode)  
            }

            // let append a node for the N-M associations
            // check wether the structName finish with "Use"
            let associationStructName = sliceOfPointerField.GongStruct!.Name
            if ( associationStructName.endsWith('Use') ) {
              console.log("N-M association field: " + sliceOfPointerField.Name, 
              " association struct: " + associationStructName)
              let assocStruc = this.gongFrontRepo.GongStructs_batch.get(sliceOfPointerField.GongStructID.Int64!)
              if ( assocStruc ) {
                // check that there is a pointer field that point to the end part of the association
                if (assocStruc.PointerToGongStructFields?.length != 1) {
                  console.log("Problem with N-M association fiels, there should be only one pointer field")
                } else {
                  var pointerField = assocStruc.PointerToGongStructFields![0]
                  // console.log("end part of association is " + pointerField.GongStruct!.Name)

                  
                  let n_m_AssocNode: GongNode = new GongNode
                  n_m_AssocNode.name = sliceOfPointerField.Name + " (" + pointerField.GongStruct!.Name + ")"
                  n_m_AssocNode.type = gongdoc.GongdocNodeType.M_N_ASSOCIATION_FIELD
                  n_m_AssocNode.id = sliceOfPointerField.ID
                  n_m_AssocNode.uniqueIdPerStack =// godel numbering (thank you kurt)
                    7 * gong.getGongStructUniqueID(assocStruc.ID)
                    + 11 * gong.getSliceOfPointerToGongStructFieldUniqueID(sliceOfPointerField.ID)
                  n_m_AssocNode.structName = gongstructDB.Name
                  n_m_AssocNode.children = new Array<GongNode>()
                  n_m_AssocNode.gongSliceOfPointerToGongStructField = sliceOfPointerField
                  
                  let destinationIsPresent = arrayOfDisplayedClassshape.has(pointerField.GongStruct!.Name)
                  let canBeIncluded = sourceIsPresent && destinationIsPresent
                  let key = gongstructDB.Name + 
                  "." + sliceOfPointerField.Name +
                  "-" + pointerField.GongStruct!.Name
                  // console.log("key for has, N-M associations: " + key)
                  let presentInDiagram = arrayOfDisplayedLink.has(key)
                  n_m_AssocNode.presentInDiagram = presentInDiagram
                  n_m_AssocNode.canBeIncluded = canBeIncluded
  
                  rootOfN_M_AssocNode.children!.push(n_m_AssocNode)
                }

              }
            }
          })

        })


      let rootOfGongNotesNode: GongNode = new GongNode
      rootOfGongNotesNode.name = "Notes"
      rootOfGongNotesNode.type = gongdoc.GongdocNodeType.ROOT_OF_GONG_NOTES
      rootOfGongNotesNode.id = 0
      rootOfGongNotesNode.uniqueIdPerStack = 41 * nonInstanceNodeId
      nonInstanceNodeId = nonInstanceNodeId + 1

      rootOfGongNotesNode.structName = "GongNote"
      rootOfGongNotesNode.children = new Array<GongNode>()
      // the root node is neither present not draggable
      rootOfGongNotesNode.presentInDiagram = false

      this.gongNodeTree.push(rootOfGongNotesNode)

      this.gongFrontRepo.GongNotes_array.forEach(
        gongNodeDB => {

          let gongnoteNode: GongNode = new GongNode
          gongnoteNode.name = gongNodeDB.Name
          gongnoteNode.type = gongdoc.GongdocNodeType.GONG_NOTE
          gongnoteNode.id = gongNodeDB.ID
          gongnoteNode.uniqueIdPerStack = gong.getGongStructUniqueID(gongNodeDB.ID)
          gongnoteNode.structName = gongNodeDB.Name
          gongnoteNode.children = new Array<GongNode>()

          let gongNote = arrayOfDisplayedNote.get(gongNodeDB.Name)
          if (gongNote) {
            gongnoteNode.canBeIncluded = true
            gongnoteNode.presentInDiagram = true
          } else {
            gongnoteNode.canBeIncluded = true
            gongnoteNode.presentInDiagram = false

          }
          rootOfGongNotesNode.children!.push(gongnoteNode)
        }
      )


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

  removeN_M_AssociationsFromDiagram(gongFlatNode: GongFlatNode) {

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

          var assocStruc = gongFlatNode.gongSliceOfPointerToGongStructField.GongStruct!
          var pointerField = assocStruc.PointerToGongStructFields![0]
          // console.log("end part of association is " + pointerField.GongStruct!.Name)
          gongdocCommandSingloton.FieldTypeName = pointerField.GongStruct!.Name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }

  addN_M_AssociationsToDiagram(gongFlatNode: GongFlatNode) {

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

          var assocStruc = gongFlatNode.gongSliceOfPointerToGongStructField.GongStruct!
          var pointerField = assocStruc.PointerToGongStructFields![0]
          // console.log("end part of association is " + pointerField.GongStruct!.Name)
          gongdocCommandSingloton.FieldTypeName = pointerField.GongStruct!.Name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand updated")
            }
          )
        }
      }
    )
  }
  addNoteToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.NoteName = gongFlatNode.name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand for creation of note updated")
            }
          )
        }
      }
    )
  }

  removeNoteFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton: GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand

        if (gongdocCommandSingloton != undefined) {
          gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
          gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
          gongdocCommandSingloton.Date = Date.now().toString()
          gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
          gongdocCommandSingloton.NoteName = gongFlatNode.name

          this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
            GongdocCommand => {
              console.log("GongdocCommand for creation of note updated")
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

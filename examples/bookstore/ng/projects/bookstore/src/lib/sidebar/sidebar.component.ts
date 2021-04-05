import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { AreaService } from '../area.service'
import { getAreaUniqueID } from '../front-repo.service'
import { BookService } from '../book.service'
import { getBookUniqueID } from '../front-repo.service'
import { EditorService } from '../editor.service'
import { getEditorUniqueID } from '../front-repo.service'

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
  selector: 'app-bookstore-sidebar',
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
    private areaService: AreaService,
    private bookService: BookService,
    private editorService: EditorService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.areaService.AreaServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.bookService.BookServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.editorService.EditorServiceChanged.subscribe(
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
      * fill up the Area part of the mat tree
      */
      let areaGongNodeStruct: GongNode = {
        name: "Area",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Area",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(areaGongNodeStruct)

      this.frontRepo.Areas_array.forEach(
        areaDB => {
          let areaGongNodeInstance: GongNode = {
            name: areaDB.Name,
            type: GongNodeType.INSTANCE,
            id: areaDB.ID,
            uniqueIdPerStack: getAreaUniqueID(areaDB.ID),
            structName: "Area",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          areaGongNodeStruct.children.push(areaGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Subarea
          */
          let SubareaGongNodeAssociation: GongNode = {
            name: "(Area) Subarea",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: areaDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Area",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          areaGongNodeInstance.children.push(SubareaGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Subarea
            */
          if (areaDB.Subarea != undefined) {
            let areaGongNodeInstance_Subarea: GongNode = {
              name: areaDB.Subarea.Name,
              type: GongNodeType.INSTANCE,
              id: areaDB.Subarea.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAreaUniqueID(areaDB.ID)
                + 5 * getAreaUniqueID(areaDB.Subarea.ID),
              structName: "Area",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            SubareaGongNodeAssociation.children.push(areaGongNodeInstance_Subarea)
          }

        }
      )

      /**
      * fill up the Book part of the mat tree
      */
      let bookGongNodeStruct: GongNode = {
        name: "Book",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Book",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(bookGongNodeStruct)

      this.frontRepo.Books_array.forEach(
        bookDB => {
          let bookGongNodeInstance: GongNode = {
            name: bookDB.Name,
            type: GongNodeType.INSTANCE,
            id: bookDB.ID,
            uniqueIdPerStack: getBookUniqueID(bookDB.ID),
            structName: "Book",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          bookGongNodeStruct.children.push(bookGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Area
          */
          let AreaGongNodeAssociation: GongNode = {
            name: "(Area) Area",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: bookDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Book",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          bookGongNodeInstance.children.push(AreaGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Area
            */
          if (bookDB.Area != undefined) {
            let bookGongNodeInstance_Area: GongNode = {
              name: bookDB.Area.Name,
              type: GongNodeType.INSTANCE,
              id: bookDB.Area.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getBookUniqueID(bookDB.ID)
                + 5 * getAreaUniqueID(bookDB.Area.ID),
              structName: "Area",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AreaGongNodeAssociation.children.push(bookGongNodeInstance_Area)
          }

        }
      )

      /**
      * fill up the Editor part of the mat tree
      */
      let editorGongNodeStruct: GongNode = {
        name: "Editor",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Editor",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(editorGongNodeStruct)

      this.frontRepo.Editors_array.forEach(
        editorDB => {
          let editorGongNodeInstance: GongNode = {
            name: editorDB.Name,
            type: GongNodeType.INSTANCE,
            id: editorDB.ID,
            uniqueIdPerStack: getEditorUniqueID(editorDB.ID),
            structName: "Editor",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          editorGongNodeStruct.children.push(editorGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Books
          */
          let BooksGongNodeAssociation: GongNode = {
            name: "(Book) Books",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: editorDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Editor",
            associatedStructName: "Book",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          editorGongNodeInstance.children.push(BooksGongNodeAssociation)

          editorDB.Books?.forEach(bookDB => {
            let bookNode: GongNode = {
              name: bookDB.Name,
              type: GongNodeType.INSTANCE,
              id: bookDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getEditorUniqueID(editorDB.ID)
                + 11 * getBookUniqueID(bookDB.ID),
              structName: "Book",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            BooksGongNodeAssociation.children.push(bookNode)
          })

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

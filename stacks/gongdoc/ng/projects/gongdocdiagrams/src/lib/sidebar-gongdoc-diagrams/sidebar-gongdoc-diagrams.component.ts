import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from 'gongdoc'
import { ClassdiagramService } from 'gongdoc'

import { Observable, timer } from 'rxjs';

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ASSOCIATION = 'ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[];
  type: GongNodeType;
  structName: string;
  id: number;
  bdId: number;
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
  id: number;
  bdId: number;
}


@Component({
  styleUrls: ['./sidebar-gongdoc-diagrams.css'],
  selector: 'app-sidebar-gongdoc-diagrams',
  templateUrl: './sidebar-gongdoc-diagrams.component.html'
})
export class SidebarGongdocDiagramsComponent implements OnInit {

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
      bdId: node.bdId,
    }
  }

  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

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

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,

    private classdiagramService: ClassdiagramService,
  ) { }
  ngOnInit(): void {
    this.refresh()

    this.expandRootNodesTimer.subscribe(
      currTime => {
        // expand nodes that were exapanded before
        if (this.treeControl.dataNodes != undefined) {
          this.treeControl.dataNodes.forEach(
            node => {
              if (node.id == 0) {
                this.treeControl.expand(node)
              }
            }
          )
        }
      }
    )

    // observable for changes in structs
    this.classdiagramService.ClassdiagramServiceChanged.subscribe(
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

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (this.treeControl.isExpanded(node)) {
              memoryOfExpandedNodes[node.id] = true
            } else {
              memoryOfExpandedNodes[node.id] = false
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      let nodeId = 0;
      /**
      * fill up the Classdiagram part of the mat tree
      */
      let classdiagramGongNodeStruct: GongNode = {
        name: "Classdiagram",
        type: GongNodeType.STRUCT,
        id: 2 * nodeId++,
        structName: "Classdiagram",
        children: new Array<GongNode>(),
        bdId: 0,
      }
      this.gongNodeTree.push(classdiagramGongNodeStruct)

      this.frontRepo.Classdiagrams_array.forEach(
        classdiagramDB => {
          let classdiagramGongNodeInstance: GongNode = {
            name: classdiagramDB.Name,
            type: GongNodeType.INSTANCE,
            id: 3 * classdiagramDB.ID,
            structName: "Classdiagram",
            children: new Array<GongNode>(),
            bdId: classdiagramDB.ID,
          }
          classdiagramGongNodeStruct.children.push(classdiagramGongNodeInstance)

          // insertion point for per field code 
          /**
          * let append a node for the slide of pointer Classshapes
          */
          let ClassshapesGongNodeAssociation: GongNode = {
            name: "Classshapes",
            type: GongNodeType.ASSOCIATION,
            id: 2 * nodeId++,
            structName: "Classdiagram",
            children: new Array<GongNode>(),
            bdId: 0,
          }
          classdiagramGongNodeInstance.children.push(ClassshapesGongNodeAssociation)

          classdiagramDB.Classshapes?.forEach(classshapeDB => {
            let classshapeNode: GongNode = {
              name: classshapeDB.Name,
              type: GongNodeType.INSTANCE,
              id: 5 * classshapeDB.ID,
              structName: "Classshape",
              children: new Array<GongNode>(),
              bdId: classshapeDB.ID,
            }
            ClassshapesGongNodeAssociation.children.push(classshapeNode)
          })

        })

      /**
      * fill up the Umlsc part of the mat tree
      */
      let umlscGongNodeStruct: GongNode = {
        name: "Umlsc",
        type: GongNodeType.STRUCT,
        id: 2 * nodeId++,
        structName: "Umlsc",
        children: new Array<GongNode>(),
        bdId: 0,
      }
      this.gongNodeTree.push(umlscGongNodeStruct)

      this.frontRepo.Umlscs_array.forEach(
        umlscDB => {
          let umlscGongNodeInstance: GongNode = {
            name: umlscDB.Name,
            type: GongNodeType.INSTANCE,
            id: 7 * umlscDB.ID,
            structName: "Umlsc",
            children: new Array<GongNode>(),
            bdId: umlscDB.ID
          }
          umlscGongNodeStruct.children.push(umlscGongNodeInstance)

          // insertion point for per field code 
          /**
          * let append a node for the slide of pointer States
          */
          let StatesGongNodeAssociation: GongNode = {
            name: "States",
            type: GongNodeType.ASSOCIATION,
            id: 2 * nodeId++,
            structName: "Umlsc",
            children: new Array<GongNode>(),
            bdId: 0,
          }
          umlscGongNodeInstance.children.push(StatesGongNodeAssociation)

          umlscDB.States?.forEach(stateDB => {
            let stateNode: GongNode = {
              name: stateDB.Name,
              type: GongNodeType.INSTANCE,
              id: 11 * stateDB.ID,
              structName: "State",
              children: new Array<GongNode>(),
              bdId: stateDB.ID
            }
            StatesGongNodeAssociation.children.push(stateNode)
          })

        })


      this.dataSource.data = this.gongNodeTree

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes[node.id] != undefined) {
              if (memoryOfExpandedNodes[node.id]) {
                this.treeControl.expand(node)
              }
            }
          }
        )
      }

    });
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

    console.log("setEditorRouterOutlet " + path)

    this.router.navigate([{
      outlets: {
        elementeditor: [path.toLowerCase()]
      }
    }]);
  }

  // set editor outlet
  setEditorRouterOutletClassdiagram(classdiagramID: number) {

    console.log("setEditorRouterOutletClassdiagram " + classdiagramID)

    this.router.navigate([{
      outlets: {
        diagrameditor: ["classdiagram-detail", classdiagramID]
      }
    }]).catch(
      reason => {
        console.log(reason)
      }
    );
  }

  delete(node: GongFlatNode) {

    console.log("delete " + node.name)

    if (node.structName == "Classdiagram") {

      this.classdiagramService.deleteClassdiagram(node.bdId).subscribe(
        classdiagram => {
          this.classdiagramService.ClassdiagramServiceChanged.next("delete")

          console.log("classdiagram deleted " + classdiagram.Name)
        }
      )
    }

  }
}

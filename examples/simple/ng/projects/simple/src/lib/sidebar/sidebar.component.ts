import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { AclassService } from '../aclass.service'
import { getAclassUniqueID } from '../front-repo.service'
import { BclassService } from '../bclass.service'
import { getBclassUniqueID } from '../front-repo.service'

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
  selector: 'app-simple-sidebar',
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
    private aclassService: AclassService,
    private bclassService: BclassService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.aclassService.AclassServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.bclassService.BclassServiceChanged.subscribe(
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
      * fill up the Aclass part of the mat tree
      */
      let aclassGongNodeStruct: GongNode = {
        name: "Aclass",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Aclass",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(aclassGongNodeStruct)

      this.frontRepo.Aclasss_array.forEach(
        aclassDB => {
          let aclassGongNodeInstance: GongNode = {
            name: aclassDB.Name,
            type: GongNodeType.INSTANCE,
            id: aclassDB.ID,
            uniqueIdPerStack: getAclassUniqueID(aclassDB.ID),
            structName: "Aclass",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          aclassGongNodeStruct.children.push(aclassGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Associationtob
          */
          let AssociationtobGongNodeAssociation: GongNode = {
            name: "(Bclass) Associationtob",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: aclassDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Aclass",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          aclassGongNodeInstance.children.push(AssociationtobGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Associationtob
            */
          if (aclassDB.Associationtob != undefined) {
            let aclassGongNodeInstance_Associationtob: GongNode = {
              name: aclassDB.Associationtob.Name,
              type: GongNodeType.INSTANCE,
              id: aclassDB.Associationtob.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAclassUniqueID(aclassDB.ID)
                + 5 * getBclassUniqueID(aclassDB.Associationtob.ID),
              structName: "Bclass",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AssociationtobGongNodeAssociation.children.push(aclassGongNodeInstance_Associationtob)
          }

          /**
          * let append a node for the association Anotherassociationtob_2
          */
          let Anotherassociationtob_2GongNodeAssociation: GongNode = {
            name: "(Bclass) Anotherassociationtob_2",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: aclassDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Aclass",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          aclassGongNodeInstance.children.push(Anotherassociationtob_2GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Anotherassociationtob_2
            */
          if (aclassDB.Anotherassociationtob_2 != undefined) {
            let aclassGongNodeInstance_Anotherassociationtob_2: GongNode = {
              name: aclassDB.Anotherassociationtob_2.Name,
              type: GongNodeType.INSTANCE,
              id: aclassDB.Anotherassociationtob_2.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getAclassUniqueID(aclassDB.ID)
                + 5 * getBclassUniqueID(aclassDB.Anotherassociationtob_2.ID),
              structName: "Bclass",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            Anotherassociationtob_2GongNodeAssociation.children.push(aclassGongNodeInstance_Anotherassociationtob_2)
          }

          /**
          * let append a node for the slide of pointer Anarrayofb
          */
          let AnarrayofbGongNodeAssociation: GongNode = {
            name: "(Bclass) Anarrayofb",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: aclassDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Aclass",
            associatedStructName: "Bclass",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          aclassGongNodeInstance.children.push(AnarrayofbGongNodeAssociation)

          aclassDB.Anarrayofb?.forEach(bclassDB => {
            let bclassNode: GongNode = {
              name: bclassDB.Name,
              type: GongNodeType.INSTANCE,
              id: bclassDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAclassUniqueID(aclassDB.ID)
                + 11 * getBclassUniqueID(bclassDB.ID),
              structName: "Bclass",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnarrayofbGongNodeAssociation.children.push(bclassNode)
          })

          /**
          * let append a node for the slide of pointer Anotherarrayofb
          */
          let AnotherarrayofbGongNodeAssociation: GongNode = {
            name: "(Bclass) Anotherarrayofb",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: aclassDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Aclass",
            associatedStructName: "Bclass",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          aclassGongNodeInstance.children.push(AnotherarrayofbGongNodeAssociation)

          aclassDB.Anotherarrayofb?.forEach(bclassDB => {
            let bclassNode: GongNode = {
              name: bclassDB.Name,
              type: GongNodeType.INSTANCE,
              id: bclassDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAclassUniqueID(aclassDB.ID)
                + 11 * getBclassUniqueID(bclassDB.ID),
              structName: "Bclass",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnotherarrayofbGongNodeAssociation.children.push(bclassNode)
          })

          /**
          * let append a node for the slide of pointer Anarrayofa
          */
          let AnarrayofaGongNodeAssociation: GongNode = {
            name: "(Aclass) Anarrayofa",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: aclassDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Aclass",
            associatedStructName: "Aclass",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          aclassGongNodeInstance.children.push(AnarrayofaGongNodeAssociation)

          aclassDB.Anarrayofa?.forEach(aclassDB => {
            let aclassNode: GongNode = {
              name: aclassDB.Name,
              type: GongNodeType.INSTANCE,
              id: aclassDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getAclassUniqueID(aclassDB.ID)
                + 11 * getAclassUniqueID(aclassDB.ID),
              structName: "Aclass",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnarrayofaGongNodeAssociation.children.push(aclassNode)
          })

        }
      )

      /**
      * fill up the Bclass part of the mat tree
      */
      let bclassGongNodeStruct: GongNode = {
        name: "Bclass",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Bclass",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(bclassGongNodeStruct)

      this.frontRepo.Bclasss_array.forEach(
        bclassDB => {
          let bclassGongNodeInstance: GongNode = {
            name: bclassDB.Name,
            type: GongNodeType.INSTANCE,
            id: bclassDB.ID,
            uniqueIdPerStack: getBclassUniqueID(bclassDB.ID),
            structName: "Bclass",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          bclassGongNodeStruct.children.push(bclassGongNodeInstance)

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
import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { MachineService } from '../machine.service'
import { getMachineUniqueID } from '../front-repo.service'
import { SimulationService } from '../simulation.service'
import { getSimulationUniqueID } from '../front-repo.service'
import { WasherService } from '../washer.service'
import { getWasherUniqueID } from '../front-repo.service'

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
  selector: 'app-laundromat-sidebar',
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
    private machineService: MachineService,
    private simulationService: SimulationService,
    private washerService: WasherService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.machineService.MachineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.simulationService.SimulationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.washerService.WasherServiceChanged.subscribe(
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
      * fill up the Machine part of the mat tree
      */
      let machineGongNodeStruct: GongNode = {
        name: "Machine",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Machine",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(machineGongNodeStruct)

      this.frontRepo.Machines_array.forEach(
        machineDB => {
          let machineGongNodeInstance: GongNode = {
            name: machineDB.Name,
            type: GongNodeType.INSTANCE,
            id: machineDB.ID,
            uniqueIdPerStack: getMachineUniqueID(machineDB.ID),
            structName: "Machine",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          machineGongNodeStruct.children.push(machineGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Simulation part of the mat tree
      */
      let simulationGongNodeStruct: GongNode = {
        name: "Simulation",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Simulation",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(simulationGongNodeStruct)

      this.frontRepo.Simulations_array.forEach(
        simulationDB => {
          let simulationGongNodeInstance: GongNode = {
            name: simulationDB.Name,
            type: GongNodeType.INSTANCE,
            id: simulationDB.ID,
            uniqueIdPerStack: getSimulationUniqueID(simulationDB.ID),
            structName: "Simulation",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          simulationGongNodeStruct.children.push(simulationGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Machine
          */
          let MachineGongNodeAssociation: GongNode = {
            name: "(Machine) Machine",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: simulationDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Simulation",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          simulationGongNodeInstance.children.push(MachineGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Machine
            */
          if (simulationDB.Machine != undefined) {
            let simulationGongNodeInstance_Machine: GongNode = {
              name: simulationDB.Machine.Name,
              type: GongNodeType.INSTANCE,
              id: simulationDB.Machine.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getSimulationUniqueID(simulationDB.ID)
                + 5 * getMachineUniqueID(simulationDB.Machine.ID),
              structName: "Machine",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MachineGongNodeAssociation.children.push(simulationGongNodeInstance_Machine)
          }

          /**
          * let append a node for the association Washer
          */
          let WasherGongNodeAssociation: GongNode = {
            name: "(Washer) Washer",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: simulationDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Simulation",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          simulationGongNodeInstance.children.push(WasherGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Washer
            */
          if (simulationDB.Washer != undefined) {
            let simulationGongNodeInstance_Washer: GongNode = {
              name: simulationDB.Washer.Name,
              type: GongNodeType.INSTANCE,
              id: simulationDB.Washer.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getSimulationUniqueID(simulationDB.ID)
                + 5 * getWasherUniqueID(simulationDB.Washer.ID),
              structName: "Washer",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            WasherGongNodeAssociation.children.push(simulationGongNodeInstance_Washer)
          }

        }
      )

      /**
      * fill up the Washer part of the mat tree
      */
      let washerGongNodeStruct: GongNode = {
        name: "Washer",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Washer",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(washerGongNodeStruct)

      this.frontRepo.Washers_array.forEach(
        washerDB => {
          let washerGongNodeInstance: GongNode = {
            name: washerDB.Name,
            type: GongNodeType.INSTANCE,
            id: washerDB.ID,
            uniqueIdPerStack: getWasherUniqueID(washerDB.ID),
            structName: "Washer",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          washerGongNodeStruct.children.push(washerGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Machine
          */
          let MachineGongNodeAssociation: GongNode = {
            name: "(Machine) Machine",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: washerDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Washer",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          washerGongNodeInstance.children.push(MachineGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Machine
            */
          if (washerDB.Machine != undefined) {
            let washerGongNodeInstance_Machine: GongNode = {
              name: washerDB.Machine.Name,
              type: GongNodeType.INSTANCE,
              id: washerDB.Machine.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getWasherUniqueID(washerDB.ID)
                + 5 * getMachineUniqueID(washerDB.Machine.ID),
              structName: "Machine",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MachineGongNodeAssociation.children.push(washerGongNodeInstance_Machine)
          }

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

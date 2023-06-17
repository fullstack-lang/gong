import { Component, OnInit, Input } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { debounceTime } from 'rxjs/operators';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { Router, RouterState } from '@angular/router';


import * as gongtree from 'gongtree'

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface Node {
  name: string;

  gongNode: gongtree.NodeDB;
  children?: Node[];
}

/** Flat node with expandable and level information */
interface FlatNode {
  expandable: boolean;

  name: string;
  gongNode: gongtree.NodeDB;
  level: number;
}

// Tree component is a diagram selector
@Component({
  selector: 'lib-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css']
})
export class TreeComponent implements OnInit {

  @Input() name: string = ""
  @Input() GONG__StackPath: string = ""

  private _transformer = (node: Node, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,

      gongNode: node.gongNode,
      name: node.name,
      level: level,
    };
  };

  treeControl = new FlatTreeControl<FlatNode>(
    node => node.level,
    node => node.expandable,
  );

  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children,
  );

  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  hasChild = (_: number, node: FlatNode) => node.expandable;

  public gongtreeFrontRepo?: gongtree.FrontRepo

  constructor(
    private gongtreeFrontRepoService: gongtree.FrontRepoService,
    private gongtreeCommitNbFromBackService: gongtree.CommitNbFromBackService,
    private gongtreePushFromFrontNbService: gongtree.PushFromFrontNbService,
    private gongtreeNodeService: gongtree.NodeService,
    private gongtreeButtonService: gongtree.ButtonService,
    private router: Router,
  ) {
  }


  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  private commutNbFromBackSubscription: Subscription = new Subscription
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date

  ngOnInit(): void {
    console.log("TreeComponent->name : ", this.name)
    console.log("TreeComponent->GONG__StackPath : ", this.GONG__StackPath)
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)
  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }


  stopAutoRefresh(): void {
    if (this.commutNbFromBackSubscription) {
      this.commutNbFromBackSubscription.unsubscribe();
    }
  }


  startAutoRefresh(intervalMs: number): void {
    this.commutNbFromBackSubscription = this.gongtreeCommitNbFromBackService
      .getCommitNbFromBack(intervalMs, this.GONG__StackPath)
      .subscribe((commitNbFromBack: number) => {
        // console.log("TreeComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

        if (this.lastCommitNbFromBack < commitNbFromBack) {
          const d = new Date()
          console.log("TreeComponent, ", this.GONG__StackPath, " name ", this.name + d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
            ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
          this.lastCommitNbFromBack = commitNbFromBack
          this.refresh()
        }
      }
      )
  }

  refresh(): void {

    this.gongtreeFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongtreesFrontRepo => {
        this.gongtreeFrontRepo = gongtreesFrontRepo

        var treeSingloton: gongtree.TreeDB = new (gongtree.TreeDB)
        var selected: boolean = false
        for (var tree of this.gongtreeFrontRepo.Trees_array) {
          if (tree.Name == this.name) {
            treeSingloton = tree
            selected = true
          }
        }
        if (!selected) {
          console.log("no tree matching with name \"" + this.name + "\"")
          return
        }

        if (treeSingloton.RootNodes == undefined) {
          console.log("no nodes on tree " + this.name)
          return
        }

        // sort the nodes by their index, and sort their buttons as well
        treeSingloton.RootNodes.sort((a, b) =>
          (a.Tree_RootNodesDBID_Index.Int64 >
            b.Tree_RootNodesDBID_Index.Int64) ? 1 : -1)

        var rootNodes = new Array<Node>()

        if (treeSingloton.RootNodes != undefined) {
          for (var nodeDB of treeSingloton.RootNodes) {
            rootNodes.push(this.gongNodeToMatTreeNode(nodeDB))
          }
        }

        this.dataSource.data = rootNodes

        // expand nodes that were exapanded before
        this.treeControl.dataNodes?.forEach(
          node => {

            let gongNode = node.gongNode
            if (gongNode && gongNode.Buttons) {

              if (gongNode.Buttons.length > 0) {
                console.log("Button", gongNode.Name)
              }

              gongNode.Buttons.sort((a, b) =>
                (a.Node_ButtonsDBID_Index.Int64 >
                  b.Node_ButtonsDBID_Index.Int64) ? 1 : -1)
            }

            if (node.gongNode.IsExpanded) {
              this.treeControl.expand(node)
            }
          }
        )

      }
    )
  }

  gongNodeToMatTreeNode(nodeDB: gongtree.NodeDB): Node {
    var matTreeNode: Node = { name: nodeDB.Name, gongNode: nodeDB, children: [] }
    if (nodeDB.Children != undefined) {
      matTreeNode.children = nodeDB.Children.map(child => this.gongNodeToMatTreeNode(child))
    }

    return matTreeNode
  }

  toggleNodeExpansion(node: FlatNode): void {
    console.log(node.name)

    node.gongNode.IsExpanded = !node.gongNode.IsExpanded

    this.gongtreeNodeService.updateNode(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
        console.log("toggleNodeExpansion: updated node")
      }
    )
  }

  // toggling behavior is controlled from the back
  toggleNodeCheckbox(node: FlatNode): void {

    const d = new Date()
    console.log("TreeComponent ", this.GONG__StackPath, " name ", this.name, " toggleNodeCheckbox, " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` + " " + this.name)
    node.gongNode.IsChecked = !node.gongNode.IsChecked
    this.gongtreeNodeService.updateNode(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
        const d = new Date()
        console.log("toggleNodeCheckbox: updated node " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` + " " + this.name)
      }
    )
  }

  onButtonClick(node: FlatNode, button: gongtree.ButtonDB) {

    this.gongtreeButtonService.updateButton(button, this.GONG__StackPath).subscribe(
      gongtreeButton => {
        console.log("button pressed")
      }
    )
  }

  update(node: FlatNode) {
    node.gongNode.IsInEditMode = false
    this.gongtreeNodeService.updateNode(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
        console.log("node.gongNode.IsInEditMode = false, updated node")
      }
    )
  }
}

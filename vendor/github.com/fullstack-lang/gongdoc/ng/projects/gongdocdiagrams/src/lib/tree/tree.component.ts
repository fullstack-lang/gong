import { Component, OnInit, Input } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { debounceTime } from 'rxjs/operators';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { Router, RouterState } from '@angular/router';


import * as gongdoc from 'gongdoc'

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface Node {
  name: string;

  gongNode: gongdoc.NodeDB;
  children?: Node[];
}

/** Flat node with expandable and level information */
interface FlatNode {
  expandable: boolean;

  name: string;
  gongNode: gongdoc.NodeDB;
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

  // the package can be editable or not
  editable: boolean = false

  public classDiagram: gongdoc.ClassdiagramDB | undefined = undefined
  public stateDiagram: gongdoc.UmlscDB | undefined = undefined

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

  public gongdocFrontRepo?: gongdoc.FrontRepo

  constructor(
    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private gongdocCommitNbFromBackService: gongdoc.CommitNbFromBackService,
    private gongdocPushFromFrontNbService: gongdoc.PushFromFrontNbService,
    private gongdocNodeService: gongdoc.NodeService,
    private router: Router,
  ) {
  }


  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram

  checkCommitNbFromBackTimer: Observable<number> | undefined // = timer(1000);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0
  dateOfLastTimerEmission: Date = new Date

  subscribeInProgress: boolean = false

    // Since this component is not reused when a new diagram is selected, there can be many
  // instances of the diagram and each instance will stay alive. For instance,
  // the instance will be in the control flow if an observable the component subscribes to emits an event.
  // Therefore, it is mandatory to manage subscriptions in order to unscribe them on the ngOnDestroy hook
  checkGongdocCommitNbFromBackTimerSubscription: Subscription = new Subscription
  gongdocCommitNbFromBackService_getCommitNbFromBack: Subscription = new Subscription

  ngOnDestroy() {
    // console.log("on destroy")
    this.checkGongdocCommitNbFromBackTimerSubscription.unsubscribe()
    this.gongdocCommitNbFromBackService_getCommitNbFromBack.unsubscribe()
  }

  ngOnInit(): void {

    this.checkCommitNbFromBackTimer = timer(500, 500);

    this.checkGongdocCommitNbFromBackTimerSubscription = this.checkCommitNbFromBackTimer?.subscribe(
      currTime => {
        this.currTime = currTime
        this.dateOfLastTimerEmission = new Date
        // console.log("Timer emission " + this.name + " " + " currTime "+ currTime + " " + 
        // this.dateOfLastTimerEmission.toLocaleString() + `.${this.dateOfLastTimerEmission.getMilliseconds()}`)

        // see above for the explanation
        this.subscribeInProgress = true
        this.gongdocCommitNbFromBackService.getCommitNbFromBack()
          .subscribe(
            commitNbFromBack => {
              this.subscribeInProgress = false
              // console.log("TreeComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

              if (this.lastCommitNbFromBack < commitNbFromBack) {
                const d = new Date()
                console.log("TreeComponent, " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` +
                  ", last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack + " " + this.name)
                this.lastCommitNbFromBack = commitNbFromBack
                this.refresh()
              }
            }
          )

        // see above for the explanation
        // this.gongdocPushFromFrontNbService.getPushFromFrontNb().subscribe(
        //   pushFromFrontNb => {
        //     if (this.lastPushFromFrontNb < pushFromFrontNb) {

        //       console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
        //       this.refresh()
        //       this.lastPushFromFrontNb = pushFromFrontNb
        //     }
        //   }
        // )
      }
    )
  }

  refresh(): void {

    this.gongdocFrontRepoService.pull().subscribe(
      gongdocsFrontRepo => {
        this.gongdocFrontRepo = gongdocsFrontRepo

        this.gongdocFrontRepo.DiagramPackages_array.forEach(
          pkgElt => {
            this.editable = pkgElt.IsEditable
          }
        )

        var treeSingloton: gongdoc.TreeDB = new (gongdoc.TreeDB)
        var selected: boolean = false
        for (var tree of this.gongdocFrontRepo.Trees_array) {
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

        // sort the nodes by their index
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
            if (node.gongNode.IsExpanded) {
              this.treeControl.expand(node)
            }
          }
        )

      }
    )
  }

  gongNodeToMatTreeNode(nodeDB: gongdoc.NodeDB): Node {
    var matTreeNode: Node = { name: nodeDB.Name, gongNode: nodeDB, children: [] }
    if (nodeDB.Children != undefined) {
      matTreeNode.children = nodeDB.Children.map(child => this.gongNodeToMatTreeNode(child))
    }

    return matTreeNode
  }

  toggleNodeExpansion(node: FlatNode): void {
    console.log(node.name)

    node.gongNode.IsExpanded = !node.gongNode.IsExpanded

    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("toggleNodeExpansion: updated node")
      }
    )
  }

  // toggling behavior is controlled from the back
  toggleNodeCheckbox(node: FlatNode): void {

    const d = new Date()
    console.log("TreeComponent, toggleNodeCheckbox, " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` + " " + this.name)
    node.gongNode.IsChecked = !node.gongNode.IsChecked
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        const d = new Date()
        console.log("toggleNodeCheckbox: updated node " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` + " " + this.name)
      }
    )
  }

  addNewItem(node: FlatNode) {

    var gongNode: gongdoc.NodeDB = new (gongdoc.NodeDB)
    gongNode.Name = "NewDiagram"
    gongNode.HasEditButton = true
    gongNode.IsInEditMode = true
    gongNode.Node_ChildrenDBID.Valid = true
    gongNode.Node_ChildrenDBID.Int64 = node.gongNode.ID
    this.gongdocNodeService.postNode(gongNode).subscribe(
      gongdocNode => {
        console.log("post node")
      }
    )

    node.gongNode.IsExpanded = true
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("node.gongNode.IsExpanded updated node")
      }
    )


  }

  setInEditMode(node: FlatNode) {
    node.gongNode.IsInEditMode = true
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("node.gongNode.IsInEditMode = true, updated node")
      }
    )
  }

  update(node: FlatNode) {
    node.gongNode.IsInEditMode = false
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("node.gongNode.IsInEditMode = false, updated node")
      }
    )
  }


  cancelEditMode(node: FlatNode) {

    // fetch the value from the server
    this.gongdocNodeService.getNode(node.gongNode.ID).subscribe(
      gongdocNode => {
        node.gongNode.Name = gongdocNode.Name

        // and set the edit mode
        node.gongNode.IsInEditMode = false
        this.gongdocNodeService.updateNode(node.gongNode).subscribe(
          gongdocNode => {
            console.log("node.gongNode.IsInEditMode = false, updated node")
          }
        )
      }
    )
  }

  cancelDrawMode(node: FlatNode) {

    // fetch the value from the server
    // <to do>

    // and set the edit mode
    node.gongNode.IsInDrawMode = false
    node.gongNode.IsSaved = false
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("node.gongNode.IsInDrawMode = false, updated node")
      }
    )

  }

  updateDiagram(node: FlatNode) {

    node.gongNode.IsSaved = true
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("node.gongNode.IsSaved = true, updated node")

        if (gongdocNode.IsSaved) {
          // and set the edit mode
          node.gongNode.IsInDrawMode = false
          this.gongdocNodeService.updateNode(node.gongNode).subscribe(
            gongdocNode => {
              console.log("gongdocNode.IsSaved, updated node")
            }
          )
        }
      }
    )


  }

  deleteNode(node: FlatNode) {
    this.gongdocNodeService.deleteNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("delete node")
      }
    )
  }

  setInDrawMode(node: FlatNode) {
    node.gongNode.IsInDrawMode = true
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("setInDrawMode, updated node")
      }
    )
  }
}

import { Component, OnInit, Input } from '@angular/core';
import { Observable, timer } from 'rxjs';
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

  checkCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  subscribeInProgress: boolean = false

  ngOnInit(): void {

    this.checkCommitNbFromBackTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.subscribeInProgress = true
        this.gongdocCommitNbFromBackService.getCommitNbFromBack()
          .subscribe(
            commitNbFromBack => {
              this.subscribeInProgress = false
              // console.log("TreeComponent, last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)

              if (this.lastCommitNbFromBack < commitNbFromBack) {
                console.log("TreeComponent, last commit increased nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
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

        // get the diagram id from the node that is selected (if it is selected)
        for (var treeDB of this.gongdocFrontRepo.Trees_array) {
          if (treeDB.Type == gongdoc.TreeType.TREE_OF_DIAGRAMS) {
            // console.log("Tree: " + treeDB.Name)
            for (var nodeDB of treeDB.RootNodes!) {
              if (nodeDB.Children == undefined) {
                continue
              }
              switch (nodeDB.Type) {
                case gongdoc.GongdocNodeType.ROOT_OF_CLASS_DIAGRAMS:
                  for (var childNodeDB of nodeDB.Children) {
                    if (childNodeDB.IsChecked) {
                      if (childNodeDB.Classdiagram == undefined) {
                        console.log("Tree: classdiagram is undefined")
                        continue
                      }
                      this.classDiagram = childNodeDB.Classdiagram
                      this.router.navigate([{
                        outlets: {
                          diagrameditor: ["classdiagram-detail", this.classDiagram.ID]
                        }
                      }]).catch(
                        reason => {
                          console.log(reason)
                        }
                      );
                    }
                  }
                  break
                case gongdoc.GongdocNodeType.ROOT_OF_STATE_DIAGRAMS:
                  for (var childNodeDB of nodeDB.Children) {
                    if (childNodeDB.IsChecked) {
                      if (childNodeDB.Umlsc == undefined) {
                        console.log("Tree: State Chart diagram is undefined")
                      } else {
                        this.stateDiagram = childNodeDB.Umlsc
                        this.router.navigate([{
                          outlets: {
                            diagrameditor: ["umlsc-detail", this.stateDiagram.ID]
                          }
                        }]).catch(
                          reason => {
                            console.log(reason)
                          }
                        );
                      }
                    }

                  }
                  break
                default:
                  console.log("Tree: unknown node type: " + nodeDB.Type)
              }
            }
          }
        }
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
        console.log("updated node")
      }
    )
  }

  // toggling behavior is controlled from the back
  toggleNodeCheckbox(node: FlatNode): void {

    node.gongNode.IsChecked = !node.gongNode.IsChecked
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("updated node")
      }
    )
  }

  addNewItem(node: FlatNode) {

    switch (node.gongNode.Type) {
      case gongdoc.GongdocNodeType.ROOT_OF_CLASS_DIAGRAMS:
        var gongNode: gongdoc.NodeDB = new (gongdoc.NodeDB)
        gongNode.Name = "NewDiagram"
        gongNode.Type = gongdoc.GongdocNodeType.CLASS_DIAGRAM
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
            console.log("updated node")
          }
        )
    }

  }

  setInEditMode(node: FlatNode) {
    node.gongNode.IsInEditMode = true
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("updated node from set in edit mode")
      }
    )
  }

  update(node: FlatNode) {
    node.gongNode.IsInEditMode = false
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("updated node")
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
            console.log("updated node")
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
        console.log("updated node")
      }
    )

  }

  updateDiagram(node: FlatNode) {

    node.gongNode.IsSaved = true
    this.gongdocNodeService.updateNode(node.gongNode).subscribe(
      gongdocNode => {
        console.log("updated node")

        if (gongdocNode.IsSaved) {
          // and set the edit mode
          node.gongNode.IsInDrawMode = false
          this.gongdocNodeService.updateNode(node.gongNode).subscribe(
            gongdocNode => {
              console.log("updated node")
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
        console.log("updated node")
      }
    )
  }
}

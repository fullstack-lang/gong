import { Component, OnInit, Input } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { debounceTime } from 'rxjs/operators';

import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTreeModule } from '@angular/material/tree';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';


import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { Router, RouterState } from '@angular/router';


import * as gongtree from '../../../../gongtree/src/public-api'
import { IconService } from '../icon-service.service';

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface Node {
  name: string;

  gongNode: gongtree.Node;
  children?: Node[];
}

/** Flat node with expandable and level information */
interface FlatNode {
  expandable: boolean;

  name: string;
  gongNode: gongtree.Node;
  level: number;
}

// Tree component is a diagram selector
@Component({
  selector: 'lib-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css'],
  imports: [
    CommonModule,

    FormsModule,
    ReactiveFormsModule,

    MatIconModule,
    MatButtonModule,
    MatTreeModule,
    MatCheckboxModule,
    MatFormFieldModule,
  ],
  standalone: true,
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
    private iconService: IconService,
  ) {
  }


  ngOnInit(): void {

    this.gongtreeFrontRepoService.connectToWebSocket(this.GONG__StackPath).subscribe(
      gongtreesFrontRepo => {
        this.gongtreeFrontRepo = gongtreesFrontRepo

        var treeSingloton: gongtree.Tree = new (gongtree.Tree)
        var selected: boolean = false
        for (var tree of this.gongtreeFrontRepo.getFrontArray<gongtree.Tree>(gongtree.Tree.GONGSTRUCT_NAME)) {
          if (tree.Name == this.name) {
            treeSingloton = tree
            selected = true
          }
        }
        if (!selected) {
          // console.log("no tree matching with name \"" + this.name + "\"")
          return
        }

        if (treeSingloton.RootNodes == undefined) {
          // console.log("no nodes on tree " + this.name)
          return
        }

        var rootNodes = new Array<Node>()

        // register all icons
        for (let svgIcon of this.gongtreeFrontRepo.getFrontArray<gongtree.SVGIcon>(gongtree.SVGIcon.GONGSTRUCT_NAME)) {
          this.iconService.registerIcon(svgIcon.Name, svgIcon.SVG)
        }

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
                // console.log("Button", gongNode.Name)
              }
            }

            if (gongNode && gongNode.IsWithPreceedingIcon) {
              // console.log("Node with preceeding icon", gongNode.Name, gongNode.PreceedingIcon)
            }

            if (node.gongNode.IsExpanded) {
              this.treeControl.expand(node)
            }
          }
        )

      }
    )
  }

  gongNodeToMatTreeNode(nodeDB: gongtree.Node): Node {
    var matTreeNode: Node = { name: nodeDB.Name, gongNode: nodeDB, children: [] }
    if (nodeDB.Children != undefined) {
      matTreeNode.children = nodeDB.Children.map(child => this.gongNodeToMatTreeNode(child))
    }

    return matTreeNode
  }

  toggleNodeExpansion(node: FlatNode): void {
    // console.log(node.name)

    node.gongNode.IsExpanded = !node.gongNode.IsExpanded

    this.gongtreeNodeService.updateFront(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
        // console.log("toggleNodeExpansion: updated node")
      }
    )
  }

  // toggling behavior is controlled from the back
  toggleNodeCheckbox(node: FlatNode): void {
    let buttons = node.gongNode.Buttons

    const d = new Date()
    // console.log("TreeComponent ", this.GONG__StackPath, " name ", this.name, " toggleNodeCheckbox, " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` + " " + this.name)
    node.gongNode.IsChecked = !node.gongNode.IsChecked
    this.gongtreeNodeService.updateFront(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
      }
    )
  }
  // toggling behavior is controlled from the back
  toggleNodeSecondCheckbox(node: FlatNode): void {
    let buttons = node.gongNode.Buttons

    const d = new Date()
    // console.log("TreeComponent ", this.GONG__StackPath, " name ", this.name, " toggleNodeCheckbox, " + d.toLocaleTimeString() + `.${d.getMilliseconds()}` + " " + this.name)
    node.gongNode.IsSecondCheckboxChecked = !node.gongNode.IsSecondCheckboxChecked
    this.gongtreeNodeService.updateFront(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
      }
    )
  }

  onButtonClick(event: Event, node: FlatNode, button: gongtree.Button) {

    // Stop the click event from propagating to the parent node
    event.stopPropagation();

    this.gongtreeButtonService.updateFront(button, this.GONG__StackPath).subscribe(
      gongtreeButton => {
        // console.log("button pressed")
      }
    )
  }

  update(node: FlatNode) {

    node.gongNode.IsInEditMode = false
    this.gongtreeNodeService.updateFront(node.gongNode, this.GONG__StackPath).subscribe(
      gongtreeNode => {
        // console.log("node.gongNode.IsInEditMode = false, updated node")
      }
    )
  }

  onNodeClick(node: FlatNode): void {

    if (node.gongNode.IsNodeClickable) {
      this.gongtreeNodeService.updateFront(node.gongNode, this.GONG__StackPath).subscribe(
        gongtreeNode => {
          // console.log("onNodeClick: updated node")
        }
      )
    }
  }

  getNodeBackgroundColor(node: FlatNode): string {
    if (node.gongNode) {
      if (node.gongNode.BackgroundColor != "") {
        return node.gongNode.BackgroundColor
      } else {
        return 'default'
      }
    }
    return 'default'
  }

  getStyle(node: FlatNode): { 'font-style': string } {
    if (node.gongNode.FontStyle == gongtree.FontStyleEnum.ITALIC) {
      return {
        'font-style': 'italic', // You can also use 'normal' or 'oblique'
      }
    } else {
      return {
        'font-style': 'normal', // You can also use 'normal' or 'oblique'
      }
    }
  }

}

import { Component, OnInit, Input, ViewChild, ElementRef, AfterViewChecked } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { debounceTime } from 'rxjs/operators';

import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTreeModule } from '@angular/material/tree';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatRadioModule } from '@angular/material/radio';
import { MAT_TOOLTIP_DEFAULT_OPTIONS, MatTooltipDefaultOptions, MatTooltipModule } from '@angular/material/tooltip';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { Router, RouterState } from '@angular/router';

import * as tree from '../../../../tree/src/public-api'
import { IconService } from '../icon-service.service';

// Define the custom default options for the tooltips
export const myCustomTooltipDefaults: MatTooltipDefaultOptions = {
  showDelay: 500,
  hideDelay: 100,
  position: 'below',
  touchGestures: 'off',
  touchendHideDelay: 1500
}

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface Node {
  name: string;
  gongNode: tree.Node;
  children?: Node[];
}

/** Flat node with expandable and level information */
interface FlatNode {
  expandable: boolean;
  name: string;
  gongNode: tree.Node;
  level: number;
}

@Component({
  selector: 'lib-tree-specific',
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatIconModule,
    MatButtonModule,
    MatTreeModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatInputModule,
    MatRadioModule,
    MatTooltipModule,
  ],
  providers: [
    { provide: MAT_TOOLTIP_DEFAULT_OPTIONS, useValue: myCustomTooltipDefaults }
  ],
  templateUrl: './tree-specific.component.html',
  styleUrl: './tree-specific.component.css'
})
export class TreeSpecificComponent implements OnInit, AfterViewChecked {

  @Input() Name: string = ""

  // Reference to the input element for node editing
  @ViewChild('nodeNameInput') nodeNameInput?: ElementRef;

  // Track which node is currently being edited
  private currentEditingNode?: FlatNode;

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

  public gongtreeFrontRepo?: tree.FrontRepo

  constructor(
    private gongtreeFrontRepoService: tree.FrontRepoService,
    private gongtreeCommitNbFromBackService: tree.CommitNbFromBackService,
    private gongtreePushFromFrontNbService: tree.PushFromFrontNbService,
    private gongtreeNodeService: tree.NodeService,
    private gongtreeButtonService: tree.ButtonService,
    private iconService: IconService,
  ) {
  }

  ngOnInit(): void {
    this.gongtreeFrontRepoService.connectToWebSocket(this.Name).subscribe(
      gongtreesFrontRepo => {
        this.gongtreeFrontRepo = gongtreesFrontRepo

        var treeSingloton: tree.Tree = new (tree.Tree)

        for (let instance of this.gongtreeFrontRepo.getFrontArray<tree.Tree>(tree.Tree.GONGSTRUCT_NAME)) {
          treeSingloton = instance
        }


        if (treeSingloton.RootNodes == undefined) {
          // console.log("no nodes on tree " + this.name)
          return
        }

        var rootNodes = new Array<Node>()

        // register all icons
        for (let svgIcon of this.gongtreeFrontRepo.getFrontArray<tree.SVGIcon>(tree.SVGIcon.GONGSTRUCT_NAME)) {
          this.iconService.registerIcon(svgIcon.Name, svgIcon.SVG)
        }

        if (treeSingloton.RootNodes != undefined) {
          for (var nodeDB of treeSingloton.RootNodes) {
            rootNodes.push(this.gongNodeToMatTreeNode(nodeDB))
          }
        }

        this.dataSource.data = rootNodes

        // expand nodes that were expanded before
        this.treeControl.dataNodes?.forEach(
          node => {
            if (node.gongNode && node.gongNode.IsInEditMode) {
              this.currentEditingNode = node;
            }

            if (node.gongNode.IsExpanded) {
              this.treeControl.expand(node)
            }
          }
        )
      }
    )
  }

  // Focus on input field when a node enters edit mode
  ngAfterViewChecked() {
    if (this.nodeNameInput && this.currentEditingNode) {
      this.nodeNameInput.nativeElement.focus();
    }
  }

  gongNodeToMatTreeNode(nodeDB: tree.Node, visited: Set<tree.Node> = new Set()): Node {
    // Check if we've already visited this node
    if (visited.has(nodeDB)) {
      // Return a node without children to break the cycle
      return { name: nodeDB.Name + ' (cycle detected)', gongNode: nodeDB, children: [] };
    }

    // Mark this node as visited
    visited.add(nodeDB);

    var matTreeNode: Node = { name: nodeDB.Name, gongNode: nodeDB, children: [] };

    if (nodeDB.Children != undefined) {
      matTreeNode.children = nodeDB.Children.map(child =>
        this.gongNodeToMatTreeNode(child, visited)
      );
    }

    // Remove from visited set after processing (allows same node in different branches)
    visited.delete(nodeDB);

    return matTreeNode;
  }

  toggleNodeExpansion(node: FlatNode): void {
    node.gongNode.IsExpanded = !node.gongNode.IsExpanded

    this.gongtreeNodeService.updateFront(node.gongNode, this.Name).subscribe(
      gongtreeNode => {
        // console.log("toggleNodeExpansion: updated node")
      }
    )
  }

  // toggling behavior is controlled from the back
  toggleNodeCheckbox(node: FlatNode): void {
    node.gongNode.IsChecked = !node.gongNode.IsChecked
    this.gongtreeNodeService.updateFront(node.gongNode, this.Name).subscribe(
      gongtreeNode => {
        // Success
      }
    )
  }

  // toggling behavior is controlled from the back
  toggleNodeSecondCheckbox(node: FlatNode): void {
    node.gongNode.IsSecondCheckboxChecked = !node.gongNode.IsSecondCheckboxChecked
    this.gongtreeNodeService.updateFront(node.gongNode, this.Name).subscribe(
      gongtreeNode => {
        // Success
      }
    )
  }

  onButtonClick(event: Event, node: FlatNode, button: tree.Button) {
    // Stop the click event from propagating to the parent node
    event.stopPropagation();

    this.gongtreeButtonService.updateFront(button, this.Name).subscribe(
      gongtreeButton => {
        // Success
      }
    )
  }

  update(node: FlatNode) {
    // Update node name from edit field
    node.name = node.gongNode.Name;

    // Exit edit mode
    node.gongNode.IsInEditMode = false;
    this.currentEditingNode = undefined;

    this.gongtreeNodeService.updateFront(node.gongNode, this.Name).subscribe(
      gongtreeNode => {
        console.log("updated")
      }
    )
  }

  onNodeClick(node: FlatNode): void {
    if (node.gongNode.IsNodeClickable) {
      this.gongtreeNodeService.updateFront(node.gongNode, this.Name).subscribe(
        gongtreeNode => {
          // Success
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
    if (node.gongNode.FontStyle == tree.FontStyleEnum.ITALIC) {
      return {
        'font-style': 'italic',
      }
    } else {
      return {
        'font-style': 'normal',
      }
    }
  }
}
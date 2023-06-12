import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { AnimateService } from '../animate.service'
import { getAnimateUniqueID } from '../front-repo.service'
import { CircleService } from '../circle.service'
import { getCircleUniqueID } from '../front-repo.service'
import { EllipseService } from '../ellipse.service'
import { getEllipseUniqueID } from '../front-repo.service'
import { LayerService } from '../layer.service'
import { getLayerUniqueID } from '../front-repo.service'
import { LineService } from '../line.service'
import { getLineUniqueID } from '../front-repo.service'
import { LinkService } from '../link.service'
import { getLinkUniqueID } from '../front-repo.service'
import { LinkAnchoredTextService } from '../linkanchoredtext.service'
import { getLinkAnchoredTextUniqueID } from '../front-repo.service'
import { PathService } from '../path.service'
import { getPathUniqueID } from '../front-repo.service'
import { PointService } from '../point.service'
import { getPointUniqueID } from '../front-repo.service'
import { PolygoneService } from '../polygone.service'
import { getPolygoneUniqueID } from '../front-repo.service'
import { PolylineService } from '../polyline.service'
import { getPolylineUniqueID } from '../front-repo.service'
import { RectService } from '../rect.service'
import { getRectUniqueID } from '../front-repo.service'
import { RectAnchoredRectService } from '../rectanchoredrect.service'
import { getRectAnchoredRectUniqueID } from '../front-repo.service'
import { RectAnchoredTextService } from '../rectanchoredtext.service'
import { getRectAnchoredTextUniqueID } from '../front-repo.service'
import { RectLinkLinkService } from '../rectlinklink.service'
import { getRectLinkLinkUniqueID } from '../front-repo.service'
import { SVGService } from '../svg.service'
import { getSVGUniqueID } from '../front-repo.service'
import { TextService } from '../text.service'
import { getTextUniqueID } from '../front-repo.service'

import { RouteService } from '../route-service';

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
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
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
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongsvg-sidebar',
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
      associationField: node.associationField,
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
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private animateService: AnimateService,
    private circleService: CircleService,
    private ellipseService: EllipseService,
    private layerService: LayerService,
    private lineService: LineService,
    private linkService: LinkService,
    private linkanchoredtextService: LinkAnchoredTextService,
    private pathService: PathService,
    private pointService: PointService,
    private polygoneService: PolygoneService,
    private polylineService: PolylineService,
    private rectService: RectService,
    private rectanchoredrectService: RectAnchoredRectService,
    private rectanchoredtextService: RectAnchoredTextService,
    private rectlinklinkService: RectLinkLinkService,
    private svgService: SVGService,
    private textService: TextService,

    private routeService: RouteService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    console.log("Sidebar init: " + this.GONG__StackPath)

    // add the routes that will used by this side panel component and
    // by the component that are called from this component
    this.routeService.addDataPanelRoutes(this.GONG__StackPath)

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.animateService.AnimateServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.circleService.CircleServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.ellipseService.EllipseServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.layerService.LayerServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.lineService.LineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.linkService.LinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.linkanchoredtextService.LinkAnchoredTextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.pathService.PathServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.pointService.PointServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.polygoneService.PolygoneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.polylineService.PolylineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rectService.RectServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rectanchoredrectService.RectAnchoredRectServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rectanchoredtextService.RectAnchoredTextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rectlinklinkService.RectLinkLinkServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.svgService.SVGServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.textService.TextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Animate part of the mat tree
      */
      let animateGongNodeStruct: GongNode = {
        name: "Animate",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Animate",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(animateGongNodeStruct)

      this.frontRepo.Animates_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Animates_array.forEach(
        animateDB => {
          let animateGongNodeInstance: GongNode = {
            name: animateDB.Name,
            type: GongNodeType.INSTANCE,
            id: animateDB.ID,
            uniqueIdPerStack: getAnimateUniqueID(animateDB.ID),
            structName: "Animate",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          animateGongNodeStruct.children!.push(animateGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Circle part of the mat tree
      */
      let circleGongNodeStruct: GongNode = {
        name: "Circle",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Circle",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(circleGongNodeStruct)

      this.frontRepo.Circles_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Circles_array.forEach(
        circleDB => {
          let circleGongNodeInstance: GongNode = {
            name: circleDB.Name,
            type: GongNodeType.INSTANCE,
            id: circleDB.ID,
            uniqueIdPerStack: getCircleUniqueID(circleDB.ID),
            structName: "Circle",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          circleGongNodeStruct.children!.push(circleGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animations
          */
          let AnimationsGongNodeAssociation: GongNode = {
            name: "(Animate) Animations",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: circleDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Circle",
            associationField: "Animations",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          circleGongNodeInstance.children.push(AnimationsGongNodeAssociation)

          circleDB.Animations?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getCircleUniqueID(circleDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimationsGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Ellipse part of the mat tree
      */
      let ellipseGongNodeStruct: GongNode = {
        name: "Ellipse",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Ellipse",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(ellipseGongNodeStruct)

      this.frontRepo.Ellipses_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Ellipses_array.forEach(
        ellipseDB => {
          let ellipseGongNodeInstance: GongNode = {
            name: ellipseDB.Name,
            type: GongNodeType.INSTANCE,
            id: ellipseDB.ID,
            uniqueIdPerStack: getEllipseUniqueID(ellipseDB.ID),
            structName: "Ellipse",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          ellipseGongNodeStruct.children!.push(ellipseGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: ellipseDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Ellipse",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          ellipseGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          ellipseDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getEllipseUniqueID(ellipseDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Layer part of the mat tree
      */
      let layerGongNodeStruct: GongNode = {
        name: "Layer",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Layer",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(layerGongNodeStruct)

      this.frontRepo.Layers_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Layers_array.forEach(
        layerDB => {
          let layerGongNodeInstance: GongNode = {
            name: layerDB.Name,
            type: GongNodeType.INSTANCE,
            id: layerDB.ID,
            uniqueIdPerStack: getLayerUniqueID(layerDB.ID),
            structName: "Layer",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          layerGongNodeStruct.children!.push(layerGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Rects
          */
          let RectsGongNodeAssociation: GongNode = {
            name: "(Rect) Rects",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Rects",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(RectsGongNodeAssociation)

          layerDB.Rects?.forEach(rectDB => {
            let rectNode: GongNode = {
              name: rectDB.Name,
              type: GongNodeType.INSTANCE,
              id: rectDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getRectUniqueID(rectDB.ID),
              structName: "Rect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RectsGongNodeAssociation.children.push(rectNode)
          })

          /**
          * let append a node for the slide of pointer Texts
          */
          let TextsGongNodeAssociation: GongNode = {
            name: "(Text) Texts",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Texts",
            associatedStructName: "Text",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(TextsGongNodeAssociation)

          layerDB.Texts?.forEach(textDB => {
            let textNode: GongNode = {
              name: textDB.Name,
              type: GongNodeType.INSTANCE,
              id: textDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getTextUniqueID(textDB.ID),
              structName: "Text",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TextsGongNodeAssociation.children.push(textNode)
          })

          /**
          * let append a node for the slide of pointer Circles
          */
          let CirclesGongNodeAssociation: GongNode = {
            name: "(Circle) Circles",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Circles",
            associatedStructName: "Circle",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(CirclesGongNodeAssociation)

          layerDB.Circles?.forEach(circleDB => {
            let circleNode: GongNode = {
              name: circleDB.Name,
              type: GongNodeType.INSTANCE,
              id: circleDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getCircleUniqueID(circleDB.ID),
              structName: "Circle",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CirclesGongNodeAssociation.children.push(circleNode)
          })

          /**
          * let append a node for the slide of pointer Lines
          */
          let LinesGongNodeAssociation: GongNode = {
            name: "(Line) Lines",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Lines",
            associatedStructName: "Line",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(LinesGongNodeAssociation)

          layerDB.Lines?.forEach(lineDB => {
            let lineNode: GongNode = {
              name: lineDB.Name,
              type: GongNodeType.INSTANCE,
              id: lineDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getLineUniqueID(lineDB.ID),
              structName: "Line",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinesGongNodeAssociation.children.push(lineNode)
          })

          /**
          * let append a node for the slide of pointer Ellipses
          */
          let EllipsesGongNodeAssociation: GongNode = {
            name: "(Ellipse) Ellipses",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Ellipses",
            associatedStructName: "Ellipse",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(EllipsesGongNodeAssociation)

          layerDB.Ellipses?.forEach(ellipseDB => {
            let ellipseNode: GongNode = {
              name: ellipseDB.Name,
              type: GongNodeType.INSTANCE,
              id: ellipseDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getEllipseUniqueID(ellipseDB.ID),
              structName: "Ellipse",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            EllipsesGongNodeAssociation.children.push(ellipseNode)
          })

          /**
          * let append a node for the slide of pointer Polylines
          */
          let PolylinesGongNodeAssociation: GongNode = {
            name: "(Polyline) Polylines",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Polylines",
            associatedStructName: "Polyline",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(PolylinesGongNodeAssociation)

          layerDB.Polylines?.forEach(polylineDB => {
            let polylineNode: GongNode = {
              name: polylineDB.Name,
              type: GongNodeType.INSTANCE,
              id: polylineDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getPolylineUniqueID(polylineDB.ID),
              structName: "Polyline",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PolylinesGongNodeAssociation.children.push(polylineNode)
          })

          /**
          * let append a node for the slide of pointer Polygones
          */
          let PolygonesGongNodeAssociation: GongNode = {
            name: "(Polygone) Polygones",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Polygones",
            associatedStructName: "Polygone",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(PolygonesGongNodeAssociation)

          layerDB.Polygones?.forEach(polygoneDB => {
            let polygoneNode: GongNode = {
              name: polygoneDB.Name,
              type: GongNodeType.INSTANCE,
              id: polygoneDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getPolygoneUniqueID(polygoneDB.ID),
              structName: "Polygone",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PolygonesGongNodeAssociation.children.push(polygoneNode)
          })

          /**
          * let append a node for the slide of pointer Paths
          */
          let PathsGongNodeAssociation: GongNode = {
            name: "(Path) Paths",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Paths",
            associatedStructName: "Path",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(PathsGongNodeAssociation)

          layerDB.Paths?.forEach(pathDB => {
            let pathNode: GongNode = {
              name: pathDB.Name,
              type: GongNodeType.INSTANCE,
              id: pathDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getPathUniqueID(pathDB.ID),
              structName: "Path",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            PathsGongNodeAssociation.children.push(pathNode)
          })

          /**
          * let append a node for the slide of pointer Links
          */
          let LinksGongNodeAssociation: GongNode = {
            name: "(Link) Links",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "Links",
            associatedStructName: "Link",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(LinksGongNodeAssociation)

          layerDB.Links?.forEach(linkDB => {
            let linkNode: GongNode = {
              name: linkDB.Name,
              type: GongNodeType.INSTANCE,
              id: linkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getLinkUniqueID(linkDB.ID),
              structName: "Link",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinksGongNodeAssociation.children.push(linkNode)
          })

          /**
          * let append a node for the slide of pointer RectLinkLinks
          */
          let RectLinkLinksGongNodeAssociation: GongNode = {
            name: "(RectLinkLink) RectLinkLinks",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: layerDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Layer",
            associationField: "RectLinkLinks",
            associatedStructName: "RectLinkLink",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          layerGongNodeInstance.children.push(RectLinkLinksGongNodeAssociation)

          layerDB.RectLinkLinks?.forEach(rectlinklinkDB => {
            let rectlinklinkNode: GongNode = {
              name: rectlinklinkDB.Name,
              type: GongNodeType.INSTANCE,
              id: rectlinklinkDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLayerUniqueID(layerDB.ID)
                + 11 * getRectLinkLinkUniqueID(rectlinklinkDB.ID),
              structName: "RectLinkLink",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RectLinkLinksGongNodeAssociation.children.push(rectlinklinkNode)
          })

        }
      )

      /**
      * fill up the Line part of the mat tree
      */
      let lineGongNodeStruct: GongNode = {
        name: "Line",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Line",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(lineGongNodeStruct)

      this.frontRepo.Lines_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Lines_array.forEach(
        lineDB => {
          let lineGongNodeInstance: GongNode = {
            name: lineDB.Name,
            type: GongNodeType.INSTANCE,
            id: lineDB.ID,
            uniqueIdPerStack: getLineUniqueID(lineDB.ID),
            structName: "Line",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          lineGongNodeStruct.children!.push(lineGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: lineDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Line",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          lineGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          lineDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLineUniqueID(lineDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Link part of the mat tree
      */
      let linkGongNodeStruct: GongNode = {
        name: "Link",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Link",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(linkGongNodeStruct)

      this.frontRepo.Links_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Links_array.forEach(
        linkDB => {
          let linkGongNodeInstance: GongNode = {
            name: linkDB.Name,
            type: GongNodeType.INSTANCE,
            id: linkDB.ID,
            uniqueIdPerStack: getLinkUniqueID(linkDB.ID),
            structName: "Link",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          linkGongNodeStruct.children!.push(linkGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Start
          */
          let StartGongNodeAssociation: GongNode = {
            name: "(Rect) Start",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: linkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Link",
            associationField: "Start",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children!.push(StartGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Start
            */
          if (linkDB.Start != undefined) {
            let linkGongNodeInstance_Start: GongNode = {
              name: linkDB.Start.Name,
              type: GongNodeType.INSTANCE,
              id: linkDB.Start.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getLinkUniqueID(linkDB.ID)
                + 5 * getRectUniqueID(linkDB.Start.ID),
              structName: "Rect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            StartGongNodeAssociation.children.push(linkGongNodeInstance_Start)
          }

          /**
          * let append a node for the association End
          */
          let EndGongNodeAssociation: GongNode = {
            name: "(Rect) End",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: linkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Link",
            associationField: "End",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children!.push(EndGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation End
            */
          if (linkDB.End != undefined) {
            let linkGongNodeInstance_End: GongNode = {
              name: linkDB.End.Name,
              type: GongNodeType.INSTANCE,
              id: linkDB.End.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getLinkUniqueID(linkDB.ID)
                + 5 * getRectUniqueID(linkDB.End.ID),
              structName: "Rect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            EndGongNodeAssociation.children.push(linkGongNodeInstance_End)
          }

          /**
          * let append a node for the slide of pointer TextAtArrowEnd
          */
          let TextAtArrowEndGongNodeAssociation: GongNode = {
            name: "(LinkAnchoredText) TextAtArrowEnd",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: linkDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Link",
            associationField: "TextAtArrowEnd",
            associatedStructName: "LinkAnchoredText",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children.push(TextAtArrowEndGongNodeAssociation)

          linkDB.TextAtArrowEnd?.forEach(linkanchoredtextDB => {
            let linkanchoredtextNode: GongNode = {
              name: linkanchoredtextDB.Name,
              type: GongNodeType.INSTANCE,
              id: linkanchoredtextDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLinkUniqueID(linkDB.ID)
                + 11 * getLinkAnchoredTextUniqueID(linkanchoredtextDB.ID),
              structName: "LinkAnchoredText",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TextAtArrowEndGongNodeAssociation.children.push(linkanchoredtextNode)
          })

          /**
          * let append a node for the slide of pointer TextAtArrowStart
          */
          let TextAtArrowStartGongNodeAssociation: GongNode = {
            name: "(LinkAnchoredText) TextAtArrowStart",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: linkDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Link",
            associationField: "TextAtArrowStart",
            associatedStructName: "LinkAnchoredText",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children.push(TextAtArrowStartGongNodeAssociation)

          linkDB.TextAtArrowStart?.forEach(linkanchoredtextDB => {
            let linkanchoredtextNode: GongNode = {
              name: linkanchoredtextDB.Name,
              type: GongNodeType.INSTANCE,
              id: linkanchoredtextDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLinkUniqueID(linkDB.ID)
                + 11 * getLinkAnchoredTextUniqueID(linkanchoredtextDB.ID),
              structName: "LinkAnchoredText",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TextAtArrowStartGongNodeAssociation.children.push(linkanchoredtextNode)
          })

          /**
          * let append a node for the slide of pointer ControlPoints
          */
          let ControlPointsGongNodeAssociation: GongNode = {
            name: "(Point) ControlPoints",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: linkDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Link",
            associationField: "ControlPoints",
            associatedStructName: "Point",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkGongNodeInstance.children.push(ControlPointsGongNodeAssociation)

          linkDB.ControlPoints?.forEach(pointDB => {
            let pointNode: GongNode = {
              name: pointDB.Name,
              type: GongNodeType.INSTANCE,
              id: pointDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLinkUniqueID(linkDB.ID)
                + 11 * getPointUniqueID(pointDB.ID),
              structName: "Point",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ControlPointsGongNodeAssociation.children.push(pointNode)
          })

        }
      )

      /**
      * fill up the LinkAnchoredText part of the mat tree
      */
      let linkanchoredtextGongNodeStruct: GongNode = {
        name: "LinkAnchoredText",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "LinkAnchoredText",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(linkanchoredtextGongNodeStruct)

      this.frontRepo.LinkAnchoredTexts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.LinkAnchoredTexts_array.forEach(
        linkanchoredtextDB => {
          let linkanchoredtextGongNodeInstance: GongNode = {
            name: linkanchoredtextDB.Name,
            type: GongNodeType.INSTANCE,
            id: linkanchoredtextDB.ID,
            uniqueIdPerStack: getLinkAnchoredTextUniqueID(linkanchoredtextDB.ID),
            structName: "LinkAnchoredText",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          linkanchoredtextGongNodeStruct.children!.push(linkanchoredtextGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: linkanchoredtextDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "LinkAnchoredText",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linkanchoredtextGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          linkanchoredtextDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLinkAnchoredTextUniqueID(linkanchoredtextDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Path part of the mat tree
      */
      let pathGongNodeStruct: GongNode = {
        name: "Path",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Path",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(pathGongNodeStruct)

      this.frontRepo.Paths_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Paths_array.forEach(
        pathDB => {
          let pathGongNodeInstance: GongNode = {
            name: pathDB.Name,
            type: GongNodeType.INSTANCE,
            id: pathDB.ID,
            uniqueIdPerStack: getPathUniqueID(pathDB.ID),
            structName: "Path",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          pathGongNodeStruct.children!.push(pathGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: pathDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Path",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pathGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          pathDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getPathUniqueID(pathDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Point part of the mat tree
      */
      let pointGongNodeStruct: GongNode = {
        name: "Point",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Point",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(pointGongNodeStruct)

      this.frontRepo.Points_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Points_array.forEach(
        pointDB => {
          let pointGongNodeInstance: GongNode = {
            name: pointDB.Name,
            type: GongNodeType.INSTANCE,
            id: pointDB.ID,
            uniqueIdPerStack: getPointUniqueID(pointDB.ID),
            structName: "Point",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          pointGongNodeStruct.children!.push(pointGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Polygone part of the mat tree
      */
      let polygoneGongNodeStruct: GongNode = {
        name: "Polygone",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Polygone",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(polygoneGongNodeStruct)

      this.frontRepo.Polygones_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Polygones_array.forEach(
        polygoneDB => {
          let polygoneGongNodeInstance: GongNode = {
            name: polygoneDB.Name,
            type: GongNodeType.INSTANCE,
            id: polygoneDB.ID,
            uniqueIdPerStack: getPolygoneUniqueID(polygoneDB.ID),
            structName: "Polygone",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          polygoneGongNodeStruct.children!.push(polygoneGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: polygoneDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Polygone",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          polygoneGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          polygoneDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getPolygoneUniqueID(polygoneDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Polyline part of the mat tree
      */
      let polylineGongNodeStruct: GongNode = {
        name: "Polyline",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Polyline",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(polylineGongNodeStruct)

      this.frontRepo.Polylines_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Polylines_array.forEach(
        polylineDB => {
          let polylineGongNodeInstance: GongNode = {
            name: polylineDB.Name,
            type: GongNodeType.INSTANCE,
            id: polylineDB.ID,
            uniqueIdPerStack: getPolylineUniqueID(polylineDB.ID),
            structName: "Polyline",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          polylineGongNodeStruct.children!.push(polylineGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: polylineDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Polyline",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          polylineGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          polylineDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getPolylineUniqueID(polylineDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the Rect part of the mat tree
      */
      let rectGongNodeStruct: GongNode = {
        name: "Rect",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Rect",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rectGongNodeStruct)

      this.frontRepo.Rects_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Rects_array.forEach(
        rectDB => {
          let rectGongNodeInstance: GongNode = {
            name: rectDB.Name,
            type: GongNodeType.INSTANCE,
            id: rectDB.ID,
            uniqueIdPerStack: getRectUniqueID(rectDB.ID),
            structName: "Rect",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rectGongNodeStruct.children!.push(rectGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animations
          */
          let AnimationsGongNodeAssociation: GongNode = {
            name: "(Animate) Animations",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: rectDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Rect",
            associationField: "Animations",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rectGongNodeInstance.children.push(AnimationsGongNodeAssociation)

          rectDB.Animations?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getRectUniqueID(rectDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimationsGongNodeAssociation.children.push(animateNode)
          })

          /**
          * let append a node for the slide of pointer RectAnchoredTexts
          */
          let RectAnchoredTextsGongNodeAssociation: GongNode = {
            name: "(RectAnchoredText) RectAnchoredTexts",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: rectDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Rect",
            associationField: "RectAnchoredTexts",
            associatedStructName: "RectAnchoredText",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rectGongNodeInstance.children.push(RectAnchoredTextsGongNodeAssociation)

          rectDB.RectAnchoredTexts?.forEach(rectanchoredtextDB => {
            let rectanchoredtextNode: GongNode = {
              name: rectanchoredtextDB.Name,
              type: GongNodeType.INSTANCE,
              id: rectanchoredtextDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getRectUniqueID(rectDB.ID)
                + 11 * getRectAnchoredTextUniqueID(rectanchoredtextDB.ID),
              structName: "RectAnchoredText",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RectAnchoredTextsGongNodeAssociation.children.push(rectanchoredtextNode)
          })

          /**
          * let append a node for the slide of pointer RectAnchoredRects
          */
          let RectAnchoredRectsGongNodeAssociation: GongNode = {
            name: "(RectAnchoredRect) RectAnchoredRects",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: rectDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Rect",
            associationField: "RectAnchoredRects",
            associatedStructName: "RectAnchoredRect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rectGongNodeInstance.children.push(RectAnchoredRectsGongNodeAssociation)

          rectDB.RectAnchoredRects?.forEach(rectanchoredrectDB => {
            let rectanchoredrectNode: GongNode = {
              name: rectanchoredrectDB.Name,
              type: GongNodeType.INSTANCE,
              id: rectanchoredrectDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getRectUniqueID(rectDB.ID)
                + 11 * getRectAnchoredRectUniqueID(rectanchoredrectDB.ID),
              structName: "RectAnchoredRect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RectAnchoredRectsGongNodeAssociation.children.push(rectanchoredrectNode)
          })

        }
      )

      /**
      * fill up the RectAnchoredRect part of the mat tree
      */
      let rectanchoredrectGongNodeStruct: GongNode = {
        name: "RectAnchoredRect",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "RectAnchoredRect",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rectanchoredrectGongNodeStruct)

      this.frontRepo.RectAnchoredRects_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.RectAnchoredRects_array.forEach(
        rectanchoredrectDB => {
          let rectanchoredrectGongNodeInstance: GongNode = {
            name: rectanchoredrectDB.Name,
            type: GongNodeType.INSTANCE,
            id: rectanchoredrectDB.ID,
            uniqueIdPerStack: getRectAnchoredRectUniqueID(rectanchoredrectDB.ID),
            structName: "RectAnchoredRect",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rectanchoredrectGongNodeStruct.children!.push(rectanchoredrectGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the RectAnchoredText part of the mat tree
      */
      let rectanchoredtextGongNodeStruct: GongNode = {
        name: "RectAnchoredText",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "RectAnchoredText",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rectanchoredtextGongNodeStruct)

      this.frontRepo.RectAnchoredTexts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.RectAnchoredTexts_array.forEach(
        rectanchoredtextDB => {
          let rectanchoredtextGongNodeInstance: GongNode = {
            name: rectanchoredtextDB.Name,
            type: GongNodeType.INSTANCE,
            id: rectanchoredtextDB.ID,
            uniqueIdPerStack: getRectAnchoredTextUniqueID(rectanchoredtextDB.ID),
            structName: "RectAnchoredText",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rectanchoredtextGongNodeStruct.children!.push(rectanchoredtextGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: rectanchoredtextDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "RectAnchoredText",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rectanchoredtextGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          rectanchoredtextDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getRectAnchoredTextUniqueID(rectanchoredtextDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )

      /**
      * fill up the RectLinkLink part of the mat tree
      */
      let rectlinklinkGongNodeStruct: GongNode = {
        name: "RectLinkLink",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "RectLinkLink",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rectlinklinkGongNodeStruct)

      this.frontRepo.RectLinkLinks_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.RectLinkLinks_array.forEach(
        rectlinklinkDB => {
          let rectlinklinkGongNodeInstance: GongNode = {
            name: rectlinklinkDB.Name,
            type: GongNodeType.INSTANCE,
            id: rectlinklinkDB.ID,
            uniqueIdPerStack: getRectLinkLinkUniqueID(rectlinklinkDB.ID),
            structName: "RectLinkLink",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rectlinklinkGongNodeStruct.children!.push(rectlinklinkGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Start
          */
          let StartGongNodeAssociation: GongNode = {
            name: "(Rect) Start",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: rectlinklinkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "RectLinkLink",
            associationField: "Start",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rectlinklinkGongNodeInstance.children!.push(StartGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Start
            */
          if (rectlinklinkDB.Start != undefined) {
            let rectlinklinkGongNodeInstance_Start: GongNode = {
              name: rectlinklinkDB.Start.Name,
              type: GongNodeType.INSTANCE,
              id: rectlinklinkDB.Start.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getRectLinkLinkUniqueID(rectlinklinkDB.ID)
                + 5 * getRectUniqueID(rectlinklinkDB.Start.ID),
              structName: "Rect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            StartGongNodeAssociation.children.push(rectlinklinkGongNodeInstance_Start)
          }

          /**
          * let append a node for the association End
          */
          let EndGongNodeAssociation: GongNode = {
            name: "(Link) End",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: rectlinklinkDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "RectLinkLink",
            associationField: "End",
            associatedStructName: "Link",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rectlinklinkGongNodeInstance.children!.push(EndGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation End
            */
          if (rectlinklinkDB.End != undefined) {
            let rectlinklinkGongNodeInstance_End: GongNode = {
              name: rectlinklinkDB.End.Name,
              type: GongNodeType.INSTANCE,
              id: rectlinklinkDB.End.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getRectLinkLinkUniqueID(rectlinklinkDB.ID)
                + 5 * getLinkUniqueID(rectlinklinkDB.End.ID),
              structName: "Link",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            EndGongNodeAssociation.children.push(rectlinklinkGongNodeInstance_End)
          }

        }
      )

      /**
      * fill up the SVG part of the mat tree
      */
      let svgGongNodeStruct: GongNode = {
        name: "SVG",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "SVG",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(svgGongNodeStruct)

      this.frontRepo.SVGs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.SVGs_array.forEach(
        svgDB => {
          let svgGongNodeInstance: GongNode = {
            name: svgDB.Name,
            type: GongNodeType.INSTANCE,
            id: svgDB.ID,
            uniqueIdPerStack: getSVGUniqueID(svgDB.ID),
            structName: "SVG",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          svgGongNodeStruct.children!.push(svgGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Layers
          */
          let LayersGongNodeAssociation: GongNode = {
            name: "(Layer) Layers",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "SVG",
            associationField: "Layers",
            associatedStructName: "Layer",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children.push(LayersGongNodeAssociation)

          svgDB.Layers?.forEach(layerDB => {
            let layerNode: GongNode = {
              name: layerDB.Name,
              type: GongNodeType.INSTANCE,
              id: layerDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getSVGUniqueID(svgDB.ID)
                + 11 * getLayerUniqueID(layerDB.ID),
              structName: "Layer",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LayersGongNodeAssociation.children.push(layerNode)
          })

          /**
          * let append a node for the association StartRect
          */
          let StartRectGongNodeAssociation: GongNode = {
            name: "(Rect) StartRect",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "SVG",
            associationField: "StartRect",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children!.push(StartRectGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation StartRect
            */
          if (svgDB.StartRect != undefined) {
            let svgGongNodeInstance_StartRect: GongNode = {
              name: svgDB.StartRect.Name,
              type: GongNodeType.INSTANCE,
              id: svgDB.StartRect.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getSVGUniqueID(svgDB.ID)
                + 5 * getRectUniqueID(svgDB.StartRect.ID),
              structName: "Rect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            StartRectGongNodeAssociation.children.push(svgGongNodeInstance_StartRect)
          }

          /**
          * let append a node for the association EndRect
          */
          let EndRectGongNodeAssociation: GongNode = {
            name: "(Rect) EndRect",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "SVG",
            associationField: "EndRect",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children!.push(EndRectGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation EndRect
            */
          if (svgDB.EndRect != undefined) {
            let svgGongNodeInstance_EndRect: GongNode = {
              name: svgDB.EndRect.Name,
              type: GongNodeType.INSTANCE,
              id: svgDB.EndRect.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getSVGUniqueID(svgDB.ID)
                + 5 * getRectUniqueID(svgDB.EndRect.ID),
              structName: "Rect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            EndRectGongNodeAssociation.children.push(svgGongNodeInstance_EndRect)
          }

        }
      )

      /**
      * fill up the Text part of the mat tree
      */
      let textGongNodeStruct: GongNode = {
        name: "Text",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Text",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(textGongNodeStruct)

      this.frontRepo.Texts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Texts_array.forEach(
        textDB => {
          let textGongNodeInstance: GongNode = {
            name: textDB.Name,
            type: GongNodeType.INSTANCE,
            id: textDB.ID,
            uniqueIdPerStack: getTextUniqueID(textDB.ID),
            structName: "Text",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          textGongNodeStruct.children!.push(textGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Animates
          */
          let AnimatesGongNodeAssociation: GongNode = {
            name: "(Animate) Animates",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: textDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Text",
            associationField: "Animates",
            associatedStructName: "Animate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          textGongNodeInstance.children.push(AnimatesGongNodeAssociation)

          textDB.Animates?.forEach(animateDB => {
            let animateNode: GongNode = {
              name: animateDB.Name,
              type: GongNodeType.INSTANCE,
              id: animateDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getTextUniqueID(textDB.ID)
                + 11 * getAnimateUniqueID(animateDB.ID),
              structName: "Animate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            AnimatesGongNodeAssociation.children.push(animateNode)
          })

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    })
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path
    this.router.navigate([{
      outlets: {
        outletName: [fullPath]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
      let outletConf: any = {}
      outletConf[outletName] = [fullPath, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }

    if (type == GongNodeType.INSTANCE) {
      let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + structName.toLowerCase() + "-detail"

      let outletConf: any = {}
      outletConf[outletName] = [fullPath, id, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf : any = {}
    outletConf[outletName] = [fullPath, this.GONG__StackPath]
    this.router.navigate([{ outlets: outletConf }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + node.associatedStructName.toLowerCase() + "-adder"
    this.router.navigate([{
      outlets: {
        outletName: [fullPath, node.id, node.structName, node.associationField, this.GONG__StackPath]
      }
    }]);
  }
}

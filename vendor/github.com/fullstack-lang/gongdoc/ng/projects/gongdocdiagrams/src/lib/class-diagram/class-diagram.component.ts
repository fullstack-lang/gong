import { Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import * as joint from 'jointjs';

import { ActivatedRoute, Router } from '@angular/router';

import * as gongdoc from 'gongdoc'

import { newUmlClassShapeFromGongStructShape } from './newUmlClassShapeFromGongStructShape'
import { ClassdiagramContextSubject, ClassdiagramContext } from '../diagram-displayed-gongstruct'
import { newUmlNote } from './newUmlNote';
import { NONE_TYPE } from '@angular/compiler';

import { onClassshapeMove, onLinkMove, onNoteMove } from './on-move-functions'
import { shapeIdentifierToShapeName } from './shape-identifier-to-shape-name';
import { informBackEndOfSelection } from './on-pointer-down-function';
import { ClassdiagramDB } from 'gongdoc';
import { newUmlClassShapeFromGongEnumShape } from './newUmlClassShapeFromGongEnumShape';
import { IdentifierToReceiverAndFieldName, IdentifierToStructname } from './identifier-function';

@Component({
  selector: 'lib-class-diagram',
  templateUrl: './class-diagram.component.html',
  styleUrls: ['./class-diagram.component.css']
})
export class ClassDiagramComponent implements OnInit, OnDestroy {


  /**
   * the class diagram component is refreshed both by direct input when the user moves vertices or positions
   * otherwise, modification are gotten from the back repo 
   * 
   * the checkCommitNbFromBackTimer polls the commit number of the back repo
   * if the commit number has increased, it pulls the front repo and redraw the diagram
   */
  checkGongdocCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  currTime: number = 0

  /**
   * jointjs stuff
   */
  namespace = joint.shapes
  private paper?: joint.dia.Paper
  private graph?: joint.dia.Graph

  // the gong diagram of interest ot be drawn
  public classdiagram: gongdoc.ClassdiagramDB = new gongdoc.ClassdiagramDB

  /**
   * gongdoc is for CRUDing the diagram elements (shapes, links, ...)
   */
  gongdocFrontRepo: gongdoc.FrontRepo = new gongdoc.FrontRepo

  // map for storing which gong struct have a classshape
  // it is important for drawing links between shapes
  public Map_GongStructName_JointjsUMLClassShape = new Map<string, joint.shapes.uml.Class>();
  public Map_GongStructName_Joint_shapes_standard_Link = new Map<string, joint.shapes.standard.Link>();

  constructor(
    private activatedRoute: ActivatedRoute,
    private router: Router,

    private positionService: gongdoc.PositionService,
    private noteService: gongdoc.NoteShapeService,
    private verticeService: gongdoc.VerticeService,
    private gongStructShapeService: gongdoc.GongStructShapeService, // for selection of the classshape
    private gongEnumShapeService: gongdoc.GongEnumShapeService, // for selection of the classshape

    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private gongdocCommitNbFromBackService: gongdoc.CommitNbFromBackService,
  ) {
  }

  // Since this component is not reused when a new diagram is selected, there can be many
  // instances of the diagram and each instance will stay alive. For instance,
  // the instance will be in the control flow if an observable the component subscribes to emits an event.
  // Therefore, it is mandatory to manage subscriptions in order to unscribe them on the ngOnDestroy hook
  checkGongdocCommitNbFromBackTimerSubscription: Subscription = new Subscription
  gongdocCommitNbFromBackService_getCommitNbFromBack: Subscription = new Subscription

  // neccessary to unsubscribe
  ngOnDestroy() {
    // console.log("on destroy")
    this.checkGongdocCommitNbFromBackTimerSubscription.unsubscribe()
    this.gongdocCommitNbFromBackService_getCommitNbFromBack.unsubscribe()
  }

  ngOnInit(): void {

    // check loop for refresh from the back repo
    this.checkGongdocCommitNbFromBackTimerSubscription = this.checkGongdocCommitNbFromBackTimer.subscribe(
      currTime => {
        this.currTime = currTime

        this.gongdocCommitNbFromBackService_getCommitNbFromBack = this.gongdocCommitNbFromBackService.getCommitNbFromBack().subscribe(
          commitNbFromBack => {

            console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
            // condition for refresh
            if (this.lastCommitNbFromBack < commitNbFromBack) {

              // console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
              // console.log("last diagram id " + this.lastDiagramId + " new: " + id)
              // console.log("last drawn diagram id " + this.idOfDrawnClassDiagram + " new: " + id)
              this.pullGongdocAndDrawDiagram()
              this.lastCommitNbFromBack = commitNbFromBack
            }
          }
        )
      }
    )
  }

  /**
   * pullGongdocAndDrawDiagram refresh the front repo and calls redraw
   */
  pullGongdocAndDrawDiagram() {
    this.gongdocFrontRepoService.pull().subscribe(
      frontRepo => {
        this.gongdocFrontRepo = frontRepo
        console.log("gongdoc front repo pull returned")

        // find the selected classdiagram
        for (let packageDiagram of frontRepo.DiagramPackages_array) {
          if (packageDiagram.SelectedClassdiagram != undefined) {
            this.classdiagram = packageDiagram.SelectedClassdiagram
          }
        }

        console.log("drawing classdiagram" + this.classdiagram.Name)
        this.drawClassdiagram();
        this.paper!.setInteractivity(this.classdiagram.IsInDrawMode)
      }
    )
  }

  //
  // make a jointjs umlclass from a gong Classshape object
  //
  addGongStructShapeToGraph(gongStructShape: gongdoc.GongStructShapeDB): joint.shapes.uml.Class {

    let umlClassShape = newUmlClassShapeFromGongStructShape(gongStructShape, this.positionService, this.gongStructShapeService)
    umlClassShape.addTo(this.graph!);

    // add a backbone event handler to update the position
    umlClassShape.on('change:position', onClassshapeMove)

    // update the map that is used for find shapes
    this.Map_GongStructName_JointjsUMLClassShape.set(
      shapeIdentifierToShapeName(gongStructShape.Identifier), umlClassShape)

    return umlClassShape
  }

  addGongEnumShapeToGraph(gongEnumShape: gongdoc.GongEnumShapeDB): joint.shapes.uml.Class {

    let umlClassShape = newUmlClassShapeFromGongEnumShape(gongEnumShape, this.positionService, this.gongEnumShapeService)
    umlClassShape.addTo(this.graph!);

    // add a backbone event handler to update the position
    umlClassShape.on('change:position', onClassshapeMove)

    // update the map that is used for find shapes
    this.Map_GongStructName_JointjsUMLClassShape.set(
      shapeIdentifierToShapeName(gongEnumShape.Identifier), umlClassShape)

    return umlClassShape
  }

  //
  // make a jointjs umlclass from a gong Note object
  //
  addNoteToGraph(note: gongdoc.NoteShapeDB): joint.shapes.basic.Rect {

    let umlNote = newUmlNote(note, this.noteService)
    umlNote.addTo(this.graph!);
    return umlNote
  }

  drawClassdiagram(): void {

    const namespace = joint.shapes;

    //
    // one hairy stuff is the computation of the drawing size
    //
    // this is a work in progress
    //
    let diagramWidth = 1000
    if (this.classdiagram != undefined) {
      if (this.classdiagram.GongStructShapes != undefined) {
        diagramWidth = (this.classdiagram.GongStructShapes.length + 2) * 300
      }
    }


    //
    // a jointjs diagram is a Graph instance with a Paper instance
    //
    // the graph stores the logical elements
    // the paper stores the SVG elements
    this.graph = new joint.dia.Graph(
      {},
      { cellNamespace: this.namespace } // critical piece of code. 
    );

    let paperOptions: joint.dia.Paper.Options = {}
    paperOptions.el = document.getElementById('jointjs-holder')!
    paperOptions.model = this.graph
    paperOptions.width = diagramWidth
    paperOptions.height = 1000
    paperOptions.gridSize = 10
    paperOptions.drawGrid = false
    paperOptions.cellViewNamespace = namespace

    this.paper = new joint.dia.Paper(paperOptions)

    if (this.classdiagram == undefined) {
      return
    }

    if (this.classdiagram.IsInDrawMode) {
      this.paper.setInteractivity(true)
    } else {
      this.paper.setInteractivity(false)

      // when a user selects a shape, we would like the backend to know
      // for example, it might display in another component the list
      // of instances of the struct pointed
      this.paper.on('cell:pointerdown', informBackEndOfSelection)
    }

    // draw class shapes from the gong classshapes
    if (this.classdiagram?.GongStructShapes != undefined) {
      for (let classshape of this.classdiagram.GongStructShapes) {
        let umlClassShape = this.addGongStructShapeToGraph(classshape)
      }


      // draw links of the diagram shapes
      for (let classshape of this.classdiagram.GongStructShapes) {
        if (classshape.Links != undefined) {
          for (let linkDB of classshape.Links) {

            // does from & to shapes exists?
            //
            // a gong st

            let id: { receiver: string, fieldName: string }
            id = IdentifierToReceiverAndFieldName(linkDB.Identifier)
            var fromShape = this.Map_GongStructName_JointjsUMLClassShape.get(id.receiver)

            let toShapeName = IdentifierToStructname( linkDB.Fieldtypename)
            var toShape = this.Map_GongStructName_JointjsUMLClassShape.get(toShapeName)

            var strockWidth = 2

            let parts = linkDB.Identifier.split(".")

            let LinkEndlabel = parts[parts.length - 1]
            let distanceEndLabel = 0.75
            let linkTargetMuliplicity = linkDB.TargetMultiplicity
            let distanceTargetMultiplicity = 0.95
            let linkSourceMuliplicity = linkDB.SourceMultiplicity
            let distanceSourceMultiplicity = 0.10

            if (toShape == undefined) {
              // the destination shape is not in the diagram
              continue;
            }

            let xFrom = fromShape!.get('position')!.x
            let yFrom = fromShape!.get('position')!.y
            let xTo = toShape.get('position')!.x
            let yTo = toShape.get('position')!.y
            var vertices = [{ x: (xFrom + yTo) / 2, y: (yFrom + yTo) / 2 }]

            if (linkDB.Middlevertice != undefined) {
              vertices = [{ x: linkDB.Middlevertice.X, y: linkDB.Middlevertice.Y }]
            }

            if (fromShape != undefined && toShape != undefined) {
              var link = new joint.shapes.standard.Link({
                source: fromShape,
                target: toShape,
                vertices: vertices,
                attrs: {
                  line: {
                    stroke: '#3c4260',
                    strokeWidth: strockWidth,
                    vertexMarker: {
                      'type': 'circle',
                      'r': 3,
                      'stroke-width': 2,
                      'fill': 'white'
                    },
                    targetMarker: { // no arrow at the end
                      'type': 'path',
                      'd': 'M 10 -5 0 0 10 5 z'
                    },
                  },
                },
                labels: [
                  {
                    attrs: { text: { text: LinkEndlabel } },
                    position: {
                      offset: 15,
                      distance: distanceEndLabel
                    }
                  },
                  {
                    attrs: { text: { text: linkTargetMuliplicity } },
                    position: {
                      offset: 15,
                      distance: distanceTargetMultiplicity
                    }
                  },
                  {
                    attrs: { text: { text: linkSourceMuliplicity } },
                    position: {
                      offset: 15,
                      distance: distanceSourceMultiplicity
                    }
                  }
                ],
                // store relevant attributes for working when callback are invoked
                middleVertice: linkDB.Middlevertice,
                verticeService: this.verticeService,
              })

              // add a backbone event handler to update the position
              link.on('change:vertices', onLinkMove)

              this.Map_GongStructName_Joint_shapes_standard_Link.set(id.receiver + "." + linkDB.Name, link)

              link.addTo(this.graph);
            }
          }
        }
      }
    }

    // draw class shapes from the gong classshapes
    if (this.classdiagram?.GongEnumShapes != undefined) {
      for (const gongEnumShape of this.classdiagram.GongEnumShapes) {
        this.addGongEnumShapeToGraph(gongEnumShape);
      }
    }

    // draw notes from the gong notes
    if (this.classdiagram?.NoteShapes != undefined) {
      for (let noteShape of this.classdiagram.NoteShapes) {
        let rectShape = this.addNoteToGraph(noteShape)

        // add a backbone event handler to update the position
        rectShape.on('change:position', onNoteMove)


        if (noteShape.NoteShapeLinks != undefined) {
          for (let noteShapeLink of noteShape.NoteShapeLinks) {

            let fromShape = rectShape

            let xFrom = fromShape!.get('position')!.x
            let yFrom = fromShape!.get('position')!.y

            switch (noteShapeLink.Type) {

              case gongdoc.NoteShapeLinkType.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE: {
                var toShape = this.Map_GongStructName_JointjsUMLClassShape.get(noteShapeLink.Name)

                if (toShape == undefined) {
                  console.log("target shape not found: " + noteShapeLink.Name)
                  continue
                }
                let xTo = toShape!.get('position')!.x
                let yTo = toShape!.get('position')!.y
                var strockWidth = 1

                var link = new joint.shapes.standard.Link({
                  source: fromShape,
                  target: toShape,
                  attrs: {
                    line: {
                      stroke: '#3c4260',
                      strokeWidth: strockWidth,
                      strokeDasharray: '2 2',
                      targetMarker: NONE_TYPE,
                    },

                  },
                })

                link.addTo(this.graph);
                break
              }
              case gongdoc.NoteShapeLinkType.NOTE_SHAPE_LINK_TO_GONG_FIELD: {
                let toLink = this.Map_GongStructName_Joint_shapes_standard_Link.get(noteShapeLink.Name)
                if (toLink == undefined) {
                  break
                }

                var strockWidth = 1

                var link = new joint.shapes.standard.Link({
                  source: fromShape,
                  target: toLink,
                  attrs: {
                    line: {
                      stroke: '#3c4260',
                      strokeWidth: strockWidth,
                      strokeDasharray: '2 2',
                      targetMarker: NONE_TYPE,
                    },

                  },
                })

                link.addTo(this.graph);
                break
              }
            }
          }
        }
      }
    }

    // allow some observers to know what are the displayed structs
    if (this.classdiagram) {
      let classdiagramContext = new ClassdiagramContext()
      classdiagramContext.ClassdiagramID = this.classdiagram.ID
      ClassdiagramContextSubject.next(classdiagramContext)
    }
  }


}


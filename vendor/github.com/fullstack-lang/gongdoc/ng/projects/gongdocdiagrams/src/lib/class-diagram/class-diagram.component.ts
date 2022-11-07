import { Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import { ThemePalette } from '@angular/material/core';

import * as joint from 'jointjs';

import { ActivatedRoute, Router } from '@angular/router';

import * as gongdoc from 'gongdoc'
import * as gong from 'gong'

import { newUmlClassShape } from './newUmlClassShape'
import { ClassdiagramContextSubject, ClassdiagramContext } from '../diagram-displayed-gongstruct'
import { MatSlideToggleChange } from '@angular/material/slide-toggle';
import { newUmlNote } from './newUmlNote';

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
  lastDiagramId = 0
  currTime: number = 0

  /**
   * jointjs stuff
   */
  namespace = joint.shapes;
  private paper?: joint.dia.Paper
  private graph?: joint.dia.Graph

  /**
   * slider management
   */
  color: ThemePalette = 'primary';

  // the gong diagram of interest ot be drawn
  public classdiagram: gongdoc.ClassdiagramDB = new gongdoc.ClassdiagramDB

  /**
   * gongdoc manages 2 stacks: gongdoc and gong
   * 
   * gongdoc is for CRUDing the diagram elements (shapes, links, ...)
   * 
   * gong is for getting the elements of the go package (it has to be of an isntance of a gong model)
   * that are being modeled
   * 
   * instances in gongdoc refers to the gong model via the names of the elements.
   * a classshape instance modeling a "foo" gong struct will have a the value "foo" for the Structname field 
   */
  gongdocFrontRepo: gongdoc.FrontRepo = new gongdoc.FrontRepo
  gongFrontRepo: gong.FrontRepo = new gong.FrontRepo

  // map of Classhapes according to the joint.shapes.uml.Class
  // it is used to save the position of the classhape in the diagram (which only know the ids)
  // into the Classshape object (via its Position field)
  public Map_CellId_ClassshapeDB = new Map<string, gongdoc.ClassshapeDB>();

  // map for storing which gong struct have a classshape
  // it is important for drawing links between shapes
  public Map_GongStructName_JointjsUMLClassShape = new Map<string, joint.shapes.uml.Class>();

  // important for storing the relation from the jointjs link to the 
  // gongdoc Link. When the user saves the position of the vertice, this enables
  // the saving it into the Link object
  public Map_CellId_LinkDB = new Map<string, gongdoc.LinkDB>();

  // idOfDrawnClassDiagram stores the id of the drawn classdiagram. Usefull for knowing if
  // one has to redraw the diagram by comparaison with the route
  public idOfDrawnClassDiagram: number = 0

  // editable
  editable: boolean = false

  constructor(
    private route: ActivatedRoute,
    private router: Router,

    private positionService: gongdoc.PositionService,
    private noteService: gongdoc.NoteShapeService,
    private verticeService: gongdoc.VerticeService,
    private classshapeService: gongdoc.ClassshapeService, // for selection of the classshape

    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private gongdocCommitNbFromBackService: gongdoc.CommitNbFromBackService,

    private ClassdiagramService: gongdoc.ClassdiagramService,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  // Since this component is not reused when a new diagram is selected, there can be many
  // instances of the diagram and each instance will stay alive. For instance,
  // the instance will be in the control flow if an observable the component subscribes to emits an event.
  // Therefore, it is mandatory to manage subscriptions in order to unscribe them on the ngOnDestroy hook
  checkGongdocCommitNbFromBackTimerSubscription: Subscription = new Subscription
  gongdocCommitNbFromBackService_getCommitNbFromBack: Subscription = new Subscription

  subscriptionToDragAndDropEvent: Subscription = new Subscription
  subscriptionToRemoveFromDiagramEvent: Subscription = new Subscription

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

            const id = +this.route.snapshot.paramMap.get('id')!;

            // console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
            // console.log("last diagram id " + this.lastDiagramId + " new: " + id)
            // console.log("last drawn diagram id " + this.idOfDrawnClassDiagram + " new: " + id)

            // condition for refresh
            if (this.lastCommitNbFromBack < commitNbFromBack || this.lastDiagramId != id || this.idOfDrawnClassDiagram != id) {

              // console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
              // console.log("last diagram id " + this.lastDiagramId + " new: " + id)
              // console.log("last drawn diagram id " + this.idOfDrawnClassDiagram + " new: " + id)
              this.pullGongdocAndDrawDiagram()
              this.lastCommitNbFromBack = commitNbFromBack
              this.lastDiagramId = id
              this.idOfDrawnClassDiagram = id
            }
          }
        )
      }
    )
  }

  // onMove is called each time the shape is moved
  onClassshapeMove(umlClassShape: joint.shapes.uml.Class) {
    // console.log(umlClassShape.id, ':', umlClassShape.get('position'));

    let classhape = umlClassShape.attributes['classshape'] as gongdoc.ClassshapeDB
    let positionService = umlClassShape.attributes['positionService'] as gongdoc.PositionService
    let position = classhape.Position
    position!.X = umlClassShape.get('position')!.x
    position!.Y = umlClassShape.get('position')!.y

    positionService.updatePosition(position!).subscribe(
      position => {
        // console.log("position updated")
      }
    )
  }

  // onMove is called each time the shape is moved
  onNoteMove(umlNote: joint.shapes.standard.Rectangle) {


    let note = umlNote.attributes['note'] as gongdoc.NoteShapeDB
    let noteService = umlNote.attributes['noteService'] as gongdoc.NoteShapeService
    note.X = umlNote.get('position')!.x
    note.Y = umlNote.get('position')!.y
    noteService.updateNoteShape(note!).subscribe(
      note => {

      }
    )
    // console.log(note.Name, ':', umlNote.get('position'));
  }

  // onMove is called each time the shape is moved
  onLinkMove(standardLink: joint.shapes.standard.Link) {
    // console.log(standardLink.id, ':', standardLink.get('vertices'));

    let middleVertice = standardLink.attributes['middleVertice'] as gongdoc.VerticeDB
    let verticeService = standardLink.attributes['verticeService'] as gongdoc.VerticeService

    middleVertice!.X = standardLink.get('vertices')![0].x
    middleVertice!.Y = standardLink.get('vertices')![0].y

    verticeService.updateVertice(middleVertice!).subscribe(
      middleVertice => {
        // console.log("middleVertice updated")
      }
    )
  }

  //
  // make a jointjs umlclass from a gong Classshape object
  //
  addClassshapeToGraph(classshape: gongdoc.ClassshapeDB): joint.shapes.uml.Class {


    // back pointers: 
    // stores  as an attribute in the jointjs uml class shape :
    // - the position service to be able to update the position
    let umlClassShape = newUmlClassShape(classshape, this.positionService, this.classshapeService)

    // structRectangle.attributes = ['firstName: String']
    umlClassShape.addTo(this.graph!);

    this.Map_CellId_ClassshapeDB.set(umlClassShape.id.toString(), classshape)
    this.Map_GongStructName_JointjsUMLClassShape.set(classshape.ReferenceName, umlClassShape)

    return umlClassShape
  }

  //
  // make a jointjs umlclass from a gong Note object
  //
  addNoteToGraph(note: gongdoc.NoteShapeDB): joint.shapes.basic.Rect {

    //
    // creates the UML shape
    //

    // back pointers: 
    // stores  as an attribute in the jointjs uml class shape :
    // - the position service
    let umlNote = newUmlNote(note, this.noteService)

    // structRectangle.attributes = ['firstName: String']
    umlNote.addTo(this.graph!);

    // this.Map_CellId_NoteDB.set(umlNote.id.toString(), note)
    // this.Map_GongStructName_JointjsUMLNote.set(note.Structname, umlNote)

    return umlNote
  }
  //
  // turn gong instances into a jointjs diagram
  //
  drawClassdiagram(): void {
    // console.log("draw diagram")

    const namespace = joint.shapes;

    //
    // one hairy stuff is the computation of the drawing size
    //
    // this is a work in progress
    //
    let diagramWidth = 1000
    if (this.classdiagram != undefined) {
      if (this.classdiagram.Classshapes != undefined) {
        diagramWidth = (this.classdiagram.Classshapes.length + 2) * 300

        this.idOfDrawnClassDiagram = this.classdiagram.ID
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

    if (this.classdiagram.IsInDrawMode) {
      this.paper.setInteractivity(true)
    } else {
      this.paper.setInteractivity(false)
    }

    this.paper.on('cell:pointerdown',
      function (cellView, evt, x, y) {

        // console.log(cellView)

        // if paper is interactive, this means we are in dev mode
        // and we do not have to take click on shapes into account
        if (cellView.paper?.options.interactive) {
          return
        }

        let umlClassShape = cellView.model

        let classhape = umlClassShape.attributes['classshape'] as gongdoc.ClassshapeDB
        let classshapeService = umlClassShape.attributes['classshapeService'] as gongdoc.ClassshapeService

        // if selected object is not a classshape, move on
        if (classhape == undefined) {
          return
        }

        classhape.IsSelected = true
        classshapeService.updateClassshape(classhape).subscribe(
          classhape => {
            console.log("classhape updated")
          }
        )
      }
    )

    // draw class shapes from the gong classshapes
    if (this.classdiagram?.Classshapes != undefined) {
      for (let classshape of this.classdiagram.Classshapes) {
        let umlClassShape = this.addClassshapeToGraph(classshape)

        // add a backbone event handler to update the position
        umlClassShape.on('change:position', this.onClassshapeMove)
      }


      // draw links of the diagram shapes
      for (let classshape of this.classdiagram.Classshapes) {
        if (classshape.Links != undefined) {
          for (let linkDB of classshape.Links) {

            // does from & to shapes exists?
            //
            // a gong st
            var fromShape = this.Map_GongStructName_JointjsUMLClassShape.get(linkDB.Structname)
            var toShape = this.Map_GongStructName_JointjsUMLClassShape.get(linkDB.Fieldtypename)

            var strockWidth = 2
            let LinkEndlabel = linkDB.Fieldname
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
              link.on('change:vertices', this.onLinkMove)

              link.addTo(this.graph);

              // later, we need to save the diagram
              // 
              // algo is 
              // - for each cells of the diagram, 
              //      get its id & position
              //      find the original LinkDB and updates its position
              //
              // Because each cell has an unique id
              // we create a map of cell id to LinkDB
              this.Map_CellId_LinkDB.set(link.id.toString(), linkDB)
            }
          }
        }
      }
    }

    // draw notes from the gong notes
    if (this.classdiagram?.Notes != undefined) {
      for (let note of this.classdiagram.Notes) {
        let umlNote = this.addNoteToGraph(note)

        // add a backbone event handler to update the position
        umlNote.on('change:position', this.onNoteMove)
      }
    }

    // allow some observers to know what are the displayed structs
    if (this.classdiagram) {
      let classdiagramContext = new ClassdiagramContext()
      classdiagramContext.ClassdiagramID = this.classdiagram.ID
      ClassdiagramContextSubject.next(classdiagramContext)
    }
  }

  /**
   * saving the classdiagram
   * 
   * the challenge is to update the positions of classshapes and vertices
   */

  pullGongdocAndDrawDiagram() {
    this.gongdocFrontRepoService.pull().subscribe(
      frontRepo => {
        this.gongdocFrontRepo = frontRepo
        // console.log("gongdoc front repo pull returned")

        const id = +this.route.snapshot.paramMap.get('id')!
        this.editable = false
        this.classdiagram = frontRepo.Classdiagrams.get(id)!

        // intercept click on shapes when in production mode
        if (this.classdiagram == undefined) {
          console.log("unkwnown classdiagram id : " + id)
        } else {

          this.drawClassdiagram();
          this.paper!.setInteractivity(this.classdiagram.IsInDrawMode)

        }
      }
    )
  }

  public toggle(event: MatSlideToggleChange) {
    console.log('toggle', event.checked);

    if (!event.checked) {
      this.paper!.setInteractivity(false)
      this.classdiagram.IsInDrawMode = false
      this.ClassdiagramService.updateClassdiagram(this.classdiagram).subscribe(
        classdiagram => {
          console.log("classdiagram edition mode set to PROD")
        }
      )
    } else {
      this.paper!.setInteractivity(true)
      this.classdiagram.IsInDrawMode = true
      this.ClassdiagramService.updateClassdiagram(this.classdiagram).subscribe(
        classdiagram => {
          console.log("classdiagram edition mode set to DEV")
        }
      )
    }
  }
}


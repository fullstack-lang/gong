import { Component, OnInit } from '@angular/core';
import { Observable, combineLatest, timer, Subscription } from 'rxjs';
import * as joint from 'jointjs';

import { ActivatedRoute, Router } from '@angular/router';

import * as gongdoc from 'gongdoc'

import { state } from '@angular/animations';

@Component({
  selector: 'app-umlsc-diagram',
  templateUrl: './umlsc-diagram.component.html',
  styleUrls: ['./umlsc-diagram.component.css']
})
export class UmlscDiagramComponent implements OnInit {

  /**
   * the uml state diagram component is refreshed both by direct input when the user moves vertices or positions
   * otherwise, modification are gotten from the back repo 
   * 
   * the checkCommitNbTimer polls the commit number of the back repo
   * if the commit number has increased, it pulls the front repo and redraw the diagram
   */
  checkGongdocCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  lastDiagramId = 0
  currTime: number = 0

  namespace = joint.shapes

  private paper?: joint.dia.Paper
  private graph?: joint.dia.Graph

  public founded = false

  // the diagram of interest
  public stateChartDiagram?: gongdoc.UmlscDB;

  // map states of the diagram
  mapStateDBIDStateDBs = new Map<number, gongdoc.UmlStateDB>()

  // map of States according to the joint.shapes.uml.State
  // it is used to save the diagram (which only know the ids)
  public MapJointjsIdsStates = new Map<string, gongdoc.UmlStateDB>();
  public MapStateDBIDJointjsStateID = new Map<number, string>()

  public MapNamesStates = new Map<string, gongdoc.UmlStateDB>()

  // front repo, that will be used to access backend elements
  gongdocFrontRepo?: gongdoc.FrontRepo

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private UmlscService: gongdoc.UmlscService,
    private UmlStateService: gongdoc.UmlStateService,
    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private GongdocCommandService: gongdoc.GongdocCommandService,
    private gongdocCommitNbService: gongdoc.CommitNbService,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    // this.router.routeReuseStrategy.shouldReuseRoute = function () {
    //   return false;
    // }
  }

  // if true the save button will appear
  public savebutton: boolean = false

  // neccessary to unsubscribe
  subscriptionToCheckCommitTimer?: Subscription
  ngOnDestroy() {
    console.log("on destroy")
    this.subscriptionToCheckCommitTimer?.unsubscribe()
  }

  ngOnInit(): void {
    // wait for all fetch to combine
    const id = +this.route.snapshot.paramMap.get('id')!
    if (this.route.snapshot.paramMap.has('savebutton')) {
      this.savebutton = (this.route.snapshot.paramMap.get('savebutton') == "true")
    }

    //
    // init of the this.stateChartDiagram
    // in order for the front html template to compute stateChartDiagram.Name
    this.gongdocFrontRepoService.pull().subscribe(
      gongdocFrontRepo => {
        this.gongdocFrontRepo = gongdocFrontRepo
        console.log("gongdoc front repo pull returned")

        const id = +this.route.snapshot.paramMap.get('id')!
        this.stateChartDiagram = this.gongdocFrontRepo.Umlscs.get(id)
      }
    )

    this.subscriptionToCheckCommitTimer = this.checkGongdocCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        const id = +this.route.snapshot.paramMap.get('id')!

        this.gongdocCommitNbService.getCommitNb().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb || this.lastDiagramId != id) {

              console.log("id " + id + ": last commit nb " + this.lastCommitNb + " new: " + commitNb)
              console.log("id " + id + ": last diagram id " + this.lastDiagramId + " new: " + id)

              this.lastCommitNb = commitNb
              this.lastDiagramId = id
              this.pullGongdocAndDrawDiagram()
            }
          }
        )
      }
    )
  }

  pullGongdocAndDrawDiagram() {
    this.gongdocFrontRepoService.pull().subscribe(
      gongdocFrontRepo => {
        this.gongdocFrontRepo = gongdocFrontRepo
        console.log("gongdoc front repo pull returned")

        const id = +this.route.snapshot.paramMap.get('id')!
        this.stateChartDiagram = this.gongdocFrontRepo.Umlscs.get(id)

        this.drawStateChartDiagram();
      }
    )
  }

  drawStateChartDiagram(): void {
    const namespace = joint.shapes;
    this.graph = new joint.dia.Graph(
      {},
      { cellNamespace: this.namespace } // critical piece of code. 
    );

    // console.log("state chart diagram name " + this.stateChartDiagram.Name)
    // let htmlDocument = document.getElementById(this.stateChartDiagram.Name)
    // // console.log("html document " + htmlDocument.id)

    // let len = document.getElementsByTagName('*').length
    // for (let idx = 0; idx < len; idx++) {
    //   let element = document.getElementsByTagName('*').item(idx)
    //   console.log("element id :" + element.id)
    // }

    this.paper = new joint.dia.Paper(
      {
        el: document.getElementById(this.stateChartDiagram!.Name)!,
        model: this.graph,
        width: 400,
        // height: window.innerHeight,
        height: 800,
        gridSize: 10,
        drawGrid: false,
        cellViewNamespace: namespace
      }
    )

    // redraw all states
    for (let stateDB of this.stateChartDiagram!.States!) {

      var color = 'rgba(48, 208, 198, 0.1)'

      if (stateDB.Name == this.stateChartDiagram!.Activestate) {
        color = 'rgba(248, 0, 0, 0.3)'
      }

      var umlState = new joint.shapes.uml.State(
        {
          position: {
            x: stateDB.X,
            y: stateDB.Y
          },
          size: { width: 240, height: 40 },
          name: [stateDB.Name],
          attrs: {
            '.uml-state-body': {
              fill: color,
            },
          }
        })

      umlState.addTo(this.graph)

      var id: any;
      id = umlState.id;
      var idstring: string
      idstring = id;
      this.MapJointjsIdsStates.set(idstring, stateDB)
      this.MapStateDBIDJointjsStateID.set(stateDB.ID, idstring)
    }
  }

  // later, we should do this with
  //
  // Listening for changes of the position to a single element
  // element1.on('change:position', function(element1, position) {
  //   alert('element1 moved to ' + position.x + ',' + position.y);
  // });
  saveClassdiagram(): void {
    console.log("save diagram")

    // parse shapes positions
    var cells = this.graph!.getCells()
    console.log(cells.length)

    cells.forEach(
      (cell: joint.dia.Cell) => {
        // ugly hack because cell.id is considered a Dimension by the ts compiler
        // vive golang
        var cellId: any
        cellId = cell.id;
        if (this.MapJointjsIdsStates.get(cellId) != undefined) {

          // retrieve the shape.
          var stateDB = this.MapJointjsIdsStates.get(cellId)

          let position = cell.attributes["position"]

          stateDB!.X = position["x"]
          stateDB!.Y = position["y"]

          // update position to DB
          this.UmlStateService.updateUmlState(stateDB!).subscribe(
            state => {
              console.log("state updated")
            }
          )
        }
      }
    )

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo!.GongdocCommands.get(1)
    gongdocCommandSingloton!.Command = gongdoc.GongdocCommandType.MARSHALL_DIAGRAM
    gongdocCommandSingloton!.DiagramName = this.stateChartDiagram!.Name
    gongdocCommandSingloton!.Date = Date.now().toString()

    this.GongdocCommandService.updateGongdocCommand(gongdocCommandSingloton!).subscribe(
      GongdocCommand => {
        console.log("GongdocCommand updated")
      }
    )
  }
}


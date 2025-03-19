import { Component, OnInit, Input, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';

import * as sim from '../../../../sim/src/public-api'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { Observable, combineLatest, timer } from 'rxjs'

import { MatToolbarModule } from '@angular/material/toolbar'
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@Component({
  selector: 'lib-sim-specific',
  imports: [
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    FormsModule,
    CommonModule,
    AngularSplitModule,
  ],
  templateUrl: './sim-specific.component.html',
  styleUrl: './sim-specific.component.css'
})
export class SimSpecificComponent {

  @Input() Name: string = ""

  public engine: sim.Engine = new sim.Engine
  public engineID: number = 0

  lastEvent: string = ""
  lastEventAgent: string = ""

  nextEventAgent: string = ""
  nextEventName: string = ""
  nextEventTime: string = ""

  engineEventNumber: number = 0

  // animation fof the simulation
  speed = 36;
  clientState = "PAUSED";

  frontRepo = new sim.FrontRepo

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  diagramIDForSamocStates: number = 0
  diagramIDForNatoStates: number = 0

  // engineUpdated is the call function
  @Input() engineUpdatedCallbackFunction?: (updateDisplay: boolean) => void
  @Input() simulationTitle: string = ""

  // UpdateDisplay is set to true if the simulation is running or if 
  // the user has provided an input such as AdvanceTillStateChange
  UpdateDisplay = true

  commandSingloton: sim.Command = new sim.Command

  constructor(
    private frontRepoService: sim.FrontRepoService,
    private engineService: sim.EngineService,
    private commandService: sim.CommandService,
    private changeDetectorRef: ChangeDetectorRef,

    private router: Router) {
  }

  ngOnInit(): void {
    this.frontRepoService.connectToWebSocket(this.Name).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        let engines = this.frontRepo.getFrontArray<sim.Engine>(sim.Engine.GONGSTRUCT_NAME)
        // console.log("Nb of Engines is ", engines.length)

        this.engine = engines[0]
        console.log("date", this.engine.CurrentTime)

        let commands = this.frontRepo.getFrontArray<sim.Command>(sim.Command.GONGSTRUCT_NAME)
        // console.log("Nb of Commands is ", commands.length)

        this.commandSingloton = commands[0]

        // this is the callback function from the generic engine to the specific engine 
        if (this.engineUpdatedCallbackFunction) {

          if (this.engine.State == sim.EngineState.RUNNING) {
            this.UpdateDisplay = true
          }
          this.engineUpdatedCallbackFunction(this.UpdateDisplay)

          // reset the need to display updates
          this.UpdateDisplay = false
        }

        // Force change detection
        this.changeDetectorRef.detectChanges();
        console.log("change detection")
      }
    )
  }

  fireEventTillStateChange(): void {
    this.commandSingloton.Command = sim.CommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log("FIRE_EVENT_TILL_STATES_CHANGE sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }

  fireEvent(): void {
    this.commandSingloton.Command = sim.CommandType.COMMAND_FIRE_NEXT_EVENT
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log("FIRCOMMAND_FIRE_NEXT_EVENT sent to the backRepo")
      }
    )
  }

  reset(): void {
    this.UpdateDisplay = true

    this.commandSingloton.Command = sim.CommandType.COMMAND_RESET
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log("RESCOMMAND_RESET sent to the backRepo")
      }
    )
  }

  play(): void {
    this.UpdateDisplay = true

    this.commandSingloton.Command = sim.CommandType.COMMAND_PLAY
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log("PLAY sent to the backRepo")
      }
    )
  }

  pause(): void {
    this.commandSingloton.Command = sim.CommandType.COMMAND_PAUSE
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log("PAUSE sent to the backRepo")
      }
    )
  }

  increaseSpeed100percent(): void {
    this.commandSingloton.Command = sim.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log(sim.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS, "sent")
      }
    )

    this.UpdateDisplay = true
  }

  decreaseSpeed50percent(): void {
    this.commandSingloton.Command = sim.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log(sim.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS, "sent")
      }
    )
    this.UpdateDisplay = true
  }

  publishAgentStates(): void {
    this.UpdateDisplay = true
  }

  publishAdvance10Minutes(): void {
    this.commandSingloton.Command = sim.CommandType.COMMAND_ADVANCE_10_MIN
    this.commandSingloton.CommandDate = Date.now().toString()
    this.commandService.updateFront(this.commandSingloton, this.Name).subscribe(
      command => {
        console.log("ADCOMMAND_ADVANCE_10_MIN sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }

}

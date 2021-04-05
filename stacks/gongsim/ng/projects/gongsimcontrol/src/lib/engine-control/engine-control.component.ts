import { Component, OnInit, Input } from '@angular/core';

import * as gongsim from 'gongsim'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { Observable, combineLatest, timer } from 'rxjs'
import { GongsimCommandTypeList } from 'gongsim';


@Component({
  selector: 'lib-engine-control',
  templateUrl: './engine-control.component.html',
  styleUrls: ['./engine-control.component.css']
})
export class EngineControlComponent implements OnInit {

  public engine: gongsim.EngineDB
  public engineID: number

  lastEvent: string;
  lastEventAgent: string;

  nextEventAgent: string;
  nextEventName: string;
  nextEventTime: string;

  engineEventNumber: number;

  // animation fof the simulation
  speed = 36;
  clientState = "PAUSED";

  currTime: number;
  obsTimer: Observable<number> = timer(1000, 1000);

  diagramIDForSamocStates: number
  diagramIDForNatoStates: number

  // engineUpdated is the call function
  @Input() engineUpdatedCallbackFunction: (updateDisplay: boolean) => void;
  @Input() simulationTitle: string;

  // UpdateDisplay is set to true if the simulation is running or if 
  // the user has provided an input such as AdvanceTillStateChange
  UpdateDisplay = true

  gongsimCommandSingloton: gongsim.GongsimCommandDB

  constructor(
    private engineService: gongsim.EngineService,
    private gongsimCommandService: gongsim.GongsimCommandService,
    private router: Router) {
  }


  ngOnInit(): void {

    // get the current engine
    this.engineService.getEngines().subscribe(
      engines => {
        engines.forEach(
          engine => {
            this.engine = engine
            this.engineID = this.engine.ID
          }
        )
      }
    )

    // get the gonogsim command singloton
    this.gongsimCommandService.getGongsimCommands().subscribe(
      gongsimCommands => {

        if (gongsimCommands.length != 1) {
          console.log("problem with fetching gongsim command singloton")
        }

        gongsimCommands.forEach(
          gongsimCommand => {
            this.gongsimCommandSingloton = gongsimCommand
          }
        )
      }
    )

    this.obsTimer.subscribe(
      currTime => {
        this.currTime = currTime

        if (this.engine) {
          this.engineService.getEngine(this.engineID).subscribe(
            engine => {
              this.engine = engine

              // this is the callback function from the generic engien to the specific engine 
              if (this.engineUpdatedCallbackFunction) {

                if (this.engine.State == gongsim.EngineState.RUNNING) {
                  this.UpdateDisplay = true
                }
                this.engineUpdatedCallbackFunction(this.UpdateDisplay)

                // reset the need to display updates
                this.UpdateDisplay = false
              }
            }
          )
        }
      }
    )
  }

  fireEventTillStateChange(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("FIRE_EVENT_TILL_STATES_CHANGE sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }

  fireEvent(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_FIRE_NEXT_EVENT
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("FIRCOMMAND_FIRE_NEXT_EVENT sent to the backRepo")
      }
    )
  }

  reset(): void {
    this.UpdateDisplay = true

    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_RESET
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("RESCOMMAND_RESET sent to the backRepo")
      }
    )
  }

  play(): void {
    this.UpdateDisplay = true

    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_PLAY
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("PLAY sent to the backRepo")
      }
    )
  }

  pause(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_PAUSE
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("PAUSE sent to the backRepo")
      }
    )
  }

  increaseSpeed100percent(): void {
    this.gongsimCommandSingloton.SpeedCommandType = gongsim.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS
    this.gongsimCommandSingloton.DateSpeedCommand = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("INCCOMMAND_INCREASE_SPEED_100_PERCENTS sent to the backRepo")
      }
    )

    this.UpdateDisplay = true
  }

  decreaseSpeed50percent(): void {
    this.gongsimCommandSingloton.SpeedCommandType = gongsim.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS
    this.gongsimCommandSingloton.DateSpeedCommand = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("DECCOMMAND_DECREASE_SPEED_50_PERCENTS sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }

  publishAgentStates(): void {
    this.UpdateDisplay = true
  }

  publishAdvance10Minutes(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_ADVANCE_10_MIN
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateGongsimCommand(this.gongsimCommandSingloton).subscribe(
      gongsimCommand => {
        console.log("ADCOMMAND_ADVANCE_10_MIN sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }
}

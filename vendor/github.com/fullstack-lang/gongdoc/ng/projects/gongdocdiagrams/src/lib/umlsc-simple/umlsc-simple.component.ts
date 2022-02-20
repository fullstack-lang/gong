import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState, ActivatedRoute, ParamMap } from '@angular/router';

import { UmlscDB } from 'gongdoc';
import { UmlscService } from 'gongdoc'
import { UmlStateDB, UmlStateService } from 'gongdoc'


// generated table component
@Component({
  selector: 'app-umlsc-simple',
  templateUrl: './umlsc-simple.component.html',
  styleUrls: ['./umlsc-simple.component.css']
})
export class UmlscSimpleTableComponent implements OnInit {

  // the data source for the table
  umlscs: UmlscDB[] = []

  // the selected umlsc 
  umlsc?: UmlscDB

  states: UmlStateDB[] = []

  @Input() ID: number = 0 // ID of the caller when component called from struct in reverse relation
  @Input() struct: string = "" // struct with pointer to Umlsc
  @Input() field: string = "" // field to display

  displayedColumns: string[] = ['ID', 'Name', 'Delete'];

  constructor(
    private umlscService: UmlscService,
    private umlStateService: UmlStateService,

    private router: Router,
    private route: ActivatedRoute,
  ) {
    // observable for changes in structs
    this.umlscService.UmlscServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getUmlscs()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getUmlscs()

    this.umlStateService.getUmlStates().subscribe(
      states => {
        this.states = states
      }
    )
  }

  getUmlscs(): void {
    if (this.ID == null) {
      this.umlscService.getUmlscs().subscribe(
        umlscs => {
          this.umlscs = umlscs
        }
      )
    }

  }

  // newUmlsc initiate a new umlsc
  // create a new Umlsc objet
  newUmlsc() {
  }

  deleteUmlsc(umlscID: number, umlsc: UmlscDB) {
    // la liste des umlscs est amputée du umlsc avant le delete afin
    // de mettre à jour l'IHM
    this.umlscs = this.umlscs.filter(h => h !== umlsc);

    this.umlscService.deleteUmlsc(umlscID).subscribe();
  }

  editUmlsc(umlscID: number, umlsc: UmlscDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(umlscID: number) {

    this.umlscService.getUmlsc(umlscID).subscribe(
      umlsc => {
        this.umlsc = umlsc
      }
    )

    this.router.navigate([{
      outlets: {
        diagrameditor: ["umlsc-detail", umlscID, { savebutton: "true" }]
      }
    }]);
  }

  switchState() {

    if (this.umlsc) {
      // parse all states
      var activeStateChanged = false
      this.states.forEach(
        state => {
          if (this.umlsc!.Activestate != state.Name && !activeStateChanged) {
            this.umlsc!.Activestate = state.Name
            activeStateChanged = true
            this.umlscService.updateUmlsc(this.umlsc!).subscribe(
              umlsc => {
                console.log("state diagram updated")
              }
            )
          }
        }
      )
    }
  }
}

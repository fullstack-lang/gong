import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';

import { ClassdiagramDB } from 'gongdoc';
import { ClassdiagramService } from 'gongdoc'


// generated table component
@Component({
  selector: 'app-classdiagrams-simple-table',
  templateUrl: './classdiagrams-simple-table.component.html',
  styleUrls: ['./classdiagrams-simple-table.component.css']
})
export class ClassdiagramsSimpleTableComponent implements OnInit {

  // the data source for the table
  classdiagrams: ClassdiagramDB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to Classdiagram
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', 'Name'];

  constructor(
    private classdiagramService: ClassdiagramService,

    private router: Router,
  ) {
    // observable for changes in structs
    this.classdiagramService.ClassdiagramServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getClassdiagrams()
        }
      }
    )
  }

  ngOnInit(): void {
    this.getClassdiagrams()
  }

  getClassdiagrams(): void {
    if (this.ID == null) {
      this.classdiagramService.getClassdiagrams().subscribe(
        Classdiagrams => {
          this.classdiagrams = Classdiagrams;
        }
      )
    }
  
  }

  // newClassdiagram initiate a new classdiagram
  // create a new Classdiagram objet
  newClassdiagram() {
  }

  // set editor outlet
  setEditorRouterOutlet(classdiagramID: number) {

    console.log("setEditorRouterOutlet " + classdiagramID)

    this.router.navigate([{
      outlets: {
        diagrameditor: ["classdiagram-detail", classdiagramID]
      }
    }]).catch(
      reason => {
        console.log(reason)
      }
    );
  }
}

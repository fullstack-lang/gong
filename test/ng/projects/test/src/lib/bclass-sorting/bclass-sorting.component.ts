import { Component, OnInit, Inject, Optional } from '@angular/core';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
@Component({
  selector: 'lib-bclass-sorting',
  templateUrl: './bclass-sorting.component.html',
  styleUrls: ['./bclass-sorting.component.css']
})
export class BclassSortingComponent implements OnInit  {

  movies = [
    'Episode I - The Phantom Menace',
    'Episode II - Attack of the Clones',
    'Episode III - Revenge of the Sith',
    'Episode IV - A New Hope',
    'Episode V - The Empire Strikes Back',
    'Episode VI - Return of the Jedi',
    'Episode VII - The Force Awakens',
    'Episode VIII - The Last Jedi',
    'Episode IX – The Rise of Skywalker'
  ];

  frontRepo: FrontRepo

  // array of Bclass instances that are in the association
  associatedBclasss = new Array<BclassDB>();

  constructor(
    private bclassService: BclassService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of bclass instances
    public dialogRef: MatDialogRef<BclassSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getBclasss()
  }

  getBclasss(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        if (this.dialogData != undefined) {
          this.frontRepo.Bclasss_array.forEach(
            bclass => {
              let ID = this.dialogData.ID
              let revPointer = bclass[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.associatedBclasss.push(bclass)
              }
            }
          )
        }
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedBclasss, event.previousIndex, event.currentIndex);
  }

}

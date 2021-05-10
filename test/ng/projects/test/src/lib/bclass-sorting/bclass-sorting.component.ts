import { Component, OnInit, Inject, Optional } from '@angular/core';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { BclassDB } from '../bclass-db'
import { BclassService } from '../bclass.service'

import { FrontRepoService, FrontRepo, NullInt64 } from '../front-repo.service'
@Component({
  selector: 'lib-bclass-sorting',
  templateUrl: './bclass-sorting.component.html',
  styleUrls: ['./bclass-sorting.component.css']
})
export class BclassSortingComponent implements OnInit {

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

        let index = 0
        for (let bclass of this.frontRepo.Bclasss_array) {
          let ID = this.dialogData.ID
          let revPointer = bclass[this.dialogData.ReversePointer]
          if (revPointer.Int64 == ID) {
            if (bclass.Aclass_AnarrayofbDBID_Order == undefined) {
              bclass.Aclass_AnarrayofbDBID_Order = new NullInt64
              bclass.Aclass_AnarrayofbDBID_Order.Valid = true
              bclass.Aclass_AnarrayofbDBID_Order.Int64 = index++
            }
            this.associatedBclasss.push(bclass)
          }
        }

        // sort associated bclass according to order
        this.associatedBclasss.sort((t1, t2) => {
          if (t1.Aclass_AnarrayofbDBID_Order.Int64 > t2.Aclass_AnarrayofbDBID_Order.Int64) {
            return 1;
          }
          if (t1.Aclass_AnarrayofbDBID_Order.Int64 < t2.Aclass_AnarrayofbDBID_Order.Int64) {
            return -1;
          }
          return 0;
        });

        console.log("front repo pull returned")
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedBclasss, event.previousIndex, event.currentIndex);

    // set the order of Bclass instances
    let index = 0
    for (let bclass of this.associatedBclasss) {
      bclass.Aclass_AnarrayofbDBID_Order.Int64 = index++
    }
    console.log("after drop")
  }

  save() {

    this.associatedBclasss.forEach(
      bclass => {
        this.bclassService.updateBclass(bclass)
          .subscribe(bclass => {
            this.bclassService.BclassServiceChanged.next("update")
            console.log("bclass saved")
          });
      }
    )

    this.dialogRef.close('Sorting of Aclass.Anarrayofb done');
  }
}

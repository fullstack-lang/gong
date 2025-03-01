import { Component, OnInit } from '@angular/core';

import * as test from '../../../test/src/public-api'
import { MatTableModule } from '@angular/material/table';

@Component({
  selector: 'lib-testspecific',
  standalone: true,
  imports: [
    MatTableModule
  ],
  template: `

<table mat-table [dataSource]="frontRepo.array_Astructs" class="mat-elevation-z8">

    <!-- ID Column -->
    <ng-container matColumnDef="ID">
      <th mat-header-cell *matHeaderCellDef> ID </th>
      <td mat-cell *matCellDef="let astruct"> {{astruct.ID}} </td>
    </ng-container>

    <!-- Name Column -->
    <ng-container matColumnDef="Name">
      <th mat-header-cell *matHeaderCellDef> Name </th>
      <td mat-cell *matCellDef="let astruct"> {{astruct.Name}} </td>
    </ng-container>


    <!-- Header Row -->
    <tr mat-header-row *matHeaderRowDef="['ID', 'Name']"></tr>
    <!-- Data Row -->
    <tr mat-row *matRowDef="let row; columns: ['ID', 'Name'];"></tr>

</table>
  `,
  styles: ``
})
export class TestspecificComponent implements OnInit {

  array_Astructs = new Array<test.Astruct>() // array of front instances

  StacksNames = test.StacksNames
  public frontRepo = new (test.FrontRepo)

  constructor(
    private frontRepoService: test.FrontRepoService,
  ) { }

  ngOnInit(): void {

    this.frontRepoService.connectToWebSocket(this.StacksNames.Test).subscribe(
      gongtablesFrontRepo => {
        console.log("front repo updated")
        this.frontRepo = gongtablesFrontRepo
      }
    )

    console.log("ngOnInit")
  }

}

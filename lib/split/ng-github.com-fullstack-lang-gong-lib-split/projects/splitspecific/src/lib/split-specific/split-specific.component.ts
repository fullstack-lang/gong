import { Component, Input, OnInit } from '@angular/core';

import * as split from '../../../../split/src/public-api'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

import { AngularSplitModule } from 'angular-split';


@Component({
  selector: 'lib-split-specific',
  imports: [
    MatRadioModule,
    CommonModule,
    FormsModule,

    AngularSplitModule
  ],
  templateUrl: './split-specific.component.html',
  styleUrl: './split-specific.component.css'
})
export class SplitSpecificComponent implements OnInit {
    @Input() GONG__StackPath: string = ""

    public frontRepo?: split.FrontRepo;

    view = ""

    constructor(
      private frontRepoService: split.FrontRepoService,
    ) { }

    ngOnInit(): void {
      console.log("ngOnInit");
  
      this.frontRepoService.connectToWebSocket(this.GONG__StackPath).subscribe({
        next: (frontRepo) => {
          this.frontRepo = frontRepo;
  
          if (this.frontRepo.array_Views.length > 0) {

            this.frontRepo.array_Views.sort((a: split.View, b : split.View) => {
              return a.ID - b.ID
            })

            this.view = this.frontRepo.array_Views[0].Name
          }

        }
      }
      )
    }
}
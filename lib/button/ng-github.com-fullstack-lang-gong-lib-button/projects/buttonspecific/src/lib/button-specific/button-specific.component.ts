import { Component, Input, OnInit } from '@angular/core';

import * as button from '../../../../button/src/public-api'

import { AngularSplitModule } from 'angular-split';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon'

@Component({
  selector: 'lib-button-specific',
  imports: [
    AngularSplitModule,

    MatButtonModule,
    MatIconModule,
  ],
  templateUrl: './button-specific.component.html',
  styleUrl: './button-specific.component.css'
})
export class ButtonSpecificComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  StacksNames = button.StacksNames;
  public frontRepo?: button.FrontRepo;
  splitAreaSize = 0

  layout: button.Layout | undefined

  constructor(
    private frontRepoService: button.FrontRepoService,
    private buttonService: button.ButtonService,
  ) { }

  formatLabel(value: number): string {
    if (value >= 1000) {
      return Math.round(value / 1000) + 'k';
    }

    return `${value}`;
  }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.GONG__StackPath).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        for (let layout_ of this.frontRepo.array_Layouts) {
          this.layout = layout_
        }

        this.splitAreaSize = 100.0 / this.frontRepo.array_Groups.length
      }
    }
    )
  }

  onClick(button: button.Button) {
    this.buttonService.updateFront(button,this.GONG__StackPath).subscribe(
      () => {
        console.log("checkbox updated")
      }
    )
  }

}

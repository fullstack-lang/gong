import { Component, Input } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as slider from '../../../../slider/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { MatRadioChange, MatRadioModule } from '@angular/material/radio';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon'


@Component({
  selector: 'lib-slider-specific',
  imports: [
    AngularSplitModule,
    MatSliderModule,
    MatRadioModule,
    MatCardModule,
    MatCheckboxModule,
    FormsModule,
    MatFormFieldModule,
    MatIconModule,
    MatGridListModule,
  ],
  templateUrl: './slider-specific.component.html',
  styleUrl: './slider-specific.component.css'
})
export class SliderSpecificComponent {

  @Input() Name: string = ""

  rowHeight: string = "30px"

  StacksNames = slider.StacksNames;
  public frontRepo?: slider.FrontRepo;
  splitAreaSize = 0

  layout: slider.Layout | undefined

  constructor(
    private frontRepoService: slider.FrontRepoService,
    private sliderService: slider.SliderService,
    private checkboxService: slider.CheckboxService,
  ) { }

  formatLabel(value: number): string {
    if (value >= 1000) {
      return Math.round(value / 1000) + 'k';
    }

    return `${value}`;
  }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
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

  input($event: Event, slider: slider.Slider) {
    this.sliderService.updateFront(slider, this.Name).subscribe(
      () => {
        console.log("slider updated")
      }
    )
  }

  inputMatRadio($event: MatRadioChange, checkbox: slider.Checkbox) {
    this.checkboxService.updateFront(checkbox, this.Name).subscribe(
      () => {
        console.log("checkbox updated")
      }
    )
  }


}

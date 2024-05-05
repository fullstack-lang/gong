import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import { CommonModule } from '@angular/common';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatCardModule } from '@angular/material/card'
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatRadioModule } from '@angular/material/radio';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

import { FormsModule } from '@angular/forms';

import { AngularSplitModule } from 'angular-split';

import * as test from '../../projects/test/src/public-api'

import { TestspecificComponent } from '../../projects/testspecific/src/public-api'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [

    CommonModule,

    MatSliderModule,
    MatSelectModule,
    MatDatepickerModule,
    MatFormFieldModule,
    MatInputModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatListModule,
    MatCardModule,
    MatTooltipModule,
    MatRadioModule,
    MatSlideToggleModule,

    FormsModule,

    AngularSplitModule,

    // gongtreespecific.GongtreespecificModule,
    // gongtablespecific.GongtablespecificModule,

    TestspecificComponent

  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  test = 'Test'
  probe = 'Test Data/Model'
  view = this.test

  views: string[] = [this.test, this.probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "test"
  StackType = test.StackType

  // TableExtraPathEnum = gongtable.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}

import { NgModule } from '@angular/core';
import { GongsimcontrolComponent } from './gongsimcontrol.component';

import { GongsimModule } from 'gongsim';
import { EngineControlComponent } from './engine-control/engine-control.component'

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

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

import { FormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    GongsimcontrolComponent,
    EngineControlComponent
  ],
  imports: [
    GongsimModule,

    MatSliderModule,
    MatSelectModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatCardModule,
    MatTooltipModule,
    MatRadioModule,

    BrowserAnimationsModule,
    MatListModule,

    FormsModule
  ],
  exports: [
    GongsimcontrolComponent,
    EngineControlComponent
  ],
})
export class GongsimcontrolModule { }

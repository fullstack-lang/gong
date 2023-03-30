import { NgModule } from '@angular/core';

import { GongspecificComponent } from './gongspecific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongModule } from 'gong'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    GongspecificComponent,
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongModule,
  ],
  exports: [
    GongspecificComponent,
    DataModelPanelComponent
  ]
})
export class GongspecificModule { }

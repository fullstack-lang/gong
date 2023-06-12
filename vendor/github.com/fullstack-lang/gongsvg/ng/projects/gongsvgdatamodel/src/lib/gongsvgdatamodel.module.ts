import { NgModule } from '@angular/core';

import { GongsvgdatamodelComponent } from './gongsvgdatamodel.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { AngularSplitModule } from 'angular-split';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongsvgModule } from 'gongsvg';

@NgModule({
  declarations: [
    GongsvgdatamodelComponent,
    DataModelPanelComponent
  ],
  imports: [

    FormsModule,
    CommonModule,

    AngularSplitModule,
    MatRadioModule,

    GongdocModule,
    GongdocdiagramsModule,

    GongsvgModule
  ],
  exports: [
    GongsvgdatamodelComponent,
    DataModelPanelComponent
  ]
})
export class GongsvgdatamodelModule { }

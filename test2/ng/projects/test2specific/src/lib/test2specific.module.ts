import { NgModule } from '@angular/core';

import { Test2specificComponent } from './test2specific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { Test2Module } from 'test2'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    Test2specificComponent,
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    Test2Module,
  ],
  exports: [
    Test2specificComponent,
    DataModelPanelComponent
  ]
})
export class Test2specificModule { }

import { NgModule } from '@angular/core';

import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongrouterModule } from 'gongrouter'
import { GongrouterspecificModule } from 'gongrouterspecific'

import { GongtableModule } from 'gongtable'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocspecificModule,

    GongtreeModule,
    GongtreespecificModule,

    GongrouterModule,
    GongrouterspecificModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongtableModule,
  ],
  exports: [
    DataModelPanelComponent
  ]
})
export class GongtabledatamodelModule { }

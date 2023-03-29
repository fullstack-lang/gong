package angular

const NgFileModuleSpecific = `import { NgModule } from '@angular/core';

import { {{TitlePkgName}}specificComponent } from './{{pkgname}}specific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { {{TitlePkgName}}Module } from '{{pkgname}}'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    {{TitlePkgName}}specificComponent,
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    {{TitlePkgName}}Module,
  ],
  exports: [
    {{TitlePkgName}}specificComponent,
    DataModelPanelComponent
  ]
})
export class {{TitlePkgName}}specificModule { }
`

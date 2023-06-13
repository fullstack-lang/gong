package angular

const NgFileModuleSpecific = `import { NgModule } from '@angular/core';

import { {{TitlePkgName}}specificComponent } from './{{pkgname}}specific.component';

import { {{TitlePkgName}}Module } from '{{pkgname}}'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    {{TitlePkgName}}specificComponent
  ],
  imports: [
    GongdocModule,
    GongdocspecificModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    {{TitlePkgName}}Module,
  ],
  exports: [
    {{TitlePkgName}}specificComponent
  ]
})
export class {{TitlePkgName}}specificModule { }
`

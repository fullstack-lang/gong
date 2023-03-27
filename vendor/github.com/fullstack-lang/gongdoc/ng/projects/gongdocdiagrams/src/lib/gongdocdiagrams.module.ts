import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module'

import { GongdocModule } from 'gongdoc'
import { GongModule } from 'gong'

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatSlideToggleModule } from '@angular/material/slide-toggle'
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
import { MatExpansionModule } from '@angular/material/expansion';
import { MatTreeModule } from '@angular/material/tree';
import { MatRadioModule } from '@angular/material/radio';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule } from 'angular-split';

import { GongdocdiagramsComponent } from './gongdocdiagrams.component';
import { ClassDiagramComponent } from './class-diagram/class-diagram.component';
import { PkgeltDocsComponent } from './pkgelt-docs/pkgelt-docs.component';

import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component';
import { TreeComponent } from './tree/tree.component';
import { PanelComponent } from './panel/panel.component'

@NgModule({
  declarations: [
    GongdocdiagramsComponent,
    ClassDiagramComponent,
    PkgeltDocsComponent,
    ClassdiagramDetailComponent,
    TreeComponent,
    PanelComponent
  ],
  imports: [

    BrowserAnimationsModule,
    HttpClientModule,
    AppRoutingModule,


    AngularSplitModule,

    MatSliderModule,
    MatSlideToggleModule,
    MatSelectModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatExpansionModule,
    MatListModule,
    MatRadioModule,
    MatTreeModule,
    DragDropModule,

    FormsModule,
    ReactiveFormsModule,

    GongdocModule,
    GongModule

  ],
  exports: [
    PkgeltDocsComponent,
    ClassDiagramComponent,
    ClassdiagramDetailComponent,
    PanelComponent
  ],
})
export class GongdocdiagramsModule { }

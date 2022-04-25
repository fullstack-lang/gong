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
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule } from 'angular-split';

import { GongdocdiagramsComponent } from './gongdocdiagrams.component';
import { UmlscDiagramComponent } from './umlsc-diagram/umlsc-diagram.component';
import { ClassDiagramComponent } from './class-diagram/class-diagram.component';
import { PkgeltDocsComponent } from './pkgelt-docs/pkgelt-docs.component';
import { UmlscSimpleTableComponent } from './umlsc-simple/umlsc-simple.component';

import { SidebarGongDiagramsComponent } from './sidebar-gong-diagrams/sidebar-gong-diagrams.component'
import { SidebarGongdocDiagramsComponent } from './sidebar-gongdoc-diagrams/sidebar-gongdoc-diagrams.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'

@NgModule({
  declarations: [
    GongdocdiagramsComponent,
    UmlscDiagramComponent,
    ClassDiagramComponent,
    PkgeltDocsComponent,
    UmlscSimpleTableComponent,
    SidebarGongDiagramsComponent,
    SidebarGongdocDiagramsComponent,
    ClassdiagramDetailComponent
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
    MatTreeModule,
    DragDropModule,

    FormsModule,
    ReactiveFormsModule,

    GongdocModule,
    GongModule

  ],
  exports: [
    PkgeltDocsComponent,
    UmlscDiagramComponent,
    ClassDiagramComponent,
    ClassdiagramDetailComponent
  ],
})
export class GongdocdiagramsModule { }

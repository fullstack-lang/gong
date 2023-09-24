import { NgModule } from '@angular/core';

import { MaterialTableComponent } from './material-table/material-table.component';
import { MaterialFormComponent } from './material-form/material-form.component';

import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatTableModule } from '@angular/material/table';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatSortModule } from '@angular/material/sort';
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatSelectModule } from '@angular/material/select'
import { DragDropModule } from '@angular/cdk/drag-drop';
import { MatDialogModule } from '@angular/material/dialog';

@NgModule({
  declarations: [
    MaterialTableComponent,
    MaterialFormComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatInputModule,
    MatSortModule,
    MatPaginatorModule,
    MatSelectModule,
    DragDropModule,
    MatDialogModule,
  ],
  exports: [
    MaterialTableComponent,
    MaterialFormComponent,
  ]
})
export class GongtablespecificModule { }

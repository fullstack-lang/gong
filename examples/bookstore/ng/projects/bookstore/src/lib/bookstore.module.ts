import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

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
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

// insertion point for imports 
import { AreasTableComponent } from './areas-table/areas-table.component'
import { AreaDetailComponent } from './area-detail/area-detail.component'
import { AreaPresentationComponent } from './area-presentation/area-presentation.component'

import { BooksTableComponent } from './books-table/books-table.component'
import { BookDetailComponent } from './book-detail/book-detail.component'
import { BookPresentationComponent } from './book-presentation/book-presentation.component'

import { EditorsTableComponent } from './editors-table/editors-table.component'
import { EditorDetailComponent } from './editor-detail/editor-detail.component'
import { EditorPresentationComponent } from './editor-presentation/editor-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		AreasTableComponent,
		AreaDetailComponent,
		AreaPresentationComponent,

		BooksTableComponent,
		BookDetailComponent,
		BookPresentationComponent,

		EditorsTableComponent,
		EditorDetailComponent,
		EditorPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

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
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		AreasTableComponent,
		AreaDetailComponent,
		AreaPresentationComponent,

		BooksTableComponent,
		BookDetailComponent,
		BookPresentationComponent,

		EditorsTableComponent,
		EditorDetailComponent,
		EditorPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class BookstoreModule { }

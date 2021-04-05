import { BrowserModule } from '@angular/platform-browser';
import { NgModule, ModuleWithProviders, SkipSelf, Optional  } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { HttpClientModule } from '@angular/common/http'; 

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

import { FormsModule } from '@angular/forms';

// for the tree
import { MatTreeModule } from '@angular/material/tree'
import { ReactiveFormsModule } from '@angular/forms';
import { MatNativeDateModule } from '@angular/material/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

// split
import { AngularSplitModule } from 'angular-split';

// devlopment in this module
import { SplitterComponent } from './splitter/splitter.component';
import { SidebarComponent } from './sidebar/sidebar.component'


import { AreasTableComponent } from './areas-table/areas-table.component'
import { AreaAdderComponent } from './area-adder/area-adder.component'
import { AreaDetailComponent } from './area-detail/area-detail.component'
import { AreaPresentationComponent } from './area-presentation/area-presentation.component'

import { BooksTableComponent } from './books-table/books-table.component'
import { BookAdderComponent } from './book-adder/book-adder.component'
import { BookDetailComponent } from './book-detail/book-detail.component'
import { BookPresentationComponent } from './book-presentation/book-presentation.component'

import { EditorsTableComponent } from './editors-table/editors-table.component'
import { EditorAdderComponent } from './editor-adder/editor-adder.component'
import { EditorDetailComponent } from './editor-detail/editor-detail.component'
import { EditorPresentationComponent } from './editor-presentation/editor-presentation.component'


@NgModule({
  declarations: [
	SplitterComponent,
	

	AreasTableComponent,
	AreaAdderComponent,
	AreaDetailComponent,
	AreaPresentationComponent,

	BooksTableComponent,
	BookAdderComponent,
	BookDetailComponent,
	BookPresentationComponent,

	EditorsTableComponent,
	EditorAdderComponent,
	EditorDetailComponent,
	EditorPresentationComponent,


    SidebarComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,

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

    MatTreeModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatListModule,

    FormsModule,

    AngularSplitModule.forRoot()
  ],
  providers: [],
  bootstrap: []
})
export class bookstoreMatModule {

    constructor( @Optional() @SkipSelf() parentModule: bookstoreMatModule,
                 @Optional() http: HttpClient) {
        if (parentModule) {
            throw new Error('bookstoreMatModule is already loaded. Import in your base AppModule only.');
        }
        if (!http) {
            throw new Error('You need to import the HttpClientModule in your AppModule! \n' +
            'See also https://github.com/angular/angular/issues/20575');
        }
    }
}

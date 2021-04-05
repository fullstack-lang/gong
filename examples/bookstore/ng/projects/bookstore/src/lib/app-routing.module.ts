import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

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


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'areas', component: AreasTableComponent, outlet: 'table' },
	{ path: 'area-adder', component: AreaDetailComponent, outlet: 'editor' },
	{ path: 'area-adder/:id/:association', component: AreaDetailComponent, outlet: 'editor' },
	{ path: 'area-detail/:id', component: AreaDetailComponent, outlet: 'editor' },
	{ path: 'area-presentation/:id', component: AreaPresentationComponent, outlet: 'presentation' },
	{ path: 'area-presentation-special/:id', component: AreaPresentationComponent, outlet: 'areapres' },

	{ path: 'books', component: BooksTableComponent, outlet: 'table' },
	{ path: 'book-adder', component: BookDetailComponent, outlet: 'editor' },
	{ path: 'book-adder/:id/:association', component: BookDetailComponent, outlet: 'editor' },
	{ path: 'book-detail/:id', component: BookDetailComponent, outlet: 'editor' },
	{ path: 'book-presentation/:id', component: BookPresentationComponent, outlet: 'presentation' },
	{ path: 'book-presentation-special/:id', component: BookPresentationComponent, outlet: 'bookpres' },

	{ path: 'editors', component: EditorsTableComponent, outlet: 'table' },
	{ path: 'editor-adder', component: EditorDetailComponent, outlet: 'editor' },
	{ path: 'editor-adder/:id/:association', component: EditorDetailComponent, outlet: 'editor' },
	{ path: 'editor-detail/:id', component: EditorDetailComponent, outlet: 'editor' },
	{ path: 'editor-presentation/:id', component: EditorPresentationComponent, outlet: 'presentation' },
	{ path: 'editor-presentation-special/:id', component: EditorPresentationComponent, outlet: 'editorpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

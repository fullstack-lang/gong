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
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramSortingComponent } from './classdiagram-sorting/classdiagram-sorting.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'
import { ClassdiagramPresentationComponent } from './classdiagram-presentation/classdiagram-presentation.component'

import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
import { ClassshapeSortingComponent } from './classshape-sorting/classshape-sorting.component'
import { ClassshapeDetailComponent } from './classshape-detail/classshape-detail.component'
import { ClassshapePresentationComponent } from './classshape-presentation/classshape-presentation.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldSortingComponent } from './field-sorting/field-sorting.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'
import { FieldPresentationComponent } from './field-presentation/field-presentation.component'

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructSortingComponent } from './gongstruct-sorting/gongstruct-sorting.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'
import { GongStructPresentationComponent } from './gongstruct-presentation/gongstruct-presentation.component'

import { GongdocCommandsTableComponent } from './gongdoccommands-table/gongdoccommands-table.component'
import { GongdocCommandSortingComponent } from './gongdoccommand-sorting/gongdoccommand-sorting.component'
import { GongdocCommandDetailComponent } from './gongdoccommand-detail/gongdoccommand-detail.component'
import { GongdocCommandPresentationComponent } from './gongdoccommand-presentation/gongdoccommand-presentation.component'

import { GongdocStatussTableComponent } from './gongdocstatuss-table/gongdocstatuss-table.component'
import { GongdocStatusSortingComponent } from './gongdocstatus-sorting/gongdocstatus-sorting.component'
import { GongdocStatusDetailComponent } from './gongdocstatus-detail/gongdocstatus-detail.component'
import { GongdocStatusPresentationComponent } from './gongdocstatus-presentation/gongdocstatus-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkSortingComponent } from './link-sorting/link-sorting.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { NotesTableComponent } from './notes-table/notes-table.component'
import { NoteSortingComponent } from './note-sorting/note-sorting.component'
import { NoteDetailComponent } from './note-detail/note-detail.component'
import { NotePresentationComponent } from './note-presentation/note-presentation.component'

import { PkgeltsTableComponent } from './pkgelts-table/pkgelts-table.component'
import { PkgeltSortingComponent } from './pkgelt-sorting/pkgelt-sorting.component'
import { PkgeltDetailComponent } from './pkgelt-detail/pkgelt-detail.component'
import { PkgeltPresentationComponent } from './pkgelt-presentation/pkgelt-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionSortingComponent } from './position-sorting/position-sorting.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

import { UmlStatesTableComponent } from './umlstates-table/umlstates-table.component'
import { UmlStateSortingComponent } from './umlstate-sorting/umlstate-sorting.component'
import { UmlStateDetailComponent } from './umlstate-detail/umlstate-detail.component'
import { UmlStatePresentationComponent } from './umlstate-presentation/umlstate-presentation.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscSortingComponent } from './umlsc-sorting/umlsc-sorting.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'
import { UmlscPresentationComponent } from './umlsc-presentation/umlsc-presentation.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeSortingComponent } from './vertice-sorting/vertice-sorting.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'
import { VerticePresentationComponent } from './vertice-presentation/vertice-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		ClassdiagramsTableComponent,
		ClassdiagramSortingComponent,
		ClassdiagramDetailComponent,
		ClassdiagramPresentationComponent,

		ClassshapesTableComponent,
		ClassshapeSortingComponent,
		ClassshapeDetailComponent,
		ClassshapePresentationComponent,

		FieldsTableComponent,
		FieldSortingComponent,
		FieldDetailComponent,
		FieldPresentationComponent,

		GongStructsTableComponent,
		GongStructSortingComponent,
		GongStructDetailComponent,
		GongStructPresentationComponent,

		GongdocCommandsTableComponent,
		GongdocCommandSortingComponent,
		GongdocCommandDetailComponent,
		GongdocCommandPresentationComponent,

		GongdocStatussTableComponent,
		GongdocStatusSortingComponent,
		GongdocStatusDetailComponent,
		GongdocStatusPresentationComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,
		LinkPresentationComponent,

		NotesTableComponent,
		NoteSortingComponent,
		NoteDetailComponent,
		NotePresentationComponent,

		PkgeltsTableComponent,
		PkgeltSortingComponent,
		PkgeltDetailComponent,
		PkgeltPresentationComponent,

		PositionsTableComponent,
		PositionSortingComponent,
		PositionDetailComponent,
		PositionPresentationComponent,

		UmlStatesTableComponent,
		UmlStateSortingComponent,
		UmlStateDetailComponent,
		UmlStatePresentationComponent,

		UmlscsTableComponent,
		UmlscSortingComponent,
		UmlscDetailComponent,
		UmlscPresentationComponent,

		VerticesTableComponent,
		VerticeSortingComponent,
		VerticeDetailComponent,
		VerticePresentationComponent,


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
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		ClassdiagramsTableComponent,
		ClassdiagramSortingComponent,
		ClassdiagramDetailComponent,
		ClassdiagramPresentationComponent,

		ClassshapesTableComponent,
		ClassshapeSortingComponent,
		ClassshapeDetailComponent,
		ClassshapePresentationComponent,

		FieldsTableComponent,
		FieldSortingComponent,
		FieldDetailComponent,
		FieldPresentationComponent,

		GongStructsTableComponent,
		GongStructSortingComponent,
		GongStructDetailComponent,
		GongStructPresentationComponent,

		GongdocCommandsTableComponent,
		GongdocCommandSortingComponent,
		GongdocCommandDetailComponent,
		GongdocCommandPresentationComponent,

		GongdocStatussTableComponent,
		GongdocStatusSortingComponent,
		GongdocStatusDetailComponent,
		GongdocStatusPresentationComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,
		LinkPresentationComponent,

		NotesTableComponent,
		NoteSortingComponent,
		NoteDetailComponent,
		NotePresentationComponent,

		PkgeltsTableComponent,
		PkgeltSortingComponent,
		PkgeltDetailComponent,
		PkgeltPresentationComponent,

		PositionsTableComponent,
		PositionSortingComponent,
		PositionDetailComponent,
		PositionPresentationComponent,

		UmlStatesTableComponent,
		UmlStateSortingComponent,
		UmlStateDetailComponent,
		UmlStatePresentationComponent,

		UmlscsTableComponent,
		UmlscSortingComponent,
		UmlscDetailComponent,
		UmlscPresentationComponent,

		VerticesTableComponent,
		VerticeSortingComponent,
		VerticeDetailComponent,
		VerticePresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongdocModule { }

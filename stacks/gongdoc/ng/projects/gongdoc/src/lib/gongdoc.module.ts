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
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'
import { ClassdiagramPresentationComponent } from './classdiagram-presentation/classdiagram-presentation.component'

import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
import { ClassshapeDetailComponent } from './classshape-detail/classshape-detail.component'
import { ClassshapePresentationComponent } from './classshape-presentation/classshape-presentation.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'
import { FieldPresentationComponent } from './field-presentation/field-presentation.component'

import { GongdocCommandsTableComponent } from './gongdoccommands-table/gongdoccommands-table.component'
import { GongdocCommandDetailComponent } from './gongdoccommand-detail/gongdoccommand-detail.component'
import { GongdocCommandPresentationComponent } from './gongdoccommand-presentation/gongdoccommand-presentation.component'

import { GongdocStatussTableComponent } from './gongdocstatuss-table/gongdocstatuss-table.component'
import { GongdocStatusDetailComponent } from './gongdocstatus-detail/gongdocstatus-detail.component'
import { GongdocStatusPresentationComponent } from './gongdocstatus-presentation/gongdocstatus-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { PkgeltsTableComponent } from './pkgelts-table/pkgelts-table.component'
import { PkgeltDetailComponent } from './pkgelt-detail/pkgelt-detail.component'
import { PkgeltPresentationComponent } from './pkgelt-presentation/pkgelt-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

import { StatesTableComponent } from './states-table/states-table.component'
import { StateDetailComponent } from './state-detail/state-detail.component'
import { StatePresentationComponent } from './state-presentation/state-presentation.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'
import { UmlscPresentationComponent } from './umlsc-presentation/umlsc-presentation.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'
import { VerticePresentationComponent } from './vertice-presentation/vertice-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		ClassdiagramsTableComponent,
		ClassdiagramDetailComponent,
		ClassdiagramPresentationComponent,

		ClassshapesTableComponent,
		ClassshapeDetailComponent,
		ClassshapePresentationComponent,

		FieldsTableComponent,
		FieldDetailComponent,
		FieldPresentationComponent,

		GongdocCommandsTableComponent,
		GongdocCommandDetailComponent,
		GongdocCommandPresentationComponent,

		GongdocStatussTableComponent,
		GongdocStatusDetailComponent,
		GongdocStatusPresentationComponent,

		LinksTableComponent,
		LinkDetailComponent,
		LinkPresentationComponent,

		PkgeltsTableComponent,
		PkgeltDetailComponent,
		PkgeltPresentationComponent,

		PositionsTableComponent,
		PositionDetailComponent,
		PositionPresentationComponent,

		StatesTableComponent,
		StateDetailComponent,
		StatePresentationComponent,

		UmlscsTableComponent,
		UmlscDetailComponent,
		UmlscPresentationComponent,

		VerticesTableComponent,
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
		ClassdiagramsTableComponent,
		ClassdiagramDetailComponent,
		ClassdiagramPresentationComponent,

		ClassshapesTableComponent,
		ClassshapeDetailComponent,
		ClassshapePresentationComponent,

		FieldsTableComponent,
		FieldDetailComponent,
		FieldPresentationComponent,

		GongdocCommandsTableComponent,
		GongdocCommandDetailComponent,
		GongdocCommandPresentationComponent,

		GongdocStatussTableComponent,
		GongdocStatusDetailComponent,
		GongdocStatusPresentationComponent,

		LinksTableComponent,
		LinkDetailComponent,
		LinkPresentationComponent,

		PkgeltsTableComponent,
		PkgeltDetailComponent,
		PkgeltPresentationComponent,

		PositionsTableComponent,
		PositionDetailComponent,
		PositionPresentationComponent,

		StatesTableComponent,
		StateDetailComponent,
		StatePresentationComponent,

		UmlscsTableComponent,
		UmlscDetailComponent,
		UmlscPresentationComponent,

		VerticesTableComponent,
		VerticeDetailComponent,
		VerticePresentationComponent,


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
export class GongdocModule { }

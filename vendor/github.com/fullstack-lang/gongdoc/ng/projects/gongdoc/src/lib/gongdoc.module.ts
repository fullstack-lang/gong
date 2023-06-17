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

import { DiagramPackagesTableComponent } from './diagrampackages-table/diagrampackages-table.component'
import { DiagramPackageSortingComponent } from './diagrampackage-sorting/diagrampackage-sorting.component'
import { DiagramPackageDetailComponent } from './diagrampackage-detail/diagrampackage-detail.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldSortingComponent } from './field-sorting/field-sorting.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'

import { GongEnumShapesTableComponent } from './gongenumshapes-table/gongenumshapes-table.component'
import { GongEnumShapeSortingComponent } from './gongenumshape-sorting/gongenumshape-sorting.component'
import { GongEnumShapeDetailComponent } from './gongenumshape-detail/gongenumshape-detail.component'

import { GongEnumValueEntrysTableComponent } from './gongenumvalueentrys-table/gongenumvalueentrys-table.component'
import { GongEnumValueEntrySortingComponent } from './gongenumvalueentry-sorting/gongenumvalueentry-sorting.component'
import { GongEnumValueEntryDetailComponent } from './gongenumvalueentry-detail/gongenumvalueentry-detail.component'

import { GongStructShapesTableComponent } from './gongstructshapes-table/gongstructshapes-table.component'
import { GongStructShapeSortingComponent } from './gongstructshape-sorting/gongstructshape-sorting.component'
import { GongStructShapeDetailComponent } from './gongstructshape-detail/gongstructshape-detail.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkSortingComponent } from './link-sorting/link-sorting.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'

import { NoteShapesTableComponent } from './noteshapes-table/noteshapes-table.component'
import { NoteShapeSortingComponent } from './noteshape-sorting/noteshape-sorting.component'
import { NoteShapeDetailComponent } from './noteshape-detail/noteshape-detail.component'

import { NoteShapeLinksTableComponent } from './noteshapelinks-table/noteshapelinks-table.component'
import { NoteShapeLinkSortingComponent } from './noteshapelink-sorting/noteshapelink-sorting.component'
import { NoteShapeLinkDetailComponent } from './noteshapelink-detail/noteshapelink-detail.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionSortingComponent } from './position-sorting/position-sorting.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'

import { UmlStatesTableComponent } from './umlstates-table/umlstates-table.component'
import { UmlStateSortingComponent } from './umlstate-sorting/umlstate-sorting.component'
import { UmlStateDetailComponent } from './umlstate-detail/umlstate-detail.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscSortingComponent } from './umlsc-sorting/umlsc-sorting.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeSortingComponent } from './vertice-sorting/vertice-sorting.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		ClassdiagramsTableComponent,
		ClassdiagramSortingComponent,
		ClassdiagramDetailComponent,

		DiagramPackagesTableComponent,
		DiagramPackageSortingComponent,
		DiagramPackageDetailComponent,

		FieldsTableComponent,
		FieldSortingComponent,
		FieldDetailComponent,

		GongEnumShapesTableComponent,
		GongEnumShapeSortingComponent,
		GongEnumShapeDetailComponent,

		GongEnumValueEntrysTableComponent,
		GongEnumValueEntrySortingComponent,
		GongEnumValueEntryDetailComponent,

		GongStructShapesTableComponent,
		GongStructShapeSortingComponent,
		GongStructShapeDetailComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,

		NoteShapesTableComponent,
		NoteShapeSortingComponent,
		NoteShapeDetailComponent,

		NoteShapeLinksTableComponent,
		NoteShapeLinkSortingComponent,
		NoteShapeLinkDetailComponent,

		PositionsTableComponent,
		PositionSortingComponent,
		PositionDetailComponent,

		UmlStatesTableComponent,
		UmlStateSortingComponent,
		UmlStateDetailComponent,

		UmlscsTableComponent,
		UmlscSortingComponent,
		UmlscDetailComponent,

		VerticesTableComponent,
		VerticeSortingComponent,
		VerticeDetailComponent,


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

		DiagramPackagesTableComponent,
		DiagramPackageSortingComponent,
		DiagramPackageDetailComponent,

		FieldsTableComponent,
		FieldSortingComponent,
		FieldDetailComponent,

		GongEnumShapesTableComponent,
		GongEnumShapeSortingComponent,
		GongEnumShapeDetailComponent,

		GongEnumValueEntrysTableComponent,
		GongEnumValueEntrySortingComponent,
		GongEnumValueEntryDetailComponent,

		GongStructShapesTableComponent,
		GongStructShapeSortingComponent,
		GongStructShapeDetailComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,

		NoteShapesTableComponent,
		NoteShapeSortingComponent,
		NoteShapeDetailComponent,

		NoteShapeLinksTableComponent,
		NoteShapeLinkSortingComponent,
		NoteShapeLinkDetailComponent,

		PositionsTableComponent,
		PositionSortingComponent,
		PositionDetailComponent,

		UmlStatesTableComponent,
		UmlStateSortingComponent,
		UmlStateDetailComponent,

		UmlscsTableComponent,
		UmlscSortingComponent,
		UmlscDetailComponent,

		VerticesTableComponent,
		VerticeSortingComponent,
		VerticeDetailComponent,


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

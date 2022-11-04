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

import { DiagramPackagesTableComponent } from './diagrampackages-table/diagrampackages-table.component'
import { DiagramPackageSortingComponent } from './diagrampackage-sorting/diagrampackage-sorting.component'
import { DiagramPackageDetailComponent } from './diagrampackage-detail/diagrampackage-detail.component'
import { DiagramPackagePresentationComponent } from './diagrampackage-presentation/diagrampackage-presentation.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldSortingComponent } from './field-sorting/field-sorting.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'
import { FieldPresentationComponent } from './field-presentation/field-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkSortingComponent } from './link-sorting/link-sorting.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeSortingComponent } from './node-sorting/node-sorting.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'
import { NodePresentationComponent } from './node-presentation/node-presentation.component'

import { NotesTableComponent } from './notes-table/notes-table.component'
import { NoteSortingComponent } from './note-sorting/note-sorting.component'
import { NoteDetailComponent } from './note-detail/note-detail.component'
import { NotePresentationComponent } from './note-presentation/note-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionSortingComponent } from './position-sorting/position-sorting.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

import { ReferencesTableComponent } from './references-table/references-table.component'
import { ReferenceSortingComponent } from './reference-sorting/reference-sorting.component'
import { ReferenceDetailComponent } from './reference-detail/reference-detail.component'
import { ReferencePresentationComponent } from './reference-presentation/reference-presentation.component'

import { TreesTableComponent } from './trees-table/trees-table.component'
import { TreeSortingComponent } from './tree-sorting/tree-sorting.component'
import { TreeDetailComponent } from './tree-detail/tree-detail.component'
import { TreePresentationComponent } from './tree-presentation/tree-presentation.component'

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

		DiagramPackagesTableComponent,
		DiagramPackageSortingComponent,
		DiagramPackageDetailComponent,
		DiagramPackagePresentationComponent,

		FieldsTableComponent,
		FieldSortingComponent,
		FieldDetailComponent,
		FieldPresentationComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,
		LinkPresentationComponent,

		NodesTableComponent,
		NodeSortingComponent,
		NodeDetailComponent,
		NodePresentationComponent,

		NotesTableComponent,
		NoteSortingComponent,
		NoteDetailComponent,
		NotePresentationComponent,

		PositionsTableComponent,
		PositionSortingComponent,
		PositionDetailComponent,
		PositionPresentationComponent,

		ReferencesTableComponent,
		ReferenceSortingComponent,
		ReferenceDetailComponent,
		ReferencePresentationComponent,

		TreesTableComponent,
		TreeSortingComponent,
		TreeDetailComponent,
		TreePresentationComponent,

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

		DiagramPackagesTableComponent,
		DiagramPackageSortingComponent,
		DiagramPackageDetailComponent,
		DiagramPackagePresentationComponent,

		FieldsTableComponent,
		FieldSortingComponent,
		FieldDetailComponent,
		FieldPresentationComponent,

		LinksTableComponent,
		LinkSortingComponent,
		LinkDetailComponent,
		LinkPresentationComponent,

		NodesTableComponent,
		NodeSortingComponent,
		NodeDetailComponent,
		NodePresentationComponent,

		NotesTableComponent,
		NoteSortingComponent,
		NoteDetailComponent,
		NotePresentationComponent,

		PositionsTableComponent,
		PositionSortingComponent,
		PositionDetailComponent,
		PositionPresentationComponent,

		ReferencesTableComponent,
		ReferenceSortingComponent,
		ReferenceDetailComponent,
		ReferencePresentationComponent,

		TreesTableComponent,
		TreeSortingComponent,
		TreeDetailComponent,
		TreePresentationComponent,

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

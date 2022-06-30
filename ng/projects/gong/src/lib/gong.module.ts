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
import { GongBasicFieldsTableComponent } from './gongbasicfields-table/gongbasicfields-table.component'
import { GongBasicFieldSortingComponent } from './gongbasicfield-sorting/gongbasicfield-sorting.component'
import { GongBasicFieldDetailComponent } from './gongbasicfield-detail/gongbasicfield-detail.component'
import { GongBasicFieldPresentationComponent } from './gongbasicfield-presentation/gongbasicfield-presentation.component'

import { GongEnumsTableComponent } from './gongenums-table/gongenums-table.component'
import { GongEnumSortingComponent } from './gongenum-sorting/gongenum-sorting.component'
import { GongEnumDetailComponent } from './gongenum-detail/gongenum-detail.component'
import { GongEnumPresentationComponent } from './gongenum-presentation/gongenum-presentation.component'

import { GongEnumValuesTableComponent } from './gongenumvalues-table/gongenumvalues-table.component'
import { GongEnumValueSortingComponent } from './gongenumvalue-sorting/gongenumvalue-sorting.component'
import { GongEnumValueDetailComponent } from './gongenumvalue-detail/gongenumvalue-detail.component'
import { GongEnumValuePresentationComponent } from './gongenumvalue-presentation/gongenumvalue-presentation.component'

import { GongNotesTableComponent } from './gongnotes-table/gongnotes-table.component'
import { GongNoteSortingComponent } from './gongnote-sorting/gongnote-sorting.component'
import { GongNoteDetailComponent } from './gongnote-detail/gongnote-detail.component'
import { GongNotePresentationComponent } from './gongnote-presentation/gongnote-presentation.component'

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructSortingComponent } from './gongstruct-sorting/gongstruct-sorting.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'
import { GongStructPresentationComponent } from './gongstruct-presentation/gongstruct-presentation.component'

import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
import { GongTimeFieldSortingComponent } from './gongtimefield-sorting/gongtimefield-sorting.component'
import { GongTimeFieldDetailComponent } from './gongtimefield-detail/gongtimefield-detail.component'
import { GongTimeFieldPresentationComponent } from './gongtimefield-presentation/gongtimefield-presentation.component'

import { ModelPkgsTableComponent } from './modelpkgs-table/modelpkgs-table.component'
import { ModelPkgSortingComponent } from './modelpkg-sorting/modelpkg-sorting.component'
import { ModelPkgDetailComponent } from './modelpkg-detail/modelpkg-detail.component'
import { ModelPkgPresentationComponent } from './modelpkg-presentation/modelpkg-presentation.component'

import { PointerToGongStructFieldsTableComponent } from './pointertogongstructfields-table/pointertogongstructfields-table.component'
import { PointerToGongStructFieldSortingComponent } from './pointertogongstructfield-sorting/pointertogongstructfield-sorting.component'
import { PointerToGongStructFieldDetailComponent } from './pointertogongstructfield-detail/pointertogongstructfield-detail.component'
import { PointerToGongStructFieldPresentationComponent } from './pointertogongstructfield-presentation/pointertogongstructfield-presentation.component'

import { SliceOfPointerToGongStructFieldsTableComponent } from './sliceofpointertogongstructfields-table/sliceofpointertogongstructfields-table.component'
import { SliceOfPointerToGongStructFieldSortingComponent } from './sliceofpointertogongstructfield-sorting/sliceofpointertogongstructfield-sorting.component'
import { SliceOfPointerToGongStructFieldDetailComponent } from './sliceofpointertogongstructfield-detail/sliceofpointertogongstructfield-detail.component'
import { SliceOfPointerToGongStructFieldPresentationComponent } from './sliceofpointertogongstructfield-presentation/sliceofpointertogongstructfield-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		GongBasicFieldsTableComponent,
		GongBasicFieldSortingComponent,
		GongBasicFieldDetailComponent,
		GongBasicFieldPresentationComponent,

		GongEnumsTableComponent,
		GongEnumSortingComponent,
		GongEnumDetailComponent,
		GongEnumPresentationComponent,

		GongEnumValuesTableComponent,
		GongEnumValueSortingComponent,
		GongEnumValueDetailComponent,
		GongEnumValuePresentationComponent,

		GongNotesTableComponent,
		GongNoteSortingComponent,
		GongNoteDetailComponent,
		GongNotePresentationComponent,

		GongStructsTableComponent,
		GongStructSortingComponent,
		GongStructDetailComponent,
		GongStructPresentationComponent,

		GongTimeFieldsTableComponent,
		GongTimeFieldSortingComponent,
		GongTimeFieldDetailComponent,
		GongTimeFieldPresentationComponent,

		ModelPkgsTableComponent,
		ModelPkgSortingComponent,
		ModelPkgDetailComponent,
		ModelPkgPresentationComponent,

		PointerToGongStructFieldsTableComponent,
		PointerToGongStructFieldSortingComponent,
		PointerToGongStructFieldDetailComponent,
		PointerToGongStructFieldPresentationComponent,

		SliceOfPointerToGongStructFieldsTableComponent,
		SliceOfPointerToGongStructFieldSortingComponent,
		SliceOfPointerToGongStructFieldDetailComponent,
		SliceOfPointerToGongStructFieldPresentationComponent,


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
		GongBasicFieldsTableComponent,
		GongBasicFieldSortingComponent,
		GongBasicFieldDetailComponent,
		GongBasicFieldPresentationComponent,

		GongEnumsTableComponent,
		GongEnumSortingComponent,
		GongEnumDetailComponent,
		GongEnumPresentationComponent,

		GongEnumValuesTableComponent,
		GongEnumValueSortingComponent,
		GongEnumValueDetailComponent,
		GongEnumValuePresentationComponent,

		GongNotesTableComponent,
		GongNoteSortingComponent,
		GongNoteDetailComponent,
		GongNotePresentationComponent,

		GongStructsTableComponent,
		GongStructSortingComponent,
		GongStructDetailComponent,
		GongStructPresentationComponent,

		GongTimeFieldsTableComponent,
		GongTimeFieldSortingComponent,
		GongTimeFieldDetailComponent,
		GongTimeFieldPresentationComponent,

		ModelPkgsTableComponent,
		ModelPkgSortingComponent,
		ModelPkgDetailComponent,
		ModelPkgPresentationComponent,

		PointerToGongStructFieldsTableComponent,
		PointerToGongStructFieldSortingComponent,
		PointerToGongStructFieldDetailComponent,
		PointerToGongStructFieldPresentationComponent,

		SliceOfPointerToGongStructFieldsTableComponent,
		SliceOfPointerToGongStructFieldSortingComponent,
		SliceOfPointerToGongStructFieldDetailComponent,
		SliceOfPointerToGongStructFieldPresentationComponent,


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
export class GongModule { }

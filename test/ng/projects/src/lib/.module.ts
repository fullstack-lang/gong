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
import { AstructsTableComponent } from './astructs-table/astructs-table.component'
import { AstructSortingComponent } from './astruct-sorting/astruct-sorting.component'
import { AstructDetailComponent } from './astruct-detail/astruct-detail.component'
import { AstructPresentationComponent } from './astruct-presentation/astruct-presentation.component'

import { AstructBstruct2UsesTableComponent } from './astructbstruct2uses-table/astructbstruct2uses-table.component'
import { AstructBstruct2UseSortingComponent } from './astructbstruct2use-sorting/astructbstruct2use-sorting.component'
import { AstructBstruct2UseDetailComponent } from './astructbstruct2use-detail/astructbstruct2use-detail.component'
import { AstructBstruct2UsePresentationComponent } from './astructbstruct2use-presentation/astructbstruct2use-presentation.component'

import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
import { AstructBstructUseSortingComponent } from './astructbstructuse-sorting/astructbstructuse-sorting.component'
import { AstructBstructUseDetailComponent } from './astructbstructuse-detail/astructbstructuse-detail.component'
import { AstructBstructUsePresentationComponent } from './astructbstructuse-presentation/astructbstructuse-presentation.component'

import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
import { BstructSortingComponent } from './bstruct-sorting/bstruct-sorting.component'
import { BstructDetailComponent } from './bstruct-detail/bstruct-detail.component'
import { BstructPresentationComponent } from './bstruct-presentation/bstruct-presentation.component'

import { DstructsTableComponent } from './dstructs-table/dstructs-table.component'
import { DstructSortingComponent } from './dstruct-sorting/dstruct-sorting.component'
import { DstructDetailComponent } from './dstruct-detail/dstruct-detail.component'
import { DstructPresentationComponent } from './dstruct-presentation/dstruct-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		AstructsTableComponent,
		AstructSortingComponent,
		AstructDetailComponent,
		AstructPresentationComponent,

		AstructBstruct2UsesTableComponent,
		AstructBstruct2UseSortingComponent,
		AstructBstruct2UseDetailComponent,
		AstructBstruct2UsePresentationComponent,

		AstructBstructUsesTableComponent,
		AstructBstructUseSortingComponent,
		AstructBstructUseDetailComponent,
		AstructBstructUsePresentationComponent,

		BstructsTableComponent,
		BstructSortingComponent,
		BstructDetailComponent,
		BstructPresentationComponent,

		DstructsTableComponent,
		DstructSortingComponent,
		DstructDetailComponent,
		DstructPresentationComponent,


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
		AstructsTableComponent,
		AstructSortingComponent,
		AstructDetailComponent,
		AstructPresentationComponent,

		AstructBstruct2UsesTableComponent,
		AstructBstruct2UseSortingComponent,
		AstructBstruct2UseDetailComponent,
		AstructBstruct2UsePresentationComponent,

		AstructBstructUsesTableComponent,
		AstructBstructUseSortingComponent,
		AstructBstructUseDetailComponent,
		AstructBstructUsePresentationComponent,

		BstructsTableComponent,
		BstructSortingComponent,
		BstructDetailComponent,
		BstructPresentationComponent,

		DstructsTableComponent,
		DstructSortingComponent,
		DstructDetailComponent,
		DstructPresentationComponent,


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
export class Module { }

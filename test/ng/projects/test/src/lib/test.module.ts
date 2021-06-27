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

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

// insertion point for imports 
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { AclassSortingComponent } from './aclass-sorting/aclass-sorting.component'
import { AclassDetailComponent } from './aclass-detail/aclass-detail.component'
import { AclassPresentationComponent } from './aclass-presentation/aclass-presentation.component'

import { AclassBclassUsesTableComponent } from './aclassbclassuses-table/aclassbclassuses-table.component'
import { AclassBclassUseSortingComponent } from './aclassbclassuse-sorting/aclassbclassuse-sorting.component'
import { AclassBclassUseDetailComponent } from './aclassbclassuse-detail/aclassbclassuse-detail.component'
import { AclassBclassUsePresentationComponent } from './aclassbclassuse-presentation/aclassbclassuse-presentation.component'

import { BclasssTableComponent } from './bclasss-table/bclasss-table.component'
import { BclassSortingComponent } from './bclass-sorting/bclass-sorting.component'
import { BclassDetailComponent } from './bclass-detail/bclass-detail.component'
import { BclassPresentationComponent } from './bclass-presentation/bclass-presentation.component'

import { DclasssTableComponent } from './dclasss-table/dclasss-table.component'
import { DclassSortingComponent } from './dclass-sorting/dclass-sorting.component'
import { DclassDetailComponent } from './dclass-detail/dclass-detail.component'
import { DclassPresentationComponent } from './dclass-presentation/dclass-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		AclasssTableComponent,
		AclassSortingComponent,
		AclassDetailComponent,
		AclassPresentationComponent,

		AclassBclassUsesTableComponent,
		AclassBclassUseSortingComponent,
		AclassBclassUseDetailComponent,
		AclassBclassUsePresentationComponent,

		BclasssTableComponent,
		BclassSortingComponent,
		BclassDetailComponent,
		BclassPresentationComponent,

		DclasssTableComponent,
		DclassSortingComponent,
		DclassDetailComponent,
		DclassPresentationComponent,


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

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		AclasssTableComponent,
		AclassSortingComponent,
		AclassDetailComponent,
		AclassPresentationComponent,

		AclassBclassUsesTableComponent,
		AclassBclassUseSortingComponent,
		AclassBclassUseDetailComponent,
		AclassBclassUsePresentationComponent,

		BclasssTableComponent,
		BclassSortingComponent,
		BclassDetailComponent,
		BclassPresentationComponent,

		DclasssTableComponent,
		DclassSortingComponent,
		DclassDetailComponent,
		DclassPresentationComponent,


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
export class TestModule { }
